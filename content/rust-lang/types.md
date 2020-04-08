+++
title = "Rust の型に関する覚え書き"
date =  "2020-04-08T16:33:06+09:00"
description = "思いつくまま脈絡なく随時更新予定。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "type" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事では [Rust] の型（type）に関することを思いつくまま脈絡なく書き記しておく。
随時更新予定。

## 組み込みデータ型

組み込みデータ型には大きく分けてスカラ型と複合型がある。

### **スカラ型**

スカラ型はスカラ値を有する型である。
型によって値のサイズが決まっていて，受け渡しは copy semantics で行われる。

スカラ型に分類される型は以下の通り。

#### 1. 整数型

|  サイズ | 符号あり | 符号なし |         既定         |
| -------:|:--------:|:--------:|:--------------------:|
|  8 bits |   `i8`   |   `u8`   |                      |
| 16 bits |  `i16`   |  `u16`   |                      |
| 32 bits |  `i32`   |  `u32`   | {{< icons "check">}} |
| 64 bits |  `i64`   |  `u64`   |                      |
| &mdash; | `isize`  | `usize`  |                      |

`isize`/`usize` のサイズはアーキテクチャに依存する（C 言語等でいう `int` と同じ）。

リテラル表記時の既定は `i32` 型となる。

```rust
let x = 123; //type i32
```

ただしリテラル表記の後ろに型名を付けることで型を指定することが可能。

```rust
let x = 123i64; //type i64
```

もちろん変数に型注釈を付けて明示することも可能。

```rust
let x: u8 = 123; //type u8
```

#### 2. 浮動小数点数型

|                   サイズ | 型名  |         既定         |
| ------------------------:|:-----:|:--------------------:|
| 32 bits (仮数部 23 bits) | `f32` |                      |
| 64 bits (仮数部 52 bits) | `f64` | {{< icons "check">}} |

型の推定・評価は整数型と同じで，リテラル表記時の既定は `f64` 型となる。

```rust
let x = 1.23; //type f64
```

#### 3. 文字型

文字型 `char` は Unicode 符号点を示す。
リテラル表記はこんな感じ。

```rust
let c = '♡';
```

[Rust] における文字列と文字の関係については「[Rust の文字列操作（1）]({{< relref "./character-string-1.md" >}})」を参照のこと。

#### 4. 論理値型

論理値型 `bool` は `true` と `false` の2値のみ取り得る。

### **複合型**

複合型は複数の要素を組み合わせた型で，スカラ型と配列型がある。

タプル型は（名前の通り）複数の型を組み合わせた型で，こんな感じに記述できる。

```rust {hl_lines=[2]}
fn main() {
	let tup = (500, 6.4, 1u8);
    println!("tup = {:?}", tup); //Output: tup = (500, 6.4, 1)
}
```

配列は，単一の型で構成される複数要素の型で[^enum1]，こんな感じに記述できる。

[^enum1]: 配列型でも列挙型（`enum`）と組み合わせることで複数の型に対応させることは可能。

```rust {hl_lines=[2]}
fn main() {
	let ary = [1, 2, 3];
	println!("ary = {:?}", ary); //Output: ary = [1, 2, 3]
}
```

複合型は定義毎に型と要素数が決まっている。
要素の型が全てスカラ型であれば copy semantics で値の受け渡しが可能。

## 型エイリアス（Type Alias）

型の別名定義。
名前は変わっても機能は変わらない。
総称型と組み合わせると吉？

```rust {hl_lines=[1]}
type Strings = Vec<String>;

fn main() {
    let planets: Strings = vec![
        "Mercury".to_string(),
        "Venus".to_string(),
        "Earth".to_string(),
        "Mars".to_string(),
    ];
    planets.iter().for_each(|p| {
        println!("{}", p);
    });
}
```

## 構造体

- 0個以上のフィールド（型と名前）を含むデータ構造
- 構造体に紐づく関連関数およびメソッドを定義可能
- トレイトからの実現（realization）が可能。構造体間の継承（inheritance）は不可

```rust
struct Planet {
	name: String,
    mass: f64,
    distance: f64,
}
```

### 構造体のインスタンス化

リテラル表現を使ったインスタンス化。

```rust {hl_lines=["9-13"]}
#[derive(Debug)]
struct Planet {
    name: String,
    mass: f64,
    distance: f64,
}

fn main() {
    let p = Planet {
        name: "Earth".to_string(),
        mass: 1.0,
        distance: 1.0,
    };
    println!("{:?}", p); //Output: Planet { name: "Earth", mass: 1.0, distance: 1.0 }
}
```

インスタンス化関数を定義してみる。

```rust {hl_lines=["9-15", 19]}
#[derive(Debug)]
struct Planet {
    name: String,
    mass: f64,
    distance: f64,
}

impl Planet {
    fn new(s: &str, mass: f64, distance: f64) -> Planet {
        Planet {
            name: s.to_string(),
            mass,
            distance,
        }
    }
}

fn main() {
    let p = Planet::new("Earth", 1.0, 1.0);
    println!("{:?}", p); //Output: Planet { name: "Earth", mass: 1.0, distance: 1.0 }
}
```

### 構造体のコピー

コピー用のメソッドを書いてみた。

```rust
impl Planet {
    fn copy(&self) -> Planet {
        Planet {
            name: self.name.clone(),
            ..*self
        }
    }
}
```

あるいは `derive` 構文を使って `Clone` トレイトから `clone()` メソッドを自動生成する。

```rust {hl_lines=[1, 10]}
#[derive(Debug, Clone)]
struct Planet {
    name: String,
    mass: f64,
    distance: f64,
}

fn main() {
    let p = Planet::new("Earth", 1.0, 1.0);
    let cpy = p.clone();
    println!("{:?}", p); //Output: Planet { name: "Earth", mass: 1.0, distance: 1.0 }
    println!("{:?}", cpy); //Output: Planet { name: "Earth", mass: 1.0, distance: 1.0 }
}
```

なお `String` 等の `Copy` トレイトを実装できないフィールドを含む構造体は `Copy` トレイトを実装できないため，代入構文で値のコピーが発生しない。
この場合は move semantics により所有権が移動する。

### タプル構造体

構造体の特殊なパターンで，名前のないフィールドを定義する。
フィールドがタプルで定義される以外は通常の構造体と同じ。
関連関数およびメソッドも定義可能。

```rust
#[derive(Debug, Copy, Clone)]
struct Point(i64, i64);

impl Point {
    fn new(x: i64, y: i64) -> Point {
        Point(x, y)
    }
}

fn main() {
    let p1 = Point(1, 2);
    let p2 = Point::new(3, 4);
    let p3 = p1; //copy semantics
    println!("{:?}", p1); //Output: PPoint(1, 2)
    println!("{:?}", p2); //Output: PPoint(3, 4)
    println!("{:?}", p3); //Output: PPoint(1, 2)
}
```






[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
