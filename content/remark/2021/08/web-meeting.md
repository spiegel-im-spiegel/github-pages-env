+++
title = "また散財してしまった..."
date =  "2021-08-09T14:36:17+09:00"
description = "『プログラミング言語Go』読書会の反省文と感想文"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "book", "communication", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## ついカッとなってやった。

今は反省している。

先日，久しぶりに『[プログラミング言語Go]』の読書会に参加したのだが

- [第15回『プログラミング言語Go』オンライン読書会（第2サイクル） - connpass](https://gpl-reading.connpass.com/event/218308/)

ちょっと思い付いてパソコン（Ubuntu 機）で繋ごうとか考えてしまったのが運の尽き（今までは Android タブレットで繋いでた）。

で，[ヘッドセット]({{< ref "/remark/2021/06/aftershokz-aeropex.md" >}} "Amazon Prime Day 2021 の戦利品")を繋ぐ Bluetooth アダプタは勤務先に置きっぱなしにしてたので「じゃあ家電量販店で買ってしまうか」と読書会イベント当日の午前中に徒歩圏内の家電量販店に出掛けたわけさ（田舎の「徒歩圏内」は広い。片道5〜10kmくらい？）。

で，考えたら Web カメラも必要だし，セキュリティ・リスクも考えたら USB ハブもスイッチ付きにしたい。
というわけで以下を購入した。

{{% review-paapi "B078J9MR75" %}} <!-- Web カメラ -->
{{% review-paapi "B01K8XRQ6O" %}} <!-- Bluetooth 4 アダプタ -->
{{% review-paapi "B01MQDYCB0" %}} <!-- USB ハブ（スイッチ付き） -->

いや，これ， Amazon で買ったら8千円ほどだけど（Prime 会員なので送料無料），家電量販店で買ったから諭吉先生がまるっと飛んでいったぞ。
ホンマ，家電量販店で買うと碌なことにならないな。
いや，我ながら魔が差したとしか思えない。
なにやってんだ，私 `orz`

{{< ruby "今度はもっと上手くやる" >}}もう二度としない{{< /ruby >}}。
といいつつ，またやらかすんだろうなぁ。
ヒトとは何度でも間違う生き物である。

ちなみに上の機器でちゃんと Zoom 接続できた。
Zoom アプリは使わず（Linux 版のアプリがあるかどうかも知らないが） Firefox のプライベートウィンドウで Zoom サイトを開いて問題なくサインインできた。
次回以降もこれで行こう。

## 第15回 読書会の感想

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

読書会の感想だが，今回もなかなか面白かった。

### 「参照型」などない

『[プログラミング言語Go]』ではポインタ，スライス，マップ，チャネル，関数の各型を「参照型」と呼んでいるが “[The Go Programming Language Specification][Golang Specs]” (以降 "[Golang Specs]” と略称する) には「参照型」に相当する記述はない。
[Go] には C++ や Java で言うところの「参照（reference）」の仕組みは持っていない。
あるのはポインタを使った直接的なアドレッシングである。
Interface 型のようなボックス化（boxing）の仕組みでさえポインタで操作するのだ。

私はぺーぺーの新人の頃に アセンブラ→C言語 とアドレッシングについて叩き込まれたので今更 [Go] のポインタに躓くことはないが，慣れてない人はやっぱ戸惑うよねぇ。

### new など要らぬ？

私は C++ や Java から来たということもあり，最初の頃は `new()` 関数を多用していた。
たとえば可変長バッファの [`bytes`]`.Buffer` を生成する際も

```
buf := new(bytes.Buffer)
```

と書かいていたが，これはリテラル表現を使って

```
buf := &bytes.Buffer{}
```

と書いても全く同じなのだ。
これを知ってからは `new()` 関数はほぼ使わなくなった（古いコードには残ってるがw）。

もっと言うと `new()` 関数の機能は

{{< fig-quote type="markdown" title="The Go Programming Language Specification" link="https://golang.org/ref/spec" lang="en" >}}
{{% quote %}}The built-in function new takes a type `T`, allocates storage for a [variable](https://golang.org/ref/spec#Variables) of that type at run time, and returns a value of type `*T` [pointing](https://golang.org/ref/spec#Pointer_types) to it. The variable is initialized as described in the section on [initial values](https://golang.org/ref/spec#The_zero_value){{% /quote %}}.
{{< /fig-quote >}}

であって，割り当てる領域がヒープであるとは記されてないんだよね。
ある関数スコープ内でのみ生成・利用されるインスタンスであれば `new()` 関数で生成してもスタック上に割り当てられるかもしれない。
インスタンスがどこに生成されるかはコンパイル時の最適化で決まる。

とするとますます `new()` 関数は要らない子になってしまう。
ただし Generics が正式に実装されると `new()` 関数の使いみちが増えるんじゃないか，ということらしい。

### 「名前付き型」はなくなった

『[プログラミング言語Go]』2.5章には

{{< fig-quote type="markdown" title="『プログラミング言語Go』2.5章" link="https://www.amazon.co.jp/dp/B094PRR7PZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}`type` 宣言は、既存の型と同じ*基底型*（<i>underlying type</i>）を持つ新たな*名前付き型*（<i>named type</i>）を定義します{{% /quote %}}
{{< /fig-quote >}}

と書かれてある。
たとえば `src/syscall/syscall_unix.go` の

```go
type Errno uintptr
```

みたいな感じ。

でも [Go] 1.9 で型エイリアス（type alias）がサポートされたことで「名前付き型」という文言は [Golang Specs] からなくなったんだそうな。
ちなみに型エイリアスは，たとえば `golang.org/x/sys/unix` パッケージで

```go {hl_lines=["13-15"]}
// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos) && go1.9
// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris zos
// +build go1.9

package unix

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
```

みたいに定義されているやつ。
文字通りの別名定義で「新たな型」でなく「全く同じ型」として扱うことができる。

### レキシカルブロックの罠

『[プログラミング言語Go]』2.7章では `{ ... }` 等で囲まれた記述ブロックをレキシカルブロック（lexical block）と呼んでいる（厳密にはちょっと違うが）。
で，レキシカルブロック間が親子関係にあるとき，親が宣言する変数などの識別名を子が上書き再宣言して親の識別子を隠蔽することができる。
いわゆる「シャドウイング（shadowing）」である。
なので

{{< fig-quote class="nobox" type="markdown" title="『プログラミング言語Go』2.7章" link="https://www.amazon.co.jp/dp/B094PRR7PZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go
func main() {
    x := "hellow!"
    for i := 0; i < len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Printf("%c", x) // "HELLO"
        }
    }
}
```
{{< /fig-quote >}}

みたいな鬼畜なコードも書けてしまう（`:=` は代入ではなく初期化付き宣言である点に注意）。
まぁ，この程度なら（鬼畜ではあるが） `main()` 関数のスコープの中に収まっているのでまだマシだが，うっかり `init()` 関数を使った初期化で

{{< fig-quote class="nobox" type="markdown" title="『プログラミング言語Go』2.7章" link="https://www.amazon.co.jp/dp/B094PRR7PZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go
var cwd string

func init() {
    cwd, err := os.Getwd() //注意: 誤り！
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}
```
{{< /fig-quote >}}


みたいなコードを書いて[初期化に失敗](https://play.golang.org/p/QR0ROgM2XjH)したりするわけだ。
こわいこわい（笑）

## ブックマーク

- [Goの言語仕様書精読のススメ & 英語彙集](https://zenn.dev/hsaki/articles/gospecdictionary)
- [golangではスタックとヒープを気にする必要が無い](https://zenn.dev/rookxx/articles/golang-stack-and-heap)

[Go]: https://golang.org/ "The Go Programming Language"
[`bytes`]: https://pkg.go.dev/bytes "bytes · pkg.go.dev"
[プログラミング言語Go]: https://www.amazon.co.jp/dp/B094PRR7PZ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1
[Golang Specs]: https://golang.org/ref/spec "The Go Programming Language Specification - The Go Programming Language"

## 参考図書

{{% review-paapi "B094PRR7PZ" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B07RRQ59JR" %}} <!-- AfterShokz Aeropex 骨伝導ヘッドセット -->
