+++
title = "AI がハッキングを始める日"
date =  "2021-05-15T17:12:16+09:00"
description = "日本語翻訳および解説。ありがたや"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Bruce Schneier 先生の AI ハッキングに関する一連の記事をブックマークとして挙げておく。

- [The Coming AI Hackers | Belfer Center for Science and International Affairs](https://www.belfercenter.org/publication/coming-ai-hackers)
  - [ブルース・シュナイアーが予言する「AIがハッカーになり人間社会を攻撃する日」 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20210511/coming-ai-hackers)
- [Hackers Used to Be Humans. Soon, AIs Will Hack Humanity | WIRED](https://www.wired.com/story/opinion-hackers-used-to-be-humans-soon-ais-will-hack-humanity/)
  - [When AIs Start Hacking - Schneier on Security](https://www.schneier.com/blog/archives/2021/04/when-ais-start-hacking.html)
  - [ブログ: AIがハッキングを始めるとき (ブルース・シュナイアー)](https://okuranagaimo.blogspot.com/2021/04/ai.html)

私は所謂「[シンギュラリティ]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")」にはもはや微塵も関心がないが，道具・手段としての「AI 技術」についてリスク・ベネフィットをきちんと評価して規制すべきところは規制すべき，とは考えている。

「理学」が「真理の追求」なら「工学」は「善の実装」である。
何を以って「善」とするかは議論の余地があるだろうが，エンジニアが自らの首を絞めることにならないよう気をつけなければ。

## 人と機械の責務分担

（これで思い出したが，以前に Facebook の TL に書きなぐった戯れ言を以下に記しておく）

{{< fig-quote type="markdown" title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}自立とは、仮想代理人ソフトウェアであるところのエージェントが自ら動き、誰の力も借りずに意思決定できることを言う。 [...] 一方、自律というのは哲学的な意味であり、自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう{{% /quote %}}
{{< /fig-quote >}}

この定義で考えると，今の世の中に「自律」機械はまだ登場していないことになる。
もしそのようなものが登場すればシンギュラリティ（特異点）到達ということになるのだろう。
まっ（少なくとも「現在」からの延長線上の未来では）ありえへんけど（笑）

もう少し踏み込んで考えるなら，機械に「自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができる」ことを期待してはいけないし，そのような期待を前提としたルール・メイキングは意味がない，と思っている。

であるならルールの適用と運用は人間側に任せるべきで，機械はあくまで運用補佐として振る舞うに留める。
つまり機械に「最終判断」させるべきではない，ということだ。

ハードでもソフトでもこうした「人と道具・機械との責務分担」をどう設計するかというのはかなり重要で，現状として人の側が最終判断を行うのであれば，システム上のルール設定も（もっと広く）法規制も人の側に寄せた設計をすべきと思っている。

## 参考図書

{{% review-paapi "B01J1I8PRQ" %}} <!-- 社会は情報化の夢を見る -->
{{% review-paapi "B00JB3F73M" %}} <!-- マップス -->
{{% review-paapi "B00F5P454W" %}} <!-- キカイダー02 -->
