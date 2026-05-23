+++
title = "『カーニハンのUNIX回顧録』オンライン読書会第5回目"
date =  "2026-05-23T16:03:30+09:00"
description = "『カーニハンのUNIX回顧録』オンライン読書会第5回のメモ。Unix 年表，Xenix，Minix/Linux，Unix 哲学を振り返る。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "book", "unix", "engineering", "linux" ]
pageType = "text"
draft = false

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

『[カーニハンのUNIX回顧録]』オンライン読書会の第5回目に参加してきた。

- [第5回『カーニハンのUNIX回顧録』オンライン読書会 (2026/05/23 13:00〜)](https://technical-book-reading.connpass.com/event/392162/)

今回は7章から9章の途中まで。
最後まで終わらなかった。
印象に残った部分をかいつまんで紹介する。

## Unix 年表

8章に載っている「Unix 年表」は Wikipedia からの転載みたいなので探してみたらあった。

{{< fig-img class="lightmode" src="./Unix_timeline.en.svg.png" title="File:Unix timeline.en.svg - Wikimedia Commons (public domain)" link="https://commons.wikimedia.org/wiki/File:Unix_timeline.en.svg" width="952" lang="en" >}}

すべての Unix 派生版が載っているわけではないが，大まかな系譜は見て取れるだろう。

研究所版 Unix が系譜の根であり，そこから大きくバークレー版（Berkeley Software Distribution; BSD）と商用版（System III & V）に分かれる。
両者は同じ祖先を持ちながら互換性がなく，また AT&T が BSD に対して知財権侵害で訴えたこともあり（BSD は AT&T からのコードを削除・書き換え，最終的に勝訴する）まるで似て非なるものとなった。

### Xenix

読んで面白かったのは Microsoft が1980年代に Xenix と呼ぶ Unix 派生バージョンを供給してたことで，1980年代後半の Unix インストールマシン台数ではこれが一番多かったらしい。

{{< fig-quote type="markdown" title="『カーニハンのUNIX回顧録』 p.172" link="https://www.maruzen-publishing.co.jp/book/b10152370.html" >}}
もし Microsoft が自身の MS-DOS の代わりに Xenix を推していたら，そして，もし AT&T が見かけによらず取引しやすかったなら，今日の世界はどれだけ違っただろうか．
{{< /fig-quote >}}

ホンマにな。
Xenix はのちに SCO (Santa Cruz Operation) に売却される。

### Minix と Linux

AT&T と BSD の間の不毛な争いの間に「もっと自由に使える Unix」の機運が高まる。
ここで登場したのが Minix である。
Minix はシステムコールレベルでは互換だがゼロからコードを書いたもので，カーネルの構成が Unix とは異なる。
作者の Andrew Tanenbaum は Minix についての教科書を書き，ソースコードを無料公開した。
Minix は教育・研究目的で今も健在らしい。

そして Linus Torvalds による Linux が登場する。
この辺までくると私にもリアルな記憶になってくるな。

『[カーニハンのUNIX回顧録]』には Usenet の `comp.os.minix` に投稿された Linus Torvalds 投稿（の翻訳）が載っている。

{{< fig-quote type="markdown" title="『カーニハンのUNIX回顧録』 p.175" link="https://www.maruzen-publishing.co.jp/book/b10152370.html" >}}
私は，386 (486) AT クローン向けの（無料の）オペレーティングシステムを作っています（単なる趣味で，gnu のように大きく，プロ向けなものにはしたくない）．
{{< /fig-quote >}}

この部分を読んで密かに笑ってしまった。
フィンランドの大学生が「単なる趣味で」作った Unix もどきが（ゼロからコードを書き，最初は数千行のサイズだったらしい），今やパソコンを除く多くの機器でメジャーな OS となった（スパコンも今や Linux 駆動だし Android も Linux ベースである）。
もっと言えば [Linux を礼賛する本](https://amzn.to/3RFXanf "リナックスの革命 ― ハッカー倫理とネット社会の精神 | ペッカ ヒマネン, リーナス トーバルズ, マニュエル カステル, 安原 和見, 山形 浩生 |本 | 通販 | Amazon")も出たりして，世の中というのは実に不思議である[^l1]。

[^l1]: 私の印象でも1990年代までは「Linux はオモチャ」だった。それでも Unix ワークステーションの X 端末としては優秀だったので仕事でもよく使っていたが。

## Unix の技術がもたらしたもの

『[カーニハンのUNIX回顧録]』では Unix の技術がもたらしたものとして以下を挙げている。

- 階層型ファイルシステム
- 高水準実装言語
- ユーザレベルのプログラム可能なシェル
- パイプ
- ツールとしてのプログラム
- 普通のテキストが標準データ形式である
- プログラムを書くプログラム
- 専用言語

詳しくは９章を参照あれ。

### Unix 哲学

Unix の技術がもたらしたものとして最後に「コンピュータを使った仕事にどのように取り組むか」というのが挙げられている[^up1]。
曰く

[^up1]: “[BSTJ 57: 6. July-August 1978: UNIX Time-Sharing System: Forward. (McIlroy, M.D.; Pinson, E.N.; Tague, B.A.) : Free Download, Borrow, and Streaming : Internet Archive](https://archive.org/details/bstj57-6-1899/page/n3/mode/2up)” p.1902 参照。

{{< fig-quote type="markdown" title="『カーニハンのUNIX回顧録』 p.188" link="https://www.maruzen-publishing.co.jp/book/b10152370.html" >}}
1. それぞれのプログラムは一つのことをうまくやるようにせよ．新しい仕事をするには，新たな「機能」を追加して古いプログラムを複雑にするのではなく，改めて組み立てよ．
2. すべてのプログラムの出力は，未知の別のプログラムの入力となることを想定せよ．出力を無関係な情報で散らかさないこと．厳格なカラム化やバイナリの入力形式は避けよ．対話型の入力を強要しないこと．
3. ソフトウェアは，たとえそれがオペレーティングシステムであっても，早期に，できれば数週間以内に試用できるように設計し構築せよ．ぎこちない部分はためらわずに捨てて作り直すこと．
4. たとえツール構築のために迂回する必要があったり，使い終わったらその一部を捨てることになると予想されようとも，プログラミング作業を楽にするためには，未熟な人の手助けよりもむしろツールの使用を優先せよ．
{{< /fig-quote >}}

これは，今日では「[Unix 哲学（Unix Philosophy）](https://en.wikipedia.org/wiki/Unix_philosophy "Unix philosophy - Wikipedia")」として知られているものの，元になったものである。
4番目とか道具としての生成 AI を指してるような気がして，ちょっと面白かった。

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[前回]: {{< ref "/remark/2026/04/reading-unix-a-history-and-a-memoir-4/" >}} "『カーニハンのUNIX回顧録』オンライン読書会第4回目"
[カーニハンのUNIX回顧録]: https://www.maruzen-publishing.co.jp/book/b10152370.html "カーニハンのUNIX回顧録 - 丸善出版 理工・医学・人文社会科学の専門書出版社"

## 参考

{{% review-paapi "4309242456" %}} <!-- リナックスの革命 Hacker Ethic -->
{{% review-paapi "4796880011" %}} <!-- それがぼくには楽しかったから -->

<!-- eof -->
