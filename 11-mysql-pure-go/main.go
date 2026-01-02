package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:dockerpwd@/tests")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	createTableStmt := `
	CREATE TABLE IF NOT EXISTS foo (
		id INT AUTO_INCREMENT PRIMARY KEY,
		bar VARCHAR(255)
	);
	`
	if _, err := db.Exec(createTableStmt); err != nil {
		panic(err)
	}

	insertStmt := `
	INSERT INTO foo (bar) values (?);
	`
	if _, err := db.Exec(insertStmt, "foobar"); err != nil {
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
	if err := db.QueryRow(queryStmt).Scan(&res.id, &res.bar); err != nil {
		panic(err)
	}

	fmt.Printf("%#+v\n", res)
}

