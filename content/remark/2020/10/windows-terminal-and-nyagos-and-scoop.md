+++
title = "Windows Terminal × NYAGOS × Scoop ＝ ♥"
date =  "2020-10-27T22:17:46+09:00"
description = "とりあえず ATOM エディタと NYAGOS を入れたい。"
image = "/images/attention/kitten.jpg"
tags = ["atom", "editor", "nyagos", "shell", "windows", "terminal", "gnupg", "git", "install"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

アルファ・ベータ・オメガの割り振りはご想像にお任せします {{< emoji ":smile:">}} って，そういう話ではない。

さて，職場で Windows 10 機を支給されたので，とりあえず [ATOM] エディタと [NYAGOS] を入れたいわけですよ。

## Windows 10 に [ATOM] を入れる

まぁこちらはサラッと。

[公式ページ][ATOM]からインストーラを取ってきて普通にインストールすればいいだけ。
なのだが，何故か Windows 版はコマンドラインからの起動が冷遇されているのよ。
古いバージョンにはあった `atom.cmd` がなくなっている。

じゃあ，どうやって起動するのかというと “System Settings” で

{{< fig-img src="./atom-system-settings.png" link="./atom-system-settings.png" title="System Settings in ATOM" width="734" >}}

てな感じにエクスプローラ等のコンテキストメニューから開けるように設定できる。
また [ATOM] が起動している状態でタスクバーの [ATOM] アイコンをピン留めしておけば

{{< fig-img src="./atom-in-task-bar.png" link="./atom-in-task-bar.png" title="ATOM icon in task bar" width="525" >}}

みたいな感じにコンテキストメニューを展開できる。

困るのが `apm` コマンドなのだが，これにはパスを通さずに直接カレントまで降りて操作するしかないだろう。
v1.52.0 なら

```text
$ cd C:\Users\username\AppData\Local\atom\app-1.52.0\resources\app\apm\bin
```

まで降りる。
この場所で apm を使ってログインできる。

```text
$ apm login
Welcome to Atom!

Before you can publish packages, you'll need an API token.

Visit your account page on Atom.io https://atom.io/account,
copy the token and paste it below when prompted.

Press [Enter] to open your account page on Atom.io.
```

ここで `[Enter]` キーを押すと Web ブラウザでアカウントページが開く（要 GitHub アカウント）。
開かない場合は Web ブラウザで直接 [`https://atom.io/account`](https://atom.io/account) を開く。
アカウントページに表示されたアクセス・トークンを入力すればログイン完了。

```text
Token> ****************
Saving token to Keychain done
```

あとは

```text
$ apm stars --install
```

で {{< emoji ":star:" >}} を付けたパッケージをまとめてインストールできる。
ちなみに，どのパッケージに {{< emoji ":star:" >}} を付けたかは

```text
$ apm stars
```

で確認できる。
その他の細々した設定は [Gist に貼り付けておいた](https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c "ATOM Editor の設定（カスタマイズ, Windows 環境用）")のを流用した。
なんでも取っておくものである {{< emoji ":smile:">}}

## Windows Terminal のインストール

これでようやく「メモ帳」から脱することができたので，本格的に環境を作っていこう。

かつて Windows 7 を使ってた頃はターミナル・エミュレータに ConEmu を使っていたが，本家 Microsoft からようやくまともなターミナル・エミュレータがリリースされた。

- [Windows ターミナルの概要 | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows/terminal/)

Windows Terminal は Microsoft Store からインストールできる。
ひとまずこれをインストールして起動し，既定の PowerShell 上で作業する。

## [Scoop] のインストール

[Scoop] は Windows 用のパッケージ・マネージャである。
Windows 用のパッケージ・マネージャといえば [Chocolatey] が有名だが，今回は [Scoop] で。
[Scoop] の特徴は以下の通り。

{{< fig-quote type="markdown" title="scoop / nyagos で始めるコマンドライン生活" link="https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146" >}}
- 利用に管理者権限は不要！
- 実行ファイルは ~\scoop\shims に集約され、環境変数 PATH の肥大化が抑制される
- インストーラのセットアップが簡単
- nyagos が公式レポジトリにある (Chocolatey にもあるけど）

{{< /fig-quote >}}

では早速。

まずは PowerShell でスクリプトを起動できるよう権限を取得する。

```text
PS > Set-ExecutionPolicy RemoteSigned -scope CurrentUser
```

既に許可を取得している場合はこの操作は不要である。
そうしておいて [Scoop] をインストールする。

```text
PS > iwr -useb get.scoop.sh | iex
```

これで OK。
簡単！

## [NYAGOS] のインストール

[NYAGOS] のインストールは `scoop install` コマンドで一発 OK。

```text
PS > scoop install nyagos
```

よーし，うむうむ，よーし。

### Windows Terminal に [NYAGOS] を登録する

Windows Terminal では PowerShell を含む複数の shell を登録できる。

まずは `[Ctrl+,]` で `config.json` ファイル（のコピー）がメモ帳で表示される（コメントを端折っているので注意）。
これを編集して [NYAGOS] を登録するのだ。
具体的には以下の部分を追記する。

```json {hl_lines=["30-37"]}
{
    "$schema": "https://aka.ms/terminal-profiles-schema",
    "defaultProfile": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
    "copyOnSelect": false,
    "copyFormatting": false,
    "profiles":
    {
        "defaults":
        {
        },
        "list":
        [
            {
                "guid": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
                "name": "Windows PowerShell",
                "commandline": "powershell.exe",
                "hidden": false
            },
            {
                "guid": "{0caa0dad-35be-5f56-a8ff-afceeeaa6101}",
                "name": "コマンド プロンプト",
                "commandline": "cmd.exe",
                "hidden": false
            },
            {
                "guid": "{b453ae62-4e3d-5e58-b989-0a998ec441b8}",
                "hidden": false,
                "name": "Azure Cloud Shell",
                "source": "Windows.Terminal.Azure"
            },
            {
                "guid": "{19ddaf5e-e045-481a-bf88-37f7ebe66292}",
                "hidden": false,
                "name": "Nihongo Yet Another GOing Shell",
                "commandline": "%USERPROFILE%\\scoop\\apps\\nyagos\\current\\nyagos.exe",
                "cursorShape": "vintage",
                "startingDirectory": "%USERPROFILE%"
            }
        ]
    },
    "schemes": [],
    "actions":
    [
        { "command": {"action": "copy", "singleLine": false }, "keys": "ctrl+c" },
        { "command": "paste", "keys": "ctrl+v" },
        { "command": "find", "keys": "ctrl+shift+f" },
        { "command": { "action": "splitPane", "split": "auto", "splitMode": "duplicate" }, "keys": "alt+shift+d" }
    ]
}
```

ちなみに `guid` 項目はローカルマシン内で一意であればなんでもいいのだが（上の記述をそのままコピペしても無問題），気になるのであれば PowerShell の以下のコマンドで取得できる。

```text
PS > New-Guid

Guid
----
6c48ee13-e32b-4937-95a5-7e95a2e88613
```

更に `config.json` ファイル（のコピー）の

```json
{
    "defaultProfile": "{61c54bbd-c2c6-5271-96e7-009a87ff44bf}",
}
```

の部分を [NYAGOS] の GUID に書き換えれば

```json
{
    "defaultProfile": "{19ddaf5e-e045-481a-bf88-37f7ebe66292}",
}
```

Windows Terminal の既定の shell を [NYAGOS] にできる。

編集した `config.json` ファイル（のコピー）を保存すれば設定が反映される。
ふぃー，よーやくここまでたどり着いた。

[NYAGOS] を使った楽しいアレコレは以下の Zenn 本に書かれている。

- [scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146)

上述のセットアップについても，もう少し丁寧に書かれているので，是非どうぞ。

### [Go] で GUID を取得する

そういや [Go] の [`github.com/google/uuid`](https://pkg.go.dev/github.com/google/uuid) パッケージを使えば [UUID][RFC 4122] を取得可能だが， UUID は実質 GUID と同じなので，これを利用して

```go
package main

import (
    "fmt"

    "github.com/google/uuid"
)

func main() {
    fmt.Println(uuid.New())
}
```

とすれば簡単に UUID/GUID が取れる。

```text
$ go run sample.go 
f6bdc505-e417-4b7d-a247-a06504cf03a9
```

[Go]: https://golang.org/ "The Go Programming Language"
[RFC 4122]: https://www.rfc-editor.org/rfc/rfc4122.html "RFC 4122: A Universally Unique IDentifier (UUID) URN Namespace"

## [Scoop] を使うなら git は必須

これでインストールは全て完了だが，今後のためにもう少し [Scoop] で遊んでみよう。

まず [Scoop] のバージョンを見ようとしたら

```text
$ scoop -v
Current Scoop version:
git : 用語 'git' は、コマンドレット、関数、スクリプト ファイル、または操作可能なプログラムの名前として認識されません。
名前が正しく記述されていることを確認し、パスが含まれている場合はそのパスが正しいことを確認してから、再試行してください
。
発生場所 行:1 文字:1
+ git --no-pager log --oneline HEAD -n 1
+ ~~~
    + CategoryInfo          : ObjectNotFound: (git:String) [], CommandNotFoundException
    + FullyQualifiedErrorId : CommandNotFoundException
```

って，どエラ怒られた（笑） どうやら [Scoop] はパッケージ管理を git ベースで行っているらしく，ちゃんと使うなら git の導入が必須のようだ。
ほんじゃあ，まぁ

```text
$ scoop install git
Scoop uses Git to update itself. Run 'scoop install git' and try again.
...
```

おっ，アップデートした後にもっかいやれって言ってるな。
ならば，アップデートしよう。

```text
$ scoop update
Updating Scoop...
Updating 'main' bucket...
Checking repo... ok
The main bucket was added successfully.
Scoop was updated successfully!
'itcode"' は、内部コマンドまたは外部コマンド、
操作可能なプログラムまたはバッチ ファイルとして認識されていません。
exit status 1
```

おりょ，また変なエラーが出たよ。
`itcode"` 云々というのは PowerShell 絡みらしい。
なので，同じコマンドを PowerShell 上で走らせれば問題なく完了する。

うむ，次回から気をつけよう。

これでもっかい

```text
$ scoop update git
```

とすれば，最新版が取れるわけだ。

## ついでに [Scoop] で [GnuPG] もインストールする

ついでに [Scoop] で [GnuPG] もインストールしてみよう。
あるかな？

```text
$ scoop search gnupg
'main' bucket:
    gnupg (2.2.23)
    gnupg1 (1.4.23)

$ scoop search gpg
'main' bucket:
    gnupg1 (1.4.23) --> includes 'gpg.exe'
    gpg (2.2.23)
```

どっちだよ（笑）

実はこれ，中身は同じなのだが別々のパッケージとしてインストールされちゃうようだ。
ふむむー。

まぁ，いいや。
今回は `gnupg` で。

```text
$ scoop install gnupg
Installing 'gnupg' (2.2.23) [64bit]

...

Linking ~\scoop\apps\gnupg\current => ~\scoop\apps\gnupg\2.2.23
Persisting home
'gnupg' (2.2.23) was installed successfully!
```

[Scoop] で提供されるパッケージのうち，シングルバイナリのコマンドは `%USERPROFILE%\scoop\shims` フォルダに集められるのだが， [GnuPG] のように複数のバイナリで構成されているものは専用のフォルダを作り環境変数 `PATH` にインストールしたフォルダを追加するようだ。
その際に

```text
Linking ~\scoop\apps\gnupg\current => ~\scoop\apps\gnupg\2.2.23
```

のように `current` フォルダをシンボリックリンクとして設置することで複数のバージョンに対応できるようにしているみたい。

[GnuPG] の動作確認をしておこう。

```
$ gpg --version
gpg (GnuPG) 2.2.23
libgcrypt 1.8.6
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:/Users/username/scoop/apps/gnupg/current/home
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

上述の `Home` だが，これもシンボリックリンクが切られていて，実体は `%USERPROFILE%\scoop\persist\gnupg\home` にある。

{{< div-box type="markdown" >}}
**【2020-11-02 追記】**
[GnuPG] については以下の記事でもう少し掘り下げて紹介している。

- [GnuPG の HOME はどこにある？]({{< ref "/openpgp/gnupg-home-in-windows.md" >}})

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
{{< /div-box >}}

## 今回はここまで

おっと。
昼休憩が終わったようだ。
次回があれば，また今度。

## ブックマーク

- [lukesampson/scoop: A command-line installer for Windows.](https://github.com/lukesampson/scoop)
- [Windows用のコマンドラインインストーラScoop - Qiita](https://qiita.com/iakio/items/78e7f098047ea0a47d70)
- [Scoopを使ったWindows環境構築のススメ - Super!! - Qiita](https://qiita.com/Dooteeen/items/12dc8fb14042888113d0)
- [GUID生成ツール](https://hogehoge.tk/guid/)
- [/bin/shに慣れた人に贈るバッチファイルの書き方](https://zenn.dev/zetamatta/books/c84cbe23093eee1b5830)
- [ScoopでWindowsにおける開発環境構築を最適化しよう | さにあらず](https://blog.satotaichi.info/scoop/index.html)
- [scoopを使ったwindows環境構築の実例 - Qiita](https://qiita.com/sozaiya/items/fd7ec3000939f0697939)
- [Scoop で Git と SourceTree をインストール - ありふれた備忘録](https://blog.isonishi.com/posts/scoop-git-sourcetree/)

[ATOM]: https://atom.io/ "Atom"
[Scoop]: https://scoop.sh/ "Scoop"
[Chocolatey]: https://chocolatey.org/ "Chocolatey Software | Chocolatey - The package manager for Windows"
[NYAGOS]: https://github.com/zetamatta/nyagos "zetamatta/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
