+++
title = "GitHub Copilot は貢献者から貢献を奪うか？"
date =  "2022-07-02T16:42:46+09:00"
description = "この件を拗らせるようならプロアカウント契約を解除することも考えないといけないだろう。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "intellectual-property", "github", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この前から話題になっている [GitHub Copilot] だが，ついに SFC (Software Freedom Conservancy) が GitHub と敵対する？ 声明を出したようだ。

- [Give Up GitHub: The Time Has Come! - Conservancy Blog - Software Freedom Conservancy](https://sfconservancy.org/blog/2022/jun/30/give-up-github-launch/)
- [「GitHubの利用を中止しよう」 SFCが提言、AI開発ツールに疑念 | TECH+](https://news.mynavi.jp/techplus/article/20220701-2385316/)

引き合いに出された SourceForge に佐渡秀治さんが反応して

{{< fig-gen class="nobox" >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">唐突に呼ばれた気がしたのでSFCの記事を読んでみたが、今は不使用のキャンペーンだけだけど、GitHubがこのままの姿勢ならSFCはいずれ訴訟を仕掛けることも考えそうというぐらいの勢いだな。Amazon CodeWhispererは適切な帰属とライセンス情報を提供するので問題ないというスタンスか。まあ、そこだよね <a href="https://t.co/xmphKprdoS">https://t.co/xmphKprdoS</a></p>&mdash; Shuji Sado (佐渡 秀治) (@shujisado) <a href="https://twitter.com/shujisado/status/1542733986417770496?ref_src=twsrc%5Etfw">July 1, 2022</a></blockquote>
{{< /fig-gen >}}

と tweet しておられるのが面白かった。

私が今回の一連を見て最初に連想したのは，ゼロ年代の Google Book Search だったのだが

- [Google Books の Library Book Scan すら Fair Use と言われたのに...]({{< ref "/remark/2015/google-books-library-project.md" >}})

これとはちょっと違うなとは思っていた。
うまく言語化できないが，つらつらと書いてみる。

Pekka Himanen 等による古典 “The Hacker Ethic” (邦題『[リナックスの革命](https://www.amazon.co.jp/dp/4309242456?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "リナックスの革命 ― ハッカー倫理とネット社会の精神 | ペッカ ヒマネン, リーナス トーバルズ, マニュエル カステル, 安原 和見, 山形 浩生 |本 | 通販 | Amazon")』) を引くまでもなく FOSS でコードを書く人にとって最も価値があるのは「自由なコード」であり，その FOSS プロジェクトにコミットする人々にとって一番の報酬はコード（＝自由）に「貢献」したという事実そのものである。

にも関わらず [GitHub Copilot] はコードに付随する貢献を洗い落としてしまう。
貢献を洗い落とされたコードは「[GitHub Copilot] のもの」としてしれっと使われる。
上の tweet の retweet 元にもあるように，これは2010年代で散々言われた「あなたが客でないとすれば、商品だ」という搾取の構造と言える。

言うなれば「コード{{% ruby "ロンダリング" %}}洗浄{{% /ruby %}}」といったところだろうか。
上手いフレーズを思いつかなくてゴメン {{< emoji "ペコン" >}}

そう考えると

{{< fig-quote type="markdown" title="Give Up GitHub: The Time Has Come!" link="https://sfconservancy.org/blog/2022/jun/30/give-up-github-launch/" lang="en" >}}
One recent preliminary finding was that [AI-assisted software development tools can be constructed in a way that by-default respects FOSS licenses](https://lists.copyleft.org/pipermail/ai-assist/2022-June/000015.html). We will continue to support the committee as they explore that idea further, and, with their help, we are actively monitoring this novel area of research. While Microsoft's GitHub was the first mover in this area, by way of comparison, early reports suggest that Amazon's new CodeWhisperer system ([also launched last week](https://www.theregister.com/2022/06/23/amazon_codewhisperer/)) seeks to provide proper attribution and licensing information for code suggestions.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="「GitHubの利用を中止しよう」 SFCが提言、AI開発ツールに疑念" link="https://news.mynavi.jp/techplus/article/20220701-2385316/" >}}
Software Freedom Conservancyの委員会はAI支援ソフトウェア開発ツールに関する調査を続けていると説明するとともに、FOSSライセンスを尊重する形でこうした機能を実現することは可能であることが最近の調査からわかったことなどにも言及している。

同委員会は今後もこの分野への取り組みを続けるとしているほか、Amazon CodeWhispererのようにコード提案時に適切な帰属とライセンス情報を表示しようとする取り組みがあることについても説明している。
{{< /fig-quote >}}

というのも単にライセンス処理の問題だけではないと思えてくる。

GitHub をはじめとするサービスは単なる「git リポジトリ・サービス」ではなく，コードを中心としたコミュニティを構成し貢献を可視化するプラットフォームとして機能している。
このことを軽視してはいけないと思う。

私は GitHub にはかなりお世話になってるが，この件を拗らせるようなら（乗り換えまではしなくても）プロアカウント契約を解除することも考えないといけないだろう。

## 【おまけ1】 [GitHub Copilot] は Fair Use か

[GitHub Copilot] が fair use かどうかについては微妙だと思う。
まぁ「公正な利用」の概念がない日本では一発アウトだろうが（笑）

先に紹介した [Google Book Search の件]({{< ref "/remark/2015/google-books-library-project.md" >}} "Google Books の Library Book Scan すら Fair Use と言われたのに...")や2021年まで争われていた [Google vs Oracle の訴訟]({{< ref "/remark/2021/04/google-vs-oracle-final.md" >}} "Google vs Oracle の訴訟の行方（最終章）")を見ても分かるように，あるサービスによる知財の取り扱いが fair use に基づいているか判断するのは難しい。
というか米国の場合，訴訟が始まってからがスタートラインなんだよな。

訴訟が始まって GitHub やその親玉である Microsoft が FOSS コードやその貢献についてどのように考えているか明らかになれば，もっと色々と語ることがあるかも知れない。

## 【おまけ2】目的によって利用を制限するのであればそれはもう FOSS ではない

ゼロ年代に 9.11 から始まる一連のテロや戦争のときもそうだったし，2022年のロシアがウクライナに仕掛けた侵略戦争勃発時にも見られたが，抗議目的でソフトウェアの（著作権的な意味での）利用を制限することは FOSS の目的や手段にマッチしない。
[「自由とはそういうことだ」]({{< ref "/remark/2019/03/ye-not-guilty.md" >}})。

たとえば，今回の [GitHub Copilot] に抗議・拒否する目的で既出の FOSS 製品の利用を制限したいのであれば「FOSS でないなにか」に変更しなければならない。
それは今まで積み上げてきたエコシステムの基盤を崩すことになる。

[GitHub Copilot]  には反対するけど今のエコシステムを壊したくないなら SFC が言うように「GitHub をあきらめる（give up GitHub）」しかない。
それとも（今まで積み上げてきたものを壊してでも）断固とした対決姿勢をとるか。
「リポジトリに Copilot 拒否オプションを付ければいいぢゃん」とかいった簡単な話ではないのである。

さて，どうなる？ どうする？

## ブックマーク

- [GitHub Copilotを正式リリース。すべての開発者が利用できるようになりました。 - GitHubブログ](https://github.blog/jp/2022-06-22-github-copilot-is-generally-available-to-all-developers/)
- [戦争に抗議する「プロテストウェア」が、オープンソースエコシステムの信頼を揺るがしている | WIRED.jp](https://wired.jp/article/open-source-sabotage-protestware/)

[GitHub Copilot]: http://copilot.github.com/ "GitHub Copilot · Your AI pair programmer"

## 参考図書

{{% review-paapi "4309242456" %}} <!-- リナックスの革命 Hacker Ethic -->
{{% review-paapi "B07FMV7HMD" %}} <!-- 夏をあきらめて（♪あきらめの夏♪） -->
