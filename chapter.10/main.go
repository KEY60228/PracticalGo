package main

import (
	"log"
	"net/http"
)

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

func Healthz(w http.ResponseWriter, r *http.Request) {
	panic("panic!")
}

func main() {
	http.Handle("/healthz", Recovery(http.HandlerFunc(Healthz)))
	http.ListenAndServe("localhost:8888", nil)
}
