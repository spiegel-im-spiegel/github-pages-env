+++
title = "WPA2 脆弱性（KRACKs）に関する覚え書き"
date =  "2017-10-17T20:00:30+09:00"
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
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨夜は早めに寝落ちしてしまったのだが，その間に TL が随分賑やかになっていた。
そこで WPA2 脆弱性の話と中性子星衝突を重力は望遠鏡で観測した話をまとめておく。

今回は，世界中で大騒ぎになっている Wi-Fi の WPA2 認証に関する脆弱性について。
[Bluetooth でやらかした]({{< relref "remark/2017/09/vulnerability-in-bluetooth-implementation.md" >}} "Bluetooth 実装の脆弱性に関する覚え書き")ばっかりなのに追い打ちですやん。

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

Nonce （ノンスまたはナンスと読むらしい）ってのは暗号通信の最初にやり取りする使い捨てのランダムな値で，これを認証情報に混ぜることで第三者による「リプレイ攻撃」を防ぐことができる。
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

**CVSSv3 基本評価値 5.0 (` CVSS:3.0/AV:A/AC:H/PR:N/UI:N/S:U/C:L/I:L/A:L`)**

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | 関節（A）         |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし（U）     |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）           |
| 情報改ざんの可能性（完全性への影響, I） | 低（L）           |
| 業務停止の可能性（可用性への影響, A）   | 低（L）           |

CVSS については[解説ページ]({{< relref "remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

## 影響を受ける製品

Wi-Fi 通信が可能なあらゆる機器が対象となる。
たとえば無線ルータ，スマホ等の携帯端末，ネットワーク家電などが考えられる。


## 対策・回避策

- 日本の [JVN](http://jvn.jp/vu/JVNVU90609033/)，および [CERT/CC](https://www.kb.cert.org/vuls/id/228519) の各ページにはベンダ企業のステータスが掲載されているので確認すること
    - [Multiple Vulnerabilities in Wi-Fi Protected Access and Wi-Fi Protected Access II](https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20171016-wpa)
    - [WPA2の脆弱性に関する弊社調査・対応状況について | IODATA アイ・オー・データ機器](http://www.iodata.jp/support/information/2017/wpa2/)
- iOS や Android を提供する Apple や Google では修正版を準備中とのこと
    - ただし Android に関しては，端末を提供しているベンダ企業やキャリア企業が対応しない限り対応されないと思われる。特に古い機種や古いバージョンの OS を使い続けている場合は対応されないと思ったほうがいい
- Windows については Microsoft からの10月のアップデートで修正されている
    - [CVE-2017-13080 | Windows Wireless WPA Group Key Reinstallation Vulnerability](https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/CVE-2017-13080)
- Linux や FreeBSD 等については対応が始まっている。ディストリビュータの情報に注意すること
    - [WPA2の脆弱性(KRACK Attacks / KRACKs )とCVE情報(CVE-2017-13077 - CVE-2017-13088) — | サイオスOSS | サイオステクノロジー](https://oss.sios.com/security/wpa-security-vulnerability-20171016)

サポートが受けられない場合，回避方法としては以下が挙げられるだろう。

- Wi-Fi を使用しない
    - 公衆空間で Wi-Fi を使わない場合は無効にしておく
    - 公衆無線 LAN はリスクが高いので利用しない
    - 可能であれば有線で接続する
- 暗号化通信を利用する
    - Web でのやりとりには HTTPS を使う（[HSTS (HTTP Strict Transport Security)](https://developer.mozilla.org/ja/docs/Web/Security/HTTP_Strict_Transport_Security "HTTP Strict Transport Security - Web セキュリティ | MDN") でちゃんと HTTPS にリダイレクトされること）
    - VPN サービスを利用する（慌ててよく分からないサービスに飛びつかないこと）

WPA2 がダメだからと言って WEP を使うのは事態を悪化させるだけなので使わないこと。

はっきりいってネットワーク機器のアップデートはあまり期待できない。
とくに安ものの無線ルータとか，古い機種のスマホとかは事実上の放置プレイである。
先月の [Bluetooth 脆弱性]({{< relref "remark/2017/09/vulnerability-in-bluetooth-implementation.md" >}} "Bluetooth 実装の脆弱性に関する覚え書き")に対応しなかったベンダ企業やその製品については，今回も何もないと考えたほうがいい。

はっきり言って Android 端末はもう潮時かなと思ってる。
他の選択肢が Apple 製品しかないってのが業腹だけど（個人的に嫌いなんだよ）。

まぁ，この機会にゆっくり考えることにしよう。

## ブックマーク

- [Serious flaw in WPA2 protocol lets attackers intercept passwords and much more | Ars Technica](https://arstechnica.com/information-technology/2017/10/severe-flaw-in-wpa2-protocol-leaves-wi-fi-traffic-open-to-eavesdropping/)
- [New KRACK Attack Against Wi-Fi Encryption - Schneier on Security](https://www.schneier.com/blog/archives/2017/10/new_krack_attac.html)

- [WPA2の脆弱性「KRACKs」公開、多数のWi-Fi機器に影響の恐れ - CNET Japan](https://japan.cnet.com/article/35108859/)
- [WPA2のWiFi脆弱性から身を守る方法――KRACK攻撃の内容と対策 | TechCrunch Japan](http://jp.techcrunch.com/2017/10/17/20171016heres-what-you-can-do-to-protect-yourself-from-the-krack-wifi-vulnerability/) : 山奥に引っ越せとか周囲の土地を買い取って建物を潰せとか，ネタでも笑えないよ
- [Wi-Fiを脅かす脆弱性「KRACK」、各社の対応状況は--MS、アップル、グーグルなど - ZDNet Japan](https://japan.zdnet.com/article/35108863/)
- [WPA2の脆弱性「KRACK」対処パッチ、Microsoftは対応済み、AppleのOSとAndroidは数週間中 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1710/17/news044.html)
    - [「WPA2」の脆弱性情報、セキュリティ機関が公開　パッチ適用を呼び掛け - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1710/17/news048.html)

<!-- eof -->
