package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strconv"
)

func main() {
	director := func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = "api.mystery-logger.com"
		req.URL.Path = "/v1/products"
	}
	modifier := func(res *http.Response) error {
		body := make(map[string]interface{})
		dec := json.NewDecoder(res.Body)
		dec.Decode(&body)
		body["fortune"] = "Lucky"

		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(&body)

		res.Body = ioutil.NopCloser(&buf)
		res.Header.Set("Content-Length", strconv.Itoa(buf.Len()))
		return nil
	}
	rp := &httputil.ReverseProxy{
		Director:       director,
		ModifyResponse: modifier,
	}

	http.ListenAndServe(":9000", rp)
}
