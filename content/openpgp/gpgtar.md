+++
title = "複数ファイルをまとめて署名・暗号化する"
date = "2018-05-07T20:46:47+09:00"
description = "暗号化したデータの復号パスワードをメールで送るなどというアホなことをやらかすのは日本独自の慣習らしいが，このような悪習が一刻も早く根絶されることを願う。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "tools", "gnupg" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = true
+++

最近の [GnuPG] には [gpgtar] コマンドが付いていて，これは Windows 版でも使える。

```text
$ gpgtar -h
gpgtar (GnuPG) 2.2.7
Copyright (C) 2018 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Syntax: gpgtar [options] [files] [directories]
Encrypt or sign files into an archive

コマンド:

     --create         create an archive
     --extract        extract an archive
 -e, --encrypt        create an encrypted archive
 -d, --decrypt        extract an encrypted archive
 -s, --sign           create a signed archive
 -t, --list-archive   list an archive

オプション:

 -c, --symmetric      use symmetric encryption
 -r, --recipient USER-ID
                      USER-ID用に暗号化
 -u, --local-user USER-ID
                      署名や復号にこのUSER-IDを使用
 -o, --output FILE    出力をFILEに書き出す
 -v, --verbose        冗長
 -q, --quiet          いくらかおとなしく
     --skip-crypto    skip the crypto processing
     --dry-run        無変更

Tar options:

 -C, --directory DIRECTORY
                      extract files into DIRECTORY
 -T, --files-from FILE
                      get names to create from FILE
     --null           -T reads null-terminated names

バグは <https://bugs.gnupg.org> までご報告ください。
```

[gpgtar] は複数のファイルや指定したフォルダをまとめて暗号化する場合には便利なコマンドである。

たとえば，フォルダ `my-folder` 以下の全てのフォルダ・ファイルをまとめて暗号化するには

```text
$ gpgtar -es -r Alice -u Bob -o hoge my-folder
```

などとすればよい。
ここで `Alice` は暗号用の鍵を示すユーザ ID， `Bob` は署名用の鍵を示すユーザ ID とする。
出力先の書庫ファイル名には `hoge` を指定している。

書庫ファイル `hoge` の中身を見たいなら

```text
$ gpgtar -t hoge
```

とすればよい。
更に書庫ファイルを復号・展開するには

```text
$ gpgtar -d -C . hoge
```

とする。
`-C` オプションで出力先のフォルダ（存在しない場合はエラーになる）を指定している。
`-C` オプションがないとカレント・フォルダに勝手にフォルダを掘って展開してしまうのでご注意を。

[gpgtar] による暗号化手順は以下の通り。

{{< fig-mermaid >}}
graph TD
  tar["tar 形式で書庫化"]
  sign["書庫データへ署名"]
  comp["データの圧縮"]
  enc["圧縮データの暗号化"]

  tar-->sign
  sign-->comp
  comp-->enc
{{< /fig-mermaid >}}

署名を行わない場合は署名処理はスキップされる。

圧縮処理は [GnuPG] の機能を使って行われる。
[GnuPG] で利用可能な圧縮形式は以下の通り。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">ZIP</td>
<td><a href="https://tools.ietf.org/html/rfc1951">RFC1951</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">ZLIB</td>
<td><a href="https://tools.ietf.org/html/rfc1950">RFC1950</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">BZip2</td>
<td><a href="http://www.bzip.org/">bzip2</a></td>
</tr>
</tbody>
</table>
<figcaption>GnuPG で使用可能なデータ圧縮アルゴリズム</figcaption>
</figure>
{{< /div-gen >}}

ちなみに先程の `hoge` ファイルを


```text
$ gpg -d -o hoge.tar hoge
```

とすれば tar 形式のファイル `hoge.tar` が出力される。

暗号化したデータの復号パスワードをメールで送るなどというアホなことをやらかすのは日本独自の慣習らしいが，このような悪習が一刻も早く根絶されることを願う。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgtar]: https://www.gnupg.org/documentation/manuals/gnupg/gpgtar.html "Using the GNU Privacy Guard: gpgtar"
