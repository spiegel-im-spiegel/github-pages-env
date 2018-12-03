+++
date = "2015-12-12T21:16:03+09:00"
description = "強すぎる光は影を濃くするのみだ。あるいは強すぎる結界は更なる魔を引き寄せる。"
draft = false
tags = ["security", "cryptography", "risk", "politics"]
title = "強すぎる結界は更なる魔を引き寄せる"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

- [我匿す、ゆえに我あり - WirelessWire News](https://wirelesswire.jp/2015/12/48567/ "我匿す、ゆえに我あり - WirelessWire News（ワイヤレスワイヤーニュース）")

今回はこれをネタにつらつらと世迷いごとを書いてみる。
いつも以上に戯れ言なので，あんまり真面目に読まないように。

## 本当に秘密なことは誰とも共有してはならない

優れたアルゴリズムを使えば暗号システムそのものは必ずしも隠す必要はないが，鍵はどうしても秘匿しなくてはならない。
しかし秘密を共有すれば必ず漏えいリスクが発生する。
公開鍵暗号方式が「発明」されるまで，暗号技術が軍事または大企業の技術として専有されていたのは「秘密の共有」に莫大なコストがかかるからだ。

本当に秘密なことは誰とも共有してはならない。

## 秘密なんてないさ

[個人的に今年イチオシ](http://www.baldanders.info/spiegel/log2/000833.shtml)の映画である『イミテーション・ゲーム』に印象的なシーンがある。
ドイツの暗号機械「エニグマ」の暗号鍵を短時間で解読することに成功したアラン等のチームは，この事実を（政府や軍にも）知らせないことに決めた。
理由は（政治的な思惑もあるだろうが）事実を敵国に知られないためである。
敵国に知られれば相手は暗号システムを変えるだろうから。

秘密を隠すためには秘密が存在することをも隠す必要がある。

フランスの警察当局がテロ事件の犯人を捕まえることができたのは暗号を解読したからではない。

- [警察がテロリストの捨てた携帯からわかること : ギズモード・ジャパン](http://www.gizmodo.jp/2015/11/what-police-can-learn-from-a-terrorist.html)

警察・諜報組織は一般の人がアクセスできない情報にもアクセスできる。
たとえ暗号データそのものを解読できなくても，いったん対象を絞り込むことができれば，行動履歴をかき集め「名寄せ」することで追い詰めることができるのだ。
これを回避するのは並大抵ではない。

- [「オーウェルが描いた悪夢のような監視社会をさまざまな点で超えてしまっているこの世界」で私たちはいかにして生き残るか — Baldanders.info](http://www.baldanders.info/spiegel/log2/000768.shtml)

だから問題はどうやって「対象を絞り込む」か，つまりフィルタリングあるいはスクリーニングの問題である。

## 前門の虎，後門の狼少年

現在「パリのテロ」は西欧諸国にとって「錦の御旗」になっている。

名前が示す通り「テロ（terrorism）」は誰だって怖い。
そして，それによって大勢の人や近しい人が命を落とすのであれば，怒りを感じて当然である（恐怖は怒りを駆動する）。

テロの目的は相手に「恐怖」を刷り込み戦争状態を維持・活性化することにある。
でも「恐怖」を利用しているのはテロリストだけではない。
このことを 9.11 以後の10年あまりの間に私たちは嫌というほど見せつけられた筈である。
西欧諸国は「パリのテロ」を利用して 9.11 と同じ状況を作り出そうとしている。
目の前の恐怖に竦みあがっていると後ろから刺される。

{{< fig-quote title="FBIによる令状なしの個人情報開示要求、実態が明らかに–あらゆる通信記録が対象 - ZDNet Japan" link="http://japan.zdnet.com/article/35074261/" >}}
<q>ニューヨーク州南部地区連邦地方裁判所のVictor Marrero判事は判決文の中で、FBIの見解は「極端であり、範囲を過度に拡大している」と述べている。<br>
同判事はまた、Merrill氏に対する発言禁止命令の適用範囲が広すぎる点について、「米国憲法修正第1項（言論の自由）、および国民に対する政府の説明責任という両方の観点から、重大な問題をはらんでいる」との判断を示した。<br>
NSLの発言禁止命令を完全に解くことに成功したのは、Merrill氏が初めてだ。<br>
米国愛国者法（US Patriot Act）は、2001年9月11日の同時多発テロ事件から1カ月後に成立したときに、NSLの適用範囲を拡大した。</q>
{{< /fig-quote >}}

## 秘密なんてないYO

「[我匿す、ゆえに我あり]」では [Bruce Schneier](https://www.schneier.com/) さんの “[Why We Encrypt](https://www.schneier.com/essays/archives/2015/06/why_we_encrypt.html)” の一部を翻訳されている。
以下に少し紹介する。

{{< fig-quote title="Essays: Why We Encrypt - Schneier on Security" link="https://www.schneier.com/essays/archives/2015/06/why_we_encrypt.html" lang="en" >}}
<q>If we only use encryption when we’re working with important data, then encryption signals that data's importance. If only dissidents use encryption in a country, that country's authorities have an easy way of identifying them. But if everyone uses it all of the time, encryption ceases to be a signal. No one can distinguish simple chatting from deeply private conversation. The government can't tell the dissidents from the rest of the population. Every time you use encryption, you're protecting someone who needs to use it to stay alive.</q>
{{< /fig-quote >}}

{{< fig-quote title="我匿す、ゆえに我あり" link="https://wirelesswire.jp/2015/12/48567/" >}}
<q>重要なデータに携わる場合にだけ暗号を使うと、暗号化はそのデータが重要である合図になってしまう。ある国で反体制派しか暗号を使わないと、その国の権力者は反体制派を特定する容易な手段を手にすることになる。でも、誰もがいつでも暗号を使うなら、暗号化は合図でなくなる。単なるお喋りと深刻な密談の区別を誰もつけられなくなる。政府は反体制派とそれ以外の国民を区別できなくなる。あなたは暗号を使うたびに、生きのびるために暗号を使う必要のある誰かを守っているのだ。</q>
{{< /fig-quote >}}

「木の葉を隠すなら森のなか[^br]」ということだろう。
たぶんこれは通信・ネットワークに於いては正しい。

[^br]: 正しくは「賢い人は葉をどこへ隠す？ 森の中だ。森がない時は、自分で森を作る。一枚の枯れ葉を隠したいと願う者は、枯れ葉の林をこしらえあげるだろう」らしい。 Gilbert Keith Chesterton 著『[ブラウン神父の童心（The Innocence of Father Brown）](https://ja.wikipedia.org/wiki/%E3%83%96%E3%83%A9%E3%82%A6%E3%83%B3%E7%A5%9E%E7%88%B6%E3%81%AE%E7%AB%A5%E5%BF%83)』より「折れた剣（The Sign of the Broken Sword）」の中の一節。

現代社会においては善人と悪人，内部告発者と犯罪者，あるいは愛国者とテロリストを区別することはできない。
これらは私たちの日常空間に埋め込まれているからだ。

特に今回のパリのテロや先日の[米国カリフォルニア州のテロ](http://www.fsight.jp/articles/-/40737)のような「ローン・ウルフ型」と言われるテロは，活動主体の武装組織などと事前に negotiation する必要はなく，勝手に立案して勝手に実行して勝手に名乗りを上げることができてしまう[^a]。
こういったことを警察・諜報組織が（事後ならともかく）事前に察知することは，たとえ全ての通信履歴を解読できたとしても，難しいだろう。
何故なら誰とも秘密を共有する必要はない（または共有する範囲がごく限られる）し，秘密の存在自体を秘匿しやすいからだ[^b]。

[^a]: ちゃんと名乗りを上げないとただの犯罪と見分けがつかないからね。そのせいで米国の例では「テロ」の判断が遅れたわけだが（笑）
[^b]: さらに言えば，自爆テロなら「事後」を考慮する必要もない。

テロのような犯罪を事前に防ぐことが難しいのは暗号技術のせいではない。
彼ら為政者はただ，森に隠された木の葉を探すために「焼き払え！」と叫んでいるだけである。

## 強すぎる結界は更なる魔を引き寄せる

しかし，実際問題として暗号技術が正しく機能しているのかというと甚だ疑問と言わざるを得ない。

「[我匿す、ゆえに我あり]」では電子メールを挙げていた[^mua] が， HTTPS も基盤である [PKI が揺らいでいる]({{< ref "/remark/2015/1128-diary.md#pki" >}})。
もっと言うと利便性やセキュリティの名のもとに HTTPS 通信に「中間者攻撃（man-in-the-middle-attack）」を仕掛けて通信内容を「監査」する製品がセキュリティ企業などから提供され，ユーザ企業側もそれを使うことをよしとしている。
しかも「中間者攻撃」の仕組みがあることが分かっているのなら，犯罪者だってそれを利用できる可能性があるわけだ。
知らぬは末端のユーザのみである。

[^mua]: 私も同じネタで[記事を書いた]({{< ref "/remark/2015/use-the-signal-luke.md" >}})。実は TextSecure が Signal に統合されてからウチのスマホでは動かなくなってしまった。とほほ orz

こう考えると本当に「あらゆる通信に」暗号化を適用すべきなのか疑問に感じてしまう。
政治的圧力には同じ圧力で対抗できるかもしれないが，市場はそんな政治的思惑も「面倒」なルールも全部迂回してしまう。
かといって他に冴えたやり方もないので，とりあえずはせっせと「森を作る」しかないのだが。

強すぎる光は影を濃くするのみだ。
あるいは強すぎる結界は更なる魔を引き寄せる。

これはいつも言っていることだが，ルールが守られないのはルール自体に問題がある。
そろそろ排除するだけのセキュリティも暗号技術も曲がり角に来ているのではないかと思ったりする。

[我匿す、ゆえに我あり]: https://wirelesswire.jp/2015/12/48567/ "我匿す、ゆえに我あり - WirelessWire News（ワイヤレスワイヤーニュース）"