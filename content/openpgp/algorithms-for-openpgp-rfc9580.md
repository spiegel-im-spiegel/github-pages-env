+++
title = "OpenPGP で利用可能なアルゴリズム（RFC 9580 対応版）"
date = "2025-03-12T17:04:54+09:00"
description = "2024年7月に RFC 9580 が発行された。この記事では最新の OpenPGP で利用可能なアルゴリズムを列挙していく。"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "openpgp"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

OpenPGP の標準化について2024年7月に [RFC 9580] が発行された。
この記事では拙作 [gpgpdump] の改定作業の目的のため OpenPGP で利用可能なアルゴリズムを列挙していく。

なお [RFC 9580] の発行に伴い，以前の [RFC 4880], [RFC 5581], [RFC 6637] は obsolete になった。
このブログでも前の「[OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< relref "./algorithms-for-openpgp.md" >}})」は obsolete とする。
また各アルゴリズムの横に {{< emoji "チェック" >}} が付いているものは [RFC 9580] で新たに追加されたものである。

## パケットバージョン

アルゴリズムとは直接関係ないが [OpenPGP] では暗号文や鍵や署名といったデータのかたまりを「パケット」と呼んでいる。
パケットは PGP/[OpenPGP] によって幾つかのバージョンがある。

現在有効なパケットバージョンの組み合わせは以下の通り。

{{< fig-gen type="markdown" title="OpenPGP 暗号化メッセージのバージョン" >}}
| Version of <br>Encrypted Data <br>Payload | Version of <br>Preceding <br>Symmetric Key ESK <br>(If Any) | Version of <br>Preceding Public <br>Key ESK (If Any) | Generate? |
| :---: | :---: | :---: | :---: |
| SED | - | v2 PKESK ([RFC 2440]) | No |
| SED | v4 SKESK | v3 PKESK | No |
| v1 SEIPD | v4 SKESK | v3 PKESK | Yes |
| v2 SEIPD  {{% emoji "チェック" %}} | v4 SKESK | v3 PKESK | Yes |

[RFC 2440]: https://datatracker.ietf.org/doc/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
{{< /fig-gen >}}

v2 SEIPD は [RFC 9580] で加わったもので認証付き暗号（AEAD）をサポートしている。

[RFC 9580] では “Generate?” が Yes のパケット組み合わせのみ生成が許容される（MUST）。
それ以外は後方互換性のために残される。

{{< fig-gen type="markdown" title="OpenPGP 鍵および署名のバージョン" >}}
| Signing Key <br>Version | Signature Packet <br>Version | One-Pass Signature <br>Packet Version | Generate? |
| :---: | :---: | :---: | :---: |
| v3 | v3 | v3 | No |
| v4 | v3 | v3 | No |
| v4 | v4 | v3 | Yes |
| v6 {{% emoji "チェック" %}} | v6 {{% emoji "チェック" %}} | v6 {{% emoji "チェック" %}} | Yes |
{{< /fig-gen >}}

ここにない v1 や v2 は最初期の PGP ([RFC 1991]) の頃のもの。
さらに，ここにない v5 は [RFC 4880bis] の議論の頃にあったもので [RFC 9580] ではなくなってしまったようだ。
おそらく [GnuPG] には残ってると思うけど... 困ったね（笑）

v3 は [RFC 2440] で， v4 は [RFC 4880] で登場したもの。
さらに v6 は [RFC 9580] で新たに加わったものである。

[RFC 9580] では “Generate?” が Yes のパケット組み合わせのみ生成が許容される（MUST）。
それ以外は後方互換性のために残される。

## 共通鍵暗号アルゴリズム（Symmetric-Key Algorithms）

[OpenPGP] で利用可能な共通鍵暗号は以下の通り。
[RFC 9580] で新たに追加されたアルゴリズムはない。
なお「鍵長」項目の括弧内はブロック長を指す。
いずれも単位は “bit” である。

{{< div-gen >}}
<figure lang="en">
<table>
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>鍵長</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>0</td>
<td class="nowrap" colspan="3">Plaintext or unencrypted data</td>
</tr><tr>
<td class='right'>1</td>
<td class="nowrap">IDEA</td>
<td class='right nowrap'>128 (64)</td>
<td><a href="https://link.springer.com/chapter/10.1007%2F978-3-642-29011-4_24">Narrow-Bicliques: Cryptanalysis of Full IDEA</a></td>
</tr><tr>
<td class='right'>2</td>
<td>TripleDES (or DES-EDE) with 168-bit key derived from 192</td>
<td class='right nowrap'>168 (64)</td>
<td><a href="https://doi.org/10.6028/NIST.SP.800-67r2">SP800-67 Rev.2</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">CAST5 with 128-bit key</td>
<td class='right nowrap'>128 (64)</td>
<td><a href="http://tools.ietf.org/html/rfc2144">RFC2144</a></td>
</tr><tr>
<td class='right'>4</td>
<td>Blowfish with 128-bit key, 16 rounds</td>
<td class='right nowrap'>128 (64)</td>
<td><a href="http://www.schneier.com/paper-blowfish-fse.html">Description of a New Variable-Length Key, 64-Bit Block Cipher (Blowfish)</a></td>
</tr><tr>
<td class='right'>5,6</td>
<td colspan="3">(Reserved)</td>
</tr><tr>
<td class='right'>7</td>
<td class="nowrap">AES with 128-bit key</td>
<td class='right nowrap'>128 (128)</td>
<td rowspan="3"><a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197-upd1.pdf">FIPS PUB 197</a></td>
</tr><tr>
<td class='right'>8</td>
<td class="nowrap">AES with 192-bit key</td>
<td class='right nowrap'>192 (128)</td>
</tr><tr>
<td class='right'>9</td>
<td class="nowrap">AES with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
</tr><tr>
<td class='right'>10</td>
<td class="nowrap">Twofish with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
<td></td>
</tr><tr>
<td class='right'>11</td>
<td class="nowrap">Camellia with 128-bit key</td>
<td class='right nowrap'>128 (128)</td>
<td rowspan="3"><a href="http://tools.ietf.org/html/rfc3713">RFC3713</a></td>
</tr><tr>
<td class='right'>12</td>
<td class="nowrap">Camellia with 192-bit key</td>
<td class='right nowrap'>192 (128)</td>
</tr><tr>
<td class='right'>13</td>
<td class="nowrap">Camellia with 256-bit key</td>
<td class='right nowrap'>256 (128)</td>
</tr><tr>
<td class='right nowrap'>100-110</td>
<td colspan="3">Private/Experimental algorithm</td>
</tr><tr>
<td class='right nowrap'>253-255</td>
<td colspan="3">(Reserved to avoid collision with Secret Key Encryption)</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な共通鍵暗号アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので “sym 1” のように表記する。

- AES-128 (sym 7) が “MUST implement” で AES-256 (sym 9) が “SHOULD implement” となった
- 後方互換性のために IDEA, TripleDES, CAST5 による復号を行ってもよい（MAY）が機密性の漏洩が疑われる非推奨のアルゴリズムであることを警告すべき（SHOULD）

ちなみに sym 11 から sym 13 の [Camellia 暗号は日本製](https://baldanders.info/blog/000451/ "The Camellia Cipher in OpenPGP — Baldanders.info")である。

[RFC 9580] では共通鍵暗号を使った暗号化について，従来からある CFB モード[^cfb1] に加えて認証付き暗号の暗号モードを使用できる。

[^cfb1]: 厳密には CFB モードの変形である。

### 認証付き暗号の暗号モード（AEAD Algorithms）

[RFC 9580] で追加された認証付き暗号（Authenticated Encryption with Associated Data; AEAD）について利用可能な暗号モードは以下の通り。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th>ID</th><th colspan="2">暗号モード</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">EAX</td>
<td>{{< emoji "チェック" >}}</td>
<td><a href="https://seclab.cs.ucdavis.edu/papers/eax.pdf">EAX: A Conventional Authenticated-Encryption Mode</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">OCB</td>
<td>{{< emoji "チェック" >}}</td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc7253">RFC 7253</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">GCM</td>
<td>{{< emoji "チェック" >}}</td>
<td><a href="https://nvlpubs.nist.gov/nistpubs/legacy/sp/nistspecialpublication800-38d.pdf">SP800-38D</a></td>
</tr><tr>
<td class='right nowrap'>100-110</td>
<td colspan="3">Private/Experimental algorithm</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な認証付き暗号アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

[RFC 9580] では OCB が “MUST implement” となる。

## 公開鍵暗号・署名アルゴリズム（Public-Key Algorithms）

[OpenPGP] で利用可能な公開鍵暗号・署名は以下の通り。

{{< div-gen >}}
<figure lang="en">
</style>
<table>
<thead>
<tr><th>ID</th><th colspan="2">アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">RSA (Encrypt or Sign)</td>
<td></td>
<td rowspan="3"><a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">RSA Encrypt-Only</td>
<td></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">RSA Sign-Only</td>
<td></td>
</tr><tr>
<td class='right nowrap'>4-15</td>
<td colspan="3">(Reserved)</td>
</tr><tr>
<td class='right'>16</td>
<td class="nowrap">Elgamal<br>(Encrypt-Only)</td>
<td></td>
<td><a href="https://doi.org/10.1109/TIT.1985.1057074">A public key cryptosystem and a signature scheme based on discrete logarithms</a></td>
</tr><tr>
<td class='right'>17</td>
<td class="nowrap">DSA</td>
<td></td>
<td><a href="hhttps://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-4.pdf">FIPS PUB 186-4</a></td>
</tr><tr>
<td class='right'>18</td>
<td class="nowrap">ECDH public key algorithm</td>
<td></td>
<td><a href="https://nvlpubs.nist.gov/nistpubs/SpecialPublications/nist.sp.800-56Ar3.pdf">SP800-56A Revision 3</a></td>
</tr><tr>
<td class='right'>19</td>
<td class="nowrap">ECDSA public key algorithm</td>
<td></td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc6090">RFC 6090</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a>, <a href="https://www.secg.org/sec1-v2.pdf">SEC 1: Elliptic Curve Cryptography</a></td>
</tr><tr>
<td class='right'>20</td>
<td colspan="3">(Reserved; formerly Elgamal Encrypt or Sign)</td>
</tr><tr>
<td class='right'>21</td>
<td colspan="3">(Reserved for Diffie-Hellman (X9.42, as defined for IETF-S/MIME))</td>
</tr><tr>
<td class='right'>22</td>
<td class="nowrap">EdDSALegacy (deprecated)</td>
<td>{{< emoji "チェック" >}}</td>
<td>(Ed25519 を参照のこと)</td>
</tr><tr>
<td class='right'>23</td>
<td colspan="3">(Reserved for AEDH)</td>
</tr><tr>
<td class='right'>24</td>
<td colspan="3">(Reserved for AEDSA)</td>
</tr><tr>
<td class='right'>25</td>
<td class="nowrap">X25519</td>
<td>{{< emoji "チェック" >}}</td>
<td rowspan="2"><a href="https://datatracker.ietf.org/doc/html/rfc7748">RFC 7748</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='right'>26</td>
<td class="nowrap">X448</td>
<td>{{< emoji "チェック" >}}</td>
</tr><tr>
<td class='right'>27</td>
<td class="nowrap">Ed25519</td>
<td>{{< emoji "チェック" >}}</td>
<td rowspan="2"><a href="https://datatracker.ietf.org/doc/html/rfc8032">RFC 8032</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='right'>28</td>
<td class="nowrap">Ed448</td>
<td>{{< emoji "チェック" >}}</td>
</tr><tr>
<td class='right nowrap'>100-110</td>
<td colspan="3">Private/Experimental algorithm</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な公開鍵暗号・署名アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので “pub 1” のように表記する。

- RSA (pub 1,2,3), Elgamal (16), DSA (17) は非推奨 (deprecated) となった。更に暗号化専用および署名専用の RSA (pub 2,3) と Elgamal (16), DSA (17) の鍵は新たな生成が禁止になった（MUST NOT）。ただし [RFC 4880] では Elgamal (16) と DSA (17) の実装が MUST なので，後方互換性を確保するのであれば，これらを実装する必要がある
- X25519 (pub 25) および Ed25519 (pub 27) の実装が MUST になった
- X448 (pub 26) および Ed448 (pub 28) の実装が MUST になった
- EdDSALegacy (pub 22) は [RFC 4880bis] の頃に定義されたが，最終的に Ed25519 (pub 27) に置き換えられた。おそらく v4 鍵および署名パケットに対応するために残されているが非推奨（deprecated）になっている。新たに EdDSA 鍵を生成するなら Ed25519 (pub 27) または Ed448 (pub 28) を選択すべき
- ElGamal Encrypt or Sign (pub 20) は，元々暗号化と署名の両方できるものだったが，脆弱性が見つかったため [OpenPGP] では使用禁止になった[^elg1]

[^elg1]: pub 20 が禁止になった経緯については “[GnuPG's ElGamal signing keys compromised](https://lists.gnupg.org/pipermail/gnupg-users/2003-November/020772.html)” を参照のこと。

余談だが，現在対量子コンピュータ暗号を OpenPGP に組み込む[議論][openpgp-pqc]が行われている。
興味がある方はそちらもどうぞ。

### 楕円曲線（ECC Curves for OpenPGP）

前節で挙げた楕円曲線暗号（Elliptic Curve Cryptography; ECC）アルゴリズムに対して [RFC 9580] で利用可能な楕円曲線（Elliptic Curve）は以下の通り。
なお「鍵長」の単位はオクテット（byte）である。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th colspan="2">楕円曲線名</th><th>適用アルゴリズム</th><th>鍵長</th></tr>
</thead>
<tbody>
<tr>
<td class='nowrap'>NIST P-256</td>
<td></td>
<td>ECDSA, ECDH</td>
<td class='right'>32</td>
<td rowspan="3"><a href="https://nvlpubs.nist.gov/nistpubs/SpecialPublications/nist.sp.800-56Ar3.pdf">SP800-56A Revision 3</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='nowrap'>NIST P-384</td>
<td></td>
<td>ECDSA, ECDH</td>
<td class='right'>48</td>
</tr><tr>
<td class='nowrap'>NIST P-521</td>
<td></td>
<td>ECDSA, ECDH</td>
<td class='right'>66</td>
</tr><tr>
<td class='nowrap'>brainpoolP256r1</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>32</td>
<td rowspan="3"><a href="https://datatracker.ietf.org/doc/html/rfc5639">RFC 5639</a></td>
</tr><tr>
<td class='nowrap'>brainpoolP384r1</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>48</td>
</tr><tr>
<td class='nowrap'>brainpoolP512r1</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>64</td>
</tr><tr>
<td class='nowrap'>Ed25519Legacy</td>
<td>{{< emoji "チェック" >}}</td>
<td>EdDSALegacy (EdDSA)</td>
<td class='right'>32</td>
<td rowspan="3"><a href="https://datatracker.ietf.org/doc/html/rfc8032">RFC8032</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='nowrap'>Ed25519</td>
<td>{{< emoji "チェック" >}}</td>
<td>Ed25519 (EdDSA)</td>
<td class='right'>32</td>
</tr><tr>
<td class='nowrap'>Ed448</td>
<td>{{< emoji "チェック" >}}</td>
<td>Ed448 (EdDSA)</td>
<td class='right'>57</td>
</tr><tr>
<td class='nowrap'>Curve25519Legacy</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDH</td>
<td class='right'>32</td>
<td rowspan="3"><a href="https://datatracker.ietf.org/doc/html/rfc7748">RFC7748</a></td>
</tr><tr>
<td class='nowrap'>X25519</td>
<td>{{< emoji "チェック" >}}</td>
<td>X25519 (ECDH)</td>
<td class='right'>32</td>
</tr><tr>
<td class='nowrap'>X448</td>
<td>{{< emoji "チェック" >}}</td>
<td>X448 (ECDH)</td>
<td class='right'>56</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な楕円曲線一覧</figcaption>
</figure>
{{< /div-gen >}}

- Ed25519Legacy と Ed25519 は同じ楕円曲線で Ed448 と併せて [RFC 8032] EdDSA アルゴリズムで用いる（ECDSA では使えない）
- Curve25519Legacy と X25519 は同じ楕円曲線で X448 と併せて [RFC 7748] で規定されている。これらの楕円曲線を使った ECDH 鍵は同じサイズの EdDSA 鍵と組み合わせて使う
- Ed25519, Ed448, X25519, X448 以外の楕円曲線は Curve OID で管理される。 Ed25519, Ed448, X25519, X448 については同名のアルゴリズム ID と紐付いているため Curve OID による管理外となる
- Ed25519Legacy は同名のアルゴリズム ID で用いる。また v4 鍵および署名パケットでのみ使用可能。 v6 パケットの場合は Ed25519 を用いる
- Curve25519Legacy は ECDH アルゴリズムかつ v4 鍵および署名パケットでのみ使用可能。 v6 パケットの場合は X25519 を用いる

具体的な実装例については「[そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})」で紹介している

### ECDH 鍵導出のパラメータ（ECDH Parameters）

ECDH (pub 18) では鍵導出のためのハッシュアルゴリズムと鍵カプセル化のための共通鍵暗号アルゴリズムをパラメータとして持つ。
各アルゴリズムの組み合わせは以下の通り。

{{< fig-gen type="markdown" title="ECDH KDF/KEK パラメータ" >}}
| Curve | Hash Algorithm | symmetric Algorithm |
| :--- | :--- | :--- |
| NIST P-256 | SHA2-256 | AES-128 |
| NIST P-384 | SHA2-384 | AES-192 |
| NIST P-521 | SHA2-512 | AES-256 |
| brainpoolP256r1 | SHA2-256 | AES-128 |
| brainpoolP384r1 | SHA2-384 | AES-192 |
| brainpoolP512r1 | SHA2-512 | AES-256 |
| Curve25519Legacy | SHA2-256 | AES-128 |
{{< /fig-gen >}}

- v6 鍵パケットの場合，上のアルゴリズムの組み合わせを使用しなければならない（MUST）
- v4 鍵パケットの場合，上のアルゴリズムの組み合わせを使用すべき（SHOULD）

## 一方向ハッシュ関数アルゴリズム（Hash Algorithms）

[OpenPGP] で利用可能なハッシュ関数は以下の通り。
[RFC 9580] で新たに追加されたアルゴリズムはない。

{{< div-gen >}}
<figure lang="en">
<table>
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>1</td>
<td class="nowrap">MD5</td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc1321">RFC1321</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">SHA-1</td>
<td><a href="https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.180-4.pdf">FIPS PUB 180-4</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">RIPE-MD/160</td>
<td><a href="http://homes.esat.kuleuven.be/~bosselae/ripemd160.html">The hash function RIPEMD-160</a></td>
</tr><tr>
<td class='right'>4-7</td>
<td colspan="2">(Reserved)</td>
</tr><tr>
<td class='right'>8</td>
<td class="nowrap">SHA2-256</td>
<td rowspan="4"><a href="https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.180-4.pdf">FIPS PUB 180-4</a></td>
</tr><tr>
<td class='right'>9</td>
<td class="nowrap">SHA2-384</td>
</tr><tr>
<td class='right'>10</td>
<td class="nowrap">SHA2-512</td>
</tr><tr>
<td class='right'>11</td>
<td class="nowrap">SHA2-224</td>
</tr><tr>
<td class='right'>12</td>
<td class="nowrap">SHA3-256</td>
<td rowspan="3"><a href="https://nvlpubs.nist.gov/nistpubs/fips/nist.fips.202.pdf">FIPS PUB 202</a></td>
</tr><tr>
<td class='right'>13</td>
<td>(Reserved)</td>
</tr><tr>
<td class='right'>14</td>
<td class="nowrap">SHA3-512</td>
</tr><tr>
<td class='right'>100-110</td>
<td colspan="2">Private/Experimental algorithm</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な一方向ハッシュ関数一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので “hash 1” のように表記する。

- SHA2-256 (hash 8) が “MUST implement” となった。また SHA2-384 (hash 9), SHA2-512 (hash 10) の実装が推奨される（SHOULD）
- v4 鍵および署名パケットのの鍵指紋や MDC (Modification Detection Code) を除き SHA-1 (hash 2) を必要とするメッセージを作成してはならない（SHOULD NOT）
- MD5 (hash 1), SHA-1 (hash 2), RIPE-MD/160 (hash 3) を用いて署名を作成してはならない（MUST NOT）
- ECDH KDF および S2K KDF のハッシュ関数として MD5 (hash 1), SHA-1 (hash 2), RIPE-MD/160 (hash 3) を用いてパケットを生成してはならない（MUST NOT）
- v6 以降のパケットで S2K KDF のハッシュ関数として MD5 (hash 1), SHA-1 (hash 2), RIPE-MD/160 (hash 3) を用いて秘密鍵を復号してはならない（MUST NOT）
- MD5 (hash 1), SHA-1 (hash 2), RIPE-MD/160 (hash 3) に依存する最近の署名を検証してはならない（MUST NOT）
- MD5 (hash 1), SHA-1 (hash 2), RIPE-MD/160 (hash 3) に依存する古い署名を検証してはならない（SHOULD NOT）
  - 署名の作成日が使用されたアルゴリズムの既知の脆弱性が発見された日より前で，メッセージが常に安全な管理下にあったと確信できる場合は除く

### 鍵指紋作成で使われるハッシュ関数アルゴリズム（Fingerprints）

鍵指紋作成で使われるハッシュ関数アルゴリズムは以下の通り。

{{< fig-gen type="markdown" title="鍵指紋作成で使われるハッシュ関数アルゴリズム" >}}
| Key Version | Fingerprint Algorithm |
| :---: | :--- |
| v3 | MD5 |
| v4 | SHA1 |
| v6 {{% emoji "チェック" %}} | SHA256 |
{{< /fig-gen >}}

## その他のアルゴリズム

### S2K (String-to-Key)

S2K はパスフレーズからセッション鍵を生成するためのハッシュ化の手順である。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th>ID</th><th colspan="2">S2K タイプ</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>0</td>
<td>Simple S2K</td>
<td></td>
<td></td>
</tr><tr>
<td class='right'>1</td>
<td>Salted S2K</td>
<td></td>
<td></td>
</tr><tr>
<td class='right'>2</td>
<td>Reserved value</td>
<td></td>
<td></td>
</tr><tr>
<td class='right'>3</td>
<td>Iterated and Salted S2K</td>
<td></td>
<td></td>
</tr><tr>
<td class='right'>4</td>
<td>Argon2</td>
<td>{{< emoji "チェック" >}}</td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc9106">RFC 9106</a></td>
</tr><tr>
<td class='right'>100-110</td>
<td>Private/Experimental S2K</td>
<td></td>
<td></td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な S2K アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので “s2k 1” のように表記する。
Simple S2K (s2k 0) と Reserved value (s2k 2) は下位互換のために残されているもので，新たに暗号データや電子署名を作成する際に使用すべきではない。

[OpenPGP] では，パスフレーズ自体はいかなる形（ハッシュ値を含む）でも保存しない。
このため，パスフレーズを紛失してしまった場合は復元できない[^pp1]。

[^pp1]: ただし [OpenPGP] では試行回数によるロックアウトは定義されないため，無限にパスフレーズ解読を試みることができる。

### データ圧縮（Compression Algorithms）

暗号化メッセージや電子署名を圧縮するためのアルゴリズムである。
[RFC 9580] で新たに追加されたアルゴリズムはない。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th>ID</th><th>アルゴリズム</th><th>参考文献</th></tr>
</thead>
<tbody>
<tr>
<td class='right'>0</td>
<td class="nowrap">Uncompressed</td>
<td>&nbsp;</td>
</tr><tr>
<td class='right'>1</td>
<td class="nowrap">ZIP</td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc1951">RFC1951</a></td>
</tr><tr>
<td class='right'>2</td>
<td class="nowrap">ZLIB</td>
<td><a href="https://datatracker.ietf.org/doc/html/rfc1950">RFC1950</a></td>
</tr><tr>
<td class='right'>3</td>
<td class="nowrap">BZip2</td>
<td><a href="http://www.bzip.org/">bzip2</a></td>
</tr><tr>
<td class='right'>100-110</td>
<td colspan="2">Private/Experimental algorithm</td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能なデータ圧縮アルゴリズム一覧</figcaption>
</figure>
{{< /div-gen >}}

ID は [OpenPGP] で定義されるもので “comp 1” のように表記する。

- 非圧縮（comp 0）は “MUST implement” である
- ZLIB (comp 2) は “SHOULD implement” となっていて， ZIP (comp 1) については “SHOULD be able to decompress using” と復号時の後方互換性のみ確保されていればいいようだ

### 乱数生成器（Random Number Generator）

基本的に `getrandom()` システムコールのような OS が標準で提供している暗号学的に安全な疑似乱数生成器（CSPRNG）を使うべき，とある。
もしそれらが使えない（信用できない）場合でも，新しいものを自作するのではなく [RFC 4086] を参照して実装すべきと書かれている。

## ブックマーク

- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})
- [NIST FIPS 186-5 および SP 800-186 正式版がリリースされた]({{< ref "/remark/2023/02/nist-fips-186-5.md" >}})


[OpenPGP]: https://datatracker.ietf.org/doc/html/rfc9580 "RFC 9580 - OpenPGP"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"

[openpgp-pqc]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/ "Post-Quantum Cryptography in OpenPGP"
[RFC 9580]: https://datatracker.ietf.org/doc/html/rfc9580 "RFC 9580 - OpenPGP"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[RFC 6637]: https://datatracker.ietf.org/doc/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 5581]: https://datatracker.ietf.org/doc/html/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 4880]: https://datatracker.ietf.org/doc/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 2440]: https://datatracker.ietf.org/doc/html/rfc2440 "RFC 2440 - OpenPGP Message Format"
[RFC 1991]: https://datatracker.ietf.org/doc/html/rfc1991 "RFC 1991 - PGP Message Exchange Formats"

[RFC 1951]: https://datatracker.ietf.org/doc/html/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[RFC 4086]: https://datatracker.ietf.org/doc/html/rfc4086 "RFC 4086 - Randomness Requirements for Security"
[RFC 8032]: https://datatracker.ietf.org/doc/html/rfc8032 "RFC 8032 - Edwards-Curve Digital Signature Algorithm (EdDSA)"
[RFC 5639]: https://datatracker.ietf.org/doc/html/rfc5639 "RFC 5639 - Elliptic Curve Cryptography (ECC) Brainpool Standard Curves and Curve Generation"
[RFC 7748]: https://datatracker.ietf.org/doc/html/rfc7748 "RFC 7748 - Elliptic Curves for Security"
[RFC 9106]: https://datatracker.ietf.org/doc/html/rfc9106 "RFC 9106 - Argon2 Memory-Hard Function for Password Hashing and Proof-of-Work Applications"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
