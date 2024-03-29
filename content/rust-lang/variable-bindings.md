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

[Rust] の所有権（ownership）のルールは以下の3つ（公式ドキュメントの「[所有権とは？](https://doc.rust-jp.rs/book/second-edition/ch04-01-what-is-ownership.html "所有権とは？ - The Rust Programming Language")」より抜粋）。

1. [Rust] の各値は、所有者と呼ばれる変数と対応している
2. いかなる時も所有者は一つである
3. 所有者がスコープから外れたら、値は破棄される

「値」は所有者である「変数」に束縛される。
これを「変数束縛（variable bindings）」と呼ぶ。

実際に変数束縛がどのように機能しているかを見てみよう。

## Copy Semantics

[Rust] において組み込みで定義されるデータ型は以下の通り。

- スカラ型
    - 整数 (`i8`, `u8`, ..., `isize`, `usize`)
    - 浮動小数点数型（`f32`, `f64`）
    - 論理値型（`bool`）
    - 文字型（`char`）: Unicode 符号点
- 複合型
    - タプル型
    - 配列型

（詳しくは「[型に関する覚え書き]({{< relref "./types.md" >}})」を参照のこと）

スカラ型およびスカラ型で構成される複合型は値と変数が一体になっていて（値が固定長で直接スタックに積まれるため），代入時に値のコピーが発生する。

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

{{< fig-img class="lightmode" src="./array.png" link="./array.svg" >}}

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
    - ハッシュマップ（`HashMap<K, V>`）

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

組み込みのデータ型との違いは，変数で示しているものが，値そのものではなく，値への参照（のようなもの[^ref1]）という点である。

[^ref1]: 厳密には [Rust] で[「参照」は別のものを指す]({{< relref "./references-and-borrowing.md" >}} "参照と借用")ので，ここではユルく「のようなもの」という感じでご容赦を。

{{< fig-img class="lightmode" src="./variable-bindings1.png" link="./variable-bindings1.svg" >}}

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

{{< fig-img class="lightmode" src="./variable-bindings2.png" link="./variable-bindings2.svg" >}}

構造体やコレクションの各型の値はヒープ領域に置かれる。
ヒープ上の値に対して代入等を行うたびにコピーを行うのは高コストだし，かといって野放図に参照を増やすとヒープ管理が煩雑になってしまう。

そこで所有権を「移動」することでヒープ管理の最適化を行うわけだ。

## Copy/Clone Trait

じゃあ，構造体やコレクションのコピーはできないのかというと，ちゃんと救済措置はある。

[Rust] の標準ライブラリには `Clone` および `Copy` トレイトが用意されていてこれらを実装することで値のコピーが可能になる[^trait1]。

[^trait1]: 「トレイト（trait）」とは，ここでは C++ や Java で言うところの interface クラスのようなものだと思っておけばよい。トレイトに関しては多分その内ちゃんとした記事を書くと思う（笑）

たとえば先程の `Person` 構造体に `Clone` トレイトを実装すると

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

のように `clone()` メソッドでコピーを生成できるようになる。
図にすると，こんな感じか。

{{< fig-img class="lightmode" src="./variable-bindings3.png" link="./variable-bindings3.svg" >}}

あるいは `derive` 構文[^derive] を使えば

[^derive]: `derive` はマクロに似た拡張構文の一種。詳細については割愛する。ていうか，これから勉強する（笑）

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

のようにコンパイラが自動的に `Clone` トレイトを実装してくれる。

実を言うと

```rust
#[derive(Copy, Clone)]
```

とすれば代入時に値をコピーできるようになるのだが， `Person` 構造体ではコンパイルエラーになってしまい上手く行かなかった。
色々試してみたが，どうも `String` 等のコレクションを含む構造体の `Copy` トレイトは実装できないようだ[^cpy1]。

[^cpy1]: `derive` 構文を使った `Copy` トレイトの実装はあくまでも bitwise なもので，たとえ `String` 等で `Copy` トレイトを実装してもスタック上のポインタ値等がコピーされるだけなので意味がないそうな。

属性にコレクション型を含まないのであれば

```rust {hl_lines=[1, 12]}
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

てな感じに簡単に `Copy` トレイトを実装できる。

ふむむー。

## ブックマーク

- [RustのコピーセマンティクスをCopyトレイトを実装して確認する ｜ Developers.IO](https://dev.classmethod.jp/server-side/language/rust-copy-semantics-use-copy-trait/)

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
