+++
title = "Bluesky における新しい認証方法"
date =  "2025-04-22T10:51:09+09:00"
description = "Bluesky 自身が個々人のユーザアカウントを認証（verification）するのではないようだ。"
image = "/images/attention/kitten.jpg"
tags = [ "bluesky", "security", "risk", "management" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

また Bluesky が面白いことを始めたらしい。

- [A New Form of Verification on Bluesky - Bluesky](https://bsky.social/about/blog/04-21-2025-verification)

この記事について [Kagi Assistant] に要約してもらった。
LLM の言うことだから鵜呑みにしないで原文をチェックしてね。

{{< fig-quote type="markdown" >}}
Blueskyは2025年4月21日に新しい認証の仕組みを発表しました。主なポイントは以下の通りです。

1. **青いチェックマークによる認証の導入**
   - これまでのドメインをユーザー名に設定する認証（ドメインハンドル）に加え、公式で著名なアカウントに対して青いチェックマークを表示する新しい認証レイヤーを導入。
   - Blueskyが積極的に認証を行い、青いチェックを付与する。

2. **Trusted Verifiers（信頼された認証機関）機能**
   - 独立した組織が直接アカウントを認証できる仕組み。
   - Trusted Verifiersは波型の青いチェックマークで表示される。
   - 例としてニューヨーク・タイムズが自社のジャーナリストに直接認証を付与可能。
   - Blueskyのモデレーションチームが認証の真偽を確認する。

3. **認証の詳細表示と設定**
   - 認証マークをタップすると、どの組織が認証したかが見られる。
   - ユーザーは設定で認証表示をオフにすることも可能。

4. **認証の申請について**
   - 現時点では直接の認証申請は受け付けていない。
   - 今後、安定化した段階で申請フォームを公開予定。

5. **自己認証の推奨**
   - 公式組織や個人は引き続きドメインをユーザー名に設定する自己認証を推奨。

この新しい認証システムは、信頼性を高め、ユーザーが誰と交流しているかをより明確にすることを目的としています[^1]。

[^1]: [A New Form of Verification on Bluesky - Bluesky](https://bsky.social/about/blog/04-21-2025-verification)
{{< /fig-quote >}}

元記事を見るに Bluesky 自身が個々人のユーザアカウントを認証（verification）するのではなく，企業・組織アカウントを verifier として Bluesky が認証した上で，認証された verifier が麾下のユーザアカウントを認証するという仕組みのようだ。

{{< fig-img-quote src="./verification-diagram.png" title="A New Form of Verification on Bluesky" link="https://bsky.social/about/blog/04-21-2025-verification" width="720" lang="en" >}}

X.509 な証明書管理の仕組み（hierarchical な信用モデル）とちょっと似てるかもね。
ユーザアカウントの認証については verifier に丸投げではなく

{{< fig-quote title="A New Form of Verification on Bluesky" link="https://bsky.social/about/blog/04-21-2025-verification" lang="en" >}}
Bluesky’s moderation team reviews each verification to ensure authenticity.
{{< /fig-quote >}}

と Bluesky も監査を行う感じだろうか。

SNS ユーザのアイデンティティを客観的に評価するのは難しい。
というか多分無理。

昔（ゼロ年代）はリアルのユーザと SNS アカウントを（実名を強制するなどして）強引に紐づけようとして，多くの批判にさらされ失敗した。
旧 Twitter はサービス提供者がアイデンティティを認証しバッヂを付与することでユーザ間の差別化を図ろうとしたが，そこから二転三転し，現在の {{% emoji "X" %}} のバッヂは単に有料ユーザであることを示しているに過ぎずアイデンティティについて何も担保しない。

Bluesky による今回の認証方法が上手く回るかどうかは（面白そうではあるが）分からない。
Verifier に対する認証・監査は簡単じゃないと思うし，恣意的な判断が入る余地もある。
それでも...

昨年末に「[Bluesky 上で起きたなりすましとサイバースクワッティング]({{< ref "/remark/2024/12/cybersquatting-on-bluesky.md" >}})」という話を書いたが，少なくともこうしたなりすましが抑制されることを期待したい。

[Kagi Assistant]: https://kagi.com/assistant "The Assistant"

## ブックマーク

- [Xそっくりの青い認証マーク導入をBlueskyが発表するも、一部で批判の声が上がる理由とは【やじうまWatch】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/yajiuma/2008789.html)

## 参考図書

{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
