package main

import (
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {
	out := excelize.NewFile()
	out.SetCellValue("Sheet1", "A1", "Hello Excel")
	if err := out.SaveAs("Book1.xlsx"); err != nil {
		log.Fatal(err)
	}
}
