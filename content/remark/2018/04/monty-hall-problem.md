+++
title = "モンティ・ホール問題で遊ぶ"
date = "2018-04-20T10:31:29+09:00"
update = "2018-04-21T03:44:59+09:00"
description = "「モンティ・ホール問題」は確かに直感に反するが，こうやって具体的なコードで記述していくと「何故そうなるのか」が何となく分かってくる。"
image = "https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Monty_open_door_chances.svg/590px-Monty_open_door_chances.svg.png"
tags = [ "math", "games", "golang", "programming", "floating-point" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

有名な数学問題なので知ってる人も多いと思うけど一応解説すると，「モンティ・ホール問題」というのは Monty Hall という人が司会をつとめるアメリカの番組 “[Let's make a deal](https://en.wikipedia.org/wiki/Let%27s_make_a_deal)” 内で行われたゲームに関する問題である。

簡単にルールを紹介する。

1. プレイヤーの前に閉まった3つのドアがある。1つのドアの後ろには景品の新車が，2つのドアの後ろには（はずれを意味する）ヤギがいる
2. プレイヤーが1つのドアを選択
3. 司会者が残りのドアのうちヤギがいるドアを開けてヤギ（＝はずれ）を見せる
4. プレイヤーは最初に選んだドアを変更してもよい

このときプレイヤーは最初に選択したドアを変更したほうが得（つまり当たる確率が高い）かどうか，というのが「モンティ・ホール問題」のあらましである。

この問題について「ドアを変更したほうが当たる確率が高い」と示した女性がいて，これに対して猛烈な反発が起こり大論争になったらしい（女性蔑視な発言もあったそうで，今なら確実にハラスメント案件だよな`w`）。

ポイントは3番目の「はずれのドアを開ける」部分で（司会者はどのドアが当たりかあらかじめ知っている），残りの2つのドアのどちらかが当たりなのだからというので多くの人が直感的に「当たる確率は半々」だと思った，博士号保持者も含めて。

この問題は事後確率（posterior probability）と呼ばれるものの一種である[^b1]。
最初の選択で当たる確率は $1 / 3$ なので後半に変更しなければ $1 / 3$ のままだが，「はずれのドアを開ける」ことにより選択肢が既に選択したドアともうひとつのドアの2つになるため，ドアを変更した場合の当たる確率が $2 / 3$ （つまり半々より確率値が大きい）になるのである。

[^b1]: 事後確率についての説明は割愛する。ベイズ主義とか頻度主義とかちゃんと説明しようとするととてつもなく面倒臭いので。ただし事前確率とか事後確率とかいったものは統計学（ベイズ統計学）ではよく出てくる話なので，興味のある方は自力でどうぞ。

{{< fig-img src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9e/Monty_open_door_chances.svg/590px-Monty_open_door_chances.svg.png" title="Monty open door chances.svg - Wikimedia Commons" link="https://commons.wikimedia.org/wiki/File:Monty_open_door_chances.svg" width="590" >}}

「モンティ・ホール問題」は直感と論理との乖離を示す好例として挙げられることが多く，また簡単な確率シミュレーションとしてプログラミングの演習問題に使われることもあるようだ。

というわけで「モンティ・ホール問題」をシミュレートするコードを [Go 言語]で組んでみよう（[Go 言語]で書くことに特に意味はない。皆さんはお好きな言語でどうぞ）。

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func selectDoor(limit, max int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < max; i++ {
            ch <- rand.Intn(limit)
        }
    }()
    return ch
}

func challenges(shuldChange bool, max int) int {
    doors := []bool{true, false, false}
    ch := selectDoor(len(doors), max)
    count := 0
    for n := range ch {
        result := doors[n]
        if shuldChange {
            result = !result
        }
        if result {
            count++
        }
    }
    return count
}

func main() {
    max := 10000
    fmt.Println("nochange:", float64(challenges(false, max))/float64(max))
    fmt.Println("  change:", float64(challenges(true, max))/float64(max))
}
```

[Go 言語]的に面白いトピックはないので詳細は省くが，ざっくり説明すると `selectDoor()` 関数は（疑似乱数を使って）ドアを選択する関数で選択したドアを指定回数だけ channel で返す。
`challenges()` 関数は「モンティ・ホール問題」を指定回数だけ試行するもので，当たりの回数を返している。
ドアを変更するか否かは引数 `shuldChange` で指定する。
変数 `doors` がプレイヤーが選択するドアで，当たりが `true` に相当する。
for-range 構文の中身が実際に「モンティ・ホール問題」を試行している部分で

```go
for n := range ch {
    result := doors[n]
    if shuldChange {
        result = !result
    }
    if result {
        count++
    }
}
```

ドアを変更しない場合はそのまま，変更する場合は真偽が反転する。
これは「はずれのドアを開ける」と当たりとはずれのドアが **必ず** ひとつづつになるため，最初に選択したドアとそうでないドアで真偽が反対になるからである。

では実際にコードを実行してみる。

```text
$ go run monty-hall-problem.go
nochange: 0.3464
  change: 0.6675
```

というわけで，それっぽい値になった。
めでたし。

「モンティ・ホール問題」は確かに直感に反するが[^p1]，こうやって具体的なコードで記述していくと「何故そうなるのか」が何となく分かってくる。
どうせプログラミング教育をやるのならこういう感じでやっていただきたいものである（笑）

[^p1]: 「モンティ・ホール問題」はジレンマやパラドックスの一種と見なされることがあるらしい。実際には論理的な矛盾はないので疑似パラドックスといったところだろうけど。

## おまけ：N個のドアでモンティ・ホール問題

応用として $n$ 個のドアで「モンティ・ホール問題」を検証してみよう。
ルールはこんな感じ。

1. プレイヤーの前に閉まった $n$ 個のドアがある（$n \ge 3$）。当たりのドアは $1$ つ。残り $n-1$ 個ははずれ
2. プレイヤーが $1$ つのドアを選択
3. 司会者が残りのドアのうち $n-2$ 個のはずれドアを開ける
4. プレイヤーは最初に選んだドアを変更してもよい

まず先ほどのコードのうち `challenges()` 関数と `main()` 関数を以下のように変更する。

```go
func challenges(shuldChange bool, dct, max int) int {
    doors := make([]bool, dct, dct) //initialized by false
    doors[0] = true
    ch := selectDoor(len(doors), max)
    count := 0
    for n := range ch {
        result := doors[n]
        if shuldChange {
            result = !result
        }
        if result {
            count++
        }
    }
    return count
}

func main() {
    dct := 3
    max := 10000
    fmt.Println("nochange:", float64(challenges(false, dct, max))/float64(max))
    fmt.Println("  change:", float64(challenges(true, dct, max))/float64(max))
}
```

これで最初のコードとまったく同じ機能になる。
ここで `main()` 関数内の `dct` を `100` にして実行してみる。

```text
$ go run monty-hall-problem.go
nochange: 0.0114
  change: 0.9785
```

何故こういう値になるか考えてみませう。
簡単だよね。

## おまけ2：有理数表現

[Go 言語]では [math/big] パッケージを使った有理数表現 [`big`]`.Rat` が使える。
これを使って `main()` 関数を書き直すと

```go
func main() {
    dct := 3
    max := 10000
    fmt.Println("nochange:", big.NewRat(int64(challenges(false, dct, max)), int64(max)).FloatString(4))
    fmt.Println("  change:", big.NewRat(int64(challenges(true, dct, max)), int64(max)).FloatString(4))
}
```

てな感じになる。
[math/big] パッケージを使うと，計算コストは高くなるが，[浮動小数点数型特有の計算誤差]({{< relref "golang/loop-counter.md" >}} "1を1億回足して1億にならない場合")を緩和できるメリットがある。
まぁ，今回はあんまり関係ないけどね。

おおっ。
なんか [Go 言語] っぽい（笑）

[Go 言語]: https://golang.org/ "The Go Programming Language"
[math/big]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"
[`big`]: https://golang.org/pkg/math/big/ "big - The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
