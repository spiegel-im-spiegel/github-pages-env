+++
title = "Ubuntu 19.04 へのアップグレードを試す"
date =  "2019-04-20T09:08:40+09:00"
description = "こういうアップグレードとかやると，いかに Windows がシステムとして駄目か分かるよな。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "openssl", "ssh", "gnupg", "git", "mono" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] 19.04 がリリースされたようだ。

[Ubuntu 19.04 その15 - Ubuntu 19.04がリリースされました・ディスクイメージのダウンロード - kledgeb](https://kledgeb.blogspot.com/2019/04/ubuntu-1904-15-ubuntu-1904.html)

私が検証中の仮想環境[^vb1] でもソフトウェアの更新[^upgrd1] 後に以下のメッセージが出た。

[^upgrd1]: [Ubuntu] では「ソフトウェアの更新」で GUI によるソフトウェアのアップデートが可能である。また「ソフトウェアとアップデート」を使って更新の自動化や LivePatch の設定も可能だ。
[^vb1]: 余談だが [VirtualBox] でも 6.0.6 がリリースされている。アップデートは計画的に。

{{< fig-img src="./upgrade-sign.png" link="./upgrade-esign.png" width="645" >}}

折角なのでアップグレードを試してみるとするか（仕事用のマシンじゃないので LTS バージョンを使う気は更々ない）。

## [Ubuntu] のアップグレード

「アップグレード」ボタンを押すとリリースノートが表示される。

{{< fig-img src="./upgrade-relnote.png" link="./upgrade-relnote.png" width="600" >}}

文章中の空白文字が詰められていて「なんじゃこりゃ」な画面だが（何かの署名かと思ったぜ`w`），気にせず「アップグレード」ボタンを押す。

{{< fig-img src="./upgrade-prepare.png" link="./upgrade-prepare.png" >}}

これでしばらく経過を眺めていると以下のワーニングが出る。

{{< fig-img src="./upgrade-warning1.png" link="./upgrade-warning1.png" width="1533" >}}

どうも [Ubuntu Japanese Team] や他のサードパーティのリポジトリは外されてしまうらしい。
まぁ，これはしょうがない。
[あとで繋げばいいし]({{< relref "#srclist" >}})。

処理を続けると確認画面が表示されるので追加・変更・削除されるソフトウェアを確認しておく。
どうやら [GnuPG], [OpenSSL], [OpenSSH] といったセキュリティ関連のツールもアップグレードされるようだ。

{{< fig-img src="./upgrade-confirm2.png" link="./upgrade-confirm2.png" width="1026" >}}

{{< fig-img src="./upgrade-confirm3.png" link="./upgrade-confirm3.png" width="1026" >}}

確認して問題ないようなら「アップグレードを開始」ボタンを押して処理を続行する。
アップグレードは時間がかかるのでお茶でも淹れてこよう。

途中でワーニングが出たりもしたが[^w1]

[^w1]: あとで確認したが，ちゃんと最新版が入ってるっぽい。

{{< fig-img src="./upgrade-warning2.png" link="./upgrade-warning2.png" width="1190" >}}

気にせず最後まで終わらせる。
終わったら（再起動しろとは言われなかったが）一応再起動しておく。

再起動後，気になっていたセキュリティ・ツールのバージョンをチェックしてみた。

```text
$ gpg --version
gpg (GnuPG) 2.2.12
libgcrypt 1.8.4
Copyright (C) 2018 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: /home/username/.gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2

$ openssl version
OpenSSL 1.1.1b  26 Feb 2019

$ ssh -V
OpenSSH_7.9p1 Ubuntu-10, OpenSSL 1.1.1b  26 Feb 2019
```

おおっ。
[GnuPG] はまだちょっと古いが [OpenSSL] と [OpenSSH] は最新になっている。
まぁこのくらいなら許容範囲だろう[^ver1]。

[^ver1]: [OpenSSL] は 1.1.1 系の最新， [OpenSSH] は 7.9 系の最新になっていた。 Facebook の TL で教えてもらったが， [GnuPG] はパッケージマネージャでパッケージの完全性検証に使われるため特に保守的な管理になっているらしい。まぁ脆弱性や不具合等は随時バックポートされているそうなので，バージョン番号であまり神経質にならないほうがいいのかも知れない。ただ [GnuPG] の動向を追いかけている身としては何とか改善したい。でもそれは後々のお楽しみということで。

## サードパーティのリポジトリを再び有効にする{#srclist}

アップグレード時に外されたサードパーティのリポジトリをチェックしておく。
「ソフトウェアとアップデート」を開いて「他のソフトウェア」タブを見てみる。

{{< fig-img src="./upgrade-srclist.png" link="./upgrade-srclist.png" width="972" >}}

うん，外れてるね。
[Git] と [Mono] のリポジトリは有効にしておく。
[Ubuntu Japanese Team] のリポジトリも有効にしておけばいいかな。

これで `apt update` と，必要に応じて `apt upgrade` や `apt autoremove` を行えば OK。

## しかし，なんだな。

こういうアップグレードとかやると，いかに Windows がシステムとして駄目か分かるよな。
パソコンでもスマホのような携帯端末でも定期的にアップデートされることを前提にハードウェアもソフトウェアも構成しないと駄目だということだよねぇ。

いまや Office ツールは Microsoft の独占というほどでもないし，大抵のことは Web インタフェースがあるのでブラウザがあればなんとかなる，と考えるとやはり Windows を捨てる選択は正解だなと改めて思う。
他人に薦められるかと言えば，それは別問題だが（笑）

## ブックマーク

- [PPA から Git をインストールする]({{< relref "./install-git-from-ppa.md" >}})
- [Ubuntu に Mono を導入する]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Ubuntu Japanese Team]: http://www.ubuntulinux.jp/
[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSL]: https://www.openssl.org/
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[Git]: https://git-scm.com/
[Mono]: https://www.mono-project.com/
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Git Extensions]: https://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"
