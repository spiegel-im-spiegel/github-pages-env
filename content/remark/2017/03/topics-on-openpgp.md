+++
title = "OpenPGP に関する話題"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "x509",
  "pki",
  "sha-1",
]
description = "GnuPG 2.1.19 がリリース / 映像の証明 / 電子メールの暗号化 / SHA-1 の危殆化と OpenPGP V5"
date = "2017-03-05T17:19:39+09:00"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さてさて。
2月も逃げちゃいましたよ。

1. [GnuPG 2.1.19 がリリース]({{< relref "#gpg" >}})
1. [映像の証明]({{< relref "#pm" >}})
1. [電子メールの暗号化]({{< relref "#em" >}})
1. [SHA-1 の危殆化と OpenPGP V5]({{< relref "#v5" >}})

## GnuPG 2.1.19 がリリース{#gpg}

- [[Announce] GnuPG 2.1.19 released](https://lists.gnupg.org/pipermail/gnupg-announce/2017q1/000402.html)

セキュリティアップデートはなし。
主な修正・変更点は以下の通り。

* gpg: Print a warning if Tor mode is requested but the Tor daemon is not running.
* gpg: New status code `DECRYPTION_KEY` to print the actual private key used for decryption.
* gpgv: New options `--log-file` and `--debug`.
* gpg-agent: Revamp the prompts to ask for card PINs.
* scd: Support for multiple card readers.
* scd: Removed option `--debug-disable-ticker`.  Ticker is used only when it is required to watch removal of device/card.
* scd: Improved detection of card inserting and removal.
* dirmngr: New option `--disable-ipv4`.
* dirmngr: New option `--no-use-tor` to explicitly disable the use of Tor.
* dirmngr: The option `--allow-version-check` is now required even if the option `--use-tor` is also used.
* dirmngr: Handle a missing nsswitch.conf gracefully.
* dirmngr: Avoid PTR lookups for keyserver pools.  The are only done for the debug command "`keyserver --hosttable`".
* dirmngr: Rework the internal certificate cache to support classes of certificates.  Load system provided certificates on startup.  Add options `--tls`, `--no-crl`, and `--systrust` to the "`VALIDATE`" command.
* dirmngr: Add support for the ntbtls library.
* wks: Create mails with a "WKS-Phase" header.  Fix detection of Draft-2 mode.
* The Windows installer is now build with limited TLS support.
* Many other bug fixes and new regression tests.

```text
$ gpg --version
gpg (GnuPG) 2.1.19
libgcrypt 1.7.6
Copyright (C) 2017 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

## 映像の証明{#pm}

[Guardian Project] が提供している [CameraV] というカメラアプリがあるが，この中の Proof Mode という機能を有効にすることで写した写真やビデオの「証明」を作成することができるらしい。

- [Combating “Fake News” With a Smartphone “Proof Mode” – Guardian Project](https://guardianproject.info/2017/02/24/combating-fake-news-with-a-smartphone-proof-mode/)
- ["Proof Mode" for your Smartphone Camera - Schneier on Security](https://www.schneier.com/blog/archives/2017/03/proof_mode_for_.html)

{{< fig-quote title="Combating “Fake News” With a Smartphone “Proof Mode”" link="https://guardianproject.info/2017/02/24/combating-fake-news-with-a-smartphone-proof-mode/" lang="en" >}}
<q>On the technical front, what the app is doing is automatically generating an OpenPGP key for this installed instance of the app itself, and using that to automatically sign all photos and videos at time of capture. A sha256 hash is also generated, and combined with a snapshot of all available device sensor data, such as GPS location, wifi and mobile networks, altitude,  device language, hardware type, and more. This is also signed, and stored with the media. All of this happens with no noticeable impact on battery life or performance, every time the user takes a photo or video.</q>
{{< /fig-quote >}}

まぁ，いまどき映像であってもいくらでも捏造できるからねぇ。
こういう仕組みも必要になってくるというわけだ。
Proof Mode の設計目標は以下の通り。

- Run all of the time in the background without noticeable battery, storage or network impact
- Provide a no-setup-required, automatic new user experience that works without requiring training
- Use strong cryptography for strong identity and verification features, but not encryption
- Produce “proof” sensor data formats that can be easily parse, imported by existing tools (CSV)
- Do not modify the original media files; all proof metadata storied in separate file
- Support chain of custody needs through automatic creation of sha256 hashes and PGP signatures
- Do not require a persistent identity or account generation

内部では映像毎に OpenPGP 電子署名を作成するが，鍵の生成や運用は自動でやってくれるようだ。
便利。
なのだが，イマイチ使い勝手がよく分からない。
使い方はおいおい覚えていくとしよう。

ちなみに [CameraV] では撮った映像を暗号化したり panic 時に映像を全部消去できたりする機能もあるらしい。
最近いろいろ物騒だからねぇ。

- [フォトジャーナリストたちがプロ用カメラに暗号化機能を求めた | TechCrunch Japan](https://jp.techcrunch.com/2016/12/15/20161214photojournalists-demand-encryption-light-is-giving-it-to-them/)
- [暗号化機能をカメラに追加しても写真ジャーナリストたちの問題は解決しない | TechCrunch Japan](https://jp.techcrunch.com/2016/12/16/20161214onpx-n-ovg-onpx-n-ovg-zber/)

## 電子メールの暗号化{#em}

ここのところ暗号化電子メールの話題をよく見かける。

- [グーグルのメール暗号化Chromeアプリケーション「E2EMail」がオープンソースに - ZDNet Japan](https://japan.zdnet.com/article/35097359/)
- [Android版Gmail v7.2でS/MIMEの暗号化機能をサポート ｜ ガジェット通信](http://getnews.jp/archives/1638152)
- [「Chromebook」で効率的に電子メールを暗号化する方法--「K-9 Mail」と「APG」を利用 - TechRepublic Japan](https://japan.techrepublic.com/article/35097236.htm)

「[グーグルのメール暗号化Chromeアプリケーション「E2EMail」がオープンソースに](https://japan.zdnet.com/article/35097359/)」の最後の方に出てくる [Key Transparency](https://security.googleblog.com/2017/01/security-through-transparency.html "Google Online Security Blog: Security Through Transparency") については以前に本家サイトで紹介したので参考にどうぞ。

- [Google による OpenPGP 鍵配送の解決提案 — Baldanders.info](https://baldanders.info/blog/000785/)

個人的には（OpenPGP ではないが）「[Android版Gmail v7.2でS/MIMEの暗号化機能をサポート](http://getnews.jp/archives/1638152)」のほうに注目している。
S/MIME は X.509 型の PKI で鍵を運用するもので Web メールやケータイアプリでの鍵管理がやりやすいのが特徴。
こちらは暗号化より電子署名によるメールの証明がやりやすくなるのではないだろうか。

企業からのプロモーションメールを S/MIME 形式で電子署名することによって HTTPS の EV SSL と同様な効果が狙えると思う。
S/MIME 形式で電子署名が一般的になれば spam メールの排除もやりやすくなるかもしれない。

電子署名ではなく暗号化メールでいうなら，当面は [ProtonMail](https://protonmail.com/ "Secure email: ProtonMail is free encrypted email.") をお勧めする。

- [暗号化メールサービスProtonMailの新規ユーザーが選挙後に急増、トランプ新大統領の不寛容を懸念 | TechCrunch Japan](https://jp.techcrunch.com/2016/11/12/20161111signups-for-encrypted-mail-client-protonmail-double-after-election/)

私もしばらく前にアカウントを作っている。

- [ProtonMail のアカウントを作りました]({{< ref "/remark/2016/03/protonmail.md" >}})

## SHA-1 の危殆化と OpenPGP V5{#v5}

現在の OpenPGP は鍵指紋に SHA-1 を使用しているが， [SHA-1 の危殆化]({{< ref "/remark/2017/02/sha-1-collision.md" >}})に伴い「どーする？」ってな議論になってるようだ。

- [[openpgp] V5 Fingerprint again](https://mailarchive.ietf.org/arch/msg/openpgp/_uV_coJ0CYayv_2ptJMuSraJhws)

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Guardian Project]: https://guardianproject.info/ "Guardian Project – People, Apps and Code You Can Trust"
[CameraV]: https://guardianproject.info/apps/camerav/ "CameraV: Secure Verifiable Photo & Video Camera – Guardian Project"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
