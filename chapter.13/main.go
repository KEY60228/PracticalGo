package main

import (
	"context"
	"log"
	"net/http"
)

func main() {
	client := http.Client{}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://example.com/movies/1985", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
}
