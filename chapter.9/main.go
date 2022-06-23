package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "password"
	PostgresDB       = "go"
)

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

func main() {
	// dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s", PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB))
	if err != nil {
		log.Fatal(err)
	}

	txn, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal(err)
	}

	users := [][]interface{}{
		{"0003", "Duke", time.Now()},
		{"0004", "KEY", time.Now()},
		{"0005", "vessy", time.Now()},
	}

	i, err := txn.CopyFrom(ctx, pgx.Identifier{"users"}, []string{"user_id", "user_name", "created_at"}, pgx.CopyFromRows(users))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(i)
}
