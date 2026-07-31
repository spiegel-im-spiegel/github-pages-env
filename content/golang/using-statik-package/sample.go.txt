package main

import (
	"fmt"
	"net/http"
	"os"

	"sample/statik"
)

//go:generate statik -src html

func main() {
	fmt.Println("Open http://localhost:3000/")
	fmt.Println("Press ctrl+c to stop")

	statikFs, err := statik.New()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	http.Handle("/", http.FileServer(statikFs))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
