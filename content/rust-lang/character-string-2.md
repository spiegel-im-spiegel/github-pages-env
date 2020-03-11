+++
title = "Rust の文字列操作（2）"
date =  "2020-03-11T20:24:35+09:00"
description = "文字列と他の型の値との相互変換について。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "string" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./character-string-1.md" >}} "Rust の文字列操作（1）")の続き。
[Rust] の文字列操作についてオベンキョ中です。

## 他の型の値を文字列に変換する

ある型が `ToString` トレイトを実装していれば `to_string()` メソッドを使って文字列に変換できる。
整数や浮動小数点数などの基本型は `ToString` トレイトを実装している。
以下のようにリテラル表現に対して直接 `to_string()` メソッドを起動することもできる。

```rust
fn main() {
    println!("{}", (123).to_string()); //Output: 123
    println!("{}", (1.23).to_string()); //Output: 1.23
    println!("{}", (true).to_string()); //Output: true
    println!("{}", ('あ').to_string()); //Output: あ
}
```

同じ基本型でもタプルや配列は `ToString` トレイトを実装していないため `to_string()` メソッドは使えない。

```rust
fn main() {
    println!("{}", ([1, 2, 3]).to_string()); //Error: no method named `to_string` found for type `[{integer}; 3]` in the current scope
}
```

### 変換する文字列を整形する

今までさんざん出てきたので今更だが， `format!` 等のマクロを使えば他の型の値を変換する際の文字列を整形できる。

| マクロ名             | 機能                                         |
| -------------------- | -------------------------------------------- |
| `format!`            | 整形文字列を `String` として出力する         |
| `print!`, `println!` | 整形文字列を標準出力に出力する               |
| `write!`, `writeln!` | 整形文字列を指定した出力ストリームに出力する |
| `panic!`             | 整形文字列を panic 出力に付加する            |

たとえば，こんな感じ。

```rust
fn main() {
    println!("{{{} = {num:#04x} = {num:#010b}}}", num = 15); //Output: {15 = 0x0f = 0b00001111}
}
```

プレースホルダ `{which:how}` の中身の `how` の部分が C/C++ や Go の `printf` 文の書式（`%d` など）に相当する。
意味もだいたい同じ。

| 書式 | 意味                               |
|:----:| ---------------------------------- |
| なし | 既定のテキスト                     |
| `b`  | 2進数                              |
| `o`  | 8進数                              |
| `x`  | 16進数（小文字）                   |
| `X`  | 16進数（大文字）                   |
| `e`  | 浮動小数点数科学技術記法（小文字） |
| `E`  | 浮動小数点数科学技術記法（大文字） |
| `?`  | デバッグ出力                       |
| `p`  | ポインタ値                         |

実はこれらの書式には対応するトレイトがあって，独自に作成した型でもトレイトを実装することで書式に対応できる。

| 書式 | トレイト             |
|:----:| -------------------- |
| なし | `std::fmt::Display`  |
| `b`  | `std::fmt::Binary`   |
| `o`  | `std::fmt::Octal`    |
| `x`  | `std::fmt::LowerHex` |
| `X`  | `std::fmt::UpperHex` |
| `e`  | `std::fmt::LowerExp` |
| `E`  | `std::fmt::UpperExp` |
| `?`  | `std::fmt::Debug`    |
| `p`  | `std::fmt::Pointer`  |

たとえば，以下の構造体 `Person` に対して

```rust
struct Person {
    age: u32,
    name: String,
}
```

`std::fmt::Display` トレイトを実装してみる。
こんな感じ。

```rust {hl_lines= [1, "8-12", "19-20"]}
use std::fmt;

struct Person {
    age: u32,
    name: String,
}

impl fmt::Display for Person {
    fn fmt(&self, dest: &mut fmt::Formatter) -> fmt::Result {
        write!(dest, "{} ({})", self.name, self.age)
    }
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "alice".to_string(),
    };
    println!("p1 = {}", p1); //Output: p1 = alice (24)
    println!("p1 = {}", p1.to_string()); //Output: p1 = alice (24)
}
```

`std::fmt::Display` トレイトを実装すると `ToString` トレイトもセットで実装される。
便利！

`std::fmt::Debug` トレイトの実装はもっと簡単で， `derive` 構文を使って

```rust {hl_lines= [1, 12]}
#[derive(Debug)]
struct Person {
    age: u32,
    name: String,
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "alice".to_string(),
    };
    println!("p1 = {:?}", p1); //Output: p1 = Person { age: 24, name: "alice" }
}
```

とすればよい。

ちなみに全ての組み込み型と標準ライブラリで定義される型は `std::fmt::Debug` トレイトを実装済みなので，さきほど `to_string()` メソッドが使えないと書いたタプルや配列でも

```rust
fn main() {
    println!("{:?}", (123, 1.13)); //Output: (123, 1.13)
    println!("{:?}", [1, 2, 3]); //Output: [1, 2, 3]
}
```

のようにできる。

## 文字列から他の型の値を取得する

`std::str::FromStr` トレイトを実装している型であれば文字列からその型の値を取得できる。

```rust
pub trait FromStr: Sized {
    type Err;
    fn from_str(s: &str) -> Result<Self, Self::Err>;
}
```

整数や浮動小数点数などの基本型は `std::str::FromStr` トレイトを実装している。
返り値は `Result` 型なので何らかのエラーハンドリングが必要だが，以下のコードでは端折っている。
ゴメンペコン。

```rust
use std::str::FromStr;

fn main() {
    println!("{:?}", i64::from_str("123")); //Output: Ok(123)
    println!("{:?}", f64::from_str("1.23")); //Output: Ok(1.23)
    println!("{:?}", bool::from_str("true")); //Output: Ok(true)
}
```

ちなみに大文字の `"TRUE"` はダメらしい（笑）

```rust
use std::str::FromStr;

fn main() {
    println!("{:?}", bool::from_str("TRUE")); //Output: Err(ParseBoolError { _priv: () })
}
```

小文字に変換すればいいかな。

```rust
use std::str::FromStr;

fn main() {
    println!("{:?}", bool::from_str(&"TRUE".to_lowercase())); //Output: Ok(true)
}
```

文字列に対して `parse()` メソッドを使う方法もある。
`parse()` メソッド内部で `from_str()` を呼び出しているらしい。

```rust
fn main() {
    println!("{:?}", "123".parse::<i64>()); //Output: Ok(123)
    println!("{:?}", "TRUE".to_lowercase().parse::<bool>()); //Output: Ok(true)
}
```

記述としてはどっちもどっちだな。

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
