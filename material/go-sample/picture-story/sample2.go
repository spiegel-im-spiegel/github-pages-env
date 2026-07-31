// +build run

package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
)

//go:embed html
var assets embed.FS

func main() {
	addr := "localhost:3000"
	fmt.Printf("Open http://%s/\n", addr)
	fmt.Println("Press ctrl+c to stop")

	root, _ := fs.Sub(assets, "html")
	http.Handle("/", http.FileServer(http.FS(root)))
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
