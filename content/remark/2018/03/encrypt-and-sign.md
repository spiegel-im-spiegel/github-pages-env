+++
title = "電子署名を暗号ってゆーな！"
date = "2018-03-12T19:48:30+09:00"
description = "言われてみればその通りで，公開鍵暗号アルゴリズムをもとに組み上げられた電子署名アルゴリズムというのは RSA 署名くらいしかない。"
image = "/images/attention/kitten.jpg"
tags = [ "cryptography" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

面白い記事を見つけた。

- [「電子署名=『秘密鍵で暗号化』」という良くある誤解の話 - Qiita](https://qiita.com/angel_p_57/items/d7ffb9ec13b4dde3357d)

内容を自分の中で咀嚼するのに1時間くらいかかってしまった。
年寄りはこれだから...

要するにこれって

**「電子署名を暗号ってゆーな！」**

ってことだよね。

言われてみればその通りで，公開鍵暗号アルゴリズムをもとに組み上げられた電子署名アルゴリズムというのは RSA 署名くらいしかない。
ElGamal 署名は同じ鍵が使えるというだけでアルゴリズム自体は別物である[^elg1]。

[^elg1]: ちなみに [OpenPGP] では同じ ElGamal 鍵で暗号化と署名を行うことを[禁止した](https://lists.gnupg.org/pipermail/gnupg-users/2003-November/020772.html)。

まぁ，でも，公開鍵暗号といえば今でも RSA なイメージだし，そうなると「電子署名は公開鍵暗号の一種」という印象に引きずられるんだよなぁ。
今後は気を付けよう。

というわけで「[OpenPGP で利用可能なアルゴリズム]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})」の文言を少し変えてみた。
本家で絶賛放置プレイになっている「[わかる！ OpenPGP 暗号](https://baldanders.info/spiegel/cc-license/)」も内容が色々アレなのでいい加減書き直さないといけないんだけど，モチベーションが上がらないんだよねぇ[^openpgp1]。
どうせ [RFC 4880bis] が正式版になったら書き直さないといけないし，もうしばらく放置でいいか。

[^openpgp1]: いや，ほら，最近は『[暗号技術入門](https://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/)』みたいな良書もあるし，暗号技術に対する要件も多様化してるしね。

ところで『[クラウドを支えるこれからの暗号技術](http://herumi.github.io/ango/)』は修正版がまるっと GitHub に上がってる気がするのだがいいのだろうか。
著者の方がいいなら外野がとやかく言うことではないが，私まだ読んでないぞ。
取り敢えず本は買って PDF 版を読むのがいいのか？

[OpenPGP]: http://openpgp.org/
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## ブックマーク

- [OpenPGP の実装](/openpgp/)

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "479804413X" %}} <!-- クラウドを支えるこれからの暗号技術 -->
