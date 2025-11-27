+++
title = "”Pay-To-Crawl Issue Brief” の翻訳を公開した"
date =  "2025-11-27T21:49:43+09:00"
description = "私は機械のメンターじゃないぞ（笑）"
image = "/images/attention/kitten.jpg"
tags = ["creative-commons", "artificial-intelligence", "engineering"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[cc-licenses]({{< rlnk "/cc-licenses/" >}}) セクションに「[「クローリング課金」問題の概要]({{< ref "/cc-licenses/07-pay-to-crawl-issue-brief.md" >}})」を公開した。

Bluesky の TL で

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:mocbsa4drcbiiqnc7cjbve2u/app.bsky.feed.post/3m6krtdczdu2o" data-bluesky-cid="bafyreidzxuiincd7whzsuo6uznwsixe7mku6ovrtlwmp5um7betzz7t7qe" data-bluesky-embed-color-mode="system"><p lang="">Pay-to-crawl systems have emerged as a possible way to help sustain the creation of content online in the wake of AI. However, overbroad use of pay-to-crawl could represent a shift towards a more tightly controlled and monetised content ecosystem. 

Read the full brief for more!

buff.ly/h0wdhXS<br><br><a href="https://bsky.app/profile/did:plc:mocbsa4drcbiiqnc7cjbve2u/post/3m6krtdczdu2o?ref_src=embed">[image or embed]</a></p>&mdash; Creative Commons (<a href="https://bsky.app/profile/did:plc:mocbsa4drcbiiqnc7cjbve2u?ref_src=embed">@creativecommons.bsky.social</a>) <a href="https://bsky.app/profile/did:plc:mocbsa4drcbiiqnc7cjbve2u/post/3m6krtdczdu2o?ref_src=embed">November 27, 2025 at 6:19 AM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

ポストを見かけて，気まぐれに「久しぶりに翻訳してみるか」と思い立ったのだ。
具体的には [GitHub Copilot] (model: GPT-5 mini を選択[^g1]) に翻訳させたらどうなるかと思って。

[^g1]: [GitHub Copilot] のモデルに GPT-5 mini を選択した理由は特にない。強いて言うなら，ゼロコストで使いたい放題だからと，個人的に使い慣れてるから，かな。

まずは Copilot に「この PDF からテキストを抽出できるか？」と聞いたら「`pdftotext` コマンドを使います。実行しますか？」と言ってきた。
いや，ソレかい！ まぁ AI にテキストを**創作**されるよりはマシか。
じゃあ，実行してくれ。

Copilot にやってもらったのは

1. PDF からテキストを抽出（`pdftotext` コマンドを使用）
2. テキストを markdown 形式に整形（別ファイルに保存）
3. Markdown テキストを日本語に翻訳（別ファイルに保存）

というところまで。
実行結果はこんな感じ。

1. [`Pay-to-Crawl-Issue-Brief-Nov-2025.txt`]({{< rlnk "/cc-licenses/07-pay-to-crawl-issue-brief/Pay-to-Crawl-Issue-Brief-Nov-2025.txt" >}})
2. [`Pay-to-Crawl-Issue-Brief-Nov-2025.md`]({{< rlnk "/cc-licenses/07-pay-to-crawl-issue-brief/Pay-to-Crawl-Issue-Brief-Nov-2025.md.txt" >}})
3. [`07-pay-to-crawl-issue-brief.draft.md`]

綺麗に markdown に整形してくれたのは偉い！

翻訳された markdown テキストを（原文と比べながら）ざっと眺めてみたのだが「意訳が過ぎん？」ってなったので [Kagi Transrate] にも翻訳させてみた。
最終的に [`07-pay-to-crawl-issue-brief.draft.md`] には原文と Copilot 訳と [Kagi Transrate] 訳を併記させている。

ここまでは簡単。
ここからが割と大変だった。
いや，私が英語不得手だから大変なんだろうけど。

傾向として Copilot の翻訳は意訳が多く，意味は大体合ってる気がするが「そんなこと書いとらんじゃろ！」とツッコむこと多し。
それに対して[Kagi Transrate] の翻訳は比較的直訳寄りなのだが，微妙な言い回しが多く「？？？」となることが多かった。
定量評価ではなく，私の印象による評価で申し訳ないが，機械による翻訳が概ね７〜8割程度の出来で，残り2〜3割を人間がチマチマ修正していく感じ。
この「残り2〜3割」が果てしなく大変なのだ。

開発の成功事例で「生成 AI で生産性爆上がり」なんてな話を聞くようになったが，たとえ生成 AI でコードの9割を上手く生成してくれたとしても残りの1割が無限イテレーションの地獄だし，そこが全体の品質を決定づけることも往々にしてあると思うのだが，成功者はどうやってるのだろう。
そもそも人間側は機械の出力を精査するところから始まるので，途中で何度か「何やってんだ，俺」ってなったよ。
私は機械のメンターじゃないぞ（笑）

機械相手に「コミュニケーション」するのって虚しくない？ 人間のチームならお互いの技術やスキルの向上を期待してやり取りできるけど，機械相手にプロンプト等をチューニングしても，機械自体が成長するわけではない。
しかもモデルごとにクセが違うから応用があまり効かない。
ぶっちゃけ，機械とのやり取りや成果物は「使い捨て」に等しいのだ。

これでホンマに（プロジェクトやミッション全体で）生産性が上がるの？ 機械がスパスポスパってやってくれるところだけ見て「生産性爆上がり」にはならんと思うのだが。

...という感想を抱きながらの作業だった。
感想というか愚痴だな，コレは。

真面目に生成 AI の勉強したほうがいいんだろうなぁ...

[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Kagi Transrate]: https://translate.kagi.com/ "Kagi Transrate"
[`07-pay-to-crawl-issue-brief.draft.md`]: {{< rlnk "/cc-licenses/07-pay-to-crawl-issue-brief/07-pay-to-crawl-issue-brief.draft.md.txt" >}}
