+++
tags = ["golang", "programming", "type"]
description = "浮動小数点数型の変数をループカウンタにするのは止めましょうね。約束だよ！"
date = "2017-01-18T21:45:30+09:00"
title = "1を1億回足して1億にならない場合"
draft = false

[author]
  license = "by-sa"
  name = "Spiegel"
  flickr = "spiegel"
  flattr = "spiegel"
  facebook = "spiegel.im.spiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  instagram = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  linkedin = "spiegelimspiegel"
  tumblr = "spiegel-im-spiegel"
+++

（この記事は [Qiita に投稿した記事](http://qiita.com/spiegel-im-spiegel/items/74a49773413c62721189 "1を1億回足して1億にならない場合 - Qiita")の転載です）

今回は軽めのネタで。

- [C > 浮動小数点型変数はループカウンタとして使用しない - Qiita](http://qiita.com/7of9/items/438a43bf53d60eab59e3)

まぁ浮動小数点数型の仕様を知れば当たり前の話なのだが，面白そうなので「1を1億回足す」ってのを [Go 言語]でも書いてみる。

```go
package main

import "fmt"

func main() {
	var d float32 = 0.0
	for i := 0; i < 100000000; i++ {
		d += 1.0
	}
	fmt.Println(d)
}
```

実行結果は予想通り

```text
$ go run loop1.go
1.6777216e+07
```

となる[^f32]。
念のため `float64` でも試してみよう。

[^f32]: `float32` は32ビットサイズの浮動小数点数型で，符号部1ビット，指数部8ビット，仮数部23ビット，という内訳になっている（仮数部は仮数の小数点以下を表す）。つまり有効桁数が24ビット（10進数で約7桁）しかない。したがって今回のような「1づつ加算する動作を繰り返す」処理では16,777,216（`=0xffffff+1`）以降は「情報落ち」が発生する。ちなみに `float64` は64ビットサイズで仮数部は52ビットあり，10進数にして約15桁の有効桁数になる。


```go
package main

import "fmt"

func main() {
	var d float64 = 0.0
	for i := 0; i < 100000000; i++ {
		d += 1.0
	}
	fmt.Println(d)
}
```

結果は

```text
$ go run loop2.go
1e+08
```

で，ちゃんと1億になる。
[Go 言語]では基本型のサイズが厳密に決まってるので（int, uint, uintptr は除く），浮動小数点数型の計算誤差についてもきちんと見積もれるはずである。

ちなみに

```go
package main

import "fmt"

func main() {
	for d := 0.0; d < 1.0; d += 0.1 {
		fmt.Println(d)
	}
}
```

とすると[^var]

[^var]: “`d := 0.0`” と記述した場合，変数 `d` は `float64` として宣言・初期化される。厳密には定数 “`0.0`” は，いったん「型付けなし」の浮動小数点数として評価された後，変数宣言時に `float64` に暗黙的に変換される。 [Go 言語]におけるこの定数の機能は何かと便利なので覚えておくとよいだろう。

```text
$ go run loop3.go
0
0.1
0.2
0.30000000000000004
0.4
0.5
0.6
0.7
0.7999999999999999
0.8999999999999999
0.9999999999999999
```

ってなことになる[^r] ので浮動小数点数型の変数をループカウンタにするのは止めましょうね。
約束だよ！

[^r]: このような結果になるのは `float32`/`float64` の浮動小数点数型の内部表現が2進数になっているため。たとえば 0.1 を2進数で表すと「0.000110011...」と循環しキリのいい値にならない。このため 0.1 を加算していくと「丸め誤差」が蓄積していくのである。

## ブックマーク

- [浮動小数点数型と誤差](http://www.cc.kyoto-su.ac.jp/~yamada/programming/float.html)
- [情報落ち、桁落ち、丸め誤差、打切り誤差の違い](http://tooljp.com/jyosho/docs/ketaochi-jyohoochi/ketaochi-jyohoochi.html)

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
