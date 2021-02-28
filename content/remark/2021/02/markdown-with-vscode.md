+++
title = "Markdown と VS Code"
date =  "2021-02-28T12:45:02+09:00"
description = "Preview 機能は個人的に必要ないのでレビューしないが，仕事で使うようなことがあれば，そのうち記事にすることもあるだろう。"
image = "/images/attention/kitten.jpg"
tags = ["vscode", "editor", "markdown"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})
2. [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}})
3. [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}}) ← イマココ

今回は [VS Code] で markdown テキストを入出力する話。

Markdown 関連の拡張機能は色々あるようだが，入力支援に関しては以下のもので必要十分ぽい。

- [Markdown All in One - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one)

インストールはこちら。

```text
$ code --install-extension yzhang.markdown-all-in-one
```

お気に入りはテーブル整形の機能で， `[Ctrl+Shift+I]` キー押下で綺麗に整形してくれる。

## Markdown Preview 機能は必要か

私個人で言うなら No で，仕事なら場合によっては Yes かな。

そもそも markdown テキストってのは，見出しや段落や箇条書き等の文書構造がそのままでも human-readable である点が利点と言える。
さらに言えば，ここのブログは [Hugo] の shortcodes 等で入力自体をカスタマイズしまくってるので，ただの markdown preview なんか使いものにならないのだ（[Hugo] はサーバ・モードで起動できるので，リアルタイムでブラウザ表示を確認しながら記事を書いている）。

一方で， Office ツールなどレガシーな環境を捨てて markdown 等の構造化テキストをベースにしたドキュメンテーションをしようとするなら，それなりにリッチな markdown preview 機能と PDF 等へ「最終出力」するためのツールチェーンが必要となる。

そのための手段（または道具立て）として [VS Code] をベースに環境を整えるというのは合理的と言えるかもしれない。

というわけで，以降では PDF 等への「最終出力」を念頭に置いたドキュメントツールとして幾つかの拡張機能を紹介してみる。
上述したように，私個人は全く必要ないのでレビューしないが，仕事で使うようなことがあれば，そのうち記事にすることもあるだろう。

### [Markdown Preview Enhanced](https://marketplace.visualstudio.com/items?itemName=shd101wyy.markdown-preview-enhanced "Markdown Preview Enhanced - Visual Studio Marketplace")

```text
$ code --install-extension shd101wyy.markdown-preview-enhanced
```

TeX の数学記法あるいは [PlantUML] や [mermaid] 等の記法も認識して preview 表示できるらしい。
[DOT] 言語も使えるのか。
CSS をカスタマイズ可能。

HTML や PDF へ出力できるようだ（PDF 出力は Chrome 経由）。

### [Markdown PDF](https://marketplace.visualstudio.com/items?itemName=yzane.markdown-pdf "Markdown PDF - Visual Studio Marketplace")

```text
$ code --install-extension yzane.markdown-pdf
```

コマンド一発で PDF 変換してくれる。
簡易的な出力しか出来ないのかと思ったら，意外にもかなりカスタマイズできるらしい。

- [VSCodeとMarkdownで技術同人誌書いたので拡張機能とかまとめ - Qiita](https://qiita.com/reona396/items/40b234108f7664267db8)
- [【Visual Studio Code】Markdown PDF のスタイル(CSS)を変える方法 - Nekonote](https://h-s-hige.hateblo.jp/entry/20190405/1554467885)

### [Marp for VS Code](https://marketplace.visualstudio.com/items?itemName=marp-team.marp-vscode "Marp for VS Code - Visual Studio Marketplace")

```text
$ code --install-extension marp-team.marp-vscode
```

[Marp] を使って markdown テキストからスライドを生成する。
PDF へエクスポートできるらしい。

- [【VS Code + Marp】Markdownから爆速・自由自在なデザインで、プレゼンスライドを作る - Qiita](https://qiita.com/tomo_makes/items/aafae4021986553ae1d8)

### [Draw.io Integration](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio "Draw.io Integration - Visual Studio Marketplace")

```text
$ code --install-extension hediet.vscode-drawio
```

[Draw.io (diagrams.net)](https://app.diagrams.net/) を利用した作図ツール。
データはテキストで保持して PNG や SVG へエクスポート可能って感じなのかな。

- [VSCodeでDraw.ioが使えるようになったらしい！ - Qiita](https://qiita.com/riku-shiru/items/5ab7c5aecdfea323ec4e)

### [PlantUML](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml "PlantUML - Visual Studio Marketplace")

```text
$ code --install-extension jebbs.plantuml
```

[PlantUML] 作図・出力支援。
あらかじめ [PlantUML] 作図環境を用意する必要がある（ただし `plantuml.jar` ファイルは拡張機能内にあらかじめ格納されている？）。

- [真面目に PlantUML (1) : PlantUML のインストール]({{< ref "/remark/2018/12/plantuml-1.md" >}})

## ブックマーク

- [VS Codeエディタ入門](https://zenn.dev/karaage0703/books/80b6999d429abc8051bb)
- [(2020年12月8日追記)VSCodeでDraw.io Integration使用時にエクスポートできないことがある問題への対処 - Qiita](https://qiita.com/tfukumori/items/0f2b52088cd39f5c124e)

[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."
[mermaid]: https://mermaidjs.github.io/
[DOT]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[Marp]: https://marp.app/ "Marp: Markdown Presentation Ecosystem"

## 参考図書

{{% review-paapi "B08CZ2C3NZ" %}} <!-- Software Design (2020年8月号) -->
