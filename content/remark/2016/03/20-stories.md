+++
date = "2016-03-20T19:09:04+09:00"
update = "2016-03-26T13:11:22+09:00"
description = "NIST SP 800-175B Draft / 翻案は「二次創作」を指すものではない / 引っ越しの際に真っ先に捨てるのは「本」"
draft = false
tags = ["security", "privacy", "cryptography", "nist", "copyright", "derivative-works", "book"]
title = "週末スペシャル： NIST SP 800-175B Draft"

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

1. [NIST SP 800-175B Draft]({{< relref "#nist" >}})
1. [翻案は「二次創作」を指すものではない]({{< relref "#adapt" >}})
1. [引っ越しの際に真っ先に捨てるのは「本」]({{< relref "#book" >}})

## NIST SP 800-175B Draft{#nist}

- {{< pdf-file title="NIST Special Publication 800-175B Guideline for Using Cryptographic Standards in the Federal Government: Cryptographic Mechanisms" link="http://csrc.nist.gov/publications/drafts/800-175/sp800-175b_draft.pdf" >}}

519行目より

{{< fig-quote title="NIST Special Publication 800-175B" link="http://csrc.nist.gov/publications/drafts/800-175/sp800-175b_draft.pdf" lang="en" >}}
<q>SKIPJACK is referenced in FIPS 185 and s pecified in a classified document. SKIPJACK is no longer considered adequate for the protection of federal information and has been withdrawn as a FIPS.  The use of SKIPJACK for applying cryptographic protection (e.g., encryption) is not approved, although it is permissible to use the algorithm for decrypting information.</q>
{{< /fig-quote >}}

SKIPJACK 暗号は NSA が1990年代に考えた 80bits 鍵長のブロック暗号で，当時「強い」と言われていた DES よりも暗号強度が高かった。
NSA はこの暗号アルゴリズムと「鍵預託（key escrow）」機能を持つ公開鍵暗号システムを組み合わせたものをコンピュータ・チップに内蔵させた。
これが Clipper Chip である。

AT&T は自社の電話暗号化システムに Clipper Chip を搭載する契約を政府と結んだ（見返りに AT&T は政府という大口の顧客を得ることになる）。
一方，1992年末にクリントンが大統領に選出され，その副大統領ゴアによる「情報ハイウェイ」構想が打ち立てられた。
FBI や NSA はこの機に乗じて Clipper Chip を売り込んだ[^a]。

[^a]: クリントン政権は最初は暗号技術の輸出規制を撤廃しようと考えていた。それに対抗する方便として暗号チップに「鍵預託」機能を埋め込む方法を提案したわけだ。「鍵預託」を使えば暗号強度を損なうことなく，かつ米国政府はいつでも暗号通信の中身を解読できると売り込んだ。1993年に起きた世界貿易センター破壊テロも追い風になったようだ。

実際には Clipper Chip は大きな批判にさらされた。
「自宅の玄関の鍵のコピーを警察署に預けなければならないとしたらどうするか？」（Steven Levy 著『[暗号化](http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/)』より）というわけだ。
Clipper Chip は Big Brother Chip と揶揄された。
最終的に鍵預託の仕様に欠陥が見つかったことなどがきっかけで Clipper Chip は凋落し電話やパソコンに使われることはなかった。

鍵預託の仕様は1994年に {{< pdf-file title="FIPS 185" link="http://csrc.nist.gov/publications/fips/fips185/fips185.pdf" >}} として標準化されたが，これが正式に obsolete となったのはようやく2015年である。
SKIPJACK 暗号も当時は強力と言われていたが，今や 80bits ぽっちの暗号が（古い資産を復号する以外の）役に立たないことは皆さんご存知の通り。

しかしこの20年余りの間，大規模テロが起こるたびに米国政府（というか FBI）は同じような試みを繰り返しては失敗している。
国防上の理由と称して国民をテロリストと同列に監視対象に入れようとしているわけだ。

- [Cryptography Is Harder Than It Looks - Schneier on Security](https://www.schneier.com/blog/archives/2016/03/cryptography_is.html)

## 翻案は「二次創作」を指すものではない{#adapt}

- [北本かおり×ドミニク・チェン×山内康裕：二次創作とライセンス 「ただコンテンツを眺めるだけじゃなく、自由に使うことによって文化が成長する。」 - DOTPLACE](http://dotplace.jp/archives/21269)
- [北本かおり×ドミニク・チェン×山内康裕：二次創作とライセンス 「非親告罪化には、著者の『許可したいです』という意思を発露する手段が奪われていく感じがします。」 - DOTPLACE](http://dotplace.jp/archives/21441)
- [北本かおり×ドミニク・チェン×山内康裕：二次創作とライセンス 「海賊版を禁止して広がるのを止めてしまうのか、『いいから出しちゃえ』って進んだ人が結局最後に勝つのか。」 - DOTPLACE](http://dotplace.jp/archives/21458)

現行の（ベルヌ条約にもとづく）著作権法はあらゆる「表現」に対して自動的に著作権を付与する。
それこそ風呂場の鼻歌やコースターの裏の落書きにも著作権が付与されてしまう。

「著作権」がクリエイターや出版社だけの問題であるならば，このリンク先の議論は有効だと思う。
しかし実際にはそうではないのだ。
著作権の非親告罪化を「二次創作」への脅威として捉えるのは問題の矮小化である。

山田奨治さんの『[〈海賊版〉の思想](http://www.amazon.co.jp/exec/obidos/ASIN/4622073455/baldandersinf-22/)』のあとがきにこんな記述がある。

{{< fig-quote title="〈海賊版〉の思想" >}}
<q>図工の時間に、ある児童が行詰まってしまった。いいアイデアが、どうしても浮かばないのだ。先生は「お友だちが作っているのをみて、参考にしてごらん」という。そこで、図工の得意な子が何を作っているかをみにいったら、こういわれた。「ぼくのまねをすると、著作権のしんがいになるよ。」</q>
{{< /fig-quote >}}

非親告罪化した著作権法のもとで，「ある児童」が「図工の得意な子」の作品を翻案すればその子は犯罪者で「先生」も共犯者だ，「図工の得意な子」がそれに対してどう思うかに関わらず。
あるいは子どもがテレビの前でプリキュアダンスを踊ってそれをお父さんがビデオで撮ったなら，その親子も犯罪者確定である。

「[著作権と著作権法]({{< ref "/cc-licenses/01-copyright.md" >}})」でも書いたが，ユーザにとって表現は他者とのコミュニケーション手段であり，コミュニケーションとしての表現はすべて翻案なのである。
もし非親告罪化を認めるのなら，それは明確に「表現の自由」の侵害である。

米国には fair use doctrine があり[^fud]，上記のケースも fair use として認められるかもしれない。
しかし日本にはそのような仕組みはない。
それどころか

{{< fig-quote title="デジタル時代の「孤児作品」問題を解消するには--権利者団体が議論 - CNET Japan" link="http://japan.cnet.com/news/business/35079134/" >}}
<q>一般社団法人日本音楽著作権協会常務理事の浅石道夫氏が「エンドユーザーと権利者の対立を表にだして自らは存在感を消している。他人の財産（著作物）を使って利益を生み出す事業者による詭弁」と真っ向から切り捨てた。</q>
{{< /fig-quote >}}

などと言ってはばからないのが現状である。

[^fud]: ある利用が Fair use に該当するかどうかは「利用の目的や本質」「原作品の本質」「抜粋の量や実質性」「原作品の価値への影響」という4つの観点で議論される。

## 引っ越しの際に真っ先に捨てるのは「本」{#book}

- [引っ越しの際に出た不用品、トップは本、モノの所有に固執しない時代に？ -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160314_748044.html)

紙の本の愛好者は様々な理由をつけて紙の本を擁護しようとするが，実際にはこんなもんである。
うちの母親は年金をもらう歳になってから積極的に本を読むようになったが（今まで我慢してたんだねぇ，ゴメン），読み終わった本は片っ端から売りに出しているらしい。

もはや多くの人にとって「本」は知識や教養の象徴ではなく，「読み捨て」の消費財である。
また「本」をそのような位置づけにしたのは他ならぬ作家と出版社である。
「本」がそういう位置付けなら，そりゃあブックオフや図書館を利用するほうが合理的だよね。
個人的には「本」も音楽や映像と同じく「水のように」なるべきだと思うな[^c]。
まぁ，もしそうなったら本屋は今のレコード・ショップ並にニッチに特化するしかないけど。

[^c]: 漫画アプリとか一部は既にそうなってるけど。でもそれってもう「本」のメタファーから外れてる気はするけどね。

- [北海道のシャッター通りに本屋をつくる « マガジン航[kɔː]](http://magazine-k.jp/2016/03/02/little-bookstore-in-northern-street/)
- [アメリカのインディペンデント書店が強いわけ « マガジン航[kɔː]](http://magazine-k.jp/2016/03/17/beyond-cool-japan-04/)

## その他の気になる記事{#other}

- [オープンアクセス（OA）とクリエイティブコモンズ（CC） - 水野祐（@TasukuMizuno）のブログ](http://tasukumizuno.hatenablog.com/entry/2016/03/14/190655)
- [セキュリティ| レポート2015年度版ダウンロード | ソニーデジタルネットワークアプリケーションズ株式会社](http://www.sonydna.com/sdna/solution/report02.html)
- [誰もあなたの製品を使いたいと思ってはいない : 製品をデザインするための考え方 | プロダクト・サービス | POSTD](http://postd.cc/nobody-wants-use-your-product/)
- [電子書籍の利用率が2割弱で頭打ち、「利用意向なし」が増加、「関心なし」と合わせると6割以上に -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160314_747958.html) : 米国でも3割ちょっとらしいので，日本ならそんなもんだろう
- [シリアのクルド勢力がシリア北部の自治を宣言へ：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41026)
    - [Vladimir Putin, Godfather of Kurdistan? | The National Interest](http://nationalinterest.org/feature/vladimir-putin-godfather-kurdistan-15358)
    - [シリアのクルド人勢力の自治への動きに関する報道：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41028)
    - [シリアのクルド人勢力の自治政府宣言の背後に米露の協調・支持はあるのか？：池内恵 | 中東の部屋 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41027)
- [自動運転車のAIが「ドライバー」であるとした米国運輸省の回答の意味（前編）「NHTSAの発表は無人運転車に関するFMVSSの解釈を示しただけ」とは - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2016/03/51254/)
    - [自動運転車AIが「ドライバー」であるとした米国運輸省の回答の意味（後編）完全自動運転車が問う「人間とテクノロジー」の関係 - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2016/03/51258/)

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4314009071/baldandersinf-22/">暗号化 プライバシーを救った反乱者たち</a></dt><dd>スティーブン・レビー 斉藤 隆央 </dd><dd>紀伊國屋書店 2002-02-16</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/487593100X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/487593100X.09._SCTHUMBZZZ_.jpg"  alt="ハッカーズ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4105393022/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4105393022.09._SCTHUMBZZZ_.jpg"  alt="暗号解読―ロゼッタストーンから量子暗号まで"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4484111160/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4484111160.09._SCTHUMBZZZ_.jpg"  alt="グーグル ネット覇者の真実 追われる立場から追う立場へ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/410215972X/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/410215972X.09._SCTHUMBZZZ_.jpg"  alt="暗号解読〈上〉 (新潮文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4102159738/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4102159738.09._SCTHUMBZZZ_.jpg"  alt="暗号解読 下巻 (新潮文庫 シ 37-3)"  /></a> </p>
<p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-03-09">2015/03/09</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4622073455/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41WpTRWCAvL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4622073455/baldandersinf-22/">〈海賊版〉の思想‐18世紀英国の永久コピーライト闘争</a></dt><dd>山田 奨治 </dd><dd>みすず書房 2007-12-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4121023390/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4121023390.09._SCTHUMBZZZ_.jpg"  alt="ロラン・バルト -言語を愛し恐れつづけた批評家 (中公新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087207846/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087207846.09._SCTHUMBZZZ_.jpg"  alt="盗作の言語学 表現のオリジナリティーを考える (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480689281/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480689281.09._SCTHUMBZZZ_.jpg"  alt="18歳の著作権入門 (ちくまプリマー新書)"  /></a> </p>
<p class="description">「コピーライト永久独占を目論む大書店主に挑む〈海賊出版者〉ドナルドソンの肖像。法廷闘争を軸に著作権を史的に考察する。」（帯文より）</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-10-17">2015-10-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
