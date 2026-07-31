+++
title = "Typst 用のテンプレートを使う"
date =  "2025-02-24T15:23:33+09:00"
description = "Typst Universe / LaTeX jsarticle/jsbook ベースのテンプレート / 履歴書テンプレート / その他"
image = "/images/attention/tools.png"
tags  = [ "typst", "pdf" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[前回]に引き続きページ設定等について調べようと思ったのだが，いいものを見つけてしまった。

- [okumuralab/typst-js - Typst template based on LaTeX jsarticle/jsbook](https://github.com/okumuralab/typst-js)

[Typst] について調べていくうちに「思ったより面倒が多いし，独りでフルスクラッチで組むのは不毛じゃないの？」と思い始めていて，使い勝手のよさそうなテンプレートかパッケージを探していたところであった。

そこで今回は予定を変更して [`typst-js`][okumuralab/typst-js] を導入して簡単な文書を作ってみようと思う。
その上で次回以降から更に [Typst] の機能について調べていくことにする。

## [Typst Universe]

そのまえに [Typst Universe] について。

[Typst] では [Typst Universe] でパッケージやテンプレートを収集している。
たとえば「工学系の日本語の学会論文テンプレート」である [`jaconf-mscs`] なら

```text
$ typst init @preview/jaconf-mscs my-project
Successfully created new project from @preview/jaconf-mscs:0.1.0 🎉
To start writing, run:
> cd my-project
> typst watch main.typ
```

という感じに雛形のファイルを生成してくれる。
生成された `main.typ` の最初の部分をちょっと見てみると

```typst
// MIT No Attribution
// Copyright 2024, 2025 Shunsuke Kimura

#import "@preview/jaconf-mscs:0.1.0": jaconf, definition, lemma, theorem, corollary, proof, appendix
```

という記述があり，ここでパッケージをリモートからインポートしている。
続く `jaconf, definition, ...` はパッケージ内で定義されている関数である。

これでテンプレートの導入は大体分かったかな。

## LaTeX jsarticle/jsbook ベースのテンプレート

[okumuralab/typst-js] は『[LaTeX美文書作成入門](https://www.amazon.co.jp/dp/4297138891?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "［改訂第9版］LaTeX美文書作成入門 | 奥村 晴彦, 黒木 裕介 |本 | 通販 | Amazon")』でおなじみ[奥村晴彦](https://okumuralab.org/~okumura/ "奥村 晴彦 | Haruhiko Okumura")さんによる [Typst] 用のテンプレート。
奥村晴彦さんは $\mathrm{\LaTeX}$ 用のドキュメントクラスである `jsarticle`/`jsbook` の作者でもあり，日本語組版のノウハウが反映されることが期待できる（実際には [`typst-js`][okumuralab/typst-js] はまだ試行錯誤中とのこと）。
ありがたや {{% emoji "ペコン" %}}

[okumuralab/typst-js] リポジトリにある `template/example.typ` を見ると

```typst
#import "@preview/js:0.1.0": *
```

となっていたので [Typst Universe] に登録されているのかと思ったが `typst init` しようとしたら怒られてしまった。

{{< div-box type="markdown" >}}
**【2025-03-03 追記】**

[okumuralab/typst-js] は 2025-02-24 時点で [Typst Universe] に[登録](https://typst.app/universe/package/js "js – Typst Universe")されていたらしい。
タッチの差で記事に間に合わなかったか（笑）

以下のコマンドで雛形を生成できる。

```text
$ typst init @preview/js js-sample
downloading @preview/js:0.1.2
  5.3 KiB /   5.3 KiB (100 %)   5.3 KiB/s in 193.39 µs ETA: 0 s

Successfully created new project from @preview/js:0.1.2 🎉
To start writing, run:
> cd js-sample
> typst watch example.typ
```

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Universe]: https://typst.app/universe/ "Typst Universe"
[okumuralab/typst-js]: https://github.com/okumuralab/typst-js "okumuralab/typst-js - Typst template based on LaTeX jsarticle/jsbook"
{{< /div-box >}}

というわけで `template/example.typ` の記述を見ながら手動でファイルを作ってみる。
こんな感じ。

```typst
#import "js.typ": *

#show: js.with(
  lang: "ja",
  seriffont: "New Computer Modern",
  seriffont-cjk: "BIZ UDMincho",
  sansfont: "Liberation Sans",
  sansfont-cjk: "BIZ UDGothic",
  paper: "a4",
  fontsize: 10pt,
  baselineskip: auto,
  textwidth: auto,
  lines-per-page: auto,
  book: false,
  cols: 1,
  non-cjk: "latin-in-cjk",
  cjkheight: 0.88,
)

#maketitle(
  title: "Go のエラーハンドリング",
  authors: "Spiegel",
  abstract: [
    Go のエラーハンドリングについて今まで書いた駄文を「全部入り」で本の形にしてみようと思う。
  ],
  keywords: ("golang", "error", "programming"),
  date: "2025年2月24日"
)

#outline() #v(1em)

= はじめに

プログラミングにおいて，正常系は基本的に「一本道」だが，異常系は（予期しないものも含めて）無数にある。
エラーハンドリングは巨大迷路パズルを袋小路から順に塗りつぶして「正解」をあぶり出していく作業に似ていると思う。
下手くそな迷路攻略はただの「作業」だが，*よく考えられた迷路は袋小路の配置も美しい*。

こんなふうに考えるなら，プログラム設計の肝は#ruby[エラー][袋小路]をどう記述するかにかかっている，と言えるだろう。

私が Go のエラーハンドリングについて最初に記事にしたのは2015年のことだが，あれから Go も少しずつ変わっているし，私も当時よりは多少なりと理解が進んだと思うので，今まで書き散らかした駄文を「全部入り」で本の形にしてみようと思う。
```

文章の元ネタは拙文「[Go のエラーハンドリング](https://zenn.dev/spiegel/books/error-handling-in-golang)」より（いつか PDF 化したいと思っている）。
インポートは

```typst
#import "js.typ": *
```

としてローカルのカレントディレクトリにある `js.typ` を示すようにした。
使用する関数を列挙しなくてもワイルドカード（`*`）が使えるんだな。

なお `#`[`v`]`(1em)` は縦方向に1文字分の空き行を作る関数。
ちなみに水平方向の空き指定には [`h`] 関数を使う。

フォントの選択については前に書いた「[ローカルで Typst 環境を整える]({{< relref "./2-setup-typst-in-local-machine.md" >}})」を参照のこと。
Book 形式ではないので `book` の値は `false` とした。
ページ内カラム数は1。
`non-cjk` は [Typst] で定義されている `"latin-in-cjk"` とした。
[正規表現][`regex`]での指定もできるらしい。
あとはデフォルトのままかな。

蛇足だが著者は複数記述でき，かつ所属とアドレスも記述できる。
こんな感じ。

```typst
authors: (("何山 何某", "某大", "username@example.org"), ("何野 何某", "某大"))
```

さて，これをコンパイルしてみる。
このときのコマンドは

```text
$ typst compile --pdf-standard a-2b sample.typ
```

として PDF/A として構成するようにする[^p1]。
結果は以下の通り。

[^p1]: PDF/A については拙文「[LuaLaTeX で PDF/A を構成する]({{< ref "/remark/2020/06/pdfa-with-luatex.md" >}})」を参考にどうぞ。

{{< fig-img src="./pdf-sample.png" title="PDF 出力結果" link="./sample.pdf" width="693" >}}

出力した PDF ファイルのプロパティはこんな感じ。

{{< fig-img src="./pdf-sample-property.png" title="PDF 出力結果（プロパティ）" link="./pdf-sample-property.png" width="800" >}}

フォントの埋め込みは問題なし。
日本語のタイトルの文字化けもなし。
形式も PDF/A-2b で構成されてるっぽい。
念のため Windows 環境の Adobe Reader でも開いてみたけど，大丈夫かな。

{{< fig-img src="./pdf-sample-win.png" title="PDF 出力結果" link="./sample-win.pdf" width="1008" >}}

よし。
これでOKにしよう。

## 履歴書テンプレート

もうひとつテンプレートを紹介しておこう。

- [Nikudanngo/typst-ja-resume-template - Typst履歴書テンプレート](https://github.com/Nikudanngo/typst-ja-resume-template)
- [Typstで履歴書を書く #Typst - Qiita](https://qiita.com/Nikudanngo/items/ed9a452b5f63101fb26b)

データとコードの分離が甘い気もするが，組版の{{< pdf-file title="仕上がり" link="https://github.com/Nikudanngo/typst-ja-resume-template/blob/main/main.pdf" >}}を見る限り，かなりよく出来ている。
近年は $\mathrm{\LaTeX}$ では履歴書しか書いてないので [Typst] で書けるのならいよいよ $\mathrm{\TeX}$ 環境は要らなくなるかなぁ。
フォントファイルだけ何処かに退避させて他は削除してもいいかもしれない。

## その他

[日本語組版情報](https://typst-jp.github.io/docs/japanese/ "日本語組版情報 – Typstドキュメント日本語版")にある「[日本語テンプレート](https://typst-jp.github.io/docs/japanese/templates/ "日本語テンプレート – Typstドキュメント日本語版")」を見ると学会による公式テンプレートが既に幾つか公開されているようだ。
他には[ソフトウェア設計書のテンプレート](https://github.com/ctenopoma/SoftwareDesignTypst "ctenopoma/SoftwareDesignTypst")もあるみたいでちょっと気になるところである。

探せば他にも色々あるだろうが，ぼちぼち探していこう。

## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[document]: https://typst.app/docs/reference/model/document/ "Document Function – Typst Documentation"
[Typst Universe]: https://typst.app/universe/ "Typst Universe"
[`v`]: https://typst.app/docs/reference/layout/v/ "Spacing (V) Function – Typst Documentation"
[`h`]: https://typst.app/docs/reference/layout/h/ "Spacing (H) Function – Typst Documentation"
[`regex`]: https://typst.app/docs/reference/foundations/regex/ "Regex Type – Typst Documentation"

[前回]: {{< relref "./3-typst-document-elements-1.md" >}} "Typst のドキュメント要素（1）"
[`jaconf-mscs`]: https://typst.app/universe/package/jaconf-mscs "jaconf-mscs – Typst Universe"
[okumuralab/typst-js]: https://github.com/okumuralab/typst-js "okumuralab/typst-js - Typst template based on LaTeX jsarticle/jsbook"


## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
