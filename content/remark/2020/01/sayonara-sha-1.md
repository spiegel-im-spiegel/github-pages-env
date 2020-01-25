+++
title = "（何度目かの）さようなら SHA-1"
date =  "2020-01-09T23:02:25+09:00"
description = "思ったより騒がれてる感じなので記事を立ててみた。"
image = "/images/attention/openpgp.png"
tags = [
  "security",
  "risk",
  "cryptography",
  "hash",
  "sha-1",
  "openpgp",
  "gnupg",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2019年11月に書いた「[GnuPG 2.2.18 リリース： さようなら SHA-1]({{< ref "/release/2019/11/gnupg-2_2_18-is-released.md" >}})」に追記してお終いにしようかと思ったが，思ったより騒がれてる感じなので記事を立ててみた。

というわけで：

- [Cryptology ePrint Archive: Report 2020/014 - SHA-1 is a Shambles - First Chosen-Prefix Collision on SHA-1 and Application to the PGP Web of Trust](https://eprint.iacr.org/2020/014)
- [SHA-1 is a Shambles](https://sha-mbles.github.io/)
- [PGP keys, software security, and much more threatened by new SHA1 exploit | Ars Technica](https://arstechnica.com/information-technology/2020/01/pgp-keys-software-security-and-much-more-threatened-by-new-sha1-exploit/)
- [New SHA-1 Attack - Schneier on Security](https://www.schneier.com/blog/archives/2020/01/new_sha-1_attac.html)

要約としては [Bruce Schneier さんの記事](https://www.schneier.com/blog/archives/2020/01/new_sha-1_attac.html "New SHA-1 Attack - Schneier on Security")が分かりやすいので，これを起点に紹介していく。

## ハッシュ値の衝突問題

（「[SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})」からの抜粋）

暗号技術におけるハッシュ関数とは，以下の機能を持ったアルゴリズムである

1. 任意のデータ列を一定の長さのデータ列（ハッシュ値）に「要約」する
1. ハッシュ値から元のデータ列を推測できない
1. ひとつのハッシュ値に対して複数のデータ列が（実時間で）見つからない

ハッシュ関数はメッセージ認証符号（Message Authentication Code; MAC）や電子署名（digital signature）の中核技術のひとつであり，データの「完全性（Integrity）」を担保する重要な要素である。
特に3番目の「ひとつのハッシュ値に対して複数のデータ列が（実時間で）見つからない」という機能が破られると，そのハッシュ関数では完全性を担保できなくなってしまう。
これを「ハッシュ値の衝突問題」という。

## SHA-1 ハッシュ値を力づくで攻略してみる

SHA-1 における「ハッシュ値の衝突問題」は[2004年まで遡る](https://baldanders.info/blog/000048/)が，当時の攻略法は SHA-1 アルゴリズムの危殆化を狙ったもので，しかもその後の進展は殆どなかった。

この状況が変わったのが2015年の “[freestart collisions for SHA-1](https://sites.google.com/site/itstheshappening/)” 論文である。

- [SHA-1 衝突問題： 廃止の前倒し]({{< ref "/remark/2015/problem-of-sha1-collision.md" >}})

ここで初めて「SHA-1 ハッシュ値を力づくで攻略」できる可能性が示された。
さらに2017年には Google によって最初の SHA-1 衝突例が公表された。

- [最初の SHA-1 衝突例]({{< ref "/remark/2017/02/sha-1-collision.md" >}})

今回の “[SHA-1 is a Shambles](https://sha-mbles.github.io/)” の注目点は

1. “chosen-prefix collision for SHA-1” なる手法により，衝突可能なデータを用意する際の自由度が高い
2. ハッシュ値を攻略する際の計算機パワーの調達コストが比較的実用的なレベルまで下がった

の2つである。
特に2番目が重要で， “Nvidia GTX 1060GPU” × 900 の構成で2ヶ月ほどで攻略できたらしい。
コストにして 45k USD だそうだ[^cost1]。

[^cost1]: 単純に 1 USD = 110 JPY とするなら 45k USD = 4.95M JPY ほど。まぁ五百万円以下で攻略できてしまうわけですな。

## [OpenPGP] / [GnuPG] は既に SHA-1 を捨てつつある

“[SHA-1 is a Shambles](https://sha-mbles.github.io/)” では [GnuPG] が生成する公開鍵への電子署名をターゲットにしているが，実は現行の [GnuPG] 2.2 系は既に SHA-1 を捨てつつあり，電子署名に使うハッシュ関数の既定は SHA256 である（“[SHA-1 is a Shambles](https://sha-mbles.github.io/)” で指摘されているのはレガシー・バージョンの 1.4 系のほう）。

たとえば私が git コミットへの署名などで[普段遣いしている OpenPGP 鍵](https://baldanders.info/pubkeys/)は2013年に作ったものだが，鍵への署名は SHA256 で行っている。
私の公開鍵をチェックするには以下のコマンドを実行すればよい（何気に拙作の [gpgpdump] を宣伝しておく`w`）。

```text
$ wget https://baldanders.info/pubkeys/spiegel.asc -O - | gpgpdump
```

また，次期 OpenPGP 標準となる [RFC 4880bis] では新しい V5 パケットを策定中だが，鍵指紋（key fingerprint）に使うハッシュ関数の既定が SHA256 になるようだ。
[RFC 4880bis] が正式に RFC 標準として公開されれば，かつての MD5 と同じく， SHA-1 は後方互換性のためだけに残されることになるだろう[^md5a]。

[^md5a]: [RFC 4880bis] では SHA-1 は “SHOULD NOT create messages” となる。[GnuPG] 2.2.18 以降では，これが前倒しで実装されたわけだ。ちなみに [GnuPG] 2.2 系では MD5 は既にサポートされていない（1.4 系があるため）。

[OpenPGP] の実装は [GnuPG] だけではなく，今や[シマンテック社のおまけプロダクトに成り下がった PGP](https://www.symantec.com/products/encryption) をはじめ [JsvaScript](http://openpgpjs.org/ "OpenPGP.js | OpenPGP JavaScript Implementation") や [Rust](https://sequoia-pgp.org/ "Sequoia-PGP") などによる実装がある。
これらの実装が SHA-1 をどのように取り扱っているかは分からない。
どなたか教えて！

## [Git] のコミット・ハッシュはどうなるのか

おそらく今回の件でもっとも議論を呼ぶのは [git] なんじゃないだろうか。
[Git] のコミット・ハッシュは SHA-1 を使って行われるが，ハッシュ値の衝突が比較的簡単にできるようになれば，ひとつのコミット・ハッシュに対して複数のコミットが重複してしまうという問題が現実的になるかもしれない（ならないかもしれない）。

ただ，コミットへの電子署名が正しく運用されているなら，[なりすましに対してはある程度は抑止（防止ではない）できる]({{< ref "/openpgp/web-of-trust.md" >}} "OpenPGP の電子署名は「ユーザーの身元を保証し」ない")だろう。

なので，個人的にはあまり心配はしていない。
[Git] の今後の活躍にご期待ください，といったところだろうか（笑）

## ブックマーク

- ["SHA-1 is a Shambles" and forging PGP WoT signatures](https://mailarchive.ietf.org/arch/msg/openpgp/h-6vCMDFFKhVXpXLC6gAt9tK7r8)

- [OpenPGP で利用可能なアルゴリズム（RFC 4880bis 対応版）]({{< ref "/openpgp/algorithms-for-openpgp.md" >}})

[OpenPGP]: https://www.openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgpdump]: https://github.com/spiegel-im-spiegel/gpgpdump "spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[git]: https://git-scm.com/
[Git]: https://git-scm.com/

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
