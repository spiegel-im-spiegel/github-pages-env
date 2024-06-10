+++
title = "Kagi Search を試してみる 〜 検索サービスも有料の時代？"
date =  "2024-06-07T21:26:58+09:00"
description = "日本は現在，自国通貨の価値が相対的に低いので国外の有料サービスは微妙に勧め辛いのがネックではあるが，興味が湧いたのなら選択肢のひとつとしてどうぞ。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "search", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## Kagi Search とは

[Kagi Search] なる検索サービスがあるらしい。
知ったきっかけはSF作家 Cory Doctorow の以下の記事。

- [Pluralistic: Too big to care (04 Apr 2024) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2024/04/04/teach-me-how-to-shruggie/)
- [なぜGoogleは“あなたの不満”を無視できるのか | p2ptk[.]org](https://p2ptk.org/monopoly/4535)

この記事で [Kagi Search] は以下のように評価されている。

{{< fig-quote type="markdown" title="なぜGoogleは“あなたの不満”を無視できるのか" link="https://p2ptk.org/monopoly/4535" >}}
私も試してみた。魔法だった。

いや、マジで。Googleではもう見つからなくなったものは？ 全部検索結果の最上位に出てくる。Googleで検索したらスパムがページ単位で出てくるような検索ワードなら？ クッソキレイな検索結果じゃないか。Kagiは何度でも、適切な検索結果を吐き出した。

ちょっと触っただけでも感動モノだった。さらにKagiのレンズや高度な機能を使うと、検索体験は「魔法」から「魔術」のレベルに上がった。

[...]

すぐにファミリープランに加入した。1ヶ月使ってみた。Google検索は基本的に使わなくなった。
{{< /fig-quote >}}

大絶賛ぢゃん（笑）

日本語圏で [Kagi Search] に関する記事がないかな，と探してみた。
この辺かな。

- [Kagi Searchをメインの検索エンジンとして使っている | Web Scratch](https://efcl.info/2024/03/15/kagi-search/)

この記事によると，利点としては

{{< fig-quote type="markdown" title="Kagi Searchをメインの検索エンジンとして使っている" link="https://efcl.info/2024/03/15/kagi-search/" >}}
- Googleより良いと感じる検索結果が出しやすい(人による) 
  - 日本語の検索結果はそこまで変わらない感じもする(後述する漢字だけ検索した時の問題はまだある) 
  - 英語で検索した時に英語のリソースにマッチさせやすいので良い結果と感じることが多い 
  - また、ビルトインでBlockや検索結果の優先度を変更する機能が入ってる
  - [Redirects (URL Rewrites) | Kagi’s Docs](https://help.kagi.com/kagi/features/redirects.html) という機能で、検索結果のURLを書き換える機能も持っている
- Lensesで特定のサイトからの検索結果だけにフィルターできるのが便利
  - プログラミング関係(GitHubやStackoverflowなど)のサイトだけに絞ったりが、1 clickで切り替えできる
  - レビューとか検索したいときに個人のブログを検索したい といった感じの用途でよく使う
- 検索結果のOrder ByとTimeが素直な感じ
  - Order By: recentで検索のソート順を新しいもの順にできる
  - Timeで24時間以内の結果にできる(ここはGoogleでもできる)
  - Verbatim modeを使うと、クエリの文字列が含まれているサイトだけが表示される
{{< /fig-quote >}}

という感じらしい（[原文記事](https://efcl.info/2024/03/15/kagi-search/ "Kagi Searchをメインの検索エンジンとして使っている | Web Scratch")の箇条書きを抜粋して列挙してるので詳しくは原文をどうぞ）。

[Kagi Search] は有料のサービスで料金体系は以下のようになっている。

{{< fig-img src="./plans.png" title="Kagi Search Pricing and Plans" link="https://kagi.com/pricing" width="1432" >}}

お試しで100回までは無料でいけるようだ。
検索クエリを無制限で使いたいなら Professional (10USD/月 or 108USD/年 ＋ 税金) 以上で契約する必要がある。

## Kagi Search を試してみる

（私の環境では [uBlock Origin](https://ublockorigin.com/ "uBlock Origin - Free, open-source ad content blocker.") で広告をブロックしているので，その辺を割り引いてね）

取り敢えず無料の Trial プランでサインアップし「Kagi Search とは」で検索してみた。

{{< fig-img src="./search-by-kagi-1.png" title="「Kagi Search とは」検索結果 by Kagi" link="./search-by-kagi-1.png" width="834" >}}

比較のために Google で検索した結果を挙げる。

{{< fig-img src="./search-by-google-1.png" title="「Kagi Search とは」検索結果 by Google" link="./search-by-google-1.png" width="857" >}}

あー，要らん要約が付いてるな。
`&udm=14` を付加して再表示する。

{{< fig-img src="./search-by-google-2.png" title="「Kagi Search とは」検索結果 by Google (要約なし)" link="./search-by-google-2.png" width="857" >}}

これなら多少は見れるかな。
Forbes の「[ついにGoogleをしのいだ？ 5つの未来型検索エンジン](https://forbesjapan.com/articles/detail/49605 "ついにGoogleをしのいだ？ 5つの未来型検索エンジン | Forbes JAPAN 公式サイト（フォーブス ジャパン）")」が上位にないのがいい感じだな（笑）

ついでに DuckDuckGo で検索した結果も挙げておこう。

{{< fig-img src="./search-by-duckduckgo.png" title="「Kagi Search とは」検索結果 by DuckDuckGo" link="./search-by-duckduckgo.png" width="834" >}}

んー。
細かい違いはあるにせよ，やっぱ日本語圏では違いが分かりにくいか？ 

今度は “Order by” を “Recency” にしてみよう。
時系列に並べ替えるわけやね。

{{< fig-img src="./search-by-kagi-2.png" title="「Kagi Search とは」検索結果 by Kagi (order by Recency)" link="./search-by-kagi-2.png" width="834" >}}

おー。
だいぶ変わるんだな。

「[Kagi Searchをメインの検索エンジンとして使っている](https://efcl.info/2024/03/15/kagi-search/ "Kagi Searchをメインの検索エンジンとして使っている | Web Scratch")」では日本の主要ブログに絞り込むレンズを紹介している。

{{< fig-quote type="markdown" title="Kagi Searchをメインの検索エンジンとして使っている" link="https://efcl.info/2024/03/15/kagi-search/" >}}
次の[Lenses](https://help.kagi.com/kagi/features/lenses.html)を使って、日本のブログだけに絞り込める

- https://kagi.com/lenses/0Q9bHFmidnH3TfNAR3OYQKb0gyqDEzM7
{{< /fig-quote >}}

この設定を取り込んでみる。

{{< fig-img src="./edit-lens.png" title="Edit Lens" link="./edit-lens.png" width="1096" >}}

以下のサイトのみに絞り込む設定らしい。

```text
*.hatenablog.com, *.hatenablog.jp, *.hateblo.jp, *.hatenadiary.com, *.hatenadiary.jp, note.com, ameblo.jp, sizu.me, zenn.com, qiita.com
```

この設定を保存して検索してみる。

{{< fig-img src="./search-by-kagi-3.png" title="「Kagi Search とは」検索結果 by Kagi (with Lens:ブログ)" link="./search-by-kagi-3.png" width="840" >}}

なるほど。
こんな感じになるのか。
こりゃあカスタマイズしがいがあるな。

## 検索サービスも有料の時代？

最初に述べたように [Kagi Search] について知ったのは4月のGW前だったのだが「そのうち試してみるか」という感じだった。
私は普段 DuckDuckGo を使ってるのだが，日本語の情報を探す限りにおいては特に不満がなかったからだ。

ところが先月5月に Microsoft のサービスで障害が発生し，その余波で DuckDuckGo も使えなくなった。

- [Microsoftの大規模障害によりChatGPT・Copilot・Bing・DuckDuckGoなどがダウン - GIGAZINE](https://gigazine.net/news/20240524-microsoft-outage-affects-chatgpt/)

このとき久しぶりに Google 検索を使ったのだが，欲しい情報に全然当たらないのね。
大人気なく癇癪を起こしちゃったよ。
そんなわけで重い腰を上げて [Kagi Search] へサインアップしてみたのだった。

Google 検索サービスの「メタクソ化（[enshittification](https://en.wikipedia.org/wiki/Enshittification "Enshittification - Wikipedia")）」は許容できるレベルを超えつつある。
「メタクソ化」の造語を作った Cory Doctorow は，最初に挙げた記事以外にも

- [Pluralistic: Google reneged on the monopolistic bargain; The Bezzle excerpt (Part IV) (21 Feb 2024) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2024/02/21/im-feeling-unlucky/)
  - [かくしてGoogleはスパマーに敗北した | p2ptk[.]org](https://p2ptk.org/monopoly/4515)
- [Pluralistic: The specific process by which Google enshittified its search (24 Apr 2024) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2024/04/24/naming-names/)
  - [Google検索を殺した男――Googleはいつ、どこでメタクソ化に舵を切ったのか | p2ptk[.]org](https://p2ptk.org/monopoly/4541)

などとボロクソに書いている。

{{< fig-quote type="markdown" title="Google検索を殺した男――Googleはいつ、どこでメタクソ化に舵を切ったのか" link="https://p2ptk.org/monopoly/4541" >}}
ジットロンは、検索の質をめぐる取締役会の闘争の物語を綴っている。その取締役会で、Googleの全盛期を支えた古参Googlerであるベン・ゴームズが、コンピュータサイエンティストからマネージャーに転身したプラバカール・ラガバンとの闘争に敗れたのだ。ラガバンの戦術は、検索の質を下げることで検索クエリの数（つまり、同社が検索者に表示できる広告の数）を増やすこと（訳注：「エンゲージメント・ハッキング」）だった。そうすれば、ユーザは目的のものを見つけるまでにGoogleでより長く時間を過ごさなければならなくなる。
{{< /fig-quote >}}

とか書いてあって納得しちゃったよ（笑）

更に spam 広告に関しては2022年末の時点で [FBI が警告]({{< ref "/remark/2022/12/ad-blocker.md" >}} "米国 FBI は広告ブロッカーを推奨している？")を出してるし，最近でも

- [「Google広告からの誘導が6割」との分析結果。より巧妙化し、高齢者を狙う「サポート詐欺」に注意！【被害事例に学ぶ、高齢者のためのデジタルリテラシー】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/column/dlis/1592999.html)

みたいな話がある。

主要な検索サービスにおいて検索結果の並び順は検閲のように働く。
2022年に DuckDuckGo はロシアの一部のサイトについて政治的理由から検索結果のランクを引き下げた。
この措置に対して検閲であるとの批判が上がったらしい。

- [DuckDuckGo will demote Russian propaganda in search results](https://www.engadget.com/duck-duck-go-reverses-course-will-demote-russian-propaganda-in-search-results-014336389.html)

そういやゼロ年代には「Google 八分」とかあったな（若い人は知らんかな）。

私は「**目的の如何にかかわらずユーザがコントロールできないフィルタリングはすべて検閲である**」という不遜な考えを持っている。
検索結果における（ユーザがコントロールできない）恣意的な順位操作は（フィルタリングではないので）検閲とまでは言わないが UX 的には「[ダークパターン]({{< ref "/remark/2023/04/the-psychology-of-video-games.md" >}} "『はじめて学ぶ ビデオゲームの心理学』は読んどけ！")」なんじゃないかと思っている。

[Kagi Search] は検索結果についてかなりの部分をユーザがコントロールできるっぽい。
有料サービスなので検索結果のトップが広告で溢れかえることもないだろうし「検索の質を下げることで検索クエリの数を増やす」などというアホなことはしでかさないだろうという期待感もある。

また [Kagi Search] は自前のクローラーや巨大なデータセンターを持っているわけではなく

{{< fig-quote type="markdown" title="なぜGoogleは“あなたの不満”を無視できるのか" link="https://p2ptk.org/monopoly/4535" >}}
Kagiで検索すると、Google、Yandex、Mojeek、Braveなどの従来の検索インデックスや、Wikimedia Commons、Flickrなどの少数の専門検索エンジンに「匿名化されたAPIコール」が行われる。Kagiはこれを独自のウェブインデックスとニュースインデックス（ニュース検索用）と組み合わせて、表示される検索結果ページを構築する。つまり基本的には、Google検索結果と他のインデックスからの検索結果とミックスしているのだ。
{{< /fig-quote >}}

なんだそうだ（引用の引用ですまん）。
つまり，ユーザから見れば VRM (Vendor Relationship Management) として機能する 4th party と見なすこともできるわけだ（流石に褒め過ぎか？）。
本当にそうならお金を払う価値はあるかもしれない。
まぁ，細かい評価はこれからなので「やっぱあかん」ってなるかも知れんけど。

日本は現在，自国通貨の価値が相対的に低いので国外の有料サービスは微妙に勧め辛いのがネックではあるが，興味が湧いたのなら選択肢のひとつとしてどうぞ。

## ブックマーク

- [Google検索結果からAIによるまとめを排除するフィルタ「&udm=14」 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20240527/udm14)
- [Googleからウェブサイトへのトラフィックがゼロになる日 – WirelessWire News](https://wirelesswire.jp/2024/06/86654/)

- [広告の曲がり角]({{< ref "/remark/2023/11/blocking-ad-blocker.md" >}})

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"

## 参考図書

{{% review-paapi "4296001574" %}} <!-- ハッキング思考 -->
{{% review-paapi "4571210450" %}} <!-- はじめて学ぶ ビデオゲームの心理学 -->
{{% review-paapi "B00DIM6BE6" %}} <!-- インテンション・エコノミー -->
