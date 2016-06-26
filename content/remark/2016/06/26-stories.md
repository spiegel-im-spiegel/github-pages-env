+++
date = "2016-06-26T20:48:08+09:00"
description = "英国の EU 離脱 / CRYPTREC Report 2015 が出てた / openpgp-rfc4880bis-02 も出てた / 七夕プロジェクト 2016 / その他の気になる記事"
draft = false
tags = ["international", "security", "cryptography", "cryptrec", "openpgp", "astronomy"]
title = "週末スペシャル： 英国の EU 離脱，他"

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

最近のカープは何事ですか。
交流戦で3位につけただけでも驚きなのに，その流れのまま9連勝ですよ。
なにこれ，ひょっとして死亡フラグなの？ 四半世紀のあいだ裏切られ続けてきたカープファンはこの事態を全く信用していない。
でも，勝ちはボれるときにボっとかないとね。

1. [英国の EU 離脱]({{< relref "#eu" >}})
1. [CRYPTREC Report 2015 が出てた]({{< relref "#cr" >}})
1. [openpgp-rfc4880bis-02 も出てた]({{< relref "#op" >}})
1. [七夕プロジェクト 2016]({{< relref "#tanabata" >}})
1. [その他の気になる記事]({{< relref "#other" >}})


## 英国の EU 離脱{#eu}

まぁ，今週最大のニュースはこれだろうね。

- [英国「EU離脱」の深層とそれに続く「予備軍」：田中直毅 | 経済の頭で考えたこと | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41288)
- [岩瀬昇のエネルギーブログ 208．英国、国民投票でEU離脱へ｜岩瀬昇のエネルギーブログ](http://ameblo.jp/nobbypapa/entry-12174176404.html)

EU というのは「包摂」型の政治システムだと思う。
「包摂」型の政治システムで一番わかり易いのが「多数決」だ。
その多数決（国民投票）によって英国は「包摂」型システムから逃れることを選択したことになる。

もちろん EU 離脱がリスキーであることは間違いないし，英国もそのことを承知でなお離脱を選択したと解釈することもできる。
これが英国にとって吉と出るか凶と出るかなんて誰にもわからない。

{{< fig-quote title="英国、国民投票でEU離脱へ" link="http://ameblo.jp/nobbypapa/entry-12174176404.html" >}}
<q>もっとも興味深かったのは「英国のEU離脱により長期的に価格は下落するか」という質問に対し、エコノミストは「Yes」、アナリストは「No」と応えている箇所だった。<br>
エコノミストは、今後の見通しがきわめて不透明になったことにより、世界景気の減速は免れない、ドミノ現象が起こればさらに不透明となり、不況は長期化する、したがって石油需要が減少するので価格は下落する、と判断している。<br>
一方のオイル・アナリストは、通貨価値、世界景気への影響がないとは言わないが、需給を含む基本的な石油市場の構造は不変だ、リバランスへの道筋に影響はなく、長期的に価格が下落することにはつながらない、と主張しているのだ。</q>
{{< /fig-quote >}}

将来の不透明感を「凶」と見なす人もいれば「吉」と見なす人もいる。
でも「ピンチはチャンス」と言える人こそが最終的に生き残れるのではないだろうか。

## CRYPTREC Report 2015 が出てた{#cr}

久しぶりに [CRYPTREC] のページを覗いたら「CRYPTREC Report 2015」が出てた。
あと「CRYPTREC 暗号リスト」も更新されていた。

- [CRYPTREC Report 2015の公開](http://www.cryptrec.go.jp/topics/cryptrec_20160617_c15report.html)
- {{< pdf-file title="電子政府における調達のために参照すべき暗号のリスト（ＣＲＹＰＴＲＥＣ暗号リスト）" link="http://www.cryptrec.go.jp/images/cryptrec_ciphers_list_2016.pdf" >}}

推奨候補暗号リストとして新たに `SHA-512/256`, `SHA3`, `SHAKE256` が加わったようだ。
レポートの方は後でゆっくり読む（暇があるかなぁ）。

お願いですから [CRYPTREC] の feed を付けてください。

## openpgp-rfc4880bis-02 も出てた{#op}

- [draft-koch-openpgp-rfc4880bis-02 - OpenPGP Message Format](https://tools.ietf.org/html/draft-koch-openpgp-rfc4880bis-02)

これも後でゆっくり読む（暇があるかなぁ）。

ところで [GitLab.com] に rfc4880bis のリポジトリができているようだ。

- [openpgp-wg / rfc4880bis · GitLab](https://gitlab.com/openpgp-wg/rfc4880bis)

なんで（[GitHub](https://github.com/) じゃなくて） [GitLab.com] なのかは分からないが，まぁいいか。
こちらでもドラフト案や関連資料を見ることができる。

## 七夕プロジェクト 2016{#tanabata}

- [七夕プロジェクト2016](https://tanabata-project.jp/)

なんか短冊に書く「願い」を募集中らしい。
しめきりは今年の「[伝統的七夕](http://www.nao.ac.jp/faq/a0310.html "伝統的七夕について教えて | 国立天文台")」である8月9日まで。
興味のある方はどうぞ。

私は今のところ神仏にお願いするようなことは何もないのでパスします。

## その他の気になる記事{#other}

- [aozorahack hackathon #1 - connpass](http://aozorahack.connpass.com/event/33921/)
- [「太陽系インターネット」構築へ　NASAが新ネットワーク「DTN」導入 | sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/030201/2016_06_24_nasa.html)
- [Research Blog: Bringing Precision to the AI Safety Discussion](https://research.googleblog.com/2016/06/bringing-precision-to-ai-safety.html)
    - [グーグル、安全な人工知能開発のために問うべき課題に関する論文を発表 - CNET Japan](http://japan.cnet.com/news/service/35084726/)
- {{< pdf-file title="Windows 10への無償アップグレードに関し、確認・留意が必要な事項について" link="http://www.caa.go.jp/adjustments/pdf/160622adjustments_1.pdf" >}}
    - [消費者庁がWindows 10への自動アップグレードについてWindows 7/8.1ユーザーへ注意喚起 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1006675.html)
- [マイナンバーカードでSSHする - AAA Blog](https://www.osstech.co.jp/~hamano/posts/jpki-ssh/)

[CRYPTREC]: http://www.cryptrec.go.jp/
[GitLab.com]: https://gitlab.com/

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4106037866/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51QsC2WBr5L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4106037866/baldandersinf-22/">【中東大混迷を解く】 サイクス=ピコ協定 百年の呪縛 (新潮選書)</a></dt><dd>池内 恵 </dd><dd>新潮社 2016-05-27</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4120048349/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4120048349.09._SCTHUMBZZZ_.jpg"  alt="増補新版 イスラーム世界の論じ方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4140884908/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4140884908.09._SCTHUMBZZZ_.jpg"  alt="中東から世界が崩れる―イランの復活、サウジアラビアの変貌 (NHK出版新書 490)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4106037890/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4106037890.09._SCTHUMBZZZ_.jpg"  alt="世界地図の中で考える (新潮選書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484162164/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484162164.09._SCTHUMBZZZ_.jpg"  alt="アステイオン84 【特集】帝国の崩壊と呪縛"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4106037874/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4106037874.09._SCTHUMBZZZ_.jpg"  alt="憲法改正とは何か: アメリカ改憲史から考える (新潮選書)"  /></a> </p>
<p class="description">まだ読んでる途中だが，欧州および中東の近代史を「サイクス=ピコ協定」を特異点として網羅的に解説していいる。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-06-26">2016-06-26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

