+++
title = "Windows 11 のカスタマイズ"
date =  "2025-01-11T18:40:16+09:00"
description = "Firefox 導入 / 最新の PowerShell を入れる / Scoop でトラブる / NYAGOS を導入して WIndows Terminal の shell として構成する / 自作ツールを導入する / GnuPG を導入する / PuTTY の導入 / パッケージマネージャは管理者権限下で使う / カスタマイズ作業はまだまだ続く"
image = "/images/attention/tools.png"
tags = [ "windows", "scoop", "nyagos", "git", "gnupg", "tools", "putty", "nyagos" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ Windows をゼロからカスタマイズするのって久しぶりでさ。
色々忘れちゃってるわけさ。

というわけでこの記事は，[前回]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした")セットアップしたミニPCのカスタマイズで試行錯誤した作業記録である。

## Firefox 導入

とりあえず既定ブラウザは Firefox で。

- [デスクトップ、モバイル、エンタープライズ用 Firefox をダウンロード - Mozilla](https://www.mozilla.org/ja/firefox/)

アドオンは [uBlock Origin](https://ublockorigin.com/ "uBlock Origin - Free, open-source ad content blocker.") と [Kagi Search for Firefox](https://addons.mozilla.org/ja/firefox/addon/kagi-search-for-firefox/ "Kagi Search for Firefox – Get this Extension for 🦊 Firefox") を入れる。
[Kagi Search] については以下を参照のこと。

- [Kagi Search を試してみる 〜 検索サービスも有料の時代？]({{< ref "/remark/2024/06/kagi-search.md" >}})

## 最新の PowerShell を入れる

まずは [Scoop] を入れたいところだが Windows Terminal を開くと

{{< fig-img src="./windows-terminal-powershell-5.png" title="PowerShell on Windows Terminal (1)" link="./windows-terminal-powershell-5.png" width="1115" >}}

とか言われる。
折角なので最新の PowerShell を入れておこう。

- [Windows PowerShell 5.1 から PowerShell 7 への移行 - PowerShell | Microsoft Learn](https://learn.microsoft.com/ja-jp/powershell/scripting/whats-new/migrating-from-windows-powershell-51-to-powershell-7?view=powershell-7.4)

インストールは Microsoft Store からできる。
インストールすると Windows Terminal にも設定が反映されて

{{< fig-img src="./windows-terminal-powershell-7.png" title="PowerShell on Windows Terminal (2)" link="./windows-terminal-powershell-7.png" width="1115" >}}

てな感じにできる。

## Scoop でトラブる

では，いよいよ [Scoop] をインストールする。
Windows Terminal で PowerShell を開いて

```text
PS > Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
PS > Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

とすればOK。
これで環境変数 `PATH` への登録も完了している。

動作確認も兼ねて git をインストールしてみる。

```text
PS C:> scoop install git
Installing '7zip' (24.09) [64bit] from 'main' bucket
7z2409-x64.msi (1.9 MB) [========================================] 100%
Checking hash of 7z2409-x64.msi ... ok.
Extracting 7z2409-x64.msi ... ERROR Exit code was 1603!
```

Git をインストールするために解凍ツールの [7-Zip] をインストールするのだが [7-Zip] の解凍に失敗してるやん `orz`

ググってみたら [Issue](https://github.com/ScoopInstaller/Scoop/issues/2225 "Installation failed for 7zip · Issue #2225 · ScoopInstaller/Scoop · GitHub") で上がっていたのだが，その通りにやっても解決せず。
しょうがないので [Scoop] の PowerShell スクリプトを眺めてみると，どうやら

```text
PS > scoop config use_external_7zip true
```

と唱える[^sc1] ことで [Scoop] 外の手動インストールした [7-Zip] を使ってくれるらしい。
それじゃあ [7-Zip] を手動でインストールして `use_external_7zip` を有効にしてみよう。

[^sc1]: `scoop config <name> <value>` で設定した値を削除するには `scoop config rm <name>` を唱える。詳しくは `scoop help config` を参照のこと。

```text
PS C:\Users\username> scoop config use_external_7zip true
'use_external_7zip' has been set to 'true'
PS C:\Users\username> scoop install git
Installing 'git' (2.47.1) [64bit] from 'main' bucket
PortableGit-2.47.1-64-bit.7z.exe (60.0 MB) [===================================] 100%
Checking hash of PortableGit-2.47.1-64-bit.7z.exe ... ok.
Extracting PortableGit-2.47.1-64-bit.7z.exe ... done.
Linking ~\scoop\apps\git\current => ~\scoop\apps\git\2.47.1
Creating shim for 'sh'.
Creating shim for 'bash'.
Creating shim for 'git'.
Creating shim for 'gitk'.
Making C:\Users\username\scoop\shims\gitk.exe a GUI binary.
Creating shim for 'git-gui'.
Making C:\Users\username\scoop\shims\git-gui.exe a GUI binary.
Creating shim for 'scalar'.
Creating shim for 'tig'.
Creating shim for 'git-bash'.
Making C:\Users\username\scoop\shims\git-bash.exe a GUI binary.
Creating shortcut for Git Bash (git-bash.exe)
Creating shortcut for Git GUI (git-gui.exe)
Running post_install script...done.
'git' (2.47.1) was installed successfully!
Notes
-----
Set Git Credential Manager Core by running: "git config --global credential.helper manager"

To add context menu entries, run 'C:\Users\username\scoop\apps\git\current\install-context.reg'

To create file-associations for .git* and .sh files, run
'C:\Users\username\scoop\apps\git\current\install-file-associations.reg'
```

今度は上手くいった。
Git の設定はのちほど（commit/push とかしない限り，とりあえず無問題）。

これで [7-Zip] をコントロールできるってことは [zstd をサポートしていない](https://zenn.dev/zetamatta/scraps/b21750b7ac7c06 "scoopシステム崩壊の序曲")問題もどうにかなるんじゃないだろうか。
機会があれば試してみようか。

ちなみに [Scoop] 自体をアンインストールしたい場合は

```text
PS > scoop uninstall scoop
```

でいけるようだ。
これで環境変数 `PATH` への記述も綺麗に消してくれる。
これで上手く行かない場合は，最悪 `scoop` フォルダを力技で削除する（笑）

```text
PS C:\Users\username> del .\scoop -Force
```

この状態で再度インストールした直後に `scoop uninstall scoop` すれば綺麗にアンインストールできる。
でも `scoop config` で設定した内容は削除できないんだよなぁ。
レジストリにもないっぽいし，どこに保存されてるんだろう。

## NYAGOS を導入して WIndows Terminal の shell として構成する

（詳しくは [NYAGOS] 作者による「[scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146)」を参照のこと）

ようやく [NYAGOS] をインストールできる。

```text
PS > scoop install nyagos
Installing 'nyagos' (4.4.16_0) [64bit] from 'main' bucket
nyagos-4.4.16_0-windows-amd64.zip (2.4 MB) [=================================] 100%
Checking hash of nyagos-4.4.16_0-windows-amd64.zip ... ok.
Extracting nyagos-4.4.16_0-windows-amd64.zip ... done.
Linking ~\scoop\apps\nyagos\current => ~\scoop\apps\nyagos\4.4.16_0
Creating shim for 'nyagos'.
```

おー。
ちゃんと最新バージョンだな。
ありがたや {{% emoji "ペコン" %}}

次に [NYAGOS] を プロファイルに追加する。
こんな感じ。

{{< fig-img src="./setup-windows-terminal-1.png" title="Settings Windows Terminal (1)" link="./setup-windows-terminal-1.png" width="1115" >}}

JSON だとこんな感じ。

```json
{
    "commandline": "%USERPROFILE%\\scoop\\shims\\nyagos.exe",
    "guid": "{ce45bb84-5d8b-49a4-be7d-47d0449207c5}",
    "hidden": false,
    "icon": "none",
    "name": "nyagos",
    "startingDirectory": "%USERPROFILE%"
}
```

`guid` は勝手に割り振られる値なのでご注意を。

ついでに既定の shell を [NYAGOS] にする。

{{< fig-img src="./setup-windows-terminal-2.png" title="Settings Windows Terminal (2)" link="./setup-windows-terminal-2.png" width="1115" >}}

では改めて Windows Terminal を起動してみよう。

{{< fig-img src="./windows-terminal-nyagos.png" title="NYAGOS on Windows Terminal" link="./windows-terminal-nyagos.png" width="1115" >}}

よーし，うむうむ，よーし。
やっとこれで馴染みのあるコンソールになったよ。
ところで [NYAGOS] って2024年で10周年だったんだな。
おめでとうございます。

## 自作ツールを導入する

詳しくは「[オレオレ Scoop Bucket を作ってみた](/release/2023/01/my-scoop-bucket.md)」を参照のこと。

ここでは作業記録のみ残しておく。

```text
$ scoop bucket add goark https://github.com/goark/scoop-bucket
Checking repo... OK
The goark bucket was added successfully.

$ scoop search gpgpdump
Results from local buckets...

Name     Version Source Binaries
----     ------- ------ --------
gpgpdump 0.15.7  goark

$ scoop install gpgpdump
Installing 'gpgpdump' (0.15.7) [64bit] from 'goark' bucket
gpgpdump_0.15.7_Windows_64bit.zip (3.6 MB) [=================================] 100%
Checking hash of gpgpdump_0.15.7_Windows_64bit.zip ... ok.
Extracting gpgpdump_0.15.7_Windows_64bit.zip ... done.
Linking ~\scoop\apps\gpgpdump\current => ~\scoop\apps\gpgpdump\0.15.7
Creating shim for 'gpgpdump'.
'gpgpdump' (0.15.7) was installed successfully!

$ gpgpdump.exe version
gpgpdump v0.15.7
repository: https://github.com/goark/gpgpdump
```

問題なさそうだな。

## GnuPG を導入する

[Scoop] で [GnuPG] のバージョンを確認する。

```text
$ scoop search gnupg
Results from local buckets...

Name   Version Source Binaries
----   ------- ------ --------
gnupg  2.4.7   main
gnupg1 1.4.23  main
```

よしよし。
ちゃんと安定版 2.4 系の最新バージョンになってるな。
それではインストールしよう。

```text
$ scoop install gnupg
Installing 'gnupg' (2.4.7) [64bit] from 'main' bucket
Loading gnupg-w32-2.4.7_20241125.exe from cache
Checking hash of gnupg-w32-2.4.7_20241125.exe ... ok.
Extracting gnupg-w32-2.4.7_20241125.exe ... done.
Running installer script...done.
Linking ~\scoop\apps\gnupg\current => ~\scoop\apps\gnupg\2.4.7
Adding ~\scoop\apps\gnupg\current\bin to your path.
Persisting home
'gnupg' (2.4.7) was installed successfully!

$ gpg --version
gpg (GnuPG) 2.4.7
libgcrypt 1.11.0
Copyright (C) 2024 g10 Code GmbH
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:\Users\username\scoop\apps\gnupg\current\home
Supported algorithms:
Pubkey: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
Cipher: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
        CAMELLIA128, CAMELLIA192, CAMELLIA256
Hash: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
Compression: Uncompressed, ZIP, ZLIB, BZIP2

$ gpg-connect-agent /bye
gpg-connect-agent: gpg-agentが動いていません - 開始します'C:\\Users\\username\\scoop\\apps\\gnupg\\current\\bin\\gpg-agent.exe'
gpg-connect-agent: agent の起動のため、8秒待ちます...
gpg-connect-agent: agentへの接続が確立しました
```

おっ。
ちゃんと `gpg-agent` が動く。
以前 [Scoop] 版の `gpg-agent` が[動かない]({{< ref "" >}})現象があったが，今は問題なさそうだな。

ここで ssh 接続にも使える鍵を生成しておこう。
鍵の作成方法は以下を参照のこと。

- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})
- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

作成した鍵の ssh 公開鍵を GitHub に上げておく。

- [GitHub アカウントへの新しい SSH キーの追加 - GitHub Docs](https://docs.github.com/ja/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account)

## PuTTY の導入

[winget] を使って [PuTTY] を導入しようと思ったのだが

```text
PS > winget search putty
ソースを更新できませんでした: winget
ソースの検索中にエラーが発生しました;結果は含まれません: winget
名前                                             ID             バージョン ソース
-----------------------------------------------------------------------------------
PuTTY                                            XPFNZKSKLBP7RJ Unknown    msstore
SSH Client: Terminal, SFTP, Mosh, Telnet & Putty 9PKGBV7S35T0   Unknown    msstore
```

あれ？ [winget] ってまともに動いてないのか？ しょうがない。
手動で入れよう。
と思って[ダウンロードページ](https://www.chiark.greenend.org.uk/~sgtatham/putty/latest.html)から msi ファイルをダウンロードしてインストーラを起動しようとしたら失敗した。
これ [Scoop] で [7-Zip] をインストールして失敗した症状と同じか？ そもそも msi でのインストールができなくなってる？ うーん，分からん。

それじゃあ [PuTTY-ranvis](https://www.ranvis.com/putty "PuTTYrv (PuTTY-ranvis) - Ranvis software") 版にするか。
`7z` 圧縮されたファイルをダウンロードして展開して手動で `PATH` を通した。

前節で鍵を作って GitHub に ssh 公開鍵を登録してあることが前提。
[GnuPG] のホームディレクトリに `gpg-agent.conf` ファイルを作成して以下の記述を行う。

```text
enable-putty-support
```

`gpg-agent` を再起動して `gpg-agent.conf` ファイルの設定を読み込ませる。

```text
$ gpg-connect-agent killagent /bye
OK closing connection

$ gpg-connect-agent killagent /bye
gpg-connect-agent: no running gpg-agent - starting 'C:\\Users\\username\\scoop\\apps\\gnupg\\current\\bin\\gpg-agent.exe'
gpg-connect-agent: waiting for the agent to come up ... (8s)
gpg-connect-agent: connection to the agent established
OK closing connection
```

これで [PuTTY] の `pageant` の代わりに `gpg-agent` が秘密鍵の管理を行う（厳密には `gpg-agent` は秘密鍵を保持せずパスフレーズのみ短期間保持する）。
試しに GitHub に接続してみよう。

```text
$ plink git@github.com
Using username "git".
Access granted. Press Return to begin session.
```

大丈夫そうかな。

git から [PuTTY] を使って ssh 接続できるように環境変数 `GIT_SSH` に `plink.exe` のフルパスをセットすること。

## パッケージマネージャは管理者権限下で使う

どうも msi ファイルを起動してインストールを行う際は管理者権限が必要らしい。
Windows Terminal を管理者権限で起動して

```text
$ msiexec /i installer.msi
```

などとやったら問題なくインストールできた。
エクスプローラから msi ファイルを起動する場合は管理者権限にならないので注意。
またアンインストールの際には

```text
$ msiexec /x installer.msi
```

でアンインストールする。
設定の `[アプリ]` → `[インストールされているアプリ]` のアプリ一覧からアンインストールしようとしても失敗するので，こちらも注意。

[Scoop] も管理者権限下であれば [7-Zip] をインストールできた。
[winget] も管理者権限下ならまともに動く。

```text
$ winget search putty
名前                                             ID                          バージョン  一致       ソース
-----------------------------------------------------------------------------------------------------------
PuTTY                                            XPFNZKSKLBP7RJ              Unknown                msstore
SSH Client: Terminal, SFTP, Mosh, Telnet & Putty 9PKGBV7S35T0                Unknown                msstore
PuTTY                                            PuTTY.PuTTY                 0.82.0.0               winget
NETworkManager                                   BornToBeRoot.NETworkManager 2024.6.15.0 Tag: putty winget
PortX                                            NetSarangComputer.PortX     2.2.12      Tag: putty winget
PuTTY CAC                                        NoMoreFood.PuTTY-CAC        0.82.0.1    Tag: putty winget
MTPuTTY                                          TTYPlus.MTPutty             1.8         Tag: putty winget
SuperPuTTY                                       JimRadford.SuperPuTTY       1.5.0                  winget

$ winget install PuTTY.PuTTY
見つかりました PuTTY [PuTTY.PuTTY] バージョン 0.82.0.0
このアプリケーションは所有者からライセンス供与されます。
Microsoft はサードパーティのパッケージに対して責任を負わず、ライセンスも付与しません。
ダウンロード中 https://the.earth.li/~sgtatham/putty/0.82/w64/putty-64bit-0.82-installer.msi
  ██████████████████████████████  3.62 MB / 3.62 MB
インストーラーハッシュが正常に検証されました
パッケージのインストールを開始しています...
インストールが完了しました
```

**パッケージマネージャを使う場合は管理者権限にしろ**ってことか。
と考えるなら納得できなくもないが，こんな面倒臭かったっけ？ とりあえず msi ファイルを直に扱うのはやめたほうがいいということは分かった。

## カスタマイズ作業はまだまだ続く

けど，セットアップが面倒そうなのは終わったかな。
一休みして続きをしよう。

### インストールメモ

- [Git Extensions | Git Extensions is a standalone UI tool for managing Git repositories](https://gitextensions.github.io/) : [winget] の `GitExtensionsTeam.GitExtensions` でインストール可。勝手に `Git.Git` をインストールするので個別に `winget uninstall` した
- [KDiff3 - Homepage](https://kdiff3.sourceforge.net/) : [winget] の `KDE.KDiff3` でインストール可
- [Amazon.co.jp: Kindle無料アプリ: Kindleストア](https://www.amazon.co.jp/b?ie=UTF8&node=26197586051) : [winget] の `Amazon.Kindle` はバージョンが古い
- [Steam, The Ultimate Online Game Platform](https://store.steampowered.com/about/) : [winget] の `Valve.Steam` でいいのか分からなかったので，サイトからインストーラをダウンロードした。リモートデスクトップ越しでは流石に入力遅延があるが，パズルゲームなら違和感なくできるか

- [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})

## ブックマーク

- [SSH 接続をテストする - GitHub Docs](https://docs.github.com/ja/authentication/connecting-to-github-with-ssh/testing-your-ssh-connection)

- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}})

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[winget]: https://github.com/microsoft/winget-cli "microsoft/winget-cli: Windows Package Manager CLI (aka winget)"
[Scoop]: https://scoop.sh/ "Scoop"
[7-Zip]: https://www.7-zip.org/ "7-Zip"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[Git Extensions]: https://gitextensions.github.io/ "Git Extensions | Git Extensions is a standalone UI tool for managing Git repositories"

## 作業中の BGV (メン限配信以外)

- [【日本史】今年はヘビ年！古代日本のヘビエピソード集めてみました🐍【古代日本史VTuber きら子】 - YouTube](https://www.youtube.com/watch?v=JCnOfhBETlI)
- [【雑談】お正月が終わったり、とある美術館に行ってきたり、【北白川かかぽ/VEE】 - YouTube](https://www.youtube.com/watch?v=P4KrXKc2E1k)
- [【60万人耐久】年内目標達成にむけて✨歌を歌うぞおおお！！！【一条莉々華 / ホロライブ ReGLOSS】 - YouTube](https://www.youtube.com/watch?v=ZU8jJqWsN2I)
