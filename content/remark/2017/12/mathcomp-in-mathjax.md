+++
title = "ちょこっと MathJax 番外編： mathcomp パッケージの代替え"
date =  "2017-12-22T19:15:52+09:00"
description = "MathJax では \\tcdegree コマンドを持っていないようだ。ググってみたら「Unicode を使え」という身も蓋もない回答が見つかった。"
image = "/images/attention/remark.jpg"
tags = [ "math", "tex", "mathjax", "javascript" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は横道に外れて小ネタでいきます。

1. [ちょこっと MathJax： 初期設定]({{< ref "/remark/2017/09/getting-started-mathjax-1.md" >}})
2. [ちょこっと MathJax： 基本的な数式表現]({{< ref "/remark/2017/09/getting-started-mathjax-2.md" >}})
3. [ちょこっと MathJax： インライン数式と別行立て数式]({{< ref "/remark/2017/10/getting-started-mathjax-3.md" >}})
4. [ちょこっと MathJax 番外編： mathcomp パッケージの代替え]({{< ref "/remark/2017/12/mathcomp-in-mathjax.md" >}}) ← イマココ

たとえば

{{< div-box >}}
直角は ±90°。
{{< /div-box >}}

をインライン数式で書くことを考える。

昔の $\mathrm{\LaTeX}$ は「°」に相当する記号がなく， `90^{\circ}` などと表記していた。
すなわち

```text
直角は $\pm90^{\circ}$。
```

と書いて

{{< div-box >}}
直角は $\pm90^{\circ}$。
{{< /div-box >}}

と表示させていたわけだ。
しかし `\circ` ($\circ$) は角度を表す記号ではなく $+$ や $-$ と同じ2項演算子の記号である。
そこで今は [mathcomp] パッケージの `\tcdegree` コマンドを使って

```html
直角は $\pm90\tcdegree$。
```

とするのが正しいそうだ[^ang1]。

[^ang1]: 他の方法としては [siunitx] パッケージを使って `\ang{90}` のように記述する手がある。

ところが [MathJax] では `\tcdegree` コマンドを持っていないようで，上の記述ではエラーになる。
[MathJax] の[機能拡張](https://docs.mathjax.org/en/latest/input/tex/extensions/ "The TeX/LaTeX Extension List — MathJax 3.0 documentation")でもそれらしいものは見当たらず，誰かが公開してるってわけでもないっぽい。

みんなどうやってるんだ？

ググってみたら「Unicode の「°」文字を使うか `\unicode{xb0}` で指定しろ」という身も蓋もない回答が見つかった。
どうもそれしかないらしい... sigh

そこで $\mathrm{\LaTeX}$ との互換性を保つため

```js
MathJax = {
  tex: {
    macros: {
      tcdegree: ['\\unicode{xb0}']
    }
  }
};
```

という感じでマクロを組んでみた。
これで，以下のように表示可能になる。

{{< div-box >}}
直角は $\pm90\tcdegree$。
{{< /div-box >}}

[mathcomp] パッケージで使える記号はかなりあるのだが，私がよく使いそうなもののみ以下に挙げる。

| Macros                                | $\mathrm{\LaTeX}$ Format       | Output                |
|:------------------------------------- |:----------------------- |:--------------------- |
| `tcdegree: ['\\unicode{xb0}']`        | `$35\tcdegree$`         | $35\tcdegree$         |
| `tccelsius: ['\\unicode{x2103}']`     | `$35\,\tccelsius$`      | $35\,\tccelsius$      |
| `tcperthousand: ['\\unicode{x2030}']` | `$35\,\tcperthousand$`  | $35\,\tcperthousand$  |
| `tcmu: ['\\unicode{x3bc}']`           | `$35\,\tcmu\mathrm{s}$` | $35\,\tcmu\mathrm{s}$ |
| `tcohm: ['\\unicode{x3a9}']`          | `$35\,\tcohm$`          | $35\,\tcohm$          | 

- `\tccelsius` は `\tcdegree\mathrm{C}` でもいける（`$35\,\tcdegree\mathrm{C}$` →  $35\,\tcdegree\mathrm{C}$）。これなら華氏にも応用できる（`$35\,\tcdegree\mathrm{F}$` →  $35\,\tcdegree\mathrm{F}$）
- `\tcmu` は `\symup{\mu}` みたいにしたいのだが `\symup` コマンドに相当するものが [MathJax] にないっぽい
- `\tcohm` は `\Omega` と等価，というかフォントのバランスを考えると `\Omega` を使ったほうがいいと思う（`$35\,\Omega$` →  $35\,\Omega$）

## ブックマーク

- [siunitx - TeX Wiki](https://texwiki.texjp.org/?siunitx)

[MathJax]: https://www.mathjax.org/
[mathcomp]: https://ctan.org/pkg/mathcomp "https://ctan.org/pkg/mathcomp"
[siunitx]: https://ctan.org/pkg/siunitx "CTAN: Package siunitx"

## 参考図書 {#books}

{{% review-paapi "4774187054" %}} <!-- [改訂第7版]LaTeX2ε美文書作成入門 -->
