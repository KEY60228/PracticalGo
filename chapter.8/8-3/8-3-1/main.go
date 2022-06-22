package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	in, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	cell, err := in.GetCellValue("Sheet1", "A1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cell)
}
