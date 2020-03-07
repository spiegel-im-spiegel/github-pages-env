+++
title = "変数束縛"
date =  "2020-03-01T21:08:53+09:00"
description = "「値」は所有者である「変数」に束縛される。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "ownership", "variable-bindings" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Rust] の所有権（ownership）のルールは以下の3つ。

1. Each value in Rust has a variable that’s called its owner (Rustの各値は、所有者と呼ばれる変数と対応している)
2. There can only be one owner at a time (いかなる時も所有者は一つである)
3. When the owner goes out of scope, the value will be dropped (所有者がスコープから外れたら、値は破棄される)

「値」は所有者である「変数」に束縛される。
これを「変数束縛（variable bindings）」と呼ぶ。

実際に変数束縛がどのように機能しているかを見てみよう。

## Copy Semantics

[Rust] で定義されるデータ型は以下の通り。

- スカラー型
    - 整数 (`i8`, `u8`, ..., `isize`, `usize`)
    - 浮動小数点型（`f32`, `f64`）
    - 論理値型（`bool`）
    - 文字型（`char`）: Unicode コードポイント？
- 複合型
    - タプル型
    - 配列型

タプル型はこんな感じで記述できる。

```rust
fn main() {
	let tup = (500, 6.4, 1);
    println!("tup = {:?}", tup); //Output: tup = (500, 6.4, 1)
}
```

配列はこんな感じ。

```rust
fn main() {
	let ary = [1, 2, 3];
	println!("ary = {:?}", ary); //Output: ary = [1, 2, 3]
}
```

これらのデータ型は値と変数が一体になっていて（値が直接スタックに積まれるため），代入時に値のコピーが発生する。

```rust
fn main() {
	let mut ary1 = [1, 2, 3];
	let ary2 = ary1;
	ary1[0] *= 100;
	println!("ary1 = {:?}", ary1); //Output: ary1 = [100, 2, 3]
	println!("ary2 = {:?}", ary2); //Output: ary2 = [1, 2, 3]
}
```

イメージとしてはこんな感じ。

{{< fig-img src="./array.png" link="./array.svg" >}}

このようにデータ型では値をコピーすることによって変数束縛を担保している。

ちなみに `mut` は変数が可変（mutable）であることを示す。
[Rust] では，全ての変数は宣言時に `mut` キーワードを付けない限り不変（immutable）である。

## Move Semantics

[Rust] には上述のデータ型以外にもいくつかの型が存在する。
以下に主なものを挙げる。

- 構造体（`struct`）
- コレクション
    - 文字列（`String`）
    - ベクタ（`Vec<T>`）

たとえば構造体なら

```rust
struct Person {
	age: u32,
	name: String,
}

fn main() {
	let p1 = Person {
		age: 24,
		name: "alice".to_string(),
	};
	println!("p1 = {} ({})", p1.name, p1.age); //Output: p1 = alice (24)
}
```

てな感じに書ける。

データ型との違いは，変数で示しているものが，値そのものではなく，値への参照（のようなもの）という点である。

{{< fig-img src="./variable-bindings1.png" link="./variable-bindings1.svg" >}}

今度は変数 `p1` の値を別の変数に代入してみよう。

```rust {hl_lines=["12-13"]}
struct Person {
	age: u32,
	name: String,
}

fn main() {
	let p1 = Person {
		age: 24,
		name: "alice".to_string(),
	};
	println!("p1 = {} ({})", p1.name, p1.age); //Output: p1 = alice (24)
	let p2 = p1;
	println!("p2 = {} ({})", p2.name, p2.age); //Output: p2 = alice (24)
}
```

一見うまく言っているようだが，`println!` マクロ[^pr1] の位置を変えると

[^pr1]: `println!` って関数じゃなくてマクロなんですよ，奥さん。

```rust {hl_lines=["12"]}
struct Person {
	age: u32,
	name: String,
}

fn main() {
	let p1 = Person {
		age: 24,
		name: "alice".to_string(),
	};
	let p2 = p1;
	println!("p1 = {} ({})", p1.name, p1.age); //Error: value borrowed here after move
	println!("p2 = {} ({})", p2.name, p2.age);
}
```

コンパイルエラーになる。

これは

```rust
let p2 = p1;
```

の部分で値がコピーされず，所有権のみ移動してしまうため。

{{< fig-img src="./variable-bindings2.png" link="./variable-bindings2.svg" >}}

構造体やコレクションの各型の値はヒープ領域に置かれる。
ヒープ上の値に対して代入等を行うたびにコピーを行うのは高コストだし，かといって野放図に参照を増やすとヒープ管理が煩雑になってしまう。

そこで所有権を「移動」することでヒープ管理の最適化を行うわけだ。

## Copy/Clone Trait

じゃあ，構造体やコレクションのコピーはできないのかというと，ちゃんと救済措置はある。

[Rust] の標準ライブラリには `Clone` および `Copy` トレイトが用意されていてこれらを実装することで値のコピーが可能になる。

ちなみに「トレイト（trait）」とは，ここでは C++ や Java で言うところの interface クラスのようなものだと思っておけばよい。
トレイトに関しては多分その内ちゃんとした記事を書くと思う（笑）

たとえば先程の `Person` 構造体であれば

```rust {hl_lines=["6-13", 20]}
struct Person {
	age: u32,
	name: String,
}

impl Clone for Person {
	fn clone(&self) -> Self {
		Person {
			age: self.age,
			name: self.name.clone(),
		}
	}
}

fn main() {
	let p1 = Person {
		age: 24,
		name: "alice".to_string(),
	};
	let p2 = p1.clone();
	println!("p1 = {} ({})", p1.name, p1.age); //Output: p1 = alice (24)
	println!("p2 = {} ({})", p2.name, p2.age); //Output: p2 = alice (24)
}
```

とすればよい（トレイトの実装詳細については割愛する）。

{{< fig-img src="./variable-bindings3.png" link="./variable-bindings3.svg" >}}

あるいはもっと簡単に

```rust {hl_lines=[1, 12]}
#[derive(Clone)]
struct Person {
	age: u32,
	name: String,
}

fn main() {
	let p1 = Person {
		age: 24,
		name: "alice".to_string(),
	};
	let p2 = p1.clone();
    println!("p1 = {} ({})", p1.name, p1.age); //Output: p1 = alice (24)
	println!("p2 = {} ({})", p2.name, p2.age); //Output: p2 = alice (24)
}
```

といった記述もできる[^derive]。

[^derive]: `derive` はマクロに似た拡張構文の一種。詳細については割愛する。ていうか，これから勉強する（笑）

実を言うと

```rust
#[derive(Copy, Clone)]
```

とすれば代入時に自動でコピーが走るのだが， `Person` 構造体ではコンパイルエラーになってしまい上手く行かなかった。
色々試してみたが，どうも `String` を含む構造体の `Copy` トレイトは実装できないらしい？

なので，たとえば

```rust
#[derive(Copy, Clone)]
struct Complex {
	real: f64,
	image: f64,
}

fn main() {
	let c1 = Complex {
		real: 1.23,
		image: 4.5,
	};
	let c2 = c1;
	println!("c1 = {} + {}i", c1.real, c1.image); //Output: c1 = 1.23 + 4.5i
	println!("c2 = {} + {}i", c2.real, c2.image); //Output: c2 = 1.23 + 4.5i
}
```

とかいったコードは問題なく書ける。

ふむむー。

## ブックマーク

- [RustのコピーセマンティクスをCopyトレイトを実装して確認する ｜ Developers.IO](https://dev.classmethod.jp/server-side/language/rust-copy-semantics-use-copy-trait/)

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
