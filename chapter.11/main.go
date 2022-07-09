package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 2

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "http://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := retryClient.StandardClient().Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.Status)
}
