+++
date = "2015-09-11T17:58:42+09:00"
update = "2015-12-16T17:09:55+09:00"
description = "本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。"
draft = false
tags = ["golang", "bookmark"]
title = "Go 言語に関するブックマーク"

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

本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。なお，この記事は以下の2箇所にマルチポストでメンテしている。内容はほぼ同じなのでお好きな方でどうぞ。

- [Go 言語に関するブックマーク - Qiita](http://qiita.com/spiegel-im-spiegel/items/98d49ac456485b007a15)
- [Go 言語に関するブックマーク - プログラミング言語 Go | text.Baldanders.info](http://text.baldanders.info/golang/bookmark/)

## 公式サイト

- [The Go Programming Language](https://golang.org/) : 2015年12月2日時点での最新は 1.5.2
    - [git repositories (Google)](https://go.googlesource.com/)
    - [git repositories (GitHub)](https://github.com/golang) : mirror
- [golang-jp - The Go Programming Language](http://golang-jp.org/) : 本家の日本語訳サイト。（[golang.jp](http://golang.jp/) は参考にしない方がいい，らしい）

### Go 1.5 Released.

- [Go 1.5 is released - The Go Blog](https://blog.golang.org/go1.5)
- [Go 1.5 Release Notes - The Go Programming Language](https://golang.org/doc/go1.5)
    - [Go 1.4 "Internal" Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit) : Internal Packages は 1.5 で GOPATH まで拡張された
    - [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)

## 言語仕様に関すること

- [はじめてのGo―シンプルな言語仕様，型システム，並行処理：特集｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/feature/01/go_4beginners)
- [Golangの基本文法をおさえてみる - Qiita](http://qiita.com/kazusa-qooq/items/40f9ea3e72406d845b10)
- [Golang の 数値型 - Qiita](http://qiita.com/tanaka0325/items/9c61a022cd32be0c65a6)
- [忙しい人のためのA Tour of Go - Qiita](http://qiita.com/makoto_kw/items/0638c0af1002647e3f7a)
- [Goのfor rangeで思った値が取れなかった話 - Qiita](http://qiita.com/modal_soul/items/e49480e5692597fda975) : ちょっとしたミス
- [Go - unsafe が unsafe なケース (1) - Qiita](http://qiita.com/kwi/items/185bb3fe0d60ca765ab0)
- [Big Sky :: Go言語でインタフェースの変更がそれ程問題にならない理由](http://mattn.kaoriya.net/software/lang/go/20130919023425.htm)
- [import 書き方まとめ - Qiita](http://qiita.com/taji-taji/items/5a4f17bcf5b819954cc1)
- [ErlangとGolangを比較してみる - Qiita](http://qiita.com/soranoba/items/68d57b4635a2917f3c73)
- [GoのEnumイディオム - Qiita](http://qiita.com/awakia/items/c81c7816b9aea5422250) : あらかじめ Enum 用の type を作成し、その type に対する `String()` メソッドを定義する
- [Big Sky :: golang では for ループの中で defer してはいけない。](http://mattn.kaoriya.net/software/lang/go/20151212021608.htm) : ループ内で defer が必要になるということは refactoring のチャンス

### Go では「継承（inheritance）」ではなく「委譲（delegation）」を実装する

- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)

### 文字列操作または変換

[string](http://golang.org/ref/spec#String_types), [rune](http://blog.golang.org/strings)

- [Strings, bytes, runes and characters in Go - The Go Blog](http://blog.golang.org/strings)
- [Go言語のstring, runeの正体とは？ - golang - The Round](http://knightso.hateblo.jp/entry/2014/06/24/090719)
- [Go言語は空文字に対してstrings.splitを掛けると1要素の配列を返す - Qiita](http://qiita.com/Sheile/items/ba51ac9091e09927b95c) : コメントに別解あり
- [golang - Goでマルチバイトが混在した文字列をtruncateする - Qiita](http://qiita.com/hokaccha/items/3d3f45b5927b4584dbac)
- [Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](http://matope.hatenablog.com/entry/2014/04/22/101127)
- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
- [Go で euc-jp や sjis の csv ファイルを読み込むには変換用のリーダーを1つかませるだけでよかった - Qiita](http://qiita.com/ikawaha/items/540c2af34b1c381c37c1)
- [Goでは文字列連結はコストの高い操作 - Qiita](http://qiita.com/ruiu/items/2bb83b29baeae2433a79)
- [Goの文字列結合のパフォーマンス - Qiita](http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37)
- [Go言語で SplitMultiSep (複数種の区切り文字列で分解) - Qiita](http://qiita.com/yoya/items/23ac2c490625c5d47ad9)
- [Go で UTF-8 の文字列を扱う - Qiita](http://qiita.com/masakielastic/items/01a4fb691c572dd71a19)

### 配列

[slice](http://golang.org/ref/spec#Slice_types), [map](http://golang.org/ref/spec#Map_types), [make](http://golang.org/ref/spec#Making_slices_maps_and_channels)

- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang - go言語のslice操作をまとめてみた（shiftしたりpushしたり） - Qiita](http://qiita.com/egnr-in-6matroom/items/282aa2fd117aab9469bd)

[slice](http://golang.org/ref/spec#Slice_types), [map](http://golang.org/ref/spec#Map_types), および後述の [channel](http://golang.org/ref/spec#Channel_types) は組み込みの型だが内部構造と状態を持つため， [new](http://golang.org/ref/spec#Allocation) ではなく [make](http://golang.org/ref/spec#Making_slices_maps_and_channels) を使う。

### 並行処理と並列処理

[goroutine](http://golang.org/ref/spec#Go_statements), [channel](http://golang.org/ref/spec#Channel_types)

- [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
- [Go: 計算なしのFizzBuzz - Qiita](http://qiita.com/suin/items/eca21ed935115e5da2e8) : channel の説明するのにいいかも
- [Goのchannelの送受信用の型について - Qiita](http://qiita.com/yuki2006/items/3f90e53ce74c6cff1608)
- [Go言語のChannelは送信時にもブロックする - Qiita](http://qiita.com/hondata/items/64776c79063e93bea9ed) : 意外と見落とす channel 送信時のブロック
- [Go - select loop の小ネタ - Qiita](http://qiita.com/Jxck_/items/da3ca2db58734a966cac)
- [Goのforとgoroutineでやりがちなミスとたった一つの冴えたgo vetと - Qiita](http://qiita.com/sudix/items/67d4cad08fe88dcb9a6d)
- [golang - x/net/context の実装パターン - Qiita](http://qiita.com/tutuming/items/c0ffdd28001ee0e9320d) : [golang.org/x/net/context](https://godoc.org/golang.org/x/net/context) を使って並行処理を細かく制御。 Domain-Driven あるいは Context-Driven な設計でも使えそう。
- [Go言語でチャネルとselect - Qiita](http://qiita.com/najeira/items/71a0bcd079c9066347b4)
- [golangでシグナルを拾ってgracefulにgoroutineを停めたい - Qiita](http://qiita.com/arc279/items/c44d4a18a851ff454c64)
- [golang の channel のブロックがよくわからん - Qiita](http://qiita.com/arc279/items/bc55cdf436c544e91c05)
- [GoのChannelを使いこなせるようになるための手引 - Qiita](http://qiita.com/awakia/items/f8afa070c96d1c9a04c9)
- [Goでスレッド（goroutine）セーフなプログラムを書くために必ず注意しなければいけない点 - Qiita](http://qiita.com/ruiu/items/54f0dbdec0d48082a5b1) : `sync.Mutex` にも言及
- [GAE/Goでもgoroutine使おうぜ！というハナシ - Qiita](http://qiita.com/hogedigo/items/f0f409ee944c4b2107c3) : GAE/Go の API には非同期版がないので積極的に goroutine 使おうよ，という話

Go 言語で複数 CPU を使った並列処理を行うには明示的な設定が必要。

- [Go言語でCPU数に応じて並列処理数を制限する | SOTA](http://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/)
    - [やはり俺のgolangがCPUを一つしか使わないのはまちがっている。 - Qiita](http://qiita.com/ymko/items/554e3630fefdc29393a8)
- [Goでお手軽に行列の積を爆速並列計算 - Qiita](http://qiita.com/hama_du/items/fce4ee1e4b5c2c2d24df)

#### プロセス間同期

- [Go言語でプロセス間同期処理 - Qiita](http://qiita.com/shanxia1218/items/7fb15f50ec645f114bc7) : Windows の Mutex を使ってプロセス間通信を行う

### エラーハンドリング

- [または私は如何にして例外するのを止めて golang を愛するようになったか — KaoriYa](http://www.kaoriya.net/blog/2014/04/17/)
- [Big Sky :: golang で複数のエラーをハンドリングする方法](http://mattn.kaoriya.net/software/lang/go/20140416212413.htm)
- [DSAS開発者の部屋:Go ではエラーを文字列比較する？という話について](http://dsas.blog.klab.org/archives/go-errors.html) : エラーハンドリングには，定数との比較， conversion 構文による型の比較，エラー文字列の比較がある
- [panicはともかくrecoverに使いどころはほとんどない - Qiita](http://qiita.com/ruiu/items/ff98ded599d97cf6646e)

### ビルド時に情報を各種埋め込みたい

- [Go言語: ビルド時にバージョン情報を埋め込みたい - Qiita](http://qiita.com/suin/items/d643a0ccb6270e8e3734)
- [Golangビルド時に、サブパッケージ内の変数をいじる - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/09/08/211626)
- [Goでビルドバージョン情報を参照できるようにする(Go1.5) - Qiita](http://qiita.com/reiki4040/items/6b32370532c3eafe1f0e)
- [go-bindata でコンパイル時にリソースを埋め込んじゃおう！ - Qiita](http://qiita.com/ikawaha/items/c02d84cfd00f8f442500)
- [ソースを実行ファイルに埋め込む方法 - Qiita](http://qiita.com/ymko/items/c2e3c8fe25bce425136d)
- [Golangで静的ファイルをバイナリに含めるライブラリを書いてみた - Qiita](http://qiita.com/konohazuku/items/131b251a5fa29213ac5e)

## 開発ツールおよびパッケージ

### デバッガ

- [golang でビルド時に最適化をオフにする - tetsuok の旅 blog](http://tetsuok.hatenablog.com/entry/2012/07/01/062325) : gdb でデバッグする際は最適化を off にするといいという話

### Go 言語のテスト・フレームワーク

- [Go の Test に対する考え方 - Qiita](http://qiita.com/Jxck_/items/8717a5982547cfa54ebc)
- [Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)
- [golang 1.4で追加されたtestingの便利機能(テストの初期化とお片づけ) - Qiita](http://qiita.com/umisama/items/0d589cca7e89b89c29a8)
- [gojiのレスポンス結果をテストする - Qiita](http://qiita.com/taizo/items/11897f6284159919f65a)
- [Go Mockでインタフェースのモックを作ってテストする #golang - Qiita](http://qiita.com/tenntenn/items/24fc34ec0c31f6474e6d)
- [Go でベンチマーク - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20131123/1385189088)
- [go言語でベンチマーク - Qiita](http://qiita.com/Mulyu/items/ed585f2777496f29a725)
- [プロダクト開発でのGoのテストとモック活用事例 - Qiita](http://qiita.com/peroli-hirayama/items/f1419db7264fa9f9fe8f)

### Go 言語のドキュメント・フレームワーク

[godoc]

- [Go で godoc を使えるようにする〜godoc のインストール〜 - Qiita](http://qiita.com/megu_ma/items/2066aef2f8c7f0ce2cc3)
- [Go言語のコードレビュー | SOTA](http://deeeet.com/writing/2014/05/26/go-code-review/)
- [Go コードのレビュー時によくされるコメント - build error](http://papaeye.tumblr.com/post/92328649161/go)
- [testingパッケージのExamplesについて - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20131101/1383285018)
- [GoのExampleテストが便利 : swdyh](http://swdyh.tumblr.com/post/55771477125/go-example)
- [godoc.org への掲載方法を調べた - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20140225/1393302743)

[godoc]: http://godoc.org/golang.org/x/tools/cmd/godoc "godoc - GoDoc"

### Refactoring

リファクタリングには gofmt よりも gorename が使えるという話。

- [Big Sky :: golang のリファクタリングには gofmt ではなく、gorename を使おう。](http://mattn.kaoriya.net/software/lang/go/20150113141338.htm)
- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。

### クロス環境

- [Goはクロスコンパイルが簡単 - unknownplace.org](http://unknownplace.org/archives/golang-cross-compiling.html)
- [Go のクロスコンパイル環境構築 - Qiita](http://qiita.com/Jxck_/items/02185f51162e92759ebe)
- [golang でのクロスコンパイルの留意事項 — KaoriYa](http://www.kaoriya.net/blog/2015/03/06/) : Windows 環境でクロス環境を構築する際の注意点。
- [Golang + Raspberry Pi + LPS331AP で気圧・温度を測定してみた - Qiita](http://qiita.com/yanolab/items/5a6dfb3c07c94f7c760d) : Raspberry Pi 用のクロス環境例。
- [Gobot - Golang framework for robotics, physical computing, and the Internet of Things](http://gobot.io/)
- [Go1.5はクロスコンパイルがより簡単 | SOTA](http://deeeet.com/writing/2015/07/22/go1_5-cross-compile/)
- [MacOS X でGo言語のクロスコンパイルを試したらハマった - Qiita](http://qiita.com/ttsuzo/items/64e29dd7caa635ac7863) : [gox](https://github.com/mitchellh/gox) を使う方法
- [Goで64bitと32bitの実行ファイルを同一Windows機で作成するために講じたこと - Qiita](http://qiita.com/zetamatta/items/e44961a8bcbb2578cfe7)
- [Travis-CI で Go の Windows 用バイナリを Github release に登録する - Qiita](http://qiita.com/methane/items/f8c5a5f2209739daf44e)
- [gopherjs + electron テスト - Qiita](http://qiita.com/shizu/items/c8a28e0d2299868dafa9) : [`gopherjs/gopherjs`](https://github.com/gopherjs/gopherjs) を使って Go のコードから javaScript コードを生成できるらしい
- [Raspberry PI ２ 用の consul を作る (201512版 - Qiita](http://qiita.com/rerofumi/items/d6a8ba08270acb61b31c) : Raspberry PI 上でビルドするより Linux のクロス環境を使ったほうが速いらしい

### C 言語との連携

[cgo]

- [cgoでGoのコードからCの関数を利用する - 1000ch.net](https://1000ch.net/posts/2014/c-in-golang-with-cgo.html)
- [cgoでGolangとC++ライブラリをリンクするとき、何が起きているのか - beatsync.net](https://beatsync.net/main/log20141022.html)
- [GO 1.5 と C++ を SWIG でブリッジさせる方法 - Qiita](http://qiita.com/Satachito/items/5a0d7dd228d3272e0780)

[cgo]: https://golang.org/cmd/cgo/ "cgo - The Go Programming Language"

## Integrated Development Environment

- [EclipseでGoプログラミング！ GoClipseのインストールとGojiフレームワークを使ったWeb APIの作成 （1/6）：CodeZine](http://codezine.jp/article/detail/8374)
- [WindowsでGolang開発環境構築　IntelliJ IDEA - Qiita](http://qiita.com/ngsm3/items/67620fc4e39219235a23)
    - [第1回　Android Studio，そしてベースとなる「IntelliJ IDEA」とは何か？：Android Studio最速入門～効率的にコーディングするための使い方｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/serial/01/android_studio/0001) : IntelliJ IDEA についての解説
- [Go入門@環境構築編(IntelliJ IDEA 14.1.4) - Qiita](http://qiita.com/makoto2468/items/3ced77de947ea997e72f)

### ATOM で Go

- [AtomでGoを書く環境を整える（Windows） - Technically, technophobic.](http://notchained.hatenablog.com/entry/2014/09/20/104829)
- [AtomでのGo言語開発環境セットアップ - Qiita](http://qiita.com/MakoTano/items/3d807a96c3933ac8aa13)

### Visual Studio Code で Go

- [VisualStudioCode - Visual Studio CodeでGo言語の設定 - Qiita](http://qiita.com/evalphobia/items/f68396d573c7caf2065b)
- [VSCode+Golang](https://gist.github.com/llaughlin/16305f0b91356e93a1c4) : task.json の設定

### Build Tool for Golang

[constabulary/gb](https://github.com/constabulary/gb) を使ってプロジェクトベースの環境構築

- [golang - gbを知ったのでgojiを使ったWEBアプリケーションプロジェクトを管理してみた - Qiita](http://qiita.com/shinofara/items/ac0591fef95c2c6e936e)
- [Building Go projects with gb - Supermighty](https://walledcity.com/supermighty/building-go-projects-with-gb)
- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[FiloSottile/gvt](https://github.com/FiloSottile/gvt) というのがあるらしい。

[Masterminds/glide](https://github.com/Masterminds/glide) と Go 1.5 の Vendoring 機能を組み合わせてパッケージ管理できる。

- [glide - パッケージ管理のお困りの方へ - - Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)

### direnv で開発環境を切り替える

[direnv - unclutter your .profile](http://direnv.net/)

- [direnv/direnv](https://github.com/direnv/direnv)
- [direnvで解決するGOPATHの3つの問題点 - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/06/18/010449)
- [改めて、direnvを使いましょう！ - HDE BLOG](http://blog.hde.co.jp/entry/2015/02/27/182117)
- [さくら - homeにgolang, direnv とvirtualenvを入れて動かす - Qiita](http://qiita.com/aminamid/items/5a0e9461385c80d0c8a6)

### Continuous Integration

- [Go + Travis CI + Coveralls でCI環境を作る - Qiita](http://qiita.com/dmnlk/items/3fb4e0abb98e39fee275)
- [GithubにあるリポジトリをTravis CI連携する手順 #junitbook - くりにっき](http://sue445.hatenablog.com/entry/2013/06/01/170607)
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [golangでTravis CIを使ってクロスコンパイルするときにハマったところ #golang #travisci - uchimanajet7のメモ](http://uchimanajet7.hatenablog.com/entry/2015/03/20/211352)
- [Go言語のビルド生活を drone.ioで幸せに暮らす #golang - Qiita](http://qiita.com/atotto/items/b796c31c1755dbec13db)
- [Go+Webアプリケーション+CircleCIで静的解析・テスト・バイナリリリースを効率良く行なう - Qiita](http://qiita.com/kaneshin/items/163626c09c1ad9818c6c)

## パッケージやサンプルコード

- [Go用のGoogle製のGUIツールキットgxuiのインストール(Windows版) - Qiita](http://qiita.com/sago35/items/cc9ed3dc38d0b2f19bf9)
- [goでLチカの練習　その２ - Qiita](http://qiita.com/ohisama@github/items/bfc1eb6407cbdfebbd18)
- [Revel templatesを使ったサンプルアプリケーション - Qiita](http://qiita.com/rubytomato@github/items/638299aabb7922cbef59)
- [goでwindowsでwindow - Qiita](http://qiita.com/ohisama@github/items/20bc61175ce4a33b7365)
- [Go言語でパッケージを作成して世界に公開する方法 at ミネルヴァの梟は黄昏とともに飛び始める（山下 大介 公式ブログ）](http://blog.daisukeyamashita.com/post/1209.html) : パッケージの作り方なんだけど、情報が古すぎた orz
- [goでwindowsでキースキャン - Qiita](http://qiita.com/ohisama@github/items/9f05679f25cfc9c3ecfc)
- [Goのencoding/jsonでタグが反映されなくてハマったしょうもない話 - Qiita](http://qiita.com/modal_soul/items/83b0930d90d44e006768)
- [gocron でジョブスケジューリング - Qiita](http://qiita.com/mitsuse/items/8669bf54d2310b3e68a1) : [gcarlescere/scheduler](https://github.com/carlescere/scheduler) のほうがおすすめらしいw
- [Google ChromeのテキストエリアをEmacsで編集する'Edit with Emacs'から任意のエディタを起動するデーモンをGo 1.4 for Windowsで書いてみたわけだが、エディタがブラウザの後ろに出てしまってダメかもしれない - Qiita](http://qiita.com/zetamatta/items/51b0f45496e5143e2e63)
- [golangでprivateなエイリアスのポインタを元の型に戻す - Qiita](http://qiita.com/shibukawa/items/9db22c9684cc0586b737)
- [Go の expvar パッケージを使ってアプリケーションのメトリクスを公開する - Qiita](http://qiita.com/methane/items/8f56f663d6da4dee9f64)
- [ジェネレートしたコードを gofmt / goimports する - Qiita](http://qiita.com/minodisk/items/d96a0673223f36315ce7)
- [Go言語: ファイルの存在をちゃんとチェックする実装? - Qiita](http://qiita.com/suin/items/b9c0f92851454dc6d461)
- [Golangで、ファイル一覧取得（最新順出力） - Qiita](http://qiita.com/shinofara/items/e5e78e6864a60dc851a6)
- [Go小ネタ: 正規表現を全角半角問わずマッチするよう変換する - Qiita](http://qiita.com/tutuming/items/d8aaaf5442d84a7961e1)
- [Goのバッチで統計を取得するAPIを用意しておくと便利 - Qiita](http://qiita.com/sudix/items/c542e1b59bc94dc741e3)
- [Mithril＋golang Gin を試す - Qiita](http://qiita.com/masatsugumatsus/items/e28254ff52963705ce7f)
- [大学入試問題をGoで解いてみる - Qiita](http://qiita.com/qube81/items/c47b9e3ea8d028e95588) : [math/big](https://golang.org/pkg/math/big/) パッケージを使って大きな数を計算する。
- [cmd.Envを設定してexecしたらコケた - Qiita](http://qiita.com/8845musign/items/5c4b32f82c2df08acd93) : [`exec`](https://golang.org/pkg/os/exec/) パッケージを使う際には環境変数に注意
- [golangで数独を解いた - Qiita](http://qiita.com/ciruzzo/items/144bc1874947441f9fb8)
- [Go言語でBigQueryのクエリを実行してみる - Qiita](http://qiita.com/najeira/items/8310fecf4b70c918f855)
- [sliceのシャッフル - Qiita](http://qiita.com/sugyan/items/fd7138a756c1a409f5fd) : [Fisher–Yates shuffle](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) というアルゴリズムらしい
- [今後イケそうなデスクトップGUIフレームワーク - Qiita](http://qiita.com/hachi8833/items/463264f531474a856064) : [go-thrust](https://github.com/miketheprogrammer/go-thrust/) について言及
- [Goで外部コマンドをパイプして実行する - Qiita](http://qiita.com/yuroyoro/items/9358cd25b5f7fe9dd37f)
    - [Big Sky :: golang で UNIX コマンドパイプラインを扱う](http://mattn.kaoriya.net/software/lang/go/20151030131242.htm)
- [1l0/sumeragi](https://github.com/1l0/sumeragi) : 皇紀や元号を出力するパッケージ
- [RubyからGoの関数をつかう → はやい - Qiita](http://qiita.com/grj_achm/items/679b3f3af2cf377f0f02)
    - [Perl6からGoの関数をつかう → はやい? - Qiita](http://qiita.com/B73W56H84/items/20a67b74bb646d140f7d)
    - [GroovyからGoの関数を使う→はやい - Qiita](http://qiita.com/yujiorama_at_github/items/3f7cab906969764cc805)
    - [Big Sky :: RubyからGoの関数をつかわなくても再帰をやめる → はやい](http://mattn.kaoriya.net/software/20151106194958.htm)
- [Go ライブラリによる CGIプログラム内ルーチング - Qiita](http://qiita.com/jyagaimo_qiita_/items/93d195ca65982b75e205)
- [Go言語と暗号技術（AESからTLS） | SOTA](http://deeeet.com/writing/2015/11/10/go-crypto/)
    - [tcnksm/go-crypto](https://github.com/tcnksm/go-crypto) : サンプルコード
- [マイナンバーのチェックデジットをGoで計算する - Qiita](http://qiita.com/qube81/items/f66a38b28ec58bc5c4da)
- [golangでImageMagickを触りたい - Qiita](http://qiita.com/arc279/items/5f277aa5cce3de5247e5)
- [GoでJsonファイルを読み込んで構造体として扱う。 - Qiita](http://qiita.com/niiyz/items/3f522c0e5a32de916ec6)
- [BurntSushi/tomlを使ってハマったこと - Qiita](http://qiita.com/reiki4040/items/6556d4eba797329e9f51) : [BurntSushi/toml](https://github.com/BurntSushi/toml) にバグがあるという話
- [Go言語でRedshiftとつなぐ（というかただのPostgreSQL） - Qiita](http://qiita.com/otiai10/items/83b186596897705ce392)
- [GolangのOpenGL事情(WebGLも含むよ) - Qiita](http://qiita.com/shibukawa/items/58f6a421462b93dec471)
- [goのmgoでfindするときの処理 - Qiita](http://qiita.com/enokidoK/items/a3aff4c05e494b004ef8)
- [Go と reflect と generate - Qiita](http://qiita.com/naoina/items/7966f73f3a807b3d25d6)
- [Go言語でバイトニックソート実装してみた - Qiita](http://qiita.com/nyamadandan/items/2c82011801b148c98e52)
- [Windows+GoでNFC/Felicaにアクセスしてみた - Qiita](http://qiita.com/mau4x/items/424fe7964e70a3a99965)

### コマンドライン・ツール

- [Go言語によるCLIツール開発とUNIX哲学について - ゆううきブログ](http://yuuki.hatenablog.com/entry/go-cli-unix)
- [開発者から見た UNIX 哲学とコマンドラインツールと Go言語 - TELLME.TOKYO](http://tellme.tokyo/post/2015/06/23/unix_cli_tool_go/) （[Qiita 版](http://qiita.com/b4b4r07/items/df660d82e2de715acda5)）

[flag](http://golang.org/pkg/flag/)

- [Go - コマンドライン引数 - Qiita](http://qiita.com/uokada/items/f0e069a751679dcf616d)
- [Go言語のflagパッケージを使う - uragami note](http://ryochack.hatenablog.com/entry/2013/04/17/232753)
- [Go を使ってコマンドラインツール wordc を作ってみた - Qiita](http://qiita.com/flaflasun/items/df5ebb057697da062a08) : [codegangsta/cli](https://github.com/codegangsta/cli) についての言及あり
- [golang - GoでCLIツール作るのに便利そうなパッケージを集めてみました - Qiita](http://qiita.com/isaoshimizu/items/71dd2ca2a08ddb607e31)
- [C言語とGo言語で標準出力が端末を参照しているかどうかを判定する - uragami note](http://ryochack.hatenablog.com/entry/2013/07/15/232207)
- [Go言語のCLIツールのpanicをラップしてクラッシュレポートをつくる | SOTA](http://deeeet.com/writing/2015/04/17/panicwrap/)
- [flag 並にシンプルでより強力な CLI パーサ kingpin の紹介 - Qiita](http://qiita.com/kumatch/items/258d7984c0270f6dd73a)

### 入出力処理

[io](http://golang.org/pkg/io/), [bufio](http://golang.org/pkg/bufio/)

- [Go ファイルや標準入力から一行ずつ読み込む - Qiita](http://qiita.com/hnakamur/items/a53b701c8827fe4bfec7)
- [Goでファイルを読んで別のgoroutineに渡す - Qiita](http://qiita.com/ono_matope/items/5c8cfce81933c5eb9fd0)
- [「連結されたgzipを1行ずつ見る」をgolangでやったらハマった - Qiita](http://qiita.com/kroton/items/431e6dad9e5e4dbc44cf) : [compress/gzip](https://golang.org/pkg/compress/gzip/) と入出力処理の関係
- [bufio.Scannerのend-of-line判断を変更してみる - Qiita](http://qiita.com/curious-eyes/items/2d4b6c20ea47e3efc47b)
- [KOBE GDG: Go言語　バイナリファイルを扱う](http://kobegdg.blogspot.jp/2013/05/go.html) : 任意のオブジェクトをバイト列に変換してファイルに格納
- [ファイル書き込みの度にファイルを開いたらどれくらい遅いのか - Qiita](http://qiita.com/catatsuy/items/41d3c49248b517b5af96)

### Logging

- [Golang logging library - Qiita](http://qiita.com/kosuda/items/988c505c2abc5321aba8)
- [go言語におけるロギングについて](http://blog.satotaichi.info/logging-frameworks-for-go/)
- [Golangで簡単にログを吐くことを考える - Qiita](http://qiita.com/Ets/items/49e8f781990a3b0b3821) : [seelog](https://github.com/cihub/seelog) について解説している。私は XML には全くアレルギーがないので無問題
- [Go言語でdebugログの実現方法 - Qiita](http://qiita.com/sbjib/items/2cef51e572eef0795bc2)
- [loggingについて話そう - Qiita](http://qiita.com/methane/items/cedbf546ae2db8a63c3d)

### Web Microframework for Golang

- [Goji](https://goji.io/)
- [Big Sky :: golang で最近お気に入りの WAF「Goji」](http://mattn.kaoriya.net/software/lang/go/20141021134209.htm)
- [gojiのMiddlewareの使い方 - Qiita](http://qiita.com/reiki4040/items/a038f1b99e0caee97d3e)
- [Gojiを使ってWebビーコン作る - Qiita](http://qiita.com/sys_cat/items/1b8581de1344cc5db6bb)
- [Google App EngineでGoのウェブアプリケーションをまず動かしてみる - Qiita](http://qiita.com/taizo/items/845fcfc58cfd0b30020a)
- [Go言語(Go-Json-Rest)でAPIサーバーを立てるときのCORS設定 (Basic認証機能付きも) - Qiita](http://qiita.com/n0bisuke/items/45ab414fc11959fc27c7)
- [Goでwebサーバー作るときに考えたこと - Qiita](http://qiita.com/taizo/items/e6597c66c3494d545686)
- [Golang の net/url で # を含む文字列を Parse() する - Qiita](http://qiita.com/tchssk/items/cb208f9ccd0a1819bbfa)
- [Golang で OpenStreetMap ファイル（osm.pbf）の読み込み（osmpbf 利用） - Qiita](http://qiita.com/kkdd/items/bd653e3d02546d1aa340)
- [GoのWEBアプリケーション運用について - Qiita](http://qiita.com/masahikoofjoyto/items/f60188f4252e455541d4)
- [go-qmlのWebViewでローカルWebサイトを表示するときの注意 - Qiita](http://qiita.com/hachi8833/items/315642ceecb378cabcb4)
- [http.Clientのタイムアウトの時間を変更 - Qiita](http://qiita.com/eielh/items/2e5fabb707355253b187)
- [golangでhttpを監視するscriptを書いて結果をslackにpostする - Qiita](http://qiita.com/kenjiszk/items/7ae842415ec392822612)
- [Go言語(Go-Json-Rest)のCORSでのハマり。Safariだけでハマった話。 - Qiita](http://qiita.com/n0bisuke/items/65c4a473a1fdbbf931f0)
- [GoでGoogleCalendarAPIv3を叩く（APIキー） - Qiita](http://qiita.com/yodatomato/items/8013f728bbf4358e9425)
- [Golang で iso-2022-jp メールのデコード - Qiita](http://qiita.com/curious-eyes/items/3dae99e5e0feb6b9f642)
- [Go言語の練習用にTwitterのOAuth認証をフルスクラッチしてみた - Qiita](http://qiita.com/mpyw/items/cb0f824d618d8fed384e)
- [Goとtesseractで簡易OCRサーバを作る - Qiita](http://qiita.com/fumizp/items/63243cf418d27898f208)
- [Go言語で簡単にHTTPリクエストを送ってJSONをパースするサンプル - Qiita](http://qiita.com/ikuwow/items/c8f494bbd16adf6db142)
- [UnixListener.Closeでソケットファイルが消えて困っている - Qiita](http://qiita.com/hiratara/items/0f0b6103a0dc9280cea9)
- [Goフレームワークのパフォーマンス比較 - Qiita](http://qiita.com/ToruFukui/items/eb0d3593b20a5e2f15c6)
- [一定時間だけ立ち上がるサーバーを書いた - Qiita](http://qiita.com/uokada/items/c30e26cd8bdee7dfe7eb)
- [Big Sky :: golang のミドルウェアとして組み込むだけでパフォーマンス改善が見込める「HTTP Coala」](http://mattn.kaoriya.net/software/lang/go/20151120113437.htm)
- [Go付属のツールでオレオレ証明書を生成する - Qiita](http://qiita.com/ono_matope/items/6c0de3e31642dfd17695) : なるほど（笑）
- [Go http.RoundTripper 実装ガイド - Qiita](http://qiita.com/tutuming/items/6006e1d8cf94bc40f8e8)
- [Go で静的 HTTP サーバを作る - Qiita](http://qiita.com/skitaoka/items/a2b55cb08060aa8d6a52)

### 組み込み開発

- [goでwindowsでRS232C - Qiita](http://qiita.com/ohisama@github/items/e6fa2bd1527c257bb9c9)
    - [ohisama/serial](https://github.com/ohisama/serial) (forkd from [lnmx/serial](https://github.com/lnmx/serial))
    - [goでwindowsでRS232C その２ - Qiita](http://qiita.com/ohisama@github/items/12dccdcdfc5082c22e72)
- [Goでのシリアル通信でハマった事 - Qiita](http://qiita.com/tomoya0x00/items/d957dc00682c57f96771)
- [go 1.5でgomobile(android) - unokun’s blog](http://unokun.hatenablog.jp/entry/2015/08/01/150628)
- [gomobileでiOSアプリをビルドする手順まとめ - GolangRdyJp](http://golang.rdy.jp/2015/09/21/ios-gomobile/)
- [gomobileで日本語フォントを扱ってみる - Qiita](http://qiita.com/bowz_standard/items/5a9c987f9242777fff30)

## その他 四方山話

- [動的言語だけやってた僕が、38日間Go言語を書いて学んだこと - Qiita](http://qiita.com/suin/items/22662f43b6a6e8728798)
- [Go言語で幸せになれる10のテクニック - Qiita](http://qiita.com/ksato9700/items/6228d4eb6d5b282f82f6)
- [これからGoを始める人のためのTips集 | The Wacul Blog](http://blog.wacul.co.jp/blog/2014/08/22/go/) : ちゃんと開発環境を整備したい場合には参考になる
- [Golang ファーストインプレッション - Qiita](http://qiita.com/mrpepper/items/95c428f2b3c25be6b3e2)
- [2014年夏、カヤックは、Go言語を積極的に推進していきます。 ｜ニュース｜面白法人カヤック](http://www.kayac.com/news/2014/07/golang)
- [(go report) Goが本当はすごかったので, 謝罪する - テストステ論](http://akiradeveloper.hatenadiary.com/entry/2014/07/22/205515)
- [Golang のインスコでハマった件 - Qiita](http://qiita.com/TakaakiFuruse/items/0fec78e5ecbcbe4337e8)
- [Go Conference 2015 summer - Qiita](http://qiita.com/t-sato/items/a5d1a309733e765533ce)
- [最新versionのgolangをぶち込む方法について - Qiita](http://qiita.com/yamadagenki/items/5032cf853f6b136b533f)
    - [7 Easy Steps to Install Go (Golang) on Ubuntu - HostingAdvice.com](http://www.hostingadvice.com/how-to/install-golang-on-ubuntu/)
- [Heroku、Go言語の正式サポートを発表 － Publickey](http://www.publickey1.jp/blog/15/herokugo.html)

> 最近ではGo言語の採用例が増えてきており、例えばDropboxは性能が重要な部分ではPythonに代わりGoで記述することを昨年11月に表明しており、オープンソースのPaaS基盤であるCloud Foundryも主要部分での開発でGo言語の採用を進めていると言われています。

- [Google App Engineも「Go言語」の正式サポートを発表 － Publickey](http://www.publickey1.jp/blog/15/google_app_enginego_1.html)
- [Sensuでネットワーク監視やってみた - Qiita](http://qiita.com/hiconyan/items/7656e9fb2d5bf5c794be)
- [i18n4go - developerWorks Open](https://developer.ibm.com/open/i18n4go/) : IBM の [developerWorks Open](https://developer.ibm.com/open/) プロジェクトのひとつ。「[プログラムを国際化するための汎用ツール](http://www.atmarkit.co.jp/ait/articles/1507/23/news058.html)」らしい。

### Go 1.5 に関する話題

- [GVM で go1.5rc1 のインストール - Qiita](http://qiita.com/msaito3/items/3aef86e9864990b16b4c)
- [goを1.5にアップデートして1.4とベンチを取る - Qiita](http://qiita.com/masahikoofjoyto/items/4ced298989e6ab346f15)
- [Go 1.3 から 1.5 へのアップデートでエラー - Qiita](http://qiita.com/taji-taji/items/4c43e126e67d65a219e3) : 古いバージョンからアップデートする際は，いったん 1.4 に上げてから 1.5 にアップデートするとよい
- [Big Sky :: golang 1.5 の internal パッケージの使い方。](http://mattn.kaoriya.net/software/lang/go/20150820102400.htm)
    - [「golang 1.5 の internal パッケージの使い方。」を試してみた - Qiita](http://qiita.com/qt-luigi/items/d0f52b3b0906b35e6027)

### はじめての Go 言語 (on Windows)

[Qiita](http://qiita.com/) で書いてる拙文の目次。

1. [はじめての Go 言語 (on Windows)](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa) : インストールと動作確認（“Hello World” まで）
1. [はじめての Go 言語 (on Windows) その2](http://qiita.com/spiegel-im-spiegel/items/047a9bd6436e6391ddd4) : 素数探索アルゴリズムで遊ぶ（slice, make, goroutine, channel）
1. [はじめての Go 言語 (on Windows) その3](http://qiita.com/spiegel-im-spiegel/items/a52a47942fd3946bb583) : get コマンドを使う
1. [はじめての Go 言語 (on Windows) その4](http://qiita.com/spiegel-im-spiegel/items/556166b6631c0369754f) : string と rune
1. [はじめての Go 言語 (on Windows) その5](http://qiita.com/spiegel-im-spiegel/items/e743d63ef5165d750eff) : 暦で遊ぶ（math, time）
1. [はじめての Go 言語 (on Windows) その6](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90) : パッケージ化について
1. [はじめての Go 言語 (on Windows) その7](http://qiita.com/spiegel-im-spiegel/items/64224f22ef17d916dc2d) : テストを書く
1. [はじめての Go 言語 (on Windows) その8](http://qiita.com/spiegel-im-spiegel/items/5f9e96f226f46089388f) : パッケージ化したのならドキュメントを書きましょう
1. [はじめての Go 言語 (on Windows) その9](http://qiita.com/spiegel-im-spiegel/items/ef15a48542e043b32c99) : プロジェクト・ベースのビルド・ツールを使ってみる
1. [はじめての Go 言語 (on Windows) その10](http://qiita.com/spiegel-im-spiegel/items/5d2878596360af8dd753) : 最終回。コマンドライン・ツールを作ってみる。エラーハンドリングについてもちょっとだけ

ちなみに「はじめての...」以降の Go 言語に関する拙文は以下の通り。

- [tcnksm/gcli を使った golang によるコマンドライン・ツール開発について](http://qiita.com/spiegel-im-spiegel/items/69d1166225d88b1faf66)
- [そろそろ真面目に Golang 開発環境について考える — GOPATH 汚染問題](http://qiita.com/spiegel-im-spiegel/items/73ebc684b5807277b7e2)
- [Golang による文字エンコーディング変換](http://qiita.com/spiegel-im-spiegel/items/2e475b48226330aa5570)
- [Golang の文字列連結はどちらが速い？](http://qiita.com/spiegel-im-spiegel/items/16ab7dabbd0749281227)
- [そろそろ真面目に Golang 開発環境について考える — Internal Packages と Vendoring](http://qiita.com/spiegel-im-spiegel/items/baa3671c7e1b8a6594a9)
