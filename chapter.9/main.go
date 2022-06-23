package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "password"
	PostgresDB       = "go"
)

type PgTable struct {
	SchemaName string
	TableName  string
}

type User struct {
	UserID    string
	UserName  string
	CreatedAt time.Time
}

var _ pgx.Logger = (*logger)(nil)

type logger struct{}

func (l *logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	if msg == "Query" {
		log.Printf("SQL:\n%v\nARGS:%v\n", data["sql"], data["args"])
	}
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	ctx := context.Background()

	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("parse config: %v\n", err)
	}
	config.Logger = &logger{}

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	rows, err := conn.Query(ctx, sql, args)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pgtables []PgTable
	for rows.Next() {
		var s string
		var t string
		if err := rows.Scan(&s, &t); err != nil {
			log.Fatalf("Scan: %v\n", err)
		}
		pgtables = append(pgtables, PgTable{SchemaName: s, TableName: t})
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, table := range pgtables {
		fmt.Println(table)
	}
}
