package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type retryableRoundTripper struct {
	transport http.RoundTripper
	attempts  int
	waitTime  time.Duration
}

func (rt *retryableRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	var res *http.Response
	var err error

	for count := 0; count < rt.attempts; count++ {
		res, err = rt.transport.RoundTrip(req)

		if !rt.shouldRetry(res, err) {
			return res, err
		}

		select {
		case <-req.Context().Done():
			return nil, req.Context().Err()
		case <-time.After(rt.waitTime):
			// リトライのため待機
		}
	}

	return res, err
}

func (rt *retryableRoundTripper) shouldRetry(res *http.Response, err error) bool {
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Temporary() {
			return true
		}
	}

	if res != nil {
		if res.StatusCode == 429 || (500 <= res.StatusCode && res.StatusCode <= 504) {
			return true
		}
	}

	return false
}

func main() {
	client := &http.Client{
		Transport: &retryableRoundTripper{
			transport: http.DefaultTransport,
			attempts:  5,
			waitTime:  1 * time.Minute,
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
