package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
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
		r := rand.New(mt19937.New())
		r.Seed(time.Now().UnixNano())
		for {
			ch <- zundoko[r.Intn(2)]
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
