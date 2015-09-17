+++
date = "2015-09-12T11:26:44+09:00"
update = "2015-09-12T11:26:44+09:00"
description = "みんなだいすき Hello World!"
draft = false
tags = ["golang", "install", "helloworld" ]
title = "インストールから Hello World まで"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) - Qiita](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa)）

最初に言い訳しておくと，現時点での私の主要マシンは Windows 機なため， Windows 上での動作を前提としている。
ただし Go 言語コンパイラはマルチプラットフォームに対応しているため，操作に関してはほぼ同じと考えてよい。
Linux 等のマシンを使っている方は適当に読み替えて欲しい（簡単でしょ）。

## Go コンパイラのインストール（on Windows）

Go言語はコンパイル言語である。
プラットフォームとして FreeBSD, Lunux, Mac OS X, Windows などがある。
また，クロスコンパイルが比較的容易なことでも知られている。

現時点（2015-09-08）での Go コンパイラの最新版は 1.5.1。
Windows 版では[ダウンロードページ](https://golang.org/dl/)にインストール・パッケージが用意されているので，ダウンロードしてインストールすればよい。

Go コンパイラが最低限動作するのに必要な環境変数は（`PATH` を除けば） `GOROOT` のみである。
Windows 版の場合は，インストール・パッケージからインストールすれば自動的に環境変数もセットされる。
セットされていない場合は手動で `GOROOT` にインストール先のフォルダを指定すればいい。
（参考： [Windows - SETX コマンドで環境変数を永続的に設定する - Qiita](http://qiita.com/rohinomiya/items/cf5236678b3459da9017)）

インストールができたら動作確認。

```
C:>go version
go version go1.5 windows/amd64
```

## みんなだいすき Hello World!

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}
```

このソースコード `hello.go` を compile and run。

```
C:>go run hello.go
Hello World!
```

おおっ，動いた動いた。じゃあこれを build して，実行モジュールを起動してみる。

```
C:>go build hello.go

C:>hello.exe
Hello World!
```

よし。ちゃんと動くようだな。
今回はここまで。

## 【おまけ】 1.5 系へのアップグレードに関する注意点

Windows 版（64bit）のインストールパッケージで 1.4 系から 1.5 系へ上書きインストールしようとしたらエラーになった。

{{< img src="https://farm6.staticflickr.com/5759/20692976166_a38bee50d8_o.png" alt="Install Error" title="Install Error" caption="Install Error" link="https://www.flickr.com/photos/spiegel/20692976166/" >}}

この場合は，コントロールパネルの「プログラムと機能」で既存のバージョンをアンインストールしてから最新バージョンをインストールし直せば OK。

1.5 系から Go コンパイラ自身で自身をコンパイルできるようになった。
この影響で， Linux 環境などでは 1.5 系を導入するために 1.4 系の Go コンパイラをあらかじめインストールする必要がある。

- [Go 1.3 から 1.5 へのアップデートでエラー - Qiita](http://qiita.com/taji-taji/items/4c43e126e67d65a219e3)

## ブックマーク

- [The Go Programming Language](https://golang.org/)
    - [git repositories (Google)](https://go.googlesource.com/)
    - [git repositories (GitHub)](https://github.com/golang) : mirror
- [golang-jp - The Go Programming Language](http://golang-jp.org/) : 本家の日本語訳サイト（一部のみ）。（[golang.jp](http://golang.jp/) は古いので参考にしない方がいい，らしい）

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。
