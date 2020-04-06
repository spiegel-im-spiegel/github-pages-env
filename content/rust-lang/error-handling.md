+++
title = "エラー・ハンドリングのキホン"
date =  "2020-04-06T21:11:43+09:00"
description = "やっぱ例外処理は要らんよね（笑）"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "error", "enumeration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今日の起点となるコードはこれにしよう。

```rust
fn parse_string(s: &str) -> Result<u32, std::num::ParseIntError> {
    s.parse::<u32>()
}

fn main() {
    println!("{:?}", parse_string("-1")); //Output: Err(ParseIntError { kind: InvalidDigit })
}
```

`parse_string()` 関数の返り値の型 `Result` は列挙型（enum）で，以下の2つの列挙子（variant）で構成されている。

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

`Result` 型の評価には `match` 式が使える。
たとえば `parse()` に失敗した際に `0` をセットしたいなら

```rust
fn main() {
    let n = match parse_string("-1") {
        Ok(x) => x,
        Err(_) => 0,
    };
    println!("{}", n); //Output: 0
}
```

などと書くことができる。

## Result 型を使ったエラー・ハンドリング

`Result` 型を使うことで基本的なエラー・ハンドリングが可能になる。
以降で例を挙げていこう。

### Panic を投げる

`parse_string()` 関数が `Err` を返した場合に強制終了したいのであれば

```rust
fn main() {
    let n = match parse_string("-1") {
        Ok(x) => x,
        Err(e) => panic!(e), //Output: thread 'main' panicked at 'Box<Any>', src/main.rs:8:19
    };
    println!("{}", n); //do not reach
}
```

と `panic!` マクロで panic を投げればよい。
ちなみに，環境変数 `RUST_BACKTRACE` に `1` をセットすると panic 時にスタックトレース情報も吐く。

単に panic を投げればいいのであれば `Result::unwrap()` メソッドを使えば上述のコードとほぼ同じ結果が得られる。

```rust
fn main() {
    let n = parse_string("-1").unwrap(); //Output: thread 'main' panicked at 'called `Result::unwrap()` on an `Err` value: ParseIntError { kind: InvalidDigit }', src/main.rs:6:13
    println!("{}", n); //do not reach
}
```

Panic を投げる際のメッセージを指定するには `Result::expect()` メソッドを使う。

```rust
fn main() {
    let n = parse_string("-1").expect("error in parse_string() function"); //Output: thread 'main' panicked at 'error in parse_string() function: ParseIntError { kind: InvalidDigit }', src/main.rs:6:13
    println!("{}", n); //do not reach
}
```

### Panic 以外のハンドリング

エラー時に単に panic を投げるのではなく，何らかの処理を行って普通にプロセスを終了したいのであれば `Result::unwrap_or_else()` メソッドを使って

```rust
fn main() {
    let n = parse_string("-1").unwrap_or_else(|e| {
        println!("Error in parse_string() function: {:?}", e); //Output: Error in parse_string() function: ParseIntError { kind: InvalidDigit }
        std::process::exit(1);
    });
    println!("{}", n); //do not reach
}
```

などとすることもできる。
プロセスを終了するのではなく，最初の例のように解析失敗時に `0` をセットしたいなら

```rust
fn main() {
    let n = parse_string("-1").unwrap_or_else(|_| 0);
    println!("{}", n); //Output: 0
}
```

てな感じにも書ける。
あるいはもっと簡単に `Result::unwrap_or()` メソッドを使って

```rust
fn main() {
    let n = parse_string("-1").unwrap_or(0);
    println!("{}", n); //Output: 0
}
```

とも書ける。

### エラーの委譲

次に以下の関数を考える。

```rust
fn parse_pair_strings(s1: &str, s2: &str) -> Result<(u32, u32), std::num::ParseIntError> {
    let x = match s1.parse::<u32>() {
        Ok(n) => n,
        Err(e) => return Err(e),
    };
    let y = match s2.parse::<u32>() {
        Ok(n) => n,
        Err(e) => return Err(e),
    };
    Ok((x, y))
}

fn main() {
    println!("{:?}", parse_pair_strings("1", "-1")); //Output: Err(ParseIntError { kind: InvalidDigit })
}
```

`parse_pair_strings()` の仮引数 `s1`, および `s2` に対してそれぞれ `parse()` を行うのだが，結果がエラーの際には， `Err` をそのまま返して関数元にハンドリングを委譲している。

関数の返り値が `Result` 型なら， `?` 演算子を使って，エラーの委譲をもっと簡単に書くことができる。
こんな感じ。

```rust {hl_lines=[2]}
fn parse_pair_strings(s1: &str, s2: &str) -> Result<(u32, u32), std::num::ParseIntError> {
    Ok((s1.parse::<u32>()?, s2.parse::<u32>()?))
}
```

### エラーの汎化

今度は，引数に文字列を渡すのではなく，標準入力から文字列を取得して `parse()` してみよう。

```rust {hl_lines=[3]}
fn parse_from_stdin() -> Result<u32, std::num::ParseIntError> {
    let mut buf = String::new();
    std::io::stdin().read_line(&mut buf)?; //Compile error: `?` couldn't convert the error to `std::num::ParseIntError`
    buf.trim().parse::<u32>()
}
```

当然ながらこれはコンパイルエラーになる。
`read_line()` 関数がエラーの際に吐く型は `std::io::Error` なので `std::num::ParseIntError` 型とはマッチしないためだ。

解決するには，これらの型の汎化である `std::error::Error` 型を使えばよい。

{{< fig-img src="./error-trait.png" link="./error-trait.puml" width="1059" >}}

標準ライブラリで定義される各種エラー型はどれも `std::error::Error` 型からの特化である。

`std::error::Error` 型を使う際は `Box<dyn Trait>` を使うようだ。
こんな感じ。

```rust {hl_lines=[1]}
fn parse_from_stdin() -> Result<u32, Box<dyn std::error::Error>> {
    let mut buf = String::new();
    std::io::stdin().read_line(&mut buf)?;
    let n = buf.trim().parse::<u32>()?;
    Ok(n)
}

fn main() {
    println!("{:?}", parse_from_stdin());
}
```

実行するとこんな感じになる。

```text
$ echo -1 | cargo run
Err(ParseIntError { kind: InvalidDigit })
```

`main()` 関数側でもう少し細かくエラーを見てみよう。
こんな感じかなぁ。

```rust
fn main() {
    let n = match parse_from_stdin() {
        Ok(x) => x,
        Err(err) => {
            match err.downcast_ref::<std::io::Error>() {
                Some(e) => {
                    println!("io::Error in parse_from_stdin(): {}", e);
                    std::process::exit(1);
                }
                _ => (),
            };
            match err.downcast_ref::<std::num::ParseIntError>() {
                Some(e) => {
                    println!("ParseIntError in parse_from_stdin(): {}", e);
                    std::process::exit(1);
                }
                _ => (),
            };
            println!("Other error in parse_from_stdin(): {}", err);
            std::process::exit(1);
        }
    };
    println!("{}", n);
}
```

実行するとこんな感じになる。

```text
$ echo -1 | cargo run
ParseIntError in parse_from_stdin(): invalid digit found in string
```

## Option 型を使ったエラー・ハンドリング

上述のコードにある `downcast_ref()` メソッドは `Option` 型の値を返す。
`Option` 型も列挙型で，以下の2つの列挙子で構成されている。

```rust
enum Option<T> {
    Some(T),
    None,
}
```

`None` はいわゆる Null 値，つまり値がないことを示す。
関数の実行結果が Null 値を取り得る場合に `Option` 型を返すことで，呼び出し元に Null 値のハンドリングを促すことができる。

起点のコードはこんな感じでどうだろうか。

```rust
fn main() {
    let nihongo = "日本語";
    for i in 0..4 {
        println!("{}: {:?}", i, nihongo.chars().nth(i));
    }
}
```

実行するとこんな感じになる。

```text
$ cargo run
0: Some('日')
1: Some('本')
2: Some('語')
3: None
```

### Panic を投げる

結果が `None` なら panic を投げるようにしてみる。

```rust
fn main() {
    let nihongo = "日本語";
    for i in 0..4 {
        let ch = match nihongo.chars().nth(i) {
            Some(c) => c,
            None => panic!("Out of bounds"),
        };
        println!("{}: {}", i, ch)
    }
}
```

実行するとこんな感じ。

```text
$ cargo run
0: 日
1: 本
2: 語
thread 'main' panicked at 'Out of bounds', src/main.rs:6:21
```

また `Option` 型にも `unwrap()` や `expect()` といったメソッドがあり， `None` が返ったら panic を投げる。

```rust
fn main() {
    let nihongo = "日本語";
    for i in 0..4 {
        println!("{}: {}", i, nihongo.chars().nth(i).expect("Out of bounds"));
    }
}
```

```text
$ cargo run
0: 日
1: 本
2: 語
thread 'main' panicked at 'Out of bounds', src/main.rs:4:31
```

### Panic 以外のハンドリング

たとえば，こんな感じかな。

```rust {hl_lines=[6]}
fn main() {
    let nihongo = "日本語";
    for i in 0..4 {
        let ch = match nihongo.chars().nth(i) {
            Some(c) => c,
            None => break,
        };
        println!("{}: {}", i, ch)
    }
}
```

```text
$ cargo run
0: 日
1: 本
2: 語
```

まぁイテレータやコレクションでこんな書き方する人はおらんじゃろうけど。
あンまりいい例示じゃなくてすまん。

## やっぱ例外処理は要らんよね

[Rust] の列挙型は（C/C++ や Java などと異なり）型の列挙を行い，パターン・マッチングによって処理を振り分ける。
この仕組みを上手く使ってエラー・ハンドリングを行っているわけだ。

こうやってみると，やっぱ例外処理は要らんよね（笑）

## ブックマーク

- [Rustのエラーハンドリングの基本 - hideoka’s blog](https://hideoka.hateblo.jp/entry/2019/11/17/224953)
- [Rustで複数のimpl Traitを返す - Qiita](https://qiita.com/taiki-e/items/39688f6c86b919988222)

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
