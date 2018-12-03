+++
date = "2016-07-17T08:26:43+09:00"
update = "2016-08-19T20:19:25+09:00"
description = "そもそもセキュリティ以前に「自分で面倒をみる」ことのできない人は自前で CMS を導入するのはやめたほうがいい。ペットを飼うのと同じ（笑）"
draft = false
tags = ["cms", "security", "management"]
title = "「自分で面倒見られる子」だけが CMS を導入しなさい"

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

私の場合，それまで使っていた[オープンソース版 Movable Type (MTOS)](https://www.movabletype.jp/opensource/ "Movable Type Open Source | MovableType.jp - CMSプラットフォーム Movable Type -") を捨ててローカルでサイト管理するようにした[^0]。
直接のきっかけは開発元が [MTOS](https://www.movabletype.jp/opensource/ "Movable Type Open Source | MovableType.jp - CMSプラットフォーム Movable Type -") のサポートを打ち切ったことだが，自前でサーバ・サイドの CMS (Content Management System) を導入・運用することが高コストになり，それを個人レベルで運用することに意義を見いだせなくなってきたというのが大きい。

[^0]: 経緯については本家サイトの「[ブログ移転準備中または Take the Hugo!](http://www.baldanders.info/spiegel/log2/000870.shtml)」を参照のこと。

というわけで以下の話題。
いつもは「セキュリティクラスタ」の話題はまるっと無視しているのだが，今回は気になる点があったので。

- [セキュリティクラスタ まとめのまとめ 2016年6月版：初心者でも気軽にWordPressサイト作成！ して大丈夫なの？ (1/3) - ＠IT](http://www.atmarkit.co.jp/ait/articles/1607/15/news015.html)

そもそもセキュリティ以前に「自分で面倒をみる」ことのできない人は自前でサーバ・サイド CMS を導入するのはやめたほうがいい。
ペットを飼うのと同じ[^1]（笑） 「初心者」かどうかはさしたる問題ではない。
誰だって最初は初心者だし，分からなければ訊くなり調べるなりすればいいからだ。

[^1]: でも，好きでペットを飼ってる人はそれをコストだと思わないだろうけど。コンピュータ・オタクやセキュリティ・オタクには変態が多いので，そういう「面倒」をわざわざ引き受けて楽しんでいるフシがある。仕事で使うツールを自分のフィールドでも使い倒すってのはあるけど。

プログラムは家電製品とは違う。
導入すればそれで終わり，とはいかない。
多くの人に使われる製品こそ常に不具合や脆弱性の改修が行われているし[^ss]，もちろん機能改善も行われる。
「プログラムは常にベータ版」という名言は一般の方も覚えておいたほうがいいだろう[^iot]。

[^ss]: 不具合や脆弱性の報告が全くない製品ってのは誰も使ってないか製品の提供者が怠慢かのどちらかである。
[^iot]: とはいえ IoT の時代になれば家電製品も「自分で面倒をみる」ことが必要になってくるかもしれない。本当に面倒な時代になったものである。

面倒を見る気はないが気軽にブログをやりたい，って程度なら SaaS (Software as a Service) 型のサービスがいくらでもある。
さすがに ameblo みたいなのはデザインが古すぎてアレだが， [WordPress](https://wordpress.com/) も [Movable Type](https://movabletype.net/) も SaaS を提供している。
最近なら [Medium](https://medium.com/) や 日本の [note](https://note.mu/) みたいに気軽にオン書きできるものもある。
[Tumblr](https://www.tumblr.com/) でブログサイトを運用している人は日本語圏にも結構いる[^t]。
最近では [esa.io](https://esa.io/ "esa.io - Expertise Sharing Archives for motivated teams.") のようにチームで使えるものが流行りらしい。
あるいは私のようにサーバ・サイドの CMS を捨てて [GitHub Pages](https://pages.github.com/ "GitHub Pages - Websites for you and your projects, hosted directly from your GitHub repository. Just edit, push, and your changes are live.") ＋ [Hugo](https://gohugo.io/ "Hugo :: A fast and modern static website engine") で静的なページとして運用すればセキュリティ上の煩わしさから解放される。
やり方なんていくらでもあるのだ。

[^t]: [Tumblr](https://www.tumblr.com/) は[ホワイトハウスの公式サイトがある](http://whitehouse.tumblr.com/ "The Official White House Tumblr")ことで有名である。

今でこそ「セキュリティは投資」とみなされ PDCA サイクルで回されるようになってきたが[^s]，ちょっと前まではセキュリティ管理は典型的なコストセンターだったのだ。
JTB の情報漏えい騒ぎに関する今回の事後の振る舞いには批判が多いが，結局のところ「セキュリティ」を経営レベルでどう捉えるかということに尽きる。
セキュリティ管理をコストとみなすうちは[昨年の日本年機構](http://www.baldanders.info/spiegel/log2/000850.shtml)や今回の JTB のような失敗は必ずまた繰り返される。

[^s]: そしてセキュリティへの投資につけ入る「輩」も増えてきた。

しかし企業や組織はそれでいいとしても個人レベルでそれを求めるのは酷な話である。
セキュリティと利便性はトレードオフにならない。
利便性を犠牲にしてセキュリティを強化してもユーザはただそこを迂回するだけだ（日本年金機構の事例はまさにこれと言える）。

WordPress のようなサーバ・サイド CMS がセキュリティ的に高コストに見えるのは，製品デザインとセキュリティ管理が齟齬を起こしていて，そこを運用で強引にカバーしようという駄目プロジェクトの典型みたいになっているからである。
あるいは「[奥が深い症候群](http://0xcc.net/misc/bad-knowhow.html "バッドノウハウと「奥が深い症候群」")」の末期症状と言うべきか。
SaaS ならそうした面倒をサービス・プロバイダに押しつけることが出来るという意味で次善の策になる。

もう一度繰り返そう。

「自分で面倒見られる子」だけが CMS を導入しなさい。

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00LBPGFJS/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51r6kpV26GL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00LBPGFJS/baldandersinf-22/">フルスクラッチから1日でCMSを作る シェルスクリプト高速開発手法入門 (アスキー書籍)</a></dt><dd>上田 隆一 後藤 大地 ＵＳＰ研究所 </dd><dd>KADOKAWA / アスキー・メディアワークス 2014-07-02</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00MPDUQQI/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00MPDUQQI.09._SCTHUMBZZZ_.jpg"  alt="サーバ／インフラエンジニア養成読本 ログ収集〜可視化編 [現場主導のデータ分析環境を構築！] (Software Design plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00ME9TTMA/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00ME9TTMA.09._SCTHUMBZZZ_.jpg"  alt="フロントエンドエンジニア養成読本［HTML ，CSS，JavaScriptの基本から現場で役立つ技術まで満載！］ (Software Design Plus)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00LHFOTF4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00LHFOTF4.09._SCTHUMBZZZ_.jpg"  alt="絵で見てわかるシステムパフォーマンスの仕組み"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00J4KDYV4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00J4KDYV4.09._SCTHUMBZZZ_.jpg"  alt="はじめてUNIXで仕事をする人が読む本 (アスキー書籍)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00MLUGZIS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00MLUGZIS.09._SCTHUMBZZZ_.jpg"  alt="すごいErlangゆかいに学ぼう！"  /></a> </p>
<p class="description" >既存の常識に凝り固まったソフトウェア・エンジニアに「痛恨の一撃」を加える快書もしくは怪書。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-21">2014/09/21</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
