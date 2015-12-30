+++
date = "2015-12-30T13:30:45+09:00"
description = "description"
draft = true
tags = ["golang", "cli", "facade"]
title = "コマンドライン・インタフェースとファサード・パターン"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

[Go 言語]コンパイラには [`flag`] パッケージが標準で組み込まれており，いわゆるコマンドライン・インタフェース（(Command line interface; CLI）の操作はこれでまかなうことができる。
ただし [`flag`] パッケージではサブコマンドをサポートしていないためサブコマンドを構成したい場合は少し工夫が必要となる。
ちなみにサブコマンドとは，以下のようなコマンドラインの構成になっているアプリケーションである。

```
$ command [golabal options] <sub-command> [sub-options] [arguments]
```

たとえば [Go 言語]コンパイラの `go run` もサブコマンドだし， [git] の `git commit` とかもサブコマンドである。

## コマンドライン・インタフェースと UNIX Philosophy

ところで CLI でよく引き合いに出されるのが “[UNIX Philosophy]” と呼ばれるアプリケーションを作る際の哲学というか指針のようなものである。
曰く

1. Small is beautiful. （小さいものは美しい）
2. Make each program do one thing well. （各プログラムが一つのことをうまくやるようにせよ）
3. Build a prototype as soon as possible. （できる限り早くプロトタイプを作れ）
4. Choose portability over efficiency. （効率よりも移植しやすさを選べ）
5. Store data in flat text files. （単純なテキストファイルにデータを格納せよ）
6. Use software leverage to your advantage. （ソフトウェアの効率を優位さとして利用せよ）
7. Use shell scripts to increase leverage and portability. （効率と移植性を高めるためにシェルスクリプトを利用せよ）
8. Avoid captive user interfaces. （拘束的なユーザーインターフェースは作るな）
9. Make every program a Filter. （全てのプログラムはフィルタとして振る舞うようにせよ）

の9項目である[^up]。
昨今は UNIX 互換環境でも GUI が普通になってきたので対話型のインタフェースも増えてきたが，それでも従来の CUI shell 上で動作するアプリケーションの需要が減ったわけではなく，サーバサイドの環境ではむしろ需要は大きくなっていると言ってもいい。

[^up]: 翻訳は [Wikipedia の記事](https://ja.wikipedia.org/wiki/UNIX%E5%93%B2%E5%AD%A6)から拝借させてもらった。ちなみに [Wikipedia のコンテンツは基本的には by-sa ライセンスで公開](https://ja.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License)されている。

[Go 言語]で CLI アプリケーションを作る際に気をつける点としては

- 他のツールと連携できるよう標準入出力を使ったフィルタプログラムとする
- 外部データの入出力は JSON, YAML, TOML といったテキストを用い UTF-8 文字エンコーディングに統一する
- コードの可搬性（または移植性）を考慮し，プラットフォーム依存を避けるようにする

といったところだろうか。
もともと [Go 言語]はクロスプラットフォーム開発に強いため， [Go 言語]の作法からの逸脱ができるだけないようにすれば，それほど難しい要件ではないはずである。

## サブコマンドとファサード・パターン

サブコマンド方式は一見 “[UNIX Philosophy]” に反しているように見えるが， [Go 言語]の場合は全てのパッケージをひとつの実行モジュールに結合してしまうため，関連する機能をサブコマンドとして組み込むのは悪くないやりかたである。

サブコマンドを構成する場合は「ファサード・パターン（facade pattern）」で考えるとよい。
「ファサード」は「建物の正面」という意味だそうで，システム内の各機能（サブシステム）の窓口のように機能する。

{{< fig-img src="/images/facade-pattern.svg" title="Facade Pattern" link="/images/facade-pattern.svg" >}}

この図のようにファサード・パターンは DDD (Domain-Driven Design) と相性がよい。
普通は Web アプリケーションのような多様なサブシステムを持つシステムを設計する際に導入する考え方だが， CLI の場合でもサブコマンドを構成するのであればファサード・パターンで設計するべきである。

## ファサード・パターンとしての mitchellh/cli

CLI をサポートするパッケージはいくつか公開されているのだが，この中で今回は [mitchellh/cli] を紹介する。

[mitchellh/cli] はサブコマンドをファサード・パターンで実装するのに便利な機能を実装している。


## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`flag`]: https://golang.org/pkg/flag/ "flag - The Go Programming Language"
[git]: https://git-scm.com/ "Git"
[UNIX Philosophy]: http://www.ru.j-npcs.org/usoft/WWW/LJ/Articles/unixtenets.html "Tenets of the UNIX Philosophy"
[mitchellh/cli]: https://github.com/mitchellh/cli "mitchellh/cli"
