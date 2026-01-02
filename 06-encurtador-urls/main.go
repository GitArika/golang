package main

import (
	"6-encurtador-urls/api"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		slog.Error("failed to execute code", "error", err)
		os.Exit(1)
	}
	slog.Info("all systems offline")
}

func run() error {
	db := make(map[string]string, 10)
	handler := api.NewHandler(db)

	s := http.Server{
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 1 * time.Minute,
		Addr: ":8080",
		Handler: handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

