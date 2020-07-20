+++
title = "「門番」としての VPN"
date =  "2020-07-20T18:28:17+09:00"
description = "専用 IP アドレスによる利便性を得る引き換えとして匿名性をある程度手放さざるを得ないわけやね。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "vpn", "tracking", "surveillance-capitalism" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## VPN とは

まずは VPN についておさらい。

VPN（Virtual Private Network）とは，物理的な回線・ネットワークの上に構築された「{{% ruby "virtual Network" %}}実質的な網{{% /ruby %}}」である。
この記事では VPN の機能として大まかに

- カプセル化とトンネリング
- 経路の暗号化
- 認証とアクセス制御

を挙げ，これらの要件を満たした製品・サービスを VPN と呼ぶことにする。

VPN は目的別に以下の3つに分類できるだろう。

1. 拠点間 VPN
2. リモートアクセス
3. バイパス型 VPN

「拠点間 VPN」は昔でいうところの「専用線」の代わりになるものだ。
「リモートアクセス」は企業イントラネットやプライベート・ネットワーク内部に外からアクセスするもので，近年では BYOD (Bring Your Own Device) とも呼ばれている。

「バイパス型 VPN」は[6年前に私が勝手につけた名前](https://baldanders.info/blog/000773/ "VPN に関する基礎学習")だが，他の2つと異なり，アクセスを行う利用者から見て，出口ノードがインターネットになっている点が特徴である。

{{< fig-img src="./vpn.png" title="「IPAテクニカルウォッチ「公衆無線LAN利用に係る脅威と対策」」より" link="https://www.ipa.go.jp/security/technicalwatch/201600330.html" width="820" >}}

たとえば，上の図のように，公衆無線 LAN などの信用度の低いネットワークを経由してインターネットに接続したい場合に「バイパス型 VPN」が使える。
また異なるリージョンのコンテンツにアクセスするために使われることもある[^vpn1]。

[^vpn1]: つか，元々の「バイパス型 VPN」の動機は他リージョンのコンテンツにアクセスするためなのだが（笑）

## 「門番」としての VPN

先程の図を見ると分かるように「バイパス型 VPN」は利用者とインターネット上のコンテンツとを仕切る「門番」として機能していることが分かる。
したがって「バイパス型 VPN」を運営するサービス・プロバイダが「門番」として信用できるか否かがとても重要である。

### ログ収集ポリシー

技術的な観点はひとまず置いておいて，「バイパス型 VPN」サービス・プロバイダの信用度を計る目安としてよく引き合いに出されるのが「ログ収集ポリシー」である。

たとえばサービス・プロバイダが収集した利用者のアクセスログを第三者に売ったり，あるいは公的機関から提出を求められる（大抵は拒否できない）場合がある。
なので，サービス・プロバイダが利用者に関するどのような情報を持っていてそれらをどのようにして管理しているか，を知ることが大事である。
そしてその最善は「ログ収集しない」ことである。

[TorrentFreak] では以下の質問を「バイパス型 VPN」のサービス・プロバイダ各社に送って

{{< fig-quote type="markdown" title="Which VPN Providers Really Take Anonymity Seriously in 2020?" link="https://torrentfreak.com/best-vpn-anonymous-no-logging/" lang="en" >}}

1. Do you keep (or share with third parties) ANY data that would allow you to match an IP-address and a timestamp to a current or former user of your service? If so, exactly what information do you hold/share and for how long?
2. What is the name under which your company is incorporated (+ parent companies, if applicable) and under which jurisdiction does your company operate?
3. What tools are used to monitor and mitigate abuse of your service, including limits on concurrent connections if these are enforced?
4. Do you use any external email providers (e.g. Google Apps), analytics, or support tools ( e.g Live support, Zendesk) that hold information provided by users?
5. In the event you receive a DMCA takedown notice or a non-US equivalent, how are these handled?
6. What steps would be taken in the event a court orders your company to identify an active or former user of your service? How would your company respond to a court order that requires you to log activity for a user going forward? Have these scenarios ever played out in the past?
7. Is BitTorrent and other file-sharing traffic allowed on all servers? If not, why? Do you provide port forwarding services? Are any ports blocked?
8. Which payment systems/providers do you use? Do you take any measures to ensure that payment details can’t be linked to account usage or IP-assignments?
9. What is the most secure VPN connection and encryption algorithm you would recommend to your users?
10. Do you provide tools such as “kill switches” if a connection drops and DNS/IPv6 leak protection? Do you support Dual Stack IPv4/IPv6 functionality?
11. Are any of your VPN servers hosted by third parties? If so, what measures do you take to prevent those partners from snooping on any inbound and/or outbound traffic? Do you use your own DNS servers?
12. In which countries are your servers physically located? Do you offer virtual locations?

{{< /fig-quote >}}

その結果を公開している。

- [Which VPN Providers Really Take Anonymity Seriously in 2020? * TorrentFreak](https://torrentfreak.com/best-vpn-anonymous-no-logging/)

まぁ，結果はリンク先を見ていただくとして，実はこの話には続きがある。

- [Most Dedicated VPN IP-addresses Are Not Anonymous * TorrentFreak](https://torrentfreak.com/most-dedicated-vpn-ip-addresses-are-not-anonymous-200719/)

「バイパス型 VPN」で「ログ収集しない」と謳うサービス・プロバイダは，匿名性を高めるために，接続ごとに共有 IP アドレスを割り当てるのだが，どうも拡張サービスとして専用あるいは固定の IP アドレスを付与することができるそうで，この場合は（当たり前だが）完全に「収集しない」とは行かないらしい。

{{< fig-quote type="markdown" title="Most Dedicated VPN IP-addresses Are Not Anonymous" link="https://torrentfreak.com/most-dedicated-vpn-ip-addresses-are-not-anonymous-200719/" lang="en" >}}
Broadly speaking, we would say that the “no logs” policies of VPN providers don’t apply to dedicated IPs. That conclusion is backed up by several VPN providers we reached out to, which include VPNArea, NordVPN, CyberGhost, and Torguard.

These providers all have a no-logging policy for their regular VPN service, which relies on shared IP-addresses. However, they see dedicated IP-addresses as a separate and different service, which is treated differently anonymity-wise.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Most Dedicated VPN IP-addresses Are Not Anonymous" link="https://torrentfreak.com/most-dedicated-vpn-ip-addresses-are-not-anonymous-200719/" lang="en" >}}
By connecting through a single IP-address, monitoring outfits can build up a profile of the user’s online activity. The real anonymity tradeoff, however, is that the VPN provider knows the user’s IP-address and can connect it to other account information it has on record. This sometimes includes an email address.
{{< /fig-quote >}}

いくつかのインターネット・サービスでは，リージョンを跨いでアクセスされないよう VPN 接続のブラックリストを持っているそうだが，専用 IP アドレスのオプションを利用することによりこれらを排除できるそうだ。

{{< fig-quote type="markdown" title="Most Dedicated VPN IP-addresses Are Not Anonymous" link="https://torrentfreak.com/most-dedicated-vpn-ip-addresses-are-not-anonymous-200719/" lang="en" >}}
With a dedicated IP-address, which is often sold as an add-on, users get a unique IP-address as opposed to a shared one. This can be very convenient as it reduces annoying captchas and can bypass regular VPN blacklists.
{{< /fig-quote >}}

こうした利便性を得る{{% ruby "tradeoff" %}}引き換え{{% /ruby %}}として匿名性をある程度手放さざるを得ないわけやね。

### F-Secure FREEDOME は止めとけ？

ところで，先程の {{% quote lang="en" %}}[Which VPN Providers Really Take Anonymity Seriously in 2020?](https://torrentfreak.com/most-dedicated-vpn-ip-addresses-are-not-anonymous-200719/){{% /quote %}} の中に私が Android 端末で愛用している F-Secure 社の [FREEDOME] が見当たらなかったので軽くググってみたが，どうも最近は [FREEDOME] は忌避されているらしい。

- [Read This F-Secure Freedome Review and Test Before You Buy It!](https://the-bestvpn.com/f-secure-freedome-review-test/)

[FREEDOME] はログ収集を行ってると言われていて，そこが嫌われている原因のひとつかも（笑）

これについては，一応 F-Secure 社側の釈明もあるようで

{{< fig-quote type="markdown" title="VPNサービス:プロバイダのログ収集以外に考えるべきこと" link="https://blog.f-secure.com/ja/vpn-logging/" >}}
{{% quote %}}ほとんどのVPNと同様に、FREEDOMEは接続ログを作成します。これらのログは、サービスの提供と改善のために使用されます。当社のプライバシーポリシーに記載されているように、サービスを提供し、データ転送をクリーンな状態に保つだけの目的で通信トラフィックを分析します。トラフィックは仮名化されていますので、私たちにはどれがあなたのトラフィックかを知ることはできません。**FREEDOMEはトラフィックログを作成しません**{{% /quote %}}。
{{< /fig-quote >}}

というわけで，接続ログはあるけどトラフィックログは作らないので，ログから（閲覧・購入履歴，メッセージなどの）センシティブ・データを抜いたりできないよー，ということらしい。

とは言え，昔に比べて [FREEDOME] の優位性が薄れているのも確かなんだよなぁ。
まぁ，来年3月まで [FREEDOME] の契約が残ってるので，それまでに [NordVPN] とか他のサービスに乗り換えるべきかじっくり考えてみるとしよう。

## 【おまけ】 NSA による VPN セキュリティ管理の5箇条

米国 NSA から VPN セキュリティに関する以下のドキュメントが公開されている。
どちらかと言うと組織内のネットワーク管理者あるいはサービス・プロバイダ側の話かな。

- {{< pdf-file title="ConfiguringIPsecVirtual Private Networ" link="https://media.defense.gov/2020/Jul/02/2002355501/-1/-1/0/CONFIGURING_IPSEC_VIRTUAL_PRIVATE_NETWORKS_2020_07_01_FINAL_RELEASE.PDF" >}}
- {{< pdf-file title="ConfiguringIPsecVirtual Private Networ" link="https://media.defense.gov/2020/Jul/02/2002355625/-1/-1/0/SECURING_IPSEC_VIRTUAL_PRIVATE_NETWORKS_EXECUTIVE_SUMMARY_2020_07_01_FINAL_RELEASE.PDF" >}} （要約）

これによると，ネットワーク管理者は VPN の運営について以下の5つを定期的に確認・実行する必要がある，と謳っている。

{{< fig-quote type="markdown" title="ConfiguringIPsecVirtual Private Networ" link="https://media.defense.gov/2020/Jul/02/2002355501/-1/-1/0/CONFIGURING_IPSEC_VIRTUAL_PRIVATE_NETWORKS_2020_07_01_FINAL_RELEASE.PDF" lang="en" >}}
- Reduce the VPN gateway attack surface
- Verify that cryptographic algorithms are Committee on National Security Systems Policy (CNSSP) 15-compliant
- Avoid using default VPN settings
- Remove unused or non-compliant cryptography suites
- Apply vendor-provided updates (i.e. patches) for VPN gateways and client
{{< /fig-quote >}}

詳しい内容はドキュメントを参照のこと。

まぁ，当たり前の話なんだけど {{% quote lang="en" %}}Avoid using default VPN settings{{% /quote %}} はちょっと面白かったので。
そんなに既定の設定はあかんのか？

## ブックマーク

- [Best VPN Services For Anonymous Torrenting? * TorrentFreak](https://torrentfreak.com/anonymous-vpn-service-provider-review-torrent/)
- [Mozilla’s VPN launches out of beta on Windows and Android - The Verge](https://www.theverge.com/platform/amp/2020/7/15/21325316/mozilla-vpn-android-windows-launch-firefox-private-network-price)
- [NSA on Securing VPNs - Schneier on Security](https://www.schneier.com/blog/archives/2020/07/nsa_on_securing.html)

- [テレワークを行う際のセキュリティ上の注意事項：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/announce/telework.html)
- [監視をコントロールする](https://baldanders.info/blog/000490/)

[TorrentFreak]: https://torrentfreak.com/
[FREEDOME]: https://www.f-secure.com/en/home/products/freedome "F-Secure FREEDOME VPN — Protect your privacy | F-Secure"
[NordVPN]: https://nordvpn.com/ "Best VPN service. Online security starts with a click. | NordVPN"

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
