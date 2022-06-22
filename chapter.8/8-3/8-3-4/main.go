package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gocarina/gocsv"
	"github.com/xuri/excelize/v2"
)

type excelCSVReader struct {
	rows *excelize.Rows
}

func NewExcelCSVReader(filename string, sheet string) (*excelCSVReader, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, err
	}

	rows, err := f.Rows(sheet)
	if err != nil {
		return nil, err
	}

	return &excelCSVReader{rows}, nil
}

func (r excelCSVReader) Read() ([]string, error) {
	if r.rows.Next() {
		return r.rows.Columns()
	}
	return nil, io.EOF
}

func (r excelCSVReader) ReadAll() ([][]string, error) {
	var res [][]string
	for r.rows.Next() {
		columns, err := r.rows.Columns()
		if err != nil {
			return nil, err
		}
		res = append(res, columns)
	}
	return res, nil
}

type Country struct {
	Name       string `csv:"国名"`
	ISOCode    string `csv:"ISOコード"`
	Population int    `csv:"人口"`
}

func main() {
	reader, err := NewExcelCSVReader("Book2.xlsx", "Sheet1")
	if err != nil {
		log.Fatal(err)
	}

	var countries []Country
	if err := gocsv.UnmarshalCSV(reader, &countries); err != nil {
		log.Fatal(err)
	}

	fmt.Println(countries)
}
