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

我ながらもの知らずにも程があると思うが `net/`[`http`] パッケージにある [`http`].`Client` って既定でプロキシに対応してるんだねぇ。
プロキシサーバを指定するには [`http`].`Client` の `Transport` 要素を弄ればいいんだけど 既定の `DefaultTransport` ってのが

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

と定義されていて，この中で指定されている [`http`].`ProxyFromEnvironment()` 関数は 環境変数 `HTTP_PROXY`, `HTTPS_PROXY` および `NO_PROXY` を見てプロキシ情報を適切にセットしてくれるらしい（ちなみに，これらの環境変数名は大文字でも小文字でもちゃんと認識してくれるようだ）。

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>Kindle版</dd>
    <dd>B015643CPE (ASIN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%AF%E3%83%AC%E3%82%A4%E3%82%B8%E3%83%BC%E3%82%AD%E3%83%A3%E3%83%83%E3%83%84%E3%83%BB%E3%82%B9%E3%83%BC%E3%83%91%E3%83%BC%E3%83%BB%E3%83%87%E3%83%A9%E3%83%83%E3%82%AF%E3%82%B9-%E5%B9%B3%E6%88%90%E7%84%A1%E8%B2%AC%E4%BB%BB%E5%A2%97%E8%A3%9C%E7%9B%A4-%E3%82%AF%E3%83%AC%E3%82%A4%E3%82%B8%E3%83%BC%E3%82%AD%E3%83%A3%E3%83%83%E3%83%84/dp/B07MPK2F11?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07MPK2F11"><img src="https://images-fe.ssl-images-amazon.com/images/I/613qa71Ld7L._SL160_.jpg" width="160" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%AF%E3%83%AC%E3%82%A4%E3%82%B8%E3%83%BC%E3%82%AD%E3%83%A3%E3%83%83%E3%83%84%E3%83%BB%E3%82%B9%E3%83%BC%E3%83%91%E3%83%BC%E3%83%BB%E3%83%87%E3%83%A9%E3%83%83%E3%82%AF%E3%82%B9-%E5%B9%B3%E6%88%90%E7%84%A1%E8%B2%AC%E4%BB%BB%E5%A2%97%E8%A3%9C%E7%9B%A4-%E3%82%AF%E3%83%AC%E3%82%A4%E3%82%B8%E3%83%BC%E3%82%AD%E3%83%A3%E3%83%83%E3%83%84/dp/B07MPK2F11?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B07MPK2F11">クレイジーキャッツ・スーパー・デラックス(平成無責任増補盤)</a></dt>
    <dd>青島幸男 (その他), 大瀧詠一 (その他)</dd>
    <dd>ユニバーサル ミュージック 2019-03-26 (Release 2019-03-27)</dd>
    <dd>CD</dd>
    <dd>B07MPK2F11 (ASIN), 4988031321218 (EAN)</dd>
  </dl>
  <p class="description">クレイジーキャッツは私の原点です。子供の頃の刷り込みは恐ろしい（笑）</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-07-22">2019-07-22</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
