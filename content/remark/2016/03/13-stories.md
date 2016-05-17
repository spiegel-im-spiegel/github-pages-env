+++
date = "2016-03-13T13:03:59+09:00"
update = "2016-05-17T22:33:36+09:00"
description = "サイバーテロ企業 Microsoft / 3月9日は皆既日食でした / いまさら「2033年問題」 / 『プログラミング言語 Go』"
draft = false
tags = ["risk", "windows", "astronomy", "eclipse", "calendar", "book", "golang"]
title = "週末スペシャル： サイバーテロ企業 Microsoft"

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
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [サイバーテロ企業 Microsoft]({{< relref "#win" >}})
1. [3月9日は皆既日食でした]({{< relref "#eclipse" >}})
1. [いまさら「2033年問題」]({{< relref "#cal" >}})
1. [『プログラミング言語 Go』]({{< relref "#prog" >}})

## サイバーテロ企業 Microsoft{#win}

あぁ，ついに知り合いに被害者が出てしまいましたよ。

- [Windows 10 への強引アップデートというサイバーテロについて - 電気ウナギ的○○](http://blog.netandfield.com/shar/2016/03/windows-10.html)
- [Windows 10 への強引アップデートというサイバーテロについて（つづき） - 電気ウナギ的○○](http://blog.netandfield.com/shar/2016/03/windows-10-1.html)

しかも [Windows 7 に復元した後もトラブってる](https://www.instagram.com/p/BC22bDzI4vI/)ようだ。
はっきりいって昔から Windows のメジャーバージョンのアップグレードはトラブることが多い。
やるなら真っさらな状態からクリーン・インストールしないと。

今月の Windows Update は特に酷かったようで

- [IEのセキュリティパッチにWindows 10アップグレード広告を挿入＆システム管理者にアップグレードを促す広告も現れいよいよMicrosoftの本気度が明らかに - GIGAZINE](http://gigazine.net/news/20160311-windows10-nag-in-ie-patch/)
- [2016年3月のWindows Updateの注意事項](http://freesoft.tvbok.com/cat97/2016/2016_03_windows_update.html)
- [(2016年3月)Windows10への半強制アップグレードが再び猛威を奮っている模様？(追記あり)](http://freesoft.tvbok.com/cat97/2016/gwx_app_2016_03.html)

これ集団訴訟起こしたら勝てるんじゃねーの。
私は[事前に対策]({{< relref "remark/2015/windows-10-upgrade-problem.md" >}} "また Windows 10 にヤラレタ（KB3112343 の恐怖） — しっぽのさきっちょ")してるので実害ないけど（でも危ないので Windows Update の自動更新は無効にした）。
もう Microsoft 社は「サイバーテロ企業」と認定していいんじゃないかな。
セキュリティ企業各社におかれましては何卒 Microsoft 社の蛮行から防御するツールを開発していただきたい。
買うよ，今なら。

[前にも書いた]({{< relref "remark/2015/windows-10-upgrade-problem.md" >}} "また Windows 10 にヤラレタ（KB3112343 の恐怖） — しっぽのさきっちょ")が， Windows 10 が悪いとは言わない（敢えてオススメはしないけど）。
しかし，その気がない人にまで無理やりアップグレードさせようというのは悪質極まる。
しかも今回はセキュリティ・アップデートに広告パッチを混ぜるという蛮行を犯した。
これが evil じゃないとは言わせない。

Microsoft は Windows を捨てるべきである。
もう OS を売って儲ける時代ではない。
Windows の（マネタイズを含めた）思想は20年以上も**大昔**のものでインターネットを含む現代の状況にマッチしていない。
自前でカーネルを作るのは諦めて [X11] 上に Windows 風の GUI shell を構築するとか [Wine] にコミットするとかしたほうがいい。
[.NET だってオープンソース](https://msdn.microsoft.com/ja-jp/library/dn878908%28v=vs.110%29.aspx ".NET Core とオープン ソース")にできたんだから，できないことはないだろう。
どうしても Windows を残したいのなら携帯端末でやってくれ。

2020年までに自宅のメインマシンを [Ubuntu] か何かに換装しようと思ってるのだけど，こりゃあ計画を前倒しにした方がいいかなぁ。

そうそう， OS X は個人的に「アウト・オブ・眼中」です。

- [「PCをWin7のままにしておきたいのに強制的にWin10にするMSが嫌だ！Macに行く！」という方へMacユーザーとして言っておきたいこと](http://stocker.jp/500W/2016/03/12/switch2mac/)

### 参考

- [Tech TIPS：まだWindows 10へアップグレードしたくない人のための設定まとめ - ＠IT](http://www.atmarkit.co.jp/ait/articles/1603/18/news047.html)
- [「Windows 10」自動アップデートをオフにするツール「Never 10」が公開 - ZDNet Japan](http://japan.zdnet.com/article/35080272/)
- [Windows 10への自動アップグレードスケジュールの通知がさらに凶悪化してWindows Updateと一体化、キャンセル方法はコレ - GIGAZINE](http://gigazine.net/news/20160517-windows-10-auto-upgrade/)

## 3月9日は皆既日食でした{#eclipse}

（以下の動画は音楽が流れるので再生時には周りに注意）

{{< fig-youtube id="VOdVQnUKmE0" width="500" height="281" title="the total solar eclipse time laps movie 2016/皆既日食まとめ動画2016 weathernews - YouTube" >}}

日本では部分日食だったけど，あいにくの天気で見れなかった人が多そうだ。
今回は人工衛星からの映像が話題になった。
人工衛星からは皆既食の地域が丸い陰となっている。

{{< fig-img src="http://www.jma-net.go.jp/sat/data/web89/parts89/himawari8_sample_data/thumbnail/201603091110_TRC_SolarEclipse_s.png" title="「ひまわり８号のサンプル画像」より" link="http://www.jma-net.go.jp/sat/data/web89/himawari8_sample_data.html#nisshoku" width="500" >}}

今年は9月1日に金環日食があるが，これも日本からは見れない。

## いまさら「2033年問題」{#cal}

- [琉球新報コラム「2033年問題」 | TOPICS](http://hoshisora.jp/topics/?p=2029)

大丈夫だ。
問題ない（笑）

「2033年問題」について詳しくは国立天文台暦計算室による以下を参照のこと。

- [旧暦2033年問題について](http://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2014.html)

2033年から2034年の春分にかけて中気[^c]を含まない朔望月が3回もあり，いわゆる「旧暦」のルール

[^c]: 「中気」は二十四節気のうち冬至を0起点として偶数番目のものを指す。奇数番目のものは「節気」と呼ぶ。現代の二十四節気は天球上の太陽の視黄経（太陽黄経）で決まる（定気法）。例えば太陽が黄経0度を通過した瞬間を含む日が春分となる。ただし二十四節気の暦上の基点は冬至で，古い暦では太陽観測により冬至を確定し次の冬至までの1年を24等分することで二十四節気が決まっていた（平気法）。

1. 冬至を含む朔望月を11月とする
2. 春分を含む朔望月を2月とする
3. 夏至を含む朔望月を5月とする
4. 秋分を含む朔望月を8月とする
5. 1年が13ヶ月ある場合，中気を含まない朔望月を閏月とする

のうち 1 と 4 が衝突するため，どの月を閏月とすべきか決められないということのようだ。

ルールにない事態なのだからルールを作ればよい。
具体的には 1 から 4 のルールの中で優先順位を決めればよい。
先程の国立天文台暦計算室の記事では「時憲暦」の「冬至を含む月から次に冬至を含む月までに13か月ある場合に，中気が入らない最初の月を閏月とする」というルールを紹介している。

ポイントは誰がそれを決めるのか，ということだが

**あなたが「旧暦」と思うものが「旧暦」です。ただし、他人の賛同を得られるとは限りません**

でいいんじゃないかな。
得意ぢゃん，日本人って，そういうの。

「旧暦」は民間暦のひとつに過ぎない。
民間暦なんてものは宗教・宗派あるいは地域で勝手に決めているものなのだから，それぞれのコミュニティで暦を合わせていれば深刻な問題にはならない。
「2033年問題」の場合は「中秋の名月」がいつになるかでちょっともめるくらいだろう。
ちなみに国立天文台には暦を決定する大事な役割がある[^naoj] が民間暦は関知しない。

[^naoj]: 国立天文台では毎年2月1日に翌年の暦を発表する。

## 『プログラミング言語 Go』{#prog}

6月に出るという『[プログラミング言語 Go](http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/)』を買おうかどうか考え中。
[原書](http://www.amazon.co.jp/exec/obidos/ASIN/0134190440/baldandersinf-22/)の目次を見るかぎり基本的な部分はちゃんと押さえてるみたいだし「買ってもいいかな」と思ってはいるのだけど。

ちなみに『[Go 言語による Web アプリケーション開発](http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/)』は既に買った。
ざくっと斜め読みしただけだけど，そのうち役に立つ日が来るかもしれない。

この手の本は，昔は重宝したけど今は賞味期限が早いので手にした時にはもう骨董品レベルだったりするんだよなぁ。
[Go 言語]はネットにあるドキュメント類が充実しているので正直「書籍」はなくてもいいんだが，英語が中心なので，英語不得手の私としては日本語のまとまったドキュメントが欲しいと思ってしまうわけなのさ。

## その他の気になる記事{#other}

- [GitBookの使用例 - Qiita](http://qiita.com/syui/items/429962c76f6c3117f351)
    - [Gitbookを高速化する方法 - Qiita](http://qiita.com/nacika_ins/items/24b91f654cf2dc59dd81)
- [Big Sky :: コマンドラインから JSON が簡単に作れるツール jo](http://mattn.kaoriya.net/software/20160309091404.htm)
- [福島原発事故から5年：水産物の放射性セシウム汚染の現状 « WIRED.jp](http://wired.jp/2016/03/10/five-years-after-fukushimas-contamination/)
    - [Fukushima Radioactivity in U.S. West Coast Tuna - SWFSC](https://swfsc.noaa.gov/textblock.aspx?Division=FRD&id=20593)
    - [Model simulations on the long-term dispersal of 137Cs released into the Pacific Ocean off Fukushima - IOPscience](http://iopscience.iop.org/1748-9326/7/3/034004/)
- [『広島平和研究』第3号](http://www.hiroshima-cu.ac.jp/modules/peace_j/content0252.html)
- [TPPでよみがえる“マジコンプレイ違法化”の亡霊、「みなし侵害」で成仏するか？　著作権法改正案が明らかに -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160310_747743.html)
- [違法コピーに悩まされる美少女ゲームメーカー、「世界最高レベル」のコピープロテクト実装へ | スラド セキュリティ](http://security.srad.jp/story/16/03/10/1530233/) : 1ヶ月かそこらで突破されちゃうようなシロモノが「世界最高レベル」なのか？ つか，今時「コピー出来ないソフト」とか，悪人には関係ないし善人には使い勝手が悪いだけだと思うのだが，悪人が利するシステムに意味があるのか？
- [GitHub、コメントに「いいね!」など6種類の感情表現が可能に -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160311_747871.html)

[X11]: http://www.x.org/ "X.Org"
[Wine]: https://www.winehq.org/ "WineHQ - Run Windows applications on Linux, BSD, Solaris and Mac OS X"
[Ubuntu]: http://www.ubuntu.com/ "The leading OS for PC, tablet, phone and cloud | Ubuntu"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4416115458/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51XoyiTnmFL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4416115458/baldandersinf-22/">天文年鑑2016年版</a></dt><dd>天文年鑑編集委員会 </dd><dd>誠文堂新光社 2015-11-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4805208899/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4805208899.09._SCTHUMBZZZ_.jpg"  alt="天文手帳 2016年版: 星座早見盤付 天文ポケット年鑑"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/462108965X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/462108965X.09._SCTHUMBZZZ_.jpg"  alt="理科年表 平成28年"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B016YNZD0I/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B016YNZD0I.09._SCTHUMBZZZ_.jpg"  alt="天文ガイド 2016年 01 月号  [特大号 付録付き]"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B017VPHTX2/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B017VPHTX2.09._SCTHUMBZZZ_.jpg"  alt="月刊 星ナビ 2016年 1月号"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4416115512/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4416115512.09._SCTHUMBZZZ_.jpg"  alt="藤井 旭の天文年鑑 2016年版: スターウォッチング完全ガイド"  /></a> </p>
<p class="description">天文ファン必携。2016年版</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-12-29">2015-12-29</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41aCueik45L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-15</dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117402/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117402.09._SCTHUMBZZZ_.jpg"  alt="ハイパフォーマンスPython"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/0134190440/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/0134190440.09._SCTHUMBZZZ_.jpg"  alt="The Go Programming Language (Addison-Wesley Professional Computing Series)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166340/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774166340.09._SCTHUMBZZZ_.jpg"  alt="Vim script テクニックバイブル ~Vim使いの魔法の杖"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。買おうかどうか悩み中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-12">2016-03-12</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51UoREcNrnL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/">Go言語によるWebアプリケーション開発</a></dt><dd>Mat Ryer 鵜飼 文敏 </dd><dd>オライリージャパン 2016-01-22</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4621300253.09._SCTHUMBZZZ_.jpg"  alt="プログラミング言語Go"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774178667/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774178667.09._SCTHUMBZZZ_.jpg"  alt="nginx実践入門 (WEB+DB PRESS plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4863541783/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4863541783.09._SCTHUMBZZZ_.jpg"  alt="改訂2版 基礎からわかる Go言語"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774179930/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774179930.09._SCTHUMBZZZ_.jpg"  alt="サーバ/インフラエンジニア養成読本 DevOps編 [Infrastructure as Code を実践するノウハウが満載! ] (Software Design plus)"  /></a> </p>
<p class="description">日本語監訳者による解説（付録 B）が意外に役に立つ感じ。 Web アプリケーションだけでなく，サーバサイドで動く CLI アプリへの言及もある。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-13">2016-03-13</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
