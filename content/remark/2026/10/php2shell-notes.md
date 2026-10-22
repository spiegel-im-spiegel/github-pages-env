+++
title = "php2shell に関する覚え書き"
date =  "2026-10-22T12:25:02+09:00"
description = "個人的には公開スケジュールが前倒しになったのが気になる。対策はきちんと行っておくとして，もうしばらくは注視しておいたほうがよさそうだ。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "vulnerability", "php", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

正直に言って今回は対岸の火事だったのだが，世間的には大事になってる。
既にあちこちでまとめページができてるが，自分用の覚え書きとして，ここでも記しておく。

通称 [php2shell](https://www.wiz.io/blog/php2shell-php-rce "PHP2Shell: AI uncovers critical RCE vulnerabilities affecting 70% of the Internet | Wiz Blog") は PHP 本体に見つかった2件の重大なリモートコード実行（RCE）脆弱性を指す。

- [PHP2Shell: AI uncovers critical RCE vulnerabilities affecting 70% of the Internet | Wiz Blog](https://www.wiz.io/blog/php2shell-php-rce)
- [RCE via OOM corruption](https://github.com/php/php-src/security/advisories/GHSA-w8hw-w9fx-j79w) ([CVE-2026-109618](https://nvd.nist.gov/vuln/detail/cve-2026-109618 "NVD-CVE-2026-109618"))
- [RCE via stale global state](https://github.com/php/php-src/security/advisories/GHSA-6wmf-4rv5-p3rp) ([CVE-2026-109619](https://nvd.nist.gov/vuln/detail/cve-2026-109619 "NVD-CVE-2026-109619"))
- [Security advisory: Critical remote code execution vulnerabilities, update now - PHP: News Archive - 2026](https://www.php.net/archive/2026.php#2026-10-21-5)

CVE ID の末尾番号が十万台に達してるよ。
脆弱性の内容を [Kagi Assistant] に簡単に要約してもらった。

{{< div-ai type="markdown" >}}
| 脆弱性 | 概要 |
|---|---|
| CVE-2026-109618 | メモリ不足（OOM）時の処理で，初期化途中の内部データが不正な状態で残り，リクエスト終了処理がそれを利用することでRCEにつながる可能性があります。[詳細](https://github.com/php/php-src/security/advisories/GHSA-w8hw-w9fx-j79w) |
| CVE-2026-109619 | リクエスト終了後に解放されたメモリを指す内部状態がワーカーに残り，同じワーカーが処理する後続リクエストで使われることで，use-after-freeやRCEにつながる可能性があります。[詳細](https://github.com/php/php-src/security/advisories/GHSA-6wmf-4rv5-p3rp) |
{{< /div-ai >}}

両者とも CVSSv4 スコアが 9.2 (Critical) と非常に高い。
PHP 開発チームは，Webリクエストを処理するPHP環境を対象として挙げていて，これには PHP-FPM, Apache mod_php, php-cgi, FrankenPHP, LiteSpeed などが含まれる。
製品としては WordPress, Drupal, Joomla で実証済みだそうな。
もちろん他にもあるだろう。

修正バージョンは以下の通り。

| PHP系列 | 修正版 |
|---|---:|
| PHP 8.5 | 8.5.12 |
| PHP 8.4 | 8.4.27 |
| PHP 8.3 | 8.3.36 |
| PHP 8.2 | 8.2.35 |

ただし PHP 5.4 以降の全てのバージョンに影響があるらしいので，古いバージョンを使っている場合は注意が必要である。
国内の主要なホスティング事業者の対応状況については [piyolog さんのまとめ](https://piyolog.hatenadiary.jp/entry/2026/10/21/174516#%E5%9B%BD%E5%86%85%E3%81%AE%E4%B8%BB%E3%81%AA%E3%83%9B%E3%82%B9%E3%83%86%E3%82%A3%E3%83%B3%E3%82%B0%E4%BA%8B%E6%A5%AD%E8%80%85%E3%81%AE%E5%AF%BE%E5%BF%9C%E7%8A%B6%E6%B3%81 "PHPの2件の深刻な脆弱性(CVE-2026-109618、CVE-2026-109619)についてまとめてみた - piyolog")が参考になる。

今回の修正版は，元々 2026-10-22 (EST) に公開する予定だったらしい。
日本時間だと23日早々か。
ところが PoC を含む脆弱性情報が[事前に漏れてしまった](https://www.bleepingcomputer.com/news/security/php-warns-admins-of-two-critical-flaws-emergency-patches-coming-tomorrow/ "PHP warns admins of two critical flaws, emergency patches coming tomorrow")らしく，2026-10-21 の 14時 UTC に前倒しになったようだ。
これが日本時間の23時。
深夜残業確定の瞬間である。

テレビも {{% emoji "X" %}} も見ないので知らなかったが，総理大臣自ら特別国会閉会後の記者会見で珍しく言及したらしい。
どうやら公的機関に PHP システムが多くあるようで（実際にかなりの数のサービスがメンテナンスで止まった），必要な措置だったのだろう。
緊急時に（通常パスをすっ飛ばして）トップダウンで指示が出せるのはいいことである。

個人的には公開スケジュールが前倒しになったのが気になる。
対策はきちんと行っておくとして，もうしばらくは注視しておいたほうがよさそうだ。

[Internet Archive の話](https://mastodon.archive.org/@brewsterkahle/117481215431503956 "Brewster Kahle: “Internet Archive ops staff wer…” - Internet Archive")を見ても凄い時代になったなぁ，と思う。
AI が脆弱性を見つけ攻撃し，AI が防御し修正する。
ウロボロスかよ！ これに適応していかないといけないんだぜ。

なんにせよ，現場のエンジニアの方々におかれましてはお疲れ様でした。
ゆっくり寝て下さい。

そうそう。
今回の脆弱性は7月にあった [wp2shell](https://www.ipa.go.jp/security/security-alert/2026/alert20260722.html "WordPressの脆弱性対策について(CVE-2026-60137、CVE-2026-63030：wp2shell) | 情報セキュリティ | IPA 独立行政法人 情報処理推進機構") とは全く別である。
私は少し混同して，最初「また WordPress かよ」と思ってしまった（笑）

## ブックマーク

- [【事前準備のお願い】PHPのセキュリティリリース公開予定について | 情報セキュリティ | IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/security-alert/2026/pre-alert20261021.html)
- [PHP、重大な脆弱性2件を修正 ——認証なしのリモートコード実行に対処。直ちに更新するよう強く推奨（10月22日更新） | gihyo.jp](https://gihyo.jp/article/2026/10/php-critical-security-release)
- [PHPの2件の深刻な脆弱性(CVE-2026-109618、CVE-2026-109619)についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2026/10/21/174516) : よいまとめ
- [PHPの重大な脆弱性「php2shell」（CVE-2026-109618、CVE-2026-109619）の概要と対応指針 - GMO Flatt Security Blog](https://blog.flatt.tech/entry/php2shell)
- [PHP Core RCE Zero Day CVE-2026-109618/109619 Critical Fix](https://phoenix.security/php-core-rce-cve-2026-109618-109619-critical-remote-code-execution/)

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

## 参考

{{< linkcard "51d15914bb38536540b3b670f086f3b05f9b7c13" >}} <!-- 体系的に学ぶ 安全なWebアプリケーションの作り方 第2版 -->
{{< linkcard "f49db55e98f0eb56c864acd2ee6f4da8da10016c" >}} <!-- ハッキング思考 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
