+++
title = "天下無敵のプー太郎"
date =  "2020-04-01T21:59:18+09:00"
description = "次の仕事が得られるまでどのくらい時間がかかるか分からないが，あり余る時間を趣味に没頭していく所存である。"
image = "/images/attention/kitten.jpg"
tags = [ "privacy", "risk", "programming", "language" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

皆さんご機嫌はいかがですか。
悪夢の{{< ruby "エイプリルフール" >}}4月1日{{< /ruby >}}を今年も無事やり過ごせたようなので，通常運転を再開しよう。

[WHO の situation report](https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports "Situation reports") を見る限り，世界レベルではようやく峠を超えた感じだけど日本は今が正念場だよね（まぁ第3第4のピークが来ないとも限らないが）。

私個人は3月いっぱいで目出度く職場をクビになったので，しばらく NEET (Not in Employment, Education or Training) でもしようかなぁ，という感じ。
せっかく自粛ムードだし，私は「人に会わない」ことに全くストレスを感じないので（むしろ実家で家族と顔をつきあわせる生活を1年以上続けられている自分を褒めたい），不謹慎ながら今の状況は大歓迎である。

## “The trade-offs are changing”

SARS 2 関連については思うところは沢山あるし愚痴を書いたらキリがないが[^fb1]，今回は以下の記事を紹介するに留める。

[^fb1]: 詳しくは Facebook の私の記述を参考に。友人以外は読めないようにしてるけど（笑）

- [EFF and COVID-19: Protecting Openness, Security, and Civil Liberties | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2020/03/eff-and-covid-19-protecting-openness-security-and-civil-liberties)
- [As Coronavirus Surveillance Escalates, Personal Privacy Plummets - The New York Times](https://www.nytimes.com/2020/03/23/technology/coronavirus-surveillance-tracking-privacy.html)
- [Privacy vs. Surveillance in the Age of COVID-19 - Schneier on Security](https://www.schneier.com/blog/archives/2020/03/privacy_vs_surv.html)

例によって Bruce Schneier 先生の要約記事が分かりやすいので，そちらを参考にするが EFF の

{{< fig-quote type="markdown" title="Protecting Openness, Security, and Civil Liberties" link="https://www.eff.org/deeplinks/2020/03/eff-and-covid-19-protecting-openness-security-and-civil-liberties" lang="en" >}}
{{< quote >}}We must be sure that measures taken in the name of responding to COVID-19 are, in the language of international human rights law, “necessary and proportionate” to the needs of society in fighting the virus. Above all, we must make sure that these measures end and that the data collected for these purposes is not re-purposed for either governmental or commercial ends{{< /quote >}}.
{{< /fig-quote >}}

という一文には激しく同意するものである。
ついでに以下の tweet も挙げておくか。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">かつてThe Economistは、「テロ対策と称して国家統制を許し、自由を捨てるな。テロで人が死ぬ。それがどうした」と言い放つ強さを持っていたけれど、いまコロナを前にして「国家統制を認めて自由を捨てるな。コロナで人が死ぬ。それがどうした」と言うだけの勇気はないらしい。でもそう言うべきだ。</p>&mdash; Hiroo Yamagata (@hiyori13) <a href="https://twitter.com/hiyori13/status/1240969104409817091?ref_src=twsrc%5Etfw">March 20, 2020</a></blockquote>
{{< /fig-gen >}}

現代の人々にとって「健康」とは{{< ruby "のろ" >}}呪{{< /ruby >}}いのようなものかもしれない。

リスク評価には「科学的リスク」と「政治的リスク」がある。
「政治的リスク」で重要なのは「その後」についてきちんと目配せできているかどうかである。
いくら危機的状況であっても「その後」に配慮しないその場しのぎの「政策」は悪手だ。

- [リスクは事前と事後をセットで考える]({{< ref "/remark/2020/02/risk-of-infection.md" >}})

そういう観点で是非とも議論していただきたいものである。
まぁ，私は（そういうのに巻き込まれたくないので）今年も政治的無監視を貫くけどね。
政治がどうこう言うような精神的余裕はない。

## そろそろ本格的に...

せっかく時間があり余ってるので，色々遊んでみたい。
プログラミング関連では

1. [Go]
2. [Rust]
3. [Haskell]

の優先順位で ~~勉強~~ もとい遊んでみようかな，と。

[Go] と [Rust] は同じ制御系言語なのに思想が真逆で実に面白い。

[Go] は典型的な手続き型言語で並列処理やメモリ管理を隠蔽することでシンプルな記述を可能にしている。
[Rust] は並列処理やメモリ管理については剥き出しだが，その分（標準ライブラリを含めて）豊かな語彙を持っている。

あと [Rust] を[勉強し始めて]({{< rlnk "/rust-lang">}} "Rust のキホン")痛感したが，やっぱ関数型プログラミング言語についてきちんと学ばないとダメだわ。
そういう意味で [Haskell] をもう一度勉強し直そうかと思っている。

次の仕事が得られるまでどのくらい時間がかかるか分からないが，あり余る時間を趣味に没頭していく所存である（笑）

## ブックマーク

- [クーリエ連載；エコノミスト紹介、自由のためなら人が死んでもいい](http://cruel.org/economist/courier200712.html)（この中で紹介されている[「自由の本当のコスト」の原文](https://www.economist.com/leaders/2007/09/20/the-real-price-of-freedom "The real price of freedom - Civil liberties under threat")）
- [専門家会議の「クラスター対策」の解説　――新型コロナウイルスに対処する最後の希望｜吉峯 耕平｜note](https://note.com/kyoshimine/n/n6bf078a369f9)
- [ある医師がエンジニアに寄せた“コロナにまつわる現場の本音” (1/10) - EE Times Japan](https://eetimes.jp/ee/articles/2003/25/news053.html)

[Go]: https://golang.org/ "The Go Programming Language"
[Rust]: https://www.rust-lang.org/ "Rust Programming Language"
[Haskell]: https://www.haskell.org/ "Haskell Language"
