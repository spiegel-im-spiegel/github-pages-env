+++
title = "golangci-lint に叱られる"
date = "2019-02-06T23:15:43+09:00"
description = "そんなことまで知っている golangci-lint は偉いねぇ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "atom", "tools", "lint" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は [Go 言語]コードを [ATOM] エディタおよび [go-plus] パッケージで書いているのだが，最近の [go-plus] は lint に以下のツールを選択できるようだ。

- [gometalinter]
- [golangci-lint]
- [revive]

特に [golangci-lint] は [gometalinter] より5倍も速いと豪語してるので，こちらを試してみることにした。
[GolangCI] も気になるが，それはまたいつか。

いやぁ，最近の lint は賢いんだね。
特に古いコードのまま放置している部分についてどえら叱られたですよ。
ボーっと生きててすみません（笑）

## error を無視すんな

例えば [`io`]`.Reader` から [`io`]`.Writer` へデータを流し込むのにやっつけコードで

```go
io.Copy(writer, reader)
```

とか書くことがあるが [golangci-lint] にかけたら「返り値の error を無視すんなや」って叱られた。
ちゃんと書くなら

```go
if err := io.Copy(writer, reader); err != nil {
    ...
}
```

とかすべき，ということだろう。
明示的に返り値の error を無視するなら

```go
_ = io.Copy(writer, reader)
```

などと書けば，とりあえず叱られない。
まぁ安直にこう書いてしまうのは問題だけど。

## true は不要

無限ループについて昔は何となく

```go
for true {
    ...
}
```

とか書いていて，またも「true とか要らんけぇ」と叱られた。
正しくは

```go
for {
    ...
}
```

でよい。
こういう「若気の至り」なコードがそこかしこに残ってて，黒歴史を見せられてるようでちょっと恥ずかしい。

## SIGNAL のバッファリング

SIGNAL を channel として登録する際に誤って

```go
ch := make(chan os.Signal)
signal.Notify(ch, sig...)
```

とか書いていて「ちゃんとバッファリングさせろや，ゴラァ」とまたまた叱られた。
正解は

```go
ch := make(chan os.Signal, 1)
signal.Notify(ch, sig...)
```

そんなことまで知っている [golangci-lint] は偉いねぇ。

## Lint は知見のかたまり

というわけで過去の恥ずかしいコードが次々と発覚してしまい，公開しているコードのリファクタリングを行っている真っ最中である。
色々やりたいことがあるのに横道に逸れてばっかり。

C 言語で書いてた頃は MISRA-C とかいった「事実上の標準」みたいなのがあって lint ツールとかも（少なくともエンタープライズ向けは）そういったものに準拠したものが色々あった。
Lint ってのはそのプログラミング言語に関する知見のかたまりなので上手に使いこなしていきたいものである。

## ブックマーク

- [Big Sky :: Lint ツールを Language Server に対応させるコマンド efm-langserver 作った。](https://mattn.kaoriya.net/software/lang/go/20190205190203.htm)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[ATOM]: https://atom.io/ "Atom"
[go-plus]: https://atom.io/packages/go-plus
[gometalinter]: https://github.com/alecthomas/gometalinter "alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"
[golangci-lint]: https://github.com/golangci/golangci-lint "golangci/golangci-lint: Linters Runner for Go. 5x faster than gometalinter. Nice colored output. Can report only new issues. Fewer false-positives. Yaml/toml config."
[GolangCI]: https://golangci.com/ "Automated code review for Go"
[revive]: https://github.com/mgechev/revive "mgechev/revive: 🔥 ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for golint."
[`io`]: https://golang.org/pkg/io/ "io - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E7%B5%84%E8%BE%BC%E3%81%BF%E9%96%8B%E7%99%BA%E8%80%85%E3%81%AB%E3%81%8A%E3%81%8F%E3%82%8BMISRA%E2%80%90C-2004%E2%80%95C%E8%A8%80%E8%AA%9E%E5%88%A9%E7%94%A8%E3%81%AE%E9%AB%98%E4%BF%A1%E9%A0%BC%E5%8C%96%E3%82%AC%E3%82%A4%E3%83%89-MISRA%E2%80%90C%E7%A0%94%E7%A9%B6%E4%BC%9A/dp/4542503461?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4542503461"><img src="https://images-fe.ssl-images-amazon.com/images/I/51CAFNAdZPL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E7%B5%84%E8%BE%BC%E3%81%BF%E9%96%8B%E7%99%BA%E8%80%85%E3%81%AB%E3%81%8A%E3%81%8F%E3%82%8BMISRA%E2%80%90C-2004%E2%80%95C%E8%A8%80%E8%AA%9E%E5%88%A9%E7%94%A8%E3%81%AE%E9%AB%98%E4%BF%A1%E9%A0%BC%E5%8C%96%E3%82%AC%E3%82%A4%E3%83%89-MISRA%E2%80%90C%E7%A0%94%E7%A9%B6%E4%BC%9A/dp/4542503461?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4542503461">組込み開発者におくるMISRA‐C:2004―C言語利用の高信頼化ガイド</a></dt>
	<dd>MISRA‐C研究会 (編集)</dd>
    <dd>日本規格協会 2006-10</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4542503461, EAN: 9784542503465</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">私が持っているのはこれよりひとつ古い版だが，まぁいいか。むかし，車載用の組み込みエンジニアをやっていた頃は必読書として読まされました。今はもっと包括的な内容のものがあるはず。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-06">2019-02-06</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
