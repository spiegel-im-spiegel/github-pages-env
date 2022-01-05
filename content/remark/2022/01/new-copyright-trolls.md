+++
title = "CC Licenses を悪用した新手の著作権トロルについて"
date =  "2022-01-05T21:23:50+09:00"
description = "このリスクを事前に回避したいのであれば「クレジットを明記する」に限る。"
image = "/images/attention/kitten.jpg"
tags = [ "copyright", "creative-commons", "risk", "code", "law" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いつものように yomoyomo さんの記事から：

- [Beware the Copyleft Trolls | by Chip Stewart | OneZero](https://onezero.medium.com/beware-the-copyleft-trolls-a8b85c66b7eb)
- [Rise of the Copyleft Trolls: When Photographers Sue After Creative Commons Licenses Go Awry by Daxton Stewart :: SSRN](https://papers.ssrn.com/sol3/papers.cfm?abstract_id=3844180)
- [著作権トロールの新種としての「コピーレフト・トロール」 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20220105/copyleft-trolls)

この記事に Twitter で脊髄反射したときは内容を誤解してたので，改めて簡単に紹介しておく。

現在，最新の CC Licenses 4.0 では，ライセンスの被許諾者が条件に違反した場合の措置について以下のように書いている。

{{< fig-quote title="クリエイティブ・コモンズ (Creative Commons) — 表示 4.0 国際 — CC BY 4.0" link="https://creativecommons.org/licenses/by/4.0/legalcode.ja" >}}
<p id="s6"><strong>第6条　期間および終了</strong></p>
<ol type="a">
<li id="s6a">本パブリック・ライセンスは、ここでライセンスされた著作権およびそれに類する権利が有効な期間、適用されます。しかし、もしあなたが本パブリック・ライセンスに違反すると、本パブリック・ライセンスに定めるあなたの権利は自動的に終了します。</li>
<li id="s6b">
<p>ライセンス対象物をあなたが利用する権利が<a href="#s6a">第6条(a)</a>の事由により終了した場合でも：</p>
<ol>
<li id="s6b1">あなたが違反を発見してから30日以内に違反を是正した場合に限り、違反を是正したその日に、自動的に復活します。または、</li>
<li id="s6b2">許諾者により権利の復活を明示された場合に、復活します。</li>
</ol>
誤解を避けるために記すと、<a href="#s6b">本第6条(b)</a>は、許諾者が、あなたの本パブリック・ライセンスに関する違反に対する救済を求めるために有するであろういかなる権利にも影響を及ぼしません。</li>
</ol>
{{< /fig-quote >}}

この30日の猶予期間は 4.0 になってから付加されたもので，それより前のバージョンには書かれていない。
たとえば，写真共有サービスの [Flickr](https://www.flickr.com/) では公開した写真に CC Licenses を付与できるが，そのバージョンは 2.0 であり “Termination” の条項には以下のように書かれているのみだ。

{{< fig-quote title="Creative Commons Legal Code -- Attribution 2.0" link="https://creativecommons.org/licenses/by/2.0/legalcode" lang="en" >}}
<p><strong>7. Termination</strong> </p>
<ol type="a">
<li>
This License and the rights granted hereunder will terminate automatically upon any breach by You of the terms of this License. Individuals or entities who have received Derivative Works or Collective Works from You under this License, however, will not have their licenses terminated provided such individuals or entities remain in full compliance with those licenses. Sections 1, 2, 5, 6, 7, and 8 will survive any termination of this License.
</li>
</ol>
{{< /fig-quote >}}

いわゆる「著作権トロル」はバージョン間のこの差異を悪用して善良でうっかりな利用者から金を巻き上げているらしい。
具体的には

{{< fig-quote type="markdown" title="Beware the Copyleft Trolls" link="https://onezero.medium.com/beware-the-copyleft-trolls-a8b85c66b7eb" lang="en" >}}
I became interested in this topic because it happened to me — or at least, to people I try to help, the students in the non-profit online news publication here at my university. A student posted a stock image from a German photographer named Marco Verch that was under a CC-BY (version 2.0) license. They used it as a thumbnail to link to a news story without the attribution. Months later, they got a demand letter requesting $750 to retroactively license the photo, ending with the line, “failure to resolve this matter of unlicensed use within 21 days will result in escalation to one of our partner attorneys for legal proceedings.”

I checked and, indeed, Verch had filed dozens of lawsuits in U.S. federal courts in recent years, including 41 cases in 2019 and 2020 alone. He employed Pixsy, a company that aids photographers in extracting cash from users who posted their photos online. Pixsy, in the demand letter, noted that taking the photo down or correcting the attribution wasn’t enough: “Removal of the image from your website does not resolve the period of unlicensed use, and it remains that our client be compensated for the previous use of their work.”
{{< /fig-quote >}}

という感じ。
訴えられてからクレジットを追記したり該当の写真を削除しても「掲載した期間の損害分を払いやがれ！」と迫るわけだ。

ただ，実際に裁判まで行っても賠償金を払うとは限らないようで

{{< fig-quote type="markdown" title="Beware the Copyleft Trolls" link="https://onezero.medium.com/beware-the-copyleft-trolls-a8b85c66b7eb" lang="en" >}}
Courts are also suggesting that Philpot doesn’t suffer economic harm because his photos have no value outside of litigation; he once testified that he had made “tens of thousands of dollars” from his Willie Nelson photo, but admitted that it was entirely from “settlements he had extracted” for failing to attribute under the CC-BY license. As Western District of Texas Judge Andrew Austin noted in a 2019 case, Philpot “is more in the business of litigation (or threatening litigation) than selling his product or licensing his photograph to third parties” who “seems to want to use the courts as a blunt object with which to coerce nuisance value settlements from unsuspecting parties.”
{{< /fig-quote >}}

と書かれている。
fair use として回避できるものもあるということだ。
“his photos have no value outside of litigation” の部分でバカウケしてしまったよ。

{{< fig-quote type="markdown" title="Beware the Copyleft Trolls" link="https://onezero.medium.com/beware-the-copyleft-trolls-a8b85c66b7eb" lang="en" >}}
In the paper, I look at three issues in particular: (1) expansive fair use arguments courts are entertaining to make copyright litigation harder for copyleft trolls like Philpot; (2) the limited damages courts are awarding in these cases; and (3) the general distaste they express for this kind of litigation. In short, courts are finding “transformative purposes” by secondary users (instead of “transformative uses”) under the first factor of the fair use analysis, and they are finding little to no economic harm under the fourth factor, expressing skepticism that lack of attribution is the equivalent of financial loss, thus requiring Philpot and his ilk to make a more robust showing on the record of actual economic harm outside of nuisance lawsuits. Courts are also pushing back at automatic awards of attorney fees when Philpot has won his infringement lawsuits, finding enough merit in defendants’ fair use claims that they do not want to deter such arguments in the future.
{{< /fig-quote >}}

このリスクを事前に回避したいのであれば「クレジットを明記する」に限る。
もしクレジットを書く余白がないのであればクレジットが明記されている URI を指示するだけでもいい。
たとえば [Flickr] なら写真が掲載されているページに全ての情報が表示されている。

あとは SaaS の運営者側がとっとと 4.0 にアップグレードしてくれることを祈る。
今回は「写真」だったけど，たとえば Wikipedia 上のコンテンツにだって同様のリスクがある。
Wikipedia の内容をそのままパクって論文に掲載している学生は多そうだ（笑）

yomoyomo さんも[苦言](https://yamdas.hatenablog.com/entry/20220105/copyleft-trolls "著作権トロールの新種としての「コピーレフト・トロール」 - YAMDAS現更新履歴")を呈しておられるが “Copyleft Trolls” という言い回しは止めたほうがいいんじゃないかなぁ。
キャッチーな名前が横滑りして元々の意味が乖離してしまう現象にはウンザリである。
たとえば[ノートン先生まで Crypto を某デジタル資産の意味で使い始めてる](https://community.norton.com/en/blogs/product-service-announcements/introducing-norton-crypto "Introducing Norton Crypto! | Norton Community")みたいで，ホンマ `orz` な気分だよ。

## 余談だが...

私が勘違いしたのは CC Licenses 1.0 から 2.0 あたりで議論になったいわゆる「Creative Commons 実施権」である。

- {{< pdf-file title="著作権と Creative Commons 実施権" link="http://k.lenz.name/j/r/CC.pdf" >}}

ただし，これはライセンスの（被許諾者ではなく）許諾者側のリスクの話。
ちなみに「Creative Commons 実施権」に絡む条項は 4.0 ではまるっと削除されている。
また日本版 2.1 にも存在しない。

いやぁ，私の英語不得手っぷりがまたもや露呈してしまいましたねぇ。
とほほ...

## ブックマーク

- [改訂3版： CC Licenses について]({{< rlnk "/cc-licenses/" >}})

[Flickr]: https://www.flickr.com/

## 参考文献

{{% review-paapi "475710152X" %}} <!-- クリエイティブ・コモンズ―デジタル時代の知的財産権 -->
{{% review-paapi "4641243336" %}} <!-- 著作権法 第3版 -->
