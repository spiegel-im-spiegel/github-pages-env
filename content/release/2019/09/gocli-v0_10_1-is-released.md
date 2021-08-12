+++
title = "spiegel-im-spiegel/gocli v0.10.1 のリリース"
date = "2019-09-07T11:55:58+09:00"
description = "書き直した結果 os.UserConfigDir() 関数で取得したパスにアプリケーション名と設定ファイル名をくっ付けただけの簡単なお仕事になった（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "package", "module", "golang", "cli" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Go 言語用 CLI プログラミング支援パッケージ [`gocli`] v0.10.1 をリリースした。

- [Release v0.10.1 · spiegel-im-spiegel/gocli · GitHub](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.10.1)

実は『[改訂2版 みんなのGo言語](https://www.amazon.co.jp/exec/obidos/ASIN/B07VPSXF6N/baldandersinf-22/)』を読んで「そろそろ設定ファイルの置き場所を何とかしなくちゃ」と思ってこっそり [v0.10.0](https://github.com/spiegel-im-spiegel/gocli/releases/tag/v0.10.0 "Release v0.10.0 · spiegel-im-spiegel/gocli") をリリースしたのだが， [Go] 1.13 で [`os`]`.UserConfigDir()` 関数が用意されているのを知って「もうこれでいいぢゃん」てな感じになり，書き直した。こんな感じに使える。

```go
package main

import (
	"fmt"

	"github.com/spiegel-im-spiegel/gocli/config"
)

func main() {
	path := config.Path("app", "config.json")
	fmt.Println(path)
	// Output:
	// /home/username/.config/app/config.json
}
```

書き直した結果 [`os`]`.UserConfigDir()` 関数で取得したパスにアプリケーション名と設定ファイル名をくっ付けただけの簡単なお仕事になった（笑）
私は設定ファイルと CLI との連携に [spf13/viper] を使っているので，これで必要十分。

[`gocli`] の使い方について詳しくは以下を参照のこと。

- [Go 言語用 CLI プログラミング支援パッケージ]({{< ref "/release/gocli-package-for-golang.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`gocli`]: https://github.com/spiegel-im-spiegel/gocli "spiegel-im-spiegel/gocli: Minimal Packages for Command-Line Interface"
[`os`]: https://golang.org/pkg/os/ "os - The Go Programming Language"
[spf13/viper]: https://github.com/spf13/viper "spf13/viper: Go configuration with fangs"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->

{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
