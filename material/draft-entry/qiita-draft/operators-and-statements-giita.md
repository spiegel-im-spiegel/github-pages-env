# Go 言語の `++` や `--` は演算子ではない

面白い記事を見つけたので。

- [web制作者にもわかる、Swift 3が++と--を削除した理由 - Qiita](http://qiita.com/tonkotsuboy_com/items/0adc5dac54e690fcf706)

[Go 言語]にも `++` や `--` はあるが，実はこれは演算子ではなく [statement](https://go.dev/ref/spec#IncDec_statements) を構成する special token である。

```
IncDecStmt = Expression ( "++" | "--" ) .
```

故に，うっかり

```go
package main

import "fmt"

func main() {
	i := 1
	a := i++
	fmt.Println(a)
}
```

などと書いてコンパイル・実行しようとすると

> syntax error: unexpected ++ at end of statement

とか言われて怒られる[^svd]。この場合は素直に

[^svd]: `:=` で構成される statement は [short variable declarations](https://go.dev/ref/spec#Short_variable_declarations) と呼ばれる「宣言」（代入ではない）を構成するもので，その右辺は [Expression](https://go.dev/ref/spec#Expression) （のリスト）でなければならない。したがって statement は記述できない。

```go
package main

import "fmt"

func main() {
	i := 1
	i++
	a := i
	fmt.Println(a)
}
```

と書くか

```go
package main

import "fmt"

func main() {
	i := 1
	a := i
	i++
	fmt.Println(a)
}
```

と書けばよい。「前置」とか「後置」とか考える必要はないのだ[^op]。

[^op]: 個人的な感想で申し訳ないが，私は「演算子は変数の状態を変えてはいけない」と考えている。C/C++ は多くの token を演算子として扱っていて，その中には `++/--` や `=` などの変数の状態を変えるものも含まれている。 C/C++ では `++/--` を前置にすべきか後置にすべきかというのでよく議論になるが，私は変数の状態を変える演算子は他の演算子と混ぜるべきではないと思っている（つまり statement を分ける。やりすぎると可読性が落ちるけど）。そういう意味では [Go 言語]のように演算子と見なさないようにしたり Swift のようにいっそ `++/--` を言語仕様から排除するというのは妥当な割り切りだと思う。ちなみに [Go 言語]では channel 受信を示す `<-` を単項演算子としている。変数の状態を変える演算子はこれが唯一である。

ちなみに `++` や `--` は以下の [assignment statements](https://go.dev/ref/spec#Assignments) と同義である。

| IncDec statement | Assignment |
|------------------|------------|
| `x++`            | `x += 1`   |
| `x--`            | `x -= 1`   |

当然ながら

```go
package main

import "fmt"

func main() {
	i := 1
	a := i += 1
	fmt.Println(a)
}
```

も以下のエラーを吐いて怒られる。

> syntax error: unexpected += at end of statement

最初のコードと比較すれば `++` や `--` が statement であるということの意味が実感できるかもしれない。

## 補足：演算子について

ブコメで頭の悪いコメントが見られたので一応補足しておく（さすがはてなは酷いいんたーねっつですねw）。

[Go 言語]では明確に言語仕様が決められている（大抵のプログラミング言語はそうだと思うけど）。

- [The Go Programming Language Specification - The Go Programming Language](https://go.dev/ref/spec)

この中で[演算子（operator）](https://go.dev/ref/spec#Operators)についても厳密に決められている。この点において誤解や独自解釈の余地はない。 [Go 言語]における式（expression）と演算子の定義は以下の通り。

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```

またこの中にある [PrimaryExpr (primary expression)](https://go.dev/ref/spec#PrimaryExpr) の定義は以下のとおりである[^el]。

[^el]: ちなみに ExpressionList は Expression をカンマで繋いで列挙したものである。

```
PrimaryExpr =
	Operand |
	Conversion |
	PrimaryExpr Selector |
	PrimaryExpr Index |
	PrimaryExpr Slice |
	PrimaryExpr TypeAssertion |
	PrimaryExpr Arguments .

Selector       = "." identifier .
Index          = "[" Expression "]" .
Slice          = "[" [ Expression ] ":" [ Expression ] "]" |
                 "[" [ Expression ] ":" Expression ":" Expression "]" .
TypeAssertion  = "." "(" Type ")" .
Arguments      = "(" [ ( ExpressionList | Type [ "," ExpressionList ] ) [ "..." ] [ "," ] ] ")" .
```

更に二項演算子（binary_op）には優先順位が決められているが，それは[言語仕様](https://go.dev/ref/spec#Operators)を読みなはれ。

たとえば C/C++ 言語では括弧も演算子の一種であり優先順位が決められているが， [Go 言語]では括弧は演算子ではない（故に二項演算子のような優先順位はない）。代入（`=` 等）も同様だ。私のように C/C++ や Java などから来た人間にとってはこの仕様上の差異をきちんと把握しておくことは結構重要で，これを怠ると（今回の `++/--` のように）時に躓きポイントになる。新しい言語を学ぶ際は（チュートリアルで満足するのではなく）まずひと通りは言語仕様を眺めておくことをお勧めする。

こういった仕様上の差異はどちらが正しいというのでもない（もちろん解釈の問題でも言葉遊びでもない）。そういう仕様だと受け入れた上でどう記述していくかなのである。「制約は構造を生む」（「数学ガール」より）というやつだ。

[Go 言語]: https://golang.org/ "The Go Programming Language"

----
