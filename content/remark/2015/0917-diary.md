+++
date = "2015-09-18T03:11:44+09:00"
update = "2015-09-23T19:57:00+09:00"
description = "セキュリティ・暗号関連 / 機械支配待望論"
draft = false
tags = ["security", "cryptography", "cryptrec", "misty1", "rc4", "engineering"]
title = "今日の戯れ言"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"

+++

日を跨いじゃったけど，溜まってたものを吐き出す。

## セキュリティ・暗号関連

とりあえずメモ。

- [エフセキュアブログ : 最悪のプライバシー保護機能を持つ、WINDOWS 10にアップグレードする必要はなし](http://blog.f-secure.jp/archives/50754420.html)
- [デフォルトのままは危険？ 「Windows 10」のプライバシー設定はこう変えよう ｜ ライフハッカー［日本版］](http://www.lifehacker.jp/2015/08/150817win10_privacy.html)
- [自堕落な技術者の日記 : 「RFC 7525 TLSとDTLSの安全な利用に関する推奨事項」の公開 - livedoor Blog（ブログ）](http://blog.livedoor.jp/k_urushima/archives/1768181.html)
- [Windows標準機能でいますぐできる標的型攻撃対策：続・設定を見直すだけ、いますぐ簡単にできる「標的型メール攻撃対策」 (1/4) - ＠IT](http://www.atmarkit.co.jp/ait/articles/1509/16/news007.html) : 「[すぐ実践可能！：設定を見直すだけ、いますぐ簡単にできる「標的型メール攻撃対策」 (1/4) - ＠IT](http://www.atmarkit.co.jp/ait/articles/1409/05/news006.html)」の続編らしい

### 「64ビットブロック暗号MISTY1の安全性について」

- [64ビットブロック暗号MISTY1の安全性について](http://cryptrec.go.jp/topics/cryptrec_20150716_misty1_cryptanalysis.html)
- [64ビットブロック暗号MISTY1の安全性について（続報）](http://cryptrec.go.jp/topics/cryptrec_20150812_misty1_cryptanalysis.html)

随分前の話で恐縮だが，国産のブロック暗号 MISTY1 がちょっとだけ攻略された，という話題。
といっても，鍵空間が $2^{64}$ に対して解読に必要なデータ量が $2^{63.58}$ ということで，ほとんど全件探索と変わらないのであるが。
続報にある改良版にしても計算量は $2^{69.5}$ で大幅に少なくなっているものの，解読に必要なデータ量は $2^{64}$ であり，これも現実的な脅威ではないようである。

MISTY1 は [CRYPTREC 暗号リスト](http://cryptrec.go.jp/list.html)において「推奨候補暗号リスト」のひとつとして挙げられている。
なおブロック暗号については，可能であれば $128\,\mathrm{bits}$ のブロック長のもの（AES や Camellia など）を選択ことが推奨されている。

そうそう。
CRYPTREC といえば [CRYPTREC Report 2014](http://cryptrec.go.jp/topics/cryptrec_20150716_c14report.html) が出てるのだった。
あとでチェックしないと。

### RC4 Overkill

- [RC4 NOMORE](https://www.rc4nomore.com/)
- [短時間でcookie解読、RC4暗号通信を破る新手法 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1507/17/news058.html)

{{< fig-youtube id="d8MtmKrXlKQ" title="The RC4 NOMORE Attack: Demonstration in Practice - YouTube" width="500" height="281" >}}

これも随分前の話でゴメン。
既に[死に体の RC4](http://www.baldanders.info/spiegel/log2/000810.shtml) に追い打ち。

### IETF-94 で OpenPGP WG がなんかやるらしい？

- [[openpgp] Should we try to meet in Yokohama?](https://mailarchive.ietf.org/arch/search/?email_list=openpgp)

[11月に横浜で行われる IETF-94](https://www.ietf.org/meeting/94/index.html) で OpenPGP WG もなにかやるらしい？

## 機械支配待望論

- [We'll Make Great Pets: 機械支配待望論 - 山形浩生の「経済のトリセツ」](http://cruel.hatenablog.com/entry/2015/08/28/161912)

この中で

{{% fig-quote lang="ja" title="機械支配待望論" link="http://cruel.hatenablog.com/entry/2015/08/28/161912" %}}
さらに人間と機械は、出自がちがうので、同じリソースをめぐって争う必要がない。
これが動物だと、居住空間とか食べ物とか毛皮や肉とか、競合する資源がある。
だけど機械とは競合しない……完全にしないとは言わないけれど、他の動物と比べれば大幅にちがう。
人間はお金や女や権力を巡って争ってきたけど、機械はお金とか関係ないし、セックスもしないし（人間のほうはしたがる人もいるけど）、権力も関係ない。
だから機械やAIが賢くなっても、別に人間なんか滅ぼす必要なんかまったくない。
人間にそういうものをエサとして差し出せば（あるいは実物なくても画面にその絵を描いてやるだけで）人間はホイホイ動くし掃除もするしメンテもするし。機械にとってこんな便利なものはないよ。
{{% /fig-quote %}}

というのは全く以ってそのとおりだと思う。

「機械が仕事を奪う」というのは[前回書いた話]({{< ref "remark/2015/information-oriented-society.md" >}})と相似形で，自分たちがやってることをモノに転嫁しているに過ぎない。
奪ってるのはモノではなくヒトである。

もし本当に「ターミネーター脳」の人達が言うように機械が人類を滅ぼすのなら，（今のところただの「自然現象」の延長にすぎない）地球上の知性にとって，間違いなく「進化」である。
そもそも機械なら地球に縛られる必要もなかろう。
どちらかと言うと機械に見捨てられる方を心配したほうがいいんじゃないのか？
