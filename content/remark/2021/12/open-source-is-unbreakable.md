+++
title = "オープンソースは砕けない"
date =  "2021-12-14T21:59:02+09:00"
description = "タイトルは釣り"
image = "/images/attention/kitten.jpg"
tags = [ "code", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Log4j の対応に追われている現場エンジニアの皆様にお見舞い申し上げます。

こういうことが起きると

- ["Open Source" is Broken - Xe](https://christine.website/blog/open-source-broken-2021-12-11) ([日本語訳](https://okuranagaimo.blogspot.com/2021/12/blog-post_13.html "ブログ: 「オープンソース」は壊れている"))
- ["Open source" is not broken - nadh.in](https://nadh.in/blog/open-source-is-not-broken/) ([日本語訳](https://okuranagaimo.blogspot.com/2021/12/blog-post_89.html "ブログ: 「オープンソース」は壊れていない"))

とか言い出す連中がいるわけで，ネタとしてはまぁまぁ面白いけど，食傷気味なところではある。
これらに対する反応としては Twitter TL で見かけた

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">20年以上みんなずっと同じ話してるなと思ってしまうが、オープンソースが壊れている、壊れていないの話がやたらに流れている。この文脈ならフリーソフトウェアの時代からずっと壊れてるんだよ。それでも動いているのは自由だからだよ。</p>&mdash; Shuji Sado (佐渡 秀治) (@shujisado) <a href="https://twitter.com/shujisado/status/1470662395228225536?ref_src=twsrc%5Etfw">December 14, 2021</a></blockquote>
{{< /fig-gen >}}

あたりを強く支持する。

自由には責任が伴う。
何故なら社会集団を構成する私達にとって自由とは自明のものではないから。
「個人の自由」は私たちの人類社会が長い時間をかけて「勝ち取った」ものであり，近代の夢で，ネットが社会と接続した現代でもそれは同じである。

でも「フリーソフトウェア」をわざわざ「オープンソース・ソフトウェア」と言い直したのにはちゃんと意味があると思う。

ソフトウェアを自由に公開する意図は様々だ。
「それがぼくには楽しかったから」で産業構造をひっくり返す製品を送り出した人もいる。
また，何らかの社会的使命のもとに製品を作り続けている人もいるかもしれない。
あるいは私のように「独占する価値もないから」オープンソース・ライセンスを付けて公衆に放り出してるだけの人もいるだろう。

オープンソースの功績・貢献をお金に換えることについては昔から議論があるが，上述したようにオープンソースで括れるプレイヤーは本当に多様で，もちろん商業的な成功例はあるが，一般化できないように思える。

そもそも「オープンソースの功績・貢献をお金に換える」というのは問いの立て方として正しいのだろうか。

いま話題の log4j とか脆弱性発生時点でスポンサーが３人しかいなかったと騒がれている。

{{< fig-quote type="markdown" title="“Open Source” is Broken" link="https://christine.website/blog/open-source-broken-2021-12-11" lang="en" >}}
{{% quote %}}As of yesterday, there were a grand total of three sponsors for this person's work{{% /quote %}}.
{{< /fig-quote >}}

GnuPG の作者である Werner Koch さんは経済的に困窮して，[一時期 GnuPG プロジェクトを放棄しようと考えていた](https://arstechnica.com/information-technology/2015/02/once-starving-gnupg-crypto-project-gets-a-windfall-but-can-it-be-saved/ "Once-starving GnuPG crypto project gets a windfall. Now comes the hard part | Ars Technica")ことがあるそうな[^gpg1]。
Linux Kernel なんてイレギュラー中のイレギュラーで，あんな奇跡は二度と起こらないと断言しておこう。
一時期はデジタル補完通貨[^cc1] と FinTech で皆ウハウハみたいな話もあったけど，あれって本当に儲かってるのは miner と山師だけでしょ（笑）

[^gpg1]: [GnuPG は現在も寄付を受け付けている](https://gnupg.org/donate/index.ja.html)。
[^cc1]: 暗号通貨とか暗号資産とか言ってやらん（笑）

この業界に入るきっかけになった会社では色々なことを叩き込まれたが，その中のひとつに「仕事は終わらせるもの」というのがある。
どんな仕事にも終わりがあるし，真っ当に終わらせることで「次」に進むことができる。

離陸したならちゃんと着陸しないとね。
まぁ，エンジニアと違って経営者は着陸した先を考えるものだけど（笑）

一方で個人の趣味や興味で作ったものに対して，その趣味や興味が他所に移ったとして，そのことで批判を受ける義理はないよね。
でもそういったものでも社会的な影響力を持つと「やーめた！」で終わらせることが難しくなる。

私はオープンソース・プロジェクトの問題のひとつは「終わらせる」ことが難しいことにあるんじゃないかと思っている。

今の GitHub はアーカイブ機能でリポジトリを凍結させられる。
この機能によってプロジェクトを明示的に「終わらせる」ことができる。
これは結構重要な機能だと思う。

個人で無限の責任を負うことはできない。
それがたとえ自由の代償であってもだ。
責任を有限にするには「境界」を設定しなければならない。
それは時間軸についても同じ。
興味本位で始めたことであっても，死ぬまで同じプロジェクトに囚われるとか嫌過ぎる（笑）

## 参考図書

{{% review-paapi "4309242456" %}} <!-- リナックスの革命 Hacker Ethic -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "4621045938" %}} <!-- いかにして問題をとくか -->
{{% review-paapi "B01DSV9KDU" %}} <!-- ダイヤモンドは砕けない -->
