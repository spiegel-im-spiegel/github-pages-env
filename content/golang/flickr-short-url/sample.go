package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/itchyny/base58-go"
)

func FlickrShortURL(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}
	if !strings.HasSuffix(u.Hostname(), "flickr.com") {
		return s
	}
	var photoId string
	elms := strings.Split(strings.TrimSuffix(u.EscapedPath(), "/"), "/")[1:]
	if len(elms) == 3 && strings.EqualFold(elms[0], "photos") {
		photoId = elms[2]
	}
	if len(photoId) == 0 || !isDigitString(photoId) {
		return s
	}
	b, err := base58.FlickrEncoding.Encode([]byte(photoId))
	if err != nil {
		return s
	}
	return "https://flic.kr/p/" + string(b)
}

func isDigitString(s string) bool {
	return strings.IndexFunc(s, func(c rune) bool {
		return c < '0' || '9' < c
	}) < 0
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	fmt.Println(FlickrShortURL(args[0]))
}
