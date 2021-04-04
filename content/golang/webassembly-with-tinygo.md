+++
title = "TinyGo で WebAssembly"
date =  "2021-03-11T21:01:29+09:00"
description = "Go および TinyGo を使って WebAssembly へのコンパイルを行い Web ブラウザ上で動作させるところまでやってみる"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "tinygo", "webassembly" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[TinyGo] は本家 [Go] のサブセットと言えるもので [LLVM] を使った組み込み用途特化のコンパイラである。
しかも [LLVM] が [WebAssembly] バイナリを直接出力できるということもあって [TinyGo] と [WebAssembly] の相性は本家 [Go] 以上と言える。

というわけで今回は， [Go] および [TinyGo] を使って [WebAssembly] へのコンパイルを行い， Web ブラウザ上で動作させるところまでやってみることにする。

## [TinyGo] のインストール

[TinyGo] が動作するためには，あらかじめ本家 [Go] のツールチェーンが導入済みであることが前提となる。
この記事では [Go] は導入済みであるとして話を進める。

[TinyGo] は以下のリポジトリから最新版をダウンロード&インストールする。

- [tinygo-org/tinygo: Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.](https://github.com/tinygo-org/tinygo)

2021-03-11 時点での最新版は 0.17.0 である。
[Go] 1.16 以降が推奨らしい。

### Ubuntu の場合

Ubuntu の APT や Snap の公式リポジトリにはないので， deb ファイルをダウンロードし，手動でインストールする。

```text
$ curl -L https://github.com/tinygo-org/tinygo/releases/download/v0.17.0/tinygo_0.17.0_amd64.deb -O
$ sudo dpkg -i tinygo_0.17.0_amd64.deb
$ tinygo version
tinygo version 0.17.0 linux/amd64 (using go version go1.16.1 and LLVM version 11.0.0)
```

ちなみに WSL/WSL2 上の Ubuntu にもインストール可能だそうだ。

### Windows の場合

Windows なら [Scoop] を使うのが最も簡単である。
[Scoop] なら本家 [Go] も簡単にインストールできるし [TinyGo] 用の周辺ツールも [Scoop] で簡単に導入できる。

```text
$ scoop install tinygo
$ tinygo version
tinygo version 0.17.0 windows/amd64 (using go version go1.16.1 and LLVM version 11.0.0)
```

### Docker の場合

Docker 環境も用意されているそうだ。
詳しくはこちら。

- [Docker :: TinyGo - Go on Microcontrollers and WASM](https://tinygo.org/getting-started/using-docker/)

## ファイル構成

今回のファイル構成はこんな感じ。

```text
$ tree .
.
├── hello
│   ├── hello.go
│   ├── index.html
│   ├── wasm.js
│   └── wasm_exec.js
└── main.go
```

`main.go` は簡易サーバのコードで，こんな感じになっている。

```go
package main

import (
    "embed"
    "fmt"
    "io/fs"
    "net/http"
    "os"
    "strings"
)

//go:embed hello
var assets embed.FS

func main() {
    addr := "localhost:3000"
    fmt.Printf("Open http://%s/\n", addr)
    fmt.Println("Press ctrl+c to stop")

    root, _ := fs.Sub(assets, "hello")
    fs := http.FileServer(http.FS(root))
    http.Handle("/", http.FileServer(http.FS(root)))
    if err := http.ListenAndServe(addr, http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
        resp.Header().Add("Cache-Control", "no-cache")
        if strings.HasSuffix(req.URL.Path, ".wasm") {
            resp.Header().Set("content-type", "application/wasm")
        }
        fs.ServeHTTP(resp, req)
    })); err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
}
```

[`embed`] パッケージと `//go:embed` ディレクティブが便利！ 簡易サーバのコードについては拙文「[紙芝居用の簡易サーバを書く【Go 1.16 版】]({{< relref "./embeded-filesystem.md" >}}) 」を参照のこと。
今回用の設定としては `*.wasm` ファイルの Content-Type を `application/wasm` にすることくらいかな。
あとは `no-cache` の設定ね。

`wasm_exec.js` ファイルは [Go] および [TinyGo] が用意しているファイルで，以下からコピってそのまま使えばよい。

| 処理系   | パス                               |
| -------- | ---------------------------------- |
| [Go]     | `$GOROOT/misc/wasm/wasm_exec.js`   |
| [TinyGo] | `$TINYGOROOT/targets/wasm_exec.js` |

`$GOROOT` および `$TINYGOROOT` の値は，以下のコマンドで取得できる。

```text
$ tinygo env
GOOS="linux"
GOARCH="amd64"
GOROOT="/usr/local/go1.16.1"
GOPATH="/home/username/go"
GOCACHE="/home/username/.cache/tinygo"
CGO_ENABLED="1"
TINYGOROOT="/usr/local/lib/tinygo"
```

## みんな大好き Hello World

さて `hello/hello.go` ファイルの中身だが，まずはこんな感じで。

```go
// +build js,wasm

package main

import "syscall/js"

func main() {
    ch := make(chan struct{})
    js.Global().Get("document").Call("getElementById", "hello").Set("innerHTML", "Hello, World!")
    <-ch // Code must not finish
}
```

JavaScript の DOM 構造に慣れている人ならそんなに難しくないだろう。
ID 名 `hello` の要素に文字列 `Hello, World!` を突っ込むだけの簡単なお仕事（笑）

これを [Go] および [TinyGo] の各コンパイラでコンパイルしてみる。

```text
$ GOOS=js GOARCH=wasm go build -o hello1.wasm -trimpath
$ tinygo build -o hello2.wasm -target wasm ./hello.go
```

前者が本家 [Go] によるコンパイルで，後者が [TinyGo] によるコンパイルだ。
コンパイル結果は以下の通り。

```text
$ ll *.wasm
-rwxrwxr-x 1 username username 1364695  3月 10 23:59 hello1.wasm*
-rwxrwxr-x 1 username username   67375  3月 10 23:59 hello2.wasm*
```

おうふ。
こんなに違うのか。

本家 [Go] のコードが大きいのは，良くも悪くも POSIX 互換環境への依存度が高く組み込み用途に使うには余計なコードを抱え込んでしまうという事情がある。

一方 [TinyGo] は [LLVM] の制約を受けるため，ガベージコレクションや並行処理などで本家 [Go] とは異なる挙動になる（他にもいくつかの標準パッケージが使えない場合があるらしい）。

たとえば先ほどの

```go {hl_lines=[4]}
func main() {
    ch := make(chan struct{})
    js.Global().Get("document").Call("getElementById", "hello").Set("innerHTML", "Hello, World!")
    <-ch // Code must not finish
}
```

の最後を見ると，チャネル受信状態で処理が止まっているが（というか止めるためにわざわざこのように書いているのだが），これがないと [TinyGo] コンパイラがエラーを吐く場合があるようだ。

### [WebAssembly] コードのバインド

`hello/wasm.js` ファイルは生成した [WebAssembly] コードを JavaScript 側にバインドするものである。
今回は以下のように書いてみた。

```js
function initWASM(url) {
    const go = new Go();

    if ('instantiateStreaming' in WebAssembly) {
        WebAssembly.instantiateStreaming(fetch(url), go.importObject).then(function (obj) {
            go.run(obj.instance);
        })
    } else {
        fetch(WASM_URL).then(resp =>
            resp.arrayBuffer()
        ).then(bytes =>
            WebAssembly.instantiate(bytes, go.importObject).then(function (obj) {
                go.run(obj.instance);
            })
        )
    }
}
```

これで

```js
initWASM('hello2.wasm');
```

という感じに任意の [WebAssembly] ファイルを取り込める。

### HTML の内容

以上を踏まえて `hello/index.html` ファイルの内容は以下のようにした。

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"/>
<title>Hello</title>
<script src="wasm_exec.js"></script>
<script src="wasm.js"></script>
</head>

<body>
    <p id="hello"></p>
</body>

<script>
    initWASM("hello2.wasm");
</script>
</html>
```

### 実行結果

では，いよいよ動かしてみよう。

```text
$ go run main.go
Open http://localhost:3000/
Press ctrl+c to stop
```

該当の URL を開くと

{{< fig-img src="./hello1.png" link="./hello1.png" title="Hello" >}}

よーし，ちゃんと表示されているな。
ここまでは楽勝。

## [WebAssembly] の機能を JavaScript から呼び出す

以上のコードは [WebAssembly] 側から HTML 要素に値をセットしていたが，これではあまり応用できないだろう。
なので，今度は JavaScript 側から [WebAssembly] の機能を呼び出すことを考える。

まずは `hello/hello.go` ファイルの内容を以下のように変更する。

```go
// +build js,wasm

package main

import (
    "strings"
    "syscall/js"
)

func say(this js.Value, args []js.Value) interface{} {
    ss := []string{}
    for _, jss := range args {
        if s := jsString(jss); s != "" {
            ss = append(ss, s)
        }
    }
    return js.ValueOf("Hello, " + strings.Join(ss, ", "))
}

func jsString(j js.Value) string {
    if j.IsUndefined() || j.IsNull() {
        return ""
    }
    return j.String()
}

func main() {
    ch := make(chan struct{})
    js.Global().Set("say", js.FuncOf(say))
    <-ch // Code must not finish
}
```

JavaScript から呼び出す関数は

```go
func(this js.Value, args []js.Value) interface{}
```

の関数型にする決まりのようだ。
また返り値は [`js`]`.Value` 型にして返すのだが，実際の [Go] の型と JavaScript の型の対応は以下のようになっているらしい。

{{< fig-quote class="nobox" type="markdown" title="js - The Go Programming Language" link="https://golang.org/pkg/syscall/js/#ValueOf" lang="en" >}}
| Go                       | JavaScript   |
| ------------------------ | ------------ |
| `js.Value`               | [its value]  |
| `js.Func`                | `function`   |
| `nil`                    | `null`       |
| `bool`                   | `boolean`    |
| integers and floats      | `number`     |
| `string`                 | `string`     |
| `[]interface{}`          | `new array`  |
| `map[string]interface{}` | `new object` |
{{< /fig-quote >}}


さらに `main()` 関数内の

```go
js.Global().Set("say", js.FuncOf(say))
```

によって [Go] の `say()` 関数は JavaScript の `global.say()` 関数に紐付けられる。

`hello/wasm.js` ファイルはそのままでOK。
`hello/index.html` ファイルの内容を以下のように書き換える。

```html
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8"/>
<title>Hello</title>
<script src="wasm_exec.js"></script>
<script src="wasm.js"></script>
</head>

<body>
	<button>Hello</button>
	<p id="hello"></p>
</body>

<script>
	initWASM('hello2.wasm');
	document.querySelector('button').addEventListener('click', event => {
		document.getElementById("hello").innerHTML = global.say("Alice", "Bob", "Chris");
	});
</script>
</html>
```

これで `[Hello]` ボタン押下で `global.say()` 関数が発火するはずである。
実行してみよう。

```text
$ go run main.go
Open http://localhost:3000/
Press ctrl+c to stop
```

この状態で画面を表示すると

{{< fig-img src="./hello-button1.png" link="./hello-button1.png" title="Hello Button (1)" >}}

さらにボタンを押下すると

{{< fig-img src="./hello-button2.png" link="./hello-button2.png" title="Hello Button (2)" >}}

よーし，うむうむ，よーし。

これなら応用が効きそうかな。
今回はここまで。

## ブックマーク

- [WebAssembly · golang/go Wiki · GitHub](https://github.com/golang/go/wiki/WebAssembly)
- [Wasmer Go embedding 1.0 lift-off](https://wasmer.io/posts/wasmer-go-embedding-1.0)
  - [GitHub - wasmerio/wasmer-go: 🐹🕸️ WebAssembly runtime for Go](https://github.com/wasmerio/wasmer-go)
- [Wasm By Example](https://wasmbyexample.dev/examples/webassembly-linear-memory/webassembly-linear-memory.go.en-us.html)
- [Go 製 WebAssembly ホスト環境パッケージ wax のご紹介 - bearmini's blog](https://bearmini.hatenablog.com/entry/2019/12/23/173924)
- [Go × WebAssemblyで電卓のWebアプリを作ってみた - Sansan Builders Blog](https://buildersbox.corp-sansan.com/entry/2019/02/14/113000)
- [WASI (WebAssembly system interface) を Wasmtime と Node.js で試す - Qiita](https://qiita.com/takewell/items/c99b44d912448e9170e6)
- [Go の形態素解析器を wasm で利用する](https://zenn.dev/ikawaha/articles/20210331-e661ac866f5ff0fc5eb8) : このコード例は参考になる

[Go]: https://golang.org/ "The Go Programming Language"
[TinyGo]: https://tinygo.org/ "TinyGo - Go on Microcontrollers and WASM"
[WebAssembly]: https://webassembly.org/ "WebAssembly"
[LLVM]: https://llvm.org/ "The LLVM Compiler Infrastructure Project"
[Scoop]: https://scoop.sh/ "Scoop"
[`embed`]: https://golang.org/pkg/embed/ "embed - The Go Programming Language"
[`js`]: https://golang.org/pkg/syscall/js/ "js - The Go Programming Language"

## 参考図書

{{% review-paapi "B08T7D2LFR" %}} <!-- ソフトウェアデザイン 2021年3月号 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
