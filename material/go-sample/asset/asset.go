package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets1d921dc85677928c7457b1a7687f82618b5cfd6b = "Hello World!\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"data"}, "/data": []string{"hello.txt"}}, map[string]*assets.File{
	"/data": &assets.File{
		Path:     "/data",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518509705, 1518509705225360300),
		Data:     nil,
	}, "/data/hello.txt": &assets.File{
		Path:     "/data/hello.txt",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1518512089, 1518512089919917400),
		Data:     []byte(_Assets1d921dc85677928c7457b1a7687f82618b5cfd6b),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1518507214, 1518507214658243100),
		Data:     nil,
	}}, "")
