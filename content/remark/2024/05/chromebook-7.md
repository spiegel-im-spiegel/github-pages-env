+++
title = "Chromebook を導入する 7 — VS Code の導入"
date =  "2024-05-12T17:48:44+09:00"
description = "多分シリーズ最終回。"
image = "/images/attention/tools.png"
tags = [ "chromebook", "linux", "golang", "vscode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}})
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}})
- [Chromebook を導入する 5 — APT で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-5.md" >}})
- [Chromebook を導入する 6 — Git & Go コンパイラ]({{< ref "/remark/2024/05/chromebook-6.md" >}})
- [Chromebook を導入する 7 — VS Code の導入]({{< ref "/remark/2024/05/chromebook-7.md" >}}) ← イマココ

このシリーズとしては多分最終回。
スキマ時間でちまちまやってたから，やたら時間がかかってしまった。

## VS Code のインストール

ググってみると Chromebook の古いバージョンの記事が多く悩ましいのだが，現状では APT 版をインストールするので問題ないようだ。

- [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)

上の記事を参考に以下の手順でインストールした[^apt1]。

[^apt1]: Ubuntu の話であるが APT v2.8 から仕様が色々と変わるらしい（[第812回　aptの新機能あれこれ ［Ubuntu 24.04 LTS版］](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0812)）。特にリポジトリ情報の記述が deb822 形式に変わるようで，おそらく ChromeOS にも影響が出てくるはず。ちなみに現行（2024-05 時点）の APT のバージョンを見たら 2.6.1 だった。

```text
$ cd /etc/apt/keyrings/
$ sudo curl -L https://packages.microsoft.com/keys/microsoft.asc -O
$ echo "deb [arch=amd64,arm64,armhf signed-by=/etc/apt/keyrings/microsoft.asc] https://packages.microsoft.com/repos/code stable main" | sudo tee /etc/apt/sources.list.d/vscode.list > /dev/null
$ sudo apt update
$ sudo apt install code
```

これでランチャーに VS Code のアイコンが入る。

{{< fig-img src="./launcher.png" title="Chromebook ランチャー" link="./launcher.png" width="647" >}}

さっそく[前回]作った `hello` プロジェクトのディレクトリに入って `code .` で起動してみる。

{{< fig-img src="./vscode.png" title="VS Code" link="./vscode.png" width="1024" >}}

うんうん。
問題ないかな。
[自宅機]({{< ref "/remark/2021/06/new-machine-here.md" >}} "自宅マシンを買うた（これで私も人並みに...）")に比べれば若干もったりしてるが，気になるほどではない。
ChromeOS 標準のエディタに比べればキビキビ動く（笑）

細かい設定や [Go] や markdown 関連の設定については以前に書いた記事を参照のこと。

- [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})
- [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}})
- [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}})

一部記述が古くなってるが，まだ使えるな。
何でも取っておくものである。

## Inconsolata フォントのインストール

私はコード表示用のフォントに昔から Inconsolata を使っている。

今回の Chromebook には既定で入ってないようなので入れておく。
フォントは GitHub のリポジトリからダウンロードできる。

- [Releases · googlefonts/Inconsolata · GitHub](https://github.com/googlefonts/Inconsolata/releases)

この中の OTF フォントを `/usr/local/share/fonts` ディレクトリに突っ込んで `fc-cache -fv` とすればOK。
確認は `fc-list` コマンドでできる。

VS Code で使う場合は設定の “Editor: Font Family” の項目に設定すればよい。

{{< fig-img src="./vscode-settings-font-family.png" title="Editor: Font Family" link="./vscode-settings-font-family.png" width="1200" >}}

VS Code は複数のフォントを（優先順位を付けて）指定できるのがいいよね。

## ファンクションキーがない？

VS Code のキーアサインの設定をしようとして気がついたのだが Chromebook にはファンクションキーがなかった。
いや，いまさら気づくとか `orz`

- [Chromebook でファンクションキーを使う方法。F1 から F12 までどこに割り振られているのか？！](https://nj-clucker.com/chromebook-function-key/)

なるほど {{% icons "mglass-key" %}} キーと組み合わせて使うのか。
覚えておこう。
しかし `Shift+Ctrl+F1` とかって組み合わせたらエラいことになるな。

[前回]: {{< ref "/remark/2024/05/chromebook-6.md" >}} "Chromebook を導入する 6 — Git & Go コンパイラ"
[Go]: https://go.dev/ "The Go Programming Language"

## ようやく

Chromebook のセットアップが終わったよ。
これで外出先でも遊べるようになった。

でも1kg強の重量は重いよなぁ。
このスペックでせめて1kg未満にならんもんかねぇ。
それとも Chromebook を担ぐためにもう少し容量の大きいメッセンジャーバッグかリュックを買うべきだろうか。
それも本末転倒だよなー。

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B07KJWYQJW" %}} <!-- ANKER PowerExpand USB メディアハブ -->
{{% review-paapi "B08P54PQDB" %}} <!-- メッセンジャーバッグ -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
