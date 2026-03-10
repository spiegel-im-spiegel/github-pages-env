+++
title = "そろそろ PDS を個人で管理する時代か？"
slug = "pds-for-everyone"
date =  "2026-03-10T13:13:04+09:00"
description = "Eurosky と Periwinkle / コリイ・ドクトロウ 𝕏 やめるってよ / 欧州版 𝕏 じゃない，キラーアプリじゃない，万能アプリじゃない / ツールを直さなきゃ"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "atproto", "activitypub", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## Eurosky と Periwinkle

ここのところ私のご近所 TL で話題の [Eurosky] がローンチしたのは昨年の11月で，今月に入って一般公開されたらしい。

- [EU entrepreneurs launch social media infrastructure to liberate Europe from US tech monopolies — eurosky](https://www.eurosky.tech/blog/eurosky-launch-social-media-infrastructure)
- [A Eurosky Account is just the start - Eurosky's blog](https://eurosky.leaflet.pub/3mgnjewfva22q)

私がこれを知ったのは[骨折入院]中で「面白そうなことをやってるなぁ」と指をくわえて見てたのだが，記事を読んでみるに，かなり政治的動機の強いサービスのようだ。
ちなみに資金は個人や組織からの寄付で賄っている[そうな](https://www.eurosky.tech/faq "FAQ — eurosky")。
ユーザに割り当てられるリソースの詳細は明記されていない。

一方，今日はじめて知ったのだが， [Periwinkle] というサービスがあるらしい。
こちらは早期アクセスの段階。
TechCrunch が取り上げていた。

- [Introducing Periwinkle: Truly own your social identity on Bluesky and AT Protocol - Periwinkle.social Blog](https://blog.periwinkle.social/3mgni4wwqks22)
- [Periwinkle is making self-hosted social media on Bluesky's AT Protocol even easier | TechCrunch](https://techcrunch.com/2026/03/09/periwinkle-at-protocol-bluesky-self-hosted-social-media/)

TechCrunch はこれをサードパーティのホスティング・プロバイダによって管理されている WordPress インストールみたいだと評している。

{{< fig-quote type="markdown" title="Periwinkle is making self-hosted social media on Bluesky's AT Protocol even easier | TechCrunch" link="https://techcrunch.com/2026/03/09/periwinkle-at-protocol-bluesky-self-hosted-social-media/" lang="en" >}}
This idea is similar to how people who want to create their own blog can choose between an easy-to-use hosted blog on WordPress.com, a DIY installation using the open source software from WordPress.org, or something in between — like a WordPress installation that’s managed by a third-party hosting provider. Periwinkle is offering the same, but for decentralized social media.
{{< /fig-quote >}}

んで，こちらはちゃんとスペックと料金プランが示されている。

{{< fig-img-quote src="./plans.png" title="Periwinkle - Get Started" link="https://periwinkle.social/plans" width="1434" >}}

Basic プランで月600円くらいか？ それなら個人でも無理なく行けそう？

[Kagi Assistant] に [Eurosky] と [Periwinkle] を比較してもらった。

{{< div-box type="markdown" >}}
PeriwinkleとEuroskyは、どちらもBlueskyの基盤である**AT Protocol**を活用し、ユーザーがデータを自己所有する「分散型ソーシャルメディア」の普及を目指していますが、そのアプローチや目的には明確な違いがあります。

### 共通点

- **AT Protocolの採用**
    - 両者ともAT Protocolを基盤としており、Blueskyなどの他のネットワークと相互運用性があります。
- **データの自己所有（自己主権型アイデンティティ）**
    - ユーザーが自身のデータやフォロワーとの繋がりを所有し、特定のプラットフォームに縛られず（ロックインされず）に移動できる仕組みを重視しています。
- **運用の簡略化**
    - サーバー構築などの技術的なハードルを下げ、一般ユーザーが容易に「自分のデータサーバー（PDS）」や「自分のアカウント」を持てるサービスを提供しています。

### 違い

| 項目 | Periwinkle | Eurosky |
| :--- | :--- | :--- |
| **主な役割** | **マネージドPDSホスティング** | **ソーシャル・インフラストラクチャ** |
| **重点を置く点** | 個人のサーバー運営の代行と管理ツールの提供。 | 欧州のデータ主権確立と独自の社会的ウェブ構築。 |
| **政治・地理的背景** | 特定の地域に限定せず、個人やチームの利便性を追求。 | **欧州（EU）** の自立を掲げ、米国テック企業の独占に対抗。 |
| **主なターゲット** | 自分のドメインで発信したい個人、クリエイター、小規模チーム。 | 欧州のユーザー、開発者、および民主的な対話空間を求める層。 |
| **提供価値** | ~~24/7監視~~、バックアップ、複数ハンドル管理などの「運用保守」。 | ~~GDPR準拠~~、欧州内でのデータ処理、~~独自のコンテンツ審査体制~~。 |

### まとめ

- **Periwinkle**は、いわば「AT Protocol界のWordPress.com」のような存在で、技術的な手間を省いて**自分の城（PDS）を持ちたい人向けのツール**を提供しています。
- **Eurosky**は、より大きな枠組みで**欧州独自のデジタル公共圏**を作ろうとしており、インフラの提供だけでなく、政治的な独立性や民主的な対話の場を重視しています。
{{< /div-box >}}

パッと見で「そんなこと書いとらんやろ」って部分は打ち消し線をしている。
これだから生成 AI は... それ以外はだいたい合ってると思う。
各自でも確認してね。

なお PDS リポジトリデータの移動には [PDS MOOver] が使える。
これで `bsky.social` にあるデータを失わずに新しい PDS に移せるし，万一利用してる PDS がポシャってもデータ移行はできそうである（突然死はどうしようもないが）。

ところで，これは完全に余談なのだが [Eurosky] も [Periwinkle] もブログ運営を [Leaflet] で行ってるのが面白い。
やはり時代は [Leaflet] なのか？

## コリイ・ドクトロウ {{% emoji "X" %}} やめるってよ

今まで散々 Bluesky を dis ってた SF 作家のコリイ・ドクトロウ氏が，ついに Bluesky に参加し {{% emoji "X" %}} からは完全に足を洗うようだ。

{{< fig-quote type="markdown" title="Twitterとの決別――マスクはカエルを火炎放射器で丸焼きにした » p2ptk[.]org" link="https://p2ptk.org/digital-rights/5452" >}}
だが――Blueskyの名誉のために言えば――彼らはやがてそれを理解し、独自のBlueskyサーバを構築するためのツールと手順を公開した。ケンがすぐに調べてくれて、すべて可能だと教えてくれたが、トロントのコロケーションケージに置いてあるLinuxボックスのハードウェアアップグレードが完了するまでは無理だという。そのアップグレードが数ヶ月前に終わり、昨日、ケンが私専用のBlueskyサーバの設置が完了したと知らせてくれた。こうして私はBlueskyに参加した。@doctorow.pluralistic.netである。

https://bsky.app/profile/doctorow.pluralistic.net
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Twitterとの決別――マスクはカエルを火炎放射器で丸焼きにした » p2ptk[.]org" link="https://p2ptk.org/digital-rights/5452" >}}
もう何年もの間、私はほぼ毎日、Twitterにこの告知を投稿してきた。

>Twitterは日に日に悪化している。いつか使えなくなるまで劣化するだろう。メタクソ化に強い代替手段として、Fediverseが我々の最善の希望だ。私は@pluralistic@mamot.frにいる。

今日、私はその修正版を投稿する。以下の一文が追加される。

>Blueskyでフォローしたい方は、@doctorow.pluralistic.netへ。これが私がTwitterに投稿する最後のスレッドだ。
{{< /fig-quote >}}

どうやら彼は自前の PDS を立てたようだ。
DID が `did:web:pluralistic.net` になっていて [PDSls](https://pdsls.dev/at://did:web:pluralistic.net) からは見れなかった。

自前の PDS を立てる方法は[公式ドキュメントにある](https://atproto.com/ja/guides/self-hosting "Self-hosting - AT Protocol Docs - AT Protocol")んだが，これを読んでもできる気がしない。
個人でやるなら [Eurosky] なり [Periwinkle] なり，あるいは他の PDS を提供しているサービスを利用するほうが現実的だと思う。
さくらインターネットあたりがこういうサービスやらないかなぁ...

## 欧州版 {{% emoji "X" %}} じゃない，キラーアプリじゃない，万能アプリじゃない

そういえば，最近見たこのポストが面白かった。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:ad5bwszsc5m5jpj2sfa6uzjk/app.bsky.feed.post/3mgkfxwwtfs2y" data-bluesky-cid="bafyreifh43gtpxx4mpjaxbnlemqwdbyo7o7pgr[Eurosky]: https://www.eurosky.tech/ "eurosky"
so3mirv3bgpw43ctf44i" data-bluesky-embed-color-mode="system"><p lang="en">AT protocol doesn’t play zero-sum, winner-takes-all games.

“European X?”
No, European data sovereignty for all of x/y/z in amicable competition.

“Killer app?”
No, interoperable apps as the monopoly-killer.

“Everything-app?”
No, everything-account in a pluriverse of user-owned software.<br><br><a href="https://bsky.app/profile/did:plc:ad5bwszsc5m5jpj2sfa6uzjk/post/3mgkfxwwtfs2y?ref_src=embed">[image or embed]</a></p>&mdash; Erlend Sogge Heggen (<a href="https://bsky.app/profile/did:plc:ad5bwszsc5m5jpj2sfa6uzjk?ref_src=embed">@erlend.sh</a>) <a href="https://bsky.app/profile/did:plc:ad5bwszsc5m5jpj2sfa6uzjk/post/3mgkfxwwtfs2y?ref_src=embed">March 8, 2026 at 10:08 PM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

Blusky が提供する `bsky.app` も `bsky.social` も，大きな[エコシステム]({{< ref "/remark/2026/02/leaflet-and-the-atmosphere.md" >}} "Leaflet と The Atmosphere")の一部になりつつある。

{{< fig-img-quote src="/remark/2026/02/leaflet-and-the-atmosphere/Screenshot_20260204_121413_Samsung-Notes.jpg" title="The Everything Account" link="https://www.augment.ink/the-everything-account/" width="1000" lang="en">}}


[Eurosky] のドキュメントを見ても分かるように，欧州は [AT Protocol][atproto] を使ってリスク塗れのアメリカからの脱出を目指してるように見える。
Instagram の代替とされる [Flashes](https://techcrunch.com/2025/07/03/everything-you-need-to-know-about-flashes-the-bluesky-based-instagram-alternative/ "Everything you need to know about Flashes, the Bluesky-based Instagram alternative | TechCrunch") や GitHub の代替とされる [Tangled](https://tangled.org/ "tangled · tightly-knit social coding") や Google ニュースの代替とされる [Sill](https://sill.social/ "Sill | Top news shared by the people you trust") も [atproto] エコシステムの一部で，「[万能アカウント（Everything Account）](https://augment.ink/the-everything-account/ "The Everything Account · augment")」を **所有する** 私達はこれらのアプリやサービスの間を繋がりは維持したまま自由に渡り歩くことができる。

{{< fig-quote type="markdown" title="Twitterとの決別――マスクはカエルを火炎放射器で丸焼きにした » p2ptk[.]org" link="https://p2ptk.org/digital-rights/ 5452" >}}
危機は変化を引き起こす。ウラジーミル・プーチンやイーロン・マスクやドナルド・トランプがいなければ、世界はもっとマシだっただろう。だが、この我慢できない、気短な、最低の男たちには1つ用途がある。行動を促すには緩慢すぎる危機を、無視できない緊急事態へと変えてくれた。プーチンはEUを化石燃料からの脱却に押しやった。マスクは何百万の人々を連合ソーシャルメディアに押し出した。トランプはポスト・アメリカのインターネットへの道を開いた。

https://pluralistic.net/2026/01/01/39c3/#the-new-coalition
{{< /fig-quote >}}

私達は既存のソーシャル・メディアから押し出され Fediverse や Atmosphere に流れ着いたが，今になって見ればソーシャル・メディアの外にいるのは未だ既存のサービスに留まってる集団の方かも知れない。

{{< fig-quote type="markdown" title="『ラフ＆レディ』デマコーヴァ" link="https://comic.k-manga.jp/title/69737/" >}}
世界がジャングルを囲んでいるんじゃない。ジャングルが世界を囲んでいるのです
{{< /fig-quote >}}

## ツールを直さなきゃ

個人的には [Periwinkle] のほうはちょっと興味があるのだが，ボット運用に使ってる自作ツールが `bsky.social` 以外の PDS に対応してないのね。
だからまず，そこから手を付けないといけない。

あと，常用している[ぞーぺん（ZonePane）][ZonePane]も `bsky.social` 以外の PDS に対応してない。
これが一番困ってる。
[`@zonepane@fedibird.com`] に要望を出したほうがいいかなぁ...

## ブックマーク

- [W3C Invites Implementations of Decentralized Identifiers (DIDs) v1.1 | 2026 | News | W3C](https://www.w3.org/news/2026/w3c-invites-implementations-of-decentralized-identifiers-dids-v1-1/)

[骨折入院]: {{< ref "/remark/2025/12/benkei-iwa-and-fracture-and-hospital-stay.md" >}} "弁慶岩を見に行って骨折入院した話（前編）"
[Kagi Assistant]: https://kagi.com/assistant "Kagi Assistant"
[atproto]: https://atproto.com/ "The AT Protocol"
[Eurosky]: https://www.eurosky.tech/ "eurosky"
[Periwinkle]: https://periwinkle.social/ "Periwinkle - Get Started"
[PDS MOOver]: https://pdsmoover.com/ "PDS MOOver"
[Leaflet]: https://leaflet.pub/ "Leaflet"
[ZonePane]: https://twitpane.com/ "TwitPane – Twitter with Panes"
[`@zonepane@fedibird.com`]: https://fedibird.com/@zonepane "ZonePane(ぞーぺん) (@zonepane@fedibird.com) - Fedibird"
