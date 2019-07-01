+++
title = "しょうがないので Amazon アフィリエイト・リンク作成ツールを作ったですよ"
date = "2019-01-13T22:45:09+09:00"
update = "2019-01-27T17:45:46+09:00"
description = "ただし CLI (Command-Line Interface) ツール。"
image = "/images/attention/go-logo_blue.png"
tags = [ "tools", "amazon", "market", "golang", "engineering", "site", "amazon-item", "template" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨年「[Amazon アフィリエイトリンク作成サービスを Amakuri へ移行]({{< ref "/remark/2018/10/amazon-affiliate-with-amakuri.md" >}})」したと書いたのだが Amazon の締めつけが厳しくて閉鎖するらしい。

- [【Amakuri】2019年01月16日のPA-API利用ポリシー変更で気に掛かること - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/159)
- [【Amakuri】Amakuri は来年1月16日以降使えなくなるだろうという話 - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/194)
- [【Amakuri】来年15日にAmakuriは閉鎖いたします - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/235)

何してくれやがるですか Amazon さん `orz`

しょうがないので，自前で Amazon アフィリエイト・リンク作成ツールを作ったですよ。
ただし CLI (Command-Line Interface) ツール。
下手にサービスなんか作って BAN されるの嫌だし，なにより Web サービスは管理・運用が面倒臭い。

## ツールのインストール

- [GitHub - spiegel-im-spiegel/amazon-item: Searching Amazon Items, Powered by PA-API](https://github.com/spiegel-im-spiegel/amazon-item)

まだテストもしていない出来たてホヤホヤ（笑） まぁ，でも，ひと通りは動く。

バイナリを用意していないので，インストールには [Go 言語] 1.11 以降のビルド環境（git を含む）が必要。
適当なフォルダで以下のコマンドを実行すれば `$GOPATH/bin` フォルダにインストールされる。

```text
$ export GO111MODULE=on
$ go mod init tools
$ go get github.com/spiegel-im-spiegel/amazon-item@latest
```

一応，動作確認。

```text
$ amazon-item -h
Searching Amazon Items, Powered by Product Advertising API

Usage:
  amazon-item [flags]
  amazon-item [command]

Available Commands:
  help        Help about any command
  lookup      Lookup Amazon Item
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

### 【追記 2019-01-27】 バイナリを含む v0.2.1 をリリースした

v0.2.1 をリリースした。

- [Amazon アイテム検索・情報取得ツール v0.2.1 をリリースした]({{< relref "./amazon-item-v0_2_1-is-released.md" >}})

各種プラットフォーム向けのバイナリも併せてリリースしているので，自前でビルドしなくてもよくなった。
改めてよろしく。

## 使用準備

[amazon-item] を使うには Amazon Associate ID を持っていて Amazon [Product Advertising API] の Access Key および Secret Key まで取得済みであることが前提となる。
取得できたら `$HOME` ディレクトリに `.paapi.yaml` ファイルを作って値を格納しておく（ファイル・アクセス管理に注意）。

たとえばこんな感じ。

```text
associate-tag: mytag-20
access-key: AKIAIOSFODNN7EXAMPLE
secret-key: 1234567890
```

なお `associate-tag` とは Amazon Associate ID のことである。

これで準備 OK

## 商品の検索

商品の検索には `amazon-item search` コマンドを使う。

```text
$ amazon-item search -h
Search Amazon Items by ItemSearch Method

Usage:
  amazon-item search [flags] keyword

Flags:
  -h, --help                    help for search
  -g, --response-group string   ResponseGroup (default "ItemAttributes,Small")
  -s, --search-index string     SearchIndex (default "All")
  -t, --template string         Template file

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key
```

たとえば結城浩さんの『数学ガール/フェルマーの最終定理』の Kindle 版を探すなら

```text
amazon-item search 数学ガール+フェルマーの最終定理+kindle
```

などとすればよい。
検索結果は JSON 形式で標準出力に出力される。
[jq] コマンドなどで整形すれば見やすくなるだろう。

```text
amazon-item search 数学ガール+フェルマーの最終定理+kindle | jq .
```

テンプレート・ファイルを用意すれば出力を任意に整形できる。
たとえば

```text
$ cat template/item-list.md
| ASIN | Title | Author | Binding | EAN | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --- | --------- | --------------- | --- |
{{ range .Items }}| {{ .ASIN }} | {{ .ItemAttributes.Title }} | {{ .ItemAttributes.Author }} | {{ .ItemAttributes.Binding }} | {{ .ItemAttributes.EAN }} | {{ .ItemAttributes.Publisher }} | {{ .ItemAttributes.ReleaseDate }} | {{ .URL }} |
{{ end }}
```

のようなテンプレートを用意して以下のように使う。

```text
$ amazon-item search -t template/item-list.md 数学ガール+フェルマーの最終定理+kindle
| ASIN | Title | Author | Binding | EAN | Publisher | PublicationDate | URL |
| ---- | ----- | ------ | ------- | --- | --------- | --------------- | --- |
| B00AXUD4EQ | 数学ガール　フェルマーの最終定理　1 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 |  | KADOKAWA / メディアファクトリー | 2012-12-19 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4EQ/mytag-20 |
| B00I8AT1CM | 数学ガール／フェルマーの最終定理 |  結城 浩 | Kindle版 |  | SBクリエイティブ | 2014-03-12 | https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20 |
| B00DONBQAI | 数学ガール　フェルマーの最終定理　3 (MFコミックス　フラッパーシリーズ) |  春日 旬 | Kindle版 |  | KADOKAWA / メディアファクトリー | 2013-06-27 | https://www.amazon.co.jp/exec/obidos/ASIN/B00DONBQAI/mytag-20 |
| B00AXUD4F0 | 数学ガール　フェルマーの最終定理　2 (MFコミックス　フラッパーシリーズ) |  春日旬 | Kindle版 |  | KADOKAWA / メディアファクトリー | 2012-12-19 | https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4F0/mytag-20 |
| B0756XMQBN | 数学ガール フェルマーの最終定理 |  春日旬 春日 旬 | Kindle版 |  |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B0756XMQBN/mytag-20 |
| B00ZEIEY1E | [まとめ買い] 数学ガール　フェルマーの最終定理（コミックフラッパー） |  春日旬 春日 旬 | Kindle版 |  |  |  | https://www.amazon.co.jp/exec/obidos/ASIN/B00ZEIEY1E/mytag-20 |
```

これで出力結果を markdown の表形式に整形できた。

## Amazon アフィリエイト・リンクを作成する

ある商品に対して Amazon アフィリエイト・リンクを作成するには `amazon-item lookup` コマンドを使う。

```text
$ amazon-item lookup -h
Lookup Amazon Item by ItemLookup Method

Usage:
  amazon-item lookup [flags]

Flags:
  -h, --help                    help for lookup
  -p, --id-type string          IdType (default "ASIN")
  -d, --item-id string          ItemId
  -g, --response-group string   ResponseGroup (default "Images,ItemAttributes,Small")
  -t, --template string         Template file
  -x, --xml                     Output with XML format

Global Flags:
      --access-key string      Access Key ID
      --associate-tag string   Associate Tag
      --config string          config file (default $HOME/.paapi.yaml)
      --marketplace string     Marketplace (default "webservices.amazon.co.jp")
      --secret-key string      Secret Access Key
```

商品を指定するには ASIN コードを使う。

```text
$ amazon-item lookup -d B00I8AT1CM
```

特に指定しなければ JSON 形式で標準出力に出力される。
また `-x` オプションをつければ XML 形式で（というか [PA-API] の返す結果がそのまま）出力される。

`amazon-item lookup` コマンドもテンプレート・ファイルで出力を整形できる。
たとえば

```text
$ cat template/lookup.md
{{ range .Items }}<div class="hreview">
  <div class="photo"><a class="item url" href="{{ .URL }}"><img src="{{ .MediumImage.URL }}" width="{{ .MediumImage.Width }}" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="{{ .URL }}">{{ .ItemAttributes.Title }}</a></dt>
    {{ with .ItemAttributes.Author }}<dd>{{ . }}</dd>{{ end }}
    <dd>{{ .ItemAttributes.Publisher }}{{ with .ItemAttributes.PublicationDate }} {{ . }}{{ end }}{{ with .ItemAttributes.ReleaseDate }} (Release {{ . }}){{ end }}</dd>
    <dd>{{ .ItemAttributes.ProductGroup }} {{ .ItemAttributes.Binding }}</dd>
    <dd>ASIN: {{ .ASIN }}{{ with .ItemAttributes.EAN }}, EAN: {{ . }}{{ end }}</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="{{ $.Today }}">{{ $.Today }}</abbr> (powered by <a href="{{ $.AppURL }}" >{{ $.AppName }}</a> {{ $.AppVersion }})</p>
</div>{{ end }}
```

といった感じにテンプレートファイルを作って

```text
$ amazon-item lookup -d B00I8AT1CM -t template/lookup.html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/mytag-20">数学ガール／フェルマーの最終定理</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-13">2019-01-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.1.0)</p>
</div>
```

などと整形することができる。

これで [PA-API] の利用を Amazon から拒絶されなければ自前で何とかいけそうかな。
でも売上実績がないと [PA-API] が使えなくなるんだよなぁ。
ホンマに勘弁してよ Amazon さん `orz`

## 【追記】 go-amazon-product-api パッケージの置き換え

[amazon-item] では [DDRBoxman/go-amazon-product-api] パッケージを使っているのだが `ItemAttributes` の `Author` および `Creator` 項目の取り扱いに難があって複数の `Author` や `Creator` を上手く取り込めない。
これを改良したものを [spiegel-im-spiegel/go-amazon-product-api] で公開している。

パッケージを置き換えるには，最初のインストールのときにできる `go.mod` ファイルに以下の行を追記する。

```text
replace github.com/DDRBoxman/go-amazon-product-api v0.0.0-20190113062856-6736abd38089 => github.com/spiegel-im-spiegel/go-amazon-product-api v0.0.0-20190113075218-1369f41b1e57
```

この状態で再度

```text
$ go get github.com/spiegel-im-spiegel/amazon-item@latest
```

とすれば [spiegel-im-spiegel/go-amazon-product-api] に置き換えてくれる。
やぁ 1.11 の[モジュール・モード]({{< ref "/golang/go-module-aware-mode.md" >}} "モジュール対応モードへの移行を検討する")は本当に便利だなや。

これで

```html
{{ range .Items }}<div class="hreview">
  <div class="photo"><a class="item url" href="{{ .URL }}"><img src="{{ .MediumImage.URL }}" width="{{ .MediumImage.Width }}" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="{{ .URL }}">{{ .ItemAttributes.Title }}</a></dt>
    {{ if .ItemAttributes.Author }}<dd>{{ range $i, $v := .ItemAttributes.Author }}{{ if ne $i 0 }}, {{ end }}{{ $v }}{{ end }}</dd>{{ end }}
    {{ if .ItemAttributes.Creator }}<dd>{{ range $i, $v := .ItemAttributes.Creator }}{{ if ne $i 0 }}, {{ end }}{{ $v.Value }}{{ with $v.Role }} ({{ . }}){{ end }}{{ end }}</dd>{{ end }}
    <dd>{{ .ItemAttributes.Publisher }}{{ with .ItemAttributes.PublicationDate }} {{ . }}{{ end }}{{ with .ItemAttributes.ReleaseDate }} (Release {{ . }}){{ end }}</dd>
    <dd>{{ .ItemAttributes.ProductGroup }} {{ .ItemAttributes.Binding }}</dd>
    <dd>ASIN: {{ .ASIN }}{{ with .ItemAttributes.EAN }}, EAN: {{ . }}{{ end }}</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="{{ $.Today }}">{{ $.Today }}</abbr> (powered by <a href="{{ $.AppURL }}" >{{ $.AppName }}</a> {{ $.AppVersion }})</p>
</div>{{ end }}
```

のようにテンプレートを記述すれば

```text
$ amazon-item lookup -d B01IGW5IIW -t template/lookup.html
<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B01IGW5IIW/mytag-20"><img src="https://images-fe.ssl-images-amazon.com/images/I/51gC8Tmq1kL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01IGW5IIW/mytag-20">リーン開発の現場 カンバンによる大規模プロジェクトの
運営</a></dt>
    <dd>ＨｅｎｒｉｋＫｎｉｂｅｒｇ, 角谷信太郎</dd>
    <dd>市谷聡啓 (翻訳), 藤原大 (翻訳)</dd>
    <dd>オーム社 2013-10-25 (Release 2017-07-15)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01IGW5IIW</dd>
  </dl>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-13">2019-01-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.1.0)</p>
</div>
```

のように整形される。

[DDRBoxman/go-amazon-product-api] に pull request を送るか悩み中。
割と破壊的変更なんだよね。
もう少し考えてから判断しよう。

## ブックマーク

- [アソシエイト・セントラル](https://affiliate.amazon.co.jp/)
    - [Product Advertising API](https://affiliate.amazon.co.jp/assoc_credentials/home)
    - [Product Advertising API Scratchpad (beta)](https://webservices.amazon.co.jp/scratchpad/)
- [Product Advertising API](https://docs.aws.amazon.com/ja_jp/AWSECommerceService/latest/DG/Welcome.html)
- [Product Advertising API](https://images-na.ssl-images-amazon.com/images/G/09/associates/paapi/dg/index.html) : 内容は古いが日本語

- [AmazonのProduct Advertising APIの使い方 | Apitore blog](http://blog.apitore.com/2016/08/01/amazon-product-advertising-api/)
- [「Amazon API」の使い方を紹介します！最安値やランキング取得できるよ①－アソシエイトID（タグ）登録編－ | HPcode](https://haniwaman.com/amazon-api-1/)
- [Amakuriからの乗り換え先に、Amazon公式のWordPressプラグイン「Amazon Associates Link Builder」が良さそう - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/331)

[amazon-item]: https://github.com/spiegel-im-spiegel/amazon-item "spiegel-im-spiegel/amazon-item: Searching Amazon Items, Powered by PA-API"
[DDRBoxman/go-amazon-product-api]: https://github.com/DDRBoxman/go-amazon-product-api "DDRBoxman/go-amazon-product-api: Wrapper for the Amazon Product Advertising API"
[spiegel-im-spiegel/go-amazon-product-api]: https://github.com/spiegel-im-spiegel/go-amazon-product-api "spiegel-im-spiegel/go-amazon-product-api: Wrapper for the Amazon Product Advertising API"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Product Advertising API]: https://affiliate.amazon.co.jp/assoc_credentials/home
[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[jq]: https://stedolan.github.io/jq/

## 出力例

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E3%83%95%E3%82%A7%E3%83%AB%E3%83%9E%E3%83%BC%E3%81%AE%E6%9C%80%E7%B5%82%E5%AE%9A%E7%90%86-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1CM?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1CM"><img src="https://images-fe.ssl-images-amazon.com/images/I/41vT2D6sERL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%EF%BC%8F%E3%83%95%E3%82%A7%E3%83%AB%E3%83%9E%E3%83%BC%E3%81%AE%E6%9C%80%E7%B5%82%E5%AE%9A%E7%90%86-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00I8AT1CM?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00I8AT1CM">数学ガール／フェルマーの最終定理</a></dt>
    <dd>結城 浩</dd>
    
    <dd>SBクリエイティブ 2008-07-29 (Release 2014-03-12)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00I8AT1CM</dd>
	<dd>評価&nbsp;<abbr class="rating fa-sm" title="5">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </abbr></dd>
  </dl>
  <p class="description">「フェルマーの最終定理」というサブタイトルをみたとき「なんちう大風呂敷を広げるねん」と思ったものだが，実際に読んでみるとぐいぐい引き込まれる。ひっさびさに頭を使ったような気がする。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-13">2019-01-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.1.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%B3%E9%96%8B%E7%99%BA%E3%81%AE%E7%8F%BE%E5%A0%B4-%E3%82%AB%E3%83%B3%E3%83%90%E3%83%B3%E3%81%AB%E3%82%88%E3%82%8B%E5%A4%A7%E8%A6%8F%E6%A8%A1%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%81%8B%E5%96%B6-%EF%BC%A8%EF%BD%85%EF%BD%8E%EF%BD%92%EF%BD%89%EF%BD%8B%EF%BC%AB%EF%BD%8E%EF%BD%89%EF%BD%82%EF%BD%85%EF%BD%92%EF%BD%87-ebook/dp/B01IGW5IIW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01IGW5IIW"><img src="https://images-fe.ssl-images-amazon.com/images/I/51gC8Tmq1kL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%B3%E9%96%8B%E7%99%BA%E3%81%AE%E7%8F%BE%E5%A0%B4-%E3%82%AB%E3%83%B3%E3%83%90%E3%83%B3%E3%81%AB%E3%82%88%E3%82%8B%E5%A4%A7%E8%A6%8F%E6%A8%A1%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%81%8B%E5%96%B6-%EF%BC%A8%EF%BD%85%EF%BD%8E%EF%BD%92%EF%BD%89%EF%BD%8B%EF%BC%AB%EF%BD%8E%EF%BD%89%EF%BD%82%EF%BD%85%EF%BD%92%EF%BD%87-ebook/dp/B01IGW5IIW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01IGW5IIW">リーン開発の現場 カンバンによる大規模プロジェクトの運営</a></dt>
    <dd>ＨｅｎｒｉｋＫｎｉｂｅｒｇ, 角谷信太郎</dd>
    <dd>市谷聡啓 (翻訳), 藤原大 (翻訳)</dd>
    <dd>オーム社 2013-10-25 (Release 2017-07-15)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01IGW5IIW</dd>
	<dd>評価&nbsp;<abbr class="rating fa-sm" title="4">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="far fa-star"></i>
    </abbr></dd>
  </dl>
  <p class="description">私はこれで勉強しました。もう一回読み直すかな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-01-13">2019-01-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.1.0)</p>
</div>
