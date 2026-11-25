+++
title = "#ホロじゃん をやるために Windows に Android ゲームを入れた話"
date =  "2026-11-25T17:32:31+09:00"
description = "久しぶりに遊んでみたい麻雀ゲームに出会ったよ。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "games", "android", "windows", "tools", "vtuber" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

ホロライブのコンテンツ（VTuber）を使った麻雀ゲーム「[ホロじゃん]」がリリースされた。

ところがこれ Android 端末か iOS 端末にしか対応していない。
切実に Steam 版が欲しいところであるが，こればっかりは仕方ない。
でもスマホやタブレットでやる気はないので，どうにかパソコン上で動かせないものか。

## スマホ画面をミラーリングする

最初に考えたのはスマホ画面をパソコン上にミラーリングする方法。
可能なら操作もパソコン上でやりたい。

この辺の知識については[ミリしら](https://dic.pixiv.net/a/%E3%83%9F%E3%83%AA%E3%81%97%E3%82%89 "ミリしら (みりしら)とは【ピクシブ百科事典】")状態なので [Kagi Assistant] に訊いてみたら [scrcpy] か， Windows 環境であればスマートフォン連携機能を勧められた。

[scrcpy] は Ubuntu でも使えるので魅力的だが，端末側を USB デバッグモードにする必要があるため少し躊躇する。
幸い Windows が動く[ミニ PC]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした") があるので，こいつでスマートフォン連携のほうを試してみることにした。

- [スマートフォン連携の要件と設定 | Microsoft Support](https://support.microsoft.com/ja-jp/windows/apps/phonelink/phone-link-requirements-and-setup)

設定自体は特に苦もなく完了。
スマホに「[ホロじゃん]」をインストールして起動した状態でミラーリングしてみる。

{{< fig-img-quote src="./hololive-mahjong-1.png" title="ホロじゃん (1) ミラーリング画面" link="./hololive-mahjong-1.png" width="1220" >}}

とりあえず問題なさそう。
操作もマウスで問題なくできる。
ただ，スマホが結構熱くなるのな。
どんだけ負荷かかってるのか。

ミラーリングはスジが悪いようだ。

## Google Play Games を導入する

それから公式 Discord チャネル[^d1] を眺めてたのだが，どうやら Windows では Google Play Games というアプリを導入すると Android ゲームをパソコン上で動かせるらしい。
「[ホロじゃん]」もこれで動くそうだ。

[^d1]: 「[ホロじゃん]」の公式 Discord チャネルは，[公式 Web ページ][ホロじゃん] の {{% icons "discord" %}} アイコンのリンクから招待を受けられる。 Discord 側で {{% emoji "X" %}} 公式アカウントのポストも閲覧できるのが素晴らしい。

- [PC 版 Google Play Games のインストールと設定 - Google Play ヘルプ](https://support.google.com/googleplay/answer/11358888?hl=ja-JP)

さっそく Google Play Games をインストールして，「[ホロじゃん]」を動かしてみた。

{{< fig-img-quote src="./hololive-mahjong-2.png" title="ホロじゃん (2) Google Play Games での画面" link="./hololive-mahjong-2.png" width="964" >}}

おぉ！ できたできた。
操作も問題なし。
{{% keytop %}}F11{{% /keytop %}} キーで全画面表示にもできる。
今の Windows ってこういうことができるんだねぇ。
便利になったもんだ。

私は一応でん同士なので，彼女を選択。

{{< fig-img-quote src="./hololive-mahjong-3.png" title="ホロじゃん (3) 儒烏風亭らでんを選択" link="./hololive-mahjong-3.png" width="1300" >}}

ただ Google Play Games ではアイテム購入等が上手くできない感じ？ 私はほぼ無課金ユーザなので困らないが。
まぁ，いざとなったらスマホ側で課金すればいいか（アカウント連携や引き継ぎをすれば，複数の端末からアクセスできる）。

さて，心置きなく遊びますかね。
久しぶりに遊んでみたい麻雀ゲームに出会ったよ。

## ブックマーク

- [スマートフォン連携でデバイスを削除する | Microsoft Support](https://support.microsoft.com/ja-jp/windows/apps/phonelink/remove-a-device-in-the-phone-link)

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[ホロじゃん]: https://www.holo-jong.com/ "ホロライブ麻雀-ホロじゃん-｜今度の推し活は、麻雀だ！【公式】"
[scrcpy]: https://github.com/Genymobile/scrcpy "Genymobile/scrcpy: Display and control your Android device · GitHub"
