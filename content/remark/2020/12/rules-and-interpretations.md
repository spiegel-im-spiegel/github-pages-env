+++
title = "ルールを解釈で捻じ曲げる"
date =  "2020-12-29T16:06:27+09:00"
description = "“People don't want to be educated, they want to be entertained”"
image = "/images/attention/kitten.jpg"
tags = [ "math", "management", "code" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「数学ガール」シリーズでおなじみの結城浩さんが [Twitter でゼロの偶奇を問うアンケート](https://twitter.com/hyuki/status/1343313840889593859)をされていて，その結果について

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">「0は偶数です」を正しいとした方は58%で、驚くほどの低正解率でした。<br><br>「0は偶数でも奇数でもありません」を正しいとした方が41%もいらっしゃいました。この回答をした方の理由をお聞きしたいです。<a href="https://t.co/reMDPiPM9f">https://t.co/reMDPiPM9f</a></p>&mdash; 結城浩 (@hyuki) <a href="https://twitter.com/hyuki/status/1343677705712123904?ref_src=twsrc%5Etfw">December 28, 2020</a></blockquote>
{{< /fig-gen >}}

と感想を tweet しておられた。

実はこのアンケートには元ネタがあるそうで，それがこれ。

- [The Parity Of Zero](https://www.solipsys.co.uk/new/TheParityOfZero.html)
- [「0は偶数ではない」と多くの人が信じているのは教育に問題があるという指摘 - GIGAZINE](https://gigazine.net/news/20201228-0-even-education/)

わざわざこんな辺境の記事を読むような人には自明だろうが，念のためにいうと，偶数とは「2の整数倍数」あるいは「2で割り切れる整数」として **定義** されるものである。
プログラマ風に言うなら「2進数に展開した際に最下位ビットが0になる整数値」または「`n&1==0` が `true` となる `int n`」でもいいだろう（笑） いずれにしてもゼロは偶数と言える，議論の余地なく。

この話のポイントは，偶奇の色分けは「定義」であるということだ。
言い方を変えるなら「2で割り切れる整数を偶数としましょう」という「ルール」である。
つまり「0は偶数でも奇数でもありません」と考えた人は，思い込みの解釈で偶奇のルールを捻じ曲げてしまったわけだ。

これで思い出したのが数年前に話題になった「[掛け算は順序が大事](https://twitter.com/genkuroki/status/515350512305049600)」という話である。
もちろん掛け算に順序に関するルールはない。
その上で数の掛け算には「交換法則」が成り立つという点が算数の算数たる所以なのだが，そういうのを全部チャイして「順序が大事」と言っちゃってるわけだ。

当時はこの話を聞いて「日本の学校教育 `＼(^o^)／ｵﾜﾀ`」と[思った](https://baldanders.info/blog/000744/ "日本の「算数」は壊れてる？ — 旧メイン・ブログ | Baldanders.info")ものだが，似たような話は世界中どこにでも転がっているということなのかもしれない。
これに関して，最初に挙げた記事の

{{< fig-quote type="markdown" title="The Parity Of Zero" link="https://www.solipsys.co.uk/new/TheParityOfZero.html" lang="en" >}}
{{% quote %}}It doesn't necessarily need to be made "relevant." Yes, for some people you can create the motivation that way, but for others, they will engage purely for the pleasure of finding things out, and the satisfaction in being able to see a reason behind things that were previously stated without justification{{% /quote %}}.
{{< /fig-quote >}}

という部分は結構重要なポイントに見える。
ものに喩えるのは私もよくやるが，喩えを行う場合はその「差異」を常に意識しないと，分かりやすい喩えの方に意識が引きずられてしまう。

まぁ，日本の，いや世界の数学教育の話はここまでにしておいて，注目したい点は「ルールを解釈で捻じ曲げる」人が一定数いるということだろう。

典型例は日本の「憲法九条」かな。
「解釈」で捻れまくっているよね。
あるいは知財やプライバシー・セキュリティ関連の法律など「ガイドライン」という名の解釈によって為政者に都合よく運用される風景が当たり前になっている。

Bruce Schneier 先生原著の『[セキュリティはなぜやぶられたのか](https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』では「抑止[^deter1]」における教育の必要性を説く。

[^deter1]: 『[セキュリティはなぜやぶられたのか](https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』では「抑止」を「将来の攻撃をセキュリティシステムが防止するやり方」と定義している。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}抑止が効果を持つためには、「教育」が必要だ。社会が犯罪から守られているのは、基本的に、攻撃に対する直接的な防御があるからではない。人々が法律を守るからだ。ほとんどの人は倫理や道徳を重んじる。倫理は人が生まれながらに持つ性質で、これがなかったら人は文明化できなかっただろう。道徳は、何が道徳的で善良な市民とはどういうものかという教育によって身につけるものだ{{% /quote %}}
{{< /fig-quote >}}

倫理や道徳はともかく，ルールを作ってそれを守ってもらいたいなら，教える側にしても教わる側にしても，それをどのように解釈するか（あるいはされるか）については常に注意を払うべきかもしれない。
なにせ偶数の定義すら捻じ曲げて解釈してしまうのが人間なのだから（笑）

あちこちの職場を渡り歩いて，その度に「セキュリテイ研修」を受けるが，ぶっちゃけルールを押し付けるだけのところが多く「何故」がないんだよね（まぁ「傭兵」相手に時間をかける気はないってことなんだろうけど）。
特にセキュリティ管理はルールの根拠をきちんと示さなければ守ろうとは思わないし，しぶしぶ守るとしても「解釈で捻じ曲げる」余地を与えてしまう。

{{< fig-quote type="markdown" title="The Parity Of Zero" link="https://www.solipsys.co.uk/new/TheParityOfZero.html" lang="en" >}}
{{% quote %}}People don't want to be educated, they want to be entertained{{% /quote %}}.
{{< /fig-quote >}}

私は「**守られないルールは，ルール自体に問題がある**」と考えるが「ルールを解釈で捻じ曲げる」事態が常態化するなら（教育を含めて）マネジメント全体を見直す最初のアラームとなるだろう。

## 参考図書

{{% review-paapi "B00NAQA33A" %}} <!-- 数学ガールの誕生 -->
{{% review-paapi "4621045938" %}} <!-- いかにして問題をとくか -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
