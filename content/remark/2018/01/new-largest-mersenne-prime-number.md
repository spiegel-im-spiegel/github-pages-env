+++
title = "2^77232917-1 is Prime!"
date =  "2018-01-23T21:06:29+09:00"
update =  "2018-02-15T07:19:42+09:00"
description = "酔狂なことに今回のメルセンヌ素数 2^77232917-1 をそのまま書き記した本が出ているらしい。"
image = "/images/attention/remark.jpg"
tags = [ "book", "math", "prime-number" ]

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

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%E3%81%AE%E7%A7%98%E5%AF%86%E3%83%8E%E3%83%BC%E3%83%88%EF%BC%8F%E6%95%B4%E6%95%B0%E3%81%A7%E9%81%8A%E3%81%BC%E3%81%86-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00L0PDMJ0?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00L0PDMJ0"><img src="https://images-fe.ssl-images-amazon.com/images/I/4186Q-UqrDL._SL160_.jpg" width="111" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%95%B0%E5%AD%A6%E3%82%AC%E3%83%BC%E3%83%AB%E3%81%AE%E7%A7%98%E5%AF%86%E3%83%8E%E3%83%BC%E3%83%88%EF%BC%8F%E6%95%B4%E6%95%B0%E3%81%A7%E9%81%8A%E3%81%BC%E3%81%86-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B00L0PDMJ0?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B00L0PDMJ0">数学ガールの秘密ノート／整数で遊ぼう</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2013-12-17 (Release 2014-07-24)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B00L0PDMJ0</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description"><a href='https://baldanders.info/blog/000670/'>小中学生にお薦め</a>。小学生高学年くらいならギリで理解可能と思われ。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2014-09-26">2014-09-26</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
