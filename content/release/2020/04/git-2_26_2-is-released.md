+++
title = "Git v2.26.2 のリリース【セキュリティ・アップデート】"
date =  "2020-04-21T09:12:19+09:00"
description = " 前回の CVE-2020-5260 が直りきってなかったってことでいいのかな？"
image = "/images/attention/tools.png"
tags  = [ "git", "tools", "security", "vulnerability", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Git] v2.26.2 を含む複数のバージョンがリリースされた。

- [[Announce] Git v2.26.2 and others](https://lore.kernel.org/git/xmqq4kterq5s.fsf@gitster.c.googlers.com/T/)

対象となるのは 2.17.x から 2.26.x までの各マイナーバージョン。
前回の [CVE-2020-5260] が直りきってなかったってことでいいのかな？

{{< fig-quote type="markdown" title="Git v2.26.2 and others" link="https://lore.kernel.org/git/xmqq4kterq5s.fsf@gitster.c.googlers.com/T/" lang="en" >}}
{{< quote >}}These releases address the security issue CVE-2020-11008, which is similar to the recently addressed CVE-2020-5260{{< /quote >}}.
{{< /fig-quote >}}

[CVE-2020-5260] と同様ということは深刻度も「緊急」ちうことかな。
[CVE-2020-5260] については，[前のリリース記事]({{< relref "./git-2_26_1-is-released.md" >}} "Git v2.26.1 のリリース【セキュリティ・アップデート】")を参考にどうぞ。

## [CVE-2020-11008](https://nvd.nist.gov/vuln/detail/CVE-2020-11008)

未稿

## アップデートは...

アップデートは計画的に。

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/
[Git for Windows]: https://gitforwindows.org/
[PPA]: https://launchpad.net/ubuntu/+ppas "Personal Package Archives : Ubuntu"
[CVE-2020-5260]: https://nvd.nist.gov/vuln/detail/CVE-2020-5260
