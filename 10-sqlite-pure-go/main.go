package main

import (
	"fmt"
	"log/slog"
	"os"

	"zombiezen.com/go/sqlite"
	"zombiezen.com/go/sqlite/sqlitex"
)

func main() {
	conn, err := sqlite.OpenConn(":memory:", sqlite.OpenReadWrite)
	if err != nil {
		slog.Error("não foi possível estabelecer a conexão com db", "error", err)
		os.Exit(1)
	}
	defer conn.Close()

	err = sqlitex.ExecuteTransient(conn, "SELECT 'hello, world';", &sqlitex.ExecOptions{
		ResultFunc: func(stmt *sqlite.Stmt) error {
			fmt.Println(stmt.ColumnText(0))
			return nil
		},
	})
	if err != nil {
		slog.Error("não foi possível executar consulta sql", "error", err)
		os.Exit(1)
	}
}

