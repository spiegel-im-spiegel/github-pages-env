package main

import (
	"fmt"
	"strings"

	"github.com/mattn/go-gimei"
	"github.com/spiegel-im-spiegel/krconv"
)

func main() {
	p := gimei.NewName()
	fmt.Println("氏名：", p.Kanji())
	fmt.Println("カナ：", p.Katakana())
	fmt.Println("ローマ字：", strings.ToTitle(krconv.Convert(p.Last.Hiragana())), strings.Title(krconv.Convert(p.First.Hiragana())))
	fmt.Println("Email：", string([]rune(krconv.Convert(p.First.Hiragana()))[0:1])+"."+krconv.Convert(p.Last.Hiragana())+"@example.com")
}
