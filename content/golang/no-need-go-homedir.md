+++
title = "go-homedir はもう要らない"
date =  "2020-01-26T21:48:22+09:00"
description = "もし使っているなら標準パッケージに代替できないか検討することをオススメしたい。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "programming", "package", "module" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Shimane.go#03] の LT に

- [GoのMakefileとgo.modを調べてみた](https://speakerdeck.com/ryer/shimanego3-lt)

というのがあって，これは GitHub の公開リポジトリから `go.mod` ファイルを探して使われているパッケージを数え上げてランキングにするという非常に面白い内容だったのだが，この中で [mitchellh/go-homedir] が割と使われているのが気になったので，記事にしてみる。

[mitchellh/go-homedir] はプラットフォーム非依存でユーザのホーム・ディレクトリを取得するパッケージである。

もともとホーム・ディレクトリを取得する手段として

```go
package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
        return
	}
	fmt.Println(u.HomeDir)
}
```

みたいな感じで標準の [`os`]`/`[`user`] パッケージが用意されているが，このパッケージは pure [Go] ではないためユーザの間で微妙に評判が悪いのだ。
この不満を解消する手段として pure [Go] 実装である [mitchellh/go-homedir] が使われ続けた経緯がある。

しかし，実は [Go] 1.12 から [`os`]`.UserHomeDir()` 関数が pure [Go] で実装されたため上述のパッケージは（少なくともホーム・ディレクトリを取得する手段としては）不要になった。

こんな感じ。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
        return
	}
	fmt.Println(home)
}
```

もし [mitchellh/go-homedir] パッケージを使っているなら [`os`]`.UserHomeDir()` 関数で代替できないか検討することをオススメしたい。

## 設定ファイル格納ディレクトリの取得

ところで，『[改訂2版 みんなのGo言語]』の2.8章では設定ファイルの置き場所について解説されている。
特に UNIX 系のプラットフォームでは [XDG Base Directory] の仕様に準拠した構成が求められることが多い。
たとえば `my-app` の設定ファイル `config.json` は

```text
$HOME/.config/my-app/config.json
```
 
 に置く，という感じ。

[Go] 1.13 では [`os`]`.UserConfigDir()` 関数が追加され，設定ファイルのパスの組み立てが簡単になった。
『[改訂2版 みんなのGo言語]』で書かれているような条件分けは，もはや不要である。

こんな感じ。

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(filepath.Join(conf, "my-app", "config.json"))
}
```

ちなみに，このコードを “[The Go Playground](https://play.golang.org/)” 上で実行すると

```text
neither $XDG_CONFIG_HOME nor $HOME are defined
```

と怒られる（笑）

というわけで， [Go 言語]は日々進歩してるんだよ，というお話でした。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Shimane.go#03]: https://shimane-go.connpass.com/event/159977/ "Shimane.go#03 - connpass"
[mitchellh/go-homedir]: https://github.com/mitchellh/go-homedir "mitchellh/go-homedir: Go library for detecting and expanding the user's home directory without cgo."
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[`user`]: https://golang.org/pkg/os/user/ "user - The Go Programming Language"
[改訂2版 みんなのGo言語]: https://www.amazon.co.jp/dp/B07VPSXF6N?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "改訂2版 みんなのGo言語 | 松木 雅幸, mattn, 藤原 俊一郎, 中島 大一, 上田 拓也, 牧 大輔, 鈴木 健太 | コンピュータ・IT | Kindleストア | Amazon"
[XDG Base Directory]: https://standards.freedesktop.org/basedir-spec/latest/ "XDG Base Directory Specification"

## ブックマーク

- [Go 言語の環境変数管理]({{< relref "./go-env.md" >}})
- [Go 言語用 CLI プログラミング支援パッケージ]({{< ref "/release/gocli-package-for-golang.md" >}})

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
