+++
title = "gpgpdump v0.12.1 をリリースした"
date =  "2021-02-23T21:25:43+09:00"
description = "今回は Go 1.16 への対応が主。 機能上の修正・変更はない。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.12.1 をリリースした。

- [Release v0.12.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.12.1)

今回は [Go] 1.16 への[対応]({{< ref "/golang/deprecation-of-ioutil.md" >}} "io/ioutil の非推奨化について")が主。
機能上の修正・変更はない。

あと `go.mod` ファイルや GitHub Actions の設定等を調整して

```text
$ go install github.com/spiegel-im-spiegel/gpgpdump@latest
```

で最新版をダウンロード&ビルド&インストールできるようにしてみた。

[gpgpdump] が依存する自作パッケージも併せてバージョンアップしているが，そっちはまぁいいか。

コンパイラのおかげなのか，依存している外部パッケージのおかげなのか分からないが，前のバージョンより実行バイナリのサイズがちょびっとだけ小さくなっている。
コンパイラのおかげだといいな。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
