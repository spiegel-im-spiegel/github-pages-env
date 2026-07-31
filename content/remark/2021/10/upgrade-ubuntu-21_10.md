+++
title = "Ubuntu 21.10 へのアップグレードを試してみた"
date =  "2021-10-17T19:16:37+09:00"
description = "やっぱり日本語 Remix が出るまでは待ったほうがよさそうだ。待って状況が改善するとは限らないけど。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "firefox", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] が「Ubuntu 21.10 にアップグレードできるけど，どーする」（←超意訳）と訊いてきたのだが

- [Ubuntu 21.10で日本語入力ができない問題](https://zenn.dev/ikuya/articles/788626c3ae6ade)
- [Ubuntu 21.10 その20 - deb版FirefoxからSnap版Firefoxへ移行するパッケージの提供 - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2110-20-debfirefoxsnapfirefox.html)

みたいな怖い話を聞くので躊躇している。
でも，まぁ，せっかくなので自宅のサブ PC で試してみることにした。

{{< fig-img src="./release-note.png" title="Ubuntu 21.10 リリースノート" link="./release-note.png" width="620" >}}

アップグレード自体は特に問題なく終了したのだが，やはり日本語入力が出来ないようだ。
[Zenn の記事](https://zenn.dev/ikuya/articles/788626c3ae6ade "Ubuntu 21.10で日本語入力ができない問題")で紹介されている通り

```text
$ ibus-daemon -rxd
```

で Mozc も動くようになった。
また Gunnar Hjalmarsson さんが [PPA で公開している im-config](https://launchpad.net/~gunnarhj/+archive/ubuntu/im-config "Work around LP 1946969 : Gunnar Hjalmarsson") を導入すればとりあえず行けるようだ。

```text
$ sudo add-apt-repository ppa:gunnarhj/im-config
$ sudo apt upgrade
```

この状態で Snap 版 Firefox を起動してみたが，問題なく日本語も入力できた。
よしよし。

{{< fig-img src="./firefox-snap-version.png" title="Mozilla Firefox Snap for Ubuntu" link="./firefox-snap-version.png" width="707" >}}

ただフォントの設定が壊れているみたいで，以前[遊びで入れた楷書体フォント]({{< ref "/remark/2019/05/gestaltzerfall-of-reiwa.md" >}} "「令」のゲシュタルト崩壊")が既定フォントになっていたので慌てて修正した。

ちなみに Snap 版 LiberOffice は全く使いものにならなかった。
まぁ，今は APT 版 LiberOffice で問題なく OpenPGP 署名・暗号化ができるので困ってないんだけど。
ホンマ Snap って日本語が冷遇されてるよなぁ。

というわけで，やっぱり[日本語 Remix](https://www.ubuntulinux.jp/products/JA-Localized) が出るまでは待ったほうがよさそうだ。
待って状況が改善するとは限らないけど。

そろそろ[日本語コミュニティ](https://www.ubuntulinux.jp/ "Homepage | Ubuntu Japanese Team")にでも参加したほうがいいかねぇ。

[後半]({{< ref "/remark/2021/11/upgrade-ubuntu-21_10-part-2.md" >}} "Ubuntu 21.10 へのアップグレード（本番）")へ続く...

## ブックマーク

- [Bug #1946969 “ibus-x11 does not start automatically” : Bugs : mozc package : Ubuntu](https://bugs.launchpad.net/bugs/1946969)
- [Ubuntu Fridge | Ubuntu 21.10 (Impish Indri) released](https://ubuntu-news.org/2021/10/14/ubuntu-21-10-impish-indri-released/)
- [Impish Indri Release Notes - Release - Ubuntu Community Hub](https://discourse.ubuntu.com/t/impish-indri-release-notes/21951)
- [Ubuntu 21.10 その22 - Ubuntu 21.10がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2110-22-ubuntu-2110.html)
- [Ubuntu 21.10 その23 - Ubuntu 21.10の新機能と変更点・既知の問題 - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2110-23-ubuntu-2110.html)
- [Ubuntu 21.04 その56 - Ubuntu 21.10へアップグレードするには・アップグレードの注意事項 - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2104-56-ubuntu-2110.html)
- [Ubuntu 21.10 その25 - セキュリティーの強化とセキュリティーの新機能 - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2110-25.html)
- [Ubuntu 21.10 その30 - インストーラーの仕様が変わる・BIOS環境でもEFIシステムパーティションが必要に - kledgeb](https://kledgeb.blogspot.com/2021/10/ubuntu-2110-30-biosefi.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
