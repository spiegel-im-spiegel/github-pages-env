+++
title = "Chromebook を導入する（2）"
date =  "2024-04-02T19:45:48+09:00"
description = "Linux サブシステムを有効にする"
image = "/images/attention/kitten.jpg"
tags = [ "chromebook", "linux" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Chromebook を導入する（1）]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する（2）]({{< ref "/remark/2024/04/chromebook-2.md" >}}) ← イマココ

## Linux サブシステムを有効にする

[前回]の続き。
Linux サブシステムを有効にするには設定から 詳細設定 → デベロッパー と辿って「Linux 開発環境」を「オンにする」ボタンを押下する。

{{< fig-img src="./linux-1.png" title="デベロッパー" link="./linux-1.png" width="1100" >}}

以降は画面のとおりにインストールを進めればよい。

{{< fig-img src="./linux-2.png" title="Linux 開発環境をセットアップする (1)" link="./linux-2.png" width="782" >}}

{{< fig-img src="./linux-3.png" title="Linux 開発環境をセットアップする (2)" link="./linux-3.png" width="782" >}}

ユーザ名はお好みで。
ディスクサイズは推奨値のまま「次へ」進む。

{{< fig-img src="./linux-4.png" title="Linux 開発環境をセットアップする (3)" link="./linux-4.png" width="782" >}}

インストール処理はマジで時間がかかるので，ちょっとコーヒーブレイク。
インストールが完了するとターミナルが立ち上がる。

{{< fig-img src="./linux-5.png" title="ターミナルが起動" link="./linux-5.png" width="736" >}}

Chromebook の Linux サブシステムは Debian ベースらしい。
確認してみよう。

{{< fig-img src="./linux-6.png" title="ディストリビューションとバージョンを確認" link="./linux-6.png" width="736" >}}

ふむ。
Debian 12 か。
今の最新って12.5 だっけ？ まぁいいや。
Debian のパッケージ管理は APT なので `apt update/upgrade` を行って最新状態にすること。

Linux サブシステムを導入した状態で 設定 → 詳細設定 → デベロッパー → Linux 開発環境 と辿ると以下のような設定画面が表示される。

{{< fig-img src="./linux-7.png" title="Linux 開発環境" link="./linux-7.png" width="1100" >}}

この画面を使ってディスクサイズを変更できる。
また Linux 開発環境の削除も簡単なので，何度でもやり直せる（実際に結構やり直した）。

Files を見てみると Linux サブシステムがひとつのファイルとして見えていることがわかる。

{{< fig-img src="./linux-in-chromebook.png" title="Linux in Chromebook" link="./linux-in-chromebook.png" width="1100" >}}

ちなみに Files からは「Linux ファイル」の中身は見えない。

今回買った [Chromebook] はメインストレージが32GBしかないため Linux サブシステム用に10GBも使うと残り容量が心もとない感じになる。

{{< fig-img src="./storage.png" title="ストレージ管理" link="./storage.png" width="1100" >}}

なので慌てて追加で microSD カードを買った。

{{% review-paapi "B0CH2X5LBX" %}} <!-- micoroSD カード -->

一時的なデータの置き場はこれでいいけどメインストレージの不足はどうしようもない。
Linux サブシステムを使うなら他サブシステムのアプリは使えないと割り切ったほうがいいかもなぁ。
まぁ Chromebook でゲームをする気はないし， Kindle アプリも別に端末を持ってる（またはブラウザで [Cloud Reader](https://read.amazon.co.jp/) を見る）から要らんかな。

Chromebook のディレクトリを共有ディレクトリとして設定できる。

{{< fig-img src="./shared-directory.png" title="Linux と共有" link="./shared-directory.png" width="1100" >}}

以下は指定した共有ディレクトリに対してシンボリックリンクを張ったところ。

{{< fig-img src="./shared-directory-2.png" title="共有ディレクトリに対してシンボリックリンクを張る" link="./shared-directory-2.png" width="736" >}}

今回はここまでかな。

[前回]: {{< ref "/remark/2024/03/chromebook-1.md" >}} "Chromebook を導入する（1）"
[Chromebook]: https://www.amazon.co.jp/gp/product/B0BKKF7JHV?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Amazon.co.jp: 【Amazon.co.jp限定】 ASUS Chromebook クロームブック Flip CX1 11.6インチ 2-in-1 タッチスクリーン 日本語キーボード 重量1.32kg トランスペアレントシルバー CX1102FKA-MK0037 : パソコン・周辺機器"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B00G9NIL7G" %}} <!-- エレコム マウス Bluetooth -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->