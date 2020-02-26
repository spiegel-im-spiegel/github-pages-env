+++
title = "みんな大好き Hello World"
date =  "2020-02-26T00:00:15+09:00"
description = "というわけで， Rust の勉強を始めることにしました。ゆっくりのんびりとね。"
image = "/images/attention/rustacean-flat-gesture.png"
tags = [ "programming", "rust", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[というわけで]({{< ref "/remark/2020/02/hello-rust-world.md" >}} "Rust 勉強会，面白かった！")， [Rust] の勉強を始めることにしました。
どぞよろしく。
まぁ，仕事じゃないので，ゆっくりのんびりとね。

余談だが，英語不得手な私は “[Rust]” を「ルスト」と読んでいた。
マジンガーZ[^z1] か！

[^z1]: マジンガーZの「ルスト・ハリケーン」は “rust hurricane”  という説があるらしい。しかも “rust” を「ラスト」と読むと “last” と勘違いするからわざと「ルスト」と読んでいるという説まであるそうな。まぁ，真実は命名者に聞かないと分からないが（笑）

## [Rust] ツールチェーンのインストール

まずは [Rust] ツールチェーンのインストールから。

インストールの方法は[いくつかある](https://forge.rust-lang.org/infra/other-installation-methods.html "Other Installation Methods - Rust Forge")が， [Ubuntu] 等の Linux プラットフォームであれば `rustup-init.sh` をダウンロードして実行するのが手っ取り早い。

```text
$ curl https://sh.rustup.rs -sSf | sh
info: downloading installer

Welcome to Rust!

This will download and install the official compiler for the Rust
programming language, and its package manager, Cargo.

It will add the cargo, rustc, rustup and other commands to
Cargo's bin directory, located at:

  /home/username/.cargo/bin

This can be modified with the CARGO_HOME environment variable.

Rustup metadata and toolchains will be installed into the Rustup
home directory, located at:

  /home/username/.rustup

This can be modified with the RUSTUP_HOME environment variable.

This path will then be added to your PATH environment variable by
modifying the profile file located at:

  /home/username/.profile

You can uninstall at any time with rustup self uninstall and
these changes will be reverted.

Current installation options:


   default host triple: x86_64-unknown-linux-gnu
     default toolchain: stable
               profile: default
  modify PATH variable: yes

1) Proceed with installation (default)
2) Customize installation
3) Cancel installation
>
```

ここで 1 を選択する。

```text
info: profile set to 'default'
info: default host triple is x86_64-unknown-linux-gnu
info: syncing channel updates for 'stable-x86_64-unknown-linux-gnu'
info: latest update on 2020-01-30, rust version 1.41.0 (5e1a79984 2020-01-27)
info: downloading component 'cargo'
info: downloading component 'clippy'
info: downloading component 'rust-docs'
 12.0 MiB /  12.0 MiB (100 %)   5.7 MiB/s in  2s ETA:  0s
info: downloading component 'rust-std'
 17.5 MiB /  17.5 MiB (100 %)   7.2 MiB/s in  2s ETA:  0s
info: downloading component 'rustc'
 57.9 MiB /  57.9 MiB (100 %)   6.5 MiB/s in  9s ETA:  0s
info: downloading component 'rustfmt'
info: installing component 'cargo'
info: installing component 'clippy'
info: installing component 'rust-docs'
 12.0 MiB /  12.0 MiB (100 %)  10.0 MiB/s in  1s ETA:  0s
info: installing component 'rust-std'
info: installing component 'rustc'
 57.9 MiB /  57.9 MiB (100 %)  11.6 MiB/s in  5s ETA:  0s
info: installing component 'rustfmt'
info: default toolchain set to 'stable'

  stable installed - rustc 1.41.0 (5e1a79984 2020-01-27)


Rust is installed now. Great!

To get started you need Cargo's bin directory ($HOME/.cargo/bin) in your PATH
environment variable. Next time you log in this will be done
automatically.

To configure your current shell run source $HOME/.cargo/env
```

これで [Rust] のツールチェーン一式が `~/.cargo/` および `~/.rustup/` ディレクトリ以下にインストールされた。
いや，どうせなら [XDG Base Directory] に対応してくれよ `orz` ...まぁいいや。

ちなみに `~/.profile` ファイルの末尾に

```bash
export PATH="$HOME/.cargo/bin:$PATH"
```

と追記されるが，私の好みで以下のように書き換えた。

```bash
# set PATH so it includes user's private bin if it exists
if [ -d "$HOME/.cargo/bin" ] ; then
  PATH="$PATH:$HOME/.cargo/bin"
fi
```

[Rust] ツールチェーンのうち主なコマンドを以下に挙げる。

| コマンド名 | 機能                           |
| ---------- | ------------------------------ |
| `rustup`   | ツールチェーンの導入および更新 |
| `rustc`    | コンパイラ本体                 |
| `rustfmt`  | ソースコードのフォーマット     |
| `cargo`    | パッケージ管理                 |
| `rls`      | Rust Language Server           |

たとえば

```text
$ rustup update
info: syncing channel updates for 'stable-x86_64-unknown-linux-gnu'
info: checking for self-updates

  stable-x86_64-unknown-linux-gnu unchanged - rustc 1.41.0 (5e1a79984 2020-01-27)

info: cleaning up downloads & tmp directories
```

てな感じにツールチェーンのアップデートも簡単にできる。

### リンカに注意

[Rust] ではリンカについては自前のものを持っていないようだ。
[Ubuntu] 等の Linux プラットフォームであれば [GCC] が標準で入っているため問題ないが， Windows の場合は Visual Studio 等の開発環境が必要になる。

[Rust] ではクロスコンパイルも可能だが，リンカはプラットフォーム毎に用意する必要があるため注意が必要である。
まぁ Docker 等を使えばいいんだろうけど。

### エディタ機能の拡張

[Rust] の [Tools](https://www.rust-lang.org/tools "Tools - Rust Programming Language") には各種エディタや IDE で使える拡張機能が紹介されている。
私は [ATOM] を使うので [ide-rust](https://atom.io/packages/ide-rust) を使えばいいようだ。
`rls` に対応してるぞ。

## みんな大好き Hello World

適当なディレクトリに `hello.rs` ファイルを作って，以下のコードを書く。

```rs
fn main() {
    println!("Hello, world!");
}
```

これをコンパイル&実行する。

```text
$ rustc hello.rs 
$ ./hello 
Hello, world!
```

よーし，うむうむ，よーし。

### Cargo でプロジェクト・ベースの環境を作る

`cargo` コマンドを使ってプロジェクト・ベースの環境の雛形を作ることができる。

```text
$ cargo new hello --vcs none
     Created binary (application) `hello` package
```

このコマンドにより `hello` ディレクトリが作成され，以下のディレクトリ・ファイルが作成される。

```text
$ tree hello
hello
├── Cargo.toml
└── src
    └── main.rs
```

`src/main.rs` の中身は先程の `main()` 関数と同じ内容である。
`Cargo.toml` の中身は以下の通り。

```toml
[package]
name = "hello"
version = "0.1.0"
authors = ["username <username@example.com>"]
edition = "2018"

# See more keys and their definitions at https://doc.rust-lang.org/cargo/reference/manifest.html

[dependencies]
```

どうも `authors` の項目は [git] のグローバル設定を読み取ってるらしい。

ちなみに，先程のコマンドの `--vcs none` オプションを外すと [git] のリポジトリとして構築される。
また GitHub に空のリポジトリを作ってローカルに clone し

```text
$ git clone https://github.com/spiegel-im-spiegel/rust-hello.git hello
$ cargo init hello
     Created binary (application) package
```

てな感じに初期化することもできる。

`cargo` コマンドで作った環境下では `cargo` コマンド経由でビルドやデバッグなどができる。

たとえば `cargo run` コマンドを使えば

```text
$ cd hello
$ cargo run
   Compiling hello v0.1.0 (/home/username/rust/hello)
    Finished dev [unoptimized + debuginfo] target(s) in 0.32s
     Running `target/debug/hello`
Hello, world!
```

てな感じにコンパイル&実行が一気にできる。
リリース用のビルドは

```text
$ cargo build --release
   Compiling hello v0.1.0 (/home/username/rust/hello)
    Finished release [optimized] target(s) in 0.18s

$ ./target/release/hello 
Hello, world!

```

という感じ。

### 外部関数インターフェイス

[Rust] は外部関数インターフェイス（Foreign Function Interface；FFI）の仕組みを持っていて，他言語の関数を呼び出したり，逆に [Rust] の関数を他言語から呼び出したりすることができる。

今回は詳細な説明は省いて [medimatrix/rust-plus-golang] リポジトリにあるサンプルコードを紹介するに留める。

```text
$ git clone https://github.com/medimatrix/rust-plus-golang.git

$ tree rust-plus-golang/
rust-plus-golang/
├── LICENSE
├── Makefile
├── README.md
├── lib
│   ├── hello
│   │   ├── Cargo.toml
│   │   └── src
│   │       └── lib.rs
│   └── hello.h
└── main.go
```

`lib/hello/Cargo.toml` の中身はこんな感じ。

```toml
[package]
name = "hello"
version = "0.1.0"

[lib]
crate-type = ["cdylib"]

[dependencies]
libc = "0.2.2"
```

更に `lib/hello/src/lib.rs` の中身はこんな感じ。

```rust
extern crate libc;
use std::ffi::CStr;

#[no_mangle]
pub extern "C" fn hello(name: *const libc::c_char) {
    let buf_name = unsafe { CStr::from_ptr(name).to_bytes() };
    let str_name = String::from_utf8(buf_name.to_vec()).unwrap();
    println!("Hello {}!", str_name);
}
```

この `hello()` 関数を呼び出す `main.go` の中身はこんな感じ。

```go
package main

/*
#cgo LDFLAGS: -L./lib -lhello
#include "./lib/hello.h"
*/
import "C"

func main() {
    C.hello(C.CString("John Smith"))
}
```

ほんじゃあ，まぁ，ビルドしてみよっか。

```text
$ make
cd lib/hello && cargo build --release
    Updating crates.io index
  Downloaded libc v0.2.67
   Compiling libc v0.2.67
   Compiling hello v0.1.0 (/home/username/rust/rust-plus-golang/lib/hello)
    Finished release [optimized] target(s) in 27.05s
cp lib/hello/target/release/libhello.so lib/
go build -ldflags="-r /home/username/rust/rust-plus-golang/lib" main.go

$ ./main
Hello John Smith!
```

よしよし。

今回は C 言語のインタフェースを使って [Go] から [Rust] の関数を呼び出していたが，他にも Ruby や Python などとも連携可能らしい。

## ブックマーク

- [Rustのクロスコンパイル（Linux, MacOSで作ったものをWindowsで動かす）｜IT会計キャリア](https://media.itkaikei.com/2019/01/24/rust-cross-compile/)
- [Rustでクロスコンパイルしてみた | 思案試行](https://blog.varwww.com/201703-rust-cross-compile.html)

[Rust]: https://www.rust-lang.org/ "Rust Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[XDG Base Directory]: https://standards.freedesktop.org/basedir-spec/latest/ "XDG Base Directory Specification"
[GCC]: https://gcc.gnu.org/ "GCC, the GNU Compiler Collection - GNU Project - Free Software Foundation (FSF)"
[Git]: https://git-scm.com/
[git]: https://git-scm.com/
[medimatrix/rust-plus-golang]: https://github.com/medimatrix/rust-plus-golang "medimatrix/rust-plus-golang: Rust + Go"
[Go]: https://golang.org/ "The Go Programming Language"
[ATOM]: https://atom.io/

## 参考図書

{{% review-paapi "4048930702" %}} <!-- プログラミング言語Rust 公式ガイド -->
