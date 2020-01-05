+++
title = "AI 時代に真っ先に駆逐される職業は（コピペ）プログラマなんじゃないだろうか"
date = "2018-06-02T21:07:48+09:00"
description = "私はこれをセンスの（優劣ではなく）差異と考える。プログラマに向いてないなら他の職業につけばいいのだ。"
image = "/images/attention/kitten.jpg"
tags = [ "programming", "language" ]

[scripts]
  mathjax = true
  mermaidjs = false
+++

面白い記事を見かけた。

- [本の虫: 世の中にはプログラミングを理解できない人間が存在する](https://cpplover.blogspot.com/2018/05/blog-post_29.html)

この手の話はこの時期特有の風物詩みたいなものだ。
春の季語といってもいいかもしれない（笑）
似たような話で印象深かったのが[3年前に読んだ](https://baldanders.info/blog/000868/ "プログラミングは「損得勘定」で考える — Baldanders.info")この記事。

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

{{% fig-quote type="markdown" title="AI vs. 教科書が読めない子どもたち" link="https://www.amazon.co.jp/exec/obidos/ASIN/B0791XCYQG/baldandersinf-22/" %}}
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

{{% review-paapi "B079JLW5YN" %}} <!-- プログラマの数学 第2版 -->
{{% review-paapi "4621045938" %}} <!-- いかにして問題をとくか -->
