+++
date = "2016-05-30T04:14:11+09:00"
description = "どうやらオープンソース・ライセンスに関する情報（メタデータ）を取り出せる仕組みのようだ。"
draft = false
tags = ["license", "tools"]
title = "Open Source License API ?"

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

[yomoyomo さんの記事](http://d.hatena.ne.jp/yomoyomo/20160529/ossapi)から。

- [Announcing the Open Source License API | Open Source Initiative](https://opensource.org/node/822)
- [OpenSourceOrg/api: Light API based on the licenses spec](https://github.com/OpenSourceOrg/api)

う～ん？ どうやらオープンソース・ライセンスに関する情報（メタデータ）を取り出せる仕組みのようだ。

たとえば Apache License 2.0 の情報を取り出すには [https://api.opensource.org/license/Apache-2.0](https://api.opensource.org/license/Apache-2.0) にリクエストを投げる。
すると結果が JSON 形式で返ってくる。
（以下は見やすいように適当に改行を入れている。
各項目の Schema はリポジトリの [doc/endpoints.md](https://github.com/OpenSourceOrg/api/blob/master/doc/endpoints.md) に説明がある）

```json
{
    "id":"Apache-2.0",
    "identifiers":
        [
            {
                "identifier":"Apache-2.0",
                "scheme":"DEP5"
            },
            {
                "identifier":"Apache-2.0",
                "scheme":"SPDX"
            },
            {
                "identifier":"License :: OSI Approved :: Apache Software License",
                "scheme":"Trove"
            }
        ],
    "links":
        [
            {
                "note":"tl;dr legal",
                "url":"https://tldrlegal.com/license/apache-license-2.0-%28apache-2.0%29"
            },
            {
                "note":"Wikipedia page",
                "url":"https://en.wikipedia.org/wiki/Apache_License"
            },
            {
                "note":"OSI Page",
                "url":"https://opensource.org/licenses/Apache-2.0"
            }
        ],
    "name":"Apache License, Version 2.0",
    "other_names":[],
    "superseded_by":null,
    "keywords":
        [
            "osi-approved",
            "popular",
            "permissive"
        ],
    "text":
        [
            {
                "media_type":"text/html",
                "title":"HTML",
                "url":"https://www.apache.org/licenses/LICENSE-2.0"
            }
        ]
}
```

あるいは [SPDX] の識別子を使って [https://api.opensource.org/license/SPDX/Apache-2.0](https://api.opensource.org/license/SPDX/Apache-2.0) にリクエストを投げても同じ結果が帰ってくる。

キーワードで検索することもできるようだ。
たとえば Copyleft なライセンスを探すのであれば [https://api.opensource.org/licenses/copyleft](https://api.opensource.org/licenses/copyleft) とリクエストを投げれば，先程のようなライセンス情報が配列で返ってくる。
ちなみに licenses/ のうしろに何もキーワードを付けないと[登録されている全ライセンス](https://github.com/OpenSourceOrg/licenses "OpenSourceOrg/licenses: machine readable OSI license information")の情報[^0] が返ってくるようだ。

[^0]: ライセンス・データおよび API は CC0 で提供されているようである。 API を利用するなら金よこせとか言うどこぞの企業とは違うらしい（笑）

[リポジトリ](https://github.com/OpenSourceOrg/api "OpenSourceOrg/api: Light API based on the licenses spec")を見ると [Go 言語]による実装がある[^a]。
こんな感じで使える。

[^a]: 他にも [Python](https://github.com/opensourceorg/python-opensource "OpenSourceOrg/python-opensource: Python bindings to the Open Source License API"), [Ruby](https://github.com/opensourceorg/ruby-opensourceapi "OpenSourceOrg/ruby-opensourceapi: Ruby API Bindings to the OSI License API"), [Haskell](https://github.com/OpenSourceOrg/haskell-opensource "OpenSourceOrg/haskell-opensource: Haskell API Bindings to the Open Source License API") による実装がある。

```go
package main

import (
    "fmt"

    "github.com/opensourceorg/api/client"
)

func ohshit(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    license, err := client.Get("Apache-2.0")
    ohshit(err)
    fmt.Printf("%s\n", license.Name)
}
```

これを実行すると以下のような出力になる。

```text
$ go run sample.go
Apache License, Version 2.0
```

キーワードで探すならこんな感じかな。

```go
package main

import (
    "fmt"

    "github.com/opensourceorg/api/client"
)

func ohshit(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    licenses, err := client.Tagged("copyleft")
    ohshit(err)
    for _, license := range licenses {
        fmt.Printf("%s\n", license.Name)
    }
}
```

このコードの実行結果は以下のとおり。

```text
$ go run sample2.go
GNU General Public License, Version 2.0
GNU General Public License, Version 3.0
GNU Lesser General Public License, Version 2.1
GNU Lesser General Public License, Version 3.0
Licence Libre du Québec – Réciprocité forte, Version 1.1
Licence Libre du Québec – Réciprocité, Version 1.1
Mozilla Public License, Version 2.0
```

う～ん。
正直に言って，既に [SPDX] のようなサービスがあるのに，わざわざこのような仕組みを作る意味がよく分からない。

{{< fig-quote title="Announcing the Open Source License API" link="https://opensource.org/node/822"  lang="en">}}
<q>The concept behind this API is to be a "hub" to store a central list of crosswalks and common identifiers to other services, allowing third parties who are already license-aware to provide their mappings, and pull OSI approval status programatically. As a proof of concept, SPDX identifiers have been added, trivially allowing cross-walks to SPDX datasets. This allows anyone to take an SPDX license ID, and determine if it's OSI approved by asking the OSI API.</q>
{{< /fig-quote >}}

OSI がやることに意味があるということだろうか。
あるいは [API に著作権があるとされた昨年の判決](http://www.baldanders.info/spiegel/log2/000861.shtml "Google vs Oracle の訴訟の行方 — Baldanders.info")を受けての防衛措置だったりして。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[SPDX]: https://spdx.org/ "SPDX | Software Package Data Exchange"
