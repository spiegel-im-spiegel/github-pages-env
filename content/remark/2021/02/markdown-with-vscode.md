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

## [Markdown All in One]

[Markdown All in One]: https://marketplace.visualstudio.com/items?itemName=yzhang.markdown-all-in-one "Markdown All in One - Visual Studio Marketplace"

```text
$ code --install-extension yzhang.markdown-all-in-one
```

Markdown 関連の拡張機能は色々あるようだが，入力支援に関してはこれで必要十分ぽい。

お気に入りはテーブル整形の機能で， Linux/Ubuntu なら `[Ctrl+Shift+I]` キー（Format Document）押下で綺麗に整形してくれる。

ところが Windows 版では `[Shift+Alt+F]` キーが Format Document に割り当てられているようだ。
プラットフォームによって違うのかよ。

というわけで Windows 版の方にキー割当を合わせることにした。

```json
// Place your key bindings in this file to override the defaults
[
    {
        "key": "shift+alt+f",
        "command": "editor.action.formatDocument",
        "when": "editorHasDocumentFormattingProvider && editorTextFocus && !editorReadonly && !inCompositeEditor"
    },
    {
        "key": "ctrl+shift+i",
        "command": "-editor.action.formatDocument",
        "when": "editorHasDocumentFormattingProvider && editorTextFocus && !editorReadonly && !inCompositeEditor"
    },
    {
        "key": "shift+alt+f",
        "command": "editor.action.formatDocument.none",
        "when": "editorTextFocus && !editorHasDocumentFormattingProvider && !editorReadonly"
    },
    {
        "key": "ctrl+shift+i",
        "command": "-editor.action.formatDocument.none",
        "when": "editorTextFocus && !editorHasDocumentFormattingProvider && !editorReadonly"
    }
]
```

## [Prettier - Code formatter][prettier] との競合

[Prettier]: https://marketplace.visualstudio.com/items?itemName=esbenp.prettier-vscode "Prettier - Code formatter - Visual Studio Marketplace"

```text
$ code --install-extension esbenp.prettier-vscode
```

Markdown 専用というわけではないが JavaScript/TypeScript, CSS/SCSS/Less, HTML, JSON, GraphQL, YAML など幅広い言語に対応している整形ツールで，しかも plugin 拡張もできるらしい。
もちろん markdown テキストにも対応している。

で，これと [Markdown All in One] の整形機能（Format Document）が被るわけですよ。
そこで，どちらの機能を使うか言語ごとに設定できるようになっている。
私は [Markdown All in One] 優先でこんな感じ。

```json
{
    "[markdown]": {
        "editor.defaultFormatter": "yzhang.markdown-all-in-one"
    }
}
```

さらにファイル保存時に変更した箇所だけを整形する，なんてな設定も言語ごとにできるようだ。

```json {hl_lines=[4,5]}
{
    "[markdown]": {
        "editor.defaultFormatter": "yzhang.markdown-all-in-one",
        "editor.formatOnSave": true,
        "editor.formatOnSaveMode": "modifications"
    }
}
```

ちなみに [EditorConfig for VS Code] が有効な場合は `.editorconfig` の設定（インデントや改行コードなど）を考慮してくれるようだ。
これを無効にするには “Use Editor Config” の項目を OFF にする。

```json {hl_lines=[4,5]}
{
    "prettier.useEditorConfig": false
}
```

ただし， [EditorConfig][EditorConfig for VS Code] が有効な場合でも `.prettierrc` ファイルなどによる [Prettier] 独自の設定がある場合は，そちらのほうが優先されるようだ。
ややこしい...

## 自動補完を有効にする

[VS Code] の売りのひとつはスニペットを含む強力な自動補完機能だが，何故か markdown ファイルには自動補完が効かない。
と思ったら，既定で無効になっているらしい。
いや，有効にしとけよ。

というわけで `settings.json` に以下の設定を手動で書き込む。

```json
{
    "[markdown]": {
        "editor.quickSuggestions": true,
        "editor.snippetSuggestions": "top"
    },
}
```

これでスニペットを優先して自動補完候補に挙げてくれる。

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

- [VSCode と Markdown で技術同人誌書いたので拡張機能とかまとめ - Qiita](https://qiita.com/reona396/items/40b234108f7664267db8)
- [【Visual Studio Code】Markdown PDF のスタイル(CSS)を変える方法 - Nekonote](https://h-s-hige.hateblo.jp/entry/20190405/1554467885)

### [Marp for VS Code](https://marketplace.visualstudio.com/items?itemName=marp-team.marp-vscode "Marp for VS Code - Visual Studio Marketplace")

```text
$ code --install-extension marp-team.marp-vscode
```

[Marp] を使って markdown テキストからスライドを生成する。
PDF へエクスポートできるらしい。

- [【VS Code + Marp】Markdown から爆速・自由自在なデザインで、プレゼンスライドを作る - Qiita](https://qiita.com/tomo_makes/items/aafae4021986553ae1d8)

### [Draw.io Integration](https://marketplace.visualstudio.com/items?itemName=hediet.vscode-drawio "Draw.io Integration - Visual Studio Marketplace")

```text
$ code --install-extension hediet.vscode-drawio
```

[Draw.io (diagrams.net)](https://app.diagrams.net/) を利用した作図ツール。
データはテキストで保持して PNG や SVG へエクスポート可能って感じなのかな。

- [VSCode で Draw.io が使えるようになったらしい！ - Qiita](https://qiita.com/riku-shiru/items/5ab7c5aecdfea323ec4e)

### [PlantUML](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml "PlantUML - Visual Studio Marketplace")

```text
$ code --install-extension jebbs.plantuml
```

[PlantUML] 作図・出力支援。
あらかじめ [PlantUML] 作図環境を用意する必要がある（ただし `plantuml.jar` ファイルは拡張機能内にあらかじめ格納されている？）。

- [真面目に PlantUML (1) : PlantUML のインストール]({{< ref "/remark/2018/12/plantuml-1.md" >}})

## ブックマーク

- [VS Code エディタ入門](https://zenn.dev/karaage0703/books/80b6999d429abc8051bb)
- [(2020 年 12 月 8 日追記)VSCode で Draw.io Integration 使用時にエクスポートできないことがある問題への対処 - Qiita](https://qiita.com/tfukumori/items/0f2b52088cd39f5c124e)
- [Visual Studio Code で markdown のスニペットを登録する  |  コーヒー飲みながら仕事したい](https://coffee-nominagara.com/2019-01-25-094628)
- [インデントおよび行末は EditorConfig で始末する](https://zenn.dev/spiegel/articles/20200922-editorconfig)

[Vs Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."
[Mermaid]: https://mermaidjs.github.io/
[dot]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[Marp]: https://marp.app/ "Marp: Markdown Presentation Ecosystem"
[EditorConfig for VS Code]: https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig "EditorConfig for VS Code - Visual Studio Marketplace"

## 参考図書

{{% review-paapi "B08CZ2C3NZ" %}} <!-- Software Design (2020年8月号) -->
