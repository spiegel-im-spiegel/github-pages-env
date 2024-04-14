+++
title = "Chromebook を導入する 1"
date =  "2024-03-28T22:59:51+09:00"
description = "「頑張ったご褒美」第2弾"
image = "/images/attention/tools.png"
tags = [ "chromebook", "tools", "gear", "nas", "tailscale", "cloud", "samba" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

給与明細を見たら残業代が凄いことになってたのですよ。
ちうわけで「頑張ったご褒美」第2弾。
ついに Chromebook を導入することにした。

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}}) ← イマココ
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}})
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}})

## 「頑張ったご褒美」にノートPCを物色中

私は（当時の）東芝の Dynabook が登場したばかりの頃にそれを持たされて，現地調整と称して客先に缶詰になったことが何度かあってノート PC にいいイメージがない。
（今で言うところの）ブラックの象徴だったり。
まぁ，でも，仕事以外のサイクリング等で出掛けた先で軽く調べ物とかできたらいいなぁ，とちょっと思い直した。
そんでノートPCの物色を始めたんだけど...

私がプライベートで使ってる[メッセンジャーバッグ]({{< ref "http://localhost:1313/remark/2022/10/messenger-bag/" >}} "遅走り，早歩き")はサイズ小さめで13インチ以下のノートPCしか入らないのね。
なのでその線で探してみたんだけど，
15インチノートPCのほうが断然安いってどういうことやねん！ 進学・就職・異動の季節で各メーカーが安売りしてるみたいだけど，13インチノートPCは恩恵を受けられないようだ。
中古PCは前に酷い目にあったので買いたくないし。

...などとしばらく悩んで「そうだ！ Chromebook なら要件を満たすんじゃね？」と気がついた。
しかも想定よりだいぶ安上がりになるし（笑）

## Chromebook / ChromeOS とは

今更だけど，念のために記しておく。

Chromebook は ChromeOS を搭載したパソコンである。
日本円で数万円程度の廉価 PC として提供されていることが多い。

ChromeOS は，もの凄く簡単に言うと， Chrome ブラウザを GUI とした OS である。
したがって，基本的に ChromeOS のアプリケーションは Chrome ブラウザの拡張機能として実装されている（[Chrome ウェブストア](https://chromewebstore.google.com/)から導入可能）。
ただし，近年の ChromeOS は Android と Linux のサブシステムを標準で装備している。

つまり Chromebook では Chrome ベースのアプリと Andorid アプリと Linux のアプリケーションが動くわけだ。
これならネットワーク端末としては申し分ないだろう。

## Chromebook を購入した

前述したように，今回は13インチ以下のノートPCという縛りがある。
あちこち探し回って疲れたので Amazon にあるものから選択することにした。
日和ったとも言う（笑）

要件を満たす候補として [ASUS 製](https://www.amazon.co.jp/gp/product/B0BKKF7JHV?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)と [Lenovo 製](https://www.amazon.co.jp/dp/B0CNSWLJFM?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)の2つまで絞り込んだ。
[Lenovo 製](https://www.amazon.co.jp/dp/B0CNSWLJFM?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)のほうが安くて軽くてストレージ容量が大きいのだが，私が ASUS のほうが好みなのとメモリ容量が大きいということで [ASUS 製](https://www.amazon.co.jp/gp/product/B0BKKF7JHV?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)に決めた。
ストレージは後から調達できるからね。

というわけで到着。

{{< fig-img src="./53614190956_f3a103a7a4_e.jpg" title="先月の汗と涙とストレスと（削られた）睡眠時間の対価 | Flickr" link="https://www.flickr.com/photos/spiegel/53614190956/" >}}

早速開梱してセットアップを始める。
以降は作業中の覚え書である。

## ChromeVox ？

セットアップが始まって最初に ChromeVox を有効にするか訊かれる。
ChromeVox って何？

- [組み込みのスクリーン リーダーを使用する - Chromebook ヘルプ](https://support.google.com/chromebook/answer/7031755?hl=ja)

んー。
どうやら視覚障害者のためのスクリーンリーダーらしい。
いや，それならそうとちゃんと説明してくれ。
分かんないからググったやんけ。

ありがたいことに私は[眼鏡があれば普通に見える]({{< ref "/remark/2021/08/age-of-reading-glasses.md" >}} "老眼鏡の季節")ので，今回はスルーした。

## PIN コードを設定する

そんなこんなでセットアップも無事完了して，色々と弄り始めたのだが，スクリーンロックからの復帰とか，ことあるごとに Google アカウントのパスワードを入れさせようとするのね。
いや，パスワードは覚えない主義なので，それじゃ困るわけさ。
せめて PIN コードの設定とかできないの？ と探してみたらあった。

設定から「セキュリティとプライバシー」を選択し，さらに「ロック画面とログイン」を開く。

{{< fig-img src="./setting-pin-code.png" title="ロック画面とログイン" link="./setting-pin-code.png" width="1100" >}}

設定後の状態ですまん。

あとで調べたら，最初のセットアップのときに PIN コードも設定させるみたいなのだが，そんな画面あったかな？ まぁ，いいや。
無事に PIN コードが設定できたのでよしとする。

## Dropbox を Files アプリにマウントする

Chromebook に標準で入ってる「ファイル（Files）」アプリはかなり出来がよくて，最初から Google Drive がマウントされた状態になっている。
ただ，私の場合はクラウドストレージは Dropbox を常用してるのね。
調べてみたら Dropbox も Google Drive のように Files にマウントできるようだ[^cs1]。

[^cs1]: あとで調べたら Box や OneDrive も Android アプリを使って Files にマウントできるらしい。今回は試さないが。

- [Chromebook の Files アプリに Dropbox を追加する - Dropbox ヘルプ](https://help.dropbox.com/ja-jp/integrations/google-files-app)

といっても全然難しくなく Android アプリの Dropbox アプリを導入してサインインすればいいらしい。

- [Dropbox: Secure Cloud Storage - Apps on Google Play](https://play.google.com/store/apps/details?id=com.dropbox.android&hl=en_US)

これで自動的に Files に Dropbox が表示される。

{{< fig-img src="./dropbox-in-files.png" title="Dropbox をマウントする" link="./dropbox-in-files.png" width="1100" >}}

もし，いつまで待っても Dropbox が表示されないなら，一度再起動すると上手く行くことがある。

## NAS 共有ディレクトリを Files アプリにマウントする

更に Files では SMB プロトコルで NAS の共有ディレクトリもマウントできるようだ。
設定から 詳細設定 → ファイル → ネットワークファイル共有 → ファイル共有を追加 と辿ると設定画面が表示される。

{{< fig-img src="./shared-directory.png" title="ファイル共有を追加" link="./shared-directory.png" width="1100" >}}

「ファイル共有 URL」に共有ディレクトリを `\\server\shared_directory` または `smb://server/shared_directory` 形式で指定し，ユーザ名とパスワードをセットする。

ちなみに今回買った ASUS の Chromebook ではキーボード右上の {{% keytop %}}&yen;{{% /keytop %}} キーと右下の {{% keytop %}}`\`{{% /keytop %}} キーに割り当てられている文字コードが正しく &yen; (`U+00A5`) と `\` (`U+005C`) になっていた。
ちょっとしたことなんだけど時代を感じるねぇ。

## Tailscale も設定できた

Andorid アプリの Tailscale アプリを導入することで tailscale にも対応できる。

- [Setting up Tailscale on a Chromebook · Tailscale Docs](https://tailscale.com/kb/1267/install-chromebook)
- [Tailscale - Apps on Google Play](https://play.google.com/store/apps/details?id=com.tailscale.ipn&hl=en_US)

当然ながらセットアップは [Android 版]({{< ref "/remark/2021/10/tailscale-with-synology-nas.md#android" >}} "Android 端末に Tailscale をインストールする")と同じ。
これで外から自宅 NAS に繋がるようになった。

よーし，うむうむ，よーし。

## その他のインストール

随時更新予定。

- [Keepass2Android Password Safe - Apps on Google Play](https://play.google.com/store/apps/details?id=keepass2android.keepass2android&hl=en_US)

これで，最低限の環境は揃ったかな。

{{< fig-img src="./53613544342_db09981a76_e.jpg" title="とりあえず最低限のセットアップ完了 | Flickr" link="https://www.flickr.com/photos/spiegel/53613544342/" >}}

自宅 Linux/Ubuntu 機と壁紙を合わせてみた。

次回からいよいよ Linux サブシステムのセットアップを行う。


## ブックマーク

- [画面のスクリーンショットを撮影または録画する - Chromebook ヘルプ](https://support.google.com/chromebook/answer/10474268?hl=ja)
- [Chromebookを購入しました｜Mina](https://note.com/minakonono/n/n30b173adc169)
- [【Chromebook】Dropboxを利用する方法（ファイルアプリからDropboxフォルダを利用できるようにする／Dropboxフォルダが表示されない場合の対処）](https://did2memo.net/2020/10/19/chromebook-dropbox-app-setup/)
- [ChromebookとノートPC、どっちが買い？ 使い方で決まる最適解 | ライフハッカー・ジャパン](https://www.lifehacker.jp/article/2404chromebook-vs-laptop-how-to-choose/)

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B0CH2X5LBX" %}} <!-- micoroSD カード -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B00G9NIL7G" %}} <!-- エレコム マウス Bluetooth -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "B08P54PQDB" %}} <!-- メッセンジャーバッグ -->
