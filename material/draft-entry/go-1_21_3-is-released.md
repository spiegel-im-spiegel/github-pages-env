+++
title = "Go 1_21_3 Is Released"
date =  "2023-10-14T11:38:01+09:00"
description = "description"
image = "/images/attention/tools.png"
tags  = [ "tools" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度遅まきながらでごめんペコン。
[予告](https://groups.google.com/g/golang-announce/c/UXJQvKffcao "[security] Go 1.21.1 and Go 1.20.8 pre-announcement")どおり [Go] 1.21.1 がリリースされた。

- [[security] Go 1.21.2 and Go 1.20.9 are released](https://groups.google.com/g/golang-announce/c/XBa1oHDevAo)
- [[security] Go 1.21.3 and Go 1.20.10 are released](https://groups.google.com/g/golang-announce/c/iNNxDTCjZvo)

今回は CVE ID ベースで5件の脆弱性修正を含んでいる。


## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://go.dev/dl/ "Downloads - go.dev")からバイナリ（[`go1.21.3.linux-amd64.tar.gz`](https://go.dev/dl/go1.21.3.linux-amd64.tar.gz)）を取ってきてインストールすることを推奨する。
以下は完全手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://go.dev/dl/go1.21.3.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.21.3.linux-amd64.tar.gz
$ sudo mv go go1.21.3
$ sudo ln -s go1.21.3 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.21.3 linux/amd64
```

Windows はインストールパッケージを取ってきて直接インストールする。
[Scoop] 経由でも OK

複数バージョンの Go コンパイラを扱いたい場合は

```text
$ go install golang.org/dl/go1.21.1@latest
$ go1.21.1 download
$ go1.21.1 version
go version go1.21.1 linux/amd64
```

てな感じに導入できる。

アップデートは計画的に。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[CVE-2023-39320]: https://nvd.nist.gov/vuln/detail/CVE-2023-39320
[CVE-2023-39318]: https://nvd.nist.gov/vuln/detail/CVE-2023-39318
[CVE-2023-39319]: https://nvd.nist.gov/vuln/detail/CVE-2023-39319
[CVE-2023-39321]: https://nvd.nist.gov/vuln/detail/CVE-2023-39321
[CVE-2023-39322]: https://nvd.nist.gov/vuln/detail/CVE-2023-39322

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4814400047" %}} <!-- 初めてのGo言語 -->
{{% review-paapi "B0CFL1DK8Q" %}} <!-- Go言語 100Tips -->
