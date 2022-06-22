package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type record struct {
	Number  int    `csv:"number"`
	Message string `csv:"message"`
}

func main() {
	f, err := os.OpenFile("sample.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c := make(chan interface{})
	go func() {
		defer close(c)
		for i := 0; i < 1000; i++ {
			c <- record{
				Message: "Hello",
				Number:  i + 1,
			}
		}
	}()

	if err := gocsv.MarshalChan(c, gocsv.DefaultCSVWriter(f)); err != nil {
		log.Fatal(err)
	}
}
