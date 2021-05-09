+++
title = "Amazon アイテム検索・情報取得ツール v0.2.1 をリリースした"
date = "2019-01-27T17:45:46+09:00"
description = "amazon-item review コマンドは v0.2.0 で追加した。せめてバイナリをリリースしておこう。"
image = "/images/attention/tools.png"
tags = [ "tools", "amazon", "market", "amazon-item", "pa-api", "golang", "template" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

なんか Amakuri の作者の方が [amazon-item] について言及してくださっていて

- [Amakuriからの乗り換え先にこれは良いんじゃないか Part2 - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/387)

公開しているツールとはいえ万人に優しい作りではないので，少々あせってたり。
せめてバイナリをリリースしておこう。

というわけで [amazon-item] v0.2.1 をリリースした。

- [Release v0.2.1 · spiegel-im-spiegel/amazon-item · GitHub](https://github.com/spiegel-im-spiegel/amazon-item/releases/tag/v0.2.1)

各種プラットフォーム用の実行バイナリを用意したので [Go 言語]のビルド環境なしに利用できる。
コードには [MIT ライセンスを付与](https://raw.githubusercontent.com/spiegel-im-spiegel/amazon-item/master/LICENSE)している。
ソースコード，実行バイナリともに自由にご利用下さい。

[前の記事]({{< ref "/release/2019/01/amazon-product-advertising-api.md" >}} "しょうがないので Amazon アフィリエイト・リンク作成ツールを作ったですよ")と重複するが，以下に使い方を簡単に説明しておこう。

## 事前準備

[amazon-item] を使うには Amazon Associate ID を持っていて Amazon [Product Advertising API] の Access Key および Secret Key まで取得済みであることが前提となる。
[PA-API] の各キーの取得方法はググると色々出てくるので，ここでは割愛する。

- [AmazonのProduct Advertising APIの使い方 | Apitore blog](http://blog.apitore.com/2016/08/01/amazon-product-advertising-api/)
- [「Amazon API」の使い方を紹介します！最安値やランキング取得できるよ①－アソシエイトID（タグ）登録編－ | HPcode](https://haniwaman.com/amazon-api-1/)

2019年1月23日より [PA-API] の[利用ポリシーが変更](https://affiliate.amazon.co.jp/help/topic/t52/ "[重要] Product Advertising API 利用ポリシーの変更について")になったため，状況によってキーが取得できないことがあるが悪しからず。
これについては[後述]({{< relref "#api-policy" >}})する。

必要な情報が取得できたら `$HOME` ディレクトリ[^home1] に `.paapi.yaml` ファイルを作って値を格納しておく（ファイル・アクセス管理に注意）。
たとえばこんな感じ。

[^home1]: Windows の場合は `%USERPROFILE%` フォルダ直下に `.paapi.yaml` ファイルを置くとよい。

```text
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
```

なお `associate-tag` とは Amazon Associate ID のことである。

## [amazon-item] のインストール

[amazon-item] の[リリースページ](https://github.com/spiegel-im-spiegel/amazon-item/releases "Releases · spiegel-im-spiegel/amazon-item")から各種プラットフォーム向けの書庫ファイルをダウンロードする。
たとえば Windows 64bit 版なら `amazon-item_0.2.1_Windows_64bit.zip` をダウンロードすればよい。

書庫ファイルを展開すると実行ファイル（`amazon-item` または `amazon-item.exe`）があるので，これを `$PATH` の通っているディレクトリに置けば完了である[^path1]。

[^path1]: Windows で `PATH` 環境変数を参照・設定するには[こちらのページ](http://www.atmarkit.co.jp/ait/articles/1805/11/news035.html "Windows 10でPath環境変数を設定／編集する：Tech TIPS - ＠IT")を参考にするとよいだろう。

[amazon-item] はコマンドライン・ツールなので各種 shell あるいはコマンド・プロンプト上で使用する。
試しに以下のコマンドを叩いて動作確認しておこう。

```text
$ amazon-item version
amazon-item 0.2.1
```

また `-h` オプションで簡単な使い方が表示される。

```text
$ amazon-item -h
Searching Amazon Items, Powered by Product Advertising API

Usage:
  amazon-item [flags]
  amazon-item [command]

Available Commands:
  help        Help about any command
  lookup      Lookup Amazon Item
  review      Make review data for Amazon item
  search      Search Amazon Items
  version     Print the version number

Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
  -h, --help                   help for amazon-item
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key

Use "amazon-item [command] --help" for more information about a command.
```

なお，書庫ファイルに入っている `template/` フォルダ内のファイルは出力整形用のテンプレートファイルのサンプルである。

## Amazon アイテムの検索

キーワードをもとに Amazon アイテムを検索するには `amazon-item search` コマンドを使う。
たとえば以下のような感じでキーワードを指定する。

```text
$ amazon-item search "数学ガール+フェルマーの最終定理+kindle"
```

検索結果は JSON 形式で標準出力に出力される。

たとえば [jq] コマンドを使えば JSON 出力を任意にフィルタリング・整形できる。

```text
$ amazon-item search "数学ガール+フェルマーの最終定理+kindle" | jq .
```

またテンプレート・ファイルを指定することで任意のテキスト形式に整形することもできる。
たとえばサンプルの `template/item-list.md` ファイルを使えば

```text
$ amazon-item search -t template/item-list.md 数学ガール+フェルマーの最終定理+kindle
| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
| B00AXUD4EQ | 数学ガール　フェルマーの最終定理　1 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-11-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4EQ/mytag-20 |
| B00I8AT1CM | 数学ガール／フェルマーの最終定理 |  結城 浩 | Kindle版 | SBクリエイティブ | 2008-07-29 | https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20 |
| B00DONBQAI | 数学ガール　フェルマーの最終定理　3 (MFコミックス　フラッパーシリーズ) |  春日 旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2013-06-27 | https://www.amazon.co.jp/exec/obidos/ASIN/B00DONBQAI/mytag-20 |
| B00AXUD4F0 | 数学ガール　フェルマーの最終定理　2 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-11-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4F0/mytag-20 |
| B009T4TN7Y | 数学ガール　上 (MFコミックス　フラッパーシリーズ) |  日坂 水柯 結城 浩 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-09-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B009T4TN7Y/mytag-20 |
| B009T4TNB0 | 数学ガール　下 (MFコミックス　フラッパーシリーズ) |  日坂 水柯 結城 浩 | Kindle版 | KADOKAWA / メディアファクトリー | 2012-09-01 | https://www.amazon.co.jp/exec/obidos/ASIN/B009T4TNB0/mytag-20 |
| B0756XMQBN | 数学ガール フェルマーの最終定理 |  春日旬 春日 旬 | Kindle版 |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B0756XMQBN/mytag-20 |
| B00ZEIEY1E | [まとめ買い] 数学ガール　フェルマーの最終定理（コミックフラッパー） |  春日旬 春日 旬 | Kindle版 |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B00ZEIEY1E/mytag-20 |
```

のように markdown の表形式で出力する。
ちなみに `template/item-list.md` ファイルの中身は以下の通り。

```text
| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.PublicationDate }} | {{ .URL }} |
{{ end }}
```

テンプレート制御は [Go 言語]の標準パッケージ [`text/template`] を利用している。
文法にちょっとクセがあるが，最近はググれば色々と情報が出てくるのでサンプルと併せて各自でチューニングしてほしい。

なお文字エンコーディングは UTF-8 を前提としている[^char1]。
ご注意を。

[^char1]: [PA-API] が吐き出す XML データが UTF-8 というのもあるし [Go 言語]のコード自体が UTF-8 を前提にしているので他の文字エンコーディングを扱うのは[少し面倒]({{< ref "/golang/transform-character-encoding.md" >}} "文字エンコーディング変換")なのですよ。そういえば以前，遊びで[「nkf っぽいなにか」を作った]({{< ref "/remark/2017/12/like-nkf.md" >}})な。

## レビュー用のメタデータを作成する

個々の Amazon アイテムの情報を取得するには `amazon-item lookup` または `amazon-item review` コマンドを使う[^rv1]。
両者の機能はよく似ているが `amazon-item review` コマンドはレビュー用の情報をセットできるのが特徴である。
今回は `amazon-item review` コマンドを使ってレビューを含む Amazon アフィリエイト・リンクを作成する方法を紹介する。

[^rv1]: `amazon-item review` コマンドは v0.2.0 で追加した。

`amazon-item review` コマンドの使い方は以下の通り。

```text
$ amazon-item review -h
Make review data for Amazon item, lookup item by ItemLookup Method

Usage:
  amazon-item review [flags] description

Flags:
  -h, --help                    help for review
  -p, --id-type string          IdType (default "ASIN")
  -d, --item-id string          ItemId
  -r, --rating int              Rating of product
  -g, --response-group string   ResponseGroup (default "Images,ItemAttributes,Small")
      --review-date string      Date of review
  -t, --template string         Template file

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key
```

たとえば ASIN コード `B00I8AT1CM` の情報を出力するならコマンドラインを

```text
$ amazon-item review -d B00I8AT1CM -r 4 "数学ガールめっさ面白い！"
```

などとする。
これも出力は JSON 形式だがテンプレート・ファイルを指定して出力を整形できる。
たとえばサンプルの `template/review.html` ファイルを使うと

```text
$ amazon-item review -d B00I8AT1CM -r 4 -t template/review.html "数学ガールめっさ面白い！"
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20">数学ガール／フェルマーの最終定理</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">数学ガールめっさ面白い！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="2019-01-27">2019-01-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.1)</p>
</div>
```

という感じになる。

レビューの description は標準入力から指定することもできる。
たとえば以下のように cat (Windows なら type) コマンドの出力をパイプで繋いで 

```text
$ cat B00I8AT1CM.txt
数学ガールめっさ面白い！

$ cat B00I8AT1CM.txt | amazon-item review -d B00I8AT1CM -r 4 --review-date 2019-01-27
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20">数学ガール／フェルマーの最終定理</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">数学ガールめっさ面白い！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>mytag-20</a> on <abbr class="dtreviewed" title="2019-01-27">2019-01-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.1)</p>
</div>
```

などとする。
これで Amazon アイテムデータの作成をバッチ処理化できる。

本当は SQLite か何かでデータベース化できればいいんだろうけど，私自身がそんなに頻繁に利用するツールではないので，それは後回しということで。

## [PA-API] 利用ポリシーに注意{#api-policy}

最初に述べたように2019年1月23日より [PA-API] の利用ポリシーが変更になった。

- [[重要] Product Advertising API 利用ポリシーの変更について](https://affiliate.amazon.co.jp/help/topic/t52/)

この記事から少し引用すると

{{% fig-quote type="markdown" title="Product Advertising API 利用ポリシーの変更について" link="https://affiliate.amazon.co.jp/help/topic/t52/" %}}
- 初期リクエスト可能数 : 1日あたり 8,640リクエスト（API利用開始より60日間）
- PA-APIより取得した商品リンクより発生した、過去30日間の発送済み商品売上$0.05（日本円で約5円）ごとに1リクエスト追加
- 1日の最大リクエスト可能数は、1日あたり864,000リクエスト

また、発送済み商品売上が過去30日以内に発生していない場合、PA-APIをご利用いただけなくなる恐れがございます。
{{% /fig-quote %}}

となるらしい[^wp1]。
また商品へのリンク URL についても `https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20` のような従来の形式ではなく [PA-API] が吐き出す長ったらしい URL を使う必要がある。

[^wp1]: WordPress プラグインの Amazon Link Builder は制限の例外らしい。よく分からない。

先程のテンプレート `template/item-list.md` を例にすると

```text
| ASIN | Title | Author | Binding | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ range .ItemAttributes.Author }} {{ . }}{{ end }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.PublicationDate }} | {{ .URL }} |
{{ end }}
```

`{{ range .Items }} ... {{ .URL }} ... {{ end }}` の部分で `{{ .URL }}` ではなく `{{ .DetailPageURL }}` を使う必要があるわけだ。

うちのサイトは積極的に商売をしているわけではなく，書影などのメタデータを合法的に利用したいだけなので「発送済み商品売上が過去30日以内に発生していない」と [PA-API] の利用が停止されるというのは結構厳しい条件だったり[^aa1]。

[^aa1]: Amazon アソシエイト ID があるなら [PA-API] が利用できなくてもアソシエイト・ツールバー等が使えるそうだが，今更あんなの使えるかい。

また Amazon アソシエイトや [PA-API] の利用資格を一度剥奪されると再取得が難しいみたいな話も聞く。
ホンマ個人ユーザにとっては面倒なことである。

## ブックマーク

- [Amakuriからの乗り換え先に、Amazon公式のWordPressプラグイン「Amazon Associates Link Builder」が良さそう - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/331)

[amazon-item]: https://github.com/spiegel-im-spiegel/amazon-item "spiegel-im-spiegel/amazon-item: Searching Amazon Items, Powered by PA-API"
[Product Advertising API]: https://affiliate.amazon.co.jp/assoc_credentials/home
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[jq]: https://stedolan.github.io/jq/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`text/template`]: https://golang.org/pkg/text/template/ "template - The Go Programming Language"
