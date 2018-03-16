+++
title = "XML データの Unmarshalling"
date =  "2018-03-16T13:41:24+09:00"
description = "description"
image = "/images/attention/go-code2.png"
tags = ["golang", "struct", "tags", "marshal", "xml"]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は XML の Unmarshalling について。
つっても Unmarshalling 自体は encoding/[xml] パッケージを使って

```go
stdata := &StructData{}
if err := xml.Unmarshal([]byte(xmldata), stdata); err != nil {
    fmt.Fprintln(os.Stderr, err)
    return
}
```

とすればよい。
問題は XML データを受け入れる構造体（上述のコードで言うなら `StructData`）をどう定義するかだ。

今回のお題となる XML データとして以下を考える。

```xml
<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF
  xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
  xmlns:foaf="http://xmlns.com/foaf/0.1/"
  xmlns:dc="http://purl.org/dc/terms/"
  xmlns:cc="http://web.resource.org/cc/"
>
  <foaf:Document rdf:about="http://text.baldanders.info/">
    <dc:title lang="ja">text.Baldanders.info</dc:title>
    <dc:creator>Spiegel</dc:creator>
    <dc:date>2018-03-13T11:57:52+09:00</dc:date>
    <cc:license rdf:resource="http://creativecommons.org/licenses/by-sa/4.0/"/>
  </foaf:Document>
</rdf:RDF>
```

この XML データの特徴は以下の通り。

1. 要素名や属性名が名前空間（namespace）を含んでいる
2. 要素に対する値として属性値と要素の内容（要素タグで囲まれているエリア）がある
3. 時刻情報（[RFC 3339] 形式）を含んでいる

これを [struct] タグでどのように記述するか。

## 名前空間について

たとえば名前空間を含まない

```xml
<creator>Spiegel</creator>
```

であれば [struct] タグは

```go
Creator string `xml:"creator"`
```

と記述できる。
しかし実際には

```xml
<dc:creator>Spiegel</dc:creator>
```

と名前空間を含んでいるため， [struct] タグを厳密に記述するなら

```go
Creator string `xml:"http://purl.org/dc/terms/ creator"`
```

とする必要がある。
名前空間の指定は短縮名の `dc` ではなく `http://purl.org/dc/terms/` を指定する点に注意。
なお（名前空間を除いた）要素名や属性の混濁がないのであれば

```go
Creator string `xml:"creator"`
```

のままでも問題なく parse できる。

## 値の取得

`<dc:title>` には属性値と要素の内容の複数が含まれている。

```xml
<dc:title lang="ja">text.Baldanders.info</dc:title>
```

これらの値を全て取得するなら [struct] タグの記述は以下のようになる。

```go
Title struct {
    Language string `xml:"lang,attr"`
    Name     string `xml:",chardata"`
} `xml:"http://purl.org/dc/terms/ title"`
```

ちなみに要素の内容が `<![CDATA[ ... ]]>` で囲まれるデータであれば `chardata` ではなく `cdata` を指定するとよい。

## 時刻情報の処理

[`xml`]`.Unmarshal()` 関数には時刻情報を [`time`]`.Time` 型に変換するロジックは用意されていない（基本型ではないので当然だが）。
したがって文字列から [`time`]`.Time` 型に変換するロジックを自前で組み込む必要がある。

今回は簡単に以下のようにした。

```go
type Time struct {
    time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var v string
    if err := d.DecodeElement(&v, &start); err != nil {
        return err
    }
    tm, err := time.Parse(time.RFC3339, v)
    if err != nil {
        return err
    }
    *t = Time{tm}
    return nil
}
```

`Time` 型が [`time`]`.Time` 型のラッパー・クラスになっている点に注目。
この `Time` 型に `UnmarshalXML()` 関数を定義してあげればよい。
これで構造体の要素は以下のように記述できる。

```go
Date Time `xml:"http://purl.org/dc/terms/ date"`
```

[`xml`]`.Unmarshal()` 関数は `Time` 型を解析するために `Time.UnmarshalXML()` 関数を呼び出すわけだ。
encoding/[xml] パッケージでは `UnmarshalXML()` 関数は以下のように定義されている。

```go
type Unmarshaler interface {
    UnmarshalXML(d *Decoder, start StartElement) error
}
```

[`xml`]`.Unmarshaler` インタフェースを持つ型であれば XML Unmarshalling 可能である[^marsh1]。

[^marsh1]: もちろん [`xml`]`.Marshaler` インタフェースも存在する。

## サンプルコード

以上を踏まえてお題の XML データの Unmarshalling コードは以下のようになる。

```go
package main

import (
    "encoding/xml"
    "fmt"
    "os"
    "time"
)

var xmldata = `
<?xml version="1.0" encoding="utf-8"?>
<rdf:RDF
  xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
  xmlns:foaf="http://xmlns.com/foaf/0.1/"
  xmlns:dc="http://purl.org/dc/terms/"
  xmlns:cc="http://web.resource.org/cc/"
>
  <foaf:Document rdf:about="http://text.baldanders.info/">
    <dc:title lang="ja">text.Baldanders.info</dc:title>
    <dc:creator>Spiegel</dc:creator>
    <dc:date>2018-03-13T11:57:52+09:00</dc:date>
    <cc:license rdf:resource="http://creativecommons.org/licenses/by-sa/4.0/"/>
  </foaf:Document>
</rdf:RDF>
`

type Time struct {
    time.Time
}

func (t *Time) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var v string
    if err := d.DecodeElement(&v, &start); err != nil {
        return err
    }
    tm, err := time.Parse(time.RFC3339, v)
    if err != nil {
        return err
    }
    *t = Time{tm}
    return nil
}

type RDF struct {
    XMLName  xml.Name `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# RDF"`
    Document struct {
        URI   string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# about,attr"`
        Title struct {
            Language string `xml:"lang,attr,omitempty"`
            Name     string `xml:",chardata"`
        } `xml:"http://purl.org/dc/terms/ title"`
        Creator string `xml:"http://purl.org/dc/terms/ creator"`
        Date    Time   `xml:"http://purl.org/dc/terms/ date"`
        License struct {
            URI string `xml:"http://www.w3.org/1999/02/22-rdf-syntax-ns# resource,attr"`
        } `xml:"http://web.resource.org/cc/ license"`
    } `xml:"http://xmlns.com/foaf/0.1/ Document"`
}

func main() {
    rdf := &RDF{}
    if err := xml.Unmarshal([]byte(xmldata), rdf); err != nil {
        fmt.Fprintln(os.Stderr, err)
        return
    }
}
```

自前で解析するより遥かにマシだが，面倒くさいことには変わりない。
XML とか○ねばいいのに（笑）

たとえば `<dc:title>` 要素の内容だけ取れればええんじゃあ！ という場合は 

```go
type RDF struct {
    Title string `xml:"Document>title"`
}
```

と要素を `>` で繋ぐ記述も可能。
ただこの記述では名前空間をうまく指定できなかった。
うーん。

XML の Marshalling については機会があれば。
つか，複雑なデータを XML Marshalling するのは不毛な気がする。
フォーマットが決まってるのであればテンプレートを使ったほうが早いんじゃないかなぁ...

## ブックマーク

- [Struct タグについて]({{< relref "golang/struct-tag.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
[struct]: https://golang.org/ref/spec#Struct_types "Struct types"
[xml]: https://golang.org/pkg/encoding/xml/ "xml - The Go Programming Language"
[`xml`]: https://golang.org/pkg/encoding/xml/ "xml - The Go Programming Language"
[`time`]: https://golang.org/pkg/time/ "time - The Go Programming Language"
[RFC 3339]: https://tools.ietf.org/html/rfc3339 "RFC 3339 - Date and Time on the Internet: Timestamps"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/410V3ulwP5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798142417/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798142417.09._SCTHUMBZZZ_.jpg"  alt="スターティングGo言語 (CodeZine BOOKS)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4865940391/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4865940391.09._SCTHUMBZZZ_.jpg"  alt="Kotlinスタートブック -新しいAndroidプログラミング"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4839959234/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4839959234.09._SCTHUMBZZZ_.jpg"  alt="Docker実戦活用ガイド"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4274218961/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4274218961.09._SCTHUMBZZZ_.jpg"  alt="グッド・マス ギークのための数・論理・計算機科学"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-13">2016-07-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
