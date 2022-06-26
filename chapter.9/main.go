package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "password"
	PostgresDB       = "go"

	TestHost     = "localhost"
	TestPort     = 5432
	TestUser     = "postgres"
	TestPassword = "password"
	TestDB       = "testdb"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := NewUserService(db)
	user, err := s.FetchUser(context.Background(), "0001")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user)
}
