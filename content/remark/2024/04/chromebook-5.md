+++
title = "Chromebook を導入する 5 — APT で Firefox を導入する"
date =  "2024-04-28T01:04:20+09:00"
description = "Chromebook にまともな Firefox を入れたい その2"
image = "/images/attention/tools.png"
tags = [ "chromebook", "linux", "firefox", "flatpak" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}})
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}})
- [Chromebook を導入する 5 — APT で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-5.md" >}}) ← イマココ

[前回]は Flatpak を使って Firefox を導入したが，今回は APT を使って導入してみる。

## Linux サブシステムの日本語化

その前に Linux サブシステムの日本語化を済ませておく。
一応[「Chromebook を導入する 2」にも追記]({{< relref "./chromebook-2.md#jp" >}})したのだが，こちらにも書いておく。
つっても

```text
$ sudo localectl set-locale LANG=ja_JP.UTF-8 LANGUAGE="ja_JP:ja"
```

を入力するだけでよい。
さらに Chrome ブラウザで `chrome://flags` を開き `crostini` で検索し ”Crostini IME support for QT applications” を enabled にすると幸せになれるかもしれない。

{{< fig-img src="../chromebook-2/experiments.png" title="Experiments" link="../chromebook-2/experiments.png" width="1100" >}}

あちこちググって `fcitx-mozc` パッケージを入れろという記述を多く見かけたが，最新版[^v1] では既に日本語 input method が有効になっているため不要，というか無用。
日本語フォントも既定で NOTOフォントの CJK JP フォントが一通り入ってるため，追加インストールは不要のようだ。

[^v1]: 今回の Chromebook では ChromeOS バージョン 123 が入っていた。

## APT 版 Firefox を導入する

APT 版 Firefox の導入手順は以下のページの「[Debian ベースのディストリビューションに Firefox の .deb パッケージをインストールする](https://support.mozilla.org/ja/kb/install-firefox-linux#w_debian-be-sunodeisutoribiyu-shiyonni-firefox-no-deb-patsuke-ziwoinsuto-rusuru)」を参考にした。

- [Linux に Firefox をインストールする | Firefox ヘルプ](https://support.mozilla.org/ja/kb/install-firefox-linux)

まずは，電子署名検証用の公開鍵をダウンロードする。

```text
$ wget -q https://packages.mozilla.org/apt/repo-signing-key.gpg -O- | sudo tee /etc/apt/keyrings/packages.mozilla.org.asc > /dev/null
```

次に Firefox のリポジトリを登録する。

```text
$ echo "deb [signed-by=/etc/apt/keyrings/packages.mozilla.org.asc] https://packages.mozilla.org/apt mozilla main" | sudo tee -a /etc/apt/sources.list.d/mozilla.list > /dev/null
```

更に，このリポジトリを優先して参照するよう設定する。

```text
$ echo '
Package: *
Pin: origin packages.mozilla.org
Pin-Priority: 1000
' | sudo tee /etc/apt/preferences.d/mozilla
```

これ準備 OK。
`sudo apt update` で最新にアップデートした後，インストールを行う。

```text
$ sudo apt-get update && sudo apt-get install firefox
```

日本語のパッケージも入れておこう。

```text
$ sudo apt-get install firefox-l10n-ja
```

Firefox のアイコンがランチャーに登録されているので，ここから起動する。

{{< fig-img src="./firefox.png" title="Experiments" link="./firefox.png" width="1366" >}}

よーし，うむうむ，よーし。

Flatpak 版は右上のボタンが表示されなくて往生したが APT 版は問題なく表示されているようだ。
Firefox を入れるなら APT 版のほうがいいかも知れない。

## 日本語が入力できない

実は Firefox のアドレスバーやページ内のテキストボックスなどに文字を入力する際に input method の状態に関わらず日本語が入力できない。
これは Flatpak 版も APT 版も同じ症状。
ターミナル等では問題なく日本語を入力できるのに。
いろいろ試してみた（それこそ  `fcitx-mozc` パッケージも入れてみた）が上手くいかない。
ChromeOS の設定というより Firefox 側の問題な気がする。
今回は諦めた。
致命傷ではないので先に進もう。

求む！ 情報！

## ブックマーク

- [[2023年5月版]Chromebook Linux でIMEの設定がめっちゃ楽になってた](https://zenn.dev/asopitech/articles/20230516-103621_1)
- [Chromebookを快適な開発環境にするためのプチノウハウ群（2024年2月版） #Linux - Qiita](https://qiita.com/komde/items/25b4c80598d7e2b679f6)

[前回]: {{< ref "/remark/2024/04/chromebook-4.md" >}} "Chromebook を導入する 4 — APT で Firefox を導入する"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
