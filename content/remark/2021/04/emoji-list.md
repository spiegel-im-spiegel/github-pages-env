+++
title = "「情報交換用に推奨される絵文字のリスト」を作ってみた"
date =  "2021-04-02T22:26:53+09:00"
description = "とりあえずチョー面倒くさいのは分かった。"
image = "/images/attention/kitten.jpg"
tags = [ "character", "unicode", "engineering", "emoji" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]，絵文字についてちょろんと調べた余波で，絵文字の一覧を作ってみた。

- [情報交換用に推奨される絵文字のリスト]({{< ref "/emoji-list.md" >}})
- [絵文字と Markdown Shortcodes]({{< ref "/emoji-shortcodes.md" >}})

本当はひとつの表にしたかったのだが，ページのレイアウトが崩れまくるので2つに分けた。
まぁ，それでも携帯端末で見たら崩れてると思うけど。

今回の一覧を作るにあたって Unicode の公式ページから以下の情報を拾って使っている。

- [`https://www.unicode.org/Public/UCD/latest/ucd/emoji/`](https://www.unicode.org/Public/UCD/latest/ucd/emoji/)
  - [`ReadMe.txt`](https://www.unicode.org/Public/UCD/latest/ucd/emoji/ReadMe.txt)
  - [`emoji-data.txt`](https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-data.txt)
  - [`emoji-variation-sequences.txt`](https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-variation-sequences.txt)
- [`https://unicode.org/Public/emoji/`](https://unicode.org/Public/emoji/)
  - [13.1/](https://unicode.org/Public/emoji/13.1/)
    - [`ReadMe.txt`](https://unicode.org/Public/emoji/13.1/ReadMe.txt)
    - [`emoji-sequences.txt`](https://unicode.org/Public/emoji/13.1/emoji-sequences.txt)
    - [`emoji-zwj-sequences.txt`](https://unicode.org/Public/emoji/13.1/emoji-zwj-sequences.txt)

なお markdown shortcodes の情報の取得には以下のパッケージを使った。

- [kyokomi/emoji: emoji terminal output for golang](https://github.com/kyokomi/emoji)

このパッケージは Hugo でも使われているので「[絵文字と Markdown Shortcodes]({{< ref "/emoji-shortcodes.md" >}})」にある markdown shortcodes は全て [Hugo で使える]({{< ref "/remark/2020/10/emoji-variation-and-markdown.md" >}} "絵文字と異体字と Markdown")。

なにせ量が量なので手作業でやるわけにもいかず，データ取得用のパッケージを作った。

- [spiegel-im-spiegel/emojis: List of Emoji-Sequences](https://github.com/spiegel-im-spiegel/emojis)

中身は JSON データ作成用のモジュールと作成した JSON ファイルを読み込むモジュールに別れている。
JSON データ作成用のモジュールで作成した JASON ファイルは，パッケージの [`json/`](https://github.com/spiegel-im-spiegel/emojis/tree/main/json) ディレクトリに置いている。
JSON データを読むための構造体定義も同じところに置いているので，ご利用はご自由にどうぞ。

## 絵文字の構造

今回利用したデータの中身については以下のページに公式の解説がある。

- [UTS #51: Unicode Emoji](http://www.unicode.org/reports/tr51/)

以降で簡単に紹介してみる。

### Emoji Characters

まず Unicode 符号点単位の「絵文字文字（emoji characters）」には文字種ごとに以下のプロパティが付与されている。

| プロパティ | 種別 |
| --- | --- |
| `Emoji` | emoji character |
| `Extended_Pictographic` | extended pictographic character |
| `Emoji_Presentation` | default emoji presentation character |
| `Emoji_Component` | emoji component |
| `Emoji_Modifier` | emoji modifier |
| `Emoji_Modifier_Base` | emoji modifier base |

これらのプロパティは独立に付与されていて，複数のプロパティが付与されている符号点コードもある。

[spiegel-im-spiegel/emojis](https://github.com/spiegel-im-spiegel/emojis "spiegel-im-spiegel/emojis: List of Emoji-Sequences") パッケージでは符号点コードごとにプロパティをチェックできるようにした。

```go
type EmojiData struct {
    Code                 rune
    Name                 string
    Emoji                bool   `json:",omitempty"`
    EmojiPresentation    bool   `json:"Emoji_Presentation,omitempty"`
    EmojiModifier        bool   `json:"Emoji_Modifier,omitempty"`
    EmojiModifierBase    bool   `json:"Emoji_Modifier_Base,omitempty"`
    EmojiComponent       bool   `json:"Emoji_Component,omitempty"`
    ExtendedPictographic bool   `json:"Extended_Pictographic,omitempty"`
    RegionalIndicator    bool   `json:"Regional_Indicator,omitempty"`
    VariationTextStyle   string `json:",omitempty"`
    VariationEmojiStyle  string `json:",omitempty"`
}
```

### Emoji Presentation Sequences

`Emoji` プロパティが付与されている符号点コードの直後に絵文字表示セレクタ `U+FE0F VARIATION SELECTOR-16` を付けることで絵文字であることを明示できる。
といっても何でも組み合わせればいいというわけではないようで，先ほどの [`emoji-variation-sequences.txt`](https://www.unicode.org/Public/UCD/latest/ucd/emoji/emoji-variation-sequences.txt) で定義されているシーケンスのみ有効らしい。

`U+FE0F VARIATION SELECTOR-16` には `Emoji_Component` プロパティが付与されている。

### Emoji Modifiers

`Emoji_Modifier_Base` プロパティが付与されている符号点コードに `Emoji_Modifier` プロパティが付与されている符号点コードを付けることで`Emoji_Modifier_Base` プロパティが付与されている符号点コードの絵文字を装飾できる。

といって現在は肌色のトーンを変更する符号点コードしかないのだが。
肌色の異体字については[前回]の記事を参照のこと。

### Emoji Flag Sequence

`Emoji_Component` プロパティを持つ符号点コードの中に Regional Indicator に分類されるコードがあるのだが，このコードのうち2文字を組み合わせて国別コードを構成すると国旗の絵文字になるというトンデモ仕様がある。
これについても[前回]の記事を参照のこと。

### Emoji Tag Sequence (ETS)

まず，文字種を以下のように定義する

```text
tag_base := emoji_character | emoji_modifier_sequence | emoji_presentation_sequence
tag_spec := [\x{E0020}-\x{E007E}]+
tag_end  := \x{E007F}
```

これを使って

```text
emoji_tag_sequence := tag_base tag_spec tag_end
```

と定義されるのが emoji tag sequence の構成である。
`tag_spec` および `tag_end` に含まれる符号点コードには `Emoji_Component` プロパティが付与されている。

Unicode v13 現在は subdivision-flag のみ実装されているようだ。
これについても[前回]の記事を参照のこと。

### Emoji Keycap Sequence

```text
emoji_keycap_sequence := [0-9#*] \x{FE0F 20E3}
```

たったこれだけのためのシーケンス。
なんだかなぁ。

ちなみに `[0-9#*]` はただの半角文字だが `Emoji_Component` プロパティが付与されている。
また `U+20E3` にも `Emoji_Component` プロパティが付与されている。

### Emoji ZWJ Sequence

```text
emoji_zwj_element := emoji_character | emoji_presentation_sequence | emoji_modifier_sequence
```

としたときに

```text
emoji_zwj_sequence := emoji_zwj_element ( ZWJ emoji_zwj_element )+
```

で構成される絵文字。
ちなみに `U+200D ZWJ` は `Emoji_Component` プロパティが付与された結合子で，これを使っていくらでも文字を繋げられるのが恐ろしい点である。
といっても情報交換用として推奨される組み合わせが [`emoji-zwj-sequences.txt`](https://unicode.org/Public/emoji/13.1/emoji-zwj-sequences.txt) で定義されているので，この中から選択することになるだろう。

## 絵文字の分類

結局，絵文字の分類は以下のようになるらしい。

- emoji sequence
  - emoji core sequence
    - emoji character
    - emoji presentation sequence
    - emoji keycap sequence
    - emoji modifier sequence
    - emoji flag sequence
  - emoji zwj sequence
  - emoji tag sequence

emoji character が emoji core sequence に含まれている点に注目。

で， “[UTS #51: Unicode Emoji](http://www.unicode.org/reports/tr51/)” に絵文字を判別するための正規表現ってのが載っていたのだが

{{< fig-quote type="markdown" class="nobox" title="UTS #51: Unicode Emoji" link="http://www.unicode.org/reports/tr51/" lang="en" >}}
```text
\p{Regional_Indicator} \p{Regional_Indicator} 
| \p{Emoji} 
  ( \p{Emoji_Modifier} 
  | \x{FE0F} \x{20E3}? 
  | [\x{E0020}-\x{E007E}]+ \x{E007F} )?
  (\x{200D} \p{Emoji}
    ( \p{Emoji_Modifier} 
    | \x{FE0F} \x{20E3}? 
    | [\x{E0020}-\x{E007E}]+ \x{E007F} )?
  )*
```
{{< /fig-quote >}}

いやいやいや。
符号点コードごとに絵文字プロパティをチェックせんとアカンのかい！ これは面倒くさい。

[`emoji-sequences.txt`](https://unicode.org/Public/emoji/13.1/emoji-sequences.txt) および [`emoji-zwj-sequences.txt`](https://unicode.org/Public/emoji/13.1/emoji-zwj-sequences.txt) ファイルでは収録している絵文字を以下のように分類している。


| 分類名 | 内容 |
| --- | --- |
| `Basic_Emoji` | `Emoji_Presentation` プロパティを含む emoji character または `Emoji_Presentation` プロパティを含まない emoji character の emoji presentation sequence で構成された絵文字 |
| `Emoji_Keycap_Sequence`   | emoji keycap sequence で構成された絵文字 |
| `RGI_Emoji_Flag_Sequence` | emoji flag sequence で構成される推奨絵文字 |
| `RGI_Emoji_Tag_Sequence`  | emoji tag sequence で構成される推奨絵文字 |
| `RGI_Emoji_Modifier_Sequence` | emoji modifier sequence で構成される推奨絵文字 |
| `RGI_Emoji_ZWJ_Sequence` | emoji zwj sequence で構成される推奨絵文字 |

ちなみに RGI は “Recommended for General Interchange” の略だそうだ。
超意訳するなら「情報交換用の推奨絵文字」ってところかねぇ。
なお「[情報交換用に推奨される絵文字のリスト]({{< ref "/emoji-list.md" >}})」ではこの分類を表示している。
また [spiegel-im-spiegel/emojis](https://github.com/spiegel-im-spiegel/emojis "spiegel-im-spiegel/emojis: List of Emoji-Sequences") パッケージでは絵文字シーケンスごとにこの分類をセットしている。

```go
type SequencesType int

const (
    TypeUnknown SequencesType = iota
    TypeBasicEmoji
    TypeEmojiKeycapSequence
    TypeRGIEmojiFlagSequence
    TypeRGIEmojiTagSequence
    TypeRGIEmojiModifierSequence
    TypeRGIEmojiZWJSequence
)
```

```go
type EmojiSequence struct {
    Sequence     string
    Name         string
    SequenceType types.SequencesType
    Shortcodes   []string `json:",omitempty"`
}
```


これで絵文字の調査は一段落かな。
とりあえずチョー面倒くさいのは分かった。

[前回]: {{< ref "/remark/2021/03/terrible-emoji.md" >}} "おそるべき絵文字"
<!-- eof -->
