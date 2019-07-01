+++
description = "今回は「つまみ食い」的に演算子（operator）とステートメント（statement）について解説する。"
draft = false
date = "2017-02-20T21:07:33+09:00"
update = "2018-01-23T20:13:49+09:00"
title = "演算子とステートメント"
tags = ["golang", "programming", "language"]

[author]
  license = "by-sa"
  tumblr = ""
  url = "https://baldanders.info/profile/"
  avatar = "/images/avatar.jpg"
  instagram = "spiegel_2007"
  name = "Spiegel"
  flickr = "spiegel"
  twitter = "spiegel_2007"
  github = "spiegel-im-spiegel"
  flattr = ""
  linkedin = "spiegelimspiegel"
  facebook = "spiegel.im.spiegel"
+++

（この記事は [Qiita](http://qiita.com/) に投稿した「[Go 言語の `++` や `--` は演算子ではない - Qiita](http://qiita.com/spiegel-im-spiegel/items/2c6cf5ff44d816d1be7b)」を大幅に修正して再構成したものです）

あるプログラミング言語を習得する際に最も早道なのは「たくさんの（他人の）コードを読むこと」であり「たくさんのコードを（コピペではなく自分で）書く」ことである。
これは間違いない。
しかし，その言語の仕様をきちんと把握してないとコードを読んでも間違って理解するかもしれないし，何より実際に自分でコードを書く際に躓く原因になる。

というわけで，少なくとも学ぶ言語の言語仕様を一度は眺めておくことをお勧めする。
[Go 言語]の場合は以下のページで言語仕様を見ることができる（“[A Tour of Go](https://tour.golang.org/)” の後で読むと頭に入りやすいかもしれない）。

- [The Go Programming Language Specification - The Go Programming Language](https://golang.org/ref/spec)
- [Goプログラミング言語仕様 - golang.jp](http://golang.jp/go_spec) : 日本語だが内容が古いので注意

今回は「つまみ食い」的に演算子（operator）とステートメント（statement）[^stmt] について軽く紹介してみる。

[^stmt]: [言語仕様]における “statement” の適切な日本語訳が思いつかなかったので，今回はカタカナにのばして「ステートメント」と表記する。教えて，英語得手の人。

## ステートメント{#stmnt}

[Go 言語]においては「ステートメント」は以下のように定義されている。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
Statement =
    Declaration | LabeledStmt | SimpleStmt |
    GoStmt | ReturnStmt | BreakStmt | ContinueStmt | GotoStmt |
    FallthroughStmt | Block | IfStmt | SwitchStmt | SelectStmt | ForStmt |
    DeferStmt .

SimpleStmt = EmptyStmt | ExpressionStmt | SendStmt | IncDecStmt | Assignment | ShortVarDecl .
```
{{% /fig-gen %}}

まぁ名前で何か大体わかると思う。
ここでは `SimpleStmt` (simple statement) に絞って紹介しよう。

### Empty Statements

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
EmptyStmt = .
```
{{% /fig-gen %}}

文字通り空のステートメント。

### Expression Statements

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
ExpressionStmt = Expression .
```
{{% /fig-gen %}}

式（expression）を表すステートメント。
関数呼び出しや受信操作のコンテキスト内に記述できる。

さらに式は以下のように定義される。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```
{{% /fig-gen %}}

（`PrimaryExpr` (primary expression) については割愛する。詳細は「[言語仕様]」で確かめてみてください。ここでは `Expression` を構成する要素にはステートメントが含まれないことに注目）

`binary_op`, `rel_op`, `add_op`, `mul_op`, `unary_op` は演算子である。
演算子については[後述]({{< relref "#op" >}})する。

### Send Statements

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
SendStmt = Channel "<-" Expression .
Channel  = Expression .
```
{{% /fig-gen %}}

[channel] 送信のステートメント。

### IncDec Statements

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
IncDecStmt = Expression ( "++" | "--" ) .
```
{{% /fig-gen %}}

インクリメント（increment）およびデクリメント（decrement）のステートメント。
C/C++ のように `++x` みたいな記述はできないので注意。

ちなみに `IncDecStmt` は次の代入ステートメントの以下の記述と同じである。

| IncDec statement | Assignment |
|------------------|------------|
| `x++`            | `x += 1`   |
| `x--`            | `x -= 1`   |

### Assignments

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
Assignment = ExpressionList assign_op ExpressionList .
assign_op = [ add_op | mul_op ] "=" .
```
{{% /fig-gen %}}

代入。
`add_op`, `mul_op` は先ほど出た `Expression` の演算子を指す。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .
```
{{% /fig-gen %}}

定義だと `assign_op` は演算子っぽく見える。
そもそも代入を “assignment operation” と表記しているのだ。
どうなんだろう。
まぁ，いずれにしろ代入自体は間違いなくステートメントであり式の中には含められない。

ちなみに `ExpressionList` は `Expression` を列挙したものである。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
ExpressionList = Expression { "," Expression } .
```
{{% /fig-gen %}}

これにより代入の左辺・右辺を組（tuple）で記述できる。
たとえば2つの変数の値を入れ替える場合は以下のように記述する。

```go
x, y = y, x
```

### Short Variable Declarations

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
ShortVarDecl = IdentifierList ":=" ExpressionList .
```
{{% /fig-gen %}}

変数宣言の短縮表現。
`var` キーワードを使った以下の表現と同じ。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
"var" IdentifierList = ExpressionList .
```
{{% /fig-gen %}}

`IdentifierList` は `identifier` を列挙したもので

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
IdentifierList = identifier { "," identifier } .
```
{{% /fig-gen %}}

これにより `identifier` で記述される複数の変数をまとめて宣言・初期化できる。
`identifier` の定義は以下の通り

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
identifier = letter { letter | unicode_digit } .
```
{{% /fig-gen %}}

ちなみに変数名となる `identifier` は全ての Unicode 文字を許容する。
なので日本語交じりでこんな書き方もできる。

```go
package main

import "fmt"

func main() {
    わーい := "わーい！ たのしー！"
    fmt.Println(わーい)
}
```

## 演算子{#op}

さて，式と演算子の定義を再び掲げる。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
```text
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```
{{% /fig-gen %}}

[Go 言語]で式に使える演算子はここに挙げられているものが全てである。
このうち二項演算子（`binary_op`）には優先順位が付けられている。

{{% fig-gen type="md" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" %}}
| Precedence | Operator |
|:----------:|:---------|
|          5 | `*  /  %  <<  >>  &  &^` |
|          4 | `+  -`  \|  `^` |
|          3 | `==  !=  <  <=  >  >=` |
|          2 | `&&` |
|          1 | \|\| |
{{% /fig-gen %}}

なお単項演算子（`unary_op`）は二項演算子よりも高い優先順位で機能する。
したがって全体としてはこんな感じだろうか。

| Precedence | Operator |
|:----------:|:---------|
|          6 | `unary_op` |
|          5 | `mul_op` |
|          4 | `add_op` |
|          3 | `rel_op` |
|          2 | `&&` |
|          1 | \|\| |

### インクリメント／デクリメントは演算子ではない

たとえば [C 言語の演算子](http://www.bohyoh.com/CandCPP/C/operator.html "BohYoh.com【Ｃ言語講座】演算子一覧表")と比較すると [Go 言語]ではインクリメント（`++`）／デクリメント（`--`）が演算子として扱われていないことに気付く[^op]。
[Go 言語]ではインクリメント／デクリメント（および代入）はステートメントである。

[^op]: 「[言語仕様]」では文章上の表現として operator と記述しているところが幾つかあるが定義としては演算子として扱われていない。

これはどういうことかというと，たとえば C 言語のコードに似せて

```go
package main

import "fmt"

func main() {
    i := 1
    fmt.Println(i++)
}
```

と書いてコンパイルしようとしても

```text
syntax error: unexpected ++, expecting comma or )
```

とコンパイルエラーになるということである（式を構成する要素にステートメントは含まれないことを思い出してほしい）。
これはコードを，以下のように，代入に置き換えたほうが直感的で分かりやすいかもしれない。
（この場合も「`syntax error: unexpected +=, expecting comma or )`」でコンパイルエラーになる）

```go
package main

import "fmt"

func main() {
    i := 1
    fmt.Println(i+=1)
}
```

私見で申し訳ないが，私は「式中の演算子は変数の状態を変えるべきではない」と考えている。
たとえば C/C++ では `++`, `--` 演算子を前置にすべきか後置にすべきかというのでよく議論になる[^pp]。
しかし，これはそもそも `++`, `--` 演算子が式の中で対象の変数の状態を変えてしまうことに問題があるのだ。

[^pp]: C/C++ で `++`, `--` 演算子を前置にするか後置にするかという問題は挙動の分かりにくさと実行時パフォーマンスの2つの論点がある。いずれにしろ前置に統一する方がよいと言われているが，[パフォーマンスに関しては異論もある](http://cpp.aquariuscode.com/preincriment-vs-postincriment "前置インクリメント vs 後置インクリメント | 闇夜のC++")ようだ。

[Go 言語]ではインクリメントやデクリメント（あるいは代入）といった変数の状態を変える操作をステートメントとし，式の中に埋め込むことを禁止することでこの問題を回避しているように見える。
式の中で変数の状態が変わらないのであれば副作用を気にすることなく安全にコードを書くことができる。

ただし例外がある。

### [channel] 操作

[channel] 操作では，送信はステートメントだが受信は `<-` 単項演算子を使う。
したがって，こんな記述もできる（意味があるかどうかはともかく）。

```go
ch2 <- <-ch1
```

[channel] 受信を含んだ式では [channel] 変数の状態が変わる副作用（特に deadlock 関連）に注意を払う必要がある。

## とまぁ，こんな感じで

手を動かしながら「[言語仕様]」を眺めていくと，いろいろ発見があって楽しいと思う。

ではまた。

## ブックマーク

- [web制作者にもわかる、Swift 3が++と--を削除した理由 - Qiita](http://qiita.com/tonkotsuboy_com/items/0adc5dac54e690fcf706) : Swift 3 では `++`, `--` 演算子を仕様から削除したらしい
- [else ifにも代入文が書ける #golang - Qiita](https://qiita.com/tenntenn/items/791bb47f4cee178b52c3) : if ステートメントに関する話

[Go 言語]: https://golang.org/ "The Go Programming Language"
[言語仕様]: https://golang.org/ref/spec "The Go Programming Language Specification - The Go Programming Language"
[channel]: http://golang.org/ref/spec#Channel_types
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"

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
