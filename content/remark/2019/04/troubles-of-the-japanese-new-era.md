+++
title = "まだ改元で消耗してるの？"
date = "2019-04-17T13:47:50+09:00"
description = "もう役所の書類で元号使うのやめようや。今回の改元は絶好の機会だったのに，これを逃したらまた数十年も「令和」に付き合わされるわけだ。"
image = "/images/attention/kitten.jpg"
tags = [ "calendar", "engineering", "time" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

全国の職業エンジニアの皆さん，改元に伴うあれこれでお疲れ様でした。
もうとっくに改修作業は終わって5月の改元をドキワクで待っていることと思います。

と思ってたらこの期に及んで「テスト」と称して本番環境を弄って案の定トラブらせている自治体があったらしい。
他でもない，我が故郷の松江市だよ（日本の新聞サイトにはリンクしないようにしてるので元記事は省略）。

なんでも本番環境でテストするのにコンビニに通達するのを忘れていて「令和」の元号で証明書類を出力させちゃったらしい。
松江市は回収しようとしているようだが，そんなレアなエラーを手放すわけないぢゃん（笑）

金払ってでもコンビニで証明書を出すべきだったか。
いや，よく考えたら私はマイナンバーカードで手続きしてるから，そもそも紙の証明書は不要だった。

ていうか改元まであと2週間というタイミングで何でテストなんかしてるかなぁ。
しかも改元前の日付で「令和」で出てしまうというのはバグなのか仕様なのか。

もしかして松江市は改元のタイミングでプログラムを入れ替えるつもりなのだろうか。
それって絶対にトラブルの元だよな。

あー，ツッコミが止まんねー！

もう役所の書類で元号使うのやめようや。
これから先，改元の度にこんな馬鹿騒ぎをするつもりなのだろうか。
今回の改元は（期間的な余裕もあったし）ドキュメントに記載する年号を和暦から西暦に切り替える絶好の機会だったのに，これを逃したらまた数十年も「令和」に付き合わされるわけだ。

## 元号で遊ぶ

ところで

- [令和へ対応せよ！元号のアルゴリズム - Qiita](https://qiita.com/sakdor/items/e13421c60438408e0e5b)

という記事を見かけたのだが，みんなそんな難しいことをしているのだろうか。
つか年月日で桁を揃えて比較するってのは汎用機時代の発想だよな。
そういうのって脈々と受け継がれるものなのかねぇ。

今どきは大抵の言語で時間クラスないしは時間関数を標準ライブラリで持ってるんだから，それを使えば簡単にできるぢゃん。

まずは各元号の起点を調べておく。

| 元号       | 起点           |
| ---------- | -------------- |
| 明治の改暦 | 1873年1月1日   |
| 大正       | 1912年7月30日  |
| 昭和       | 1926年12月25日 |
| 平成       | 1989年1月8日   |
| 令和       | 2019年5月1日   |

なんで明治は「元年」を起点にしないかというと，明治6年（1873年）より前は暦が異なるため現行暦[^c1] の加減算が使えないからである。

[^c1]: ちなみに日本の現行暦は「グレゴリオ暦と同じ」だがグレゴリオ暦ではない。暦の原点が異なるからだ。詳しくは拙文「[「暦」日本史]({{< ref "/remark/2015/japanese-koyomi.md" >}})」を参考にどうぞ。

じゃあ，この情報を元に西暦を和暦に変換する簡単なコードを書いてみよう。
[Go 言語]でね。

こんな感じかな。

```go
package main

import (
    "flag"
    "fmt"
    "os"
    "strconv"
    "time"
)

type EraName int

const (
    Unknown EraName = iota
    Meiji
    Taisho
    Showa
    Heisei
    Reiwa
)

var (
    eraString = map[EraName]string{
        Unknown: "",
        Meiji:   "明治",
        Taisho:  "大正",
        Showa:   "昭和",
        Heisei:  "平成",
        Reiwa:   "令和",
    }
    locJST     = time.FixedZone("JST", 9*60*60)
    eraTrigger = map[EraName]time.Time{
        Meiji:  time.Date(1873, time.January, 1, 0, 0, 0, 0, locJST),
        Taisho: time.Date(1912, time.July, 30, 0, 0, 0, 0, locJST),
        Showa:  time.Date(1926, time.December, 25, 0, 0, 0, 0, locJST),
        Heisei: time.Date(1989, time.January, 8, 0, 0, 0, 0, locJST),
        Reiwa:  time.Date(2019, time.May, 1, 0, 0, 0, 0, locJST),
    }
    eraSorted = []EraName{Reiwa, Heisei, Showa, Taisho, Meiji}
)

func (e EraName) String() string {
    if s, ok := eraString[e]; ok {
        return s
    }
    return ""
}

type JapaneseEra struct {
    time.Time
}

func New(t time.Time) JapaneseEra {
    return JapaneseEra{t.In(locJST)}
}

func (e JapaneseEra) Era() EraName {
    for _, es := range eraSorted {
        if !e.Before(eraTrigger[es]) {
            return es
        }
    }
    return Unknown

}

func (e JapaneseEra) YearEra() (EraName, int) {
    era := e.Era()
    if era == Unknown {
        return Unknown, 0
    }
    year := e.Year() - eraTrigger[era].Year() + 1
    if era == Meiji {
        return era, year + 5
    }
    return era, year
}

func (e JapaneseEra) YearEraString() (string, error) {
    era, year := e.YearEra()
    if era == Unknown {
        return "", fmt.Errorf("out of range: %v", e)
    }
    if year == 1 {
        return fmt.Sprintf("%v元年", era), nil
    }
    return fmt.Sprintf("%v%d年", era, year), nil
}

func main() {
    flag.Parse()
    argsStr := flag.Args()
    tm := time.Now()
    if len(argsStr) > 0 {
        if len(argsStr) < 3 {
            fmt.Fprintln(os.Stderr, "年月日を指定してください")
            return
        }
        args := make([]int, 3)
        for i := 0; i < 3; i++ {
            num, err := strconv.Atoi(argsStr[i])
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                return
            }
            args[i] = num
        }
        tm = time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.Local)
    }
    ye := New(tm)
    ys, err := ye.YearEraString()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
    fmt.Printf("%s%d月%d日\n", ys, ye.Month(), ye.Day())
}
```

さっそく動かしてみる。

```text
$ go run main.go
平成31年4月17日

$ go run main.go 2019 4 30
平成31年4月30日

$ go run main.go 2019 5 1
令和元年5月1日
```

ほら，簡単でしょ。
余暇のやっつけコードでもこの程度は書けるってことだね。

やぁ，遊んだ遊んだ。

## ブックマーク

- [1l0/sumeragi](https://github.com/1l0/sumeragi) : 皇紀や元号を出力するパッケージ
- [新元号「文字」という技術的負債]({{< ref "/remark/2017/12/character-of-the-new-era-name.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
