+++
title = "Secretary as a Service"
date =  "2026-05-31T12:43:36+09:00"
description = "既に AI サービスはメタクソ化が始まってると思う。早めに逃げれるようにしておかないと。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "vscode", "github", "tools", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## Agents Window

最近気がついたのだが [VS Code] 1.120 から [Agents Window] なるものが追加されたそうだ。

- [VS Code 1.120リリース、Agentsウィンドウの安定版向けプレビューを開始 | gihyo.jp](https://gihyo.jp/article/2026/05/vscode-1.120)
- [VS Codeの「Agents window」とは？複数AIエージェントを並列で動かす新機能をやさしく解説します｜しゃり](https://note.com/shali_note/n/n47d9fa3e4a46)

仕事でコードを書くときは分担して作業するものなのであまり意識することはないが，独りで作業する場合，特に [Go] で複数のパッケージを同時に弄る場合は頭の中がとっ散らかってしまうことがある。
我ながら歳をとったなぁとは思うけど。

というわけで，少しでも作業が軽減できればと [Agents Window] を試そうとしたんだけど，これって既存の AI エージェントとのセッションを他セッションに引き継げないのな。
「どうにかならん？」と，それこそ [Copilot][GitHub Copilot] に訊いてみたんだが，セッション履歴のインポート・エクスポート機能はあるが（つまりバックアップ・リストアはできる？），他セッションに引き継いで fork することはできないと言われた。
できることは AI に引き継ぎ資料を書かせて，それを別のセッションに読み込ませるくらいだそうな。

この辺の話を Mastodon/Bluesky で愚痴ったのだが

{{< fig-gen >}}
<blockquote class="mastodon-embed" data-embed-url="https://mstdn.world/@rino/116650070084518151/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://mstdn.world/@rino/116650070084518151" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M63 45.3v-20c0-4.1-1-7.3-3.2-9.7-2.1-2.4-5-3.7-8.5-3.7-4.1 0-7.2 1.6-9.3 4.7l-2 3.3-2-3.3c-2-3.1-5.1-4.7-9.2-4.7-3.5 0-6.4 1.3-8.6 3.7-2.1 2.4-3.1 5.6-3.1 9.7v20h8V25.9c0-4.1 1.7-6.2 5.2-6.2 3.8 0 5.8 2.5 5.8 7.4V37.7H44V27.1c0-4.9 1.9-7.4 5.8-7.4 3.5 0 5.2 2.1 5.2 6.2V45.3h8ZM74.7 16.6c.6 6 .1 15.7.1 17.3 0 .5-.1 4.8-.1 5.3-.7 11.5-8 16-15.6 17.5-.1 0-.2 0-.3 0-4.9 1-10 1.2-14.9 1.4-1.2 0-2.4 0-3.6 0-4.8 0-9.7-.6-14.4-1.7-.1 0-.1 0-.1 0s-.1 0-.1 0 0 .1 0 .1 0 0 0 0c.1 1.6.4 3.1 1 4.5.6 1.7 2.9 5.7 11.4 5.7 5 0 9.9-.6 14.8-1.7 0 0 0 0 0 0 .1 0 .1 0 .1 0 0 .1 0 .1 0 .1.1 0 .1 0 .1.1v5.6s0 .1-.1.1c0 0 0 0 0 .1-1.6 1.1-3.7 1.7-5.6 2.3-.8.3-1.6.5-2.4.7-7.5 1.7-15.4 1.3-22.7-1.2-6.8-2.4-13.8-8.2-15.5-15.2-.9-3.8-1.6-7.6-1.9-11.5-.6-5.8-.6-11.7-.8-17.5C3.9 24.5 4 20 4.9 16 6.7 7.9 14.1 2.2 22.3 1c1.4-.2 4.1-1 16.5-1h.1C51.4 0 56.7.8 58.1 1c8.4 1.2 15.5 7.5 16.6 15.6Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @rino@mstdn.world</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://mstdn.world/" async src="https://mstdn.world/embed.js"></script>
{{< /fig-gen >}}

みたいなポストをいただいたり。
でも，ここに書かれている Claude のやり方も裏技っぽいよねぇ。
他に「[ZedのMCPサーバでコンテキスト共有したい](https://rio.st/2026/02/26/zed%e3%81%aemcp%e3%82%b5%e3%83%bc%e3%83%90%e3%81%a7%e3%82%b3%e3%83%b3%e3%83%86%e3%82%ad%e3%82%b9%e3%83%88%e5%85%b1%e6%9c%89%e3%81%97%e3%81%9f%e3%81%84/ "ZedのMCPサーバでコンテキスト共有したい - rio.st")」みたいな運用もあるよ，と教えてもらった。

セッションの複製や fork は必須の機能だと思うんだけど，なんで [Copilot][GitHub Copilot] の標準機能にないの？ やっぱそろそろ [VS Code] や [Copilot][GitHub Copilot] は卒業すべきかねぇ。
だからといって，個人利用で Claude Code とか使う気はないが。

## 秘書としての AI エージェント

他の記事でも言ってるけど，コーディング支援として使う AI はどちらかというとペアプログラミングしている感じ。
それも新人くんとのペアプロ（笑） 最初の頃は「自分で書いたほうが速い」って感じだったんだけど，反応を見ながら指示を出すことで，なんとなく AI の性格のようなものが分かってきた。
[GitHub Copilot] の Pro 契約で使える主なモデルはこんな感じ（個人の意見です。あしからず）。

- **GPT-5-mini**: コストゼロで使えるのが最大のメリット。生真面目なおバカ。コード補完やコメント補完では重宝するが，エージェント・モードだと最初から明後日の方向に全力疾走する。「けもフレ」のアライさんみたいなモデル
- **GPT-5.3-Codex**: 個人で使うエージェントとしては一番バランスが取れている。こちらの意見に[追従（ついしょう）]({{< relref "./reading-living-with-ai.md#sycophancy" >}} "追従する AI は向社会的な姿勢を減退させる？")せず，ちゃんと反論できる。人間のペースに合わせてくれる盲導犬モデル
- **Claude Sonnet 4.6**: 計画立案は上手いが，一度走り出すと融通が利かず失敗するまで突っ走るチョロQモデル。コード分析等で無闇にトークンを消費してる印象。金に糸目をつけず地球資源を濫用する富裕層向き

エージェントの使い方がちょっと変わったきっかけになったが [`goark/pa-api`] パッケージに来た PR の評価に [VS Code] の [Copilot][GitHub Copilot] を使ったこと。
PA-API から Creator API に移行するための大きな変更だったので，自分だけで評価・判断するのがちょっと不安だったのだ。

実をいうと PR ページでアサインする AI レビューが使いものにならなかったのであまり期待してなかったのだが，思った以上に使えた。
私は，上述の性格診断から GPT-5.3-Codex を常用しているのだが， PR に対するお互いの評価を突き合わせて「議論」が成立したことに驚いた。
結城浩さんの『AIと生きる』に出てくる「[AI との対話]({{< relref "./reading-living-with-ai.md" >}} "『AIと生きる』を読む")」ってこういう感じなんだなぁ，と実感したのだった。

これで我が家では [Copilot][GitHub Copilot] はペアプロ相手の新人くんから秘書に格上げになった。
AI で秘書的な役割を演じるものとして思いつくのはグレッグ・イーガンの『[万物理論](https://amzn.to/4viO0LY "万物理論 | グレッグ・イーガン, 山岸 真 |本 | 通販 | Amazon")』に出てくるシジフォスだな。
私としてはあーゆーのが理想なんだけどねぇ。

## サービスとしての秘書

個人で [Go] 製ツールを書くときは複数の自作パッケージを同時に扱うことが多い。
そういう意味で [Agents Window] は便利そうなのだが，前に述べたようにセッションの複製や fork ができないのはいただけない。
あと，複数のセッションを統括する AI セッションが要るよね，たぶん（セッション間の通信？）。
秘書はなんぼあってもいいですからね（by ミルクボーイ）。

で，思ったんだけど，これ [VS Code] から分離したほうがよくね？ そもそもエディタや IDE にエージェント機能を組み込むのは筋が悪いと思うのだ。
だからこそ CLI インタフェースの AI エージェントが人気なんだろうけど GUI デスクトップでいまさら CUI ターミナル上で動かす意味はないよね。
ブラウザ上で動かすのはもっとないけど。

というわけで [Agents Window] には密かに期待を寄せている。
どう発展するのか分からないけど。

あと，やっぱローカル LLM は真面目に勉強しないといけないなぁ。
既に AI サービスは[メタクソ化](https://en.wikipedia.org/wiki/Enshittification "Enshittification - Wikipedia")が始まってると思う。
早めに逃げれるようにしておかないと。

## ブックマーク

- [VS Code 1.121 で何が変わったのか？ Agents・Mermaid・HTMLプレビュー・ターミナル最適化を整理する - Qiita](https://qiita.com/yasuo-dev/items/88cc099d3e133ab87cc1)

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

[Go]: https://go.dev/
[Agents Window]: https://code.visualstudio.com/docs/copilot/agents/agents-window "Use the Agents window (Preview)"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - The open source AI code editor | Your home for multi-agent development"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[`goark/pa-api`]: https://github.com/goark/pa-api "goark/pa-api: Go client for the Amazon Creators API"

## 参考

{{% review-paapi "4488711022" %}} <!-- 万物理論 -->
