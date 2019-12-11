+++
title = "Git v2.24.1 のリリース【セキュリティ・アップデート】"
date =  "2019-12-11T10:25:03+09:00"
description = "今回は多くのセキュリティ・アップデートを含むため，必ず対応すること。"
image = "/images/attention/tools.png"
tags  = [ "git", "tools", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git] v2.24.1 を含む複数のバージョンがリリースされた。
今回は多くのセキュリティ・アップデートを含むため，必ず対応すること。

- [[ANNOUNCE] Git v2.24.1 and others](https://public-inbox.org/git/xmqqr21cqcn9.fsf@gitster-ct.c.googlers.com/T/)

対象となるバージョンは v2.14.x から v2.24.x までの各マイナーバージョン。
改修対象となる脆弱性の CVE ID はバージョンごとに異なるが v2.24.1 の場合は以下の通り（{{% quote lang="en" %}}[Git v2.24.1 and others](https://public-inbox.org/git/xmqqr21cqcn9.fsf@gitster-ct.c.googlers.com/T/){{% /quote %}} より抜粋）。

- [CVE-2019-1348](https://nvd.nist.gov/vuln/detail/CVE-2019-1348): The `--export-marks` option of git fast-import is exposed also via the in-stream command feature `export-marks=...` and it allows overwriting arbitrary paths.
- [CVE-2019-1349](https://nvd.nist.gov/vuln/detail/CVE-2019-1349): When submodules are cloned recursively, under certain circumstances Git could be fooled into using the same Git directory twice. We now require the directory to be empty.
- [CVE-2019-1350](https://nvd.nist.gov/vuln/detail/CVE-2019-1350): Incorrect quoting of command-line arguments allowed remote code execution during a recursive clone in conjunction with SSH URLs.
- [CVE-2019-1351](https://nvd.nist.gov/vuln/detail/CVE-2019-1351): While the only permitted drive letters for physical drives on
   Windows are letters of the US-English alphabet, this restriction does not apply to virtual drives assigned via subst `<letter>: <path>`. Git mistook such paths for relative paths, allowing writing outside of the worktree while cloning.
- [CVE-2019-1352](https://nvd.nist.gov/vuln/detail/CVE-2019-1352): Git was unaware of NTFS Alternate Data Streams, allowing files inside the `.git/` directory to be overwritten during a clone.
- [CVE-2019-1353](https://nvd.nist.gov/vuln/detail/CVE-2019-1353): When running Git in the Windows Subsystem for Linux (also known as "WSL") while accessing a working directory on a regular Windows drive, none of the NTFS protections were active.
- [CVE-2019-1354](https://nvd.nist.gov/vuln/detail/CVE-2019-1354): Filenames on Linux/Unix can contain backslashes. On Windows, backslashes are directory separators. Git did not use to refuse to write out tracked files with such filenames.
- [CVE-2019-1387](https://nvd.nist.gov/vuln/detail/CVE-2019-1387): Recursive clones are currently affected by a vulnerability that is caused by too-lax validation of submodule names, allowing very targeted attacks via remote code execution in recursive clones.
- [CVE-2019-19604](https://nvd.nist.gov/vuln/detail/CVE-2019-19604): The change to disallow `submodule.<name>.update=!command` entries in `.gitmodules` which was introduced v2.15.4 (and for which v2.17.3 added explicit fsck checks) fixes the vulnerability in v2.24.x[^git1] where a recursive clone followed by a submodule update could execute code contained within the repository without the user explicitly having asked for that.

[^git1]: [CVE-2019-19604](https://nvd.nist.gov/vuln/detail/CVE-2019-19604) の説明文の {{% quote lang="en" %}}v2.24.x{{% /quote %}} は 2.20.x 以降の各バージョンに適宜読み替えていただきたい。

[Git for Windows] は既に v2.24.1.windows.2 をリリースしている。

- [Release Git for Windows 2.24.1(2) · git-for-windows/git · GitHub](https://github.com/git-for-windows/git/releases/tag/v2.24.1.windows.2)

[Ubuntu] の APT の標準リポジトリは最新版を提供してないが， [USN-4220-1](https://usn.ubuntu.com/4220-1/ "USN-4220-1: Git vulnerabilities | Ubuntu security notices") によると各バージョンの [git] でアップデートを行っているようだ。
なお Ubuntu Git Maintainers が提供している [PPA] 版では v2.24.1 がリリースされている（2019-12-11 追記）。

- [Git stable releases : “Ubuntu Git Maintainers” team](https://launchpad.net/~git-core/+archive/ubuntu/ppa)

アップデートは計画的に。

**以下は愚痴：** やっぱ Windows のサブシステムに Linux を入れるってのは筋が悪いよな。
ライセンス云々を抜きにしてもファイルシステムの根底が違うんだから両者は相容れないだろう。
Windows が心を入れ替えてファイルシステムを UNIX 互換に履き替えるってのならともかく，サブシステムでお茶を濁そうとする WSL は（Windows Vista 以来の）最大の失敗作な気がする。そもそも文字エンコードディングからして今だに [DBCS の呪い](https://twitter.com/mattn_jp/status/1198638950480433153)から抜け出せないようだし[^ce1]。

[^ce1]: 文字集合および文字エンコーディングの歴史については[拙文](https://github.com/spiegel-im-spiegel/charset_document "spiegel-im-spiegel/charset_document: 「文字コードとその実装」 upLaTeX ドキュメント")を参考にどうぞ。ただし原文（非公開）は20年以上前に書かれたもので既に内容が古いのはあしからずご了承の程を。つまり文字エンコーディングの問題はそれだけ根深いってことなんだけど。

## ブックマーク

- [「Git」に複数の脆弱性、Windowsユーザーはとくに注意 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1223826.html)

- [PPA から Git をインストールする]({{< ref "/remark/2019/04/install-git-from-ppa.md" >}})
- [Ubuntu アプリケーションにおけるセキュリティ・アップデート一覧]({{< ref "/release/vuln-list.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
