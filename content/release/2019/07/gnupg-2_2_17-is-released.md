+++
title = "GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり"
date =  "2019-07-10T21:29:53+09:00"
description = "今回の変更で公開鍵サーバ上の公開鍵について付帯する電子署名は（自己署名を除いて）捨てられることになった。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "certification",
  "pki",
  "risk",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

くっそー。
Gmail の野郎が [GnuPG] からのアナウンスを迷惑メールとして処理してくれやがったので気づくのが遅れた。

[GnuPG] 2.2.17 のリリースである。

- [[Announce] GnuPG 2.2.17 released to mitigate attacks on keyservers](https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000439.html)

先日紹介した「[OpenPGP 公開鍵サーバにおける公開鍵の汚染問題]({{< ref "/remark/2019/07/openpgp-certificate-flooding.md" >}})」を受けて公開鍵サーバからのインポートをかなり制限することにしたようだ。

詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.17 released to mitigate attacks on keyservers" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000439.html" lang="en" >}}
* gpg: Ignore all key-signatures received from keyservers.  This change is required to mitigate a DoS due to keys flooded with faked key-signatures.  The old behaviour can be achieved by adding `keyserver-options no-self-sigs-only,no-import-clean` to your `gpg.conf`.  [#4607]
* gpg: If an imported keyblocks is too large to be stored in the keybox (`pubring.kbx`) do not error out but fallback to an import using the options "`self-sigs-only,import-clean`".  [#4591]
* gpg: New command `--locate-external-key` which can be used to refresh keys from the Web Key Directory or via other methods configured with `--auto-key-locate`.
* gpg: New import option "`self-sigs-only`".
* gpg: In `--auto-key-retrieve` prefer WKD over keyservers.  [#4595]
* dirmngr: Support the "`openpgpkey`" subdomain feature from `draft-koch-openpgp-webkey-service-07`. [#4590].
* dirmngr: Add an exception for the "`openpgpkey`" subdomain to the CSRF protection.  [#4603]
* dirmngr: Fix endless loop due to http errors 503 and 504.  [#4600]
* dirmngr: Fix TLS bug during redirection of HKP requests.  [#4566]
* gpgconf: Fix a race condition when killing components.  [#4577]

Release-info: https://dev.gnupg.org/T4606
{{< /fig-quote >}}

今回の変更で公開鍵サーバ上の公開鍵について付帯する電子署名は（自己署名を除いて）捨てられることになった。
以前の動作に戻したければ `gpg.conf` に以下のオプションを追加する。

```
keyserver-options no-self-sigs-only,no-import-clean
```

しかし現時点でこれはおすすめできない。

また公開鍵のインポートで `--self-sigs-only` オプションが使えるようだ。
状況に応じて使い分けるといいだろう。

ちなみに拙作の [gpgpdump] を使って，公開鍵をインポートする前に鍵の構成をチェックできる。
v0.6.0 から公開鍵サーバ上の公開鍵を直接チェックできるようにしたので上手く使っていただけるとありがたい。

- [gpgpdump v0.6.0 をリリースした]({{< ref "/release/2019/07/gpgpdump-v0_6_0-is-released.md" >}})

後半の WKD (Web Key Directory) は 2019-07-10 現在ドラフト08が出ている。

- [OpenPGP Web Key Directory draft-koch-openpgp-webkey-service-07](https://tools.ietf.org/html/draft-koch-openpgp-webkey-service-07)
- [OpenPGP Web Key Directory draft-koch-openpgp-webkey-service-08](https://tools.ietf.org/html/draft-koch-openpgp-webkey-service-08)

今回の一連を受けて標準化と実装が早まるかもしれない。
私もちょっと検討してみようかな。
OpenPGP のメーリングリストでも議論が行われているんだけどスルーしてたんだよなぁ。

さて [Ubuntu] が今回のバージョンをちゃんとリリースしてくれるといいんだけど。
しないのなら本気で私的ビルドを{{< ruby "どげんか" >}}検討{{< /ruby >}}せんといけんかもしれん。

## ブックマーク

- [OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Libgcrypt]: https://gnupg.org/software/libgcrypt/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
