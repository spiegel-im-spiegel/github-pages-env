+++
date = "2017-04-25T22:16:55+09:00"
update = "2017-04-30T21:34:04+09:00"
description = "この手の攻撃の回避法だが， Chrome についてはバージョン 58 以降であれば対応済みである。"
draft = false
tags = ["security", "risk", "web", "unicode", "punycode"]
title = "Punycode によるホモグラフ攻撃例とその回避"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

- [Phishing with Unicode Domains - Xudong Zheng](https://www.xudongz.com/blog/2017/idn-phishing/)

たとえば，以下の [Go 言語]コードで2つの “apple" を考える（元ネタは[ここ](https://play.golang.org/p/BzJVWN78pA "The Go Playground")）。

```go
package main

import "fmt"

func main() {
    for _, value := range "apple" {
        fmt.Printf("%#U\n", value)
    }
    fmt.Println()
    for _, value := range "аррӏе" {
        fmt.Printf("%#U\n", value)
    }
}
```

見た目では分かりにくいかもしれないが，最初の “apple" は US ASCII で2番目の “аррӏе" はキリル文字なんだそうだ[^grp]。
このコードの実行結果は以下の通り。

[^grp]: キリル文字の “аррӏе" の並びに単語としての意味はない。ここでは単純に字形の類似性のみに着目して考える。

```text
U+0061 'a'
U+0070 'p'
U+0070 'p'
U+006C 'l'
U+0065 'e'

U+0430 'а'
U+0440 'р'
U+0440 'р'
U+04CF 'ӏ'
U+0435 'е'
```

現在，国際化ドメイン名（Internationalized Domain Name; IDN）については `xn--` から始まる [punycode] を使った表記が認められている。
更に [punycode] を使った「ホモグラフ攻撃（homograph attack）」については以前から議論があり，少なくとも複数の言語の文字が混在する場合はブラウザ側で Unicode 文字による表記がキャンセルされる。
たとえば `xn-pple-43d.com` は Unicode 表記では `аpple.com` （先頭の а がキリル文字）だが， Chrome や Firefox といった主要ブラウザでは [punycode] のまま `xn-pple-43d.com` と表記される（試さないように）。

しかし複数言語が混在しない場合，つまり最初に挙げたキリル文字だけの “аррӏе" のような場合にはこの制約は効かない。
その言語による真っ当な名前なのかホモグラフ攻撃なのか見分けがつかないからである。
たとえば `xn--80ak6aa92e.com` は `аррӏе.com` だが “аррӏе" の部分は全てキリル文字なので主要ブラウザでも `аррӏе.com` と表示される。

PoC として https://www.xn--80ak6aa92e.com/ が用意されているので，皆さんが使っているブラウザで（証明書の詳細情報も含めて）ドメイン名がどう表示されるか確認して欲しい。

さて，この手の攻撃の回避法だが， Chrome についてはバージョン 58 以降であれば `xn--80ak6aa92e.com` も [punycode] 表記になる。
どういうロジックなのかは不明[^idn]。
Firefox については，設定の ` network.IDN_show_punycode` 項目[^cfg] を true にすれば強制的に [punycode] 表記になる。

[^idn]: Firefox のように [punycode] をまるっと無視するのではないようだ。たとえば Chrome 58 でも「情報処理試験.jp（`xn--n9q36mh1hnxuksz7wt.jp`）」はちゃんと Unicode 表記になる。
[^cfg]: `about:config` から設定する。 “punycode" で検索すれば一発で出てくる。

個人的には国際化ドメイン名は要らんのじゃないかと思うのだが，どうなんだろうねぇ。
果てしなく紛らわしい。

## ブックマーク

- [Chrome, Firefox, and Opera users beware: This isn’t the apple.com you want | Ars Technica](https://arstechnica.com/security/2017/04/chrome-firefox-and-opera-users-beware-this-isnt-the-apple-com-you-want/)
- [Chrome And Firefox Adding Protection Against This Nasty Phishing Trick | McAfee Blogs](https://securingtomorrow.mcafee.com/business/neutralize-threats/chrome-and-firefox-adding-protection-against-this-nasty-phishing-trick/)

- [683314 - Security: Whole-script confusable domain label spoofing (Cyrillic) - chromium - Monorail](https://bugs.chromium.org/p/chromium/issues/detail?id=683314)
- [1332714 - IDN Phishing using whole-script confusables on Windows and Linux](https://bugzilla.mozilla.org/show_bug.cgi?id=1332714)
- [「Google Chrome 58」が正式版に ～“Indexed DB 2.0”対応と29件の脆弱性修正 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1055935.html)
- [本物と偽物の区別がつかないホモグラフ攻撃 | マルウェア情報局](https://eset-info.canon-its.jp/malware_info/special/detail/151001.html)
- [日本語JPドメイン名のPunycode変換・逆変換 - 日本語.jp](http://punycode.jp/)
- [日本語ドメイン→Punycode表記への変換 | IPラーニング](http://www.arearesearch.co.jp/learn/program/06.html)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[punycode]: https://en.wikipedia.org/wiki/Punycode "Punycode - Wikipedia"
