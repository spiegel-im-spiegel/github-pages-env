package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/goark/csvdata"
	"github.com/goark/utf8bom"
)

func openCsvFile(path string) (*csvdata.Rows, error) {
	file, err := os.Open("./sample3.csv")
	if err != nil {
		return nil, err
	}
	return csvdata.NewRows(csvdata.New(utf8bom.Strip(file)), true), nil
}

func main() {
	file, err := openCsvFile("./sample3.csv")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer file.Close()

	for {
		if err := file.Next(); err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			break
		}
		fmt.Println(file.Row())
	}
}
