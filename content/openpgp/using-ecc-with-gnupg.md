+++
title = "そろそろ GnuPG でも ECC を標準で使うのがいいんじゃないかな"
date =  "2017-12-02T16:20:26+09:00"
update =  "2018-01-28T18:00:19+09:00"
description = "手っ取り早く ECC 鍵を作りたいなら --quick-generate-key コマンドでアルゴリズムに future-default を指定すればよい。"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "gnupg",
  "ecc",
]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = true
  mermaidjs = false
+++

最初に言っておくと [OpenPGP] では（秘密鍵の漏洩や暗号アルゴリズム等の危殆化がない限り）永続的に使われるのがよい鍵とされている[^km1]。
なので，無理に新しい鍵に切り替える必要はないのだが，もし [GnuPG] で新しい鍵を作るのであれば ECC (Elliptic Curve Cryptography; 楕円曲線暗号) を選択するのがいいんじゃないかな？ という提案である。

[^km1]: この辺については「[OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})」を参照のこと。

## 手っ取り早く ECC 鍵を作る

手っ取り早く ECC 鍵を作りたいなら `--quick-generate-key` コマンドでアルゴリズムに `future-default` を指定すればよい。
（以下は有効期限を2年で指定する場合。無期限でいいなら `0` を指定する）

```text
$ gpg --quick-gen-key "Alice <alice@example.com>" future-default - 2y
たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。

たくさんのランダム・バイトの生成が必要です。キーボードを打つ、マウスを動か
す、ディスクにアクセスするなどの他の操作を素数生成の間に行うことで、乱数生
成器に十分なエントロピーを供給する機会を与えることができます。
gpg: 失効証明書を 'C:/Users/slice/AppData/Roaming/gnupg/openpgp-revocs.d\F79B411717B86902177EA08200693613E2B68F5A.rev' に保管しました。
公開鍵と秘密鍵を作成し、署名しました。

pub   ed25519 2017-12-02 [SC] [有効期限: 2019-12-02]
      F79B411717B86902177EA08200693613E2B68F5A
uid                      Alice <alice@example.com>
sub   cv25519 2017-12-02 [E]
```

作成した鍵を私作の [gpgpdump] [^gpd] で見てみるとこんな感じになっている。

[^gpd]: [gpgpdump] v0.3.4 以降なら多分大丈夫。

{{< highlight text "hl_lines=2-9 52-62" >}}
$ gpg -a --export alice | gpgpdump
Public-Key Packet (tag 6) (51 bytes)
    Version: 4 (current)
    Public key creation time: 2017-12-02T10:37:17+09:00
        5a 22 03 cd
    Public-key Algorithm: EdDSA (pub 22)
    ECC Curve OID: ed25519
        2b 06 01 04 01 da 47 0f 01
    EdDSA EC point (04 || X || Y) (263 bits)
User ID Packet (tag 13) (25 bytes)
    User ID: Alice <alice@example.com>
Signature Packet (tag 2) (150 bytes)
    Version: 4 (current)
    Signiture Type: Positive certification of a User ID and Public-Key packet (0x13)
    Public-key Algorithm: EdDSA (pub 22)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (62 bytes)
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                f7 9b 41 17 17 b8 69 02 17 7e a0 82 00 69 36 13 e2 b6 8f 5a
        Signature Creation Time (sub 2): 2017-12-02T10:37:17+09:00
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to certify other keys.
            Flag: This key may be used to sign data.
        Key Expiration Time (sub 9): 730 days after (2019-12-02T10:37:17+09:00)
        Preferred Symmetric Algorithms (sub 11) (4 bytes)
            Symmetric Algorithm: AES with 256-bit key (sym 9)
            Symmetric Algorithm: AES with 192-bit key (sym 8)
            Symmetric Algorithm: AES with 128-bit key (sym 7)
            Symmetric Algorithm: TripleDES (168 bit key derived from 192) (sym 2)
        Preferred Hash Algorithms (sub 21) (5 bytes)
            Hash Algorithm: SHA2-256 (hash 8)
            Hash Algorithm: SHA2-384 (hash 9)
            Hash Algorithm: SHA2-512 (hash 10)
            Hash Algorithm: SHA2-224 (hash 11)
            Hash Algorithm: SHA-1 (hash 2)
        Preferred Compression Algorithms (sub 22) (3 bytes)
            Compression Algorithm: ZLIB <RFC1950> (comp 2)
            Compression Algorithm: BZip2 (comp 3)
            Compression Algorithm: ZIP <RFC1951> (comp 1)
        Features (sub 30) (1 bytes)
            Flag: Modification Detection (packets 18 and 19)
        Key Server Preferences (sub 23) (1 bytes)
            Flag: No-modify
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x00693613e2b68f5a
    Hash left 2 bytes
        7e b5
    EdDSA compressed value r (253 bits)
    EdDSA compressed value s (251 bits)
Public-Subkey Packet (tag 14) (56 bytes)
    Version: 4 (current)
    Public key creation time: 2017-12-02T10:37:17+09:00
        5a 22 03 cd
    Public-key Algorithm: ECDH public key algorithm (pub 18)
    ECC Curve OID: cv25519
        2b 06 01 04 01 97 55 01 05 01
    ECDH EC point (04 || X || Y) (263 bits)
    KDF parameters (3 bytes)
        Hash Algorithm: SHA2-256 (hash 8)
        Symmetric Algorithm: AES with 128-bit key (sym 7)
Signature Packet (tag 2) (120 bytes)
    Version: 4 (current)
    Signiture Type: Subkey Binding Signature (0x18)
    Public-key Algorithm: EdDSA (pub 22)
    Hash Algorithm: SHA2-256 (hash 8)
    Hashed Subpacket (32 bytes)
        Issuer Fingerprint (sub 33) (21 bytes)
            Version: 4 (need 20 octets length)
            Fingerprint (20 bytes)
                f7 9b 41 17 17 b8 69 02 17 7e a0 82 00 69 36 13 e2 b6 8f 5a
        Signature Creation Time (sub 2): 2017-12-02T10:37:17+09:00
        Key Flags (sub 27) (1 bytes)
            Flag: This key may be used to encrypt communications.
            Flag: This key may be used to encrypt storage.
    Unhashed Subpacket (10 bytes)
        Issuer (sub 16): 0x00693613e2b68f5a
    Hash left 2 bytes
        29 27
    EdDSA compressed value r (254 bits)
    EdDSA compressed value s (254 bits)
{{< /highlight >}}

長々とした出力で申し訳ないが `Public-Key Packet` と `Public-Subkey Packet` の項目に注目していただくと，電子署名用に EdDSA (ed25519)，暗号用に ECDH (cv25519) が使われているのが分かる。
鍵長はいずれも 256bits である。

通常はこれで問題ない。

## 楕円曲線と鍵長

ここでちょっと横道に逸れて鍵長の話をしておく。

たとえば RSA の（平文 $M$ から暗号文 $C$ への）暗号化は以下の式で表される。

{{< div-box >}}
\[
    C = M^e \bmod n
\]
{{< /div-box >}}

$(e, n)$ のセットが公開鍵で， $n$ のサイズが鍵長に相当する。
実際の計算はともかく，感覚としては分かりやすいよね。

ECC の場合は暗号化の前に，まずベースとなる楕円曲線の（素数 $q$ で決められる）素体（prime fields）を決めなければならない[^pf1]。
これは以下の式で表される。

[^pf1]: 素体ではなく「標数2の体（binary fields）」を使う場合もあるが，ここでは割愛する。

{{< div-box >}}
\[
    E : y^2 \equiv x^3 + ax + b \pmod{q}
\]
{{< /div-box >}}

この素体上の有理点の数（位数）を $\\#E$ とした時の $\\#E$ のサイズが鍵長に相当する。
$(a,b,q)$ といったパラメータを眺めただけでは鍵長が分からないというのが面倒臭い感じである[^ecc1]。

[^ecc1]: 楕円曲線と楕円曲線暗号については結城浩さんの『[暗号技術入門 第3版](http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/)』の付録に比較的分かりやすい解説が載っている。同じく結城浩さんの『[数学ガール ガロア理論](http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/)』が何となく分かれば楕円曲線についても何となく分かると思う。大丈夫。私も何となくしか分かっていない（笑）

そこで ECC の場合は楕円曲線を表すパラメータのセットが標準化されていて，そのセットを識別する ECC Curve OID (Object IDentifier) も決められている。
先程書いた [gpgpdump] 出力結果の "`ECC Curve OID`” の項目がそれに該当する。

つまり ECC Curve OID から楕円曲線の種類を特定し，そこから鍵長も分かる，というわけだ。

## [GnuPG] で利用できる楕円曲線

そこで， [GnuPG] で利用できる楕円曲線をリストアップしておく。

表の左端列の「GnuPG」は [GnuPG] で表示される楕円曲線名を示す。
楕円曲線名に鍵長が含まれている（ただし `cv25519` と `ed25519` は 256bits）。
右端列の「強度」は鍵長ではなく，セキュリティ強度なのでご注意を。

### SECG/NIST 推奨パラメータ

| GnuPG | SECG | NIST | アルゴリズム | 強度 bits |
|:------|:-----|:-----|:------------:|:--------:|
| `nistp256` | secp256r1 | NIST curve P-256 | ECDH, ECDSA | 128 |
| `nistp384` | secp384r1 | NIST curve P-384 | ECDH, ECDSA | 192 |
| `nistp521` | secp521r1 | NIST curve P-521 | ECDH, ECDSA | 256 |
| `secp256k1` | secp256k1 |  | ECDH, ECDSA | 128 |

SECG/NIST の楕円曲線名とのマッピングは [RFC 4492] を参考にした。
`nistp256`, `nistp384`, `nistp521` は既に [RFC 6637] で定義済みなので正式に使える。
[RFC 6637] では NIST curve P-256 については "MUST implement”， NIST curve P-521 については "SHOULD implement”， NIST curve P-384 については "MAY implement” となっている。

### Brainpool

| GnuPG | アルゴリズム | 強度 bits |
|:------|:------------:|:--------:|
| `brainpoolP256r1` | ECDH, ECDSA | 128 |
| `brainpoolP384r1` | ECDH, ECDSA | 192 |
| `brainpoolP512r1` | ECDH, ECDSA | 256 |

Brainpool については [RFC 5639] を参照のこと。
[RFC 4880bis] では `brainpoolP256r1` および `brainpoolP512r1` については "MAY implement” となっている。

### その他

| GnuPG | アルゴリズム | 強度 bits |
|:------|:------------:|:--------:|
| `cv25519` | ECDH | 128 |
| `ed25519` | EdDSA | 128 |

EdDSA については [RFC 8032] を参照のこと。
[RFC 4880bis] では Curve25519 (`cv25519`) および Ed25519 (`ed25519`) については "SHOULD implement” となっている。

## セキュリティ強度と楕円曲線

セキュリティ強度と暗号アルゴリズムと鍵長の関係は以下の表の通り（単位は全てbit）。

{{< div-gen >}}
<figure lang="en">
<style>
main table.nist2 th  {
  vertical-align:middle;
  text-align: center;
}
main table.nist2 td  {
  vertical-align:middle;
  text-align: center;
}
</style>
<table class="nist2">
<thead>
<tr>
<th>Security<br>Strength</th>
<th>Symmetric<br> key<br> algorithms</th>
<th>FFC<br>(e.g., DSA, D-H)</th>
<th>IFC<br>(e.g., RSA)</th>
<th>ECC<br>(e.g., ECDSA)</th>
</tr>
</thead>
<tbody>
<tr><td>$\le 80$</td><td>2TDEA</td><td>$L=1024$<br>$N=160$</td><td>$k=1024$</td> <td>$f = 160\text{ - }223$</td></tr>
<tr><td>$112$</td><td>3TDEA</td><td>$L=2048$<br>$N=224$</td> <td>$k=2048$</td> <td>$f = 224\text{ - }255$</td></tr>
<tr><td>$128$</td><td>AES-128</td><td>$L=3072$<br>$N=256$</td> <td>$k=3072$</td> <td>$f = 256\text{ - }383$</td></tr>
<tr><td>$192$</td><td>AES-192</td><td>$L=7680$<br>$N=384$</td> <td>$k=7680$</td> <td>$f = 384\text{ - }511$</td></tr>
<tr><td>$256$</td><td>AES-256</td><td>$L=15360$<br>$N=512$</td><td>$k=15360$</td><td>$f=512+$</td></tr>
</tbody>
</table>
<figcaption>Comparable strengths (via <q><a href='https://doi.org/10.6028/NIST.SP.800-57pt1r4'>SP800-57 Part 1 Revision 4 <sup><i class='far fa-file-pdf'></i></sup></a></q>)</figcaption>
</figure>
{{< /div-gen >}}

2030年以降も Acceptable な鍵が要件なら 128bits 以上のセキュリティ強度が必要だが， [GnuPG] で利用できる楕円曲線は全て問題ないことが分かるだろう。

## 【付録】 対話モードによる ECC 鍵の選択

対話モードで ECC 鍵を選択して作成するには以下のコマンドラインで起動する。

```text
$ gpg --full-gen-key --expert
```

このうち ECC 鍵を作成できるのは，以下の選択肢の内 (9) から (11) を選択した場合である。

{{< highlight text "hl_lines=8-10" >}}
ご希望の鍵の種類を選択してください:
   (1) RSA と RSA (デフォルト)
   (2) DSA と Elgamal
   (3) DSA (署名のみ)
   (4) RSA (署名のみ)
   (7) DSA (機能をあなた自身で設定)
   (8) RSA (機能をあなた自身で設定)
   (9) ECC と ECC
  (10) ECC (署名のみ)
  (11) ECC (機能をあなた自身で設定)
  (13) 既存の鍵
あなたの選択は? 
{{< /highlight >}}

(9) を選択すると電子署名用の主鍵と暗号化用の副鍵をセットで作成する。
(10) と (11) を選択すると電子署名用の主鍵のみ作成する（(11) では主鍵の機能を選択できる） 。

いずれを選択した場合でも以下の楕円曲線の選択肢が表示される。

```text
ご希望の楕円曲線を選択してください:
   (1) Curve 25519
   (3) NIST P-256
   (4) NIST P-384
   (5) NIST P-521
   (6) Brainpool P-256
   (7) Brainpool P-384
   (8) Brainpool P-512
   (9) secp256k1
あなたの選択は? 
```

実際にどのような鍵が作成されるか以下に示す。

### 「ECC と ECC」を選んだ場合

| 選択した楕円曲線 | 主鍵（電子署名用） | 副鍵（暗号化用） | 強度 bit |
|:-----------------|:-------------------|:-----------------|:--------:|
| `Curve 25519` | EdDSA (`ed25519`) | ECDH (`cv25519`) | 128 |
| `NIST P-256` | ECDSA (`nistp256`) | ECDH (`nistp256`) | 128 |
| `NIST P-384` | ECDSA (`nistp384`) | ECDH (`nistp384`) | 192 |
| `NIST P-521` | ECDSA (`nistp521`) | ECDH (`nistp521`) | 256 |
| `Brainpool P-256` | ECDSA (`brainpoolP256r1`) | ECDH (`brainpoolP256r1`) | 128 |
| `Brainpool P-384` | ECDSA (`brainpoolP384r1`) | ECDH (`brainpoolP384r1`) | 192 |
| `Brainpool P-512` | ECDSA (`brainpoolP512r1`) | ECDH (`brainpoolP512r1`) | 256 |
| `secp256k1` | ECDSA (`secp256k1`) | ECDH (`secp256k1`) | 128 |

### 「ECC (署名のみ)」を選んだ場合

「ECC (機能をあなた自身で設定)」でも同じ。

| 選択した楕円曲線 | 主鍵（電子署名用） | 副鍵 | 強度 bit |
|:-----------------|:-------------------|:-----|:--------:|
| `Curve 25519` | EdDSA (`ed25519`) | なし | 128 |
| `NIST P-256` | ECDSA (`nistp256`) | なし | 128 |
| `NIST P-384` | ECDSA (`nistp384`) | なし | 192 |
| `NIST P-521` | ECDSA (`nistp521`) | なし | 256 |
| `Brainpool P-256` | ECDSA (`brainpoolP256r1`) | なし | 128 |
| `Brainpool P-384` | ECDSA (`brainpoolP384r1`) | なし | 192 |
| `Brainpool P-512` | ECDSA (`brainpoolP512r1`) | なし | 256 |
| `secp256k1` | ECDSA (`secp256k1`) | なし | 128 |

## ブックマーク

- {{< pdf-file title="楕円曲線暗号の整備動向＋楕円暗号の実装状況" link="http://www.jnsa.org/seminar/pki-day/2011/data/02_kanaoka.pdf">}}
- [楕円曲線暗号の鍵長に512bitがなく521bitがある理由 | maidsphere](http://www.maidsphere.jp/archive/%E6%A5%95%E5%86%86%E6%9B%B2%E7%B7%9A%E6%9A%97%E5%8F%B7%E3%81%AE%E9%8D%B5%E9%95%B7%E3%81%AB512bit%E3%81%8C%E3%81%AA%E3%81%8F521bit%E3%81%8C%E3%81%82%E3%82%8B%E7%90%86%E7%94%B1)

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})
- [GnuPG チートシート（鍵作成から失効まで）]({{< ref "/openpgp/gnupg-cheat-sheet.md" >}})

[OpenPGP]: http://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"

[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"

[RFC 4492]: https://tools.ietf.org/html/rfc4492 "RFC 4492 - Elliptic Curve Cryptography (ECC) Cipher Suites for Transport Layer Security (TLS)"
[RFC 8032]: https://tools.ietf.org/html/rfc8032 "RFC 8032 - Edwards-Curve Digital Signature Algorithm (EdDSA)"
[RFC 5639]: https://tools.ietf.org/html/rfc5639 "RFC 5639 - Elliptic Curve Cryptography (ECC) Brainpool Standard Curves and Curve Generation"
[RFC 6637]: https://tools.ietf.org/html/rfc6637 "RFC 6637 - Elliptic Curve Cryptography (ECC) in OpenPGP"


## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
