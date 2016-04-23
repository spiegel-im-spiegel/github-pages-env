+++
date = "2016-04-23T15:17:49+09:00"
description = "Git で commit する際に OpenPGP 署名を付加できるらしい。いやぁ，今まで知らなかったよ。物知らずでごめん。"
draft = false
tags = ["git", "github", "gnupg", "openpgp", "cryptography", "certification"]
title = "Git Commit で OpenPGP 署名を行う"

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

```text
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

```text
[gpg]
	program = C:/path/to/GnuPG/bin/gpg.exe
```

Windows 環境以外でも [GnuPG] を classic version と stable/modern version で使い分けている場合は，この設定が役に立つだろう。

なお Windows 環境であれば，特に理由がないかぎり， modern version を使うことをお薦めする。
Windows 環境へ [GnuPG] を導入する方法については以下の拙文を参考にして欲しい。

- [GnuPG Modern Version for Windows ― インストール編]({{< relref "remark/2016/03/using-gnupg-modern-version-1.md" >}})
- [GnuPG Modern Version for Windows ― gpg-agent について]({{< relref "remark/2016/03/using-gnupg-modern-version-2.md" >}})

## Commit に OpenPGP 署名を付加する

`git commit` 時に OpenPGP 署名を付加したい場合は， version 2 以降なら，以下の設定で常時署名を行うようになる。

```text
$ git config --global commit.gpgsign true
```

`.gitconfig` ファイルには以下の記述が追加される。

```text
[commit]
	gpgsign = true
```

この状態で逆に commit 時に署名を付加してほしくない場合は `--no-gpg-sign` オプションを付加すればいいようだ。

ちなみに [ATOM] エディタの [git-plus] パッケージで commit してみたが，きちんと署名もできていた。
「[GitHubでGPGにより署名されたコミットにバッジが表示されるようになったので設定してみる](http://qiita.com/prince_0203/items/ef0e12f2f6d150ff0485)」には「GitHub DesktopはGPGによる署名をサポートしていません」とあるが， Facebook で教えてもらっとところによれば， `gpg-agent` と Pinentry が正しく設定されていれば使えるそうだ。

きちんと署名されているか確認するには `git log` コマンドに --show-signature オプションを付けるとよい。

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

[Git]: https://git-scm.com/ "Git"
[Git for Windows]: http://git-for-windows.github.io/ "Git for Windows"
[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"
[ATOM]: https://atom.io/ "Atom"
[git-plus]: https://atom.io/packages/git-plus "git-plus"
[GitHub]: https://github.com/ "GitHub"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
