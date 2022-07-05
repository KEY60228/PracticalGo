package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type basicAuthRoundTripper struct {
	transport http.RoundTripper
	username  string
	password  string
}

func (rt *basicAuthRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(rt.username, rt.password)
	return rt.transport.RoundTrip(req)
}

func main() {
	client := &http.Client{
		Transport: &basicAuthRoundTripper{
			transport: http.DefaultTransport,
		},
	}

	req, err := http.NewRequestWithContext(context.Background(), "GET", "http://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
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
