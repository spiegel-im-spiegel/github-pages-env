+++
title = "松江に出戻り5度目の冬（Advent Calendar）"
date =  "2022-12-09T00:00:01+09:00"
description = "「年寄りの冷や水」という感じで眺めていただければと思います。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "calendar", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（本ページは「[山陰 Advent Calendar 2022](https://adventar.org/calendars/8245 "山陰 Advent Calendar 2022 - Adventar")」9日目の記事です）

はじめましての方ははじめまして。
Spiegel と申します。
簡単なプロフィールは[こちら](https://baldanders.info/profile/ "公開履歴書 | Baldanders.info")。

長いこと広島市に住んでいたのですが4年前（1998年末）に生まれ故郷の松江市に ~~撤退~~ おっと転進しまして，以後は「死の行軍」で本当に死にそうになることもなく，かといって歓楽街で夜遊びすることもなく，比較的のんびり暮らしています。
（体調面の不安もあり）帰郷直後は IT 業界への復帰を諦めていたのですが，幸運にも拾ってくださった地元企業がありまして，こちらにも帰ってまいりました。

今回参加する Advent Calendar は「せっかくだから、ゆるっと記事を書いてみませんか？」ということなので，アドバイス通り，ゆる～く行きたいと思います。
とはいえ，特に後半の話は以前から書きたかった内容ですので「年寄りの冷や水」という感じで眺めていただければと思います。

----

プログラミング言語にも「母国語」と言えるものがあります。
たとえば「数学ガール」シリーズやデザパタ本などでおなじみ結城浩さんの「連ツイ」にも以下のような記述があります。

{{< fig-quote type="markdown" title="考えてみると結城はC, Perl, Javaの本は書いたけれどRubyやJavaScriptの本はまだ書いていないですね（出版社からは、Ruby!とかJavaScript!と言われていますけれど）。でも、結城が現在日常的に書いている言語はRubyとJavaScriptになります。 - 結城浩の連ツイ" link="http://rentwi.hyuki.net/?666213569055166464s" >}}
プログラミング言語との付き合い方というのはいろいろあってですね。自分の母国語という言語はある。それから現在学んでいる最中の言語というのもある。そして、仕事用の言語やら、他の人とのコミュニケーション用言語というのもある。そのあたりは、自然言語とちょっと似ている。

ITな業界で仕事をしているひとというのは、だいたいそういう感じでプログラミング言語とつきあっていると思っている。たった一つしかプログラミング言語ができないという人は少なくて、二つくらいは読み書きできる。三つ四つくらいはなんとなく読むのはできる。五つ六つくらいは何のソースか言える。
{{< /fig-quote >}}

私の場合は出自が組込みソフトウェア・エンジニアなのでアセンブラと C 言語が「母国語」と言えるものでしたが，2015年頃から趣味で [Go 言語][Go]を習い始めまして，近年は Go が「母国語」になってきています。

物凄く簡単に言うと，昔は

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
}
```

を C で書いたらどうなるかというのを脳内でほぼ無意識に「翻訳」してましたが，今は逆に K&R の

```c
#include <stdio.h>

main()
{
    printf(“hello, world\n”);
}
```

を Go で書くとどうなるか，という感じでしょうか。

仕事で様々なプログラミング言語を渡り歩いて来ましたが「母国語」が変わるほどのインパクトは初めての経験で，今も楽しく Go で遊んでいます。
願わくば，もう少し仕事に繋げられるといいのですが。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

2020年のパンデミック宣言以後[なかなか収まる気配のない COVID-2019](https://github.com/spiegel-im-spiegel/covid-2019-report) ですが，おかげさまというか何というか，いわゆる「リモートワーク」が普通になりつつありオンライン・イベントもたくさん開催されるようになったのは，田舎に引っ込んだ身としてはありがたい機会でした。

特に『[プログラミング言語Go](https://www.amazon.co.jp/dp/B099928SJD?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』や『[Go言語による分散サービス](https://www.amazon.co.jp/dp/4873119979?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の翻訳者でもある[柴田芳樹](https://yshibata.blog.ss-blog.jp/)さん主催のオンライン読書会は勉強になっています。

その柴田芳樹さんのブログ記事で興味深いものがありまして。

- [ソフトウェアエンジニアの成長カーブ（再掲載）：柴田 芳樹 (Yoshiki Shibata)：SSブログ](https://yshibata.blog.ss-blog.jp/2013-10-10)

この記事で示されている図

{{< fig-img-quote src="https://yshibata.c.blog.ss-blog.jp/_images/blog/_d3f/yshibata/E68890E995B7E382ABE383BCE38396.JPG" title="ソフトウェアエンジニアの成長カーブ（２）：柴田 芳樹 (Yoshiki Shibata)：SSブログ" link="https://yshibata.blog.ss-blog.jp/2010-01-27" width="1057" >}}

にはなかなか考えさせられるものがあります。

この図が示すのは，会社から与えられている仕事に慣れて自律[^ai1] 的に「学ぶ」ことを止めてしまうと，手持ちのスキルは（時代の流れで）先細りし，さらに学習習慣が失われることにより新たな技術・スキルを得る機会を逸してしまうため，全体としてスキルレベルが低下してしまう，というものです。

[^ai1]: 『[そろそろ、人工知能の真実を話そう](https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』によると「自律」というのは元々哲学用語で「自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう」のだそうです。

特に日本企業は従業員の「今」のスキルレベルに合わせて「できそうな仕事」を見繕って割り振る傾向があり，所属部署から与えられる仕事をこなすことに満足しているとあっという間にスキルが先細りしてしまう，というような話を読書会でもしていました。

この辺は私も耳が痛い部分があります。
お金をもらって技術とスキルを発揮するプロのエンジニアであれば，いくつになっても何処にいても，学ぶ姿勢と習慣を捨ててはいけないということでしょう。
でも，どうせなら楽しく学びたいものです。

{{< fig-quote title="数学ガール" link="https://www.amazon.co.jp/dp/B00EYXMA9I?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
僕たちは好きで学んでいる。
<br>先生を待つ必要はない。授業を待つ必要はない。
<br>本を探せばいい。本を読めばいい。
<br>広く、深く、ずっと先まで勉強すればいい
{{< /fig-quote >}}

[Go]: https://go.dev/

## ブックマーク

- [チャンスは待たずに自分で作る ─ ソフトウェアエンジニアが「好きな技術」で生きていくための技術とは - Findy Engineer Lab - ファインディエンジニアラボ](https://engineer-lab.findy-code.io/tenntenn-go)

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "4873119693" %}} <!-- 実用 Go 言語 -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4873119979" %}} <!-- Go言語による分散サービス -->
{{% review-paapi "4908686122" %}} <!-- Goならわかるシステムプログラミング 第2版 -->
{{% review-paapi "B09C2XBC2F" %}} <!-- Golang Tシャツ -->
