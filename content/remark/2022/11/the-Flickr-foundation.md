+++
title = "The Flickr Foundation 100年の計"
date =  "2022-11-29T21:26:19+09:00"
description = "まずはドキュメントを読み込まないとなぁ。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "flickr", "intellectual-property", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[先日の記事]({{< relref "./new-spark-of-competition-between-platforms-and-federations.md" >}} "「プラットフォームとフェデレーションとの競争」")で Flickr が [ActivityPub] に対応するかも，という話を書いたが，実際に [Mastodon 連合](https://sfba.social/@d0n/109422647995225732)と [Twitter](https://twitter.com/DonMacAskill/status/1597256480519966720) で投票を行っている。
投票の内容は同じで，いずれの場所でも [ActivityPub] に対応することに対してはポジティブな反応が多いが，どのように導入していくかについて傾向が分かれているのが面白い。

この辺の話を TechCrunch が取り上げているのだが

- [Flickr weighs support for ActivityPub, the social protocol powering Twitter alternative Mastodon • TechCrunch](https://techcrunch.com/2022/11/28/flickr-weighs-support-for-activitypub-the-social-protocol-powering-twitter-alternative-mastodon/)

この記事の最後の方で [Flickr Foundation] についてサラリと言及している。
元ネタの tweet がこれ。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">We have! <a href="https://twitter.com/Flickr?ref_src=twsrc%5Etfw">@Flickr</a> is healthy and growing again. And we’ve even set up a non-profit to preserve humanity’s photography for 100+ years <a href="https://t.co/rFYdVyR2A3">flickr.org</a> (preventing something like what may be happening at Twitter if Flickr falls on hard times again)</p>&mdash; Don MacAskill (@DonMacAskill) <a href="https://twitter.com/DonMacAskill/status/1594001639219769344?ref_src=twsrc%5Etfw">November 19, 2022</a></blockquote>
{{< /fig-gen >}}

いや，こっちの話のほうが重要ぢゃん。
“preventing something like what may be happening at Twitter if Flickr falls on hard times again” の部分でホロリとしちゃったよ（笑） 備えは大事。

[Flickr Foundation] って2021年から始動してるんだね。
件のサイトには

{{< fig-quote type="markdown" title="The Flickr Foundation" link="https://www.flickr.org/" lang="en" >}}
We believe the establishment of a non-profit Flickr Foundation will combine with Flickr to properly preserve and care for the Flickr Commons archive, support Commons members to collaborate in a true 21st-century Commons, and plan for the very long-term health and longevity of the entire Flickr collection. We’re also in the early stages of imagining other educational and curatorial initiatives to highlight and share the power of photography for decades to come.
{{< /fig-quote >}}

とある。
「The Flickr Foundation 100年の計」って感じ。

デジタルデータは失われやすい。
ある「表現」が公有になるまで（70年！）待ってたら消失して何処にも見当たらない，なんてなこともありうる話だ。

今年に入ってからも TechCrunch Japan 終了の際に，それまでの日本語記事を（翻訳記事以外も含めて）[あっさり捨て去った]({{< ref "/remark/2022/02/the-nation-of-amnesia.md" >}} "記憶喪失の国")ことなど，まだ記憶に新しいだろう。
これに比べれば，[公開が終了した記事を著者に渡して後を委ねた cakes](https://yamdas.hatenablog.com/entry/20221128/cakes_data_journalism "cakesに掲載された「ネット×ジャーナリズムの歴史とその最新潮流としてのデータジャーナリズム」をサルベージした - YAMDAS現更新履歴") とか良心的なほうである。

あるいは，どこぞの GitHub みたいに[北極に千年封印する]({{< ref "/remark/2020/07/github-archive-program-2020.md" >}} "私的コード黒歴史が北極に千年封印される")なんてな手もあるかもしれない。
でも，それって封印が解かれるまで照会も再利用もできないってことだよね。
常に変容するプログラム・コードのある時点を封印することに，どれほどの意味があるか知らんけど。

デジタルな「表現」が，倫理・道徳・法・政治・市場・社会構造を超えて常に「使用」または「利用」できる状態を維持しつつ（100年単位で）保持し続けるというのは，かなり大変であろうことは想像に難くない。

とりあえず [Flickr Foundation] は2023年までの3年計画で色々と準備をするみたいなので，それらを注視するところから始めますかね。
まずはドキュメントを読み込まないとなぁ。

## ブックマーク

- [Research Report - Flickr Commons Revitalization - Google ドキュメント](https://docs.google.com/document/d/19eHlYDCJS2F4KH-Eml8mW9aVqP5WVOaKN_-5XFW4-wc)
- [Strategy 2021-2023 - Flickr Commons Revitalization  - Google ドキュメント](https://docs.google.com/document/d/1zm6HRq6ryDP1RKNHckUsGBwsn2sZaQsP_woxo4HTAws)
- [Revitalising the Flickr Commons program - CC Summit - Google スライド](https://docs.google.com/presentation/d/1t_vousKMZrGXx3auoVVQkflibqY5En9c0bL1ihgK9bc/edit) : 2021 Creative Commons Summit

{{< fig-youtube id="rsqUv3JU32g" title="Coming of age in the Digital Decade: A 17 year old software company and what it means to look ahead" >}}

[Flickr Foundation]: https://www.flickr.org/ "The Flickr Foundation"
[ActivityPub]: https://www.w3.org/TR/activitypub/

## 参考文献

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->
