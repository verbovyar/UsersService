package main

import (
	"MiddleApp/config"
	"MiddleApp/internal/app"
)

func main() {
	conf, err := config.LoadConfig("./config")
	if err != nil {
		println(err.Error())
	}

	repo := app.RunRepo(&conf)
	app.RunHttp(&conf, repo)
}
