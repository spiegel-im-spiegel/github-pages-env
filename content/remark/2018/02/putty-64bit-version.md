+++
title = "64ビット版 PuTTY を導入する"
date = "2018-02-24T12:03:11+09:00"
description = "今回たまたまインストールする機会があったので覚え書きとして残しておく。"
image = "/images/attention/kitten.jpg"
tags = ["security", "cryptography", "windows", "tools", "openssh", "putty"]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

（この記事を書くにあたって古い記事を見直していたのだが昨年のセキュリティ・アップデートの話とか全く書いてなかった。
手元の [PuTTY] はちゃんと追従してたのに。
ちうわけで [PuTTY] の現時点（2018-02-24）のバージョンは 0.70 だが，過去のセキュリティ・アップデートを含んでいるので，**必ず** アップデートすること）

昨年あたりから [PuTTY] の64ビット版が出ているのは気付いていたが，SSH クライアントは日常的に使うもので，万一動かなくなったりしたら非常に困ることになるので手を出しかねていた。
今回たまたま自機以外でインストールする機会があったので覚え書きとして残しておく。

64ビット版 [PuTTY] は以下の公式ダウンロードページ[^pt1] にある。

[^pt1]: Google 検索で [PuTTY] を探すと putty.org とかメチャメチャ怪しげなページが最上位に来るのだが何とかならないのだろうか。ちゃんと公式ページをトップに持ってこいよ。機械任せのフィルタリングに頼ってるからこんなことになるんだよ。

- [Download PuTTY: latest release](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)

で，このままでは日本語環境（ISO-2022-JP や EUC-JP）で苦労するので日本語対応パッチを当てたものを上書きコピーするのだが，私がいつも使っている [PuTTYjp] は64ビット版に対応してないのだ。

- [hdk の自作ソフトの紹介 | PuTTYjp](http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html)
    - [hdk_2 / puttyjp — Bitbucket](https://bitbucket.org/hdk_2/puttyjp)

自力でビルドするのもナニだしなぁと思い，他のものを物色してみることにした。
以下は64ビット版のバイナリを公開しているようだ。

- [iceiv+putty](http://ice.hotmint.com/putty/)
- [PuTTY-ranvis - Ranvis software](http://www.ranvis.com/putty) : HTTPS で繋ぐとうまくいかないのだが...
    - [ranvis/putty: PuTTY custom](https://github.com/ranvis/putty)

[iceiv+putty] のほうはテスト不足ということで何となく消極的だが [PuTTY-ranvis] は commit 履歴を見ても割と積極的な感じがするのでいいかもしれない。
どちらも実行環境をフルセットでパッケージングしているので，取り敢えず試してみるにはいいと思う。

日本語対応版はあくまでも私的なビルドなので依存するのは拙いかなとも思うのだが，自分で面倒見るのは大変なので，今回も有難く使わせていただきます。

[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free telnet/ssh client"
[PuTTYjp]: http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html "hdk の自作ソフトの紹介 | PuTTYjp"
[iceiv+putty]: http://ice.hotmint.com/putty/
[PuTTY-ranvis]: http://www.ranvis.com/putty "PuTTY-ranvis - Ranvis software"
