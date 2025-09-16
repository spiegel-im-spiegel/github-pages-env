+++
title = "Intel One Mono フォントに換装する"
date =  "2025-09-16T12:31:07+09:00"
description = "Google Fonts にも収録されていた"
image = "/images/attention/kitten.jpg"
tags = [ "font", "ubuntu", "web", "site" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
  jsx = false
+++

五十路に入ってからから近視・乱視・老眼の三重苦になり，[2種類の眼鏡]({{< ref "/remark/2021/08/age-of-reading-glasses.md" >}} "老眼鏡の季節")を使い分けてるのだが，特にモニタでプログラム・コードを読むのがしんどくなりつつある。

というところで以下の記事を見かけた。

- [Intelが開発したフォント「Intel One Mono」、目が悪くても読みやすいコーディング向けフォント【レビュー】 - 窓の杜](https://forest.watch.impress.co.jp/docs/review/2046577.html)

ひょっとしてこれは私向きか？

{{< fig-img-quote src="./239321698-6c921cf4-f614-41bd-a909-363bb19f9a30.png" title="intel/intel-one-mono: Intel One Mono font repository" link="https://github.com/intel/intel-one-mono" width="5334" lang="en" >}}

私は古い人間なので，ゼロは “$\emptyset$” (empty set) みたいに斜線が入ってないと嫌なのよ（だから今までずっと [Inconsolata] を使っていた）。
数字の `1` と大文字の `I` と小文字の `l`，小文字の `i` と `j` の区別も分かりやすいし，記号類もデフォルメされてる感じで見やすい。
これならイケそうかな。
早速ローカルにインストールしてみよう。

リポジトリの[リリースページ](https://github.com/intel/intel-one-mono/releases "Releases · intel/intel-one-mono")にある最新版の `otf.zip` (zip 圧縮されている) を取ってくる。
Ubuntu の場合は，展開して取得した `*.otf` ファイルを全て `/usr/local/share/fonts/` または `~/.local/share/fonts/` ディレクトリに放り込めばOK。
その後

```text
$ fc-cache -fv
```

でフォントキャッシュを更新するのを忘れないように。
上手くインストールできれば

```text
$ fc-list | grep Intel
/usr/local/share/fonts/IntelOneMono-Light.otf: Intel One Mono,Intel One Mono Light:style=Light,Regular
/usr/local/share/fonts/IntelOneMono-Italic.otf: Intel One Mono:style=Italic
/usr/local/share/fonts/IntelOneMono-Medium.otf: Intel One Mono,Intel One Mono Medium:style=Medium,Regular
/usr/local/share/fonts/IntelOneMono-Regular.otf: Intel One Mono:style=Regular
/usr/local/share/fonts/IntelOneMono-MediumItalic.otf: Intel One Mono,Intel One Mono Medium:style=Medium Italic,Italic
/usr/local/share/fonts/IntelOneMono-LightItalic.otf: Intel One Mono,Intel One Mono Light:style=Light Italic,Italic
/usr/local/share/fonts/IntelOneMono-Bold.otf: Intel One Mono:style=Bold
/usr/local/share/fonts/IntelOneMono-BoldItalic.otf: Intel One Mono:style=Bold Italic
```

みたいな感じにリスト表示できる。

VS Code で [Intel One Mono] フォントを使うには “Font Family” 項目で

{{< fig-img src="./vscode-setting.png" title="Font Family の設定例" link="./vscode-setting.png" >}}

という感じに `'Intel One Mono'` または `'Intel One Mono Medium'` を最初に指定すればよい。
VS Code は複数フォントを（並び順で優先順位を決めて）指定できるのが嬉しい。

[リポジトリ][Intel One Mono]を見ると現在最新の V1.4.0 のリリースは 2024-07-26 らしい。
これって Google Fonts にも収録されている？ と思って探したらありました。

- [Intel One Mono - Google Fonts](https://fonts.google.com/specimen/Intel+One+Mono)

ってことは，Web フォントとしても使えるってことか。
このブログサイトでの設定は以下のようにしてみた。

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=BIZ+UDGothic:wght@400;700&family=BIZ+UDMincho:wght@400;700&family=Intel+One+Mono:ital,wght@0,300..700;1,300..700&family=Noto+Color+Emoji&family=Noto+Sans:ital,wght@0,100..900;1,100..900&family=Noto+Serif:wght@100..900&display=swap" rel="stylesheet">
```

試しに何かコードを表示してみよう。
今年の始めに紹介した2025年の干支を求める Go プログラムを表示してみる。

{{< fig-quote class="nobox" type="markdown" title="2025年も心臓リハビリ＠がんばらない" link="/remark/2025/01/cardiac-rehabilitation-2025/" >}}
```go
package main

import (
    "fmt"

    "github.com/goark/koyomi/zodiac"
)

func main() {
    year := 2025
    干, 支 := zodiac.ZodiacYearNumber(year)
    fmt.Printf("%d年は%v%v，恵方は%v (%v°)", year, 干, 支, 干.DirectionJp(), 干.Direction())
}
```
{{< /fig-quote >}}

上手く表示できてるかな。
皆さんの環境ではどうだろう。

括弧などの記号が見やすいのがよい。
エディタと Web フォントは [Intel One Mono] でいこう。
私文書は [Inconsolata] のままでいいかな。
そもそも紙（含PDF）に書くような文書でコードを書くことは少ないしな。

## ブックマーク

- [結局 Google Fonts に巻き戻した。そしてモリサワ BIZ UD フォント採用へ]({{< ref "/remark/2022/09/rollback-web-fonts.md" >}})

[Intel One Mono]: https://github.com/intel/intel-one-mono "intel/intel-one-mono: Intel One Mono font repository"
[Inconsolata]: http://www.levien.com/type/myfonts/inconsolata.html "Inconsolata"
