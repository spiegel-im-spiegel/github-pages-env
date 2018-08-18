+++
title = "go-myjvn パッケージ v0.2.0 をリリースした"
date = "2018-04-08T18:20:56+09:00"
description = "オプション周りを少しいじった。あと CVSSv3 用のサブ・パッケージを作った。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "jvn", "cvss" ]

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
  mathjax = false
  mermaidjs = false
+++

[JVN] が提供する [MyJVN API] のラッパークラス [spiegel-im-spiegel/go-myjvn] の v0.2.0 をリリースした。

- [Release v0.2.0 · spiegel-im-spiegel/go-myjvn](https://github.com/spiegel-im-spiegel/go-myjvn/releases/tag/v0.2.0)

オプション周りを少しいじった。
「[脆弱性対策概要情報一覧の取得](https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html "MyJVN - API: getVulnOverviewList")」の発見日，発行日，更新日の各条件をそれぞれ独立で指定できるようにした。
どうも発見日，発行日，更新日の各条件は論理積で有効になるみたいなので（つか，各オプションは全て論理積で有効になるようだ）。

あと CVSSv3 用のサブ・パッケージを作った。
とりあえず Base Metrics のみ。
こんな感じで使える。

```go
import (
    "fmt"

    "github.com/spiegel-im-spiegel/go-myjvn/cvssv3/base"
)

m, err := base.Decode("CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N") //CVE-2015-8252
if err != nil {
    return
}
fmt.Println("Score =", m.Score())
fmt.Println("Severity =", m.GetSeverity())
v, err := m.Encode()
if err != nil {
    return
}
fmt.Println("CVSS Vector =", v)
//Output
//Score = 7.5
//Severity = High
//CVSS Vector = CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N
```

残りの評価はおいおい。

## ブックマーク

- [MyJVN API に関する覚え書き]({{< ref "/remark/2018/03/myjvn-api.md" >}})

[spiegel-im-spiegel/go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[MyJVN API]: https://jvndb.jvn.jp/apis/
