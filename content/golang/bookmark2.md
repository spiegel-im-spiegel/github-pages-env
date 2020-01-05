+++
title = "Go 言語に関するブックマーク（未整理分）"
date = "2015-09-11T23:59:59+09:00"
description = "「Go 言語に関するブックマーク」に整理しきれない未整理分。"
tags = ["golang", "bookmark"]
image = "/images/attention/go-logo_blue.png"
  
[scripts]
  mathjax = false
  mermaidjs = false
+++

「[Go 言語に関するブックマーク]({{< relref "bookmark.md" >}})」に整理しきれない未整理分だけど，よろしければ参考にどうぞ。

## ニュース関連

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
- [Go言語の初心者が見ると幸せになれる場所 - Qiita](http://qiita.com/tenntenn/items/0e33a4959250d1a55045)
- [Go言語で利用するLLVM入門 | プログラミング | POSTD](http://postd.cc/an-introduction-to-llvm-in-go/)
- [Big Sky :: 「みんなのGo言語」の執筆に参加させて頂きました。](http://mattn.kaoriya.net/software/lang/go/20160808013725.htm)
- [ASCII.jp：Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/235/1235262/)
- [グーグルの「Go」、2016年のプログラミング言語大賞に輝く - ZDNet Japan](http://japan.zdnet.com/article/35094856/)
- [Go言語でWebアプリを作りかけて辞めた話 - ぼっち勉強会](http://kannokanno.hatenablog.com/entry/2017/02/28/011159)
- [Go 2016 Survey Results - The Go Blog](https://blog.golang.org/survey2016-results)
- [Big Sky :: GoCon に初参加してきた。](https://mattn.kaoriya.net/etc/gocon2017autumn.htm)
- [［速報］AWS Lambdaが機能強化。.NETとGo言語をサポート、サーバレスアプリケーションのリポジトリも登場。AWS re:Invent 2017 － Publickey](http://www.publickey1.jp/blog/17/aws_lambdanetgoaws_reinvent_2017.html)
- [サーバレスアーキテクチャのAWS Lambda、Go言語とC#の.NET Core 2.0サポートを開始 － Publickey](http://www.publickey1.jp/blog/18/aws_lambdagoc.html)

## その他，分類困難で未整理（もしくは一時保管場所）

- [goでLチカの練習　その２ - Qiita](http://qiita.com/ohisama@github/items/bfc1eb6407cbdfebbd18)
- [Revel templatesを使ったサンプルアプリケーション - Qiita](http://qiita.com/rubytomato@github/items/638299aabb7922cbef59)
- [Go言語でパッケージを作成して世界に公開する方法 at ミネルヴァの梟は黄昏とともに飛び始める（山下 大介 公式ブログ）](http://blog.daisukeyamashita.com/post/1209.html) : パッケージの作り方なんだけど情報が古い。「昔はこうだった」くらいの感じで
- [goでwindowsでキースキャン - Qiita](http://qiita.com/ohisama@github/items/9f05679f25cfc9c3ecfc)
- [Google ChromeのテキストエリアをEmacsで編集する'Edit with Emacs'から任意のエディタを起動するデーモンをGo 1.4 for Windowsで書いてみたわけだが、エディタがブラウザの後ろに出てしまってダメかもしれない - Qiita](http://qiita.com/zetamatta/items/51b0f45496e5143e2e63)
- [golangでprivateなエイリアスのポインタを元の型に戻す - Qiita](http://qiita.com/shibukawa/items/9db22c9684cc0586b737)
- [Go の expvar パッケージを使ってアプリケーションのメトリクスを公開する - Qiita](http://qiita.com/methane/items/8f56f663d6da4dee9f64)
- [cmd.Envを設定してexecしたらコケた - Qiita](http://qiita.com/8845musign/items/5c4b32f82c2df08acd93) : [`exec`](https://golang.org/pkg/os/exec/) パッケージを使う際には環境変数に注意
- [1l0/sumeragi](https://github.com/1l0/sumeragi) : 皇紀や元号を出力するパッケージ
- [goのmgoでfindするときの処理 - Qiita](http://qiita.com/enokidoK/items/a3aff4c05e494b004ef8)
- [Go + QML + QChart.js で素敵なチャートを表示する - Qiita](http://qiita.com/miyabishi/items/09a55b10953c9dbe7ee3)
- [Goで帳票をPDFに作成するライブラリ。 請求書などの複雑なフォーマットにも対応 - Qiita](http://qiita.com/mikeshimura/items/d149bbd869683e820579)
- [GoでANSIエスケープコードを扱うライブラリを作った(色付け・カーソル移動等) - Qiita](http://qiita.com/morikuni/items/ad8d900f56ddeb223101)
- [Golangでreduce関数を提供しているライブラリugoを眺めてみた - Qiita](http://qiita.com/letusfly85/items/5f479e5b072a05dbcf53)
- [Go 言語で rm 用ごみ箱ツール gomi を作った - Qiita](http://qiita.com/b4b4r07/items/3a790fe7e925b4ba14f3)
- [パッケージのimport pathを好みのURLにする - Qiita](http://qiita.com/lufia/items/8f3cc32f26168702e2f4)
- [Terraform for さくらのクラウド - Qiita](http://qiita.com/yamamoto-febc/items/0ce30e2dba32c60bbf66)
- [Big Sky :: golang で slim テンプレートエンジン書いてる。](http://mattn.kaoriya.net/software/lang/go/20160910001214.htm)
- [独自のfmt.Formatterを実装する - Qiita](http://qiita.com/deeeet/items/1e2164f89ccfc29d7b11)
- [Go 言語で wc を実装してみた - takatoshiono's blog](http://takatoshiono.hatenablog.com/entry/2016/09/22/024605)
- [簡単な式の評価機を作ってみる #golang - Qiita](http://qiita.com/tenntenn/items/a312d2c5381e36cf4cd3)
- [Goのコード生成のためのテンプレートエンジン seyfert を書いてみた - Qiita](http://qiita.com/mackee_w/items/71d7685852bb5bdda465)
- [nagomeのplugin ngm-polly 作った - Qiita](http://qiita.com/bamchoh/items/49e230db51e7237b1ce9)
- [Go言語でコレクション処理のメソッドを作ってみた　#golang - Qiita](http://qiita.com/yagitatsu/items/264aa3e167bf4650e705)
- [ASTを取得する方法を調べる #golang - Qiita](http://qiita.com/tenntenn/items/13340f2845316532b55a)
- [goパッケージで簡単に静的解析して世界を広げよう #golang - Qiita](http://qiita.com/tenntenn/items/868704380455c5090d4b)
- [こわくない！今日からはじめるGo言語コード生成 - Qiita](http://qiita.com/nirasan/items/bb0a239641028312b4db)
- [Big Sky :: レーベンシュタイン距離を使ったあいまい grep コマンド「lsdgrep」作ってみた](http://mattn.kaoriya.net/software/lang/go/20170227010706.htm)
- [Re:ゼロから始めないAPNGエンコーダ - Qiita](http://qiita.com/cia_rana/items/18c78e0233e117b22af6)
- [Go言語を使用して簡単なLineBotを作る - Qiita](http://qiita.com/sao_rio/items/8801b78ba60acbb0ae41)
- [Go言語でコマンドを実行し、一定時間内に終了しなかったらプロセスを強制終了する - Qiita](https://qiita.com/pyjama/items/a61844b11086ab6cbd76)
- [golangでQRコードを生成するパッケージを作ってみた - Qiita](https://qiita.com/toas555/items/763bddbd1992502e62b5)
- [goでwindowsのWMI経由からディスクIOPSを取得してみた - Qiita](https://qiita.com/sky_jokerxx/items/091a70d7b51fc33fe71e)
- [Big Sky :: Golang で優先度を変えてプロセスを起動する。](https://mattn.kaoriya.net/software/lang/go/20171108182710.htm)
- [GoでORMライブラリまわりを綺麗に書く - Qiita](https://qiita.com/tetsuyanh/items/7d807110f602ab150d46)
- [ECHO+GORMでJWTとGraphQLの環境を構築 - Qiita](https://qiita.com/hiroykam/items/31862832a562388d876b)
- [Goの抽象構文木（AST）を手入力してHello, Worldを作る #golang - Qiita](https://qiita.com/tenntenn/items/0cbc6f1f00dc579fcd8c)
- [go-bindata が awesome-go から削除された - Qiita](https://qiita.com/pinzolo/items/5bb88f0fc7343d3a59c6)
- [go-twitterでUserStreamingを取得する - Qiita](https://qiita.com/tiechel/items/b81305694424b3bf6b3c)
- [GolangでのUDP処理メモ - Qiita](https://qiita.com/craftone/items/aa05a104440529b27cdb)
- [ぼくが かんがえた さいきょうの でーたすとあ らっぱー - Qiita](https://qiita.com/vvakame/items/9310bcb5a4e87888d505)
- [Goからlocalのtest用DB(MySQL)をdockerで起動する - Qiita](https://qiita.com/YmgchiYt/items/cc97142614f5b61a69e9)
- [GoでHTMLをPDFに出力する - Qiita](https://qiita.com/kurkuru/items/65614fd3524fefccf576)
- [golang 3ways to iterate - Qiita](https://qiita.com/YmgchiYt/items/fe5936ccbc440cbb6214)
- [flagdayという日本の祝日を Go で扱うライブラリを作った - Qiita](https://qiita.com/pinzolo/items/970b0b980396a1ba0fa0)
- [machineryについて - Qiita](https://qiita.com/yellow/items/829863d7344e7808d8ac)
- [go-prompt v0.2.0の新機能紹介 - Qiita](https://qiita.com/c-bata/items/54eee079cfe3cda02eee)
- [GoアプリケーションをSupervisorでデーモン化😘 - Qiita](https://qiita.com/gericass/items/fa794bfac5c6bd3e0aab)
- [Facebook の Graph API で自分の投稿を取得する (go) - Qiita](https://qiita.com/ekzemplaro/items/b306a3c08c5fd83b5208)
    - [Facebook の Graph API で 団体の情報を得る (go) - Qiita](https://qiita.com/ekzemplaro/items/c269d3e43463b82a81b9)
- [GoでHTMLをPDFに出力する - Qiita](https://qiita.com/kurkuru/items/65614fd3524fefccf576)
- [GoでLet's Encryptの証明書を自動で取得するサーバーを作る - Qiita](https://qiita.com/ruyoumo/items/699634f6c62447669f2b)
- [How to Tar and Un-tar files in Golang – Steve Domino – Medium](https://medium.com/@skdomino/taring-untaring-files-in-go-6b07cf56bc07)
- [[Golang]リアルタイムログ転送+閲覧ツール作った | ブログ :: Web notes.log](http://blog.wnotes.net/blog/article/golang-realtime-log-tailing-tool)
- [Big Sky :: Go言語で ping を打って「にゃーん」を表示させる](https://mattn.kaoriya.net/software/lang/go/20180315230112.htm)
- [Linuxのユーザーランドをinitから全てまるごとgolangで書く - Qiita](https://qiita.com/tetsu_koba/items/aa2d245a61db98299a89)
- [Go言語でAWKを作ってみました - Qiita](https://qiita.com/hideshi/items/3280ae6616319a78c8e3)
- [パズルゲームをGoで作ってみる - Qiita](https://qiita.com/secondarykey/items/2a5bbd35a98153e1b72f)
- [Goでゼロからニューラルネットワークを組んでみた - Qiita](https://qiita.com/shotasakamoto/items/97c17f37c152bb83c654)
- [ブラックジャックをGoで実装してみた。 - Qiita](https://qiita.com/aimof/items/2220bc1f1f0754f62870)
- [golang.org/x/oauth2で色々な認可フローや方言に対応する - Qiita](https://qiita.com/lufia/items/e5344596975676865c3b)
- [io.Pipe関数の２つのdeadlockポイント - Qiita](https://qiita.com/iwaaya/items/59a51706644a6b86b5d6)
- [golangで計算量オーダーを実感する - Qiita](https://qiita.com/soy-curd/items/f5757f6a654c51e75deb)
- [CLI で esa.io の記事を作成するツール esautils を作りました - Qiita](https://qiita.com/uobikiemukot/items/4a34de27a694d9e33649)
- [KAMINASHI（カミナシ） | 誰でも使える食品工場管理アプリ](https://kaminashi.jp/) : バックエンドは Go 言語で書かれているらしい
- [Go GoCSVでShift_JISでCRLFなCSVを作る](https://qiita.com/hiro9/items/ff9333fcf66c8a3f1c3c) : `github.com/gocarina/gocsv` というパッケージが便利っぽい
- [Golangでグラフを描く](https://qiita.com/yutsuki/items/7de97e09289a915f86b9)
- [grpcの練習がてらgoのディレクトリ転送パッケージを作成した](https://qiita.com/youtanagai/items/73557a8be4f643d044b8)
- [Goでods(Open Office Spreadsheet)ファイルを読み込む - Qiita](https://qiita.com/jp_ibis/items/506911b0deaa5ff94687)
- [golangでi18n - Qiita](https://qiita.com/shibukawa/items/f0e4df597e62372fe7d5)
- [Golangで日本語PDFを出力する方法 - Qiita](https://qiita.com/tobita0000/items/f0c2e69a00773cdac9c0)
- [Go初心者が書くarXiv APIを使って論文リストから論文を取ってくるアプリ - Qiita](https://qiita.com/k4saNova/items/61f41a2e56a786cd75b8)
- [Golangでメソッド呼び出しによる部分適用 - Qiita](https://qiita.com/k-motoyan/items/89755685349cbfa956b4)
- [Excelize version 2.0.0 Released - Qiita](https://qiita.com/xuri/items/73488c5f5f8aa02f240c)
- [golangでtemplateの読み込みパスを複数定義する - Qiita](https://qiita.com/kazu22002/items/3e7e167fb238cc4779d7)
- [Dropbox SDK for Goを使ったファイルのアップロードやダウンロード - Qiita](https://qiita.com/maniju/items/0dd9c4c1ae901ce60605)
- [Go言語で日本語フォーマットをpdf出力 - Qiita](https://qiita.com/ikeponsu/items/bb2fc22a2a7969cb622c)
- [Go 1.12でgopdf使ってテンプレートのPDFに色々埋め込み - Qiita](https://qiita.com/tao_s/items/be145dc85169689a2a4f)
- [Big Sky :: 如何にしてけしからん画像を超高速でダウンロードするか](https://mattn.kaoriya.net/software/lang/go/20120322190727.htm)
- [goceleryを使ってみた - Qiita](https://qiita.com/es-h-sugihara/items/a376380dd2ef9b353d99)
- [golangにおけるJSONPの解析 - Qiita](https://qiita.com/shohei-ojs/items/dcc24bf1928fff575838)
- [[Go] レイヤードアーキテクチャの階層構造を守らないimportを警告するlinterを作った - My External Storage](https://budougumi0617.github.io/2019/10/18/launch-layer-for-the-layered-achitectures/)
- [Go: ElasticSearch Clients Study Case - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-elasticsearch-clients-study-case-dbaee1e02c7)
- [Go の標準パッケージにないシステムコールを使う - Qiita](https://qiita.com/navel3/items/2f464163cc0a07458bb1)
