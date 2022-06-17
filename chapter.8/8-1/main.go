package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type ip struct {
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func main() {
	s := `{"origin": "255.255.255.255", "url": "https://httpbin.org/get"}`
	var res ip
	if err := json.Unmarshal([]byte(s), &res); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
}
