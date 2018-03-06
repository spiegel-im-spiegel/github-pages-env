+++
draft = false
date = "2017-04-06T22:25:55+09:00"
update = "2018-03-06T11:00:35+09:00"
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
    1. 自身はアカウント管理を行ってない。サインインは Facebook, Twitter, GitHub, Dropbox, Google のアカウントを利用できる。ちなみにサインインしなくても「ゲスト」として利用することも可能
    1. Markdown 記法で記述するが，かなり強力にカスタマイズされている
        - [YAML 形式によるページ制御](https://hackmd.io/yaml-metadata "Supported YAML metadata - HackMD")ができる
        - [MathJax](www.mathjax.org) による数式表現が可能
        - [graphviz]（DOT 言語）, [mermaid]，[js-sequence-diagrams]，[flowchart.js] の記法で作図ができる
        - [abc] 記法で楽譜が書ける
        - 簡単な記述により YouTube, Vimeo, Gist, SlideShare, Speakerdeck のコンテンツを埋め込める
        - タグを設定できる。タグをキーにした検索が可能
        - その他，詳しくは「[機能紹介](https://hackmd.io/s/4JbKDCN1hx "機能紹介 - HackMD")」で[^ed1]
    1. [imgur] と連動している。アップロードした画像は自動的に [imgur] に格納される[^ig]
    1. Dropbox, Google Drive, Gist へエクスポート可能
    1. Dropbox, Google Drive, Gist およびクリップボードからインポート可能
    1. markdown または HTML 形式でローカルにダウンロード可能
    1. 基本的に誰でも編集でき誰でも閲覧できる。なお，編集・閲覧許可範囲を「サインイン・ユーザのみ」「オーナーのみ」に絞ることはできる。許可するユーザを指定したりはできないようだ（まぁアカウント管理をしてないからね）[^pv]

[^ed1]: 編集画面で見出し単位で表示の畳み込みができるのが地味に便利。アウトライン編集に使える。
[^ig]: [imgur] への画像の登録は匿名アカウントで行われ後から削除できない。したがって間違ってアップロードしても取り消せないし，ましてや公開できない画像をアップロードするのは以っての外である。ご注意を。
[^pv]: 言うまでもないが，パスワード等の秘密情報，プライバシーに関わる情報（個人情報を含む），その他公開できない情報をこのサービスに載せないこと。一応 private モードにすればオーナー以外は編集・閲覧できないが，この手のサービスは信用しすぎないのが肝要である。

試しにちょろんと落書きしてみた。
（楽譜は無理。復活の呪文を唱えているようにしか見えん）

- [MathJax による数式表現。 - HackMD](https://hackmd.io/s/S1thQI76e)
- [シーケンス図を描こう - HackMD](https://hackmd.io/s/ByuxOLQag)
- [クラス図を描こう - HackMD](https://hackmd.io/s/S19e0LXTe)
- [フローチャートを描こう - HackMD](https://hackmd.io/s/H1iq2i76e#)
- [Go 言語で Hello World - HackMD](https://hackmd.io/s/Hkrec_Nae)

もうしばらく遊んでみて，よさげなら常用してみようかな，と。

## 【追記】 メタデータについて

上で述べたように [HackMD] では YAML 形式によるページ制御ができる。

- [Supported YAML metadata - HackMD](https://hackmd.io/yaml-metadata)

具体的には先頭行に以下の記述を加える。

```yaml
---
YAML metas
---
```

設定項目は色々あるが，私は必ず以下の設定をするようにしている。

```yaml
---
robots: noindex, nofollow
lang: ja
dir: ltr
breaks: false
---
```

`robots` 項目は `<meta>` 要素に robots を設定する。
検索エンジンや他ページの referer に拾われたくない場合は `noindex, nofollow` をセットしておけばいい。
ただし行儀のいい crawler や Web サイトばかりではないので，その辺はあしからず。

`lang` 項目はページに国・言語情報を設定する，筈なのだが利いてないようである（既定の `en` のまま）。
まぁそのうち有効になると信じて。

`dir` 項目は文字の向き（右向き・左向き）を指定する。
既定は `ltr` （左→右向き）なのでなくてもいいのだが，一応設定しておく。

`breaks` 項目は markdown 入力の改行をそのまま HTML 表示に反映（hard break）させるかどうか指定する。
既定は `true`。
この辺は好みで。

## ブックマーク

- [リアルタイム共同編集可能なMarkdownエディタ「HackMD」をハックしてみた| Nulab (Japanese)](https://nulab-inc.com/ja/blog/nulab/hackmd-hack/)
- [esa.ioと HackMDでつくるいい感じの議事録&ドキュメント管理 | ShareWis Blog](http://blog.share-wis.com/esa-and-hackmd)
- [クラウド上にMarkdownで手軽にメモを残せる無料ツール「HackMD」が便利](https://nelog.jp/hackmd)

[Simplenote]: https://simplenote.com/
[HackMD]: https://hackmd.io/ "HackMD - 共同編集できるMarkdownノート"
[imgur]: http://imgur.com/ "Imgur: The most awesome images on the Internet"
[Graphviz]: http://graphviz.org/ "Graphviz - Graph Visualization Software"
[js-sequence-diagrams]: https://bramp.github.io/js-sequence-diagrams/ "js-sequence-diagrams by bramp"
[mermaid]: http://knsv.github.io/mermaid/ "mermaid - Generation of diagrams and flowcharts from text in a similar manner as markdown."
[flowchart.js]: http://flowchart.js.org/
[abc]: http://abcnotation.com/
