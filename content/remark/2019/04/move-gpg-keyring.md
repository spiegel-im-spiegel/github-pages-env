+++
title = "Windows 環境で作った GnuPG の鍵束を Ubuntu に移行する"
date = "2019-04-07T00:01:53+09:00"
description = "鍵束を Ubuntu 環境にコピーして OpenSSH 連携を確認するところまで。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "cryptography", "openpgp", "gnupg", "ssh", "openssh", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さて NAS に接続できてファイルのバックアップや移動が楽にできるようになった。
次は [GnuPG] で遊んでみよう。

## なんぼなんでも古すぎるじゃろ

[Ubuntu] のディストリビューションには最初から [GnuPG] が入っているのだが

```text
$ gpg --version
gpg (GnuPG) 2.2.8
libgcrypt 1.8.3
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
```

いやいや。
なんぼなんでも古すぎるじゃろ[^u1]。
確かに最後にセキュリティアップデートが行われたのは[2018年の 2.2.8 のとき]({{< ref "/release/2018/06/gnupg-2_2_8-and-libgcrypt-1_8_3-are-released.md" >}} "GnuPG 2.2.8 および Libgcrypt 1.8.3 がリリース（セキュリティ・アップデート）")だけどさ。
あれから不具合の修正とかも結構あったのよ。

[^u1]: [Ubuntu] 19.04 では v2.2.12 までサポートしていた。それでもまだ古いけど [GnuPG] はパッケージ管理などで中核のツールとして使われるため運用が特に保守的なんだそうだ。まぁしょうがないか。 [GnuPG] は複数のコンポーネントで成り立っているのでうかつに弄りたくないというのはよく分かる。

調べてみたら [Ubuntu] や [Debian] では [GnuPG] のアップデートに積極的ではない様子。
まぁ使えないことはないので，最新版の導入は~~諦めて~~もとい後々のお楽しみにとっておいて，今回はこの 2.2.8 を使っていろいろやってみることにする。

## なにはともあれ鍵を作らないと

[GnuPG] の使い方の概要は拙文「[GnuPG チートシート]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})」を参考にどうぞ。
さっそく鍵を作ってみる。

```text
$ gpg --quick-gen-key "Alice <alice@example.com>" future-default - 1w
gpg: ディレクトリ'/home/username/.gnupg'が作成されました
gpg: keybox'/home/username/.gnupg/pubring.kbx'が作成されました
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: /home/username/.gnupg/trustdb.gpg: 信用データベースができました
gpg: 鍵7979662E697CDD8Aを究極的に信用するよう記録しました
gpg: ディレクトリ'/home/username/.gnupg/openpgp-revocs.d'が作成されました
gpg: 失効証明書を '/home/username/.gnupg/openpgp-revocs.d/E23197776C20290F831D97747979662E697CDD8A.rev' に保管しました。
公開鍵と秘密鍵を作成し、署名しました。

pub   ed25519 2019-04-04 [SC] [有効期限: 2019-04-11]
      E23197776C20290F831D97747979662E697CDD8A
uid                      Alice <alice@example.com>
```

途中でパスフレーズを定義するために Pinentry が起動する。

{{< fig-img src="./new-passphrase.png" link="./new-passphrase.png" >}}

おおっ。
これが [Ubuntu] の Pinentry か。
シンプルでよろしい。

鍵作成後の `~/.gnupg` ディレクトリの中身はこうなっている。

```text
$ ls -al .gnupg/
合計 28
drwx------  4 username users 4096  4月  4 21:23 .
drwxr-xr-x 16 username users 4096  4月  4 21:23 ..
drwx------  2 username users 4096  4月  4 21:23 openpgp-revocs.d
drwx------  2 username users 4096  4月  4 21:23 private-keys-v1.d
-rw-r--r--  1 username users  634  4月  4 21:23 pubring.kbx
-rw-------  1 username users   32  4月  4 21:23 pubring.kbx~
-rw-------  1 username users 1240  4月  4 21:23 trustdb.gpg

$ ls -al .gnupg/private-keys-v1.d/
合計 16
drwx------ 2 username users 4096  4月  4 21:23 .
drwx------ 4 username users 4096  4月  4 21:23 ..
-rw------- 1 username users  340  4月  4 21:23 6103D260692806ECD8FF194DD5A42A166ECFCA7C.key
-rw------- 1 username users  332  4月  4 21:23 D08B6C062AC8FA95D4526EDADBDA906A74D26D77.key

$ ls -al .gnupg/openpgp-revocs.d/
合計 12
drwx------ 2 username users 4096  4月  4 21:23 .
drwx------ 4 username users 4096  4月  4 21:23 ..
-rw------- 1 username users 1401  4月  4 21:23 E23197776C20290F831D97747979662E697CDD8A.rev
```

ちなみに `openpgp-revocs.d` ディレクトリには秘密鍵が， `openpgp-revocs.d` ディレクトリには鍵の失効証明書が入っている。
失効証明書は作った鍵を執行させるために使うもので，重要なデータなので取り扱いは慎重に。

これを見ると `~/.gnupg` ディレクトリ以下のファイル・ディレクトリは他ユーザから見えないようにするのがいいみたいだね。

作った鍵を使って試しに適当なテキストを署名してみよう。

```text
$ echo Hello world | gpg -u alice --clear-sign
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA256

Hello world
-----BEGIN PGP SIGNATURE-----

iIgEARYIADAWIQTiMZd3bCApD4Mdl3R5eWYuaXzdigUCXKX4dBIcYWxpY2VAZXhh
bXBsZS5jb20ACgkQeXlmLml83YogagD6A6tRm+uEgphlYrORJBZ9oRmFpRzImXO6
22I/HeiXTGkBALU1QLnv/xcPf1pMQRjpxM9SO/IupO/Zt9lryx4lFOAC
=8DRi
-----END PGP SIGNATURE-----
```

このときも Pinentry が起動する。

{{< fig-img src="./unlock-private-key.png" link="./unlock-private-key.png" >}}

よーし，うむうむ，よーし。

## [GnuPG] の環境変数

[GnuPG] の環境変数は以下のコマンドで取得できるようだ。

```text
$ gpgconf --list-dirs
sysconfdir:/etc/gnupg
bindir:/usr/bin
libexecdir:/usr/lib/gnupg
libdir:/usr/lib/x86_64-linux-gnu/gnupg
datadir:/usr/share/gnupg
localedir:/usr/share/locale
socketdir:/run/user/1000/gnupg
dirmngr-socket:/run/user/1000/gnupg/S.dirmngr
agent-ssh-socket:/run/user/1000/gnupg/S.gpg-agent.ssh
agent-extra-socket:/run/user/1000/gnupg/S.gpg-agent.extra
agent-browser-socket:/run/user/1000/gnupg/S.gpg-agent.browser
agent-socket:/run/user/1000/gnupg/S.gpg-agent
homedir:/home/username/.gnupg
```

また shell 側の環境変数は

```text
$ env | grep GPG
GPG_AGENT_INFO=/run/user/1000/gnupg/S.gpg-agent:0:1
```

となっていて `GPG_AGENT_INFO` に `agent-socket` の値がセットされているのが分かる。
古いバージョンで要ると言われた `GPG_TTY` 環境変数は要らなそうである。

もし `GPG_AGENT_INFO` に手動で値をセットするなら

```text
$ export GPG_AGENT_INFO=$(gpgconf --list-dirs agent-socket):0:1
```

って感じでいいかな。

通信ソケットが置かれている `/run/user/1000/gnupg/` ディレクトリはどうやら [XDG Base Directory Specification] に準拠したディレクトリらしい。
`XDG_RUNTIME_DIR` 環境変数が

```text
$ env | grep XDG_RUNTIME_DIR
XDG_RUNTIME_DIR=/run/user/1000
```

と定義されているので多分そうだろう（`1000` は `username` ユーザの UID）。
環境変数 `XDG_RUNTIME_DIR` で示されるディレクトリはユーザがログインしている間のみ存在が保証されているそうだ。

### Pinentry の定義

[Ubuntu] では Pinentry として `pinentry-gnome3` と `pinentry-curses` が組み込まれているようだ。
[Ubuntu] インストール直後の既定は `pinentry-gnome3` のようだが，これを切り替えるには

```text
$ sudo update-alternatives --config pinentry
alternative pinentry (/usr/bin/pinentry を提供) には 2 個の選択肢があります。

  選択肢    パス                    優先度  状態
------------------------------------------------------------
* 0            /usr/bin/pinentry-gnome3   90        自動モード
  1            /usr/bin/pinentry-curses   50        手動モード
  2            /usr/bin/pinentry-gnome3   90        手動モード

現在の選択 [*] を保持するには <Enter>、さもなければ選択肢の番号のキーを押してください: 
```

とすればいいらしい。
実際には

```text
$ ls -lF /usr/bin/pinentry* 
lrwxrwxrwx 1 root root    26  4月  2 21:13 /usr/bin/pinentry -> /etc/alternatives/pinentry*
-rwxr-xr-x 1 root root 63992  7月  6  2018 /usr/bin/pinentry-curses*
-rwxr-xr-x 1 root root 72184  7月  6  2018 /usr/bin/pinentry-gnome3*
lrwxrwxrwx 1 root root    30  4月  2 21:13 /usr/bin/pinentry-x11 -> /etc/alternatives/pinentry-x11*

$ ls -lF /etc/alternatives/pinentry*
lrwxrwxrwx 1 root root 24  4月  2 21:13 /etc/alternatives/pinentry -> /usr/bin/pinentry-gnome3*
lrwxrwxrwx 1 root root 24  4月  2 21:13 /etc/alternatives/pinentry-x11 -> /usr/bin/pinentry-gnome3*
lrwxrwxrwx 1 root root 40  4月  2 21:13 /etc/alternatives/pinentry-x11.1.gz -> /usr/share/man/man1/pinentry-gnome3.1.gz
lrwxrwxrwx 1 root root 40  4月  2 21:13 /etc/alternatives/pinentry.1.gz -> /usr/share/man/man1/pinentry-gnome3.1.gz
```

てな構成になっているみたい。

## Windows 環境で作った [GnuPG] の鍵束を [Ubuntu] にコピーする

では Windows 環境で作成・運用している [GnuPG] の鍵束を Unintu にコピーしてみよう。

必要なファイルは以下の通り。

- `pubring.kbx`： 公開鍵の鍵束
- `trustdb.gpg`： [OpenPGP] 用の信用データベース・ファイル
- `private-keys-v1.d/*.rev`： 秘密鍵ファイル
- `openpgp-revocs.d/*.key`： 失効証明書
- `tofu.db`： [TOFU] 用の信用データベース・ファイル（もしあれば）
- `gpg.conf`： [GnuPG] 用の設定ファイル（もしあれば）
- `gpg-agent.conf`： `gpg-agent` 用の設定ファイル（もしあれば）
- `sshcontrol`： [OpenSSH] 認証用鍵の定義ファイル（もしあれば）

これらのファイルを `~/.gnupg` ディレクトリにディレクトリ構造ごとコピーする（以前のファイルは削除しておいてね）。
`chmod` コマンドでアクセス権を設定するのも忘れずに。

`gpg.conf`, `gpg-agent.conf`, `sshcontrol` 各ファイルの中身はテキストなのだが，改行コードが `CRLF` になっているかもしれないのであらかじめ始末しておくこと[^gnkf1]。

[^gnkf1]: 改行コードを始末するなら拙作の「[nkf っぽいなにか](https://github.com/spiegel-im-spiegel/text "spiegel-im-spiegel/text: Encoding/Decoding Text Package by Golang")」をどうぞ。コマンドラインで `gonkf nwline gpg.conf` とすれば改行コードを LF に変換してくれる。以上宣伝でした（笑）

これで完了。
試しに私の鍵束に入っている [JPCERT/CC の公開鍵](https://www.jpcert.or.jp/jpcert-pgp.html "JPCERT コーディネーションセンター PGP公開鍵")を表示してみた。

```text
$ gpg --list-keys jpcert
pub   rsa2048 2009-06-02 [SCE]
      FC8953BBDC65BD974BDAD1BD317D97A469ECE048
uid           [  充分  ] JPCERT/CC <info@jpcert.or.jp>
sub   rsa2048 2009-06-02 [E]
```

よーし，うむうむ，よーし。

## [OpenSSH] との連携

Windows 環境では [OpenSSH] の認証用鍵を [GnuPG] にインポートして [PuTTY] と連携させていたが，もちろん [Ubuntu] の [OpenSSH] とも連携できる。

[OpenSSH] と連携させるには `gpg-agent.conf` ファイルに `enable-ssh-support` オプションをセットすればいいのだが，これだけでは上手くいかないようだ。

いろいろ試行錯誤した挙げ句，なんの気なしに環境変数を見てみたら

```text
$ env | grep SSH_AUTH_SOCK
SSH_AUTH_SOCK=/run/user/1000/keyring/ssh
```

と定義されていた。
原因はお前かよ！ もっと早く気付け，自分 `orz`

環境変数 `SSH_AUTH_SOCK` を `gpg-agent` のソケットに置き換えるには以下のコマンドラインを叩く。

```text
$ export SSH_AUTH_SOCK=$(gpgconf --list-dirs agent-ssh-socket)
```

これで `SSH_AUTH_SOCK` の値は

```text
$ env | grep SSH_AUTH_SOCK
SSH_AUTH_SOCK=/run/user/1000/gnupg/S.gpg-agent.ssh
```

となった。

この状態でリモートホストに ssh 接続してみよう。

```text
$ ssh username@remotehost
```

これで Pinentry が起動して

{{< fig-img src="./ssh-login.png" link="./ssh-login.png" >}}

パスフレーズを入力後ログインできれば成功。
`SSH_AUTH_SOCK` 値の書き換えコマンドは `~/.bashrc` ファイルにでも書いておけばいいだろうか。

```bash
export -n SSH_AGENT_PID
export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
```

ちなみに `gpg-agent.conf` ファイルに設定できる [OpenSSH] 連携関連のオプションは以下の通り。

| オプション名 | 既定値 | 内容 |
| ------------------------ | ------ | ---- |
| `enable-ssh-support`    | ―      | `ssh-agent` 互換プロトコルを有効にする |
| `default-cache-ttl-ssh` | `1800`  | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する |
| `max-cache-ttl-ssh`     | `7200`  | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる |

[OpenSSH] との連携を調べ始め，試行錯誤してここまでたどり着くまでに3時間くらいかかっちまったよ。
[GnuPG] の公式サイトも含めて資料がなさすぎるんだよ。
特に日本語の記事は内容が古すぎて壊滅状態。
いかに [GnuPG] が使われてないか分かるよなぁ。

まっ，とにかく， [Ubuntu] でも（バージョンの問題に目をつぶれば）普通に [GnuPG] が使えることが分かったので今回はよしとする。

これでようやく git を使う準備ができた。
まだまだ道のりは遠い。

## ブックマーク

- [Ubuntu フォルダー構造 その10 - XDG Base Directory Specificationについて - kledgeb](https://kledgeb.blogspot.com/2013/04/ubuntu-10-xdg-base-directory.html)
- [GNOME/Keyring - ArchWiki](https://wiki.archlinux.org/index.php/GNOME/Keyring)

- [Change pinentry program temporarily with gpg-agent - Unix & Linux Stack Exchange](https://unix.stackexchange.com/questions/236746/change-pinentry-program-temporarily-with-gpg-agent)
- [Using GnuPG (2.1) for SSH authentication](https://incenp.org/notes/2015/gnupg-for-ssh-authentication.html)

- [Gnupg Download (DEB, RPM, TGZ, TXZ, XZ)](https://pkgs.org/download/gnupg) : こうやって見ると最新バージョンに追随しているディストリビューションの方が珍しいのか

- [GnuPG for Windows インストール編]({{< ref "/openpgp/using-gnupg-for-windows-1.md" >}})
- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})
- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Debian]: https://www.debian.org/ "Debian -- The Universal Operating System"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[TOFU]: https://en.wikipedia.org/wiki/Trust_on_first_use "Trust on first use - Wikipedia"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[XDG Base Directory Specification]: https://specifications.freedesktop.org/basedir-spec/latest/

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
