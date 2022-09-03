+++
title = "モリサワ BIZ UD 明朝/ゴシック Web フォント"
date =  "2022-09-03T19:39:48+09:00"
description = "これって JIS X 0208 にない文字は収録されてない？"
image = "/images/attention/kitten.jpg"
tags = [ "web", "font" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter 経由で知ったのだが [Google Fonts](https://www.google.com/fonts/) にモリサワ BIZ UD 明朝/ゴシック フォントってのがあるらしい。

- [Morisawa BIZ UDMincho](https://github.com/googlefonts/morisawa-biz-ud-mincho)
- [Morisawa BIZ UDGothic](https://github.com/googlefonts/morisawa-biz-ud-gothic)

おおっ，2022年春に公開されたのか。
それなら [Bunny Fonts] にもあるかなぁ，と検索してみたら

{{< fig-img src="./morisawa-biz-ud-fonts-in-bunny.png" title="Bunny Fonts | Explore Faster & GDPR friendly Fonts" link="https://fonts.bunny.net/" lang="en" width="763" >}}

おー，あるやないかい！
というわけで，手元の環境でさっそく試してみた。

こっちは従来の NOTO フォント（クリックで拡大）

{{< fig-img src="./japanese-with-noto-font.png" link="./japanese-with-noto-font.png" lang="en" width="1072" >}}

こっちが モリサワ BIZ UD 明朝 フォント（クリックで拡大）

{{< fig-img src="./japanese-with-morisawa-font.png" link="./japanese-with-morisawa-font.png" lang="en" width="1072" >}}

んー，やっぱモリサワのほうが見やすいかなぁ。
全体的に詰まってる感じだが窮屈には見えない。
Web ブラウザで見やすいようにチューニングされている？
役物はモリサワのほうはポロポーショナル・フォントっぽく詰めているようだ。

英文でも試してみよう。

こっちは従来の NOTO フォント（クリックで拡大）

{{< fig-img src="./english-with-noto-font.png" link="./english-with-noto-font.png" lang="en" width="1072" >}}

こっちが モリサワ BIZ UD 明朝 フォント（クリックで拡大）

{{< fig-img src="./english-with-morisawa-font.png" link="./english-with-morisawa-font.png" lang="en" width="1072" >}}

英字だとモリサワのほうが幅広に見えるんだな。
日本語フォントとのバランスを取ってるからか？ この辺は好みだろうな。

というわけでモリサワ BIZ UD 明朝/ゴシックにしようかな，と一瞬思ったのだが，どうも一部表示できない文字があるらしい。
「神（`U+FA19`）」とか「㈱（`U+3231`）」とか。
これって JIS X 0208 にない文字は収録されてないってことかな。
ひょっとして収録グリフ数を制限することで有料ライセンス版[^l1] との差別化を図ってる？ うーむ。

[^l1]: モリサワ BIZ UD 明朝/ゴシックは [SIL Open Font License, Version 1.1](https://scripts.sil.org/cms/scripts/page.php?site_id=nrsi&id=OFL) でライセンスされている。このライセンス下のフォントは販売目的でなければ自由に再配布できる（ただし Attribution & Share Alike 条件がつく）。なお，ソフトウェアにバンドルするのであれば販売も許可される。ドキュメントへの埋め込みもOK。

このブログは Unicode の変わった文字もたまに使うのでモリサワ BIZ UD 明朝/ゴシックは向かないか。
というわけで，今回は諦めた。
「機種依存文字（古語）なんか使わないぜ」という方は検討してみてもいいだろう。

## ブックマーク

- [Google Fonts から Bunny Fonts に乗り換える]({{< ref "/remark/2022/06/migrate-to-bunny-fonts-from-google-fonts.md" >}})

[Bunny Fonts]: https://fonts.bunny.net/ "Bunny Fonts | Explore Faster & GDPR friendly Fonts"
<!-- eof -->
