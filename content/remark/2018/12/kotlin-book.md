+++
title = "Kotlin の予備学習"
date = "2018-12-02T15:07:36+09:00"
description = "「Kotlin をふわっと覚えられる」教材を探していた。 そしたら偶々 Kindle unlimited で『速習 Kotlin』という本を見かけたので予備学習教材として読んでみた。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "kotlin", "programming", "language", "object-oriented" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [The State of the Octoverse: top programming languages of 2018 | The GitHub Blog](https://blog.github.com/2018-11-15-state-of-the-octoverse-top-programming-languages/)

この記事の中で “Fastest growing languages by contributors as of September 30, 2018” のトップが [Kotlin] なのを見て「やっぱ時代は [Kotlin] だ！」と確信したのだが，今は忙しいので，まずは「[Kotlin] をふわっと覚えられる」教材を探していた。
そしたら偶々 Kindle unlimited で『[速習 Kotlin]』という本を見かけたので予備学習教材として読んでみた。

[Kotlin] は Android アプリ用言語として有名になったせいか，言語そのものを解説する本はまだ少ない感じ。
個人的には [Kotlin]/Native に期待しているので，もっと言語そのものを解説する本がたくさん出るといいな，と思ったりする。

さて『[速習 Kotlin]』を読んだ第一印象だが，昔 .NET Framework が登場したときに真っ先に「覚えられるか！ こんなもん」と感じたのを思い出した。

[Kotlin] に限らないが，コード記述の全てを「オブジェクト」およびそれらの関係で表そうとすると膨大な語彙や文脈が発生する[^go1]。
故に「言語を覚える」というのは言語仕様を理解するだけでは全然足りなくて，言語に付随する語彙つまりフレームワークを習得する必要がある。
無理やり喩えるなら，昔の（C言語などの）言語は「数独[^sd1]」でオブジェクト指向言語は「クロスワードパズル」といった感じだろうか[^pz1]。

[^go1]: この点で [Go 言語]における全てを「オブジェクト指向」に包摂しようとしない設計上の割り切り方は個人的に好印象である。スクリプト言語ならともかく，バイナリ・コードにコンパイルする言語は（泥臭さを許容してでも）パフォーマンスが要求されることも多く，無理にオブジェクト指向でマニピュレータを書くより「for 文で回したほうがはやい」なんてなこともあるのだ。
[^sd1]: 「数独（「数字は独身に限る」の略称）」は[ニコリ](https://www.nikoli.co.jp/)の商標である。ちなみに数独のオリジナルは米国の方が考えられた ”Number Place” なのだが，海外では “Sudoku” の名称のほうが有名である。詳しくは日経サイエンス 2006年9月号「[数独の科学](https://www.nikkei-science.net/modules/flash/index.php?id=200609_052)」を参照のこと。
[^pz1]: クロスワードパズルは出題者と解答者との間で共通の語彙がなければ絶対に解けない。

この手の解説本でありがちなミスは，言語「仕様」と言語周辺の「知識」がゴッチャになって記述が混沌としてしまうことだ。
『[速習 Kotlin]』にもこの傾向がある。
とはいえ Int や Double といったプリミティブな型でさえクラス・オブジェクトとして定義されているのだから，この点で著者の方を責めるのは酷というものだろう。

もうひとつ気になったのが糖衣構文（syntax sugar）。
既にその言語に馴染んでいる人にとっては糖衣構文は書きやすくて読みやすいのだろうが，基本のキから言語仕様を学ぼうという人にとっては毒が強すぎる気がする。
これって習熟度が異なる者同士で peer review すると支障が出たりしないのか？ 私なら研修である程度言語に慣れるまでは糖衣構文は禁止にするぞ。

というわけで，私なら

1. 言語仕様の理解
2. 糖衣構文に慣れる練習
3. フレームワークの知識を深める

の三段構えで学習を進めていきたいところである。
まずは目の前の雑事を片付けていかないといけないのだが...

[Kotlin]: https://kotlinlang.org/ "Kotlin Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[速習 Kotlin]: https://www.amazon.co.jp/exec/obidos/ASIN/B07HQMNLCV/baldandersinf-22 "Amazon.co.jp： 速習 Kotlin: Javaより簡単！新Android開発言語を今すぐマスター 速習シリーズ eBook: 山田祥寛: Kindleストア"

## 参考図書

{{% review-paapi "B07HQMNLCV" %}} <!-- 速習 Kotlin -->
{{% review-paapi "B077PDFQ4S" %}} <!-- 数独 -->
