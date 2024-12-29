+++
title = "書きたくないときには書かなくていい"
date =  "2024-12-28T17:39:01+09:00"
description = "大事なのは好奇心を錆びつかせないこと"
image = "/images/attention/kitten.jpg"
tags = [ "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

皆さま，年末をいかがお過ごしでしょうか。
勤務先は[昨年]({{< ref "/remark/2023/12/bonenkai.md" >}} "はじめての一泊忘年会")と同じく温泉旅館で一泊忘年会を行いましたよ。
ただし今年は皆生温泉へ。

{{< fig-img src="./54229882215_3c32d3255e_e.jpg" title="皆生温泉上がりのコーヒー牛乳（白バラ） | Flickr" link="https://www.flickr.com/photos/spiegel/54229882215/" >}}

{{< fig-img src="./54229724428_5a643384fe_e.jpg" title="宴会場 机と椅子助かる | Flickr" link="https://www.flickr.com/photos/spiegel/54229724428/" >}}

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

今回は Bluesky で見かけた記事をネタに戯れ言を書いてみる。

- [なぜ俺たち中年はアウトプットできなくなったのか #ポエム - Qiita](https://qiita.com/moroi/items/e39fb42172bfe2d4d209)
- [「本を読んでいれば、教養が身に付く」は勘違い　本物の教養とは | 日経BOOKプラス](https://bookplus.nikkei.com/atcl/column/121100449/121100003/)

「俺たち」ゆーな（笑） ぶっちゃけ「時間」はつくるものだし。
以上！

...では面白くないので続きを。

この辺の記事を読んでて思い出したのは『[Effective Java](https://www.amazon.co.jp/dp/4621303252?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』や『[プログラミング言語Go](https://www.amazon.co.jp/dp/4621300253?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の翻訳でおなじみ柴田芳樹さんが挙げておられる以下のイメージ図だ（グラフっぽいけど定量評価ではない）。

{{< fig-img-quote src="https://yshibata.c.blog.ss-blog.jp/_images/blog/_d3f/yshibata/E68890E995B7E382ABE383BCE38396.JPG" title="ソフトウェアエンジニアの成長カーブ（２）：柴田 芳樹 (Yoshiki Shibata)：SSブログ" link="https://yshibata.blog.ss-blog.jp/2010-01-27" width="1057" >}}

この図について教えてもらったのは，とある読書会の雑談でだったが，その時の話がこんな感じだった。

{{< fig-quote type="markdown" title="松江に出戻り5度目の冬（Advent Calendar）" link="/remark/2022/12/09-san-in/" >}}
この図が示すのは，会社から与えられている仕事に慣れて自律的に「学ぶ」ことを止めてしまうと，手持ちのスキルは（時代の流れで）先細りし，さらに学習習慣が失われることにより新たな技術・スキルを得る機会を逸してしまうため，全体としてスキルレベルが低下してしまう，というものです。

特に日本企業は従業員の「今」のスキルレベルに合わせて「できそうな仕事」を見繕って割り振る傾向があり，所属部署から与えられる仕事をこなすことに満足しているとあっという間にスキルが先細りしてしまう，というような話を読書会でもしていました。
{{< /fig-quote >}}

ちなみに縦軸の「[ソフトウェア・スキル・インデックス](http://yshibata.blog.so-net.ne.jp/2010-01-11 "技術者のレベルとソフトウェア開発の難易度（２）")」については『[プログラマー“まだまだ”現役続行](https://www.amazon.co.jp/dp/B0CCY7VJ4C?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』に詳しく書かれているので是非ご一読を[^p1]。

[^p1]: 『[プログラマー“まだまだ”現役続行](https://www.amazon.co.jp/dp/B0CCY7VJ4C?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』が最初に出たのは2006年。当時は「プログラマ35歳定年説」とか言われてた頃で，仕事でも「私はコード読めないので」とか平気で言う SE がいたりした。この本はそういう風潮に対するアンチテーゼになっていたわけだ。ただし中身は今でも十分通用すると思う。奇抜な内容ではないが，こういうのは継続することが重要だからなぁ。ちなみに著者の柴田芳樹さんは既に[65歳](https://note.com/yoshiki_shibata/n/n86d730f758e5 "65歳になりました｜柴田 芳樹")とのことだが，今も「[現役](https://note.com/yoshiki_shibata/n/n1489871a51ac "カウシェでのAPI仕様整備と効率的なバックエンド開発の実現｜柴田 芳樹")」だそうだ。

自身の話で恐縮だが，私は思考が言語的じゃないのよ。
だから「文章を書く」というのは思考を言語に翻訳する作業なのだ（ブログ記事は意図的に文体を崩して書いてるけど[^w]）。
と同時にインプットした（と思ってる）ことを振り返る機会でもある。
これはプログラミングでも同じ。
思考を言語化することで理解の解像度が上がるし，見落としてたことにも気付ける。

[^w]: かっちりした文章は仕事でアホほど書いてるし。仕事以外で真面目な文章なんて書きたくない。

最初に挙げたイメージ図にもある通り，エンジニアの学習・勉強は現役でいるかぎり続くものだけど，続けるための最大のモチベーションは「それ」を楽しめるか否かだと思う。
ただ（仕事はともかく）個人の趣味・興味・関心は歳を重ねるごとに移り変わって当たり前である。
なのにアウトプットは若い頃と同じなんてありえない。

大事なのは好奇心を錆びつかせないことなので，一時的にアウトプットを止めてインプットに専念していいと思うし疲れたならインプットも休んでいいと思う。
まぁ，休むのなら期限を切ったほうがいいだろうけど。

私なんて今年このブログで書いた Go 関連の記事は[一本]({{< ref "/golang/pseudo-random-number-generator-v2.md" >}} "Go 1.22 における疑似乱数生成器")だけだぜ。
年間の記事数はそう変わらないと思うけど。
松江に {{% ruby "かえ" %}}帰郷{{% /ruby %}}って自転車を再開して明らかに趣味が変わったもんな（笑）

『[はじめて学ぶ ビデオゲームの心理学](https://www.amazon.co.jp/dp/B0C9Z7KGRN?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』にも書いてあるぢゃん。

{{< fig-quote type="markdown" title="『はじめて学ぶ ビデオゲームの心理学』 p.71" link="https://www.amazon.co.jp/dp/4571210450?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
年齢に関わらず、遊びは私たちの精神を鋭敏に保つために重要です。 [...] 遊ぶことは学ぶことです。
{{< /fig-quote >}}

って。

来年は『[Go言語で学ぶ並行プログラミング](https://www.amazon.co.jp/dp/B0DNYMMBBQ?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』の[読書会]が始まるし CVSSv4 の実装もいい加減始めないとなぁ...

[読書会]: https://technical-book-reading-2.connpass.com/event/337562/ "第1回『Go言語で学ぶ並行プログラミング』オンライン読書会 - connpass"

## 参考図書

{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B0CCY7VJ4C" %}} <!-- プログラマー"まだまだ"現役続行 -->
{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
