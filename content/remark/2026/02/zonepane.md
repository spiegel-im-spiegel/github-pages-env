+++
title = "ぞーぺんのユーザ体験"
date =  "2026-02-18T20:17:33+09:00"
description = "ぞーぺんのスタイルだと，最初に表示されるポストから下は既に見たポストだと確信が持てる。そしてスクロールアップして TL の先頭に辿り着けば閲覧終了である。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "android", "apple", "mastodon", "bluesky", "communication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

「ぞーぺん（ZonePane）」については以前に[クロスポスト機能を紹介]({{< ref "/remark/2025/06/multi-post-with-zonepane.md" >}})した。
あれから iOS 版もリリースされたそうで

- [ZonePane iOS版をリリースしました！｜takke](https://note.com/panecraft/n/n2b5e96633032)
- [ZonePane for iOS が Mastodon に対応しました！｜takke](https://note.com/panecraft/n/n0bfcfffdcb62)

CMP[^cmp1] (Compose Multiplatform) で共通化した部分については Android 版も幾らかこれの恩恵を受けているっぽい。
前にぞーぺんの利用を止めてしまった人も改めて試したら印象が変わってるかもしれない。

[^cmp1]: CMP (Compose Multiplatform) は Android と iOS の両方に対応したアプリを開発できる Kotlin ベースのフレームワーク。

さて，2ヶ月ちょっとの入院生活だったわけだが，困ったのが通信環境。
転院後は Wi-Fi 環境も使えるようになったが，回線が細くて微妙な感じ。
自宅から[ノート PC]({{< ref "/remark/2024/05/get-a-used-pc-from-workplace.md" >}} "勤務先からの払い下げ PC") を持ってきてもらったが，それなりに場所を取るし，特にリハビリ中などは貴重品入れにしまっておく必要があり常時広げておくわけにはいかないため実際にはあまり使わなかった。

結局 Mastodon や Bluesky のタイムライン（TL）の新着チェックは専らスマホのキャリア回線で行い，アプリはぞーぺんを使っていた。
いや，ホンマぞーぺん入れててよかったよ。

それから，おかげさまで無事退院して自宅の広いモニタで Web ブラウザベースで作業できるようになったんだけど TL を眺めててどうにも違和感を感じる。
ちょっと考えて分かった。
スクロールの向きが違うんだ！

Mastodon にせよ Bluesky にせよ Web アプリではまず TL のトップが表示され過去を遡るようにスクロールダウンする。
一方，ぞーぺんは前回閲覧したポストを覚えていて，そこからスクロールアップして新しいポストを読む，という動作になる。

ぞーぺんのスタイルだと，最初に表示されるポストから下は既に見たポストだと確信が持てる。
そしてスクロールアップして TL の先頭に辿り着けば閲覧終了である。

それ以上の作業はない。

一方 Web アプリの場合は過去にスクロールダウンしていくため「どこまで読んだっけ？」と考えながら読み進める。
それで「この辺かな？」というところまで辿り着いてもまだ下にスクロールできるため，ついその先（過去）も見てしまう。

それでもちゃんと TL が時系列に並んでいるなら「止めどき」がまだ分かりやすいけど，これが「おすすめ」順になっていたら目も当てられない。
「やめられない，とまらない」というやつだ。
まさにメンタル・ジャンクフード，注目の搾取，（可処分）時間の搾取，である。
だから「[子どもの脳内中毒を設計](https://www.cnn.co.jp/tech/35243762.html "SNSが「子どもの脳内中毒を設計」、ユーザーがメタとYouTube訴えた裁判始まる(1/2) - CNN.co.jp")」とか言われて訴えられるんだよ（「脳内中毒」という言い草には？？？だけど）。

というわけで，今後も TL の新着チェックはスマホでぞーぺんで行うことにした。
過去に「おすすめ」や「ブックマーク」したポストの確認とそこから発生する調べものはパソコンのブラウザで。
何故ならスマホのブラウザはタブ操作の使い勝手が悪いから。

マジで UX (User eXperience) って大事だなって思ったよ。

## ブックマーク

- [ZonePane(ぞーぺん) (@zonepane@fedibird.com) - Fedibird](https://fedibird.com/@zonepane)

- {{< pdf-file link="https://gwern.net/doc/culture/2010-dobelli.pdf" title="Avoid News: Towards a Healthy News Diet" >}}

## 参考

{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->

[ZonePane]: https://play.google.com/store/apps/details?id=com.zonepane "ZonePane for Bluesky&Mastodon - Google Play"
[@takke]: https://fedibird.com/@takke "たけうちひろあき (@takke@fedibird.com) - Fedibird"
[mixi2]: https://play.google.com/store/apps/details?id=social.mixi "mixi2 - Google Play"
