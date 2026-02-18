+++
title = "ぞーぺん（ZonePane）によるクロスポスト"
date =  "2025-06-01T21:16:42+09:00"
description = "複数のサービスを利用できる Android アプリ「ぞーぺん」がサービスを跨いだクロスポストに対応した。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "android", "mastodon", "bluesky" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Bluesky / Mastodon / Misskey の複数のサービスを利用できる Android アプリ「ぞーぺん（[ZonePane]）」がサービスを跨いだクロスポストに対応した。

- [ZonePane for Bluesky&Mastodon - Apps on Google Play](https://play.google.com/store/apps/details?id=com.zonepane)

実は「[ぞーぺん][ZonePane]」は昨年の秋に導入してたことがあるんだけど，それまで使ってたサードパーティの Mastodon アプリ[^ma1] に比べてイマイチな気がして削除してしまったのよ。
でも「[ぞーぺん][ZonePane]」自体は気になっていて[作者のアカウント][@takke]をフォローして様子見していたのだった。

[^ma1]: Mastodon アプリは一応[公式](https://joinmastodon.org/ja/apps "Mastodonアプリを入手 - Mastodon")のものがあるんだけど，イマイチなんだよな。で，サードパーティのアプリを使ってる人も多いらしい。 Android 用としては [Subway Tooter](https://play.google.com/store/apps/details?id=jp.juggler.subwaytooter "Subway Tooter - Google Play") とか人気だよね。私は [Fedilab](https://play.google.com/store/apps/details?id=app.fedilab.android "Fedilab - Google Play") を使っていた。

ツールバーの {{% keytop %}}<img src="./button-1.png" alt="投稿" width="25" />{{% /keytop %}} ボタンを長押しするとクロスポストを含む投稿先の選択ダイアログが出る。

{{< fig-img src="./zonepane-01.png" title="投稿" link="./zonepane-01.png" >}}

ここで「クロスポスト」を選択すると

{{< fig-img src="./zonepane-03.png" title="アカウント選択" link="./zonepane-03.png" >}}

という画面になるのでアカウントを選択する。
なお，以前にクロスポストを行っていれば，選択したアカウントの履歴が残る。

{{< fig-img src="./zonepane-04.png" title="アカウント選択 履歴" link="./zonepane-04.png" >}}

履歴からアカウントを選択すると簡単である。

投稿の作成・編集は普通にできる。

{{< fig-img src="./zonepane-05.png" title="投稿の作成" link="./zonepane-05.png" >}}

作成できたら右上の {{% keytop %}}＞{{% /keytop %}} をタップして次の画面へ。

{{< fig-img src="./zonepane-06.png" title="公開範囲の設定" link="./zonepane-06.png" >}}

ここで投稿先ごとに公開範囲を設定できる。
一見まだるっこしい操作に見えるけど，公開範囲の指定方法はサービスによって異なるので，確認も兼ねて妥当な操作だと思う。

{{< fig-quote class="nobox" >}}
<iframe src="https://fedibird.com/@takke/114605044959694244/embed" width="400" allowfullscreen="allowfullscreen" sandbox="allow-scripts allow-same-origin allow-popups allow-popups-to-escape-sandbox allow-forms"></iframe>
{{< /fig-quote >}}

公開範囲を確認・設定できたら {{% keytop %}}＞{{% /keytop %}} で次の画面へ。

{{< fig-img src="./zonepane-07.png" title="投稿準備完了" link="./zonepane-07.png" >}}

確認して問題なければ {{% keytop %}}投稿開始{{% /keytop %}} をタップして処理を開始する。

{{< fig-img src="./zonepane-08.png" title="投稿処理中" link="./zonepane-08.png" >}}

{{< fig-img src="./zonepane-09.png" title="投稿完了" link="./zonepane-09.png" >}}

さてここからが「[ぞーぺん][ZonePane]」の面白いところ。
端末に {{% emoji "X" %}}, Threads, mixi2 のアプリがインストールされていれば Android の共有機能を使って，これらのサービスにもクロスポストできる。

以下は [mixi2] アプリで投稿しているところ。

{{< fig-img src="./zonepane-11.png" title="mixi2 アプリで投稿" link="./zonepane-11.png" >}}

mixi2 は API が公開されてないし， {{% emoji "X" %}} は API 利用料が馬鹿高い上そもそもサードパーティアプリを許容してないので「[ぞーぺん][ZonePane]」のアプローチはとてもいいと思う。
まぁ，私は携帯端末に {{% emoji "X" %}} のアプリは入れてないけどね。
Threads は知らん（笑）

「[ぞーぺん][ZonePane]」は基本無料だがアプリ内課金でサブスクリプション契約をすれば広告が出なくなる他，上述の Android の共有機能を使ったクロスポストの制限がなくなる。

{{< fig-img src="./zonepane-10.png" title="クロスポストの制限" link="./zonepane-10.png" >}}

あと「[ぞーぺん][ZonePane]」の偉いところは複数のアカウントのタイムラインをタブで並べることができることだ。
これで明示的にアカウント切り替えしなくても左右のスワイプでタブを切り替えるだけで複数アカウントの TL を見ることができる。
めがっさ便利！

少なくとも Mastodon アプリは「[ぞーぺん][ZonePane]」に切り替えてもいいかな。
もう少し使ってみて問題なさそうならサブスクリプション契約しよう。

## ブックマーク

- [ZonePane(ぞーぺん) (@zonepane@fedibird.com) - Fedibird](https://fedibird.com/@zonepane)
- [ぞーぺん、ついにクロスポスト対応！｜takke](https://note.com/panecraft/n/n55b33ad1496e)
- [ZonePaneのクロスポスト機能のこだわり｜takke](https://note.com/panecraft/n/n504805dcb70c)

[ZonePane]: https://play.google.com/store/apps/details?id=com.zonepane "ZonePane for Bluesky&Mastodon - Google Play"
[@takke]: https://fedibird.com/@takke "たけうちひろあき (@takke@fedibird.com) - Fedibird"
[mixi2]: https://play.google.com/store/apps/details?id=social.mixi "mixi2 - Google Play"
