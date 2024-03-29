+++
title = "「ネットの中立性」と「後出しジャンケン」と「メディアによる検閲を迂回する」"
date =  "2017-12-23T18:37:50+09:00"
update = "2018-01-04T20:41:28+09:00"
description = "自分用にメモしておくことは大事なので，覚え書きとして書いておく。"
image = "/images/attention/remark.jpg"
tags = [ "code", "law", "intellectual-property", "hacker-ethic", "internet", "engineering", "media", "censorship" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

年も押し迫って初日の出が云々とか聞かれる今日この頃にこんな暑苦しい話もどうかと思ったんだけど，[自分用にメモしておく](http://www.hyuki.com/dream/log.html "あなた自身の航海日誌")ことは大事なので，覚え書きとして書いておく。

1. [ネットの中立性]({{< relref "#network-neutrality" >}})
1. [後出しジャンケン]({{< relref "#opensource" >}})
1. [メディアによる検閲を迂回する]({{< relref "#media-censorship" >}})

## ネットの中立性{#network-neutrality}

- [FCC、オバマ前大統領導入の「ネット中立性」廃止勧告　12月に採決へ - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1711/22/news062.html)
    - [「ネット中立性」が米国で廃止へ　日本のサービスにも影響はある？ - ITmedia PC USER](http://www.itmedia.co.jp/pcuser/articles/1711/26/news011.html)
    - [FCC現委員長のネット中立性悪者説はすべて正しくない、と同僚の委員が指摘  |  TechCrunch Japan](https://techcrunch.com/2017/11/30/fcc-commissioner-clyburn-takes-down-chairmans-net-neutrality-doom-and-gloom/)
    - [ネット中立性の擁護者たちが12月7日に全国のVerizonショップで抗議集会を開く  |  TechCrunch Japan](https://techcrunch.com/2017/11/22/net-neutrality-advocates-plan-protests-for-december-7-at-verizon-stores/)
    - [インターネットはすでに死んでいる* | 辺境社会研究室](https://youkoseki.tumblr.com/post/167900801310/network-unneutrality)
    - [ネット中立性撤廃に反対する公開書簡をテクノロジ分野の先駆者らが投稿 - CNET Japan](https://japan.cnet.com/article/35111780/)
- [The battle for the net continues after FCC erases net neutrality rules - Creative Commons](https://creativecommons.org/2017/12/15/battle-net-continues-fcc-erases-net-neutrality-rules/)
- [ニュース - 米でオバマ時代の「ネット中立性」の撤廃が承認：ITpro](http://itpro.nikkeibp.co.jp/atcl/news/17/121502874/?rt=nocnt)
- [チーム・インターネットの戦いは終わらない：ネット中立性のためにできること – P2Pとかその辺のお話R](http://p2ptk.org/net-neutrality/662)
- [「ネット中立性」の撤廃は決まれど、通信業界との攻防は続く｜WIRED.jp](https://wired.jp/2017/12/21/net-neutrality-fight-moves-to-courts/)

週明けくらいから FCC の決定に対する日本語の記事もボチボチ見えだしたが，誰も「俺たちの戦いはこれからだ！」の決まり文句を言ってくれないので（打ち切りフラグだから？），ちょっと寂しい。

それはともかく，「ネット中立性」については本当に長～いあいだ議論されている。
個人的には中立性を守ろうとする側の言い分はもっともだし私も「激しく同意」するところである。

今のインターネットは実質的にキャリア企業に回線を牛耳られている状態であり，ネットの「停止スイッチ」を握っているのが彼等である。
最近は「核のボタン」を現米国大統領が握っていることを危惧する意見があるが，「ネット中立性」を放棄するのはこれに近い脅威であると言えよう。
残念ながら今のインターネットは米国を中心に回ってるので，今回の FCC の決定は世界にとっても他人事ではない。

個人的には，日本のキャリア企業がどう動くかに注目している。

## 後出しジャンケン{#opensource}

- [ニュース - エヌビディアが消費者向けGPUのライセンスを変更、データセンターへの導入を禁止：ITpro](http://itpro.nikkeibp.co.jp/atcl/news/17/122002893/?rt=nocnt)
- [さくら、TITAN X搭載GPUサーバサービスの新規提供を一時停止　NVIDIAのライセンス条件変更で - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1712/22/news071.html)
- [NVIDIAが規約変更によりGeForceのデータセンター利用を制限。大学などの研究活動にも大ブレーキ - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2017/12/62658/)
- [本の虫: nVidia、GeForceのデータセンターでの利用を禁止する](https://cpplover.blogspot.jp/2017/12/nvidiageforce.html)
- [みんな、これからは深層学習にはGeForceではなくRadeonを使おう - shi3zの長文日記](http://d.hatena.ne.jp/shi3z/20171223/1513980907)
- [GeForceのデータセンター利用を禁止する使用許諾契約に対してNVIDIAが声明  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1099/481/index.html)

恥ずかしながら GPU 市場の動向には全く疎い。
NVIDIA と AMD で一種の「住み分け」ができてることも知らなかった[^gpu1]。
道理で両者が全く異なるスタンスをとってるのに競合しないはずである。

[^gpu1]: ブロックチェーン関連では AMD の独壇場らしい。なので [NVIDIA の修正 EULA](http://www.nvidia.com/content/DriverDownload-March2009/licence.php?lang=us&type=GeForce "License For Customer Use of NVIDIA GeForce Software") でもブロックチェーン処理は制限の対象から外れている（2.1.3 章）。

- [AMD、オープンソースのグラフィックスドライバの開発を支援へ - CNET Japan](https://japan.cnet.com/article/20356020/)
- [【後藤弘茂のWeekly海外ニュース】AMDがドライバを含めたGPUソフトウェアをオープンソース化  - PC Watch](https://pc.watch.impress.co.jp/docs/column/kaigai/736647.html)

オープンソースと言っても上から下まで全部「オープンソース化」されているわけではなく，ハードウェアに近いところほどベンダ企業による統制が強くなる。
ハードウェアまで全てオープンソース化しようという動きはあるにはあるが市場にインパクトを与えるほどではないのが現状である。

このような「後出しジャンケン」がライセンスとして有効なのかどうか甚だ疑問なのだが，問題はそこに留まらない。
他のハードウェア企業の幾つかはこれに便乗するだろう。
こんなことが許されるのならオープンソースなんか鼻糞である。

## メディアによる検閲を迂回する{#media-censorship}

（この節は [Twitter のスレッドで書きなぐった tweets](https://twitter.com/spiegel_2007/status/943788617268207616) の再構成です）

- [Japanese doctor wins global prize for standing up to anti-vaccine activists | Cornell Alliance for Science](https://allianceforscience.cornell.edu/blog/japanese-doctor-wins-global-prize-standing-anti-vaccine-activists)
- [ジョン・マドックス賞受賞スピーチ全文「10万個の子宮」｜村中璃子　Riko Muranaka｜note](https://note.mu/rikomuranaka/n/n64eb122ac396)
- [子宮頸がんワクチン副作用とマスコミの役割｜新・山形月報！｜山形浩生｜cakes（ケイクス）](https://cakes.mu/posts/18760)
- [高橋洋一の霞ヶ関ウォッチ 「ワクチンめぐる表彰」問題で再認識　国内新聞・テレビの「異様さ」 : J-CASTニュース](https://www.j-cast.com/2017/12/14316605.html?p=all)

[J-CAST の記事](https://www.j-cast.com/2017/12/14316605.html?p=all "高橋洋一の霞ヶ関ウォッチ 「ワクチンめぐる表彰」問題で再認識　国内新聞・テレビの「異様さ」 : J-CASTニュース")の論点は2つあると思う。
ひとつは間違った報道による「被害」をどう補償するか，または「被害」をどう防ぐか。
もうひとつはマスコミによる報道の選択（＝検閲）である。

前者については[山形浩生さんの記事](https://cakes.mu/posts/18760 "子宮頸がんワクチン副作用とマスコミの役割｜新・山形月報！｜山形浩生｜cakes")のほうが参考になるのではないだろうか。
はっきり言ってこれは「集団訴訟」案件だと思う。
被害者が自分が被害を受けてると自覚するかどうかにもよるけど。

後者については色々意見があるだろうが，私個人としては「どうでもいい」。
テレビ・新聞の報道が，仕事を含む，日常生活の関心領域について（社会性・公益性の有無に関わらず）報道することのほうが稀だからだ。
加えて言うなら，私はイエロー・ジャーナリズムには関心がない。
故にテレビ・新聞の報道をスルーしても全く支障がないのだ。

私は社会人になって数年で新聞の購読を止めた。
理由は3つ。
勤務先や客先が買ってる新聞が読めたから。
長期出張で自宅を空けることが増えたから（若い頃の話だよ）。
そして，新聞はテレビより情報が遅いから，である。

更に地上アナログ波の停波タイミングで本格的にテレビを見ることを止めた。
これも理由が3つ。
情報の事実性に関する S/N 比が悪すぎるから。
テレビ以外に享受できる娯楽が増えたから（アニメもネットで見れる）。
そして，テレビはネットより情報が遅いから[^media1]，でる。

[^media1]: おそらくある分野の最先端にいる人から見れば「ネットの情報は遅い」と感じるのではないだろうか。たとえばセキュリティ分野であれば，インシデントが公衆空間にて顕在化するか「対応済み」となった後でなければ公開されない。更にそれをテレビ・新聞等のメディアが報道するか否かは相当に恣意的である。

- [「テレビという共同体」 — Baldanders.info](https://baldanders.info/blog/000307/)

これはネットが情報摂取に関して優位であるということではない。
「情報の事実性に関する S/N 比が悪すぎる」というのはネットにも当てはまる。
今回の件は，たまたまネットのほうが「1次情報」に近いところにいる，と言うこともできる。

問題は，あるメディア（テレビでも新聞でもネットメディアでも）が特定の情報を選択（＝検閲）することではなく，検閲を迂回してちゃんと「事実」に手が届く手段があるかどうかである。
**Unreachable であることは故障と同一であり，故障を迂回するのがインターネット本来の機能である**。
その上で（メディアではなく）エンド・ユーザが膨大な情報を如何に効率よく偏りなく誤配を残しつつフィルタリングできるかが難問だったりする。
そのためにも「[ネットの中立性]({{< relref "#network-neutrality" >}})」を法的に担保しておくことは重要なのである。
（おおっ。ちゃんと話が繋がった`w`）

そう言えば Facebook は「偽ニュースに印をつける」という頭の悪い考えからようやく脱却できたようである。

- [Facebook、偽ニュースの赤フラグ表示を廃止--関連記事を強化へ - CNET Japan](https://japan.cnet.com/article/35112356/)

時間的空間的な制約の多いテレビ・新聞ではできないことのひとつと言っていいだろう。
個人的には，あるトピックについて虚実にかかわらず機械的に記事を集約できるようになれば面白いと思う。
つか，集約して欲しい。
似たような記事が TL にずらずらと並ぶのは鬱陶しいことこの上ない。

Twitter も（検閲に血道を上げるんじゃなく）これを真似してくれんかね。
自分の tweets をスレッドで繋げることができたんだから，関連 tweets のカスケード表示もできそうなものなんだけど。

## ブックマーク{#bookmark}

- [“The Shadow Web” — Baldanders.info](https://baldanders.info/blog/000599/)

## 参考図書

{{% review-paapi "B01CYDGUV8" %}} <!-- CODE VERSION 2.0 -->
{{% review-paapi "B00DIM6BE6" %}} <!-- インテンション・エコノミー -->
{{% review-paapi "4150504598" %}} <!-- フィルターバブル -->
