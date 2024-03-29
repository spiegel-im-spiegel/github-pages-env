+++
title = "人の靭性"
date =  "2021-10-23T19:43:26+09:00"
description = "本の紹介を兼ねてちょろんと書いておく。"
image = "/images/attention/kitten.jpg"
tags = [ "trust", "security", "risk", "management", "book" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter での脊髄反射 tweet に微妙に反応があったみたいなので，本の紹介を兼ねてちょろんと書いておく。

セキュリティ・システムにおける「{{% ruby "じんせい" %}}靭性{{% /ruby %}}」への言及を見かけたのは我らが Bruce Schneier 先生の『[セキュリティはなぜやぶられたのか](https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "セキュリティはなぜやぶられたのか | ブルース・シュナイアー, 井口 耕二 |本 | 通販 | Amazon")』第9章である。
ちょっと引用してみよう。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}セキュリティシステムには障害がよくおきるので、障害の性質を理解することが大事である。硬直し剛性が高いシステムは悲惨な障害がおき、靭性が高くねばり強いシステムはうまく壊れていく。靭性の高いシステムは変化することができる。障害が一部にとどまったり、少しずつ劣化したりする。あるいは環境の変化に対応できる。自動化されたシステムは、基本的に剛性が高く、ひとつのことしかできない。部分的な障害や想定外の攻撃、新しい攻撃方法、ずるい攻撃者などに対応できないのだ。画一的や秘密に強く依存するシステムも剛性が高くなりがちで、その秘密が漏れると一気に潰れたりする{{% /quote %}}
{{< /fig-quote >}}

このように「システムの壊れ方」について 剛性⇔靭性 という対比で説明している。
具体例については書籍を見ていただきたいが，この説明の後，第10章「セキュリティの中心は人である」に続くわけだ。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}システムを機能させるためには一部の人を信頼しなければならない。そのような人々&mdash;信頼を託される人々&mdash;は、セキュリティシステムの一部だといえる。重要な構成要素、いや要となる要素だといえるだろう。システムで最も靭性が高い部分であり、臨機応変の対応や即断即決ができる部分であり、攻撃者の存在を感じとる能力が最も高い部分だからだ。しかし一方、セキュリティシステムの構成要素として考えたとき、人間な両刃の剣である。居眠りもすれば気がそれることもある。だまされることもある。敵側に寝返ることさえある。すぐれたセキュリティシステムとするためには信頼を託す人の長所を活用するとともに、信用が乱用されない防止策を講じなければならない{{% /quote %}}
{{< /fig-quote >}}

「人はセキュリティの最弱点」とはよく言われるが，システムの 剛性⇔靭性 の対比で考えた場合は，人はセキュリティの最強点にもなり得る。
ある意味で最後の切り札（ace in the hole）と言ってもいいかも知れない。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}最近は、技術力でセキュリティ問題を解決できると信じ、技術によって人を減らそうという傾向がある [...] しかし人を技術で置き換えると、人が持つ創造性や創意工夫、適応力が利用できなくなる。人間がすることに制限を加えすぎると、人は無気力となり、同じ状況が生まれる。技術側では、人とその行動を非現実的なほどに理想化してしまう傾向がある。人は同じことしかしない、するはずのことを必ずして、するはずのないことはしないと仮定してしまうのだ。しかし実際には、そんな一定の行動を人がするはずがない。同時に、そんな融通が利かないこともない。技術的な対策は剛性が高く、障害がおきると大変なことになることが多い{{% /quote %}}
{{< /fig-quote >}}

どうだ，耳が痛かろう（笑） これが2003年に（原書が）出版された本の内容だよ。

Twitter でも書いたが，人の靭性を軽視すると，いったん破られたときに止める手段がなくなり class break を引き起こしやすい。
かといって人に依存しすぎて「運用でカバー」でも人が疲弊するだけだけど。
お互いが有機的に連携できる落とし所を探すのが「改善」というやつである。

## ブックマーク

- [NIST SP 800-207: “Zero Trust Architecture”]({{< ref "/remark/2020/09/nist-sp-800-207-zero-trust-architecture.md" >}})

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "B07ND6QTN4" %}} <!-- OODA LOOP -->
