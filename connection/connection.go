package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnection() {
	var err error
	databaseURL := "postgres://postgres:125438@localhost:5432/B48S1"

	Conn, err = pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// fmt.Println("Database is connected")
}