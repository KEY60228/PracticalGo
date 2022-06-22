package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func main() {
	f, err := os.Open("country.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []Country
	if err := gocsv.UnmarshalFile(f, &lines); err != nil {
		log.Fatal(err)
	}

	for _, v := range lines {
		fmt.Printf("%+v\n", v)
	}
}
