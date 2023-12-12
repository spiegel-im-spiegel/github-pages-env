+++
title = "「人工知能と信頼」〜 鉄腕アトムとは友達になれない"
date =  "2023-12-05T22:01:31+09:00"
description = "我らが Bruce Schneier 先生の最新エッセイが面白かったので紹介するふりをして戯れ言を書いてみる。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "code", "market", "politics", "trust" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

我らが Bruce Schneier 先生の最新エッセイが面白かったので紹介するふりをして戯れ言を書いてみる。

- [AI and Trust - Schneier on Security](https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html)

このエッセイでは以下の4つの内容について取り上げている。

1. 信頼には対人信頼（interpersonal trust）と社会的信頼（social trust）の2つがあり，私達はしばしば両者を混同してしまう
2. 人工知能（AI）によってこの混同が拡大する
3. AI システムを管理している企業はこの混同に便乗して私達を利用する
4. 社会への信頼を生み出すのが政府の役割（＝規制）である

その上で

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Not regulating AI, but regulating the organizations that control and use AI.
{{< /fig-quote >}}

と言うのだ。
もうのっけから面白い！

まずは対人信頼（interpersonal trust）について。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
When we say that we trust a friend, it is less about their specific actions and more about them as a person. IIt’s a general reliance that they will behave in a trustworthy manner. We trust their intentions, and know that those intentions will inform their actions. Let’s call this “interpersonal trust.”
{{< /fig-quote >}}

つまり「◯◯さんは人として信頼できる」というのが対人信頼なわけだ。
それに対し，社会的信頼（social trust）については

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
We might not know someone personally, or know their motivations—but we can trust their behavior. We don’t know whether or not someone wants to steal, but maybe we can trust that they won’t. It’s really more about reliability and predictability. We’ll call this “social trust.” It’s the ability to trust strangers.
{{< /fig-quote >}}

と書かれている。
たとえば企業が「顧客サービスは信頼が大事」などというときの「信頼」は対人信頼ではなく社会的信頼ということだ。
そして

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Because of how large and complex society has become, we have replaced many of the rituals and behaviors of interpersonal trust with security mechanisms that enforce reliability and predictability—social trust.
{{< /fig-quote >}}

と続く。
しかし対人信頼と社会的信頼はひと括りに「信頼」と呼んでるため両者を混同してしまう。
カテゴリーエラーが発生するわけだ。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
We might think of them as friends, when they are actually services. Corporations are not moral; they are precisely as immoral as the law and their reputations let them get away with.
{{< /fig-quote >}}

私は上に引用した文章で爆笑してしまったよ。
企業は法律と同じくらい非道徳的だってさ。
言われてみればそうだよね（笑）

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Corporations like that we make this category error—see, I just made it myself—because they profit when we think of them as friends. They use mascots and spokesmodels. They have social media accounts with personalities. They refer to themselves like they are people.
{{< /fig-quote >}}

ここでようやく AI が出てくる。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Science fiction author Ted Chiang writes about it. Instead of solving all of humanity’s problems, or wandering off proving mathematical theorems that no one understands, the AI single-mindedly pursues the goal of maximizing production. Chiang’s point is that this is every corporation’s business plan. And that our fears of AI are basically fears of capitalism.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
And near-term AIs will be controlled by corporations. Which will use them towards that profit-maximizing goal. They won’t be our friends. At best, they’ll be useful services. More likely, they’ll spy on us and try to manipulate us.
{{< /fig-quote >}}

そして，企業に制御される AI は以下の2つの理由で悪い方向に転ぶと予測する。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
The first is that these AI systems will be more relational. We will be conversing with them, using natural language. As such, we will naturally ascribe human-like characteristics to them.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
The second reason to be concerned is that these AIs will be more intimate. One of the promises of generative AI is a personal digital assistant. Acting as your advocate with others, and as a butler with you.
{{< /fig-quote >}}

その結果

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
It’s no accident that these corporate AIs have a human-like interface. There’s nothing inevitable about that. It’s a design choice. It could be designed to be less personal, less human-like, more obviously a service—like a search engine . The companies behind those AIs want you to make the friend/service category error. It will exploit your mistaking it for a friend. And you might not have any choice but to use it.
{{< /fig-quote >}}

私達は AI によって「信頼」を混同したまま強制されるわけだ。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Corporations are profit maximizers, at the expense of society. And the incentives of surveillance capitalism are just too much to resist.
{{< /fig-quote >}}

そして，これに規制をかけるのは政府の役割となる。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
The more you can trust that your societal interactions are reliable and predictable, the more you can ignore their details. Places where governments don’t provide these things are not good places to live.
{{< /fig-quote >}}

“not good places to live” って今の日本社会のことか（笑）

それはともかく

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Government can do this with AI. We need AI transparency laws. When it is used. How it is trained. What biases and tendencies it has. We need laws regulating AI—and robotic—safety. When it is permitted to affect the world. We need laws that enforce the trustworthiness of AI. Which means the ability to recognize when those laws are being broken. And penalties sufficiently large to incent trustworthy behavior.
{{< /fig-quote >}}

そのために多くの国が AI 規制について議論しているが，そこに大きな間違いがあるという。
つまり

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
Many countries are contemplating AI safety and security laws—the EU is the furthest along—but I think they are making a critical mistake. They try to regulate the AIs and not the humans behind them.
{{< /fig-quote >}}

ということだ。
何故なら AI は人間ではないから。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
AIs are not people; they don’t have agency. They are built by, trained by, and controlled by people. Mostly for-profit corporations. Any AI regulations should place restrictions on those people and corporations. Otherwise the regulations are making the same category error I’ve been talking about. At the end of the day, there is always a human responsible for whatever the AI’s behavior is. And it’s the human who needs to be responsible for what they do—and what their companies do. Regardless of whether it was due to humans, or AI, or a combination of both.
{{< /fig-quote >}}

AI とは友達になれない。
でも AI を善いサービスにすることはできる。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
We can never make AI into our friends. But we can make them into trustworthy services—agents and not double agents. But only if government mandates it. We can put limits on surveillance capitalism. But only if government mandates it.
{{< /fig-quote >}}

そのための政府ってことだよね。

{{< fig-quote type="markdown" title="AI and Trust" link="https://www.schneier.com/blog/archives/2023/12/ai-and-trust.html" lang="en" >}}
To the extent a government improves the overall trust in society, it succeeds. And to the extent a government doesn’t, it fails.
{{< /fig-quote >}}

あれ？ やっぱり日本政府って失敗してる？

このエッセイを読みながら5年前に書いた記事を思い出してた。

{{< fig-quote type="markdown" title="AI アシスタントはユーザをアシストしない" link="/remark/2018/04/ai-assistant/" >}}
Amazon や Google などが提供する AI アシスタントの最大の問題はユーザがコントロールできないことにある。
喩えるなら百貨店の外商がリアルタイムで店と連携をとりつつ24時間家の中に張り付いているようなものだ。
「外商」が忠誠を誓うのは所属する店であり，客との関係は利害の一致を基本とした信頼関係に過ぎない。
これが「執事」なら向きが逆になる。
「執事」は主人に忠誠を誓うからこそ執事として機能するし主人は執事をコントロールできる。
{{< /fig-quote >}}

結局そこに帰着するんだよなぁ。

そうそう。
関係ないが ~~「地上最大のロボット」~~ じゃない「PLUTO」のアニメ版が Netflix で公開されているらしい。
当時コミックをあれほど[絶賛した](https://baldanders.info/spiegel/log/200605.html#d07_t2)のに，アニメはまだ観てないとか。
平日はアニメを観る気がしなくて，週末はチャリンコを乗り回して遊んでるので，やっぱりアニメは観なくなった。
このまま年を越してしまうのか...

## ブックマーク

- [AI and Mass Spying - Schneier on Security](https://www.schneier.com/blog/archives/2023/12/ai-and-mass-spying.html)
- [オープンソースの失われた10年と「オープンソースAI」の行方 – WirelessWire News](https://wirelesswire.jp/2023/12/85735/)

## 参考図書

{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
{{% review-paapi "B0BHYJKB5N" %}} <!-- PLUTO -->
