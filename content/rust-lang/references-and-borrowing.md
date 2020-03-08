+++
title = "参照と借用"
date =  "2020-03-08T00:51:00+09:00"
description = "ちょっとした数理パズルだと思えばいい（笑）"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "ownership", "reference", "lifetime" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./variable-bindings.md" >}} "変数束縛")の続きで，所有権に絡む話。

## 参照と借用

まずはこのコードを出発点にしよう。

```rust
struct Person {
    age: u32,
    name: String,
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    println!("{} ({})", p1.name1, p.age); //Output: Alice (24)
}
```

毎回 `println!` マクロで出力を整形するのはかったるいので，整形を行う関数を考えてみる。
こんな感じ。

```rust {hl_lines= ["6-12"]}
struct Person {
    age: u32,
    name: String,
}

fn display_person(p: Person) -> String {
    let mut s = p.name;
    s.push_str(" (");
    s.push_str(&p.age.to_string());
    s.push_str(")");
    s
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    println!("{}", display_person(p1)); //Output: Alice (24)
}
```

一見うまく行っているようだが，次の1行を足すとコンパイル・エラーになる。

```rust {hl_lines=["7"]}
fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    println!("{}", display_person(p1));
    println!("{}", p1.name); //Error: borrow of moved value: `p1`
}
```

これは 変数の `p1` の値の所有権が `display_person()` 関数の引数 `p` に移動したからである。

変数を引数にセットするたびに所有権が移るのは鬱陶しいので，変数の「参照」を引数にセットする。
こんな感じに書き換えてみよう。

```rust {hl_lines=["1-2"]}
fn display_person(p: &Person) -> String {
    let mut s = p.name.clone();
    s.push_str(" (");
    s.push_str(&p.age.to_string());
    s.push_str(")");
    s
}
```

2行目の宣言文も変わってることに注意。
実は2行目を

```rust {hl_lines=[2]}
fn display_person(p: &Person) -> String {
    let mut s = p.name; //Error: cannot move out of `p.name` which is behind a shared reference
    s.push_str(" (");
    s.push_str(&p.age.to_string());
    s.push_str(")");
    s
}
```

のままにするとコンパイル・エラーになる。
引数 `p` は元の変数の値を「借用」しているに過ぎないので，所有権の移動はできないのである。
ていうか，最初の `display_person()` 関数は引数の値を壊してたのか。
コワイコワイ（笑）

全体のコードは以下の通り。

```rust
struct Person {
    age: u32,
    name: String,
}

fn display_person(p: &Person) -> String {
    let mut s = p.name.clone();
    s.push_str(" (");
    s.push_str(&p.age.to_string());
    s.push_str(")");
    s
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    println!("{}", display_person(&p1)); //Output: Alice (24)
    println!("{}", p1.name); //Output: Alice
}
```

変数と参照の関係を図式化するとこんな感じだろうか。

{{< fig-img src="./references-and-borrowing.png" link="./references-and-borrowing.png" >}}

`display_person()` 関数の引数である参照変数 `p` は，値を直接参照しているのではなく，値の所有権を持つ変数 `p1` を参照している点がポイントである。
このことを示す例を挙げてみよう。

### 参照されている変数の値の所有権は移動できない

先程の `main()` 関数を少し弄って

```rust {hl_lines=["6-7"]}
fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    let p2 = &p1;
    let p3 = p1; //Error: cannot move out of `p1` because it is borrowed
    println!("{}", display_person(p2));
    println!("{}", p3.name);
}
```

とするとコンパイル・エラーになる。
参照されている変数の値の所有権は移動できないようだ。
まぁ，当たり前か。

ただし `println!` マクロの位置を少し変えると

```rust {hl_lines=[7]}
fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    let p2 = &p1;
    println!("{}", display_person(p2)); //Output: Alice (24)
    let p3 = p1;
    println!("{}", p3.name); //Output: Alice
}
```

問題なく動く。
これは参照変数 `p2` のライフタイム（lifetime）が `display_person()` 関数実行完了をもって満了しているため。
その後の `let p3 = p1;` 宣言文では `p1` を参照している変数はないので問題なく所有権を移動できる。

### 変数のライフタイムを超えた参照はできない

ちょっとへんてこりんなコードだが

```rust
fn main() {
    let long_p;
    {
        let p1 = Person {
            age: 24,
            name: "Alice".to_string(),
        };
        long_p = &p1; //Error: `p1` does not live long enough
        println!("{}", p1.name);
    }
    println!("{}", display_person(long_p));
}
```

これもコンパイル・エラーになる。

変数 `p1` のライフタイムは

```rust {hl_lines=["4-9"]}
fn main() {
    let long_p;
    {
        let p1 = Person {
            age: 24,
            name: "Alice".to_string(),
        };
        long_p = &p1; //Error: `p1` does not live long enough
        println!("{}", p1.name);
    }
    println!("{}", display_person(long_p));
}
```

このスコープだが，参照変数 `long_p` のライフタイムは `main()` 関数全域に及ぶため，最後の

```rust {hl_lines=[11]}
fn main() {
    let long_p;
    {
        let p1 = Person {
            age: 24,
            name: "Alice".to_string(),
        };
        long_p = &p1; //Error: `p1` does not live long enough
        println!("{}", p1.name);
    }
    println!("{}", display_person(long_p));
}
```

`println!` マクロ実行時で，元の変数 `p1` のライフタイムは満了しているのに参照変数 `long_p` はまだ生きているという不整合が発生しているわけだ。

コンパイル・エラーのメッセージだけ見ても分かりにくいかもしれないが，変数のライフタイムを頭に入れながら考えれば分かりやすいだろう。
ちょっとした数理パズルだと思えばいい（笑）

## ライフタイム注釈

ここでもうひとつ関数を考えてみる。

欲しいのは2つの `Person` インスタンスの `age` が小さい方を選択する関数である。
とりあえず何も考えずに書いてみよう。

```rust
fn younger(l: &Person, r: &Person) -> &Person { //Error: missing lifetime specifier
    if l.age < r.age {
        l
    } else {
        r
    }
}
```

残念。
コンパイルエラーになってしまった。

この関数では参照変数 `l` と `r` のうちどちらかを返すが，どちらを返すかは実行時にしか分からない。
したがって，以下の `main()` 関数内で

```rust {hl_lines=[10]}
fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    let p2 = Person {
        age: 18,
        name: "Bob".to_string(),
    };
    let p3 = younger(&p1, &p2);
    println!("{}", display_person(p3));
}
```

参照変数 `p3` が参照しているのが `p1` なのか `p2` なのか（コンパイル時に）決定できないためライフタイムも決まらないのである。

こういうときは `younger()` 関数に「ライフタイム注釈（lifetime annotation）」を付けるとよい。
こんな感じ。

```rust {hl_lines=[1]}
fn younger<'a>(l: &'a Person, r: &'a Person) -> &'a Person {
    if l.age < r.age {
        l
    } else {
        r
    }
}
```

ここでは `'a` がライフタイム注釈に相当する。
アポストロフィ（`'`）から続く文字列で構成されている。
新しい `younger()` 関数では引数と返り値の `&Person` 参照変数が同じライフタイムであることをコンパイラに知らせている。

全体のコードはこんな感じ。

```rust
struct Person {
    age: u32,
    name: String,
}

fn display_person(p: &Person) -> String {
    let mut s = p.name.clone();
    s.push_str(" (");
    s.push_str(&p.age.to_string());
    s.push_str(")");
    s
}

fn younger<'a>(l: &'a Person, r: &'a Person) -> &'a Person {
    if l.age < r.age {
        l
    } else {
        r
    }
}

fn main() {
    let p1 = Person {
        age: 24,
        name: "Alice".to_string(),
    };
    let p2 = Person {
        age: 18,
        name: "Bob".to_string(),
    };
    println!("{}", display_person(younger(&p1, &p2))); //Output: Bob (18)
}
```

よーし，うむうむ，よーし。






[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
