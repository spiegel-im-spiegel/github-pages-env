+++
title = "Typst について"
date =  "2025-02-20T22:09:42+09:00"
description = "そろそろ Typst の勉強を始めようか / Typst とは / ちょっと触ってみる"
image = "/images/attention/tools.png"
tags  = [ "typst", "tex", "markdown" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

## そろそろ Typst の勉強を始めようか

Bluesky の TL を眺めてて気がついたのだが [Typst] の Bluesky アカウントがあるらしい。
Mastodon のアカウントもあるな。
ブログもある。

- [Typst (@typst.app) — Bluesky](https://bsky.app/profile/typst.app)
- [Typst (@typst@mastodon.social) - Mastodon](https://mastodon.social/@typst)
- [Typst: Blog](https://typst.app/blog/)

最新の[ブログ記事](https://typst.app/blog/2025/typst-0.13 "Typst 0.13 is out now – Typst Blog")を見ると，バージョン 0.13 のハイライトとして

{{< fig-quote type="markdown" title="Typst 0.13 is out now – Typst Blog" link="https://typst.app/blog/2025/typst-0.13" lang="en" >}}
- [Paragraphs and first line indent](https://typst.app/blog/2025/typst-0.13#paragraphs-and-first-line-indent)
- [Better-looking outlines](https://typst.app/blog/2025/typst-0.13#better-looking-outlines)
- [New curves](https://typst.app/blog/2025/typst-0.13#new-curves)
- [Files and bytes](https://typst.app/blog/2025/typst-0.13#new-curves)
- [Generating images](https://typst.app/blog/2025/typst-0.13#generating-images)
- [Faster plugins](https://typst.app/blog/2025/typst-0.13#faster-plugins)
- [Single-letter strings in math](https://typst.app/blog/2025/typst-0.13#single-letter-strings-in-math)
- [Font coverage control](https://typst.app/blog/2025/typst-0.13#font-coverage-control)
- [PDF file embedding](https://typst.app/blog/2025/typst-0.13#pdf-file-embedding)
- [A first look at HTML export](https://typst.app/blog/2025/typst-0.13#a-first-look-at-html-export)
- [Migrating to Typst 0.13](https://typst.app/blog/2025/typst-0.13#migrating)
- [Community Call](https://typst.app/blog/2025/typst-0.13#community-call)
{{< /fig-quote >}}

と書かれていた。
どうやらこのバージョンから HTML 形式へのエクスポート機能を搭載していくらしい（まだ完全ではないようだが）。
またフォント制御もよくなっているみたいで（たとえば日本語だと仮名・漢字と英数字では別のフォントを割り当てたくなったりするだろうし）ちょっと期待できる。

というわけで新たに [typst]({{< rlnk "/typst/" >}} "Typst のお勉強") セクションを設け，お勉強を通して [Typst] で出来ること出来ないことを探ってみることにする。

## Typst とは

“Typst” は「タイプスト」と読むらしい。

{{< fig-quote type="markdown" title="typst/typst: A new markup-based typesetting system that is powerful and easy to learn." link="https://github.com/typst/typst" lang="en" >}}
IPA: /taɪpst/. "Ty" like in **Ty**pesetting and "pst" like in Hi**pst**er. When writing about Typst, capitalize its name as a proper noun, with a capital "T".
{{< /fig-quote >}}

イケてる組版ソフトってこと？

簡単に言うと [Typst] は独自のマークアップ言語とそれを解釈して組版を行うツールおよびサービスを指す。
組版ソフトとしては定番の $\mathrm{\LaTeX}$ や Markdown テキストと比較して以下のような特徴があるらしい。

- マークアップについては $\mathrm{\LaTeX}$ より簡易で Markdown より高度な表現が出来る
  - マクロを使った制御が可能
- 組版結果を PDF形式のファイルに直接出力できる
- 組版結果を HTML 形式へエクスポートができる（0.13 より）
- コマンドライン・ツールは Rust 言語で組まれている

コマンドライン・ツールは [GitHub リポジトリ][typst/typst] でコードとバイナリが公開されている。
ライセンスは Apache-2。

しかし [Typst] ではドキュメント環境を Web サービスとして提供していて，どうやらこちらがメインっぽい。

{{< fig-img src="./signup-3.png" title="Typst: dashboard" link="https://typst.app/" width="1540" >}}

{{< fig-img src="./pricing.png" title="Typst: Pricing" link="https://typst.app/pricing/" width="1099" >}}

とりあえず言語仕様や組版機能を確認するだけなら Free でも問題なさそうなので，しばらくはこれを使うことにする（Free 版を使うだけなら決済情報の入力は不要）。
まぁ $\mathrm{\LaTeX}$ にも[クラウドサービス](https://cloudlatex.io/ "Cloud LaTeX | Build your own LaTeX environment, in seconds")があったりするし，こういうところで収益を確保するのはアリかもしれない（実際に収益があるかどうかは知らないが）。

なお，コマンドライン・ツールについては Windows 環境なら [Winget][winget] または [Scoop] でインストールできる。

```text
$ winget search typst
名前     ID                    バージョン 一致       ソース
-----------------------------------------------------------
Typst    Typst.Typst           0.12.0                winget
typstyle Enter-tainer.typstyle 0.12.15    Tag: typst winget

$ scoop search typst
Results from local buckets...

Name  Version Source Binaries
----  ------- ------ --------
typst 0.13.0  main
```

[Winget][winget] のほうがタイムラグがあるのかな？ 手元の Ubuntu 機では Snap でインストール可能だった。

```text
$ snap search typst
Name   Version  Publisher  Notes  Summary
typst  0.13.0   reknih     -      A new markup-based typesetting system that is powerful and easy to learn.
```

macOS は [Homebrew からインストール可能](https://formulae.brew.sh/formula/typst "typst — Homebrew Formulae")かな。
知らんけど。

VS Code 用の拡張機能や vim 用のプラグインもあるらしい。

## ちょっと触ってみる {#ope}

先程のダッシュボードを開いて {{% keytop %}}`Empty document`{{% /keytop %}} を押してみる。

{{< fig-img src="./empty-document.png" title="Create project" link="./empty-document.png" width="1489" >}}

適当に “`sample`” とかで始めてみるか。
以下の文言を入力する。

```typst
= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

構文について今は後回しにして，組版結果はこんな感じ。

{{< fig-img src="./sample.png" title="Project: sample" link="./sample.png" width="1263" >}}

うわぁ。
これは酷い（笑）

まずエディタのフォントが `"Cascadia Mono", monospace` になってるな。

{{< fig-img src="./settings.png" title="Settings" link="./settings.png" width="1264" >}}

画面左上に {{% keytop %}}↓{{% /keytop %}} ボタンがあるので，これを押下して PDF ファイルを[ダウンロード](./sample.pdf)し，プロパティを見てみる。

{{< fig-img src="./property.png" title="property" link="./property.png" >}}

なんか中華フォントっぽい。
句読点も明らかにおかしいし。
やっぱ，まともに日本語のドキュメント環境が欲しいならローカルで作らないとダメかな。
ちょっと考えてみよう。

### 【追記 2025-02-21】

Mastodon で教えていただきました。

{{< fig-gen >}}
<blockquote class="mastodon-embed" data-embed-url="https://social.vivaldi.net/@adbird/114041736926058126/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://social.vivaldi.net/@adbird/114041736926058126" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M74.7135 16.6043C73.6199 8.54587 66.5351 2.19527 58.1366 0.964691C56.7196 0.756754 51.351 0 38.9148 0H38.822C26.3824 0 23.7135 0.756754 22.2966 0.964691C14.1319 2.16118 6.67571 7.86752 4.86669 16.0214C3.99657 20.0369 3.90371 24.4888 4.06535 28.5726C4.29578 34.4289 4.34049 40.275 4.877 46.1075C5.24791 49.9817 5.89495 53.8251 6.81328 57.6088C8.53288 64.5968 15.4938 70.4122 22.3138 72.7848C29.6155 75.259 37.468 75.6697 44.9919 73.971C45.8196 73.7801 46.6381 73.5586 47.4475 73.3063C49.2737 72.7302 51.4164 72.086 52.9915 70.9542C53.0131 70.9384 53.0308 70.9178 53.0433 70.8942C53.0558 70.8706 53.0628 70.8445 53.0637 70.8179V65.1661C53.0634 65.1412 53.0574 65.1167 53.0462 65.0944C53.035 65.0721 53.0189 65.0525 52.9992 65.0371C52.9794 65.0218 52.9564 65.011 52.9318 65.0056C52.9073 65.0002 52.8819 65.0003 52.8574 65.0059C48.0369 66.1472 43.0971 66.7193 38.141 66.7103C29.6118 66.7103 27.3178 62.6981 26.6609 61.0278C26.1329 59.5842 25.7976 58.0784 25.6636 56.5486C25.6622 56.5229 25.667 56.4973 25.6775 56.4738C25.688 56.4502 25.7039 56.4295 25.724 56.4132C25.7441 56.397 25.7678 56.3856 25.7931 56.3801C25.8185 56.3746 25.8448 56.3751 25.8699 56.3816C30.6101 57.5151 35.4693 58.0873 40.3455 58.086C41.5183 58.086 42.6876 58.086 43.8604 58.0553C48.7647 57.919 53.9339 57.6701 58.7591 56.7361C58.8794 56.7123 58.9998 56.6918 59.103 56.6611C66.7139 55.2124 73.9569 50.665 74.6929 39.1501C74.7204 38.6967 74.7892 34.4016 74.7892 33.9312C74.7926 32.3325 75.3085 22.5901 74.7135 16.6043ZM62.9996 45.3371H54.9966V25.9069C54.9966 21.8163 53.277 19.7302 49.7793 19.7302C45.9343 19.7302 44.0083 22.1981 44.0083 27.0727V37.7082H36.0534V27.0727C36.0534 22.1981 34.124 19.7302 30.279 19.7302C26.8019 19.7302 25.0651 21.8163 25.0617 25.9069V45.3371H17.0656V25.3172C17.0656 21.2266 18.1191 17.9769 20.2262 15.568C22.3998 13.1648 25.2509 11.9308 28.7898 11.9308C32.8859 11.9308 35.9812 13.492 38.0447 16.6111L40.036 19.9245L42.0308 16.6111C44.0943 13.492 47.1896 11.9308 51.2788 11.9308C54.8143 11.9308 57.6654 13.1648 59.8459 15.568C61.9529 17.9746 63.0065 21.2243 63.0065 25.3172L62.9996 45.3371Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @adbird@vivaldi.net</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://social.vivaldi.net/" async src="https://social-cdn.vivaldi.net/embed.js"></script>
{{< /fig-gen >}}

有難うございます {{% emoji "ペコン" %}}

エディタ画面の左上に {{% keytop %}}`Ag`{{% /keytop %}} ボタンがあってフォントを選択できるのだが Noto Sans/Serif CJK JP フォントが見当たらなくて（リージョン未指定の Noto Sans/Serif はある）対応してないのかと思って諦めていた。
フォントのアップロードはしたくなかったし。

さっそくフォントを指定してみる。

{{< fig-img src="./sample-2.png" title="Project: sample" link="./sample-2.png" width="1142" >}}

生成した [PDF](./sample-2.pdf) をダウンロードしてプロパティを見てみると

{{< fig-img src="./property-2.png" title="property" link="./property-2.png" >}}

おー。
ちゃんと Noto Sans/Serif CJK JP フォントが埋め込まれている。
フォントの指定の仕方については[次回]({{< relref "./2-setup-typst-in-local-machine.md" >}} "ローカルで Typst 環境を整える")の記事を参照のこと。

## ブックマーク

- [話題の組版エンジン Typst を触ってみた](https://zenn.dev/monaqa/articles/2023-04-19-typst-introduction)
- [Typstの環境構築 Windows編 #Windows - Qiita](https://qiita.com/denkiuo604/items/21e8758ab160bf895e34)
- [Typst 備忘録 - adbird（広告鳥） 備忘録](https://adbird.hatenablog.com/entry/2024/03/21/015335)

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[typst/typst]: https://github.com/typst/typst "typst/typst: A new markup-based typesetting system that is powerful and easy to learn."
[winget]: https://github.com/microsoft/winget-cli "microsoft/winget-cli: Windows Package Manager CLI (aka winget)"
[Scoop]: https://scoop.sh/ "Scoop"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
