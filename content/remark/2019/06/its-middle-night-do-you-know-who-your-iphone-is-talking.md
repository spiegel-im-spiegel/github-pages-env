+++
title = "午前3時のケータイ・トーク"
date =  "2019-06-26T14:53:22+09:00"
description = "“It’s 3 a.m. Do you know what your iPhone is doing?”"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "k-tai", "surveillance-capitalism" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## {{< quote lang="en" >}}It’s 3 a.m. Do you know what your iPhone is doing?{{< /quote >}}

ワシントン・ポストが面白い記事を掲載しているらしい。

- [Apple promises privacy, but iPhone apps share your data with trackers, ad companies and research firms - The Washington Post](https://www.washingtonpost.com/technology/2019/05/28/its-middle-night-do-you-know-who-your-iphone-is-talking/)

残念なことに私はワシントン・ポストの記事は読めない（お金を払ってない）ので，こちらの記事で代替することにする。

- [It’s the middle of the night. Do you know who your iPhone is talking to?](https://www.msn.com/en-us/news/technology/its-the-middle-of-the-night-do-you-know-who-your-iphone-is-talking-to/ar-AAC1Wvl)

記事の内容はユーザが寝ている深夜にスマホのアプリがプライバシー情報を含む個人データを「報告」しているというものである（それ以外の Apple 礼賛文章は華麗にスルーする）。

{{< fig-quote type="md" title="It’s the middle of the night. Do you know who your iPhone is talking to?" link="https://www.msn.com/en-us/news/technology/its-the-middle-of-the-night-do-you-know-who-your-iphone-is-talking-to/ar-AAC1Wvl" lang="en" >}}
{{< quote >}}On a recent Monday night, a dozen marketing companies, research firms and other personal data guzzlers got reports from my iPhone. At 11:43 p.m., a company called Amplitude learned my phone number, email and exact location. At 3:58 a.m., another called Appboy got a digital fingerprint of my phone. At 6:25 a.m., a tracker called Demdex received a way to identify my phone and sent back a list of other trackers to pair up with.{{< /quote >}}
{{< /fig-quote >}}

こうした「報告」を行う iPhone アプリのいくつかは特定されていて

{{< fig-quote type="md" title="It’s the middle of the night. Do you know who your iPhone is talking to?" link="https://www.msn.com/en-us/news/technology/its-the-middle-of-the-night-do-you-know-who-your-iphone-is-talking-to/ar-AAC1Wvl" lang="en" >}}
{{< quote >}}IPhone apps I discovered tracking me by passing information to third parties — just while I was asleep — include Microsoft OneDrive, Intuit’s Mint, Nike, Spotify, The Washington Post and IBM’s the Weather Channel. One app, the crime-alert service Citizen, shared personally identifiable information in violation of its published privacy policy.{{< /quote >}}
{{< /fig-quote >}}

と書かれている。
ワシントン・ポストもか（笑）
もちろんこれらの例は網羅的ではないだろうが，網羅的に調べるのは骨が折れるだろうな。

そしてこれは iPhone だけの問題ではない，当然ながら。

{{< fig-quote type="md" title="It’s the middle of the night. Do you know who your iPhone is talking to?" link="https://www.msn.com/en-us/news/technology/its-the-middle-of-the-night-do-you-know-who-your-iphone-is-talking-to/ar-AAC1Wvl" lang="en" >}}
{{< quote >}}Yes, trackers are a problem on phones running Google’s Android, too. Google [won’t even let Disconnect’s tracker-protection software](https://disconnect.me/blog/update-android-app-is-still-banned-from-play-and-google-wont-talk-about-it) into its Play Store. (Google’s rules prohibit apps that might interfere with another app displaying ads.){{< /quote >}}
{{< /fig-quote >}}

そういえば先日，[ケータイが異常動作を起こした]({{< relref "./programmable-controller.md" >}} "【憶測記事】 マイニング・コードを仕込まれたかもしれない")ときに簡単なネットワーク監視ツールを入れて調べてみたのだが，相手先を調べるのも面倒なほどのトラフィックがあり「だめだこりゃ」と思ったものである。
あれを調査するのはシロートでは無理。

私は[先日の件]({{< relref "./programmable-controller.md" >}} "【憶測記事】 マイニング・コードを仕込まれたかもしれない")を教訓としてケータイを使わないときは積極的に節電モードに切り替えるようにした[^no1]。
私のケータイは節電モードにするとデータ通信が一切シャットダウンされるためメールも取れない状態になるが（電話はつながる），どうせ日中や深夜に緊急メールをよこすような人は（今は）いないので無問題。
あとは [FREEDOME](https://www.f-secure.com/en/home/products/freedome "F-Secure FREEDOME VPN — Protect your privacy | F-Secure") でトラッキング・コードをブロックすれば多少はマシかな。
マシになるといいなぁ。

[^no1]: でもたまにケータイを点けたまま値落ちしてることがあるんだよなぁ。

## {{< quote lang="en" >}}Cambridge Analytica was just the beginning{{< /quote >}}

1年ほど前に

{{< fig-quote type="md" title="「情報弱者」を再々定義する" link="/remark/2018/05/information-illiterate/" >}}
{{< quote >}}確かに Facebook は [Campbridge Analytica の件]({{< ref "/remark/2018/03/name-identification.md" >}} "善悪の葛藤")以来さまざまな点で批判の矢面に立っているが Facebook に論点を絞るのは問題の矮小化といえる。
もしメディアが意図的に矮小化を行っているのなら，これこそがフィルターバブルの典型例と言えるだろう。{{< /quote >}}
{{< /fig-quote >}}

と書いたが，今回紹介した記事でも

{{< fig-quote type="md" title="It’s the middle of the night. Do you know who your iPhone is talking to?" link="https://www.msn.com/en-us/news/technology/its-the-middle-of-the-night-do-you-know-who-your-iphone-is-talking-to/ar-AAC1Wvl" lang="en" >}}
{{< quote >}}Part of Jackson’s objection to trackers is that many feed the personal data economy, used to target us for marketing and political messaging. Facebook’s fiascos have made us all more aware of how our data can be passed along, stolen and misused — but Cambridge Analytica was just the beginning.{{< /quote >}}
{{< /fig-quote >}}

と書かれており，よーやくメディアもそこに至ったかと感慨深い。
まぁ周回遅れの日本は今だに GAFA 礼賛だそうだが（笑）

## ブックマーク

- [iPhone Apps Surreptitiously Communicated with Unknown Servers - Schneier on Security](https://www.schneier.com/blog/archives/2019/06/iphone_apps_sur.html)

- [搾取と狂狷]({{< ref "/remark/2019/06/kyoken.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%B6%85%E7%9B%A3%E8%A6%96%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC-ebook/dp/B01MZGVHOA?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01MZGVHOA"><img src="https://images-fe.ssl-images-amazon.com/images/I/51T6PBdGbyL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%B6%85%E7%9B%A3%E8%A6%96%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC-ebook/dp/B01MZGVHOA?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01MZGVHOA">超監視社会</a></dt>
	<dd>ブルース・シュナイアー</dd>
	<dd>池村 千秋 (翻訳)</dd>
    <dd>草思社 2016-12-13 (Release 2017-02-03)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01MZGVHOA</dd>
  </dl>
  <p class="description">実は積ん読のまま読んでない。そろそろちゃんと最後まで読まないと。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-23">2019-03-23</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>

<div class="hreview" >
	<div class="photo"><a class="item url" href="https://tatsu-zine.com/books/infoshare2"><img src="https://tatsu-zine.com/images/books/877/cover_s.jpg" alt="photo"></a></div>
    <dl class="fn">
      <dt><a href="https://tatsu-zine.com/books/infoshare2">もうすぐ絶滅するという開かれたウェブについて 続・情報共有の未来</a></dt>
      <dd>yomoyomo</dd>
      <dd>達人出版会 2017-12-25</dd>
      <dd>評価&nbsp;<abbr class="rating fa-sm" title="4">
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="fas fa-star"></i>
        <i class="far fa-star"></i>
      </abbr></dd>
    </dl>
    <p class="description"><a href="https://wirelesswire.jp/author/yomoyomo/">WirelessWire News 連載</a>の書籍化。感想は<a href="/remark/2019/01/infoshare2/">こちら</a></p>
	<p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018-12-31</abbr></p>
</div>
