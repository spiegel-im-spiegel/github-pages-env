+++
title = "GnuPG 2.2.18 リリース： さようなら SHA-1"
date =  "2019-11-26T22:12:19+09:00"
description = "2019-01-19 以降に鍵に付与された SHA-1 ベースの電子署名を全て削除する（CVE-2019-14855）。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "cryptography",
  "tools",
  "openpgp",
  "gnupg",
  "risk",
  "hash",
  "sha-1",
  "gpgpdump",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GnuPG] 2.2.18 がリリースされた。

- [[Announce] GnuPG 2.2.18 released](https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html)

今回は（[GnuPG] 自体には）脆弱性もなく通常のメンテナンス・リリースなのだが，ひとつ大きな修正というか対策があって

{{< fig-quote type="markdown" title="GnuPG 2.2.18 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html" lang="en" >}}
{{< quote >}}This release also retires the use of SHA-1 key signatures created since this year.{{< /quote >}}
{{< /fig-quote >}}

らしい。
厳密には 2019-01-19 以降に鍵に付与された SHA-1 ベースの電子署名を全て削除するというもの（[CVE-2019-14855](https://nvd.nist.gov/vuln/detail/CVE-2019-14855)）。

とはいえ，ずいぶん前から [GnuPG] が生成する電子署名は SHA2-256 ベースが既定なので影響は限定的だと思うが[^sig1]， **わざわざ** SHA-1 ベースの電子署名を鍵に付与している方はご注意を。

[^sig1]: ちなみに私が[普段遣いしている鍵](https://baldanders.info/pubkeys/ "OpenPGP 公開鍵リスト")は2013年に作ったものだが， SHA2-256 ベースの電子署名が付与されている。

一応 `--allow-weak-key-signatues` オプションを付けることで今回の措置を回避できるようだが，腹を括ったほうがいいだろう。
なお，鍵への電子署名にどのようなアルゴリズムが使われているかを調べるために拙作の [gpgpdump] を宣伝しておく（笑）

- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

例えば，こんな感じで鍵を取り出して調べることができる。

```text
$ gpg -a --export alice@example.com | gpgpdump
```

その他，詳細はこちら。

{{< fig-quote type="markdown" title="GnuPG 2.2.18 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2019q4/000442.html" lang="en" >}}
* gpg: Changed the way keys are detected on a smartcards; this allows the use of non-OpenPGP cards.  In the case of a not very likely regression the new option `--use-only-openpgp-card` is available.  [#4681]
* gpg: The commands `--full-gen-key` and `--quick-gen-key` now allow direct key generation from supported cards.  [#4681]
* gpg: Prepare against chosen-prefix SHA-1 collisions in key signatures.  This change removes all SHA-1 based key signature newer than 2019-01-19 from the web-of-trust.  Note that this includes all key signature created with dsa1024 keys.  The new option `--allow-weak-key-signatues` can be used to override the new and safer behaviour.  [#4755,CVE-2019-14855]
* gpg: Improve performance for import of large keyblocks.  [#4592]
* gpg: Implement a keybox compression run.  [#4644]
* gpg: Show warnings from dirmngr about redirect and certificate problems (details require `--verbose` as usual).
* gpg: Allow to pass the empty string for the passphrase if the '`--passphase=`' syntax is used.  [#4633]
* gpg: Fix printing of the KDF object attributes.
* gpg: Avoid surprises with `--locate-external-key` and certain `--auto-key-locate` settings.  [#4662]
* gpg: Improve selection of best matching key.  [#4713]
* gpg: Delete key binding signature when deletring a subkey.  [#4665,#4457]
* gpg: Fix a potential loss of key sigantures during import with `self-sigs-only` active.  [#4628]
* gpg: Silence "marked as ultimately trusted" diagnostics if option `--quiet` is used.  [#4634]
* gpg: Silence some diagnostics during in key listsing even with option `--verbose`.  [#4627]
* gpg, gpgsm: Change parsing of agent's pkdecrypt results.  [#4652]
* gpgsm: Support AES-256 keys.
* gpgsm: Fix a bug in triggering a keybox compression run if `--faked-system-time` is used.
* dirmngr: System CA certificates are no longer used for the SKS pool if GNUTLS instead of NTBTLS is used as TLS library.  [#4594]
* dirmngr: On Windows detect usability of IPv4 and IPv6 interfaces to avoid long timeouts.  [#4165]
* scd: Fix BWI value for APDU level transfers to make Gemalto Ezio Shield and Trustica Cryptoucan work.  [#4654,#4566]
* wkd: gpg-wks-client `--install-key` now installs the required policy file.

Release-info: https://dev.gnupg.org/T4684
{{< /fig-quote >}}

アップデートは計画的に。

## ブックマーク

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})

- [SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})
- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/4314009071?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号化 プライバシーを救った反乱者たち</a></dt>
    <dd>スティーブン・レビー (著), 斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>単行本</dd>
    <dd>4314009071 (ASIN), 9784314009072 (EAN), 4314009071 (ISBN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1"><img src="https://m.media-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩 (著)</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>Kindle版</dd>
    <dd>B015643CPE (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-APIv5</a>)</p>
</div>
