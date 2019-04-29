+++
title = "「《命》に関わる確率」を疑似乱数を使って解いてみる"
date = "2019-04-23T21:06:30+09:00"
description = "結果だけを見れば直感に反するかも知れないが，こうやって実際にコードを書いてみると納得感が強まる。 やっぱプログラマはコードを書いてナンボだよね（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "math", "random", "programming", "golang", "probability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[cakes で連載中の「数学ガールの秘密ノート」](https://cakes.mu/series/339 "数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）")。
今回の「[確率の冒険：《命》に関わる確率（前編）](https://cakes.mu/posts/25349 "第257回　確率の冒険：《命》に関わる確率（前編）｜数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）")」はなかなか面白かった。
確率や統計は直感に反する場合があって，そういう事例を考えるのは本当に楽しい。

今回のプロブレムを整理してみよう。

{{< fig-quote type="md" title="確率の冒険：《命》に関わる確率（前編）" link="https://cakes.mu/posts/25349" >}}
- 《○○菌》に感染している人は《全人口の1%》である
- ○○菌に感染しているか否かを調べる《判定キット》がある。判定キットは感染しているか否かを《90%の確率》で正しく判定する
    - 《感染している》という判定のことを《陽性》と呼ぶ
    - 《感染していない》という判定のことを《陰性》と呼ぶ
- ある人を判定キットで調べたら《陽性》だった。この人が○○菌に実際に感染している確率を求めよ
{{< /fig-quote >}}

きちんとした解法は本編を読んでいただくとして，この記事では疑似乱数を用いたシミュレーションで解いてみる。

## 指定した確率で真偽を出力するクラス

まずは指定した確率で真偽を出力するクラスを考えてみる。
[Go 言語]でね（笑）

こんな感じでどうだろう。

```go
package prob

import (
	"math"
	"math/rand"
	"time"
)

func New(p float64) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		max := 1000000
		limit := percent(p, max)
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			n := rnd.Intn(max) + 1
			ch <- n < limit
		}
	}()
	return ch
}

func percent(f float64, max int) int {
	if f < 0 {
		return 0
	}
	if f > 1.0 {
		return max
	}
	return int(math.Floor(f*float64(max) + 0.5))
}
```

これで

```go
ch := prob.New(0.1)
```

とすれば channel `ch` から10%の確率で true が出力される筈である。
サンプル・コードを書いて試してみよう。

```go
func probability(try int, ch <-chan bool) int {
	count := 0
	for i := 0; i < try; i++ {
		if <-ch {
			count++
		}
	}
	return count
}

func main() {
	ch := prob.New(0.1)
	min := float64(1)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	try := 10000
	tryf := float64(try)
	ps := make([]float64, 0, try)
	for i := 0; i < try; i++ {
		count := probability(try, ch)
		p := float64(count) / tryf
		ps = append(ps, p)
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
		sum += p
		sum2 += p * p
	}
	fmt.Printf("minimum value: %7.5f\n", min)
	fmt.Printf("maximum value: %7.5f\n", max)
	ave := sum / tryf
	fmt.Printf("average: %7.5f\n", ave)
	devi := math.Sqrt(sum2/tryf - ave*ave) //standard deviation
	ct := 0
	for _, p := range ps {
		if ave-devi <= p && p <= ave+devi {
			ct++
		}
	}
	fmt.Printf("standard deviation: %7.5f (%4.1f%%)\n", devi, float64(ct)*100.0/tryf)
}
```

10,000回の試行で割合を求める処理をワンセットとしてこれを10,000セット繰り返し，最小値・最大値・平均値・標準偏差を求めている。

これを実行するとこんな感じになる。

```text
$ go run prob/sample/sample.go
minimum value: 0.08930
maximum value: 0.11160
average: 0.09999
standard deviation: 0.00299 (67.9%)
```

まぁこんなもんかな（笑）

## 感染者と検査キットを定義する。

ではこの `prob` パッケージを使って感染者と検査キットを定義してみる。

こんな感じでどうだろう。

```go
package note257

import (
	"github.com/spiegel-im-spiegel/mathgirl-problem/prob"
)

type People struct {
	infect <-chan bool
}

type Person bool

func NewPeople() *People {
	return &People{infect: prob.New(0.01)}
}

func (ppl *People) SelectPersion() Person {
	return Person(<-ppl.infect)
}

func (psn Person) Infection() bool {
	return bool(psn)
}

type TestKit struct {
	probability <-chan bool
}

func NewTestKit() *TestKit {
	return &TestKit{probability: prob.New(0.9)}
}

func (tk *TestKit) Inspect(psn Person) bool {
	if psn.Infection() {
		return <-tk.probability
	}
	return !<-tk.probability
}
```

まず `People` を定義し，この中から `People.SelectPersion()` でサンプル `Person` を選び出す。
このサンプルは1%の確率で感染している。
感染しているかどうかは `Person.Infection()` 関数で分かる。

判定キットは `TestKit` で定義され `TestKit.Inspect()` 関数で検査結果が出る。
このとき

- （全体の1%存在する）感染者は90%の確率で陽性（true）になる
- （全体の99%存在する）非感染者は10%の確率で陽性（true）になる

のがポイントである。

このパッケージを使って実際に検査を行ってみる。

```go
func probability(ppl *note257.People, tk *note257.TestKit, try int) float64 {
	total := 0
	count := 0
	for i := 0; i < try; i++ {
		psn := ppl.SelectPersion()
		if tk.Inspect(psn) {
			total++
			if psn.Infection() {
				count++
			}
		}
	}
	return float64(count) / float64(total)
}

func main() {
	flag.Parse()
	try, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ppl := note257.NewPeople()
	tk := note257.NewTestKit()
	min := float64(1)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	tryf := float64(try)
	ps := make([]float64, 0, try)
	for i := 0; i < try; i++ {
		p := probability(ppl, tk, try*10)
		ps = append(ps, p)
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
		sum += p
		sum2 += p * p
	}
	fmt.Printf("minimum value: %7.5f\n", min)
	fmt.Printf("maximum value: %7.5f\n", max)
	ave := sum / tryf
	fmt.Printf("average: %7.5f\n", ave)
	devi := math.Sqrt(sum2/tryf - ave*ave) //standard deviation
	ct := 0
	for _, p := range ps {
		if ave-devi <= p && p <= ave+devi {
			ct++
		}
	}
	fmt.Printf("standard deviation: %7.5f (%4.1f%%)\n", devi, float64(ct)*100.0/tryf)
}

```

`probability()` 関数で判定キットの結果が陽性だった人の中で実際に感染している人の割合を返している。

たとえば

```text
$ go run note257/sample/sample.go 10000
```

とすれば `probability()` 関数による試行を10,000回繰り返すわけだ。
私のマシンは遅いので，これを実行するとめっさ時間がかかるのだが，まぁやってみよう。

```text
$ go run note257/sample/sample.go 10000
minimum value: 0.08061
maximum value: 0.10279
average: 0.09092
standard deviation: 0.00272 (68.5%)
```

平均値が少し高めに出たけど，こんなものかな。

結果だけを見れば直感に反するかも知れないが，こうやって実際にコードを書いてみると納得感が強まる。
やっぱプログラマはコードを書いてナンボだよね（笑）

## ブックマーク

- [第257回　確率の冒険：《命》に関わる確率（前編）｜数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）](https://cakes.mu/posts/25349)
    - [第258回　確率の冒険：《命》に関わる確率（後編）｜数学ガールの秘密ノート｜結城浩｜cakes（ケイクス）](https://cakes.mu/posts/25465)

- [モンテカルロ法による円周率の推定（その3 Gaussian）]({{< ref "/golang/estimate-of-pi-3-gaussian.md" >}})
- [モンティ・ホール問題で遊ぶ]({{< ref "/remark/2018/04/monty-hall-problem.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%81%AE%E6%95%B0%E5%AD%A6-%E7%AC%AC2%E7%89%88-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B079JLW5YN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079JLW5YN"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QDhrqqEtL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9E%E3%81%AE%E6%95%B0%E5%AD%A6-%E7%AC%AC2%E7%89%88-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B079JLW5YN?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B079JLW5YN">プログラマの数学 第2版</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2018-01-16 (Release 2018-02-08)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B079JLW5YN</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">タイトル通りプログラマ必読書。第2版では機械学習に関する章が付録に追加された。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-19">2018-03-19</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938"><img src="https://images-fe.ssl-images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" width="112" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%84%E3%81%8B%E3%81%AB%E3%81%97%E3%81%A6%E5%95%8F%E9%A1%8C%E3%82%92%E3%81%A8%E3%81%8F%E3%81%8B-G-%E3%83%9D%E3%83%AA%E3%82%A2/dp/4621045938?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621045938">いかにして問題をとくか</a></dt>
	<dd>G. ポリア</dd>
	<dd>G. Polya (原著), 柿内 賢信 (翻訳)</dd>
    <dd>丸善 1975-04-01</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4621045938, EAN: 9784621045930</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

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


[Go 言語]: https://golang.org/ "The Go Programming Language"
