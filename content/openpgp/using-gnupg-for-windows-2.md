+++
date = "2017-12-01T17:48:43+09:00"
update = "2017-12-02T21:55:18+09:00"
title = "GnuPG for Windows : gpg-agent について"
description = "今回は gpg-agent について解説する。"
draft = false
tags = ["security", "cryptography", "openpgp", "gnupg", "tools", "ssh", "putty"]
image = "/images/attention/openpgp.png"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"

[scripts]
  mathjax = false
  mermaidjs = true
+++

## gpg-agent

`gpg-agent` は [GnuPG] の中核コンポーネントで，秘密鍵の管理[^sr] を行い一定期間キャッシュする。
`gpg-agent` は `gpg`, `gpgsm`, `gpgconf`, `gpg-connect-agent` といったコンポーネントから常駐プロセスとして起動されお互いに通信を行う[^od]。

[^sr]: [前回]も書いたが， classic version と現行バージョンでは鍵（特に秘密鍵）の管理の仕方が異なるため両者を混在させる場合は注意が必要である。 Classic version で作成した鍵を現行バージョンにも反映させたいのであれば `gpg-v21-migrated` ファイルを削除すると再度移行処理が走るらしい。 Classic version を使わなければならない状況（Linux などではパッケージ管理ツールがアプリケーションの証明用に [GnuPG] の classic version を使うことがある）でないのなら現行バージョンに一本化するほうがお勧めである。
[^od]: 現行バージョンでは `gpg-agent` が必須である。したがって，かつての `--use-agent`, `--no-use-agent`, `--gpg-agent-info` 各オプションは無効（ダミーオプション）になっている。また UNIX 互換プラットフォームで `gpg-agent` 利用する際は `GPG_TTY` 環境変数をセットする必要があるが， Windows では不要なためここでは割愛する。

`gpg-agent` が稼働中かどうかは `gpg-agent` を引数なしで起動すれば分かる。
以下は既に起動している場合。

```text
$ gpg-agent
gpg-agent[3996]: gpg-agent running and available
```

`gpg-agent` が稼働していない場合は

```text
$ gpg-agent
gpg-agent[9552]: no gpg-agent running in this session
```

などと表示される。

手動で `gpg-agent` を起動する場合は以下のコマンドで起動する。

```text
$ gpg-connect-agent /bye
gpg-connect-agent: no running gpg-agent - starting 'C:\path\to\GnuPG\bin\gpg-agent.exe'
gpg-connect-agent: waiting for the agent to come up ... (5s)
gpg-connect-agent: connection to agent established
```

逆に `gpg-agent` を手動で停止したい場合は

```text
$ gpg-connect-agent killagent /bye
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
さらに [PuTTY] には Pageant と呼ばれるエージェントツールもあり，認証用の秘密鍵をキャッシュすることができる。

今回は Pageant を `gpg-agent` で置き換えることを考える。

### gpg-agent のオプション

`gpg-agent.conf` ファイルに以下のオプションを追加する（`enable-putty-support` 以外は任意）。

| オプション名            | 内容 |
|:------------------------|:-----|
| `enable-putty-support`  | Pageant プロトコルを有効にする |
| `default-cache-ttl-ssh` | 直前にアクセスしたキャッシュ・エントリの有効期間を秒単位で指定する。 既定値は 1800 |
| `max-cache-ttl-ssh`     | キャッシュ・エントリの有効期間の最大値を秒単位で指定する。 アクセスの有無にかかわらずこの期間が過ぎるとキャッシュがクリアされる。 既定値は 7200 |

設定を保存したら `gpg-connect-agent` を使って `gpg-agent` を再起動する。

```text
$ gpg-connect-agent killagent /bye
OK closing connection

$ gpg-connect-agent /bye
gpg-connect-agent: no running gpg-agent - starting 'C:\path\to\GnuPG\bin\gpg-agent.exe'
gpg-connect-agent: waiting for the agent to come up ... (5s)
gpg-connect-agent: connection to agent established
```

なお， [GnuPG] の各コンポーネントは必要に応じて自動的に `gpg-agent` を起動するので問題ないのだが， [PuTTY] は [GnuPG] と連動しているわけではないため， [PuTTY] 起動時に `gpg-agent` が起動していない状況もありうる。
そこで， Windows ログイン時に `gpg-connect-agent` を使って `gpg-agent` を起動しておくことをお薦めする。

### SSH 鍵のインポート

SSH 鍵のインポートには2通りの方法あるようだが，今回は簡単な方でいく[^imp]。

[^imp]: 今回は PPK ファイルを直接読み込む方法をとったが， PPK ファイルから OpenSSH 形式にエクスポートし，それを更に X.509 形式に変換した後 `gpgsm` でインポートすることもできるらしい。（参考： [Convert keys between GnuPG, OpenSsh and OpenSSL](http://www.sysmic.org/dotclear/index.php?post/2010/03/24/Convert-keys-betweens-GnuPG%2C-OpenSsh-and-OpenSSL)）

鍵ファイル（ここでは `id_rsa.PPK` とする）を Pageant で開く。
ファイルの関連付けがされている場合はエクスプローラから該当の PPK ファイルをダブルクリックすればいい。
そうでない場合は以下のコマンドで PPK ファイルを開く。

```
$ pageant.exe id_rsa.PPK
```

すると Pageant のプロンプトが1回， `gpg-agent` のプロンプトが2回表示され，都合3回パスフレーズを入力させられる。

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1454/25558118892_045f0a9b8f_o.png" title="Pageant" link="https://www.flickr.com/photos/spiegel/25558118892/" >}}

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

[^kg]: [OpenPGP] 鍵以外の鍵にも対応するためらしい。 [OpenPGP] 鍵の keygrip は `--with-keygrip` オプションを付けて鍵を表示すると見ることができる。ちなみに `private-keys-v1.d` フォルダにある秘密鍵のファイルは，この keygrip 値がそのままファイル名になっている。

これで鍵のインポートができたので [PuTTY] で実際に SSH 接続してみると

{{< fig-img flickr="true" src="https://farm2.staticflickr.com/1586/25585468551_0447584a65_o.png" title="GnuPG Pinentry (3)" link="https://www.flickr.com/photos/spiegel/25585468551/" >}}

とプロンプトが表示された。
めでたし！

### [Git for Windows] との連携

[Git for Windows] と [PuTTY] を連携するには，環境変数 `GIT_SSH` に Plink へのフルパスをセットする。

```text
setx GIT_SSH=C:\path\to\PuTTY\plink.exe
```

ただし，最近の [Git for Windows] ではインストール時の SSH 接続設定で [PuTTY] を使うよう指定すれば自動的に環境変数をセットしてくれるので，手動で設定する必要はないと思われる。

## Windows 版 gpg-agent は [OpenSSH] と相性が悪い？

`gpg-agent` は [OpenSSH] の `ssh-agent` と置き換えることもできる[^sa]。
`gpg-agent` への SSH 鍵のインポートには `ssh-add` を使うのだが， Windows 環境では上手く動かない。
どうやらファイル・ディスクリプタ `S.gpg-agent.ssh` が上手く機能しないようだ。

[MSYS2] 版[^m] と [PowerShell 用](https://github.com/PowerShell/Win32-OpenSSH "PowerShell/Win32-OpenSSH: Win32 port of OpenSSH")の [OpenSSH] で試してみたのだが，いずれも上手くいかなかった。
[MSYS2] 版については [MSYS2] の [GnuPG] を使えば上手くいくのかもしれないが，面倒なので試してない。

[^sa]: `gpg-agent.conf` ファイルに `enable-ssh-support` オプションをセットする。
[^m]: [Git for Windows] に同梱されている bash も [MSYS2] である。

### ssh-pageant 経由で [OpenSSH] と連携できる！

フィードバックで教えていただいたのだが `ssh-pageant` というツールがあって，これを経由して `gpg-agent` と [OpenSSH] を繋げられるようだ。

- [GitHub - cuviper/ssh-pageant: An SSH authentication agent for Cygwin/MSYS to PuTTY's Pageant.](https://github.com/cuviper/ssh-pageant)

`ssh-pageant` は最近の [Git for Windows] には同梱されている。
`ssh-pageant` が常駐している状態では， [OpenSSH] からは `ssh-agent` が起動しているように見える。
一方， `gpg-agent` には [PuTTY] から要求があるように見える。

{{< fig-mermaid >}}
graph LR
  OpenSSH-- request key -->ssh-pageant
  ssh-pageant-- store key -->OpenSSH
  ssh-pageant-- request key -->gpg-agent
  gpg-agent-- store key -->ssh-pageant
{{< /fig-mermaid >}}

この方法の利点は， `enable-putty-support` オプションを有効にしておけば， `gpg-agent` は [PuTTY] とも [OpenSSH] とも全く等価にアクセスできる点だろう。
欠点は，結局のところ [PuTTY] は手放せないということだろうか（だって bash 以外の環境では今まで通りだし）。
[Git for Windows] に関しては... もう Plink での接続でいいんじゃないかな。

まぁ Windows だしね。

`ssh-pageant` の起動は bash から以下のコマンドを起動する。

```text
$ eval $(/usr/bin/ssh-pageant -r -a "/tmp/.ssh-pageant-$USERNAME")
```

`ssh-agent` と似たようなやりかただな。
`ssh-pageant` は完全に土管として機能するので，上のコマンドを `.bash_profile` か `.bashrc` に書いておいて bash 起動時に常駐させておけばいいだろう[^bash1]。

[^bash1]: うちの [Git for Windows] 付属の bash では何故か `.bash_profile` を読んでくれない。ので `.bashrc` に書いている。

 `ssh-add` で鍵のインポートができるかどうかは試してないが（Pageant からインポートできてるし）， `ssh-add -l` ってやったらちゃんと鍵情報が取れたので，多分できるんじゃないかな？

## 参考になる（かもしれない） Web ページ

- [hdk の自作ソフトの紹介 | PuTTYjp](http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html)
- [iceiv+putty](http://ice.hotmint.com/putty/)
- [公開鍵認証によるSSH接続 - PuTTYの使い方 - Linux入門 - Webkaru](http://webkaru.net/linux/putty-ssh-login-public-key/)
- [Windowsでのssh agent - Qiita](http://qiita.com/tsuyoshi_cho/items/79c09905ae3f192b3a0f)
- [SSH authentication using a YubiKey on Windows](https://developers.yubico.com/PGP/SSH_authentication/Windows.html)
- [Git/Git for Windows/SSHにPuTTYを使う - yanor.net/wiki](http://yanor.net/wiki/?Git%2FGit%20for%20Windows%2FSSH%E3%81%ABPuTTY%E3%82%92%E4%BD%BF%E3%81%86)
- [GPG Agent under Windows as SSH Agent for  - Super User](https://superuser.com/questions/911496/gpg-agent-under-windows-as-ssh-agent-for-git-bash)
- [Using GnuPG (2.1) for SSH authentication](http://incenp.org/notes/2015/gnupg-for-ssh-authentication.html)
- [Convert keys between GnuPG, OpenSsh and OpenSSL - Sysmic.org](http://www.sysmic.org/dotclear/index.php?post/2010/03/24/Convert-keys-betweens-GnuPG%2C-OpenSsh-and-OpenSSL)
- [Git for WindowsのシェルからPageantでSSH - Qiita](https://qiita.com/jkr_2255/items/f1ebd3fa4a9bf8ee1b03)
- [Windowsでのssh agent - Qiita](https://qiita.com/tsuyoshi_cho/items/79c09905ae3f192b3a0f)
- [Big Sky :: Windowsでもssh-agentとssh-addを使ってパスフレーズ入力を省略する。](https://mattn.kaoriya.net/software/20081106192615.htm)
- [KeeAgent – lechnology.com](https://lechnology.com/software/keeagent/) : パスワード管理ツール [KeePass] のプラグインで， [KeePass] のパスワードデータベースを使って SSH 鍵を管理し Agent 機能で SSH に鍵を渡す仕組みらしい。  [PuTTY] と [OpenSSH] に対応しているようだ


[前回]: {{< ref "/openpgp/using-gnupg-for-windows-1.md" >}} "GnuPG for Windows インストール編"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[Git for Windows]: https://git-for-windows.github.io/ "Git for Windows"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[KeePass]: https://keepass.info/ "KeePass Password Safe"
