+++
title = "簡単な算数をしませう：場合の数と確率"
date =  "2020-07-12T15:11:28+09:00"
description = "「場合の数」は小学6年生で習うらしい。"
image = "/images/attention/kitten.jpg"
tags = [ "math", "infection" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

最近の義務教育のカリキュラムを知らないのだが，「場合の数」は小学6年生で習うらしい。
だから簡単だよね！

## 問題

架空の《〇〇ウイルス》に関する問題。

{{< div-box type="markdown" >}}
1. 《〇〇ウイルス》に感染している人は《全人口の $0.001=1\\,\\tcperthousand$》である
2. 《〇〇ウイルス》に感染しているか否かを調べる《判定キット》がある
    - 《判定キット》は感染しているか否かを《$0.999=99.9\\,\\%$ の確率》で正しく判定する
    - 感染している という判定のことを《陽性》と呼ぶ
    - 感染していない という判定のことを《陰性》と呼ぶ
{{< /div-box >}}

この条件の下，**無作為に選んだ** 人を《判定キット》で調べたら《陽性》だった。この人が《〇〇ウイルス》に実際に感染している確率を求めよ。

## 解答

まずは場合の数を求める。
被験者について以下の4つの場合が考えられるだろう。

|               場合 | 全人口に対する割合              |
| ------------------:| ------------------------------- |
|   感染者で《陽性》 | $0.001 \times 0.999 = 0.000999$ |
|   感染者で《陰性》 | $0.001 \times 0.001 = 0.000001$ |
| 非感染者で《陽性》 | $0.999 \times 0.001 = 0.000999$ |
| 非感染者で《陰性》 | $0.999 \times 0.999 = 0.998001$ |

この場合分けを使って《陽性》反応が出る人のうち「実際に感染している割合」を求めると以下のようになる。

{{< fig-math >}}
\[
    \frac{感染者で《陽性》}{感染者で《陽性》+非感染者で《陽性》} = \frac{0.000999}{0.000999+0.000999} = 0.500 = 50.0\,\%
\]
{{< /fig-math >}}

{{< div-gen type="markdown" class="center" >}}
**やったね！ 半々の確率で当たりだよ☆**
{{< /div-gen >}}

ポイントは母集団中の感染者の割合と《判定キット》の正答率の2つの変数を合わせて考える必要がある，ということだろう。

実際には，その辺の人を無作為に捕まえて検査を行うなどという馬鹿なことはしない（筈）。
「疑わしい人」を対象に検査を実施するのだから，母集団の構成は，当然ながら，大幅に感染者に偏っている筈である。

## 余談だが...

現在の日本の人口は130M人弱。
このうち COVID-2019 発症者の累計は（[WHO の situation report](https://www.who.int/emergencies/diseases/novel-coronavirus-2019/situation-reports) によると） 2020-07-11 時点で21K人ほどである。
これを単純に割り算すると，国内人口の $0.16\\,\\tcperthousand$ ほどが発症した（している）ことになる。

感染しても無症状のままというケースが多いという話も聞くので，実際の感染者の割合はもっと上がるだろうけど，10倍を上回るということはないだろう（希望的観測）。

軽くググってみると SARS-CoV-2 向けの PCR (Polymerase Chain Reaction) 検査の感度は $97\\,\\%$ ほどらしい。
ただ技術的な巧拙によって $70\\,\\%$ 程度まで下がるとかいう話も聞くのだが，実質はどうなんだろう。

なお「感度」は「感染者を検査して陽性になる確率」で「特異度」は「非感染者を検査して陰性になる確率」なんだそうだ。
問題の《判定キット》では感度と特異度を同じ値とみなしているが，実際には異なるらしい。

## ブックマーク

- [「《命》に関わる確率」を疑似乱数を使って解いてみる]({{< ref "/remark/2019/04/mathgirl-note257.md" >}})

## 参考図書

{{% review-paapi "B01MSJMKMW" %}} <!-- 数学ガールの秘密ノート／やさしい統計 -->
{{% review-paapi "B079JLW5YN" %}} <!-- プログラマの数学 第2版 -->
