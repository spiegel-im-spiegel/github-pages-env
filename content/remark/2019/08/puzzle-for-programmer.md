+++
title = "パズル好きこそプログラマに向いている"
date =  "2019-08-29T22:29:01+09:00"
description = "与えられたルールに沿って問題を作ること，あるいはルールそのものを作ること，これはまさに「プログラミング」である。"
image = "/images/attention/kitten.jpg"
tags = [ "math", "games", "programming", "joke" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter で

<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">先日、某講演を聞きに言った時に某大の講師の方が「世の中にはパズルが得意な人とレゴが得意な人がいる。パズルが得意な人は完成形にまでしか仕上げられない事が多い。我々はレゴが得意な人が欲しい。｣と言ってて分かりやすい例えだなと思った。プログラマにも言えそう。</p>&mdash; mattn＠有益情報 (@mattn_jp) <a href="https://twitter.com/mattn_jp/status/1154184157305589760?ref_src=twsrc%5Etfw">July 25, 2019</a></blockquote>

という tweet を見かけたので脊髄反射してみる。

私はレゴが嫌いである。

ブロック玩具が嫌いなわけではない。
[ナノブロック](http://www.diablock.co.jp/nanoblock/ "nanoblock")とかビンボー人に優しいし面白いと思う。
でもレゴはダメだ！ あれは「資本主義」の象徴である。
「持たざるもの」はスタートラインにすら立てないし，持っていてもその優劣は持っているレゴの「量」で決まるからだ。
ビンボー人のルサンチマンを煽る玩具。
それがレゴである。

Tweet された方や「某大の講師の方」がどのようなつもりで言ったのかは知らないが，パズルの本当の面白さはパズルを「解く」ことではなく「作る」ことにある。
これは12ピースのジグソーパズルから難解な数理パズルまで一貫して言えることだ。

私（の実家）はビンボーでジグソーパズルもめったに買ってもらえなかったが，パズル雑誌は少ない小遣いの中から自力で買って楽しんでいた。
私が学生時代の1980年代にめちゃめちゃ流行ったパズル雑誌に「ニコリ」というのがある（いや，今もあるけど）。
当時の「ニコリ」が他のパズル雑誌と一線を画していたのは読者が「問題」を投稿できることだった（現在はどうなっているか知らない）。

{{< fig-img src="https://photo.baldanders.info/instagram/photos/201812/49f53d2d5e84b37688e4099694c6c544.jpg" title="発掘品。ニコリ便箋。1回だけ懸賞に当たったのよ。" link="https://photo.baldanders.info/instagram/49f53d2d5e84b37688e4099694c6c544/" width="1080" >}}

特に「数独（数字は独身に限る）」は本当に面白かった。
商標の関係で日本では「ナンバープレイス」の本名のほうが有名だが，海外では “sudoku” という単語があり，[数学論文のネタにさえなっている](http://www.nikkei-science.com/page/magazine/0609/sudoku.html "数独の科学 | 日経サイエンス")。

与えられたルールに沿って問題を作ること，あるいはルールそのものを作ること，これはまさに「プログラミング」である。
「問題を解く」なんてのは，ただプログラムを実行しているだけに過ぎない[^sudoku1]。

[^sudoku1]: ちなみに数独を自動で解くプログラムは既に存在している。解くことができるということは作ることもできるということだ。でも機械が作った「問題」って **美しくない** んだよなぁ（笑）

というわけで私は

{{< div-gen type="markdown" class="center" >}}
**パズル好きこそプログラマに向いている**
{{< /div-gen >}}


と声高に主張しよう（笑）

## 余談だが...

私は（設計を除く）プログラミング作業の中で「デバッグ」が一番好きだ。
特に他人の書いたコードのデバッグは極上の数理パズルである。
他人のコードを読むのは勉強になるしね。

デバッグに比べればコーディングなんて脳内のコードを書き写すだけの「写経」に過ぎない。
最も退屈で苦痛な時間。

## ブックマーク

- [「イミテーション・ゲーム」が面白かった。](https://baldanders.info/blog/000833/)
- [Project Euler で遊ぶ]({{< ref "/remark/2018/03/project-euler.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/STARTSIDE-%E3%81%99%E3%81%86%E3%81%A9%E3%81%8F-%E8%84%B3%E3%83%88%E3%83%AC-9%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF%E3%83%91%E3%82%BA%E3%83%AB-%E3%83%96%E3%83%A9%E3%83%83%E3%82%AF/dp/B077PDFQ4S?psc=1&SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B077PDFQ4S"><img src="https://images-fe.ssl-images-amazon.com/images/I/513eTXxTlJL._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/STARTSIDE-%E3%81%99%E3%81%86%E3%81%A9%E3%81%8F-%E8%84%B3%E3%83%88%E3%83%AC-9%E3%83%96%E3%83%AD%E3%83%83%E3%82%AF%E3%83%91%E3%82%BA%E3%83%AB-%E3%83%96%E3%83%A9%E3%83%83%E3%82%AF/dp/B077PDFQ4S?psc=1&SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B077PDFQ4S">STARTSIDE 数独 すうどく 脳トレ 卓上 ボード ゲーム 9ブロックパズル (ブラック)</a></dt>
    <dd>STARTSIDE</dd>
    <dd>おもちゃ＆ホビー</dd>
    <dd>B077PDFQ4S (ASIN), 4562440102073 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">「数独」がボードゲームに（笑）</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-02">2018-12-02</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
