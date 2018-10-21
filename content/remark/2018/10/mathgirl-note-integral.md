+++
title = "『数学ガールの秘密ノート／積分を見つめて』の感想を書くのを忘れていた"
date = "2018-10-21T17:44:16+09:00"
update = "2018-10-21T18:21:51+09:00"
description = "今さら読み直して感想をアップしてみる。1年以上も放置とか orz"
image = "/images/attention/kitten.jpg"
tags = ["book", "math", "integral"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

（若干のネタバレあり）

前回「[数学ガールの秘密ノート]」シリーズ[最新刊の感想]({{<ref "/remark/2018/10/mathgirl-note-matrix.md" >}} "『数学ガールの秘密ノート／行列が描くもの』を読んで子供の頃を思い出す")を書いた後に気がついた。
『[積分を見つめて]』の感想を書いてないぢゃん。

ちうわけで，今さら読み直して感想をアップしてみる。
ちなみに『[積分を見つめて]』は2017年6月に出版されている。
1年以上も放置とか `orz`

## 『[微分を追いかけて]』の感想（再掲載）

ニュートン以来，微分と積分はペアで考えるのが正しい。

というわけで，以前に[本家サイトで書いた『微分を追いかけて』の感想](http://www.baldanders.info/spiegel/log2/000839.shtml "『数学ガールの秘密ノート』 微分で遊ぶ — Baldanders.info")を再掲載する。

{{% div-box title="『数学ガールの秘密ノート』 微分で遊ぶ — Baldanders.info" link="http://www.baldanders.info/spiegel/log2/000839.shtml" %}}
毎度言っているが，「数学ガールの秘密ノート」シリーズは数学成分多めで中学生以上を対象にしているが，小学生高学年なら頑張れば理解できるはず。 
てか，是非挑戦して欲しい。
「かけ算の順序」とか「$6\div2\sqrt{3}$」とか瑣末なことに躓いている場合ではないのだ。 

個人的には第2章かな。 

```
ユーリ「テストに出てきたから使うんじゃなくて《自分の計算》で使ったのは初めて、ってこと」
僕「なるほど……なるほど」
```

このやりとりに痺れまくりですよ。
若いっていいなぁ。 

[「丸い三角関数」の感想](http://www.baldanders.info/spiegel/log2/000685.shtml "「丸い三角関数」を読む 他 — Baldanders.info")でも書いたけど，三角関数や微分方程式は物理学，特に天文学を理解するには必須の道具。
学校で教えてくれるのを待ってる暇はないのだよ。
そして三角関数や微分（と積分）を理解すると理科も数学も抜群に面白くなる。
「分かった！」の振れ幅が桁違いになるのだ。 

たとえば，円周の長さを表す $2{\pi}r$ が円の面積 ${\pi}r^2$ の $r$ に対する微分だと気づけば円のイメージがより明確になるし，何よりおぼえなくてはならない「公式」が劇的に減る。
第1章および第2章で出てくる位置と速度と加速度の関係もそうだよね。
位置と速度と加速度の関係が理解できると，たとえば「人工衛星は地球の重力に引っ張られて「落ち続けてる」のに，何故地球にぶつからないのか」といったことも理解できるようになる（と思う）。 
{{% /div-box %}}

## 『[積分を見つめて]』の感想

では本題。

つか，円周の長さと円の面積の関係は『[積分を見つめて]』で思いっきり登場する。
実際に5章で円の面積を求めるし。

微積分の話なので当然ながら「極限」が頻繁に登場する[^lim1]。
実は「極限」の話で個人的にお気に入りなのは茉崎ミユキさんのコミック版『[数学ガール ゲーデルの不完全性定理]』だったり。
テトラちゃんが可愛いのよ。
学園ラブコメはこうあるべきだよね（笑）

[^lim1]: 「極限」の話は『[数列の広場]』にも出てくる。

そういえば大昔の受験生時代に，某大学の入試問題で方程式を解くのにニュートン法（Newton-Raphson method）を使えってのがあって往生した覚えがある。
『[積分を見つめて]』2章のはさみうちの話で何故かそれを思い出して[^ni1]，ちょっと遠い目になってしまった（笑）

[^ni1]: 『[積分を見つめて]』2章のはさみうちの話は「区分求積法」と呼ばれる解法。念のため。

『[積分を見つめて]』で特に面白かったのは3章の最後にある「付録」。
「定積分を定義する二つのスタイル」だ。
以下のスタイルに分けられるらしい。

1. 面積で定積分を定義する
2. 原始関数で定積分を定義する

学校教育では 2 のスタイルで教えることが多いそうな。
多分このスタイルは数式の暗記を前提としたもので，理解などすっ飛ばして「とにかく問題を解く」ことに向いているのではないかと思う。

でも，微分と積分をきちんと関連付けて体系的に理解を得たいなら 1 から始めるのが断然お薦めである。
『[積分を見つめて]』は 1 のスタイルで話が進んでいく。

すでにたくさん出ている「[数学ガールの秘密ノート]」シリーズだが，個人的なお薦めは

> 『[数列の広場]』 → 『[微分を追いかけて]』 → 『[積分を見つめて]』

と繋げていく読み方である。
『[数列の広場]』で「極限」の感覚に慣れ『[微分を追いかけて]』と『[積分を見つめて]』で微積分について理解を進めるのがいいのではないだろうか。

『[微分を追いかけて]』の感想でも書いたが，微積分そのものを理解することは難しくない。
学校で教えてくれるのを待つ必要はないのだ。
まぁ，公式や公式から導き出されるものを丸呑みしようとするとけっこう大変だとは思うけど（そゆときは公式集を持ち歩けばよい。試験以外であんちょこを落ち歩くのは悪ではない）。

ところで今回は「僕」とミルカさんとテトラちゃんが上手く役割分担してて，いい感じに対話している。
まぁミルカさんの出番は少なめだったけど（「秘密ノート」はテトラちゃんの持ち物なのでテトラちゃんの比率が多いのは仕方ない）。

個人的にはもっとリサちゃんを出して欲しいです。

## ブックマーク

- [微分積分法(びぶんせきぶんほう)とは - コトバンク](https://kotobank.jp/word/%E5%BE%AE%E5%88%86%E7%A9%8D%E5%88%86%E6%B3%95-1400205)
- [はじめMath! Javaでコンピュータ数学：連載｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/serial/01/java-calculation)
    - [第72回　微分・積分の数学　数値積分　区分求積法・台形公式［前編］ ：はじめMath! Javaでコンピュータ数学｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/serial/01/java-calculation/0072)
    - [第73回　微分・積分の数学　数値積分　区分求積法・台形公式［後編］ ：はじめMath! Javaでコンピュータ数学｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/serial/01/java-calculation/0073)
    - [第67回　微分・積分の数学　ニュートン・ラフソン法　［前編］ ：はじめMath! Javaでコンピュータ数学｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/serial/01/java-calculation/0067)
    - [第68回　微分・積分の数学　ニュートン・ラフソン法　［後編］ ：はじめMath! Javaでコンピュータ数学｜gihyo.jp … 技術評論社](https://gihyo.jp/dev/serial/01/java-calculation/0068)

[数学ガールの秘密ノート]: http://www.hyuki.com/girl/#note "『数学ガール』シリーズ"
[積分を見つめて]: https://www.amazon.co.jp/exec/obidos/ASIN/B073F45B97/baldandersinf-22 "数学ガールの秘密ノート／積分を見つめて | 結城 浩 | 数学 | Kindleストア | Amazon"
[微分を追いかけて]: https://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22 "数学ガールの秘密ノート／微分を追いかけて | 結城 浩 | 数学 | Kindleストア | Amazon"
[数列の広場]: https://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22 "数学ガールの秘密ノート／数列の広場 | 結城 浩 | 数学 | Kindleストア | Amazon"
[数学ガール ゲーデルの不完全性定理]: https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4MS/baldandersinf-22 "数学ガール　ゲーデルの不完全性定理　1 (MFコミックス　アライブシリーズ) | 茉崎 ミユキ | 少年マンガ | Kindleストア | Amazon"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B073F45B97/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41WTBPbSEtL._SL160_.jpg" width="115" height="160" alt="数学ガールの秘密ノート／積分を見つめて"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B073F45B97/baldandersinf-22">数学ガールの秘密ノート／積分を見つめて</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">三角関数や微積分は物理学，特に天文学を理解するには必須の道具。学校で教えてくれるのを待ってる暇はないのだよ。そして，これらを通してものの位置や形について理解が進むと理科も数学も抜群に面白くなる。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.21</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41n9NfuGsIL._SL160_.jpg" width="113" height="160" alt="数学ガールの秘密ノート／数列の広場"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22">数学ガールの秘密ノート／数列の広場</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">「無限」と「極限」について。頑張れば小学生の高学年くらいなら理解可能。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.21</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41pgiwRb0zL._SL160_.jpg" width="111" height="160" alt="数学ガールの秘密ノート／微分を追いかけて"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22">数学ガールの秘密ノート／微分を追いかけて</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">三角関数や微分方程式は物理学，特に天文学を理解するには必須の道具。学校で教えてくれるのを待ってる暇はないのだよ。そして三角関数や微分（と積分）を理解すると理科も数学も抜群に面白くなる。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.21</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4MS/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/51Hz%2BE6YF0L._SL160_.jpg" width="113" height="160" alt="数学ガール　ゲーデルの不完全性定理　1 (MFコミックス　アライブシリーズ)"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00AXUD4MS/baldandersinf-22">数学ガール　ゲーデルの不完全性定理　1 (MFコミックス　アライブシリーズ)</a></dt>
    <dd>茉崎 ミユキ</dd>
    <dd>KADOKAWA / メディアファクトリー</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="far fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">原作『<a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22">数学ガール／ゲーデルの不完全性定理</a>』のコミカライズ版。茉崎ミユキさんの描くテトラちゃんがめっさ可愛いの！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.10.21</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
