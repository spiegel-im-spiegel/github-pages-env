+++
title = "「仮想通貨」と公開鍵基盤"
date = "2018-02-05T19:10:06+09:00"
description = "Bitcoin が気にするのは Blockchain に記載されるアドレスの一貫性と無矛盾性である。今回はこの部分についてもう少し詳しく書いてみる。"
image = "/images/attention/kitten.jpg"
tags = ["blockchain", "fintech", "pki", "openpgp", "x509", "trust", "management", "security"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
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
  mermaidjs = true
+++

Twitter で見かけた記事。

- [Satoshiが注意深く設定した世界の境界線 – Shin'ichiro Matsuo – Medium](https://medium.com/@ShinichiroMatsuo/-cde3f8ffa0e4)

Satoshi Nakamoto 氏の論文を引いていてかなり面白い内容だと思うが，言いたいことは単純で，私がこれまで[述べてきた]({{< ref "/remark/2018/01/cryptocurrency-are-not-crypto.md" >}} "「暗号通貨」ってゆーな！")通り

- Bitcoin のアドレスの帰属先について Bitcoin/Blockchain は関知しない。Bitcoin が気にするのは Blockchain に記載されるアドレスの一貫性と無矛盾性である。アドレスの証明が必要な場合は外部の PKI を利用するか新たに組み込む必要がある

ということに尽きる。

今回はこの部分についてもう少し詳しく書いてみる。

## まずは定義から

Blockchain もしくは Blockchain に準ずる技術を用い，価値可換なトークンによって取引を行うシステムを括弧書きで[「仮想通貨」]または「[「仮想通貨」]システム」と命名する。
この時の「価値可換なトークン」を「コイン」と命名する。
コインは量で表すことができるものとする。

また[「仮想通貨」]システムで発生する取引を記録する追記型データベースを「元帳」と命名する。
もちろん元帳は「Blockchain もしくは Blockchain に準ずる技術」を用いて実装されているわけだ。

ここで，ある[「仮想通貨」]システム上でユーザ $A$ からユーザ $B$ へコインを移転[^t1] する「取引」を考える。

[^t1]: $A$ から見ると $B$ への「出金」， $B$ からみると $A$ からの「入金」と言える。

- 「ある[「仮想通貨」]システム」を $COIN$ とする
- 取引を $T$ とし，取引の際に移転するコインの量を $c$ とする
- $A$ が持つ[「仮想通貨」]のアドレスを $a$ とし， $B$ が持つ[「仮想通貨」]のアドレスを $b$ とする

このときの取引全体を示す図式[^math1] を以下のように記述してみる。

[^math1]: 数式じゃなくて図式。数式記号を使ってるけどあくまで図式と言い張ってみる。

{{< fig-quote >}}
\[
    COIN : A[a] \xrightarrow{T(c)} B[b]
\]
{{< /fig-quote >}}

このとき取引 $T$ を元帳に追記する内容は

{{< fig-quote >}}
\[
    a \xrightarrow{c} b
\]
{{< /fig-quote >}}

であり，取引関係者である $A$ や $B$ は一切登場しないのがポイントである。

## [「仮想通貨」]はアドレスの帰属先を証明（Certificate）しない

[「仮想通貨」]は $a$ の帰属先が $A$ であること，あるいは $b$ の帰属先が $B$ であることを証明しないし証明できない。
もう少し厳密にいうなら「[「仮想通貨」]はアドレスの帰属先を証明する責務を負わない」と言える。

このことが何をもたらすか，いくつかシナリオを考えてみよう。

- $a$ の帰属先が $A$ であると証明できない
    - $B$ は入金 $c$ を $A$ によるものではないと主張できる。 $B$ は $A$ からの入金を否認し $A$ に尚も $c$ を請求するかもしれない
    - $a$ は別の誰か（たとえば $E$）に乗っ取られているかもしれない。これにより $B$ は取引不成立とみなし $A$ に何らかのペナルティを課すかもしれない \\[ COIN : E[a] \xrightarrow{T( c )} B[b] \\]
- $b$ の帰属先が $B$ であると証明できない
    - $B$ は $b$ が自身に帰属しないと主張できる。これにより $B$ は $A$ からの入金を否認できる
    - $b$ は別の誰かに乗っ取られているかもしれない。これにより $A$ は取引不成立とみなして出金を拒否した上で $B$ に何らかのペナルティを課すかもしれない（出金した $c$ を上回る量の賠償請求を行うなど） \\[ COIN : A[a] \xrightarrow{T( c )} E[b] \\]

Coincheck 事例の事実関係は（今のところ）よく分かってない部分もあるが，知らない誰かがアドレスを乗っ取って知らない誰かへ「流出」したということであれば

{{< fig-quote >}}
\[
    COIN : E[a] \xrightarrow{T(c)} E[b]
\]
{{< /fig-quote >}}

という図式も成り立つ。

しかし実態がどのようなものであれ[「仮想通貨」]の元帳には $a \xrightarrow{c} b$ という記録が事実として残るのみで，それが望んだ取引なのか何らかの不正を含んでいるのかといった点について[「仮想通貨」]は一切関知しない。

## [「仮想通貨」]は P2P を前提とする

アドレスの帰属先を証明できないというのは実際の取引において致命的な問題となるが，それでもそれなりにまわっているのは[「仮想通貨」]がユーザ同士の P2P (peer-to-peer) な関係を前提にしているからである。
つまり「$a$ の帰属先は $A$ である」であり「$b$ の帰属先は $B$ である」であることを $A,B$ 相互に「信用」していることが取引の前提になっている。

しかし，見知った者同士の取引ならともかく，不特定の誰かを何の担保もなく「信用」するのは無理だし，その「信用」そのものを数学的に示す方法は存在しない。
存在しないのであれば，それに代わる「運用でカバー」するしかない。

この「運用」のロジックのことを「信用モデル（trust model）」と呼ぶ。
[「仮想通貨」]自体はアドレスに対する信用モデルを持たないが，[「仮想通貨」]を利用するサービスが何らかの信用モデルと組み合わせることによりアドレスの帰属先を証明することが可能になる。
また，出来のよい信用モデルを導入することにより不正取引を働くインセンティブが低下することも期待できるだろう。

おそらく[「仮想通貨」]を利用するユーザの多くはウォレット・サービスや通貨交換所が信用モデルを組み込むことを期待しているんじゃないかと思うが（投機目的で[「仮想通貨」]を運用している人はどうでもいいと思ってるかも知れない），実際にそれらのサービスがアドレスをどうやって「運用」してるのかは（私は現在の[「仮想通貨」]への興味が薄いので）知らない。

## 公開鍵基盤の信用モデル

ここで少し目先を変えて公開鍵基盤（Public Key Infrastructure; PKI）の信用モデルを2つほど紹介してみる。
公開鍵基盤というのは公開鍵が誰に帰属するかをサービスを横断して証明するための技術基盤である。

なぜ公開鍵基盤かというと，公開鍵を使った暗号通信の要件が[「仮想通貨」]による取引の要件によく似ていると考えられるからだ。
公開鍵を使った暗号通信には以下の4つの要件がある。

1. 機密性（Confidentiality）
2. 完全性（Integrity）
3. 認証（Authentication）
4. 否認防止（Non-repudiation）

暗号なので1番目は言わずもがなだが，2番目は電子署名によって達成できる。
そして3番目を達成する手段として公開鍵基盤がある。

ちなみに2番目と3番目が達成できれば4番目も達成可能なのだが，否認防止の重要性は前節までを見ればお分かりいただけるだろう。

### X.509 の信用モデル

典型的な hierarchical PKI として有名なのが X.509 である。
Web の HTTPS 通信で必要な「電子証明書」は X.509 下で運用されている。

X.509 では認証局（Certification Authority; CA）が公開鍵（の帰属先）を証明する電子証明書を発行する。
電子証明書は具体的には，ユーザの公開鍵に対して認証局の鍵で電子署名を付与したものである。
では，認証局の鍵はどうやって証明するかというと，さらに上位の認証局が証明する。
ただし最上位のルート認証局は誰も証明してくれない（自己証明のみ）。

{{< fig-mermaid >}}
graph TD
  CA1["root CA"]-- Digital Sign -->CA2
  CA1-- Digital Sign -->CA3

  CA2-- Digital Sign -->Aa(("A[a]"))
  CA3-- Digital Sign -->Bb(("B[b]"))
{{< /fig-mermaid >}}

X.509 は「認証局は信用できる」という前提に立った信用モデルと言える。
言い方を変えると，ある認証局が信用できるのであれば配下の認証局やユーザは総て信用できる。

X.509 は大規模かつ安定的な運用に向いているが，いったん認証局の信用が崩れると配下の認証局やユーザの信用が一気に崩れることになる。
そのため，認証局，特にルート認証局では高いセキュリティが要求される。

### [OpenPGP] の信用モデル

[OpenPGP] における典型的な信用モデルは「信用の輪（web of trust）」と呼ばれている[^tm1]。
信用の輪はユーザ間の P2P な関係がベースになっている。

[^tm1]: [OpenPGP] の実装である [GnuPG] では信用の輪以外にも [TOFU (Trust On First Use)](https://en.wikipedia.org/wiki/Trust_on_first_use "Trust on first use - Wikipedia") などの信用モデルを実装している（参考： “{{< pdf-file title="TOFU for OpenPGP" link="ftp://ftp.gnupg.org/people/neal/tofu.pdf" >}}”）。


信用の輪ではユーザ同士がお互いの公開鍵に電子署名を付与する。
たとえば $A$ と $B$ に面識があるなら，相互に電子署名を付与することができる。

{{< fig-mermaid >}}
graph LR
  Aa(("A[a]"))
  Bb(("B[b]"))

  Aa-- Digital Sign -->Bb
  Bb-- Digital Sign -->Aa
{{< /fig-mermaid >}}

ここで3人目の $D$ に登場してもらおう。
$B$ と $D$ は面識があって電子署名を交わしているが， $A$ と $D$ は面識がないものとする。

{{< fig-mermaid >}}
graph LR
  Aa(("A[a]"))
  Bb(("B[b]"))
  Dd(("D[d]"))

  Aa-- Digital Sign -->Bb
  Bb-- Digital Sign -->Aa

  Bb-- Digital Sign -->Dd
  Dd-- Digital Sign -->Bb
{{< /fig-mermaid >}}

この場合でも $A$ と $B$ との関係， $B$ と $D$ との関係をもとに $A$ から見て $D$ も信用できるとみなすのだ。

{{< fig-mermaid >}}
graph LR
  Aa(("A[a]"))
  Bb(("B[b]"))
  Dd(("D[d]"))

  Aa-- Digital Sign -->Bb
  Aa-. validate! .->Dd
  Aa-. trust .->Bb
  Bb-- Digital Sign -->Aa

  Bb-- Digital Sign -->Dd
  Bb-. trust .->Dd
  Dd-- Digital Sign -->Bb
{{< /fig-mermaid >}}

信用の輪はコミュニティ内のアドホックな鍵管理に向いているが，全く関係のない第3者を証明するのは難しい。

### X.509 と [OpenPGP]

山根信二さん等の「{{< pdf-file title="OpenPGPとPKI" link="http://www.baldanders.info/spiegel/archive/pgpdump/PGP-001.pdf" >}}」では X.509 と OpenPGP の PKI の比較を行っている[^pki1]。
以下に比較表を示す。

[^pki1]: この論文は2002年に旧 OpenPKSD.org で公開されたが，サイトそのものが消失したため [Internet Archive](https://web.archive.org/web/20110907063003/http://www.openpksd.org/) からサルベージした。

|                 特徴 | X.509            | OpenPGP            |
| --------------------:| ---------------- | ------------------ |
|           PKI の形態 | hierarchical PKI | trust-file PKI     |
|       公開鍵の認証者 | 専門機関（CA）   | 各ユーザ           |
|               信頼点 | ルート CA        | 利用者自身（面識） |
|       認証の連鎖構造 | ツリー型         | ユーザ中心型       |
| 認証者を認証する根拠 | 利用者による選択 | 利用者自身         |
|   証明書の破棄と管理 | あり             | 不完全             |
|               コスト | 高い             | 低い               |

X.509 と [OpenPGP] の信用モデルはコンセプトが直交しているためどちらが正解とは言えない。
また相互補完的に運用することも可能である。

## [「仮想通貨」]の信用モデルは？

[「仮想通貨」]のアドレスの運用についても，おそらく正解はひとつではなく，さまざまな信用モデルがありうると思う。
また信用モデルを[「仮想通貨」]システム自体に埋め込むのか周辺の（ウォレットや交換所などの）サービスで提供するのかというのも議論の余地があると思う。

[「仮想通貨」]が単なる投機物件ではなく generative な経済活動の技術基盤として生き残っていくことを期待したい。

## ブックマーク

- [OpenPGP 鍵管理に関する考察]({{< ref "/openpgp/openpgp-key-management.md" >}})

[「仮想通貨」]: {{< ref "/remark/2018/01/cryptocurrency-are-not-crypto.md" >}} "「暗号通貨」ってゆーな！"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00FONW2V8/baldandersinf-22/"><img src="https://images-fe.ssl-images-amazon.com/images/I/51AT2LqRIsL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00FONW2V8/baldandersinf-22/">UNDERGROUND MARKET　ヒステリアン・ケース</a></dt><dd>藤井太洋 </dd><dd>朝日新聞出版 2013-11-07</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00H91YQGW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00H91YQGW.09._SCTHUMBZZZ_.jpg"  alt="UNDERGROUND MARKET　アービトレーター"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00V7Y7DUS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00V7Y7DUS.09._SCTHUMBZZZ_.jpg"  alt="ビッグデータ・コネクト (文春文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00B71TZ9S/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00B71TZ9S.09._SCTHUMBZZZ_.jpg"  alt="UNDER GROUND MARKET"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01IGR0TA4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01IGR0TA4.09._SCTHUMBZZZ_.jpg"  alt="アンダーグラウンド・マーケット (朝日文庫)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00CHIFA1M/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00CHIFA1M.09._SCTHUMBZZZ_.jpg"  alt="Gene Mapper -full build-"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0753H9CDS/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0753H9CDS.09._SCTHUMBZZZ_.jpg"  alt="公正的戦闘規範 (ハヤカワ文庫JA)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01FFKMSMC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01FFKMSMC.09._SCTHUMBZZZ_.jpg"  alt="オービタル・クラウド 下 (ハヤカワ文庫JA)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B01FFKMSP4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B01FFKMSP4.09._SCTHUMBZZZ_.jpg"  alt="オービタル・クラウド 上 (ハヤカワ文庫JA)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B078JT42J7/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B078JT42J7.09._SCTHUMBZZZ_.jpg"  alt="アリスマ王の愛した魔物 (ハヤカワ文庫JA)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B077Q9WM8M/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B077Q9WM8M.09._SCTHUMBZZZ_.jpg"  alt="SF飯：宇宙港デルタ３の食料事情 (ハヤカワ文庫JA)"  /></a> </p>
<p class="description">日本で「仮想通貨」が流行る前に登場した傑作。つかエンジニアは全員「UNDERGROUND MARKET」シリーズを読め！</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-02-04">2018-02-04</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
