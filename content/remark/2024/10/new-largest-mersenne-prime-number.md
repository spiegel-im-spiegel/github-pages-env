+++
title = "2^136279841-1 is Prime!"
date =  "2024-10-23T07:59:26+09:00"
description = "41,024,320桁の素数！"
image = "/images/attention/kitten.jpg"
tags = [ "math", "prime-number" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

発見された素数の最大数が更新されたようだ。
ついに $p$ 値が1億超えたか。

- [Mersenne Prime Discovery - 2^136279841-1 is Prime!](//www.mersenne.org/primes/?press=M136279841)

「メルセンヌ素数（Mersenne prime number）」というのは $M_p = 2^p-1$ で表されるメルセンヌ数[^mn1] $M_p$ のうち素数であるものを指す。
$M_p$ が素数なら $p$ も素数であるという面白い性質がある（逆は成り立たない）。
ちなみに $M_p = 2^p-1$ が素数なら $2^{p-1}(2^p-1)$ は完全数[^pn] になる。

[^mn1]: メルセンヌ数の定義（$M_p = 2^p-1$）を見るとなにやら難しそうだが，2進数で考えると簡単である。つまり $p$ 桁の2進数のビットが全て立っている状態がメルセンヌ数である。たとえば 1B, 11B, 111B,... はメルセンヌ数（の定義そのもの）だ。
[^pn]: 「完全数（perfect number）」とは「その数自身を除く約数の和がその数自身と等しい自然数」を指す。たとえば $6$ の素因数は $2\times3$ なので $6$ 自身を除く約数の和は $1+2+3=6$ となり完全数と言える。ちなみに $3$ は $2^2-1$ とメルセンヌ素数になっていて， $2^1(2^2-1)=6$ である。

素数探索プロジェクトで最大のものが [GIMPS (Great Internet Mersenne Prime Search)](https://www.mersenne.org/) プロジェクトである。
[GIMPS] は分散コンピューティング・プロジェクトで，世界中のコンピュータを使ってメルセンヌ素数を探索している。
メルセンヌ数に対しては効率的な素数判定法が知られており分散コンピューティング向きの題材と言える。
今回のトピックは GPU を活用した分散コンピューティングソフトウェアを使った初の発見というところだろうか。

なお，メルセンヌ素数 $2^{136,279,841}-1$ を書き記したテキスト・ファイルを zip 圧縮したものが [GIMPS] のサイトから入手できる。
ダウンロードしたら 19MB 以上あったよ（笑）

## ブックマーク

- [“最も大きい素数”更新　「2^1億3627万9841-1」元NVIDIA社員が発見　文字に起こすと4000万字超え【訂正あり】 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2410/22/news186.html)
- [完全数について12歳の長男と話す（思い出の日記）｜結城浩](https://mm.hyuki.net/n/nb17107856e67)

[GIMPS]: https://www.mersenne.org/ "Great Internet Mersenne Prime Search - PrimeNet"

## 参考図書

{{% review-paapi "B00L0PDMJ0" %}} <!-- 数学ガールの秘密ノート／整数で遊ぼう -->
