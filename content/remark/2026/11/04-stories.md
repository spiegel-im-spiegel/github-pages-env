+++
title = "四方山話（2026-11-04 の戯れ言）"
date =  "2026-11-04T15:18:59+09:00"
description = "お名前 ID の削除 / 𝕏 アカウントを BAN された / ぞーぺんが mixi2 に対応した！ / Foursquare Placemaker Tools / 主権 AI"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "communication", "tools", "artificial-intelligence", "politics", "media", "twitter" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

ブログネタが溜まってるが，いずれも小ネタばかりなので，今回まとめて放出する。

1. [お名前 ID の削除](#onamae)
2. [𝕏 アカウントを BAN された](#x-ban)
3. [ぞーぺんが mixi2 に対応した！](#zonepane)
4. [Foursquare Placemaker Tools](#foursquare)
5. [主権 AI](#sovereign-ai)

## お名前 ID の削除 {#onamae}

GMO の一連はセキュリティ・インシデントが云々というよりも「企業として大丈夫そう？」なレベルでツッコミどころ満載だが，私個人としては「対岸の火事」のつもりだった...

Mastodon とかの TL でドメインやレンタルサーバの移管・移転のポストを眺めていて「あれ？ そういえば私 お名前.com にアカウントがなかったっけ？」と思い出した。
随分前にドメインを取得したことがあるんだけど，上手く活用できず，継続せずに解約したんだった。

というわけで お名前.com のアカウントも削除することにした。

- [お名前IDの削除｜お名前.com Navi ガイド｜ドメイン取るならお名前.com](https://www.onamae.com/guide/p/6)

削除自体は問題なく終了。

削除の Web ページはレスポンシブじゃないので，スマホだと見づらい。
横着せずパソコンのブラウザで行うことをオススメする。
あと VPN 越しにアクセスすると Cloudflare がアクセスをブロックするみたいなのでご注意を。
なんか最近 Cloudflare にアクセスをブロックされることが増えた気がする。
めっちゃ面倒。

アカウントを作ったのは本当に随分前でカード情報も期限切れだった気がするので，情報が漏れててもそれほどインパクトはないと思う（希望的観測）。

君子危うきに近寄らず，である。

## 𝕏 アカウントを BAN された {#x-ban}

先日 {{% emoji "X" %}} に[人間と認められなかった話]({{< ref "/remark/2026/10/x-verification-failed.md" >}} "𝕏 に人間と認められなかったので活動を停止します")を書いたが，昨日付けで BAN されたようだ。

{{< fig-img src="./x-ban.png" title="𝕏 アカウントを BAN された" link="./x-ban.png" >}}

2007年からの古参ユーザに偽装行為とか何言ってんだか（笑）

あれ以来，サインインしなくても全く困らないことに気付いた（他人の公開ポストはアカウントなしでも見れる）。
決済情報も入れてない。
異議申し立てとか面倒くさいし，老兵は消え去るのみである。

概ね20年か。
長いお付き合いでした。

## ぞーぺんが mixi2 に対応した！ {#zonepane}

これも先日の話だが，マルチ SNS クライアント [ZonePane] (ぞーぺん) が [mixi2] に対応した。
マルチカラムで表示させるとこんな感じにできる。

{{< fig-img src="./zonepane.png" title="ZonePane デスクトップ版" link="./zonepane.png" >}}

今まで [mixi2] は，アカウントは取ったものの放置プレイ状態だったのだが，ぞーぺんが対応してくれたおかげで最近は TL をチェックしたり短文ポストを投げるようになった。
ありがたや {{% emoji "ペコン" %}}

{{< div-gen class="center" type="markdown" >}}
アカウント名は [`@spiegel`](https://mixi.social/@spiegel "Spiegel (@spiegel) | mixi2") です。<br>
興味のある方はどうぞ。
{{< /div-gen >}}

ところで Bluesky 上に「[`#さいきんのBlueskyってどうよ`](https://bsky.app/hashtag/%E3%81%95%E3%81%84%E3%81%8D%E3%82%93%E3%81%AEBluesky%E3%81%A3%E3%81%A6%E3%81%A9%E3%81%86%E3%82%88 "ハッシュタグ — Bluesky")」というハッシュタグがあって，私も少し書いた。

{{< fig-quote type="markdown" >}}
Mastodon や Bluesky が Twitter/X の「出口」となっていた季節はとうの昔に終わっているが，だからといって Bluesky が覇権を取る必要は全くない。というのが最近の感想である。

こう考えるようになったのは #ぞーぺん を使うようになってから。

個人的な感想として #ぞーぺん の重要な機能は主に2つ。

1. アプリ内で複数のサービスを横断的に閲覧・投稿できること（ただし X を除く）
2. タイムラインを過去→現在の方向にスクロールアップして見れること

である。特にスクロールアップで TL を見ることで「作業の終わり」が明確になるので，安心してサービスから離れることができる。これ，本当に重要。

ぞーぺんを通してサービス（群）を見ることで Bluesky も one of them に過ぎないし，それでいいと思うようになった。自分が居心地がいいと思える場所に行けばいいのである。X も含めて

あとは Danah Boyd さんの [parasocial の考え方]({{< ref "/remark/2026/05/facebook-is-para-not-meta.md" >}} "Facebook は meta じゃなくて para だった")か。やっぱ TL 型のメディアは social じゃないよね，と確信を抱くようになった。

というか social media なプラットフォームが欲しいなら Discord みたいなほうがいんじゃね？ まぁ，あれはあれで昨今の政治情勢の煽りで面倒くさくなってるけど，私は Discord に決済情報は入れてないし年齢確認処理も行ってない。それでも大きな問題はない。

Bluesky というか Atmosphere の強みは "[Everything Account](https://www.augment.ink/the-everything-account "The Everything Account - augment")" なので，その意味でも Bluesky は one of them に過ぎない，と言える
{{< /fig-quote >}}

やっぱ「サービスを横断できる」ってのは重要だと思うのよ。
Web なんだから。
旧 Twitter や Facebook など黎明期の SNS が Web 特性を殺してユーザや情報を囲い込んでしまった（そしてユーザがそれをよしとしてしまった）のは重大な失敗だった，と今更ながら思う。

## Foursquare Placemaker Tools {#foursquare}

しばらく前から [Foursquare Placemaker Tools] の案内が来るようになった。
アプリの [Swarm] は今のスマホには入れてなくて [Foursquare] がどうなってるのか全く知らなかったのだが，2024年末に一般ユーザ向けの City Guide が終了していたようだ。

- [City Guide Sunset](https://foursquare.com/city-guide-sunset/)

さらに今年9月にデータブローカーの Infillion が Foursquare を買収するとのニュースがあった。

- [Infillion to Acquire Foursquare, Closing the Loop Between Advertising and Real-World Results](https://infillion.com/blog/infillion-to-acquire-foursquare-closing-the-loop-between-advertising-and-real-world-results/)

Infillion 社はこんな感じの会社らしい[^i1]。

[^i1]: カリフォルニア州プライバシー保護機関（CPPA）は，2024年に Infillion がデータブローカーとしての登録と年会費の支払いを怠ったとして，54,200ドルの支払いと是正措置に合意したと[報じている](https://privacy.ca.gov/2024/12/additional-data-brokers-settle-with-cppa/ "Additional Data Brokers Settle With CPPA - privacy.ca.gov")。

{{< fig-quote type="markdown" title="Infillion to Acquire Foursquare, Closing the Loop Between Advertising and Real-World Results" link="https://infillion.com/blog/infillion-to-acquire-foursquare-closing-the-loop-between-advertising-and-real-world-results/" lang="en" >}}
Infillion is the first fully composable advertising platform, built to reduce complexity, waste, and opacity in digital media. Formed by unifying category-defining technologies including MediaMath, TrueX, Gimbal, Drawbridge, and Catalina, Infillion connects demand, supply, data, and creative in a single platform with the transparency of the open web and the scale of walled gardens. Powered by deterministic purchase intelligence and agentic-ready infrastructure, Infillion supports managed service, self-service, and automated execution for media buyers, sellers, agencies, retailers, and enterprise clients.
{{< /fig-quote >}}

んー。
そのせいで [Placemaker Tools][Foursquare Placemaker Tools] への案内メールが増えたのかねぇ。

買収話は気になるところではあるが， [Swarm] は再開してもいいかもなぁ。
[Placemaker Tools][Foursquare Placemaker Tools] はどうするか...

## 主権 AI {#sovereign-ai}

起点は，先月のこの記事から。

- [‘You want to control your own AI’: Pegasus spyware founder tells governments to push for sovereign AI and national security over reliance on US AI](https://www.techradar.com/pro/security/you-want-to-control-your-own-ai-pegasus-spyware-founder-tells-governments-to-push-for-sovereign-ai-and-national-security-over-reliance-on-us-ai)
- [「自分たちのAIは自分たちで管理すべき」――Pegasus開発元の創業者、各国政府にソブリンAIの推進を促す](https://blackhatnews.tokyo/archives/153848)

Pegasus は NSO Group が政府機関向けに販売していたスマートフォン用のスパイウェアで[^p1]，著者の Shalev Hulio 氏は NSO の共同創業者で CEO も務めていた。
現在の彼は NSO 離脱後 [Dream](https://dreamgroup.com/ "Sovereign AI and National Cyber-Defense Platform for Nation States | Dream") を立ち上げ CEO におさまっている。

[^p1]: Pegasus は，端末に侵入すれば持ち主が知らないうちに生活や通信を広く覗ける強力な監視ツールとして機能する。実際にアラブ首長国連邦の人権活動家が標的になったり WhatsApp のユーザに大規模な攻撃を仕掛けた事例がある。 WhatsApp の件については2025年に NSO に対して損害賠償を命じる陪審評決が出ている。最初に立ち上げた NSO もその後に立ち上げた Dream もイスラエル発のサイバー企業である。

{{< fig-quote type="markdown" title="Israeli cyber startup Dream raises $260 million, valued at $3 billion | Reuters" link="https://www.reuters.com/technology/israeli-cyber-startup-dream-raises-260-million-valued-3-billion-2026-06-18/" lang="en" >}}
Unlike NSO, which licensed ⁠its Pegasus software to sovereign states and law enforcement to fight terrorism and serious ​crime, Dream focuses on protecting governments and critical infrastructure - such as water, oil and ​gas facilities - from cyber attacks. It unveiled its sovereign AI platform Atlas on Wednesday.
{{< /fig-quote >}}

それで最初の記事に戻るわけやね。
なんとなく「お前がゆーな！」って気分なんだけど，深くはツッコむまい。
広告記事というのはこういうものだし。

ここのところ「デジタル主権」とか「AI 主権」とかいった話が頻繁に聞こえてきて姦しい。
EFF も昨今の情勢に警告を発しているようだ。

- [「デジタル主権」とは何か、それは何をもたらしうるのか » p2ptk[.]org](https://p2ptk.org/privacy/5659)
- [デジタル主権の要諦は「AI振興」でも「AIデータセンターを立てまくる」ことでもない » p2ptk[.]org](https://p2ptk.org/ai/5668)
- [AI規制は報復的であってはならないし、AI企業の誇大広告に踊らされてもいけない » p2ptk[.]org](https://p2ptk.org/freedom-of-speech/5663)

Cory Doctorow 氏もいつもの調子でぼやいてるな（笑）

- [「デジタル主権」にAIを絡めてくる連中にはウンザリさせられる » p2ptk[.]org](https://p2ptk.org/monopoly/5665)

一方で OpenAI の CEO は「[個人 PC のメモリは無駄](https://japan.cnet.com/article/35254269/ "個人PCで動く「ローカルAI」はメモリの無駄？ OpenAIアルトマン氏「クラウドの方が効率的」 - CNET Japan")」とか日本の土地バブル時代の地上げ屋みたいなことを言ってるし。

最近は，かなりマジで

{{< div-gen class="center caution" type="markdown" >}}
弾けろ！ AI バブル！！
{{< /div-gen >}}

と思っている。

[ZonePane]: https://twitpane.com/ "ZonePane - マルチSNSクライアント"
[mixi2]: https://mixi.social/ "mixi2"
[Foursquare]: https://foursquare.com/ "Location Intelligence Platform | Foursquare"
[Foursquare Placemaker Tools]: https://ja.foursquare.com/placemakers/ "Foursquare Placemaker Tools"
[Swarm]: https://swarmapp.com/ "Check in and track places | Swarm by Foursquare"

## 参考

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{< linkcard "80446738630e3d6e2fd656927bc64838731bc1e6" >}} <!-- THIS IS FOR EVERYONE Tim Berners-Lee -->
