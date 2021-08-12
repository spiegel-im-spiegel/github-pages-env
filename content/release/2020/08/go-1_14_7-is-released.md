+++
title = "Go 1.14.7 のリリース【セキュリティ・アップデート】"
date =  "2020-08-07T09:21:59+09:00"
description = "1件のセキュリティ・アップデートを含んでいる。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

週明けに[予告](https://groups.google.com/g/golang-announce/c/_ulYYcIWg3Q "[security] Go 1.14.7 and Go 1.13.15 pre-announcement")されていたとおり， [Go] 1.14.7 がリリースされた。

- [[security] Go 1.14.7 and Go 1.13.15 are released](https://groups.google.com/g/golang-announce/c/NyPIaucMgXo)

以下のセキュリティ・アップデートを含んでいる。

## 【[CVE-2020-16845]】encoding/binary: ReadUvarint and ReadVarint can read an unlimited number of bytes from invalid inputs

{{< fig-quote type="markdown" title="[security] Go 1.14.7 and Go 1.13.15 are released" link="https://groups.google.com/g/golang-announce/c/NyPIaucMgXo" lang="en" >}}
{{% quote %}}Certain invalid inputs to `ReadUvarint` or `ReadVarint` could cause those functions to read an unlimited number of bytes from the `ByteReader` argument before returning an error. This could lead to processing more input than expected when the caller is reading directly from a network and depends on `ReadUvarint` and `ReadVarint` only consuming a small, bounded number of bytes, even from invalid inputs{{% /quote %}}.
{{< /fig-quote >}}

- `CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:N/I:N/A:H`
- 深刻度: 重要 (7.5)

| 基本評価基準     | 評価値       |
| ---------------- | ------------ |
| 攻撃元区分       | ネットワーク |
| 攻撃条件の複雑さ | 低           |
| 必要な特権レベル | 不要         |
| ユーザ関与レベル | 不要         |
| スコープ         | 変更なし     |
| 機密性への影響   | なし         |
| 完全性への影響   | なし         |
| 可用性への影響   | 高           |

## 例によって

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.14.7.linux-amd64.tar.gz`](https://golang.org/dl/go1.14.7.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する（もしくは自力でコンパイルするか）。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.14.7.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.14.7.linux-amd64.tar.gz
$ sudo mv go go1.14.7
$ sudo ln -s go1.14.7 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.14.7 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[CVE-2020-16845]: https://nvd.nist.gov/vuln/detail/CVE-2020-16845
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
