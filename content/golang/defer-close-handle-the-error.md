+++
title = "Close 関数のエラーを無視しない（『効率的なGo』読書会より）"
date =  "2025-11-15T20:04:13+09:00"
description = "最終回の範囲（11章）から個人的に気になった部分を取り上げる。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "book" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

『[効率的なGo]』の[読書会]は無事に最終回を迎えた。
今回は，[読書会]最終回の範囲（11章 最適化パターン）から個人的に気になった部分を取り上げる。

## Close 関数のエラーを無視しない

以下のサンプルコードを起点にする。

```go {hl_lines=[13]}
package main

import (
	"fmt"
	"os"
)

func doWithFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Perform file operations here...

	return nil
}

func main() {
	err := doWithFile("example.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err)
	}
}
```

注目するのは `defer f.Close()` の部分。
`f.Close()` メソッドはエラーを返す可能性があるが，このコードではエラーを捨てている。
この手のサンプルコードはよく見るし私も過去記事で書いたような気がするが，エラーハンドリングの観点からは悪いコードということになる。

ファイルアクセスではクローズ処理の際にバッファに残ってるデータをフラッシュする場合がある。
特に zip ファイルの書き込み処理とかでは終端レコードをクローズ時に書き込んだりするので，絶対にエラーを無視してはいけない。

一般的にクロージング・メソッドを含むパッケージや構造体ではクローズ処理時のエラーを適切に処理する必要がある。
だからといって `defer` を使わないのはもったいない。
今回の場合は以下のように書き換える。

```go {hl_lines=[1,"6-10"]}
func doWithFile(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && cerr != os.ErrClosed {
			err = errors.Join(err, cerr)
		}
	}()

	// Perform file operations here...

	return err
}
```

ポイントは2つ。

1. 返り値を `err error` と名前付きにすること
2. `defer` 処理を無名関数 `func() { ... }()` で囲み，内部で適切にエラー処理を行うこと

最近の [Go] は（[`errors`]`.Join()` など）複数エラーを記述できるようになったのでハンドリングが簡潔になった[^e1]。
よしよし。

[^e1]: 拙作の [`goark/errs`](https://github.com/goark/errs "goark/errs: Error handling for Golang") パッケージもよろしく（宣伝）。こちらも複数エラーに対応している他，エラーの出力を JSON 形式に出来る。

## 最後まで読み切ってからクローズする

次のサンプルコードはこれ。
HTTP レスポンスデータを処理する関数である。

```go
func doWithResp(resp *http.Response) (err error) {
	defer func() {
		if cerr := resp.Body.Close(); cerr != nil && cerr != os.ErrClosed {
			err = errors.Join(err, cerr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Perform operations with the response

	return err
}
```

[`net/http`] パッケージのドキュメントには冒頭に

{{< fig-quote type="markdown" title="http package - net/http - Go Packages" link="https://pkg.go.dev/net/http" lang="en" >}}
The caller must close the response body when finished with it:

```
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := io.ReadAll(resp.Body)
// ...
```
{{< /fig-quote >}}

と書かれていて，必ず [`http`][`net/http`]`.Response.Body.Close()` メソッドを呼び出すよう指示されている。
クローズ処理を行わないと内部でゴルーチン・リークが発生するらしい。

ただ [`http`][`net/http`]`.Response.Body` に関しては，ただクローズするだけでは駄目で，クローズする前にデータを最後まで読み切ることが推奨されている。

{{< fig-quote type="markdown" title="http package - net/http - Go Packages" link="https://pkg.go.dev/net/http" lang="en" >}}
If the Body is not both read to EOF and closed, the [`Client`](https://pkg.go.dev/net/http#Client)'s underlying [`RoundTripper`](https://pkg.go.dev/net/http#RoundTripper) (typically [`Transport`](https://pkg.go.dev/net/http#Transport)) may not be able to re-use a persistent TCP connection to the server for a subsequent "keep-alive" request. 
{{< /fig-quote >}}

これについては『[Go言語 100Tips]』にも解説がある。

{{< fig-quote type="markdown" title="『Go言語 100Tips』 10. No.79：一時的な資源をクローズしない" link="https://www.amazon.co.jp/dp/B0CFL1DK8Q?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
もう 1 つ覚えておくべき重要なことは、ボディをクローズするときの動作は、そのボディから読み込んだかどうかによって異なることです。

- 読み込まずにボディをクローズすると、デフォルトの HTTP トランスポートがコネクションをクローズするかもしれません。
- 読み込んだ後にボディをクローズしても、デフォルトの HTTP トランスポートはコネクションをクローズしないので、再利用できる可能性があります。

したがって、`getStatusCode` が繰り返し呼ばれ、キープアライブ（keep-alive）のコネクションを使いたい場合、必要がないと思ってもボディを読み込むべきです。
{{< /fig-quote >}}

これを踏まえて，最初のサンプルコードを書き換えると以下のようになる。

```go {hl_lines=[3,5,7]}
func doWithResp(resp *http.Response) (err error) {
	defer func() {
		_, rerr := io.Copy(io.Discard, resp.Body)
		if cerr := resp.Body.Close(); cerr != nil && cerr != os.ErrClosed {
			err = errors.Join(err, rerr, cerr)
		}
		err = errors.Join(err, rerr)
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Perform operations with the response

	return err
}
```

## 余談

『[効率的なGo]』は，ちぃっと難しかった。
何が難しいって書籍内参照が前の章や後の章に飛びまくって「何だっけこれ？」ってなることが多くて。
まるで longjump だらけの[スパゲッティコード]({{< ref "/remark/2016/06/code-by-hardware-engineer.md" >}} "ハード屋が書くCソースコードが凄まじかった思い出")を読んでる気分だった。
これを曲がりなりにも理解するにはあと2,3回は周回する必要があるかな。

『[効率的なGo]』では標準パッケージの中身にも踏み込んでカスタマイズしたパッケージ置き換えるといったことも行っている。
今回挙げた例についても，エラーハンドリングに関しては[独自のパッケージ](https://github.com/efficientgo/core "efficientgo/core: Set of core packages every Go project needs. Minimal API, strictly versioned and with ~no dependencies.")が用意されている。
参考になるだろう。

今回の[読書会]の最後の方で話題になったのだが，2025年12月に『[実用 Go言語 第2版]』が出るそうで。
買わなきゃ。
先日買った『[初めてのGo言語 第2版]』もまだ読み終わってないんだけど。

[Go]: https://go.dev/
[効率的なGo]: https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: 効率的なGo ―データ指向によるGoアプリケーションの性能最適化 : Bartłomiej Płotka, 山口 能迪: 本"
[Go言語 100Tips]: https://www.amazon.co.jp/dp/B0CFL1DK8Q?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Go言語 100Tips ありがちなミスを把握し、実装を最適化する impress top gearシリーズ | Teiva Harsanyi, 柴田 芳樹 | 工学 | Kindleストア | Amazon"
[実用 Go言語 第2版]: https://www.oreilly.co.jp/books/9784814401369/ "実用 Go言語 第2版 - O'Reilly Japan"
[初めてのGo言語 第2版]: https://www.oreilly.co.jp/books/9784814401192/ "初めてのGo言語 第2版 - O'Reilly Japan"
[読書会]: https://yokohama-go-reading.connpass.com/event/373111/ "第75回横浜Go読書会（オンライン） - connpass"
[`net/http`]: https://pkg.go.dev/net/http "http package - net/http - Go Packages"
[`errors`]: https://pkg.go.dev/errors "errors package - errors - Go Packages"

## 参考図書

{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
{{% review-paapi "4814401191" %}} <!-- 初めてのGo言語 第2版 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
