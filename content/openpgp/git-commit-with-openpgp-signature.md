+++
date = "2017-12-01T17:51:38+09:00"
update = "2018-02-14T23:26:48+09:00"
description = "Git で commit する際に OpenPGP 署名を付加できるらしい。いやぁ，今まで知らなかったよ。物知らずでごめん。"
draft = false
tags = ["git", "github", "gnupg", "openpgp", "cryptography", "certification"]
title = "Git Commit で OpenPGP 署名を行う"
image = "/images/attention/openpgp.png"

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
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"

[scripts]
  mathjax = false
  mermaidjs = false
+++

(move from [{{< ref "/remark/2016/04/git-commit-with-openpgp-signature.md" >}}]({{< ref "/remark/2016/04/git-commit-with-openpgp-signature.md" >}} "Git Commit で OpenPGP 署名を行う"))


[Git] で commit する際に OpenPGP 署名を付加できるらしい。
いやぁ，今まで知らなかったよ。
物知らずでごめん。

- [Git - Signing Your Work](https://git-scm.com/book/uz/v2/Git-Tools-Signing-Your-Work)
- [GitHubでGPGにより署名されたコミットにバッジが表示されるようになったので設定してみる - Qiita](http://qiita.com/prince_0203/items/ef0e12f2f6d150ff0485)

## OpenPGP 鍵の登録

[Git] に OpenPGP 鍵を設定するには以下のコマンドでいいようだ。

```text
$ git config --global user.signingkey 7E20B81C
```

“`7E20B81C`” は OpenPGP 鍵の鍵 ID である。
鍵ごとに異なる値になるので注意。
今回は「[OpenPGP 公開鍵リスト](http://www.baldanders.info/spiegel/pubkeys/)」で公開している鍵の鍵 ID を登録している。
この設定で `.gitconfig` ファイルに以下の記述が追加される。

```ini
[user]
    signingkey = 7E20B81C
```

## Gpg を直接指定する場合

[Git for Windows] の場合， `git bash` に同梱されている `gpg.exe` を使うのだが，困ったことにこれが classic version なのである。

```text
$ gpg --version
gpg (GnuPG) 1.4.20
Copyright (C) 2015 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
Supported algorithms:
Pubkey: RSA, RSA-E, RSA-S, ELG-E, DSA
Cipher: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
        CAMELLIA128, CAMELLIA192, CAMELLIA256
Hash: MD5, SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
Compression: Uncompressed, ZIP, ZLIB, BZIP2
```

既に [GnuPG] の Windows 版をインストールしているのであれば，以下の設定で `gpg.exe` をフルパス指定できる。

```text
$ git config --global gpg.program C:/path/to/GnuPG/bin/gpg.exe
```

ちなみに `.gitconfig` ファイルには以下の記述が追加される。

```ini
[gpg]
    program = C:/path/to/GnuPG/bin/gpg.exe
```

なお Windows 環境であれば，特に理由がないかぎり， stable/modern version を使うことをお薦めする。
Windows 環境以外でも [GnuPG] を classic version と stable/modern version で使い分けている場合は，この設定が役に立つだろう。

## Commit に OpenPGP 署名を付加する

`git commit` 時に OpenPGP 署名を付加したい場合は， version 2 以降なら，以下の設定で常時署名を行うようになる。

```text
$ git config --global commit.gpgsign true
```

`.gitconfig` ファイルには以下の記述が追加される。

```ini
[commit]
    gpgsign = true
```

この状態で逆に commit 時に署名を付加してほしくない場合は `--no-gpg-sign` オプションを付加すればいいようだ。

ちなみに [ATOM] エディタの [git-plus] パッケージで commit してみたが，きちんと署名もできていた。
[Git Extensions] でも問題ないようだ。
「[GitHubでGPGにより署名されたコミットにバッジが表示されるようになったので設定してみる](http://qiita.com/prince_0203/items/ef0e12f2f6d150ff0485)」には「GitHub DesktopはGPGによる署名をサポートしていません」とあるが， Facebook で教えてもらった話によれば， `gpg-agent` と Pinentry が正しく設定されていれば使えるそうだ。

きちんと署名されているか確認するには `git log` コマンドに `--show-signature` オプションを付けるとよい。

### Tag にも Merge にも署名できる

タグに署名を付加する場合は `-s` オプションを付けて

```text
$ git tag -s -a v0.1.0
```

とすればいい。
同様に merge でも `-S` オプションを付けて

```text
$ git merge -S branch
```

でいいようだ。

## GitHub が OpenPGP 署名に対応した

[GitHub] で commit や tag に対する署名を表示できるようになったらしい。

- [GPG signature verification](https://github.com/blog/2144-gpg-signature-verification)

たとえばこんな感じに表示される。

{{< fig-img src="https://farm2.staticflickr.com/1671/26315000570_ba79ae50b1.jpg" title="OpenPGP Key in GitHub (4)" link="https://www.flickr.com/photos/spiegel/26315000570/" >}}

これを有効にするには [GitHub] に OpenPGP 公開鍵を登録して署名検証可能にしなければならない。
公開鍵の登録は設定画面で行う。

{{< fig-img src="https://farm2.staticflickr.com/1679/26494073882_e53d80376b.jpg" title="OpenPGP Key in GitHub (1)" link="https://www.flickr.com/photos/spiegel/26494073882/" >}}

この画面で「New GPG key」ボタンを押すと以下の入力画面になる。

{{< fig-img src="https://farm2.staticflickr.com/1598/26520705641_81e21edd32.jpg" title="OpenPGP Key in GitHub (2)" link="https://www.flickr.com/photos/spiegel/26520705641/" >}}

ここに公開鍵の armor テキストを貼り付けて「Add GPG key」ボタンを押せばよい。
これで OpenPGP 公開鍵の登録は完了である。

{{< fig-img src="https://farm2.staticflickr.com/1493/26586667165_9b7d41f16d.jpg" title="OpenPGP Key in GitHub (4)" link="https://www.flickr.com/photos/spiegel/26586667165/" >}}

なお公開鍵の armor テキストは以下のコマンドで取得できる。

```text
$ gpg --armor --export 7E20B81C
```

“`7E20B81C`” は先ほど説明した鍵 ID である。

Tag や commit に署名することで「なりすまし」に対する抑止になる。
これは特にチームで開発を行う場合に威力を発揮するだろう。
ぜひ習慣付けていきたいものである。

## ブックマーク

- [git(GitHub)でGPGを使った署名をおこなう - Qiita](http://qiita.com/pontago/items/5867b6492e09c34084fe)
- [Yubikeyを使って、Githubのcommitをverifyする - Qiita](https://qiita.com/akashima/items/4b40ccb13ad13dee5cdb)

[Git]: https://git-scm.com/ "Git"
[Git for Windows]: http://git-for-windows.github.io/ "Git for Windows"
[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"
[ATOM]: https://atom.io/ "Atom"
[git-plus]: https://atom.io/packages/git-plus "git-plus"
[Git Extensions]: http://gitextensions.github.io/ "Git Extensions"
[GitHub]: https://github.com/ "GitHub"
