+++
title = "本当のチャイナ・リスク"
date =  "2019-05-22T12:52:58+09:00"
description = "モノやサービスで溢れかえる時代に生きている私達の自由は意外に貧弱なものである。それを握っているのは「私達」ではないのだから。"
image = "/images/attention/kitten.jpg"
tags = [ "risk", "management", "market" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「チャイナ・リスク」というフレーズが出回ったのはゼロ年代，ブッシュ共和党政権の時である。
その後，オバマ民主党政権に代わって「チャイナ・リスク」というフレーズは聞こえなくなり，むしろ中国を重要なビジネス・パートナーと見做すようになっていく。
それがトランプ共和党政権でまた反転する。
この辺の見事な色分けは流石米国としか言いようがない。
官僚主導で硬直しきった日本の政治状況と比べれば羨ましいことである。

- [Googleがファーウェイに対してAndroidのサポートを中止へ  |  TechCrunch Japan](https://techcrunch.com/2019/05/19/google-reportedly-suspends-select-business-with-huawei-following-u-s-ban/)
- [MIT Tech Review: グーグルが一部取引を中止、「ファーウェイ排除」の波紋](https://www.technologyreview.jp/nl/google-has-blocked-huawei-from-using-android-in-any-new-phones/)
- [既存ファーウェイ端末はGoogle Playストアを継続利用可能とグーグルが声明  |  TechCrunch Japan](https://techcrunch.com/2019/05/21/google-says-its-app-store-will-continue-to-work-for-existing-huawei-smartphone-owners/)
- [Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も  |  TechCrunch Japan](https://techcrunch.com/2019/05/20/trumps-huawei-ban-also-causing-tech-shocks-in-europe/)

Google の思惑は分からない。
「[（大統領の）行政命令を遵守するため](https://techcrunch.com/2019/05/20/trumps-huawei-ban-also-causing-tech-shocks-in-europe/ "Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も  |  TechCrunch Japan")」というのはいかにも後付けくさい。
ただ Huawei (華為) はこの決定にあまり困らないんじゃないのかなぁ。
そして「あまり困らない」ことこそが本当のリスクと言える。
何故なら Huawei 製端末においてはコントロールの主体が Google から Huawei に移ることを意味するから。
まぁ Google と Huawei のどっちがマシかというリスク・トレードオフの問題でもあるのだが。

{{< fig-quote title="Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も" link="https://techcrunch.com/2019/05/20/trumps-huawei-ban-also-causing-tech-shocks-in-europe/" >}}
<q>スマートフォンにおけるEUの反トラスト措置の核心はAndroidデバイスとグーグルのポピュラーなアプリをアンバンドルすることだ。これによりスマートフォンメーカーはグーグルのブランドを維持したまま完全にグーグルの支配下にあるのではないデバイスを販売できる。例えば、Playストアをプレロードするものの、デフォールトの検索エンジンやブラウザにグーグル以外のプロダクトを設定するなどだ。<br>
しかしグーグルが（現行モデルでは継続されるとしても）ファーウェイに新しいAndroid OSやGoogle Playストアを提供しないとなれば、こうした構想は崩れてしまう。</q>
{{< /fig-quote >}}

OS ないしはハードウェア・ベンダが「アプリケーション・ストア」を通じてソフトウェア供給を支配下に置く，いわゆる[「アプリケーション経済」は故 Steve Jobs の考えたビジネスモデルと聞いたことがある](https://japan.zdnet.com/article/35108430/ "「スティーブ・ジョブズがいなければセールスフォースはなかった」：ベニオフCEO - ZDNet Japan")が，結局のところそれはアプリケーション流通を利用した CRM (Customer Relationship Management) 覇権競争にほかならない。

「アプリケーション・ストア」は APT (Advanced Package Tool) のようなパッケージ管理ツールとは根本的に異なる。
たとえば APT は簡単にサードパーティのリポジトリを組み込むことができ，その信用モデルは [OpenPGP] の「信用の輪（web of trust）」で構成されている[^wot1]。
どちらかというとこれは VRM (Vendor Relationship Management) に近い。

[^wot1]: 「信用の輪」については拙文「[OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})」で少し紹介しているので参考にどうぞ。

今回の件をきっかけに「アプリケーション経済」がもっと VRM 寄りになっていくと面白いのだが，流石にそれは夢を見すぎか（笑） 折角の AI 技術も「[バーチャル外商]({{< ref "/remark/2018/04/ai-assistant.md" >}} "AI アシスタントはユーザをアシストしない")」としてしか使われないのなら宝の持ち腐れだよなぁ。

モノやサービスで溢れかえる時代に生きている私達の自由は意外に貧弱なものである。
それを握っているのは「私達」ではないのだから。

## ブックマーク

- [ARMが米方針に従いファーウェイとの取引を停止  |  TechCrunch Japan](https://techcrunch.com/2019/05/23/arm-halts-huawei-relationship-following-us-ban/)
- [中国最大のチップメーカーがニューヨーク証券取引所上場廃止へ  |  TechCrunch Japan](https://techcrunch.com/2019/05/24/smic-nasdaq-delisting/)

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"

## 参考図書

{{% review-paapi "B00DIM6BE6" %}} <!-- インテンション・エコノミー -->
{{% review-paapi "4532149029" %}} <!-- クルートレイン宣言 -->
