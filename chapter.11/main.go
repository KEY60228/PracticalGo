package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(errors.New(res.Status))
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
