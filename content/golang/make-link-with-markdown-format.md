+++
title = "Markdown 形式のリンクを生成するツールを作ってみた"
date =  "2017-11-08T18:37:57+09:00"
update = "2017-12-14T15:51:24+09:00"
description = "あれ？ これ Go 言語でも簡単に作れるんじゃないかな。ちうわけで作ってみた。"
tags        = [ "golang", "programming", "tools" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

きっかけはこれ。

- [Markdownのリンクフォーマットを生成するコマンドを自作した - Qiita](https://qiita.com/Struuuuggle/items/08e6c1bf70df55ecc7b5)

仕様はこんな感じ。

{{< fig-quote title="Markdownのリンクフォーマットを生成するコマンドを自作した" link="https://qiita.com/Struuuuggle/items/08e6c1bf70df55ecc7b5" >}}
<q>URLを入力すると、Markdownに最適な形式で、webページタイトルとURLを出力します。出力結果はコンソールに表示すると同時に、クリップボードにもコピーされます。</q>
{{< /fig-quote >}}

あれ？ これ [Go 言語]でも簡単に作れるんじゃないかな。
ちうわけで作ってみた。

- [spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)

（Windows では `mklink` はシンボリック・リンクを作るコマンドだと後で気づいたが，後悔先に立たず`w`）

パッケージの使い方としてはこんな感じである。

```go
link, err := mklink.New("https://git.io/vFR5M")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(link.Encode(mklink.StyleMarkdown))
// Output:
// [GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

コマンドライン・インタフェースはこんな感じ。

```text
$ mklink -h
Usage:
  mklink [flags] [URL [URL]...]

Flags:
  -h, --help           help for mklink
  -i, --interactive    interactive mode
      --log string     output log
  -s, --style string   link style (default "markdown")
  -v, --version        output version of mklink
```

リンクの形式が markdown だけなのはアレなので `-s` オプションで選べるようにした。
今のところ `markdown`, `wiki`, `html`, `csv` の4つに対応している。

`-i` オプションを付けると対話モードになる。

```text
$ mklink -i
Input 'q' or 'quit' to stop
mklink> https://git.io/vFR5M
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
mklink>
```

作成したリンクを標準出力に出力すると同時にクリップボードにもコピーする。
いやぁ，これめっさ便利だわ。

## Web Scraping

URL からページのタイトルを取得するには HTML の解析を行うスクレイピング（Web scraping）機能が必要だが，好都合なことに [PuerkitoBio/goquery] という便利なパッケージが公開されている。

- [PuerkitoBio/goquery: A little like that j-thing, only in Go.](https://github.com/PuerkitoBio/goquery)

[PuerkitoBio/goquery] が優れているのは jQuery っぽい仕掛けでとても簡単に HTML の解析ができる点にある。
今回は `<head>` 要素からタイトルと説明（description）を抜き出すだけだが，こんな感じに記述できる。

```go
//New returns new Link instance
func New(url string) (*Link, error) {
    link := &Link{URL: strings.Trim(url, "\t \n")}
    doc, err := goquery.NewDocument(link.URL)
    if err != nil {
        return link, err
    }
    link.Location = doc.Url.String()

    doc.Find("head").Each(func(_ int, s *goquery.Selection) {
        s.Find("title").Each(func(_ int, s *goquery.Selection) {
            link.Title = strings.Trim(s.Text(), "\t \n")
        })
        s.Find("meta[name='description']").Each(func(_ int, s *goquery.Selection) {
            if v, ok := s.Attr("content"); ok {
                link.Description = strings.Trim(v, "\t \n")
            }
        })
    })

    return link, nil
}
```

ね， jQuery ぽいでしょ。
まさか，この期に及んで jQuery （ぽいもの）を触ることになるとは思わなかったぜ。
jQuery の本は納戸に仕舞っちゃったんだけどなぁ。

## ターミナルの判定

[mklink] はパイプでも動作する。

```text
$ echo https://git.io/vFR5M | mklink
[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
```

この時にうっかり `-i` オプションを付けて（パイプのつもりが）対話モードになっては困るので標準入出力がターミナルかどうかを判定するロジックを入れている。

```go
func isTerminal() bool {
    if !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd()) {
        return false
    }
    if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
        return false
    }
    return true
}
```

この機能は [mattn/go-isatty] パッケージで実装した[^ssh1]。

[^ssh1]: 最初は `golang.org/x/crypto/ssh/terminal` パッケージを使っていたのだが「たしかもう少し軽いパッケージあったよなぁ」と思ってググったら思い出した。

## クリップボードの操作

クリップボードの操作といっても今回は書き込みだけだが [atotto/clipboard] を使って実装している。

{{< highlight go "hl_lines=5" >}}
buf := new(bytes.Buffer)
io.Copy(c.writer, io.TeeReader(lnk.Encode(c.linkStyle), buf))
strLink := buf.String()
if c.clipbrdFlag {
    clipboard.WriteAll(strLink)
}
if c.log != nil {
    fmt.Fprint(c.log, strLink)
}
{{< /highlight >}}

Windows 環境では問題なく動作しているが，他の OS ではどうなのかよく分からない。

## ブックマーク

- [Big Sky :: Go言語で jQuery ライクな操作が出来る goquery を試した。](https://mattn.kaoriya.net/software/lang/go/20120914184828.htm)
- [[golang]遅れながらgoqueryを使ってみた - Qiita](https://qiita.com/pokochi/items/042e91a2e724c336d02d)
- [シェルで短縮URLの展開 - nyaocatのがんばるブログ](http://nyaocat.hatenablog.jp/entry/2012/12/10/235259)
- [ASCII.jp：Go言語で知るプロセス（2）｜Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/459/1459279/)
- [テキストファイルの中身をクリップボードにコピーするGolang製CLIツールを作った | タイトル未定(仮)](https://kogai.github.io/2016/08/25/create-golip/)
- [Go x goqueryでwebスクレイピング - Qiita](https://qiita.com/akif999/items/2d6428c2377e020ce904)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[mklink]: https://github.com/spiegel-im-spiegel/mklink "spiegel-im-spiegel/mklink: Make Link with Markdown Format"
[PuerkitoBio/goquery]: https://github.com/PuerkitoBio/goquery "PuerkitoBio/goquery: A little like that j-thing, only in Go."
[mattn/go-isatty]: https://github.com/mattn/go-isatty
[atotto/clipboard]: https://github.com/atotto/clipboard "atotto/clipboard: clipboard for golang"
