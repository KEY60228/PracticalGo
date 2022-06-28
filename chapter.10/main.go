package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/file", handleFile)
	http.ListenAndServe("127.0.0.1:8888", nil)
}

func handleFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 * 1024 * 1024)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println(h.Filename)

	o, err := os.Create(h.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer o.Close()

	_, err = io.Copy(o, f)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	value := r.PostFormValue("data")
	log.Printf("value = %s\n", value)
}
