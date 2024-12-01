package app

import (
	"Application/ValidationService/api/api/ServiceApiPb"
	"Application/ValidationService/internal/handlers"
	"Application/ValidationService/kafka"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

const port = ":3030"

func RunHttp() {
	connect, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := ServiceApiPb.NewUsersServiceClient(connect)

	producer := kafka.NewProducer()
	consumer := kafka.NewConsumer()
	h := handlers.New(producer, consumer, client)
	http.Handle("/", h.Router)

	fmt.Println("Service 3030 is listening...")

	err = http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
