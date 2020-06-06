+++
title = "中古 PC に Ubuntu 環境を導入する"
date =  "2019-12-01T02:22:24+09:00"
description = "ついカッとなってやった。反省はしない。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "install", "tools", "gnupg", "openssh", "git", "golang", "atom" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ついカッとなってやった。
反省はしない。

久しぶりの衝動買いですよ。

{{< fig-img src="./49144529442_3655cab86e_e.jpg" title="玩具が来た！ 早速 Ubuntu をぶち込んで遊ぶナリ" link="https://www.flickr.com/photos/spiegel/49144529442/" >}}

一応，理由はあって

1. システムが壊れても日常生活でダメージにならない余剰パソコンが欲しかった
2. 技術系のイベントに参加するのにそろそろ Andorid タブレットではキツくなってきた

といったところ。

最初の理由に関しては，最初は来年あたりにラズパイのキットでも買って遊ぼうかとも思っていたのだが，一通り揃ってるキットを買うと数千円はするみたいだし，それなら1万円前後の中古 PC でもいいんじゃないかって感じ。

今回は[ドスパラ](https://www.dospara.co.jp/)で中古の DELL Lititude E5540 を購入した。
送料やら消費税やら全部込み込みで13K円ほど。
OS は入っていない。
もっと安いものもあったが，動作保証しないとか怖いことが書いてあったので止めた。
あーゆーのって部品取りするためのものなのかねぇ。

今回は [Ubuntu] を導入して [ATOM] エディタ上で [Go 言語]コードをコンパイル&テストできるところまでが目標。
メインの PC に [Ubuntu] 環境が既にあって，そこから設定をコピーする。

## [Ubuntu] のインストール

いざというときのために [USB メモリでブータブル・メディアを作成済み]({{< ref "/remark/2019/04/bootable-usb.md" >}} "Ubuntu インストール用のブート可能 USB メモリを作成する")なので，これを使って [Ubuntu] をインストールする。

DELL パソコンの BIOS 画面がなかなか出なくて往生したが，電源投入直後に表示される「DELL」ロゴ画面で F2 キーを連打すればいいらしい。
なんじゃそら（笑）

BIOS 画面で USB からのブートを最優先にし，改めてブータブル・メディアを指してインストールを開始する[^bios1]。
[日本語 Remix](https://www.ubuntulinux.jp/japanese) 版を使えばインストーラも日本語だし悩むところはないのだが，ノート PC なので（紛失・盗難のリスクを考慮して）ストレージの暗号化を忘れないこと。

[^bios1]: BIOS 設定はあとで戻しておくこと。ついでに BIOS を[パスワードロック](https://www.dell.com/community/%E4%B8%80%E8%88%AC-%E3%83%8E%E3%83%BC%E3%83%88%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3/BIOS%E3%81%A7%E8%A8%AD%E5%AE%9A%E3%81%A7%E3%81%8D%E3%82%8B%E3%83%91%E3%82%B9%E3%83%AF%E3%83%BC%E3%83%89%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6/td-p/6225663 "BIOSで設定できるパスワードについて - Dell Community")しておかないと。

私の場合はセットアップに必要なファイルを NAS に置いているので，インストールが完了したら NAS への接続設定をしておく。
前のときは「[CIFS 経由で NAS に接続]({{< ref "/remark/2019/03/common-internet-file-system.md" >}})」していたが，今回は（外に持ち出すことも考慮して）ファイルマネージャ Nautilus 上で動的に接続できるよう構成してみた。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## KeePassXC のインストール

パスワード・マネージャ [KeePassXC] のインストールについては以下を参照のこと。

- [Ubuntu に KeePassXC を導入する]({{< ref "/remark/2019/08/installing-keepassxc-in-ubuntu.md" >}})

これでようやく無線 LAN に繋げられる。
有線からの解放。

[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"

## [GnuPG] 環境の移行

[GnuPG] 環境の移行および [OpenSSH] との連携については以下を参照のこと。

- [Windows 環境で作った GnuPG の鍵束を Ubuntu に移行する]({{< ref "/remark/2019/04/move-gpg-keyring.md" >}})

つっても，今回は `~/.gnupg` ディレクトリの中身をまるっとコピればいいだけなんだけど。

[OpenSSH] と連携するなら `gpg-agent.conf` ファイルで以下のオプションを設定し

```text
# for ssh
enable-ssh-support
```

環境変数 `SSH_AUTH_SOCK` の設定も忘れないこと。

```bash
export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
```

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest/

## 基本ツールのインストール

APT で基本ツールをインストールしておく。

```text
$ sudo apt install net-tools curl kdiff3 graphviz vim xsel
```

[ATOM] エディタで使いたいので [Inconsolata] フォントも入れておく。

`Inconsolata.otf` ファイルを取ってきて `/usr/local/share/fonts` ディレクトリに入れてキャッシュを更新する。

```text
$ fc-cache -fv
```

更新されたかどうかは以下で確認できる。

```text
$ fc-list | grep Inconsolata
/usr/local/share/fonts/Inconsolata.otf: Inconsolata:style=Medium
```

[Inconsolata]: https://www.levien.com/type/myfonts/inconsolata.html

## [Git] のインストール

[Git] をインストールする前に `~/.config/git/config` ファイルをコピっておく。
既に `~/.gitconfig` ファイルがあるなら `~/.config/git/config` へ `mv` する。

インストールは以下を参照のこと。

- [PPA から Git をインストールする](https://text.baldanders.info/remark/2019/04/install-git-from-ppa/)

[git]: https://git-scm.com/
[Git]: https://git-scm.com/

## [Go] コンパイラのインストール

APT で管理される [Go] コンパイラはバージョンが古すぎるので公式サイトの[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")から直接バイナリを取ってきて展開する。

たとえば [Go] 1.13.4 をインストールしたいなら

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.4.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.4.linux-amd64.tar.gz
$ sudo mv go go1.13.4
$ sudo ln -s go1.13.4 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.4 linux/amd64
```

などとする。

これでようやく [GnuPG], [OpenSSH], [Git], [Go] コンパイラと必要なツールのセットアップが完了した。

[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## [ATOM] エディタのインストール

[ATOM] エディタは公式の APT リポジトリが用意されているので，そちらを利用する。

まずはリポジトリの登録。

```text
$ wget -qO - https://packagecloud.io/AtomEditor/atom/gpgkey | sudo apt-key add -
$ sudo sh -c 'echo "deb [arch=amd64] https://packagecloud.io/AtomEditor/atom/any/ any main" > /etc/apt/sources.list.d/atom.list'
$ sudo apt update
```

これであとは

```text
$ sudo apt install atom
```

とすれば，インストール完了。

### パッケージをまとめてインストールする

[ATOM] エディタのサード・パーティ・パッケージにあらかじめ star (☆) を付けておくことでまとめてインストールすることができる。

まずはログインから

```text
$ apm login
Welcome to Atom!

Before you can publish packages, you'll need an API token.

Visit your account page on Atom.io https://atom.io/account,
copy the token and paste it below when prompted.

Press [Enter] to open your account page on Atom.io.
```

ここで `[Enter]` キーを押すと Web ブラウザでアカウントページが開く。
開かない場合は Web ブラウザで直接 https://atom.io/account を開く。
アカウントページに表示されたアクセス・トークンを入力すればログイン完了。

```text
Token> ****************
Saving token to Keychain done
```

Star を付けたパッケージは以下のコマンドで見れる。

```text
$ apm stars
```

Star を付けたパッケージをまとめてインストールするなら

```text
$ apm stars --install
```

でOK。

### 設定ファイルの移行

`~/.atom` ディレクトリ直下にある以下のファイルをコピれば簡単に設定を移行できる。

- `config.cson`
- `init.coffee`
- `keymap.cson`
- `styles.less`
- `snipets.cson`

これで完了！

[ATOM]: https://atom.io/

## その他お好みで

- [Ubuntu に Mono を導入する](https://text.baldanders.info/remark/2019/04/mono-in-ubuntu/)
- [TeX Live を Ubuntu に（APT を使わずに）導入する]({{< ref "/remark/2019/04/install-texlive-in-ubuntu.md" >}})
    - [TeX Live 2018 から 2019 へのアップグレード]({{< ref "/remark/2019/06/upgrade-texlive-from-2018-to-2019.md" >}})
- [BOINC による学術分散コンピューティング・プロジェクトでの活動を再開した]({{< ref "/remark/2019/04/academic-distributed-computing-projects-by-boinc.md" >}})
- [Ununtu に ClamAV を導入する]({{< ref "/remark/2019/05/clamav-for-ubuntu.md" >}})
- [Ubuntu に LibreOffice をインストールする3つの方法]({{< ref "/remark/2019/05/installing-libreoffice-in-ubuntu.md" >}})
    - [LibreOffice 6.3 へのアップグレード]({{< ref "/release/2019/08/upgrade-libreoffice-6_3.md" >}})
- [Ubuntu でも Kindle 本が読みたい]({{< ref "/remark/2019/05/kindle-for-wine.md" >}})
- [メール・サービスを立てずにコマンドラインでメールを送信する]({{< ref "/remark/2019/06/send-mail-without-mail-service.md" >}})
- [そろそろ Vuls を唱えるか]({{< ref "/remark/2019/06/cast-vuls.md" >}})
- [Ubuntu で音楽 CD のリッピング]({{< ref "/remark/2019/06/ripping-cd-music-in-ubuntu.md" >}})
- [音楽プレイヤー Lollypop を試す](https://text.baldanders.info/remark/2019/06/lollypop-music-player/)
- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})
- [結局 Thunderbird もインストールし直すことにした]({{< ref "/remark/2019/11/reinstalling-thunderbird.md" >}})

## ブックマーク

- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})
- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})
- [Advanced Package Tool に関する覚え書き]({{< ref "/remark/2019/05/advanced-package-tool.md" >}})
- [Ubuntu アプリケーションにおけるセキュリティ・アップデート一覧]({{< ref "/release/vuln-list.md" >}})

{{% review-paapi "B01NBU1OS5" %}} <!-- シリコンパワー USBメモリ 32GB USB3.1 -->
