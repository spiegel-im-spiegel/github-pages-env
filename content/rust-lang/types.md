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

1. [組み込みデータ型]({{< relref "#data-types" >}})
1. [型エイリアス]({{< relref "#type-alias" >}})
1. [構造体]({{< relref "#structures" >}})
1. [列挙型]({{< relref "#enumeration" >}})
1. [総称型]({{< relref "#generics" >}})

## 組み込みデータ型 {#data-types}

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

`isize`/`usize` のサイズはアーキテクチャに依存する[^int1]。
また，リテラル表記時等での型推測の既定は `i32` 型となる[^int2]。

[^int1]: `isize`/`usize` は C 言語等でいう `int` と同じと思って構わない。まぁ，どちらかと言うと `size_t` みたいな感じかもしれないけど。
[^int2]: 現在のアーキテクチャの主流が 64 bits にも関わらず整数型の既定が `i32` なのは， 32 bits アクセスのほうが「速い」かららしい。

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

型の推定・評価は整数型と同じで，リテラル表記時等での型推測の既定は `f64` 型となる[^float1]。

[^float1]: 浮動小数点数型の場合，単精度（`f32`）と倍精度（`f64`）の間でパフォーマンス上の違いは殆どないそうだ。それならサイズが大きい方を既定にすればいいよね，ってことらしい。

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

複合型は複数の要素を組み合わせた型で，タプル型と配列型がある。

複合型は定義毎に型と要素数が決まっている。
要素の型が全てスカラ型であれば copy semantics で値の受け渡しが可能。

#### 1. タプル型

タプル型は（名前の通り）複数の型を組み合わせた型で，こんな感じに記述できる。

```rust {hl_lines=[2]}
fn main() {
	let tup = (500, 6.4, 1u8);
    println!("tup = {:?}", tup); //Output: tup = (500, 6.4, 1)
}
```

#### 2. 配列型

配列型は，単一の型で構成される複数要素の型で[^enum1]，こんな感じに記述できる。

[^enum1]: 配列型でも列挙型（`enum`）と組み合わせることで複数の型に対応させることは可能。

```rust {hl_lines=[2]}
fn main() {
	let ary = [1, 2, 3];
	println!("ary = {:?}", ary); //Output: ary = [1, 2, 3]
}
```

## 型エイリアス {#type-alias}

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

## 構造体 {#structures}

- 0個以上のフィールド（型と名前）を含むデータ構造
- 構造体に紐づく関連関数およびメソッドを実装可能
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

インスタンス化関数を実装してみる。

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
関連関数およびメソッドも実装可能。

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

## 列挙型 {#enumeration}

[Rust] の列挙型 `enum` はどちらかというと関数型プログラミング言語の影響を強く受けているらしい。

```rust
enum VariantType {
    Int(i32),
    Float(f64),
    Text(String),
}
```

列挙型の評価には `match` 式を用いる。

```rust {hl_lines=["9-13"]}
fn main() {
    let list: Vec<VariantType> = vec![
        VariantType::Int(123),
        VariantType::Float(1.23),
        VariantType::Text("hello".to_string()),
    ];

    list.iter().for_each(|e| {
        match e {
            VariantType::Int(v) => println!("{:#x}", v), //Output: 0x7b
            VariantType::Float(v) => println!("{:e}", v), //Output: 1.23e0
            VariantType::Text(v) => println!("{:?}", v), //Output: "hello"
        };
    });
}
```

列挙型も構造体と同じく関連関数やメソッドを実装できる。

```rust {hl_lines=["2-7", 18]}
impl VariantType {
    fn is_integer(&self) -> bool {
        match self {
            VariantType::Int(_) => true,
            _ => false, //others
        }
    }
}

fn main() {
    let list = vec![
        VariantType::Int(123),
        VariantType::Float(1.23),
        VariantType::Text("hello".to_string()),
    ];

    list.iter().for_each(|e| {
        println!("Is {:?} an integer?: {}", e, e.is_integer());
    });
}
```

## 総称型 {#generics}

[Rust] ではオブジェクトの抽象化の手段として総称型をサポートしている。

### 関数における総称型

以下の `max()` 関数で使われる型 `T` が総称型と呼ばれているもの。

```rust
fn max<T: std::cmp::PartialOrd>(left: T, right: T) -> T {
    if left >= right {
        left
    } else {
        right
    }
}

fn main() {
    let x = 1;
    let y = 2;
    println!("max({}, {}) = {}", x, y, max(x, y)); //Output: max(1, 2) = 2
}
```

`<T: std::cmp::PartialOrd>` の表現は型 `T` の制約条件を示すもので `std::cmp::PartialOrd` トレイトを実装する型のみ `max()` 関数が有効になる。
このようなトレイトを使った制約の定義を「トレイト境界（trait bound）」と呼ぶらしい。

トレイト境界は `where` 節を使って記述することもできる。

```rust {hl_lines=["2-3"]}
fn max<T>(left: T, right: T) -> T
where
    T: std::cmp::PartialOrd,
{
    if left >= right {
        left
    } else {
        right
    }
}
```

トレイト境界の記述は煩雑になりがちなので，これで多少はスッキリするだろう。

ちなみに `std::cmp::PartialOrd` トレイトは `>` `>=`, `<`, `<=` の順序比較演算子を使えるためのもので，少なくとも組み込みデータ型は全て `std::cmp::PartialOrd` トレイトを実装している。

たとえば比較可能でない構造体

```rust
#[derive(Debug)]
struct Person {
    age: u32,
    name: String,
}
```

に `max()` 関数を使おうとしても

```rust {hl_lines=[18]}
fn max<'a, T: std::cmp::PartialOrd>(left: &'a T, right: &'a T) -> &'a T {
    if left >= right {
        left
    } else {
        right
    }
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    let p2 = Person {
        age: 24,
        name: "Bob".to_string(),
    };
    println!("max({:?}, {:?}) = {:?}", p1, p2, max(&p1, &p2)); //Error: can't compare `Person` with `Person`
}
```

コンパイルエラーになる。

### 構造体における総称型

構造体のフィールドやメソッドも総称型で記述することができる。

```rust
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn new(x: T, y: T) -> Self {
        Self { x, y }
    }
}

impl<T: std::ops::Add<Output = T> + Copy> Point<T> {
    fn add(&self, p: &Self) -> Self {
        Self::new(self.x + p.x, self.y + p.y)
    }
}

impl<T: std::fmt::Display> std::fmt::Display for Point<T> {
    fn fmt(&self, dest: &mut std::fmt::Formatter) -> std::fmt::Result {
        write!(dest, "<{}, {}>", self.x, self.y)
    }
}

fn main() {
    let p1 = Point::new(1, 2);
    let p2 = Point::new(3, 4);
    println!("{} + {} = {}", p1, p2, p1.add(&p2)); //Output: <1, 2>
}
```

上述のように `impl` 構文毎に個別にトレイト境界を設定できる。

### 列挙型における総称型

列挙型も構造体と同じように総称型が使える。
列挙型と総称型の組み合わせでもっとも有名なのが `Result` 型だろう。

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```





























[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
