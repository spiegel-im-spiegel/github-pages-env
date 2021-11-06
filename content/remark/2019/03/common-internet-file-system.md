+++
title = "CIFS 経由で NAS に接続する"
date = "2019-03-31T20:24:45+09:00"
description = "今回は CIFS (Common Internet File System) クライアントを使って NAS に接続を試みる。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "nas", "samba", "cifs" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] から[自宅の簡易 NAS]({{<ref "/remark/2018/04/nas.md" >}} "NAS を導入した") に接続することを考える。

相手の NAS には [Samba] サーバとして接続できることは分かっているので，こちらは対応するクライアント・ツールを用意すればいいわけだ。

というわけで今回は CIFS (Common Internet File System) クライアントを使って接続を試みる。

## CIFS クライアントの導入

[Ubuntu] であれば APT (Advanced Package Tool) で導入できる。

```text
$ sudo apt install cifs-utils smbclient
```

念のため動作確認しておこう。

```text
$ mount.cifs -V
mount.cifs version: 6.11
```

バージョンは 6.11 か（2021-11-06 Ubuntu 21.10 にて）。

## 事前準備（2019-04-03 追記）

実際に作業に入る前にいくつか事前設定をしておく。

### ホスト名の解決

IP アドレス指定のままでもいいのだが，もし LAN 内で NAS の名前解決が出来ないなら `/etc/hosts` ファイルで名前を定義しておくといいだろう。

```text
10.0.0.100  nas01
```

### Samba グループを作成する

共有ディレクトリをマウントできるグループを作成する。

```text
$ sudo groupadd samba
```

さらに

```text
$ sudo visudo
```

で `sudo` 用の設定ファイルを開き以下の行を追加する。

```text
%samba  ALL=(ALL) /bin/mount,/bin/umount,/sbin/mount.cifs,/sbin/umount.cifs
```

これで `samba` グループに追加したユーザは自身のホーム・ディレクトリ以下に NAS の共有ディレクトリをマウントできる。
さっそく追加しておこう。

```text
$ sudo adduser username samba
```

これで準備完了。

## 一時的な接続

まずマウント先のディレクトリを作成する（ディレクトリの場所や名前は他と被らなければなんでもいい）。

```text
$ mkdir ~/nas
```

この `~/nas` ディレクトリに対して以下のコマンドを実行する。

```text
sudo mount -t cifs //nas01/shared ~/nas -o username=sambauser,noexec,nosuid,nodev
[sudo] username のパスワード: 
Password for sambauser@//nas01/shared:  ********
```

これで `~/nas` ディレクトリに対して共有ディレクトリ `//nas01/shared` を一時的にマウントできる。
マウント時に指定できる主なパラメータ（値がある場合は `param=value` 形式で指定）は以下の通り[^cifs1]。

[^cifs1]: オプションの詳細は `man mount.cifs` でマニュアルを参照するとよい。

| パラメータ名  | 既定値    | 内容                                                                                                |
| ------------- | --------- | --------------------------------------------------------------------------------------------------- |
| `username`    | `root`    | 接続先のユーザ名                                                                                    |
| `password`    |           | 接続先ユーザのパスワード。指定しない場合は接続時に訊いてくる                                        |
| `domain`      |           | 接続先のドメイン名またはワークグループ名（必要な場合のみ）                                          |
| `credentials` |           | 接続情報が書かれたファイル（詳しくは後述する）                                                      |
| `uid`         | `root`    | 共有ディレクトリ・ファイルの（見かけの）オーナー名                                                  |
| `gid`         | `root`    | 共有ディレクトリ・ファイルの（見かけの）グループ名                                                  |
| `file_mode`   | `0755`    | 共有ファイルの（見かけの）ファイルモード                                                            |
| `dir_mode`    | `0755`    | 共有ディレクトリの（見かけの）ディレクトリモード                                                    |
| `vers`        | `3.0`     | SMB プロトコル・バージョン（`1.0`/`2.0`/`3.0`/`3.1.1`）                                             |
| `sec`         | `ntlmssp` | セキュリティ・モード (`none`/`krb5`/`krb5i`/`ntlm`/`ntlmi`/`ntlmv2`/`ntlmv2i`/`ntlmssp`/`ntlmsspi`) |
| `noexec`      | ―         | 共有ディレクトリ上のファイルの実行をさせない                                                        |
| `nosuid`      | ―         | 共有ディレクトリ・ファイルの SUID ビットを認識しない                                                |
| `nodev`       | ―         | 共有ディレクトリ上のデバイスファイルを作用させない                                                  |
| `ro`          | ―         | 読み込み専用でマウントする                                                                          |
| `rw`          | ―         | 読み書き可でマウントする                                                                            |

コマンドラインでマウントする際は絶対に `password` オプションを付けないこと（履歴に残っちゃうからね）。
あるいは以下の内容の接続情報ファイル `~/.nascred` を作って（ファイル名は適当）

```ini
username=sambauser
password=password_string
domain=domainname
```

ファイルモードを `0400` に設定して他ユーザから見られないようにし

```text
$ sudo chamod 0400 ~/.nascred
```

`credentials` オプションで指定してマウントする。

```text
sudo mount -t cifs //nas01/shared ~/nas -o credentials=~/.nascred,noexec,nosuid,nodev
```

他ユーザから見えなくてもパスワードを平文で保存するのはちょっと... という人は `password` の指定行を外してしまえばコマンド実行時にパスワードを訊いてくる。

マウントを解除するには

```text
$ sudo umount ~/nas
```

で OK。

{{< div-box type="markdown" >}}
**【2021-11-06 追記】**
ドメインではなくワークグループを構成している NAS の場合は

```ini
username=sambauser
password=password_string
workgroup=WORKGROUP
```

などとワークグループ名を指定する。
{{< /div-box >}}

## /etc/fstab を使って起動時にマウントする

`/etc/fstab` ファイルに以下の行を追加することで起動時に共有ファイルをマウントできる。

```text
# <file system> <mount point>      <type> <options>                                                 <dump> <pass>
//nas01/shared  /home/username/nas cifs   _netdev,credentials=/home/username/.nascred,noexec,nosuid 0      0
```

ちなみに `_netdev` はネットワークの設定が終わったあとにマウントするよう指示するオプションである（先頭のアンダーバーを忘れずに）。

起動時のマウントではないが以下の2つのオプションを追加する手もある。

| パラメータ名 | 内容                         |
| ------------ | ---------------------------- |
| `noauto`     | 起動時にはマウントを行わない |
| `user`       | 一般ユーザもマウントできる   |

```text
# <file system> <mount point>      <type> <options>                                                     <dump> <pass>
//nas01/shared  /home/username/nas cifs   noauto,user,credentials=/home/username/.nascred,noexec,nosuid 0      0
```

これなら管理者モードでなくともログイン後に任意のタイミングで

```text
$ mount ~/nas
```

という感じにコマンドラインを少し簡略化してマウントできる（`umount` も同様）。

よし。
これで次に進める。
次は [GnuPG] かな。

## ブックマーク

- [mount.cifs](http://www.samba.gr.jp/project/translation/3.5/htmldocs/manpages-3/mount.cifs.8.html)
- [Samba/SambaClientGuide - Community Help Wiki](https://help.ubuntu.com/community/Samba/SambaClientGuide)
- [mountコマンドでCIFSをファイル共有 - profaim.jp](http://www.profaim.jp/tools/soft/linux/mnt_cifs.php)
- [mount.cifsに付けるオプション - Qiita](https://qiita.com/kakinaguru_zo/items/af0122f79af0aa0913b7)
- [Ubuntu から Windows の共有フォルダをマウントして利用する - Qiita](https://qiita.com/mdstoy/items/54925cdcbca6d558b666)
- [samba-client／cifs-utilsを使ってWindowsの共有フォルダをマウントする - Qiita](https://qiita.com/You_name_is_YU/items/85ffbffee744f6f494ed)
- [Windows 10 に cifs 接続する - Qiita](https://qiita.com/office-itou/items/0f8df0c5f8c7022b7fad)
- [Linuxはサンバ（Samba）で踊る――WindowsとLinuxのファイル共有のいま (1/2)：その知識、ホントに正しい？ Windowsにまつわる都市伝説（25） - ＠IT](https://www.atmarkit.co.jp/ait/articles/1502/13/news042.html)
- [fstab - ArchWiki](https://wiki.archlinux.jp/index.php/Fstab)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Samba]: https://www.samba.org/ "Samba - opening windows to a wider world"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "B01CXL8NL6" %}} <!-- 【改訂新版】サーバ構築の実例がわかるSamba［実践］入門 -->
