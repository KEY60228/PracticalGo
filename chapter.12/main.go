package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	tp, err := initJaegerProvider("http://localhost:14268/api/traces")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down jaeger tracer provider: %v", err)
		}
	}()

	r := mux.NewRouter()
	r.Use(otelmux.Middleware("my-server"))

	r.HandleFunc("/users/{id:[0-9]+}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		name := getUser(r.Context(), id)

		time.Sleep(time.Microsecond)

		reply := fmt.Sprintf("user %s (id %s)\n", name, id)
		_, _ = w.Write([]byte(reply))
	}))

	http.Handle("/", r)
	log.Println("start listening at :8888")
	_ = http.ListenAndServe("localhost:8888", nil)
}

func getUser(ctx context.Context, id string) string {
	time.Sleep(time.Microsecond)

	_, span := tracer.Start(ctx, "getUser", trace.WithAttributes(attribute.String("id", id)))
	defer span.End()

	time.Sleep(time.Microsecond)

	if id == "123" {
		return "otelmux tester"
	}
	return "unknown"
}
