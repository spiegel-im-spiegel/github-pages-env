+++
date = "2015-11-05T20:39:13+09:00"
description = "おおっ！ 次期 OpenPGP の最初のドラフトが登場している。"
draft = false
tags = ["security", "cryptography", "openpgp", "rfc"]
title = "OpenPGP: First RFC4880bis Draft"

[scripts]
  mathjax = false
  mermaidjs = false
+++

おおっ！ 次期 OpenPGP の最初のドラフトが登場している。

- [First 4880bis drafts](https://mailarchive.ietf.org/arch/msg/openpgp/uUKa8eQzWOh3quNElu0BDNrKi2o)
    - [OpenPGP Message Format draft-koch-openpgp-rfc4880bis-00](https://tools.ietf.org/id/draft-koch-openpgp-rfc4880bis-00.txt) （[差分](https://tools.ietf.org/rfcdiff?url2=draft-koch-openpgp-rfc4880bis-00.txt)）
    - [OpenPGP Message Format draft-koch-openpgp-rfc4880bis-01](https://tools.ietf.org/id/draft-koch-openpgp-rfc4880bis-01.txt) （[差分](https://tools.ietf.org/rfcdiff?url2=draft-koch-openpgp-rfc4880bis-01.txt)）

差分を簡単に見れるのは便利だな。

Repository も公開されている。
以下の URI で取得すればよい。

```bash
$ git clone git://git.gnupg.org/people/wk/rfc4880bis.git
```

- [git.gnupg.org Git - people/wk/rfc4880bis.git/summary](http://git.gnupg.org/cgi-bin/gitweb.cgi?p=people/wk/rfc4880bis.git)

現行の [RFC 4880](https://tools.ietf.org/html/rfc4880) が[2007年にリリース](https://baldanders.info/blog/000356/)され，その後に追加された日本国産の Camellia 暗号（[RFC 5581](https://tools.ietf.org/html/rfc5581)）や ECC（Elliptic Curve Cryptography; [RFC 6637](https://tools.ietf.org/html/rfc6637)）が今回ひとつに統合される。
以前 [SHA-3 が OpenPGP に組み込まれるかも](https://baldanders.info/blog/000866/)，という話があったが，今のところドラフト版には記述がない感じ。
Fingerprint を SHA-2 ベースにするとかいう議論もあったような気がするが，これもないか？

まぁ週末にでもゆっくり読むか。

## 参考ページ

- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
