+++
title = "golangci-lint に叱られる"
date = "2019-02-06T23:15:43+09:00"
description = "そんなことまで知っている golangci-lint は偉いねぇ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "atom", "tools", "lint" ]

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

例えば [`io`].`Reader` から [`io`].`Writer` へデータを流し込むのにやっつけコードで

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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4542503461?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51CAFNAdZPL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4542503461?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">組込み開発者におくるMISRA‐C:2004―C言語利用の高信頼化ガイド</a></dt>
    <dd>MISRA‐C研究会 (編集)</dd>
    <dd>日本規格協会 2006-10-01</dd>
    <dd>単行本</dd>
    <dd>4542503461 (ASIN), 9784542503465 (EAN), 4542503461 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">私が持っているのはこれよりひとつ古い版だが，まぁいいか。むかし，車載用の組み込みエンジニアをやっていた頃は必読書として読まされました。今はもっと包括的な内容のものがあるはず。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-02-06">2019-02-06</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
