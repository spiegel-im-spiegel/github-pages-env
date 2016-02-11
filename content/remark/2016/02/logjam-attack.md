+++
date = "2016-02-11T03:59:41+09:00"
description = "以前 Qiita に書いた記事を再掲載する。ちなみに元記事は2015年5月に公開している。"
draft = false
tags = ["security", "vulnerability", "cryptography", "openssl"]
title = "TLS における Diffie-Hellman 鍵交換の脆弱性（再掲載）"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

なんか [node.js](https://nodejs.org/) が今さら Logjam 攻撃に対応したとか言ってるので

- [Node v0.10.42 (LTS) | Node.js](https://nodejs.org/en/blog/release/v0.10.42/)
- [Node v0.12.10 (LTS) | Node.js](https://nodejs.org/en/blog/release/v0.12.10/)
- [Node v4.3.0 (LTS) | Node.js](https://nodejs.org/en/blog/release/v4.3.0/)
- [Node v5.6.0 (Stable) | Node.js](https://nodejs.org/en/blog/release/v5.6.0/)

{{< fig-quote title="Node v5.6.0 (Stable)" link="https://nodejs.org/en/blog/release/v5.6.0/" lang="en" >}}
<q>upgrade from 1.0.2e to 1.0.2f. To mitigate against the Logjam attack, TLS clients now reject Diffie-Hellman handshakes with parameters shorter than 1024-bits, up from the previous limit of 768-bits.</q>
{{< /fig-quote >}}

以前 [Qiita に書いた記事](http://qiita.com/spiegel-im-spiegel/items/af0cdb620ad79c4d0f36)をこちらでも再掲載する。
ちなみに元記事は2015年5月に公開している。

ちなみに OpenSSL 1.0.2f では Logjam 攻撃への追加対応のほか，安全でない素数を生成する問題や SSLv2 をブロックできない問題に対応している。
むしろこちらの方が重要か。

- [[openssl-announce] OpenSSL Security Advisory](https://mta.openssl.org/pipermail/openssl-announce/2016-January/000061.html)
- [Vulnerability Note VU#257823 - OpenSSL re-uses unsafe prime numbers in Diffie-Hellman protocol](http://www.kb.cert.org/vuls/id/257823)
- [JVNVU#95668716: OpenSSL の DH プロトコルにおける脆弱性](http://jvn.jp/vu/JVNVU95668716/)

安全でない素数を生成する問題については CVSSv3 基本値 `CVSS:3.0/AV:N/AC:H/PR:N/UI:N/S:U/C:H/I:H/A:N` で 7.4 と見積もられている。

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 高（H）           |
| 必要な特権レベル（PR）                  | 不要（N）         |
| ユーザ関与レベル（UI）                  | 不要（N）         |
| スコープ（S）                           | 変更なし (U)      |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | 高（H）           |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |


## 脆弱性の内容

通称 “Logjam” 攻撃。

かつて騒がれた FREAK 脆弱性と同じく， TLS 経路上に「中間者」がいる場合， Diffie-Hellman（DH）鍵交換で使われる鍵を輸出用の脆弱なものにダウングレードさせられる。
FEAK のときとは異なり，特定の実装の脆弱性ではなく TLS プロトコルの欠陥。

ちなみに DH 鍵交換アルゴリズムは公開鍵暗号の一種で，お互いに（もちろん第3者にも）秘密情報（秘密鍵）を知られることなく安全にセッション鍵を生成することができる。
ベースとなるロジックは「離散対数問題」と呼ばれるもので ElGamal や DSA と同系統のロジック。なので，鍵長の管理も ElGamal や DSA と同等のものが要求される。

暗号強度と各暗号方式の鍵長の関係は以下のとおり（単位はすべて bit）

<figure>
<table>
<thead>
<tr>
<th>Bits of<br> security</th>
<th>Symmetric key<br> algorithms</th>
<th>FFC<br>(e.g., DSA, D-H)</th>
<th>IFC<br>(e.g., RSA)</th>
<th>ECC<br>(e.g., ECDSA)</th>
</tr>
</thead>
<tbody>
<tr><td class='right'>80</td> <td>2TDEA</td>  <td>$L=1024$<br>$N=160$</td> <td>$k=1024$</td> <td>$160 \le f \lt 224$</td></tr>
<tr><td class='right'>112</td><td>3TDEA</td>  <td>$L=2048$<br>$N=224$</td> <td>$k=2048$</td> <td>$224 \le f \lt 256$</td></tr>
<tr><td class='right'>128</td><td>AES-128</td><td>$L=3072$<br>$N=256$</td> <td>$k=3072$</td> <td>$256 \le f \lt 384$</td></tr>
<tr><td class='right'>192</td><td>AES-192</td><td>$L=7680$<br>$N=384$</td> <td>$k=7680$</td> <td>$384 \le f \lt 512$</td></tr>
<tr><td class='right'>256</td><td>AES-256</td><td>$L=15360$<br>$N=512$</td><td>$k=15360$</td><td>$512 \le f$</td></tr>
</tbody>
</table>
<figcaption>Comparable strengths (via <q lang='en'><a href='http://csrc.nist.gov/publications/nistpubs/800-57/sp800-57_part1_rev3_general.pdf'>SP800-57 Part 1 (Revision 3) <sup><i class='fa fa-file-pdf-o'></i></sup></a></q>)</figcaption>
</figure>

2030年以降も安全に使える暗号強度は $128\,\mathrm{bits}$ 以上だと言われている。
Logjam 攻撃では $L=512\,\mathrm{bits}$ にダウングレードさせられるが全くお話にならない強度だということが分かるだろう。

暗号について詳しくは拙文「[わかる！ OpenPGP 暗号](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml)」の「[暗号に関する雑多な話](http://www.baldanders.info/spiegel/archive/pgpdump/openpgp.shtml#appendix)」あたりをどうぞ。
また [FREAK については Gist にまとめている](https://gist.github.com/spiegel-im-spiegel/47f340122c895ccc8bb8)ので，そちらも参考にどうぞ。

## 影響度（CVSS）

[CVE-2015-1716] より

CVSSv2 基本値 5.0 (AV:N/AC:L/Au:N/C:P/I:N/A:N)

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 低（L）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | 部分的（P）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）         |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

[CVE-2015-4000] より

CVSSv2 基本値 4.3 (AV:N/AC:M/Au:N/C:N/I:P/A:N)

| 基本評価基準                            | 評価値            |
|----------------------------------------:|:------------------|
| 攻撃元区分（AV）                        | ネットワーク（N） |
| 攻撃条件の複雑さ（AC）                  | 中（M）           |
| 攻撃前の認証要否（Au）                  | 不要（N）         |
| 情報漏えいの可能性（機密性への影響, C） | なし（N）         |
| 情報改ざんの可能性（完全性への影響, I） | 部分的（P）       |
| 業務停止の可能性（可用性への影響, A）   | なし（N）         |

CVSS については[デモページ](http://www.baldanders.info/spiegel/archive/cvss/cvss2.html)を参照のこと。

[CVE-2015-1716]: https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2015-1716
[CVE-2015-4000]: https://web.nvd.nist.gov/view/vuln/detail?vulnId=CVE-2015-4000

## 影響を受ける実装

- Microsoft : 影響あり。 [MS15-055] で修正済み
- OpenSSL 規定で輸出グレード暗号は無効化されているので，実質的には大丈夫？ : [Logjam, FREAK and Upcoming Changes in OpenSSL - OpenSSL Blog](https://www.openssl.org/blog/blog/2015/05/20/logjam-freak-upcoming-changes/)
    - [脆弱性を修正した「OpenSSL」の最新版が公開、“Logjam”脆弱性の修正も - 窓の杜](http://www.forest.impress.co.jp/docs/news/20150615_706966.html)

> また、「OpenSSL」v1.0.2b/v1.0.1nに関しては、TLS通信で暗号強度の弱い輸出グレードの暗号へ意図せずダウングレードされてしまう“Logjam”脆弱性（CVE-2015-4000）も修正されている。

- Apple Safari : 影響あり。修正版なし
- Android : 影響あり。修正版なし
- Google Chrome : 影響あり。修正版なし。 Chrome 43 でも解消されてないらしい
- Mozilla Firefox : 影響あり。修正版なし
    - [脆弱性 Logjam Attack対策 for Firefox - ふらっと 気の向くままに](http://datyotosanpo.blog.fc2.com/blog-entry-69.html) : Firefox が更新されるまで DH(E) を無効化する
- OpenVPN : 影響はほとんどない？
    - [TLSの脆弱性「Logjam」のOpenVPNへの影響 « yamata::memo](http://yamatamemo.blogspot.jp/2015/05/tlslogjamopenvpn.html)
        - OpenVPNでは、サーバーセットアップ時に openssl dhparam コマンドを使ってOpenVPN専用のDHパラメータを生成しているため、DHパラメータを別個に生成しない使用方法よりは安全といえる。
        - openssl dhparam コマンドの実行時に鍵長を 2048ビット以上にしていれば安全。1024ビットの場合は攻撃される可能性は否定できないが、それでも簡単ではない。
        - TLS-Authが有効になっていればこの種のTLSのダウングレード攻撃は回避できる。
- OpenSSH : 影響あり。 [On OpenSSH and Logjam – Technology & Policy – Jethro Beekman](https://jbeekman.nl/blog/2015/05/ssh-logjam/)

- [Logjam: PFS Deployment Guide](https://weakdh.org/sysadmin.html) : サーバ側での回避例

[MS15-055]: https://technet.microsoft.com/library/security/MS15-055 "マイクロソフト セキュリティ情報 MS15-055 - 重要"

## 影響の有無を確認する方法

ブラウザで “[The Logjam Attack](https://weakdh.org/)” のサイトを訪れると自動的に判定してくれる。

> Warning! Your web browser is vulnerable to Logjam and can be tricked into using weak encryption. You should update your browser.

と表示されると影響を受ける可能性がある。

サーバ側は “[Guide to Deploying Diffie-Hellman for TLS](https://weakdh.org/sysadmin.html)” のページにある “Server Test” で確認できる。

## 参考ページ

- [Logjam TLS Attack](https://weakdh.org/logjam.html)
- [Logjam: How Diffie-Hellman Fails in Practice](https://weakdh.org/)
    - [Imperfect Forward Secrecy: How Diffie-Hellman Fails in Practice](https://weakdh.org/imperfect-forward-secrecy.pdf) (PDF)
- [セキュリティホール memo の記事](http://www.st.ryukoku.ac.jp/~kjm/security/memo/2015/05.html#20150521_Logjam)
- [Logjam Attackについてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20150521/1432219012)
- [The Logjam (and Another) Vulnerability against Diffie-Hellman Key Exchange - Schneier on Security](https://www.schneier.com/blog/archives/2015/05/the_logjam_and_.html)
- [OpenSSH環境に対するLogjam脆弱性の対応 | NaviPlus Engineers' Blog](http://tech.naviplus.co.jp/2015/05/25/openssh%E7%92%B0%E5%A2%83%E3%81%AB%E5%AF%BE%E3%81%99%E3%82%8Blogjam%E8%84%86%E5%BC%B1%E6%80%A7%E3%81%AE%E5%AF%BE%E5%BF%9C/) : “[On OpenSSH and Logjam](https://jbeekman.nl/blog/2015/05/ssh-logjam/)” の日本語解説
- [Logjam, Part 1: Why the Internet is Broken Again (an Explainer) | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2015/05/logjam-internet-breaks-again)
- [Logjam, Part 2: Did the NSA Know the Internet Was Broken? | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2015/05/logjam-part-2-did-nsa-know-years-internet-was-broken)

