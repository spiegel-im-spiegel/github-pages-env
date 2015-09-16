+++
date = "2015-09-16T10:48:40+09:00"
description = "description"
draft = true
tags = ["golang", "algorithm", "math", "prime-number", "slice", "make"]
title = "素数探索アルゴリズムで遊ぶ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[Go 言語]は公式のドキュメントがとても充実していて（ただしほぼ英語だけど），私のような初学者に易しい環境といえる。

- [Documentation - The Go Programming Language](https://golang.org/doc/) : 言語仕様に関するドキュメントはこちら（[一部日本語化](http://golang-jp.org/doc/)されている）
- [Packages - The Go Programming Language](https://golang.org/pkg/) : 標準パッケージのドキュメントはこちら（[一部日本語化](http://golang-jp.org/pkg/)されている）

とはいえ，コードが実際にどのように機能するかは書いていないとわからない部分もある。
なので，今回からは実際にコードを書きながら言語の癖のようなものを調べていくことにする。
仕事に使うなら厳密な評価が必要だけど，今のところはそんな予定もないし，まずはテキトーで（笑）

早速，みんな大好き素数探索アルゴリズムで遊ぶ。

## 素数の定義

一応，素数（prime number）の定義を以下に示す。

> 1 と自分自身以外に正の約数を持たない 1 より大きい自然数

ここで自然数（natural number）は「ペアノの公理」に従う（0 が自然数に含まれるかどうかについては色々あるみたいだが，素数の定義には影響がないので，ここでは無視する）。

- [ペアノの公理 - Wikipedia](https://ja.wikipedia.org/wiki/%E3%83%9A%E3%82%A2%E3%83%8E%E3%81%AE%E5%85%AC%E7%90%86)
- [Swiftで自然数を作ってみた（ペアノの公理） - Qiita](http://qiita.com/taketo1024/items/2ab856d21bf9b9f30357)

ちなみに結城浩さんの『数学ガール／ゲーデルの不完全性定理』にペアノの公理について分かりやすく解説した章がある。
お勧め。


## 素数探索アルゴリズム（その1）{#alg1}

素数の定義を愚直にコードで表すなら以下のようになる。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	max := 100
	fmt.Printf("%v 以下の素数:", max)

	start := time.Now() //Start
	for n := 2; n <= max; n++ {
		flag := true
		for m := 2; m < n; m++ {
			if (n % m) == 0 { // n が m で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf(" %v", n)
		}
	}
	goal := time.Now()                     //Goal
	fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
```

実行結果はこんな感じ。

```
C:>go run prime01.go
100 以下の素数: 2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97
5.0002ms 経過
```

この値を覚えておいてね。
検算に使うから。

## 素数探索アルゴリズム（その2: エラトステネスの篩の変形）{#alg2}

もう少しだけ効率的に素数を探すアルゴリズムとして「[エラトステネスの篩]」と呼ばれる方法がある。
ただし「[エラトステネスの篩]」は決まった範囲を探索するものなので少々使いづらい。
そこで「[エラトステネスの篩]」で使われている以下の素数の特徴を最初のアルゴリズムに加えてみる。

1. 2 以上の全ての自然数はひとつ以上の素数の積で構成される（この素数の集合を素因数（prime factor）という）。したがってある数が素数か否かの判定は，その数より小さい素数のみで調べればよい
1. 更に，積の可換則（commutative property あるいは交換法則）により，自然数 $n$ が素数か否か判定する場合は $\sqrt{n}$ 以下の素数で調べればよい  
   （たとえば 35 の素因数は 5 と 7 だが， $5 \times 7 = 7 \times 5 = 35$ なので，直前の素数 31 まで回さずとも 3 および 5 ($\le \sqrt{35}$) まで調べれば判定できる）
1. 素数の定義から 2 が素数であることは自明なので（1 と 2 の間に自然数は存在しない）， 2 より大きい 2 の倍数（すなわち偶数）については判定しなくてもよく，対象となる自然数は 3 以上の奇数のみでいいことになる。

では，この特徴を加えたコードを書いてみよう。

```go
package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	primes := make([]int64, 1)     // 素数のリスト
	primes_f := make([]float64, 1) // 素数のリスト（浮動小数点へのキャスト）
	primes[0] = 2                  // 2 は素数
	primes_f[0] = 2.0              // 2 は素数（浮動小数点）
	var max int64 = 100

	start := time.Now() // Start
	var n int64 = 3
	for n = 3; n < max; n += 2 { // 3 から始まる奇数のみを探索
		flag := true
		f := float64(n)                    // 浮動小数点に cating
		rf := math.Sqrt(f)                 // n に対して √n をとる
		for i := 1; i < len(primes); i++ { // 2 より大きい既知の素数でチェックする
			if primes_f[i] > rf { // n に対して √n 以下の素数まで探索すればよい
				break
			} else if (n % primes[i]) == 0 { // n が既知の素数で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, n)     // 素数を追加
			primes_f = append(primes_f, f) // 素数を追加（浮動小数点）
		}
	}
	goal := time.Now() // Goal
	fmt.Printf("%v 以下の素数: %v\n", max, primes)
	fmt.Printf("%v 経過", goal.Sub(start)) // 経過時間を表示
}
```

実行結果はこんな感じ。

```
C:>go run prime02.go
100 以下の素数: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
0 経過
```

### slice と make() と append()

コード中

```go
primes := make([]int64, 1)     // 素数のリスト
primes_f := make([]float64, 1) // 素数のリスト（浮動小数点へのキャスト）
primes[0] = 2                  // 2 は素数
primes_f[0] = 2.0              // 2 は素数（浮動小数点）
```

の変数 `primes` および `primes_f` は slice と呼ばれる可変長の配列型である。
更に組み込み関数 `make()` は slice, map, channel のみ使用可能なメモリ割り当て関数である。
ちなみに slice, map, channel 以外は `new()` を使う。
slice, map, channel のみ特別なのは，これらの型は初期値と内部状態を持つためである。

slice に要素を追加する場合は `append()` 関数を使えばいいのだが，これが結構クセがある。
`append()` 関数では slice の容量（capacity）がいっぱいになると新たにメモリを確保してオリジナルの内容をコピーする。
つまりポインタが変わってしまうのだ。（メモリの割り当て方のパターンにも注目）

```go
package main

import "fmt"

func main() {
	a := make([]int, 0)                                         // 空の配列を用意
	fmt.Printf("Slice(%02d) : %p : %v (%v)\n", 0, a, a, cap(a)) // 配列の表示（初期状態）
	for num := 1; num <= 17; num++ {
		a = append(a, num)                                            //配列要素の追加
		fmt.Printf("Slice(%02d) : %p : %v (%v)\n", num, a, a, cap(a)) //配列の表示
	}
}
```

```
C:>go run slice.go
Slice(00) : 0x5cebb8 : [] (0)
Slice(01) : 0xc082002340 : [1] (1)
Slice(02) : 0xc082002380 : [1 2] (2)
Slice(03) : 0xc082006740 : [1 2 3] (4)
Slice(04) : 0xc082006740 : [1 2 3 4] (4)
Slice(05) : 0xc0820083c0 : [1 2 3 4 5] (8)
Slice(06) : 0xc0820083c0 : [1 2 3 4 5 6] (8)
Slice(07) : 0xc0820083c0 : [1 2 3 4 5 6 7] (8)
Slice(08) : 0xc0820083c0 : [1 2 3 4 5 6 7 8] (8)
Slice(09) : 0xc082050000 : [1 2 3 4 5 6 7 8 9] (16)
Slice(10) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10] (16)
Slice(11) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11] (16)
Slice(12) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12] (16)
Slice(13) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13] (16)
Slice(14) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14] (16)
Slice(15) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15] (16)
Slice(16) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16] (16)
Slice(17) : 0xc082056000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] (32)
```

容量が変化するごとにポインタ値も変化していることがお分かりだろうか。

ちなみに Go コンパイラは，返値を無視するコーディングに対してエラーを吐く。

```go
a = append(a, num)
```

の部分を

```go
append(a, num)
```

とすると

```
C:>go run prime02.go
# command-line-arguments
.\slice.go:9: append(a, num) evaluated but not used
```

とエラーになる。

また，容量はあらかじめ指定できる。

```go
a := make([]int, 0)
```

の部分を

```go
a := make([]int, 0, 32)
```

とすると

```
C:>go run slice.go
Slice(00) : 0xc082050000 : [] (32)
Slice(01) : 0xc082050000 : [1] (32)
Slice(02) : 0xc082050000 : [1 2] (32)
Slice(03) : 0xc082050000 : [1 2 3] (32)
Slice(04) : 0xc082050000 : [1 2 3 4] (32)
Slice(05) : 0xc082050000 : [1 2 3 4 5] (32)
Slice(06) : 0xc082050000 : [1 2 3 4 5 6] (32)
Slice(07) : 0xc082050000 : [1 2 3 4 5 6 7] (32)
Slice(08) : 0xc082050000 : [1 2 3 4 5 6 7 8] (32)
Slice(09) : 0xc082050000 : [1 2 3 4 5 6 7 8 9] (32)
Slice(10) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10] (32)
Slice(11) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11] (32)
Slice(12) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12] (32)
Slice(13) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13] (32)
Slice(14) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14] (32)
Slice(15) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15] (32)
Slice(16) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16] (32)
Slice(17) : 0xc082050000 : [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17] (32)
```

となる。
メモリ割り当ては意外に高コストの操作なので， slice を扱う場合はこの辺がチューニング・ポイントになるだろう。

## 100万個目の素数

上のコードを少し修正して $x$ 個目の素数を調べることにしよう。

```go
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	//コマンドライン引数の解析
	algno := flag.Int("alg", 0, "0: Basic algorithm , 1: Sieve of Eratosthenes")
	flag.Parse()
	args := flag.Args()
	if *algno < 0 || *algno > 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, os.ErrInvalid)
		return
	}
	max, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if max <= 0 {
		max = 1
	}

	//素数探索
	prime := int64(0)
	start := time.Now() // Start
	switch *algno {
	case 1:
		prime = LastPrimeE(max)
	default:
		prime = LastPrimeB(max)
	}
	goal := time.Now()                       // Goal
	fmt.Printf("%v 個目の素数: %v\n", max, prime) // max 個目の素数
	fmt.Printf("%v 経過\n", goal.Sub(start))   // 経過時間を表示
}

func LastPrimeB(max int64) int64 {
	count := int64(0)

	for n := int64(2); ; n++ {
		flag := true
		for m := int64(2); m < n; m++ {
			if (n % m) == 0 { // n が m で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
		}
	}
}

func LastPrimeE(max int64) int64 {
	if max <= 1 {
		return 2
	}
	primes := make([]int64, 1, max)     // 素数のリスト
	primes_f := make([]float64, 1, max) // 素数のリスト（浮動小数点へのキャスト）
	primes[0] = 2                       // 2 は素数
	primes_f[0] = 2.0                   // 2 は素数（浮動小数点）

	count := int64(1)
	for n := int64(3); ; n += 2 { // 3 から始まる奇数のみを探索
		flag := true
		f := float64(n)                    // 浮動小数点に cating
		rf := math.Sqrt(f)                 // n に対して √n をとる
		for i := 1; i < len(primes); i++ { // 2 より大きい既知の素数でチェックする
			if primes_f[i] > rf { // n に対して √n 以下の素数まで探索すればよい
				break
			} else if (n % primes[i]) == 0 { // n が既知の素数で割り切れる → 素数ではない
				flag = false
				break
			}
		}
		if flag {
			count++
			if count >= max {
				return n
			}
			primes = append(primes, n)     // 素数を追加
			primes_f = append(primes_f, f) // 素数を追加（浮動小数点）
		}
	}
}
```

今後のためにコマンドライン解析の部分と実際の素数探索アルゴリズムを分けている。
まず検算。
25個目の素数が97なら OK。

```
C:>go run prime03.go -alg=0 25
25 個目の素数: 97
0 経過

C:>go run prime03.go -alg=1 25
25 個目の素数: 97
0 経過
```

では実際に動かしてみよう。
まずは「[その1]({{< ref "#alg1" >}})」のアルゴリズムから。

```
C:>go run prime03.go -alg=0 100
100 個目の素数: 541
0 経過

C:>go run prime03.go -alg=0 10000
10000 個目の素数: 104729
4.4072521s 経過
```

100万個目の素数は有意の時間で見つかりませんでした orz

次に「[その2]({{< ref "#alg2" >}})」のアルゴリズムで。

```
C:>go run prime03.go --alg=1 100
100 個目の素数: 541
0 経過

C:>go run prime03.go --alg=1 10000
10000 個目の素数: 104729
7.0004ms 経過

C:>go run prime03.go --alg=1 1000000

1000000 個目の素数: 15485863
4.9042805s 経過

C:>go run prime03.go --alg=1 10000000
10000000 個目の素数: 179424673
2m13.8686568s 経過
```

というわけで，100万個目の素数探索に5秒弱，1000万個目の素数探索に2分ちょっとかかってしまった。
まぁ，でも，こんなもんか。

## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[エラトステネスの篩]: https://ja.wikipedia.org/wiki/%E3%82%A8%E3%83%A9%E3%83%88%E3%82%B9%E3%83%86%E3%83%8D%E3%82%B9%E3%81%AE%E7%AF%A9 "エラトステネスの篩 - Wikipedia"
