+++
title = "Firefox の DoH は無効にすべきか（もしくは水売りと水道局）"
date =  "2019-09-15T10:31:39+09:00"
description = "これは ISPA の言いがかりみたいな話ではなく，現在のインターネットの構造，ひいては社会システムに連動する問題提起だ。"
image = "/images/attention/kitten.jpg"
tags = [ "internet", "privacy", "firefox", "grigori" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Mozilla が DoH (DNS-over-HTTPS) を Firefox に正式に実装するとアナウンスがあったが

- [What’s next in making Encrypted DNS-over-HTTPS the Default - Future Releases](https://blog.mozilla.org/futurereleases/2019/09/06/whats-next-in-making-dns-over-https-the-default/)
- [Mozilla Firefox to begin slow rollout of DNS-over-HTTPS by default at the end of the month • The Register](https://www.theregister.co.uk/2019/09/09/mozilla_firefox_dns/)
- [FirefoxがDNSとの通信を暗号化する「DNS over HTTPS(DoH)」を正式に実装すると発表 - GIGAZINE](https://gigazine.net/news/20190910-mozilla-firefox-dns-over-https/)

これに懸念を呈する面白い記事がある。

- [ungleich blog - Turn off DoH, Firefox. Now.](https://ungleich.ch/en-us/cms/blog/2019/09/11/turn-off-doh-firefox/) （[邦訳版](https://okuranagaimo.blogspot.com/2019/09/firefoxdoh.html "ブログ: Firefoxよ、DoHをオフにしろ、今すぐ")）

これは [ISPA の言いがかり]みたいな話ではなく，現在のインターネットの構造，ひいては社会システムに連動する問題提起だ。

そもそも件のブログ記事は DoH 自体には反対していない。

{{< fig-quote type="md" title="Turn off DoH, Firefox. Now." link="https://ungleich.ch/en-us/cms/blog/2019/09/11/turn-off-doh-firefox/" lang="en" >}}
{{< quote >}}DoH and DoT (DNS over TLS) are in general good technologies as they add encryption to an important process of daily life. However the approach Mozilla takes is simply wrong. The correct way would be to standardise DoH and DoT and add support into it into automatic address configurations and operating systems. Not in applications!{{< /quote >}}
{{< /fig-quote >}}

そして，問題は DoH が組み込まれる Firefox がアプリケーションに過ぎないこと， DoH のホストとして米国の Cloudflare を利用していることだと主張している。

{{< fig-quote type="md" title="Turn off DoH, Firefox. Now." link="https://ungleich.ch/en-us/cms/blog/2019/09/11/turn-off-doh-firefox/" lang="en" >}}
{{< quote >}}It means people outside the US can now be fully tracked by US government{{< /quote >}}
{{< /fig-quote >}}

{{< fig-quote type="md" title="Turn off DoH, Firefox. Now." link="https://ungleich.ch/en-us/cms/blog/2019/09/11/turn-off-doh-firefox/" lang="en" >}}
{{< quote >}}whether you trust Cloudflare or not, you'll end up directly supporting centralisation by using DoH in Firefox. Centralisation makes us depend on one big player, which results in fewer choices and less innovation. Centralisation affects everybody by creating a dangerous power and resource imbalance between the center and the rest.{{< /quote >}}
{{< /fig-quote >}}

これで思い出すのが，かつて言われた「水のような音楽」というやつだ。

もともと「水のような音楽」は DRM (Digital Right Management) で貞操帯のごとくがんじ搦めにされたコンテンツに対するアンチテーゼのようなものだった。

{{< fig-quote type="md" title="EMIのDRMとの決別は「水のような音楽」への大きな一歩か - YAMDAS現更新履歴" link="https://yamdas.hatenablog.com/entry/20070404/nodrm" >}}
{{< quote >}}もう一つは既存のサービスを残したままで、新しい DRM フリーの高品質サービスが提供されること。手軽にアクセスできる安価な水道水がある一方で、より質を重視する人にはそれより値がはるミネラルウォーターも提供されるという「水のような音楽」モデルじゃないですか。{{< /quote >}}
{{< /fig-quote >}}

しかし，音楽にしろ映像にしろ，無料または定額制のストリーミングサービスが一般化し「水のような音楽」が合法的に利用できるようになって分かったことは
**「水道水はミネラルウォーターより統制しやすい」**
といういうことだった。

言ってみれば「水」を売るのが「水売り」から「水道局」に代わっただけで，むしろ「水道局」のほうが中央集権的で統制に向いているのは明らかである。
いずれにしろ利用者に自由なんてものはないのだ。

これは音楽や映像といったコンテンツに限らず電子メール等のメッセージング・サービスや（マイクロ）ブログにも言えることで，今回の Firefox への DoH 実装は **インターネットによる統制** を更に更に推し進めていくだろう，というわけだ。

そういう意味でも [ISPA の言いがかり]はホンマに言いがかりなんだなぁ，と思ってしまう。

ちなみに件のブログでは

{{< fig-quote type="md" title="Turn off DoH, Firefox. Now." link="https://ungleich.ch/en-us/cms/blog/2019/09/11/turn-off-doh-firefox/" lang="en" >}}
{{< quote >}}It is clear what Mozilla needs to do: Mozilla can and should revert the change and allow users to easily opt-in. And to select or enter the DoH provider instead of defaulting to Cloudflare. Also Mozilla can take real responsibility and work together with the Internet community and create RFCs to make DHCPv4, DHCPv6 and Router Advertisements support DNS URLs instead of just IP addresses.{{< /quote >}}
{{< /fig-quote >}}

と書かれていて，これに関しては激しく同意する。


そういえば最近 Firefox 69 がリリースされて

- [Today’s Firefox Blocks Third-Party Tracking Cookies and Cryptomining by Default - The Mozilla Blog](https://blog.mozilla.org/blog/2019/09/03/todays-firefox-blocks-third-party-tracking-cookies-and-cryptomining-by-default/)

一瞬喜んだが，私のケータイは 68.x から一向にアップグレードされる気配がない。

はっきり言おう。

{{< div-gen type="md" class="center" >}}
**Mozilla がユーザのプライバシーを重視しているというのは嘘っぱちである**
{{< /div-gen >}}

本当に Mozilla がユーザのプライバシーを重視しているというのなら，検索サービスの既定を DuckDuckGo にすべき。
話はそれからだ。

## ブックマーク

- [Blocking Firefox DoH with Bind - SANS Internet Storm Center](https://isc.sans.edu/forums/diary/Blocking+Firefox+DoH+with+Bind/25316/)

- [監視をコントロールする](https://baldanders.info/blog/000490/)

[ISPA の言いがかり]: https://www.ispa.org.uk/ispa-announces-finalists-for-2019-internet-heroes-and-villains-trump-and-mozilla-lead-the-way-as-villain-nominees/ "ISPA announces finalists for 2019 Internet Heroes and Villains: Trump and Mozilla lead the way as Villain nominees » Press Releases | The Internet Service Providers Association"

## 参考文献

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%B6%85%E7%9B%A3%E8%A6%96%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC-ebook/dp/B01MZGVHOA?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01MZGVHOA"><img src="https://images-fe.ssl-images-amazon.com/images/I/51T6PBdGbyL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%B6%85%E7%9B%A3%E8%A6%96%E7%A4%BE%E4%BC%9A-%E3%83%96%E3%83%AB%E3%83%BC%E3%82%B9%E3%83%BB%E3%82%B7%E3%83%A5%E3%83%8A%E3%82%A4%E3%82%A2%E3%83%BC-ebook/dp/B01MZGVHOA?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01MZGVHOA">超監視社会</a></dt>
    <dd>ブルース・シュナイアー</dd>
    <dd>池村 千秋 (翻訳)</dd>
    <dd>草思社 2016-12-13 (Release 2017-02-03)</dd>
    <dd>Kindle版</dd>
    <dd>B01MZGVHOA (ASIN)</dd>
  </dl>
  <p class="description">実は積ん読のまま読んでない。そろそろちゃんと最後まで読まないと。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-23">2019-03-23</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%87%E3%82%B8%E3%82%BF%E3%83%AB%E9%9F%B3%E6%A5%BD%E3%81%AE%E8%A1%8C%E6%96%B9-David-Kusek/dp/4798110035?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4798110035"><img src="https://images-fe.ssl-images-amazon.com/images/I/51HINm31mtL._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%87%E3%82%B8%E3%82%BF%E3%83%AB%E9%9F%B3%E6%A5%BD%E3%81%AE%E8%A1%8C%E6%96%B9-David-Kusek/dp/4798110035?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4798110035">デジタル音楽の行方</a></dt>
    <dd>David Kusek, Gerd Leonhard, yomoyomo</dd>
    <dd>yomoyomo (翻訳), 津田 大介 (翻訳)</dd>
    <dd>翔泳社 2005-12-06</dd>
    <dd>単行本</dd>
    <dd>4798110035 (ASIN), 9784798110035 (EAN)</dd>
  </dl>
  <p class="description">読んでない。絶賛絶版中？ もはや歴史的価値の有無さえ不明。「デジタル」とタイトルに冠されているにも関わらず Kindle 化すらされてないことに出版側のやる気を感じる（笑）</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-07-30">2017-07-30</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
