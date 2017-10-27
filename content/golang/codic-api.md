+++
title = "Codic API を利用するパッケージを作ってみた"
date =  "2017-10-25T15:46:59+09:00"
update = "2017-10-27T18:19:23+09:00"
description = "spf13/viper を使ってみたかったのだ。"
tags        = [ "golang", "programming", "package", "api", "curl", "language", "cli" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

つい最近まで知らなかったのだが [codic] というサービスがあるらしい。

- [プログラマーのためのネーミング辞書 | codic](https://codic.jp/)

簡単に言うと日本語の「名前」を英語に変換してくれるサービスなのだが，プログラマ向けに変数名やメソッド名として使いやすいよう提案してくれる優れものである。
まさに[英語不得手な私]({{< relref "remark/2017/10/programmin-language-in-english.md" >}} "プログラミング言語の暗黙ルール")のためにあるようなサービスじゃないか！ 何故今までこのサービスに辿り着けなかったのか `orz`

Web 画面はこんな感じ。

{{< fig-img src="https://farm5.staticflickr.com/4459/37176009973_5026b4c303.jpg" title="codic service" link="https://www.flickr.com/photos/spiegel/37176009973/" >}}

あの時このサービスのことを知っていたら[メソッド名に `regist`]({{< relref "remark/2017/04/regist-dose-not-exist.md" >}} "“regist” という単語は存在しない") とか付けようとして赤っ恥をかかなくて済んだのに。
とほほ。

というわけで早速サインアップしましたよ。
GitHub のアカウントでもサインアップできるのが素敵（最終確認にメールアドレスを要求されるけど）。

で， [codic] では API を公開しているようだ。

- [API | codic](https://codic.jp/docs/api)

で，これを使うための [Go 言語]パッケージも既に見られるんだけど

- [codic-project/Codic_cli](https://github.com/codic-project/Codic_cli)
- [39e/go-codic](https://github.com/39e/go-codic)

CLI しか用意されてない，っていうか何で GET で取ろうとするんだよ！ というわけで自作することにした。

→ 自作しました。

- [spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ](https://github.com/spiegel-im-spiegel/gocodic)

すみません。
勢いで作ったのでテストを書いてません。
そのうちなんとかします。
日本語圏向けのサービスなんだから [README] もガッツリ日本語でいいよね（笑）

## Curl で API を確認する

RESTfull API なんだから [curl] で説明してくれよ，と思う私は贅沢なのでしょうか。

- [cURL as DSL — cURL as DSL 1.0 documentation](https://shibukawa.github.io/curl_as_dsl/)
- [Shibu's Diary: cURL as DSLとは何だったのか。あるいは細かすぎて伝わらないcURL as DSL。](http://blog.shibu.jp/article/115602749.html)

とりあえず，[翻訳用の API](https://codic.jp/docs/api/engine/translate) は [curl] を使うと以下のように記述できる。

```text
curl "https://api.codic.jp/v1/engine/translate.json" -H "Authorization: Bearer YOUR_ACCESS_TOKEN" -F "text=hello" -F "casing=camel"
```

`text=hello` って日本語やないやないかい！ というのはとりあえずスルーしていただいて，これを [cURL as DSL] で [Go 言語]コードに変換すると以下のようになる。

```go
package main

import (
    "bytes"
    "io/ioutil"
    "log"
    "mime/multipart"
    "net/http"
)

func main() {

    client := &http.Client{}
    var buffer bytes.Buffer
    writer := multipart.NewWriter(&buffer)
    writer.WriteField("text", "hello")
    writer.WriteField("casing", "camel")
    writer.Close()

    request, err := http.NewRequest("POST", "https://api.codic.jp/v1/engine/translate.json", &buffer)
    request.Header.Add("Authorization", "Bearer YOUR_ACCESS_TOKEN")
    request.Header.Add("Content-Type", "multipart/form-data; boundary="+writer.Boundary())

    resp, err := client.Do(request)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Print(string(body))
}
```

これが雛形で出発点である。
`YOUR_ACCESS_TOKEN` に正しいアクセス・トークン（[codic] にサインアップするともらえる）をセットすればちゃんと動く。
動くコードってのは大事だよね。

最終的にどうなったかは [README] を見ていただきたい。

## spf13/viper を使ってみたかったのだ

外部パッケージは [dep で管理]({{< relref "golang/consider-switching-from-glide-to-dep.md" >}} "Glide から Dep への移行を検討する")することにした。
こんな感じである。

```toml
[[constraint]]
  name = "github.com/spf13/cobra"
  branch = "master"

[[constraint]]
  name = "github.com/spf13/viper"
  version = "^1.0.0"

[[constraint]]
  name = "github.com/spiegel-im-spiegel/gocli"
  version = "^0.4.0"

[[constraint]]
  name = "github.com/pkg/errors"
  version = "^0.8.0"
```

[pkg/errors] パッケージ以外は CLI (Command-Line Interface) 用のパッケージである。
この中でも今回は特に [spf13/viper] を使ってみたかったのだ。
だって毎回アクセス・トークンをコマンドラインに書く訳にはいかないでしょ。
呼び出しバッチやスクリプトに書くとか以ての外だし。

[spf13/viper] は設定ファイルにアクセスするためのパッケージで，特に [spf13/cobra] との相性がいいのが特徴である。
というか同じ作者なのだが。
[spf13/cobra] の使い方は以前紹介したので，そちらを参考にして欲しい。

- [モンテカルロ法による円周率の推定（その2 CLI）]({{< relref "golang/estimate-of-pi-2-cli.md" >}})

[spf13/cobra] が生成してくれる `cmd/root.go` に [spf13/viper] 初期化のコードがある。

```go
// initConfig reads in config file and ENV variables if set.
func initConfig() {
    if cfgFile != "" {
        // Use config file from the flag.
        viper.SetConfigFile(cfgFile)
    } else {
        // Find home directory.
        home, err := homedir.Dir()
        if err != nil {
            fmt.Print(err)
            os.Exit(1)
        }

        // Search config in home directory with name ".gocodic" (without extension).
        viper.AddConfigPath(home)
        viper.SetConfigName(".gocodic")
    }

    viper.AutomaticEnv() // read in environment variables that match

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Print("Using config file:", viper.ConfigFileUsed())
    }
}
```

この関数は [spf13/cobra] が生成した `cmd` パッケージの `init()` 関数内で呼び出される。
このまま弄らなくても問題ないが，個人的にはエラーを標準出力に出してるのが気に入らなかったので少し変えている。

その後， [spf13/viper] で読み込む設定項目を記述していくのだが

```go
rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gocodic.yaml)")
rootCmd.PersistentFlags().BoolP("json", "j", false, "output by JSON format (raw data)")
rootCmd.PersistentFlags().StringP("token", "t", "", "access token of codic.jp")
viper.BindPFlag("json", rootCmd.PersistentFlags().Lookup("json"))
viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
```

こんな感じで [spf13/cobra] 側のフラグ（厳密には [spf13/pflag]）と [spf13/viper] をバインドしてしまう。
これで `cmd` パッケージ側からはフラグ情報を透過的に扱える。
フラグ情報を取り出す場合には

```go
jsonFlag := viper.GetBool("json")
token := viper.GetString("token")
```

などとすればよい。
分かれば簡単。

## ところで

[2016年4月からブログが更新されてない](http://blog.codic.jp/)けど，そのうちサービスが止まるなんてないよね？
[Twitter アカウントは生きてる](https://twitter.com/codic_project)みたいだし。

## ブックマーク

- [【codic】プログラマ必見！もう変数名や関数名に困らない！プログラマのためのネーミングツールを紹介 - プログラミング向上雑記](http://niisi.hatenablog.jp/entry/2016/08/17/171000)
- [関数や変数のネーミングに悩んだら「codic」に日本語名を入力するとある程度解決するかも](https://nelog.jp/codic)
- [よく使うcurlコマンドのオプションまとめ（12個） - Qiita](https://qiita.com/shtnkgm/items/45b4cd274fa813d29539)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[codic]: https://codic.jp/ "プログラマーのためのネーミング辞書 | codic"
[curl]: http://curl.haxx.se/ "curl and libcurl"
[cURL as DSL]: https://shibukawa.github.io/curl_as_dsl/ "cURL as DSL — cURL as DSL 1.0 documentation"
[gocodic]: https://github.com/spiegel-im-spiegel/gocodic "spiegel-im-spiegel/gocodic: codic の API を利用するための Go 言語パッケージ"
[README]: https://github.com/spiegel-im-spiegel/gocodic/blob/master/README.md "gocodic/README.md at master · spiegel-im-spiegel/gocodic"
[pkg/errors]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"
[spf13/cobra]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"
[spf13/viper]: https://github.com/spf13/viper "spf13/viper: Go configuration with fangs"
[spf13/pflag]: https://github.com/spf13/pflag "spf13/pflag: Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags."
