package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := "postgres://docker:dockerpwd@localhost:5432/golang"
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	createTableStmt := `
	CREATE TABLE IF NOT EXISTS foo (
		id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
		bar VARCHAR(255)
	);
	`
	if _, err := db.Exec(context.Background(),createTableStmt); err != nil {
		panic(err)
	}

	insertStmt := `
	INSERT INTO foo (bar) values ($1);
	`
	if _, err := db.Exec(context.Background(), insertStmt, "foobar"); err != nil {
		panic(err)
	}

	queryStmt := `
	SELECT * FROM foo LIMIT 1;
	`
	type foobar struct {
		id int
		bar string
	}
	var res foobar
	if err := db.QueryRow(context.Background(), queryStmt).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v\n", res)
}
