+++
title = "OpenPGP 公開鍵サーバにおける公開鍵の汚染問題"
date =  "2019-07-05T23:46:33+09:00"
description = "新しい OpenPGP 公開鍵サーバや Autocrypt について調査したほうがいいかしらん。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "pki", "security", "certification", "risk", "pgpdump", "gpgpdump" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

7pay のセキュリティ事故があまりにバカすぎるのでブログネタにしてやろうかと思っていたが，個人的にもっと重大な案件が出てきたので，先にこちらの話を書くことにする。

- [PGPのSKSキーサーバーネットワークへの証明書ポイズニング--攻撃を受け開発者らが警鐘 - ZDNet Japan](https://japan.zdnet.com/article/35139514/)

かなりヤバいというか「ついにやっちゃったか」って感じの話なのだが，記事の後半に書かれている

{{< fig-quote title="PGPのSKSキーサーバーネットワークへの証明書ポイズニング--攻撃を受け開発者らが警鐘" link="https://japan.zdnet.com/article/35139514/" >}}
<q>キーサーバーはPGPと、PGPプロトコルにおけるユーザー認証の要となるコンポーネントだ</q>
{{< /fig-quote >}}

というのはかなり言い過ぎである。

というのも，そもそも [OpenPGP] の元祖である [PGP] は必ずしも公開鍵サーバを要件としていたわけではなく（[PGP 本](https://www.amazon.co.jp/exec/obidos/ASIN/4900900028/baldandersinf-22/)を読めば分かるが，当時はフロッピー運用とか当たり前の時代だった），後継である [OpenPGP] においてもそのコンセプトが踏襲されているからだ。
[OpenPGP] の信用モデル（web of trust; 信用の輪）については拙文ながら以下を参考にしてほしい。

- [OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})

この信用モデルの下では

- 多くの電子署名が集まっていること
- 同じ鍵が永続的に使われ続けていること

が鍵の信用を高める条件となっている[^pki1]。
[OpenPGP] 公開鍵サーバにおいて鍵や署名の追記しかできないのにはちゃんと理由があるのだ。

[^pki1]: そういう意味で [OpenPGP] 公開鍵への電子署名は厳密な「証明（certification）」というよりは小切手の裏書きのようなものを連想してもらうのがいいだろう。つまり今回の「公開鍵の汚染問題」は「裏書き spam」と考えると分かりやすい。

とは言え [OpenPGP] 公開鍵サーバが鍵運用において大きなウエイトを占めているのは間違いないことで，鍵の所有者が（電子署名や鍵そのものの削除を含めて）制御できないというのはちょっと，いやだいぶ困った事態となっているのも確かである。

たとえば APT などのパッケージ管理ツールでは，パッケージへの電子署名に [OpenPGP] 公開鍵を使うが，鍵のインポートの際に公開鍵サーバを利用するようだ。

この前紹介した [Mono のインストール]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}} "Ubuntu に Mono を導入する
")でも

```text
$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
```

といった感じで鍵サーバから公開鍵をインポートしている[^gpg1]。

[^gpg1]: ちなみに `--keyserver` とか `--recv-keys` とかは [GnuPG] のオプションである。おそらくこれらのオプションをそのまま [GnuPG] に引き渡しているのだろう。

なので，鍵サーバに登録されている公開鍵が汚染されている（可能性がある）というのは拙いのである。

## 回避策1： [OpenPGP] 公開鍵サーバを使わない

今回のリスクを確実に回避したいのであれば [OpenPGP] 公開鍵サーバを使わないのが手っ取り早い。
以下のように，いったんテキストデータとして吐き出して

```text
$ gpg -a --export alice@example.com
-----BEGIN PGP PUBLIC KEY BLOCK-----

mQENBFofiskBCADjUvPHA3PNscg0K74/Uwxj46+oLsyIy7fYIp/4C4dHejcbbPjx
VFeic9wQ4aQFp3VKjYgONgQrRo/9p40Ei1+PtMAV7D6Oy6dxlV8zyCJcSf74ahpB
B15GyA7v4uvTf0Py+Ujyt241ik0fXeLEuwt7p4SIbgJnQs1Fb+61wo8UcCFOLJO5
An6HjXNgNs6fFoiTad+T4PfaTbRHLfFPkoqmDUKWy40hjWl+Ui0QborXH+PUeUm9
vgHbqZzS0QRDGI7rO9AeJ6LweBkP1A2qbDLyexS/F+WUEcY0b76IQM5XH0txwnnl
uCPYcQfIGWce3US1GWJhChF9s/bMGVXOEJbvABEBAAG0GUFsaWNlIDxhbGljZUBl
eGFtcGxlLmNvbT6JAVQEEwEIAD4CGwMFCwkIBwIGFQgJCgsCBBYCAwECHgECF4AW
IQR5/SuZ88YtLTuFC7yTs1CUdYINXQUCWh+LMAUJA8JnZwAKCRCTs1CUdYINXcKT
B/4tLFaPRe289GcX91yLJ/yPS0JvvJKyZzjpNqLbKHuQHPEqGromMGlP4LcaGdFL
rVZ36W3kVk+75q8JFkld0eRS22vftjz6lA9lyb3W9lU1CayF5s3IsC/Ehj55uaHc
OHnp6rl7zEeIdvca6yV0gwySs3j9VPHy58zNrpN/clHoB4Zozy6vCXFMShyLc/wF
brPySf/5LP/642Uro92M2lbkIvZpDhZCVG7s7Ilz3BzsTTNMPkPd5yvdGa5lHQzK
OmXHaxydOYbEWBgqRGqzEIIoLbEd8KHxJVIVDfcAQCjSWRUjAUSDLpBokGsKoQfp
41NjWwjkIsfyJ2tDUeRPGYRbuQENBFofiskBCACzyYfIB+/ZwJBJXw7WMDlEKdnz
L4abwVpw9rBGAWGXjaC/cu7l0svNilXyTgZNq4uKddJ6aYjs7of0SaBl20I8aj5G
nbw0pG+KkoYhfpZaAZc+bcb+6SprSbAsRhrZ810XNIBUMa8XWsUDn1uv70vGBWBv
keKZZ7FJ4kuQe0nTONmvQ4EwFekV+IXT5LwdgmPWF0QR7cO8jqeb6psHYauktuzZ
2ul4nMLmLLf/m4DwiCAbEdToBXqRA30KshtgBYYQwL1YkWYgknnAdhHyeu6ybJvv
Y57JYzotjFOlnFhtcGITESEWv+pnj0RJUUrlVwLkJhUOKMwL+sbhw0s5+m27ABEB
AAGJATwEGAEIACYCGwwWIQR5/SuZ88YtLTuFC7yTs1CUdYINXQUCWh+LhAUJA8Jn
uwAKCRCTs1CUdYINXXuvB/9IKK3SLgJ6lOc2Vq73rGYsrDqfjYt5rCDXhjIaFRE7
LYmFJcGL5CHJTae438XtAixa+mu6PYG28eknjZs58Cx/bSj9uS6NiLAPCgyTAtvg
ao6usECOm9Y0xf2+ZcZ9Uji+wsCAFmxRC9je0yUErVyuyQRqzNtdqytnszoTzvb9
iOP8sX/YNrjC83BtZ4Vg3fzAu8qvwbObgSbws5M8TBwIKd4WFTjOtSU6F8aioJ1g
mpfd8KGljHkzC0oG8l8fZiTNYqkIMbfyfPpVwsSqsysLKofifFT+mNs79DJdqNFO
HA2W4WzekYmWWmgK7J8kXHYkxUJA6VpSmNAKwUKqXbNV
=hneF
-----END PGP PUBLIC KEY BLOCK-----
```

このテキストデータで運用すればいいのだ。

たとえば私の公開鍵は[本家サイトで公開している](https://baldanders.info/pubkeys/ "OpenPGP 公開鍵リスト | Baldanders.info")が

```text
$ gpg --fetch-keys https://baldanders.info/pubkeys/spiegel.asc
```

とすれば簡単に鍵束にインポートできる。

<!--
メールの暗号化や署名検証については [Autocrypt] のような仕組みを使えば鍵サーバを経由せずに鍵のやりとりができるらしい（実はよく知らない）。
ちなみに Thunderbird の [Enigmail] は [Autocrypt] に対応している。

[Autocrypt] についてはちゃんと調べていつか記事にする予定である。
-->

## 回避策2： [OpenPGP] 公開鍵サーバ上の公開鍵をいきなりインポートしない

APT のように [OpenPGP] 公開鍵サーバを使った鍵運用が必須の場合でも，いきなり鍵束にインポートするのではなく，事前に公開鍵をチェックして問題がないか調べたほうがよさそうである。

公開鍵をチェックするのであれば [pgpdump] か（手前味噌でナニだが）その [Go 言語]版である [gpgpdump] を使うことをオススメする。

先ほどの

```text
$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
```

であれば

```text
$ wget -O - "http://keyserver.ubuntu.com/pks/lookup?search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF&op=get"
```

などとすればとすれば Armor テキストでダウンロードできる。
汚染されている公開鍵であればかなり巨大になっているだろうから，ある程度の判断はできそうである。

### 【追記 2019-07-06】 [gpgpdump] に HKP アクセスモードを追加した

[gpgpdump] の v0.6.0 をリリースしたが，このバージョンから HKP アクセスモードを追加した。

- [gpgpdump v0.6.0 をリリースした]({{<ref "/release/2019/07/gpgpdump-v0_6_0-is-released.md" >}})

この機能を使えば

```text
$ gpgpdump hkp --keyserver keyserver.ubuntu.com --port 80 0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF -u
Public-Key Packet (tag 6) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2014-08-04T15:35:03Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (58 bytes)
    User ID: Xamarin Public Jenkins (auto-signing) <releng@xamarin.com>
Signature Packet (tag 2) (312 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (34 bytes)
        Signature Creation Time (sub 2): 2014-08-04T15:35:03Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Preferred Symmetric Algorithms (sub 11) (5 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: CAST5 (128 bit key, as per) (sym 3)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
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
        Issuer (sub 16): 0xa6a19b38d3d831ef
    Hash left 2 bytes
        90 e8
    RSA signature value m^d mod n (2047 bits)
Signature Packet (tag 2) (284 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2014-09-04T15:26:28Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0xc90f9cb90e1fad0c
    Hash left 2 bytes
        9c d7
    RSA signature value m^d mod n (2046 bits)
Signature Packet (tag 2) (284 bytes)
    Version: 4 (current)
    Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (6 bytes)
        Signature Creation Time (sub 2): 2016-12-11T01:14:48Z
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x01150a655bbd8102
    Hash left 2 bytes
        7f cf
    RSA signature value m^d mod n (2048 bits)
Public-Subkey Packet (tag 14) (269 bytes)
    Version: 4 (current)
    Public key creation time: 2014-08-04T15:35:03Z
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    RSA public modulus n (2048 bits)
    RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (287 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
    Hash Algorithm: SHA-1 (hash 2)
    Hashed Subpacket (9 bytes)
        Signature Creation Time (sub 2): 2014-08-04T15:35:03Z
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0xa6a19b38d3d831ef
    Hash left 2 bytes
        ac 35
    RSA signature value m^d mod n (2048 bits)
```

といった感じに [OpenPGP] 公開鍵サーバ上の公開鍵を直接検査できる。
これなら最悪でも [gpgpdump] がコケるだけなので [OpenPGP] の鍵束にはダメージがいかないだろう。

## 回避策3： [GnuPG] 2.2.17 以降を使って電子署名のインポートを拒否する（2019-07-10 追記）

[GnuPG] 2.2.17 から公開鍵サーバ上の公開鍵について付帯する電子署名を（自己署名を除いて）捨てることにしたようだ。

- [GnuPG 2.2.17 リリース： 公開鍵サーバ・アクセスに関する過激な変更あり]({{< ref "/release/2019/07/gnupg-2_2_17-is-released.md" >}})

これなら最悪は免れるかな。
公開鍵の管理の仕方が大幅に変わるかもしれないけど。

## 新しい keys.openpgp.org を使う

今後の話になるだろうが，新しい [OpenPGP] 公開鍵サーバが登場したので，公開鍵の運用をそちらに移行する手もある。

- [新しい OpenPGP 鍵サーバが Launch したらしい]({{< ref "/remark/2019/06/launching-a-new-openpgp-key-server.md" >}})

「まだベータ運用だし，しばらくは様子見かなぁ」と思っていたが，ちょっと前倒しして調査したほうがいいかしらん。

<!-- 先ほどの [Autocrypt] の調査も併せて色々調べてみるつもりである。 -->

## ブックマーク

- [Someone Is Spamming and Breaking a Core Component of PGP’s Ecosystem - VICE](https://www.vice.com/en_us/article/8xzj45/someone-is-spamming-and-breaking-a-core-component-of-pgps-ecosystem)
- [SKS Keyserver Network Under Attack · GitHub](https://gist.github.com/rjhansen/67ab921ffb4084c865b3618d6955275f)
- [dkg's blog - OpenPGP Certificate Flooding](https://dkg.fifthhorseman.net/blog/openpgp-certificate-flooding.html)
- [Impact of SKS keyserver poisoning on Gentoo – Gentoo Linux](https://www.gentoo.org/news/2019/07/03/sks-key-poisoning.html)

- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})

[PGP]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Autocrypt]: https://autocrypt.org/ "Autocrypt 1.1.0 documentation"
[Enigmail]: https://addons.thunderbird.net/thunderbird/addon/enigmail/
[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
