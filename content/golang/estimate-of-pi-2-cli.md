+++
description = "さっそく推定結果について評価を行いたいところだが，その前に CLI (command-line interface) を整備する。今回は spf13/cobra パッケージを使うことにする。"
tags = ["golang", "cli", "facade", "circle-ratio"]
date = "2016-11-06T17:57:37+09:00"
update = "2017-01-16T09:42:12+09:00"
title = "モンテカルロ法による円周率の推定（その2 CLI）"
draft = false

[author]
  flickr = "spiegel"
  facebook = "spiegel.im.spiegel"
  tumblr = "spiegel-im-spiegel"
  linkedin = "spiegelimspiegel"
  flattr = "spiegel"
  license = "by-sa"
  instagram = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  twitter = "spiegel_2007"
  avatar = "/images/avatar.jpg"
  github = "spiegel-im-spiegel"
  name = "Spiegel"
+++

1. [モンテカルロ法による円周率の推定（その1）]({{< relref "golang/estimate-of-pi.md" >}})
2. [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "golang/estimate-of-pi-2-cli.md" >}}) ← イマココ
3. [モンテカルロ法による円周率の推定（その3 Gaussian）]({{< relref "golang/estimate-of-pi-3-gaussian.md" >}})
4. [モンテカルロ法による円周率の推定（その4 PRNG）]({{< relref "golang/estimate-of-pi-4-prng.md" >}})

## コマンドライン・インタフェース

さっそく推定結果について評価を行いたいところだが，その前に CLI (command-line interface) を整備する。
どういうことかというと，[前回]作った2つの `main()` 関数の処理

```go
package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/spiegel-im-spiegel/pi/gencmplx"
)

func main() {
    c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(10000))
    for p := range c {
        fmt.Printf("%v\t%v\n", real(p), imag(p))
    }
}
```

```go
package main

import (
    "fmt"
    "math/cmplx"
    "math/rand"
    "time"

    "github.com/spiegel-im-spiegel/pi/gencmplx"
)

func main() {
    c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), int64(100000))
    n := int64(0) // total
    m := int64(0) // plot in circle
    for p := range c {
        n++
        if cmplx.Abs(p) <= float64(1) {
            m++
        }
    }
    fmt.Printf("n = %v, m = %v, 4m/n = %v\n", n, m, float64(4*m)/float64(n))
}
```

これをひとつの CLI で呼び出せるよう統合してしまおうというわけ。

CLI については以前に解説した。

- [コマンドライン・インタフェースとファサード・パターン]({{< relref "golang/cli-and-facade-pattern.md" >}})

このときは [`mitchellh/cli`] を紹介したが，今回は [`spf13/cobra`] を使うことにする。

### spf13/cobra パッケージ

[`spf13/cobra`] パッケージの作者 [spf13 (Steve Francia)](https://github.com/spf13) さんは [Docker の中の人](https://www.linkedin.com/in/stevefrancia "Steven Francia | LinkedIn")で [Hugo] の作者としても有名な方。
もちろん [Hugo] の CLI にも [`spf13/cobra`] が使われている。

さらにありがたいことに [`spf13/cobra`] にはテンプレートコードを出力する CLI も用意されている。
インストールは `go get` コマンドで行う。

```text
$ go get -v github.com/spf13/cobra/cobra
```

これで [`spf13/cobra`] パッケージ本体と CLI がインストールされる。
テンプレートコードの生成は以下のコマンドを叩く。

```text
$ cobra init github.com/spiegel-im-spiegel/pi
Your Cobra application is ready at
C:\workspace\pi\src\github.com\spiegel-im-spiegel\pi
Give it a try by going there and running `go run main.go`
Add commands to it by running `cobra add [cmdname]`
```

既にパッケージ用のフォルダが作られている場合は，そのフォルダまで降りて

```text
$ cobra init .
Your Cobra application is ready at
C:\workspace\pi\src\github.com\spiegel-im-spiegel\pi
Give it a try by going there and running `go run main.go`
Add commands to it by running `cobra add [cmdname]`
```

でもよい。

この時点で `main.go` と `cmd/root.go` のふたつが生成される。

```go
// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "github.com/spiegel-im-spiegel/pi/cmd"

func main() {
    cmd.Execute()
}
```

```go
// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
    Use:   "pi",
    Short: "A brief description of your application",
    Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
// Uncomment the following line if your bare application
// has an action associated with it:
//    Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}

func init() {
    cobra.OnInitialize(initConfig)

    // Here you will define your flags and configuration settings.
    // Cobra supports Persistent Flags, which, if defined here,
    // will be global for your application.

    RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pi.yaml)")
    // Cobra also supports local flags, which will only run
    // when this action is called directly.
    RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" { // enable ability to specify config file via flag
        viper.SetConfigFile(cfgFile)
    }

    viper.SetConfigName(".pi") // name of config file (without extension)
    viper.AddConfigPath("$HOME")  // adding home directory as first search path
    viper.AutomaticEnv()          // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}
```

この状態でいきなり動かしてみる。

```text
$ go run main.go
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.
```

`RootCmd` で定義した説明が表示されているのがわかると思う。

次にサブコマンドを定義する。
名前は `plot` と `estmt` としようか。

```text
$ cobra add plot
plot created at C:\workspace\pi\src\github.com\spiegel-im-spiegel\pi\cmd\plot.go

$ cobra add estmt
estmt created at C:\workspace\pi\src\github.com\spiegel-im-spiegel\pi\cmd\estmt.go
```

これで `cmd/plot.go` と `cmd/estmt.go` のふたつが生成された。
`cmd/plot.go` の中身を見てみよう。

```go
// Copyright Â© 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

// plotCmd represents the plot command
var plotCmd = &cobra.Command{
    Use:   "plot",
    Short: "A brief description of your command",
    Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
    Run: func(cmd *cobra.Command, args []string) {
        // TODO: Work your own magic here
        fmt.Println("plot called")
    },
}

func init() {
    RootCmd.AddCommand(plotCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // plotCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // plotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
```

[`spf13/cobra`] パッケージで特徴的なのは，サブコマンドを追加する際に `cmd/root.go` を変更する必要が無いことである。
サブコマンドの組み込みは `cmd` パッケージ内の各ファイルに定義されている `init()` 関数によって `main()` 起動前に行われる。

この状態で動かしてみよう。

```text
$ go run main.go
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  pi [command]

Available Commands:
  estmt       A brief description of your command
  plot        A brief description of your command

Flags:
      --config string   config file (default is $HOME/.pi.yaml)
  -t, --toggle          Help message for toggle

Use "pi [command] --help" for more information about a command.

$ go run main.go plot --help
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  pi plot [flags]

Global Flags:
      --config string   config file (default is $HOME/.pi.yaml)

$ go run main.go plot
plot called
```

問題なくサブコマンドが組み込まれていることがわかる。

## CLI の作成

では，生成されたテンプレートをベースに機能を組み込んでいく。 ...というわけで，出来上がりが以下の repository にある。

- [spiegel-im-spiegel/pi: Estimate of Pi with Monte Carlo method.](https://github.com/spiegel-im-spiegel/pi)

フォルダ構成は以下の通り。

```text
github.com/spiegel-im-spiegel/pi
|   .editorconfig
|   .gitignore
|   glide.lock
|   glide.yaml
|   LICENSE
|   main.go
|   README.md
|
+---cmd
|       estmt.go
|       plot.go
|       root.go
|
+---estmt
|       estmt.go
|
+---gencmplx
|       gencmplx.go
|
+---genpi
|       genpi.go
|
\---plot
        plot.go
```

そして各パッケージの構成は以下のようになっている。

{{< fig-img src="/images/estimate-of-pi.svg" title="パッケージ構成" link="/images/estimate-of-pi.svg" width="640" >}}

[前回]と変わったところは `genpi` パッケージを追加したことだろうか。
こんな感じ。

```go
package genpi

import (
    "math/cmplx"
    "math/rand"
    "time"

    "github.com/spiegel-im-spiegel/pi/gencmplx"
)

//New returns generator of Pi
func New(pc, ec int64) <-chan float64 {
    ch := make(chan float64)
    pcf := float64(pc)
    go func(pc, ec int64) {
        for i := int64(0); i < ec; i++ {
            c := gencmplx.New(rand.NewSource(time.Now().UnixNano()), pc)
            m := int64(0) // plot in circle
            for p := range c {
                if cmplx.Abs(p) <= float64(1) {
                    m++
                }
            }
            ch <- float64(4*m) / pcf
        }
        close(ch)
    }(pc, ec)

    return ch
}
```

指定した回数だけ推定処理を行い，結果を [channel] `ch` に渡している。
たとえば100,000個の点から円周率を推定する処理を10回したければ

```text
$ go run main.go estmt -e 10 -p 100000
3.14576
3.1422
3.13716
3.14648
3.14852
3.13952
3.14824
3.13828
3.14532
3.14312
```

とすればよい。

これでようやく評価のための準備が整った。

## ブックマーク

- [GolangでCLIの場合にcobraを使うことにした件 | FLAMA技術Blog](http://lab.flama.co.jp/archives/1536/)
- [packageに複数のinitがあるときの挙動 - Qiita](http://qiita.com/astronoka/items/aa2f271d280863cedf5e)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "golang/estimate-of-pi.md" >}} "モンテカルロ法による円周率の推定（その1）"
[`mitchellh/cli`]: https://github.com/mitchellh/cli "mitchellh/cli"
[`spf13/cobra`]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"
[Hugo]: https://gohugo.io/ "Hugo :: A fast and modern static website engine"
[channel]: http://golang.org/ref/spec#Channel_types

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
