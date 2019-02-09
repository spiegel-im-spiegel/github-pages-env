+++
date = "2016-06-05T16:28:27+09:00"
update = "2016-06-05T18:14:30+09:00"
description = "古いパソコンに Ubuntu を入れようと思ったが... / 「「IoTセキュリティガイドライン」（案）に関する意見募集」 / 『BOOM TOWN』単行本未収録話が Kindle に！ / 「Googleアカウントがアップデート」 / その他の気になる記事"
draft = false
tags = ["linux", "ubuntu", "security", "risk", "politics", "e-book", "comic", "google"]
title = "週末スペシャル： 古いパソコンに Ubuntu を入れようと思ったが..."

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

1. [古いパソコンに Ubuntu を入れようと思ったが...]({{< relref "#ubnt" >}})
1. [「「IoTセキュリティガイドライン」（案）に関する意見募集」]({{< relref "#iot" >}})
1. [『BOOM TOWN』単行本未収録話が Kindle に！]({{< relref "#bt" >}})
1. [「Googleアカウントがアップデート」]({{< relref "#dm" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 古いパソコンに Ubuntu を入れようと思ったが...{#ubnt}

以前にも書いたが，2020年までに自宅のパソコン環境を Linux ベースに換えようと思っていて，試しにもう使ってない古いパソコンに [Ubuntu](http://www.ubuntu.com/ "The leading OS for PC, tablet, phone and cloud | Ubuntu") を入れようと思ったのだが...

- [Ubuntu Japanese Team](https://www.ubuntulinux.jp/home)
    - [Ubuntuの日本語環境 | Ubuntu Japanese Team](https://www.ubuntulinux.jp/japanese)

まず「日本語 Remix」版のインストールイメージは 64bit 用しかないらしい。
まぁそれはいいのだが，実は Ubuntu を入れようとしたパソコンは DVD ドライブが壊れていて DVD からインストール出来ないのであった。
このことをすっかり忘れていて，[本家](http://www.ubuntu.com/ "The leading OS for PC, tablet, phone and cloud | Ubuntu")から 32bit のイメージを取ってきて DVD に焼いて「さて」となったところでようやく思い出した。

**がぁ！ ディスク1枚無駄にしちまったじゃねーか。**

USB からインストールする手もあるらしいのだが，あいにく私は USB メモリを持ってない[^a]。
USB メモリを買ってまでやる気はないので，この件はしばらくお蔵入り。

[^a]: 職場では USB メモリの持ち込みも持ち出しも NG なのが普通だし（それなら企業アカウントで [Box](https://www.box.com/) 等を使うほうがマシ），自宅でわざわざ USB メモリを使う局面はないので（[YubiKey](https://www.yubico.com/products/yubikey-hardware/) は興味あるけど）全く用がない。

しかし DVD ドライブが壊れてるということは，今の環境が壊れたら復旧できないということなので，このパソコンはもう[廃棄](http://www.city.hiroshima.lg.jp/www/contents/1111215356400/index.html)だな。
というわけで，とりあえずパソコンから HDD とメモリを抜き取った。

{{< fig-img src="https://photo.baldanders.info/flickr/image/27191657840_m.jpg" title="使ってない古い PC から（PC の処分を前提に） HDD を抜き取った" link="https://photo.baldanders.info/flickr/27191657840/" >}}

この写真で下のほうの HDD は IDE/ATAPI だ！ なんせ[2005年に買った「牛丼パソコン」](https://baldanders.info/spiegel/log/200510.html#d18_t1)だからな。
幸運なことに，私は今まで自宅のマシンで HDD のトラブルに遭ったことは一度もない（仕事では何度かある）。
まぁでも，今時 ATAPI が繋がるマシンなんてないだろうから，このまま自宅の隅っこで朽ち果てていくんだろうな，これ。

## 「「IoTセキュリティガイドライン」（案）に関する意見募集」{#iot}

- [総務省｜「IoTセキュリティガイドライン」（案）に関する意見募集](http://www.soumu.go.jp/menu_news/s-news/01ryutsu03_02000107.html)

締切は6月14日。
意見のある方はお早めにどうぞ。

総務省が IoT なるバズワードを使うのもどうかと思うが，ガイドラインでは，ITU 勧告を引いて， IoT をこう定義しているようだ。

{{< fig-quote title="IoTセキュリティガイドライン ver 1.0 （案）" link="http://www.soumu.go.jp/main_content/000421617.pdf" >}}
<q>情報社会のために、既存もしくは開発中の相互運用可能な情報通信技術により、物理的もしくは仮想的なモノを接続し、高度なサービスを実現するグローバルインフラ</q>
{{< /fig-quote >}}

なんともボンヤリした説明である。
物理的か仮想的かを問わないなら今までだって IoT なわけで，今更こんなセキュリティ・ガイドラインが出てくるということ自体，国も企業もいかにネットのセキュリティに対して無頓着だったかという証でもある。

ネットワーク・セキュリティは「予防」から（事後の対処を睨んだ）「監視」へと比重が移りつつある。
理由は簡単で，ネットワークに接続する機器が多すぎて管理しきれなくなっているためである。
また，昨年の[日本年金機構の例](https://baldanders.info/spiegel/log2/000850.shtml "人を排除するシステムは人に殺される — Baldanders.info")を見れば分かるように，いかにルールや罰則を決めても逸脱はなくならないし人為的なミスは避けようがない。

企業やそれなりの組織ならセキュリティ対策を「投資」とみなして PDCA サイクルを回すこともできるけど[^b]，一般家庭では無理筋な話である。
ならばこのガイドライン案のターゲットは企業・組織に向けたもので，全く一般向けではないと考えることもできる[^bb]。
私たちは私たち自身でリスクに立ち向かわなければならない。
国は助けてくれないし，セキュリティ企業の多くは貴方から「みかじめ料」を巻き上げたいだけなのだ。

[^b]: 件のガイドライン案では PDCA サイクルを回すことで内部不正やミスをなくそうとしているようだが， PDCA サイクルは「投資とその評価」に対しては有効だが，セキュリティ対策を「コスト」とみなしているうちはうまく回らない。まぁ，総務省としては「もっと僕達セキュリティ・ゴロにお金を落としてYO」と言いたいのかもしれないが。
[^bb]: それなら経産省の管轄だよね（笑）

一般家庭については

{{< fig-quote title="エフセキュアブログ : スマートホームの安全を保つ方法" link="http://blog.f-secure.jp/archives/50744439.html" >}}
<q>プライバシーやセキュリティについて、非常に心配に思うのであれば、こうしたガジェットを買ったり使ったりしないことが、安全にいるための唯一の方法である。</q>
{{< /fig-quote >}}

が最善である。
それでもリスクを取って利便性を享受したいというのであれば

{{< fig-quote title="エフセキュアブログ : スマートホームの安全を保つ方法" link="http://blog.f-secure.jp/archives/50744439.html" >}}
<q>攻撃ポイントを限定すること。必要になることがないと分かっているデバイスは、導入しない。もはや必要がなく使わないデバイスは、すべてシャットダウンして撤去するとよい。最上位機種の洗濯機を購入したところ、Wi-Fi経由で接続可能なことに気付いたのなら、接続する前に本当にその必要性があるのかを検討する。実際にはオンライン機能をまったく使わないことに気付いたのなら、デバイスをネットワークから切り離すること。</q>
{{< /fig-quote >}}

をお薦めする。
そう考えると，家庭内無線 LAN ルータで安易に DHCP を有効にするのは考えものかもしれない（もしくは DHCP で接続してくる機器はネットに繋がないようにするとか）。

それにしても，今時コメントをもらうのに PDF 配布でメールで意見を受け付けるとか，どんだけアナクロなんだか。
これで IoT でセキュリティってんだから，ヘソで茶が沸いちゃうぜ。
意見が欲しいなら repository を公開して pull request を受け付けるようにしろよ。

## 『BOOM TOWN』単行本未収録話が Kindle に！{#bt}

今更な話で申し訳ないんだけど。

今朝気がついたのだが，内田美奈子さんの『BOOM TOWN』で[単行本未収録の第30話を Kindle で売ってる](https://www.amazon.co.jp/exec/obidos/ASIN/B00NWQI4N4/baldandersinf-22/ "Amazon.co.jp: BOOM TOWN　TRIP.30 電子書籍: 内田 美奈子: Kindleストア")じゃないか！

早速購入。
で，奥付を見たら今回のこれは「[マンガ図書館Z]」の仕事らしい。
偉いぞ！ いい仕事してますねぇ（中島某の声で）。

- [赤松健、これが究極の一手。・・・なぜ我々は「電子書籍版YouTube」を目指すか - （株）Ｊコミックテラスの中の人](http://d.hatena.ne.jp/KenAkamatsu/20160113/p1)
- [「マンガ図書館Z」がプチ炎上中！？公式作家として支持・静観表明 - 漫画原作者 猪原賽BLOG](http://iharadaisuke.hatenablog.com/entry/2016/01/25/121758)
- [絶版マンガを無料で読める「マンガ図書館Z」、iOS/Android専用スマホアプリをリリース -INTERNET Watch Watch](http://internet.watch.impress.co.jp/docs/news/755216.html)

いや，私はこういうの全然アリだと思うのよ。
マンガ雑誌で2,3回で打ち切られたような作品でも面白いと思う読者がいれば Kindle や他のEブック・アプリでどんどん出すべきだと思う。
個人的には[西岡さち](http://mariyaribbon.zashiki.com/)さんの作品で1話きりの「ふんわか×あぶのーまる」とか Kindle で100円程度で出せるのなら買うよ[^c]。
作者も出版社もあらゆるチャネルを利用して読者との接触機会を積極的に作るべき。
「紙かデジタルか」なんてまるでくだらない議論だよな。

[^c]: 現在連載中の「ざしきわらしと僕」も単行本化を期待しています。

おそらく，日本のマンガにおいてEブックの方向性は2つあって，ひとつは「[マンガ図書館Z]」のようなアーカイブ化（＋マネタイズ）の方向で[^cc]，もうひとつは作品の「アプリ化」の方向。

[^cc]: 歴史の浅いマンガでは公有化されているものが少ないし作家自身も存命されてることが多いため「[青空文庫](http://www.aozora.gr.jp/)」よりは「[マンガ図書館Z]」のほうが馴染むかもしれない。

{{< fig-quote title="電子書籍の未来を握るのはインディー系 - WirelessWire News" link="https://wirelesswire.jp/2016/04/52669/" >}}
<q>ゲームがコンソールからアプリに移動し、CDがストリーミングに移動したように、活字もアプリに移行します。消費者は、大半の雑誌や本を、かつてよりも安くて置き場所にも困らないアプリで消費し、本当に気に入った本のみを、限定本や豪華本として部屋に飾るようになっていきます。その移行スピードは思った以上の速さで進んでいて、止めることはできないということです。</q>
{{< /fig-quote >}}

マンガはある意味で「活字」よりもアプリ化しやすい（実際に Web マンガとかは随分前からある）。
あとはマンガが「紙」という frame からいかにして逃れるかというところだろう。

ええつと。
なんでこんな話してたんだっけ。
あぁ，そうそう。

『BOOM TOWN』のことを思い出したのは以下の記事を見たから。

- [グーグル、人工知能が作り出したメロディを初披露 - CNET Japan](http://japan.cnet.com/news/service/35083595/)

これを読んで『BOOM TOWN』に出てくる{{< ruby "かずりしせい" >}}数理紫生{{< /ruby >}}という XYZ-People を思い出したのだった。
もう何度目かの流行になっている人工知能や仮想現実（あるいは Metaverse）のアイデアの殆どは20世紀までに出尽くしている。
『BOOM TOWN』はコンピュータ・エンジニアなら必読書だと思うぞ（あとは『スノウ・クラッシュ』とか）。

ちなみに『BOOM TOWN』のスピンオフ作品とも言える『サーキットワンダラーズ』が[「画楽.mag」 VOL.5](https://www.amazon.co.jp/exec/obidos/ASIN/4834284387/baldandersinf-22/) から連載されてるらしい。
（追記：さっき教えてもらったけど『サーキットワンダラーズ』は未完で終わってるらしい。残念）

## 「Googleアカウントがアップデート」{#dm}

- [Googleアカウントがアップデート―AndroidだけでなくiOSデバイスも探してロックできる | TechCrunch Japan](http://jp.techcrunch.com/2016/06/02/20160601googles-my-account-will-now-help-both-ios-and-android-users-find-their-lost-phones/)

Android にはもともと「デバイスマネージャー」と呼ばれる機能がある。

- [Android デバイス マネージャー](https://www.google.com/android/devicemanager)
- [Androidデバイスマネージャー - Google Play の Android アプリ](https://play.google.com/store/apps/details?id=com.google.android.apps.adm)

この機能が「[アカウント情報](https://myaccount.google.com/)」のページからも利用できるようになったということらしい。
併せて iOS 機器についても端末の捜索等ができるようになったということのようだ。
iOS 機器の遠隔操作は iOS の標準機能でやったほうがいいと思うが，複数の端末を一度に捜索するなら役に立つかもしれない。

- [iPhone、iPad、または iPod touch を紛失したり盗まれたりした場合 - Apple サポート](https://support.apple.com/ja-jp/HT201472)
- [iCloud: 紛失モードの使用](https://support.apple.com/kb/PH2700?locale=ja_JP&viewlocale=ja_JP)

なお，端末を紛失したときのことを考えて定期的に「避難訓練」をすることをお薦めする。
大袈裟なのじゃなくても適当に距離のあるところに端末を置いておいてアラームを鳴らしてみればいい。
いざというときに未知の道具を使うと焦るものだが，実際に使って手順を確認しておけば多少は気が楽になるというものである。
そうそう。
端末のロックを試すときは[ロック解除の仕方](http://androidlover.net/googleapps/android-device-manager/android-forget-lock-screen-security.html)を事前に確認しておくこと。

## その他の気になる記事{#other}

- [IPAテクニカルウォッチ「増加するインターネット接続機器の不適切な情報公開とその対策」：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/technicalwatch/20160531.html)
- [News ＆ Trend - 「マイナンバーを記録したパソコンは修理できない」、PC各社の修理規定が波紋：ITpro](http://itpro.nikkeibp.co.jp/atcl/column/14/346926/052900539/?n_cid=nbpitp_fbed)
- [マイクロソフト、ブロックチェーン技術によるID管理に向け提携を発表 - ZDNet Japan](http://japan.zdnet.com/article/35083525/)
- [中国がインターネットユーザー全員の実名による登録を義務付けへ | TechCrunch Japan](http://jp.techcrunch.com/2016/06/01/20160601china-attempts-to-reinforce-real-name-registration-for-internet-users/)
- [仮想通貨、「興味はあるが購入しない」7割 --バード調査 - ZDNet Japan](http://japan.zdnet.com/article/35083495/)
- [ファンサブサイト運営者は禁錮刑が相当？ – P2Pとかその辺のお話R](http://p2ptk.org/copyright/405)
- [インターネットの歴史を変えた『バックアップ』 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/400)
- [「悪魔のように賢い」とGoogleのエンジニアが舌を巻く「悪意あるハードウェア」が登場 - GIGAZINE](http://gigazine.net/news/20160603-undetectable-backdoor-into-chip-attack/)
- [Big Sky :: peco みたいだけど peco と違うコマンドラインセレクタ cho を作った。](http://mattn.kaoriya.net/software/lang/go/20160603011620.htm)

[マンガ図書館Z]: http://www.mangaz.com/ "マンガ図書館Z - 全巻無料で読み放題"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/BOOM-TOWN-TRIP-30-%E5%86%85%E7%94%B0-%E7%BE%8E%E5%A5%88%E5%AD%90-ebook/dp/B00NWQI4N4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00NWQI4N4"><img src="https://images-fe.ssl-images-amazon.com/images/I/51Ia%2B77IpiL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/BOOM-TOWN-TRIP-30-%E5%86%85%E7%94%B0-%E7%BE%8E%E5%A5%88%E5%AD%90-ebook/dp/B00NWQI4N4?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00NWQI4N4">BOOM TOWN　TRIP.30</a></dt>
	<dd>内田 美奈子</dd>
    <dd> 2014-09-26 (Release 2014-09-26)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00NWQI4N4</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">掲載誌「コミックガンマ」が休刊になって単行本収録できなかった<del>まるぼし</del>まぼろしの30話。これが Kindle で読めるとはいい時代になったものです。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-06-05">2016-06-05</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
