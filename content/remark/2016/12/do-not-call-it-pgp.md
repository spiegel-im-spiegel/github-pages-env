+++
title = "PGP ってゆーな！"
description = "暗号ツールを比較したいのなら，最低でもこの程度は言及していただきたいものである。"
tags = [
  "cryptography",
  "openpgp",
]
date = "2016-12-23T12:57:28+09:00"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（Facebook の TL に書いた戯言の再録＋αと補足）

## [PGPは有害無益？ | スラド セキュリティ](http://security.srad.jp/story/16/12/21/0422242/)

OpenPGP が暗号化メッセージングに向いてないというのは数年前（スノーデン以降）から分かってる話で，その最大の理由は PFS (perfect foward secrecy) を満たすことができないというものだ。

現在の暗号化メッセージングでは「否認可能」（過去のメッセージについて送信者との関連を証明できないこと，つまり「それは私じゃない」と否認できること）であることが要求される。
これは OpenPGP の要件「否認防止」とは真逆のものである。
そもそもの前提が違うのだ。

見方を変えれば「否認防止」が要求されるデータ暗号化や電子署名等では OpenPGP はバリバリ現役である。
これはたとえば Linux 系のパッケージ管理ツールがどうやってパッケージの真正性を担保しているか考えればわかるだろう。

目的に応じて適切な手段を用いることが肝要。こういう話ができないスラドはやっぱクズだよねぇ。

## [暗号化チャットのSignalがAndroidアプリをアップデート、国家による検閲を迂回しステッカーや落書きの添付を可能に | TechCrunch Japan](http://jp.techcrunch.com/2016/12/23/20161222signal-for-android-egypt-uae-stickers/)

うちの端末では Signal がうまく動かないんだよ。
随分前に iPhone 版の Signal と統合されてから調子が悪い。
iPhone なんか使う連中はプライバシーなんか気にしないだろう。
iPhone ユーザは切り捨てて欲しい。
ちなみに WhatsApp や Facebook の暗号化メッセージングには Signal と同じ技術が使われてるよ。

（*【追記 2017-12-23】* 今は問題なく Signal 使えてるよ！）

[^sgnl]: 現在は [Signal](https://whispersystems.org/ "Open Whisper Systems") と SMS/MMS 連携を外している。多分 au のキャリアメールがおかしな運用になってるのが原因だと思う。

OpenPGP を PGP と呼ばわったり（昔はともかく今は PGP はシマンテック社の商品名）暗号化メッセージングに OpenPGP は使えないとか意味不明な比較をする輩にはいい加減ウンザリである。

## 補足

### OpenPGP の要件

OpenPGP は PKI (public key infrastructure; 公開鍵基盤) の一種であり，その要件は以下の4つである。

1. 機密性（Confidentiality）
2. 完全性（Integrity）
3. 認証（Authentication）
4. 否認防止（Non-repudiation）

これに対して，たとえば暗号化メッセージングのひとつである [OTR](https://wiki.xmpp.org/web/OTR "OTR - XMPP WIKI") では以下を要件としている。

1. Encryption (暗号化)
2. Authentication (認証)
3. Deniability (否認可能)
4. Perfect Forward Secrecy （前方秘匿性）

「否認可能」についてはこう説明されている。

{{< fig-quote title="OTR - XMPP WIKI" link="https://wiki.xmpp.org/web/OTR" lang="en" >}}
<q>The messages you send do not have digital signatures that are checkable by a third party. Anyone can forge messages after a conversation to make them look like they came from you. However, during a conversation, your correspondent is assured the messages he sees are authentic and unmodified.</q>
{{< /fig-quote >}}

### 前方秘匿性

実は前方秘匿性 (foward secrecy または perfect foward secrecy) の定義は難しいが NIST の [SP800-52 revision 1] では

{{< fig-quote title="SP800-52 revision 1 : Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations" link="http://dx.doi.org/10.6028/NIST.SP.800-52r1" lang="en" >}}
<q>Perfect forward secrecy is the condition in which the compromise of a long-term private key used in deriving a session key subsequent to the derivation does not cause the compromise of the session key.</q>
{{< /fig-quote >}}

とある。
大雑把に言えば「長期的に使われる秘密鍵が漏洩しても、漏洩前のセッションで交換された鍵の秘密が漏れない」という解釈でいいようだ。 

ちなみに OpenPGP ではセッション鍵を（受信者の公開鍵で暗号化して）暗号文に添付して配信する。
つまり構造的に前方秘匿性を満たせないのである。

### という感じで

暗号ツールを比較したいのなら，最低でもこの程度は言及していただきたいものである。

## ブックマーク

- [Op-ed: I’m throwing in the towel on PGP, and I work in security | Ars Technica](http://arstechnica.com/security/2016/12/op-ed-im-giving-up-on-pgp/)
    - [Op-ed: Why I’m not giving up on PGP | Ars Technica](http://arstechnica.com/information-technology/2016/12/signal-does-not-replace-pgp/)

- [RFC 4880 - OpenPGP Message Format](https://tools.ietf.org/html/rfc4880)
- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)
- [Lavabit 事件とその余波、そして Forward Secrecy - セキュリティは楽しいかね？ Part 2](http://negi.hatenablog.com/entry/2013/11/05/093606)
    - [OTRでオフレコチャット！ - セキュリティは楽しいかね？ Part 2](http://negi.hatenablog.com/entry/2013/11/09/103401)
- [メッセージングは E2E 暗号化および PFS が肝 — Baldanders.info](https://baldanders.info/blog/000675/)
- [CRYPTREC Report 2013 — Baldanders.info](https://baldanders.info/blog/000740/)
- [安全なメッセージング・アプリとは（追記あり） — Baldanders.info](https://baldanders.info/blog/000782/)

- [OpenPGP の実装](/openpgp/)

[SP800-52 revision 1]: https://www.nist.gov/node/562891?pub_id=915295 "Guidelines for the Selection, Configuration, and Use of Transport Layer Security (TLS) Implementations | NIST"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
