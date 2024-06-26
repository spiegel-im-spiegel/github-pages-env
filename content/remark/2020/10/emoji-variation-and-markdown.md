+++
title = "絵文字と異体字と Markdown"
date =  "2020-10-08T16:22:55+09:00"
description = "Hugo の markdown で絵文字を簡単に表示する。"
image = "/images/attention/kitten.jpg"
tags = [ "character", "unicode", "markdown", "hugo", "font", "emoji" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

絵文字ってある意味で漢字より難解なので，なるべく使わないようにしてきたのだが，文章に混ぜたり emoticon 代わりに使ったりするだけでなく [Zenn] みたいにページのアテンションに使ってる例もあるわけで「もう少し積極的に使ってもいいかなぁ」という気になってきた。

とはいえ input method から延々と探したりその辺のページから絵文字をコピペするのも馬鹿らしいし，折角 markdown で書いてるんだから意味のあるコードで書きたい。

[Hugo] なら `config.toml` 等の設定ファイルで

```toml
enableEmoji = true
```

とすれば markdown テキスト内の `:heart:` 等を {{< emoji ":heart:" >}} 等に変換してくれる。
あるいはテンプレートや shortcodes で

```text
{{ emojify ":heart:" }}
```

などとしてもよい。

変換可能な絵文字の一覧は以下を参照のこと。

- [🎁 Emoji cheat sheet for GitHub, Basecamp, Slack & more](https://www.webfx.com/tools/emoji-cheat-sheet/)

ただしリンク先にある全部の文字が [Hugo] で使えるわけではないらしい。
私がよく使いそうな絵文字を挙げておくか。

| 字形  | Markdown コード | CLDR Short Name |
|:-----:| --------------- | ------------ |
| {{< span class="larger" >}}{{< emoji ":smile:" >}}{{< /span >}} | `:smile:` | grinning face with smiling eyes |
| {{< span class="larger" >}}{{< emoji ":sleeping:" >}}{{< /span >}} | `:sleeping:` | sleeping face |
| {{< span class="larger" >}}{{< emoji ":pensive:" >}}{{< /span >}} | `:pensive:` | pensive face  |
| {{< span class="larger" >}}{{< emoji ":bow:" >}}{{< /span >}} | `:bow:` | person bowing ??? |
| {{< span class="larger" >}}{{< emoji ":zzz:" >}}{{< /span >}} | `:zzz:` | zzz |
| {{< span class="larger" >}}{{< emoji ":anger:" >}}{{< /span >}} | `:anger:` | anger symbol |
| {{< span class="larger" >}}{{< emoji ":sweat_drops:" >}}{{< /span >}} | `:sweat_drops:` | sweat droplets[^sweat1]  |
| {{< span class="larger" >}}{{< emoji ":star:" >}}{{< /span >}} | `:star:` | star  |
| {{< span class="larger" >}}{{< emoji ":bulb:" >}}{{< /span >}} | `:bulb:` | light bulb  |
| {{< span class="larger" >}}{{< emoji ":musical_note:" >}}{{< /span >}} | `:musical_note:` | musical note  |
| {{< span class="larger" >}}{{< emoji ":key:" >}}{{< /span >}} | `:key:` | key |
| {{< span class="larger" >}}{{< emoji ":lock:" >}}{{< /span >}} | `:lock:` | locked |
| {{< span class="larger" >}}{{< emoji ":unlock:" >}}{{< /span >}} | `:unlock:` | unlocked  |
| {{< span class="larger" >}}{{< emoji ":closed_lock_with_key:" >}}{{< /span >}} | `:closed_lock_with_key:` | locked with key |

[^sweat1]: 海外では {{< emoji ":sweat_drops:" >}} は男性の射精の意味で受け取られるそうで，この絵文字を含む8文字ほどをブロック対象にすると [Instagram でセクハラコメントの90%を防げる](https://boingboing.net/2021/02/01/blocking-8-emoji-stopped-90-of-my-unwanted-instagram-comments.html "Blocking 8 emoji stopped 90% of my unwanted Instagram comments | Boing Boing")そうな（笑） 世知辛い。

さて，皆さんの環境ではどう見えているでしょうか。
人によって見え方が異なると思います（絵文字のフォント指定については[後述]({{< relref "#fonts" >}})）。

ところで，上に挙げた `:bow:` ってどう見えてます？ 実は {{< span >}}&#x1f647;&#x2642;{{< /span >}} ってな感じの2文字に見えてません？ うちでは `:bow:` から変換した絵文字をターミナル・エミュレータとかにコピペすると，こうなるんですよ。

というわけでコードを見てみると

```text
$ echo 🙇‍♂️ | gnkf dump --unicode
0x0001f647, 0x200d, 0x2642, 0xfe0f, 0x000a
```

おいおい。
本来のコードの後ろに何か付いとるやないかいっ {{< emoji ":anger:" >}} （← 早速）

実はこれ，絵文字の異体字なんだよ。
内訳はこんな感じ。

| Unicode Point |                字形                | Unicode 名称          |
| ------------- |:----------------------------------:| --------------------- |
| `U+1F647`     | {{< span >}}&#x1f647;{{< /span >}} | PERSON BOWING DEEPLY  |
| `U+200D`      |                                    | ZERO WIDTH JOINER     |
| `U+2642`      | {{< span >}}&#x2642;{{< /span >}}  | MALE SIGN             |
| `U+FE0F`      |                                    | VARIATION SELECTOR-16 |

ZERO WIDTH JOINER (ゼロ幅接合子; ZWJ) はアラビア文字なんかで複数の文字を結合してひとつの文字にするための制御文字。
さらに VARIATION SELECTOR-16 は “Emoji Variation Selector” とも呼ばれ，絵文字の異体字であることを示す異体字セレクタである[^vs1]。

[^vs1]: 異体字セレクタのコードポイントは通常のものと区別するために `E+FE0F`と表記する場合もあるらしい。

つまり `:bow:` で表示されるのは「土下座する男性」を意味する絵文字（の異体字）なわけだ。
処理系によって異体字を上手く表示できない場合は {{< span >}}&#x1f647;&#x2642;{{< /span >}} のように2つの文字が並んでいるように見えたりするようだ。

...なんで男性なんだろうね。
[GitHub の既定ブランチ名を弄る]({{< ref "/remark/2020/08/renaming-default-branch-name-in-github-repositries.md" >}} "GitHub リポジトリの既定ブランチ名が main になるらしい")とかするくらいなら，こういうのを真っ先に改善すべきなんじゃないの？

というわけで [Hugo] 環境全体の設定は変えずに，絵文字表示用の shortcode を作って対応することにした。
こんな感じ。

```text
{{- range $i, $s := .Params -}}
{{- if gt $i 0 -}}&nbsp;{{- end -}}<abbr title="{{ $s }}">
{{- if eq $s "ゴメン" -}}&#x1f647;
{{- else if eq $s "ふむむ" -}}&#x1f914;
{{- else if eq $s "はぁと" -}}&#x1f49b;
{{- else if eq $s "キーボード" -}}&#x2328;
{{- else if eq $s "はなまる" -}}&#x1f4ae;
{{- else if eq $s "錠前と鍵" -}}{{- emojify ":closed_lock_with_key:" -}}
{{- else -}}{{ emojify $s }}{{- end -}}
</abbr>{{- end -}}
```

これで `{{</* emoji ":zzz:" */>}}` などとすれば

```html
<abbr title=":zzz:">💤</abbr>
```

という感じに展開してくれる。
あとは自分で定義名を作れば別名定義もできるし対応する絵文字を増やすこともできる。
上の shortcode だと

|                                字形                                | shortcode                        | CLDR Short Name |
|:------------------------------------------------------------------:| -------------------------------- | --------------- |
|   {{< span class="larger" >}}{{< emoji "ゴメン" >}}{{< /span >}}   | `{{</* emoji "ゴメン" */>}}`     | person bowing   |
|   {{< span class="larger" >}}{{< emoji "ふむむ" >}}{{< /span >}}   | `{{</* emoji "ふむむ" */>}}`     | thinking face   |
|   {{< span class="larger" >}}{{< emoji "はぁと" >}}{{< /span >}}   | `{{</* emoji "はぁと" */>}}`     | yellow heart    |
| {{< span class="larger" >}}{{< emoji "キーボード" >}}{{< /span >}} | `{{</* emoji "キーボード" */>}}` | keyboard        |
|  {{< span class="larger" >}}{{< emoji "はなまる" >}}{{< /span >}}  | `{{</* emoji "はなまる" */>}}`   | white flower    |
|  {{< span class="larger" >}}{{< emoji "錠前と鍵" >}}{{< /span >}}  | `{{</* emoji "錠前と鍵" */>}}`   | locked with key |

てな感じで定義している。

## 【2020-10-09 追記】 絵文字のフォント指定 {#fonts}

絵文字に関しては [Google Fonts](https://fonts.google.com/) みたいなメジャーな Web フォントはないっぽいので各自ローカル環境のフォントをあてにするしかないんだけど（もう自前で Web フォントを用意したくない），これに関して [Unicode 本家サイト](https://unicode.org/emoji/charts/full-emoji-list.html "Full Emoji List")で

```css
.chars, .charsSmall {
    text-decoration: none;
    color: #ccc;
    font-family: "Noto Color Emoji", "Apple Color Emoji", "Segoe UI Emoji",
        Times, Symbola, Aegyptus, Code2000, Code2001, Code2002, Musica, serif,
        LastResort;
    font-variant-emoji: emoji;
}
```

みたいな記述があったので，ありがたくパクらせていただくことにした（笑） すなわち

```css
abbr.emoji-chars {
  text-decoration: none;
  font-family: "Noto Color Emoji", "Apple Color Emoji", "Segoe UI Emoji", Times, Symbola, Aegyptus, Code2000, Code2001, Code2002, Musica, serif, LastResort;
  font-variant-emoji: emoji;
}
```

と定義して

```html
<abbr class="emoji-chars" title=":zzz:">💤</abbr>
```

となるよう shortcode を書き直せばいいわけだ。
当面はこれで。

{{< div-box type="markdown" >}}
**【2022-09-18 追記】**

Noto Color Emoji は [Google Fonts](https://fonts.google.com/) で Web フォントとして指定できるようになっていた。
ので，このブログサイトは Web フォントの絵文字を優先するようにした。
{{< /div-box >}}

## ブックマーク

- [Full Emoji List](https://unicode.org/emoji/charts/full-emoji-list.html) : based on Unicode CLDR data
- [emojify | Hugo](https://gohugo.io/functions/emojify/)
- [絵文字一覧 🤣 | Let's EMOJI](https://lets-emoji.com/emojilist/)
- [UnicodeのEmojiの一覧 - Wikipedia](https://ja.wikipedia.org/wiki/Unicode%E3%81%AEEmoji%E3%81%AE%E4%B8%80%E8%A6%A7)
- [twitter/twemoji: Emoji for everyone. https://twemoji.twitter.com/](https://github.com/twitter/twemoji)
- [googlefonts/noto-emoji: Noto Emoji fonts](https://github.com/googlefonts/noto-emoji)

[Zenn]: https://zenn.dev/ "Zenn｜プログラマーのための情報共有コミュニティ"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
<!-- eof -->
