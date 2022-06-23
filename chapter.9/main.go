package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gchaincl/sqlhooks"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
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

var _ sqlhooks.Hooks = (*hook)(nil)

type hook struct{}

func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	log.Printf("SQL:\n%v\nArgs:\n%v\n", query, args)
	return ctx, nil
}

func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	return ctx, nil
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	ctx := context.Background()

	sql.Register("postgres-proxy", sqlhooks.Wrap(stdlib.GetDefaultDriver(), &hook{}))

	db, err := sqlx.Connect("postgres-proxy", dsn)
	if err != nil {
		log.Fatalf("connect: %v\n", err)
	}

	sql := `SELECT schemaname, tablename FROM pg_tables WHERE schemaname = $1;`
	args := `information_schema`

	var pgtables []PgTable
	if err := db.SelectContext(ctx, &pgtables, sql, args); err != nil {
		log.Fatalf("select: %v\n", err)
	}

	for _, table := range pgtables {
		fmt.Println(table)
	}
}
