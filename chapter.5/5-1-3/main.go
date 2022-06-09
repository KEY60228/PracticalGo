package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HTTPError struct {
	StatusCode int
	URL        string
}

func (he *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url = %s", he.StatusCode, he.URL)
}

func ReadContents(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode: res.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(res.Body)
}

func main() {
	contents, err := ReadContents("https://yahoo.co.jp/")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contents))
}
