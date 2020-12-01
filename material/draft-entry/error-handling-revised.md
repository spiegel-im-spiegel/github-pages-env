# 【改訂版】 Go 言語のエラー・ハンドリング

## もはや「例外」は Legacy

今年（2020年）になって Rust の勉強を少しだけ始めたが，改めて分かった。

**もはや「例外（Exception）」は Legacy だ！**

たとえば Rust は[列挙型と match 式を組み合わせてエラーの抽出と評価を行う](https://text.baldanders.info/rust-lang/error-handling/ "エラー・ハンドリングのキホン")ことでエラー・ハンドリングを実装できる。

```rust
fn main() {
    let n = match parse_string("-1") {
        Ok(x) => x,
        Err(e) => panic!(e), //Output: thread 'main' panicked at 'Box<Any>', src/main.rs:8:19
    };
    println!("{}", n); //do not reach
}
```

実にスマート！

### 例外の問題は “goto” と同じ

「例外」では，あるオブジェクトに関する記述が少なくとも2つ（try と catch）下手をすると3つ以上のスコープに分割されてしまう。
しかもオブジェクトの状態ごと脱出するため，その状態（の可能性）の後始末をスコープ間で矛盾なく記述しきらなければならない。

この一連に不備があれば，バグやリークやその他の脆弱性のもとになる。
考えるだけで面倒である。

## Go におけるエラー・ハンドリングのキホン

[Go] のエラー・ハンドリングはとにかく「シンプル」の一言に尽きる。

### error 型

まずエラーを扱う組込み interface 型の `error` は以下のように定義される。

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
    Error() string
}
```

つまり，文字列を返す `Error()` メソッドを持つ型であれば全て `error` 型として扱うことができる。

[Go] ではエラーを普通に関数の返値として返す。

```go
file, err := os.Open(filename)
```

他に返すべき値があれば組（tuple）にして最後の要素に `error` 型のインスタンスを配置するのが慣例らしい。
検出したエラーはその場で評価してしまえばよい。

```go
file, err := os.Open(filename)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return false
}
```

If 構文は内部に構文を含めることができるので

```go
if err := file.Close(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return false
}
```

なんてな感じに書くこともできる。
