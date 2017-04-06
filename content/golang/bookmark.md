+++
date = "2015-09-11T17:58:42+09:00"
description = "本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。"
update = "2017-04-06T10:19:32+09:00"
draft = false
tags = ["golang", "bookmark"]
title = "Go 言語に関するブックマーク"

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

本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。
なおサンプルコードやパッケージを紹介しているページのブックマークは「[サンプルコード，パッケージ関連]({{< relref "golang/bookmark2.md" >}})」に移動した。

## 公式サイト

- [The Go Programming Language](https://golang.org/) : 2017年2月16日時点での最新は 1.8
    - [git repositories (Google)](https://go.googlesource.com/)
    - [git repositories (GitHub)](https://github.com/golang) : mirror
- [golang-jp - The Go Programming Language](http://golang-jp.org/) : 本家の日本語訳サイト。（[golang.jp](http://golang.jp/) は内容が古いので参考にしない方がいい，らしい）

### Go 1.5 Released.

- [Go 1.5 is released - The Go Blog](https://blog.golang.org/go1.5)
- [Go 1.5 Release Notes - The Go Programming Language](https://golang.org/doc/go1.5)
    - [Go 1.4 "Internal" Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit) : Internal Packages は 1.5 で GOPATH まで拡張された
    - [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)
- [[security] Go 1.5.3 is released - Qiita](http://qiita.com/spiegel-im-spiegel/items/83b53109f54f8fb62c1c)

### Go 1.6 is released

- [Go 1.6 is released - The Go Blog](https://blog.golang.org/go1.6)
- [Go 1.6 Release Notes - The Go Programming Language](https://golang.org/doc/go1.6)
    - [Go1.6の新機能 - Qiita](http://qiita.com/ksato9700/items/5505e506c20b6048c218)
- [Go 言語 1.6.1 および 1.5.4 のセキュリティ・アップデート — しっぽのさきっちょ | text.Baldanders.info]({{< relref "remark/2016/04/golang-1_6_1-released.md" >}})

### Go 1.7 is released

- [Go 1.7 is released - The Go Blog](https://blog.golang.org/go1.7)
- [Go 1.7 Release Notes - The Go Programming Language](https://golang.org/doc/go1.7)
- [Security Release Go 1.7.4 and 1.6.4 — しっぽのさきっちょ | text.Baldanders.info]({{< relref "remark/2017/01/go-1_7_5-released.md" >}})
- [Go 1.7.5 がリリース — しっぽのさきっちょ | text.Baldanders.info]({{< relref "remark/2017/01/go-1_7_5-released.md" >}})

### Go 1.8 is released

- [Go 1.8 is released - The Go Blog](https://blog.golang.org/go1.8)
- [Go 1.8 Release Notes - The Go Programming Language](https://golang.org/doc/go1.8)

## 言語仕様に関すること

- [はじめてのGo―シンプルな言語仕様，型システム，並行処理：特集｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/feature/01/go_4beginners)
- [Golangの基本文法をおさえてみる - Qiita](http://qiita.com/kazusa-qooq/items/40f9ea3e72406d845b10)
- [忙しい人のためのA Tour of Go - Qiita](http://qiita.com/makoto_kw/items/0638c0af1002647e3f7a)
- [Golang の 数値型 - Qiita](http://qiita.com/tanaka0325/items/9c61a022cd32be0c65a6)
- [Go言語の型とreflect - Qiita](http://qiita.com/atsaki/items/3554f5a0609c59a3e10d)
- [Goのfor rangeで思った値が取れなかった話 - Qiita](http://qiita.com/modal_soul/items/e49480e5692597fda975) : ちょっとしたミス
- [unsafe が unsafe なケース (1) - Qiita](http://qiita.com/kwi/items/185bb3fe0d60ca765ab0)
    - [unsafe が unsafe なケース (2) - Qiita](http://qiita.com/kwi/items/d06f49c9cf7e5ace8692)
- [Big Sky :: Go言語でインタフェースの変更がそれ程問題にならない理由](http://mattn.kaoriya.net/software/lang/go/20130919023425.htm)
- [import 書き方まとめ - Qiita](http://qiita.com/taji-taji/items/5a4f17bcf5b819954cc1)
- [ErlangとGolangを比較してみる - Qiita](http://qiita.com/soranoba/items/68d57b4635a2917f3c73)
- [GoのEnumイディオム - Qiita](http://qiita.com/awakia/items/c81c7816b9aea5422250) : あらかじめ Enum 用の type を作成し、その type に対する `String()` メソッドを定義する
- [Big Sky :: golang では for ループの中で defer してはいけない。](http://mattn.kaoriya.net/software/lang/go/20151212021608.htm) : ループ内で defer が必要になるということは refactoring のチャンス
- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)
- [埋込みとインタフェース #golang - Qiita](http://qiita.com/tenntenn/items/bba69d84a1e0b67c4db8)
- [Goで再帰使うと遅くなりますがそれが何だ - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/02/23/165341)
- [Go言語(Golang) はまりどころと解決策](http://www.yunabe.jp/docs/golang_pitfall.html)
- [Golang: nil Pointer Receiverの話 - Qiita](http://qiita.com/stsn/items/73714caf8458b1d973f2)
- [Go 言語の値レシーバとポインタレシーバ | Step by Step](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [Big Sky :: Golang のオフィシャルが提供するインタフェースまとめ](http://mattn.kaoriya.net/software/lang/go/20140501172821.htm)
- [Visualizing Concurrency in Go · divan's blog](http://divan.github.io/posts/go_concurrency_visualize/)
- [Go の定数の話 - Qiita](http://qiita.com/hkurokawa/items/a4d402d3182dff387674)
- [Goにatexitやグローバルなデストラクタがない理由 - Qiita](http://qiita.com/ruiu/items/22910a4bae6cb716a391)
- [init関数のふしぎ #golang - Qiita](http://qiita.com/tenntenn/items/7c70e3451ac783999b4f)
    - [packageに複数のinitがあるときの挙動 - Qiita](http://qiita.com/astronoka/items/aa2f271d280863cedf5e)
- [Goで任意のbuild tagsをつけてビルドファイルを切り替える - Qiita](http://qiita.com/ueokande/items/fac0d1219dbbc8f44db7)
- [Go と reflect と generate - Qiita](http://qiita.com/naoina/items/7966f73f3a807b3d25d6)
- [Golang 日付のフォーマットでハマった話 - Qiita](http://qiita.com/masa23/items/e781124a7e0305bc40c4)
- [Big Sky :: golang では変数の宣言位置が大事](http://mattn.kaoriya.net/software/lang/go/20170406003909.htm)

### 文字列操作または変換

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
- [Go言語 Gmailのsubjectの日本語文字化けに対応する - Qiita](http://qiita.com/yyoshiki41/items/79882e269ca6af4c2236)
- [Text normalization in Go - The Go Blog](https://blog.golang.org/normalization) : Unicode 正規化について
- [Go言語で文字列の変換(全角・半角、ひらがな・カタカナ)をする : Serendip - Webデザイン・プログラミング](http://www.serendip.ws/archives/6307)
- [Goで全角英数字を半角にする - Qiita](http://qiita.com/ktashiro/items/da5cbee3129acc74e5d7)
- [uint64型を[]bytes型に変換する - Qiita](http://qiita.com/joniyjoniy/items/cbfb7d5c49aec5bf63c0)

### 配列

[slice](http://golang.org/ref/spec#Slice_types), [map](http://golang.org/ref/spec#Map_types), および後述の [channel](http://golang.org/ref/spec#Channel_types) は組み込みの型だが内部構造と状態を持つため， [new](http://golang.org/ref/spec#Allocation) ではなく [make](http://golang.org/ref/spec#Making_slices_maps_and_channels) を使う。

- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang - go言語のslice操作をまとめてみた（shiftしたりpushしたり） - Qiita](http://qiita.com/egnr-in-6matroom/items/282aa2fd117aab9469bd)
- [sliceの重複チェックを高速化 - Qiita](http://qiita.com/hi-nakamura/items/5671eae147ffa68c4466)
- [Goのarrayとsliceを理解するときがきた - Qiita](http://qiita.com/seihmd/items/d9bc98a4f4f606ecaef7) : この説明は分かりやすい。おススメ

### 並行処理と並列処理

- [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
- [Go: 計算なしのFizzBuzz - Qiita](http://qiita.com/suin/items/eca21ed935115e5da2e8) : channel の説明するのにいいかも
- [Goのchannelの送受信用の型について - Qiita](http://qiita.com/yuki2006/items/3f90e53ce74c6cff1608)
- [Go言語のChannelは送信時にもブロックする - Qiita](http://qiita.com/hondata/items/64776c79063e93bea9ed) : 意外と見落とす channel 送信時のブロック
- [select loop の小ネタ - Qiita](http://qiita.com/Jxck_/items/da3ca2db58734a966cac)
- [Goのforとgoroutineでやりがちなミスとたった一つの冴えたgo vetと - Qiita](http://qiita.com/sudix/items/67d4cad08fe88dcb9a6d)
- [x/net/context の実装パターン - Qiita](http://qiita.com/tutuming/items/c0ffdd28001ee0e9320d) : [golang.org/x/net/context](https://godoc.org/golang.org/x/net/context) を使って並行処理を細かく制御。 Domain-Driven あるいは Context-Driven な設計でも使えそう。
- [Go言語でチャネルとselect - Qiita](http://qiita.com/najeira/items/71a0bcd079c9066347b4)
- [golangでシグナルを拾ってgracefulにgoroutineを停めたい - Qiita](http://qiita.com/arc279/items/c44d4a18a851ff454c64)
- [golang の channel のブロックがよくわからん - Qiita](http://qiita.com/arc279/items/bc55cdf436c544e91c05)
- [GoのChannelを使いこなせるようになるための手引 - Qiita](http://qiita.com/awakia/items/f8afa070c96d1c9a04c9)
- [Goでスレッド（goroutine）セーフなプログラムを書くために必ず注意しなければいけない点 - Qiita](http://qiita.com/ruiu/items/54f0dbdec0d48082a5b1) : `sync.Mutex` にも言及
- [GAE/Goでもgoroutine使おうぜ！というハナシ - Qiita](http://qiita.com/hogedigo/items/f0f409ee944c4b2107c3) : GAE/Go の API には非同期版がないので積極的に goroutine 使おうよ，という話
- [複数のgroutineが生えてるStructの安全な終了方法 - Qiita](http://qiita.com/shunsukeaihara/items/f9ef7c8d430f63d79d40)
- [Go言語の並行性を映像化する | プログラミング | POSTD](http://postd.cc/go_concurrency_visualize/)
- [Go言語でプロセス間同期処理 - Qiita](http://qiita.com/shanxia1218/items/7fb15f50ec645f114bc7) : Windows の Mutex を使ってプロセス間通信を行う
- [Goで並行処理のベンチマークをとる - Qiita](http://qiita.com/hhatto/items/c8eb987b0516a45db754)
- [go言語初心者が図を書きながらgo routineやgo channelを理解する(Part 1) - Qiita](http://qiita.com/Vermee81/items/88c9e28dec83d43e7883)
    - [go言語初心者が図を書きながらgoroutineやgo channelを理解する(Part2) - Qiita](http://qiita.com/Vermee81/items/30ad42a7265375b1b7b1)
- [Goroutineと channelから はじめるgo言語](http://www.slideshare.net/takuyaueda967/goroutine-channel-go)
- [Go言語の並行処理デザインパターン by Rob Pike 前編 - Qiita](http://qiita.com/tfutada/items/a289628d8b2d0af6152d)
    - [Go言語の並行処理デザインパターン by Rob Pike 後編 - Qiita](http://qiita.com/tfutada/items/dc8db894ac270a79ef2b)
- [Big Sky :: golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)

Go 言語で複数 CPU を使った並列処理を行うには明示的な設定が必要。

- [Go言語でCPU数に応じて並列処理数を制限する | SOTA](http://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/)
    - [やはり俺のgolangがCPUを一つしか使わないのはまちがっている。 - Qiita](http://qiita.com/ymko/items/554e3630fefdc29393a8)
- [Goでお手軽に行列の積を爆速並列計算 - Qiita](http://qiita.com/hama_du/items/fce4ee1e4b5c2c2d24df)

### Struct タグについて

- [Goのencoding/jsonでタグが反映されなくてハマったしょうもない話 - Qiita](http://qiita.com/modal_soul/items/83b0930d90d44e006768)
- [Go で struct のタグ情報を取得する - hiyosi's blog](http://hiyosi.tumblr.com/post/100922038678/go-%E3%81%A7-struct-%E3%81%AE%E3%82%BF%E3%82%B0%E6%83%85%E5%A0%B1%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Goのjson.Marshal/Unmarshalの仕様を整理してみる - I Will Survive](http://blog.restartr.com/2014/08/13/golang-json-marshal-unmarshal/)
- [GoでJsonファイルを読み込んで構造体として扱う。 - Qiita](http://qiita.com/niiyz/items/3f522c0e5a32de916ec6)
- [BurntSushi/tomlを使ってハマったこと - Qiita](http://qiita.com/reiki4040/items/6556d4eba797329e9f51) : [BurntSushi/toml](https://github.com/BurntSushi/toml) にバグがあるという話
- [GoでJSONの一部分を利用者が定義した構造体に読み込める便利な手法を見つけた - Qiita](http://qiita.com/hnakamur/items/ba363e82332d4dbdf34a)
- [golang は ゆるふわに JSON を扱えまぁす! — KaoriYa](https://www.kaoriya.net/blog/2016/06/25/)
- [Go言語でJSONに泣かないためのコーディングパターン - Qiita](http://qiita.com/minagoro0522/items/dc524e38073ed8e3831b)
- [Go 言語 1つの構造体に複数の validation を適応する - Qiita](http://qiita.com/iktakahiro/items/2e240147ca3188948a17)
- [Go で関数の引数用構造体のバリデーションと初期化をするメソッドをタグから生成するツールを作った - Qiita](http://qiita.com/nirasan/items/7b96f080cd286c324033)
- [Goでsql.NullStringを含む構造体をjson.Marshalする方法 - Qiita](http://qiita.com/oikyn/items/26808b03997797f1b342)
- [Go言語でJSON内の整数は10進数6桁しか表現できない - Qiita](http://qiita.com/toast-uz/items/52f0c86716493ad3ca12)

### エラーハンドリング

- [または私は如何にして例外するのを止めて golang を愛するようになったか — KaoriYa](http://www.kaoriya.net/blog/2014/04/17/)
- [Big Sky :: golang で複数のエラーをハンドリングする方法](http://mattn.kaoriya.net/software/lang/go/20140416212413.htm)
- [DSAS開発者の部屋:Go ではエラーを文字列比較する？という話について](http://dsas.blog.klab.org/archives/go-errors.html) : エラーハンドリングには，定数との比較， conversion 構文による型の比較，エラー文字列の比較がある
- [panicはともかくrecoverに使いどころはほとんどない - Qiita](http://qiita.com/ruiu/items/ff98ded599d97cf6646e)
- [go で AggregationException(.NET)的なことをする - Qiita](http://qiita.com/kyoh86/items/6cadd79de08cc597b65a) : ループ等でエラーを集約してからまとめて処理する方法
- [echoのAPIサーバ実装とエラーハンドリングの落とし穴 - Qiita](http://qiita.com/tienlen/items/5f2bcfe06eb83830ee55)
- [Golangのエラー処理とpkg/errors | SOTA](http://deeeet.com/writing/2016/04/25/go-pkg-errors/)
    - [Golangでエラー時にスタックトレースを表示する - Qiita](http://qiita.com/deeeet/items/dacc71932393ab35d9f8)

## 開発ツールおよびパッケージ

### デバッガ

- [golang でビルド時に最適化をオフにする - tetsuok の旅 blog](http://tetsuok.hatenablog.com/entry/2012/07/01/062325) : gdb でデバッグする際は最適化を off にするといいという話
- [Go言語のトラブルシューティング用機能](http://www.slideshare.net/satorutakeuchi18/go-53685632)
- [Go で利用できるプロファイリングツール pprof の読み方 - Qiita](http://qiita.com/ikawaha/items/e3b35f09fb49e9217924)
- [Goでfunctionが実行された順番を追いかける - sgykfjsm.github.com](http://sgykfjsm.github.io/blog/2016/01/20/golang-function-tracing/)
- [Go言語でプリント文デバッグするときのTips - Qiita](http://qiita.com/ohac/items/0aa8eb6ff8ee5f599dcd)

### テスト・フレームワーク

- [Go の Test に対する考え方 - Qiita](http://qiita.com/Jxck_/items/8717a5982547cfa54ebc)
- [Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)
- [golang 1.4で追加されたtestingの便利機能(テストの初期化とお片づけ) - Qiita](http://qiita.com/umisama/items/0d589cca7e89b89c29a8)
- [gojiのレスポンス結果をテストする - Qiita](http://qiita.com/taizo/items/11897f6284159919f65a)
- [Go Mockでインタフェースのモックを作ってテストする #golang - Qiita](http://qiita.com/tenntenn/items/24fc34ec0c31f6474e6d)
- [Go でベンチマーク - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20131123/1385189088)
- [go言語でベンチマーク - Qiita](http://qiita.com/Mulyu/items/ed585f2777496f29a725)
- [プロダクト開発でのGoのテストとモック活用事例 - Qiita](http://qiita.com/peroli-hirayama/items/f1419db7264fa9f9fe8f)
- [Go言語でファジング | SOTA](http://deeeet.com/writing/2015/12/21/go-fuzz/)
- [GAE/GoとGojiの組み合わせでテストを書く - Qiita](http://qiita.com/yosukesuzuki/items/c9e5c19df97d2ad5595a)
- [`go test -count n -timeout t` - Qiita](http://qiita.com/AkihiroSuda/items/0fd83df29182d4f5cdef)
- [Big Sky :: Re: golangでIOへのテストを行う](http://mattn.kaoriya.net/software/lang/go/20160328114637.htm)
- [Golang におけるサブテストの並行処理実装について | eureka tech blog](https://developers.eure.jp/tech/go1_7-subtests/)
- [GAE/Goのテストを実行するために必要なこと - Qiita](http://qiita.com/ttyokoyama/items/5b99299ec112b580f03b)
- [Golangでテストしやすいコードをかく - Qiita](http://qiita.com/chisso/items/1dcc52f404b88d274f29)

### ドキュメント・フレームワーク

- [Go で godoc を使えるようにする〜godoc のインストール〜 - Qiita](http://qiita.com/megu_ma/items/2066aef2f8c7f0ce2cc3)
- [Go言語のコードレビュー | SOTA](http://deeeet.com/writing/2014/05/26/go-code-review/)
- [Go コードのレビュー時によくされるコメント - build error](http://papaeye.tumblr.com/post/92328649161/go)
- [testingパッケージのExamplesについて - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20131101/1383285018)
- [GoのExampleテストが便利 : swdyh](http://swdyh.tumblr.com/post/55771477125/go-example)
- [godoc.org への掲載方法を調べた - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20140225/1393302743)

### Refactoring

- [ジェネレートしたコードを gofmt / goimports する - Qiita](http://qiita.com/minodisk/items/d96a0673223f36315ce7)
- [Big Sky :: golang のリファクタリングには gofmt ではなく、gorename を使おう。](http://mattn.kaoriya.net/software/lang/go/20150113141338.htm) : リファクタリングには gofmt よりも gorename が使えるという話
- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。
- [golintと闘いたいけど心の折れてしまった勇者のための隠しダンジョン - Qiita](http://qiita.com/kyoh86/items/1f2022b63372b84f1a27)
- [gometalinter で楽々 lint - Qiita](http://qiita.com/spiegel-im-spiegel/items/238f6f0ee27bdf1de2a0) : 各種 lint を統合的に扱える

### Continuous Integration

- [Go + Travis CI + Coveralls でCI環境を作る - Qiita](http://qiita.com/dmnlk/items/3fb4e0abb98e39fee275)
- [GithubにあるリポジトリをTravis CI連携する手順 #junitbook - くりにっき](http://sue445.hatenablog.com/entry/2013/06/01/170607)
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [golangでTravis CIを使ってクロスコンパイルするときにハマったところ #golang #travisci - uchimanajet7のメモ](http://uchimanajet7.hatenablog.com/entry/2015/03/20/211352)
- [Go言語のビルド生活を drone.ioで幸せに暮らす #golang - Qiita](http://qiita.com/atotto/items/b796c31c1755dbec13db)
- [Go+Webアプリケーション+CircleCIで静的解析・テスト・バイナリリリースを効率良く行なう - Qiita](http://qiita.com/kaneshin/items/163626c09c1ad9818c6c)

### クロス環境

- [Goはクロスコンパイルが簡単 - unknownplace.org](http://unknownplace.org/archives/golang-cross-compiling.html)
- [Go のクロスコンパイル環境構築 - Qiita](http://qiita.com/Jxck_/items/02185f51162e92759ebe)
- [golang でのクロスコンパイルの留意事項 — KaoriYa](http://www.kaoriya.net/blog/2015/03/06/) : Windows 環境でクロス環境を構築する際の注意点。
- [Golang + Raspberry Pi + LPS331AP で気圧・温度を測定してみた - Qiita](http://qiita.com/yanolab/items/5a6dfb3c07c94f7c760d) : Raspberry Pi 用のクロス環境例。
- [Gobot - Golang framework for robotics, physical computing, and the Internet of Things](http://gobot.io/)
- [複数プラットフォームにGoアプリケーションを配布する | SOTA](http://deeeet.com/writing/2014/05/19/gox/)
    - [Go1.5はクロスコンパイルがより簡単 | SOTA](http://deeeet.com/writing/2015/07/22/go1_5-cross-compile/)
- [MacOS X でGo言語のクロスコンパイルを試したらハマった - Qiita](http://qiita.com/ttsuzo/items/64e29dd7caa635ac7863) : [gox](https://github.com/mitchellh/gox) を使う方法
- [Goで64bitと32bitの実行ファイルを同一Windows機で作成するために講じたこと - Qiita](http://qiita.com/zetamatta/items/e44961a8bcbb2578cfe7)
- [Travis-CI で Go の Windows 用バイナリを Github release に登録する - Qiita](http://qiita.com/methane/items/f8c5a5f2209739daf44e)
- [gopherjs + electron テスト - Qiita](http://qiita.com/shizu/items/c8a28e0d2299868dafa9) : [`gopherjs/gopherjs`](https://github.com/gopherjs/gopherjs) を使って Go のコードから javaScript コードを生成できるらしい
- [Raspberry PI ２ 用の consul を作る (201512版 - Qiita](http://qiita.com/rerofumi/items/d6a8ba08270acb61b31c) : Raspberry PI 上でビルドするより Linux のクロス環境を使ったほうが速いらしい
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [RaspberryPi1(2とzeroも)で動かすgolang製アプリをクロスコンパイル(onMac) - Qiita](http://qiita.com/m0a/items/d933982293dcadd4998c)

### C 言語との連携

- [cgoでGoのコードからCの関数を利用する - 1000ch.net](https://1000ch.net/posts/2014/c-in-golang-with-cgo.html)
- [cgoでGolangとC++ライブラリをリンクするとき、何が起きているのか - beatsync.net](https://beatsync.net/main/log20141022.html)
- [GO 1.5 と C++ を SWIG でブリッジさせる方法 - Qiita](http://qiita.com/Satachito/items/5a0d7dd228d3272e0780)
- [cgoを使ったCとGoのリンクの裏側 (1) - Qiita](http://qiita.com/yugui/items/e71d3d0b3d654a110188)
    - [cgoを使ったCとGoのリンクの裏側 (1) - Qiita](http://qiita.com/yugui/items/e71d3d0b3d654a110188#_reference-b69de9de7311c6e17e7f)
- [Golang で Static Library を作る際、stringをparameterで受け取るならコピーしよう。 - Qiita](http://qiita.com/flowtumn/items/2df066ca776023bcc687)
- [マルチプラットフォーム対応したライブラリ Golang - Qiita](http://qiita.com/nothingcosmos/items/b5dc76aa953222bbdb5c)

### 組み込み開発

- [goでwindowsでRS232C - Qiita](http://qiita.com/ohisama@github/items/e6fa2bd1527c257bb9c9)
    - [ohisama/serial](https://github.com/ohisama/serial) (forkd from [lnmx/serial](https://github.com/lnmx/serial))
    - [goでwindowsでRS232C その２ - Qiita](http://qiita.com/ohisama@github/items/12dccdcdfc5082c22e72)
- [Goでのシリアル通信でハマった事 - Qiita](http://qiita.com/tomoya0x00/items/d957dc00682c57f96771)
- [go 1.5でgomobile(android) - unokun’s blog](http://unokun.hatenablog.jp/entry/2015/08/01/150628)
- [gomobileでiOSアプリをビルドする手順まとめ - GolangRdyJp](http://golang.rdy.jp/2015/09/21/ios-gomobile/)
- [gomobileで日本語フォントを扱ってみる - Qiita](http://qiita.com/bowz_standard/items/5a9c987f9242777fff30)

### ビルド時に情報を各種埋め込みたい

- [Go言語: ビルド時にバージョン情報を埋め込みたい - Qiita](http://qiita.com/suin/items/d643a0ccb6270e8e3734)
- [Golangビルド時に、サブパッケージ内の変数をいじる - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/09/08/211626)
- [Goでビルドバージョン情報を参照できるようにする(Go1.5) - Qiita](http://qiita.com/reiki4040/items/6b32370532c3eafe1f0e)
- [go-bindata でコンパイル時にリソースを埋め込んじゃおう！ - Qiita](http://qiita.com/ikawaha/items/c02d84cfd00f8f442500)
- [ソースを実行ファイルに埋め込む方法 - Qiita](http://qiita.com/ymko/items/c2e3c8fe25bce425136d)
- [Golangで静的ファイルをバイナリに含めるライブラリを書いてみた - Qiita](http://qiita.com/konohazuku/items/131b251a5fa29213ac5e)
- [GolangのGin/bindataでシングルバイナリを試してみた(+React) - Qiita](http://qiita.com/wadahiro/items/4173788d54f028936723)

### Build Tool for Golang

[constabulary/gb](https://github.com/constabulary/gb) を使ってプロジェクトベースの環境構築

- [golang - gbを知ったのでgojiを使ったWEBアプリケーションプロジェクトを管理してみた - Qiita](http://qiita.com/shinofara/items/ac0591fef95c2c6e936e)
- [Building Go projects with gb - Supermighty](https://walledcity.com/supermighty/building-go-projects-with-gb)
- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[FiloSottile/gvt](https://github.com/FiloSottile/gvt) というのがあるらしい。

[Masterminds/glide](https://github.com/Masterminds/glide) と Go 1.5 の Vendoring 機能を組み合わせてパッケージ管理できる。

- [glide - パッケージ管理のお困りの方へ - - Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)

パッケージ依存解決ツールというのがあるらしい。

- [Big Sky :: golang オフィシャル謹製のパッケージ依存解決ツール「dep」](http://mattn.kaoriya.net/software/lang/go/20170125023240.htm)
- [dep(Go dependency tool)を自作ライブラリに使ってみた - Qiita](http://qiita.com/kwmt@github/items/0f77e083494ca94d782b)

### direnv で開発環境を切り替える

[direnv - unclutter your .profile](http://direnv.net/)

- [direnv/direnv](https://github.com/direnv/direnv)
- [direnvで解決するGOPATHの3つの問題点 - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/06/18/010449)
- [改めて、direnvを使いましょう！ - HDE BLOG](http://blog.hde.co.jp/entry/2015/02/27/182117)
- [さくら - homeにgolang, direnv とvirtualenvを入れて動かす - Qiita](http://qiita.com/aminamid/items/5a0e9461385c80d0c8a6)

### Integrated Development Environment (IDE)

- [EclipseでGoプログラミング！ GoClipseのインストールとGojiフレームワークを使ったWeb APIの作成 （1/6）：CodeZine](http://codezine.jp/article/detail/8374)
- [WindowsでGolang開発環境構築　IntelliJ IDEA - Qiita](http://qiita.com/ngsm3/items/67620fc4e39219235a23)
    - [第1回　Android Studio，そしてベースとなる「IntelliJ IDEA」とは何か？：Android Studio最速入門～効率的にコーディングするための使い方｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/serial/01/android_studio/0001) : IntelliJ IDEA についての解説
- [Go入門@環境構築編(IntelliJ IDEA 14.1.4) - Qiita](http://qiita.com/makoto2468/items/3ced77de947ea997e72f)
- [WindowsにIntelliJ IDEAでgoの開発環境を作りましたよ　ついでにgit連携もいたしましたよ - Qiita](http://qiita.com/peka2/items/cdc9d508dcbf4a131271)
- [Cloud9にGo言語でHelloWorldを実装・実行してみた - Qiita](http://qiita.com/ryo-endo/items/e8cb4987b4af9ddc24d5)

#### ATOM で Go

- [AtomでGoを書く環境を整える（Windows） - Technically, technophobic.](http://notchained.hatenablog.com/entry/2014/09/20/104829)
- [AtomでのGo言語開発環境セットアップ - Qiita](http://qiita.com/MakoTano/items/3d807a96c3933ac8aa13)

#### Visual Studio Code で Go

- [VisualStudioCode - Visual Studio CodeでGo言語の設定 - Qiita](http://qiita.com/evalphobia/items/f68396d573c7caf2065b)
- [VSCode+Golang](https://gist.github.com/llaughlin/16305f0b91356e93a1c4) : task.json の設定
- [VisualStudioCodeでGAE/Goの環境設定 - Qiita](http://qiita.com/m0a/items/ebcf964effbc7d50a489)

#### Emacs で Go

- [emacsでGoの環境構築をやる - Qiita](http://qiita.com/koki_cheese/items/2e2ead918a1f1ac5bf6e)

#### Vim で Go

- [Vim で golang を書く環境を整える - Qiita](http://qiita.com/masa23/items/db184871c78311566434)

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
- [なぜGo言語は設計が悪いのか – Go愛好者の見地から | 未分類 | POSTD](http://postd.cc/why-go-is-a-poorly-designed-language/) : 必見！ Go のイケてないところ
- [Go言語の初心者が見ると幸せになれる場所 - Qiita](http://qiita.com/tenntenn/items/0e33a4959250d1a55045)
- [Big Sky :: Names](http://mattn.kaoriya.net/software/20160126101358.htm) : Golang の開発者 Russ Cox 氏による記事の抄訳。「変数名の長さ」について
- [Go言語で作った実行ファイルを小さくしよう！ - Qiita](http://qiita.com/circus/items/450254c59d194cbf22d7)
- [Goのプログラミングパターン](http://www.infoq.com/jp/news/2016/03/go-patterns)
- [Go言語で利用するLLVM入門 | プログラミング | POSTD](http://postd.cc/an-introduction-to-llvm-in-go/)
- [Big Sky :: golang の Windows 版が buildmode=c-archive をサポートした。](http://mattn.kaoriya.net/software/lang/go/20160405114638.htm)
- [Gopherの道を歩む – Node.jsからGo言語への移行 | プログラミング | POSTD](http://postd.cc/the-way-of-the-gopher/)
- [6年間におけるGoのベストプラクティス | プログラミング | POSTD](http://postd.cc/go-best-practices-2016/)
- [Big Sky :: 「みんなのGo言語」の執筆に参加させて頂きました。](http://mattn.kaoriya.net/software/lang/go/20160808013725.htm)
- [Go言語の低レイテンシGC実現のための取り組み | プログラミング | POSTD](http://postd.cc/gos-march-to-low-latency-gc/)
- [Go言語がダメな理由 | プログラミング | POSTD](http://postd.cc/why-go-is-not-good/)
- [ASCII.jp：Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/235/1235262/)
- [グーグルの「Go」、2016年のプログラミング言語大賞に輝く - ZDNet Japan](http://japan.zdnet.com/article/35094856/)
- [Go言語でWebアプリを作りかけて辞めた話 - ぼっち勉強会](http://kannokanno.hatenablog.com/entry/2017/02/28/011159)
- [Go 2016 Survey Results - The Go Blog](https://blog.golang.org/survey2016-results)
- [Big Sky :: golang と Generics と私](http://mattn.kaoriya.net/software/lang/go/20170309201506.htm)
    - [golang と Generics と吾 - Qiita](http://qiita.com/yuroyoro/items/6bf33f3cd4bb35469e0b)
    - [Java の Generics にもの思い - Qiita](http://qiita.com/t2y/items/139c6a38173d7750ddfc)
- [Why Everyone Hates Go · npf.io](Why Everyone Hates Go · npf.io)
    - [[翻訳]なんでGoってみんなに嫌われてるの？ - Qiita](http://qiita.com/hirokidaichi/items/adccebb41f77eaa6132f)
- [Why Go? | Dave Cheney](https://dave.cheney.net/2017/03/20/why-go)
    - [[翻訳] Why Go? - Qiita](http://qiita.com/methane/items/b627f20457873a504638)

### Go 1.5 に関する話題

- [GVM で go1.5rc1 のインストール - Qiita](http://qiita.com/msaito3/items/3aef86e9864990b16b4c)
- [goを1.5にアップデートして1.4とベンチを取る - Qiita](http://qiita.com/masahikoofjoyto/items/4ced298989e6ab346f15)
- [Go 1.3 から 1.5 へのアップデートでエラー - Qiita](http://qiita.com/taji-taji/items/4c43e126e67d65a219e3) : 古いバージョンからアップデートする際は，いったん 1.4 に上げてから 1.5 にアップデートするとよい
- [Big Sky :: golang 1.5 の internal パッケージの使い方。](http://mattn.kaoriya.net/software/lang/go/20150820102400.htm)
    - [「golang 1.5 の internal パッケージの使い方。」を試してみた - Qiita](http://qiita.com/qt-luigi/items/d0f52b3b0906b35e6027)

### Go 1.6 に関する話題

- [Goで良い感じに日時をパースするライブラリdatemakiの話とGo 1.6 - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/12/22/000011)
- [Go1.6でポインタをcgoの関数へ渡す際の注意点 - Qiita](http://qiita.com/saturday06/items/84535c61a3328c02032c)
- [Go1.6でポインタをcgoの関数へ渡す際に発生するcgoCheckPointerを回避する方法 - Qiita](http://qiita.com/mattn/items/90c8558d5fff05a2ba0c)
- [Goのバージョンを1.6rc2にアップデートしてみた - Qiita](http://qiita.com/kanuma1984/items/245f7efafeaee5728523)
- [Goのバージョンを1.4.3→1.6にアップグレードできなかった - Qiita](http://qiita.com/hirocueki2/items/3ec4b409a3ed2cbea681)
- [Go 1.6 templateパッケージ新機能 - Qiita](http://qiita.com/hoshi-k/items/f2eaff298f93f089e10d)

### Go 1.7 に関する話題

- [go1.7 の気になるところを試した - Qiita](http://qiita.com/tutuming/items/16df6fc7bf82fb5d7eb0)
- [Sierraでgo1.7.3のコンパイル - Qiita](http://qiita.com/lufia/items/8a71a29fa6e1089735f2)
- [【作業メモ】さくらインターネットの標準CGIサーバでGoをビルド【古典的レンタル鯖】 - Qiita](http://qiita.com/zetamatta/items/143849cf34db4ffaad4b)
- [Go1.7からSSAが導入された - flyhigh](http://shinpei.github.io/blog/2016/08/13/what-ssa-brings-to-go-17/)

### Go 1.8 に関する話題

- [Go 1.8のpluginパッケージを試してみる - Qiita](http://qiita.com/qt-luigi/items/47a7913145fc747da0c7)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/61EL3Dc95dL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/477418392X/baldandersinf-22/">みんなのGo言語【現場で使える実践テクニック】</a></dt><dd>松木雅幸 mattn 藤原俊一郎 中島大一 牧 大輔 鈴木健太 稲葉貴洋 </dd><dd>技術評論社 2016-09-09</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774184322/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774184322.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.95"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274219151/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274219151.09._SCTHUMBZZZ_.jpg"  alt="プログラミングElixir"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798147400/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798147400.09._SCTHUMBZZZ_.jpg"  alt="詳解MySQL 5.7 止まらぬ進化に乗り遅れないためのテクニカルガイド (NEXT ONE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774182869/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774182869.09._SCTHUMBZZZ_.jpg"  alt="WEB+DB PRESS Vol.94"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117763/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117763.09._SCTHUMBZZZ_.jpg"  alt="Docker"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/477418361X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/477418361X.09._SCTHUMBZZZ_.jpg"  alt="オブジェクト指向設計実践ガイド ~Rubyでわかる 進化しつづける柔軟なアプリケーションの育て方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4295000337/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4295000337.09._SCTHUMBZZZ_.jpg"  alt="WebデベロッパーのためのReact開発入門 JavaScript UIライブラリの基本と活用"  /></a> </p>
<p class="description">リファレンス本なのに索引が貧弱。これなら Kindle で買ってもよかったか。 1.7 への言及があるのはよかった。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-17">2016-11-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/">Go言語によるWebアプリケーション開発</a></dt><dd>Mat Ryer 鵜飼 文敏 </dd><dd>オライリージャパン 2016-01-22</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774178667/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774178667.09._SCTHUMBZZZ_.jpg"  alt="nginx実践入門 (WEB+DB PRESS plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4863541783/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4863541783.09._SCTHUMBZZZ_.jpg"  alt="改訂2版 基礎からわかる Go言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774179930/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774179930.09._SCTHUMBZZZ_.jpg"  alt="サーバ/インフラエンジニア養成読本 DevOps編 [Infrastructure as Code を実践するノウハウが満載! ] (Software Design plus)"  /></a> </p>
<p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-13">2016-03-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
