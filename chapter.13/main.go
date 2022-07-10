package main

import (
	"io"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func initServer() http.Handler {
	r := chi.NewRouter()
	r.Get("/fortune", fortuneHandler)
	return r
}

func fortuneHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"fortune": "大吉"}`)
}

func main() {
	log.Println("open at localhost:8888")
	log.Println(http.ListenAndServe("localhost:8888", initServer()))
}
