+++
date = "2017-03-25T13:57:44+09:00"
update = "2017-04-05T13:02:53+09:00"
title = "「隠すことによるセキュリティ」が何をもたらすか"
draft = false
tags = ["security", "risk", "vulnerability", "intelligence", "leak"]
description = "よく「隠すセキュリティはダメ」と言われるが，目的はどうあれ脆弱性情報を秘匿することがどういう結果をもたらすか身に染みたのではないだろうか。"

[author]
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  facebook = "spiegel.im.spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  tumblr = ""
  twitter = "spiegel_2007"
  license = "by-sa"
  avatar = "/images/avatar.jpg"
  flattr = ""
  github = "spiegel-im-spiegel"
  url = "https://baldanders.info/spiegel/profile/"
+++

例の [Vault 7 絡み]({{< ref "/remark/2017/03/cia-hacking-tools-from-wikileaks.md" >}} "WikiLeaks がリークした CIA ハッキングツール")の話。

- [ニュース解説 - CIAの機密文書で発覚、シスコ製品300種類にパッチ提供未定の危険な脆弱性：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/032300899/?rt=nocnt)

Cisco は相変わらず「遅い」なぁ。
まぁほかのネットワーク機器を販売している企業も似たり寄ったりなのだが。

それはともかく，今回の Cisco 製品の脆弱性は telnet 絡みのもののようだ。

{{< fig-quote title=" CIAの機密文書で発覚、シスコ製品300種類にパッチ提供未定の危険な脆弱性" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/032300899/?rt=nocnt" >}}
<q>CMPとは、クラスターを構成する別の機器に、TELNETを使ってメッセージやコマンドを送信するためのプロトコル。今回見つかったCMPの脆弱性を突くと、攻撃者はアクセス権限のないシスコ製品にTELNETで接続できる。その結果、その製品を乗っ取ったり、製品上でウイルスを実行したりすることが可能になる。</q>
{{< /fig-quote >}}

言うまでもないが，もはや telnet は使ってはいけない。
可能な限り無効にするべき。

{{< fig-quote title=" CIAの機密文書で発覚、シスコ製品300種類にパッチ提供未定の危険な脆弱性" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/032300899/?rt=nocnt" >}}
<q>パッチ提供前の対策として同社が挙げているのが、IOS/IOS XEでTELNETを無効にすること。これにより、脆弱性がある機器でも攻撃されることを防げる。TELNETの代わりには、SSHを利用するよう推奨している。また、TELNETを無効にできない環境では、アクセス制御で対応するよう呼びかけている。</q>
{{< /fig-quote >}}

その上で ssh を使えばいいのだが ssh でもパスワード認証は比較的破られやすいことが分かっているので必ず公開鍵暗号を使った認証を使うこと。
ついでに ssh の version 1 は使わないこと。

- {{< pdf-file title="セキュアプロトコルに対する攻撃法等に関する技術調査" link="https://www.ipa.go.jp/security/fy15/reports/protocol/documents/protocol_2003.pdf" >}}
- {{< pdf-file title="SSH パスワードユーザ認証の脆弱性とその考察" link="https://ipsj.ixsq.nii.ac.jp/ej/?action=repository_action_common_download&item_id=10325&item_no=1&attribute_id=1&file_no=1" >}}

Vault 7 では脆弱性情報について PoC (Proof of Concept) や exploit code などの詳細情報がなく解析に時間がかかっているようだ[^r1]。

[^r1]: Vault 7 のリーク元がロシアということもあって情報自体の信憑性を確認する必要もある。

{{< fig-quote title=" CIAの機密文書で発覚、シスコ製品300種類にパッチ提供未定の危険な脆弱性" link="http://itpro.nikkeibp.co.jp/atcl/column/14/346926/032300899/?rt=nocnt" >}}
<q>ところが今回のケースでは、CIAから流出したとする機密文書Vault 7から脆弱性の存在が判明した。このためシスコは、詳細な情報やエクスプロイトを入手していない可能性がある。<br>
実際、シスコの公式ブログには、「（今回の脆弱性に関連する）ツールやマルウエアは公表されていないので、シスコが取れるアクションは限られている。より多くの情報を入手するまでは、脆弱性ハンドリングの観点でシスコができることはほとんどない」としている。</q>
{{< /fig-quote >}}

{{% md-box %}}
**追記（2017-04-04）**

- [暴露されたCIAの諜報能力「Vault 7」の衝撃度（前編） | THE ZERO/ONE](https://the01.jp/p0004740/)

この記事によるとちょっと状況が異なるようである。

{{< fig-quote title="暴露されたCIAの諜報能力「Vault 7」の衝撃度（前編）" link="https://the01.jp/p0004740/" >}}
<q>それに対して今回のYear Zeroが公開したのは、CIAが利用しているサイバー兵器そのものの情報だ。「CIAはこんなツールを使っていますよ」ではなく、CIAの利用しているゼロデイ、マルウェア、エクスプロイトなどの現物の情報である。WikiLeaksは、「このスマートフォンの脆弱性を利用してインストールできるマルウェア」や、「あの暗号化通信のアプリケーションを迂回できるツール」をソースコードごと入手しており、閲覧しても問題のない形で（＝他者から悪用されない形で）それらの情報を公開している。</q>
{{< /fig-quote >}}

つまり WikiLeaks 側は脆弱性情報に対して exploit code 等の詳細情報を持っていることになる。
であれば，後述のように，ベンダ企業にはその詳細情報が渡っている筈である（多分）。
{{% /md-box %}}

WikiLeaks 側は製品のベンダ企業に対して脆弱性の一般公開まで90日間の猶予を設けているらしい（予告なしに公開はしないということ）。

{{< fig-quote title="WikiLeaks Won’t Tell Tech Companies How to Patch CIA Zero-Days Until Its Demands Are Met" link="https://motherboard.vice.com/en_us/article/wikileaks-wont-tell-tech-companies-how-to-patch-cia-zero-days-until-its-demands-are-met" lang="en" >}}
<q>WikiLeaks included a document in the email, requesting the companies to sign off on a series of conditions before being able to receive the actual technical details to deploy patches, according to sources. It's unclear what the conditions are, but a source mentioned a 90-day disclosure deadline, which would compel companies to commit to issuing a patch within three months.</q>
{{< /fig-quote >}}

90日の猶予期間というのは Google の Project Zero などでもやっているので特に無理筋な設定ではないと思うが Cisco みたいな企業だと厳しいかもしれない。
一方で Android や iOS などでは発覚した脆弱性の殆んどは修正済みだったらしい。
そういった脆弱性情報の解析・対処の能力の違いはありそうである。

- [「CIAハッキングの脆弱性のほとんどは対応済み」とGoogleやAppleは公表するものの依然として危険な状態は続いている - GIGAZINE](http://gigazine.net/news/20170310-apple-google-treat-cia-hucking/)

この問題の根源は CIA などの諜報機関が脆弱性情報を「悪用」するために秘匿していたことである。
よく「隠すことによるセキュリティはダメ」と言われるが，目的はどうあれ脆弱性情報を秘匿することがどういう結果をもたらすか身に染みたのではないだろうか。

{{< fig-quote title="WikiLeaks Not Disclosing CIA-Hoarded Vulnerabilities to Companies" link="https://www.schneier.com/blog/archives/2017/03/wikileaks_not_d.html" lang="en" >}}
<q>Honestly, at this point the CIA should do the right thing and disclose all the vulnerabilities to the companies. They're burned as CIA attack tools. I have every confidence that Russia, China, and several other countries can hack WikiLeaks and get their hands on a copy. By now, their primary value is for defense. The CIA should bypass WikiLeaks and get the vulnerabilities fixed as soon as possible.</q>
{{< /fig-quote >}}

## ブックマーク

- [Cisco IOS and IOS XE Software Cluster Management Protocol Remote Code Execution Vulnerability](https://tools.cisco.com/security/center/content/CiscoSecurityAdvisory/cisco-sa-20170317-cmp)
- [WikiLeaks Won’t Tell Tech Companies How to Patch CIA Zero-Days Until Its Demands Are Met - Motherboard](https://motherboard.vice.com/en_us/article/wikileaks-wont-tell-tech-companies-how-to-patch-cia-zero-days-until-its-demands-are-met)
- [WikiLeaks Not Disclosing CIA-Hoarded Vulnerabilities to Companies - Schneier on Security](https://www.schneier.com/blog/archives/2017/03/wikileaks_not_d.html)
- [Is Security by Obscurity a valid approach? (大論争 : 隠すことによるセキュリティ)](https://technet.microsoft.com/ja-jp/library/2008.06.obscurity.aspx)
- [暴露されたCIAの諜報能力「Vault 7」の衝撃度（前編） | THE ZERO/ONE](https://the01.jp/p0004740/)
    - [暴露されたCIAの諜報能力「Vault 7」の衝撃度（中編） | THE ZERO/ONE](https://the01.jp/p0004753/)
    - [暴露されたCIAの諜報能力「Vault 7」の衝撃度（後編） | THE ZERO/ONE](https://the01.jp/p0004767/)
