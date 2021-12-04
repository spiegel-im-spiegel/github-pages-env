+++
title = "GNKF: NKF ぽいなにか の v0.5.0 をリリースした"
date =  "2021-05-23T14:11:55+09:00"
description = "データのハッシュ値を取得・検証する機能を追加した。"
image = "/images/attention/tools.png"
tags  = [ "tools", "gnkf", "golang", "hash" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語][Go]における文字列処理の習作 [gnkf] の v0.5.0 をリリースした。

- [Release v0.5.0 · spiegel-im-spiegel/gnkf · GitHub](https://github.com/spiegel-im-spiegel/gnkf/releases/tag/v0.5.0)

このバージョンでデータのハッシュ値を取得・検証する機能を追加した。

```text
$ gnkf hash -h
Print or check hash value.
  Support algorithm:
  MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, SHA-512/256

Usage:
  gnkf hash [flags] [file]

Aliases:
  hash, h

Flags:
  -a, --algorithm string   hash algorithm (default "SHA-256")
  -c, --check              don't fail or report status for missing files
  -h, --help               help for hash
      --ignore-missing     don't fail or report status for missing files (with check option)
      --quiet              don't print OK for each successfully verified file (with check option)

Global Flags:
      --debug   for debug
```

MD5, SHA-1, SHA-224, SHA-256, SHA-384, SHA-512, SHA-512/224, および SHA-512/256 各アルゴリズムに対応している[^sha3]。
既定のアルゴリズムは SHA-256 で，それ以外を使う場合は

[^sha3]: SHA-3 は標準パッケージにないので今回は見送った。ちなみに [`golang.org/x/crypto/sha3`](https://pkg.go.dev/golang.org/x/crypto/sha3 "sha3 · pkg.go.dev") パッケージを使えば SHA-3 アルゴリズムも簡単に組み込める。

```text
$ echo Hello World | gnkf h -a md5
e59ff97941044f85df5297e1c302d260  -
```

などとアルゴリズム名を指定する。

また，リリースファイルに対して [`gnkf_0.5.0_checksums.txt`](https://github.com/spiegel-im-spiegel/gnkf/releases/download/v0.5.0/gnkf_0.5.0_checksums.txt) のようなハッシュ値の情報があれば， `-c` オプションを使って

```text
$ curl -L https://github.com/spiegel-im-spiegel/gnkf/releases/download/v0.5.0/gnkf_0.5.0_checksums.txt -O
$ curl -L https://github.com/spiegel-im-spiegel/gnkf/releases/download/v0.5.0/gnkf_0.5.0_Linux_64bit.tar.gz -O
$ gnkf h --ignore-missing -c gnkf_0.5.0_checksums.txt
gnkf_0.5.0_Linux_64bit.tar.gz: OK
WARNING in 9 items: no such file or directory
```

といった感じにダウンロードしたファイルの検証を行うことができる。

まぁ Linux とかなら `sha256sum` コマンド等を使えばいいのだけど， Windows 環境で互換性のあるコマンドが見当たらなかったので [gnkf] に組み込んでみたのだった。
以前に [BASE64 符号化のサブコマンド]({{< ref "/release/2020/10/gnkf-v0_2_0-is-released.md">}})を組み込んだことがあったので，これもありかなと（笑）

余談だが [Go] で MD5 や SHA-1 アルゴリズムを使うと [gosec](https://github.com/securego/gosec "securego/gosec: Golang security checker") に叱られる。
これを回避するには G501 および G505 のワーニングを除外すればよい。

たとえば

```text
$ golangci-lint run --enable gosec --exclude "G501|G505"
```

みたいな感じ。

## ブックマーク

- [GNKF: Network Kanji Filter by Golang]({{< ref "/release/gnkf.md" >}})

[Go]: https://go.dev/
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
