+++
title = "Facebook はもう駄目か知らん"
date = "2019-03-23T18:04:07+09:00"
description = "Facabook のことは長い目で見て「付かず離れず」で行こうと思っていたが，やめます。これから撤退戦を開始します。"
image = "/images/attention/kitten.jpg"
tags = [ "facebook", "privacy", "surveillance-capitalism" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

かつて Google は[「完全なプライバシーは存在しない」と公言](http://www.thesmokinggun.com/documents/internet/google-complete-privacy-does-not-exist)し「[Google はプライバシーに敵対的](http://cloud.watch.impress.co.jp/epw/cda/infostand/2007/06/18/10531.html)」と大きな批判を浴びていた。
しかし「プライバシーに敵対的」な企業・組織・国家は Google だけではない。

## Facebook はユーザのプライバシーを蹂躙する

[Campbridge Analytica の件]({{< ref "/remark/2018/03/name-identification.md" >}})以来 Facebook は大きな批判にさらされているが，上述したように「プライバシーに敵対的」で「監視資本主義の走狗」たる企業なんて（規模の大小はともかく）いくらでもあるし，特定の企業・サービスを名指しで批判することは問題を矮小化させるだけだろうと控えてきたつもりである。

なので「[Facebookが数億人のパスワードを平文で保存していたと認める](https://techcrunch.com/2019/03/21/facebook-plaintext-passwords/ "Facebookが数億人のパスワードを平文で保存していたと認める  |  TechCrunch Japan")」という記事のタイトルを見たときも「またやらしてるよぱすわーどへんこうしなきゃ」くらいにしか思わなかったのだが，どうやらもっととんでもない話のようである。

- [Facebook Stored Hundreds of Millions of User Passwords in Plain Text for Years —  Krebs on Security](https://krebsonsecurity.com/2019/03/facebook-stored-hundreds-of-millions-of-user-passwords-in-plain-text-for-years/)
- [Facebookが一部ユーザーのパスワードを平文記録していた問題についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/03/22/061503)

インターネットにおける認証の多くはサービスとユーザの間で「秘密」を共有・交換することで成立する。
パスワードは，それが互いの間で「秘密」である限り認証として有効である。

真っ当なサービス・プロバイダはパスワードという「秘密」を秘密として維持するために苦労している。
理想はサービス・プロバイダの誰もユーザの設定したパスワードを取得できないようにすることである。
故にパスワードが平文のまま保持されている状態はセキュリティ上の重大な瑕疵（defect）となる。

しかし今回の件は瑕疵ではない。
サービス・プロバイダである Facebook によるプライバシーの蹂躙である。
ユーザとの「秘密」を組織内部で共有していたのだから。

{{< fig-quote title="Facebook Stored Hundreds of Millions of User Passwords in Plain Text for Years" link="https://krebsonsecurity.com/2019/03/facebook-stored-hundreds-of-millions-of-user-passwords-in-plain-text-for-years/" lang="en" >}}
<q>My Facebook insider said access logs showed some 2,000 engineers or developers made approximately nine million internal queries for data elements that contained plain text user passwords.</q>
{{< /fig-quote >}}

更に Facebook 側は

{{< fig-quote title="Facebook Stored Hundreds of Millions of User Passwords in Plain Text for Years" link="https://krebsonsecurity.com/2019/03/facebook-stored-hundreds-of-millions-of-user-passwords-in-plain-text-for-years/" lang="en" >}}
<q>Renfro said the company planned to alert affected Facebook users, but that no password resets would be required.</q>
{{< /fig-quote >}}

などと言ったようだ。
まぁ当然だよね。
Facebook 側がいつでもパスワードを含むユーザの「秘密」を引き出して利用できるならパスワードの変更など「やっても無駄」なのだから。

ちなみに今回の件について Facebook の TL に「客をなめるにもほどがある」と書いたら「客じゃなくて商品ですから」とツッコミをいただきました。
そのとーり `orz`

## Facebook は「監視資本主義」へ邁進するか

そういえば Facebook の CEO である Mark Zuckerberg 氏による[プライバシーに関する記事](https://www.facebook.com/notes/mark-zuckerberg/a-privacy-focused-vision-for-social-networking/10156700570096634/)が批判されたのはつい先週のことである。
そのうちのひとつがこれ。

- [A New Privacy Constitution for Facebook – OneZero](https://onezero.medium.com/a-new-privacy-constitution-for-facebook-a7106998f904?gi=72191a38d1fa)
- [Judging Facebook's Privacy Shift - Schneier on Security](https://www.schneier.com/blog/archives/2019/03/judging_faceboo.html)

{{< fig-quote title="Judging Facebook's Privacy Shift" link="https://www.schneier.com/blog/archives/2019/03/judging_faceboo.html" lang="en" >}}
<q>Facebook users have limited control over how their data is shared with other Facebook users and almost no control over how it is shared with Facebook's advertisers, which are the company's real customers.</q>
{{< /fig-quote >}}

{{< fig-quote title="Judging Facebook's Privacy Shift" link="https://www.schneier.com/blog/archives/2019/03/judging_faceboo.html" lang="en" >}}
<q>User privacy is traditionally against Facebook's core business interests. Advertising is its business model, and targeted ads sell better and more profitably -- and that requires users to engage with the platform as much as possible. Increased pressure on Facebook to manage propaganda and hate speech could easily lead to more surveillance. But there is pressure in the other direction as well, as users equate privacy with increased control over how they present themselves on the platform.</q>
{{< /fig-quote >}}

最後に記事はこう結んでいる。

{{< fig-quote title="Judging Facebook's Privacy Shift" link="https://www.schneier.com/blog/archives/2019/03/judging_faceboo.html" lang="en" >}}
<q>Facebook talks about community and bringing people together. These are admirable goals, and there's plenty of value (and profit) in having a sustainable platform for connecting people. But as long as the most important measure of success is short-term profit, doing things that help strengthen communities will fall by the wayside. Surveillance, which allows individually targeted advertising, will be prioritized over user privacy. Outrage, which drives engagement, will be prioritized over feelings of belonging. And corporate secrecy, which allows Facebook to evade both regulators and its users, will be prioritized over societal oversight. If Facebook now truly believes that these latter options are critical to its long-term success as a company, we welcome the changes that are forthcoming.</q>
{{< /fig-quote >}}

私も最後のこの言葉には同意しかけたが，その直後にこれである。

- [Facebook、ニュージーランドモスク銃撃の動画150万件を削除 - CNET Japan](https://japan.cnet.com/article/35134319/)

{{< div-gen >}}
<blockquote class="twitter-tweet" data-lang="ja"><p lang="en" dir="ltr">In the first 24 hours we removed 1.5 million videos of the attack globally, of which over 1.2 million were blocked at upload...</p>&mdash; Facebook Newsroom (@fbnewsroom) <a href="https://twitter.com/fbnewsroom/status/1107117981358682112?ref_src=twsrc%5Etfw">2019年3月17日</a></blockquote>
{{< /div-gen >}}

いやいや。
アップロードのタイミングで{{< ruby "ブロック" >}}検閲{{< /ruby >}}したとかヤバいだろ，それ。

この件で「9.11 以後」のアメリカの状況を思い出して「まずいことになったなぁ」と思ってたら冒頭の話につながるわけだ。

## Facebook からの撤退戦を開始します

Facabook のことは長い目で見て「付かず離れず」で行こうと思っていたが，やめます。
多分 Facebook はもう駄目だと思う。
なのでこれから撤退戦を開始します。

### Instagram から（最終的に）完全撤退します

これについては実は [Flickr にお金を払った]({{< relref "./flickr-has-not-turned-to-surveillance-capitalism.md" >}})あたりから考えていたのだが，前倒しすることにする。
まぁ[2019年2月までのバックアップは取ってある](https://photo.baldanders.info/instagram/ "Photos in Instagram | photo.Baldanders.info")ので無問題。

まずはスマホから Instagram アプリをサクっと削除。

参照は出来るよう当面はアカウントを残すけど，以後は放置プレイで。
代わりに [Flickr アプリ](https://play.google.com/store/apps/details?id=com.yahoo.mobile.client.android.flickr "Flickr - Google Play")を入れた。
更に IFTTT で Flickr から Twitter にも自動で転載できるようにした。

とりあえず，こんなところかな。

- [もう Instagram はいいかな]({{< ref "/remark/2019/12/no-more-instagram.md" >}})

### Baldanders.info の Facebook ページは閉鎖します

つか，閉鎖しました。
2週間後には完全に削除される予定（何故か勝手に猶予期間が付けられた）。

影響はほとんどないだろう，多分。

### スマホから Facebook アプリを削除します

スマホに Facebook アプリを入れてるとつい見ちゃうので削除した。
まぁ PC でブラウザから見れるし。
でも，今は PC の前に毎日座ってるわけではないので更新頻度は落ちるだろう。
くだらない戯れ言が「友人」の TL から減るということで（笑）

脊髄反射の戯れ言を書く場所がなくなるので [note.mu](https://note.mu/spiegel) を復活させようか考え中。
Twitter は広告がウザすぎて常用する気になれない。

### 連絡は [Signal] で

Facebook がきっかけで知り合った「友人」も多いので，いざというときの連絡手段として Messenger アプリは残しておこうと思ったけど，たまにでも PC で見るならスマホのアプリは要らんか。
現在，私のメインのメッセージング・アプリは [Signal] なので今後はそちらでどうぞ。
[Signal] は SMS と連動できるので，私のケータイ番号を知ってる人ならアクセス可能だと思われ。

実家に帰ったので，家族に連絡がとれず「ある日，部屋から腐乱死体で発見される」心配はしなくてよくなった（腐ったら後始末が大変だものね）ので連絡手段についてはあまり気にしていない。

### 情報収集用に [Feedly] を再開しました

実は Facebook 利用のメインは情報収集だったのだが，今後は覚束なくなるので，数年ぶりに [Feedly] を復活させた。
しかも以前のアカウントではなく新規に取得し直して真っさらから構築した。

Facebook の TL と違って [Feedly] では「誤配」の確率が減るが，そこは誤配だらけの Twitter 等を組み合わせて何とかなるか？

### これで BAN されても安心かな

結構 Facebook に依存してたんだねぇ，私。
ここまでやっておけば万一 BAN されても安心だろう。
GitHub に BAN されたらメガっさ困るが Facebook なら大して未練はない（Facebook を通じて知り合った友人達にはゴメンなさいだけど）。

これこそまさに「断捨離」だな（笑）

## ブックマーク

- [FacebookにFTCが数十億ドルの罰金を課す可能性  |  TechCrunch Japan](https://techcrunch.com/2019/02/14/facebook-ftc-fine/)
- [MIT Tech Review: 「プライバシー重視」宣言、突然の大転換に潜むフェイスブックの欺瞞](https://www.technologyreview.jp/s/129792/zuckerbergs-new-privacy-essay-shows-why-facebook-needs-to-be-broken-up/)
- [MIT Tech Review: フェイスブックの個人情報「裏提供」問題、米検察が捜査へ](https://www.technologyreview.jp/nl/facebooks-data-deals-are-now-under-criminal-investigation/)
- [ニュージーランドの大量殺人事件に関するEFFの考え：恐怖が導く検閲へ道 – P2Pとかその辺のお話R](https://p2ptk.org/freedom-of-speech/1702)
- [Change your Facebook password. And don't try to remember it. - F-Secure Blog](https://blog.f-secure.com/change-your-facebook-password-and-dont-try-to-remember-it/)
- [MIT Tech Review: フェイスブックのデータがまた漏洩、5億件以上がクラウドに放置](https://www.technologyreview.jp/nl/more-than-half-a-billion-facebook-records-were-left-exposed-on-the-public-internet/)
- [サードパーティ取得のFacebookデータが公開状態だった件についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/04/09/063000)
- [Facebookは「利用解除」したユーザーも追跡し続けている - CNET Japan](https://japan.cnet.com/article/35135675/)
- [Facebookがパスワードが漏れたInstagramユーザー数を「数百万人」に訂正  |  TechCrunch Japan](https://techcrunch.com/2019/04/18/instagram-password-leak-millions/)
- [Facebookを“マシ”にするためにFTCは何をすべきか – P2Pとかその辺のお話R](https://p2ptk.org/privacy/2028)

[Signal]: https://signal.org/
[Feedly]: https://feedly.com/ "Feedly. Read more, know more."

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4588276484" %}} <!-- 闇をひらく光 -->
