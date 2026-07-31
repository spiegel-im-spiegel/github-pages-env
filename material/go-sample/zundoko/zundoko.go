package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	zun     = "ズン"
	doko    = "ドコ"
	kiyoshi = "キ・ヨ・シ！"
)

func generate() chan string {
	ch := make(chan string)
	go func() {
		var zundoko = [2]string{zun, doko}
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- zundoko[rand.Intn(2)]
		}
	}()
	return ch
}

func main() {
	zundoko := generate()
	zcount := 0
	for {
		zd := <-zundoko
		fmt.Print(zd)
		if zd == zun {
			zcount++
		} else if zcount >= 4 {
			break
		} else {
			zcount = 0
		}
	}
	fmt.Print(kiyoshi)
}
