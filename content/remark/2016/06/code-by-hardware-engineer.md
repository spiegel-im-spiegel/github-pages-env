+++
date = "2016-06-04T09:08:09+09:00"
description = "この件では，若いころのほろ苦い思い出がある。"
draft = false
tags = ["engineering", "programming"]
title = "ハード屋が書く C ソースコードが凄まじかった思い出（再掲載）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

これは[ちょうど1年前に Medium で書いた記事](https://medium.com/@spiegel/-1ca9e4895f4c)の再掲載である。
今は Medium を全く利用しないので，昔書いたもので（私が）面白いと思った記事は少しずつこっちに移転しようかな。

----

{{< fig-quote title="トヨタの車のソースコードはスパゲッティコード山盛り？ - YAMDAS現更新履歴" link="http://d.hatena.ne.jp/yomoyomo/20150604/toyota" >}}
<q>この記事で面白いのは、Michael Barr が20ヶ月以上にわたりトヨタ車で使われているソースコードを、Philip Koopman カーネギーメロン大学教授がトヨタのエンジニアリングの安全プロセスを精査した話で、両者ともトヨタのソフトウェアがスパゲッティコード山盛りなことを証言している。
<br>
トヨタの生産方式はアジャイル方面においてソフトウェア開発手法に多大な影響を与えている。ところでそのトヨタが開発するソフトウェアの品質はどうなんだろう、というのは多くの人の頭に浮かぶ疑問だろう。組み込みソフトウェアのエキスパートによると、ものすごく複雑で、複雑すぎてテストもメンテもできない関数がたくさんあるとか、グローバル変数が1万個以上あるとかなかなか壮絶らしい……。マジかよ。</q>
{{< /fig-quote >}}

この件では，若いころのほろ苦い思い出がある。

私は若いころは「システムハウス」と呼ばれる類の会社にいたのだが，そこではハードとソフトの両面で開発を進められるのが「売り」だった。
[私はハードは壊滅的にダメ](https://baldanders.info/spiegel/log2/000529.shtml "私はこうしてプログラミングを覚えた — Baldanders.info")なのでソフトウェア担当。
当時のハード屋は自分が組んだ回路の実証のために自分でもプログラムを組んで動作確認する。
私たちソフト屋はそのコードをもらって実際のコードを書くわけだが，この実証コードが凄まじかった。

まず変数は全てグローバル変数。
スコープとかカプセル化なんて知るか！ という気概が感じられる。
そして関数は果てしなく長い main 関数のみ。
無間地獄のネスト。
goto 文で飛びまくり。
なのに異常系の記述は皆無。
世に聞く「スパゲッティ・コード」とはこのようなものなのかと感嘆したよ。

一番凄かったのは，とあるチップを使った30次のバンドパスフィルタを組むのに「サンプルがあるから簡単でしょ」と言われてサンプルを見たら世にも悍ましいコードで，解析するだけで半月もかかってしまった。

まぁ，ハード屋がこういうコードを書くのは理由があって，変数を記述するときはメモリ上のマッピングをそのまま置き換えようとするし，ロジックも基本的にマシン語のインストラクションをそのまま C に置き換えようとするから「関数」という概念がそもそもないことが多い[^a]。

[^a]: ハード屋は call インストラクションを「特殊なジャンプ」程度にしか認識していない（まぁ確かにそうなのだが）。あるプロジェクトで見せてもらったアセンブラコードでは call で積んだスタックをいじって戻り先アドレスを変えて return する記述が頻出していた。これはアセンブラ・コードにパッチを当てる際の基本テクニックらしい。大昔の話だよ（笑）

おかげで私は「他人のコードを読む」ことがすんごい得意になってしまった。
アレに比べればソフト屋の書くコードなんて絵本を読むように分かりやすい（笑） でも，こんなしょうもない特技でも後年ちゃんと役に立ってるんだから世の中というのは分からないものである。

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%B3%E9%96%8B%E7%99%BA%E3%81%AE%E7%8F%BE%E5%A0%B4-%E3%82%AB%E3%83%B3%E3%83%90%E3%83%B3%E3%81%AB%E3%82%88%E3%82%8B%E5%A4%A7%E8%A6%8F%E6%A8%A1%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%81%8B%E5%96%B6-%EF%BC%A8%EF%BD%85%EF%BD%8E%EF%BD%92%EF%BD%89%EF%BD%8B%EF%BC%AB%EF%BD%8E%EF%BD%89%EF%BD%82%EF%BD%85%EF%BD%92%EF%BD%87-ebook/dp/B01IGW5IIW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01IGW5IIW"><img src="https://images-fe.ssl-images-amazon.com/images/I/51gC8Tmq1kL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%AA%E3%83%BC%E3%83%B3%E9%96%8B%E7%99%BA%E3%81%AE%E7%8F%BE%E5%A0%B4-%E3%82%AB%E3%83%B3%E3%83%90%E3%83%B3%E3%81%AB%E3%82%88%E3%82%8B%E5%A4%A7%E8%A6%8F%E6%A8%A1%E3%83%97%E3%83%AD%E3%82%B8%E3%82%A7%E3%82%AF%E3%83%88%E3%81%AE%E9%81%8B%E5%96%B6-%EF%BC%A8%EF%BD%85%EF%BD%8E%EF%BD%92%EF%BD%89%EF%BD%8B%EF%BC%AB%EF%BD%8E%EF%BD%89%EF%BD%82%EF%BD%85%EF%BD%92%EF%BD%87-ebook/dp/B01IGW5IIW?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B01IGW5IIW">リーン開発の現場 カンバンによる大規模プロジェクトの運営</a></dt>
	<dd>ＨｅｎｒｉｋＫｎｉｂｅｒｇ, 角谷信太郎</dd>
	<dd>市谷聡啓 (翻訳), 藤原大 (翻訳)</dd>
    <dd>オーム社 2013-10-25 (Release 2017-07-15)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B01IGW5IIW</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">私はこれで勉強しました。もう一回読み直すかな。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-24">2018-12-24</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
