+++
title = "GNKF: NKF ぽいなにか の v0.4.0 をリリースした"
date =  "2021-02-28T19:26:39+09:00"
description = "部首の正規化の対象に CJK 部首補助を含めるようにした。"
image = "/images/attention/tools.png"
tags  = [ "tools", "gnkf", "golang", "unicode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語][Go]における文字列処理の習作 [gnkf] の v0.4.0 をリリースした。

- [Release v0.4.0 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.4.0)

実は

- [PDFのコピペが文字化けするのはなぜか？～CID/GIDと原ノ味フォント～](https://www.slideshare.net/trueroad_jp/pdfcidgid)

というスライドを見てたんだけど，[康熙部首][Kangxi radical]だけじゃなくて [CJK 部首補助][CJK Radicals]に化けることもあるらしいぢゃん。

[gnkf] では `norm` サブコマンドに `-k` オプションを付けて

```text
$ echo ㈱埼⽟ | gnkf norm -n nfkc -k
```

とかすれば[康熙部首][Kangxi radical]の「⽟（U+2F5F）」だけを正規化できるんだけど，この対象に [CJK 部首補助][CJK Radicals]を含めるようにした。
具体的には Unicode オフィシャルの

- [EquivalentUnifiedIdeograph.txt](https://www.unicode.org/Public/UCD/latest/ucd/EquivalentUnifiedIdeograph.txt)

に従って変換している。

あとは，お試し機能として `completion` サブコマンドを追加して各 shell 用の自動補完スクリプトを吐き出せるようにした。
Bash, Zsh, Fish, PowerShell の自動補完機能に対応している。

```text
 gnkf completions -h
```

とかすれば簡単なヘルプが出るので，よろしければどうぞ。

## ブックマーク

- [やっかいな漢字 – CJK部首補助／康煕部首 – ものかの](https://tama-san.com/resolve-kanji/)

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"
[Kangxi radical]: https://en.wikipedia.org/wiki/Kangxi_radical "Kangxi radical - Wikipedia"
[CJK Radicals]: https://en.wikipedia.org/wiki/CJK_Radicals_Supplement "CJK Radicals Supplement - Wikipedia"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
