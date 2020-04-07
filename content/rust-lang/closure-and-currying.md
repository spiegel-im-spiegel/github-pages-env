+++
title = "関数閉包で遊ぶ"
date =  "2020-04-07T18:55:33+09:00"
description = "関数閉包の記述では fn 構文は使えない。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "closure", "currying", "ownership" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は小ネタとして関数閉包（closure）で遊んでみる。
起点となるコードはこれにしよう。

```rust
fn add(x: i32, y: i32) -> i32 {
    x + y
}

fn main() {
    let x = 1;
    let y = 2;
    println!("{} + {} = {}", x, y, add(x, y)); //Output: 1 + 2 = 3
}
```

`add()` 関数は見ての通り整数の足し算。
この関数を関数閉包で表してみたい。

## 関数閉包で記述する

関数閉包の記述では `fn` 構文は使えない。
こんな感じに書く。

```rust {hl_lines=[2]}
fn main() {
    let add = |x, y| {x + y};
    let x = 1;
    let y = 2;
    println!("{} + {} = {}", x, y, add(x, y)); //Output: 1 + 2 = 3
}
```

`|...|` で囲まれている部分で引数を定義し `{...}` で囲まれている部分が関数本体になる。
ちなみに，上のような記述であれば波括弧は省略できる。
ていうか `rustfmt` コマンドで整形すると問答無用で消される（笑）

```rust
let add = |x, y| x + y;
```

引数の型は推論可能であれば省略できるが，型注釈で明示することもできる。

```rust
let add = |x: i32, y: i32| x + y;
```

この関数閉包をもう少し弄ってみよう。

### スマートポインタ

関数閉包はスマートポインタを表すトレイト `Box<T>` で囲むことができる。

```rust {hl_lines=[2]}
fn main() {
    let add = Box::new(|x, y| x + y);
    let x = 1;
    let y = 2;
    println!("{} + {} = {}", x, y, add(x, y)); //Output: 1 + 2 = 3
}
```

敢えて `add` に型注釈を付けるとこんな感じになる。

```rust
let add: Box<dyn Fn(i32, i32) -> i32> = Box::new(|x, y| x + y);
```

ちなみに `Fn` は関数閉包を表すトレイトである。

これだけでは何も面白くないが，実は「高階関数（higher-order function）」で威力を発揮する。
高階関数の定番といえばアレだよね。
そう「カリー化（currying）」である[^cry1]。

[^cry1]: カリー化については拙文「[カリー化に関する覚え書き]({{< ref "/remark/2020/03/currying.md" >}})」を参考にどうぞ。

## カリー化と所有権

では早速 `add()` 関数をカリー化してみよう。
こんな感じ。

```rust {hl_lines=[2]}
fn add(x: i32) -> Box<dyn Fn(i32) -> i32> {
    Box::new(move |y| x + y)
}

fn main() {
    let x = 1;
    let y = 2;
    println!("{} + {} = {}", x, y, add(x)(y)); //Output: 1 + 2 = 3
    let increment = add(x); //partial application
    println!("add({}) -> increment({}) = {}", x, y, increment(y)); //Output: add(1) -> increment(2) = 3
}
```

`move` キーワードを使って関数閉包内の変数 `x` の所有権を明示的に移動してる点に注意。
`move` を付けないと借用とみなされるが `add()` 関数を抜けるとスコープから外れるためコンパイルエラーになる。

ちなみに `main()` 関数スコープ内で記述するのであれば，もっと簡単に記述できる。
こんな感じ。

```rust {hl_lines=[2]}
fn main() {
    let add = |x| move |y| x + y;
    let x = 1;
    let y = 2;
    println!("{} + {} = {}", x, y, add(x)(y)); //Output: 1 + 2 = 3
    let increment = add(x); //partial application
    println!("add({}) -> increment({}) = {}", x, y, increment(y)); //Output: add(1) -> increment(2) = 3
}
```

よーし，うむうむ，よーし。

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
