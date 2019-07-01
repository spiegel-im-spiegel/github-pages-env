+++
date = "2016-05-28T22:59:50+09:00"
description = "Barack Obama 米国大統領の来広 / Google vs Oracle の訴訟の行方 2 / Windows 7 用の Rollup が出た / その他の気になる記事"
tags = ["politics", "international", "hiroshima", "code", "intellectual-property", "copyright", "java", "windows"]
title = "週末スペシャル： Barack Obama 米国大統領の来広"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今参加してるプロジェクトが「徒労感」半端ない。
はずれクジ引いちゃったかなぁ。
週末はぐったりして寝てるだけのことも多い。
倒れないうちに手を打ったほうがいいか？

1. [Barack Obama 米国大統領の来広]({{< relref "#hrsm" >}})
1. [Google vs Oracle の訴訟の行方 2]({{< relref "#api" >}})
1. [Windows 7 用の Rollup が出た]({{< relref "#rlup" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## Barack Obama 米国大統領の来広{#hrsm}

いやぁ，昨日は大変だったみたいですな（以下の動画の26分頃から）。

{{< fig-youtube id="QHIPZhrma6I" title="President Obama Participates in a Wreath Laying Ceremony - YouTube" height="281" width="500" >}}

- [Remarks by President Obama and Prime Minister Abe of Japan at Hiroshima Peace Memorial | whitehouse.gov](https://www.whitehouse.gov/the-press-office/2016/05/27/remarks-president-obama-and-prime-minister-abe-japan-hiroshima-peace) （[日本語訳](http://japanese.japan.usembassy.gov/j/p/tpj-20160527-02.html)）
- [平成28年5月27日 広島訪問　日米両首脳によるステートメント | 平成28年 | 総理の演説・記者会見など | 記者会見 | 首相官邸ホームページ](http://www.kantei.go.jp/jp/97_abe/statement/2016/0527hiroshima.html)

私は職場が郊外なので（ただし通勤するために市街地を経由しなければならない[^a]）通勤時間をずらして回避できたが，市街地はほぼマヒ状態だったようだ。

- [【お願い】オバマ米国大統領広島訪問に伴う交通規制について](http://www.pref.hiroshima.lg.jp/site/police19/gyouji-kisei.html)
- [いやあ、昨日はひどい目にあった(^^; - 電気ウナギ的○○](http://blog.netandfield.com/shar/2016/05/post-2540.html)

[^a]: 以前から思ってるけど，これって広島市の都市設計における最大の欠点だよね。広島市の都市設計は前時代的で郊外の道路は基本的に市街地（広島市民の間では「旧市内」と呼ばれている）から放射状に敷設されている。したがって郊外から郊外へ移動するためには一度市街地を通らなければならない。これはバス路線や電車も同じ。自動車道については市街地を迂回するための高速道路やバイパスが造られたが（広島は地理的に関西・九州間の通り道にあるため，大昔は長距離トラックが市街地を猛スピードで走っていた），公共交通機関については全く改善されないため朝晩の通勤時には（たかだか100万人ほどしかいない）地方都市とは思えないほどの混雑になる（褒めてないよ）。海外では広島市は「コンパクト・シティ」として評価されていると聞くが，それは市街地の中だけの話である。地元 J1 チームのサンフレッチェが市街地ど真ん中にある市民球場跡地への移転にこだわるのはちゃんと理由があるのだ。アストラムラインを環状線にする案も提案されたが，アストラムラインの経営状態は単年でかろうじて黒字で初期の建設費も返せない状況。このままサンフレッチェがいなくなって施設が老朽化したらどうするつもりだろうね。

オバマ大統領は「メッセージだけの大統領」と揶揄されることもあるが，本当の変化は100年単位でゆるゆると進むものである。
今回の来広が，2009年のプラハ演説と同じく，100年先の変化へのトリガーとなれば今はそれで充分だと思う（まぁ次の大統領が誰になるかでちゃぶ台がひっくり返される可能性もあるわけだが）。

{{< fig-quote title="「核兵器なき世界」の本当の意味と「日本の役割」：池内恵 | 中東―危機の震源を読む" link="http://www.fsight.jp/5104" >}}
<q>これはまさに、核保有国と非保有国の格差を永続化させる従来の核不拡散体制の理念を再度打ち出したものである。違いがどこにあるかといえば、この格差を維持することが困難であることが明白となり、新たに核保有を行なおうとする国を抑制する新たな体制を構築する必要が出てきたことである。そのために、まず保有国側が核軍縮の努力を示すことによって倫理的優位性を確保する必要が認識された。オバマのプラハ演説は倫理的優位性の確保と、軍事的優位性の確保を両方得ようとする戦略的な言説である。</q>
{{< /fig-quote >}}

### 参考

- [「オバマ広島訪問」はなぜ感動を呼んだのか：PRESIDENT Online - プレジデント](http://president.jp/articles/-/18253)

## Google vs Oracle の訴訟の行方 2{#api}

- [Google beats Oracle—Android makes “fair use” of Java APIs | Ars Technica](http://arstechnica.com/tech-policy/2016/05/google-wins-trial-against-oracle-as-jury-finds-android-is-fair-use/)
- [グーグル、Java API使用が「フェアユース」と認められる--対オラクル訴訟 - CNET Japan](http://japan.cnet.com/news/business/35083291/)
- [ニュース - Java著作権訴訟でGoogleが勝訴、「フェアユース」が認められる：ITpro](http://itpro.nikkeibp.co.jp/atcl/news/16/052701526/?rt=nocnt)
- [Googleの「公正使用」勝訴後も残る著作権に関する疑問 | TechCrunch Japan](http://jp.techcrunch.com/2016/05/31/20160527copyright-questions-remain-after-googles-fair-use-victory/)
- [グーグルがJava API訴訟でオラクルに勝利--開発者にとって朗報か - ZDNet Japan](http://japan.zdnet.com/article/35083475/)
- [Google’s fair use victory is good for open source | Ars Technica](http://arstechnica.com/tech-policy/2016/06/googles-fair-use-victory-is-good-for-open-source/)

以前「[Google vs Oracle の訴訟の行方](https://baldanders.info/blog/000861/)」でも書いたが， Java API の著作権を巡る争いは「著作権の有無」から著作権を認めた上で Java API の（使用ではなく）利用が fair use にあたるかどうかに争点が移った。

米国で fair use が成立するためには以下の観点において「公正」であることが認められなければならない。

1. 利用の目的や本質
2. 原作品の本質
3. 抜粋の量や実質性
4. 原作品の価値への影響

連邦地方裁判所の陪審は fair use であると評したようだが，恐らくこれから本格的な議論になるだろう。
しかし私たちエンジニアの側はあらゆる可能性を考えて備えておかなければならない。

たとえば [Go 言語]は内部パッケージを含めて MIT ライセンスで提供されている。
API がこのようなことになることで，標準ライブラリやフレームワーク等の扱いが明確でない多くの言語は FLOSS にせざるを得なくなっていくだろう[^b]。
そういう意味で，私企業がコントロールを手放さない Java は高リスクな「終わってる言語」と言える。
日本人は何故か Java が大好きだが，新たにこれから何かを作るのであれば Java は忌避すべきある。

[^b]: そういう意味で Microsoft が [.NET Core をオープンソースにした](http://jp.techcrunch.com/2015/04/30/20150429microsoft-launches-its-net-distribution-for-linux-and-mac/)のは慧眼だよね。

## Windows 7 用の Rollup が出た{#rlup}

- [May 2016 update rollup for Windows 7 SP1 and Windows Server 2008 R2 SP1](https://support.microsoft.com/en-us/kb/3156417)

Windows Update で KB3156417 を導入すればよい。
これを当てることによってこれまで非表示にしていた更新が元に戻るといったことはないようである。
あとは Windows 10 のアップグレード強制さえなくなればなぁ。
早く無料キャンペーン終わってくれ！

- [また Windows 10 にヤラレタ（KB3112343 の恐怖）]({{< ref "/remark/2015/windows-10-upgrade-problem.md" >}})
- [Tech TIPS：終わらないWindows 7のWindows Updateの問題を解決する - ＠IT](http://www.atmarkit.co.jp/ait/articles/1605/26/news029.html)

## その他の気になる記事{#other}

- [Official Google Blog: Introducing Spaces, a tool for small group sharing](https://googleblog.blogspot.jp/2016/05/introducing-spaces-tool-for-small-group.html)
    - [Google、ソーシャル・サービスに再挑戦―グループチャット・アプリのSpacesをリリース | TechCrunch Japan](http://jp.techcrunch.com/2016/05/17/20160516google-tries-its-hand-at-social-again-with-launch-of-group-chat-app-spaces/)
- [ロシアの天才ハッカーによる【新人エンジニアサバイバルガイド】 - Qiita](http://qiita.com/jacksuzuki/items/b2fa6b44962e73a53d08)
- [コマンドラインからググれてもいいと思ったので作った - Qiita](http://qiita.com/ieee0824/items/13435fc6de5f22cdb2f4)
- [tsaka1's blog: ノイマンの自然数(非負整数)生成プログラムについて](http://tsaka1.blogspot.jp/2016/05/blog-post.html)
- [電子書籍の未来を握るのはインディー系 - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2016/04/52669/)
    - [英米のEブックを支えている読者は誰？ « マガジン航[kɔː]](http://magazine-k.jp/2016/05/24/beyond-cool-japan-06/)
    - [電子書籍は漫画家の希望となるか？｜佐藤秀峰｜note](https://note.mu/shuho_sato/n/n736593947e6c)
- [「フェアユースでも使用料を払え」というソニーミュージックの横暴と、それを許すYouTubeのコンテンツID – P2Pとかその辺のお話R](http://p2ptk.org/copyright/350)
- [textlintで日本語の文章をチェックする | Web Scratch](http://efcl.info/2015/09/10/introduce-textlint/)
- [スノーデン氏、グーグルの新メッセージアプリ「Allo」を「危険」と批判 - CNET Japan](http://japan.cnet.com/news/service/35083036/)
- [WPAD Name Collision Vulnerability | US-CERT](https://www.us-cert.gov/ncas/alerts/TA16-144A)
    - [JVNTA#91048063: WPAD と名前衝突の問題](http://jvn.jp/ta/JVNTA91048063/)
- [LinkedInからの流出情報、自分の被害の有無を確認可能に - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1605/25/news073.html)
    - [Have I been pwned? Check if your email has been compromised in a data breach](https://haveibeenpwned.com/)
- [SETI Gets an Upgrade | Space | Air & Space Magazine](http://www.airspacemag.com/space/new-seti-search-180959126/?is_pocket=1)

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%80%90%E4%B8%AD%E6%9D%B1%E5%A4%A7%E6%B7%B7%E8%BF%B7%E3%82%92%E8%A7%A3%E3%81%8F%E3%80%91-%E3%82%B5%E3%82%A4%E3%82%AF%E3%82%B9-%E3%83%94%E3%82%B3%E5%8D%94%E5%AE%9A-%E7%99%BE%E5%B9%B4%E3%81%AE%E5%91%AA%E7%B8%9B-%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8/dp/4106037866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4106037866"><img src="https://images-fe.ssl-images-amazon.com/images/I/51QsC2WBr5L._SL160_.jpg" width="106" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%80%90%E4%B8%AD%E6%9D%B1%E5%A4%A7%E6%B7%B7%E8%BF%B7%E3%82%92%E8%A7%A3%E3%81%8F%E3%80%91-%E3%82%B5%E3%82%A4%E3%82%AF%E3%82%B9-%E3%83%94%E3%82%B3%E5%8D%94%E5%AE%9A-%E7%99%BE%E5%B9%B4%E3%81%AE%E5%91%AA%E7%B8%9B-%E6%96%B0%E6%BD%AE%E9%81%B8%E6%9B%B8/dp/4106037866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4106037866">【中東大混迷を解く】 サイクス=ピコ協定 百年の呪縛 (新潮選書)</a></dt>
	<dd>池内 恵</dd>
    <dd>新潮社 2016-05-27</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4106037866, EAN: 9784106037863</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">欧州および中東の近代および現代を「サイクス=ピコ協定」を特異点として網羅的に解説していいる。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-07-02">2016-07-02</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
