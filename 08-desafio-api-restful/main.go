package main

import (
	"8-desafio-restful/api"
	"8-desafio-restful/database"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main () {
	if err := run(); err != nil {
		slog.Error("não foi possível executar a aplicação")
		os.Exit(1)
	}
}

func run() error {
	id, user := database.Insert(database.User{
		FirstName: "Ariel",
		LastName: "Evangelista",
		Biography: "Tech Lead",
	})
	slog.Info(id.String(), "user", user)

	handler := api.CreateHandler()

	server := http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
		Addr: ":8080",
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
