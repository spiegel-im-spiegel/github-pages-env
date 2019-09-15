+++
title = "Go 言語用エラーハンドリング・パッケージをリリースした"
date =  "2019-09-05T23:54:45+09:00"
description = "実は以前にこっそり v0.1.0 をリリースして自作ツールのエラーハンドリングに用いていたのだが， Go 1.13 のリリースに合わせて中身を作り直した。"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 言語用エラーハンドリング・パッケージ [`errs`] の v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.0)

実は以前にこっそり v0.1.0 をリリースして自作ツールのエラーハンドリングに用いていたのだが， [Go] 1.13 のリリースに合わせて中身を作り直した。

使い方は以下を参照のこと。

- [Go 言語用エラーハンドリング・パッケージ]({{< ref "/release/errs-package-for-golang.md" >}})

これで先に進めるな。

## 【2019-09-06 追記】 v0.2.1 をリリースした

- [Release v0.2.1 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.1)

つか，実はバージョンタグを付け間違えただけなのだが，バージョンタグを付け換えるとチェックサム・データベース [`sum.golang.org`] が怒って

```text
SECURITY ERROR
This download does NOT match the one reported by the checksum server.
The bits may have been replaced on the origin server, or an attacker may
have intercepted the download attempt.

For more information, see 'go help module-auth'.
```

とか言ってくさるので，しょうが無しにバージョン番号を上げることにした。
やれやれ。

バージョンタグの管理は慎重に。

[`sum.golang.org`]: https://sum.golang.org/

## 【2019-09-15 追記】 v0.2.2 をリリースした

- [Release v0.2.2 · spiegel-im-spiegel/errs · GitHub](https://github.com/spiegel-im-spiegel/errs/releases/tag/v0.2.2)

Formatting に関するバグを修正した。

## ブックマーク

- [Go 1.13 のエラー・ハンドリング]({{< ref "/golang/error-handling-in-go-1_3.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`errs`]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
    <dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
