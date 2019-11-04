+++
date = "2016-05-21T16:27:42+09:00"
description = "ちょっとしたことなんだけどね。キャラクタ端末メインで仕事してる時はこういうのが便利だったりする。"
draft = false
tags = ["golang", "search", "google"]
title = "コマンドラインからググる"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

面白いツールが公開されている。

- [コマンドラインからググれてもいいと思ったので作った - Qiita](http://qiita.com/ieee0824/items/13435fc6de5f22cdb2f4)

もちろん Windows のコマンドプロンプトでも使える。
たとえば

```text
$ ggr -i ググレカス
```

と打ち込めば，既定のブラウザが開いて

{{< fig-img src="https://photo.baldanders.info/flickr/image/27143946615_m.png" title="ググレカス" link="https://photo.baldanders.info/flickr/27143946615/" >}}

などと表示される。
ちょっとしたことなんだけどね。
キャラクタ端末メインで仕事してる時はこういうのが便利だったりする。
面白いのでうちの子として迎え入れることにした。

[コード](https://github.com/ieee0824/ggr)を見たら [main.go](https://github.com/ieee0824/ggr/blob/master/main.go) に全部入っていて，それはそれで全然いいんだけど，折角なので fork してパッケージ化してみることにした。

- [spiegel-im-spiegel/ggr](https://github.com/spiegel-im-spiegel/ggr)

これで main 関数（[ggr/cli.go](https://github.com/spiegel-im-spiegel/ggr/blob/master/ggr/cli.go)）はこうなる。

```go
package main

import (
    "flag"
    "os"

    "github.com/spiegel-im-spiegel/ggr"
    "github.com/toqueteos/webbrowser"
)

// sample url
// https://www.google.co.jp/search?q=test&ie=utf-8&oe=utf-8&hl=ja

func main() {
    var (
        imageFlag bool
        newsFlag  bool
        shopFlag  bool
    )

    f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
    f.BoolVar(&imageFlag, "i", false, "image flag")
    f.BoolVar(&newsFlag, "n", false, "new flag")
    f.BoolVar(&shopFlag, "s", false, "shop flag")

    f.Parse(os.Args[1:])
    q := f.Args()

    t := ggr.TypeNormal
    if imageFlag {
        t = ggr.TypeImage
    } else if newsFlag {
        t = ggr.TypeNews
    } else if shopFlag {
        t = ggr.TypeShop
    }

    g := ggr.NewGgr(ggr.LangJa, t, q)
    webbrowser.Open(g.GetSearchURL())
}
```

まぁ，パッケージ化していいことがあるかと言われれば微妙なんだけど，最近は CLI ツールでもロジック（DDD で言うところのドメイン・レイヤ）はパッケージとして独立させたほうが何かと使い勝手がいいような気がしている。

というわけで，絶賛中断中の [gpgpdump](https://github.com/spiegel-im-spiegel/gpgpdump) は全面的に書きなおす予定。

いやぁ，最近仕事（Java アプリケーション）で煮詰まってたので，いい気分転換になったよ。
やはり自分の自由に書けるコードは楽しい。

（ちなみに，いつものようにブランチでの作業を GitHub 上で自分に PR しようとして間違えて作者様に PR してしまったのは内緒だ。 Fork したリポジトリからの PR は必ず Fork 元に飛んじゃんだね。もうしません。反省）

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
