package app

import (
	"MiddleApp/api/api/ServiceApiPb"
	"MiddleApp/config"
	"MiddleApp/grpcHandlers"
	"MiddleApp/internal/handlers"
	db "MiddleApp/internal/repositories/db/postgres"
	"MiddleApp/internal/repositories/interfaces"
	"MiddleApp/pkg/postgres"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func RunHttp(config *config.Config, repo interfaces.Repository) {
	h := handlers.New(repo)

	fmt.Printf("Service is listening on port: %s", config.Port)

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
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Printf("Listener error: %s", err.Error())
	}

	h := grpcHandlers.New(repo)
	ServiceApiPb.RegisterUsersServiceServer(grpcServer, h)
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("serve error%v\n", err)
	}
}
