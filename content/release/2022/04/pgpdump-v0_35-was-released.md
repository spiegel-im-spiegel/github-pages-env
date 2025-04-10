+++
title = "pgpdump 0.35 がリリースされていた"
date =  "2022-04-14T19:42:39+09:00"
description = "気が付かなかったよ orz"
image = "/images/attention/openpgp.png"
tags  = [ "openpgp", "tools", "pgpdump" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2022-02-28 に本家 [pgpdump] の v0.35 がリリースされていた。
気が付かなかったよ `orz`

0.34 からの変更点は以下の通り。

{{< fig-quote class="nobox" type="markdown" title="pgpdump/CHANGES" link="https://github.com/kazu-yamamoto/pgpdump/blob/master/CHANGES" lang="en" >}}
```
0.35 2022/02/28

* Adding BrainPool-384/512 curve definitions.
	https://github.com/kazu-yamamoto/pgpdump/pull/33
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

拙作の [gpgpdump] は本家 [pgpdump] をリファレンス実装とみなして参考にさせてもらってる。
こちらもよろしく【広告】


[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenPGP]: http://openpgp.org/
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
