+++
title = "「nkf っぽいなにか」の v0.5.2 をリリースした"
date =  "2019-11-16T17:16:49+09:00"
description = "機能の変更はないが Go 1.13.x に対応させるために内部をいろいろ弄った。2017年に v0.5.1 をリリースしてから放ったらかしにしてたからなぁ。"
image = "/images/attention/tools.png"
tags = ["tools", "character", "encoding", "transform"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私はあまり使ってないが「nkf っぽいなにか」の v0.5.2 をリリースした。

- [Release v0.5.2 · spiegel-im-spiegel/text · GitHub](https://github.com/spiegel-im-spiegel/text/releases/tag/v0.5.2)

```text
$ gonkf -h
Network Kanji Filter by Golang

Usage:
  gonkf [flags]
  gonkf [command]

Available Commands:
  conv        Convert character encoding of text
  guess       Guess character encoding of text
  help        Help about any command
  norm        Unicode normalization
  nwline      Convert newline of text
  version     Print the version number of gonkf
  width       Convert character width of text

Flags:
  -h, --help   help for gonkf

Use "gonkf [command] --help" for more information about a command.
```

機能の変更はないが [Go] 1.13.x に対応させるために内部をいろいろ弄った。
2017年に v0.5.1 をリリースしてから放ったらかしにしてたからなぁ。

もともと文字エンコーディング変換の習作用に作ったツールなので，思い入れがないんだよなぁ。
まっ，よろしければ使ってやってください。

## ブックマーク

- [「nkf っぽいなにか」を作った]({{< ref "/remark/2017/12/like-nkf.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
