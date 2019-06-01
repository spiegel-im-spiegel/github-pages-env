+++
title = "メール・サービスを立てずにコマンドラインでメールを送信する"
date =  "2019-06-01T23:45:58+09:00"
description = "調べてみたら msmtp なるツールが良さげである。今回は msmtp を構成してメールを送信してみる。"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "mua", "mta" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

cron で回したプロセスの実行結果をメールで送信することを考える。

[Ubuntu] はインストール直後の既定状態ではメール・サービスは入ってない。
したがって cron の実行結果（標準出力等に吐き出されるもの）は何処にも通知されない。
これはこれで合理的な設計なのだが（ cron の実行結果でメールボックスが溢れたりすることもないし），やはり結果をメールで通知したいこともある。
かといって，そのためにわざわざデスクトップ・パソコンにメール・サービスを入れるのはナンセンスであろう。

自前でメール・サービスを立てずに外部のメール・サーバにメールを流す送信専用の簡易 MUA があれば便利である。
そこで調べてみたら [msmtp] なるツールが良さげである[^ssmtp1]。
今回は [msmtp] を構成してコマンドラインでメールを送信してみる。

[^ssmtp1]: ググってみると [sSMTP] に関する記事が頻出したが，残念なことに [sSMTP] はもはやメンテナンスされていないようである。

## [msmtp] のインストール

[msmtp] は APT で導入可能である。

```text
$ apt show msmtp
Package: msmtp
Version: 1.8.3-1
Priority: extra
Section: universe/mail
Origin: Ubuntu
Maintainer: Ubuntu Developers <ubuntu-devel-discuss@lists.ubuntu.com>
Original-Maintainer: Emmanuel Bouthenot <kolter@debian.org>
...
```

なので早速インストールする。

```text
$ sudo apt install msmtp msmtp-mta
```

`msmtp-mta` パッケージは cron から [msmtp] を利用するのに必要なので併せてインストールしておく。
一応，操作確認しておこう。

```text
$ msmtp --version
msmtp version 1.8.3
Platform: x86_64-pc-linux-gnu
TLS/SSL library: GnuTLS
Authentication library: GNU SASL
Supported authentication methods:
plain scram-sha-1 external gssapi cram-md5 digest-md5 login ntlm 
IDN support: enabled
NLS: enabled, LOCALEDIR is /usr/share/locale
Keyring support: none
System configuration file name: /etc/msmtprc
User configuration file name: /home/username/.msmtprc

Copyright (C) 2019 Martin Lambers and others.
This is free software.  You may redistribute copies of it under the terms of
the GNU General Public License <http://www.gnu.org/licenses/gpl.html>.
There is NO WARRANTY, to the extent permitted by law.
```

よしよし。

## [msmtp] の設定

まずは `~/.msmtprc` ファイルを作成して外部のメール・サーバを定義する。
`~/.msmtprc` の雛形ファイルは以下にあるので参考になるだろう。

- `/usr/share/doc/msmtp/examples/msmtprc-user.example`

このファイルを元にして Gmail のメール・サーバにアクセスする設定を記述してみる。
こんな感じ[^passwd1]。

[^passwd1]: `password` の項目にはパスワードを記述するが， Google サインイン用のパスワードではなく「アプリ・パスワード」をセットする（Googleへのサインインに2要素認証を使っていることが前提）。アプリ・パスワードは「[Google アカウント](https://myaccount.google.com/)」のページで設定できる。アプリやプラットフォームごとに異なるアプリ・パスワードを設定するのがコツである。

```text
# Set default values for all following accounts.
defaults

# Use the mail submission port 587 instead of the SMTP port 25.
port 587

# Always use TLS.
tls on

# Gmail service
account gmail

# Host name of the SMTP server
host smtp.gmail.com

# Envelope-from address
from username@gmail.com

# Authentication. The password is given using one of five methods, see below.
auth on
user username@gmail.com

# Password method 2: Store the password in an encrypted file, and tell msmtp
# which command to use to decrypt it. This is usually used with GnuPG, as in
# this example. Usually gpg-agent will ask once for the decryption password.
#passwordeval gpg2 --no-tty -q -d ~/.msmtp-password.gpg

# Password method 3: Store the password directly in this file. Usually it is not
# a good idea to store passwords in plain text files. If you do it anyway, at
# least make sure that this file can only be read by yourself.
password password_string

# Set a default account
account default : gmail
```

`~/.msmtprc` ファイルの権限を変更するのも忘れずに。

```text
$ chmod 0600 ~/.msmtprc
```

これで準備 OK。
試しにコマンドラインでメールを送信してみよう。

```text
$ echo "hello there." | msmtp -a gmail username@gmail.com
```

これで `username@gmail.com` 宛に “hello there.” という内容でメールが届けば成功である。
ちなみに `-a` オプションは省略できる。
この場合 `account default` で指定されたアカウント情報で送信される。

## cron との連携

では cron との連携を試してみよう。

まずはパソコンのユーザ・アカウントとメールアドレスを連携させるために `/etc/aliases` ファイルを作成する。
内容はこんな感じ。

```text
username: username@gmail.com
root: username@gmail.com
default: username@gmail.com
```

さらに `~/.msmtprc` ファイルに以下の記述を追加する。

```text
# aliases file
aliases /etc/aliases
```

これで [msmtp] 側の準備は完了。

テスト用のプロセスを crontab で定義する。

```
$ crontab -e
```

内容は例えばこんな感じ。

```text
*/5 * * * * MAILTO=username ls
```

これで `username` のメールアドレス宛に5分おきに `ls` コマンドの実行結果をメール送信する。
かなりウザいので動作確認できたらソッコーで削除しないと（笑）

うまくいかない場合は `/var/log/syslog` を見てみるとヒントになるかもしれない。

```text
$ cat /var/log/syslog | grep sendmail
```

## パウワード情報の暗号化

外部メール・サーバへ送信する際に認証を行う場合は `~/.msmtprc` ファイルに認証用のパスワードを設定する必要があるが，平文で保存されるため，いかにも不用心である。
そこでパスワード情報を暗号化することを考える。

具体的には `~/.msmtprc` ファイルの `password` 項目を以下の記述で置き換える。

```text
#passwordeval gpg --no-tty -q -d ~/.msmtp-password.gpg
```

`~/.msmtp-password.gpg` が暗号化されたパスワード情報を格納するファイルで [GnuPG] で暗号化されている。
`~/.msmtp-password.gpg` ファイルを作成するには以下のコマンドラインを実行する。

```text
gpg --encrypt -o ~/.msmtp-gmail.gpg -r <user>@gmail.com -
```

最後のハイフン（`-`）を忘れずに。
これで標準入力からパスワードを入力し改行コードを入力した後 `ctrl-d` で処理を抜ければ完了である。

メール送信時には復号のために [GnuPG] がパスフレーズを要求するのでご注意を[^gpg1]。

[^gpg1]: [GnuPG] の鍵を作成する際にパスフレーズの入力を省略すれば復号処理を自動化できるが秘密鍵の中身が丸見えになってしまうので取り扱いには注意が必要である。パスワード暗号化用の専用鍵を作って運用するのが無難だろう。

## ブックマーク

- [msmtp - Debian Wiki](https://wiki.debian.org/msmtp)
- [msmtp - ArchWiki](https://wiki.archlinux.jp/index.php/Msmtp)

- [各ユーザのcrontabファイルの場所について - Qiita](https://qiita.com/iganari/items/5a61ec93d989f7c77a2c)
- [初心者向けcronの使い方 - Qiita](https://qiita.com/tossh/items/e135bd063a50087c3d6a)
- [crontab使い方まとめ。 - ばくのエンジニア日誌](http://bakunyo.hatenablog.com/entry/2013/06/20/crontab%E4%BD%BF%E3%81%84%E6%96%B9%E3%81%BE%E3%81%A8%E3%82%81%E3%80%82)
- [Cronの使い方とテクニックと詰まったところ - Qiita](https://qiita.com/UNILORN/items/a1a3f62409cdb4256219)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[msmtp]: https://marlam.de/msmtp/
[sSMTP]: https://wiki.debian.org/sSMTP "sSMTP - Debian Wiki"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
