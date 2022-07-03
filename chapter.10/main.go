package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	PostgresHost     = "localhost"
	PostgresPort     = 5432
	PostgresUser     = "postgres"
	PostgresPassword = "password"
	PostgresDB       = "go"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(writer http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{writer, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewMiddlewareTx(db *sql.DB) func(http.Handler) http.Handler {
	return func(wrappedHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, _ := db.Begin()
			lrw := NewLoggingResponseWriter(w)
			r = r.WithContext(context.WithValue(r.Context(), "tx", tx))

			wrappedHandler.ServeHTTP(lrw, r)

			statusCode := lrw.statusCode
			if 200 <= statusCode && statusCode < 400 {
				log.Println("transaction committed")
				tx.Commit()
			} else {
				log.Println("transaction rolling back due to status code:", statusCode)
				tx.Rollback()
			}
		})
	}
}

func extractTx(r *http.Request) *sql.Tx {
	tx, ok := r.Context().Value("tx").(*sql.Tx)
	if !ok {
		panic("transaction middleware is not supported")
	}
	return tx
}

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", PostgresHost, PostgresPort, PostgresUser, PostgresDB, PostgresPassword)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx := NewMiddlewareTx(db)

	http.Handle("/comments", tx(Recovery(http.HandlerFunc(Comments))))
	http.ListenAndServe("localhost:8888", nil)
}

func Comments(w http.ResponseWriter, r *http.Request) {
	tx := extractTx(r)
	// DB処理
	fmt.Print(tx)
}
