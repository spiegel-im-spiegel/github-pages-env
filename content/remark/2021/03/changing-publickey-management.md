+++
title = "Baldanders.info サイトにおける OpenPGP 鍵管理の変更"
date =  "2021-03-06T19:43:36+09:00"
description = "来年から本気出す（笑）"
image = "/images/attention/openpgp.png"
tags = [ "cryptography", "security", "openpgp", "openssh", "site" ]
pageType = "text"
draft = true

[scripts]
  mathjax = true
  mermaidjs = false
+++

このブログを含む [Baldanders.info] の各サイトでは，私個人の [OpenPGP] 公開鍵を[公開](https://baldanders.info/pubkeys/ "OpenPGP 公開鍵リスト | Baldanders.info")している。
実は2013年から[年次鍵の運用を止めて](https://baldanders.info/blog/000629/)単一で永続的な鍵運用に切り替えたのだが，来年からまた短期の運用に戻すことにした。

具体的には

- 現行の公開鍵の有効期限を 2022-03-31 に変更する（変更済）
- 2022年4月以降に使用する鍵は最長で2年の有効期限付きで作成する
- 単一の鍵で運用するのではなく目的別に鍵を分ける
  - [OpenSSH] のクライアント認証鍵は [GnuPG] ベースのものに切り替える
  - 鍵同士の相互署名はしない

という感じで行こうかと。

理由のひとつは，現在[ドラフト中の NIST FIPS 186-5](https://csrc.nist.gov/publications/detail/fips/186/5/draft "FIPS 186-5 (Draft), Digital Signature Standard (DSS) | CSRC") から DSA が削除されるというもの。

- [FIPS 186-5 (Draft), Digital Signature Standard (DSS) | CSRC](https://csrc.nist.gov/publications/detail/fips/186/5/draft)

もうひとつの理由は [GnuPG] において第3者による [OpenPGP] 鍵への電子署名の社会的信頼度が下がってきたことだ。
もともと，第3者による [OpenPGP] 鍵への電子署名は「小切手の裏書き」みたいなもので，多数の署名が永続的に存在することで鍵の信頼性が担保できていたのだが，その前提が崩れ去ってしまったのだ。

- [OpenPGP 公開鍵サーバにおける公開鍵の汚染問題]({{< ref "/remark/2019/07/openpgp-certificate-flooding.md" >}})

こうなると従来の [OpenPGP] 鍵サーバをベースにした鍵運用は大して意味がないし（せいぜい手渡しよりは便利という程度），単一の鍵を長期的に利用ドメインを跨いで使い回すのは却ってリスクになりかねない。

さらに，強いて3つ目を挙げるなら「電子メールはオワコン」ということだろうか。
少なくとも署名・暗号化しないといけないようなメッセージ・データを電子メールで配送するというのはもうナシだろう。
[PPAP]({{< ref "/remark/2020/06/security-theater.md" >}} "劇場型セキュリティと PPAP") 以前の問題。
[Thunderbird による OpenPGP 鍵利用が使いものにならない]({{< ref "/remark/2020/09/using-gnupg-in-thunderbird-78.md" >}} "Thunderbird 78 系で GnuPG を使う【ただし不完全】")というのもあるが（笑）

暗号鍵の運用については [NIST SP 800-57][SP 800-57] が参考になるだろう。
この中で暗号鍵の使用期間と暗号期間について推奨値が示されている。
こんな感じ。

{{< cryptoperiods >}} <!-- 要 MathJax -->

HTTPS など X.509 ベースで運用される公開鍵の有効期間が1年程度の短期運用に切り替わっているのに気付いている人も多いだろう。
[OpenPGP] 鍵は ad hoc 運用だし合わせる必要はないと思っていたが，永続的な鍵運用に価値がないならこっちに合わせるべきだよね。

とりあえず，現行の [OpenPGP] 公開鍵は

```text
$ gpg --fetch-keys https://baldanders.info/pubkeys/spiegel.asc
```

で有効期限付きのものに更新できる。

ぶっちゃけ個人で鍵を2,3年おきに更新するのって面倒くさいんだよね。
まぁ，実際の運用をどうするか（暗号デバイスの導入も含めて）これから3年以内には確立させる予定。
3年もあれば暗号界隈のシーンも変わっているだろうけど（笑）

## ブックマーク


- [Kernel Maintainer PGP guide — The Linux Kernel  documentation](https://www.kernel.org/doc/html/v5.8/process/maintainer-pgp-guide.html)
- [OpenPGP SSH access with Yubikey and GnuPG · GitHub](https://gist.github.com/artizirk/d09ce3570021b0f65469cb450bee5e29)
- [セキュリティ関連NIST文書：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/publications/nist/)
  - {{< pdf-file title="鍵管理における推奨事項 第一部：一般事項" link="https://www.ipa.go.jp/files/000055490.pdf" >}}
- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[Baldanders.info]: https://baldanders.info/ "Baldanders.info"
[SP 800-57]: https://csrc.nist.gov/publications/detail/sp/800-57-part-1/rev-5/final "SP 800-57 Part 1 Rev. 5, Recommendation for Key Management: Part 1 – General | CSRC"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
