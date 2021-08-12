+++
date = "2016-04-17T17:10:49+09:00"
description = "Go 言語に2つの脆弱性がある。脆弱性に対処した 1.6.1 および 1.5.4 がリリースされている。"
tags = ["golang", "security", "vulnerability"]
title = "Go 言語 1.6.1 および 1.5.4 のセキュリティ・アップデート"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]に2つの脆弱性がある。
脆弱性に対処した 1.6.1 および 1.5.4 がリリースされている。

{{< fig-quote title="[security] Go 1.6.1 and 1.5.4 are released - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/9eqIHqaWvck" lang="en" >}}
<q>On Windows, Go loads system DLLs by name with LoadLibrary, making it vulnerable to DLL preloading attacks. For instance, if a user runs a Go executable from a Downloads folder, malicious DLL files also downloaded to that folder could be loaded into that executable.<br>
This is CVE-2016-3958 and was addressed by this change: <a href="https://golang.org/cl/21428">https://golang.org/cl/21428</a></q>
{{< /fig-quote >}}

{{< fig-quote title="[security] Go 1.6.1 and 1.5.4 are released - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/9eqIHqaWvck" lang="en" >}}
<q>Go's crypto libraries passed certain parameters unchecked to the underlying big integer library, possibly leading to extremely long-running computations, which in turn makes Go programs vulnerable to remote denial of service attacks.  Programs using HTTPS client certificates or the Go SSH server libraries are both exposed to this vulnerability.<br>
This is CVE-2016-3959 and was addressed by this change: <a href="https://golang.org/cl/21533">https://golang.org/cl/21533</a></q>
{{< /fig-quote >}}

具体的に CVSS などを記述したページは見つからなかった。
が，順次更新する予定。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
