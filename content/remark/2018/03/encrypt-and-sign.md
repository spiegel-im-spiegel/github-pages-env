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

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/479804413X/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41ZOQaZu0SL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/479804413X/baldandersinf-22/">クラウドを支えるこれからの暗号技術</a></dt><dd>光成 滋生 </dd><dd>秀和システム 2015-06-24</dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4774196053/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4774196053.09._SCTHUMBZZZ_.jpg"  alt="エンジニアリング組織論への招待 ~不確実性に向き合う思考と組織のリファクタリング"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4798053775/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4798053775.09._SCTHUMBZZZ_.jpg"  alt="アプリケーションエンジニアのためのApache Spark入門"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B0788XWJQX/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/B0788XWJQX.09._SCTHUMBZZZ_.jpg"  alt="ソフトウェアデザイン 2018年3月号"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4061538314/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4061538314.09._SCTHUMBZZZ_.jpg"  alt="ブロックチェーン・プログラミング 仮想通貨入門 (KS情報科学専門書)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4797382228/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4797382228.09._SCTHUMBZZZ_.jpg"  alt="暗号技術入門 第3版"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4822258424/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4822258424.09._SCTHUMBZZZ_.jpg"  alt="ブロックチェーン技術の未解決問題"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4757103670/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4757103670.09._SCTHUMBZZZ_.jpg"  alt="ビットコインとブロックチェーン:暗号通貨を支える技術"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4797395451/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4797395451.09._SCTHUMBZZZ_.jpg"  alt="プログラマの数学第2版"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4873118255/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4873118255.09._SCTHUMBZZZ_.jpg"  alt="仕事ではじめる機械学習"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4627847513/baldandersinf-22/" target="_blank"><img src="https://images-fe.ssl-images-amazon.com/images/P/4627847513.09._SCTHUMBZZZ_.jpg"  alt="暗号理論と楕円曲線"  /></a> </p>
<p class="description">まだ読んでない。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-12">2018-03-12</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
