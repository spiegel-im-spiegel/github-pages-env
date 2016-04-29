+++
date = "2016-04-29T19:49:40+09:00"
description = "今回は小ネタ。JavaScript の数値演算における演算誤差の話。"
draft = false
tags = ["programming", "javascript", "error"]
title = "JavaScipt の演算誤差"

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

今回は小ネタ。
実は最近の JavaScript についてはよく知らなかったりするので（勉強中），間違いとか「最近はもっと簡単にできるよ」とかいった点があったら指摘していただけると助かります。

（そうそう。
[node.js v6 のリリース](https://nodejs.org/en/blog/announcements/v6-release/)おめでとうございます）

JavaScript の数値（Number）型の実体は IEEE754 浮動小数点形式である。
他の言語でよく見るような整数型や decimal/currency 型というのは存在しない。
たとえばある値の百分率を取ろうとして，うっかり

```javascript
function percent(rate) {
	let pc = 100.0 * rate;
	console.log("rate: " + pc + "%");
}

percent(0.0112);
```

なんてなコードを書くと

```text
rate: 1.1199999999999999%
```

などと表示され「はうーん」な感じになってしまう。
これは浮動小数点形式特有の演算誤差で，数値の符号化を2進数を基数として行っているためにどうしてもそうなってしまう[^fp]。

[^fp]: 最近の IEEE754-2008 では10進数を基数とした符号化も標準化されている。

表示の上でこれを回避する方法はいくつかある。
よくあるのは以下の2つ。

1. 演算部分を他の言語（Java や SQL など）で行う。
2. 有効桁数以下を丸める

ビジネスロジックを実装する Entity 以下のクラスを JavaScript 以外の言語で記述している場合は最初のやり方で問題ないだろう。
これができない場合でも，大抵の数値計算は「有効桁数」が仕様として決められている筈なので，2番目のやり方で `Math.round()`, `Math.ceil()`, `Math.floor()` といった関数を組み合わせて表示桁数を調整できる[^cf]。
たとえば先程の `percent()` 関数を小数点以下3位で四捨五入するように変更すると

[^cf]: ググってみると `Math.ceil()` を切り上げ， `Math.floor()` を切り捨てと紹介している記事を見かけるが厳密には間違いである。正しくは `Math.ceil()` は天井関数（ceiling function）で `Math.floor()` は床関数（floor function）である。切り上げや切り捨てとは負値の扱いで挙動が異なるので注意が必要。

```javascript
function percent(rate) {
	let pc = Math.round(100000.0 * rate) / 1000.0;
	console.log("rate: " + pc + "%");
}

percent(0.0112);
```

```text
rate: 1.12%
```

と表示される。

しかし他システムと連携していて有効桁が決まらないとか，要求として値を丸められては困る場合もある。
こういう場合は以下のようにするとよい。
（型を意識してもらうため冗長な書き方になっているがご容赦）

```javascript
function percent(rate) {
	let sRate = rate.toString();
	let digits = sRate.indexOf(".");
	if (digits < 0) {
		digits = 0;
	} else {
		digits = sRate.length - (digits + 1);
	}
	let pc = (100 * sRate.replace(".", "").valueOf()) / Math.pow(10, digits);
	console.log("rate: " + pc + "%");
}

percent(0.0112);
```

これで表示結果は `1.12%` になる（ちなみに入力値が `1` ならちゃんと `100%` になる）。
つまり 0.0112 は $0.0112 = 112 \times 10^{-4}$ と整数部分と小数点以下の桁数に分解できるので，演算は整数部分で行い，最後に小数点以下の桁数分だけ割り算をすればよい。

これは固定小数点形式の数値演算によく似ている。
ポイントは整数同士の演算に変換し最後に桁を揃えることである。
これなら演算誤差は生じない（ただし桁あふれに注意）。

余談だが `percent()` 関数では `100` は固定値なので

```javascript
function percent(rate) {
	let sRate = rate.toString();
	let digits = sRate.indexOf(".");
	if (digits < 0) {
		digits = 0;
	} else {
		digits = sRate.length - (digits + 1);
	}
	let pc = (sRate.replace(".", "")+"00").valueOf() / ("1"+"0".repeat(digits)).valueOf();
	console.log("rate: " + pc + "%");
}

percent(0.0112);
```

とすれば割り算1回で済む（数学関数と文字列操作のどちらがコストが高いかは微妙な気もするが）。




