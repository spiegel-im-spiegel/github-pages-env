+++
date = "2015-11-11T20:28:23+09:00"
description = "0.65 以下の PuTTY を使っている人は必ず 0.66 にアップデートすること。日本語版もリリースされている。"
draft = false
tags = ["security", "cryptography", "vulnerability", "windows", "tools", "ssh", "putty"]
title = "PuTTY 0.66 リリース（Security Fix）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

Windows 用の SSH（Secure SHell） クライアント [PuTTY] 0.66 がリリースされている。

{{< fig-quote title="PuTTY vulnerability vuln-ech-overflow" link="http://www.chiark.greenend.org.uk/~sgtatham/putty/wishlist/vuln-ech-overflow.html" lang="en" >}}
<q>Versions of PuTTY and pterm between 0.54 and 0.65 inclusive have a potentially memory-corrupting integer overflow in the handling of the ECH (erase characters) control sequence in the terminal emulator.</q>
{{< /fig-quote >}}

以前のバージョンには ECH (ERASE CHARACTER) 制御シーケンスにおいてメモリ破壊を伴う整数オーバーフロー（memory-corrupting integer overflow）脆弱性があるそうな。
ECH については以下を参照のこと。

- {{< pdf-file title="Standerd ECMA-48: Control Functions for Coded Character Sets" link="(http://www.ecma-international.org/publications/files/ECMA-ST/Ecma-048.pdf" >}}
- [Man page of CONSOLE_CODES](http://linuxjm.osdn.jp/html/LDP_man-pages/man4/console_codes.4.html)

0.65 以下の [PuTTY] を使っている人は **必ず** 0.66 にアップデートすること。
なお，以下の日本語版も 0.66 がリリースされている。

- [PuTTYjp](http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html "hdk の自作ソフトの紹介 | PuTTYjp")
- [iceiv+putty](http://ice.hotmint.com/putty/ "iceiv+putty")

そういや，最近「整数オーバーフロー」な脆弱性報告をよく見かけるな。
Google のこいつとかもそうだっけ。

- [JVNDB-2015-005816 Google Picasa における整数オーバーフローの脆弱性 - JVN iPedia](http://jvndb.jvn.jp/ja/contents/2015/JVNDB-2015-005816.html)

流行ってるのか？

[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free telnet/ssh client"
