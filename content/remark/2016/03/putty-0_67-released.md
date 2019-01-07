+++
date = "2016-03-08T14:34:05+09:00"
description = "SCP にセキュリティ脆弱性が発見されたようだ。"
draft = false
tags = ["security", "cryptography", "vulnerability", "windows", "tools", "ssh", "putty"]
title = "PuTTY 0.67 リリース（Security Fix）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

Windows 用の SSH（Secure SHell） クライアント [PuTTY] 0.67 がリリースされている。
`pscp.exe` にセキュリティ脆弱性が発見されたようだ。

{{< fig-quote title="PuTTY vulnerability vuln-pscp-sink-sscanf" link="http://www.chiark.greenend.org.uk/~sgtatham/putty/wishlist/vuln-pscp-sink-sscanf.html" lang="en" >}}
<q> Many versions of PSCP prior to 0.67 have a stack corruption vulnerability in their treatment of the 'sink' direction (i.e. downloading from server to client) of the old-style SCP protocol.<br>
In order for this vulnerability to be exploited, the user must connect to a malicious server and attempt to download any file. </q>
{{< /fig-quote >}}

0.66 以下の [PuTTY] を使っている人は **必ず** 0.67 にアップデートすること。
なお，以下の日本語版も 0.67 がリリースされている。

- [PuTTYjp](http://hp.vector.co.jp/authors/VA024651/PuTTYkj.html "hdk の自作ソフトの紹介 | PuTTYjp")
- [iceiv+putty](http://ice.hotmint.com/putty/ "iceiv+putty")

[PuTTY]: http://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free telnet/ssh client"
