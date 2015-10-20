+++
date = "2015-10-17T15:51:23+09:00"
update = "2015-10-21T05:21:00+09:00"
description = "まずは「著作権」について簡単におさらい。"
draft = false
tags = ["code", "law", "intellectual-property", "copyright"]
title = "著作権と著作権法"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

まずは「著作権」について簡単におさらい。
なお，私は法の専門家ではないので説明が大雑把だったり間違いがあったりするかもしれませんが，平にご容赦。
（事実関係の間違いは指摘していただけると助かります）

## 「新しい葡萄酒[^0]」は独占を望まない

[^0]: 「[グーグル・ブック・サーチ，あるいはバベルの図書館 新しいぶどう酒は新しい革袋に](https://www.jstage.jst.go.jp/article/johokanri/53/3/53_3_131/_article/-char/ja/)」参照。ちなみに「新しい葡萄酒」というのはマタイ伝第9章17節で出てくる言葉で，（旧いユダヤ教に対する）キリスト教を指す。

「著作権（copyright）」は「知的財産権（intellectual property）」の一種で，日本の[著作権法]第1条では以下のように書かれている。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>この法律は、著作物並びに実演、レコード、放送及び有線放送に関し著作者の権利及びこれに隣接する権利を定め、これらの文化的所産の公正な利用に留意しつつ、著作者等の権利の保護を図り、もつて文化の発展に寄与することを目的とする。</q>
{{< /fig-quote >}}

もちろんこんなものが建前だけのコンコンチキであることは皆さんご存知だろう。
そうでなければ「[著作権トロル（copyright troll）](https://www.eff.org/issues/copyright-trolls "Copyright Trolls | Electronic Frontier Foundation")」みたいなものが跋扈するはずない。
大人の世界って汚いね。

著作権の起源は「出版特許」にあると言われている。
時の為政者から出版ギルドへ出版特許を「下賜」することにより，書籍流通の安定と統制を確保しようとしたらしい。
つまり著作権のもとになったものは，著作者のためではなく，もちろん読者（ユーザ）のためなんかではなく，ひたすら出版社の特権として生み出されたものなのである。

それから紆余曲折があって現在の著作権に近いかたちに整備されてきた。
現代の著作権の枠組みは1886年にできた「ベルヌ条約」がベースになっている。
かつての出版特許との違いは，著作権が，著作者と出版社，著作者と著作者，出版社と出版社，の間のパワーバランスを図る手段として機能していることである。

ただ，この時点でもまだユーザは創作・出版のプレイヤーではなかった。
というか，割と最近までユーザは著作権システムの埒外であった。
これが数十年前くらいから変わってきた。

変化のひとつは，プログラムコードを「著作物」として含めたことにより著者とユーザの区別が曖昧になり，著作物であるプログラム・コードを巡る「独占 vs 自由」の抗争が本格化してきたこと。
もうひとつは情報の符号化・複製コストの劇的低下により既存の出版・流通による独占が意味を失いつつあること，である。
前者は FLOSS（Free/Libre and Open Source Software）というかたちでコンピュータ・ソフトウェア産業に大規模な構造変革をもたらし，発展を加速させている。
後者は現在進行形で，既存の出版・流通チャネルはこの変化に対して必死に抵抗を試みている[^a]。

[^a]: こうした抵抗のひとつが「[著作権トロル](https://www.eff.org/issues/copyright-trolls "Copyright Trolls | Electronic Frontier Foundation")」といった形で現れているわけだ。

こうした変化により著作者と出版社（者）とユーザの区別は曖昧になりつつある。
しかし，著作権の枠組みは19世紀から殆ど変わっておらず，一連の「大変化」に適合できてない。

著作者や出版社にとっての著作権システムとは，著作物の「独占」状態を如何にコントロールするかであった。
しかし，ユーザの場合は行動原理が異なる。
それはひとことで言うと「共有」である。
何故ならユーザはコミュニケーションの手段として著作物を利用するからだ。

## 著作権とは

では，現行の[著作権法]はどのようなものか簡単に説明していく。

### 「著作物」とは

まずは「著作物」の定義から。
これは[著作権法]第2条に記されている。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>著作物　思想又は感情を創作的に表現したものであつて、文芸、学術、美術又は音楽の範囲に属するものをいう。 </q>
{{< /fig-quote >}}

ちょっと分かりにくいけど，この条文のポイントは著作物が「創作的に表現」したものであるということ。
したがってアイデアは著作物に含まれない（アイデアは特許権で保護される）。
またキャラクタや名前も基本的には[^c] 著作物に含まれない（キャラクタや名前は商標権や意匠権で保護される）。

[^c]: ただし創作上の文脈としてキャラクタや名前が書（描）かれている場合は著作物の一部として認められる場合がある。

具体的に著作物とは何を指すかというと，[著作権法] 第10条から12条で以下を挙げている。

- 言語の著作物
- 音楽の著作物
- 舞踊又は無言劇の著作物
- 美術の著作物
- 建築の著作物
- 地図又は図形の著作物
- 映画の著作物
- 写真の著作物
- プログラムの著作物
- 二次的著作物
- 編集著作物
- データベースの著作物

「プログラムの著作物」や「データベースの著作物」が「文芸、学術、美術又は音楽の範囲に属するもの」かどうかは甚だ怪しいところではあるが，これらは1996年の「著作権に関する世界知的所有権期間条約」（通称「WIPO著作権条約」または WCT）によって著作物であると決められてしまった。
日本の著作権法もこれに従う形となっている。

ちなみに，過去の判例により，ゲームソフトは「プログラムの著作物」ではなく「映画の著作物」に含まれることになった[^cc] 。
（ここでは触れないが）「映画の著作物」は他の著作物に比べてかなり優遇されているため，ゲームソフトもそれに倣うことになってしまった。
やれやれ。

[^cc]: ただし，全てのゲームが「映画の著作物」となるわけではない（参考： [ビデオ・ゲームの中古販売と著作権法](http://www.ben.li/essay/Second-handed-game.html)）。

「二次的著作物」については別の記事で改めて紹介する。

例外として，以下のものには著作権は適用されない（[著作権法] 第13条）。

1. 憲法その他の法令
2. 国若しくは地方公共団体の機関、独立行政法人又は地方独立行政法人が発する告示、訓令、通達その他これらに類するもの
3. 裁判所の判決、決定、命令及び審判並びに行政庁の裁決及び決定で裁判に準ずる手続により行われるもの
4. 上の3つのものの翻訳物及び編集物で、国若しくは地方公共団体の機関、独立行政法人又は地方独立行政法人が作成するもの

ただし，各省庁が出す「◯◯白書」のようなものは上記の例外に含まれないので注意。

### 「著作者」とは

「著作者」は「著作物を創作する者」（[著作権法] 第2条）である。
以上。
では面白くないので，もう少し具体的に[著作権法] 第14条から。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>著作物の原作品に、又は著作物の公衆への提供若しくは提示の際に、その氏名若しくは名称（以下「実名」という。）又はその雅号、筆名、略称その他実名に代えて用いられるもの（以下「変名」という。）として周知のものが著作者名として通常の方法により表示されている者は、その著作物の著作者と推定する。</q>
{{< /fig-quote >}}

つまり創った人が誰か特定できるのであれば実名を名乗る必要はない。

もうひとつ。
「職務著作」という考え方がある（[著作権法] 第15条）。

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>法人その他使用者（以下この条において「法人等」という。）の発意に基づきその法人等の業務に従事する者が職務上作成する著作物（プログラムの著作物を除く。）で、その法人等が自己の著作の名義の下に公表するものの著作者は、その作成の時における契約、勤務規則その他に別段の定めがない限り、その法人等とする。</q>
{{< /fig-quote >}}

{{< fig-quote title="著作権法" link="http://law.e-gov.go.jp/htmldata/S45/S45HO048.html" >}}
<q>法人等の発意に基づきその法人等の業務に従事する者が職務上作成するプログラムの著作物の著作者は、その作成の時における契約、勤務規則その他に別段の定めがない限り、その法人等とする。 </q>
{{< /fig-quote >}}

つまり企業・組織等の法人が著作者になることができる。
たとえば私は職業プログラマだが，私が仕事で書いたプログラムはその仕事を管轄する企業・組織が著作者となる[^d] [^e]。

[^d]: 別の企業・組織が発注したプログラムを請け負っている場合は，大抵は契約により，その企業・組織に権利が帰属する。
[^e]: 法学者の[白田秀彰さん](http://orion.mt.tama.hosei.ac.jp/hideaki/indexj.htm)はこの点を突いて著作権法の矛盾を指摘している： [ほんとうの創作者利益について](http://thinkcopyright.org/shirata0115.html)

### 「著作権」は権利の束

さて，著作物に対して著作者が行使可能な「著作権」だが，これはひとつの権利ではなく複数の権利の束になっていて，それがが複雑怪奇に絡み合っている。
それぞれの権利の詳細についてはここでは述べないが，列挙だけしておく。

- 著作者人格権
    - 公表権（[著作権法] 第18条）
    - 氏名表示権（[著作権法] 第19条）
    - 同一性保持権（[著作権法] 第20条）
- 著作財産権
    - 複製権（[著作権法] 第21条）
    - 上演権・演奏権（[著作権法] 第22条）
    - 上映権（[著作権法] 第22条）
    - 公衆送信権（[著作権法] 第23条）
        - 放送権・有線放送権
        - 自動公衆送信権
        - 送信可能化権
    - 伝達権（[著作権法] 第23条）
    - 口述権（[著作権法] 第24条）
    - 展示権（[著作権法] 第25条）
    - 頒布権（[著作権法] 第26条：映画の著作物のみ）
    - 貸与権（[著作権法] 第26条：映画の著作物のみ）
    - 譲渡権（[著作権法] 第26条：映画の著作物以外）
    - 翻訳・翻案権等（[著作権法] 第27条：二次的著作物の作成に関する権利）
        - 翻訳
        - 編曲
        - 変形
        - 脚色
        - 映画化
        - 翻案
    - 二次的著作物の利用に関する原著作者の権利（[著作権法] 第28条）
- 出版権（「複製権等保有者[^r1]」が「出版」を引き受ける者に対して設定する権利で，権利者は出版の義務を伴う）
    - 複製権（[著作権法] 第80条）
    - 公衆送信権（[著作権法] 第80条）
    - 修正・増減権（[著作権法] 第82条：出版権者に対する複製権等保有者の権利）
    - 撤回権（[著作権法] 第84条：出版権者に対する複製権等保有者の権利）
- 著作隣接権
    - 実演家の権利
        - 実演家の人格権（[著作権法] 第90, 101, 102条）
            - 氏名表示権
            - 同一性保持権
        - 録音・録画権（[著作権法] 第91条）
        - 放送権・有線放送権（[著作権法] 第92条）
        - 送信可能化権（[著作権法] 第92条）
        - 二次使用料を受ける権利（[著作権法] 第95条）
        - 譲渡権（[著作権法] 第95条）
        - 貸与権（[著作権法] 第95条）
    - レコード製作者の権利
        - 複製権（[著作権法] 第96条）
        - 送信可能化権（[著作権法] 第96条）
        - 二次使用料を受ける権利（[著作権法] 第97条）
        - 譲渡権（[著作権法] 第97条）
        - 貸与権（[著作権法] 第97条）
    - 放送事業者の権利
        - 複製権（[著作権法] 第98条）
        - 再放送権・有線放送権（[著作権法] 第99条）
        - 送信可能化権（[著作権法] 第99条）
        - テレビジョン放送の伝達権（[著作権法] 第100条）
    - 有線放送事業者の権利
        - 複製権（[著作権法] 第100条）
        - 再放送権・有線放送権（[著作権法] 第100条）
        - 送信可能化権（[著作権法] 第100条）
        - テレビジョン放送の伝達権（[著作権法] 第100条）

[^r1]: [著作権法] 第79条では，著作物に対して複製権と公衆送信権を持つ著作権者を「複製権等保有者」と定義している。

### 著作権の発生と消尽

現在の著作権はベルヌ条約に基づき著作物が創作された時点で自動的に付与される。
これを「無方式主義」と言う[^ff]。
つまり思いつきの鼻歌だろうがコースターの裏の落書きだろうが，発生した時点でそれらには著作権が付与されてしまうのだ。

[^ff]: これに対して，（登録制にするなど）何らかの手続を踏むことで著作権が発生するやり方を「方式主義」という。いわゆる &copy; マークは方式主義をとる国や地域（でかつ万国著作権条約（1955年）加盟国）に対して著作物に著作権があることを明示するためのマークである。現在ではほとんどの国や地域で無方式主義をとっているため &copy; マークは必要ない。

著作権のうち著作者人格権は「著作者の一身に専属し、譲渡することができない」（[著作権法] 第59条）[^f]。
一方，著作財産権については「その全部又は一部を譲渡することができる」（[著作権法] 第61条）[^g]。
したがって「著作者」と「著作権者」が異なる場合がありうる。
[本シリーズ]では両者をまとめて「著作（権）者」と表記する場合がある。
しかしライセンスで許諾することができるのは厳密には「著作権者」の方である。

[^f]: 著作隣接権のうち「実演家の人格権」についても同様に譲渡できないし他者が行使することも出来ない（[著作権法] 第101条）。
[^g]: 著作隣接権についても，「実演家の人格権」を除き，基本的には譲渡可能（[著作権法] 第103条）。

著作者人格権または実演家の人格権には期限が設けられてなく，死後も権利が継続される[^h]。
著作財産権については，基本的に著作者の死後50年[^h2]までの期限が設けられているが，映画の著作物は例外として公表後70年（創作後70年以内に公表されなかった場合は創作後70年）となっている。
なお，現在 TPP (Trans-Pacific Partnership; 環太平洋パートナーシップ協定) においてその期限を一律70年まで引き上げようとしているのはご存知のとおりである。

[^h]: ただし，人格権の侵害があったとしても，それが本当に著作者にとって害になるかどうかは考慮されるらしい。
[^h2]: 無名・変名の著作者で死亡時期が分からないものについては公表後50年。法人の場合は公表後50年（創作後50年以内に公表されなかった場合は創作後50年）。

## ユーザから見た著作権

著作権は著作物に対するあらゆる行為を権利化し独占しているわけではない。

前節で述べた権利群にまつわる行為を「利用」と呼ぶ。
それ以外の，たとえば「本を読む」「音楽を聞く」「映画を観る」といった行為は権利化されないため，自由に行うことができる。
こういった「利用」以外の行為を著作物の「使用」と呼ぶ。
つまり，著作権者は「利用」に関して他者をコントロールできるが，「使用」に関しては関知しない。

著作権を厳密に解釈するなら事前に許諾がなければ「利用」できないのだが，実際には事後承諾だったり利用後にお金を払ったりといった感じで，その場その場で ad hoc な運用になることが多い。
もし著作権を厳密に運用しようとすると著作物を巡る取引が硬直化するため，市場的にかなりのダメージなる可能性が高い。
「著作権に違反しているからダメ」という発想ではそれこそダメなのである。

### 「公正な利用」と「著作権の制限」

こういった問題に対し，米国などでは「公正な利用（fair use）」という考え方がある。
ある利用パターンが社会的に見て「公正」であれば認めようというものだ。
実際にどのような場合に認められるかは司法の場で議論されることが多い。

一方日本では，「公正な利用」というものは存在しない。
その代わり[著作権法]では「著作権の制限」（第30-50条）を設けている。
以下に項目のみ列挙してみる[^k]。

[^k]: 「著作権の制限」には項目ごとに細かい条件があって字面ほど自由に行使できるわけではない。たとえば「パロディ」は「著作権の制限」のどの項目にも該当しないため日本の著作権法では認められていないことになる。

- 私的使用のための複製
- 付随対象著作物の利用
- 検討の過程における利用
- 技術の開発又は実用化のための試験の用に供するための利用
- 図書館等における複製等
- 引用
- 教科用図書等への掲載
- 教科用拡大図書等の作成のための複製等
- 学校教育番組の放送等
- 学校その他の教育機関における複製等
- 試験問題としての複製等
- 視覚障害者等のための複製等
- 聴覚障害者等のための複製等
- 営利を目的としない上演等
- 時事問題に関する論説の転載等
- 政治上の演説等の利用
- 時事の事件の報道のための利用
- 裁判手続等における複製
- 行政機関情報公開法 等による開示のための利用
- 公文書管理法 等による保存等のための利用
- 国立国会図書館法 によるインターネット資料及びオンライン資料の収集のための複製
- 翻訳、翻案等による利用
- 放送事業者等による一時的固定
- 美術の著作物等の原作品の所有者による展示
- 公開の美術の著作物等の利用
- 美術の著作物等の展示に伴う複製
- 美術の著作物等の譲渡等の申出に伴う複製等
- プログラムの著作物の複製物の所有者による複製等
- 保守、修理等のための一時的複製
- 送信の障害の防止等のための複製
- 送信可能化された情報の送信元識別符号の検索等のための複製等
- 情報解析のための複製等
- 電子計算機における著作物の利用に伴う複製
- 情報通信技術を利用した情報提供の準備に必要な情報処理のための利用
- 複製権の制限により作成された複製物の譲渡

クラクラする。

米国のような「公正な利用」がいいのか日本のような「著作権の制限」がいいのかは難しいところだ。
しかし，時代や価値観が大きく変化している真っ最中の現代において，日本のやりかたが「遅きに失している」ことは否めない[^m]。

[^m]: [著作権法]の標準的教科書と言われる『[著作権法](http://www.amazon.co.jp/exec/obidos/ASIN/4641144699/baldandersinf-22/)』第2版のなかで「[本書初版（2007）においては、フェアユース規定の導入には消極的な見解を述べていたが、本文で述べる理由のように、現在はフェアユース規定を早急に導入すべきであると言う見解に改めた](http://www.techvisor.jp/blog/archives/5228)」と書かれているらしい。

### 「共有」のためのライセンス

ユーザ側から見た場合，「法」が整備されていくのを悠長に待っている訳にはいかない。
ならばいっそ，あらかじめ「共有」に必要な許諾を行うライセンスを作って運用するほうが理にかなっている。
そうして作られたライセンス・ツールのひとつが Creative Commons Licenses (CC Licenses) である。

というわけで，ようやく前置き終わり。
次回から CC Licenses について解説していく。

## 参考文献

### 参考になる（かもしれない） Web ページ

- [バーチャルネット法律娘　真紀奈17歳](http://homepage3.nifty.com/machina/) : 古いコンテンツだが，このなかの「著作権法講座」は参考になる
- [著作権Q&A | 公益社団法人著作権情報センター CRIC](http://www.cric.or.jp/qa/index.html)
- [特集 : 18歳からの著作権入門 - CNET Japan](http://japan.cnet.com/sp/copyright_study/)

### 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22/">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt><dd>名和 小太郎 </dd><dd>エヌティティ出版 2010-06-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4569812902/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4569812902.09._SCTHUMBZZZ_.jpg"  alt="著作権法がソーシャルメディアを殺す (PHP新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4334037070/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4334037070.09._SCTHUMBZZZ_.jpg"  alt="「ネットの自由」vs.著作権: TPPは、終わりの始まりなのか (光文社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087202941/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087202941.09._SCTHUMBZZZ_.jpg"  alt="著作権とは何か―文化と創造のゆくえ (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4166608347/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4166608347.09._SCTHUMBZZZ_.jpg"  alt="ビジネスパーソンのための契約の教科書 (文春新書 834)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798119806/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798119806.09._SCTHUMBZZZ_.jpg"  alt="REMIX ハイブリッド経済で栄える文化と商業のあり方"  /></a> </p>
<p class="description">名著です。今すぐ買うべきです。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-08-02">2014/08/02</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4622073455/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41WpTRWCAvL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4622073455/baldandersinf-22/">〈海賊版〉の思想‐18世紀英国の永久コピーライト闘争</a></dt><dd>山田 奨治 </dd><dd>みすず書房 2007-12-20</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4121023390/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4121023390.09._SCTHUMBZZZ_.jpg"  alt="ロラン・バルト -言語を愛し恐れつづけた批評家 (中公新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4087207846/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4087207846.09._SCTHUMBZZZ_.jpg"  alt="盗作の言語学 表現のオリジナリティーを考える (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4480689281/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4480689281.09._SCTHUMBZZZ_.jpg"  alt="18歳の著作権入門 (ちくまプリマー新書)"  /></a> </p>
<p class="description">「コピーライト永久独占を目論む大書店主に挑む〈海賊出版者〉ドナルドソンの肖像。法廷闘争を軸に著作権を史的に考察する。」（帯文より）</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-10-17">2015-10-17</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00SM7G6SI/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41ZC-Qu61LL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00SM7G6SI/baldandersinf-22/">18歳の著作権入門 (ちくまプリマー新書)</a></dt><dd>福井健策 </dd><dd>筑摩書房 2015-01-10</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00R3U0XTS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00R3U0XTS.09._SCTHUMBZZZ_.jpg"  alt="誰が「知」を独占するのか ――デジタルアーカイブ戦争 (集英社新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00RF1QG82/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00RF1QG82.09._SCTHUMBZZZ_.jpg"  alt="知的財産法入門 (岩波新書)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W535LOU/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W535LOU.09._SCTHUMBZZZ_.jpg"  alt="HARD THINGS　答えがない難問と困難にきみはどう立ち向かうか"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00RKUS0IC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00RKUS0IC.09._SCTHUMBZZZ_.jpg"  alt="条文の読み方"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00J4ECEFW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00J4ECEFW.09._SCTHUMBZZZ_.jpg"  alt="「ネットの自由」ＶＳ．著作権～ＴＰＰは、終わりの始まりなのか～ (光文社新書)"  /></a> </p>
<p class="description">福井健策弁護士による「<a href="http://japan.cnet.com/sp/copyright_study/">18歳からの著作権入門</a>」の書籍化。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-05-07">2015/05/07</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

[本シリーズ]: /cc-licenses "改訂3版： CC-License について — text.Baldanders.info"
[著作権法]: http://law.e-gov.go.jp/htmldata/S45/S45HO048.html "著作権法"
