+++
title = "PuTTY 0.81 がリリースされた【セキュリティ・アップデート】"
date =  "2024-04-16T20:42:46+09:00"
description = "NIST P-521 楕円曲線を使った ECDSA 鍵を使ってる人は差し替えが必要"
image = "/images/attention/openpgp.png"
tags = [ "security", "vulnerability", "openssh", "putty", "random", "windows", "ecc" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Windows 用の SSH クライアントソフトである [PuTTY] について ECDSA 署名に絡んで秘密鍵が漏洩するリスクがあるらしい。

- [PuTTY vulnerability vuln-p521-bias](https://www.chiark.greenend.org.uk/~sgtatham/putty/wishlist/vuln-p521-bias.html)

{{< fig-quote type="markdown" title="PuTTY vulnerability vuln-p521-bias" link="https://www.chiark.greenend.org.uk/~sgtatham/putty/wishlist/vuln-p521-bias.html" lang="en" >}}
Every version of the PuTTY tools from 0.68 to 0.80 inclusive has a critical vulnerability in the code that generates signatures from ECDSA private keys which use the NIST P521 curve. (PuTTY, or Pageant, generates a signature from a key when using it to authenticate you to an SSH server.)
{{< /fig-quote >}}

上の引用の通り，対象となるのは [PuTTY] 0,68 から 0.80 までのバージョン。
最新の 0.81 で修正されている。
なお FileZilla, WinSCP, TortoiseGit, TortoiseSVN といったアプリケーションには [PuTTY] がバンドルされていることがあるので，これらのアップデート情報にも注意すること。
[PuTTY] は年に1,2回くらいの頻度でセキュリティ・アップデートが発生するので（ちゃんとメンテされてる証拠），バンドル品は使わないほうがいいと思うなぁ。

日本語カスタム版は PuTTY-ranvis がおすすめ。
こちらも 0.81 がリリースされた。

- [PuTTYrv (PuTTY-ranvis) - Ranvis software](https://www.ranvis.com/putty)

厄介なことに今回はアプリケーションを更新しただけではダメで，認証鍵を新しいものに差し替える必要があるかも知れない。
ただし対象となるのは NIST P-521 楕円曲線を使った ECDSA 鍵のみ。
それ以外はとりあえずセーフということで（笑）

今回の脆弱性は ECDSA 署名時に使用する nonce を生成するための乱数生成に起因するものらしい。
以前から DSA/ECDSA の実装上のネックとなっているのがこの乱数生成部分で，ここの実装をサボると秘密鍵の漏洩に繋がることが知られている。
まぁ，今回はサボったわけではないだろうが，何らかの不備があったということだろう。

というわけで，新たに認証鍵を作成するのなら EdDSA (ed25519 曲線) 鍵にすることをお勧めする。
EdDSA は2023年にリリースされた [NIST FIPS 186-5](https://csrc.nist.gov/pubs/fips/186-5/final "FIPS 186-5, Digital Signature Standard (DSS) | CSRC") に標準として正式に組み込まれたので，仕事用でも大手を振って使ってよい。

## [CVE-2024-31497]

{{< fig-quote type="markdown" title="CVE-2024-31497" link="https://nvd.nist.gov/vuln/detail/CVE-2024-31497" lang="en" >}}
In PuTTY 0.68 through 0.80 before 0.81, biased ECDSA nonce generation allows an attacker to recover a user's NIST P-521 secret key via a quick attack in approximately 60 signatures. This is especially important in a scenario where an adversary is able to read messages signed by PuTTY or Pageant. The required set of signed messages may be publicly readable because they are stored in a public Git service that supports use of SSH for commit signing, and the signatures were made by Pageant through an agent-forwarding mechanism. In other words, an adversary may already have enough signature information to compromise a victim's private key, even if there is no further use of vulnerable PuTTY versions. After a key compromise, an adversary may be able to conduct supply-chain attacks on software maintained in Git. A second, independent scenario is that the adversary is an operator of an SSH server to which the victim authenticates (for remote login or file copy), even though this server is not fully trusted by the victim, and the victim uses the same private key for SSH connections to other services operated by other entities. Here, the rogue server operator (who would otherwise have no way to determine the victim's private key) can derive the victim's private key, and then use it for unauthorized access to those other services. If the other services include Git services, then again it may be possible to conduct supply-chain attacks on software maintained in Git. This also affects, for example, FileZilla before 3.67.0, WinSCP before 6.3.3, TortoiseGit before 2.15.0.1, and TortoiseSVN through 1.14.6.
{{< /fig-quote >}}

以下未稿。

## ブックマーク

- [oss-security - CVE-2024-31497: Secret Key Recovery of NIST P-521 Private Keys  Through Biased ECDSA Nonces in PuTTY Client](https://www.openwall.com/lists/oss-security/2024/04/15/6)
- [「PuTTY」に秘密鍵が復元できてしまう深刻な脆弱性 ～「WinSCP」など他ツールにも影響 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1584589.html)

- [NIST FIPS 186-5 および SP 800-186 正式版がリリースされた]({{< ref "/remark/2023/02/nist-fips-186-5.md" >}})
- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})
- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}}) : GnuPG と PuTTY を連携する

[PuTTY]: https://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[CVE-2024-31497]: https://nvd.nist.gov/vuln/detail/CVE-2024-31497

## 参考

{{% review-paapi "B079NL1L9K" %}} <!-- SSH Mastery -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
