+++
title = "Chromebook を導入する 4 — Flatpak で Firefox を導入する"
date =  "2024-04-11T21:44:36+09:00"
description = "Chromebook にまともな Firefox を入れたい"
image = "/images/attention/tools.png"
tags = [ "chromebook", "linux", "firefox", "flatpak" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]「次回は git かな」と言ったな。
あれはウソだ（笑）

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}})
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}}) ← イマココ
- [Chromebook を導入する 5 — APT で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-5.md" >}})
- [Chromebook を導入する 6 — Git & Go コンパイラ]({{< ref "/remark/2024/05/chromebook-6.md" >}})
- [Chromebook を導入する 7 — VS Code の導入]({{< ref "/remark/2024/05/chromebook-7.md" >}})

## Chromebook に Firefox を入れたい

個人的な嗜好で申し訳ないが，できれば私は Chrome ブラウザは使いたくない。
というわけで Firefox を入れるのだが， Chromebook に Firefox を入れるには，以下の2通りの方法がある。

1. [Andorid 版 Firefox] を導入する
2. Linux サブシステムの Firefox を導入する

最初は安直に [Andorid 版 Firefox] を入れていたのだが，これが使い勝手が悪いのなんの。
そもそもスマホ&タブレット用のアプリなのでフル機能ではなかったりする。

というわけで，まともな Firefox を入れる方法はないの？ と探したら2番目の方法を解説した

- [ChromeOS で Firefox を起動する | Firefox ヘルプ](https://support.mozilla.org/ja/kb/run-firefox-chromeos)

という記事を見つけた。
これは Linux サブシステム上に Firefox を入れるというもので，パッケージマネージャとして [Flatpak] を使うらしい。
Chromebook の Linux サブシステムには既定で [Flatpak] は入ってないので，まずは APT を使って [Flatpak] を導入する。

## Flatpak を導入する

- [Chrome OS Quick Setup; Follow these simple steps to start using Flatpak](https://flatpak.org/setup/Chrome%20OS)

早速やってみよう。

```text
$ sudo apt install flatpak
```

実際はこんな感じ。

{{< fig-img src="./apt-install-flatpak.png" title="Install Flatpck" link="./apt-install-flatpak.png" width="736" >}}

{{% keytop %}}`y`{{% /keytop %}} を入力しインストールを完了する。
その後

```text
$ flatpak --user remote-add --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
```

として [Flathub] リポジトリを追加。

{{< fig-img src="./add-flathub.png" title="Add Flathub" link="./add-flathub.png" width="736" >}}

ん？ これでええんか？ まぁいいや。
いったん Linux サブシステムをシャットダウンして再起動する。

{{< fig-img src="./shutdown-linux.png" title="Linux をシャットダウン" link="./shutdown-linux.png" >}}

## Flatpak 版 Firefox を導入する

ターミナルを立ち上げ直したら

```text
$ flatpak install firefox
```

を起動して Firefox インストールを開始する。

{{< fig-img src="./install-firefox-1.png" title="Install Firefox (1)" link="./install-firefox-1.png" width="736" >}}

ええつと， stable でいいかな。

{{< fig-img src="./install-firefox-2.png" title="Install Firefox (2)" link="./install-firefox-2.png" width="736" >}}

23.08 ってのがよー分からんのだが，まぁいいや。
{{% keytop %}}`y`{{% /keytop %}} を入力して次へ。

{{< fig-img src="./install-firefox-3.png" title="Install Firefox (3)" link="./install-firefox-3.png" width="736" >}}

{{% keytop %}}`y`{{% /keytop %}} を入力してインストールを完了させる。

{{< fig-img src="./install-firefox-4.png" title="Install Firefox (4)" link="./install-firefox-4.png" width="736" >}}

これで Chromebook のランチャーを見ると

{{< fig-img src="./launcher.png" title="Chromebook ランチャー" link="./launcher.png" width="648" >}}

よしよし。
Firefox が追加されてるな。
さっそく起動してみよう。

{{< fig-img src="./firefox.png" title="Firefox" link="./firefox.png" width="1366" >}}

んー。
右上の 最小化・最大化・閉じる のボタンが表示されてないなぁ。
カーソルを当てるとツールチップは表示されるのでボタン自体は存在するみたいだが...

バージョンを見てみると

{{< fig-img src="./firefox-version.png" title="Firefox version" link="./firefox-version.png" width="718" >}}

となっているので，一応最新が入ったのかな。
ちなみに [Flatpak] で管理しているパッケージを更新するには

```text
$ flatpak update
```

とすればまとめて更新できるらしい。

## Firefox の日本語化

そうそう。
インストールしたての Firefox は英語表記なので[^l1] 設定（Settings） → 一般（General） → 言語（Language） で日本語を追加する。

[^l1]: Chromebook で有効にした Linux サブシステムは既定で英語表記になっている。その辺のチューニングもしないといけないのだが，この時点では困ってないので放置している。 Linux サブシステムの日本語化については「[日本語化]({{< relref "./chromebook-2.md#jp" >}})」の節を参考のこと。

{{< fig-img src="./firefox-settings.png" title="Firefox 設定" link="./firefox-settings.png" width="1267" >}}

フォント（Fonts）は Noto Sans CJK JP でいいだろう。

## ブックマーク

- [chromebook FireFoxの公式インストール Flatpak – FlatterMe](/itsupport/526/2022/05/29/)

[前回]: {{< ref "/remark/2024/04/chromebook-3.md" >}} "Chromebook を導入する 3 — GnuPG & OpenSSH"
[Andorid 版 Firefox]: https://play.google.com/store/apps/details?id=org.mozilla.firefox&hl=en_US "Firefox Fast & Private Browser - Apps on Google Play"
[Flatpak]: https://www.flatpak.org/ "Flatpak—the future of application distribution"
[Flathub]: https://flathub.org/ "Flathub - Apps for Linux"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B00WAMAKZQ" %}} <!-- コマンドー -->
