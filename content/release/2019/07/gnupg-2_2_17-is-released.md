+++
title = "GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり"
date =  "2019-07-10T21:29:53+09:00"
description = "今回の変更で公開鍵サーバ上の公開鍵について付帯する電子署名は（自己署名を除いて）捨てられることになった。"
image = "/images/attention/tools.png"
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

{{< fig-quote type="md" title="GnuPG 2.2.17 released to mitigate attacks on keyservers" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q3/000439.html" lang="en" >}}
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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
