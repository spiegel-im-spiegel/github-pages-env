+++
title = "OpenSSL 1.1.1k がリリースされた【セキュリティ・アップデート】"
date =  "2021-03-26T19:10:57+09:00"
description = "アップデートは計画的に"
image = "/images/attention/tools.png"
tags  = [ "openssl", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenSSL] 1.1.1k がリリースされた。
2件の脆弱性の改修が行われている。

- [OpenSSL Security Advisory [25 March 2021]](https://www.openssl.org/news/secadv/20210325.txt)

## [CVE-2021-3449]

[OpenSSL] 側は “Severity: High” と評価している。

{{< fig-quote type="markdown" title="CVE-2021-3449" link="https://nvd.nist.gov/vuln/detail/CVE-2021-3449" lang="en" >}}
{{% quote %}}An OpenSSL TLS server may crash if sent a maliciously crafted renegotiation ClientHello message from a client. If a TLSv1.2 renegotiation ClientHello omits the signature_algorithms extension (where it was present in the initial ClientHello), but includes a signature_algorithms_cert extension then a NULL pointer dereference will result, leading to a crash and a denial of service attack. A server is only vulnerable if it has TLSv1.2 and renegotiation enabled (which is the default configuration){{% /quote %}}.
{{< /fig-quote >}}

（以下未稿）

## [CVE-2021-3450]

[OpenSSL] 側は “Severity: High” と評価している。

{{< fig-quote type="markdown" title="CVE-2021-3450" link="https://nvd.nist.gov/vuln/detail/CVE-2021-3450" lang="en" >}}
{{% quote %}}The `X509_V_FLAG_X509_STRICT` flag enables additional security checks of the certificates present in a certificate chain. It is not set by default. Starting from OpenSSL version 1.1.1h a check to disallow certificates in the chain that have explicitly encoded elliptic curve parameters was added as an additional strict check. An error in the implementation of this check meant that the result of a previous check to confirm that certificates in the chain are valid CA certificates was overwritten. This effectively bypasses the check that non-CA certificates must not be able to issue other certificates{{% /quote %}}.
{{< /fig-quote >}}


## アップデートは計画的に

Ubuntu は各バージョンごとにバックポートパッチを当てたバージョンを APT で配布している。

- [USN-4891-1: OpenSSL vulnerability | Ubuntu security notices | Ubuntu](https://ubuntu.com/security/notices/USN-4891-1)

手動で更新するのであれば

```text
$ cd /usr/local/src/
$ sudo curl -L "https://www.openssl.org/source/openssl-1.1.1k.tar.gz" -O
$ sudo curl -L "https://www.openssl.org/source/openssl-1.1.1k.tar.gz.asc" -O
$ gpg -d openssl-1.1.1k.tar.gz.asc # 署名を確認する
$ $ sudo tar xvf openssl-1.1.1k.tar.gz
$ cd openssl-1.1.1k/
$ sudo ./config
$ sudo make
$ sudo make install
$ sudo ldconfig # 念のため
$ openssl version
OpenSSL 1.1.1k  25 Mar 2021
```

などとする。

アップデートは計画的に。

## ブックマーク

- [OpenSSLの脆弱性（CVE-2021-3450、CVE-2021-3449）に関する注意喚起](https://www.jpcert.or.jp/at/2021/at210015.html)
- [Ubuntu に最新の OpenSSL を入れる]({{< ref "/remark/2020/06/installing-openssl-in-ubuntu.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[OpenSSL]: https://www.openssl.org/
[CVE-2021-3449]: https://nvd.nist.gov/vuln/detail/CVE-2021-3449
[CVE-2021-3450]: https://nvd.nist.gov/vuln/detail/CVE-2021-3450
<!-- eof -->
