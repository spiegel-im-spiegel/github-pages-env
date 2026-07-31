+++
title = "OpenPGP 公開鍵の期限を延長した"
date =  "2022-03-27T12:29:28+09:00"
description = "GitHub の「警戒モード」についても書いてみた"
image = "/images/attention/openpgp.png"
tags = [ "cryptography", "security", "openpgp", "site", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1年ほど前に「[来年からまた短期の運用に戻す]({{< ref "/remark/2021/03/changing-publickey-management.md" >}})」と言ったな。
あれはウソだ。

厳密には有効期限を 2023-03-31 に延長した。
更新した公開鍵は以下の URI からインポートできる（既にインポート済みの場合は上書き更新される）。

```text
$ gpg --fetch-keys https://baldanders.info/pubkeys/spiegel.asc
```

または

```text
$ gpg --fetch-keys https://github.com/spiegel-im-spiegel.gpg
```

また拙作の [gpgpdump] を使えばインポートする前に公開鍵の内容をチェックできる。

```text
$ gpgpdump fetch https://baldanders.info/pubkeys/spiegel.asc
```

または

```text
$ gpgpdump fetch https://github.com/spiegel-im-spiegel.gpg
```

運用の切り替えが遅れているのは単純に本業が忙しくて腰を据えて調査ができないため。
何せ Yubikey のような暗号デバイスを使うかどうかも決めかねているのだ。

もうひとつは，昨年春に [GnuPG] 2.3 系が登場して鍵管理の仕方が若干変わっているのだが，こちらの調査も手つかずで，しかも Ubuntu のク◯野郎が頑なに 2.2 系を堅持してくれやがるという理由もある。
やっぱ Ubuntu は [GnuPG] を見捨てるのかねぇ。
APT の鍵管理もなんか妙なことになってるし。
最悪は自前でビルド&インストールかなぁ。

というわけで，来年こそは本気出す（笑） ということで。

そういえば [GitHub] に更新した公開鍵を登録し直したときに気がついたのだが「警戒モード（Vigilant mode）」てのがあるらしい。

{{< fig-img-quote src="./vigilant-mode.png" title="SSH and GPG keys" link="https://github.com/settings/keys" lang="en" width="803" >}}

今のところベータ版なんだと。

- [Displaying verification statuses for all of your commits - GitHub Docs](https://docs.github.com/authentication/managing-commit-signature-verification/displaying-verification-statuses-for-all-of-your-commits)

通常，コミットやタグに [OpenPGP] 電子署名が施されている場合は，署名に対応する公開鍵が [GitHub] に登録されていれば “Verified” のマークが付けられるが，警戒モードでは

{{< fig-quote class="nobox" type="markdown" title="Displaying verification statuses for all of your commits " link="https://docs.github.com/authentication/managing-commit-signature-verification/displaying-verification-statuses-for-all-of-your-commits" lang="en" >}}
| Status | Description |
| --- | --- |
| Verified | The commit is signed, the signature was successfully verified, and the committer is the only author who has enabled vigilant mode. |
| Partially verified | The commit is signed, and the signature was successfully verified, but the commit has an author who: a) is not the committer and b) has enabled vigilant mode. In this case, the commit signature doesn't guarantee the consent of the author, so the commit is only partially verified. |
| Unverified | Any of the following is true:<br>- The commit is signed but the signature could not be verified.<br>- The commit is not signed and the committer has enabled vigilant mode.<br>- The commit is not signed and an author has enabled vigilant mode |
{{< /fig-quote >}}

と細分化されるそうな。
チームで開発を行っている場合はいいかもね。

## ブックマーク

- [非推奨となったapt-keyの代わりにsigned-byとgnupgを使う方法 - 2021-05-05 - ククログ](https://www.clear-code.com/blog/2021/5/5.html)

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
{{% review-paapi "B00WAMAKZQ" %}} <!-- コマンドー -->
