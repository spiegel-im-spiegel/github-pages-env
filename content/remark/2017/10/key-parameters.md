+++
title = "暗号鍵関連の各種変数について"
date =  "2017-10-17T16:32:03+09:00"
description = "この記事は将来の記事で再利用するための snippet 置き場として使うことにする。"
tags = ["security", "cryptography", "hash", "risk", "management"]

[scripts]
  mathjax = true
  mermaidjs = false
+++

「[ECDSA鍵をGitHubで使う方法](https://qiita.com/darai0512/items/c7b47d1b3fe06c4dea7d)」で SP 800-57 第一部が Rev.4 になってるのを見て一瞬「ふぁ！」となったが，よく考えたら今年始めに書いた「[最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})」では既に Rev.4 を参照していたのだった。
でも古い記事を見返したら結構 Rev.3 のままになってるので，このさい全部アップデートすることにした。

この記事は将来の記事で再利用するための snippet 置き場として使うことにする。
表のレイアウトの関係で携帯端末で見ている人は見づらいかもしれないけど，そこはご容赦。

{{< div-box type="markdown" >}}
**【2020-06-07】** SP 800-57 第一部 の Rev.5 の最終版がリリースされていたので，以降 Rev.5 をベースに書き換えた。
{{< /div-box >}}

## 参照資料

ここで参照する資料は，米国 [NIST] の Special Publication (SP) 800-57 Part 1 で正式タイトルは「Recommendation for Key Management Part 1: General （鍵管理における推奨事項 第一部：一般事項）」となっている（日本語訳は IPA によるもの）。

- [SP 800-57 Part 1 Rev. 5, Recommendation for Key Management: Part 1 – General | CSRC](https://csrc.nist.gov/publications/detail/sp/800-57-part-1/rev-5/final)
    - {{< pdf-file title="NIST Special Publication 800-57 Part 1; Revision 5; Recommendation for Key Management Part 1: General" link="https://doi.org/10.6028/NIST.SP.800-57pt1r5" >}}

なお，ひとつ前の [Rev.4](https://csrc.nist.gov/publications/detail/sp/800-57-part-1/rev-4/final "SP 800-57 Part 1 Rev. 4 Recommendation for Key Management, Part 1: General") については {{< pdf-file title="IPA による日本語訳" link="https://www.ipa.go.jp/files/000055490.pdf" >}} があるので参考にどうぞ。

## セキュリティ強度と鍵長の関係

最初はセキュリティ強度と鍵長の関係を示す表。
単位は全てビットである。

{{< comparable-security-strengths >}} <!-- 要 MathJax -->

Symmetric key algorithms は共通鍵暗号アルゴリズム全般を指す。
たとえば AES とか。
IFC (Integer Factorization Cryptosystems) は素因数分解問題ベースの公開鍵暗号アルゴリズムで RSA がこれに該当する。
FFC (Finite Field Cryptosystems) は離散対数問題ベースの公開鍵暗号アルゴリズムで Diffie-Hellman や ElGamal, DSA などがこれに該当する。
ECC (Elliptic Curve Cryptosystems) は離散対数問題でも特に楕円曲線上の離散対数問題ベースの公開鍵暗号アルゴリズムを指す。
たとえば ECDH や ECDSA など。

IFC では $k$，FFC では $L$，ECC では $f$ が鍵長を示す。
たとえばセキュリティ強度が128ビットなら

- AES 128bit
- ElGamal, DSA 3072bit
- RSA 3072bit
- ECDH, ECDSA 256bit

の組み合わせで「[ベストマッチ](https://dic.pixiv.net/a/%E3%83%93%E3%83%AB%E3%83%89%E3%83%89%E3%83%A9%E3%82%A4%E3%83%90%E3%83%BC) キター！」となる。

## セキュリティ強度と Hash 関数の関係

次はセキュリティ強度とHash 関数の関係を示す表。

{{< security-strengths-for-hash >}} <!-- 要 MathJax -->

考え方は先程の暗号鍵長のときと同じ。
ただし Hash 関数の場合は使用目的ごとに要求されるアルゴリズムが異なるので注意が必要である。

## セキュリティ強度と有効期限

こちらはセキュリティ強度の有効期限を表したものだ。

{{< security-strength-time-frames >}} <!-- 要 MathJax -->

各用語はそれぞれ

| 用語       | 意味                     |
|:---------- |:------------------------ |
| Applying   | 適用                     |
| Processing | 処理                     |
| Acceptable | 許容                     |
| Legacy use | 許容（レガシー使用のみ） |
| Disallowed | 禁止                     |

という意味だ。
例を挙げると，セキュリティ強度112ビットの暗号スイート（Cipher Suites）を適用する場合は2030年までは許容するけど2031年以降は禁止。
すでに暗号化されているデータを復号したい場合でも2031年以降はレガシー・システムしか許容しない，ということになる。

たとえば ssh 認証は「適用」なので，いまだ多くの人が使ってる（かもしれない） RSA 2048ビットの鍵は2031年以降は使用禁止となるわけだ。
まぁ，そんな先まで同じシステムで同じ鍵を使い続けるかどうかは分からないが（なので今使ってる鍵を慌てて新調する必要はない。新規に作成するなら128ビット強度の鍵をお勧めするが）。

なお，これは各アルゴリズムに危殆化要因となる脆弱性等がない場合の話である。
したがって暗号製品を使うシステムの管理者やセキュリティ管理者は常に暗号関係のトピックに耳を澄ませておくべきであろう。

## OpenPGP で利用可能なアルゴリズム

この項は「[OpenPGP で利用可能なアルゴリズム]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})」に移動した。

## パスワードの強度

これは [NIST] ではなく IPA の資料だが，文字種と文字数の組み合わせによるパスワードの強度についても上げておこう。
出典は以下。

- [情報漏えいを防ぐためのモバイルデバイス等設定マニュアル：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ipg/documents/dev_setting_crypt.html)
    - {{< pdf-file title="情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編" link="https://www.ipa.go.jp/files/000026760.pdf" >}}

{{< security-strengths-for-password >}} <!-- 要 MathJax -->

測定基準は以下の通り。

{{< div-gen >}}
<figure>
<blockquote>
<q>利用できる文字種類すべてを完全にランダムに選択して作ったパスワードを一つ一つ調べる全数探索により1日で解読しようとした際にかかるおおまかな想定攻撃コストを示しています。ここでは、全数探索(暗号鍵の総数256)でDES10を1日で解読するためのコストを約250万円と仮定します。また、パスワードを1つ検査するのとDESの暗号鍵を1つ検査するコストは同じであるとし、パスワードを求めるのに必要な計算量(検査する個数)が半分になればコストも半分、2倍になればコストも2倍になるものとしています。</q>
</blockquote>
<figcaption><q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より</figcaption>
</figure>
{{< /div-gen >}}

## ブックマーク

- [CRYPTREC Report 2013 — Baldanders.info](https://baldanders.info/blog/000740/)
- [Prohibiting RC4 — Baldanders.info](https://baldanders.info/blog/000810/)
- [CRYPTREC | SHA-1の安全性について](http://www.cryptrec.go.jp/topics/cryptrec_20151218_sha1_cryptanalysis.html)
    - {{< pdf-file title="CRYPTREC暗号技術ガイドライン(SHA-1)" link="http://www.cryptrec.go.jp/report/c13_tech_guideline_SHA-1_web.pdf" >}}
- [scryptがGPUに破られる時 | びりあるの研究ノート](https://blog.visvirial.com/articles/519)

- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})
- [「パスワードのベストプラクティス」が変わる]({{< ref "/remark/2017/10/changes-in-password-best-practices.md" >}})

[NIST]: https://www.nist.gov/ "National Institute of Standards and Technology | NIST"
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
