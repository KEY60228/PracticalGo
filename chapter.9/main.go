package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
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
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	users := []User{
		{"0003", "Duke", time.Now()},
		{"0004", "KEY", time.Now()},
		{"0005", "vessy", time.Now()},
	}

	valueStrings := make([]string, 0, len(users))
	valueArgs := make([]interface{}, 0, len(users)*3)

	number := 1
	for _, u := range users {
		valueStrings = append(valueStrings, fmt.Sprintf(" ($%d, $%d, $%d)", number, number+1, number+2))
		valueArgs = append(valueArgs, u.UserID)
		valueArgs = append(valueArgs, u.UserName)
		valueArgs = append(valueArgs, u.CreatedAt)
		number += 3
	}

	query := fmt.Sprintf("INSERT INTO users (user_id, user_name, created_at) VALUES %s;", strings.Join(valueStrings, ","))
	if _, err := db.ExecContext(ctx, query, valueArgs...); err != nil {
		log.Fatal(err)
	}
}
