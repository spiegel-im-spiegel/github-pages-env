+++
title = "OpenPGP 公開鍵サーバにおける公開鍵の汚染問題"
date =  "2019-07-05T23:46:33+09:00"
description = "新しい OpenPGP 公開鍵サーバや Autocrypt について調査したほうがいいかしらん。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "pki", "security", "risk", "pgpdump", "gpgpdump" ]
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

- [OpenPGP 鍵管理に関する考察 — OpenPGP の実装 | text.Baldanders.info](https://text.baldanders.info/openpgp/openpgp-key-management/)

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

メールの暗号化や署名検証については [Autocrypt] のような仕組みを使えば鍵サーバを経由せずに鍵のやりとりができるらしい（実はよく知らない）。
ちなみに Thunderbird の [Enigmail] は [Autocrypt] に対応している。

[Autocrypt] についてはちゃんと調べていつか記事にする予定である。

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

とすればテキスト形式でダウンロードできるので，をれを [gpgpdump] に食わせれば

```text
$ wget -O - "http://keyserver.ubuntu.com/pks/lookup?search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF&op=get" | gpgpdump
Public-Key Packet (tag 6) (269 bytes)
	Version: 4 (current)
	Public key creation time: 2014-08-05T00:35:03+09:00
		53 df a8 27
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
	RSA public modulus n (2048 bits)
	RSA public encryption exponent e (17 bits)
User ID Packet (tag 13) (58 bytes)
	User ID: Xamarin Public Jenkins (auto-signing) <releng@xamarin.com>
Signature Packet (tag 2) (284 bytes)
	Version: 4 (current)
	Signiture Type: Generic certification of a User ID and Public-Key packet (0x10)
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
	Hash Algorithm: SHA-1 (hash 2)
	Hashed Subpacket (6 bytes)
		Signature Creation Time (sub 2): 2014-09-05T00:26:28+09:00
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
		Signature Creation Time (sub 2): 2016-12-11T10:14:48+09:00
	Unhashed Subpacket (10 bytes)
		Issuer (sub 16): 0x01150a655bbd8102
	Hash left 2 bytes
		7f cf
	RSA signature value m^d mod n (2048 bits)
Signature Packet (tag 2) (312 bytes)
	Version: 4 (current)
	Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
	Hash Algorithm: SHA-1 (hash 2)
	Hashed Subpacket (34 bytes)
		Signature Creation Time (sub 2): 2014-08-05T00:35:03+09:00
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
Public-Subkey Packet (tag 14) (269 bytes)
	Version: 4 (current)
	Public key creation time: 2014-08-05T00:35:03+09:00
		53 df a8 27
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
	RSA public modulus n (2048 bits)
	RSA public encryption exponent e (17 bits)
Signature Packet (tag 2) (287 bytes)
	Version: 4 (current)
	Signiture Type: Subkey Binding Signature (0x18)
	Public-key Algorithm: RSA (Encrypt or Sign) (pub 1)
	Hash Algorithm: SHA-1 (hash 2)
	Hashed Subpacket (9 bytes)
		Signature Creation Time (sub 2): 2014-08-05T00:35:03+09:00
		Key Flags (sub 27) (1 bytes)
			Flag: This key may be used to encrypt communications.
			Flag: This key may be used to encrypt storage.
	Unhashed Subpacket (10 bytes)
		Issuer (sub 16): 0xa6a19b38d3d831ef
	Hash left 2 bytes
		ac 35
	RSA signature value m^d mod n (2048 bits)
```

と公開鍵の中身をダンプ表示できる。
ちなみに [gpgpdump] は `-j` オプションを使って JSON 形式で出力できるので

```json
$ wget -O - "http://keyserver.ubuntu.com/pks/lookup?search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF&op=get" | gpgpdump -j | jq .
{
  "Packet": [
    {
      "name": "Public-Key Packet (tag 6)",
      "note": "269 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Public key creation time",
          "value": "2014-08-05T00:35:03+09:00",
          "dump": "53 df a8 27"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "RSA public modulus n",
          "note": "2048 bits"
        },
        {
          "name": "RSA public encryption exponent e",
          "note": "17 bits"
        }
      ]
    },
    {
      "name": "User ID Packet (tag 13)",
      "note": "58 bytes",
      "Item": [
        {
          "name": "User ID",
          "value": "Xamarin Public Jenkins (auto-signing) <releng@xamarin.com>"
        }
      ]
    },
    {
      "name": "Signature Packet (tag 2)",
      "note": "284 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Signiture Type",
          "value": "Generic certification of a User ID and Public-Key packet (0x10)"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "Hash Algorithm",
          "value": "SHA-1 (hash 2)"
        },
        {
          "name": "Hashed Subpacket",
          "note": "6 bytes",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2014-09-05T00:26:28+09:00"
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "note": "10 bytes",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0xc90f9cb90e1fad0c"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "9c d7"
        },
        {
          "name": "RSA signature value m^d mod n",
          "note": "2046 bits"
        }
      ]
    },
    {
      "name": "Signature Packet (tag 2)",
      "note": "284 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Signiture Type",
          "value": "Generic certification of a User ID and Public-Key packet (0x10)"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "Hash Algorithm",
          "value": "SHA2-256 (hash 8)"
        },
        {
          "name": "Hashed Subpacket",
          "note": "6 bytes",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2016-12-11T10:14:48+09:00"
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "note": "10 bytes",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0x01150a655bbd8102"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "7f cf"
        },
        {
          "name": "RSA signature value m^d mod n",
          "note": "2048 bits"
        }
      ]
    },
    {
      "name": "Signature Packet (tag 2)",
      "note": "312 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Signiture Type",
          "value": "Positive certification of a User ID and Public-Key packet (0x13)"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "Hash Algorithm",
          "value": "SHA-1 (hash 2)"
        },
        {
          "name": "Hashed Subpacket",
          "note": "34 bytes",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2014-08-05T00:35:03+09:00"
            },
            {
              "name": "Key Flags (sub 27)",
              "note": "1 bytes",
              "Item": [
                {
                  "name": "Flag",
                  "value": "This key may be used to certify other keys."
                },
                {
                  "name": "Flag",
                  "value": "This key may be used to sign data."
                }
              ]
            },
            {
              "name": "Preferred Symmetric Algorithms (sub 11)",
              "note": "5 bytes",
              "Item": [
                {
                  "name": "Symmetric Algorithm",
                  "value": "AES with 256-bit key (sym 9)"
                },
                {
                  "name": "Symmetric Algorithm",
                  "value": "AES with 192-bit key (sym 8)"
                },
                {
                  "name": "Symmetric Algorithm",
                  "value": "AES with 128-bit key (sym 7)"
                },
                {
                  "name": "Symmetric Algorithm",
                  "value": "CAST5 (128 bit key, as per) (sym 3)"
                },
                {
                  "name": "Symmetric Algorithm",
                  "value": "TripleDES (168 bit key derived from 192) (sym 2)"
                }
              ]
            },
            {
              "name": "Preferred Hash Algorithms (sub 21)",
              "note": "5 bytes",
              "Item": [
                {
                  "name": "Hash Algorithm",
                  "value": "SHA2-256 (hash 8)"
                },
                {
                  "name": "Hash Algorithm",
                  "value": "SHA-1 (hash 2)"
                },
                {
                  "name": "Hash Algorithm",
                  "value": "SHA2-384 (hash 9)"
                },
                {
                  "name": "Hash Algorithm",
                  "value": "SHA2-512 (hash 10)"
                },
                {
                  "name": "Hash Algorithm",
                  "value": "SHA2-224 (hash 11)"
                }
              ]
            },
            {
              "name": "Preferred Compression Algorithms (sub 22)",
              "note": "3 bytes",
              "Item": [
                {
                  "name": "Compression Algorithm",
                  "value": "ZLIB <RFC1950> (comp 2)"
                },
                {
                  "name": "Compression Algorithm",
                  "value": "BZip2 (comp 3)"
                },
                {
                  "name": "Compression Algorithm",
                  "value": "ZIP <RFC1951> (comp 1)"
                }
              ]
            },
            {
              "name": "Features (sub 30)",
              "note": "1 bytes",
              "Item": [
                {
                  "name": "Flag",
                  "value": "Modification Detection (packets 18 and 19)"
                }
              ]
            },
            {
              "name": "Key Server Preferences (sub 23)",
              "note": "1 bytes",
              "Item": [
                {
                  "name": "Flag",
                  "value": "No-modify"
                }
              ]
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "note": "10 bytes",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0xa6a19b38d3d831ef"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "90 e8"
        },
        {
          "name": "RSA signature value m^d mod n",
          "note": "2047 bits"
        }
      ]
    },
    {
      "name": "Public-Subkey Packet (tag 14)",
      "note": "269 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Public key creation time",
          "value": "2014-08-05T00:35:03+09:00",
          "dump": "53 df a8 27"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "RSA public modulus n",
          "note": "2048 bits"
        },
        {
          "name": "RSA public encryption exponent e",
          "note": "17 bits"
        }
      ]
    },
    {
      "name": "Signature Packet (tag 2)",
      "note": "287 bytes",
      "Item": [
        {
          "name": "Version",
          "value": "4",
          "note": "current"
        },
        {
          "name": "Signiture Type",
          "value": "Subkey Binding Signature (0x18)"
        },
        {
          "name": "Public-key Algorithm",
          "value": "RSA (Encrypt or Sign) (pub 1)"
        },
        {
          "name": "Hash Algorithm",
          "value": "SHA-1 (hash 2)"
        },
        {
          "name": "Hashed Subpacket",
          "note": "9 bytes",
          "Item": [
            {
              "name": "Signature Creation Time (sub 2)",
              "value": "2014-08-05T00:35:03+09:00"
            },
            {
              "name": "Key Flags (sub 27)",
              "note": "1 bytes",
              "Item": [
                {
                  "name": "Flag",
                  "value": "This key may be used to encrypt communications."
                },
                {
                  "name": "Flag",
                  "value": "This key may be used to encrypt storage."
                }
              ]
            }
          ]
        },
        {
          "name": "Unhashed Subpacket",
          "note": "10 bytes",
          "Item": [
            {
              "name": "Issuer (sub 16)",
              "value": "0xa6a19b38d3d831ef"
            }
          ]
        },
        {
          "name": "Hash left 2 bytes",
          "dump": "ac 35"
        },
        {
          "name": "RSA signature value m^d mod n",
          "note": "2048 bits"
        }
      ]
    }
  ]
}
```

てなこともできる（自画自賛`w`）。

これなら最悪でも [gpgpdump] がコケるだけなので [OpenPGP] の鍵束にはダメージがいかないだろう。

## 新しい keys.openpgp.org を使う

今後の話になるだろうが，新しい [OpenPGP] 公開鍵サーバが登場したので，公開鍵の運用をそちらに移行する手もある。

- [新しい OpenPGP 鍵サーバが Launch したらしい]({{< ref "/remark/2019/06/launching-a-new-openpgp-key-server.md" >}})

「まだベータ運用だし，しばらくは様子見かなぁ」と思っていたが，ちょっと前倒しして調査したほうがいいかしらん。
先ほどの [Autocrypt] の調査も併せて色々調べてみるつもりである。

## ブックマーク

- [Someone Is Spamming and Breaking a Core Component of PGP’s Ecosystem - VICE](https://www.vice.com/en_us/article/8xzj45/someone-is-spamming-and-breaking-a-core-component-of-pgps-ecosystem)
- [SKS Keyserver Network Under Attack · GitHub](https://gist.github.com/rjhansen/67ab921ffb4084c865b3618d6955275f)
- [dkg's blog - OpenPGP Certificate Flooding](https://dkg.fifthhorseman.net/blog/openpgp-certificate-flooding.html)
- [Impact of SKS keyserver poisoning on Gentoo – Gentoo Linux](https://www.gentoo.org/news/2019/07/03/sks-key-poisoning.html)

- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "http://localhost:1313/openpgp/gnupg-cheat-sheet.md" >}})

[PGP]: https://tools.ietf.org/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"
[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Autocrypt]: https://autocrypt.org/ "Autocrypt 1.1.0 documentation"
[Enigmail]: https://addons.thunderbird.net/thunderbird/addon/enigmail/
[pgpdump]: https://github.com/kazu-yamamoto/pgpdump "kazu-yamamoto/pgpdump: A PGP packet visualizer"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/PGP%E2%80%95%E6%9A%97%E5%8F%B7%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%A8%E9%9B%BB%E5%AD%90%E7%BD%B2%E5%90%8D-%E3%82%B7%E3%83%A0%E3%82%BD%E3%83%B3-%E3%82%AC%E3%83%BC%E3%83%95%E3%82%A3%E3%83%B3%E3%82%B1%E3%83%AB/dp/4900900028?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4900900028"><img src="https://images-fe.ssl-images-amazon.com/images/I/5132396FFQL._SL160_.jpg" width="124" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/PGP%E2%80%95%E6%9A%97%E5%8F%B7%E3%83%A1%E3%83%BC%E3%83%AB%E3%81%A8%E9%9B%BB%E5%AD%90%E7%BD%B2%E5%90%8D-%E3%82%B7%E3%83%A0%E3%82%BD%E3%83%B3-%E3%82%AC%E3%83%BC%E3%83%95%E3%82%A3%E3%83%B3%E3%82%B1%E3%83%AB/dp/4900900028?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4900900028">PGP―暗号メールと電子署名</a></dt>
	<dd>シムソン ガーフィンケル</dd>
	<dd>Simson Garfinkel (原著), ユニテック (翻訳)</dd>
    <dd>オライリー・ジャパン 1996-04</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4900900028, EAN: 9784900900028</dd>
    <dd>評価<abbr class="rating fa-sm" title="3">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">良書なのだが，残念ながら内容が古すぎた。 PGP の歴史資料として読むならいいかもしれない。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-10-16">2014-10-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071"><img src="https://images-fe.ssl-images-amazon.com/images/I/51ZRZ62WKCL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E5%8C%96-%E3%83%97%E3%83%A9%E3%82%A4%E3%83%90%E3%82%B7%E3%83%BC%E3%82%92%E6%95%91%E3%81%A3%E3%81%9F%E5%8F%8D%E4%B9%B1%E8%80%85%E3%81%9F%E3%81%A1-%E3%82%B9%E3%83%86%E3%82%A3%E3%83%BC%E3%83%96%E3%83%B3%E3%83%BB%E3%83%AC%E3%83%93%E3%83%BC/dp/4314009071?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4314009071">暗号化 プライバシーを救った反乱者たち</a></dt>
	<dd>スティーブン・レビー</dd>
	<dd>斉藤 隆央 (翻訳)</dd>
    <dd>紀伊國屋書店 2002-02-16</dd>
    <dd>Book 単行本</dd>
    <dd>ASIN: 4314009071, EAN: 9784314009072</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">20世紀末，暗号技術の世界で何があったのか。知りたかったらこちらを読むべし！</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-12-16">2018-12-16</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
