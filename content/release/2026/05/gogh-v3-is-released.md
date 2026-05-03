+++
title = "作業支援ゲーム Gogh v3 がリリースされた"
date =  "2026-05-03T15:03:51+09:00"
description = "Watch Party 機能およびフェイストラッキング機能の紹介。最後に VTuber ごっこもやってみた（笑）"
isCJKLanguage = true
image = "/release/2026/05/gogh-v3-is-released/obs-studio.png"
tags  = [ "games", "streaming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

日頃お世話になっている作業支援ゲーム [Steam 版 gogh][gogh] が発売1周年（おめでとうございます）を機に 3.0.0 をリリースした。

- [gogh、Steam版の発売1周年で『ルーム投稿』『YouTube同時視聴』『コーデ保存』『フェイストラッキング』など大型アップデート！ | 株式会社 ambrのプレスリリース](https://prtimes.jp/main/html/rd/p/000000072.000043299.html)
- [gogh: Focus with Your Avatar－Major Update 3.0.0: ルーム投稿、YouTube Watch Party など新機能多数！！－Steamニュース](https://store.steampowered.com/news/app/3213850/view/537759153440948744)

主なアップデート内容は以下の通り：

1. ルーム投稿
2. YouTubeの同時視聴 "Watch Party"
3. コーデ保存
4. 髪型左右反転
5. フェイストラッキング機能
6. 表情変更機能
7. 回転くん
8. 新アイテム、新ポーズ
9. 新種のクロダ42「せぱたクロ」が登場
10. Steamフレンドをマルチルームに招待できるように！
11. ルームスロットの拡大

この中で今回は Watch Party 機能とフェイストラッキング機能を紹介する。

## Watch Party 機能

アイテムに「Watch Party スクリーン」が追加され，マルチルーム・モードで YouTube 視聴が出来るようになった。
サイズが S/M/L/XL とあり XL だとほぼ壁一面の大きさになる。

せっかくなので新しく YouTube 視聴用の部屋を作ってみた。

{{< fig-img-quote src="./watch-party-0.png" title="gogh: Focus with Your Avatar" link="./watch-party-0.png" width="1282" lang="en" >}}

部屋つうか野外だけど。
スクリーンは XL サイズ。
[違法建築（合法）](https://note.com/koku_sb/n/n429fef46eb04 "続・goghで違法建築のすすめ(※合法)｜こく")で床面を拡張している。
屋根もないようなところで YouTube 視聴とかファンタジーでいいよね。

Watch Party スクリーン設定画面（マルチルーム・モード）はこんな感じ。

{{< fig-img-quote src="./watch-party-1.png" title="Setting Watch Party Screen" link="./watch-party-1.png" width="1282" lang="en" >}}

ちなみに映ってる動画は NASA の “[NASA's Artemis II Crew Launches To The Moon](https://www.youtube.com/watch?v=Tf_UjBMIzNo)” である。
この状態で「共有開始」ボタンを押すと他プレイヤーも同じものを見ることが出来る。

{{< fig-img-quote src="./watch-party-2.png" title="Watch Party Screen" link="./watch-party-2.png" width="1282" lang="en" >}}

かなり無理やり映してる感じで，色々と制限がある。

- Ubuntu/Proton 環境下では YouTube 動画が動かなかった。これは想定内
- サインオフ状態での視聴なので，とにかく広告がウザい。つか，無料視聴だとこんなに広告が入るんだねぇ
- 動画の seek ができない（らしい）。基本的に流しっぱなし

制作側がアナウンスしている制限というか注意事項は以下の通り：

{{< fig-quote type="markdown" title="gogh: Focus with Your Avatar－Major Update 3.0.0: ルーム投稿、YouTube Watch Party など新機能多数！！" link="https://store.steampowered.com/news/app/3213850/view/537759153440948744" >}}
⚠️ 利用上のご注意

- 利用にはYouTube利用規約およびGoogleプライバシーポリシーが適用されます。
- WatchPartyは、各参加者の端末でYouTube動画を再生し、goghが再生位置を再生開始からの経過時間で合わせる独自機能です。YouTubeが提供・保証する機能ではありません。
- WatchPartyはマルチルームのホストのみが利用できます。他ユーザーと同期しないソロ視聴は全員が可能です。
- WatchParty中にホストが再生位置をシークした場合、同期はされません。
- 技術上の制約により、再生停止・音量操作などはgoghのUIから制御できません。YouTube画面を操作してください。
- 技術上の制約により、一部の動画は再生できません。ご了承ください。
- YouTube以外のURLは入力できません。
- 複数箇所に投影する場合も、同時に流せる動画は1つです。
{{< /fig-quote >}}

ぶっちゃけ，機能上の制限より YouTube のポリシーによる制限のほうがキツそう。
YouTuber (VTuber 含む) の配信をマルチルームで流すのはまず無理だよね（自分で作成したコンテンツなら問題ないだろうけど）。
政府系の記者会見とか日食ライブとか公共性が高そうなものならなんとか行ける？ と思って NASA live の ”[Live Video from the International Space Station](https://www.youtube.com/watch?v=sWasdbDVNvc)” を流そうとしたら，何かの制限に引っかかるらしく，できなかった。
残念。

独りで見るのなら，既にある Stream 機能と OBS の仮想カメラを組み合わせたやり方のほうが融通がきく。

まぁ，せっかく部屋を作ったのでミラーリング可能な配信コンテンツがあればまた試してみよう。

## フェイストラッキング機能

Web カメラ映像を取り込んで顔の動きを同期できる。

{{< fig-img-quote src="./face-tracking-1.png" title="face tracking" link="./face-tracking-1.png" width="1272" lang="en" >}}

キャリブレーションはこんな感じ。

{{< fig-img-quote src="./face-tracking-2.png" title="face tracking" link="./face-tracking-2.png" width="1272" lang="en" >}}

なんで身体が斜めってるのかは不明だが，別にカメラとの同期がおかしいわけではないらしい。

あれ？ これってグリーンバックとかの配信部屋を作って OBS と組み合わせれば [gogh] のアバターで VTuber ごっこができる？ やってみよう。

{{< fig-img-quote src="./obs-studio.png" title="OBS Studio" link="./obs-studio.png" width="1282" lang="en" >}}

おー。できるぢゃん（笑） これはまたオモチャの要素が増えたねぇ。

## ブックマーク

- [マルチプレイ作業集中ゲーム『gogh』「YouTubeのウォッチパーティ」機能が追加](https://news.denfaminicogamer.jp/news/260430w)
- [マルチ対応作業集中アプリ『gogh: Focus with Your Avatar』Steam版、大規模アプデで“YouTube同時視聴”ができるように。休憩もみんなで一緒に - AUTOMATON](https://automaton-media.com/articles/newsjp/gogh-20260430-440674/)

[gogh]: https://store.steampowered.com/app/3213850/gogh_Focus_with_Your_Avatar/ "gogh: Focus with Your Avatar"
