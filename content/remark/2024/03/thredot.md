+++
title = "オンラインメモ帳 Thredot がよさげ"
date =  "2024-03-29T22:46:47+09:00"
description = "メモ書きに「分類」は不要というのは同意する。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "chromebook", "web" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Koki Sato](https://koki.me/) さんによる [Thredot] が面白い。

- [スレッド形式で雑にメモを書き散らすためのサービス「Thredot」をリリースしました](https://zenn.dev/kou_pg_0131/articles/thredot-introduction)

スレッド形式のメモサービスといえば [Zenn] のスクラップ機能が思い浮かぶが，上の記事によると [Thredot] の特徴は

{{< fig-quote type="markdown" title="スレッド形式で雑にメモを書き散らすためのサービス「Thredot」をリリースしました" link="https://zenn.dev/kou_pg_0131/articles/thredot-introduction" >}}
- メモを整理させない
- 爆速な検索機能
- 公開範囲を設定できる
- WYSIWYG エディタを採用
{{< /fig-quote >}}

ということらしい。
メモ書きに「分類」は不要というのは同意する。
「メモ」では思ったことや調べたことを深く考えずに大量に書き留めていきたい。
その代わりにスレッド名の検索が「爆速」なので，スレッド名にキーワードになる単語を列挙しておけば，後からメモを簡単に探せるわけだ。
あと，プログラムコードを syntax highlight させて書き込めるのも魅力である。

というわけで，早速アカウントを作った。

{{< fig-img src="./my-thredot.png" title="Thredot にアカウントを作った" link="./my-thredot.png" width="1100" >}}

上の絵は先日導入した Chromebook の Chrome ブラウザで表示させたもの。
更に [Thredot] サイトを PWA (Progressive Web Apps) 化してみる。

{{< fig-img src="./pwa.png" title="Thredot サイトを PWA 化する (1)" link="./pwa.png" width="1366" >}}

このように 保存して共有 → ショートカットを作成 を選択すると以下のダイアログが表示される。

{{< fig-img src="./pwa-2.png" title="Thredot サイトを PWA 化する (2)" link="./pwa-2.png" width="1366" >}}

ここで「ウィンドウとして開く」をチェックして「作成」すれば PWA 化できる。

{{< fig-img src="./pwa-3.png" title="Thredot サイトを PWA 化する (3)" link="./pwa-3.png" width="1017" >}}

よーし，うむうむ，よーし。

既に Chromebook の作業メモは [Thredot] で行っている（非公開ご容赦）。
やっぱスレッド形式でメモをバンバンぶち込めるのは気持ちいい。
今までメモサービスとしては [Simplenote](https://simplenote.com/) を常用していたが，作業メモや読書メモは [Thredot] のほうがよさそうだ。

初期 [Zenn] のときも思ったが，個人開発でこういうのをサラッと作れる人はホンマに尊敬する。
願わくば

1. チェックボックスのリストが書けると嬉しい（お買い物リストを作れる）
2. 非公開 → 公開 の切り替えができると嬉しい（非公開で書いたけど，あとで公開してもよかったなぁ，と思うことがある）

といったところかな。

## ブックマーク

- [Thredot がサポートしている書式の一覧 | Thredot](https://thredot.org/threads/FRMQVRRNRU499B12)
- [Playground | Thredot](https://thredot.org/playground)

[Thredot]: https://thredot.org/ "Thredot | シンプルに繋がる、あなたのメモ。"
[Zenn]: https://zenn.dev/ "Zenn｜エンジニアのための情報共有コミュニティ"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
