package main

import (
	"8-desafio-restful/api"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func main () {
	if err := run(); err != nil {
		slog.Error("não foi possível executar a aplicação")
		os.Exit(1)
	}
}

func run() error {
	db := make(map[uuid.UUID]api.User, 10)
	
	// mock user for test
	id := uuid.New()
	user := api.User{
		FirstName: "Ariel",
		LastName: "Evangelista",
		Biography: "Tech Lead",
	}
	db[id] = user
	slog.Info(id.String(), "user", user)

	handler := api.CreateHandler(db)

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
