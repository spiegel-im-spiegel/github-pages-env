+++
title = "Satellite Constellations : 「光害」問題は新たなステージへ"
date =  "2019-06-15T13:43:24+09:00"
description = "各団体の声明には激しく同意するところではあるが，一方で「これが宇宙時代のさきがけか」と感慨深い。今後の展開に注目しておこう。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "light-pollution", "satellite-constellation" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

まぁ，前々から ISS を始めとする低軌道人工衛星が撒き散らす「光害」を苦々しく思う人はいたわけよ（せっかく撮った天体写真に人工衛星の光跡が写り込んでいた日にはもう...）。
でも，人工衛星の軌道は予測可能でさほど実害があるわけでもなかったため愚痴程度で済んでいたのだけど，流石に今回は規模が違う。

## Satellite Constellations

2019年5月に SpaceX 社によって “Starlink” と呼ばれる60基の低軌道人工衛星群（satellite constellations）が打ち上げられた。

- [MIT Tech Review: スペースX、インターネット衛星60基を打ち上げ](https://www.technologyreview.jp/nl/spacex-has-launched-the-first-60-satellites-of-its-space-internet-system/)

人工衛星群を使った通信網の構想は昔から考えられ実際に稼働しているものもあるが，今回は夥しい数の人工衛星（SpaceX 社は最終的に1万2千基もの人工衛星を打ち上げる予定らしい）を低軌道（高度2千キロメートル未満）に打ち上げて網を構成する点が画期的である。
また SpaceX 社以外にも同様の低軌道人工衛星通信網の構築を目論んでいる企業があるらしい。

大抵の人工衛星の表面は反射率の高い金属製で太陽電池パネルも搭載されており，これらに太陽光が反射して明るく光ることがある。
これは俗に「人工衛星フレア」と呼ばれており，人工衛星オタクの観測対象となっている。
また通信衛星ということで，最終的にはたくさんの人工衛星によって常時空から通信電波が降り注ぐことになる[^r1]。

[^r1]: 念の為に言っておくが，私は携帯電話等の通信電波による人体への影響に関しては懐疑的である。ここではむしろ電波望遠鏡等を使った天文観測への影響を懸念している。先日，電波望遠鏡網をつかったブラックホールの直接観測に成功したってニュースになったばかりなのにねぇ。

というわけで，打ち上げ直後から批判が起きている。

- [Lights in the sky from Elon Musk's new satellite network have stargazers worried](https://phys.org/news/2019-05-sky-elon-musk-satellite-network.html)
- [SpaceXのインターネット衛星群「夜空に明るすぎる」と苦情。マスクCEOは対応を約束 - Engadget 日本版](https://japanese.engadget.com/2019/05/28/spacex-ceo/)
- [MIT Tech Review: 「明るすぎる」スペースXの人工衛星群、天文学者から批判](https://www.technologyreview.jp/nl/spacexs-starlink-satellites-are-clearly-visible-in-the-sky-and-astronomers-arent-happy/)

そして遂に IAU (International Astronomical Union; 国際天文学連合) から声明が出された。

- [IAU Statement on Satellite Constellations | IAU](https://www.iau.org/news/announcements/detail/ann19035/)

{{< fig-quote type="markdown" title="IAU Statement on Satellite Constellations" link="https://www.iau.org/news/announcements/detail/ann19035/" lang="en" >}}
{{< quote >}}The scientific concerns are twofold. Firstly, the surfaces of these satellites are often made of highly reflective metal, and reflections from the Sun in the hours after sunset and before sunrise make them appear as slow-moving dots in the night sky. Although most of these reflections may be so faint that they are hard to pick out with the naked eye, they can be detrimental to the sensitive capabilities of large ground-based astronomical telescopes, including the extreme wide-angle survey telescopes currently under construction. Secondly, despite notable efforts to avoid interfering with radio astronomy frequencies, aggregate radio signals emitted from the satellite constellations can still threaten astronomical observations at radio wavelengths. Recent advances in radio astronomy, such as producing [the first image of a black hole](https://iopscience.iop.org/article/10.3847/2041-8213/ab0e85/pdf) or understanding more about the formation of planetary systems, were only possible through concerted efforts in safeguarding the radio sky from interference.{{< /quote >}}
{{< /fig-quote >}}

さらに RAS (Royal Astronomical Society; 王立天文学会) や AAS (American Astronomical Society; アメリカ天文学会) も声明を出した。

- [RAS statement on Starlink satellite constellation | The Royal Astronomical Society](https://ras.ac.uk/news-and-press/news/ras-statement-starlink-satellite-constellation)
- [AAS Issues Position Statement on Satellite Constellations | American Astronomical Society](https://aas.org/media/press-releases/aas-issues-position-statement-satellite-constellations)

これらの声明は IAU のものより踏み込んだ内容になっていて

{{< fig-quote type="markdown" title="RAS statement on Starlink satellite constellation" link="https://ras.ac.uk/news-and-press/news/ras-statement-starlink-satellite-constellation" lang="en" >}}
{{< quote >}}Given the scale of these projects, there is also the prospect of a significant and lasting change to the views of the night sky until now enjoyed throughout human history and pre-history. The night sky is part of the cultural heritage of humanity, and the Society believes that it deserves protection.{{< /quote >}}
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="AAS Issues Position Statement on Satellite Constellations" link="https://ras.ac.uk/news-and-press/news/ras-statement-starlink-satellite-constellation" lang="en" >}}
{{< quote >}}"The natural night sky is a resource not just for astronomers but for all who look upward to understand and enjoy the splendor of the universe, and its degradation has many negative impacts beyond the astronomical," says Jeffrey C. Hall (Lowell Observatory), Chair of the [AAS Committee on Light Pollution, Radio Interference](https://aas.org/comms/committee-light-pollution-radio-interference-and-space-debris), and Space Debris. "I appreciate the initial conversation we have already had with SpaceX, and I look forward to working with my AAS colleagues and with all stakeholders to understand and mitigate the effects of the rapidly increasing numbers of satellites in near-Earth orbit."{{< /quote >}}
{{< /fig-quote >}}

と述べている。

個人的には人工流星の商業イベントの利用すらよく思ってないので各団体の声明には激しく同意するところではあるが，一方で「これが宇宙時代のさきがけか」と感慨深い。
今後の展開に注目しておこう。

## 【おまけの再掲載】 天文学と光害

この文章は[2003年に公開したブログ記事](https://baldanders.info/blog/000011/ "天文学と光害 -- 戯れ言++")の再掲載だ。
古い記事だが「光害問題」について考える際の参考になれば幸い。

{{< div-box type="markdown" >}}

- [回転サーチライト等禁止の法制化についての要望書](http://www.asj.or.jp/news/031212.html)
- [回転サーチライト等禁止の法制化、日本天文学会が環境省へ要望](http://www.astroarts.co.jp/news/2003/12/12nao689/index-j.shtml)

私の感じでは光害の問題は1980年代あたりから特に問題視されてきたように思います。
しかし当時は光害自体が世論を呼ぶ状態ではありませんでした。
逆に1980年代後半から脹れ上がるバブル景気とともに日本国各地における「ライトアップ運動」が非常に盛んになっていったのを憶えています。

日本で光害が問題視されるようになってきたのは，ようやく1990年代の終わり頃です。
それも天文学等の学術領域が受ける被害としてではなく，無節操な夜間照明による野生動物の生態系への影響すなわち「環境問題」としてクローズアップされてきています。

天文学にとって光害の何が問題かというと，夜間照明が空に向けられる「ライトアップ」の問題です。
従ってこれに対する解決策は照明による光を空に漏らさないようにすることです。
（街灯を含め）既にそのような照明器具は市場に出まわっています。
光を空に漏らさないということは，それだけエネルギー効率の良い照明であるとも言えます。

もちろん天文観測施設のある地域では条例や付近住民との取り決め等で光害を減らす努力が行われています。
しかし狭い日本列島ではもはや光害の影響を全く受けない場所など「ない」といっても過言ではないと思います。
たとえ近所の明かりを落としても，何百キロ先の都会の明かりが夜空を照らしていたのでは何にもなりません。

近年（2003年当時）では目に見える「光」以外にも電波による光害も問題視されるようになってきました。
一応国際協定で電波天文学用にリザーブされている電波領域というのはあるのですが，
それを破る人（？）がいるということや，商業サイドからの圧力により観測用の電波領域を確保するのが難しくなりつつあります。

電波天文観測の対象となる電波は非常に微弱なものです。
例えば今あなたが持っているケータイを月に置き，それを電波望遠鏡で受信してしまったら「針が飛ぶほどの大音量」になってしまいます。
電波天文観測というのはそれほどデリケートなものなのです。

光および電波による光害は今後も続くでしょう。
でも，ここに書いたようなことも頭の隅に入れていただけると嬉しいです 
{{< /div-box >}}

## ブックマーク

- [通信衛星群による天文観測への悪影響についての懸念表明 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2019/20190709-satellites.html) : よーやく日本からも声明を出すことにしたらしい。産業界に忖度（笑）して出さないのかと思ったぜ。でも，こういうのって「拙速を尊ぶ」なんだけどねぇ
- [アマゾンが3000機以上のブロードバンド通信衛星の打ち上げをFCCに申請  |  TechCrunch Japan](https://jp.techcrunch.com/2019/07/10/2019-07-10-amazon-seeks-fcc-approval-to-launch-over-three-thousand-broadband-satellites/)

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E3%81%86%E3%81%A1%E3%81%AE%E5%B1%85%E5%80%99%E3%81%8C%E4%B8%96%E7%95%8C%E3%82%92%E6%8E%8C%E6%8F%A1%E3%81%97%E3%81%A6%E3%81%84%E3%82%8B%EF%BC%81%EF%BC%88GA%E6%96%87%E5%BA%AB%EF%BC%89/dp/B06WD5SGXL?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B06WD5SGXL"><img src="https://images-fe.ssl-images-amazon.com/images/I/C1AChJPjnMS._SL160_.png" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%81%BE%E3%81%A8%E3%82%81%E8%B2%B7%E3%81%84-%E3%81%86%E3%81%A1%E3%81%AE%E5%B1%85%E5%80%99%E3%81%8C%E4%B8%96%E7%95%8C%E3%82%92%E6%8E%8C%E6%8F%A1%E3%81%97%E3%81%A6%E3%81%84%E3%82%8B%EF%BC%81%EF%BC%88GA%E6%96%87%E5%BA%AB%EF%BC%89/dp/B06WD5SGXL?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B06WD5SGXL">[まとめ買い] うちの居候が世界を掌握している！（GA文庫）</a></dt>
	<dd>七条 剛</dd>
	<dd>希望 つばめ (イラスト)</dd>
    <dd></dd>
    <dd>Ebook Bundle Kindle版</dd>
    <dd>ASIN: B06WD5SGXL</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">宇宙開発事業で世界を掌握する高校生の話（笑） 個人的にはルファさんには幸せになってほしい。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-06-15">2019-06-15</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
