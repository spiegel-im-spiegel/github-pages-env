+++
title = "go-myjvn パッケージを作ってみた"
date = "2018-03-17T00:20:43+09:00"
description = "spiegel-im-spiegel/go-myjvn パッケージは JVN が提供する「脆弱性対策情報共有フレームワーク」のひとつである MyJVN API を Go 言語でハンドリングするためのラッパークラスである。"
image = "/images/attention/tools.png"
tags  = [ "golang", "package", "jvn" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ちょっと思いついて [spiegel-im-spiegel/go-myjvn] パッケージを作ってみた。

- [spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang](https://github.com/spiegel-im-spiegel/go-myjvn)

[spiegel-im-spiegel/go-myjvn] パッケージは [JVN] が提供する「[脆弱性対策情報共有フレームワーク]」のひとつである [MyJVN API] を [Go 言語]でハンドリングするためのラッパークラスである。
今のところ，以下の API をサポートしている。

- [脆弱性対策概要情報一覧の取得（getVulnOverviewList）](https://jvndb.jvn.jp/apis/getVulnOverviewList_api_hnd.html "MyJVN - API: getVulnOverviewList")
    - 期間および深刻度によるフィルタリングをサポート
    - 日本語のみサポート
- [脆弱性対策詳細情報の取得（getVulnDetailInfo）](https://jvndb.jvn.jp/apis/getVulnDetailInfo_api_hnd.html "MyJVN - API: getVulnDetailInfo")
    - 日本語のみサポート

サンプルコードとして [`go-myjvn/cli/vulnlist/main.go`](https://github.com/spiegel-im-spiegel/go-myjvn/blob/master/cli/vulnlist/main.go "") も書いてみた。
中身については実際のコードを見てもらうとして（やっつけコードでゴメンペコン），実際に動かすこともできる。
以下のように2つの年月日を指定して，その期間の脆弱性情報を JSON 形式で吐き出す。

```text
$ go run main.go 2018-02-16 2018-03-16
```

ちなみに1ヶ月とかの幅で期間を指定すると100件以上取ってきてワケワカメになるのでご注意を。
[jq] コマンドをかませると JSON 出力をいい感じに整形してくれるので多少読みやすくなるだろう。

```text
$ go run main.go 2018-02-16 2018-03-16 | jq
```

[spiegel-im-spiegel/go-myjvn] パッケージは最低限の機能しかないが，最終的にはこれを利用して脆弱性情報を帳票出力するところまで持っていきたい。
まぁ，余暇でやってることなので牛の如き歩みだけど。

なお，脆弱性情報収集を含めた統合的なソリューションが望みなら [Vuls] がいいらしい。

- [Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD](https://vuls.io/)

Linux および FreeBSD 用だが [Go 言語]製で GPLv3 で公開されている。

## ブックマーク

- [MyJVN API に関する覚え書き]({{< ref "/remark/2018/03/myjvn-api.md" >}})

[spiegel-im-spiegel/go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[go-myjvn]: https://github.com/spiegel-im-spiegel/go-myjvn "spiegel-im-spiegel/go-myjvn: Handling MyJVN RESTful API by Golang"
[JVN]: https://jvn.jp/ "Japan Vulnerability Notes"
[脆弱性対策情報共有フレームワーク]: https://jvndb.jvn.jp/apis/myjvn/ "脆弱性対策情報共有フレームワーク - MyJVN"
[MyJVN API]: https://jvndb.jvn.jp/apis/
[jq]: https://stedolan.github.io/jq/
[Vuls]: https://vuls.io/ "Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD"
