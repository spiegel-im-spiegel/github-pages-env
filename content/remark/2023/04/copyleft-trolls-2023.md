+++
title = "“Copyleft Trolls” とたたかう 2023"
date =  "2023-04-08T12:21:36+09:00"
description = "「文化の発展に寄与する」ってのは，こういう活動を指すんだよね，本来は！ "
image = "/images/attention/kitten.jpg"
tags = ["code", "creative-commons", "copyright", "risk", "license", "flickr"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近すっかり疎遠になっている Twiter の TL を眺めてたら

- [ライセンスの執行についての資料を公開しました | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/2023/04/07/enforcement-principles-now-published/)
  - [ライセンスの執行 | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/license-enforcement/)
  - [執行に関する原則 | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/license-enforcement/enforcement-principles/)
  - [もしあなたのCCライセンス作品が不適切な利用をされたら | クリエイティブ・コモンズ・ジャパン](https://creativecommons.jp/misuse-of-works/)

という記事が[紹介](https://twitter.com/yomoyomo/status/1644355169231200256)されていた。
つか CCjp ってまだ活動してるんだねぇ。
もう看板だけの組織だと思ってたよ（笑）

いわゆる “copyleft trolls” は[^cl1] CC licenses の古いバージョンに起因する不備をついたもので，2022年のはじめ頃に話題になったものだ。

[^cl1]: CC licences は厳密には copyleft ではないので「[「コピーレフト・トロール」という言葉は問題があるよね](https://yamdas.hatenablog.com/entry/20220105/copyleft-trolls "著作権トロールの新種としての「コピーレフト・トロール」 - YAMDAS現更新履歴")（超意訳）」と yomoyomo さんも書いておられたが，すっかり定着してしまった感がある。

- [CC Licenses を悪用した新手の著作権トロルについて]({{< ref "/remark/2022/01/new-copyright-trolls.md" >}})
- [「トロルにエサを与えないでください」]({{< ref "/remark/2022/02/do-not-feed-the-trolls.md" >}})

1年以上前の話を今更？ 選挙時期だし，ひょっとして何かの政治宣伝か？ と思ったが，どうも最近になって本家 CC が出した

- [Statement of Enforcement Principles - Creative Commons](https://creativecommons.org/license-enforcement/enforcement-principles/)

を受けてのものらしい。
CC licences 執行の3つの原則とは以下のとおり：

{{< fig-quote type="markdown" title="Statement of Enforcement Principles" link="https://creativecommons.org/license-enforcement/enforcement-principles/" lang="en" >}}
- The primary goal of license enforcement should be getting reusers to comply with the license.
- Legal action should be taken sparingly. 
- Enforcement may involve monetary compensation, but should not be a business model.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="執行に関する原則" link="https://creativecommons.jp/license-enforcement/enforcement-principles/" >}}
- 再利用を行う者にライセンスを遵守してもらうことを、ライセンスの執行の第一の目的とすること
- 法的な措置は頻繁に行われるべきではないこと
- 金銭的補償をもたらす執行もあり得るが、それがビジネスモデルとなるべきではないこと
{{< /fig-quote >}}

もう少し言うと「ライセンスの目的は（利用条件を遵守させることにより）知的コモンズを拡大することにあり，（違反を誘発して）お金を巻き上げるビジネスモデルであってはならない」ということだ。

“Copyleft trolls” の経緯と経過については [Cory Doctorow] さんが最近書かれた “[Flickr to copyleft trolls: drop dead](https://pluralistic.net/2023/04/01/pixsynnussija/#pilkunnussija "Pluralistic: Flickr to copyleft trolls: drop dead (01 Apr 2023) – Pluralistic: Daily links from Cory Doctorow")” が詳しい。
4月1日と明記されてたのでエイプリルフールを疑ったが，マジな記事なようだ（笑）

これによると Flickr のガイドラインに以下の文言が追加された。

{{< fig-quote type="markdown" title="Flickr Community guidelines | Flickr" link="https://www.flickr.com/help/guidelines" lang="en" >}}
Give some grace. When you choose to grant permission to your photos under any [open license](https://www.flickrhelp.com/hc/en-us/articles/4404070159636-Creative-Commons) available on Flickr, we ask that you give the reuser a 30-day grace period to fix any possible mistake or misuse of your CC-licensed work with no penalty. Failure to allow a good faith reuser the opportunity to correct errors is against the intent of the license and not in line with the values of our community, and can result in your account being removed. If you use CC licenses, please understand that we support and adhere to the strategy for addressing license enforcement described in the [Creative Commons’ Statement of Enforcement Principles](https://creativecommons.org/license-enforcement/enforcement-principles/).
{{< /fig-quote >}}

これは上述した CC の「執行に関する原則」を受けてのものだね。
このガイドラインに違反する「トロル」な行為を受けた場合の報告手段も示されている。

- [How to report Community Guidelines violations – Flickr Help Center](https://www.flickrhelp.com/hc/en-us/articles/4404057906068-How-to-report-Community-Guidelines-violations)

Flickr は Yahoo! / Verizon 時代の「技術的負債」を返済するのに忙しく，なかなか前に進んでいないように見えるが

{{< fig-quote type="markdown" title="Flickr to copyleft trolls: drop dead" link="https://pluralistic.net/2023/04/01/pixsynnussija/#pilkunnussija" lang="en" >}}
But what they have done is modify their policies to create a de facto CC 4.0 environment for their users, by promising to terminate the accounts of any user who repeatedly threatens legal action over bad attribution strings without first offering a 30-day grace period.

Flickr's done more than that, actually. For one thing, they ditched Pixsy, severing their relationship with the company (Pixsy still lists them on its "partner" page). They also created the Flickr Foundation, a nonprofit devoted to providing long-term, responsible stewardship for their CC and public domain image respositories:

https://www.flickr.org/
{{< /fig-quote >}}

とも書かれている。
Flickr Foundation については私も以前に紹介したので参考にどうぞ。

- [The Flickr Foundation 100年の計]({{< ref "/remark/2022/11/the-flickr-foundation.md" >}})

あとは [Cory Doctorow] さん自身が脅迫を受けた（？）話とか書かれていてなかなかに面白いので，ぜひ件の記事をご覧あれ。

[Cory Doctorow] さんといえば

- [Tiktok's enshittification](https://pluralistic.net/2023/01/21/potemkin-ai/#hey-guys "Pluralistic: Tiktok’s enshittification (21 Jan 2023) – Pluralistic: Daily links from Cory Doctorow") （[日本語訳](https://p2ptk.org/monopoly/4366 "メタクソ化するTiktok：プラットフォームが生まれ、成長し、支配し、滅びるまで | p2ptk[.]org")）

という記事があって，その冒頭でいきなり

{{< fig-quote type="markdown" title="Tiktok's enshittification" link="https://pluralistic.net/2023/01/21/potemkin-ai/#hey-guys" lang="en" >}}
Here is how platforms die: first, they are good to their users; then they abuse their users to make things better for their business customers; finally, they abuse those business customers to claw back all the value for themselves. Then, they die.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="メタクソ化するTiktok" link="https://p2ptk.org/monopoly/4366" >}}
プラットフォームはこのように滅びていく。まず、ユーザにとって良き存在になる。次に、ビジネス顧客にとって良き存在になるために、ユーザを虐げる。最後に、ビジネス顧客を虐げて、すべての価値を自分たちに向ける。そうして死んでいく。
{{< /fig-quote >}}

とか書かれていて，あまりの納得感に思わず笑ってしまったのだが，「ビジネス顧客にとって良き存在になるために、ユーザを虐げる」をプラットフォームの第2フェーズとするなら， Flickr はこのフェーズへの移行をどうにか踏み留まっている状況か。
末永く踏み留まっていただきたいものである。

まぁ，なんちうか「文化の発展に寄与する」ってのは，こういう活動を指すんだよね，本来は！ どこぞの文化庁やら文科省やら，あるいは某 Twitter とかに於いては，よくよく考えていただきたい。

[Cory Doctorow]: https://mamot.fr/@doctorow "doctorow (@doctorow@mamot.fr) - La Quadrature du Net - Mastodon - Media Fédéré"

## 参考文献

{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->
