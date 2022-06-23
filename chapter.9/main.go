package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userID := "0001"
	var userName string
	var createdAt time.Time
	err = db.QueryRowContext(context.TODO(), `SELECT user_name, created_at FROM users WHERE user_id = $1;`, userID).Scan(&userName, &createdAt)
	if err != nil {
		log.Fatalf("query row(user_id=%s): %v", userID, err)
	}

	u := User{
		UserID:    userID,
		UserName:  userName,
		CreatedAt: createdAt,
	}

	fmt.Println(u)
}
