+++
title = "pgpdump 0.34 がリリースされた"
date =  "2021-12-11T21:47:42+09:00"
description = "Windows 用バイナリを作るか悩み中..."
image = "/images/attention/openpgp.png"
tags  = [ "openpgp", "tools", "pgpdump" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

本家 [pgpdump] の v0.34 がリリースされたようだ。

変更点は以下の通り。

{{< fig-quote class="nobox" type="markdown" title="pgpdump/CHANGES" link="https://github.com/kazu-yamamoto/pgpdump/blob/master/CHANGES" lang="en" >}}
```
0.34 2021/12/07

* Uploading modifications to support GnuPG-2.3.3 ECC curves, additional hash and algorithm names.
	https://github.com/kazu-yamamoto/pgpdump/pull/32
* Improved labels for Literal Data Packet fields.
	https://github.com/kazu-yamamoto/pgpdump/pull/29
* Indicate unknown sigtype value is displayed in hex.
	https://github.com/kazu-yamamoto/pgpdump/pull/27
* Fixing cross-building and avoiding infinite loop when invoking BZ2_bzDecompress.
	https://github.com/kazu-yamamoto/pgpdump/pull/25
```
{{< /fig-quote >}}

例によってソースコードのみの提供だが Linux 環境ならビルドは簡単である。

```text
$ ./configure 
...

$ make
gcc -c  -g -O2 -O -Wall pgpdump.c
gcc -c  -g -O2 -O -Wall types.c
gcc -c  -g -O2 -O -Wall tagfuncs.c
gcc -c  -g -O2 -O -Wall packet.c
gcc -c  -g -O2 -O -Wall subfunc.c
gcc -c  -g -O2 -O -Wall signature.c
gcc -c  -g -O2 -O -Wall keys.c
gcc -c  -g -O2 -O -Wall buffer.c
gcc -c  -g -O2 -O -Wall uatfunc.c
gcc -g -O2 -O -Wall -o pgpdump pgpdump.o types.o tagfuncs.o packet.o subfunc.o signature.o keys.o buffer.o uatfunc.o -lbz2 -lz  
```

作者の方は既に Haskell の人で [pgpdump] のメンテナンスは積極的には行っていないみたいだが pull request は歓迎のようだ。
拙作の [gpgpdump] は本家 [pgpdump] をリファレンス実装とみなして参考にさせてもらってるので，今回の [GnuPG] 2.3 系への対応はありがたい。

[前回]({{< ref "/release/2018/05/pgpdump-v0.33-is-released.md" >}} "pgpdump 0.33 がリリース")のリリースは2018年で，このときはまだ Windows 環境だったのだが，今回はどうすっかなぁ。
まぁ，折角 [Azure Virtual Desktop 環境を作った]({{< ref "/remark/2021/12/azure-virtual-desktop.md" >}} "ようやく Azure Virtual Desktop を導入できた")んだし，そっちで MSYS2 環境を作るかな。

[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
