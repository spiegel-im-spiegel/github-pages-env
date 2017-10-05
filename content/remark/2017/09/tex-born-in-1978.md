+++
title = "TeX の思ひ出"
date =  "2017-09-29T14:50:20+09:00"
update =  "2017-10-05T17:51:09+09:00"
description = "TeX の最初のバージョンって1978年なのか。それってもうすぐ40周年ぢゃん！"
tags        = [ "tex" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = true
  mermaidjs = false
+++

$\\mathrm{\TeX}$ について調べ物をしていて，何気なく [Wikipedia のページ](https://ja.wikipedia.org/wiki/TeX "TeX - Wikipedia")を見たら， $\\mathrm{\TeX}$ の最初のバージョンって1978年なのか。
それってもうすぐ40周年ぢゃん！

私は学生時代はロクに勉強もしない不良学生だったので $\\mathrm{\TeX}$ とは無縁だったのだが，バブルの恩恵でたまたま潜り込んだシステムハウスで某社製レーザープリンタ用の $\\mathrm{\TeX}$ ドライバと出会い，それが $\\mathrm{\TeX}$ というか $\\mathrm{\LaTeX}$ にハマるきっかけになった。

大昔の話だし，この程度の内容なら時効だよな，もう。

----

当時は LIPS (LBP Image Processing System) 制御の某企業のレーザープリンターが市場を席巻していた頃で，その中で件のレーザープリンタは比較的低価格[^lp] ながら PostScript を内蔵してるってんで「なら $\\mathrm{\TeX}$ ドライバと抱き合わせで売ったらいんじゃね？」というかなり打算的な考えだったような気がする。
勤めていた会社が大学方面のコネクションが弱かったので繋ぎを付けたかったというのもあったような。
実際，このプリンタは $\\mathrm{\TeX}$ ドライバとセットでそこそこ売れたらしい。

[^lp]: つっても当時のレーザープリンタは個人では買えないくらい高かったが。しかも PostScript プリンタは目玉が飛び出るくらい高かった。まぁ，まだバブルの名残が残ってたしねぇ。

その $\\mathrm{\TeX}$ ドライバは Pascal で書かれてたんだけど，オリジナルの作者の方は所有権を既に手放していて，技術サポートとして私にお鉢が回ってきたのだった。

いや，私， **Pascal なんて（仕事では）やったことないんだってば！**[^dp]

[^dp]: ちなみに私が [Delphi](https://ja.wikipedia.org/wiki/Delphi "Delphi - Wikipedia") で遊びだしたのは後のことである。

幸いなことにソフトウェアに関しては完成度が非常に高く，私自身は内部の深いところまでは触らずに済んだのだが，技術サポートの人がソフトウェアを触ったことがないなんてありえないので，私自身も $\\mathrm{\TeX}$ に手を出すことになってしまった。

しかも当時は DOS で 日本語 $\\mathrm{Micro\TeX}$ である。
フロッピーディスクでは運用できず（当時はまだ高かった）ハードディスクを専用に買ってもらったっけ。
しかし，件のプリンタと $\\mathrm{\TeX}$ ドライバによる出力結果は素晴らしいの一言で，ワープロ[^wp1] なんか目じゃない仕上がりに一目惚れしてしまったのだ。

[^wp1]: 当時の日本語ワープロといえば圧倒的に「一太郎」だったが （Windows 95 が登場してからは Microsoft Word も健闘してたけど当時の Word はとにかく日本語処理が下手くそだった），紙への出力に関しては私のような素人が見てもカスみたいな品質だった。

そこからはホントにどハマリしたね。
フォーマット・フリーの社内文書はほとんど $\\mathrm{\LaTeX}$ で書いた。
議事録すら $\\mathrm{\LaTeX}$ で書いたことがあって，お客さんに「どこの論文かと思った」と皮肉を言われたのも今は良い思い出である（笑）
ちなみに今でも履歴書は $\\mathrm{\LaTeX}$ で[書いている](https://www.tamacom.com/rireki-j.html "履歴書スタイルファイル")。

当時に社内で書いた文書の一部は[本家サイト](http://www.baldanders.info/ "Baldanders.info")でも公開している。

- {{< pdf-file title="文字コードとその実装" link="http://www.baldanders.info/spiegel/archive/charset-pdfa.pdf" >}} （[GitHub](https://github.com/spiegel-im-spiegel/charset_document "spiegel-im-spiegel/charset_document: 「文字コードとその実装」 upLaTeX ドキュメント")）
- {{< pdf-file title="SH マイコン C プログラミング虎の巻" link="http://www.baldanders.info/spiegel/archive/ProgSH.pdf" >}} （[GitHub](https://github.com/spiegel-im-spiegel/progSH_document "spiegel-im-spiegel/progSH_document: SHマイコン Cプログラミング虎の巻 (pLaTeX ドキュメント")）

その後，証券会社が倒産するほどの不況のずんどこが訪れ，勤務先もその煽りでプリンタ販売業務を止めてしまったが，私は $\\mathrm{\TeX}$ で遊び続けた。
最初の頃のバイブルだった『てくてく$\\mathrm{\TeX}$』とか『$\\mathrm{\LaTeX}$入門』とかいまだに本棚に並んでるですよ[^bs1]。

[^bs1]: 先日，納戸にしまった本を発掘してたら [$\\mathrm{\TeX}$ 関連の本やらパッケージが大量に出てきて](https://www.instagram.com/p/BZs5cUuHDNO/)我ながら笑ってしまった。どんだけ好きだったんだろう（笑）

{{< fig-img src="https://farm5.staticflickr.com/4448/36713793483_453ca1483a.jpg" title="私はコレで TeX を覚えましたw"  link="https://www.flickr.com/photos/spiegel/36713793483/" >}}

そういや昔は，テープを物理的に「回覧」して配布してたんだよねぇ。
私はテープを扱えるデバイスを（私的には）持ってなかったので回覧に参加できなかったんだけど。

そのうち FTP や HTTP で大量にダウンロードできる環境が整って（ADSL 万歳！），長いあいだ一番お世話になったのはいわゆる「[角藤版](http://w32tex.org/index-ja.html "W32TeX")」だった（今でも成果については大変恩恵を受けてますが）。

[TeX Live](http://www.tug.org/texlive/ "TeX Live - TeX Users Group") を利用するようになったのは[2013年](http://www.baldanders.info/spiegel/log2/000640.shtml "TeX Live 2013 のインストールに挑戦 — Baldanders.info")から。
いや，めっさ簡単になったものだよ。

しかし，最近いちばん衝撃を受けたのは [LuaTeX-ja](https://ja.osdn.net/projects/luatex-ja/wiki/FrontPage) か。
だって日本語環境でもようやく直接 PDF/A を吐ける処理系が登場したんだよ。

- [TeX 覚え書き（upLaTeX から PDF/A まで） — Baldanders.info](http://www.baldanders.info/spiegel/log2/000731.shtml)
- [LuaTeX-ja に関する覚え書き]({{< relref "remark/2015/luatex-ja.md" >}})

これによってようやく日本語の $\\mathrm{\TeX}$ 処理系でストレスなく PDF/A 形式のデジタル文書を作れるようになったわけだ。

----

まぁでも $\\mathrm{\TeX}$ 処理系は使わなくなったね。
理由は簡単で，デジタル文書が「紙」の制約から離れ始めたから。
それでも過去の成果は継承されていくし [MathJax] のように形を変えて残っていくものもあるんだけど。

昔は「完成図書」をどうやって作成・維持するかが重要だった。
だから最終出力が「紙」である（または「紙」を意識する）ことに意味があったわけだ。

でも今は違う。
そもそも「完成図書」など存在し得ない。
何故なら，プログラムと同様に，文書もまた常に更新される（ベータ版だ）から。
デジタル以前は更新コストが高かっただけで，もともと完全な文書など存在しないのだ（「コーディング」が紙テープやパンチカードから端末上での直接編集に変わって「コーダー」という職業が消滅したのと相似形かもしれない）。

しかも更新履歴は git などのバージョン管理ツールで自動化できる。
つまり（プログラムコードもそうだけど）今や文書ってのは「時間」というベクトルを加えた存在になりつつあるわけだ。
文書が「「紙」の制約から離れ」るってのはそういうことなのだ。

ひとつのソフトウェア・システムが40年も継続して存在するというのは凄いことだ。
「文書」の意味が変わりつつある中で，これからの40年で $\\mathrm{\TeX}$ がどう変わっていくのか。
それとも無くなってしまうのか。

## ブックマーク

- [TeXの歴史 - TeX Wiki](https://texwiki.texjp.org/?TeX%E3%81%AE%E6%AD%B4%E5%8F%B2)

[MathJax]: https://www.mathjax.org/

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51E5K7B53aL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4774187054/baldandersinf-22/">[改訂第7版]LaTeX2ε美文書作成入門</a></dt><dd>奥村 晴彦 黒木 裕介 </dd><dd>技術評論社 2017-01-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798118141/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798118141.09._SCTHUMBZZZ_.jpg"  alt="LaTeX2e辞典~用法・用例逆引きリファレンス (DESKTOP REFERENCE)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4535558752/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4535558752.09._SCTHUMBZZZ_.jpg"  alt="公共政策入門 ミクロ経済学的アプローチ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4320112415/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4320112415.09._SCTHUMBZZZ_.jpg"  alt="Rで楽しむ統計 (Wonderful R 1)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298550/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298550.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.5"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4797391383/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4797391383.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/積分を見つめて (数学ガールの秘密ノートシリーズ)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4000298569/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4000298569.09._SCTHUMBZZZ_.jpg"  alt="岩波データサイエンス Vol.6"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4798115363/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4798115363.09._SCTHUMBZZZ_.jpg"  alt="独習 LaTeX2ε"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4785315717/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4785315717.09._SCTHUMBZZZ_.jpg"  alt="具体例から学ぶ 多様体"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774193046/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774193046.09._SCTHUMBZZZ_.jpg"  alt="【改訂第3版】基礎からわかる情報リテラシー"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4768704700/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4768704700.09._SCTHUMBZZZ_.jpg"  alt="はじめて学ぶリー群 ―線型代数から始めよう"  /></a> </p>
<p class="description">ついに第7版が登場。紙の本で買って常に側に置いておくのが吉。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2017-09-27">2017-09-27</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
