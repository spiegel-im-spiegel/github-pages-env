+++
date = "2016-07-25T21:37:51+09:00"
description = "Go 1.6.3 セキュリティ・アップデート / GnuPG 2.1.14 もリリース / 「オーロラの調べ」観に行きました / Pokémon GO しばらく休眠します / その他の気になる記事"
draft = false
tags = ["security", "vulnerability", "golang", "cryptography", "openpgp", "tools", "gnupg", "astronomy", "games", "pokemon"]
title = "週明けから戯れ言： Go 1.6.3 セキュリティ・アップデート，他"

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

久しぶりに低血圧ですよ。
上が95切ると正直立ち上がるのもしんどい状態。
明日から本気出します。

で，週が明けてしまったので「週明けから戯れ言」。

1. [Go 1.6.3 セキュリティ・アップデート]({{< relref "#go" >}})
1. [GnuPG 2.1.14 もリリース]({{< relref "#gpg" >}})
1. [「オーロラの調べ」観に行きました]({{< relref "#kgy" >}})
1. [Pokémon GO しばらく休眠します]({{< relref "#pg" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## Go 1.6.3 セキュリティ・アップデート{#go}

- [[security] Go 1.6.3 and Go 1.7rc2 pre-announcement - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/7JTsd70ZAT0)
- [[security] Go 1.6.3 and 1.7rc2 are released - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/7jZDOQ8f8tM)
- [Go 1.7 Release Candidate 3 is released - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/G5N8lCAspoU)

例の “[httpoxy](https://httpoxy.org/)” に絡むセキュリティ・アップデートです。

- [Vulnerability Note VU#797896 - CGI web servers assign Proxy header values from client requests to internal HTTP_PROXY environment variables](https://www.kb.cert.org/vuls/id/797896)
- [CGI 等を利用する Web サーバの脆弱性 (CVE-2016-5385 等) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160031.html)

まぁ CVSSv2 基本値で 5.1 くらいなので「要注意」程度ですが，計画的にアップデートを行いましょう。
焦る必要はありません。

## GnuPG 2.1.14 もリリース{#gpg}

- [[Announce] Libgcrypt 1.7.2 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q3/000392.html)
- [[Announce] GnuPG 2.1.14 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q3/000393.html)
- [[Announce] OpenPGP.conf program published](https://lists.gnupg.org/pipermail/gnupg-announce/2016q3/000391.html)

機能改善やバグ修正がメイン。
以下に GnuPG の変更点のみ示す。

* gpg: Removed options `--print-dane-records` and `--print-pka-records`. The new export options "`export-pka`" and "`export-dane`" can instead be used with the export command.
* gpg: New options `--import-filter` and `--export-filter`.
* gpg: New import options "`import-show`" and "`import-export`".
* gpg: New option `--no-keyring`.
* gpg: New command `--quick-revuid`.
* gpg: New options `-f`/`--recipient-file` and `-F`/`--hidden-recipient-file` to directly specify encryption keys.
* gpg: New option `--mimemode` to indicate that the content is a MIME part.  Does only enable `--textmode` right now.
* gpg: New option `--rfc4880bis` to allow experiments with proposed changes to the current OpenPGP specs.
* gpg: Fix regression in the "`fetch`" sub-command of `--card-edit`.
* gpg: Fix regression since 2.1 in option `--try-all-secrets`.
* gpgv: Change default options for extra security.
* gpgsm: No more root certificates are installed by default.
* agent: "`updatestartuptty`" does now affect more environment variables.
* scd: The option `--homedir` does now work with scdaemon.
* scd: Support some more GEMPlus card readers.
* gpgtar: Fix handling of '`-`' as file name.
* gpgtar: New commands `--create` and `--extract`.
* gpgconf: Tweak for `--list-dirs` to better support shell scripts.
* tools: Add programs gpg-wks-client and gpg-wks-server to implement a Web Key Service.  The configure option `--enable-wks-tools` is required to build them; they should be considered Beta software.
* tests: Complete rework of the openpgp part of the test suite.  The test scripts have been changed from Bourne shell scripts to Scheme programs.  A customized scheme interpreter (gpgscm) is included. This change was triggered by the need to run the test suite on non-Unix platforms.
* The rendering of the man pages has been improved.

`--rfc4880bis` オプションができてるな。
まぁ普通は使わないけど。
次期 OpenPGP も実装フェーズに入ったってことかな。

## 「オーロラの調べ」観に行きました{#kgy}

観に行きましたよ！

- [プラネタリウム番組「ネイチャーリウム オーロラの調べ 神秘の光を探る」：広島市こども文化科学館](http://www.pyonta.city.hiroshima.jp/event/detail/id/2904.html)

{{< fig-youtube id="zQUelppNwq8" width="500" height="281" title="全天周映像『オーロラの調べ』予告編 - YouTube" >}}

前から [KAGAYA さん](http://www.kagayastudio.com/)の映像は観たかったのだが，本当に素晴らしい。
[ORIGA](http://www.r-s.co.jp/origa/) さん相変わらず素敵。
ただ映像が綺麗なだけでなく，ちゃんと科学的な解説もしてくれる。
いやぁ堪能しました。

[広島市こども文化科学館](http://www.pyonta.city.hiroshima.jp/)で観るときは投影機より少し斜め後ろの席に座るのがオススメ。
あと[入場チケットの販売が1階カウンターに変わった](http://www.pyonta.city.hiroshima.jp/event/detailnews/id/187.html)のでご注意を。

## Pokémon GO しばらく休眠します{#pg}

いや，始めたばっかりなのだが， [Pokémon GO Plus] が出るまでしばらく休眠します。

[Pokémon GO] が [Ingress] と一番異なるのはポケモンがそこら中に湧いて出ることである。
[Ingress] なら Portal の近所でアプリを起動すればいいのだが， [Pokémon GO] ではどこでポケモンに出くわすか分からないのでアプリを立ち上げっぱなしにしないといけない。
スリープ状態でもダメ。
これって「歩きスマホ」を助長してしまうので割と危ないのだ。

ちゃんと時間を作ってプレイできればいいのだが，現在仕事が忙しすぎてそれもままならない。
しかも通勤時間帯は私にとってネットからの情報摂取の時間なので [Pokémon GO] をやる暇がないのだ。

[Pokémon GO Plus] が出ればこの状況が多少は緩和されると信じて今は休眠します。
他の人と比べて出遅れる形になるけどしょうがない。

## その他の気になる記事{#other}

- [トルコのクーデタ未遂をめぐる８つの考察ポイント：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41381)
- [テロの時代の論理と倫理 – 中東・イスラーム学の風姿花伝](http://ikeuchisatoshi.com/%E3%83%86%E3%83%AD%E3%81%AE%E6%99%82%E4%BB%A3%E3%81%AE%E8%AB%96%E7%90%86%E3%81%A8%E5%80%AB%E7%90%86/)
    - [日本でテロ事件の議論が見当はずれになる背景 – アゴラ](http://agora-web.jp/archives/2020343.html)
- [JAXA | X線天文衛星ASTRO-H「ひとみ」の後継機の検討について](http://www.jaxa.jp/press/2016/07/20160714_hitomi_j.html)
- [Amazon.co.jp： なぜ核はなくならないのかII: 「核なき世界」への視座と展望: 吉川 元, 水本 和実, 佐渡 紀子, 福井 康人, 広瀬 訓, 倉科 一希, 茅原 郁生, 吉村 慎太郎, 孫 賢鎮, ロバート・ジェイコブズ, 国末 憲人, 中村 桂子, 広島市立大学 広島平和研究所: 本](http://www.amazon.co.jp/exec/obidos/ASIN/4589037858/baldandersinf-22/)
- [Adobe Reader および Acrobat の脆弱性 (APSB16-26) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160030.html)
    - [Adobe Reader および Acrobat の脆弱性対策について(APSB16-26)(CVE-2016-4210等)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160713-adobereader.html)
- [Adobe Flash Player の脆弱性 (APSB16-25) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160029.html)
    - [Adobe Flash Player の脆弱性対策について(APSB16-25)(CVE-2016-4247等)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160713-adobeflashplayer.html)
- [2016 年 7 月のセキュリティ情報 (月例) – MS16-084 ～ MS16-094 – 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/07/13/201607-security-bulletin/)
    - [2016年 7月 Microsoft セキュリティ情報 (緊急 6件含) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160028.html)
    - [Microsoft 製品の脆弱性対策について(2016年7月)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160713-ms.html)
- [JVNDB-2016-003523 Nexus 7 (2013) デバイス上で稼動する Android の Qualcomm コンポーネントの drivers/video/fbcmap.c におけるバッファオーバーフローの脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003523.html)
- [JVNDB-2016-003526 Nexus 5 および 7 (2013) デバイス上で稼動する Android の Qualcomm コンポーネントにおけるバッファオーバーフローの脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003526.html)
- [JVNDB-2016-003520 Nexus 5 および 7 (2013) デバイス上で稼動する Android の Qualcomm コンポーネントにおける権限を取得される脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003520.html)
- [JVNVU#96627087: libbpg にメモリ境界外への書き込みを行う脆弱性](http://jvn.jp/vu/JVNVU96627087/)
- [2016年7月 Oracle Java SE のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2016/at160032.html)
    - [Oracle Java の脆弱性対策について(CVE-2016-3587等)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160720-jre.html)

[Pokémon GO]: http://www.pokemongo.jp/ "『Pokémon GO』公式サイト"
[Pokémon GO Plus]: http://www.pokemongo.jp/plus/ "Pokémon GO Plus｜『Pokémon GO』公式サイト"
[Ingress]: https://www.ingress.com/