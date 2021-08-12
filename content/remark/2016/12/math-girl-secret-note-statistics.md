+++
title = "『数学ガールの秘密ノート／やさしい統計』で遊ぶ"
date = "2016-12-11T15:59:26+09:00"
description = "今回はテトラちゃん回かなぁ。1,2章のユーリちゃんとのやり取りも面白いけど，新しい用語が次々登場する状況で，言葉や名前に敏感なテトラちゃんがアワアワする感じがよかった。"
tags = [
  "book",
  "math",
  "programming",
  "golang",
]

[scripts]
  mathjax = false
  mermaidjs = false
+++

（今回はネタバレをガッツリ含むので見たくない方は静かにページを閉じてください）

本当は「この世界の片隅に」を観に行こうかと思っていたのだが，さすが地元広島は大人気のようで満席になっていた。
こりゃあ劇場で観るのは無理そうだな。

{{< fig-img src="https://photo.baldanders.info/flickr/image/31524625546_m.jpg" title="残念。こりゃあ劇場で見るのは無理そうだな。シネコンで見る気はないし。" link="https://photo.baldanders.info/flickr/31524625546/" >}}

ということで映画は諦めて，夕方の待ち合わせ時間まで喫茶店で積ん読状態だった『[数学ガールの秘密ノート／やさしい統計]』を一気読みすることにした。

ちなみに『[数学ガールの秘密ノート／やさしい統計]』の詳しい解説は以下のページが参考になる。

- [数学好きから統計好きに――『数学ガールの秘密ノート／やさしい統計』｜Colorless Green Ideas](http://id.fnshr.info/2016/11/05/secret-notebook-statistics/)

この記事によると最近の中高生は統計についてきちんと習うらしい（でも選択科目？）。
私は高校は理数科だったので一通り習ったはずなのだが，あまり憶えてなかったりする。
統計について強く意識するようになったのは大学で誤差論を習ってから。
だから純粋に数学って感じじゃなくて，あくまでもデータを解析するための道具・手段として捉えていた。
『[数学ガールの秘密ノート／やさしい統計]』を通して統計の初歩について数学体系として学ぶのは多分初めての体験で（高校で習ったのはあくまでも受験用だったからね）とても面白かった。

今回はテトラちゃん回かなぁ。
1,2章のユーリちゃんとのやり取りも面白いけど，新しい用語が次々登場する状況で，言葉や名前に敏感なテトラちゃんがアワアワする感じがよかった。

今回の『[数学ガールの秘密ノート／やさしい統計]』がいつもと毛色が違うなと感じたのは，次々と用語が登場して，その定義と意味についてユーリちゃんやテトラちゃんに延々と説明していく，という流れになっていたからだろうか。
実は最近新しいプロジェクトで要求分析をやっているところなので，それと脳内でシンクロする感じが楽しかった。

要求分析で真っ先にやるのは「用語集」を作ることだ。
「用語集」を作る目的は2つある。
ひとつは顧客と「言葉」を合わせることで，もうひとつは「言葉」を厳密に定義することだ。
これを最初にやっておかないと顧客との間で齟齬が生じてしまう場合がある。
『[数学ガールの秘密ノート／やさしい統計]』でやっていることは用語集の作成プロセスに似ている。

さて，作中で登場した村木先生の《カード》

{{< fig-quote title="『数学ガールの秘密ノート／やさしい統計』第4章" link="https://www.amazon.co.jp/exec/obidos/ASIN/B01MSJMKMW/baldandersinf-22/" >}}
<q>コインを10回投げたとき、表は何回出るだろう。</q>
{{< /fig-quote >}}

プログラマならこれを擬似乱数を使って試してみたいって思うよね。
このサイトの別セクションで，以前「[モンテカルロ法を使って円周率を求める]({{< ref "/golang/estimate-of-pi.md" >}})」というのをやったので，これを応用してコードを組んでみることにしよう。

できあがりはこちら。

- [spiegel-im-spiegel/cointoss: コインを10回投げたとき、表は何回出るだろう。](https://github.com/spiegel-im-spiegel/cointoss)

パッケージ構成は概ねこんな感じ。

{{< fig-img src="./ddd-cointoss.svg" title="パッケージ構成" link="./ddd-cointoss.svg" width="640" >}}

`gen` パッケージで擬似乱数を使ってコイントスを行う。
こんな感じのコード。

```go
/**
 * These codes are licensed under CC0.
 * https://creativecommons.org/publicdomain/zero/1.0/
 */

package gen

import "math/rand"

//New returns generator of random number [0,2)
func New(s rand.Source, ct int64) <-chan int {
    ch := make(chan int)
    r := rand.New(s)

    go func(ct int64) {
        for i := int64(0); i < ct; i++ {
            ch <- r.Intn(2)
        }
        close(ch)
    }(ct)

    return ch
}
```

コイントスを行った結果を入れる [channel] を生成し，実際のコイントスは [goroutine] の中で擬似乱数生成器を使い指定した回数だけ行っている。
1が出れば表である。
では実際にコインを10回投げてみよう。

```text
$ go run main.go toss -t 10
0
0
1
0
0
0
0
0
1
1
```

ここでは10回中表になったのは3回。
さらに「コインを10回投げる」行為を10回行って統計値を取ってみる。

```text
$ go run main.go repeat -t 10 -c 10
6
6
5
1
9
5
5
5
6
3
minimum value: 1
maximum value: 9
average value: 5.10000
standard deviation: 1.97231
```

ここでは表が出た回数の最小値が1，最大値が9，平均値が5.1，標準偏差が約1.97となった。
10回ではよくわからないので次はどどーんと1万回やってみる。

```text
$ go run main.go repeat -t 10 -c 10000 > cointoss.dat
minimum value: 0
maximum value: 10
average value: 4.99690
standard deviation: 1.57477
```

結果データは [`cointoss.dat`](./cointoss.dat) ファイルに保存している。
ここでは表が出た回数の最小値が0，最大値が10，平均値が約5.0，標準偏差が約1.57となった。
では [`cointoss.dat`](./cointoss.dat) ファイルを [gnuplot] に食わせてヒストグラムにしてみる（階級幅は1）[^h]。

[^h]: 「[gnuplot でヒストグラム（頻度分布図）を描画する - Qiita](http://qiita.com/iwiwi/items/4c7635d4c84bc785e47a)」を参考にした。

```text
gnuplot> unset key
gnuplot> set xrange [-1:11]  
gnuplot> filter(x,y)=int(x/y)*y                                                              
gnuplot> plot "cointoss.dat"  u (filter($1,1)):(1) smooth frequency with boxes lc rgb "black"
```

結果はこんな感じ。

{{< fig-img src="./cointoss-hist.png" title="ヒストグラム" link="./cointoss-hist.png" width="750" >}}

おおっ。
きれいに正規分布っぽくなっている。
さて，『[数学ガールの秘密ノート／やさしい統計]』4章で数学的に求めた値とどのくらい違うかな（笑） ちなみに $\sqrt{2.5} = 1.5811...$ である。

あー，遊んだ遊んだ。

[数学ガールの秘密ノート／やさしい統計]: https://www.amazon.co.jp/exec/obidos/ASIN/B01MSJMKMW/baldandersinf-22/ "Amazon.co.jp: 数学ガールの秘密ノート／やさしい統計 電子書籍: 結城 浩: Kindleストア"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[channel]: http://golang.org/ref/spec#Channel_types
[goroutine]: http://golang.org/ref/spec#Go_statements
[gnuplot]: http://www.gnuplot.info/ "gnuplot homepage"

{{% review-paapi "B01MSJMKMW" %}} <!-- 数学ガールの秘密ノート／やさしい統計 -->

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
