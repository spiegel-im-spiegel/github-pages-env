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

## ちょっと触ってみる

先程のダッシュボードを開いて {{% keytop %}}`Empty document`{{% /keytop %}} を押してみる。

{{< fig-img src="./empty-document.png" title="Create project" link="./empty-document.png" width="1489" >}}

適当に “`sample`” とかで始めてみるか。
以下の文言を入力する。

```typst
= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

構文についてはとりあえず後回しにして，組版結果はこんな感じ。

{{< fig-img src="./sample.png" title="Project: tutorial-1" link="./sample.png" width="1263" >}}

うわぁ。
これは酷い（笑）

まずエディタのフォントが `"Cascadia Mono", monospace` になってるな。

{{< fig-img src="./settings.png" title="Settings" link="./settings.png" width="1264" >}}

画面左上に {{% keytop %}}↓{{% /keytop %}} ボタンがあるので，これを押下して PDF ファイルを[ダウンロード](./sample.pdf)し，プロパティを見てみる。

{{< fig-img src="./property.png" title="property" link="./property.png" width="675" >}}

なんか中華フォントっぽい。
句読点も明らかにおかしいし。
やっぱ，まともに日本語のドキュメント環境が欲しいならローカルで作らないとダメかな。
ちょっと考えてみよう。

## ブックマーク

- [話題の組版エンジン Typst を触ってみた](https://zenn.dev/monaqa/articles/2023-04-19-typst-introduction)
- [Typstの環境構築 Windows編 #Windows - Qiita](https://qiita.com/denkiuo604/items/21e8758ab160bf895e34)

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[typst/typst]: https://github.com/typst/typst "typst/typst: A new markup-based typesetting system that is powerful and easy to learn."
[winget]: https://github.com/microsoft/winget-cli "microsoft/winget-cli: Windows Package Manager CLI (aka winget)"
[Scoop]: https://scoop.sh/ "Scoop"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
