+++
draft = false
date = "2017-04-06T22:25:55+09:00"
title = "エディタ以上ワープロ未満の HackMD"
tags = ["tools", "editor", "markdown"]
description = "もうしばらく遊んでみて，よさげなら常用してみようかな，と。"

[author]
  github = "spiegel-im-spiegel"
  flattr = "spiegel"
  linkedin = "spiegelimspiegel"
  url = "http://www.baldanders.info/spiegel/profile/"
  name = "Spiegel"
  flickr = "spiegel"
  instagram = "spiegel_2007"
  tumblr = "spiegel-im-spiegel"
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  twitter = "spiegel_2007"
  license = "by-sa"
+++

たとえばちょっとしたメモを取るとき。

仕事なら紙のノートで手書きで書く。
「ちょっとしたメモ」なら手書きで走り書きの方がまだ速い（キーボード打ちながらメモを書けるし）。
その代り私の走り書きは酷い悪筆で私自身も読めないことがあるのが玉に瑕である（笑）

仕事以外のことで目の前に紙のノートも PC もない場合は [Simplenote] を愛用するようになった。
[Simplenote] は SaaS 型のテキストエディタである。
Web 版のほか各種携帯端末用のアプリが揃っていて使い勝手がいい。

ただ， [Simplenote] は良くも悪くもプレーンテキストのエディタで，まとまった情報を整理して書こうとするとイマイチである。
しかし仕事でもないのにワープロなんか使いたくないし，簡単な構造化テキストであれば markdown で書けた方がいい。

で， markdown なテキストが書ける SaaS 型のエディタツールがないかなぁ，と思っていたのだが，どうも [HackMD] が良さげな感じである。

- [HackMD - 共同編集できるMarkdownノート](https://hackmd.io/)
- [hackmdio/hackmd: Realtime collaborative markdown notes on all platforms.](https://github.com/hackmdio/hackmd/)

[HackMD] の特徴を以下に列挙してみる。

- MIT ライセンス。 Docker イメージも用意されていてオンプレミスな運用もできる
- [hackmd.io](https://hackmd.io/ "HackMD - 共同編集できるMarkdownノート") で提供される SaaS 版は以下の通り
    1. 自身はアカウント管理を行ってない。サインインは Facebook, Twitter, GitHubm, Dropbox, Google のアカウントを利用できる。ちなみにサインインしなくても「ゲスト」として利用することも可能
    1. Markdown 記法で記述するが，かなり強力にカスタマイズされている
        - [YAML 形式による制御](https://hackmd.io/yaml-metadata "Supported YAML metadata - HackMD")ができる
        - [MathJax](www.mathjax.org) による数式表現が可能
        - [graphviz]（DOT 言語）, [mermaid]，[js-sequence-diagrams]，[flowchart.js] の記法で作図ができる
        - [abc] 記法で楽譜が書ける
        - 簡単な記述により YouTube, Vimeo, Gist, SlideShare, Speakerdeck のコンテンツを埋め込める
        - タグを設定できる。タグをキーにした検索が可能
        - その他，詳しくは「[機能紹介](https://hackmd.io/s/4JbKDCN1hx "機能紹介 - HackMD")」で
    1. [imgur] と連動している。アップロードした画像は自動的に [imgur] に格納される
    1. Dropbox, Google Drive, Gist へエクスポート可能
    1. Dropbox, Google Drive, Gist およびクリップボードからインポート可能
    1. markdown または HTML 形式でローカルにダウンロード可能
    1. 基本的に誰でも編集でき誰でも閲覧できる。なお，編集・閲覧許可範囲を「サインイン・ユーザのみ」「オーナーのみ」に絞ることはできる。許可するユーザを指定したりはできないようだ（まぁアカウント管理をしてないからね）

試しにちょろんと落書きしてみた。

- [MathJax による数式表現。 - HackMD](https://hackmd.io/s/S1thQI76e)
- [シーケンス図を描こう - HackMD](https://hackmd.io/s/ByuxOLQag)
- [クラス図を描こう - HackMD](https://hackmd.io/s/S19e0LXTe)
- [フローチャートを描こう - HackMD](https://hackmd.io/s/H1iq2i76e#)

もうしばらく遊んでみて，よさげなら常用してみようかな，と。

## ブックマーク

- [リアルタイム共同編集可能なMarkdownエディタ「HackMD」をハックしてみた| Nulab (Japanese)](https://nulab-inc.com/ja/blog/nulab/hackmd-hack/)
- [esa.ioと HackMDでつくるいい感じの議事録&ドキュメント管理 | ShareWis Blog](http://blog.share-wis.com/esa-and-hackmd)
- [クラウド上にMarkdownで手軽にメモを残せる無料ツール「HackMD」が便利](https://nelog.jp/hackmd)

[Simplenote]: https://simplenote.com/
[HackMD]: https://hackmd.io/ "HackMD - 共同編集できるMarkdownノート"
[imgur]: http://imgur.com/ "Imgur: The most awesome images on the Internet"
[graphviz]: http://www.graphviz.org/ "Graphviz | Graphviz - Graph Visualization Software"
[js-sequence-diagrams]: https://bramp.github.io/js-sequence-diagrams/ "js-sequence-diagrams by bramp"
[mermaid]: http://knsv.github.io/mermaid/ "mermaid - Generation of diagrams and flowcharts from text in a similar manner as markdown."
[flowchart.js]: http://flowchart.js.org/
[abc]: http://abcnotation.com/
