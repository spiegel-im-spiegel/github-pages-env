+++
title = "gpg-agent の設定： GnuPG in Ubuntu"
date = "2021-01-10T11:53:13+09:00"
description = "Zenn で書いた「Ubuntu で OpenSSH の鍵管理を gpg-agent に委譲する」をこのブログの「最新版 GnuPG をビルドする」の続編として再構成した"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "gnupg", "ubuntu", "openssh", "cryptography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

この記事は Zenn で書いた「[Ubuntu で OpenSSH の鍵管理を gpg-agent に委譲する](https://zenn.dev/spiegel/articles/20210109-gpg-agent)」をこのブログの「[最新版 GnuPG をビルドする]({{< relref "./build-gnupg-in-ubuntu.md" >}})」の続編として再構成したものである。

なお Windows 版 `gpg-agent` の設定については，古い記事で恐縮だが

- [GnuPG for Windows : gpg-agent について]({{< relref "./using-gnupg-for-windows-2.md" >}})

が参考になる。

## ビルドした gpg-agent をサービスとして登録する

[前回]ビルドした [GnuPG] を使って処理を行おうとすると `gpg-agent` が古い旨の警告が出る。
`ps` コマンドでチェックすると，どうやらオリジナルの `/usr/bin/gpg-agent` が稼働しているようだ。

稼働している `/usr/bin/gpg-agent` を一旦落として `gpg-connect-agent` コマンドで再起動すれば `/usr/local/bin/gpg-agent` が起動するのだが，ログインし直すと元に戻ってしまう。

いろいろググってみたのだが， `gpg-agent` の制御は `/usr/lib/systemd/user/gpg-agent.service` ファイルでサービスとして行っているようだ。
内容はこんな感じ。

```ini {hl_lines=[7,8]}
[Unit]
Description=GnuPG cryptographic agent and passphrase cache
Documentation=man:gpg-agent(1)
Requires=gpg-agent.socket

[Service]
ExecStart=/usr/bin/gpg-agent --supervised
ExecReload=/usr/bin/gpgconf --reload gpg-agent
```

そこで，最後の2行を

```ini
ExecStart=/usr/local/bin/gpg-agent --supervised
ExecReload=/usr/local/bin/gpgconf --reload gpg-agent
```

と変更し `gpg-agent.service` ファイルをリロードする。

```text
$ systemctl --user daemon-reload
```

その後サービスを `restart` したら `/usr/local/bin/gpg-agent` のほうで稼働してくれるようになった。

```text
$ systemctl --user restart gpg-agent
$ systemctl --user status gpg-agent
● gpg-agent.service - GnuPG cryptographic agent and passphrase cache
     Loaded: loaded (/usr/lib/systemd/user/gpg-agent.service; static)
     Active: active (running) since Sat 2021-01-09 09:38:33 JST; 9min ago
TriggeredBy: ● gpg-agent.socket
             ● gpg-agent-ssh.socket
             ● gpg-agent-extra.socket
             ● gpg-agent-browser.socket
       Docs: man:gpg-agent(1)
   Main PID: 18913 (gpg-agent)
     CGroup: /user.slice/user-1000.slice/user@1000.service/gpg-agent.service
             ├─18913 /usr/local/bin/gpg-agent --supervised
             └─19398 scdaemon --multi-server

 Jan 09 09:38:33 mymachine systemd[1616]: Started GnuPG cryptographic agent and passphrase cache.
 Jan 09 09:38:33 mymachine gpg-agent[18913]: gpg-agent (GnuPG) 2.2.26 starting in supervised mode.
 Jan 09 09:38:33 mymachine gpg-agent[18913]: using fd 3 for std socket (/run/user/1000/gnupg/S.gpg-agent)
 Jan 09 09:38:33 mymachine gpg-agent[18913]: using fd 4 for ssh socket (/run/user/1000/gnupg/S.gpg-agent.ssh)
 Jan 09 09:38:33 mymachine gpg-agent[18913]: using fd 5 for extra socket (/run/user/1000/gnupg/S.gpg-agent.extra)
 Jan 09 09:38:33 mymachine gpg-agent[18913]: using fd 6 for browser socket (/run/user/1000/gnupg/S.gpg-agent.browser)
 Jan 09 09:38:33 mymachine gpg-agent[18913]: listening on: std=3 extra=5 browser=6 ssh=4
...
```

よーし，うむうむ，よーし。

### dirmgr も同様に

OpenPGP 鍵サーバを含むディレクトリ・サービスとやり取りする `dirmngr` も `/usr/lib/systemd/user/dirmngr.service` ファイルでサービス化されている。

```ini {hl_lines=[7,8]}
[Unit]
Description=GnuPG network certificate management daemon
Documentation=man:dirmngr(8)
Requires=dirmngr.socket

[Service]
ExecStart=/usr/bin/dirmngr --supervised --honor-http-proxy
ExecReload=/usr/bin/gpgconf --reload dirmngr
```

これも最後の2行を

```ini
ExecStart=/usr/local/bin/dirmngr --supervised --honor-http-proxy
ExecReload=/usr/local/bin/gpgconf --reload dirmngr
```

と変更する。
後の操作は同じ。

```text
$ systemctl --user daemon-reload
$ systemctl --user restart dirmngr
$ systemctl --user status dirmngr
● dirmngr.service - GnuPG network certificate management daemon
     Loaded: loaded (/usr/lib/systemd/user/dirmngr.service; static)
     Active: active (running) since Sat 2021-01-09 09:38:33 JST; 9min ago
TriggeredBy: ● dirmngr.socket
       Docs: man:dirmngr(8)
   Main PID: 349126 (dirmngr)
     CGroup: /user.slice/user-1000.slice/user@1000.service/dirmngr.service
             └─349126 /usr/local/bin/dirmngr --supervised --honor-http-proxy

 Jan 09 09:38:33 mymachine systemd[2209]: dirmngr.service: Succeeded.
 Jan 09 09:38:33 mymachine dirmngr[348581]: dirmngr (GnuPG) 2.2.20 stopped
 Jan 09 09:38:33 mymachine systemd[2209]: Stopped GnuPG network certificate management daemon.
 Jan 09 09:38:33 mymachine systemd[2209]: Started GnuPG network certificate management daemon.
...
```

## [OpenSSH] の鍵管理を gpg-agent に委譲する

上述の話を調べてて初めて知ったのだが，どうも Ubuntu を含む Debian 系のディストリビューションでは [OpenSSH] の認証鍵を `ssh-agent` 固定で管理するよう構成されていて， `gpg-agent` と設定が競合してしまうらしい。
[OpenSSH] の鍵管理を gpg-agent に委譲するよう構成するには `gpg-agent.conf` ファイル以外にもいくつか設定が必要なようだ。

### Xsession.options の変更

まずは `/etc/X11/Xsession.options` ファイルの内容を変更する。
元々の内容は以下の通り。

```bash {hl_lines=[8]}
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

この中の `use-ssh-agent` の記述を `no-use-ssh-agent` に差し替える。
当然ながら変更には管理者権限が必要なのでご注意を。
バックアップを取りながら作業すること。

### autostart/gnome-keyring-ssh.desktop の変更

次に `/etc/xdg/autostart/gnome-keyring-ssh.desktop` の内容を変更するのだが，このファイルを直接変更するのではなく，いったん `~/.config/autostart/` ディレクトリにコピーしてから，コピーしたファイルに対して変更をかける（`autostart` ディレクトリがない場合は先に作成する）。

```
$ cp /etc/xdg/autostart/gnome-keyring-ssh.desktop ~/.config/autostart/
```

これでユーザ単位で設定を弄ることができる。
といっても最終行に

```ini
Hidden=true
```

を追記するだけだが。

{{< div-box type="markdown" >}}
### 【2021-06-05 追記】 Ubuntu 21.04 の場合

2021年4月にリリースされた Ubuntu 21.04 ではこの設定ではうまく行かないようだ（Wayland の問題？）。
この場合，とりあえず応急措置として `.bashrc` ファイルで

```bash
export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
dbus-update-activation-environment --systemd SSH_AUTH_SOCK
```

と言った感じに環境変数 `SSH_AUTH_SOCK` を直接指定する（情報募集）。
{{< /div-box >}}

### gpg-agent.conf の設定

最後に `~/.gnupg/gpg-agent.conf` ファイルに以下の内容を書き込む。

```text
enable-ssh-support
default-cache-ttl-ssh 1800
max-cache-ttl-ssh 7200
```

`gpg-agent.conf` ファイルがない場合は作成すること。
下2行のオプションは任意で，以下の意味を持つ。

| オプション名            | 内容 |
|-------------------------|------|
| `default-cache-ttl-ssh` | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する。 既定値は 1800 |
| `max-cache-ttl-ssh`     | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる。 既定値は 7200 |

有効期間は大きすぎると漏洩リスクが高まるのでほどほどに（笑）

これで設定は完了。念のためログインし直しておこう。

### 環境変数の確認

ログインし直したら環境変数を確認しておく。

```
$ env | grep SSH
SSH_AUTH_SOCK=/run/user/1000/gnupg/S.gpg-agent.ssh
```

てな感じに `SSH_AUTH_SOCK` 環境変数の値が `gpg-agent` のソケットになっていれば OK である。
これで `~/.bashrc` とかに要らん記述をしなくてもよくなった。
ブラボー！

## [GnuPG] による鍵管理

### 既存の [OpenSSH] 認証鍵を [GnuPG] に登録する

上述の設定が完了していれば，既存の [OpenSSH] 認証鍵を [GnuPG] の鍵束に登録するのは `ssh-add` コマンドで簡単にできる。

```
$ ssh-add ./id_ecdsa
Enter passphrase for ./id_ecdsa: 
Identity added: ./id_ecdsa (alice@example.com)
```

この時 `ssh-add` コマンドによるパスフレーズ入力とは別に [GnuPG] の Pinentry によるパスフレーズの設定が行われるので注意。
確認を含めて2箇所入力する必要がある。

{{< fig-img src="/remark/2020/06/upgrade-openssh-key/pinentry.png" link="/remark/2020/06/upgrade-openssh-key/pinentry.png" title="pinentry" >}}

[GnuPG] の鍵束に登録される認証鍵はこのパスフレーズで保護される。
登録した秘密鍵は `~/.gnupg/private-keys-v1.d/` ディレクトリ）に入る。
また `~/.gnupg/sshcontrol` ファイルに

```bash
# ECDSA key added on: 2020-06-01 14:05:35
# Fingerprints:  MD5:e4:5b:66:a6:03:9a:a4:0e:f2:1b:a5:04:72:93:f3:f0
#                SHA256:DtXgQm9rz7Dc5M5yWu/CNVo341o1rcfN9UCyYu+SZU4
A5353D587000D820669B0BD55A0B4AD6897458DB 0
```

という感じに追加した鍵の情報が入る。

ちなみに `A5353D587000D820669B0BD55A0B4AD6897458DB` は keygrip と呼ばれる値で，鍵の種類に関係なく統一的に表される ID である。
また `~/.gnupg/private-keys-v1.d/` ディレクトリに入っている鍵は `A5353D587000D820669B0BD55A0B4AD6897458DB.key` のように keygrip に紐付いたファイル名で格納されている。

さらに，末尾の `0` はキャッシュ期間（秒）を指すらしい。 
`0` より大きければ `gpg-agent.conf` ファイルの指定より優先されるってことかな。

さらにさらに，行頭に `!` マークを付けると鍵の使用を無効化できる。

### [GnuPG] 鍵を [OpenSSH] 認証鍵として設定する

[GnuPG] 鍵を [OpenSSH] 認証鍵として設定することもできる。
ただし専用の認証鍵を作る必要がある。
詳しくは

- [SSH の認証鍵を GunPG で作成・管理する]({{< relref "./ssh-key-management-with-gnupg.md" >}})

を参照のこと。
最終的には全部 [GnuPG] で管理するのがいいんだろうな。

## ブックマーク

- [Configuring gpg-agent for SSH Authentication on Ubuntu](https://curiouslynerdy.com/gpg-agent-for-ssh-on-ubuntu/)
- [Kernel Maintainer PGP guide — The Linux Kernel  documentation](https://www.kernel.org/doc/html/v5.8/process/maintainer-pgp-guide.html)
- [OpenPGP SSH access with Yubikey and GnuPG · GitHub](https://gist.github.com/artizirk/d09ce3570021b0f65469cb450bee5e29)

- [GnuPG チートシート（鍵作成から失効まで）]({{< relref "./gnupg-cheat-sheet.md" >}})
- [OpenSSH 鍵をアップグレードする（さようなら SHA-1）]({{< ref "/remark/2020/06/upgrade-openssh-key.md" >}})

[前回]: {{< relref "./build-gnupg-in-ubuntu.md" >}} "最新版 GnuPG をビルドする"
[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
