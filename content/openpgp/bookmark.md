+++
title = "OpenPGP に関するブックマーク"
date =  "2017-12-01T15:15:33+09:00"
description = "OpenPGP に関するリンクをブックマークとして集めてみた（時事ネタは除く）。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "gnupg", "bookmark" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] に関するリンクをブックマークとして集めてみた（時事ネタは除く）。

## Request for Comments

- [RFC 1847 - Security Multiparts for MIME: Multipart/Signed and Multipart/Encrypted](https://tools.ietf.org/html/rfc1847)
- [RFC 3156 - MIME Security with OpenPGP](https://tools.ietf.org/html/rfc3156) (update [RFC 2015])
- [RFC 4880 - OpenPGP Message Format](https://tools.ietf.org/html/rfc4880) (obsoletes [RFC 1991] and [RFC 2440])
    - [RFC 5581 - The Camellia Cipher in OpenPGP](https://tools.ietf.org/html/rfc5581) (update [RFC 4880])
    - [RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP](https://tools.ietf.org/html/rfc6637) (update [RFC 4880])

- [draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format](https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/)
    - [draft-bre-openpgp-samples-00 - OpenPGP Example Keys and Certificates](https://datatracker.ietf.org/doc/draft-bre-openpgp-samples/)
    - [openpgp-wg / rfc4880bis · GitLab](https://gitlab.com/openpgp-wg/rfc4880bis)
    - [openpgp-wg / openpgp-samples · GitLab](https://gitlab.com/openpgp-wg/openpgp-samples)

[RFC 1847]: https://tools.ietf.org/html/rfc1847 "RFC 1847 - Security Multiparts for MIME: Multipart/Signed and Multipart/Encrypted"
[RFC 1991]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[RFC 2015]: https://tools.ietf.org/html/rfc2015 "RFC 2015 - MIME Security with Pretty Good Privacy (PGP)"
[RFC 2440]: https://tools.ietf.org/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[RFC 3156]: https://tools.ietf.org/html/rfc3156 "RFC 3156 - MIME Security with OpenPGP"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 5581]: https://tools.ietf.org/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

## [OpenPGP] の実装および関連ツール

- [Encryption Solutions | Symantec](https://www.symantec.com/products/encryption) : PGP は現在 Symantec 社の暗号化ソリューションの一部となっていて，[商標権も Symantec 社が保持](https://www.symantec.com/about/legal/trademark-policies "Our Current Trademarks | Symantec")している
    - [日本語ページ](https://www.symantec.com/ja/jp/products/encryption)
- [The GNU Privacy Guard](https://gnupg.org/)
    - [Gpg4win - Secure email and file encryption with GnuPG for Windows](https://www.gpg4win.org/)
    - [GPG Suite](https://gpgtools.org/) : macOS 用 GPG Suite
- [OpenPGP.js | OpenPGP JavaScript Implementation](http://openpgpjs.org/)
    - [openpgpjs/openpgpjs: OpenPGP implementation for JavaScript](https://github.com/openpgpjs/openpgpjs)
    - [Mailvelope](https://www.mailvelope.com/)
    - [FlowCrypt: PGP Encryption for Gmail](https://flowcrypt.com/)
- [APG - Google Play](https://play.google.com/store/apps/details?id=org.thialfihar.android.apg) ([Thialfihar](http://thialfihar.org/))
    - [thialfihar/apg: OpenPGP for Android](https://github.com/thialfihar/apg)
- [OpenKeychain](https://www.openkeychain.org/)
    - [OpenKeychain: Easy PGP - Google Play](https://play.google.com/store/apps/details?id=org.sufficientlysecure.keychain)
    - [open-keychain/open-keychain: OpenKeychain is an OpenPGP implementation for Android.](https://github.com/open-keychain/open-keychain)
- [iPGMail](https://ipgmail.com/) : PGP for iOS
- [Sequoia-PGP](https://sequoia-pgp.org/) : Rust 言語による [OpenPGP] 実装
    - [sequoia-pgp / sequoia · GitLab](https://gitlab.com/sequoia-pgp/sequoia)
- [Keybase](https://keybase.io/)
    - [kbpgp.js](https://keybase.io/kbpgp)
    - [[メモ]: keybaseとgit/githubでgpg keyを共有する - Qiita](https://qiita.com/joemphilips/items/7e4d2941448807c4d431)
- [pgpdump]
    - [kazu-yamamoto/pgpdump: A PGP packet visualizer](https://github.com/kazu-yamamoto/pgpdump)
    - [PGPdump Interface](http://www.pgpdump.net/) : Web サービス

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/

## [OpenPGP] 鍵サーバ

- [keys.gnupg.net](http://keys.gnupg.net/ "Nebraska Wesleyan University - OpenPGP Keyserver")
- [pgp.mit.edu](https://pgp.mit.edu/ "MIT PGP Key Server")
- [pgp.nic.ad.jp](http://pgp.nic.ad.jp/ "PGP KEYSERVER")
- [pool.sks-keyservers.net](https://sks-keyservers.net/ "SKS Keyservers")
- [keys.mailvelope.com ](https://keys.mailvelope.com/ "Mailvelope Key Server")

## 参考になる（かもしれない） Web ページ

- [openpgp -- Ongoing discussion of OpenPGP issues.](https://www.ietf.org/mailman/listinfo/openpgp) : ここで次期 OpenPGP の議論などが交わされている
- [PGP User's Manual for Windows](http://www.cla-ri.net/pgp/) : 最近は更新されてないが PGP を巡る歴史資料としての価値は高い
- [GNU Privacy Guard講座](https://gnupg.hclippr.com/) : 最近の [GnuPG] に関する日本語の情報はこちらがお薦め
- [GnuPG - ArchWiki](https://wiki.archlinuxjp.org/index.php/GnuPG)
- [はじめての暗号化メール (Thunderbird編)](https://www.jpcert.or.jp/magazine/security/pgpquick.html) by [JPCERT/CC](https://www.jpcert.or.jp/ "JPCERT コーディネーションセンター")
- [CRYPTREC](http://www.cryptrec.go.jp/)
    - [CRYPTREC暗号リスト(電子政府推奨暗号リスト)](http://www.cryptrec.go.jp/list.html)
    - [CRYPTREC暗号の仕様書](http://www.cryptrec.go.jp/method.html)
    - [CRYPTREC報告書](http://www.cryptrec.go.jp/report.html)
- 旧 OpenPKSD.org（[Internet Archive](https://web.archive.org/web/20110907063003/http://www.openpksd.org/)）からドキュメントをサルベージしてみた
    - {{< pdf-file title="PGP/GPG インターネットで広く使われている暗号技術（白黒バージョン）" link="https://baldanders.info/spiegel/pgpdump/About_PGP_2002.pdf" >}} （2002年）
    - {{< pdf-file title="PGP/GPG インターネットで広く使われている暗号技術（スライド＋ノート）" link="https://baldanders.info/spiegel/pgpdump/About_PGP_2002_NOTE.pdf" >}} （2002年）
    - {{< pdf-file title="OpenPGPとPKI" link="https://baldanders.info/spiegel/pgpdump/PGP-001.pdf" >}} （2002年）
- [OpenPGP(実際はGnuPGなど)の活用事例 - Qiita](https://qiita.com/tsuyoshi_cho/items/b5b37be4f52f0adb0f1b)
- [OpenPGP interoperability test suite](https://tests.sequoia-pgp.org/)
    - [sequoia-pgp / OpenPGP interoperability test suite · GitLab](https://gitlab.com/sequoia-pgp/openpgp-interoperability-test-suite)
- [GnuPG and LDAP](https://gnupg.org/blog/20201018-gnupg-and-ldap.html)
- [GitHub - proglottis/gpgme: Go wrapper for the GPGME library](https://github.com/proglottis/gpgme)
- [Forwarding gpg-agent to a remote system over SSH - GnuPG wiki](https://wiki.gnupg.org/AgentForwarding)
- [gpg-agentをforwardingしてホスト上で署名する - Qiita](https://qiita.com/kazumaemoto/items/b95b8ec5756209b14909)
- [Yubikey5を手に入れたのでGPGを使ってみる(1/2)(GPG key作成編) - akashisnの日記](https://blog.akashisn.info/entry/2020/12/09/132652)
- [cybrkyd - GPG quick-reference cheat sheet](https://cybrkyd.com/post/gpg-quick-reference-cheat-sheet/)

## 自作ツール

- [spiegel-im-spiegel/BkGnuPG: GNU Privacy Guard Plug-in for Becky! 2](https://github.com/spiegel-im-spiegel/BkGnuPG) : 過去の技術的負債（笑）。（セキュリティ的に）危ないので使ってはいけません
- [goark/gpgpdump: OpenPGP packet visualizer](https://github.com/goark/gpgpdump) : [pgpdump] を Go 言語で replace してみた

## [本家サイト]の記事

- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)
- [pgpdump (patched version) — Baldanders.info](https://baldanders.info/spiegel/pgpdump/) : [pgpdump] の Windows 用バイナリを公開中
- [PGP/GnuPG 用語一覧](https://baldanders.info/spiegel/archive/pgp-vocabulary/) : 作りかけ放置プレイ中
- [OpenPGP で利用可能なアルゴリズム一覧（RFC4880 対応版） — Baldanders.info](https://baldanders.info/blog/000452/)
- [Google による OpenPGP 鍵配送の解決提案 — Baldanders.info](https://baldanders.info/blog/000785/)

[OpenPGP]: https://www.openpgp.org/
[本家サイト]: https://baldanders.info/ "Baldanders.info"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
