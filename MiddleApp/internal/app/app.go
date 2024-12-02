package app

import (
	"MiddleApp/api/api/ServiceApiPb"
	"MiddleApp/config"
	"MiddleApp/internal/grpcHandlers"
	"MiddleApp/internal/handlers"
	"MiddleApp/internal/kafka"
	"MiddleApp/internal/kafkaHandlers"
	db "MiddleApp/internal/repositories/db/postgres"
	"MiddleApp/internal/repositories/interfaces"
	"MiddleApp/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

func RunHttp(config *config.Config, repo interfaces.Repository) {
	h := handlers.New(repo)

	fmt.Println("Service is listening on port: 9000")

	http.Handle("/", h.Router)

	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		panic(err)
	}
}

func RunRepo(config *config.Config) *db.UsersRepository {
	pool := postgres.New(config.ConnectingString)

	repo := db.New(pool.Pool)

	return repo
}

func RunGrpc(config *config.Config, repo interfaces.Repository) {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", config.GrpcPort)
	if err != nil {
		fmt.Printf("Listener error: %s", err.Error())
	}

	h := grpcHandlers.New(repo)
	ServiceApiPb.RegisterUsersServiceServer(grpcServer, h)
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("serve error%v\n", err)
	}

	fmt.Println("GRPC service is listening on port: 9091")
}

func RunKafka(repo interfaces.Repository) {
	producer := kafka.NewProducer()
	consumerGroup := kafka.NewConsumerGroup()

	createHandler := kafkaHandlers.NewCreateHandler(producer, consumerGroup, repo)
	go kafkaHandlers.CreateClaim(createHandler)

	updateHandler := kafkaHandlers.NewUpdateHandler(producer, consumerGroup, repo)
	go kafkaHandlers.UpdateClaim(updateHandler)

	deleteHandler := kafkaHandlers.NewDeleteHandler(producer, consumerGroup, repo)
	go kafkaHandlers.DeleteClaim(deleteHandler)
}
