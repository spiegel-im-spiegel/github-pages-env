+++
title = "OpenAI API を使って GPT と遊ぶ gpt-cli をリリースした。"
date =  "2023-03-28T20:49:13+09:00"
description = "先日書いたとおり，ちょっと真面目に ChatGPT で遊んでみることにした。"
image = "/images/attention/go-logo_blue.png"
tags = [ "golang", "tools", "openai", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## OpenAI API で遊びたい

[先日書いた]({{< ref "/remark/2023/03/practice-on-the-wall.md" >}} "ChatGPT という壁打ち")とおり，ちょっと真面目に ChatGPT で遊んでみることにした。
んで，とりあえず Web ブラウザでチャットは何かと辛いので，コマンドライン・ツールで近い機能が実装できないかなぁ，と物色してみたら

- [sashabaranov/go-openai: OpenAI ChatGPT, GPT-3, GPT-4, DALL·E, Whisper API wrapper for Go][sashabaranov/go-openai]

というパッケージを使うのが[よさげ](https://zenn.dev/spiegel/scraps/fa5fcc9f7f1781 "ターミナルで ChatGPT とお話するための #golang コード")である。

というわけで，このパッケージを用いた CLI ツールを作ってみた。

- [goark/gpt-cli: CLI tool for GPT with OpenAI API](https://github.com/goark/gpt-cli)

ぶっちゃけ，この手のツールはみんな思いつくようで， [Go] 製のものだけでもそこそこ見かける。
でも，まぁ，こういうのは実際にコードを書いて動かしてみないと理解できなかったりするからねぇ。
N番煎じはご容赦を。

拙作の場合はこんな感じ。

```text
$ gpt-cli -h
CLI tool for GPT with OpenAI API.

Usage:
  gpt-cli [flags]
  gpt-cli [command]

Available Commands:
  chat        Chat with GPT-x
  help        Help about any command
  version     Print the version number

Flags:
      --api-key string     OpenAI API key
      --config string      Config file (default /home/username/.config/gpt-cli/config.yaml)
      --debug              for debug
  -h, --help               help for gpt-cli
      --log-dir string     Directory for log files (default "/home/username/.cache/gpt-cli")
      --log-level string   Log level [nop|error|warn|info|debug|trace] (default "nop")

Use "gpt-cli [command] --help" for more information about a command.
```

今のところはチャット機能しかないが， [sashabaranov/go-openai] を使えば，いま話題沸騰（笑）の GPT-4 や DALL·E, Whisper あたりもハンドリングできるみたいなので，そのうち挑戦してみたいと思っている。
ちなみに GPT-4 は申請中だが，しばらく順番は回ってこないだろうな。

## OpenAI API Key を取得する

拙作を動かす場合は [OpenAI API] Key を取得する必要がある。
取得方法は適当にググってください。
[OpenAI][OpenAI API] のアカウントを取ると，最初は無料だけど，無料枠を使い切ると有料（従量制）になるのでご注意を。

で， [OpenAI API] Key はコマンドラインで指定することもできるけど，設定ファイルに記述しておくことができる。
規定ファイルは `$XDG_CONFIG_HOME/gpt-cli/config.yaml`。
中身は YAML 形式で，こんな感じに記述できる。

```yaml
api-key: your_api_key_string
```

Linux の場合は， `$XDG_CONFIG_HOME` は `$HOME/.config/` ディレクトリに割り当てられていることが多い。
Windows だと `%AppData%` フォルダ， macOS (Darwin) だと `$HOME/Library/Application/` フォルダになる。
コマンドラインの `--config` オプションで設定ファイルを指定することもできる。

## GPT とチャットする

```text
$ gpt-cli chat -h
Chat with GPT-x, input from standard input.

Usage:
  gpt-cli chat [flags]
  gpt-cli chat [command]

Aliases:
  chat, c

Available Commands:
  history     Print chat history
  interactive Interactive mode

Flags:
  -a, --attach-file strings   Path of attach files (text file only)
  -c, --clipboard             Input message from clipboard
  -h, --help                  help for chat
  -m, --message string        Chat message
  -p, --prepare-file string   Path of prepare file (JSON format)
  -f, --save-file string      Path of save file (JSON format)

Global Flags:
      --api-key string     OpenAI API key
      --config string      Config file (default /home/username/.config/gpt-cli/config.yaml)
      --debug              for debug
      --log-dir string     Directory for log files (default "/home/username/.cache/gpt-cli")
      --log-level string   Log level [nop|error|warn|info|debug|trace] (default "nop")

Use "gpt-cli chat [command] --help" for more information about a command.
```

いちばん簡単な使い方は以下の通り（[OpenAI API] Key は設定ファイルに記述済みとする）。

```text
$ gpt-cli c -m "hello"
Hello! How can I help you today?

save to /home/username/.cache/gpt-cli/chat.2133582955.json
```

[OpenAI API] を使って “`hello`” と投げると “`Hello! How can I help you today?`” と返ってきた。
最後に，やり取りの記録を `$XDG_CACHE_HOME/gpt-cli/chat.*.json` ファイル[^c1] に格納して終了（保存先は `--save-file` または `-f` オプションで指定できる）。
そして，次に起動するときは `--prepare-file` または `-p` オプションで記録を格納したファイルを指定すれば続きから始められる。

[^c1]: Linux では `$XDG_CACHE_HOME` は `$HOME/.cache/` ディレクトリに割り当てられていることが多い。 Windows だと `%LocalAppData%` フォルダ， macOS (Darwin) だと `$HOME/Library/Caches/` フォルダになる。

`--prepare-file` オプションを使うことで初期状態をある程度いじることができる。
`--prepare-file` オプションで指定するファイルは JSON 形式で，たとえば

```json
{
  "model": "gpt-3.5-turbo-0301",
  "max_tokens": 256,
  "temperature": 0.7,
  "messages": [
    {
      "role": "system",
      "content": "これからプログラムのコードを渡すので，質問に答えてください。"
    }
  ]
}
```

などと記述しておいて，これを `--prepare-file` オプションで読み込ませて初期状態にできる。

### チャットにファイルを添付する

`--attach-file` または `-a` オプションを使ってファイルを添付して評価してもらうこともできる。
こんな感じ。

```text
$ cat sample/hello.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}

$ gpt-cli c -a sample/hello.go -m "Summarize the source code below."
The source code is a basic program in the Go programming language that prints the text "Hello, world." to the console. It includes the standard library package "fmt" and a main function that utilizes the "Println" function from the "fmt" package.

save to /home/username/.cache/gpt-cli/chat.3658387869.json
```

`--attach-file` オプションを複数使って複数のファイルを指定することが可能。
また `*.go` みたいなワイルドカード指定もできる（bash 上などではファイル名が勝手に展開されてしまうので `"*.go"` のように引用符で囲むとよい）。

### クリップボードの内容をメッセージとして送信する

`--clipboard` または `-c` オプションを指定することでクリップボードの内容をメッセージとして送信できる。

たとえば上述の `sample/hello.go` の内容をクリップボードにコピーした上で

```text
$ gpt-cli c -c -m "次のコードを要約してください。"
このコードは、mainパッケージをインポートし、"Hello, world."という文字列を出力するプログラムです。

save to /home/username/.cache/gpt-cli/chat.2878035472.json
```

などとできる。
送信の順番は `--message` オプションで指定した文字列 → クリップボードの内容 → 添付ファイル の順に送信される。

## 対話モード

```text
$ gpt-cli chat interactive -h
Interactive mode in chat.

Usage:
  gpt-cli chat interactive [flags]

Aliases:
  interactive, i

Flags:
  -h, --help               help for interactive
  -p, --profile string     Path of profile file (JSON format)
  -f, --save-file string   Path of save file (JSON format)

Global Flags:
      --api-key string     OpenAI API key
      --config string      Config file (default /home/username/.config/gpt-cli/config.yaml)
      --debug              for debug
      --log-dir string     Directory for log files (default "/home/username/.cache/gpt-cli")
      --log-level string   Log level [nop|error|warn|info|debug|trace] (default "nop")
```

こんな感じに使える。

```text
$ gpt-cli c i
Input 'q' or 'quit' to stop

Chat> こんにちは

こんにちは！私はAIのアシスタントです。何かお手伝いできることがありますか？

Chat> 科学雑誌の日経サイエンスにChatGPTの特集が載るそうです。どんな内容になるか予想してみましょう。

ChatGPTとは、人工知能による自然言語処理を利用した対話型のチャットボットです。日経サイエンスがChatGPTに特集を組むことから、以下のような予想が考えられます。

1. ChatGPTの機能や特徴の紹介
ChatGPTの性能、動作原理などの詳しい解説や、どのような機能があるのか、どのような場面で活躍するのかなどが紹介される可能性があります。

2. ChatGPTの応用例の紹介
ChatGPTが活用される応用例として、カスタマーサポート、医療支援、人工知能アシスタントなどがあります。特集では、ChatGPTがどのようにこれらの分野で活躍しているのか、具体例を交えて解説されるかもしれません。

3. ChatGPTの進化と今後の展望について
ChatGPTは、自然言語処理分野での革新的な技術であり、今後、ますます進化していくことが期待されています。特集では、今後どのような改良がされ、また、普及の度合いがどこまで広がるのか、その展望についても取り上げられるかもしれません。

Chat> ありがとうございます。雑誌を読むのが楽しみです

どういたしまして。日経サイエンスは、科学技術に関する話題や最新の研究成果などが掲載され、とても興味深く読める雑誌ですね。ChatGPTに特集が組まれるということで、ますます注目が集まりそうですね。楽しみにしていてください。

Chat> q

save to /home/username/.cache/gpt-cli/chat.2788390222.json
```

入力は1行のみ可（複数行編集の実装は保留中）。
ファイルの添付とかもできない。
本当に文章で対話するだけ（笑）

## チャット履歴を表示

```text
$ gpt-cli chat history -h
Print chat history.

Usage:
  gpt-cli chat history [flags]

Aliases:
  history, hist, h

Flags:
  -a, --assistant-name string   Assistant name (display name)
  -h, --help                    help for history
  -f, --history-file string     Path of history file (JSON format)
  -u, --user-name string        User name (display name)

Global Flags:
      --api-key string     OpenAI API key
      --config string      Config file (default /home/username/.config/gpt-cli/config.yaml)
      --debug              for debug
      --log-dir string     Directory for log files (default "/home/username/.cache/gpt-cli")
      --log-level string   Log level [nop|error|warn|info|debug|trace] (default "nop")
```

最後に出力される記録ファイルを読み込んで markdown 風のテキストに整形して出力できる。
オマケ機能。
こんな感じ。

```text
$ gpt-cli c h -u Spiegel -a ChatGPT -f /home/username/.cache/gpt-cli/chat.2133582955.json
# Chat with GPT

- `model`: gpt-3.5-turbo-0301

## Spiegel

hello

## ChatGPT

Hello! How can I help you today?
```

ChatGPT とのやりとりを [Gist に上げたり](https://gist.github.com/spiegel-im-spiegel)してるんだけど，手動で整形するのが面倒くさくて（笑）

## よっしゃ，今日はこれぐらいにしといたるわ

というわけで，今回はこのへんで。
しばらく遊んでみて，必要なら機能を足すかも知れん。
こうやって，ツールがカオスになっていくんだねぇ（笑）

## ブックマーク

- [ChatGPTに回答の参考文献を提示させるには - ZDNET Japan](https://japan.zdnet.com/article/35201375/)
- [ChatGPTのプロンプトをLispで書く](https://zenn.dev/u_u/articles/54902c757ffce5)
- [ChatGPTで開発効率アップ！askiを使ったCLIコードレビュー＆バグ特定 - Qiita](https://qiita.com/kznrluk/items/6d5ecd22c1b19d3e1d56)
- [TOP 3 open-source AI to code like a PRO 🧠 🤖 - DEV Community](https://dev.to/disukharev/top-3-open-source-ai-tools-for-programmers-4oed)
- [go-readline-ny でマルチライン編集 - 標準愚痴出力](https://zetamatta.hatenablog.com/entry/2023/03/24/233301)

[Go]: https://go.dev/
[OpenAI API]: https://platform.openai.com/
[sashabaranov/go-openai]: https://github.com/sashabaranov/go-openai "sashabaranov/go-openai: OpenAI ChatGPT, GPT-3, GPT-4, DALL·E, Whisper API wrapper for Go"

## 参考図書

{{% review-paapi "B0BXQ2HMQ5" %}} <!-- 日経サイエンス2023年5月号（特集：対話するAI ChatGPT） -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
