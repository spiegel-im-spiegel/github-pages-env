#import "js.typ": *

#show: js.with(
  lang: "ja",
  seriffont: "New Computer Modern",
  seriffont-cjk: "BIZ UDMincho",
  sansfont: "Liberation Sans",
  sansfont-cjk: "BIZ UDGothic",
  paper: "a4",
  fontsize: 10pt,
  baselineskip: auto,
  textwidth: auto,
  lines-per-page: auto,
  book: false,
  cols: 1,
  non-cjk: "latin-in-cjk",
  cjkheight: 0.88,
)

#maketitle(
  title: "Go のエラーハンドリング",
  authors: "Spiegel",
  abstract: [
    Go のエラーハンドリングについて今まで書いた駄文を「全部入り」で本の形にしてみようと思う。
  ],
  keywords: ("golang", "error", "programming"),
  date: "2025年2月24日"
)

#outline() #v(1em)

= はじめに

プログラミングにおいて，正常系は基本的に「一本道」だが，異常系は（予期しないものも含めて）無数にある。
エラーハンドリングは巨大迷路パズルを袋小路から順に塗りつぶして「正解」をあぶり出していく作業に似ていると思う。
下手くそな迷路攻略はただの「作業」だが，*よく考えられた迷路は袋小路の配置も美しい*。

こんなふうに考えるなら，プログラム設計の肝は#ruby[エラー][袋小路]をどう記述するかにかかっている，と言えるだろう。

私が Go のエラーハンドリングについて最初に記事にしたのは2015年のことだが，あれから Go も少しずつ変わっているし，私も当時よりは多少なりと理解が進んだと思うので，今まで書き散らかした駄文を「全部入り」で本の形にしてみようと思う。
