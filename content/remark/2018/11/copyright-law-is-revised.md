+++
title = "年末年始に施行される改正著作権法に関する覚え書き"
date = "2018-11-13T14:20:00+09:00"
update = "2019-01-26T22:54:06+09:00"
description = "もう今さらグダグダ言ってもしょうがないので，この記事では今回の改正ポイントについて簡単に紹介するに留める。"
image = "/images/attention/kitten.jpg"
tags = ["code", "law", "intellectual-property", "copyright", "tpp", "access-control"]

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

2018年末から2019年始にかけて改正著作権法が施行される。

1. [著作権法の一部を改正する法律（平成30年法律第30号）について | 文化庁](http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/h30_hokaisei/)
2. [環太平洋パートナーシップ協定の締結に伴う関係法律の整備に関する法律（平成28年法律第108号）及び環太平洋パートナーシップ協定の締結に伴う関係法律の整備に関する法律の一部を改正する法律（平成30年法律第70号）について | 文化庁](http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/kantaiheiyo_hokaisei/)

1番目は（一部を除いて）2019年1月1日に，2番目は2018年12月30日に施行される。
もう今さらグダグダ言ってもしょうがないので，この記事では今回の改正ポイントについて簡単に紹介するに留める。
なお「[改訂3版： CC Licenses について]({{< rlnk "/cc-licenses/" >}})」の各記事は施行後にゆるゆると改訂する予定である。

あと「[著作権の保護期間延長に反対します](https://baldanders.info/spiegel/log2/000820.shtml "Stop Fast Track TPP")」ロゴは今後も貼り付けておくので，よろしくどうぞ。
せめてもの悪あがきである。

## TPP11 締結に伴う法改正

まずは TPP11 関連から。
改正点は以下の通り。

1. 著作物等の保護期間の延長（第51条第2項，第52条第1項，第53条第1項，第101条第2項第1号及び第2号関係）
2. 著作権等侵害罪の一部非親告罪化（第123条第2項及び第3項関係）
3. 著作物等の利用を管理する効果的な技術的手段に関する制度整備（アクセスコントロールの回避等に関する措置）（第2条第1項第21号，第113条第3項，第119条第1項，第120条の2第1項第1号及び第2号関係）
4. 配信音源の二次使用に対する使用料請求権の付与（第95条第1項関係）
5. 損害賠償に関する規定の見直し（第114条第4項関係）

2018年10月31日に メキシコ，日本，シンガポール，ニュージーランド，カナダ，オーストラリア の6ヶ国で国内手続きが完了し，そこから60日後の12月30日に施行となったわけだ。
著作権の保護期間延長の絡みがあるのでギリギリ年末の施行に間に合わせたというところだろうか。

### 著作物等の保護期間の延長

日本では一部を除いて著作物等の保護期間は50年だったが，これが一律70年に延長される。

{{< fig-gen title="環太平洋パートナーシップ協定の締結及び環太平洋パートナーシップに関する包括的及び先進的な協定の締結に伴う関係法律の整備に関する法律の概要（著作権法関係）" link="http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/kantaiheiyo_hokaisei/pdf/r1408266_01.pdf" >}}
<table>
<thead><tr>
    <th colspan="2">種類</th>
    <th>改正前</th>
    <th>改正後</th>
</tr></thead>
<tbody style="text-align:left;">
    <tr>
        <td rowspan="4">著作物</td>
        <td>原則</td>
        <td>著作者の死後50年</td>
        <td>著作者の死後70年</td>
    </tr>
    <tr>
        <td>無名・変名</td>
        <td>公表後50年</td>
        <td>公表後70年</td>
    </tr>
    <tr>
        <td>団体名義</td>
        <td>公表後50年</td>
        <td>公表後70年</td>
    </tr>
    <tr>
        <td>映画</td>
        <td>公表後70年</td>
        <td>公表後70年</td>
    </tr>
    <tr>
        <td colspan="2">実演</td>
        <td>実演が行われた後50年</td>
        <td>実演が行われた後70年</td>
    </tr>
    <tr>
        <td colspan="2">レコード</td>
        <td>レコードの発行後50年</td>
        <td>レコードの発行後70年</td>
    </tr>
</tbody>
</table>
{{< /fig-gen >}}

なお2018年に公有（public domain）となったものを含めて現時点で公有化された著作物等については公有のままとなるらしい。
上表の「著作物」で考えると，1968年に亡くなられた作家の作品は少なくとも2038年までは公有化されないということになる[^cr1]。

[^cr1]: 個人の著作物の場合，著者が亡くなられた翌年初からカウントされる。

### 著作権等侵害罪の一部非親告罪化

現行法では著作権侵害罪は親告罪[^wc1] となっているが，改正後は以下の3つの条件を全て満たす場合に非親告罪の対象となる。

[^wc1]: 「親告罪」とは被害者や身内の告訴がなければ刑事裁判を起こすことができない犯罪を指す。「非親告罪」はその逆で，刑事裁判を起こすにあたり被害者や身内の告訴を必要としない犯罪を指す。通常は訴えがあって初めて捜査が行われるんだろうけど，今後知財トロルが跋扈すれば当事者たちに関係なく刑事裁判を起こされる可能性があり，それをネタに強請られる可能性もあるわけだ。

1. 対価を得る目的又は権利者の利益を害する目的があること
2. 有償著作物等[^cr2] について原作のまま譲渡・公衆送信又は複製を行うものであること
3. 有償著作物等の提供・提示により得ることが見込まれる権利者の利益が不当に害されること

[^cr2]: 「有償著作物等」とは「有償で公衆に提供又は提示されている著作物等」を指すようだ。

「利益が不当に害される」というのは著作（権）者側の判断になると思うのだが，これが非親告罪の条件になるというのは納得いかないものがある，んだよなぁ。
きっと施行後にどなたか偉い人が解説されるんだろう。

どういったものが親告罪/非親告罪の対象になるかという例示は以下に示されている。

{{% fig-gen title="環太平洋パートナーシップ協定の締結及び環太平洋パートナーシップに関する包括的及び先進的な協定の締結に伴う関係法律の整備に関する法律の概要（著作権法関係）" link="http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/kantaiheiyo_hokaisei/pdf/r1408266_01.pdf" %}}
| 非親告罪となる侵害行為の例                 | 親告罪のままとなる行為の例           |
| ------------------------------------------ | ------------------------------------ |
| 販売中の漫画や小説本の海賊版を販売する行為 | 漫画等の同人誌をコミケで販売する行為 |
| 映画の海賊版をネット配信する行為           | 漫画のパロディをブログに投稿する行為 |
{{% /fig-gen %}}

しかし，この例が条文に載せられているわけではなく，おそらくその線引で揉めることになるんだろう[^cr3]。
つか，同人誌を神聖視するのはそろそろ止めようや。

余談だが日本にはパロディに関する規定（fair use）はなく「引用か二次的著作物か」で線引される。
許可なく二次的著作物を公開すれば著作権法違反である。
引用であれば「著作権の制限」に該当するので違反にはあたらない。

[^cr3]: そして知財トロルが跋扈する。

### アクセスコントロールの回避等に関する措置

従来の「技術的保護手段」と新しく定義された「技術的利用制限手段」の回避について「著作権者等の利益を不当に害しない場合を除き，著作権等を侵害する行為とみなす」というものだ（ただし刑事罰の対象にはならない）。
「技術的利用制限手段」には複製を伴わない回避（たとえば暗号化された放送やストリーミングの復号など）も含まれているのがポイントである。

「技術的保護手段」および「技術的利用制限手段」を回避する装置やプログラムの販売等は刑事罰の対象となる。
たとえば著作（権）者が「技術的保護手段」の回避を許諾していても，回避する手段を封じられているため，合法的に回避するのは相当困難になるだろう。

ちなみに [CC Licenses] v4 では対象物に対して（許諾条件を守る限り）「技術的利用制限手段」を含む技術的保護手段を行使することを禁止し，行使されている場合は回避することを許諾している。
でもどうやって回避するんだろうねぇ（笑）

- [CC Licenses における「技術的保護手段」の扱い]({{< ref "/cc-licenses/05-technological-measures.md" >}})

### 配信音源の二次使用に対する報酬請求権の付与

{{< fig-quote title="環太平洋パートナーシップ協定の締結及び環太平洋パートナーシップに関する包括的及び先進的な協定の締結に伴う関係法律の整備に関する法律の概要（著作権法関係）" link="http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/kantaiheiyo_hokaisei/pdf/r1408266_01.pdf" >}}
<q>放送事業者等がＣＤ等の商業用レコードを用いて放送又は有線放送を行う際に，実演家及びレコード製作者に認められている使用料請求権について，対象を拡大し，配信音源を用いて放送又は有線放送を行う場合についても，使用料請求権を付与する。</q>
{{< /fig-quote >}}

これで日本のクリエイターはようやく「レコード」という名の軛から解放される（笑） 放送局とかは大変そうだけど。
これってネットでアドホックに公開されている音源についても「著作権等管理事業者」の制御下に組み入れようという意図なのかしら。

### 損害賠償に関する規定の見直し

{{< fig-quote title="環太平洋パートナーシップ協定の締結及び環太平洋パートナーシップに関する包括的及び先進的な協定の締結に伴う関係法律の整備に関する法律の概要（著作権法関係）" link="http://www.bunka.go.jp/seisaku/chosakuken/hokaisei/kantaiheiyo_hokaisei/pdf/r1408266_01.pdf" >}}
<q>侵害された著作権等が著作権等管理事業者により管理されている場合は，著作権者等は，当該著作権等管理事業者の使用料規程により算出した額（複数ある場合は最も高い額）を損害額として賠償を請求することができる。</q>
{{< /fig-quote >}}

これって揉めたりしないのかなぁ。
（いままでの[某団体の所業]({{< ref "/remark/2017/02/jasrac.md" >}} "JASRAC が音楽教室からも「著作権（みかじめ）料」をまきあげる話")を見る限り）著作権等管理事業者が管理する「使用料」が適正とは限らないぢゃん。
しかも非親告罪になる可能性が高いし。

## 著作権法の一部を改正する法律（平成30年法律第30号）について

この改正では著作権法第30条以降で規定されている「著作権の制限」を変更するもののようだ。

「著作権の制限」とは，簡単に言うと，この規定に従う限り公正な利用（使用ではない）として認めようというものだ。
米国などではそのものズバリの “fair use” があるが，日本の場合は「著作権の制限」として認められる利用を細かく規定している。
はっきり言って時代遅れの法設計。

改正点は以下の通り。

1. デジタル化・ネットワーク化の進展に対応した柔軟な権利制限規定の整備（第30条の4，第47条の4，第47条の5等関係）
2. 教育の情報化に対応した権利制限規定等の整備（第35条等関係）
3. 障害者の情報アクセス機会の充実に係る権利制限規定の整備（第37条関係）
4. アーカイブの利活用促進に関する権利制限規定の整備等（第31条，第47条，第67条等関係）

このうち2番目については2019年初の施行ではなく「公布の日から起算して3年を超えない範囲内において政令で定める日」になる模様。
どうやら「[対面授業と同時中継の遠隔合同授業のための公衆送信](https://hon.jp/news/1.0/0/14387 "著作権の保護と制限の規定がもうすぐ変わる ～ 保護期間延長、非親告罪化、柔軟な権利制限、教育の情報化対応など、まとめて解説 – HON.jp News Blog")」以外の公衆送信を利用できるように補償金決める，という随分と生臭い話のようだ。
教育教材を利権でガチガチに固めるからこんなことになるんだよな。

今回は新たに規定を追加するというより，既存の規定を再構成したものという感じだろうか。
ソフトウェア・エンジニアから見た場合，1番目の改正が大きいだろう。
AI の学習データ素材などを柔軟に利用できるよう考えられているようだ。

はっきり言って「著作権の制限」は字面ほど自由ではなく条件が事細かく決められている。
ホンマにそろそろ抜本的に見直していただきたいところである。

## ブックマーク

- [著しく短縮して語る著作権延長問題の歴史と、これからどうなり、何をしていくのか　福井健策｜コラム | 骨董通り法律事務所 For the Arts](https://www.kottolaw.com/column/181102.html)
- [著作権の保護と制限の規定がもうすぐ変わる ～ 保護期間延長、非親告罪化、柔軟な権利制限、教育の情報化対応など、まとめて解説 – HON.jp News Blog](https://hon.jp/news/1.0/0/14387)
- [シンポジウム「著作権延長後の世界で、我われは何をすべきか」 | TPPの知的財産権と協議の透明化を考えるフォーラム](http://thinktppip.jp/?p=853)
- [アメリカにパブリックドメインが帰ってくることを祝してインターネット・アーカイブに寄付した - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20181210/reopeningpublicdomain)
- [CPTPPによる著作権保護期間延長への対応と年末年始のお休みについて](https://www.aozora.gr.jp/soramoyou/soramoyou2018.html#000503)
    - [これからの20年に向けて](https://www.aozora.gr.jp/soramoyou/soramoyou2018.html#000500)
    - [三島由紀夫や志賀直哉、川端康成らの作品が青空文庫で公開されるのは20年先に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1158484.html)

- [漫画は衰退しました]({{< ref "/remark/2018/01/manga-has-delined.md" >}})
- [出版業界の自殺に日本のネットを巻き込まないで欲しい]({{< ref "/remark/2018/03/suicide-of-publishing.md" >}})
- [「技術的保護手段」と「技術的利用制限手段」]({{< relref "./copy-control-and-access-control.md" >}})
- [2019年 公有化に関する2つの話題]({{< ref "/remark/2019/01/public-domain-material.md" >}})

[CC Licenses]: https://creativecommons.org/licenses/ "ライセンスについて - Creative Commons"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E8%91%97%E4%BD%9C%E6%A8%A9%EF%BC%92%EF%BC%8E%EF%BC%90-%E3%82%A6%E3%82%A7%E3%83%96%E6%99%82%E4%BB%A3%E3%81%AE%E6%96%87%E5%8C%96%E7%99%BA%E5%B1%95%E3%82%92%E3%82%81%E3%81%96%E3%81%97%E3%81%A6-NTT%E5%87%BA%E7%89%88%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC%E2%80%95%E3%83%AC%E3%82%BE%E3%83%8A%E3%83%B3%E3%83%88-%E5%90%8D%E5%92%8C-%E5%B0%8F%E5%A4%AA%E9%83%8E/dp/4757102852?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757102852"><img src="https://images-fe.ssl-images-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E8%91%97%E4%BD%9C%E6%A8%A9%EF%BC%92%EF%BC%8E%EF%BC%90-%E3%82%A6%E3%82%A7%E3%83%96%E6%99%82%E4%BB%A3%E3%81%AE%E6%96%87%E5%8C%96%E7%99%BA%E5%B1%95%E3%82%92%E3%82%81%E3%81%96%E3%81%97%E3%81%A6-NTT%E5%87%BA%E7%89%88%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC%E2%80%95%E3%83%AC%E3%82%BE%E3%83%8A%E3%83%B3%E3%83%88-%E5%90%8D%E5%92%8C-%E5%B0%8F%E5%A4%AA%E9%83%8E/dp/4757102852?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4757102852">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt>
	<dd>名和 小太郎</dd>
    <dd>エヌティティ出版 2010-06-24</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4757102852, EAN: 9784757102859</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">名著です。今すぐ買うべきです。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-13">2018-11-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/18%E6%AD%B3%E3%81%AE%E8%91%97%E4%BD%9C%E6%A8%A9%E5%85%A5%E9%96%80-%E3%81%A1%E3%81%8F%E3%81%BE%E3%83%97%E3%83%AA%E3%83%9E%E3%83%BC%E6%96%B0%E6%9B%B8-%E7%A6%8F%E4%BA%95%E5%81%A5%E7%AD%96-ebook/dp/B00SM7G6SI?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00SM7G6SI"><img src="https://images-fe.ssl-images-amazon.com/images/I/41ZC-Qu61LL._SL160_.jpg" width="98" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/18%E6%AD%B3%E3%81%AE%E8%91%97%E4%BD%9C%E6%A8%A9%E5%85%A5%E9%96%80-%E3%81%A1%E3%81%8F%E3%81%BE%E3%83%97%E3%83%AA%E3%83%9E%E3%83%BC%E6%96%B0%E6%9B%B8-%E7%A6%8F%E4%BA%95%E5%81%A5%E7%AD%96-ebook/dp/B00SM7G6SI?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00SM7G6SI">18歳の著作権入門 (ちくまプリマー新書)</a></dt>
	<dd>福井健策</dd>
    <dd>筑摩書房 2015-01-10 (Release 2015-01-30)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00SM7G6SI</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">福井健策弁護士による「<a href="http://japan.cnet.com/sp/copyright_study/">18歳からの著作権入門</a>」の書籍化。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-11-13">2018-11-13</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
