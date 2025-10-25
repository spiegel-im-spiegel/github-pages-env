+++
title = "Typst 0.14 がリリースされた"
date =  "2025-10-25T17:03:16+09:00"
description = "個人的に気になるのは PDF/UA-1 (ISO 14289-1) をサポートした点かな。"
image = "/images/attention/tools.png"
tags  = [ "typst","pdf" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Typst] 0.14 がリリースされた。

- [Typst: Typst 0.14: Now accessible – Typst Blog](https://typst.app/blog/2025/typst-0.14/)
- [Release Version 0.14.0 (October 24, 2025) · typst/typst · GitHub](https://github.com/typst/typst/releases/tag/v0.14.0)

0.14 のハイライトは以下の通り：

{{< fig-quote type="markdown" title="Typst: Typst 0.14: Now accessible" link="https://typst.app/blog/2025/typst-0.14/" lang="en" >}}
- [Accessibility](https://typst.app/blog/2025/typst-0.14/#accessibility)
- [PDF standards](https://typst.app/blog/2025/typst-0.14/#pdf-standards)
- [PDFs as images](https://typst.app/blog/2025/typst-0.14/#pdfs-as-images)
- [Character-level justification](https://typst.app/blog/2025/typst-0.14/#character-level-justification)
- [Richer HTML export](https://typst.app/blog/2025/typst-0.14/#richer-html-export)
- [Migrating to Typst 0.14](https://typst.app/blog/2025/typst-0.14/#migrating)
- [Community Call](https://typst.app/blog/2025/typst-0.14/#community-call)
{{< /fig-quote >}}

日本語での解説は以下の記事が参考になる。

- [Typst 0.14.0 の内容を早めに深堀り](https://zenn.dev/monaqa/articles/2025-10-24-typst-updates-v0-14)

個人的に気になるのは PDF/UA-1 (ISO 14289-1) をサポートした点かな。

タグ付き PDF の標準としては PDF/A があるけど[^p1]， PDF/UA はアクセシビリティを重視したもののようで，米国の議会図書館では「望ましい文書形式」として PDF/A より PDF/UA のほうが優先順位が高いらしい。
一方 PDF/A のほうは，アーカイブ用途での長期保存を目的としていてセキュリティ設定が禁止されている。
この辺は使い分けが必要かな。

[^p1]: PDF/A については，拙文「[LuaLaTeX で PDF/A を構成する]({{< ref "/remark/2020/06/pdfa-with-luatex.md" >}})」を参考にどうぞ。

なお [Typst] は PDF/A もフルサポートしている。

他に注意点としては関数名やシンボル名の一部が deprecated になってるので “[Migrating to Typst 0.14](https://typst.app/blog/2025/typst-0.14/#migrating)” を確認したほうがいいだろう。

## ブックマーク

- [PDF/UA-1(ISO 14289-1)について | アンテナハウス PDF資料室](https://www.antenna.co.jp/PDF/reference/PDFUA-1.html)

- [Typst のお勉強]({{< rlnk "/typst/" >}})

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
