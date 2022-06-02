+++
title = "Go 1.18.2 のリリース【セキュリティ・アップデート】"
date =  "2022-06-01T08:08:23+09:00"
description = "description"
image = "/images/attention/tools.png"
tags  = [ "tools" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++






## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.18.2.linux-amd64.tar.gz`](https://go.dev/dl/go1.18.2.linux-amd64.tar.gz)）を取ってきてインストールすることを強く推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.18.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.18.2.linux-amd64.tar.gz
$ sudo mv go go1.18.2
$ sudo ln -s go1.18.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.18.1 linux/amd64
```

Windows は [Scoop] 経由で OK

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2022-24675]: https://nvd.nist.gov/vuln/detail/CVE-2022-24675
[CVE-2022-28327]: https://nvd.nist.gov/vuln/detail/CVE-2022-28327
[CVE-2022-27536]: https://nvd.nist.gov/vuln/detail/CVE-2022-27536

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
