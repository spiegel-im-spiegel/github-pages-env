+++
title = "誰がプライバシーを支配するのか"
date = "2018-04-01T14:14:00+09:00"
description = "キーワードは「情報へのアクセス性」で，それが公開されているか否かに関わらず，ユーザがアクセスをコントロールできるかどうかが重要になってくる。"
image = "/images/attention/kitten.jpg"
tags        = [ "privacy", "code", "market", "censorship", "grigori" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

こういう日に記事をアップしたくなかったが，ムカついてしょうがないので書いておく。

- [ジョブズ時代から「個人情報で商売しない」を掲げてきたAppleのブレなさ - ITmedia PC USER](http://www.itmedia.co.jp/pcuser/articles/1804/01/news021.html)

アホか！
Apple は基本的にハードウェア・メーカである。
ハードウェア・メーカが個人のプライバシーで商売することはない。
そんなの昔あった「もし Microsoft が自動車を作ったら」みたいなホラーである。
それは Apple に限らず IBM や他の老舗企業でも同じと言える[^ai1]。

[^ai1]: IBM は本格的に AI 技術に舵を切ってるし今後は分からないけど，それすらない Apple ではありえないだろう。 Apple は「個人情報で商売しない」のではなくできないのである。[自前のクラウドすら持てない](http://jp.techcrunch.com/2018/02/28/2018-02-27-apple-now-relies-on-google-cloud-platform-and-amazon-s3-for-icloud-data/ "AppleはiCloudのデータをGoogle CloudとAmazon S3に保存している  |  TechCrunch Japan")のでは競争に勝てない。

## 公開されたプライバシー

大昔はプライバシーというのは公衆から隔離された空間に存在し，そこに踏み込まれないための「方便」として成り立っていた。
しかしネットに「[つながりっぱなし]({{< ref "/remark/2016/05/its-complicated.md" >}} "『つながりっぱなしの日常を生きる』を読む")」の現代ではその意味を大きく変じている。

- [「ネットが蝕むプライバシー」を読む（再掲載）]({{< ref "/remark/2016/11/the-end-of-privacy.md" >}})

キーワードは「情報へのアクセス性」で，それが公開されているか否かに関わらず，ユーザがアクセスをコントロールできるかどうかが重要になってくる。
たとえば [GDPR (General Data Protection Regulation)](https://www.eugdpr.org/) 絡みで再燃している「忘れられる権利」も情報へのアクセス性の問題に帰着する。

## ユーザを国に売り渡しした Apple

「個人情報で商売しない」と言う Apple は iCloud のデータについて中国人ユーザのデータを（中国在住か否かに関わりなく）中国のデータセンターに移管した。

- [Appleの中国iCloudデータ移管は海外ユーザーも対象になっていた | TechCrunch Japan](https://jp.techcrunch.com/2018/01/12/2018-01-11-apple-china-icloud-international-users/)

言うまでもなく中国は個人の人権が軽い。
そんな国に私的情報を置くリスクは計り知れない。

確かに Apple は「個人情報で商売しない」かもしれない。
しかし Apple にとって個人のプライバシーは非常に軽い。
だからこそ Apple はユーザを国に売り渡したのである。

もともとプライバシーの問題は「個人 vs. 国家」の枠組みで論じられていたもので（今もそうだけど），そういう基本的な部分で Apple はユーザを裏切っているわけだ。

実は（Apple に限らず）似たような事例は過去にいくつもあった。
分かりやすいのはメッセージング・サービスの通話情報を国に売り渡すケースである。
メッセージング・サービスのデータを自国内に置くことを強要するケースが相次ぎ，サービス・プロバイダの多くはそれを承諾した（そうしなければその国で商売できないから）。
メッセージング・アプリで end-to-end の暗号化が推奨されるのはこうした事態への対抗措置と言える[^tls1]。

[^tls1]: TLS のように中間者攻撃を食らったら E2E も意味ないけどね。 TLS が中間者攻撃に弱いのは CA の権威性に頼る X.509 信用モデルが既に破綻しているからである。セキュリティ企業はセキュリティ対策の名のもとにユーザの端末に中間者攻撃用の証明書を仕込む。いわんや犯罪者をや。

## プライバシーと敵対する市場

上述の Apple 等の話は「個人 vs. 国家」だが，ゼロ年代以降に追加された問題は「個人 vs. 市場」の関係である。
「スノーデン」以前はプライバシーに対して敵対的な態度をとる企業やサービスが多かったのだ。

- [Google: "Complete Privacy Does Not Exist" | The Smoking Gun](http://www.thesmokinggun.com/documents/internet/google-complete-privacy-does-not-exist)
- [「プライバシー保護に敵対的」 最低評価を受けたGoogle](https://cloud.watch.impress.co.jp/epw/cda/infostand/2007/06/18/10531.html)

これは「スノーデン以後」も尾を引いており，当時からある企業やサービスで明確な答えを示せているところはない。

### 「検閲」と「過視」

情報のフィルタリングは誰でもやることだし，そうしなければ膨大なネットの情報を処理できない。
しかし，目的の如何に関わらず，ユーザがコントロールできないフィルタリングは全て「検閲」である。
そういう意味で Twitter や Facebook や Instagram といった流行りのタイムライン・サービスは「検閲」そのものであると言える。
また Google 検索の表示順位も「検閲」の一種である。
「検閲」を行えばそれを迂回しようとするものが現れるし，そいった連中を目当てに商売をするものも現れる。
素晴らしき生態系（笑）

ここでもキーワードは「情報へのアクセス性」である。

「検閲」と「過視[^vis1]」はコインの裏表の関係と言える。
これを利用して「情報の非対称性」あるいは「情報格差」を作り出し商売にしているのが Facebook などのタイムライン・サービスである。

[^vis1]: 「過視化」は「可視化（visualization）」よりもう少し強い意味になる。英語ではなんて言えばいいのだろう。この造語を初めて目にしたのは『[不過視なものの世界](https://www.amazon.co.jp/exec/obidos/ASIN/4022575360/baldandersinf-22/)』である。現在私が唯一手元に残している東浩紀さんの著書でインタビュー集になっている。

「スノーデン以後」の Facebook がユーザのプライバシー保護に取り組むようになったのは確かである。
しかしそれは歪な形で実装されている。

たとえば Facebook の設定には「プライバシー」という項目があるが， Facebook と連携するアプリのプライバシー設定はここではなく「アプリ」の項目にある。
このように Facebook の設定は一貫性のない奇々怪々な構成になっていて，しかもしょっちゅう変わるためユーザは定期的に設定を監視しなければならない。
つい先日も[設定が変わって](https://kaztaira.wordpress.com/2018/03/31/%e7%b1%b3%ef%bd%85%ef%bd%95%e3%83%97%e3%83%a9%e3%82%a4%e3%83%90%e3%82%b7%e3%83%bc%e5%8d%94%e5%ae%9a%e3%82%92%e7%84%a1%e5%8a%b9%e3%81%ab%e3%81%97%e3%81%9f%e7%94%b7%e3%80%8c%e3%83%95%e3%82%a7%e3%82%a4/ "米ＥＵプライバシー協定を無効にした男「フェイスブック流出問題は防げた」 | 新聞紙学的")いて，往生させられたよ `orz`

こうした Facebook の異様さには色々と憶測を立てられるかもしれないが，そうした憶測を全部スルーしたとしても，その異様さこそが「情報の非対称性」を生み出していることは断言できるだろう。

### AI によるオープンソース・インテリジェンス

「名寄せ」はシステム内の複数アカウントを（主に名前をキーに）統合する処理のことを指すが，ネットに於いては複数サービスを横断して情報を結合することを指す。
Campbridge Analytica の件は何故か Facebook だけが悪者としてやり玉に上げられているが，実際にこの問題の本質は異なるサービス・セグメントを横断する「名寄せ」にある。

名寄せのキーとなるものは名前に限らない。
画像認識技術が向上した現在では顔写真で名寄せできるし他の生体情報も使えるかもしれない。
さらに AI 技術が向上すればネット上の「振る舞い」のパターンからデータを繋げていくことも難しくなくなるだろう。
おそらく Google や Amazon や IBM といった AI 技術で一歩先んじている企業は，これからそういう方面に AI 技術の資源を振り向けていく筈だ。
これからのオープンソース・インテリジェンスは AI にお任せ♥ である。

日本でも与信審査に AI を活用する事例が増えているが，やろうと思えばそれにネットの情報をいくらでも名寄せすることができる。
Facebook や Twitter でうっかり不穏なことを書いて反社ユーザとして与信でハネられるなんてなことも技術的にはありえなくもない（またそういう印象を意図的に AI に植え付けるということもできるかもしれない）。

しかし，仮にそういうことが現実にあったとしても，これらの名寄せ情報が公になることはないだろう。
名寄せという行為そのものが「情報の非対称性」を増幅させるもので，そこが商売としての旨味だからだ。

## VRM に関する妄想

プライバシーの根源が「情報へのアクセス性」にあるのなら，それをコントロールする主体はあくまでも個人ユーザでなければならない。
しかし昨今のサービスは全て企業による CRM (Customer Relationship Management; 顧客関係管理) で語られている。
つまりユーザの情報をコントロールするのは企業・サービス側だと言っているのだ。

故に個人が武装するためにはどうしても VRM (Vendor Relationship Management; 企業関係管理) が必要となる。

VRM という言葉を初めて目にしたのは『[インテンション・エコノミー](https://www.amazon.co.jp/exec/obidos/ASIN/B00DIM6BE6/baldandersinf-22/)』が最初だが，あまり需要がなさそうである。
しかし AI が VRM を行うアシスタントとして機能するのならもう少しマシになるんじゃないだろうか。
それこそ『[万物理論](https://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/)』に出てくる「シジフォス」のように（笑）

（Amazon Echo や Google Home と何が違うんだと言われそうだが，決定的に違う。何故なら Amazon Echo や Google Home は CRM として機能するからだ。私たち個人ユーザは Amazon Echo や Google Home よりも下位の存在なのである）

## ブックマーク

- [善悪の葛藤]({{< ref "/remark/2018/03/name-identification.md" >}})

## 参考図書

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B0125TZSZ0/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/616sjle5ITL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B0125TZSZ0/baldandersinf-22/">つながりっぱなしの日常を生きる：ソーシャルメディアが若者にもたらしたもの</a></dt><dd>ダナ・ボイド 野中 モモ </dd><dd>草思社 2014-10-09</dd><dd>評価<abbr class="rating" title="5"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B0141TUJHY/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B0141TUJHY.09._SCTHUMBZZZ_.jpg"  alt="角川インターネット講座５　ネットコミュニティの設計と力　つながる私たちの時代<角川インターネット講座> (角川学芸出版全集)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01CZK0B2Y/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01CZK0B2Y.09._SCTHUMBZZZ_.jpg"  alt="これからの世界をつくる仲間たちへ"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01B1CKZQO/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B01B1CKZQO.09._SCTHUMBZZZ_.jpg"  alt="ForbesJapan (フォーブスジャパン) 2016年 03月号 [雑誌]"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B010LYGB34/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B010LYGB34.09._SCTHUMBZZZ_.jpg"  alt="「炎上」と「拡散」の考現学"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B0191AIN6W/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B0191AIN6W.09._SCTHUMBZZZ_.jpg"  alt="ぼくらの仮説が世界をつくる"  /></a> </p>
<p class="description">読むのに1年半以上かかってしまった。ネット，特に SNS 上で自身のアイデンティティやプライバシーを保つにはどうすればいいか。豊富な事例を交えて考察する。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-05-10">2016-05-10</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/41UdjkE4OpL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4150504598/baldandersinf-22/">フィルターバブル──インターネットが隠していること (ハヤカワ文庫NF)</a></dt><dd>イーライ・パリサー 井口耕二 </dd><dd>早川書房 2016-03-09</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4569762468/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4569762468.09._SCTHUMBZZZ_.jpg"  alt="インターネット的 (PHP文庫)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4140912073/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4140912073.09._SCTHUMBZZZ_.jpg"  alt="ウェブ社会のゆくえ―<多孔化>した現実のなかで (NHKブックス No.1207)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4122033985/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4122033985.09._SCTHUMBZZZ_.jpg"  alt="情報の文明学 (中公文庫)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4480062858/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4480062858.09._SCTHUMBZZZ_.jpg"  alt="ウェブ進化論 本当の大変化はこれから始まる (ちくま新書)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4152096098/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4152096098.09._SCTHUMBZZZ_.jpg"  alt="五〇億年の孤独:宇宙に生命を探す天文学者たち"  /></a> </p>
<p class="description">ネットにおいて私たちは自由ではなく，それと知らず「フィルターバブル」に捕らわれている。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-05-07">2016-05-07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00DIM6BE6/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/519%2BkIHb71L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B00DIM6BE6/baldandersinf-22/">インテンション・エコノミー～顧客が支配する経済</a></dt><dd>Doc Searls 栗原潔 </dd><dd>翔泳社 2013-03-14</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00W535LOU/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00W535LOU.09._SCTHUMBZZZ_.jpg"  alt="HARD THINGS　答えがない難問と困難にきみはどう立ち向かうか"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00TXZXE5Q/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00TXZXE5Q.09._SCTHUMBZZZ_.jpg"  alt="パーソナルデータの衝撃"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00T3YFXJM/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00T3YFXJM.09._SCTHUMBZZZ_.jpg"  alt="UX侍 スマホアプリでユーザーが使いやすいデザインとは (impress Digital Books)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00J9ZGYQQ/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00J9ZGYQQ.09._SCTHUMBZZZ_.jpg"  alt="位置情報ビッグデータ (NextPublishing)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/B00LTCR0IS/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/B00LTCR0IS.09._SCTHUMBZZZ_.jpg"  alt="起業のエクイティ・ファイナンス"  /></a> </p>
<p class="description" >時代はソーシャル CRM から VRM へ。<a href='https://baldanders.info/spiegel/log2/000794.shtml'>俺達がインターネットだ！</a></p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2015-04-26">2015/04/26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51J3DEJJ1TL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4488711022/baldandersinf-22/">万物理論 (創元SF文庫)</a></dt><dd>グレッグ・イーガン 山岸 真 </dd><dd>東京創元社 2004-10-28</dd><dd>評価<abbr class="rating" title="4"><img src="https://images-fe.ssl-images-amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4488711014/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4488711014.09._SCTHUMBZZZ_.jpg"  alt="宇宙消失 (創元SF文庫)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4150115311/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4150115311.09._SCTHUMBZZZ_.jpg"  alt="ディアスポラ (ハヤカワ文庫 SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4150113378/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4150113378.09._SCTHUMBZZZ_.jpg"  alt="祈りの海 (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4150112908/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4150112908.09._SCTHUMBZZZ_.jpg"  alt="順列都市〈下〉 (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4150118264/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4150118264.09._SCTHUMBZZZ_.jpg"  alt="プランク・ダイヴ (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4150112894/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4150112894.09._SCTHUMBZZZ_.jpg"  alt="順列都市〈上〉 (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/415011594X/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/415011594X.09._SCTHUMBZZZ_.jpg"  alt="ひとりっ子 (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4122056810/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4122056810.09._SCTHUMBZZZ_.jpg"  alt="六本指のゴルトベルク (中公文庫)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/415011451X/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/415011451X.09._SCTHUMBZZZ_.jpg"  alt="しあわせの理由 (ハヤカワ文庫SF)"  /></a> <a href="https://www.amazon.co.jp/exec/obidos/ASIN/4582760465/baldandersinf-22/" target="_top"><img src="https://images-fe.ssl-images-amazon.com/images/P/4582760465.09._SCTHUMBZZZ_.jpg"  alt="文字逍遥 (平凡社ライブラリー)"  /></a> </p>
<p class="description">グレッグ・イーガンの名作。これも singularity を巡る物語だな。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-18">2017-09-18</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
