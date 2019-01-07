+++
title = "2^77232917-1 is Prime!"
date =  "2018-01-23T21:06:29+09:00"
update =  "2018-02-15T07:19:42+09:00"
description = "酔狂なことに今回のメルセンヌ素数 2^77232917-1 をそのまま書き記した本が出ているらしい。"
image = "/images/attention/remark.jpg"
tags = [ "book", "math", "prime-number" ]

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

今月初めの話で（アンテナ感度が低くて）申し訳ないが，素数の最大桁数が更新されたようだ（発見自体は2017年末）。

- [Mersenne Prime Discovery - 2^77232917-1 is Prime!](https://www.mersenne.org/primes/?press=M77232917)

「メルセンヌ素数（Mersenne prime number）」というのは $M_p = 2^p-1$ で表されるメルセンヌ数[^mn1] $M_p$ のうち素数であるものを指す。
$M_p$ が素数なら $p$ も素数であるという面白い性質がある（逆は成り立たない）。
ちなみに $M_p = 2^p-1$ が素数なら $2^{p-1}(2^p-1)$ は完全数[^pn] になる。

[^mn1]: メルセンヌ数の定義（$M_p = 2^p-1$）を見るとなにやら難しそうだが，2進数で考えると簡単である。つまり $p$ 桁の2進数のビットが全て立っている状態がメルセンヌ数である。たとえば 1B, 11B, 111B,... はメルセン数（の定義そのもの）だ。
[^pn]: 「完全数（perfect number）」とは「その数自身を除く約数の和がその数自身と等しい自然数」を指す。たとえば $6$ の素因数は $2\times3$ なので $6$ 自身を除く約数の和は $1+2+3=6$ となり完全数と言える。ちなみに $3$ は $2^2-1$ とメルセンヌ素数になっていて， $2^1(2^2-1)=6$ である。

実は大きな素数には懸賞金がかけられている。
現在は1億桁以上の素数に対して15万ドルの懸賞金がかけられているそうだ。
まぁ，当分先の話かな。

素数探索プロジェクトで最大のものが [GIMPS (Great Internet Mersenne Prime Search)](https://www.mersenne.org/) プロジェクトである。
[GIMPS] は分散コンピューティング・プロジェクトで，世界中のコンピュータを使ってメルセンヌ素数を探索している。
メルセンヌ数に対しては効率的な素数判定法が知られており分散コンピューティング向きの題材と言える。

酔狂なことに今回のメルセンヌ素数 $2^{77,232,917}-1$ をそのまま書き記した[本が出ている](https://www.amazon.co.jp/exec/obidos/ASIN/4909045074/baldandersinf-22/ "2017年最大の素数 | 虹色社 |本 | 通販 | Amazon")らしい[^cr1]。
しかもこの記事を書いている時点（2018-01-23）で Amazon で在庫切れを起こしているようだ。
買う方も随分と酔狂なことである。

[^cr1]: Amazon の紹介文には「GIMPS (Great Internet Mersenne Prime Search)およびJonathan Pace氏により2017年12月26日に発見された50番目のメルセンヌ素数を、そのまま引用掲載」と書かれているがメルセンヌ素数自体は「表現」ではないので誰に対しても著作権は発生しない。従って引用もへったくれもない。書籍出版社なら言葉は正しく使いましょう。

なお，メルセンヌ素数 $2^{77,232,917}-1$ を書き記したテキスト・ファイルを zip 圧縮したものが [GIMPS] のサイトから入手できる。
圧縮した状態で 10MB 以上あるとか（笑）

## ブックマーク

- [1億桁の賞金まであと1歩！ FedEx社員が50番目の素数見つける。2325万桁 | ギズモード・ジャパン](https://www.gizmodo.jp/2018/01/fedex-employee-discovers-largest-known-prime-number.html)
- [完全数について12歳の長男と話す（思い出の日記）｜結城浩](https://mm.hyuki.net/n/nb17107856e67)

[GIMPS]: https://www.mersenne.org/ "Great Internet Mersenne Prime Search - PrimeNet"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMJ0/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51-nVSeXaNL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMJ0/baldandersinf-22/">数学ガールの秘密ノート／整数で遊ぼう</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ株式会社 2014-07-24</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMIQ/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMIQ.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート／式とグラフ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00L0PDMK4/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00L0PDMK4.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ガロア理論"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00NAQA33A/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00NAQA33A.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの誕生 　理想の数学対話を求めて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1FO/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1FO.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／乱択アルゴリズム"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00I8AT1D6/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00I8AT1D6.09._SCTHUMBZZZ_.jpg"  alt="数学ガール／ゲーデルの不完全性定理"  /></a> </p>
<p class="description" ><a href='https://baldanders.info/spiegel/log2/000670.shtml'>小中学生にお薦め</a>。小学生高学年くらいならギリで理解可能と思われ。</p>
<p class="gtools" >reviewed by <a href="#maker" class="reviewer">Spiegel</a> on <abbr class="dtreviewed" title="2014-09-26">2014/09/26</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html">G-Tools</a>)</p>
</div>
