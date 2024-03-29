+++
title = "WPA2 脆弱性（KRACKs）に関する覚え書き"
date =  "2017-10-17T20:00:30+09:00"
update = "2018-01-04T20:59:33+09:00"
description = "今回は，世界中で大騒ぎになっている Wi-Fi の WPA2 認証に関する脆弱性について。"
tags = [
  "security",
  "vulnerability",
  "device",
  "wireless",
  "wpa2",
  "authentication",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨夜は早めに寝落ちしてしまったのだが，その間に TL が随分賑やかになっていた。
そこで WPA2 脆弱性の話と中性子星衝突を重力波望遠鏡で観測した話をまとめておく。

今回は，世界中で大騒ぎになっている Wi-Fi の WPA2 認証に関する脆弱性について。
[Bluetooth でやらかした]({{< ref "/remark/2017/09/vulnerability-in-bluetooth-implementation.md" >}} "Bluetooth 実装の脆弱性に関する覚え書き")ばっかりなのに追い打ちですやん。

- [KRACK Attacks: Breaking WPA2](https://www.krackattacks.com/)
    - {{< pdf-file title="Key Reinstallation Attacks: Forcing Nonce Reuse in WPA2" link="https://papers.mathyvanhoef.com/ccs2017.pdf" >}}
- [Vulnerability Note VU#228519 - Wi-Fi Protected Access II (WPA2) handshake traffic can be manipulated to induce nonce and session key reuse](https://www.kb.cert.org/vuls/id/228519)
- [WPA2 における複数の脆弱性について：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20171017_WPA2.html)
- [JVNVU#90609033: Wi-Fi Protected Access II (WPA2) ハンドシェイクにおいて Nonce およびセッション鍵が再利用される問題](http://jvn.jp/vu/JVNVU90609033/)
- [WPA2の脆弱性 KRACKsについてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20171016/1488907259) : オススメ

## 脆弱性の内容

通称 KRACKs (**K**ey **R**einstallation **A**tta**cks**)。

{{< fig-quote link="http://jvn.jp/vu/JVNVU90609033/" title="Wi-Fi Protected Access II (WPA2) ハンドシェイクにおいて Nonce およびセッション鍵が再利用される問題" >}}
<q>WPA2 プロトコルには、ハンドシェイク中に Nonce およびセッション鍵が再利用される問題があります。攻撃者はアクセスポイント (AP) とクライアントの間で Man-in-the-Middle 攻撃を成功させた後、ハンドシェイク中に特定のメッセージを AP またはクライアントに再送することで、Nonce やリプレイカウンタ をリセットし、すでに使用されているセッション鍵を再利用させることが可能です。</q>
{{< /fig-quote >}}

{{< fig-youtube id="Oh4WURZoR98" title="KRACK Attacks: Bypassing WPA2 against Android and Linux - YouTube" >}}

Nonce （ノンスまたはナンスと読むらしい）ってのは暗号通信の最初にやり取りする使い捨てのランダムな値で，これを認証情報に混ぜることで第三者による「リプレイ攻撃（replay attack）」を防ぐことができる。
今回はそういうのが全部チャイされてしまうわけやね。

これが成功すると中間者攻撃（man-in-the-middle attack）が成立してしまい，通信内容の盗み見や改竄ができるようになってしまう。
ただし影響範囲は無線 LAN 内に限定されるため，インターネット越しに攻撃を受けるわけではない。
また暗号通信に使われる暗号アルゴリズムには脆弱性はない。

具体的な脆弱性の内容は以下の通り（[Vulnerability Note VU#228519](http://jvn.jp/vu/JVNVU90609033/) より抜粋）。

- 4-way ハンドシェイクにおける Pairwise Key の再利用 (CVE-2017-13077)
- 4-way ハンドシェイクにおける Group Key の再利用 (CVE-2017-13078)
- 4-way ハンドシェイクにおける Integrity Group Key の再利用 (CVE-2017-13079)
- Group-key ハンドシェイクにおける Group Key の再利用 (CVE-2017-13080)
- Group-key ハンドシェイクにおける Integrity Group Key の再利用 (CVE-2017-13081)
- Fast BSS Transition 再接続リクエストの再送許可とその処理における Pairwise Key の再利用 (CVE-2017-13082)
- PeerKey ハンドシェイクにおける STK Key の再利用 (CVE-2017-13084)
- Tunneled Direct-Link Setup (TDLS) ハンドシェイクにおける TDLS PeerKey (TPK) Key の再利用 (CVE-2017-13086)
- Wireless Network Management (WNM) Sleep Mode レスポンスフレーム処理時の Group Key (GTK) の再利用 (CVE-2017-13087)
- Wireless Network Management (WNM) Sleep Mode レスポンスフレーム処理時の Integrity Ggroup Key (IGTK) の再利用 (CVE-2017-13088)

## 影響度（CVSS）

- [JVNVU#90609033: Wi-Fi Protected Access II (WPA2) ハンドシェイクにおいて Nonce およびセッション鍵が再利用される問題](http://jvn.jp/vu/JVNVU90609033/)

**CVSSv3 基本評価値 5.0 (`CVSS:3.0/AV:A/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | 隣接（A）         |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし（U）     |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）           |
| 情報改ざんの可能性（完全性への影響, I） | 低（L）           |
| 業務停止の可能性（可用性への影響, A）   | 低（L）           |

CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

## 影響を受ける製品

Wi-Fi 通信が可能なあらゆる機器が対象となる。
たとえば無線 LAN ルータ[^rt1]，スマホ等の携帯端末，ネットワーク家電などが考えられる。

[^rt1]: 無線 LAN ルータは Wi-Fi ネットワークを中継する機能のあるものが対象となる。

## 対策・回避策

- 日本の [JVN](http://jvn.jp/vu/JVNVU90609033/)，および [CERT/CC](https://www.kb.cert.org/vuls/id/228519) の各ページにはベンダ企業のステータスが掲載されているので確認すること
    - [Multiple Vulnerabilities in Wi-Fi Protected Access and Wi-Fi Protected Access II](https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20171016-wpa)
    - [WPA2の脆弱性に関する弊社調査・対応状況について | IODATA アイ・オー・データ機器](http://www.iodata.jp/support/information/2017/wpa2/)
    - [大切なお知らせ ｜ BUFFALO バッファロー](http://buffalo.jp/support_s/t20171017.html)
    - [WPA と WPA2 の脆弱性に関する最新情報 – UTM/NGFWでマルウェア・標的型攻撃対策｜ウォッチガード・テクノロジー](https://www.watchguard.co.jp/security-news/wpa-and-wpa2-vulnerabilities-update.html)
    - [「Wi-Fi Protected Access II (WPA2) ハンドシェイクにおいて Nonce およびセッション鍵が再利用される問題」について (FAQ for YAMAHA RT Series / Security)](http://www.rtpro.yamaha.co.jp/RT/FAQ/Security/JVNVU90609033.html)
    - [サポート | D-Link Japan](http://www.dlink-jp.com/support) : 「「WPA2の脆弱性」 に関する弊社調査状況について」をクリック
    - [WPA2の脆弱性に関する弊社調査・対応状況について - 重要なお知らせ - ELECOM |](http://www.elecom.co.jp/support/news/20171018/)
    - [【重要】「WPA2」の脆弱性に関するお知らせ| お知らせ | AtermStation](http://www.aterm.jp/product/atermstation/info/2017/info1018.html)
    - [【お知らせ】Wi-Fiの暗号化技術「WPA2」脆弱性対策について - CNET Japan](https://japan.cnet.com/release/30213798/) : 日本ヒューレット・パッカード株式会社からの報道資料
    - [【重要なお知らせ】無線LAN製品のWPA2の脆弱性に関するお知らせ｜フルノシステムズ](http://www.furunosystems.co.jp/info/detail/id=820)
        - [【重要なお知らせ】無線ハンディターミナル製品のWPA2の脆弱性に関するお知らせ｜フルノシステムズ](http://www.furunosystems.co.jp/info/detail/id=822)
    - [WPA2の脆弱性への対応についてのお知らせ｜PLANEX](http://www.planex.co.jp/news/info/20171019_info.shtml)
    - [無線LAN搭載 SDメモリカード/FlashAir™における「WPA2の鍵情報の生成・管理に関する脆弱性」について｜ニュース｜会社情報｜東芝メモリ](https://www.toshiba-memory.co.jp/company/news/20171017-1.html)
    - [サポート｜Wi-Fi Protected Access II (WPA2) ハンドシェイクに関する脆弱性](https://www.allied-telesis.co.jp/support/list/faq/vuls/20171019.html) : アライドテレシス株式会社
    - [勧告: Sophos Wireless への WPA および WPA2 の脆弱性による鍵再インストール攻撃 (KRACKs) の影響 - Sophos Community](https://community.sophos.com/kb/ja-jp/127658)
    - [WPA2 セキュリティの脆弱性に関して(KRACKs) - TP-Link](http://www.tp-link.jp/faq-1970.html)
    - [WPA2の脆弱性に関するBHT製品への影響と対応｜お知らせ｜デンソーウェーブ](https://www.denso-wave.com/ja/info/detail__991.html)
- Android については2017年11月の修正で対応しているようだ
    - [Android Security Bulletin—November 2017  |  Android Open Source Project](https://source.android.com/security/bulletin/2017-11-01)
    - ただし端末を提供しているベンダ企業やキャリア企業が対応しない限り対応されない。特に古い機種や古いバージョンの OS を使い続けている場合は対応されないと思ったほうがいいだろう
- Apple 製品については修正版がリリースされている
    - [watchOS 4.1 のセキュリティコンテンツについて - Apple サポート](https://support.apple.com/ja-jp/HT208220)
    - [About the security content of iOS 11.1 - Apple サポート](https://support.apple.com/ja-jp/HT208222)
    - [About the security content of macOS High Sierra 10.13.1, Security Update 2017-001 Sierra, and Security Update 2017-004 El Capitan - Apple サポート](https://support.apple.com/ja-jp/HT208221)
    - [tvOS 11.1 のセキュリティコンテンツについて - Apple サポート](https://support.apple.com/ja-jp/HT208219)
    - [About the security content of iOS 11.2 - Apple サポート](https://support.apple.com/ja-jp/HT208334)
        - [Apple、旧iPhone/iPadのWi-Fi脆弱性を「iOS 11.2」で修正  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1095897.html)
    - [About the security content of AirPort Base Station Firmware Update 7.6.9 - Apple サポート](https://support.apple.com/ja-jp/HT208258)
    - [About the security content of AirPort Base Station Firmware Update 7.7.9 - Apple サポート](https://support.apple.com/ja-jp/HT208354)
- Windows については Microsoft からの10月のアップデートで修正されている
    - [CVE-2017-13080 | Windows Wireless WPA Group Key Reinstallation Vulnerability](https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/CVE-2017-13080)
- Linux や FreeBSD 等については対応が始まっている。ディストリビュータの情報に注意すること
    - [WPA2の脆弱性(KRACK Attacks / KRACKs )とCVE情報(CVE-2017-13077 - CVE-2017-13088) — | サイオスOSS | サイオステクノロジー](https://oss.sios.com/security/wpa-security-vulnerability-20171016)

なお GitHub にて CVE-2017-13082（Fast BSS Transition 再接続リクエストの再送許可とその処理における Pairwise Key の再利用）についてのチェックツールが公開されている模様（CVE-2017-13082 は 802.11r 方式を使うアクセスポイントに影響する）

- [vanhoefm/krackattacks-test-ap-ft](https://github.com/vanhoefm/krackattacks-test-ap-ft)

サポートが受けられない場合，回避方法としては以下が挙げられるだろう。

- Wi-Fi を使用しない
    - 公衆空間で Wi-Fi を使わない場合は無効にしておく
    - 公衆無線 LAN はリスクが高いので利用しない
    - 可能であれば有線で接続する
- 暗号化通信を利用する
    - Web でのやりとりには HTTPS を使う（[HSTS (HTTP Strict Transport Security)](https://developer.mozilla.org/ja/docs/Web/Security/HTTP_Strict_Transport_Security "HTTP Strict Transport Security - Web セキュリティ | MDN") でちゃんと HTTPS にリダイレクトされること）
    - VPN サービスを利用する（慌ててよく分からないサービスに飛びつかないこと）

WPA2 がダメだからと言って WEP を使うのは事態を悪化させるだけである。
脆弱性が発見された現時点でも WPA2 が一番まともな認証プロトコルであることには変わりない。

はっきりいってネットワーク機器のアップデートはあまり期待できない。
とくに安ものの無線 LAN ルータとか，古い機種のスマホとかは事実上の放置プレイである。
先月の [Bluetooth 脆弱性]({{< ref "/remark/2017/09/vulnerability-in-bluetooth-implementation.md" >}} "Bluetooth 実装の脆弱性に関する覚え書き")に対応しなかったベンダ企業やその製品については，今回も何もないと考えたほうがいい。

はっきり言って Android 端末はもう潮時かなと思ってる。
他の選択肢が Apple 製品しかないってのが業腹だけど（個人的に嫌いなんだよ）。

まぁ，この機会にゆっくり考えることにしよう。

## ブックマーク

- [Serious flaw in WPA2 protocol lets attackers intercept passwords and much more | Ars Technica](https://arstechnica.com/information-technology/2017/10/severe-flaw-in-wpa2-protocol-leaves-wi-fi-traffic-open-to-eavesdropping/)
- [New KRACK Attack Against Wi-Fi Encryption - Schneier on Security](https://www.schneier.com/blog/archives/2017/10/new_krack_attac.html)

- [無線LAN（Wi-Fi）暗号化における脆弱性への対応について  |  無線LANビジネス推進連絡会【WiBiz（ワイビズ）】 - お知らせ - 新着情報無線LANビジネス推進連絡会【WiBiz（ワイビズ）】](http://www.wlan-business.org/archives/11325)
- [セキュリティホール memo の情報](https://www.st.ryukoku.ac.jp/~kjm/security/memo/2017/10.html#20171017_wpa2)
- [WPA2の脆弱性「KRACKs」公開、多数のWi-Fi機器に影響の恐れ - CNET Japan](https://japan.cnet.com/article/35108859/)
- [WPA2のWiFi脆弱性から身を守る方法――KRACK攻撃の内容と対策 | TechCrunch Japan](https://techcrunch.com/2017/10/16/heres-what-you-can-do-to-protect-yourself-from-the-krack-wifi-vulnerability/) : 山奥に引っ越せとか周囲の土地を買い取って建物を潰せとか，ネタでも笑えないよ
- [Wi-Fiを脅かす脆弱性「KRACK」、各社の対応状況は--MS、アップル、グーグルなど - ZDNet Japan](https://japan.zdnet.com/article/35108863/)
- [WPA2の脆弱性「KRACK」対処パッチ、Microsoftは対応済み、AppleのOSとAndroidは数週間中 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1710/17/news044.html)
    - [「WPA2」の脆弱性情報、セキュリティ機関が公開　パッチ適用を呼び掛け - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1710/17/news048.html)
- [WPA2の脆弱性「KRACKs」、ほぼすべてのWi-Fi通信可能な端末機器に影響 | トレンドマイクロ セキュリティブログ](http://blog.trendmicro.co.jp/archives/16162)
- [WiFi Connection Vulnerability - Blog - ProtonVPN -](https://protonvpn.com/blog/wifi-vulnerability-krack/) : ProtonVPN の広告記事（笑）
- [総務省｜無線LAN（Wi-Fi）暗号化における脆弱性について（注意喚起）](http://www.soumu.go.jp/menu_kyotsuu/important/kinkyu02_000274.html) : いまさら総務省（笑）
- [Apple、WPA2暗号化の脆弱性を修正した「iOS 11.1」、ただし機種限定 - PC Watch](https://pc.watch.impress.co.jp/docs/news/1089340.html)
- [Androidの月例セキュリティ情報公開、「KRACK」の脆弱性に対処 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1711/07/news067.html)
- [Google、Wi-Fi関連の脆弱性「KRACK」に対処するAndroid向けパッチを公開 | マイナビニュース](http://news.mynavi.jp/news/2017/11/10/068/)
- [LineageOSではKRACKバグの修正は完了、Android 8.1ベースのLineageOS 15.1の開発に着手 ｜ ガジェット通信 GetNews](http://getnews.jp/archives/1992799)
