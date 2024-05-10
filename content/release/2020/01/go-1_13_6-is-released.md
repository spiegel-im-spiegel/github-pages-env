+++
title = "Go 1.13.6 がリリースされた，他"
date =  "2020-01-13T13:14:28+09:00"
description = "他にも pkg/errors v0.9.0 がリリースされたようだ。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "error" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅まきながら（笑）

[Go] 1.13.6 がリリースされた。

- [Go 1.13.6 and Go 1.12.15 are released - Google Group](https://groups.google.com/forum/#!topic/golang-announce/RLFrcJ_FZZs)

今回もセキュリティ・アップデートはなし。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release.html#go1.13.minor" lang="en" >}}
{{% quote %}}go1.13.6 (released 2020/01/09) includes fixes to the runtime and the `net/http` package. See the [Go 1.13.6 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.13.6+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

例によって [Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.13.6.linux-amd64.tar.gz`](https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。

```text
$ cd /usr/local/src
$ sudo curl "https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.13.6.linux-amd64.tar.gz
$ sudo mv go go1.13.6
$ sudo ln -s go1.13.6 go
$ go version # /usr/local/go/bin にパスが通っている場合
go version go1.13.6 linux/amd64
```

アップデートは計画的に。

[Go]: https://go.dev/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## pkg/errors v0.9.0 がリリースされた

[pkg/errors] の v0.9.0 がリリースされたようだ。

{{< fig-quote type="markdown" title="Release errors 0.9.0 · pkg/errors" link="https://github.com/pkg/errors/releases/tag/v0.9.0" lang="en" >}}
{{% quote %}}errors 0.9.0 is a preparation release for a 1.0 final release. Also we were working on removing support for Go 1.8, 1.9 and 1.10 and earlier, and become compatible this package with new way of errors on Go 1.13{{% /quote %}}.
{{< /fig-quote >}}

おおっ！

中身を見ると，標準の  [`errors`]`.Is()`, [`errors`]`.As()`, および [`errors`]`.Unwrap()` 各関数を置き換え可能になっていて，さらに [pkg/errors] の特徴である `Cause()` 関数も使えるようになっている。

```go
func Cause(err error) error {
    type causer interface {
        Cause() error
    }

    for err != nil {
        if cause, ok := err.(causer); ok {
            err = cause.Cause()
        } else if unwrapped := Unwrap(err); unwrapped != nil {
            err = unwrapped
        } else {
            break
        }
    }
    return err
}
```

ふむむー。
拙作の [spiegel-im-spiegel/errs] の `Cause()` 関数もこれに追従しておこうかな。
いや [pkg/errors] の `*.Cause()` メソッドって `*.Unwrap()` メソッドと等価みたいだし，無理に合わせないほうがいいか。

[pkg/errors]: https://github.com/pkg/errors "pkg/errors: Simple error handling primitives"
[spiegel-im-spiegel/errs]: https://github.com/spiegel-im-spiegel/errs "spiegel-im-spiegel/errs: Error handling for Golang"
[`errors`]: https://golang.org/pkg/errors/ "errors - The Go Programming Language"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
