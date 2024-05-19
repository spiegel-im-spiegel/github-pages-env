+++
title = "『効率的な Go』読書会 2回目"
date =  "2024-05-18T20:19:34+09:00"
description = "第2章だけ立ち読みするってのも手かも知れない。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "golang", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今日は読書会の日。

- [第59回横浜Go読書会（オンライン） - connpass](https://yokohama-go-reading.connpass.com/event/317024/)

というわけで，[前回]({{< ref "/remark/2024/04/efficient-go-1.md" >}})に引き続き『[効率的な Go]』の話を。
今回は第2章の真ん中辺りまで読んだので「第2章 効率的な Go 入門」の感想を書いてみる。

これから [Go] について調べたい。
あるいは「[Go] ってどがぁな言語？」って方がいたらこの第2章だけ立ち読みするってのも手かも知れない。
[Go] を礼賛するフレーズが多いが，苦手とする部分もきっちり言及してるし参考にはなると思う。
例示されているコードは少なめ且つ簡単なので「[Go] なにもわからない」って方も取り敢えず大丈夫。

以下，個人的に面白いと思ったフレーズを抜き出しててみよう。

{{< fig-quote type="markdown" title="『効率的なGo』p.41" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
指針となった原則は、安全性と繰り返しの少なさを引き換えにせず、よりシンプルなコードを可能にする言語を作ることでした。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.41" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
GoogleはいまだにGoを管理する唯一の企業であり、Goに対する最後の決定的な支配権を持っています。たとえ誰もが修正、使用、貢献できるとしても、単一のベンダーによって調整されたプロジェクトは、再ライセンスや特定の機能のブロックなど、身勝手で有害な決定を下す危険性があります
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.42" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
「gofmt のスタイルは誰の好みでもないが、gofmt はみんなの好みである（ Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. ）」
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.47" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
標準ライブラリの設計を通じて、依存関係を制御することに多大な努力が払われました。1つの機能のために大きなライブラリを引っ張ってくるよりも、少しのコードをコピーするほうが良い場合があります（システムビルドのテストでは、新しいコアの依存関係が発生すると文句を言われます）。依存関係の衛生管理は、コードの再利用に優先します。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.48" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
繰り返しになりますが、効率を考慮すると、依存関係と透明性における潜在的なミニマリズムは、非常に大きな価値をもたらします。未知なものが少ないということは、主要なボトルネックをすばやく検出し、もっとも重要な価値のある最適化にまず注力できることを意味します。私たちは、依存関係に最適化の余地があることに気づいても、それを回避する必要はないのです。そのかわり、私たちは通常、その修正を直接アップストリームに貢献することが歓迎されます。これは両者にとって有益なことです！
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.50" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
Goは、エラーを第一級市民の言語機能として扱うことで、独自の道を歩んでいます。信頼できるソフトウェアを書きたいと仮定して、エラー処理を明示的に、簡単に、そしてライブラリやインターフェイスにわたって統一的に行うのです。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.53" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
もうひとつよくある不満は、Goを書くと非常に「悲観的」になるということです。なぜなら、決して起こらないかもしれないエラーが、目に見える形で現れるからです。
{{< /fig-quote >}}

以下は読書会では未読だけど面白かったので。

{{< fig-quote title="『効率的なGo』p.71" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
<p>1960 年、米国と欧州のプログラミング言語専門家が手を組み、Algol60 が誕生しました。1970 年、Algol のツリーはCとPascal の2 つのブランチに分かれました。約 40 年後、この2 つの枝はGoで再び合流しました。</p>
<p class="right" lang="en">Robert Griesemer, <q>The Evolution of Go</q> （https://oreil.ly/a4V1e）</p>
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.71" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
効率という点では、Go にはアキレス腱があります。「5.5 Goのメモリ管理」で学ぶように、メモリ使用量はときに制御しにくいことがあります。プログラム中の割り当ては（特に新しいユーザーにとっては）驚くようなもので、ガベージコレクションの自動メモリ解放処理にはオーバーヘッドや偶発的な動作があります。特にデータ指向（data-intensive）のアプリケーションでは、メモリやCPU の効率を確保するために、メモリ容量が厳しく制限されているマシン（IoT など）と同様に努力が必要です。
{{< /fig-quote >}}

と，まぁ，こんな感じかな。
多少は雰囲気が伝わるだろうか。

そういえば [Go] 処理系のバリエーションとして [TinyGo] は知ってたけど [GoBot] は知らんかった。
いや Mastodon の TL で見かけたかな。
いわゆる IoT で [Go] 処理系を考えているならこの辺も視野に入れたほうがいいだろう。

今回はここまで。
次回も感想を書くかどうかは分からない。

## ブックマーク

- [Golang(Gobot+Firmata+Arduino)でアナログメーターを作ってみる](https://zenn.dev/_kazuya/articles/0045ef8057c0b5)

[Go]: https://go.dev/
[TinyGo]: https://tinygo.org/
[GoBot]: https://gobot.io/ "Gobot - Golang framework for robotics, drones, and the Internet of Things (IoT)"
[効率的な Go]: https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "効率的なGo ―データ指向によるGoアプリケーションの性能最適化 | Bartłomiej Płotka, 山口 能迪 |本 | 通販 | Amazon"

## 参考

{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4863544006" %}} <!-- 基礎から学ぶ TinyGo の組込み開発 -->
