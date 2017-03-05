+++
draft = false
tags = ["golang", "programming", "time", "calendar"]
date = "2017-03-04T09:40:51+09:00"
update = "2017-03-04T17:27:54+09:00"
title = "「プレミアムフライデー」を求めるパッケージを作ってみた"
description = "もちろん息抜きである。潤いは大事。でも実用性はないと思われ。"

[author]
  flickr = "spiegel"
  license = "by-sa"
  name = "Spiegel"
  flattr = "spiegel"
  linkedin = "spiegelimspiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  tumblr = "spiegel-im-spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
+++

そういえば先月の「プレミアムフライデー」，皆様はいかがお過ごしでしたか。
私は3時間も残業してしまいましたよ（笑）

ちうわけで，以下を真似して「プレミアムフライデー」を求めるパッケージを考えてみる。
もちろん息抜きである。
潤いは大事。
でも実用性はないと思われ。

- [プレミアムフライデーを求めるメソッドを作った - Qiita](http://qiita.com/neko_the_shadow/items/4ebf94a8a6d9282e7207)
- [プレミアムフライデーを求めるメソッドを作った（Java8版） - Qiita](http://qiita.com/deaf_tadashi/items/963a62072338f09f12a5)

まずはパッケージ分割しないでベタに書いてみる[^rf1]。

[^rf1]: 元記事のコードがループさせてたんでこっちもついループさせちゃったけど，考えてみれば（いや考えるまでもなく）ループを回す必要はなかった。

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

//GetPremiumFriday returns day of premium friday
func GetPremiumFriday(y int, m time.Month) (int, error) {
	//引数のチェック
    if y < 2017 || m < time.January || m > time.December {
		return 0, os.ErrInvalid
	}
	if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
		return 0, os.ErrInvalid
	}

	//指定月末（翌月0日）で初期化する
    tm := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で

    w := tm.Weekday() - time.Friday
	if w < 0 {
		w += 7
	}
	return tm.Day() - (int)(w), nil
}

func main() {
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 2 {
		fmt.Fprintln(os.Stderr, "年月を指定してください")
		return
	}
	args := make([]int, 2)
	for i := 0; i < 2; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		args[i] = num
	}
	d, err := GetPremiumFriday(args[0], time.Month(args[1]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(d)
}
```

指定月末を求めるのに「翌月0日」で初期化するのがポイント（つか，ここしかポイントになるものがない`w`）。
実行結果は以下の通り。

```text
$ go run pf.go 2017 2
24
```

期待通りの値が得られた。

関数1個だけなんでパッケージにするのもどうかと思うけど折角なのでパッケージ化してみた。

- [spiegel-im-spiegel/pf: Premium Friday](https://github.com/spiegel-im-spiegel/pf)

内容は `GetPremiumFriday()` 関数を切り出しただけ。

```go
package pf

import (
	"os"
	"time"
)

//GetPremiumFriday returns day of premium friday
func GetPremiumFriday(y int, m time.Month) (int, error) {
	//引数のチェック
    if y < 2017 || m < time.January || m > time.December {
		return 0, os.ErrInvalid
	}
	if y == 2017 && m < time.February { //2017年1月は実施前なのでエラー
		return 0, os.ErrInvalid
	}

	//指定月末（翌月0日）で初期化する
    tm := time.Date(y, m+1, 0, 0, 0, 0, 0, time.UTC) //時差は影響しないので，とりあえず UTC で

    w := tm.Weekday() - time.Friday
	if w < 0 {
		w += 7
	}
	return tm.Day() - (int)(w), nil
}
```

したがって `main()` 関数はこうなる。

```go
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spiegel-im-spiegel/pf"
)

func main() {
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 2 {
		fmt.Fprintln(os.Stderr, "年月を指定してください")
		return
	}
	args := make([]int, 2)
	for i := 0; i < 2; i++ {
		num, err := strconv.Atoi(argsStr[i])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		args[i] = num
	}
	d, err := pf.GetPremiumFriday(args[0], time.Month(args[1]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(d)
}
```

まぁこんなもんかな。
遊んだ遊んだ。

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
