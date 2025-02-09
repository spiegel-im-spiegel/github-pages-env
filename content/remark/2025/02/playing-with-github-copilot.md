+++
title = "GitHub Copilot で遊ぶ"
date =  "2025-02-08T18:04:57+09:00"
description = " 生成 AI の時代に求められる人材とはプロンプトを駆使する技能を持つ人ではなく，生成 AI の提案に NO と言える技術力と見識を持つ「狂狷の徒」ではないだろうか。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "github", "vscode", "golang", "programming", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GitHub Copilot] の無料版（制限あり）が開放されて VS Code 上で使えるようになった。

- [GitHub Copilot - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=GitHub.copilot)

といってもコメントの補完とかコミットメッセージくらいにしか使ってないけど。
英語不得手なので，これだけでもめっちゃ助かっている。

ちょっと思いついて [`github.com/goark/koyomi`]`/value` パッケージに曜日（Weekday）型を追加しようと思うのだが，試しに [GitHub Copilot] に手伝ってもらうことにした。

## GitHub Copilot で遊ぶ

まず {{% keytop %}}Ctrl{{% /keytop %}}+{{% keytop %}}I{{% /keytop %}} でプロンプトを表示し「[`time`]`.Weekday` と同等な機能で日本語の曜日名を出力する型」で問い合わせてみる。

{{< fig-img src="./github-copilot-prompt.png" title="time.Weekday と同等な機能で日本語の曜日名を出力する型" link="./github-copilot-prompt.png" width="1039" >}}

おー。
一通り提示してくれるんだな。
既に作っている [`koyomi`][`github.com/goark/koyomi`]`/value.DateJp` 型と関連するメソッドも追加されている。

これはこれで（エラーもバグもないし）使えなくもないのだが，いくつか調整したい。

1. `WeekdayJp` 型の基底型は `int` ではなく [`time`]`.Weekday` としたい
   - 機能上のメリットがあるわけではないが「文芸」的に `WeekdayJp` 型と [`time`]`.Weekday` 型が「関連」することを明示したい（[Go] に「継承」はない）
2. 曜日名の出力メソッドは `String()`, `ShortString()` ではなく `StringJp()`, `ShortStringJp()` としたい
   - `String()` メソッドは [`time`]`.Weekday` と同じ機能にする
3. `WeekdayJp` 型の値が `Sunday` 〜 `Saturday` 以外なら，曜日名の出力は [`time`]`.Weekday.String()` メソッドと同じにしたい

というわけで，最終的には以下のようなコードにした。

```go
// WeekdayJp is a type that represents the days of the week in the Japanese context.
// It is based on the time.Weekday type from the standard library.
type WeekdayJp time.Weekday

const (
    Sunday    WeekdayJp = WeekdayJp(time.Sunday) + iota // 日曜日
    Monday                                              // 月曜日
    Tuesday                                             // 火曜日
    Wednesday                                           // 水曜日
    Thursday                                            // 木曜日
    Friday                                              // 金曜日
    Saturday                                            // 土曜日
)

var weekdayNames = [7]string{"日曜日", "月曜日", "火曜日", "水曜日", "木曜日", "金曜日", "土曜日"}
var weekdayShortNames = [7]string{"日", "月", "火", "水", "木", "金", "土"}

// String returns the English name of the Japanese weekday (WeekdayJp)
// by converting it to the standard time.Weekday type and calling its String method.
func (w WeekdayJp) String() string {
    return time.Weekday(w).String()
}

// StringJp returns the Japanese name of the WeekdayJp if it is between Sunday and Saturday.
// If the WeekdayJp is out of this range, it returns the standard time.Weekday string representation.
func (w WeekdayJp) StringJp() string {
    if w < Sunday || w > Saturday {
        return time.Weekday(w).String()
    }
    return weekdayNames[w]
}

// ShortStringJp returns the short Japanese name of the WeekdayJp.
// If the WeekdayJp is not within the valid range (Sunday to Saturday),
// it returns the default string representation of the time.Weekday.
func (w WeekdayJp) ShortStringJp() string {
    if w < Sunday || w > Saturday {
        return time.Weekday(w).String()
    }
    return weekdayShortNames[w]
}
```

たとえば `String()` メソッドは

{{< fig-img src="./autocomplete-code.png" title="コード補完" link="./autocomplete-code.png" width="530" >}}

という感じにコードを提案してくる。
この状態で {{% keytop %}}Tab{{% /keytop %}} キー押下で確定する。
テストコードも関数名から推測して妥当なコードを提案してくる。
変数・定数もクラス・メソッド名も名前が大事ってことですね（笑）

コメントの補完も同様。

{{< fig-img src="./autocomplete-comment.png" title="コメント補完" link="./autocomplete-comment.png" width="530" >}}

ドキュメント生成もやってくれる。

{{< fig-img src="./generate-docs.png" title="ドキュメント生成" link="./generate-docs.png" width="1038" >}}

ATOM や VS Code の使い始めの頃はスニペットベースのコード補完機能に感動したものだが， [Copilot][GitHub Copilot] によって柔軟なコード補完や生成ができるのマジ助かる。
今回はコードレビュー機能等は紹介しないが，ちょろんと試した感じではなかなかいい感じである。
人間にレビューを投げる前に [Copilot][GitHub Copilot] 相手にレビューを行うのはアリかもしれない。

## 「自律」機械は（今のところ）存在しない

教科書に載っているサンプルコード程度ならともかく，プログラムコードに **唯一の正解** はない。
今回のような小さなコードですらそうなのだ。
[Copilot][GitHub Copilot] を含む生成 AI による提案を受け入れるか否かについては「人」による判断が不可欠だし，判断を行うためには，扱う対象に関する知識・技能・技術が要求される。
「自律」的な判断を行うのは今なお，機械ではなく，人の側なのである。

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1" >}}
<q>自立とは、仮想代理人ソフトウェアであるところのエージェントが自ら動き、誰の力も借りずに意思決定できることを言う。
[...]
一方、自律というのは哲学的な意味であり、自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう。</q>
{{< /fig-quote >}}

{{< fig-quote  title="そろそろ、人工知能の真実を話そう" link="https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&LINKCODE=OGI&TH=1&PSC=1" >}}
<q>今、世の中で懸念されているのは、自立ではなく自律の方だが、学習能力を与えられ、自らのプログラムを改善できるようになっても、機械が自律することは考えられない。
なぜなら、機械は結局、人間に与えられた理論やルールにのっとって行動することになるからである。</q>
{{< /fig-quote >}}

[GitHub Copilot] をはじめとする最近流行りの生成 AI の機能とは「翻案の大量生産」であり，その膨大な翻案空間から何をどうやって拾い上げるかについては，アルゴリズムの設計を含め（今のところ）人の側に委ねられている[^c1]。
そういう意味で [GitHub Copilot] が文字通りコパイロット（ナビゲーション）の立ち位置に徹しているのは上手い割り切り方だと思う。

[^c1]: 生成 AI の背後にある翻案元のコンテンツやアルゴリズムと知財権との[関係]({{< ref "/remark/2023/06/is-generative-ai-copyright-safe.md" >}} "Generative AI は Copyright-Safe か？")についてはここでは知らないふりをする。

生成 AI の時代に求められる人材とはプロンプトを駆使（して翻案空間を放浪）する技能を持つ人ではなく（積極的に利用しつつ）生成 AI の提案に NO と言える見識と技術を持つ「[狂狷の徒]({{< ref "/remark/2017/12/hacker-ethic.md" >}} "エンジニアこそ「狂狷の徒」たれ")」ではないだろうか。

## ブックマーク

- [GitHub、VS Code上で「Copilot Free」開始　無料で月2000回コード補完 - Impress Watch](https://www.watch.impress.co.jp/docs/news/1649089.html)
- [OpenAIの最新推論モデル「o3-mini」が「GitHub Copilot」などで利用可能に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1659601.html)
- [無償ユーザーも対象 ～「Gemini 2.0 Flash」が「GitHub Copilot」で利用できるように - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1660865.html)
- [「Visual Studio Code 1.97」が公開 ～新AI支援機能「Copilot NES」をプレビュー - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1661141.html)

- [GitHub Copilot は貢献者から貢献を奪うか？]({{< ref "/remark/2022/07/code-launderring.md" >}})
- [AI は「創作者様」を引きずり下ろすか — 『人権と文化コモンズ』を流し読む]({{< ref "/remark/2022/10/cultural-commons.md" >}})

[GitHub Copilot]: http://copilot.github.com/ "GitHub Copilot · Your AI pair programmer"
[Go]: https://go.dev/
[`github.com/goark/koyomi`]: https://github.com/goark/koyomi "GitHub - goark/koyomi: 日本のこよみ"
[`time`]: https://pkg.go.dev/time "time package - time - Go Packages"

## 参考図書

{{% review-paapi "B071FHBGW8" %}} <!-- 人工知能の真実を話そう -->
{{% review-paapi "B0DTT1L1KL" %}} <!-- 「日経サイエンス」2025年3月号 -->
{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->
{{% review-paapi "B01J1I8PRQ" %}} <!-- 社会は情報化の夢を見る -->

## 作業中の BGV (メン限配信以外)

ReGLOSS の五色のバラバラな声がパズルのようにカチッと填まると本当にかっこいいよね。

- [ReGLOSS 'フィーリングラデーション' OFFICIAL MV - YouTube](https://www.youtube.com/watch?v=u4uDiV3u-do)
- [晴る / ヨルシカ  covered by ReGLOSS 【歌ってみた / hololive DEV_IS】 - YouTube](https://www.youtube.com/watch?v=Mb2Pdf7P-a4)
- [ムーンライト / 星街すいせい covered by ReGLOSS 【歌ってみた / hololive DEV_IS】 - YouTube](https://www.youtube.com/watch?v=m6T71r7hRMI)
- [【歌ってみた】STAR TRAIN/Perfume【ReGLOSS】 - YouTube](https://www.youtube.com/watch?v=ed-0AKPWaf8)
