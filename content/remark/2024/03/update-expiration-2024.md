+++
title = "OpenPGP 公開鍵の期限を延長した"
date =  "2024-03-24T12:04:44+09:00"
description = "2024年末で GnuPG 2.2 系はサポートから外れるんだけどねぇ"
image = "/images/attention/kitten.jpg"
tags = [ "cryptography", "security", "openpgp", "site", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

3年前に「[来年からまた短期の運用に戻す]({{< ref "/remark/2021/03/changing-publickey-management.md" >}} "Baldanders.info サイトにおける OpenPGP 鍵管理の変更")」と言ったのに未だに有言不実行。
今年も現行鍵の有効期限を 2025-04-01 まで延長するに留めた。

更新した公開鍵は以下の URI からインポートできる（既にインポート済みの場合は上書き更新される）。

```text
$ gpg --fetch-keys https://baldanders.info/pubkeys/spiegel.asc
```

または

```text
$ gpg --fetch-keys https://github.com/spiegel-im-spiegel.gpg
```

拙作の [gpgpdump] を使えばインポートする前に公開鍵の内容をチェックできる。

```text
$ gpgpdump fetch https://baldanders.info/pubkeys/spiegel.asc
```

または

```text
$ gpgpdump fetch https://github.com/spiegel-im-spiegel.gpg
```

いや Ubuntu が [GnuPG] 2.4 系にアップデートしてくれないので待ち状態なのよ。
いきなり自前で最新版をインストールするのは怖いので，仮想環境か何かで動作確認をしないといけないのだが，これも遅々として進まず。
2024年末で 2.2 系はサポートから外れるんだけどねぇ。

最近は OpenPGP 鍵の短期運用は，期限切れの鍵が増えるだけで，あまり意味がない気がしてきた。
サーバを運用しているならともかく， OpenPGP 鍵を Yubikey とかのデバイスに入れるのもあまり意味がない気がするんだよな。
私は基本的にノート PC を忌避してるし。

まぁ，どのみち OpenPGP 鍵は [ECC にしたい]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}} "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな")ので，いつかは入れ替えなきゃいけないけどね。
それもこれも [GnuPG] 2.4 系の環境が整ってからだ！

...というわけで，もうしばらくはグズってます。

## ブックマーク

- [Installing GnuPG 2.4 on Ubuntu 22.04 | Pro Custodibus](https://www.procustodibus.com/blog/2023/02/gpg-2-4-on-ubuntu-22-04/)

- [OpenPGP 公開鍵リスト | Baldanders.info](https://baldanders.info/pubkeys/)
- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})
- [OpenPGP の電子署名は「ユーザーの身元を保証し」ない]({{< ref "/openpgp/web-of-trust.md" >}})

[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"
[GitHub]: https://github.com/

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B0CW1DZS1W" %}} <!-- 岩波「科学」2024年3月号 現代暗号の展開と応用 -->
