package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

//go:embed html
var assets embed.FS

func main() {
	port := flag.Int("p", 3000, "port number")
	host := flag.String("host", "", "hostname")
	flag.Parse()

	addr := net.JoinHostPort(*host, strconv.Itoa(*port))
	if len(*host) == 0 {
		fmt.Printf("Open http://localhost%s/\n", addr)
	} else {
		fmt.Printf("Open http://%s/\n", addr)
	}
	fmt.Println("Press ctrl+c to stop")

	root, err := fs.Sub(assets, "html")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	serverMux := http.NewServeMux()
	serverMux.Handle("/", http.FileServer(http.FS(root)))
	server := &http.Server{
		Addr:              addr,
		Handler:           serverMux,
		ReadHeaderTimeout: 3 * time.Second,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		defer close(idleConnsClosed)
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := server.Shutdown(context.Background()); err != nil {
			fmt.Fprintln(os.Stderr, "shutdown error:", err)
			return
		}
		fmt.Println("\ngraceful shutdown")
	}()

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		if err != nil {
			fmt.Fprintln(os.Stderr, "server error:", err)
		}
		return
	}
	<-idleConnsClosed
}
