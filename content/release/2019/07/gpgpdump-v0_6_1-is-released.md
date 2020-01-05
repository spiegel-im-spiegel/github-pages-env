+++
title = "gpgpdump v0.6.1 をリリースした"
date =  "2019-07-22T22:28:48+09:00"
description = "主な変更としては HKP アクセスモードの --proxy オプションを削除した。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の v0.6.1 をリリースした。

- [Release v0.6.1 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.1)

主な変更としては HKP アクセスモードの `--proxy` オプションを削除した。

```text
$ gpgpdump hkp -h
Dumps from OpenPGP key server

Usage:
  gpgpdump hkp [flags] <user ID or key ID>

Flags:
  -h, --help               help for hkp
      --keyserver string   OpenPGP key server (default "keys.gnupg.net")
      --port int           port number of OpenPGP key server (default 11371)
      --raw                output raw text from OpenPGP key server
      --secure             enable HKP over HTTPS

Global Flags:
  -a, --armor        accepts ASCII input only
      --debug        for debug
      --indent int   indent size for output string
  -i, --int          dumps multi-precision integers
  -j, --json         output with JSON format
  -l, --literal      dumps literal packets (tag 11)
  -m, --marker       dumps marker packets (tag 10)
  -p, --private      dumps private packets (tag 60-63)
  -t, --toml         output with TOML format
  -u, --utc          output with UTC time
```

我ながらもの知らずにも程があると思うが `net/`[`http`] パッケージにある [`http`]`.Client` って既定でプロキシに対応してるんだねぇ。
プロキシサーバを指定するには [`http`]`.Client` の `Transport` 要素を弄ればいいんだけど 既定の `DefaultTransport` ってのが

```go
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```

と定義されていて，この中で指定されている [`http`]`.ProxyFromEnvironment()` 関数は 環境変数 `HTTP_PROXY`, `HTTPS_PROXY` および `NO_PROXY` を見てプロキシ情報を適切にセットしてくれるらしい（ちなみに，これらの環境変数名は大文字でも小文字でもちゃんと認識してくれるようだ）。

なのでコマンドラインでプロキシ・サーバを指定する必要はないってこと。
勉強になりました。

## ブックマーク

- [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c)

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[`http`]: https://golang.org/pkg/net/http/ "http - The Go Programming Language"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "B07MPK2F11" %}} <!-- クレイジーキャッツ・スーパー・デラックス -->
