+++
title = "プライバシーの乱獲"
date =  "2024-06-06T20:13:28+09:00"
description = "独断と偏見に塗れているので，誤読はご容赦（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "code", "privacy", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

プライバシー（権）が指し示すものは時代とともに変わる。

{{< fig-quote type="markdown" title="「ネットが蝕むプライバシー」を読む（再掲載）" link="/remark/2016/11/the-end-of-privacy/" >}}
もともとのプライバシーは空間概念と一体になっている。 つまり私的領域と公的領域を区別することである。 しかしインターネット（特に Web 2.0 以降）は情報は空間概念から切り離される。 場所が不定になるのだ。 だから，それをコントロールしたければ情報へのアクセス性を対象にせざるを得なくなる。
{{< /fig-quote >}}

また「誰」からプライバシーを守るか，その対象も変わってきた。
インターネット上のプラットフォームを巨大企業が牛耳るようになれば，プライバシーの問題は「個人 vs. 市場」に発展する。

- [誰がプライバシーを支配するのか]({{< ref "/remark/2018/04/handling-privacy.md" >}})

一方で市場は私達個人に大いなる恩恵をもたらしてくれる。
こうした恩恵とプライバシーはトレードオフにされているのだ。
このため「どこまで自身のプライバシー（に関する情報）を渡せばいいのか」という議論に流れがちである。

こうした「シフト」を漁獲量管理の問題に喩えたエッセイが公開されている。

- [Microsoft AI Spying Scandal: Time to Rethink Privacy Standards - IEEE Spectrum](https://spectrum.ieee.org/online-privacy)
- [Online Privacy and Overfishing - Schneier on Security](https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html)

これは Bruce Schneier と Barath Raghavan の共著のようだ。
今回はこのエッセイを少し真面目に読んでみようと思う。
とはいえ独断と偏見に塗れているので，誤読はご容赦（笑）

----

話の{{< ruby トリガー >}}発端{{< /ruby >}}は，スピアフィッシング攻撃等の精度を高めるために生成 AI チャットボット・サービスを利用している国家支援のハッカー集団を発見しこれを排除したというニュースだ。

{{< fig-quote type="markdown" title="Microsoft says it caught hackers from China, Russia and Iran using its AI tools | Reuters" link="https://www.reuters.com/technology/cybersecurity/microsoft-says-it-caught-hackers-china-russia-iran-using-its-ai-tools-2024-02-14/" lang="en" >}}
Microsoft ([MSFT.O](https://www.reuters.com/markets/companies/MSFT.O)), opens new tab said in its report it had tracked hacking groups affiliated with Russian military intelligence, Iran's Revolutionary Guard, and the Chinese and North Korean governments as they tried to perfect their hacking campaigns using large language models. Those computer programs, often called artificial intelligence, draw on massive amounts of text to generate human-sounding responses.

The company announced the find as it rolled out a blanket ban on state-backed hacking groups using its AI products.
{{< /fig-quote >}}

こうした「悪用」自体は予測可能なことではあるが，問題はどうやって件の「ハッカー」を見つけ出したか，だろう。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
In the security community, the immediate questions weren’t about how hackers were using the tools (that was utterly predictable), but about how Microsoft figured it out. The natural conclusion was that Microsoft was spying on its AI users, looking for harmful hackers at work.
{{< /fig-quote >}}

まぁ，そうだよな（笑） ただこれを「スパイ」行為と言っていいのかということについて疑問を呈す人もいる，当然ながら。

{{< fig-quote type="markdown" title="Microsoft is spying on users of its AI tools | Hacker News" link="https://news.ycombinator.com/item?id=39442429" lang="en" >}}
If you call this "spying", you should read the TOS of the OpenAI API again. If you, as a "hacker", still use their API
{{< /fig-quote >}}

最初から ToS (Terms of Service) に書いてあるぢゃん，文句ゆーな！ というわけだ。
本当に？

ここで先ほどの漁獲量管理の例が出てくる。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
One scientist, [Daniel Pauly](https://oceans.ubc.ca/2023/05/19/daniel-pauly/), realized that researchers studying fish populations were making a major error when trying to determine acceptable catch size. It wasn’t that scientists didn’t recognize the declining fish populations. It was just that they didn’t realize how significant the decline was. Pauly noted that each generation of scientists had a different baseline to which they compared the current statistics, and that each generation’s baseline was lower than that of the previous one.

[...]

Pauly called this “[shifting baseline syndrome](https://en.wikipedia.org/wiki/Shifting_baseline)” in a 1995 paper. The baseline most scientists used was the one that was normal when they began their research careers. By that measure, each subsequent decline wasn’t significant, but the cumulative decline was devastating. Each generation of researchers came of age in a new ecological and technological environment, inadvertently masking an exponential decline.
{{< /fig-quote >}}

これってそのまま「シフティング・ベースライン症候群」って呼べばいいのかな。
「基準推移症候群」って訳してるところもあるみたいだけど。
まぁ，とりあえず前者で行こう。

[Bruce Schneier & Barath Raghavan のエッセイ][エッセイ]では，プライバシーにおいてもベースラインのシフトが起きていると指摘する。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
Historically, people controlled their computers, and software was standalone. The always-connected cloud-deployment model of software and services flipped the script. Most apps and services are designed to be always-online, feeding usage information back to the company. A consequence of this modern deployment model is that everyone—cynical tech folks and even ordinary users—expects that what you do with modern tech isn’t private. But that’s because the baseline has shifted.

[...]

Shifting baselines are at the heart of our collective loss of privacy. The U.S. Supreme Court has long held that our right to privacy depends on whether we have a reasonable [expectation of privacy](https://www.law.cornell.edu/wex/expectation_of_privacy). But expectation is a slippery thing: It’s subject to shifting baselines.
{{< /fig-quote >}}

ここで最初に挙げた「スパイ」の話に戻る。
生成 AI チャットボット・サービスを利用している殆どは善良なユーザで，そこで疚しい行動をしているわけではない。
疚しいことをしていないならプライバシー情報を吸い上げられたところで問題ないのではないか（ゼロ年代にそういう主張があったな）。
むしろ今回のようにサービスを「悪用」するハッカーを釣り上げられるんだからええぢゃん！ という感じ。

でも，サービスがユーザの行動を監視し情報を吸い上げるのはユーザのためではない。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
AI chatbots are the latest incarnation of this phenomenon: They produce output in response to your input, but behind the scenes there’s a complex cloud-based system keeping track of that input—both to improve the service and to [sell you ads](https://about.ads.microsoft.com/en/blog/post/february-2023/the-new-bing-creating-value-for-advertisers).
{{< /fig-quote >}}

陰謀論的な言いがかりをするなら，プラットフォーマー（笑）は利便性やセキュリティを「人質」にとって，プライバシーのベースラインを意図的に引き下げてるんじゃないか，などとゲスの勘ぐりをしてしまう。
まぁ，本当にそうなら「シフティング・ベースライン症候群」なんてもんじゃないけどな。

話が逸れた。

意図せざるベースラインのシフトによる「崩壊（collapse）」を回避するにはどうすればいいか。
それはもう議論の焦点を変えるしかない。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
Fisheries scientists, armed with knowledge of shifting-baseline syndrome, now look at the big picture. They no longer consider relative measures, such as comparing this decade with the last decade. Instead, they take a holistic, ecosystem-wide perspective to see what a healthy marine ecosystem and thus sustainable catch should look like. They then turn these scientifically derived sustainable-catch figures into limits to be codified by regulators.
{{< /fig-quote >}}

プライバシーやセキュリティについて考える場合も同じ。
その時々の「ベースライン」で考えるのではなく，系全体を俯瞰する視点が必要だろう。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
Instead of comparing to a shifting baseline, we need to step back and look at what a healthy technological ecosystem would look like: one that respects people’s privacy rights while also allowing companies to recoup costs for services they provide. Ultimately, as with fisheries, we need to take a big-picture perspective and be aware of shifting baselines.
{{< /fig-quote >}}

その上で，[件のエッセイ][エッセイ]では具体的な手段として「科学的情報に基づいた民主的な規制プロセス」が必要と主張する。

{{< fig-quote type="markdown" title="Online Privacy and Overfishing" link="https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html" lang="en" >}}
A scientifically informed and democratic regulatory process is required to preserve a heritage—whether it be the ocean or the Internet—for the next generation.
{{< /fig-quote >}}

ところで「シフティング・ベースライン症候群」でググると色んなシチュエーションで使われているっぽいね。
でも検索結果を拾い読みしてみると，思い込みや認知の問題として扱ってるところが多いような。
「[茹でガエル](https://ja.wikipedia.org/wiki/%E8%8C%B9%E3%81%A7%E3%82%AC%E3%82%A8%E3%83%AB "茹でガエル - Wikipedia")」を例に挙げてるところもあるみたいだし。
実際にはこれって生態系（ecosystem）におけるスコープの取り方の問題で，それを適切に扱わないと科学的な議論と言えど致命的な間違いをおかすよってことだと思うのだけど。

今回のプライバシーの話に引き付けて考えるのなら，最初に述べたように，オンラインプラットフォームにおいてプライバシーはセキュリティや利便性とのトレードオフの対象になることが多く，セキュリティや利便性のために「どこまで自身のプライバシー（に関する情報）を渡せばいいのか」という方向に流れがちだ。
その文脈で，その時々のプライバシーに対する「合理的な期待（reasonable [expectation of privacy](https://www.law.cornell.edu/wex/expectation_of_privacy "expectation of privacy | Wex | US Law | LII / Legal Information Institute")）」をベースラインとして議論する。
でもホンマにそんな議論の仕方でええのん？ というのが今回の[エッセイ]の核心なんじゃないかと思っている。

プライバシーは切り売りする商材ではなく，社会と個人の関係の中で「絶対に譲れないもの」があるという点が重要であり，個人が台頭する近代〜現代社会ならではの命題だ。
もちろん簡単ではない。

[エッセイ]: https://www.schneier.com/blog/archives/2024/06/online-privacy-and-overfishing.html "Online Privacy and Overfishing - Schneier on Security"

## 参考図書

{{% review-paapi "4296001574" %}} <!-- ハッキング思考 -->
{{% review-paapi "4571210450" %}} <!-- はじめて学ぶ ビデオゲームの心理学 -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "B0125TZSZ0" %}} <!-- つながりっぱなしの日常を生きる -->
