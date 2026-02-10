+++
title = "Kagi News あかよろし"
date =  "2026-01-27T20:33:44+09:00"
description = "Kagi News は生成 AI をキュレーターとして上手く使ってる印象"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "artificial-intelligence", "Kagi", "media" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

入院生活が暇でねぇ。
リハビリ以外はほとんどすることないし。
「たくさん歩け」と言われてるけど，病院内を目的もなくうろついてたら不審者だし。

というわけで，前から気になっていた [Kagi News] で遊んでみることにした。

- [Introducing Kagi News | Kagi Blog](https://blog.kagi.com/kagi-news)

## Kagi News の導入

Web ブラウザでの初期表示はこんな感じ。

{{< fig-img-quote src="./kagi-news-web.png" title="Kagi News" link="https://news.kagi.com/" width="1056" lang="en">}}

[Kagi News] はまだベータ版で，今のところアカウント無しで使えるっぽい[^k1]。
ただ，この見た目なので日本語には対応していないと今まで思っていた。
よく見ると設定で表示言語は変えられるようだ。

[^k1]: Kagi アカウントを持っている（かつサインインしている）場合は [Kagi News] の設定等をデバイス間で同期するらしい（スマホアプリを除く）。この機能は設定で解除できる。なお Kagi アカウントには（機能制限付きの）無料プランもあるので試してみるのもいいだろう。

{{< fig-img-quote src="./settings-language.png" title="Settings - Language" link="./settings-language.png" width="1056" lang="en">}}

“Interface Language” を日本語に変更し “Content Language” を “Custom” とした上で “Main Language” を日本語にすれば全て日本語表示になる。

{{< fig-img-quote src="./settings-language-custom.png" title="Settings - Language" link="./settings-language-custom.png" width="1056" lang="en">}}

“Languages I Speak” には “Main Language” を含む複数の言語を指定できる。
例えば “Main Language” を日本語にして “Languages I Speak” に（日本語の他に）英語を指定すれば，日本語と英語の記事はそのまま表示され，それ以外の言語の記事は日本語に翻訳されて表示される。
翻訳は AI で行っている（おそらく Kagi オリジナルの Quick モデル）。
翻訳精度はそこそこ良い感じかなぁ。

ニュース・カテゴリはあらかじめ用意されているものしか選べないが[^k2]，日本をはじめとする国・地域のニュースや科学やスポーツといったジャンルも選べる。
たとえばこんな感じに選んでみた。

[^k2]: カテゴリ追加の提案は GitHub Issue を通して可能。

{{< fig-img-quote src="./settings-categories.png" title="Settings - Categories" link="./settings-categories.png" width="1056" lang="en">}}

これでヘッドラインはこんな感じになる。

{{< fig-img-quote src="./kagi-news-web-2.png" title="Kagi News (日本語)" link="https://news.kagi.com/" width="1056" lang="en">}}

先頭の記事（Story）をクリックして詳細を表示してみる。

{{< fig-img-quote src="./kagi-news-web-detail.png" title="Kagi News detail" link="https://news.kagi.com/s/vcav4v" width="1056" lang="en">}}

こんな感じに 概要，出典，ハイライト，視点，歴史的背景 などと見出しが続く。
これらの見出しの順番や表示の有無は設定で変更できる。

{{< fig-img-quote src="./settings-story.png" title="Settings - Story" link="./settings-story.png" width="1056" lang="en">}}

[Kagi News] はスマホアプリも提供している。

- [Kagi News - Apps on Google Play](https://play.google.com/store/apps/details?id=com.kagi.news)
- [Kagi News App - App Store](https://apps.apple.com/us/app/kagi-news/id6748314243)

Web 版とほぼ同等の機能なので，私は入院中の暇つぶしに主にこちらを使っている。

## ニュースは壊れているから

[Kagi News] は昨今のニュースについてこう語っている。

{{< fig-quote type="markdown" title="Why Kagi News? Because news is broken." link="https://news.kagi.com/about" lang="en" >}}
Driven by relentless ad monetization, news has [become](https://gwern.net/doc/culture/2010-dobelli.pdf) mental junk food - an endless stream of clickbait that destroys our ability to [think deeply and clearly](http://www.aaronsw.com/weblog/hatethenews). We can do better, and create what news should have been all along - pure, essential information that respects your intelligence and time.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="なぜkiteなのか？それはニュースは壊れているから。" link="https://news.kagi.com/about" >}}
広告収益化の追求により、ニュースは[精神的なジャンクフード](https://gwern.net/doc/culture/2010-dobelli.pdf)と化しました - 深く明確に考える能力を破壊するクリックベイトの延々と続く流れです。私たちはより良い方法を見つけ、ニュースが本来あるべき姿 - あなたの知性と時間を尊重する、純粋で本質的な情報 - を創造できます。
{{< /fig-quote >}}

「ニュースが本来あるべき姿」のための [Kagi News] のアプローチはこうだ。

{{< fig-quote type="markdown" title="Understanding our approach" link="https://news.kagi.com/about" lang="en" >}}
Kagi News reads public RSS feeds of thousands of (community-curated) world-wide news sources and utilizes AI to distill them into one perfect daily briefing. You get every critical perspective and timeline in just 5 minutes. That's it. No endless scrolling. No attention hijacking. Because we deserve better.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="私たちのアプローチの理解" link="https://news.kagi.com/about" >}}
Kagiニュースは、（コミュニティが厳選した）世界中の何千ものニュースソースの公開RSSフィードを読み取り、AIを活用して一つの完璧なデイリーブリーフィングに要約します。わずか5分であらゆる重要な視点とタイムラインを把握できます。それだけです。無限のスクロールはありません。あなたの注意を奪うものもありません。私たちには、もっと良い体験がふさわしいからです。
{{< /fig-quote >}}

具体的には

- 更新は1日1回（12時UTC，日本時間で午後9時頃）
- AI によるキュレーション
- カテゴリごとに最大12件の記事を表示（設定で変更可能）
  - アテンションエコノミーからの解放 - 無限のスクロールの中ではなく，今この瞬間に存在する
  - ニュース消費の自然な終点を作り出す
- 意見よりも事実と視点に焦点を当てる
  - 量より質と深さを重視
- プライバシーへの配慮

といったところ。

デメリットもあって，特に日本ユーザの場合はどうしても即時性に欠ける（ニュースがほぼ1日遅れる）。
また，ワールドワイドなニュースが優先され，ローカルなニュースは拾われない感じである（ローカルニュースほどニュースソースが減るので仕方がないのだが）。
この辺は「量より質と深さを重視」する方針とトレードオフだと思うので，ツールの使い分けが必要かもしれない。

これまで見たように  [Kagi News] は生成 AI をキュレーターとして上手く使ってる印象である。
創作や翻案の濫造（乱造）で Web を汚すその辺のサービスに比べれば随分と真っ当な AI の使い方だと思う。
Kagi はもともと[検索サービス]({{< ref "/remark/2024/06/kagi-search.md" >}} "Kagi Search を試してみる 〜 検索サービスも有料の時代？")を提供している会社ということもあり，情報収集と評価に関して AI はユーザの活動を横取りするのではなく，執事的な立ち位置でユーザと協働できる。
個人的には理想的な AI の使い方なんだよなぁ。

## つまり

[Kagi News] は「あ&#x1b019;よろし[^a1]」ということやね。

[^a1]: [^a1]: 「あかよろし」は「めっちゃいい！」くらいの意味（笑）

おあとがよろしいようで。

[Kagi News]: https://news.kagi.com/ "Kagi News"
<!-- eof -->
