+++
title = "time.Time の比較が覚えれん！"
date =  "2020-09-07T11:16:53+09:00"
description = "一覧表にしておこう。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "time", "sort" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

時刻を表す [`time`]`.Time` 型は比較演算子（`==`, `<`, `>` 等）が使えないので `Equal()`, `Before()`, `After()` 各メソッドが用意されているのだけど， `Equal()` メソッドはともかく `Before()` や `After()`  は覚えれんっちうの！

まぁ[ドキュメント][`time`]を見れば済む話なのだが，毎回「どうだっけ？」と探すのもナニなので，この際，記事として纏めておくことにした。

まずはコードを書いてみる。
これ基本。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	year2000 := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	year3000 := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC)

	fmt.Printf("year3000.After(year2000) = %v\n", year3000.After(year2000))
	fmt.Printf("year3000.Before(year2000) = %v\n", year3000.Before(year2000))
	fmt.Printf("year2000.After(year3000) = %v\n", year2000.After(year3000))
	fmt.Printf("year2000.Before(year3000) = %v\n", year2000.Before(year3000))

}
```

これの実行結果は

```test
$ go run sample1.go 
year3000.After(year2000) = true
year3000.Before(year2000) = false
year2000.After(year3000) = false
year2000.Before(year3000) = true
```

となる。
`a.After(b)` は「`a` は `b` の後か？」と覚えればいいかな。

やっぱ面倒くさい。
一覧表にしておこう。

{{< fig-gen >}}
<table class="left">
<thead><tr>
    <th>関係</th>
    <th>メソッド</th>
    <th>返り値</th>
</tr></thead>
<tbody>
    <tr><td class="center" rowspan="3" style="vertical-align:middle;">$a = b$</td>
        <td><code>a.Equal(b)</code></td><td class="center"><code>true</code></td></tr>
    <tr><td><code>a.Before(b)</code></td><td class="center"><code>false</code></td></tr>
    <tr><td><code>a.After(b)</code></td><td class="center"><code>false</code></td></tr>

    <tr><td class="center" rowspan="3" style="vertical-align:middle;">$a \lt b$</td>
        <td><code>a.Equal(b)</code></td><td class="center"><code>false</code></td></tr>
    <tr><td><code>a.Before(b)</code></td><td class="center"><code>true</code></td></tr>
    <tr><td><code>a.After(b)</code></td><td class="center"><code>false</code></td></tr>

    <tr><td class="center" rowspan="3" style="vertical-align:middle;">$a \gt b$</td>
        <td><code>a.Equal(b)</code></td><td class="center"><code>false</code></td></tr>
    <tr><td><code>a.Before(b)</code></td><td class="center"><code>false</code></td></tr>
    <tr><td><code>a.After(b)</code></td><td class="center"><code>true</code></td></tr>
</tbody>
</table>
{{< /fig-gen >}}

## 【サンプル】時刻のソート

もう少し「ありそう」なサンプルを考えてみよう。
たとえば，以下のようなデータセット `eraList` があるとする。

```go
type Era struct {
	Name  string
	Start time.Time
}

var (
	jst     = time.FixedZone("JST", 9*60*60)
	eraList = []Era{
		{Name: "令和", Start: time.Date(2019, time.May, 1, 0, 0, 0, 0, jst)},
		{Name: "平成", Start: time.Date(1989, time.January, 8, 0, 0, 0, 0, jst)},
		{Name: "昭和", Start: time.Date(1926, time.December, 25, 0, 0, 0, 0, jst)},
		{Name: "大正", Start: time.Date(1912, time.July, 30, 0, 0, 0, 0, jst)},
	}
)
```

この `eraList` を時刻の昇順で並べ替えてみる。
こんな感じかな。

```go {hl_lines=[9]}
func (e Era) String() string {
	return fmt.Sprintf("%s (from %s)", e.Name, e.Start.Format("2006-01-02"))
}

func main() {
	fmt.Println(eraList)

	sort.Slice(eraList, func(i, j int) bool {
		return eraList[i].Start.Before(eraList[j].Start)
	})

	fmt.Println(eraList)
}
```

結果は以下の通り。

```text
$ go run sample2.go 
[令和 (from 2019-05-01) 平成 (from 1989-01-08) 昭和 (from 1926-12-25) 大正 (from 1912-07-30)]
[大正 (from 1912-07-30) 昭和 (from 1926-12-25) 平成 (from 1989-01-08) 令和 (from 2019-05-01)]
```

よーし，うむうむ，よーし。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`time`]: https://pkg.go.dev/time "time package · pkg.go.dev"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
