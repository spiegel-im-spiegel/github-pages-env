+++
title = "“Text to Math Converter” に Unicode の業の深さを見る"
date =  "2023-01-01T11:52:00+09:00"
description = "元日早々から Unicode の業の深さを垣間見てしまった"
image = "/images/attention/kitten.jpg"
tags = [ "math", "unicode" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

[『数学ガール』シリーズ](https://www.hyuki.com/girl/)でおなじみ結城浩さんが面白いツール（サービス）を公開されている。

- [Text to Math Converter - textmath.hyuki.net](https://textmath.hyuki.net/)

早速試してみる。

{{< fig-img-quote src="./text-to-math-converter-1.png" title="Text to Math Converter" link="https://textmath.hyuki.net/" width="765" >}}

変換結果の “𝐸 = 𝑚𝑐²” は Twitter や Mastodon の tweet や toot としてそのままコピペできる。

{{< fig-gen >}}
<iframe src="https://fedibird.com/@spiegel/109610792617725371/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://fedibird.com/embed.js" async="async"></script>
{{< /fig-gen >}}

ホンマに「どうなってるんだ？」と20秒くらい考えてしまったが，そうか Unicode 文字で表現してるのか，と気がついた（気づくのが遅い）。
試しに拙作の [gnkf] でダンプしてみる。

```text
$ echo 𝐸 = 𝑚𝑐² | gnkf dump -u
0x0001d438, 0x0020, 0x003d, 0x0020, 0x0001d45a, 0x0001d450, 0x00b2, 0x000a
```

おー，なるほど。

イタリックの英字は「[数学用英数字記号]（Mathematical Alphanumeric Symbols: `U+1D400` 〜 `U+1D7FF`）」として定義されているらしい[^u1]。
あくまで数学記号なので，通常の文章で装飾文字として使うなってあるね。
たしかに[ホモグラフ攻撃]({{< ref "/remark/2017/04/homograph-attack.md" >}} "Punycode によるホモグラフ攻撃例とその回避")とかに使えそうだもんな（笑）

[^u1]: 数学用英数字記号の詳細については公式の “{{< pdf-file title="Mathematical Alphanumeric Symbols" link="https://www.unicode.org/charts/PDF/U1D400.pdf">}}” を参照のこと。

上付き文字の `U+00B2` は ISO/IEC 8859-1 (通称 Latin-1) で定義されているものと同じである[^u2]。

[^u2]: というか Unicode の `U+00FF` までは ISO/IEC 8859-1 と同じである。

{{< fig-img-quote src="./latin-1.png" title="ISO/IEC 8859-1 - Wikipedia" link="https://ja.wikipedia.org/wiki/ISO/IEC_8859-1" width="543" >}}

この表を見ると上付き文字の数字は 1〜3 しかないな。
0 や 4 以降はどうなってるかというと “{{< pdf-file title="Superscripts and Subscripts" link="http://www.unicode.org/charts/PDF/U2070.pdf">}}” (`U+2070` 〜 `U+209F`) で定義されているらしい。

つまり `x^4` → 𝑥⁴ は

```text
$ echo 𝑥⁴ | gnkf dump -u
0x0001d465, 0x2074, 0x000a
```

と展開されているわけだ。

これらを踏まえて “[Text to Math Converter]” では利用可能な文字[^u3] を

[^u3]: 利用可能な文字や記号は随時追加されている。

{{< fig-img-quote src="./text-to-math-converter-2.png" title="Text to Math Converter" link="https://textmath.hyuki.net/" width="701" >}}

としているようだ（`a e o x` の下付き文字 `U+2090` 〜 `U+2093` がない理由は不明）。
これ以外の文字は変換されずにそのまま出力されるみたい。
たとえば `|a_m - a_n| < d` → |𝑎ₘ - 𝑎ₙ| < 𝑑 とかいった感じ。

しかし，まぁ，元日早々から Unicode の業の深さを垣間見てしまった感じである。
でも “[Text to Math Converter]” 自体は面白いので，これからお世話になるかもしれない。
もちろん，このブログでは今までどおり MathJax で `$E = mc^2$` → $E = mc^2$ と表記するので，本年もよろしくおねがいします。

## ブックマーク

- [ちょこっと MathJax]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})

[Text to Math Converter]: https://textmath.hyuki.net/ "Text to Math Converter - textmath.hyuki.net"
[gnkf]: https://github.com/goark/gnkf "goark/gnkf: Network Kanji Filter by Golang"
[数学用英数字記号]: https://ja.wikipedia.org/wiki/%E6%95%B0%E5%AD%A6%E7%94%A8%E8%8B%B1%E6%95%B0%E5%AD%97%E8%A8%98%E5%8F%B7 "数学用英数字記号 - Wikipedia"

## 参考図書

{{% review-paapi "4297117126" %}} <!-- LaTeX2ε美文書作成入門 -->

