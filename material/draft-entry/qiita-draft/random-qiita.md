# 「ズンドコキヨシ」と擬似乱数

最近はやりものらしい「ズンドコキヨシ[^zk]」を [Go 言語]で実装することを考える。既に Qiita でも様々な実装があるが，今回は素朴に

[^zk]: 「ズンドコキヨシ」「キヨシチェック」「ズンドコチェック」などと呼ばれているらしいが，今回は「ズンドコキヨシ」で統一する。

```go:zundoko.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	zun     = "ズン"
	doko    = "ドコ"
	kiyoshi = "キ・ヨ・シ！"
)

func generate() chan string {
	ch := make(chan string)
	go func() {
		var zundoko = [2]string{zun, doko}
		rand.Seed(time.Now().UnixNano())
		for {
			ch <- zundoko[rand.Intn(2)]
		}
	}()
	return ch
}

func main() {
	zundoko := generate()
	zcount := 0
	for {
		zd := <-zundoko
		fmt.Print(zd)
		if zd == zun {
			zcount++
		} else if zcount >= 4 {
			break
		} else {
			zcount = 0
		}
	}
	fmt.Print(kiyoshi)
}
```

としてみた。実行例は以下のとおり。

```
$ go run zundoko.go
ドコズンドコズンズンドコズンズンドコドコズンドコドコドコズンズンズンズンドコキ・ヨ・シ！
```

この処理の内容については本家ブログで解説している。

- [「ズンドコチェック」なるものが流行っているらしい — プログラミング言語 Go | text.Baldanders.info](https://text.baldanders.info/golang/zundoko-choir/)

今回は「ズン」「ドコ」の配列を生成している部分 `generate()` のうち擬似乱数について紹介しておく。

## 線形合同法（Linear Congruential Generators）

[`math/rand`] パッケージで既定で使われている擬似乱数アルゴリズムはちょっと特殊な線形合同法（Linear Congruential Generators; LCGs）らしい。

- [[Announce] A rand package for high quality 64bit random numbers (possibly go2) - Google グループ](https://groups.google.com/forum/#!topic/golang-nuts/RZ1G3_cxMcM)

線形合同法は実装レベルで優劣が分かれる実装なのだが，そもそも線形合同法の問題はある乱数がその前の乱数の影響を強く受けてしまう点にある。今回の「ズンドコキヨシ」に引き寄せて言うなら，「ズン」が4回続いた後に「ドコ」が来る確率が 1/2 にならないのである[^ra]。

[^ra]: 線形合同法には本当にいろいろな実装があり，古い UNIX 処理系では生成した乱数の最下位ビットが 1 と 0 交互に出現するらしい。この場合，永遠に「ズンドコキヨシ」が終わらないことになる。線形合同法の改良版（[`math/rand`] もそのひとつ）では過去に指摘された問題は緩和される傾向にあるがアルゴリズム上の欠陥はどうしようもない。

## メルセンヌ・ツイスタ（Mersenne Twister）

線形合同法の欠点を克服すべく様々な擬似乱数アルゴリズムが考えだされた。そうしたアルゴリズムでかなり性能が良いと言われているアルゴリズムのひとつが松本眞，西村拓士両氏によって発表されたメルセンヌ・ツイスタ（Mersenne Twister; MT）である[^mt]。

[^mt]: 現在は改良版である [SIMD-oriented Fast Mersenne Twister (SFMT)](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/SFMT/index-jp.html) が公開されている。またソースコードが BSD ライセンスで提供されている。

- [Mersenne Twister: A random number generator (since 1997/10)](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html)

[Go 言語]の標準パッケージはメルセンヌ・ツイスタをサポートしていないが， [`seehuhn/mt19937`] パッケージでメルセンヌ・ツイスタが使える[^sm]。

[^sm]: メルセンヌ・ツイスタのオリジナルコードは BSD または MIT のデュアル・ライセンスで提供されているが， [`seehuhn/mt19937`] パッケージのライセンスは GPLv3 なので取り扱いに注意が必要である。

実は [`math/rand`] パッケージは以下の [interface] を持つ任意の擬似乱数生成パッケージを使うことができる。

```go
// A Source represents a source of uniformly-distributed
// pseudo-random int64 values in the range [0, 1<<63).
type Source interface {
	Int63() int64
	Seed(seed int64)
}
```

したがって，最初に示した `generate()` は以下のように書き換えられる。

```go
import (
	"math/rand"
	"time"

	"github.com/seehuhn/mt19937"
)

func generate() chan string {
	ch := make(chan string)
	go func() {
		var zundoko = [2]string{zun, doko}
		r := rand.New(mt19937.New())
		r.Seed(time.Now().UnixNano())
		for {
			ch <- zundoko[r.Intn(2)]
		}
	}()
	return ch
}
```

メルセンヌ・ツイスタは科学シミュレーション（モンテカルロ法など）でよく使われるほか，ゲーム内の乱数生成器としても使われているようである。

## もうひとつの rand パッケージ

線形合同法にしろメルセンヌ・ツイスタにしろエントロピー源として与えられる seed が同じなら同じ値になってしまう。暗号用に乱数を生成するためにはエントロピー源の与え方を工夫する必要がある。また得られた乱数をそのまま使うのもよくないと言われている。

[Go 言語]では暗号用に [`crypto/math`] パッケージを用意している。 [`crypto/math`] パッケージは [`math/rand`] パッケージとは互換性がない[^cr]。

[^cr]: [`crypto/math`] は乱数生成用のエントロピー源にハードウェア・デバイスを使用する。 UNIX 系のプラットフォームでは `/dev/urandom` を参照する。 Windows プラットフォームでは [`CryptGenRandom`](https://msdn.microsoft.com/ja-jp/library/windows/desktop/aa379942(v=vs.85).aspx "CryptGenRandom function (Windows)") を使うようだが，この API の内部実装は公開されていない。やれやれ。

## ブックマーク

- [良い乱数・悪い乱数](http://www001.upp.so-net.ne.jp/isaku/rand.html)
    - [悪い乱数のリスクを視覚的に明らかにする](http://www001.upp.so-net.ne.jp/isaku/rand3.html)
- [Big Sky :: Mersenne Twister のライセンスが変更された](http://mattn.kaoriya.net/software/20130409093112.htm)
- [MSC30-C. 疑似乱数の生成に rand() 関数を使用しない](https://www.jpcert.or.jp/sc-rules/c-msc30-c.html)
- [RFC 4086 - Randomness Requirements for Security](http://tools.ietf.org/html/rfc4086) （[IPA による日本語訳](https://www.ipa.go.jp/security/rfc/RFC4086JA.html)）
- [サイコロを振るのは簡単である、しかし、ゲームでサイコロを実装するにはサイコロを知らなければならない - Qiita](http://qiita.com/isonami/items/1cc278cbf2093d2d6abd)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math/rand`]: https://golang.org/pkg/math/rand/ "rand - The Go Programming Language"
[`crypto/math`]: https://golang.org/pkg/crypto/rand/ "rand - The Go Programming Language"
[`seehuhn/mt19937`]: https://github.com/seehuhn/mt19937 "seehuhn/mt19937: An implementation of Takuji Nishimura's and Makoto Matsumoto's Mersenne Twister pseudo random number generator in Go."
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"

## 脚注
