+++
title = "「青空文庫」用の書影がほしい"
date = "2019-03-30T13:27:29+09:00"
description = "青空文庫収録作品の書影に使えるかっこいい画像を教えてください。 CC0 下で aozora-cover リポジトリに含めてもいいよ，という奇特な方がいらっしゃいましたらバンバン pull request していただいて構いません。"
image = "https://photo.baldanders.info/flickr/image/32118886_o.jpg"
tags = [ "book", "aozora", "image" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ここを読んでくださっている皆様にはありがとうございます。
[前にも書いた]({{< ref "/release/2019/01/amazon-product-advertising-api.md" >}} "しょうがないので Amazon アフィリエイト・リンク作成ツールを作ったですよ")とおり [PA-API] を利用するには月に1冊以上売り上げないといけないのだが，幸いなことにもうしばらくは続けられそうである。

というわけで，営業活動に勤しむのはこのくらいにして，今回は[青空文庫]の話。

このブログで本を紹介する際には，こんな感じでカードっぽく表記するわけだけど

{{% review-paapi "4899840721" %}} <!-- インターネット図書館 青空文庫 -->

これを[青空文庫]に収録されている作品でもやりたいわけ。
でも[青空文庫]に収録されている作品にはいわゆる「書影」がないので，ちょっと寂しい感じになってしまうのですよ。

<div class="hreview">
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/000021/card4307.html">グリゴリの捕縛</a></dt>
    <dd>白田 秀彰</dd>
    <dd> 2001-11-26 (Release 2014-09-17)</dd>
    <dd>青空文庫</dd>
    <dd>4307 (図書カードNo.)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">白田秀彰さんの「<a href="http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm">グリゴリの捕縛</a>」が青空文庫に収録されていた。
内容は <delete>怪獣大決戦</delete> おっと憲法（基本法）についてのお話。
古代社会 → 中世社会 → 近代社会 → 現代社会 と順を追って法と慣習そして力（power）との関係について解説し，その中で憲法（基本法）がどのように望まれ実装されていったか指摘してる。
その後，現代社会の次のパラダイムに顕現する「情報力」と社会との関係に言及していくわけだ。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-30">2019-03-30</abbr> (powered by <a href="https://aozorahack.org/">aozorahack</a>)</p>
</div>

で，[青空文庫]用に統一されたデザインの書影っぽい画像がないかなぁ，とググってみたんだけどそれぽいのはない様子。
底本の書影を使う手もあるけど書影の著作権は作品の権利とは異なることも多いので迂闊に使えない（だから無理しても [PA-API] を利用してるんだけど）。

ちうわけで作ってしまいました。

- [spiegel-im-spiegel/aozora-cover: Book Cover images for Aozora-bunko](https://github.com/spiegel-im-spiegel/aozora-cover)

リポジトリでは A5 サイズのアスペクト比になるように，こんな感じの背景画像を公開してみた。

{{< fig-img src="/images/aozora/cover-raw.jpg" link="/images/aozora/cover-raw.jpg" >}}

[CC0] で公開しているので，[青空文庫]にある公有（public domain）の作品と組み合わせても全く問題ないと思う。

これに[青空文庫]のロゴとタイトルを重ねて SVG 形式でこんな感じにする。

{{< fig-img src="/images/aozora/card4307.svg" link="/images/aozora/card4307.svg" >}}

これで先程のカードは以下のように表示できた。

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.aozora.gr.jp/cards/000021/card4307.html"><img src="https://text.baldanders.info/images/aozora/card4307.svg" width="110" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.aozora.gr.jp/cards/000021/card4307.html">グリゴリの捕縛</a></dt>
    <dd>白田 秀彰</dd>
    <dd> 2001-11-26 (Release 2014-09-17)</dd>
    <dd>青空文庫</dd>
    <dd>4307 (図書カードNo.)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">白田秀彰さんの「<a href="http://orion.mt.tama.hosei.ac.jp/hideaki/kenporon.htm">グリゴリの捕縛</a>」が青空文庫に収録されていた。
内容は <delete>怪獣大決戦</delete> おっと憲法（基本法）についてのお話。
古代社会 → 中世社会 → 近代社会 → 現代社会 と順を追って法と慣習そして力（power）との関係について解説し，その中で憲法（基本法）がどのように望まれ実装されていったか指摘してる。
その後，現代社会の次のパラダイムに顕現する「情報力」と社会との関係に言及していくわけだ。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-30">2019-03-30</abbr> (powered by <a href="https://aozorahack.org/">aozorahack</a>)</p>
</div>

ちなみに背景画像の元ネタはこれ。

{{< fig-img src="https://photo.baldanders.info/flickr/image/32118886_m.jpg" title="In the beach — Photos in Flickr | photo.Baldanders.info" link="https://photo.baldanders.info/flickr/32118886/" >}}

この写真は [CC BY] で公開しているが，私の作品なので，私が翻案（トリミングして彩度を調整した程度だけど）を [CC0] で公開したところで全然問題ない。

というわけで，[青空文庫]収録作品の書影に使えるかっこいい画像を教えてください。
[CC0] 下で [aozora-cover] リポジトリに含めてもいいよ，という奇特な方がいらっしゃいましたらバンバン pull request していただいて構いません。

[PA-API]: https://affiliate.amazon.co.jp/assoc_credentials/home "Product Advertising API"
[青空文庫]: https://www.aozora.gr.jp/ "青空文庫　Aozora Bunko"
[aozora-cover]: https://github.com/spiegel-im-spiegel/aozora-cover "spiegel-im-spiegel/aozora-cover: Book Cover images for Aozora-bunko"
[CC0]: http://creativecommons.org/publicdomain/zero/1.0/ "Creative Commons — CC0 1.0 Universal"
[CC BY]: https://creativecommons.org/licenses/by/4.0/ "Creative Commons — Attribution 4.0 International — CC BY 4.0"
