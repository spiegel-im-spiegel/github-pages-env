+++
title = "gpgpdump v0.12.0 をリリースした"
date =  "2021-01-23T11:14:29+09:00"
description = "各 shell 用の自動補完スクリプトを吐き出せるようにした。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "gpgpdump", "golang", "shell"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を可視化する [gpgpdump] の v0.12.0 をリリースした。

- [Release v0.12.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.12.0)

今回は，お試し機能として `completion` サブコマンドを追加して各 shell 用の自動補完スクリプトを吐き出せるようにした。
つっても [spf13/cobra][cobra] の[機能](https://github.com/spf13/cobra/blob/master/shell_completions.md)を使ってるだけだけどね。

Bash, Zsh, Fish, PowerShell の自動補完機能に対応している。
本当は [NYAGOS] でも使えるようにしたかったんだけど， [cobra] も [NYAGOS] も双方カスタマイズの仕方がよく分からなくて，今回は諦めた。
今後の課題としておこう。

たとえば Linux 上の bash であれば

```text
$ source <(gpgpdump completion bash)
```

で取り敢えず試すことができる。
また

```text
sudo sh -c "gpgpdump completion bash > /usr/share/bash-completion/completions/gpgpdump"
```

とかすればシステム全体に設定可能である。

他の shell については

```text
$ gpgpdump completion -h
```

とすれば簡単な使い方が表示されるので参考にして欲しい。

## ブックマーク

- [OpenPGP の実装]({{< rlnk "openpgp/" >}})
- [OpenPGP パケットを可視化する gpgpdump]({{< ref "/release/gpgpdump.md" >}})

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[RFC 4880]: https://tools.ietf.org/html/rfc4880
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/
[Go]: https://golang.org/ "The Go Programming Language"
[cobra]: https://github.com/spf13/cobra "spf13/cobra: A Commander for modern Go CLI interactions"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
