+++
title = "最初の SHA-1 衝突例"
date = "2017-02-25T12:38:07+09:00"
tags = ["security", "cryptography", "risk", "hash", "sha-1", "collision"]
description = "もうみんな SHA-1 とはオサラバしてるよね（笑）"

[scripts]
  mathjax = true
  mermaidjs = false
+++

いやぁ，ついにこの日が来たようです。

- [Google Online Security Blog: Announcing the first SHA1 collision](https://security.googleblog.com/2017/02/announcing-first-sha1-collision.html)
- [SHAttered](https://shattered.it/) : SHA-1 の衝突例
- [SHA-1 Collision Found - Schneier on Security](https://www.schneier.com/blog/archives/2017/02/sha-1_collision.html)
- [SHA-1衝突攻撃がついに現実に、Google発表　90日後にコード公開 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1702/24/news067.html)
- [The end of SHA-1 on the Public Web | Mozilla Security Blog](https://blog.mozilla.org/security/2017/02/23/the-end-of-sha-1-on-the-public-web/)
- [グーグル、SHA-1衝突攻撃に成功--同一ハッシュ値の2つのPDFも公開 - ZDNet Japan](https://japan.zdnet.com/article/35097102/)
- [Re: [openpgp] SHA1 collision found](https://mailarchive.ietf.org/arch/msg/openpgp/AjJ3BHzd2c9K2KQ3DTk9Ry_QVYM)
    - [[openpgp] V5 Fingerprint again](https://mailarchive.ietf.org/arch/msg/openpgp/_uV_coJ0CYayv_2ptJMuSraJhws)
- [GoogleのSHA-1のはなし](https://www.slideshare.net/herumi/googlesha1) : 分かりやすい解説
- [SHA-1 collision detection on GitHub.com](https://github.com/blog/2338-sha-1-collision-detection-on-github-com)
    - [GitHub Enterprise、SHA-1衝突を実行不能にするパッチを適用へ -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/1050486.html)

{{< fig-quote title="SHA-1 Collision Found" link="https://www.schneier.com/blog/archives/2017/02/sha-1_collision.html" lang="en" >}}
<q>This is not a surprise. We've all expected this for over a decade, watching computing power increase. This is why NIST standardized SHA-3 in 2012.</q>
{{< /fig-quote >}}

SHA-1 衝突問題については以下を参照のこと。
NIST などでは2014年以降 SHA-1 を電子署名等に使わないよう勧告している。

- [SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})
- [CRYPTREC | SHA-1の安全性について](http://www.cryptrec.go.jp/topics/cryptrec_20151218_sha1_cryptanalysis.html)

現時点で主要な CA では証明書に SHA-1 は使っていないはずである。
また，主要なブラウザについても SHA-1 を使う証明書に対して警告を発するようになっている。

- [SHA-1 ウェブサーバー証明書は 2017 年２月から警告！ウェブサイト管理者は影響の最終確認を – 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/11/25/sha1countdown/)
- [「Google Chrome」の閲覧画面にエラーが！ ～“https://”のサイトにアクセスできない - やじうまの杜 - 窓の杜](http://forest.watch.impress.co.jp/docs/serial/yajiuma/1041798.html)

もうみんな SHA-1 とはオサラバしてるよね（笑）

## 追記というか補足

たとえば git の commit hash 値は SHA-1 で付与されるが大丈夫なのか？ とかいった意見が散見されるが，当面は問題ない。

今回の件はあくまでも電子署名や hash 値そのものを何かの証明に使おうとする場合に問題となる。
git の commit hash 値はあくまで identity として付与されるものである。
改ざんされたかどうかは commit hash 値ではなく差分情報によって容易に知ることができる。

git による悪意のなりすまし等を警戒する必要があるのなら commit hash 値を気にするのではなく commit にきちんと電子署名を行うことをお勧めする（チームで作業する人は是非習慣化するべきである）。

- [Git Commit で OpenPGP 署名を行う]({{< ref "/remark/2016/04/git-commit-with-openpgp-signature.md" >}})

ただし，かつて標準として使われていた MD5 が危殆化とともに廃れていったように，今後 SHA-1 は電子署名以外でも使われなくなると思われる。
念のため， NIST による現在の SHA アルゴリズムの評価と有効期限を以下に示す。

{{< security-strengths-for-hash >}} <!-- 要 MathJax -->
{{< security-strength-time-frames >}} <!-- 要 MathJax -->

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
