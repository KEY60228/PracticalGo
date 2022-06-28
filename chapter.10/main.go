package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/params", handleParams)
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func handleParams(w http.ResponseWriter, r *http.Request) {
	word := r.FormValue("searchword")
	log.Printf("search word = %s\n", word)

	words, ok := r.Form["searchword"]
	log.Printf("search words = %v has values %v\n", words, ok)

	log.Print("== all queries ==")
	for key, values := range r.Form {
		log.Printf("%s: %v\n", key, values)
	}
}
