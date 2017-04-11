+++
date = "2015-09-11T23:59:59+09:00"
update = "2017-04-11T10:04:59+09:00"
description = "本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。"
draft = false
tags = ["golang", "bookmark", "package"]
title = "Go 言語に関するブックマーク（サンプルコード，パッケージ関連）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  license = "by-sa"
  tumblr = "spiegel-im-spiegel"
  linkedin = "spiegelimspiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
  name = "Spiegel"
  flattr = "spiegel"
  instagram = "spiegel_2007"
+++

本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。
なお以下に挙げているブックマーク以外のページは「[Go 言語に関するブックマーク]({{< relref "golang/bookmark.md" >}})」にあるかもしれない。
よろしければ参考にどうぞ。

## context パッケージ

[バージョン 1,7 から標準パッケージ](https://tip.golang.org/doc/go1.7#context)に入った [context](https://golang.org/pkg/context/) パッケージについて。

- [Goの並行パターン：コンテキスト (Go Concurrency Pattern: Context)](https://ymotongpoo.github.io/goblog-ja/post/context/)
- [Go1.7のcontextパッケージ | SOTA](http://deeeet.com/writing/2016/07/22/context/)
- [Golangのcontext.Valueの使い方 | SOTA](http://deeeet.com/writing/2017/02/23/go-context-value/)
- [contextの使い方 - Qiita](http://qiita.com/taizo/items/69d3de8622eabe8da6a2)
- [context.Contextでリクエストスコープな値を持ち回す - Qiita](http://qiita.com/hogedigo/items/a26d816395b7545ce5f8) : [context](https://golang.org/pkg/context/) の使い方って（名前からいって）本来こっちだよね
- [goroutine にシグナルを送信する - Qiita](http://qiita.com/Kei-Kamikawa/items/620f9504daf2ec53f0b5)

## Logging

- [Golang logging library - Qiita](http://qiita.com/kosuda/items/988c505c2abc5321aba8)
- [go言語におけるロギングについて](http://blog.satotaichi.info/logging-frameworks-for-go/)
- [Golangで簡単にログを吐くことを考える - Qiita](http://qiita.com/Ets/items/49e8f781990a3b0b3821) : [seelog](https://github.com/cihub/seelog) について解説している。私は XML には全くアレルギーがないので無問題
- [Go言語でdebugログの実現方法 - Qiita](http://qiita.com/sbjib/items/2cef51e572eef0795bc2)
- [loggingについて話そう - Qiita](http://qiita.com/methane/items/cedbf546ae2db8a63c3d)
- [golangでlogを標準出力とテキストファイルの2箇所の出力する - Qiita](http://qiita.com/74th/items/441ffcab80a6a28f7ee3)
- [gorpで実行されるSQLをログに出力する - Qiita](http://qiita.com/reiki4040/items/303a0bfa9f0296eb544f)
- [Goのバッチで統計を取得するAPIを用意しておくと便利 - Qiita](http://qiita.com/sudix/items/c542e1b59bc94dc741e3)
- [golangのloggerを作ってみた - Qiita](http://qiita.com/kazuma1107/items/009454fca4f56af6e411)
- [golangの高速な構造化ログライブラリ「zap」の使い方 - Qiita](http://qiita.com/emonuh/items/28dbee9bf2fe51d28153)
- [zapでログレベルでログの出力先を振り分ける方法 - Qiita](http://qiita.com/emonuh/items/cb3a730979dba7d76920)

## コマンドライン・ツール

- [Go言語によるCLIツール開発とUNIX哲学について - ゆううきブログ](http://yuuki.hatenablog.com/entry/go-cli-unix)
- [開発者から見た UNIX 哲学とコマンドラインツールと Go言語 - TELLME.TOKYO](http://tellme.tokyo/post/2015/06/23/unix_cli_tool_go/) （[Qiita 版](http://qiita.com/b4b4r07/items/df660d82e2de715acda5)）
- [コマンドライン引数 - Qiita](http://qiita.com/uokada/items/f0e069a751679dcf616d)
- [Go言語のflagパッケージを使う - uragami note](http://ryochack.hatenablog.com/entry/2013/04/17/232753)
- [Go を使ってコマンドラインツール wordc を作ってみた - Qiita](http://qiita.com/flaflasun/items/df5ebb057697da062a08) : [codegangsta/cli](https://github.com/codegangsta/cli) についての言及あり
- [GoでCLIツール作るのに便利そうなパッケージを集めてみました - Qiita](http://qiita.com/isaoshimizu/items/71dd2ca2a08ddb607e31)
- [C言語とGo言語で標準出力が端末を参照しているかどうかを判定する - uragami note](http://ryochack.hatenablog.com/entry/2013/07/15/232207)
- [Go言語のCLIツールのpanicをラップしてクラッシュレポートをつくる | SOTA](http://deeeet.com/writing/2015/04/17/panicwrap/)
- [flag 並にシンプルでより強力な CLI パーサ kingpin の紹介 - Qiita](http://qiita.com/kumatch/items/258d7984c0270f6dd73a)
- [Goで外部コマンドをパイプして実行する - Qiita](http://qiita.com/yuroyoro/items/9358cd25b5f7fe9dd37f)
    - [Big Sky :: golang で UNIX コマンドパイプラインを扱う](http://mattn.kaoriya.net/software/lang/go/20151030131242.htm)
- [Golangで外部コマンドを実行する方法まとめ - Qiita](http://qiita.com/tanksuzuki/items/9205ff70c57c4c03b5e5)
- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
    - [GolangでCLIの場合にcobraを使うことにした件 | FLAMA技術Blog](http://lab.flama.co.jp/archives/1536/)
- [GolangでwebサービスのAPIを叩くCLIツールを作ろう - Qiita](http://qiita.com/minamijoyo/items/cfd22e9e6d3581c5d81f)
- [GoでMySQLにCA証明書を使ってアクセスする - Qiita](http://qiita.com/Peranikov/items/4376633bb72492051336)

## 入出力処理

- [Go ファイルや標準入力から一行ずつ読み込む - Qiita](http://qiita.com/hnakamur/items/a53b701c8827fe4bfec7)
- [Goでファイルを読んで別のgoroutineに渡す - Qiita](http://qiita.com/ono_matope/items/5c8cfce81933c5eb9fd0)
- [「連結されたgzipを1行ずつ見る」をgolangでやったらハマった - Qiita](http://qiita.com/kroton/items/431e6dad9e5e4dbc44cf) : [compress/gzip](https://golang.org/pkg/compress/gzip/) と入出力処理の関係
- [bufio.Scannerのend-of-line判断を変更してみる - Qiita](http://qiita.com/curious-eyes/items/2d4b6c20ea47e3efc47b)
- [KOBE GDG: Go言語　バイナリファイルを扱う](http://kobegdg.blogspot.jp/2013/05/go.html) : 任意のオブジェクトをバイト列に変換してファイルに格納
- [ファイル書き込みの度にファイルを開いたらどれくらい遅いのか - Qiita](http://qiita.com/catatsuy/items/41d3c49248b517b5af96)
- [Golangで標準入力がパイプで渡されたものか判定する - Qiita](http://qiita.com/tanksuzuki/items/e712717675faf4efb07a)
- [Go言語: ファイルの存在をちゃんとチェックする実装? - Qiita](http://qiita.com/suin/items/b9c0f92851454dc6d461)
- [Golangで、ファイル一覧取得（最新順出力） - Qiita](http://qiita.com/shinofara/items/e5e78e6864a60dc851a6)
- [大きなファイルのアップロードを省メモリで行いたい(io.Pipeを使う) - Qiita](http://qiita.com/m0a/items/bba395b2fc9cd160e441)

## ソートアルゴリズム

- [sliceのシャッフル - Qiita](http://qiita.com/sugyan/items/fd7138a756c1a409f5fd) : [Fisher–Yates shuffle](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) というアルゴリズムらしい
- [Go言語でバイトニックソート実装してみた - Qiita](http://qiita.com/nyamadandan/items/2c82011801b148c98e52)
- [Goでバケットソートアルゴリズム(ビット列を使用) - Qiita](http://qiita.com/ohkawa/items/269507985b3ae10cbff9)
- [interface{} をソートする - Qiita](http://qiita.com/tchssk/items/b61f1f06d22a6232d4c8)
- [Big Sky :: golang の sort インタフェース難しい問題が解決した](http://mattn.kaoriya.net/software/lang/go/20161004092237.htm)

## Web Microframework for Golang

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
- [[golang]RevelでCSRFの対策 - Qiita](http://qiita.com/ponchi/items/7e474041484841ee1ce6)
- [net/httpでポート443のHTTPSサーバーを立ち上げるまで - Qiita](http://qiita.com/ryurock/items/f55db5944397619735bf)
- [概観からGoのWebFrameworkを選ぶ（2016/02） - Qiita](http://qiita.com/jumbOS5/items/45f86db15a5a6c8a0622)
- [Revel(Golang)でViewを理解する - Qiita](http://qiita.com/jumbOS5/items/d817bc95279877e72b24)
- [Big Sky :: golang で画像アップロードが簡単に出来る go-imageupload を使ってみた。](http://mattn.kaoriya.net/software/lang/go/20160224103638.htm)
- [Go言語のスクレイピング系ライブラリまとめ - Qiita](http://qiita.com/okataitai/items/db6999ea1ab39ec0bd3e)
- [Go の echo ってWebサーバーでサクッと REST しちゃう - Qiita](http://qiita.com/keika299/items/62e806ae42828bb3567a)
- [nginx+circus+gojiによるgolang webアプリケーションの動作環境構築 - Qiita](http://qiita.com/reiki4040/items/795a008d1b12ee657d9a)
- [Big Sky :: golang で HTTP を使った処理を4倍速くする](http://mattn.kaoriya.net/software/lang/go/20160329094503.htm)
- [GolangのフレームワークEchoの話とHelloWorldサンプルを読み解くだけ - Qiita](http://qiita.com/CST_negi/items/5ea3395b35e68fd0d3b3)
- [Golang+Echo+dbrでMySQLのCRUDをする／JSONでDBの値を返却する話 - Qiita](http://qiita.com/CST_negi/items/5e276ddc0412cefef7e3)
- [Gin(Golang)におけるHTMLテンプレート記述方法 - Qiita](http://qiita.com/lanevok/items/dbf591a3916070fcba0d)
- [GoLang で html/template を使っていて遭遇したエラーとその解決 - Qiita](http://qiita.com/kent_ocean/items/45e153d2c5467501a20b)
- [GAE/Go (echoフレームワーク)で Line Message API 使って Bot を作る。 - Qiita](http://qiita.com/naoki_koreeda/items/8c818a3e9f6138ddbb87)
- [Go製のフレームワークechoを使ってJSONを返すWebサーバーを作り、GoogleAppEngineで動かす - Qiita](http://qiita.com/qube81/items/701279c43b33ce923613)
- [go+ginでローカルで作ったWebアプリをGoogleAppEngineに載せる時に注意することまとめ - Qiita](http://qiita.com/CST_negi/items/bcb4730c6efeb838c4a7)
- [社内のバックエンド開発にgRPCを導入してみた - Qiita](http://qiita.com/nozaq/items/9cd9bf7ee6118779bda9)
- [Echoはver.3で結局のところ何が変わったのか？ - Qiita](http://qiita.com/tienlen/items/85b06c0856c33e716c75)
- [Go言語のWEBフレームワークRevelを使用してセキュアなAPIを作成 - Qiita](http://qiita.com/GushiSnow/items/97c1f64c003b27c6b98a)
- [Go でツール書くときの Makefile 晒す - Qiita](http://qiita.com/dtan4/items/8c417b629b6b2033a541)
- [goa でデザイン・ファーストをシュッとする - Qiita](http://qiita.com/ikawaha/items/6638ee8b6978aef50d65)
- [echo 初心者でも簡単!! echo で扱うアセットファイル群を簡単にバイナリにまとめて使ってみる - Qiita](http://qiita.com/Kei-Kamikawa/items/a6cb72251b95c8f5baa1)
- [Go ライブラリによる CGIプログラム内ルーチング - Qiita](http://qiita.com/jyagaimo_qiita_/items/93d195ca65982b75e205)
- [Go 言語で Apache Bench (ab) を実装してみた - takatoshiono's blog](http://takatoshiono.hatenablog.com/entry/2017/02/06/013323)
- [HTTP/2 Server Push - The Go Blog](https://blog.golang.org/h2push)
- [Go + echoでfetch APIを使ってPOSTする - Qiita](http://qiita.com/yutasuzuki/items/fa7e78a8f4d8cb3e83b3)

## 他サービスとの連携

- [マイナンバーのチェックデジットをGoで計算する - Qiita](http://qiita.com/qube81/items/f66a38b28ec58bc5c4da)
- [golangでImageMagickを触りたい - Qiita](http://qiita.com/arc279/items/5f277aa5cce3de5247e5)
- [Go言語でRedshiftとつなぐ（というかただのPostgreSQL） - Qiita](http://qiita.com/otiai10/items/83b186596897705ce392)
- [Gmail API for Goで、下書き生成ツールを作りました。 - Qiita](http://qiita.com/yyoshiki41/items/1159e1a70ffaa8fd84ed)
- [Big Sky :: Windows からも ssh でリモートコマンド実行したい、それ golang で出来るよ](http://mattn.kaoriya.net/software/lang/go/20170111165324.htm)
- [コマンドラインからググれてもいいと思ったので作った - Qiita](http://qiita.com/ieee0824/items/13435fc6de5f22cdb2f4)
- [GoでSpreadsheetを操作するパッケージを作った - Qiita](http://qiita.com/Iwark/items/726dfb2d15a883e389b6)
- [Go言語で東京メトロAPIを叩く - Qiita](http://qiita.com/nayuneko/items/ca2651e3a613c8e3256a)
- [Twilio使って、入力した電話番号をチェックする方法 - Qiita](http://qiita.com/dbsparkle/items/72fc1ad50ba6000f2630)
- [Windows+GoでNFC/Felicaにアクセスしてみた - Qiita](http://qiita.com/mau4x/items/424fe7964e70a3a99965)
- [UnityのネイティブプラグインをGoで書く #golang #unity - Qiita](http://qiita.com/tenntenn/items/4d3316490a571e5d79ed)
- [radikoの録音ツールをGoで書いた - Qiita](http://qiita.com/yyoshiki41/items/f81442d7dc2d0ddcf15b)

### DB 連携

- [Go言語でBigQueryのクエリを実行してみる - Qiita](http://qiita.com/najeira/items/8310fecf4b70c918f855)
- [Go の DB アクセス用のパッケージを作った - Qiita](http://qiita.com/chimatter/items/1a5fb2f03477f2ada520)
- [MySQLでのトランザクション処理をGolang+dbrで実現してゆく話 - Qiita](http://qiita.com/CST_negi/items/d79600c34191adb09c79)

### GAE/Go

- [GAE/GoでCMSつくった - Qiita](http://qiita.com/hogedigo/items/342217982f267ccd234d)
- [Go+GAE+Cloud Datastoreで簡単なREST APIを構築 - Qiita](http://qiita.com/silverfox/items/81769425e51f24e676d2])
- [Google App Engine SDK for Goを使ってGAE上でアプリを動かすまで - Qiita](http://qiita.com/walkers/items/e407386d7ef184ec830a)
- [GAE/Goで形態素解析してみた - Qiita](http://qiita.com/mako0715/items/259659e5e2935d2afc10)
- [GAE/Go+glide的な構成での環境構築 ~ローカルサーバー立ち上げまで~ - Qiita](http://qiita.com/ryutah/items/d864310c62f0385d876d)
- [Go言語の依存管理ツールを作って、開発環境構築を覚えた - Qiita](http://qiita.com/eaglesakura/items/b7e92281735569c528a6)
- [GAE/Go で Google Cloud Spanner を操作する（前編） - Qiita](http://qiita.com/wezardnet/items/daf520b82e2199d16f4f)

## Go で数学

- [大学入試問題をGoで解いてみる - Qiita](http://qiita.com/qube81/items/c47b9e3ea8d028e95588) : [math/big](https://golang.org/pkg/math/big/) パッケージを使って大きな数を計算する。
- [golangで数独を解いた - Qiita](http://qiita.com/ciruzzo/items/144bc1874947441f9fb8)
- [golang で AB x CD / E - F * G * H = 2016 になる全パターン洗い出し - Qiita](http://qiita.com/amanoiverse/items/06fff7b224d77517c08f)
- [Go で 0 から始まる連続する n 個の整数を重複無く k 個選んだ時の組み合わせの列挙 - Qiita](http://qiita.com/yumura_s/items/68760d6b902aee9c78f0)
- [ピーマンとハトと数学を Go 言語で試す - Qiita](http://qiita.com/nirasan/items/69643d0ddf8a7345cf7c)
- [golangでニュートン法を使って平方根の計算をする - Qiita](http://qiita.com/k-yamada@github/items/0a7baa61bd668c3cb3dc)

### Go 言語で学ぶ『暗号技術入門』

[Soshi Katsuta](https://github.com/skatsuta) さんによるテキスト。
丁寧な内容でとても参考になる。

- [Go 言語で学ぶ『暗号技術入門』Part 1 -DES, Triple DES- | Step by Step](http://skatsuta.github.io/2016/01/02/hyuki-crypt-book-go-1/)
- [Go 言語で学ぶ『暗号技術入門』Part 2 -AES- | Step by Step](http://skatsuta.github.io/2016/01/19/hyuki-crypt-book-go-2/)
- [Go 言語で学ぶ『暗号技術入門』Part 3 -CBC Mode- | Step by Step](http://skatsuta.github.io/2016/03/06/hyuki-crypt-book-go-3/)

その他，暗号関連

- [Go言語と暗号技術（AESからTLS） | SOTA](http://deeeet.com/writing/2015/11/10/go-crypto/)
    - [tcnksm/go-crypto](https://github.com/tcnksm/go-crypto) : サンプルコード

## その他のサンプルコードおよびパッケージ

- [Go用のGoogle製のGUIツールキットgxuiのインストール(Windows版) - Qiita](http://qiita.com/sago35/items/cc9ed3dc38d0b2f19bf9)
- [goでLチカの練習　その２ - Qiita](http://qiita.com/ohisama@github/items/bfc1eb6407cbdfebbd18)
- [Revel templatesを使ったサンプルアプリケーション - Qiita](http://qiita.com/rubytomato@github/items/638299aabb7922cbef59)
- [goでwindowsでwindow - Qiita](http://qiita.com/ohisama@github/items/20bc61175ce4a33b7365)
- [Go言語でパッケージを作成して世界に公開する方法 at ミネルヴァの梟は黄昏とともに飛び始める（山下 大介 公式ブログ）](http://blog.daisukeyamashita.com/post/1209.html) : パッケージの作り方なんだけど情報が古い。「昔はこうだった」くらいの感じで
- [goでwindowsでキースキャン - Qiita](http://qiita.com/ohisama@github/items/9f05679f25cfc9c3ecfc)
- [gocron でジョブスケジューリング - Qiita](http://qiita.com/mitsuse/items/8669bf54d2310b3e68a1) : [gcarlescere/scheduler](https://github.com/carlescere/scheduler) のほうがおすすめらしいw
- [Google ChromeのテキストエリアをEmacsで編集する'Edit with Emacs'から任意のエディタを起動するデーモンをGo 1.4 for Windowsで書いてみたわけだが、エディタがブラウザの後ろに出てしまってダメかもしれない - Qiita](http://qiita.com/zetamatta/items/51b0f45496e5143e2e63)
- [golangでprivateなエイリアスのポインタを元の型に戻す - Qiita](http://qiita.com/shibukawa/items/9db22c9684cc0586b737)
- [Go の expvar パッケージを使ってアプリケーションのメトリクスを公開する - Qiita](http://qiita.com/methane/items/8f56f663d6da4dee9f64)
- [Go小ネタ: 正規表現を全角半角問わずマッチするよう変換する - Qiita](http://qiita.com/tutuming/items/d8aaaf5442d84a7961e1)
- [Mithril＋golang Gin を試す - Qiita](http://qiita.com/masatsugumatsus/items/e28254ff52963705ce7f)
- [cmd.Envを設定してexecしたらコケた - Qiita](http://qiita.com/8845musign/items/5c4b32f82c2df08acd93) : [`exec`](https://golang.org/pkg/os/exec/) パッケージを使う際には環境変数に注意
- [今後イケそうなデスクトップGUIフレームワーク - Qiita](http://qiita.com/hachi8833/items/463264f531474a856064) : [go-thrust](https://github.com/miketheprogrammer/go-thrust/) について言及
- [1l0/sumeragi](https://github.com/1l0/sumeragi) : 皇紀や元号を出力するパッケージ
- [RubyからGoの関数をつかう → はやい - Qiita](http://qiita.com/grj_achm/items/679b3f3af2cf377f0f02)
    - [Perl6からGoの関数をつかう → はやい? - Qiita](http://qiita.com/B73W56H84/items/20a67b74bb646d140f7d)
    - [GroovyからGoの関数を使う→はやい - Qiita](http://qiita.com/yujiorama_at_github/items/3f7cab906969764cc805)
    - [Big Sky :: RubyからGoの関数をつかわなくても再帰をやめる → はやい](http://mattn.kaoriya.net/software/20151106194958.htm)
- [GolangのOpenGL事情(WebGLも含むよ) - Qiita](http://qiita.com/shibukawa/items/58f6a421462b93dec471)
- [goのmgoでfindするときの処理 - Qiita](http://qiita.com/enokidoK/items/a3aff4c05e494b004ef8)
- [GoでShared Libraryをビルドしてみた(簡単ドキュメント指向DB) - Qiita](http://qiita.com/umisama/items/48ba74a58f1e6530e305)
- [Goで3Dモデル変換してプレビュー - Qiita](http://qiita.com/tetuyoko/items/746599e36ca4985d9e1a)
- [Go最後の秘宝「GUI」を探しに行く - Qiita](http://qiita.com/shibukawa/items/cd8d122dfeb41e1608d1)
- [Go + QML + QChart.js で素敵なチャートを表示する - Qiita](http://qiita.com/miyabishi/items/09a55b10953c9dbe7ee3)
- [Goで帳票をPDFに作成するライブラリ。 請求書などの複雑なフォーマットにも対応 - Qiita](http://qiita.com/mikeshimura/items/d149bbd869683e820579)
- [Go で簡単に Excelを作成するライブラリ。 色、罫線、網掛けを事前定義済 - Qiita](http://qiita.com/mikeshimura/items/b60823e923fb6d0840c0)
- [Goで関数型プログラミング - Qiita](http://qiita.com/taksatou@github/items/d721a62158f554b8e399) : [reflect](https://golang.org/pkg/reflect/ "reflect - The Go Programming Language") パッケージを使って高階関数を表現できる
- [GoでANSIエスケープコードを扱うライブラリを作った(色付け・カーソル移動等) - Qiita](http://qiita.com/morikuni/items/ad8d900f56ddeb223101)
- [Go で interface {} の中身がポインタならその参照先を取得する - Qiita](http://qiita.com/chimatter/items/b0879401d6666589ab71)
- [Golangでreduce関数を提供しているライブラリugoを眺めてみた - Qiita](http://qiita.com/letusfly85/items/5f479e5b072a05dbcf53)
- [Go 言語で rm 用ごみ箱ツール gomi を作った - Qiita](http://qiita.com/b4b4r07/items/3a790fe7e925b4ba14f3)
- [パッケージのimport pathを好みのURLにする - Qiita](http://qiita.com/lufia/items/8f3cc32f26168702e2f4)
- [JSONSchemaからstructのようなコードを生成する"structr"というのを書いた - Qiita](http://qiita.com/damele0n/items/92a9b845c991b1b29aea)
- [Golang で過去の遺物的(cp932)DLLを利用する - Qiita](http://qiita.com/asm/items/6184cf5dcca637670e0e)
- [Terraform for さくらのクラウド - Qiita](http://qiita.com/yamamoto-febc/items/0ce30e2dba32c60bbf66)
- [golangのtime.Timeの当日00:00:00を取得する方法とベンチマーク - Qiita](http://qiita.com/ushio_s/items/3e270933641710bbd88e)
- [Big Sky :: golang の遅いコードをたった1行で高速化するテクニック](http://mattn.kaoriya.net/software/lang/go/20160804131744.htm) : 実際にはちょっと速くなるくらいらしい
- [Big Sky :: Golang で Windows の DLL を作る方法](http://mattn.kaoriya.net/software/lang/go/20160921010820.htm)
- [Big Sky :: golang で slim テンプレートエンジン書いてる。](http://mattn.kaoriya.net/software/lang/go/20160910001214.htm)
- [独自のfmt.Formatterを実装する - Qiita](http://qiita.com/deeeet/items/1e2164f89ccfc29d7b11)
- [Go 言語で wc を実装してみた - takatoshiono's blog](http://takatoshiono.hatenablog.com/entry/2016/09/22/024605)
- [【コンピュータ将棋】ゴルーチンでお手軽持ち時間管理＆並行探索 - Qiita](http://qiita.com/32hiko/items/3be36dad2d651399ba1b)
- [Go言語でエクセルファイル (.xlsx) を読み込む - Qiita](http://qiita.com/kaorumori/items/fa37130065d0450d6342) : [`github.com/tealeg/xlsx`](https://github.com/tealeg/xlsx) パッケージを使用
- [Go言語でExcel操作ライブラリを書いてみた - Qiita](http://qiita.com/tebakane/items/2f2ed2558357c274c478) : [`github.com/loadoff/excl`](https://github.com/loadoff/excl) パッケージの説明
- [簡単な式の評価機を作ってみる #golang - Qiita](http://qiita.com/tenntenn/items/a312d2c5381e36cf4cd3)
- [Goのコード生成のためのテンプレートエンジン seyfert を書いてみた - Qiita](http://qiita.com/mackee_w/items/71d7685852bb5bdda465)
- [nagomeのplugin ngm-polly 作った - Qiita](http://qiita.com/bamchoh/items/49e230db51e7237b1ce9)
- [Go言語でコレクション処理のメソッドを作ってみた　#golang - Qiita](http://qiita.com/yagitatsu/items/264aa3e167bf4650e705)
- [ASTを取得する方法を調べる #golang - Qiita](http://qiita.com/tenntenn/items/13340f2845316532b55a)
- [Big Sky :: golang の http.Client を速くする](http://mattn.kaoriya.net/software/lang/go/20170112181052.htm)
- [goパッケージで簡単に静的解析して世界を広げよう #golang - Qiita](http://qiita.com/tenntenn/items/868704380455c5090d4b)
- [Big Sky :: Re: Go でシングルバイナリな Web アプリを開発しているときに webpack --watch をうまいところやる](http://mattn.kaoriya.net/software/lang/go/20170119180147.htm)
- [Ruby + mecabが遅いのでGoを経由する - Qiita](http://qiita.com/EastResident/items/f41fd0285fe270e7d3d5)
- [こわくない！今日からはじめるGo言語コード生成 - Qiita](http://qiita.com/nirasan/items/bb0a239641028312b4db)
- [Big Sky :: レーベンシュタイン距離を使ったあいまい grep コマンド「lsdgrep」作ってみた](http://mattn.kaoriya.net/software/lang/go/20170227010706.htm)
- [Re:ゼロから始めないAPNGエンコーダ - Qiita](http://qiita.com/cia_rana/items/18c78e0233e117b22af6)
- [Go言語のFunctional Option Pattern - Qiita](http://qiita.com/weloan/items/56f1c7792088b5ede136)
- [Go言語を使用して簡単なLineBotを作る - Qiita](http://qiita.com/sao_rio/items/8801b78ba60acbb0ae41)
- [golangのGUIパッケージgo-gtkを試す - Qiita](http://qiita.com/intelf___/items/2207c02c306a495d99e6)

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
