+++
title = "OpenPGP の電子署名は「ユーザーの身元を保証し」ない"
date = "2019-03-21T00:28:40+09:00"
description = "つまり公開鍵や電子署名で「ユーザーの身元を保証」するのではなく「身元の保証されたユーザ」同士で鍵と電子署名を運用するのである。"
image = "/images/attention/openpgp.png"
tags = [
  "cryptography",
  "openpgp",
  "management",
  "pki",
  "trust",
  "gnupg",
  "git",
  "github",
  "certification",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

重箱の隅を突っつくような内容で申し訳ないのだが

- [「GitKraken 5.0」がリリース ～GPGコミットや“Interactive Rebase”をサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1175019.html)

という記事で

{{< fig-quote title="「GitKraken 5.0」がリリース ～GPGコミットや“Interactive Rebase”をサポート" link="https://forest.watch.impress.co.jp/docs/news/1175019.html" >}}
<q>メジャーアップデートとなる本バージョンでは、“GNU Privacy Guard (GPG)”による署名付きのコミットがサポートされた。ユーザーの身元を保証し、他のユーザーによるなりすましを防止することができる。</q>
{{< /fig-quote >}}

などと書いてあって「それはちゃうやろ」という話。

## 暗号機能の4要素

昔からよく言われる暗号機能の4要件は以下の通り。

- 機密性（Confidentiality）
- 完全性（Integrity）
- 認証（Authentication）
- 否認防止（Non-repudiation）

このうちデータへの電子署名では主に完全性と否認防止を行う。
否認防止という言葉はちょっと耳慣れないかもしれないが，要するに「あなたはこのデータに署名した」という事実を否認することが出来ない，という意味である[^pki1]。

[^pki1]: メッセージング・システムを含む暗号通信においては「否認防止」よりむしろ「否認可能 (Deniability)」が要求される場合がある。詳しくは拙文「[OTR over XMPP](https://baldanders.info/blog/000787/ "OTR over XMPP — Baldanders.info")」を参考にどうぞ。

これらを達成するためには電子署名に使う公開鍵が鍵オーナーと正しく紐付いている必要がある[^bt1]。
おそらく最初の記事は「公開鍵が鍵オーナーと正しく紐付いている」という前提で「ユーザーの身元を保証」などと書いているのかもしれないが，話はそう簡単ではないのだ。

[^bt1]: たとえば Bitcoin/Blockchain は公開鍵と鍵オーナーとの紐づけを行わない。元帳である Blockchain に記載された取引自体は改竄もなく正しいとしても誰がそれを行ったかを証明する術がない。証明するためには別の仕組みが必要となる。あるいは Bitcoin/Blockchain は「信用」を極力排除することでシステム自体の「正しさ」を担保しているとも言える。

## 公開鍵基盤

### X.509

「公開鍵が鍵オーナーと正しく紐付いている」ことを証明するために必要なのが「公開鍵基盤（Public-Key Infrastructure; PKI）」である。
公開鍵基盤として最も有名なのは HTTPS 接続で使われている X.509 であろう。
X.509 の「信用モデル（trust model）」では hierarchical な「認証局（Certification Authority; CA）」を構成し，その認証局が公開鍵を証明することで成り立っている。

でも実は認証局が証明しているのは「公開鍵が鍵オーナーと正しく紐付いている」ことだけで，鍵オーナーの「身元を保証」しているわけではない。

そこで HTTPS には [EV SSL (Extended Validation SSL) なる奇っ怪な仕組み](https://baldanders.info/blog/000277/ "Extended Validation SSL — Baldanders.info")が組み込まれた。
これは鍵オーナーの「身元を保証」するための仕組みで，鍵オーナーは認証局に対して自身の身元を証明するものを提出し認証局は公開鍵の管理をより厳格に行う，ということらしい。

正直に言って「屋上屋を架す」仕組みであり認証局の責務を逸脱していると思うのだが，まぁ深くは突っ込むまい。

### Web of Trust

これに対して PGP/[OpenPGP] が伝統的に執っている信用モデルは「信用の輪（web of trust）」と呼ばれている。

これは要するに「友達の友達は友達」というやつで，ユーザ同士がお互いの鍵を相互に認証することで信用を構成する仕組みである。

何故 PGP/[OpenPGP] が X.509 のような「権威による認証」を採用しなかったかというと，それは PGP の出自に関係がある。
PGP の作者である Phil Zimmermann は，当時は反核運動家であり国家等の「権威」に依らない信用モデルを必要としていたと言われている[^gpg2]。

[^gpg2]: 本当のことろ，作者である Phil Zimmermann がどこまで企図していたのかは知らないが， PGP の登場によって暗号は初めて（国家や大企業のものではなく）一般の人も「使える」ものになったことは確かである。なお [OpenPGP] 実装のひとつである [GnuPG] では伝統的な「信用の輪」以外にも X.509 の信用モデルや [TOFU (Trust On First Use)](https://en.wikipedia.org/wiki/Trust_on_first_use) と呼ばれる信用モデルもサポートしている。

OpenPGP は名前だけならどんな鍵でも作れる。
たとえば Bitcoin の作者と言われる Satoshi Nakamoto の公開鍵は[公開鍵サーバを探せば簡単に見つかる]({{< ref "/remark/2016/05/openpgp-key-of-satoshi-nakamoto.md" >}})が，それが「あの」 Satoshi Nakamoto 本人であると示す証拠は（少なくとも公開鍵自体には）存在しない。
OpenPGP 公開鍵やその電子署名で赤の他人の「身元を保証」することは出来ないのだ。

## Git を中心としたチーム運営に [OpenPGP] を利用する

じゃあ git commit で OpenPGP 署名を付与することにどんな意義があるかというと，それはチーム運営で威力を発揮する。
つまり公開鍵や電子署名で「ユーザーの身元を保証」するのではなく「身元の保証されたユーザ」同士で鍵と電子署名を運用するのである。
これでチーム以外からのなりすまし commit を検知（防止ではない[^d1]）しやすくなる。

[^d1]: なりすましの防止はできなくとも，きちんと鍵と電子署名を運用しているのであれば，抑止にはなるだろう。

この辺について詳しくは，拙文「[OpenPGP 鍵管理に関する考察]({{< relref "./openpgp-key-management.md" >}})」を書いたので興味があれば参照してほしい。

また GitHub のようにアカウントと公開鍵を紐つけることによってサービス内における強力なポートフォリオとして機能している点は見逃せないだろう。
GitHub 上の{{< ruby "contribution" >}}活動{{< /ruby >}}がそのまま「エンジニアとしてのユーザ」の身元を保証しているわけだ。

ホンマ git ってよく出来たシステムだよなぁ。

## ブックマーク

- [Git Commit で OpenPGP 署名を行う]({{< relref "./git-commit-with-openpgp-signature.md" >}})
- [わかる！ OpenPGP 暗号 — Baldanders.info](https://baldanders.info/spiegel/cc-license/)

- [spiegel-im-spiegel/gpgpdump: OpenPGP packet visualizer](https://github.com/spiegel-im-spiegel/gpgpdump) : OpenPGP 鍵や電子署名のダンプには拙作をどうぞ（宣伝`w`）

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
