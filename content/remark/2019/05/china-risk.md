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

- [Googleがファーウェイに対してAndroidのサポートを中止へ  |  TechCrunch Japan](https://jp.techcrunch.com/2019/05/20/2019-05-19-google-reportedly-suspends-select-business-with-huawei-following-u-s-ban/)
- [MIT Tech Review: グーグルが一部取引を中止、「ファーウェイ排除」の波紋](https://www.technologyreview.jp/nl/google-has-blocked-huawei-from-using-android-in-any-new-phones/)
- [既存ファーウェイ端末はGoogle Playストアを継続利用可能とグーグルが声明  |  TechCrunch Japan](https://jp.techcrunch.com/2019/05/21/2019-05-21-google-says-its-app-store-will-continue-to-work-for-existing-huawei-smartphone-owners/)
- [Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も  |  TechCrunch Japan](https://jp.techcrunch.com/2019/05/22/2019-05-20-trumps-huawei-ban-also-causing-tech-shocks-in-europe/)

Google の思惑は分からない。
「[（大統領の）行政命令を遵守するため](https://jp.techcrunch.com/2019/05/22/2019-05-20-trumps-huawei-ban-also-causing-tech-shocks-in-europe/ "Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も  |  TechCrunch Japan")」というのはいかにも後付けくさい。
ただ Huawei (華為) はこの決定にあまり困らないんじゃないのかなぁ。
そして「あまり困らない」ことこそが本当のリスクと言える。
何故なら Huawei 製端末においてはコントロールの主体が Google から Huawei に移ることを意味するから。
まぁ Google と Huawei のどっちがマシかというリスク・トレードオフの問題でもあるのだが。

{{< fig-quote title="Googleのファーウェイ制裁参加で欧州にショック拡大中、脱米模索も" link="https://jp.techcrunch.com/2019/05/22/2019-05-20-trumps-huawei-ban-also-causing-tech-shocks-in-europe/" >}}
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

- [ARMが米方針に従いファーウェイとの取引を停止  |  TechCrunch Japan](https://jp.techcrunch.com/2019/05/23/2019-05-23-arm-halts-huawei-relationship-following-us-ban/)

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%BB%E3%82%A8%E3%82%B3%E3%83%8E%E3%83%9F%E3%83%BC%EF%BD%9E%E9%A1%A7%E5%AE%A2%E3%81%8C%E6%94%AF%E9%85%8D%E3%81%99%E3%82%8B%E7%B5%8C%E6%B8%88-Doc-Searls-ebook/dp/B00DIM6BE6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DIM6BE6"><img src="https://images-fe.ssl-images-amazon.com/images/I/519%2BkIHb71L._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%A4%E3%83%B3%E3%83%86%E3%83%B3%E3%82%B7%E3%83%A7%E3%83%B3%E3%83%BB%E3%82%A8%E3%82%B3%E3%83%8E%E3%83%9F%E3%83%BC%EF%BD%9E%E9%A1%A7%E5%AE%A2%E3%81%8C%E6%94%AF%E9%85%8D%E3%81%99%E3%82%8B%E7%B5%8C%E6%B8%88-Doc-Searls-ebook/dp/B00DIM6BE6?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00DIM6BE6">インテンション・エコノミー～顧客が支配する経済</a></dt>
	<dd>Doc Searls</dd>
	<dd>栗原潔 (翻訳)</dd>
    <dd>翔泳社 2013-03-14 (Release 2013-06-20)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00DIM6BE6</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">時代はソーシャル CRM から VRM へ。<a href='https://baldanders.info/spiegel/log2/000794.shtml'>俺達がインターネットだ！</a> <a href='https://baldanders.info/spiegel/log2/000638.shtml'>感想文はこちら</a>。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-04-26">2015-04-26</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%93%E3%82%8C%E3%81%BE%E3%81%A7%E3%81%AE%E3%83%93%E3%82%B8%E3%83%8D%E3%82%B9%E3%81%AE%E3%82%84%E3%82%8A%E6%96%B9%E3%81%AF%E7%B5%82%E3%82%8F%E3%82%8A%E3%81%A0%E2%80%95%E3%81%82%E3%81%AA%E3%81%9F%E3%81%AE%E4%BC%9A%E7%A4%BE%E3%82%92%E7%B5%B6%E6%BB%85%E6%81%90%E7%AB%9C%E3%81%AB%E3%81%97%E3%81%AA%E3%81%8495%E3%81%AE%E6%B3%95%E5%89%87-%E3%83%AA%E3%83%83%E3%82%AF-%E3%83%AC%E3%83%90%E3%82%A4%E3%83%B3/dp/4532149029?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532149029"><img src="https://images-fe.ssl-images-amazon.com/images/I/51811H1N43L._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%93%E3%82%8C%E3%81%BE%E3%81%A7%E3%81%AE%E3%83%93%E3%82%B8%E3%83%8D%E3%82%B9%E3%81%AE%E3%82%84%E3%82%8A%E6%96%B9%E3%81%AF%E7%B5%82%E3%82%8F%E3%82%8A%E3%81%A0%E2%80%95%E3%81%82%E3%81%AA%E3%81%9F%E3%81%AE%E4%BC%9A%E7%A4%BE%E3%82%92%E7%B5%B6%E6%BB%85%E6%81%90%E7%AB%9C%E3%81%AB%E3%81%97%E3%81%AA%E3%81%8495%E3%81%AE%E6%B3%95%E5%89%87-%E3%83%AA%E3%83%83%E3%82%AF-%E3%83%AC%E3%83%90%E3%82%A4%E3%83%B3/dp/4532149029?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532149029">これまでのビジネスのやり方は終わりだ―あなたの会社を絶滅恐竜にしない95の法則</a></dt>
	<dd>リック レバイン, ドク サールズ, クリストファー ロック, デビッド ワインバーガー</dd>
	<dd>Rick Levine (原著), Doc Searls (原著), Christopher Locke (原著), David Weinberger (原著), 倉骨 彰 (翻訳)</dd>
    <dd>日本経済新聞社 2001-03</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4532149029, EAN: 9784532149024</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description"><a href='http://www2.gol.com/users/jheine/cluetrainj.html'>クルートレイン宣言</a>。タイトルがアレ過ぎる（笑） ちなみに今は <a href="https://github.com/dweinberger/newclues">New Clues</a> が登場している。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-01-14">2015-01-14</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
