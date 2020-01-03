+++
title = "Go モジュールのミラーリングとインデックス化"
date =  "2019-05-31T23:15:13+09:00"
description = "モジュールのミラーリング・サービスは各所で公開されているモジュールのミラーリングを行うためのプロキシ・サーバの一種で，現在ベータ版である Go 1.13 では既定で有効になっているそうだ。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "module", "engineering", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（この記事は [golang セクションから移動]({{< ref "/golang/mirror-index-and-checksum-database-for-go-module.md" >}} "Go モジュールのミラーリングとインデックス化")させてきたものである）

[Go] モジュールのミラーリングとインデックス化を行うためのサービスがベータ・リリースされたようだ。

- [Go Module Mirror and Checksum Database in Beta! - Google Group](https://groups.google.com/forum/#!topic/golang-announce/0wo8cOhGuAI)

{{< div-box type="markdown" >}}
**【2019-09-01 追記】**

2019年8月に [proxy.golang.org], [sum.golang.org] および [index.golang.org] が正式にサービスを開始した。

- [Module Mirror and Checksum Database Launched - The Go Blog](https://blog.golang.org/module-mirror-launch)

[proxy.golang.org]: https://proxy.golang.org/
[sum.golang.org]: https://sum.golang.org/
[index.golang.org]: https://index.golang.org/
{{< /div-box >}}

これは “[Go Modules in 2019](https://blog.golang.org/modules2019 "Go Modules in 2019 - The Go Blog")” で予告されていたものだ。

{{< fig-quote title="Go Modules in 2019" link="https://blog.golang.org/modules2019" lang="en" >}}
<q>We are planning to launch a mirror service for publicly-available modules in 2019.</q>
{{< /fig-quote >}}

モジュールのミラー・サービスは各所で公開されているモジュールのミラーリングを行うためのプロキシ・サーバの一種で，現在ベータ版である [Go] 1.13 では既定で有効になっているそうだ。
それ以前のバージョン（1.12 ?）では環境変数 `GOPROXY` に [`https://proxy.golang.org`](https://proxy.golang.org/ "Go modules services") をセットすることで有効になるらしい（試してない）。

更に，モジュールのインデックス・サービスも開始される。

- [`index.golang.org`](https://index.golang.org/)

クエリに対して JSON データを返す仕様なのかな。

更に更に，モジュールのチェックサム値をデータベース化してモジュールの検証に使えるようにするようだ。

- [`sum.golang.org`](https://sum.golang.org/)

チェックサム値をデータベースとして保持っておくことでモジュールの完全性を担保し，コードの改竄を検知しやすくする目的があると思われる[^i1]。
チェックサム・データベースの利用については [gosumcheck](https://godoc.org/golang.org/x/exp/sumdb/gosumcheck "gosumcheck - GoDoc") というツールが提供されている。

[^i1]: 本当に完全性を担保したいなら電子署名と組み合わせるべきだと思うけどね。まぁ，そこまで厳密な管理は（今のところ）要らないと考えているのだろう。今や GnuPG と OpenSSH と Git はワンセットなので（[OpenSSH の鍵は GnuPG で管理可能]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}} "GnuPG for Windows : gpg-agent について")） OpenPGP で電子署名すればいいじゃない，と思うのだが。

```text
$ go get golang.org/x/exp/sumdb/gosumcheck
$ gosumcheck /path/to/go.sum
```

これらのサービスのプライバシー・ポリシーについては [`proxy.golang.org/privacy`](http://proxy.golang.org/privacy) を参照しろとあるが，この URL を叩くと [Google のプライバシー・ポリシー](https://policies.google.com/privacy)のページに飛ばされる。
まぁ Google のサービスなんだから当然だろうけど，[プライバシーに敵対的な企業]({{< ref "/remark/2018/04/handling-privacy.md" >}} "誰がプライバシーを支配するのか")のサービスだと思うとあまり利用したくない気分[^lang1]。

[^lang1]: [Go 言語]にしろ（最近ブームが再燃している） Dart 言語にしろ，言語系のプロダクトやサービスは Google から切り離してほしいよなぁ。 Alphabet の傘下から外れろとまでは言わないから。

モジュールのミラーリングやインデックス化はそれぞれ単体のサービスではなく，モジュールを中心とした生態系（module ecosystem）を構築するための部品と考えるのがいいだろう。
[Go] 1.13 以降からはモジュール周りが更に大きく変わりそうである。

## ブックマーク

- [Proposal: Secure the Public Go Module Ecosystem](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md)

- [Go 1.13 と 1.14 （Go 2 へ向けて）]({{< ref "/release/2019/06/next-steps-toward-go-2.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Go]: https://golang.org/ "The Go Programming Language"

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
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E8%B6%85%E4%BA%BA%E3%83%AD%E3%83%83%E3%82%AF-%E3%83%9F%E3%83%A9%E3%83%BC%E3%83%AA%E3%83%B3%E3%82%B0/dp/B013URFRLQ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B013URFRLQ"><img src="https://images-fe.ssl-images-amazon.com/images/I/C1rL0TB3gnS._SL160_.png" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E8%B6%85%E4%BA%BA%E3%83%AD%E3%83%83%E3%82%AF-%E3%83%9F%E3%83%A9%E3%83%BC%E3%83%AA%E3%83%B3%E3%82%B0/dp/B013URFRLQ?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B013URFRLQ">[まとめ買い] 超人ロック　ミラーリング</a></dt>
    <dd>聖 悠紀</dd>
    <dd></dd>
    <dd>Kindle版</dd>
    <dd>B013URFRLQ (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">「エレナ」続編。分散 AI といえばこれだよね。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-12-28">2017-12-28</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
