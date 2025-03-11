+++
title = "OpenPGP で利用可能なアルゴリズム（RFC 9580 対応版）"
date = "2025-03-11T14:49:14+09:00"
description = "description"
image = "/images/attention/openpgp.png"
tags = ["security", "cryptography", "openpgp"]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

OpenPGP の標準化について2024年7月に [RFC 9580] が発行された。
この記事では拙作 [gpgpdump] の改定作業の目的のため OpenPGP で利用可能なアルゴリズムに絞って列挙していく。

なお [RFC 9580] の発行に伴い，以前の [RFC 4880], [RFC 5581], [RFC 6637] は obsolete になった。
このブログでも前の「[OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< relref "./algorithms-for-openpgp.md" >}})」は obsolete とする。
また各アルゴリズムの横に {{< emoji "チェック" >}} が付いているものは [RFC 9580] で新たに追加されたものである。

## 共通鍵暗号アルゴリズム（Symmetric-Key Algorithms）

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

[RFC 9580] では OCB mode が “MUST implement” となる。

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
<td>{{< emoji "チェック" >}}</td>
<td><a href="https://nvlpubs.nist.gov/nistpubs/SpecialPublications/nist.sp.800-56Ar3.pdf">SP800-56A Revision 3</a></td>
</tr><tr>
<td class='right'>19</td>
<td class="nowrap">ECDSA public key algorithm</td>
<td>{{< emoji "チェック" >}}</td>
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

- RSA (pub 1,2,3), Elgamal (16), DSA (17) は deprecated (非推奨) となった。更に暗号化専用および署名専用の RSA (pub 2,3) と Elgamal (16), DSA (17) の鍵は新たな生成が禁止になった（MUST NOT）。
- X25519 (pub 25) および Ed25519 (pub 27) の実装が MUST になった
- X448 (pub 26) および Ed448 (pub 28) の実装が MUST になった
- EdDSALegacy (pub 22) は [RFC 4880bis] の頃に定義されたが，最終的に Ed25519 (pub 27) に置き換えられた。 pub 22 が残っているのは OID が異なるためである。新たに EdDSA 鍵を生成するなら Ed25519 (pub 27) または Ed448 (pub 28) を選択すべき
- ElGamal (pub 20) は，元々暗号化と署名の両方できるものだったが，脆弱性が見つかったため [OpenPGP] では使用禁止になった[^elg1]

[^elg1]: pub 20 が禁止になった経緯については “[GnuPG's ElGamal signing keys compromised](https://lists.gnupg.org/pipermail/gnupg-users/2003-November/020772.html)” を参照のこと。

余談だが，現在対量子コンピュータ暗号を OpenPGP に組み込む[議論][openpgp-pqc]が行われている。
興味がある方はそちらもどうぞ。

### 楕円曲線（ECC Curves for OpenPGP）

[RFC 6637] および [RFC 4880bis] で追加される楕円曲線暗号（Elliptic Curve Cryptography; ECC）について [OpenPGP] で利用可能な楕円曲線（Elliptic Curve）は以下の通り。
なお「鍵長」の単位は “byte” である。

{{< div-gen >}}
<figure>
<table>
<thead>
<tr><th colspan="2">楕円曲線名</th><th>適用アルゴリズム</th><th>鍵長</th></tr>
</thead>
<tbody>
<tr>
<td class='nowrap'>NIST P-256</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>32</td>
<td rowspan="3"><a href="https://nvlpubs.nist.gov/nistpubs/SpecialPublications/nist.sp.800-56Ar3.pdf">SP800-56A Revision 3</a>, <a href="https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.186-5.pdf">FIPS PUB 186-5</a></td>
</tr><tr>
<td class='nowrap'>NIST P-384</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>48</td>
</tr><tr>
<td class='nowrap'>NIST P-521</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>66</td>
</tr><tr>
<td class='nowrap'>brainpoolP256r1</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDSA, ECDH</td>
<td class='right'>32</td>
<td rowspan="2"><a href="http://www.ecc-brainpool.org/">ECC-Brainpool</a></td>
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
<td>EdDSA</td>
<td class='right'>32</td>
<td>FIPS PUB 186-5, <a href="https://www.rfc-editor.org/rfc/rfc8032">RFC8032</a></td>
</tr><tr>
<td class='nowrap'>Curve25519Legacy</td>
<td>{{< emoji "チェック" >}}</td>
<td>ECDH</td>
<td class='right'>32</td>
<td><a href="https://www.rfc-editor.org/rfc/rfc7748">RFC7748</a></td>
</tr>
</tbody>
</table>
<figcaption>OpenPGP で使用可能な楕円曲線一覧</figcaption>
</figure>
{{< /div-gen >}}

- [RFC 6637] および [RFC 4880bis] では NIST curve P-256 が “MUST implement” となっている。また NIST curve P-521, Ed25519, Curve25519 が “SHOULD implement” となっている
- 具体的な実装例については「[そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな]({{< ref "/openpgp/using-ecc-with-gnupg.md" >}})」で紹介している



## 一方向ハッシュ関数アルゴリズム（Hash Algorithms）

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

### 乱数生成器（Random Number Generator）

### データ圧縮（Compression Algorithms）


## アルゴリズムの選択










[OpenPGP]: https://datatracker.ietf.org/doc/html/rfc9580 "RFC 9580 - OpenPGP"
[RFC 9580]: https://datatracker.ietf.org/doc/html/rfc9580 "RFC 9580 - OpenPGP"
[RFC 4880]: https://www.rfc-editor.org/rfc/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[openpgp-pqc]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-pqc/ "Post-Quantum Cryptography in OpenPGP"
[RFC 5581]: https://www.rfc-editor.org/rfc/rfc5581 "RFC 5581 - The Camellia Cipher in OpenPGP"
[RFC 9106]: https://datatracker.ietf.org/doc/html/rfc9106 "RFC 9106 - Argon2 Memory-Hard Function for Password Hashing and Proof-of-Work Applications"
[RFC 6637]: https://www.rfc-editor.org/rfc/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"
[RFC 8032]: https://www.rfc-editor.org/rfc/rfc8032 "RFC 8032 - Edwards-Curve Digital Signature Algorithm (EdDSA)"
[RFC 4086]: https://www.rfc-editor.org/rfc/rfc4086 "RFC 4086 - Randomness Requirements for Security"
[RFC 1951]: https://www.rfc-editor.org/rfc/rfc1951 "RFC 1951 - DEFLATE Compressed Data Format Specification version 1.3"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/goark/gpgpdump "goark/gpgpdump: OpenPGP packet visualizer"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
