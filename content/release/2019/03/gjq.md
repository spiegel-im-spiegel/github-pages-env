+++
title = "jq ぽい何か を書いてみた"
date = "2019-03-16T23:12:32+09:00"
description = "私が Windows を捨てれば simeji/jid で全然 OK なのだが，鼻の先はそういうわけにもいかないので，自分用に CLI を書いてみることにした。"
image = "/images/attention/tools.png"
tags  = [ "tools", "json", "golang", "cli" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ウソです。
ほとんど自力じゃありません（笑）

JSON データを操作できる [jq] はとても高機能で使い勝手もいいのだが，純粋にフィルタとして機能するため Windows 環境では微妙に使いづらい場合がある。
具体的には

- 入力が標準入力のみなのでリダイレクトやパイプと組み合わせる必要がある
- 対話モードがない

といったところ。
大した問題ではないのだがどうにかならないかなぁ，と思っていた。

んで，どなたか [Go 言語]で書いている人がいるんじゃないかと思って探してみたら以下のツール/パッケージを見つけた。

- [simeji/jid]
    - [JSONをインタラクティブに掘り下げるコマンド jid  - Qiita](https://qiita.com/simeji/items/dd0464b7ed91c51ee618)
- [savaki/jq]
- [apang1992/jq]

[simeji/jid] はよく出来ていると思うのだが，残念なことに [ConEmu] 上で動かすと表示がぶっ壊れる（コマンドプロンプトでは問題ない）。
[savaki/jq] と [apang1992/jq] はパッケージだが，ここ2,3年は更新されていないようである。

私が Windows を捨てれば [simeji/jid] で全然 OK なのだが（そのうち捨てる），鼻の先はそういうわけにもいかないので，自分用に CLI を書いてみることにした。

- [spiegel-im-spiegel/gjq]

コマンドライン・オプションは以下の通り。

```text
$ gjq -h
Usage:
  gjq [flags] <filter string>

Flags:
      --debug         for debug
  -f, --file string   JSON data (file path)
  -h, --help          help for gjq
  -i, --indent int    indent size for formatted JSON string (default 2)
  -I, --interactive   interactive mode
  -t, --tab           use tabs for indentation
  -u, --url string    JSON data (URL)
  -v, --version       output version of gjq
```

パーサには [savaki/jq] を使っている。
フィルタの文法は以下の通り。

| syntax     | meaning                                         |
| ---------- | ----------------------------------------------- |
| `.`        | unchanged input                                 |
| `.foo`     | value at key                                    |
| `.foo.bar` | value at nested key                             |
| `.[0]`     | value at specified element of array             |
| `.[0:1]`   | array of specified elements of array, inclusive |
| `.foo.[0]` | nested value                                    |

まずは [jq] と同じようにフィルタとして使うならこんな感じ。

```text
$ cat testdata/test.json
{
  "string": "a",
  "number": 1.23,
  "simple": ["a", "b", "c"],
  "mixed": [
    "a",
    1,
    {"hello":"world"}
  ],
  "object": {
    "first": "joe",
    "array": [1,2,3]
  }
}

$ gjq .object.array < testdata/test.json
[
  1,
  2,
  3
]

$ cat testdata/test.json | gjq .mixed
[
  "a",
  1,
  {
    "hello": "world"
  }
]
```

インデントは `-i` オプションで調整できるが 0 をセットすると

```text
$ gjq -i 0 .mixed < testdata/test.json
["a",1,{"hello":"world"}]
```

とコンパクトになる。

また `-f` オプションでファイルを直接指定することもできる。

```text
$ gjq -f testdata/test.json .object.array
[
  1,
  2,
  3
]
```

さらに `-u` オプションで URL を指定すれば Web 上のファイルも対象になる[^u1]。

[^u1]: `-u` オプションは HTTP/HTTPS の GET で対象を取得する。 POST には対応してないのであしからず。 POST で取ってくる必要があるなら潔く curl を使ったほうがいいだろう。

```text
$ gjq -u https://text.baldanders.info/index.json .entry.[0]
{
  "title": "猫を殺すに猫を以ってせよ",
  "section": "remark",
  "description": "分かるかな。分っかんねーだろうな（反語）",
  "author": "Spiegel",
  "license": "http://creativecommons.org/licenses/by-sa/4.0/",
  "url": "https://text.baldanders.info/remark/2019/03/no-cat-no-life/",
  "published": "2019-03-11T13:51:41+00:00",
  "update": "2019-03-11T13:53:40+00:00"
}
```

さらにさらに `-I` オプションを付けると対話モードになる。

```text
$ gjq -u https://text.baldanders.info/index.json -I
Press Ctrl+C to stop
Filter> .title
"text.Baldanders.info"
Filter> .entry.[0]
{
  "title": "猫を殺すに猫を以ってせよ",
  "section": "remark",
  "description": "分かるかな。分っかんねーだろうな（反語）",
  "author": "Spiegel",
  "license": "http://creativecommons.org/licenses/by-sa/4.0/",
  "url": "https://text.baldanders.info/remark/2019/03/no-cat-no-life/",
  "published": "2019-03-11T13:51:41+00:00",
  "update": "2019-03-11T13:53:40+00:00"
}
Filter> 
```

対話モードではクリップボードにも出力されるのでターミナル上でコピペする必要はない。
また JSON データは起動時に1度だけ GET するので，安心して対話モードで弄れる（笑）

さて次は色を付けてみるかな。
[Go 言語]でターミナル出力に色を付けるってどうやるんだろ。
調べてみるか。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/gjq]: https://github.com/spiegel-im-spiegel/gjq "spiegel-im-spiegel/gjq: Another Implementation of "jq" by golang"
[jq]: https://stedolan.github.io/jq/
[simeji/jid]: https://github.com/simeji/jid "simeji/jid: json incremental digger"
[savaki/jq]: https://github.com/savaki/jq "savaki/jq: A high performance Golang implementation of the incredibly useful jq command line tool."
[apang1992/jq]: https://github.com/apang1992/jq "apang1992/jq: json query with golang"
[ConEmu]: https://conemu.github.io/ "ConEmu - Handy Windows Terminal"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
    <dd>Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>単行本（ソフトカバー）</dd>
    <dd>4621300253 (ASIN), 9784621300251 (EAN), 4621300253 (ISBN), 9784621300251 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
