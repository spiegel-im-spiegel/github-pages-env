+++
title = "最近の話題（2026-06-07）"
date =  "2026-06-07T12:29:03+09:00"
description = "非会員による Kagi Translate 利用の一時停止 / Bridgy Fed からの信頼できる脱出 / Linus Torvalds は激おこらしい（いつものこと？） / 欧州は脱米国に突き進むのか"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "kagi", "artificial-intelligence", "atproto", "bluesky", "politics", "market" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

例によって Mastodon や Bluesky の TL で見かけた記事を紹介する。
今回は複数あるので，コメントは簡単に。

お品書きはこちら：

- [非会員による Kagi Translate 利用の一時停止]({{< rlnk "/remark/2026/06/07-stories/#kagi">}})
- [Bridgy Fed からの信頼できる脱出]({{< rlnk "/remark/2026/06/07-stories/#exit">}})
- [Linus Torvalds は激おこらしい（いつものこと？）]({{< rlnk "/remark/2026/06/07-stories/#ai">}})
- [欧州は脱米国に突き進むのか]({{< rlnk "/remark/2026/06/07-stories/#breakup">}})

ではどうぞ。

## 非会員による Kagi Translate 利用の一時停止 {#kagi}

- [An update on Kagi Translate](https://blog.kagi.com/translate-update) 

私は [Kagi Translate] がローンチする前にサブスクリプション契約をしていたので気が付かなかったが，今まではサインアップしなくても誰でも [Kagi Translate] を使うことができたらしい。
したら，エラい人気だったらしく，運営コストがヤバいことになってるそうな。
そこで，一時的に有料会員（無料の Trial プランも除外かな）のみ利用可能にしたとのこと。
この一時的制限がいつ解除されるのか（あるいは恒久的なものになるのか）はまだ決まってないみたい。
もしかしたら [Kagi Translate] のみ利用可能な格安プランとかできるかも？

[Kagi Translate] の日本語訳は，最初の頃こそぎこちなかったが，今では Google 翻訳を凌駕する性能だと思っている。
文脈や話者ごとに別訳も提示してくれるし，入力もテキストだけでなく PDF やオフィスドキュメントなどにも対応している。
あとは辞書機能と校正機能かな。
校正機能は結構ヘヴィに使っている人もいるようだ。

- [Kagi 翻訳のスマホアプリが登場]({{< ref "/remark/2026/02/kagi-translate-arrives-on-mobile.md" >}})
- [Kagi 翻訳によるおもしろ翻訳]({{< ref "/remark/2026/03/kagi-fun.md" >}})

個人的にはメインの [Kagi Search] にしても [Kagi Assistant] にしても（月10USD程度なら）お金を払う価値は十分あると思うけどねぇ。

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

## Bridgy Fed からの信頼できる脱出 {#exit}

- [Launches: Credible Exit, Bridged User Notifications, & Fediverse Re-Bridge](https://blog.anew.social/migrations-update-launch/)

[Bridgy Fed] 内にあるアカウントをネイティブ・アカウントに移行できる機能を提供開始したらしい。

{{< fig-quote type="markdown" title="Launches: Credible Exit, Bridged User Notifications, & Fediverse Re-Bridge" link="https://blog.anew.social/migrations-update-launch/" lang="en" >}}
If you’re bridging from the Atmosphere to the Fediverse, you can migrate the bridged account to Mastodon; if you’re bridging from the Fediverse to the Atmosphere, you can migrate the bridged account to any PDS of your choice. After that, you’ll have two native accounts that you can choose to use separately, or migrate them into a service like Wafrn that lets you use both protocols with one account.
{{< /fig-quote >}}

さらに Mastodon でアカウントを「引っ越し」した場合でも Bridgy Fed も追従できるようになったそうな。

{{< fig-quote type="markdown" title="Launches: Credible Exit, Bridged User Notifications, & Fediverse Re-Bridge" link="https://blog.anew.social/migrations-update-launch/" lang="en" >}}
From now on, we track your Fediverse moves and automatically migrate your bridged account with it. No action needed on your side – we handle all the work, and you can keep posting across the bridge to all your followers on the Atmosphere.
{{< /fig-quote >}}

個人的にこれらの機能を使うことは（今のところ）ないと思うけど，プラットフォームからの脱出方法があるというのは安心材料と言える。

[Bridgy Fed]: https://fed.brid.gy/ "Bridgy Fed"

## Linus Torvalds は激おこらしい（いつものこと？） {#ai}

- [Why Linux creator Linus Torvalds gets angry hearing "99% of code is AI"](https://thenewstack.io/torvalds-ai-programming-productivity/)

{{< fig-quote type="markdown" title="Why Linux creator Linus Torvalds gets angry hearing \"99% of code is AI\"" link="https://thenewstack.io/torvalds-ai-programming-productivity/" lang="en" >}}
“AI is a great new tool, but it’s a tool, and when I see people saying, ‘Hey, 99% of our code is written by AI,’ I literally get angry, because those same people — I can pretty much guarantee — that 100% of their code is written by compilers. But they never say that.”
{{< /fig-quote >}}

「コードの 99% が AI によって書かれている」としても AI が書けない（書かない）残り 1% が最も重要なのよ。
「コードを書かない人」はそのことに気が付かず，統計上の数字だけを見て人材を機械に置き換えようとする（そして往々にして失敗する）。
これはプログラミングに限ったことではないと思う。

昨今の「生成 AI」に対する Linus Torvalds 氏の見解は一貫していて，上の引用にもある通り「 AI はあくまでもツール」というものだ。
故に，この便利なツールをどのように使うかということには関心があるようだが， AI が人に取って代わるとか，ましてや[「自分たちは神を創造できる」](https://gigazine.net/news/20260602-anthropic-thinks-building-god/ "「自分たちは神を創造できる」とAI開発企業のAnthropicが本気で考えているのではないかとの指摘 - GIGAZINE")などと頭のおかしいことを言う連中には絶対に与しないと思う。

私もどちらかというと Linus Torvalds 氏に近いかな。
ただ AI に対して命令（command）でも指示（instruction）でもなくプロンプト（prompt）を与えるという点が従来の「道具」とは違う点だと最近は[考えている]({{< ref "/remark/2026/05/reading-living-with-ai.md#sycophancy" >}} "追従する AI は向社会的な姿勢を減退させる？")。

{{< fig-quote type="markdown" title="Why Linux creator Linus Torvalds gets angry hearing \"99% of code is AI\"" link="https://thenewstack.io/torvalds-ai-programming-productivity/" lang="en" >}}
“People who know what they’re doing to understand systems will be able to prompt tools to write good code,” Torvalds said. “People who don’t understand the complexity of systems will also prompt systems and write processes that will fail. So,  I think you do want to understand how it all works.”
{{< /fig-quote >}}

[GitHub Copilot] が登場した当初はコードを「洗浄」する生成 AI に反発が大きかったが，今や AI にコードを書かせることが当たり前になっている。
その一方で

<blockquote class="mastodon-embed" data-embed-url="https://mstdn.jp/@zetamatta/116677027811816194/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://mstdn.jp/@zetamatta/116677027811816194" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M74.7135 16.6043C73.6199 8.54587 66.5351 2.19527 58.1366 0.964691C56.7196 0.756754 51.351 0 38.9148 0H38.822C26.3824 0 23.7135 0.756754 22.2966 0.964691C14.1319 2.16118 6.67571 7.86752 4.86669 16.0214C3.99657 20.0369 3.90371 24.4888 4.06535 28.5726C4.29578 34.4289 4.34049 40.275 4.877 46.1075C5.24791 49.9817 5.89495 53.8251 6.81328 57.6088C8.53288 64.5968 15.4938 70.4122 22.3138 72.7848C29.6155 75.259 37.468 75.6697 44.9919 73.971C45.8196 73.7801 46.6381 73.5586 47.4475 73.3063C49.2737 72.7302 51.4164 72.086 52.9915 70.9542C53.0131 70.9384 53.0308 70.9178 53.0433 70.8942C53.0558 70.8706 53.0628 70.8445 53.0637 70.8179V65.1661C53.0634 65.1412 53.0574 65.1167 53.0462 65.0944C53.035 65.0721 53.0189 65.0525 52.9992 65.0371C52.9794 65.0218 52.9564 65.011 52.9318 65.0056C52.9073 65.0002 52.8819 65.0003 52.8574 65.0059C48.0369 66.1472 43.0971 66.7193 38.141 66.7103C29.6118 66.7103 27.3178 62.6981 26.6609 61.0278C26.1329 59.5842 25.7976 58.0784 25.6636 56.5486C25.6622 56.5229 25.667 56.4973 25.6775 56.4738C25.688 56.4502 25.7039 56.4295 25.724 56.4132C25.7441 56.397 25.7678 56.3856 25.7931 56.3801C25.8185 56.3746 25.8448 56.3751 25.8699 56.3816C30.6101 57.5151 35.4693 58.0873 40.3455 58.086C41.5183 58.086 42.6876 58.086 43.8604 58.0553C48.7647 57.919 53.9339 57.6701 58.7591 56.7361C58.8794 56.7123 58.9998 56.6918 59.103 56.6611C66.7139 55.2124 73.9569 50.665 74.6929 39.1501C74.7204 38.6967 74.7892 34.4016 74.7892 33.9312C74.7926 32.3325 75.3085 22.5901 74.7135 16.6043ZM62.9996 45.3371H54.9966V25.9069C54.9966 21.8163 53.277 19.7302 49.7793 19.7302C45.9343 19.7302 44.0083 22.1981 44.0083 27.0727V37.7082H36.0534V27.0727C36.0534 22.1981 34.124 19.7302 30.279 19.7302C26.8019 19.7302 25.0651 21.8163 25.0617 25.9069V45.3371H17.0656V25.3172C17.0656 21.2266 18.1191 17.9769 20.2262 15.568C22.3998 13.1648 25.2509 11.9308 28.7898 11.9308C32.8859 11.9308 35.9812 13.492 38.0447 16.6111L40.036 19.9245L42.0308 16.6111C44.0943 13.492 47.1896 11.9308 51.2788 11.9308C54.8143 11.9308 57.6654 13.1648 59.8459 15.568C61.9529 17.9746 63.0065 21.2243 63.0065 25.3172L62.9996 45.3371Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @zetamatta@mstdn.jp</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://mstdn.jp/" async src="https://mstdn.jp/embed.js"></script>

という懸念にも見られるように FOSS の価値というか優位性は，少しずつ低下しているようにも見える。

どうなるやら。

### ブックマーク

- [オープンソースブラウザ「Ladybird」が外部コード受け入れを停止、AI時代の開発体制へ転換](https://gigazine.net/news/20260608-ladybird-change-develop/)

[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"

## 欧州は脱米国に突き進むのか {#breakup}

- [欧州で進む「米国テック依存」から脱却の動き。トランプ政権下で高まる危機感](https://wired.jp/article/the-eu-is-going-through-a-trump-fueled-breakup-with-big-tech/)

ゼロ年代は（オバマ政権に替わるまで[^c1]）「チャイナ・リスク」が問題視された。
今や「US リスク」について真剣に考えなければならない時代になっている。
その意味で欧州が脱米国に走るのは自然な成り行きに思える。

[^c1]: オバマ政権（というか民主党政権）は中国を「未開拓市場」とみなして積極的に関与する政策に転換した。もちろん米国の事実上の属国である日本も追従する。でも中国のような老獪な国に対してそれは悪手だったんじゃないだろうか，と今にしては思えてならない。米国が「覇権主義」から手を引き始めたのは子ブッシュ政権時代から（政権を超えて）一貫している。その意味で現トランプ政権は，そうした米国の姿勢を加速しているだけとも言える。中国は独裁政治をやりすぎて陰りが出ている。英雄はやはり晩節を汚すのか？ まぁ，トランプ氏は最初から晩節で既に汚しまくりだけど（笑） 政治には，過去に対しても未来に対しても「もしも」はない。結果が全て。

私の観測範囲は広くないが，特に [atproto] を中心とした [The Atmosphere]({{< ref "/remark/2026/02/leaflet-and-the-atmosphere.md" >}} "Leaflet と The Atmosphere") 界隈を眺めると脱米国の動きが目立つ。
[Eurosky] とか [Streamplace] とか [tangled] とかあからさまだもんねぇ。

そういやゼロ年代， Bill Gates 氏が Microsoft に健在で「悪の帝国」などと呼ばれていた頃，欧州は脱 Microsoft として Linux などの FOSS 製品を積極的に導入しようとしたんだよね。
結局は失敗したっぽいけど。
今回はどうなるやらだよ

{{< fig-quote type="markdown" title="欧州で進む「米国テック依存」から脱却の動き。トランプ政権下で高まる危機感" link="https://wired.jp/article/the-eu-is-going-through-a-trump-fueled-breakup-with-big-tech/" >}}
たとえ欧州の企業や政府機関がMicrosoft Outlookのような米系サービスから移行したとしても、「スマートフォンでは依然として米国製OSを使っているかもしれませんし、米国企業の管理下にあるインターネット基盤を利用せざるを得ないでしょう」とヴェルディエは語る。「現時点では、こうした巨大テック企業を完全に代替できる存在はありません」
{{< /fig-quote >}}

[Streamplace]: https://stream.place/ "Streamplace"
[atproto]: https://atproto.com/ "The AT Protocol"
[Eurosky]: https://eurosky.tech/ "Eurosky - Building a thriving open social web for Europe"
[tangled]: https://tangled.org/
