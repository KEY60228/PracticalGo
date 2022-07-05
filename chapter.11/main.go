package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	transport http.RoundTripper
	logger    func(string, ...interface{})
}

func (t *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.logger == nil {
		t.logger = log.Printf
	}

	start := time.Now()
	res, err := t.transport.RoundTrip(req)

	if res != nil {
		t.logger("%s %s %s, duration: %d", req.Method, req.URL.String(), res.Status, time.Since(start))
	}

	return res, err
}

func main() {
	client := &http.Client{
		Transport: &loggingRoundTripper{
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

	// body, err := ioutil.ReadAll(res.Body)
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(body))
}
