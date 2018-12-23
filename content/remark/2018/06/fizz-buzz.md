+++
title = "AI 時代に真っ先に駆逐される職業は（コピペ）プログラマなんじゃないだろうか"
date = "2018-06-02T21:07:48+09:00"
update = "2018-06-03T16:46:16+09:00"
description = "私はこれをセンスの（優劣ではなく）差異と考える。プログラマに向いてないなら他の職業につけばいいのだ。"
image = "/images/attention/kitten.jpg"
tags        = [ "programming", "language" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

面白い記事を見かけた。

- [本の虫: 世の中にはプログラミングを理解できない人間が存在する](https://cpplover.blogspot.com/2018/05/blog-post_29.html)

この手の話はこの時期特有の風物詩みたいなものだ。
春の季語といってもいいかもしれない（笑）
似たような話で印象深かったのが[3年前に読んだ](http://www.baldanders.info/spiegel/log2/000868.shtml "プログラミングは「損得勘定」で考える — Baldanders.info")この記事。

- [ペアプログラミングして気がついた新人プログラマの成長を阻害する悪習 - Qiita](https://qiita.com/ukiuni@github/items/4a252f47a37e17d99c59)

この記事ではプログラミングができない人の問題を「習慣」に求めている。
また最初に挙げた記事では

{{< fig-quote title="世の中にはプログラミングを理解できない人間が存在する" link="https://cpplover.blogspot.com/2018/05/blog-post_29.html" >}}
<q>思うに、先天的に、あるいは幼少期の教育の結果、プログラミングに向かない学習方法に最適化されてしまった人間がいるのではないだろうか。</q>
{{< /fig-quote >}}

とし，主に教育または学習方法の問題と考えているようだ。

しかし，世の中には「文章またはコードが読み書きできない人」というのが存在するのだ[^h1]。
何故そう言い切れるかというと，私も（どっちかというと）そちら側の人間だからだ。

[^h1]: 人は生まれながらにして平等，ではない。包摂と排除は表裏一体であり「人は生まれながらにして平等」と考える人がいるのならそれは平等でない人から「機会」を奪う不遜な人間である。

## 数学の問題を解くのに公式の暗記から始める人はプログラマには向いてないので諦めた方がいい

私のようなタイプの人間が「職業プログラマ」としてかろうじて禄を食む生活が成り立っているのは「理解に時間をかけることを苦としない」性格だからに過ぎない（頻繁に間違うので人生ブレまくりだが`w`）。
その一方で，競技プログラミングとか「なんたらキャンプ」とか「なんとかジャム」みたいに，特に短期で解（理解ではない）を求められるものは本当に苦手だったりする。

学校教育の問題でも習慣の問題でもなく読解力の弱い人は一定数いる。
これについて端的に語られているのが『[AI vs. 教科書が読めない子どもたち]』であろう。
（[感想文]({{< ref "/remark/2018/02/artificial-intelligence-book.md" >}} "『AI vs. 教科書が読めない子どもたち』を読む")にも書いたが）この本に書かれている将来予測はともかく，読解力に対する調査結果は非常に興味深い。
以下に一部を挙げてみる。

{{% fig-quote title="AI vs. 教科書が読めない子どもたち" link="https://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/" %}}
- 中学校を卒業する段階で、約3割が（内容理解を伴わない）表層的な読解もできない
- 学力中位の高校でも、半数以上が内容理解を要する読解はできない
- 進学率100%の進学校でも、内容理解を要する読解問題の正答率は50%強程度である
- 読解能力値と進学できる高校の偏差値との相関は極めて高い
- 読解能力値は中学生の間は平均的に向上する
- 読解能力値は高校では向上していない
- 読解能力値と家庭の経済状況には負の相関がある
- 通塾の有無と読解能力値は無関係
- 読書の好き嫌い、科目の得意不得意、1日のスマートフォンの利用時間や学習時間などの自己申告結果と基礎的読解力には相関はない
{{% /fig-quote %}}

母国語の文章の読み書きがうまくできないというのは日常生活が辛い状況だが，プログラミング言語に関してはそうでもない。
プログラミング言語を知らなくても日常生活には困らないからだ。

私はこれをセンスの（優劣ではなく）差異と考える。
プログラマに向いてないなら他の職業に就けばいいのだ。

プログラミングってのは算法（algorithm）の塊でしかない。
それを理解し記述するのに必要なのは，昔も今も，美術（art）ではなく算術（arithmetic）である[^aa1]。
故に数学の問題を解くのに公式の暗記から始める人はプログラマには向いてないので諦めた方がいい。

[^aa1]: もちろん美術も算術も両方持ち合わせているならそれに越したことはないけど。私は美術センスが壊滅的だからなぁ。


## 【付録】 Fizz Buzz で遊ぶ

「[世の中にはプログラミングを理解できない人間が存在する](https://cpplover.blogspot.com/2018/05/blog-post_29.html)」に出てくる Fizz Buzz は以下のルールで遊ぶゲーム（暇つぶし）である。

1. プレイヤーは複数（独りでは「世界のナベアツ」になってしまう`W`）
1. プレイヤー間で順番を決め 1 から順番に数を数えていく。ただし
    - 3 の倍数の場合は “Fizz” と唱える
    - 5 の倍数の場合は “Buzz” と唱える
    - 3 と 5 の公倍数の場合は “Fizz Buzz” と唱える[^fb1]

[^fb1]: 言うまでもないが 3 と 5 は素数なので 3 と 5 の公倍数とは $3 \times 5 = 15$ の倍数を指す。しかし，この記事では知らないふりをする。

このルールのプログラミングが新人研修の演習問題とかで行われているらしい。

というわけで，まずは [Go 言語]で書いてみる。
今回は素朴にこんな感じ。

```go
package main

import (
    "fmt"
    "strconv"
)

func newCounter() func() int {
    ct := 0
    return func() int {
        ct++
        return ct
    }
}

func fizzBuzz(n int) string {
    if n%3 == 0 && n%5 == 0 {
        return "Fizz Buzz"
    } else if n%3 == 0 {
        return "Fizz"
    } else if n%5 == 0 {
        return "Buzz"
    }
    return strconv.Itoa(n)
}

func main() {
    max := 30
    c := newCounter()
    for n := c(); ; n = c() {
        if n < max {
            fmt.Print(fizzBuzz(n), ", ")
        } else {
            fmt.Println(fizzBuzz(n))
            break
        }
    }
}
```

数列を発生させるカウンタ関数[^ch1] と Fizz Buzz を唱える関数を使って for 文でぐるぐる回してるだけだ。
終了条件として30まででカウンタを止めるようにした。
これを実行するとこうなる。

[^ch1]: 最初は（いつもの如く） channel を使っていたのだが， C++ に変換することを考えて関数閉包（closure）に書き直した。ただしこのままではスレッド・セーフにならない。

```text
$ go run fizz-buzz.go
1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, Fizz, Buzz, 26, Fizz, 28, 29, Fizz Buzz
```

とりあえず正しく動いてるかな。

そして今度はこれを C++ で書き直してみる。
こんな感じかな。

```cpp
#include <iostream>
#include <string>

auto newCounter() {
    return [ct=0]() mutable {
        return ++ct;
    };
}

auto fizzBuzz(auto n) {
    if (n%3==0 && n%5==0) {
        return std::string("Fizz Buzz");
    } else if (n%3==0) {
        return std::string("Fizz");
    } else if (n%5==0) {
        return std::string("Buzz");
    }
    return std::to_string(n);
}

int main() {
    auto max = 30;
    auto c = newCounter();
    for (auto n = c(); ; n = c()) {
        if (n < max) {
            std::cout << fizzBuzz(n) << ", ";
        } else {
            std::cout << fizzBuzz(n) << std::endl;
            break;
        }
    }
    return 0;
}
```

完全に動くコード（の筈）なので[動作確認](https://code.sololearn.com/c0fNdu02zSRC)は各自でどうぞ[^gcc1]。
ロートル・エンジニアから見ると auto 変数めっさ便利（笑）
でも if 文や for 文の式を括弧で括るのかったるいし，型を前置するの違和感が半端ない。

[^gcc1]: 手元では C++ コードの動作確認を [MinGW-w64 Windows版]({{< ref "/remark/2018/03/mingw-w64.md" >}} "MinGW-w64 を導入する")の GCC 7.3 でコンパイルしたバイナリで行った。

今ふと思ったが，私の[母国語]({{< ref "/remark/2015/programming-language.md" >}})って既に [Go 言語]になってるな。
Fizz Buzz のコードを脳内で記述するときに [Go 言語]で考えてたもんな。
この歳になって母国語が変わるとか結構すごいことなんじゃないだろうか。


## ブックマーク

- [「だれもがプログラミングを学ぶべき」ではない]({{< ref "/remark/2016/05/lets-programming.md" >}})
- [プログラミングで「計算論的思考」は身につかない]({{< ref "/remark/2016/09/programming.md" >}})
- [AI の読解力，人の読解力]({{< ref "/remark/2016/11/reading-comprehension.md" >}})

[AI vs. 教科書が読めない子どもたち]: https://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/ "Amazon.co.jp： ＡＩ　ｖｓ．　教科書が読めない子どもたち eBook: 新井 紀子: Kindleストア"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B079JLW5YN/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QDhrqqEtL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B079JLW5YN/baldandersinf-22/">プログラマの数学 第2版</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2018-01-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B015643CPE.09._SCTHUMBZZZ_.jpg"  alt="暗号技術入門 第3版　秘密の国のアリス"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMK4.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ガロア理論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1CM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1CM.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／フェルマーの最終定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1FO.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／乱択アルゴリズム"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B018VE46YW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B018VE46YW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／ベクトルの真実"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B073F45B97/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B073F45B97.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／積分を見つめて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMJ0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMJ0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／整数で遊ぼう"  /></a> </p>
<p class="description">タイトル通りプログラマ必読書。第2版では機械学習に関する章が付録に追加された。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-03-19">2018-03-19</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51XGP8AFX2L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621045938/baldandersinf-22/">いかにして問題をとくか</a></dt><dd>G. ポリア 柿内 賢信 </dd><dd>丸善 1975-04-01</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621085298/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621085298.09._SCTHUMBZZZ_.jpg"  alt="いかにして問題をとくか・実践活用編"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4061497863/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4061497863.09._SCTHUMBZZZ_.jpg"  alt="数学的思考法―説明力を鍛えるヒント  講談社現代新書"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/462108819X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/462108819X.09._SCTHUMBZZZ_.jpg"  alt="数学×思考=ざっくりと  いかにして問題をとくか"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797375698/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797375698.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4185086180/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4185086180.09._SCTHUMBZZZ_.jpg"  alt="授業研究に学ぶ高校新数学科の在り方"  /></a> </p>
<p class="description" >数学書。というか問いの立てかたやものの考え方についての指南書。のようなものかな。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-26">2014/09/26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
