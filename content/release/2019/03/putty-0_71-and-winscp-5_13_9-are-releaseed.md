+++
title = "PuTTY 0.71 および WinCSP 5.13.9 がリリースされた【セキュリティ・アップデート】"
date = "2019-03-22T22:41:00+09:00"
description = "ただし PuTTY 関しては具体的な攻略方法がなく，単なる不具合として処理されたようだ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "putty", "ssh", "scp", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[PuTTY] 0.71 がリリースされていていた。

{{< fig-quote title="PuTTY: a free SSH and Telnet client" link="https://www.chiark.greenend.org.uk/~sgtatham/putty/" lang="en" >}}
<q>PuTTY 0.71, released today, includes a large number of security fixes, many of which were found by the recent EU-funded HackerOne bug bounty. There are also other security enhancements (side-channel resistance), and a few new features. It's also the first release to be built for Windows on Arm.</q>
{{< /fig-quote >}}

今回はセキュリティ・アップデートを含んでいる。

- [scp client multiple vulnerabilities](https://sintonen.fi/advisories/scp-client-multiple-vulnerabilities.txt)
- [PuTTY bug pscp-unsanitised-server-output](https://www.chiark.greenend.org.uk/~sgtatham/putty/wishlist/pscp-unsanitised-server-output.html)
- [CVE-2019-6109](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6109) : relating to file names sent by the server
- [CVE-2019-6110](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-6110) : relating to the server's stderr stream

ただし [PuTTY] 関しては具体的な攻略方法がなく，単なる不具合として処理されたようだ。

また，これに関連して [WinCSP] も 5.13.9 にアップデートされた。

アップデートは計画的に。

## ブックマーク

- [64ビット版 PuTTY を導入する]({{< ref "/remark/2018/02/putty-64bit-version.md" >}})
- [GnuPG for Windows : gpg-agent について]({{< ref "/openpgp/using-gnupg-for-windows-2.md" >}}) : GnuPG の gpg-agent を使ってログインを行う方法について

[PuTTY]: https://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client"
[WinCSP]: https://winscp.net/ "WinSCP :: Official Site :: Free SFTP and FTP client for Windows"
<!-- eof -->
