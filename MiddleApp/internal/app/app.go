package app

import (
	"MiddleApp/config"
	"MiddleApp/internal/handlers"
	db "MiddleApp/internal/repositories/db/postgres"
	"MiddleApp/internal/repositories/interfaces"
	"MiddleApp/pkg/postgres"
	"fmt"
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
