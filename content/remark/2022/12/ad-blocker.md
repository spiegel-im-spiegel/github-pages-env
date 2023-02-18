+++
title = "米国 FBI は広告ブロッカーを推奨している？"
date =  "2022-12-29T10:34:27+09:00"
description = "既に Web 広告は曲がり角に入ってるんだろう。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "media", "search", "market" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

個人的には Web ページ上の広告はあまり気にしないのだが（コンテンツをぶった切るような TL 広告や YouTube 広告はムカつくけど），どうやら最近は検索サービスに於いて有名ブランドの広告になりすます詐欺広告が横行しているらしい。

- [Internet Crime Complaint Center (IC3) | Cyber Criminals Impersonating Brands Using Search Engine Advertisement Services to Defraud Users](https://www.ic3.gov/Media/Y2022/PSA221221)

これによると

{{< fig-quote type="markdown" title="Cyber Criminals Impersonating Brands Using Search Engine Advertisement Services to Defraud Users" link="https://www.ic3.gov/Media/Y2022/PSA221221" lang="en" >}}
In instances where a user is searching for a program to download, the fraudulent webpage has a link to download software that is actually malware. The download page looks legitimate and the download itself is named after the program the user intended to download.

These advertisements have also been used to impersonate websites involved in finances, particularly cryptocurrency exchange platforms. These malicious sites appear to be real exchange platforms and prompt users to enter login credentials and financial information, giving criminal actors access to steal funds.
{{< /fig-quote >}}

てぇことらしい。

というわけで，久しぶりに Google 検索を試してみる。
たとえば Google 検索で「暗号通貨」を検索すると，まずはこんな感じに広告がずらっと並ぶ。

{{< fig-img src="./google-search-1.png" title="Google 検索（1）" link="./google-search-1.png" width="1061" >}}

そこからスクロールダウンするとニュースが表示され

{{< fig-img src="./google-search-2.png" title="Google 検索（2）" link="./google-search-2.png" width="1061" >}}

そこからさらにスクロールダウンでようやく本来の検索結果が表示される。

{{< fig-img src="./google-search-3.png" title="Google 検索（3）" link="./google-search-3.png" width="1061" >}}

えっ，待って。
皆こんなク◯みたいな UX でいいの？ つまり Google 検索を利用してる人って広告を検索してるってこと？ これって詐欺広告以前の問題じゃないの？

最初に挙げた FBI の記事では一応，個人向けの対策も挙げている。曰く

{{< fig-quote type="markdown" title="Cyber Criminals Impersonating Brands Using Search Engine Advertisement Services to Defraud Users" link="https://www.ic3.gov/Media/Y2022/PSA221221" lang="en" >}}
- Before clicking on an advertisement, check the URL to make sure the site is authentic. A malicious domain name may be similar to the intended URL but with typos or a misplaced letter.
- Rather than search for a business or financial institution, type the business’s URL into an internet browser’s address bar to access the official website directly.
- Use an ad blocking extension when performing internet searches. Most internet browsers allow a user to add extensions, including extensions that block advertisements. These ad blockers can be turned on and off within a browser to permit advertisements on certain websites while blocking advertisements on others.
{{< /fig-quote >}}

とのこと。

最初の2つは Phishing 回避策としては定番だね。
でも今時の Phishing は巧妙でドメイン名をパッと見しただけでは判別しづらいものも多い。
あと，検索サービスで探してるのに “type the business’s URL into an internet browser’s address bar to access the official website directly” っていうアドバイスはないと思う。
その URL が分からなくて検索してるのに。

で，最後のが広告ブロッカーを使えというアドバイスなんだけど，検索サービスに限るなら，こんなク◯みたいな UX を強制するサービスは使わないのが吉なんじゃないのかなぁ。

個人的には [DuckDuckGo] を推させてもらう。
たとえば先程の「暗号通貨」を [DuckDuckGo] で検索すれば

{{< fig-img src="./duckduckgo.png" title="DuckDuckGo" link="./duckduckgo.png" width="1282" >}}

という感じに多少はまともな表示になる。
少なくとも広告がページトップにドカンと表示されることはない。

昔の [DuckDuckGo] は特に日本語サイトの検索がイマイチだったが，今は全く遜色なく使えている。
個人的には情報を探す際に90%以上は [DuckDuckGo] で事足りる。
よほどのことがないと Google は使わない感じ。

とにかく Web 広告はイヤって方については GIZMODE が [uBlock Origin] を[推している](https://www.gizmodo.jp/2022/12/google-bing-fbi-ad-blocker-scam-ads.html "FBIがみんなに広告ブロッカーを使って欲しい理由 | ギズモード・ジャパン")ようなの試してみてもいいだろう。

まとめとしては

- 検索サービスは [DuckDuckGo] など，広告のないサービスを利用する
- Web ページ上の広告には（真贋に関わらず）近づかない
- どうしても見えてる広告を触っちゃう，あるいは広告が表示されること自体がウザいと思う方は [uBlock Origin] 等の広告ブロッカーを検討する

といったところだろうか。
あっ，もちろんクロスサイト cookie やトラッキング・コード，マイニング・コードはしっかりブロックした上でね。

{{< fig-img src="./settings-for-firefox.png" title="Firefox の設定" link="./settings-for-firefox.png" width="999" >}}

EFF も2019年に

- [Adblocking: How About Nah? | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2019/07/adblocking-how-about-nah)
- [広告ブロッカーは「嫌ならどうする？」の表明である | p2ptk[.]org](https://p2ptk.org/monopoly/2668)

みたいな記事を出してたりするし，既に Web 広告は曲がり角に入ってるんだろう。


## ブックマーク

- [メディアは（常に）スパムか？ « マガジン航[kɔː]](https://magazine-k.jp/2016/01/25/spam-and-media/)
- [Until further notice, think twice before using Google to download software | Ars Technica](https://arstechnica.com/information-technology/2023/02/until-further-notice-think-twice-before-using-google-to-download-software/)
  - [Malware Delivered through Google Search - Schneier on Security](https://www.schneier.com/blog/archives/2023/02/malware-delivered-through-google-search.html)

[DuckDuckGo]: https://duckduckgo.com/
[uBlock Origin]: https://ublockorigin.com/jp "uBlock Origin – フリーかつオープンソースの広告ブロッカー。"

## 参考図書

{{% review-paapi "430924744X" %}} <!-- スパム -->
