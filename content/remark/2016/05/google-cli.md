+++
date = "2016-05-21T16:27:42+09:00"
description = "ちょっとしたことなんだけどね。キャラクタ端末メインで仕事してる時はこういうのが便利だったりする。"
draft = false
tags = ["golang", "search", "google"]
title = "コマンドラインからググる"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

面白いツールが公開されている。

- [コマンドラインからググれてもいいと思ったので作った - Qiita](http://qiita.com/ieee0824/items/13435fc6de5f22cdb2f4)

もちろん Windows のコマンドプロンプトでも使える。
たとえば

```text
$ ggr -i ググレカス
```

と打ち込めば，既定のブラウザが開いて

{{< fig-img flickr="true" src="https://farm8.staticflickr.com/7117/27143946615_45a9c04842.jpg" title="ググレカス" link="https://www.flickr.com/photos/spiegel/27143946615/" >}}

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

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41aCueik45L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-15</dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117402/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117402.09._SCTHUMBZZZ_.jpg"  alt="ハイパフォーマンスPython"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/0134190440/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/0134190440.09._SCTHUMBZZZ_.jpg"  alt="The Go Programming Language (Addison-Wesley Professional Computing Series)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166340/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774166340.09._SCTHUMBZZZ_.jpg"  alt="Vim script テクニックバイブル ~Vim使いの魔法の杖"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。買おうかどうか悩み中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-12">2016-03-12</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
