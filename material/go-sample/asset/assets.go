package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets950a25363eee220f7d8ce234bcc0b349e4ea9072 = "<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset=\"utf-8\">\n\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n\t<title>Hello World!</title>\n</head>\n<body>\n\t<p>Hello World!</p>\n</body>\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518522609, 1518522609437368800),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1518523737, 1518523737979915400),
		Data:     []byte(_Assets950a25363eee220f7d8ce234bcc0b349e4ea9072),
	}}, "")
