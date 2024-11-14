package app

import (
	"MiddleApp/config"
	"MiddleApp/internal/domain"
	"MiddleApp/internal/handlers"
	"MiddleApp/internal/repositories/db/memory"
	"fmt"
	"net/http"
)

func RunHttp(config *config.Config) {
	data := make(map[uint32]*domain.User)
	repo := memory.New(data)
	h := handlers.New(repo)

	http.Handle("/", h.Router)

	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Service is listening...")
	}
}
