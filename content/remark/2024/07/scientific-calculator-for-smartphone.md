+++
title = "スマホ用関数電卓を物色する — 『関数電卓がすごい』を読んだ"
date =  "2024-07-06T10:06:41+09:00"
description = "面白い本だったが特に薦めないことにしよう（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "book", "k-tai", "tools", "math" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

Bluesky の[早川書房](https://bsky.app/profile/hayakawa-online.co.jp "早川書房公式 (@hayakawa-online.co.jp) — Bluesky")アカウントで以下の本が紹介されていて，タイトルだけ見て衝動買いしてしまった。

{{% review-paapi "B0D6QQ3VXT" %}} <!-- 関数電卓がすごい -->

反省はしない。

## 『[関数電卓がすごい]』を読んだ

タイトル通り関数電卓についての話で，関数電卓の使い方についても書かれているのだが，おそらくこの本の主題は「使い方」ではなく，もっと大きな「価値」評価の話をしてるのかなと思った。

少なくとも現代社会における「価値」の多くは，定性評価ではなく，定量評価である（まぁ value って言うくらいだし）。
「量」を評価するのだから計算可能というわけ。
だから電卓でも価値の評価ができるのだ。
『[関数電卓がすごい]』の0章では

{{< fig-quote type="markdown" title="『関数電卓がすごい』0章" link="https://www.amazon.co.jp/dp/B0D6QQ3VXT?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
この世の商売の7割は、数学を嫌がるヒトに対して数学を提供するもの
{{< /fig-quote >}}

などと書かれていてクスってなった。

## 電卓の思ひ出

自分の話で申し訳ないが，実は私は CASIO 派である。
ちょうど学生時代にポケットコンピュータ（通称ポケコン）が出始めて，周囲が SHARP のポケコンばっかりだったのに自分だけ CASIO のポケコンで北極星の時角計算プログラムを（BASIC で！）組んだりして遊んでいた[^pc]。

[^pc]: 今どきの望遠鏡はコンピュータ制御で自動制御が当たり前みたいだが，昔は赤道儀の極軸合わせも人間が手動でやってたので，北極星の時角計算は結構大事だった。

その後ひょんなことから IT 業界に潜り込むことになったのだが，当時勤めていたのが制御系システムのハードとソフトを得意とする会社で，検算のために CASIO の関数電卓を持ち歩いていた。
いやぁ，客先での「現地調整」とかで必須だったのよ，関数電卓。
つっても主に使ってたのは N 進数演算機能だったけどね。

その後いろいろあって，制御系システムの仕事とも疎遠になり，関数電卓を持ち歩くこともなくなり（うっかり水に落として壊れたのだ），代わりにスマートフォンの電卓アプリを使うようになった。
お気に入りは有料の [RealCalc Plus] で，もう10年以上愛用している。
当時はスマホの電卓アプリで N 進数演算ができるのは珍しかったのよ。

## 電卓か電卓アプリか

『[関数電卓がすごい]』の著者の方もそうみたいだが（スマホアプリではなく）電卓を愛用する人はスマホのの取り回しの悪さを挙げるよね。
まぁ，確かにスマホをスリープから叩き起こして，認証を行って，えんやらやっと電卓アプリを立ち上げる。
電卓ならその間に計算終わっとるっちうねん，て感じだろうか。

でもホンマのことを言うと電卓機能ってそんなに使わないんだよね。
更に言うと簡単な四則演算なら検索サービスの検索窓に式を入れれば答えが返ってくる。

{{< fig-img src="./calculator-by-google-search.png" title="Google 先生に計算結果を訊く" link="./calculator-by-google-search.png" width="856" >}}

ブラウザで計算させることのメリットは式をコピペできることだ。
しかもパソコンを開いてたら大抵ブラウザも開いてるしね。
ブラウザ検索の手間なんかあってないようなものである。

それでも『[関数電卓がすごい]』で示されている日常生活に密着した計算例はなかなか面白かった。
左手に Kindle 端末，右手にスマホを持ちながら楽しく読めたですよ。

ただ，今回『[関数電卓がすごい]』を読みながら [RealCalc Plus] を使っていて不満が生じてしまった。
[RealCalc Plus] は本当に電卓的な操作なので，計算履歴から式を弄って色々試してみるといったことができない。
計算履歴は呼び出せるんだけどユーザからは計算結果の値しか参照できないのだ。
これは不便！

というわけで，10年振りに関数電卓アプリを物色してみることにした。

## スマホ用関数電卓アプリの評価基準

物色を始める前に『[関数電卓がすごい]』で書かれている内容も参考にしてスマホ用関数電卓アプリの評価基準を挙げておく。
以下に示す機能・関数が搭載されていれば『[関数電卓がすごい]』で楽しく遊ぶことができる。

### 計算履歴を取得して式が変形できるか

計算履歴を参照・再利用できないアプリは論外。

### 長押しで裏機能が使えるか

大抵の関数電卓は {{% keytop %}}SHIFT{{% /keytop %}} や {{% keytop %}}2<sup>nd</sup>{{% /keytop %}} キー押下で裏機能が使えるが，関数電卓アプリによってはキーの長押しで代替できる場合がある。
めっさ便利！

### N 進数演算機能があるか

2進数・8進数・16進数への変換ができて，それぞれで四則演算ができて，可能ならシフト演算や論理和・論理積の演算ができるとありがたい。

### 単位変換ができるか

メートル法とヤード・ポンド法との相互変換とか。
『[関数電卓がすごい]』ではカロリー計算の単位変換を例示している。

私は自転車の整備でタイヤの空気圧を[調整]({{< ref "/remark/2022/09/pound-force-per-square-inch.md" >}} "はじめての空気圧チェック")する関係で圧力の単位とか結構気にしている。
おなじみの $\mathrm{Pa}$ や $\mathrm{bar}$ だけじゃなく $\mathrm{psi}$ や $\mathrm{atm}$ といった様々な単位が飛び交ってて分かりにくいったらないので，簡単に単位変換ができるのはありがたい。

単位変換で面倒なのが角度や時刻の変換。
{{% keytop %}}$\mathrm{D}{\\tcdegree}\mathrm{M}'\mathrm{S}''${{% /keytop %}} や {{% keytop %}}H:M:S{{% /keytop %}} があれば比較的簡単に入力や変換ができる。
時間の四則演算は正直に言って面倒くさいので，これはかなり助かる。

### 変数操作ができるか

{{% keytop %}}M+{{% /keytop %}} とか {{% keytop %}}M-{{% /keytop %}} とかじゃなくて任意の変数を扱えるか。
変数が使えると試行錯誤がしやすい。

### 1+2×3 が7になるか

昔ながらの電卓は入力した順に逐次評価を行う。
たとえば {{% keytop %}}$1${{% /keytop %}} {{% keytop %}}$+${{% /keytop %}} {{% keytop %}}$2${{% /keytop %}} {{% keytop %}}$\times${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} と入力した場合は，まず $1+2$ を計算し，その結果を $3$ 倍するので答えは $9$ になる。
まともな関数電卓であれば演算子の優先順位を評価して正しく $1+2\times3=7$ となる筈である。

### 1÷3×3 が1になるか

厳密に言うと {{% keytop %}}$1${{% /keytop %}} {{% keytop %}}$\div${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} で一度計算を確定させたあと {{% keytop %}}$\times${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} とした場合の計算結果がどうなるかである。

これは関数電卓によって挙動が異なるらしい。
想像だけど，もし $1\div3$ の計算結果を有理数 $1/3$ として内部的に保持しているなら，それを $3$ 倍すれば $1$ になる筈である。
一方 $1\div3=0.333\cdots$ として保持していれば $3$ 倍すると $0.999\cdots$ となる。

私としては $1\div3\times3$ が $1$ になることを期待したい。
さらに言うと {{% keytop %}}$a\frac{b}{c}${{% /keytop %}} があると計算結果を分数表現にしてくれる。
助かる。

### y<sup>x</sup> があるか

{{% keytop %}}$x^2${{% /keytop %}} や {{% keytop %}}$x^3${{% /keytop %}} はあるのに {{% keytop %}}$y^x${{% /keytop %}} がない関数電卓があるらしい。
これがないと複利計算が面倒だと思うのだが。

### x√y があるか

同じく {{% keytop %}}$\sqrt{\\;}${{% /keytop %}} や {{% keytop %}}$\sqrt[3]{\\;}${{% /keytop %}} はあるのに {{% keytop %}}$\sqrt[x]{y}${{% /keytop %}} がない関数電卓もあるらしい。
これも利息の計算で要るよなぁ。

『[関数電卓がすごい]』では「5年で売上を2倍にする」などという経営計画が降ってきたときに

{{< div-box >}}
\[
    \sqrt[5]{2}=1.148698\cdots
\]
{{< /div-box >}}

と計算して，年15%の売上拡大が必要だと評価してみせている。

まぁでも $\sqrt[x]{y}=y^{1/x}$ なので

{{< div-box >}}
\[
    \sqrt[5]{2} = 2^{\frac{1}{5}} = 1.148698\cdots
\]
{{< /div-box >}}

 {{% keytop %}}$y^x${{% /keytop %}} があれば代替は可能だから必須ではないか。

### log<sub>x</sub>y があるか

{{% keytop %}}$\log${{% /keytop %}} や {{% keytop %}}$\ln${{% /keytop %}} は大抵あるが {{% keytop %}}$\log_x{y}${{% /keytop %}} がない関数電卓は多いらしい。
たとえば $\log_2{256}=8$ みたいに使える。

ただし $\log_x{y}={\log{y}}/{\log{x}}$ で代替できるので，頻繁に使わないならなくても構わない？

『[関数電卓がすごい]』ではガチャ（ルートボックスのこと）でレアアイテムを引く確率を例示している。
たとえば排出率1%のアイテムを何回引けば50%の確率で引き当てられるかは以下の式で計算できる。

{{< div-box >}}
\[
    \log_{0.99}{0.5}=68.96756\cdots
\]
{{< /div-box >}}

なるほど。
こうやって使うのか。

...[ガチャはほどほどに]({{< ref "/remark/2023/04/the-psychology-of-video-games.md" >}} "『はじめて学ぶ ビデオゲームの心理学』は読んどけ！")（笑）

### 逆数計算があるか

{{% keytop %}}$1/x${{% /keytop %}} または {{% keytop %}}$x^{-1}${{% /keytop %}} があれば簡単に逆数が取れる。
『[関数電卓がすごい]』ではガソリンの燃費計算を例示しているが，逆数は結構使う機能。
たとえばウォーキングで $1\\,\mathrm{km}$ を $11$ 分で歩いたとすると

{{< div-box >}}
\[
    (11\div60)^{-1} = 5.454\cdots
\]
{{< /div-box >}}

と計算すれば時速約 $5.5\\,\mathrm{km}/\mathrm{h}$ で歩いたことになる。
むー，ちょっと早歩きだったか。
まぁ，逆数なので $60\div11$ で計算してもいいんだけどね。
頭の中で式変形しなくても電卓上で試すってのがいいのだろう。

### 双曲線関数があるか

{{% keytop %}}$\sinh${{% /keytop %}} {{% keytop %}}$\cosh${{% /keytop %}} {{% keytop %}}$\tanh${{% /keytop %}} または {{% keytop %}}$\mathrm{hyp}${{% /keytop %}} と {{% keytop %}}$\sin${{% /keytop %}} {{% keytop %}}$\cos${{% /keytop %}} {{% keytop %}}$\tan${{% /keytop %}} を組み合わせて双曲線関数が使えるのだが，この関数がない関数電卓もあるようだ。

『[関数電卓がすごい]』では評価値に下駄を履かせるのに双曲線関数を使ってて面白かった。

### nPr , nCr があるか

順列・組み合わせですな。
これがない関数電卓もあるらしい。

順列・組み合わせは確率計算では必須要素だが，それ以前に確率計算は前提となる命題の立て方や解釈で間違いが起きやすかったりする。

## スマホ用関数電卓アプリを紹介してみる

では，関数電卓アプリの物色をはじめよう。
Android アプリのみ。
iPhone は知らん。

### [RealCalc Plus]

{{< fig-img src="./realcalc-plus.png" title="RealCalc Plus" link="./realcalc-plus.png" >}}

広告付きの無料版もある。
{{% keytop %}}$\log_x{y}${{% /keytop %}} がない以外は欲しい機能がほぼ揃ってるので長く使っていたが，残念なことに計算履歴の再利用では式の変更ができないため，試行錯誤する際の使い勝手が悪いことに気がついた。

### [ClassWiz Calc App]

CASIO の [CLASSWIZ] シリーズのエミュレーションができる関数電卓アプリ。
電卓を無理やりアプリにしたみたいな感じで動きがぎこちないが [CLASSWIZ] を購入前に評価するアプリとして考えるのなら悪くない。
継続して使うならサブスクリプション契約をする必要がある？

ちょっと試してみたがアプリとしては微妙。
これを使うくらいなら電卓を買いなはれってことでしょうな。

### [Panecal]

{{< fig-img src="./panecal.png" title="Panecal" link="./panecal.png" >}}

（広告が下品だったので塗りつぶしている。ご容赦）

広告のない[有料版](https://play.google.com/store/apps/details?id=jp.ne.kutu.PanecalPlus "関数電卓 Panecal Plus - Google Play")もある。
ヘルプが日本語。
マジ助かる。

{{% keytop %}}$y^x${{% /keytop %}} とか {{% keytop %}}$\log_x{y}${{% /keytop %}} とか双曲線関数とかがない。
あと分数表現もできないっぽい。
裏機能を使うのに {{% keytop %}}ALT{{% /keytop %}} を組み合わせるためか，キーの長押しが使えないのが微妙。

{{% keytop %}}$1${{% /keytop %}} {{% keytop %}}$\div${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} {{% keytop %}}$\times${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} と入力すると $0.999\cdots$ になる。
これは [Panecal] の仕様みたいで {{% keytop %}}$1${{% /keytop %}} {{% keytop %}}$\div${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} で確定させたあと {{% keytop %}}ALT{{% /keytop %}} {{% keytop %}}Ans{{% /keytop %}} {{% keytop %}}$\times${{% /keytop %}} {{% keytop %}}$3${{% /keytop %}} {{% keytop %}}$=${{% /keytop %}} とすれば $1$ になる。
{{% keytop %}}Ans{{% /keytop %}} が2種類あるわけやね。

[Panecal] 自体は単位変換機能を持ってないが，同じメーカーから出ている[単位変換](https://play.google.com/store/apps/details?id=jp.appsys.unit_converter "単位変換 - Google Play")アプリと連携できる。
マジ日本語助かる（大事なことなので2回言いました）。

あと一歩なんだよなー。
せっかく日本語なのに（3回目）。

### [HiPER Calc Pro]

{{< fig-img src="./hiper-calc-pro.png" title="HiPER Calc Pro" link="./hiper-calc-pro.png" >}}

無料版もあるが，あまりの出来のよさにソッコーで有料版を買った。
前節で挙げた評価基準は全て満たしている。
その上で行列や微積分も使えるし $\sum$ も $\prod$ も使えるし複素数も使える。
乱数生成もできるとか！

変数は $x$, $y$ 以外に自由に名前を付けて定義できるっぽい。
{{% keytop %}}STAT{{% /keytop %}} や {{% keytop %}}OTHER{{% /keytop %}} を見るとアホみたいに関数が定義されている。
しかも関数の自作もできるらしい。

更に更にグラフ描画機能もある。
たとえば $\sin(x) + \cos(y)$ を定義して

{{< fig-img src="./hiper-calc-pro-2.png" title="HiPER Calc Pro (2)" link="./hiper-calc-pro-2.png" >}}

{{% keytop %}}More...{{% /keytop %}} を押下すると

{{< fig-img src="./hiper-calc-pro-3.png" title="HiPER Calc Pro (3)" link="./hiper-calc-pro-3.png" >}}

てな感じに描画してくれる。
まぁ，グラフ描画なら表計算ソフトでもできるし，なんなら Gnuplot を使う手もあるんだけど，たとえば $x^2+3x+2=x+6$ みたいな方程式を解く際

{{< fig-img src="./hiper-calc-pro-4.png" title="HiPER Calc Pro (4)" link="./hiper-calc-pro-4.png" >}}

{{% keytop %}}More...{{% /keytop %}} 押下で以下のように視覚的に示してくれるのは嬉しい。

{{< fig-img src="./hiper-calc-pro-5.png" title="HiPER Calc Pro (5)" link="./hiper-calc-pro-5.png" >}}

ちうわけで，当面はこれをメインに使っていこうと思う。

## 量は交換できる

最初の方で「現代社会における「価値」の多くは，定性評価ではなく，定量評価である」と書いたが，量で評価することのメリットは異なる価値を比較・交換できることである。
そして「価値」評価の基準として多く使われる単位がお金と時間である。
だから『[関数電卓がすごい]』ではお金と時間を単位として複数の価値を評価・比較・交換することをしつこく行っている。

5章で

{{< fig-quote type="markdown" title="『関数電卓がすごい』5章" link="https://www.amazon.co.jp/dp/B0D6QQ3VXT?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
モチベーションとは要するに野望です。ここでいう野望とは身の丈を超えた大いなる望みのことをいいます。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『関数電卓がすごい』5章" link="https://www.amazon.co.jp/dp/B0D6QQ3VXT?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
モチベーションアップのコツは、できることでやりくりしようとしないことと、逆算です。
{{< /fig-quote >}}

と書かれていて面白かった。
何故異なる価値を評価・比較・交換するかというと「身の丈を超えた大いなる望み」を叶えるためにほかならない。
身の丈にあったスローライフが望みならモチベーションなど必要ないのだから。

今どきこういうことを言う人は周りにいなくなったなぁ，というのが『[関数電卓がすごい]』を読んだ最終的な感想だった。
著者の芝村裕吏さんは「[狂狷の徒]({{< ref "/remark/2017/12/hacker-ethic.md" >}} "エンジニアこそ「狂狷の徒」たれ")」って感じだろうか。
これは読む人を選ぶ本かもしれない。
0章で「啓蒙は厳禁」と書いてあったのも宜なるかな，と思ったり。

というわけで面白い本だったが特に薦めないことにしよう（笑）

[関数電卓がすごい]: https://www.amazon.co.jp/dp/B0D6QQ3VXT?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "関数電卓がすごい (ハヤカワ新書) | 芝村 裕吏 | 数学 | Kindleストア | Amazon"
[RealCalc Plus]: https://play.google.com/store/apps/details?id=uk.co.nickfines.RealCalcPlus "RealCalc Plus - Google Play"
[CLASSWIZ]: https://web.casio.jp/dentaku/sp/classwiz/ "関数電卓CLASSWIZ - 電卓 - CASIO"
[ClassWiz Calc App]: https://play.google.com/store/apps/details?id=jp.co.casio.fx.ClassWizCalcApp "ClassWiz Calc App - Google Play"
[Panecal]: https://play.google.com/store/apps/details?id=jp.ne.kutu.Panecal "関数電卓 Panecal - Google Play"
[HiPER Calc Pro]: https://play.google.com/store/search?q=hiper+calc+pro&c=apps "hiper calc pro - Google Play の Android アプリ"

## 参考

{{% review-paapi "B08XQGKL6G" %}} <!-- ぐだふわエブリデー -->
