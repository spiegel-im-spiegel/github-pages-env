+++
date = "2016-01-09T18:59:13+09:00"
update = "2016-12-26T21:07:45+09:00"
description = "Go コンパイラのセキュリティアップデートがあるらしい / GnuPG 1.4.20 released / GitLab.com にアカウントを作った / くそな「中間者デバイス」が SHA-1 廃止の邪魔をする"
draft = false
tags = ["security", "vulnerability", "golang", "cryptography", "openpgp", "gnupg", "tools", "git", "gitlab", "firefox", "x509", "hash", "sha-1", "risk"]
title = "週末スペシャル： Go コンパイラのセキュリティアップデートがあるらしい"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"

+++

年末忙しくて書き損ねてるネタを回収中。

1. [Go コンパイラのセキュリティアップデートがあるらしい]({{< relref "#go" >}})
1. [GnuPG 1.4.20 released]({{< relref "#gpg" >}})
1. [GitLab.com にアカウントを作った]({{< relref "#gl" >}})
1. [くそな「中間者デバイス」が SHA-1 廃止の邪魔をする]({{< relref "#mitm" >}})

## Go コンパイラのセキュリティアップデートがあるらしい{#go}

- [[security] Go 1.5.3 pre-announcement - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/MLaPAPFlCNY)

詳細は示されてないが13日（日本時間だと14日かな）を待つことにしよう。
来週は月例のセキュリティ更新週間である。

## GnuPG 1.4.20 released{#gpg}

昨年の話で申し訳ないが

- [[Announce] GnuPG 1.4.20 released](https://lists.gnupg.org/pipermail/gnupg-announce/2015q4/000382.html)

セキュリティアップデートではないが既定の動作が変わるようである。

* Reject signatures made using the MD5 hash algorithm unless the new option --allow-weak-digest-algos or --pgp2 are given.
* New option --weak-digest to specify hash algorithms which should be considered weak.
* Changed default cipher for symmetric-only encryption to AES-128.
* Fix for DoS when importing certain garbled secret keys.
* Improved error reporting for secret subkey w/o corresponding public subkey.
* Improved error reporting in decryption due to wrong algorithm.
* Fix cluttering of stdout with trustdb info in double verbose mode.
* Pass a DBUS envvar to gpg-agent for use by gnome-keyring.

GnuPG というか OpenPGP が後生大事に MD5 を残しているのは過去の資産への対応のためであろう。

ちなみに PGP の最初のリリースは1991年で，実に四半世紀も前である。
作者の [Phil Zimmermann](https://www.philzimmermann.com/) は当時，米国内反核運動の活動家であった。
更に暗号技術に対する政治的圧力は現在と比較にならないほど厳しく，米国は長いあいだ彼をマークし続けた。
PGP の広まり方や改良のプロセスは実に「インターネット的」であった。
詳しくは Steven Levy の『[暗号化（Crypto）](http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/)』をどうぞ。

あれからネットも随分変わったが，当時政府が何をしたか企業は何をしたか「暗号アナーキスト」たちはどうしたか。
現在のネットは先人の努力により「勝ち取った」ものであることを私たちは絶対に忘れてはならない。

自由そのものは自由ではない。
勝ち取ったものはいつか奪われる。
奪われたくなければ勝ち続けなければならない。

## GitLab.com にアカウントを作った{#gl}

- [GitLab/GitLab.com 勉強会 (2015/12/09) レポート - Qiita](http://qiita.com/masakura/items/e679c094e8afea9a4879)

これ見て [GitLab.com] にアカウントを作ってみた。
とりあえず中身は空っぽ。
公開リポジトリを [GitLab.com] に作るメリットはない気もするが，容量が 10GB/repos あるのはありがたい。

まぁ，どう使うかはこれからおいおい考える。
今年あたり，どっかにサーバでも借りて私用リポジトリ・サービスを立ち上げてみたいのだが，これもおいおい。

[GitLab.com]: https://gitlab.com/ "GitLab"

## くそな「中間者デバイス」が SHA-1 廃止の邪魔をする{#mitm}

いやもうこれは笑うところだよね。

- [「Firefox」、SHA-1証明書のサポートを一時的に復活--HTTPSサイトのアクセスに問題 - CNET Japan](http://japan.cnet.com/news/service/35075954/)
- [FirefoxのSHA-1廃止で一部ユーザーに障害、サポート復活 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1601/08/news069.html)

{{< fig-quote title="「Firefox」、SHA-1証明書のサポートを一時的に復活--HTTPSサイトのアクセスに問題" link="http://japan.cnet.com/news/service/35075954/" >}}
<q>Barnes氏は次のように説明している。「ユーザーがHTTPSサイトへの接続を試みると、中間者デバイスがFirefoxに対し、サーバの本物の証明書でなく新規のSHA-1証明書を送信する」<br>
「Firefoxは新規のSHA-1証明書を拒否するため、サーバに接続できない」（同氏）</q>
{{< /fig-quote >}}

「[SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})」でも紹介したが， SHA-1 の危殆化はかなり現実的な問題になってきている。
いますぐどうこうというわけではないが，これ以上の先延ばしは出来ない状態である。
しかし CA やブラウザが頑張っても「中間者デバイス」なる覗き屋が邪魔をする。

やはりセキュリティ製品が暗号通信に対して「中間者攻撃」を仕掛けるのは筋が悪すぎると思うのだが，何とかならないものかねぇ。

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
