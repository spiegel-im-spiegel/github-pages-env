+++
title = "TinyGo で WASI 【失敗編】"
date =  "2021-03-21T17:38:59+09:00"
description = "どうやったら動かせるのか。どなたか教えてください 🙇"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "tinygo", "webassembly" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./webassembly-with-tinygo.md" >}})は [Go] および [TinyGo] を使って [WebAssembly] コードを生成しブラウザ上で実行するところまでやった。

しかし，クライアント側のブラウザ上で動かすだけではあまり面白くないよね。
そこで WASI (WebAssembly System Interface) という POSIX 風の標準規格があるそうな。
WASI に則った [WebAssembly] コードと，それを駆動するランタイム環境を用意することで “Write Once, Run Anywhere” の夢よもう一度，というわけ[^wora1]（笑）

[^wora1]: “Write Once, Run Anywhere” は初期の Java のキャッチフレーズだった。当時は UNIX 機のハードウェア非互換の問題が酷くて，なんとかバイナリ互換を確保する方法がないかみんな頭を悩ませていた。そこに登場したのが Sun Microsystems の Java だったわけ。でも実際にはプラットフォーム間の差異が微妙に残ってしまい，むしろ “Write Once, Debug Everywhere” などと揶揄されることもあった。それでも Virtual Machine 上で標準化されたバイトコードを駆動させるというアイデアは秀逸だったので Java 以外の処理系でも応用され，特に組み込み用途では重宝されている。

実は本家 [Go] の `wasm` アーキテクチャは WASI に対応していない。
ただし [TinyGo] のほうはイケるみたいなので，今回は [TinyGo] オンリーでお送りする。

## WASI ランタイム

スタンドアロンで動く WASI ランタイムには色々あるようで

- [Lucet](https://github.com/bytecodealliance/lucet "bytecodealliance/lucet: Lucet, the Sandboxing WebAssembly Compiler.")
- [Wasmtime]
- [Wasmer](https://docs.wasmer.io/ "Wasmer Docs")
- [WAVM](https://wavm.github.io/)
- [SSVM](https://www.secondstate.io/ssvm/ "The Second State WebAssembly VM")

といった実装があるらしい。

ただ [TinyGo] のターゲット定義が

```json {hl_lines=[20]}
{
    "llvm-target":   "wasm32--wasi",
    "build-tags":    ["wasm", "wasi"],
    "goos":          "linux",
    "goarch":        "arm",
    "compiler":      "clang",
    "linker":        "wasm-ld",
    "libc":          "wasi-libc",
    "cflags": [
        "--target=wasm32--wasi",
        "--sysroot={root}/lib/wasi-libc/sysroot",
        "-Oz"
    ],
    "ldflags": [
        "--allow-undefined",
        "--stack-first",
        "--export-dynamic",
        "--no-demangle"
    ],
    "emulator":      ["wasmtime"],
    "wasm-abi":      "generic"
}
```

と [Wasmtime] をリファレンスとしているみたいなので，今回はこれを使う。

### [Wasmtime] の導入

[Wasmtime] のリポジトリでバイナリが[リリース](https://github.com/bytecodealliance/wasmtime/releases "Releases · bytecodealliance/wasmtime")されているので，これを取ってきて PATH の通ったディレクトリに放り込んでおけばよい。

あるいは 

```text
$ curl https://wasmtime.dev/install.sh -sSf | bash
```

とすれば `$HOME/.wasmtime/bin/` ディレクトリを掘って入れてくれる。
さらに PATH を通すために `$HOME/.bashrc` ファイルを書き換えてくれやがるので，ご注意を。

なお [Wasmtime] 自体のビルドには Rust と C++ (多分 GCC の g++) のビルド環境が必要らしい。
時代は Rust なんだねぇ。

以下，動作確認。

```text
$ wasmtime --help
wasmtime 0.25.0
Wasmtime WebAssembly Runtime

USAGE:
    wasmtime <SUBCOMMAND>

FLAGS:
    -h, --help       Prints help information
    -V, --version    Prints version information

SUBCOMMANDS:
    config      Controls Wasmtime configuration settings
    help        Prints this message or the help of the given subcommand(s)
    run         Runs a WebAssembly module
    wasm2obj    Translates a WebAssembly module to native object file
    wast        Runs a WebAssembly test script file

If a subcommand is not provided, the `run` subcommand will be used.

Usage examples:

Running a WebAssembly module with a start function:

  wasmtime example.wasm

Passing command line arguments to a WebAssembly module:

  wasmtime example.wasm arg1 arg2 arg3

Invoking a specific function (e.g. `add`) in a WebAssembly module:

  wasmtime example.wasm --invoke add 1 2
```

## みんな大好き Hello World

何はともあれ，コードを用意しないとね。
いつものように，みんな大好き Hello World で。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

これを [TinyGo] で処理する。

```text
$ tinygo build -o hello.wasm -target wasi ./hello.go
```

ターゲットが `wasi` になっている点に注意。

## [Wasmtime] で WASI コードを動かす

んではビルドした `hello.wasm` ファイルを実行してみる。

```text
$ wasmtime run hello.wasm 
Hello, World!
```

よーし，うむうむ，よーし。

## [wasmtime-go] で WASI ランタイムを組み込む【失敗編】

[bytecodealliance/wasmtime-go][wasmtime-go] を使うと [Wasmtime] のランタイム機能を [Go] のコードとして埋め込めるらしい（要 cgo）。
こんな感じかな。

```go
package main

import (
    _ "embed"
    "fmt"
    "os"

    "github.com/bytecodealliance/wasmtime-go"
)

//go:embed hello.wasm
var wasm []byte

func main() {
    store := wasmtime.NewStore(wasmtime.NewEngine())

    wasiConfig := wasmtime.NewWasiConfig()
    wasiConfig.InheritStdout()
    wasi, err := wasmtime.NewWasiInstance(store, wasiConfig, "wasi_snapshot_preview1")
    if err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in wasmtime.NewWasiInstance() : %w", err))
        return
    }

    linker := wasmtime.NewLinker(store)
    if err := linker.DefineWasi(wasi); err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in wasmtime.Linker.DefineWasi() : %w", err))
        return
    }

    if err := wasmtime.ModuleValidate(store, wasm); err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in wasmtime.ModuleValidate() : %w", err))
        return
    }
    module, err := wasmtime.NewModule(store.Engine, wasm)
    if err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in wasmtime.NewModule() : %w", err))
        return
    }

    instance, err := linker.Instantiate(module)
    if err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in wasmtime.Linker.Instantiate() : %w", err))
        return
    }

    if _, err := instance.GetExport("_start").Func().Call(); err != nil {
        fmt.Fprintln(os.Stderr, fmt.Errorf("error in \"_start\" : %w", err))
        return
    }
}
```

では，これを動かしてみよう。

```text
$ go run sample.go 
error in wasmtime.Linker.Instantiate() : unknown import: `wasi_unstable::fd_write` has not been defined
```

おうふ。
なんか足らんと言っている。

[TinyGo] 側でなにか不備があるのかと思って以下の[サンプル・コード](https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go#example-package-Wasi)もそのまま動かしてみたが

```go
package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"

    "github.com/bytecodealliance/wasmtime-go"
)

const TextWat = `(module
    ;; Import the required fd_write WASI function which will write the given io vectors to stdout
    ;; The function signature for fd_write is:
    ;; (File Descriptor, *iovs, iovs_len, nwritten) -> Returns number of bytes written
    (import "wasi_unstable" "fd_write" (func $fd_write (param i32 i32 i32 i32) (result i32)))

    (memory 1)
    (export "memory" (memory 0))

    ;; Write 'hello world\n' to memory at an offset of 8 bytes
    ;; Note the trailing newline which is required for the text to appear
    (data (i32.const 8) "hello world\n")

    (func $main (export "_start")
        ;; Creating a new io vector within linear memory
        (i32.store (i32.const 0) (i32.const 8))  ;; iov.iov_base - This is a pointer to the start of the 'hello world\n' string
        (i32.store (i32.const 4) (i32.const 12))  ;; iov.iov_len - The length of the 'hello world\n' string

        (call $fd_write
            (i32.const 1) ;; file_descriptor - 1 for stdout
            (i32.const 0) ;; *iovs - The pointer to the iov array, which is stored at memory location 0
            (i32.const 1) ;; iovs_len - We're printing 1 string stored in an iov - so one.
            (i32.const 20) ;; nwritten - A place in memory to store the number of bytes written
        )
        drop ;; Discard the number of bytes written from the top of the stack
    )
)`

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    dir, err := ioutil.TempDir("", "out")
    check(err)
    defer os.RemoveAll(dir)
    stdoutPath := filepath.Join(dir, "stdout")

    engine := wasmtime.NewEngine()
    store := wasmtime.NewStore(engine)

    linker := wasmtime.NewLinker(store)

    // Configure WASI imports to write stdout into a file.
    wasiConfig := wasmtime.NewWasiConfig()
    wasiConfig.SetStdoutFile(stdoutPath)

    // Set the version to the same as in the WAT.
    wasi, err := wasmtime.NewWasiInstance(store, wasiConfig, "wasi_snapshot_preview1")
    check(err)

    // Link WASI
    err = linker.DefineWasi(wasi)
    check(err)

    // Create our module
    wasm, err := wasmtime.Wat2Wasm(TextWat)
    check(err)
    module, err := wasmtime.NewModule(store.Engine, wasm)
    check(err)
    instance, err := linker.Instantiate(module)
    check(err)

    // Run the function
    nom := instance.GetExport("_start").Func()
    _, err = nom.Call()
    check(err)

    // Print WASM stdout
    out, err := ioutil.ReadFile(stdoutPath)
    check(err)
    fmt.Print(string(out))
}
```

結果は同じで `wasi_unstable::fd_write` なんぞ知らんと言ってくさる。
えっ？ みんなこのサンプルコード動かせるの？ どうやんだ？ 多分ランタイム側で何か足らないんだろうけど，よく分からん。
`wasmtime-c-api` を組み込めばいいのかなと思ったが，違うよなぁ？

というところで挫折した `orz` どなたか教えてください {{< emoji "ペコン" >}}

## 【おまけ】 Node.js で WASI を動かす

Node.js は v13 から WASI に対応しているらしい。

```text
$ npm i wasi
```

でパッケージを組み込めば使えるようだ。
で，こんな感じのコードを書いて

```js
'use strict';
const fs = require('fs');
const { WASI } = require('wasi');

const wasi = new WASI({
  args: process.argv,
  env: process.env,
  preopens: {
  }
});
const importObject = { wasi_unstable: wasi.wasiImport };
// const importObject = { wasi_snapshot_preview1: wasi.wasiImport };

(async () => {
  try {
    const wasm = await WebAssembly.compile(fs.readFileSync('./hello.wasm'));
    const instance = await WebAssembly.instantiate(wasm, importObject);
    wasi.start(instance);
  } catch (e) {
    console.error(e)
  }
})();
```

動かしてみると

```text
$ node --experimental-wasi-unstable-preview1 --experimental-wasm-bigint wasi.js
(node:210549) ExperimentalWarning: WASI is an experimental feature. This feature could change at any time
(Use `node --trace-warnings ...` to show where the warning was created)
Hello, World!
```

おー，動いた動いた。
これで [Go] のコードを WSAI 経由で JavaScript コードに埋め込めるわけだ。

## ブックマーク

- [WASI - WebAssembly System Interface with Wasmtime - DEV Community](https://dev.to/sendilkumarn/wasi-webassembly-system-interface-with-wasmtime-4cec)
- [コンテナ技術を捨て、 WASIを試す. こんにちは、NTTの藤田です。 | by FUJITA Tomonori | nttlabs | Medium](https://medium.com/nttlabs/wasi-6060b243ac90)

[Go]: https://golang.org/ "The Go Programming Language"
[TinyGo]: https://tinygo.org/ "TinyGo - Go on Microcontrollers and WASM"
[WebAssembly]: https://webassembly.org/ "WebAssembly"
[Wasmtime]: https://wasmtime.dev/ "Wasmtime — a small and efficient runtime for WebAssembly & WASI"
[wasmtime-go]: https://github.com/bytecodealliance/wasmtime-go "bytecodealliance/wasmtime-go: Go WebAssembly runtime powered by Wasmtime"

## 参考図書

{{% review-paapi "B08T7D2LFR" %}} <!-- ソフトウェアデザイン 2021年3月号 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
