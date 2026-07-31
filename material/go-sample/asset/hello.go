package main

import (
	"fmt"
	"net/http"
	"os"
)

//go:generate go-assets-builder -s="/html" -o assets.go html/

func main() {
	fmt.Println("Open http://localhost:3000/")
	fmt.Println("Press ctrl+c to stop")

	http.Handle("/", http.FileServer(Assets))
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
