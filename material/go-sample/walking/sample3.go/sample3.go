package main

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/sys/unix"
)

func main() {
	rootPath := "/usr/local/go/src/cmd/go/testdata/script"
	fd, err := syscall.Open(rootPath, 0, 0)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer syscall.Close(fd)

	buf := make([]byte, 8<<10)
	fmt.Println(8 << 10)
	n, err := unix.ReadDirent(fd, buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	names := make([]string, 0, 100)
	offset := 0
	ct := 0
	for {
		consumed, count, names := unix.ParseDirent(buf[offset:n], 100, names[0:])
		offset += consumed
		if count <= 0 {
			break
		}
		for _, name := range names[:count] {
			ct++
			_ = name
			//fmt.Println(name)
		}
	}
	fmt.Println("count", ct)
}
