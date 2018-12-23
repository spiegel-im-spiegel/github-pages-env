+++
title = "プログラムの著作物も「ダウンロード違法化」するつもりかしら"
date = "2018-12-11T18:03:27+09:00"
update = "2018-12-12T10:28:57+09:00"
description = "そうすることに何の意味があるのか誰か教えてくれ！"
image = "/images/attention/kitten.jpg"
tags = ["code", "law", "intellectual-property", "copyright", "access-control"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

なんか Facebook の TL でキナ臭い話が聞こえてきたので覚え書きとして記しておく。

なお私は私事で忙しく，この件に関して考察する時間的・精神的余裕がないのであしからず。
まぁ簡単に「（合法・非合法に関わらず）[著作物へのアクセス規制は著作権の基本理念に反する]({{< ref "/remark/2018/11/copy-control-and-access-control.md" >}} "「技術的保護手段」と「技術的利用制限手段」")」とだけ言っておこう。
今さら言っても愚痴だけどさ。

## 「[文化審議会著作権分科会法制・基本問題小委員会中間まとめに関する意見募集の実施について](http://search.e-gov.go.jp/servlet/Public?CLASSNAME=PCMMSTDETAIL&id=185001021&Mode=0)」

このうちの「{{< pdf-file title="文化審議会著作権分科会法制・基本問題小委員会中間まとめ" link="http://search.e-gov.go.jp/servlet/PcmFileDownload?seqNo=0000180847" >}}」から見出しだけ拾っておこう。

1. リーチサイト等を通じた侵害コンテンツへの誘導行為への対応
2. ダウンロード違法化の対象範囲の見直し
3. アクセスコントロール等に関する保護の強化
4. 著作権等侵害訴訟における証拠収集手続の強化
5. 著作物等の利用許諾に係る権利の対抗制度の導入
6. 行政手続に係る権利制限規定の見直し（地理的表示法・種苗法関係）
7. その他（改正著作権法第４７条の５第１項第３号の規定に基づく政令のニーズ）

### 「リーチサイト等を通じた侵害コンテンツへの誘導行為への対応」について

「リーチサイト」というのは

{{< fig-quote title="文化審議会著作権分科会法制・基本問題小委員会中間まとめ" link="http://search.e-gov.go.jp/servlet/PcmFileDownload?seqNo=0000180847" >}}
<q>自身のウェブサイトにはコンテンツを掲載せず，他のウェブサイトに蔵置された著作権侵害コンテンツへのリンク情報等を提供して利用者を侵害コンテンツへ誘導するためのウェブサイト</q>
{{< /fig-quote >}}

を指すらしい。
（外部の）違法コンテンツへのリンクが違法になるか，という話は昔から議論されているが，欧州では既に違法として取り扱われているため，これに倣ったものと思われる。

- [欧州司法裁、侵害コンテンツへのリンクを著作権侵害と判断 – P2Pとかその辺のお話R](https://p2ptk.org/copyright/553)

まっ，分かりやすいポリシー・ロンダリングやね。

### 「ダウンロード違法化の対象範囲の見直し」について

「ダウンロード違法化の対象範囲の見直し」というのは既に「静止画ダウンロード違法化」として話題になっている。
ぶっちゃけて言うと，かねてから要望されていたサイト・ブロッキングの法案化が無理ぽかったため「ならダウンロード違法化で対応しろやゴラァ」ということのようだ。
既に前例があるなら通りやすいと思うよね，普通。

ここで言う「ダウンロード」とは

{{< fig-quote title="文化審議会著作権分科会法制・基本問題小委員会中間まとめ" link="http://search.e-gov.go.jp/servlet/PcmFileDownload?seqNo=0000180847" >}}
<q>現行著作権法第３０条第１項第３号に規定する「・・自動公衆送信・・を受信して行うデジタル方式の録音又は録画」について，対象著作物を音楽・映像以外にも広げたもの，すなわち「自動公衆送信を受信して行うデジタル方式の複製（注：有形的再製一般）」を指すものであり，例えば，ウェブサイトに掲載されたテキストをプリントアウトする行為や，そこでプリントアウトされたものを更にＰＤＦ化してコンピュータに保存する行為等を含むものではない。</q>
{{< /fig-quote >}}

とのことで[^cl1]，直に保存する（スクリーンショット等を含む）のはダメで，いったん紙に落とせばいいらしい。
なんじゃそら。
そうすることに何の意味があるのか誰か教えてくれ！

[^cl1]: [著作権法]第30条とは「著作権の制限」のうち「私的使用のための複製」について記されたもの。ダウンロード違法化は「私的使用のための複製」の例外として定義されているわけだ。

今回のまとめには

{{< fig-quote title="文化審議会著作権分科会法制・基本問題小委員会中間まとめ" link="http://search.e-gov.go.jp/servlet/PcmFileDownload?seqNo=0000180847" >}}
<q>また，関係団体からプログラム（ビジネスソフト・ゲーム）に関しても違法ダウンロード等による被害が継続的に生じているとの報告があったほか，多数の学術論文の全文を無料でダウンロードできる論文版海賊版サイトの存在が明らかとなるなど，幅広い分野の著作物について，違法にアップロードされた著作物のダウンロードによる被害が一定程度生じていることが確認された。</q>
{{< /fig-quote >}}

などと書かれていて，静止画やテキストのみならず「プログラムの著作物」も含めるつもりなんじゃないか，という噂もあるらしい。

なお「ダウンロード違法化の対象範囲の見直し」については参考資料が添えられている。

- {{< pdf-file title="「ダウンロード違法化の対象範囲の見直し」に関する留意事項" link="http://search.e-gov.go.jp/servlet/Public?CLASSNAME=PCMMSTDETAIL&id=185001021&Mode=0" >}}

念入りなことです。

## ブックマーク

- [文化審議会著作権分科会法制・基本問題小委員会（第6回） | 文化庁](http://www.bunka.go.jp/seisaku/bunkashingikai/chosakuken/hoki/h30_06/)
    - {{< pdf-file title="ダウンロード違法化の対象範囲の見直しに関する論点について" link="http://www.bunka.go.jp/seisaku/bunkashingikai/chosakuken/hoki/h30_06/pdf/r1411529_01.pdf" >}}

- [文化庁文化審議会法制・基本問題小委員会で静止画ダウンロード規制に関して意見を述べました - MIAU](https://miau.jp/ja/880)
- [第４０１回：ダウンロード違法化・犯罪化の対象範囲の拡大とリーチサイト規制（リンク規制）の法制化を含む文化庁・著作権分科会・法制・基本問題小委員会中間まとめに関する意見募集の開始（２０１９年１月６日〆切）: 無名の一知財政策ウォッチャーの独言](http://fr-toen.cocolog-nifty.com/blog/2018/12/post-ef1a.html)
- [静止画や小説等ダウンロードの違法化／処罰化に強く反対する: 弁護士山口貴士大いに語る](http://yama-ben.cocolog-nifty.com/ooinikataru/2018/12/post-8df7.html)

[著作権法]: http://elaws.e-gov.go.jp/search/elawsSearch/elaws_search/lsg0500/detail?lawId=345AC0000000048 "著作権法"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/B01CYDGUV8/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/31Q2jh%2B5SgL._SL160_.jpg" width="113" height="160" alt="CODE VERSION 2.0"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/B01CYDGUV8/baldandersinf-22">CODE VERSION 2.0</a></dt>
    <dd>ローレンス・レッシグ</dd>
    <dd>翔泳社</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="far fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">前著『CODE』改訂版。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.11.17</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/41YkbcP5IyL._SL160_.jpg" width="108" height="160" alt="著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4757102852/baldandersinf-22">著作権２．０ ウェブ時代の文化発展をめざして (NTT出版ライブラリー―レゾナント)</a></dt>
    <dd>名和 小太郎</dd>
    <dd>エヌティティ出版</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">名著です。今すぐ買うべきです。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.11.13</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/exec/obidos/ASIN/4757122349/baldandersinf-22"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ftPU2g7FL._SL160_.jpg" width="112" height="160" alt="〈反〉知的独占　―特許と著作権の経済学"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/exec/obidos/ASIN/4757122349/baldandersinf-22">〈反〉知的独占　―特許と著作権の経済学</a></dt>
    <dd>ミケーレ・ボルドリン, デイヴィッド・Ｋ・レヴァイン</dd>
    <dd>エヌティティ出版</dd>
    <dd>評価&nbsp;<span class="fa-sm" style="color:goldenrod;">
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="fas fa-star"></i>
      <i class="far fa-star"></i>
    </span></dd>
  </dl>
  <p class="description">「知的財産権は、人類進歩を阻害する！」（帯文より）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed">2018.11.17</abbr> (powered by <a href="https://dadadadone.com/amakuri/" >Amakuri</a>)</p>
</div>
