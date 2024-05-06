+++
title = "Chromebook を導入する 3 — GnuPG & OpenSSH"
date =  "2024-04-06T22:07:26+09:00"
description = "GnuPG で OpenSSH 認証鍵を作成し ssh で GitHub に繋いでみる"
image = "/images/attention/openpgp.png"
tags = [ "chromebook", "linux", "gnupg", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [Chromebook を導入する 1]({{< ref "/remark/2024/03/chromebook-1.md" >}})
- [Chromebook を導入する 2 — Linux サブシステム]({{< ref "/remark/2024/04/chromebook-2.md" >}})
- [Chromebook を導入する 3 — GnuPG & OpenSSH]({{< ref "/remark/2024/04/chromebook-3.md" >}}) ← イマココ
- [Chromebook を導入する 4 — Flatpak で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-4.md" >}})
- [Chromebook を導入する 5 — APT で Firefox を導入する]({{< ref "/remark/2024/04/chromebook-5.md" >}})

[前回]に続いて Linux サブシステムのセットアップを行う。
なお，ターミナルはランチャーから起動できる。

{{< fig-img src="./launcher.png" title="ランチャー" link="./launcher.png" width="657" >}}

ターミナルを起動すると以下の表示になってるので “penguin” を押下する。

{{< fig-img src="./terminal.png" title="ターミナル" link="./terminal.png" width="736" >}}

## GnuPG で鍵ペアを生成してみる

ターミナルが起動したところで Linux サブシステムに入ってる（であろう） [GnuPG] を確認する。

{{< fig-img src="./gnupg-1.png" title="GnuPG (1)" link="./gnupg-1.png" width="736" >}}

2.2 系かよ！ Ubuntu の [GnuPG] のバージョンが上がらないと思ったら Debian が元凶なのか `orz`

ま，まぁ，入ってはいるようなのでよしとしよう。
`gpg-agent` のサービスはどうなってるかな。

{{< fig-img src="./gnupg-2.png" title="GnuPG (2)" link="./gnupg-2.png" width="736" >}}

うんうん。
サービスもちゃんと上がってるね。
では，さっそく鍵ペアを生成してみよう。

{{< fig-img src="./gnupg-3.png" title="GnuPG (3)" link="./gnupg-3.png" width="736" >}}

とりあえず署名専用の鍵ペアを作っている。
このあと必要な項目を入力してパスフレーズを設定するのだが

{{< fig-img src="./gnupg-4.png" title="GnuPG (4)" link="./gnupg-4.png" width="736" >}}

おぅふ。
これは `pinentry-curses` かな？ OpenPGP 鍵の生成や操作については以下の拙文を参考にどうぞ。

- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})
- [そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})

なんでも文章で残しておくものである。

## GnuPG で OpenSSH 認証鍵を作成し ssh で GitHub に繋いでみる

[GnuPG] で [OpenSSH] 認証鍵を作成する手順については以下の拙文を参照のこと。

- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

上記記事にない情報としては，環境変数 `SSH_AUTH_SOCK` および `GPG_TTY` の扱いがある。
まず，以下のように環境設定を設定する（`~/.bashrc` 等に書いておくとよい）。

```bash
export SSH_AUTH_SOCK="$(gpgconf --list-dirs agent-ssh-socket)"
export GPG_TTY=$(tty)
```

さらに `ssh` で `pinentry-curses` を呼び出せるように `~/.ssh/config` ファイルに以下の記述を追加する。

```bash
Match host * exec "gpg-connect-agent UPDATESTARTUPTTY /bye"
```

`~/.ssh/config` ファイルのこの記述と `GPG_TTY` 環境変数はセットで設定しないと上手く動かないらしい。

動作確認をしてみよう。

`gpg --export-ssh-key <keyID>` または `ssh-add -L` コマンドで出力した公開鍵を GitHub の settings → SSH and GPG keys で辿れる設定画面に登録する。
これで

```text
$ ssh -T git@github.com
```

などと入力し接続テストができる。

{{< fig-img src="./github-ssh-connect.png" title="GitHub へ ssh 接続テスト" link="./github-ssh-connect.png" width="736" >}}

ログイン処理の途中でパスフレーズの入力がある。

{{< fig-img src="./input-passphrase.png" title="パスフレーズの入力" link="./input-passphrase.png" width="736" >}}

うんうん。
問題なく動作しているな。
今回はここまで。
次回は git かな。

## 【2024-05-06】 pinentry-gnome3 を導入する { #pinentry }

さすがに `pinentry-curses` のままではアレなので [`pinentry-gnome3`] に換装することを考える。
その前に `update-alternatives` コマンドで現在の pinentry の設定を見てみる。

```text
$ sudo update-alternatives --config pinentry
```

と唱えればよい。

{{< fig-img src="./update-alternatives-1.png" title="update-alternatives (1)" link="./update-alternatives-1.png" width="986" >}}

ふむむ。
`pinentry-curses` だけが入ってる状態やね。
ほんじゃあ APT で [`pinentry-gnome3`] を入れてみるか。

{{< fig-img src="./install-pinentry-gnome3-1.png" title="install pinentry-gnome3 (1)" link="./install-pinentry-gnome3-1.png" width="986" >}}

うん。
大丈夫そうだな。
このまま {{% keytop %}}`y`{{% /keytop %}} キー押下で続行する。

{{< fig-img src="./install-pinentry-gnome3-2.png" title="install pinentry-gnome3 (2)" link="./install-pinentry-gnome3-2.png" width="986" >}}

へー。
`ssh-agent` でもこいつを使うのか？ まぁ，今回は関係ないけど（多分）

もう一度 `update-alternatives` で状態を確認してみよう。

{{< fig-img src="./install-pinentry-gnome3-2.png" title="install pinentry-gnome3 (2)" link="./install-pinentry-gnome3-2.png" width="986" >}}

よしよし。
[`pinentry-gnome3`] が自動・優先で設定されてるな。
これで前節のように

```text
$ ssh -T git@github.com
```

のように接続テストを行うと，めでたく以下のプロンプトが表示された。

{{< fig-img src="./pinentry-gnome3.png" title="pinentry-gnome3" link="./pinentry-gnome3.png" >}}

おー。
ちゃんと日本語で表示されてる[^jp1]。
偉い偉い。

[^jp1]: Chromebook の Linux サブシステムを日本語化するにはひとつ前の記事の「[日本語化]({{< ref "/remark/2024/04/chromebook-2.md#jp" >}})」の節を参照のこと。

ここまでできれば前節で説明した環境変数の

```bash
export GPG_TTY=$(tty)
```

や `~/.ssh/config` ファイルの

```bash
Match host * exec "gpg-connect-agent UPDATESTARTUPTTY /bye"
```

の記述は不要になる（環境変数 `SSH_AUTH_SOCK` の指定は必要）。

## ブックマーク

- [GnuPG - ArchWiki](https://wiki.archlinux.jp/index.php/GnuPG)

[前回]: {{< ref "/remark/2024/04/chromebook-2.md" >}} "Chromebook を導入する 2 — Linux サブシステム"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[`pinentry-gnome3`]: https://manpages.debian.org/bookworm/pinentry-gnome3/pinentry-gnome3.1.en.html "pinentry-gnome3(1) — pinentry-gnome3 — Debian bookworm — Debian Manpages"

## 参考

{{% review-paapi "B0BKKF7JHV" %}} <!-- ASUS Chromebook -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B09BMPZ3BZ" %}} <!-- Chromebook仕事術 -->
{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B079NL1L9K" %}} <!-- SSH Mastery -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B0CW35PBT6" %}} <!-- ネコカブリーナ -->
