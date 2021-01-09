+++
title = "SSH の認証鍵を GunPG で作成・管理する"
date = "2021-01-06T23:06:57+09:00"
description = "新たに鍵を作成する場合や今までの鍵を破棄して作り直す場合などの状況があれば検討してもいいだろう。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "openssh", "management", "cryptography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は [GnuPG] で作成した鍵を [OpenSSH] の認証鍵として使う方法について覚え書きとして記しておく。
現時点で運用できている認証鍵を置き換えるメリットはないが，新たに鍵を作成する場合や今までの鍵を破棄して作り直す場合などの状況があれば検討してもいいだろう。

まず [GnuPG] で作成する鍵は機能別に以下の4種類に分類される（ひとつの鍵で複数の機能を持たせることもできる）。

| 機能   | 略称 |
|:------ |:----:|
| 署名   | `S`  |
| 証明   | `C`  |
| 認証   | `A`  |
| 暗号化 | `E`  |

このうち [OpenSSH] の認証鍵として使えるのは「認証」機能をもった鍵のみである。
認証機能は電子署名用の鍵であれば任意に付与することができるが，専用の副鍵を追加するのがセオリーらしい。

## 認証用の副鍵を追加する

まず以下の [OpenPGP] 鍵があるとする（主鍵には `SC`，副鍵には `E` の機能が付いてる点に注目）。

```text
$ gpg --list-keys alice
pub   ed25519 2021-01-06 [SC] [有効期限: 2021-01-13]
      011C720B03D2E1D6BCFA98391DFF44901121B61D
uid           [  究極  ] Alice <alice@example.com>
sub   cv25519 2021-01-06 [E]
```

この鍵に認証用の副鍵を追加する。
鍵の追加には `--edit` コマンドを使う。
なお `--expert` オプションと一緒に使うと幸せになれる。

```text
$ gpg --expert --edit-key alice
gpg (GnuPG) 2.2.26; Copyright (C) 2020 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

秘密鍵が利用できます。

sec  ed25519/1DFF44901121B61D
     作成: 2021-01-06  有効期限: 2021-01-13  利用法: SC  
     信用: 究極        有効性: 究極
ssb  cv25519/4FECD03BE5BE4454
     作成: 2021-01-06  有効期限: 無期限      利用法: E   
[  究極  ] (1). Alice <alice@example.com>

gpg> 
```

ここで `addkey` コマンドを入力する。

```text {hl_lines=[10]}
gpg> addkey
ご希望の鍵の種類を選択してください:
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (5) Elgamal (暗号化のみ)
   (6) RSA (暗号化のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (12) ECC (暗号化のみ)
  (13) 既存の鍵
  (14) カードに存在する鍵
あなたの選択は? 
```

今回は認証用の鍵の追加なので 7, 8, 11 のいずれかを選択する。
ここは個人的な好みで ECC 鍵を選択しよう。

```text {hl_lines=[4]}
あなたの選択は? 11

鍵ECDSA/EdDSAに認められた操作: Sign Authenticate 
現在の認められた操作: Sign 

   (S) 署名機能を反転する
   (A) 認証機能を反転する
   (Q) 完了

   あなたの選択は? 
```

現在は署名機能（Sign）のみ有効になっているが，欲しいのは認証機能のみなので `S` と `A` を一回づつ入力する。

```text {hl_lines=[4,13]}
あなたの選択は? s

鍵ECDSA/EdDSAに認められた操作: Sign Authenticate 
現在の認められた操作: 

   (S) 署名機能を反転する
   (A) 認証機能を反転する
   (Q) 完了

あなたの選択は? a

鍵ECDSA/EdDSAに認められた操作: Sign Authenticate 
現在の認められた操作: Authenticate 

   (S) 署名機能を反転する
   (A) 認証機能を反転する
   (Q) 完了

あなたの選択は? 
```

これで認証機能（Authenticate）のみ有効になった。
`Q` を入力して次に進もう。

```text {hl_lines=[3]}
あなたの選択は? q
ご希望の楕円曲線を選択してください:
   (1) Curve 25519
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
あなたの選択は? 
```

[OpenSSH] の認証用には `1` から `5` の楕円曲線のいずれかを選択する。
個人的なお勧めは `1` の “`Curve 25519`” である。
理由は以下の記事を参考のこと。

- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})

では `1` を入力して先に進む。

```text {hl_lines=[3]}
あなたの選択は? 1
鍵の有効期限を指定してください。
         0 = 鍵は無期限
      <n>  = 鍵は n 日間で期限切れ
      <n>w = 鍵は n 週間で期限切れ
      <n>m = 鍵は n か月間で期限切れ
      <n>y = 鍵は n 年間で期限切れ
鍵の有効期間は? (0)0
鍵は無期限です
```

有効期限は意味がないので無期限（`0`）を選択する。
理由は後述するのでちょっと待ってね。

最終確認をして鍵を生成する。

```text
これで正しいですか? (y/N) y
本当に作成しますか? (y/N) y         
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

sec  ed25519/1DFF44901121B61D
     作成: 2021-01-06  有効期限: 2021-01-13  利用法: SC  
     信用: 究極        有効性: 究極
ssb  cv25519/4FECD03BE5BE4454
     作成: 2021-01-06  有効期限: 無期限      利用法: E   
ssb  ed25519/230D446E390C3E49
     作成: 2021-01-06  有効期限: 無期限      利用法: A   
[  究極  ] (1). Alice <alice@example.com>

gpg> save
```

最後は `save` コマンドを入力して結果を鍵束に保存する。
これで

```text {hl_lines=[6]}
$ gpg --list-keys alice
pub   ed25519 2021-01-06 [SC] [有効期限: 2021-01-13]
      011C720B03D2E1D6BCFA98391DFF44901121B61D
uid           [  究極  ] Alice <alice@example.com>
sub   cv25519 2021-01-06 [E]
sub   ed25519 2021-01-06 [A]
```

認証用の鍵が追加できた。

## [OpenSSH] フォーマットの公開鍵を出力する。

[OpenSSH] フォーマットの公開鍵は `--export-ssh-key` コマンドで出力できる。

```text
$ gpg --export-ssh-key alice
ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIFfjejx/Saej929myfZoBQAKgusPi2iiOxdZZfpCLxh5 openpgp:0x390C3E49
```

このテキストをホスト機の `~/.ssh/authorized_keys` ファイルに追記すれば仕込みは完了である。
パーミッションの変更を忘れずに（笑）

この操作で分かると思うが [OpenSSH] フォーマットで出力する時点で [OpenPGP] 鍵の情報はほぼ脱落している。
だから「有効期限は意味がない」のよ。

ちなみにこの操作は公開鍵に対して行われる。
たとえば，認証用の鍵を付加した [OpenPGP] 公開鍵をサーバ管理者に渡せば，サーバ管理者は集めた [OpenPGP] 公開鍵に署名して完全性を確保した後， [OpenSSH] 認証用公開鍵を抽出して各ユーザのディレクトリにまとめてセットする，といったこともできるだろう。


## ローカル側の設定 【2020-01-09 変更・追記あり】

念のためローカル側の設定についても記しておく。

### ssh-agent を gpg-agent に置き換える

[OpenSSH] では `ssh-agent` を [GnuPG] の `gpg-agent` に置き換えることで鍵の管理を [GnuPG] 側に委譲できる。
`gpg-agent` は [GnuPG] 秘密鍵管理の中核コンポーネントで，自身は Pinentry で入力したパスフレーズを一定期間キャッシュすることでユーザの鍵操作を省力化できる（秘密鍵自体のキャッシュは行わない）。

この記事では Ubuntu での設定方法を挙げておく[^win1]。

[^win1]: Windows 環境での設定手順については拙文「[GnuPG for Windows : gpg-agent について]({{< relref "./using-gnupg-for-windows-2.md" >}})」を参考にどうぞ。

どうも (Ubuntu を含む) Debian 系のディストリビューションでは SSH 接続に `ssh-agent` を使うよう構成されているようで， `gpg-agent` に置き換えるためにはいくつか設定を換える必要があるようだ。

まず `/etc/X11/Xsession.options` ファイル。

```text {hl_line=[8]}
# $Id: Xsession.options 189 2005-06-11 00:04:27Z branden $
#
# configuration options for /etc/X11/Xsession
# See Xsession.options(5) for an explanation of the available options.
allow-failsafe
allow-user-resources
allow-user-xsession
use-ssh-agent
use-session-dbus
```

この中の `use-ssh-agent` を `no-use-ssh-agent` に置き換える。

次に `/etc/xdg/autostart/gnome-keyring-ssh.desktop` ファイルを `~/.config/autostart/` ディレクトリにコピーし

```text
Hidden=true
```

を書き加える。

最後に `~/.gnupg/gpg-agent.conf` ファイルに `enable-ssh-support` を書き加える。

```text
enable-ssh-support
default-cache-ttl-ssh 1800
max-cache-ttl-ssh 7200
```

下2つのオプションは任意で，以下の意味を持つ。

| オプション名            | 内容 |
|-------------------------|------|
| `default-cache-ttl-ssh` | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する。 既定値は 1800 |
| `max-cache-ttl-ssh`     | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる。 既定値は 7200 |

有効期間は大きすぎると漏洩リスクが高まるのでほどほどに（笑）

### [OpenSSH] 認証鍵の登録

[GnuPG] の鍵束の鍵を [OpenSSH] の認証鍵として使うには `~/.gnupg/sshcontrol` ファイルへの登録が必要である。
先ほど作成した鍵であれば，まず以下のコマンドで

```text {hl_lines=[9]}
$ gpg --list-keys --with-keygrip alice
pub   ed25519 2021-01-06 [SC] [有効期限: 2021-01-13]
      011C720B03D2E1D6BCFA98391DFF44901121B61D
      Keygrip = 97249ABEB2A2FD9E88F6723BB19D4F84B90E261A
uid           [  究極  ] Alice <alice@example.com>
sub   cv25519 2021-01-06 [E]
      Keygrip = 96CB831965E1A7EB4705577D6A7CB7F9E05C8192
sub   ed25519 2021-01-06 [A]
      Keygrip = F5C774B5B418B6E0B5B7942F93DE82BF2FEF4C8E
```

該当する鍵の keygrip 値を調べる。
今回の例なら “`F5C774B5B418B6E0B5B7942F93DE82BF2FEF4C8E`” が該当する keygrip 値である。
これを `~/.gnupg/sshcontrol` ファイルに追記する。

```text
$ echo F5C774B5B418B6E0B5B7942F93DE82BF2FEF4C8E 0 >> ~/.gnupg/sshcontrol
```

これで `ssh-add -L` コマンドでこの鍵の内容が表示されればOK。

ちなみに keygrip 値の後ろの `0` はキャッシュ期間（秒）を指すらしい。
`0` より大きければ `gpg-agent.conf` ファイルの指定より優先されるってことかな。
また行頭に `!` マークを付けると鍵の使用を無効化できる。

念のため `sshcontrol` ファイルも不用意に書き込みできないよう制限をかけておくとよいだろう。

## ブックマーク

- [GnuPG チートシート（鍵作成から失効まで）]({{< relref "./gnupg-cheat-sheet.md" >}})
- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< relref "./using-ecc-with-gnupg.md" >}})
- [OpenSSH 鍵をアップグレードする（さようなら SHA-1）]({{< ref "/remark/2020/06/upgrade-openssh-key.md" >}})

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[git]: https://git-scm.com/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
