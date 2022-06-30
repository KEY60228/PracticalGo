package main

import (
	"log"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start %s\n", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("finish %s\n", r.URL)
	})
}

func main() {
	http.Handle("/healthz", MiddlewareLogging(http.HandlerFunc(Healthz)))
	http.ListenAndServe("localhost:8888", nil)
}
