+++
title = "OpenPGP 公開鍵を新しくした"
date =  "2025-03-30T21:03:33+09:00"
description = "OpenPGP 公開鍵を新しくした / OpenPGP 鍵作成手順 / OpenPGP と LibrePGP"
image = "/images/attention/openpgp.png"
tags = [ "cryptography", "security", "openpgp", "site" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## OpenPGP 公開鍵を新しくした

昨年，ようやく Ubuntu 内の [GnuPG] が 2.4 系に上がったので，いよいよ OpenPGP 公開鍵を新しくすることにした。
新しい鍵は本家サイトの以下のページにある。

- [OpenPGP 公開鍵リスト](https://baldanders.info/pubkeys/)

最初「[来年からまた短期の運用に戻そう]({{< ref "/remark/2021/03/changing-publickey-management.md" >}} "Baldanders.info サイトにおける OpenPGP 鍵管理の変更")」と思っていた。
ここ3年間，有効期限の延長しかしてないけど... 正直ダルい。
というわけで2030年までの長期運用にした。
なぜ無期限にしなかったのかについては後述する。

## OpenPGP 鍵作成手順

今回は

1. 証明専用の主鍵を作成する
2. 署名用の副鍵を作成する
3. 暗号用の副鍵を作成する

という3ステップで作成した[^c0]。

[^c0]: 主鍵から署名機能を分離することで個別に有効期限を設定したりトラブルで署名鍵が（秘密鍵漏えい等で）使えなくなったときに簡単に分離できるようになる。まぁ主鍵が使えなくなったら副鍵もまとめてアウトだけど（笑） 運用する際は（オリジナルのバックアップを取った上で）主鍵の秘密鍵だけ分離して使うやりかたもある。

まずは主鍵の作成[^g1]：

[^g1]: `gpg --quick-generate-key` でコマンドライン上で ユーザID，アルゴリズム，用途，有効期限 を指定できる。楕円曲線暗号の場合は楕円曲線名を指定する。

```text
$ gpg --full-gen-key --expert
gpg (GnuPG) 2.4.4; Copyright (C) 2024 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

ご希望の鍵の種類を選択してください:
   (1) RSA と RSA
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
   (9) ECC (署名と暗号化) *デフォルト
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (13) 既存の鍵
  (14) カードに存在する鍵
あなたの選択は? 11

このECC鍵にありうる操作: Sign Certify Authenticate
現在の認められた操作: Sign Certify

   (S) 署名機能を反転する
   (A) 認証機能を反転する
   (Q) 完了

あなたの選択は? s

このECC鍵にありうる操作: Sign Certify Authenticate
現在の認められた操作: Certify

   (S) 署名機能を反転する
   (A) 認証機能を反転する
   (Q) 完了

あなたの選択は? q
ご希望の楕円曲線を選択してください:
   (1) Curve 25519 *デフォルト
   (2) Curve 448
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
あなたの選択は? 1
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0) 1828
鍵は2030年04月01日 13時21分30秒 JSTで期限切れとなります
これで正しいですか? (y/N) y

GnuPGはあなたの鍵を識別するためにユーザIDを構成する必要があります。

本名: Spiegel
電子メール・アドレス: spiegel@example.com
コメント: main
次のユーザIDを選択しました:
    "Spiegel (main) <spiegel@example.com>"

名前(N)、コメント(C)、電子メール(E)の変更、またはOK(O)か終了(Q)? o

...
```

[GnuPG] 2.4 系では [EdDSA] が既定になった[^c1]。
楕円曲線に `Curve 25519` または `Curve 448` を選択すれば [EdDSA] にできる。

[^c1]: [RFC 9580] では RSA, ElGamal, DSA は deprecated になった。多分この影響で既定から外れたのだろう。なお  [EdDSA] は [NIST FIPS 186-5](https://csrc.nist.gov/publications/detail/fips/186/5/final "FIPS 186-5, Digital Signature Standard (DSS) | CSRC") にて[標準のアルゴリズムになった]({{< ref "/remark/2023/02/nist-fips-186-5.md" >}} "NIST FIPS 186-5 および SP 800-186 正式版がリリースされた")。これで政府調達でもなんでも大手を振って [EdDSA] を使える。

署名用の副鍵の作成：

```text
$ gpg --edit-key --expert 1A69C1BE7345EDF6EF52B2EE21FD7CE9BC554D13
gpg (GnuPG) 2.4.4; Copyright (C) 2024 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

秘密鍵が利用できます。

gpg: 信用データベースの検査
gpg: marginals needed: 3  completes needed: 1  trust model: pgp
gpg: 深さ: 0  有効性:   4  署名:   4  信用: 0-, 0q, 0n, 0m, 0f, 4u
gpg: 深さ: 1  有効性:   4  署名:   0  信用: 2-, 0q, 0n, 1m, 1f, 0u
gpg: 次回の信用データベース検査は、2025-04-01です
sec  ed25519/21FD7CE9BC554D13
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: C
     信用: 究極        有効性: 究極
[  究極  ] (1). Spiegel (main) <spiegel@example.com>

gpg> addkey
ご希望の鍵の種類を選択してください:
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (5) Elgamal (暗号化のみ)
   (6) RSA (暗号化のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (12) ECC (暗号化のみ)
  (13) 既存の鍵
  (14) カードに存在する鍵
あなたの選択は? 10
ご希望の楕円曲線を選択してください:
   (1) Curve 25519 *デフォルト
   (2) Curve 448
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
あなたの選択は? 1
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0) 1828
鍵は2030年04月01日 13時27分57秒 JSTで期限切れとなります
これで正しいですか? (y/N) y
本当に作成しますか? (y/N) y

...
```

今回は `gpg --edit-key` コマンドで対話モードを起動し `addkey` コマンドで署名鍵を追加している[^g2]。
署名鍵も [EdDSA] で作成する。

[^g2]: `gpg --quick-add-key` でコマンドライン上で副鍵の アルゴリズム，用途，有効期限 を指定できる。

さくさく行こう。
次は暗号用の副鍵の作成：

```text
gpg> addkey
ご希望の鍵の種類を選択してください:
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (5) Elgamal (暗号化のみ)
   (6) RSA (暗号化のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (12) ECC (暗号化のみ)
  (13) 既存の鍵
  (14) カードに存在する鍵
あなたの選択は? 12
ご希望の楕円曲線を選択してください:
   (1) Curve 25519 *デフォルト
   (2) Curve 448
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
あなたの選択は? 1
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0) 1828
鍵は2030年04月01日 13時34分28秒 JSTで期限切れとなります
これで正しいですか? (y/N) y
本当に作成しますか? (y/N) y

...
```

[EdDSA] と対となるよう楕円曲線暗号（ECDH アルゴリズム）で `Curve 25519` を選択する。

ついでに作成した鍵に私の眠り猫アイコンを仕込んでおく。

```text
gpg> addphoto

あなたのフォトIDに使う画像を決めてください。画像はJPEGファイルである必
要があります。画像は公開鍵といっしょに格納される、ということを念頭にお
いておきましょう。もし大きな写真を使うと、あなたの鍵も同様に大きくなり
ます! 240x288くらいにおさまる大きさの画像は、使いよいでしょう。

フォトID用のJPEGファイル名を入力してください: /path/to/avatar.jpg
この写真は正しいですか (y/N/q)? y

...
```

最後に `save` コマンドで保存し `edit` モードを終了する。

## OpenPGP と LibrePGP

今回の鍵の変更や拙作 [gpgpdump] のための調査で気がついたが OpenPGP と [LibrePGP] は仕様レベルで違うものらしい。

OpenPGP は2024年7月にリリースされた [RFC 9580] で標準化されているが，それとは別に [LibrePGP] ってのがあって [draft-koch-librepgp] にて議論されているっぽい。

OpenPGP ([RFC 9580]) と [LibrePGP] の最大の違いはパケットバージョンで， [LibrePGP] が [RFC 4880bis] の V5 パケット仕様を引き継いでいるのに対し OpenPGP は [RFC 4880bis] を仕切り直した draft-ietf-openpgp-crypto-refresh の議論の中で V5 パケットを捨て V6 パケットを定義している。
この違いで影響するのが楕円曲線暗号, AEAD (Authenticated Encryption with Associated Data; 認証付き暗号), 鍵指紋, および耐量子コンピュータ暗号 (Post-Quantum cryptography) である。

ちなみに耐量子コンピュータ暗号は OpenPGP 側は [draft-ietf-openpgp-pqc] にて [LibrePGP] 側は [draft-koch-librepgp] にて検討されている。
また公開テスト中の [GnuPG] 2.5 系では [draft-koch-librepgp] 仕様ベースで耐量子コンピュータ暗号の実装を行っているようだ（確認してない）。

私は draft-ietf-openpgp-crypto-refresh で仕切り直される2021年あたりから次期 OpenPGP の議論を追い切れなくなっていて，こんなヘンテコなことになってるのに気づかなかった。
[RFC 9580] の “Authors” を見ると Proton AG や [Sequoia PGP] の中の人が名を連ねているのに対し [draft-koch-librepgp] は [g10 Code GmbH](https://g10code.com/ "g10code.com") の Werner Koch さん（[GnuPG] のメイン開発者）の名前が見える。

なんか面倒臭いことになってる？

ユーザ側が OpenPGP と [LibrePGP] との間の差異で混乱するのは面白くないので，当面の間，以下の運用を行うことをオススメする。

1. V5 および V6 パケットは使わない
2. AEAD は使わない
3. 耐量子コンピュータ暗号には手を出さない

パケットバージョンの組み合わせに関しては [RFC 9580] で以下のように決められている。

{{< fig-gen type="markdown" title="OpenPGP 鍵および署名のバージョン" >}}
| Signing Key <br>Version | Signature Packet <br>Version | One-Pass Signature <br>Packet Version | Generate? |
| :---: | :---: | :---: | :---: |
| v3 | v3 | v3 | No |
| v4 | v3 | v3 | No |
| v4 | v4 | v3 | Yes |
| v6 | v6 | v6 | Yes |
{{< /fig-gen >}}

[RFC 9580] では “Generate?” の項目が “Yes” の組み合わせのみ生成が許容される（“No” の組み合わせは後方互換のために残される）。
この表の下から2番目の組み合わせ（v4-v4-v3）であれば [GnuPG] 2.4 系の通常の操作で作成・処理が可能である。

AEAD に関しては OpenPGP と [LibrePGP] との間に互換性はない。

[GnuPG] 2.4 系で新たに生成した鍵に対し `gpg --edit-key` コマンドを起動し選好（preferences）を `showpref` コマンドで確認する。

```text {hl_lines=[18,21,24,27,30]}
$ gpg --edit-key --expert 1A69C1BE7345EDF6EF52B2EE21FD7CE9BC554D13
gpg (GnuPG) 2.4.4; Copyright (C) 2024 g10 Code GmbH
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

秘密鍵が利用できます。

sec  ed25519/21FD7CE9BC554D13
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: C
     信用: 究極        有効性: 究極
ssb  ed25519/74DCA828DC07B24C
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: S
ssb  cv25519/0E25614837E634DC
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: E
[  究極  ] (1). Spiegel (main) <spiegel@example.com>
[  究極  ] (2)  [jpeg image of size 5707]

gpg> showpref
[  究極  ] (1). Spiegel (main) <spiegel@example.com>
     暗号方式: AES256, AES192, AES, 3DES
     AEAD: OCB
     ダイジェスト: SHA512, SHA384, SHA256, SHA224, SHA1
     圧縮: ZLIB, BZIP2, ZIP, 無圧縮
     機能: MDC, AEAD, 鍵サーバ 修正しない
[  究極  ] (2)  [jpeg image of size 5707]
     暗号方式: AES256, AES192, AES, 3DES
     AEAD: OCB
     ダイジェスト: SHA512, SHA384, SHA256, SHA224, SHA1
     圧縮: ZLIB, BZIP2, ZIP, 無圧縮
     機能: MDC, AEAD, 鍵サーバ 修正しない
```

ここで機能として AEAD がリストアップされている場合は (相手が AEAD を使えるなら) AEAD を暗号化に使用する。
AEAD を使わせないようにするため `OCB` を除くアルゴリズムを `setpref` コマンドで指定する。

```text {hl_lines=[1,23,26,29,32]}
gpg> setpref AES256 AES192 AES SHA512 SHA384 SHA256 SHA224 ZLIB BZIP2 ZIP
優先指定の一覧を設定:
     暗号方式: AES256, AES192, AES, 3DES
     AEAD:
     ダイジェスト: SHA512, SHA384, SHA256, SHA224, SHA1
     圧縮: ZLIB, BZIP2, ZIP, 無圧縮
     機能: MDC, 鍵サーバ 修正しない
優先指定を本当に更新しますか? (y/N) y

sec  ed25519/21FD7CE9BC554D13
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: C
     信用: 究極        有効性: 究極
ssb  ed25519/74DCA828DC07B24C
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: S
ssb  cv25519/0E25614837E634DC
     作成: 2025-03-30  有効期限: 2030-04-01  利用法: E
[  究極  ] (1). Spiegel (main) <spiegel@example.com>
[  究極  ] (2)  [jpeg image of size 5707]

gpg> showpref
[  究極  ] (1). Spiegel (main) <spiegel@example.com>
     暗号方式: AES256, AES192, AES, 3DES
     AEAD:
     ダイジェスト: SHA512, SHA384, SHA256, SHA224, SHA1
     圧縮: ZLIB, BZIP2, ZIP, 無圧縮
     機能: MDC, 鍵サーバ 修正しない
[  究極  ] (2)  [jpeg image of size 5707]
     暗号方式: AES256, AES192, AES, 3DES
     AEAD:
     ダイジェスト: SHA512, SHA384, SHA256, SHA224, SHA1
     圧縮: ZLIB, BZIP2, ZIP, 無圧縮
     機能: MDC, 鍵サーバ 修正しない

gpg> save
```

最後の耐量子コンピュータ暗号についてだが OpenPGP にせよ [LibrePGP] にせよ，まだドラフト段階のため（公開テスト等に参加するのでなければ）急ぐ必要はない。
その辺も含め2030年頃には OpenPGP と [LibrePGP] の歪な状況もどうにか解決してないかなぁ，と希望を込めて2030年で有効期限を切ったのであった。

この記事の内容をもう少し整理して「[OpenPGP の実装]({{< rlnk "/openpgp/" >}})」セクションにまとめる予定（予定は未定）。

## ブックマーク

- [17. Migration from OpenPGP v4 to v6 — OpenPGP for application developers](https://openpgp.dev/book/migration.html)
- [OpenPGP compatibility - GnuPG - ArchWiki](https://wiki.archlinux.org/title/GnuPG#OpenPGP_compatibility)
- [OpenPGPEmailSummit202305Notes - GnuPG wiki](https://wiki.gnupg.org/OpenPGPEmailSummit202305Notes)
- [OpenPGP updates, LibrePGP and RFC 9580](https://didisoft.com/2024/10/08/openpgp-updates-librepgp-and-rfc-9580/)
- [sequoia-pgp / sequoia · GitLab](https://gitlab.com/sequoia-pgp/sequoia)
- [rpgp/rpgp: OpenPGP implemented in pure Rust, permissively licensed](https://github.com/rpgp/rpgp)

- [OpenPGP で利用可能なアルゴリズム（RFC 9580 対応版）]({{< ref "/openpgp/algorithms-for-openpgp-rfc9580.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Sequoia PGP]: https://sequoia-pgp.org/ "Sequoia-PGP"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[RFC 9580]: https://datatracker.ietf.org/doc/html/rfc9580 "RFC 9580 - OpenPGP"
[LibrePGP]: https://librepgp.org/ "LibrePGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[draft-koch-librepgp]: https://datatracker.ietf.org/doc/draft-koch-librepgp/ "draft-koch-librepgp - LibrePGP Message Format"
[draft-ietf-openpgp-pqc]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/ "draft-ietf-openpgp-pqc - Post-Quantum Cryptography in OpenPGP"

[EdDSA]: {{< ref "/remark/2020/06/eddsa.md" >}} "Edwards-curve Digital Signature Algorithm"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
