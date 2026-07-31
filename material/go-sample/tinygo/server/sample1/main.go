package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"strings"
)

//go:embed hello
var assets embed.FS

func main() {
	addr := "localhost:3000"
	fmt.Printf("Open http://%s/\n", addr)
	fmt.Println("Press ctrl+c to stop")

	root, _ := fs.Sub(assets, "hello")
	fs := http.FileServer(http.FS(root))
	http.Handle("/", http.FileServer(http.FS(root)))
	if err := http.ListenAndServe(addr, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	})); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
