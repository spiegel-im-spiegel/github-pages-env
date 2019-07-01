+++
title = "Go 言語における Singleton Pattern"
date =  "2017-10-24T14:43:01+09:00"
update = "2018-12-07T10:08:03+09:00"
description = "はっきり言って「Singleton なめんな！」ですよ。"
tags        = [ "golang", "programming", "singleton", "init", "sync", "goroutine" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は小ネタでお送りします。

いや，ネットでね，[Go 言語]での Singleton 実装をこんな感じに書く人をやたら見かけるのだが

```go
//Hello class
type Hello struct{}

var instance *Hello

//GetInstance returns singleton instance
func GetInstance() *Hello {
    if instance == nil {
        instance = &Hello{}
    }
    return instance
}
```

はっきり言って **Singleton なめんな！** ですよ。

そうそう。
プログラマで Singleton Pattern を知らない人はいないと思うけど，一応解説しておくと， Singleton Pattern というのは，あるクラスに対してプログラム全体でインスタンスがひとつだけ生成されるよう制限するプログラミング・パターンである。
たとえば，単一セッションでシリアルに通信を行う entity class なんかはインスタンスがぼこぼこできて各々勝手に処理をされると困るわけで[^c1]， singleton インスタンスの内部で同期をとっていく必要があるわけ。

[^c1]: HTTP なんかはステートレスでパラレルにやり取りできるため Singleton Pattern は不要。ちなみに [goroutine] 間の通信は [channel] を使えっちう話です，はい。

という説明からも分かると思うけど「**スレッドセーフでない singleton 実装に存在意義はない**」のである。
ちなみに結城浩さんの[デザパタ本]にあるサンプルコードも以下のようになっている（こっちは Java での記述だけど[^java1]）。

[^java1]: このコードのパターンは Singleton Pattern を説明するにはよく出来ているしちゃんと動く（ココ重要）が，同期コストが高いため，実際にはあまり使われない。 Java における Singleton Pattern には様々な実装例があるので探してみるといいだろう。ちなみに Java 使いではなくとも結城浩さんの[デザパタ本]は買って読んでおくことを強くお勧めする。 Java 使いの方から見ると古いバージョンで書かれたコードなのが難点だが，紙の本で買うとサンプルコード入りのディスクが付いてくるので若干お得？

```java
import java.util.Date;

public class MySystem {
    private static MySystem instance = null;
    private Date date = new Date();
    private MySystem() {
    }
    public Date getDate() {
        return date;
    }
    public static synchronized MySystem getInstance() {
        if (instance == null) {
            instance = new MySystem();
        }
        return instance;
    }
}
```

`static` で `synchronized` なのがポイントね。
つまり，実際にインスタンスを生成する処理では何らかの手段でスレッドセーフであることが保証されてないといけない[^st1]。
最初の [Go 言語]のパターンが何故ダメなのかは実際に動くコードを書いてみれば分かる。

[^st1]: 他にもインスタンスのコピー（コピーコンストラクタ等）を暗黙的に許容する言語ではコピーを無効にする措置が必要，とかある。そういう意味じゃ今回私が書いたコードも不完全で，実際には singleton インスタンスを隠蔽するためのラッパークラスが必要になる。ビジネス・ロジックも含めると，実は Singleton の実装ってそう甘くないのよねー

```go
package main

import (
    "fmt"
    "time"
)

//Hello class
type Hello struct{}

func (h *Hello) String() string {
    return "Hello"
}

var instance *Hello

//GetInstance returns singleton instance
func GetInstance() *Hello {
    if instance == nil {
        fmt.Println("new instance")
        time.Sleep(1 * time.Second) //delay 1sec
        instance = &Hello{}
    }
    return instance
}

func main() {
    ch := make(chan interface{})
    go run(ch, "Alice")
    go run(ch, "Bob")
    go run(ch, "Chris")
    <-ch
    <-ch
    <-ch
}

func run(ch chan<- interface{}, person string) {
    hello := GetInstance()
    ch <- hello //blocking
    fmt.Println(hello, person)
}
```

簡単に説明すると，まず `GetInstance()` 関数内部で初期化処理時間を演出するために１秒間の delay を発生させている。
`run()` 関数内で [channel] `ch` にインスタンスを食わせているのはブロッキングのため。
別に何を食わせてもいいのだが，手近に `GetInstance()` 関数で取得したインスタンスがあるので，それを食わせている。
`main()` 関数では `run()` 関数を [goroutine] で3連続起動したあと， `<-ch` でブロックを解除している。

[実行結果](https://play.golang.org/p/cL-RMmS2ev)は以下の通り。

```text
new instance
new instance
new instance
Hello Alice
Hello Chris
Hello Bob
```

並行処理下の3つの `run()` 関数に対してインスタンスが3つ生成されてしまっているのが分かると思う。

ではどう書けばいいのか。
一番簡単なのは [var] 宣言時に初期化してしまうことである。

```go
var instance = &Hello{}

//GetInstance returns singleton instance
func GetInstance() *Hello {
    return instance
}
```

次に簡単なのは [`init()`] 関数を使うことである。

```go
var instance *Hello

func init() {
    //create instance and initialize
    instance = &Hello{}
}
```

[`init()`] 関数は少し特殊な関数で，`main()` 関数がキックされる前，パッケージ内の [var] 宣言時の初期化の後に呼ばれる。
ひとつのパッケージ内またはひとつのファイル内にいくつも [`init()`] 関数を設置できるのが特徴なのだが，どういう順番に起動するかは言語仕様として明記されていないため[^init1]，パッケージ内の複数の [`init()`] 関数同士が依存また干渉するような書き方は避けるべきだろう。

[^init1]: どうもソースファイルのファイル名が影響するらしい。つまりファイル名を工夫すれば [`init()`] 関数の呼び出し順を制御できる，という噂。

`main()` 関数がキックされるまではメイン以外の [goroutine] は（生成は可能だが）起動されないことが保証されているため，記述がスレッドセーフか否か気にする必要はない。
言い方を変えると，何らかの同期を伴う初期化処理の場合はこの方法では記述できないことになる。

「どうしても `GetInstance()` 関数内で同期をとりたいんじゃ」という場合は... たとえば [`sync`] パッケージを使うとかだろうか。
Singleton Pattern におあつらえ向きの [`sync`].`Once` というのがある。
最初に挙げた例を流用するならこんな感じだろうか。

{{< highlight go "hl_lines=5 9-11" >}}
//Hello class
type Hello struct{}

var instance * Hello
var once sync.Once

//GetInstance returns singleton instance
func GetInstance() * Hello {
    once.Do(func() {
        instance = &Hello{}
    })
    return instance
}
{{< /highlight  >}}

以下のコードで実際に動かして検証してみよう。

{{< highlight go "hl_lines=17 21-25" >}}
package main

import (
    "fmt"
    "sync"
    "time"
)

//Hello class
type Hello struct{}

func (h * Hello) String() string {
    return "Hello"
}

var instance * Hello
var once sync.Once

//GetInstance returns singleton instance
func GetInstance() * Hello {
    once.Do(func() {
        fmt.Println("new instance")
        time.Sleep(1 * time.Second) //delay 1sec
        instance = &Hello{}
    })
    return instance
}

func main() {
    ch := make(chan interface{})
    go run(ch, "Alice")
    go run(ch, "Bob")
    go run(ch, "Chris")
    <-ch
    <-ch
    <-ch
}

func run(ch chan<- interface{}, person string) {
    hello := GetInstance()
    ch <- hello //blocking
    fmt.Println(hello, person)
}
{{< /highlight  >}}

[実行結果](https://play.golang.org/p/8ODOeffoyF)は以下の通り。

```text
new instance
Hello Chris
Hello Alice
Hello Bob
```

ちゃんと singleton として動作していることが分かる。

## ブックマーク

- [GASCII.jp：Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/235/1235262/)
    - [ASCII.jp：Go言語と並列処理（2）｜Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/480/1480872/)
- [Goでスレッド（goroutine）セーフなプログラムを書くために必ず注意しなければいけない点 - Qiita](https://qiita.com/ruiu/items/54f0dbdec0d48082a5b1)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[デザパタ本]: https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1BS/baldandersinf-22/ "増補改訂版 Java言語で学ぶデザインパターン入門 マルチスレッド編 | 結城 浩 | コンピュータ・IT | Kindleストア | Amazon"
[channel]: http://golang.org/ref/spec#Channel_types "The Go Programming Language Specification - The Go Programming Language"
[goroutine]: http://golang.org/ref/spec#Go_statements "The Go Programming Language Specification - The Go Programming Language"
[var]: https://golang.org/ref/spec#Variables "The Go Programming Language Specification - The Go Programming Language"
[defer]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[`init()`]: https://golang.org/doc/effective_go.html#init "Effective Go - The Go Programming Language"
[`sync`]: https://golang.org/pkg/sync/ "sync - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%AA%E3%83%95%E3%82%A1%E3%82%AF%E3%82%BF%E3%83%AA%E3%83%B3%E3%82%B0%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1EU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1EU"><img src="https://images-fe.ssl-images-amazon.com/images/I/41GPVATQiZL._SL160_.jpg" width="125" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%AA%E3%83%95%E3%82%A1%E3%82%AF%E3%82%BF%E3%83%AA%E3%83%B3%E3%82%B0%E5%85%A5%E9%96%80-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1EU?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1EU">Java言語で学ぶリファクタリング入門</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2007-01-26 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1EU</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">結城浩さんによる「リファクタリング本」の Kindle 版。意外にも Java 以外でも応用できる優れもの。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-11">2018-12-11</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E5%A2%97%E8%A3%9C%E6%94%B9%E8%A8%82%E7%89%88-Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E3%83%9E%E3%83%AB%E3%83%81%E3%82%B9%E3%83%AC%E3%83%83%E3%83%89%E7%B7%A8-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1BS?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1BS"><img src="https://images-fe.ssl-images-amazon.com/images/I/41uoMp5etSL._SL160_.jpg" width="126" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E5%A2%97%E8%A3%9C%E6%94%B9%E8%A8%82%E7%89%88-Java%E8%A8%80%E8%AA%9E%E3%81%A7%E5%AD%A6%E3%81%B6%E3%83%87%E3%82%B6%E3%82%A4%E3%83%B3%E3%83%91%E3%82%BF%E3%83%BC%E3%83%B3%E5%85%A5%E9%96%80-%E3%83%9E%E3%83%AB%E3%83%81%E3%82%B9%E3%83%AC%E3%83%83%E3%83%89%E7%B7%A8-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1BS?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1BS">増補改訂版 Java言語で学ぶデザインパターン入門 マルチスレッド編</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2004-06-18 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1BS</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">結城浩さんによる通称「デザパタ本」の Kindle 版。マルチスレッド下のデザインパターン集。こちらも Java 以外でも応用できる優れもの。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-07">2018-12-07</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
