+++
date = "2016-03-11T04:51:33+09:00"
description = "description"
draft = true
tags = ["security", "cryptography", "openpgp", "gnupg", "tools", "ssh", "putty"]
title = "GnuPG Modern Version for Windows ― gpg-agent について"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [インストール編]({{< relref "remark/2016/03/using-gnupg-modern-version-1.md" >}})
2. [gpg-agent について]({{< relref "remark/2016/03/using-gnupg-modern-version-2.md" >}}) （← イマココ）

## gpg-agent

`gpg-agent` は [GnuPG] modern version の中核コンポーネントで，秘密鍵の管理を行い，一定期間キャッシュする。
`gpg-agent` は `gpg`, `gpgsm`, `gpgconf`, `gpg-connect-agent` といったコンポーネントから常駐プロセスとして起動されお互いに通信を行う[^od]。

[^od]: Modern version では `gpg-agent` が必須となった。したがって `--use-agent`, `--no-use-agent`, `--gpg-agent-info` 各オプションは無効（ダミーオプション）になっている。また UNIX 互換プラットフォームで `gpg-agent` 利用する際は `GPG_TTY` 環境変数をセットする必要があるが， Windows では不要なためここでは割愛する。

`gpg-agent` が稼働中かどうかは `gpg-agent` を引数なしで起動すれば分かる。
以下は既に起動している場合。

```text
C:>gpg-agent
gpg-agent[3996]: gpg-agent running and available
```

`gpg-agent` が稼働していない場合は

```text
C:>gpg-agent
gpg-agent[9552]: no gpg-agent running in this session
```

などと表示される。

手動で `gpg-agent` を起動する場合は以下のコマンドで起動する。

```text
C:>gpg-connect-agent /bye
gpg-connect-agent: no running gpg-agent - starting 'C:\path\to\GnuPG\bin\gpg-agent.exe'
gpg-connect-agent: waiting for the agent to come up ... (5s)
gpg-connect-agent: connection to agent established
```

逆に `gpg-agent` を手動で停止したい場合は

```text
C:>gpg-connect-agent killagent /bye
OK closing connection
```

とすれば安全に停止できる。

### Pinentry

Pinentry はパスフレーズ[^pp] やスマートカードの PIN コードを入力する際に `gpg-agent` から呼び出される。
Windows 版では Pinentry プログラムとして `pinentry-basic.exe` が同梱されている。
同等の機能を持つものであれば他のプログラムと差し替えることもできる。

[^pp]: パスワード（password）とパスフレーズ（passphrase）の違いは，パスフレーズでは英数字以外に空白文字や記号が使え文字数の制限がないことにある。ちなみに [OpenPGP] の秘密鍵にはいかなる形でもパスフレーズを保持しない（S2K パラメータ情報は持っている）。

`gpg` に `--batch` （または `--pinentry-mode loopback`）オプションとパスフレーズ指定オプション（`--passphrase`, `--passphrase-fd`, `--passphrase-file`）をセットで指定している場合は Pinentry を迂回できることがある（`--quick-gen-key` コマンドの場合など）。

### gpg-agent のオプション

`gpg-agent` のオプションは [GnuPG] ホームディレクトリ[^hd] 直下にある `gpg-agent.conf` ファイルで設定する。
設定は以下の様なフォーマットで行う。

[^hd]: Windows では， [GnuPG] ホームディレクトリの既定は `%APPDATA%\gnupg` となっている。これを変更するには `GNUPGHOME` 環境変数または `--homedir` オプションを使う。（[前回]を参照のこと）

```text
default-cache-ttl 600
max-cache-ttl 7200
```

`gpg-agent.conf` ファイルで使いそうなオプションを以下に挙げる。

| オプション名        | 内容 |
|:--------------------|:-----|
| `log-file`          | ログの出力先をフルパスで指定する。 挙動をチェックしたい場合など |
| `default-cache-ttl` | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する。 既定値は 600 |
| `max-cache-ttl`     | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる。 既定値は 7200 |
| `pinentry-program`  | 独自に Pinentry プログラムを指定する場合はここにフルパスで指定する |
| `pinentry-timeout`  | Pinentry プログラムの表示時間を秒単位で指定する。 既定値は 0 （タイムアウトなし） |

他にも，鍵生成時にパスフレーズの文字種や最小文字長を指定したり，パスフレーズの有効期間（期間が過ぎると警告が出るらしい）を設定できたりするようだ。
オプション項目について詳しくはは[マニュアル](https://gnupg.org/documentation/manuals/gnupg/Agent-Options.html "Agent Options - Using the GNU Privacy Guard")（英語）を参照のこと。

## PuTTY with gpg-agent

[PuTTY] は Windows 用の SSH クライアント兼ターミナル・エミュレータである。
[PuTTY] には Plink と呼ばれるコマンドラインベースの SSH 接続ツールがあり，他ツール（例えば [Git for Windows]）と連携できるようになっている。
さらに [PuTTY] には Pagent と呼ばれるエージェントツールもあり，認証用の秘密鍵をキャッシュすることができる。

今回は Pagent を `gpg-agent` で置き換えることを考える。

### gpg-agent のオプション

`gpg-agent.conf` ファイルに以下のオプションを追加する（`enable-putty-support` 以外は任意）。

| オプション名            | 内容 |
|:------------------------|:-----|
| `enable-putty-support`  | Pagent プロトコルを有効にする |
| `default-cache-ttl-ssh` | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する。 既定値は 1800 |
| `max-cache-ttl-ssh`     | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる。 既定値は 7200 |

設定を保存したら `gpg-connect-agent` を使って `gpg-agent` を再起動する。

```text
C:>gpg-connect-agent killagent /bye
OK closing connection

C:>gpg-connect-agent /bye
gpg-connect-agent: no running gpg-agent - starting 'C:\path\to\GnuPG\bin\gpg-agent.exe'
gpg-connect-agent: waiting for the agent to come up ... (5s)
gpg-connect-agent: connection to agent established
```

なお， [GnuPG] の各コンポーネントは必要に応じて自動的に `gpg-agent` を起動するので問題ないのだが， [PuTTY] は [GnuPG] と連動しているわけではないため， [PuTTY] 起動時に `gpg-agent` が起動していない状況もありうる。
そこで， Windows ログイン時に `gpg-connect-agent` を使って `gpg-agent` を起動しておくことをお薦めする。

### SSH 鍵のインポート

SSH 鍵のインポートには2通りの方法あるようだが，今回は簡単な方でいく[^imp]。

[^imp]: 今回は PPK ファイルを直接読み込む方法をとったが， PPK ファイルから OpenSSH 形式にエクスポートし，それを更に X.509 形式に変換した後 `gpgsm` でインポートすることもできる。（参考： [Convert keys between GnuPG, OpenSsh and OpenSSL](http://www.sysmic.org/dotclear/index.php?post/2010/03/24/Convert-keys-betweens-GnuPG%2C-OpenSsh-and-OpenSSL)）

鍵ファイル（ここでは `id_rsa.PPK` とする）を Pagent で開く。
ファイルの関連付けがされている場合はエクスプローラから該当の PPK ファイルをダブルクリックすればいい。
そうでない場合は以下のコマンドで PPK ファイルを開く。

```
C:>pageant.exe id_rsa.PPK
```

すると Pagent のプロンプトが1回， `gpg-agent` のプロンプトが2回表示され，都合3回パスフレーズを入力させられる。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1454/25558118892_045f0a9b8f_o.png" title="Pagent" link="https://www.flickr.com/photos/spiegel/25558118892/" >}}

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1485/25558116832_dd02c5d7ec_o.png" title="GnuPG Pinentry (1)" link="https://www.flickr.com/photos/spiegel/25558116832/" >}}

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1467/25376004580_58a790993c_o.png" title="GnuPG Pinentry (2)" link="https://www.flickr.com/photos/spiegel/25376004580/" >}}

これで秘密鍵が `private-keys-v1.d` フォルダに格納される。
また `sshcontrol` ファイルが作成され，インポートした鍵の情報が書き込まれる。

```text
# List of allowed ssh keys.  Only keys present in this file are used
# in the SSH protocol.  The ssh-add tool may add new entries to this
# file to enable them; you may also add them manually.  Comment
# lines, like this one, as well as empty lines are ignored.  Lines do
# have a certain length limit but this is not serious limitation as
# the format of the entries is fixed and checked by gpg-agent. A
# non-comment line starts with optional white spaces, followed by the
# keygrip of the key given as 40 hex digits, optionally followed by a
# caching TTL in seconds, and another optional field for arbitrary
# flags.   Prepend the keygrip with an '!' mark to disable it.

# RSA key added on: 2016-03-10 21:24:32
# MD5 Fingerprint:  56:ff:fd:60:38:a1:7a:44:0c:37:86:90:94:8d:7f:6a
F65BB98767E88930612C6EABC4D4918E2A573903 0
```

この `F65B...` の長ったらしい数字列は keygrip と呼ばれる鍵の識別子で [OpenPGP] の鍵 ID とは異なるもののようだ[^kg]。

[^kg]: [OpenPGP] 鍵以外の鍵にも対応するためらしい。 [OpenPGP] 鍵の keygrip は `--with-keygrip` オプションを付けて鍵を表示すると見ることができる。

これで鍵のインポートができたので [PuTTY] で実際に SSH 接続してみると

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1586/25585468551_0447584a65_o.png" title="GnuPG Pinentry (3)" link="https://www.flickr.com/photos/spiegel/25585468551/" >}}

とプロンプトが表示された。
めでたし！

## Git for Windows との連携

[Git for Windows] と [PuTTY] を連携するには，環境変数 `GIT_SSH` に Plink へのフルパスをセットする。

```text
setx GIT_SSH=C:\path\to\PuTTY\plink.exe
```

一方，リポジトリの `.git\config` ファイルには PPK ファイルの場所をセットする。

```text
[remote "origin"]
	puttykeyfile = C:/path/to/PuTTY/id_rsa.PPK
```

この状態で `git fetch` または `git push` を行うと Plink 経由で `gpg-agent` にリクエストが発生する。

## Windows 版 gpg-agent は OpenSSH と相性が悪い？

`gpg-agent` は [OpenSSH] の `ssh-agent` と置き換えることもできる[^sa]。
`gpg-agent` への SSH 鍵のインポートには `ssh-add` を使うのだが， Windows 環境では上手く動かない。
どうやらファイル・ディスクリプタ `S.gpg-agent.ssh` が上手く機能しないようだ。

[MSYS2] 版[^m] と [PowerShell 用](https://github.com/PowerShell/Win32-OpenSSH "PowerShell/Win32-OpenSSH: Win32 port of OpenSSH")の [OpenSSH] で試してみたのだが，いずれも上手くいかなかった。
[MSYS2] 版については [MSYS2] の [GnuPG] modern version を使えば上手くいくのかもしれないが，面倒なので試してない。
今後，試す機会があればここに追記する。

[^sa]: `gpg-agent.conf` ファイルに `enable-ssh-support` オプションをセットする。
[^m]: [Git for Windows] に同梱されている `git bash` も [MSYS2] である。

## 参考になる（かもしれない） Web ページ

- [The GNU Privacy Guard](https://gnupg.org/)
- [わかる！ OpenPGP 暗号 — Baldanders.info](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)
- [Windowsでのssh agent - Qiita](http://qiita.com/tsuyoshi_cho/items/79c09905ae3f192b3a0f)
- [SSH authentication using a YubiKey on Windows](https://developers.yubico.com/PGP/SSH_authentication/Windows.html)
- [Git/Git for Windows/SSHにPuTTYを使う - yanor.net/wiki](http://yanor.net/wiki/?Git%2FGit%20for%20Windows%2FSSH%E3%81%ABPuTTY%E3%82%92%E4%BD%BF%E3%81%86)
- [GPG Agent under Windows as SSH Agent for git bash - Super User](https://superuser.com/questions/911496/gpg-agent-under-windows-as-ssh-agent-for-git-bash)
- [Using GnuPG (2.1) for SSH authentication](http://incenp.org/notes/2015/gnupg-for-ssh-authentication.html)
- [Convert keys between GnuPG, OpenSsh and OpenSSL - Sysmic.org](http://www.sysmic.org/dotclear/index.php?post/2010/03/24/Convert-keys-betweens-GnuPG%2C-OpenSsh-and-OpenSSL)

[前回]: {{< relref "remark/2016/03/using-gnupg-modern-version-1.md" >}} "GnuPG Modern Version for Windows ― インストール編"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[Git for Windows]: https://git-for-windows.github.io/ "Git for Windows"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
