package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		slog.Error("não foi possível conectar ao db", "error", err)
		os.Exit(1)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS foo (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT
	);
	`
	res, err := db.Exec(createTableSQL)
	if err != nil {
		slog.Error("ocorreu um erro ao criar tabela", "error", err)
		os.Exit(1)
	}

	fmt.Println(res.RowsAffected())

	insertSQL := `
	INSERT INTO foo (id, name) VALUES (1, "Ariel")
	`

	res, err = db.Exec(insertSQL);
	if err != nil {
		slog.Error("ocorreu um erro ao inserir dados", "error", err)
		os.Exit(1)
	}
	fmt.Println(res.RowsAffected())

	type User struct {
		ID int64
		Name string
	}

	querySQL := `
	SELECT * FROM foo WHERE id = ?;
	`

	var user User
	if err = db.QueryRow(querySQL, 1).Scan(&user.ID, &user.Name); err != nil {
		slog.Error("ocorreu um erro ao buscar dados", "error", err)
		os.Exit(1)
	}
	fmt.Println(user)

	deleteSQL := `
	DELETE FROM foo WHERE id = ?;
	`

	if res, err = db.Exec(deleteSQL, 1); err != nil {
		slog.Error("ocorreu um erro ao excluir dados", "error", err)
	}
	fmt.Println(res.RowsAffected())
}

