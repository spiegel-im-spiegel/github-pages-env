+++
title = "結局 Google Fonts に巻き戻した。そしてモリサワ BIZ UD フォント採用へ"
date =  "2022-09-18T21:31:41+09:00"
description = "とほほ orz"
image = "/images/attention/kitten.jpg"
tags = [ "web", "font" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

実はさっき2022年3月に公開された

- [「BIZ UD」フォントが「Google Fonts」へ ～モリサワのユニバーサルフォント - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1399117.html)

という記事を見つけたのだが，この中で

{{< fig-quote type="markdown" title="「BIZ UD」フォントが「Google Fonts」へ" link="https://forest.watch.impress.co.jp/docs/news/1399117.html" >}}
TrueType版に関しては教育現場で利用されることを考慮し、「IPAフォント」v003（最新）に収録されている全文字をカバーしているという
{{< /fig-quote >}}

という記述[^ipa1] を見て「え？ じゃあ [神（`U+FA19`）とか ㈱（`U+3231`）とかが表示できない]({{< relref "./morisawa-biz-ud-fonts.md" >}} "モリサワ BIZ UD 明朝/ゴシック Web フォント")のおかしくね？」と思い，もう少し真面目に調べてみた。

[^ipa1]: IPAex フォントのほうは2019年に [Ver.004.01](https://moji.or.jp/ipafont/releasenote00401/ "IPAexリリースノート Ver.004.01 | 一般社団法人 文字情報技術促進協議会") が出ている。例の令和の合字「&#x32FF;（`U+32FF`）」への対応である。ただ BIZ UD フォントはこの字もちゃんと含んでいるように見える。

したら，収録文字を制限してるのはどうやら [Bunny Fonts] 側のようなのだ。
[Google Fonts] で BIZ UD フォントを検索したら「神」も「㈱」も「①」もちゃんと表示してくれた。

たぶん [Bunny Fonts] は文字数を抑えるためにわざと JIS X 0208 の範囲しか収録してないんだろうね（邪推）。
でも異体字セレクタで指示する「&#x845B;&#xE0100;（U+845B U+E0100）」なんかは [Bunny Fonts] でも表示できるっぽいんだよなぁ。
基準が分からん。

とりあえず BIZ UD フォント側を疑ってまじすんません {{< emoji "ペコン" >}}

正直このブログでは機種依存文字（古語）や異体字等は割と使うので削られると困るし [Bunny Fonts] を使う限り NOTO JP フォントでも同じようになるみたいなので，泣く泣く [Google Fonts] に戻すことにした。
そんで，どうせ [Google Fonts] に戻すならモリサワ BIZ UD フォントで全然問題ないやろ，ということで最終的に以下の指定で Web フォントを読み込んでいる。

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=BIZ+UDGothic:wght@400;700&family=BIZ+UDMincho&family=Inconsolata:wght@400;700&family=Noto+Color+Emoji&family=Noto+Sans:wght@400;700&family=Noto+Serif&display=swap" rel="stylesheet">
```

簡単に言うと日本語部分のみ BIZ UD ゴシック/明朝フォントを使い，それ以外は従来の NOTO フォント及び Inconsolata フォントを使う感じ。
これで英文の引用や

{{< fig-quote type="markdown" title="35,000 code repos not hacked—but clones flood GitHub to serve malware" link="https://www.bleepingcomputer.com/news/security/35-000-code-repos-not-hacked-but-clones-flood-github-to-serve-malware/" lang="en" >}}
As a best practice, remember to consume software from the official project repos and watch out for potential typosquats or repository forks/clones that may appear identical to the original project but hide malware.
{{< /fig-quote >}}

プログラム・コードは

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, 世界")
}

```

従来どおりに表示するようにしている。

ほんじゃあ，まぁ，これからこれで行きますかね。
モリサワ BIZ UD フォントで表示できない文字が出てきたら最終的に NOTO JP フォントにまで巻き戻すかも知れんけど（笑）

## ブックマーク

- [Google Fonts が日本語に対応してた]({{< ref "/remark/2019/12/japanese-fonts-by-google-cdn.md" >}})
- [Google Fonts から Bunny Fonts に乗り換える]({{< ref "/remark/2022/06/migrate-to-bunny-fonts-from-google-fonts.md" >}})

[Bunny Fonts]: https://fonts.bunny.net/ "Bunny Fonts | Explore Faster & GDPR friendly Fonts"
[Google Fonts]: https://www.google.com/fonts/ "Browse Fonts - Google Fonts"
