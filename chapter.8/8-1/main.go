package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Rectangle struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

func main() {
	f, err := os.Open("square.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var rect Rectangle
	d := json.NewDecoder(f)
	d.DisallowUnknownFields()
	if err := d.Decode(&rect); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", rect)
}
