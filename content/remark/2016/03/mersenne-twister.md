+++
date = "2016-03-17T21:41:22+09:00"
description = "「ズンドコキヨシ」で興味が出たので [Mersenne Twister] について調べている。適宜追加予定。"
draft = false
tags = ["mersenne-twister", "math", "random", "programming"]
title = "Mersenne Twister に関する覚え書き"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

[「ズンドコキヨシ」で興味が出た](http://qiita.com/spiegel-im-spiegel/items/6a5bc07dbfa46a328e26 "「ズンドコキヨシ」と擬似乱数 - Qiita")ので [Mersenne Twister] について調べている。
適宜追加予定。

[Mersenne Twister] とは[松本眞](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/ "Makoto Matsumoto Home Page")・西村拓士両氏によって1996年に発表された擬似乱数生成アルゴリズムである。
他の擬似乱数生成アルゴリズムと比べて以下の特徴があるそうだ。
（「[Mersenne Twister とは?](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/what-is-mt.html "What & how is MT?")」より）

- 従来の様々な生成法の欠点を考慮して設計されています。
- 従来にない長周期, 高次元均等分布を持ちます。（周期が $2^{19937}-1$ で[^mt1]、623次元超立方体の中に 均等に分布することが証明されています。）
- 生成速度がかなり速い。（処理系にもよりますが、パイプライン処理やキャッシュメモリ のあるシステムでは、Cの標準ライブラリの `rand()` より高速なこと もあります。なお、開発当時には cokus 版は `rand()` より4倍程度高速でしたが、その後 ANSI-C の `rand()` が LCG 法から lagged-fibonacci に 変更されたこともあり、2002年現在 rand と MT の速度差はあまりありません。）
- メモリ効率が良い。（32ビット以上のマシン用に設計された `mt19937.c` は、 624ワードのワーキングメモリを消費するだけです。 1ワードは32ビット長とします。

[^mt1]: $2^{19937}-1$ はメルセンヌ素数で [Mersenne Twister] の名前の由来になっている。 [Mersenne Twister] では周期サイズごとに複数の実装があるが， $2^{19937}-1$ がポピュラーな実装として広く使われているようだ。

[Mersenne Twister] が主に使われるのは科学シミュレーション（最近流行りのモンテカルロ法とか）だが，比較的メモリ効率がよいためゲームなどでも使われるらしい[^mt2]。
また [JIS Z 9031](http://kikakurui.com/z9/Z9031-2012-01.html) の附属書 B にも [Mersenne Twister] のコードが掲載されている。
改良版の [SFMT (SIMD-oriented Fast Mersenne Twister)](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/SFMT/index-jp.html) や $2^{127}-1$ 周期の軽量版 [TinyMT](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/TINYMT/index-jp.html) などがある。

[^mt2]: [Mersenne Twister] は「予測可能」であるため暗号技術など高いセキュリティが要求される場合には使えない。暗号技術における乱数生成器の要件については [RFC 4086] （[IPA による日本語訳](https://www.ipa.go.jp/security/rfc/RFC4086JA.html)）などが参考になる。なお [Mersenne Twister] を応用した [CryptMT](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/CRYPTMT/index-jp.html) というのはある。

オリジナルのコードは [GitHub] で公開されている。

- [MersenneTwister-Lab](https://github.com/MersenneTwister-Lab)
    - [SFMT](https://github.com/MersenneTwister-Lab/SFMT "MersenneTwister-Lab/SFMT: SIMD-oriented Fast Mersenne Twister")
    - [dSFMT](https://github.com/MersenneTwister-Lab/dSFMT "MersenneTwister-Lab/dSFMT: Double precision SIMD-oriented Fast Mersenne Twister") （倍精度浮動小数点擬似乱数を直接生成できる）
    - [TinyMT](https://github.com/MersenneTwister-Lab/TinyMT "MersenneTwister-Lab/TinyMT: Tiny Mersenne Twister")

主に C 言語で記述されており BSD ライセンスで提供されている[^mt3]。
また C++, PHP, Python, Ruby などの言語では標準で組み込まれている。

[^mt3]: [MIT ライセンスでの利用も可能](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/MT2002/license.html "Mersenne Twisterの商用について")らしい。

- [std::mersenne_twister_engine - cppreference.com](http://en.cppreference.com/w/cpp/numeric/random/mersenne_twister_engine)
- [PHP: mt_srand - Manual](http://php.net/manual/en/function.mt-srand.php)
- [9.6. random — Generate pseudo-random numbers — Python 3.3.6 documentation](https://docs.python.org/3.3/library/random.html)
    - [9.6. random — Generate pseudo-random numbers — Python 2.7.11 documentation](https://docs.python.org/2.7/library/random.html)
- [Class: Random (Ruby 2.3.0)](http://ruby-doc.org/core-2.3.0/Random.html)

これら以外では Java や [Go] などによる実装がある。

- [TinyMT Java Implementation (Japanese)](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/TINYMT/JAVA/index-jp.html) （オリジナル）
- [Sean Luke : Code](http://cs.gmu.edu/~sean/research/) に [Mersenne Twister] の Java 実装が紹介されている
- [seehuhn/mt19937: An implementation of Takuji Nishimura's and Makoto Matsumoto's Mersenne Twister pseudo random number generator in Go.](https://github.com/seehuhn/mt19937) （GPL ライセンスなので取り扱いに注意）

## ブックマーク

- {{< pdf-file title="有限体の擬似乱数への応用" link="http://www.math.sci.hiroshima-u.ac.jp/~m-mat/TEACH/0407-2.pdf" >}}
- {{< pdf-file title="間違いだらけの疑似乱数選び" link="http://www.soi.wide.ad.jp/class/20010000/slides/03/sfc2002.pdf" >}}
- [良い乱数・悪い乱数](http://www001.upp.so-net.ne.jp/isaku/rand.html)
- [Multiple stream Mersenne Twister PRNG](http://theo.phys.sci.hiroshima-u.ac.jp/~ishikawa/PRNG/README.jp.html)

[Mersenne Twister]: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/mt.html "Mersenne Twister: A random number generator (since 1997/10)"
[GitHub]: https://github.com/ "GitHub"
[RFC 4086]: http://tools.ietf.org/html/rfc4086 "RFC 4086 - Randomness Requirements for Security"
[Go]: https://golang.org/ "The Go Programming Language"
