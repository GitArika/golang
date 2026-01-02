package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./bar.db")
	if err != nil {
		slog.Error("nao foi possivel conectar ao db", "error", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		slog.Error("nao foi possivel fazer um ping ao db", "error", err)
		os.Exit(1)
	}
}
