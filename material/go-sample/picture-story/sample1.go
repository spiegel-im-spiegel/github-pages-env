// +build run

package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
)

//go:embed html
var assets embed.FS

func main() {
	addr := "localhost:3000"
	fmt.Printf("Open http://%s/\n", addr)
	fmt.Println("Press ctrl+c to stop")

	http.Handle("/", http.FileServer(http.FS(assets)))
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
