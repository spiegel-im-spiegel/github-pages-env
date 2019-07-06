+++
title = "gpgpdump v0.6.0 をリリースした"
date =  "2019-07-06T21:41:07+09:00"
description = "このバージョンから OpenPGP 公開鍵サーバ上の公開鍵を直接検査できるよう HKP モードを作った。"
image = "/images/attention/openpgp.png"
tags = ["tools", "openpgp", "golang", "gpgpdump"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[OpenPGP] パケットの内容を視覚化する [gpgpdump] の v0.6.0 をリリースした。

- [Release v0.6.0 · spiegel-im-spiegel/gpgpdump · GitHub](https://github.com/spiegel-im-spiegel/gpgpdump/releases/tag/v0.6.0)

「[OpenPGP 公開鍵サーバにおける公開鍵の汚染問題]({{< ref "/remark/2019/07/openpgp-certificate-flooding.md" >}})」を受け， [OpenPGP] 公開鍵サーバ上の公開鍵を直接検査できるよう HKP アクセスモードを作った。

```text
$ gpgpdump hkp -h
Dumps from OpenPGP key server

Usage:
  gpgpdump hkp [flags] <user id>

Flags:
  -h, --help               help for hkp
      --keyserver string   OpenPGP key server (default "keys.gnupg.net")
      --port int           port number of OpenPGP key server (default 11371)
      --proxy string       URL of proxy server
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

HKP (HTTP Keyserver Protocol) に対応する公開鍵サーバであれば公開鍵パケットをダウンロードして解析できる。
たとえば `0x491F9BDF0F7BD4AD` の鍵を調べたいなら

```text
$ gpgpdump hkp 0x491F9BDF0F7BD4AD
Public-Key Packet (tag 6) (397 bytes)
    Version: 4 (current)
    Public key creation time: 2008-10-30T02:50:43Z
        49 09 21 03
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (3072 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (41 bytes)
    User ID: BuZz spacedout   <satoshin@vistomail.com>
Signature Packet (tag 2) (447 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (41 bytes)
        Signature Creation Time (sub 2): 2008-10-30T02:50:43Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Key Expiration Time (sub 9): 4016.9231134259257 days after (2019-10-30T01:00:00Z)
        Preferred Symmetric Algorithms (sub 11) (6 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
            Symmetric Algorithm: IDEA (sym 1)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA-1 (hash 2)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-224 (hash 11)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x491f9bdf0f7bd4ad
    Hash left 2 bytes
        f5 97
    RSA signature value m^d mod n (3071 bits)
User ID Packet (tag 13) (41 bytes)
    User ID: Satoshi Nakamoto <satoshin@vistomail.com>
Signature Packet (tag 2) (447 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (41 bytes)
        Signature Creation Time (sub 2): 2008-10-30T02:50:43Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Key Expiration Time (sub 9): 4016.9231134259257 days after
        Preferred Symmetric Algorithms (sub 11) (6 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
            Symmetric Algorithm: IDEA (sym 1)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA-1 (hash 2)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-224 (hash 11)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x491f9bdf0f7bd4ad
    Hash left 2 bytes
        f5 97
    RSA signature value m^d mod n (3071 bits)
Public-Subkey Packet (tag 14) (397 bytes)
    Version: 4 (current)
    Public key creation time: 2008-10-30T02:50:43Z
        49 09 21 03
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (3072 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (421 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (15 bytes)
        Signature Creation Time (sub 2): 2008-10-30T02:50:43Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
        Key Expiration Time (sub 9): 4016.9231134259257 days after (2019-10-30T01:00:00Z)
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x491f9bdf0f7bd4ad
    Hash left 2 bytes
        57 48
    RSA signature value m^d mod n (3072 bits)
```

てな感じにできる。

ちなみに `--raw` オプションを付けるとダウンロードしたテキストをそのまま表示する。

```text
$ gpgpdump hkp 0x491F9BDF0F7BD4AD --raw
<?xml version="1.0" encoding="utf-8"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd" >
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>Public Key Server -- Get "0x491f9bdf0f7bd4ad "</title>
<meta http-equiv="Content-Type" content="text/html;charset=utf-8" />
<style type="text/css">
/*<![CDATA[*/
 .uid { color: green; text-decoration: underline; }
 .warn { color: red; font-weight: bold; }
/*]]>*/
</style></head><body><h1>Public Key Server -- Get "0x491f9bdf0f7bd4ad "</h1>
<pre>
-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: SKS 1.1.6
Comment: Hostname: pgpkeys.uk

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
=LSEt
-----END PGP PUBLIC KEY BLOCK-----
</pre>
</body></html>
```

これで [OpenPGP] 公開鍵サーバ上の公開鍵の事前チェックがしやすくなるだろう。

なお，今回の機能追加により既存の機能を若干変更した。
以前ならファイルを指定して調べる場合は

```text
$ gpgpdump sig.asc
Signature Packet (tag 2) (63 bytes)
    Version: 3 (old)
    Hashed material (5 bytes)
        Signiture Type: Signature of a binary document (0x00)
        Signature creation time: 1998-11-27T09:35:42Z
    Key ID: 0xa79778e247b63037
    Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
    Hash Algorithm: SHA-1 (hash 2)
    Hash left 2 bytes
        27 ae
    DSA value r (159 bits)
    DSA value s (159 bits)
Literal Data Packet (tag 11) (45 bytes)
    Literal data format: b (binary)
    File name: hoge
    Modification time of a file: 1975-04-26T19:41:04Z
    Literal data (35 bytes)
```

でよかったが，このバージョンからは `-f` オプションでファイルを指定する必要がある。

```text
$ gpgpdump -f sig.asc
Signature Packet (tag 2) (63 bytes)
    Version: 3 (old)
    Hashed material (5 bytes)
        Signiture Type: Signature of a binary document (0x00)
        Signature creation time: 1998-11-27T09:35:42Z
    Key ID: 0xa79778e247b63037
    Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
    Hash Algorithm: SHA-1 (hash 2)
    Hash left 2 bytes
        27 ae
    DSA value r (159 bits)
    DSA value s (159 bits)
Literal Data Packet (tag 11) (45 bytes)
    Literal data format: b (binary)
    File name: hoge
    Modification time of a file: 1975-04-26T19:41:04Z
    Literal data (35 bytes)
```

いや，マジすんません。

パイプを使う場合は今までどおり

```text
$ cat sig.asc | gpgpdump
Signature Packet (tag 2) (63 bytes)
    Version: 3 (old)
    Hashed material (5 bytes)
        Signiture Type: Signature of a binary document (0x00)
        Signature creation time: 1998-11-27T09:35:42Z
    Key ID: 0xa79778e247b63037
    Public-key Algorithm: DSA (Digital Signature Algorithm) (pub 17)
    Hash Algorithm: SHA-1 (hash 2)
    Hash left 2 bytes
        27 ae
    DSA value r (159 bits)
    DSA value s (159 bits)
Literal Data Packet (tag 11) (45 bytes)
    Literal data format: b (binary)
    File name: hoge
    Modification time of a file: 1975-04-26T19:41:04Z
    Literal data (35 bytes)
```

でオッケ。

## ブックマーク

- [The OpenPGP HTTP Keyserver Protocol (HKP) draft-shaw-openpgp-hkp-00.txt](https://tools.ietf.org/id/draft-shaw-openpgp-hkp-00.txt) : HKP の元ネタになっている（らしい）ドラフト案
- [PGP HKP Keyservers](https://people.spodhuis.org/phil.pennock/pgp-keyservers)
- [HKPプロトコル](http://ccm.sherry.jp/documents/hkp_protocol.html)

[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[pgpdump]: http://www.mew.org/~kazu/proj/pgpdump/ "pgpdump"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
    <dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
