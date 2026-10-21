+++
title = "ArtCraft/PdfCraft を導入する"
date =  "2026-10-21T17:56:55+09:00"
description = "PdfCraft をはじめとするツール群は AI エージェントと連携しやすいよう MCP サーバとしても動作するらしい。この辺がイマドキだな（笑）"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "linux", "ubuntu", "pdf", "install", "flatpak" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Bluesky の TL を眺めてたら

- [「Linux」に無料の「Acrobat」クローンが登場--本格的なPDF編集アプリ](https://japan.zdnet.com/article/35253752/)

という記事を見かけた。

{{< fig-quote type="markdown" title="「Linux」に無料の「Acrobat」クローンが登場--本格的なPDF編集アプリ" link="https://japan.zdnet.com/article/35253752/" >}}
そんな中、「Adobe Photoshop」のクローンである「PhotoCraft」を開発したArtCraftが、今度は「PdfCraft」をリリースした。「Adobe Acrobat」に取って代わるこのアプリのおかげで、Linuxでもようやく本格的なPDFの編集ができるようになる。

ArtCraftの他のアプリと同じく、PdfCraftは「Rust」で書かれたもので、人工知能（AI）を用いたリバースエンジニアリングによって、Linux、「FreeBSD」、「macOS」、「Windows」向けのアプリが開発されている。このアプリを利用するのにアカウントは不要で、テレメトリーデータが送信されることは一切なく、クラウドに依存する機能もまったくない。完全にローカルで動作し、完全に無料で、すべてを自分自身で管理できる。
{{< /fig-quote >}}

ほほう。
面白そうだな。早速試してみよう。

## ArtCraft と PdfCraft

[ArtCraft] は "Controllable AI for Artists" とある通り，アーティスト向けに使いやすい AI 環境を提供しているようだ。
[ArtCraft] 本体はバイナリが [OSS](https://github.com/storytold/artcraft "storytold/artcraft: ArtCraft is an intentional crafting engine for artists, designers, and filmmakers") かつ無償で提供されているが，利用するためにはサブスクリプション契約が必要となる。

{{< fig-img-quote src="./pricing.png" title="Pricing — ArtCraft" link="https://getartcraft.com/pricing" lang="en" width="1482" >}}

一方 [PdfCraft] は [ArtCraft] とは異なる製品で [OSS][PdfCraft-GitHub] なのは同じだが，完全に無料で利用できる。
[ArtCraft] アカウントも必要ない。
[ArtCraft] の中身を知らないのだが，もしかして [PdfCraft] って [ArtCraft] のサブツールなのかなぁ（憶測）。

まぁ，でも， AI 関連機能できっちり金を取って，人間が主に使うツールは無償で提供するってのはなかなか面白い収益構造だよなぁ。

## ArtCraft Launcher と PdfCraft

Linux，特に Ubuntu では [Flatpak] でインストールするのがよさそうだ。
[Flatpak] でインストールしたパッケージはサンドボックス上で動作するため，システムを汚さず比較的安全に利用できる。

さらに [PdfCraft] は [ArtCraft Launcher] からインストールするのがよさそう。
[ArtCraft Launcher] は [PdfCraft] をはじめとする製品群の更新情報をチェックして，自動的にアップデートしてくれる。

では早速やってみよう。

## ArtCraft Launcher のインストール

[Flatpak] 版 [ArtCraft Launcher] は標準リポジトリの [Flathub] にないみたいなので [GitHub のリリースページ](https://github.com/storytold/craft-launcher/releases "Releases · storytold/craft-launcher") から `*.flatpak` ファイルをダウンロードする。

```text
$ curl -LO "https://github.com/storytold/craft-launcher/releases/download/v0.6.0/artcraft-launcher-0.6.0-linux-x86_64.flatpak"
$ flatpak install --user artcraft-launcher-0.6.0-linux-x86_64.flatpak
The application ai.storyteller.artcraft-launcher depends on runtimes from:
  https://dl.flathub.org/repo/
Configure this as new remote 'flathub' [Y/n]: y
Required runtime for ai.storyteller.artcraft-launcher/x86_64/stable (runtime/org.freedesktop.Platform/x86_64/26.08) found in remote flathub
Do you want to install it? [Y/n]: y

ai.storyteller.artcraft-launcher permissions:
    ipc                   network               fallback-x11      pulseaudio      wayland      x11      dri
    file access [1]       dbus access [2]

    [1] home
    [2] org.freedesktop.Flatpak, org.freedesktop.Notifications, org.kde.StatusNotifierWatcher

...
```

なるほど。
インストール時に専用リポジトリも併せて追加してるのか。
簡単でいいな。

インストールが成功すると，ドックパネルにある「アプリケーションを表示」ボタンを押すと開くアプリドロワー画面に [ArtCraft Launcher] アイコンが生えている（筈）。
起動するとこんな感じ。

{{< fig-img-quote src="./artcraft-launcher.png" title="ArtCraft Launcher" link="./artcraft-launcher.png" lang="en" width="1120" >}}

なお，私が今回インストールした 0.6.0 は日本語が表示できなかった。
なんでかなぁ...

## PdfCraft のインストール

[ArtCraft Launcher] 画面で [PdfCraft] を `[INSTALL]` する。

{{< fig-img-quote src="./installing-pdfcraft.png" title="Installing PdfCraft" link="./installing-pdfcraft.png" lang="en" width="1120" >}}

インストールした [PdfCraft] は [ArtCraft Launcher] から起動できる。
またアプリドロワー画面に [PdfCraft] アイコンが生えている筈なので，そこからも起動できる。

{{< fig-img-quote src="./pdfcraft.png" title="PdfCraft" link="./pdfcraft.png" lang="en" width="1440" >}}

よかった。
ちゃんと日本語で表示されている。
入力についても日本語で問題なく行えた。
フォームへの入力・保存も問題なし（申請書が PDF フォームになっているファイルとかあるし，大事）。
[PdfCraft] で編集した PDF を Windows 機に入れている Adobe Acrobat で開いてみたが，これも問題なかった。

既定で左側のサイドバーに表示されるメニュー（一部）はこんな感じ（メニュー部分だけコラージュしている）。

{{< fig-img-quote src="./menus.png" title="PdfCraft Menus" link="./menus.png" lang="en" width="1136" >}}

[GitHub リポジトリ][PdfCraft-GitHub]には "PdfCraft is **alpha**" と書かれてるし，実際に一部まだ未実装のものもあるようだが，もう十分に実用レベルなんじゃないかな。

試しに Windows 版も手元の[ミニ PC]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした") に同様の手順で入れてみたが，こちらも問題なく動作した。
というか [PdfCraft] のほうが Acrobat より軽い気がする（体感レベルだけど）。

なお [PdfCraft] をはじめとするツール群は AI エージェントと連携しやすいよう MCP サーバとしても動作するらしい。
この辺がイマドキだな（笑）

しばらく使ってみて問題なければ [PdfCraft] に置き換えてもいいかな。
ホンマ，いい仕事してますねぇ。

## ブックマーク

- [無料の「Photoshop」クローンを「Linux」で試す--不可能が可能に](https://japan.zdnet.com/article/35253562/)
- [「WordCraft」レビュー--無料で使えるオープンソースの「MS Word」クローン](https://japan.zdnet.com/article/35253759/)
- [全てのflatpakアプリのフォントの設定を一括で設定する - *nixで過ごす日々](https://okutom.hatenablog.com/entry/2024/11/07/204055)

[ArtCraft]: https://getartcraft.com/ "ArtCraft — Controllable AI for Artists"
[ArtCraft Launcher]: https://github.com/storytold/craft-launcher "storytold/craft-launcher: The Craft App Launcher"
[PdfCraft]: https://getartcraft.com/apps/pdfcraft "PdfCraft: The open-source PDF workbench. — ArtCraft"
[PdfCraft-GitHub]: https://github.com/storytold/pdfcraft "storytold/pdfcraft: PdfCraft: An open-source PDF reader and editor written by scratch in pure Rust"
[Flatpak]: https://flatpak.org/ "Flatpak—the future of application distribution"
[Flathub]: https://flathub.org/en "Flathub - Apps for Linux"
