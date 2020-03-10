+++
title = "Rust の文字列操作（1）"
date =  "2020-03-08T23:08:47+09:00"
description = " Rust では「文字列」としての操作対象を Unicode/UTF-8 に限定している。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "character", "string", "unicode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて，所有権について何とな〜く分かったところで，目先を変えて文字・文字列・文字列操作について調べていこう。

## [Rust] の「文字列」は Unicode/UTF-8 限定

文字列操作で何が面倒かって複数の文字セットと文字エンコーディングが混在している点である。
そこで今どきのプログラミング言語は「文字列」としての操作対象を Unicode/UTF-8 に限定していることが多い。
[Rust] もそうした言語のひとつである。

### こんにちは，世界！

まずは，以下のコードを書いてみる。

```rust
fn main() {
    println!("こんにちは，世界！");
}
```

ここで `"こんにちは，世界！"` はリテラル文字列で，実行すると

```text
$ cargo run
こんにちは，世界！
```

と表示される。
`println!` マクロ引数のリテラル文字列も出力される文字列も UTF-8 エンコーディングである。

ここでリテラル文字列の `"こんにちは，世界！"` を強引に Shift-JIS に書き換えて実行するとどうなるか。

```text
$ pushd src/
$ cp main.rs main.rs.utf8.txt
$ gonkf conv -s utf8 -d sjis main.rs.utf8.txt > main.rs
$ popd
$ cargo run
error: couldn't read src/main.rs: stream did not contain valid UTF-8
```

というわけでコンパイルに失敗しました（笑）

## [Rust] における文字と文字列

[Rust] では Unicode/UTF-8 文字および文字列を扱うために以下の3つの型が存在する。

| 型名     | 内部表現  | 内容                           |
| -------- | --------- | ------------------------------ |
| `char`   | `u32`     | 文字（Unicode 符号点） |
| `String` | `Vec<u8>` | 可変長文字列（UTF-8）          |
| `str`    | `[u8]`    | 文字列への参照（UTF-8）        |

`char` および `str` は組み込み型， `String` は標準ライブラリで定義される型である。

`str` 型は文字列に対するスライスで，それ自体は所有権を持たず（値を持たないため）範囲指定付きビューとして機能する。
通常は `&str` と参照の形で使う。

`String` と `str` の関係はこんな感じ。

```rust
fn main() {
    let s: String = "Hello World".to_string();
    let hw: &str = s.as_str();
    let w: &str = &s[6..11];
    println!("{}", s); //Output: Hello World
    println!("{}", hw); //Output: Hello World
    println!("{}", w); //Output: World
}
```

図にするとこんな感じだろうか。

{{< fig-img src="./character-string-1.png" link="./character-string-1.png" width="684" >}}

`"Hello World"` などのリテラル文字列は型としては `&'static str` と表現できる。
ちなみに `'static` はライフタイム注釈の特殊な表現で，生存期間がプログラム全域に及ぶという恐ろしいものである（笑）

リテラル文字列の後ろの `to_string()` メソッドは文字列のコピーを返す。
リテラル文字列そのものがセットされているわけではない。

## 文字列から文字を抽出する

[Rust] の文字列の中身は UTF-8 バイト列なので，文字列から文字（厳密には Unicode 符号点[^norm1]）を抽出するのは単純ではない。
そもそも

[^norm1]: Unicode ではひとつの符号点がひとつの文字を表すとは限らない点に注意。これに関して以前に「[Go 言語と Unicode 正規化]({{< ref "/golang/unicode-normalization.md" >}})」という記事を書いたので参考にどうぞ。

```rust
fn main() {
    let s = "日本語";
    println!("{}", s[0]); //Error: the type `str` cannot be indexed by `{integer}`
}
```

とかやってもコンパイル・エラーになるだけである。

文字列から文字を抽出するには，例えばこんな感じに書ける。

```rust
fn main() {
    let s1 = "日本語";
    println!("{:?}", s1.chars().nth(0)); //Output: Some('日')
    let s2 = s1.to_string();
    println!("{:?}", s2.chars().nth(1)); //Output: Some('本')
}
```

まず `chars()` メソッドで文字列を文字単位のイテレータ `Chars` に変換し，更に `nth()` メソッドで任意の位置の文字を抽出している。
`nth()` メソッドは `Option` 型の値を返すので本来は何らかのエラー・ハンドリングが必要だが，今回は横着している[^err1]。
ゴメンペコン。

[^err1]: エラー・ハンドリングについてはいつかどこかでやる予定。

あるいはもっと簡単に `collect()` メソッドを使って `Chars` から配列 `Vec<char>` に変換して

```rust
fn main() {
    let nihongo = "日本語";
    let chs = nihongo.chars().collect::<Vec<char>>();
    println!("{}", chs[0]); //Output: 日
}
```

てな感じにも記述できる。

## 「文字」列から「文字列」を生成する

1文字ではなく，先頭の2文字を抽出して文字列を生成してみる。
まずは素朴にこんな感じだろうか。

```rust
fn main() {
    let nihongo = "日本語";
    let mut nippon = String::new();
    for (i, c) in nihongo.chars().enumerate() {
        nippon.push(c);
        if i > 0 {
            break;
        }
    }
    println!("{}", nippon); //Output: 日本
}
```

うーん，微妙。
たとえば `take()` メソッドを使えば先頭から指定した数だけ要素を有効にできる。

```rust {hl_lines= ["4-6"]}
fn main() {
    let nihongo = "日本語";
    let mut nippon = String::new();
    for c in nihongo.chars().take(2) {
        nippon.push(c);
    }
    println!("{}", nippon); //Output: 日本
}
```

いや，そもそも `for` 文で回すってのがダサいというか... ふむむー，そっか。
`fold()` メソッドを使えばいいのか。

```rust {hl_lines= ["3-6"]}
fn main() {
    let nihongo = "日本語";
    let nippon = nihongo.chars().take(2).fold(String::new(), |mut s, c| {
        s.push(c);
        s
    });
    println!("{}", nippon); //Output: 日本
}
```

でも「先頭の2文字を抽出する」だけなのに，コードが大袈裟すぎないか？ と思ってたら，ゴメン。
`collect()` メソッドを使えば一発で `String` に変換できたわ。

```rust {hl_lines= [3]}
fn main() {
    let nihongo = "日本語";
    let nippon = nihongo.chars().take(2).collect::<String>();
    println!("{}", nippon); //Output: 日本
}
```

なんだ，簡単ぢゃん `orz`

## ブックマーク

- [Stringとstr入門 - Qiita](https://qiita.com/Mizunashi_Mana/items/db88cb0bff002abce1ae)
- [文字列操作 with Rust - Qiita](https://qiita.com/hobo/items/04846eeccb5e2004731a)
- [Rustで文字列の先頭文字や部分文字列を取得する - Qiita](https://qiita.com/7ma7X/items/7fb68395984958987a54)

- [spiegel-im-spiegel/charset_document: 「文字コードとその実装」 upLaTeX ドキュメント](https://github.com/spiegel-im-spiegel/charset_document) : 大昔に書いたドキュメントで内容が古いため，取り扱いには注意
- [spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang](https://github.com/spiegel-im-spiegel/text) : `gonkf` コマンドはここで公開しています

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
{{% review-paapi "4873118557" %}} <!-- プログラミングRust -->
