+++
title = "「OpenPGP の実装」セクションを開設した"
date =  "2017-12-01T15:43:34+09:00"
description = "「しっぽのさきっちょ」に散らばってる OpenPGP や GnuPG に関する記事を「OpenPGP の実装」セクションに移動・統合することにした。"
image = "/images/attention/remark.jpg"
tags        = [ "site", "section", "openpgp" ]

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
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[しっぽのさきっちょ](/remark)」に散らばってる [OpenPGP] や [GnuPG] に関する記事を「[OpenPGP の実装](/openpgp/)」セクションに移動・統合することにした。
とりあえず入れ物だけは作った。

前回書いた記事も移動しちゃうのでゴメンナサイ。
今の URL は削除せず残しておいて，移転先 URL をリンクで指示するようにします。

海外の事情は分からないけど，少なくとも日本では [OpenPGP] や [GnuPG] は全くメジャーではないのは分かってるし，私もセキュリティ管理者やめてから15年以上は経ってて真面目に最先端情報を追う気もないのだが（仕事でポストくれるのなら頑張るけど），それでも [GnuPG] は日常的に使ってるツールだし， [OpenPGP] も [RFC 4880bis] に向かって少しずつ動いてる感じはするし， [Go 言語]のお勉強用とはいえ [gpgpdump] とか公開してることもあるので，メンテナンスし易い位置に集約しておこうかな，と考えたのだった。

以上，いいわけ終わり。

そうそう。
セキュリティ情報やアップデート情報など時事ネタに関しては変わらず「[しっぽのさきっちょ](/remark)」で扱うので，今後ともよしなに。

[OpenPGP]: https://www.openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
