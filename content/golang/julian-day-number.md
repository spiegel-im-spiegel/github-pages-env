+++
date = "2015-09-18T14:52:22+09:00"
description = "今回は暦で遊びます。とりあえず，簡単なところで「ユリウス日」をいってみるか。"
draft = true
tags = ["golang", "julian-day-number", "math", "time", "astronomy"]
title = "「ユリウス日」で遊ぶ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その5 - Qiita](http://qiita.com/spiegel-im-spiegel/items/e743d63ef5165d750eff)）

今回は暦で遊びます。
とりあえず，簡単なところで「ユリウス日（Julian Date または Julian Day Number[^1]）」をいってみるか。

[^1]: ちなみに1日未満の端数を含む場合を「ユリウス日（Julian Date）」，端数を含まない場合を「ユリウス通日（Julian Day Number）」と呼び分けているようです。

## ユリウス日の定義

ユリウス日とは紀元前4713年1月1日正午[^2] を起点とした通日のことです。
たとえば2015年1月1日零時はユリウス日では2,457,023.5日になります。

- [ユリウス日](http://eco.mtk.nao.ac.jp/cgi-bin/koyomi/cande/date2jd.cgi) （[国立天文台暦計算室](http://eco.mtk.nao.ac.jp/koyomi/)）
- [ユリウス日(Julian Day)](http://homepage1.nifty.com/manome/astrology/julian.html) : 暦にまつわるエピソードを含めて参考になります
- [The Julian Period](http://www.tondering.dk/claus/cal/julperiod.php)

[^2]: もともと太陽暦は「子午線から太陽までの時角」が時刻のベースになってるため，ユリウス日を考えた人は正午を起点にすべきと考えたのでしょう。日常生活で昼に日付が変わったら色々面倒そうですが。あぁでも， B な企業に勤めている人には日付の起点とか関係ないかな（笑）

この際，いくつか気を付ける点があります。

- 紀元前（BC）1年の翌年は紀元後（AD）1年。つまり AD1年を $1$ とするなら BC1年は $0$，BC2年は $-1$ となります。紀元前4713年は $-4712$ です。
- ユリウス暦は紀元前45年から開始されたと言われています。つまりそれ以前は異なる暦だったわけです。しかしユリウス日では紀元前45年より前の日付もユリウス暦と見なして取り扱います。
- 欧州では西暦（紀元後）1582年に当時のローマ法王グレゴリオ13世によって（現在言われるところの）グレゴリオ暦が布告されましたが，この際に1582年10月4日の翌日を10月15日としたためギャップが生じました。

## 年月日から（修正）ユリウス日を求める

というわけで大昔のユリウス日を求めるのは西暦を使う場合でもちょっと面倒くさいのですが，範囲をグレゴリオ暦に限るなら便利な式があります。

グレゴリオ暦のある日を「$Y$ 年 $M$ 月 $D$ 日」で表せるとすると

<blockquote>
\begin{aligned}
    y   & = Y + \left\lfloor \frac{M - 3}{12} \right\rfloor \\
	m   & = \left( 12 + \left( M - 3 \right) \right) \bmod 12 \\
	d   & = D - 1 \\
    MJD & = \left\lfloor 365.25y \right\rfloor + \left\lfloor \frac{y}{400} \right\rfloor - \left\lfloor \frac{y}{100} \right\rfloor + \left\lfloor 30.60m + 0.5 \right\rfloor + d - 678881 \\
    JD  & = MJD + 2400000.5
\end{aligned}
</blockquote>


でユリウス日 $JD $ を求めることができます[^3] [^4] [^5]。

[^3]: 上の式は「Fliegel の公式」などと呼ばれることがありますが，厳密には Fliegel の公式を電卓向けに分かりやすく展開したもので，初等天文学の教科書などでよく登場する式です。
[^4]: 余談ですが365.25日を「ユリウス年」と呼びます。天文学では1年の長さが年によって変わるのは困るので，一様な長さの「ユリウス年」を考えたわけです。つまり上の式はユリウス年にうるう年の補正をかけてるわけですね。ちなみに「ユリウス世紀」は36525日です。
[^5]: 月の値から2を引くのは暦計算の基本的なテクニックだったります（式では0基点にするために3を引いてますが）。現在の1月（Ianuarius または January）を年初としたのはユリウス暦以降からで，それまでは現在の3月（Martius または March）が年初でした。だから2月だけちょっと特殊なんですねぇ。

ちなみに $MJD$ は修正ユリウス日（Modified Julian Date）と呼ばれるものです。
定義は上の式の通りで，ユリウス日から240万日分をカットして日付の起点を正午から（私たちになじみのある）正子[^6] にずらしています。

[^6]: 耳慣れないかもしれないですが，夜中の12時のことです。

$ \left\lfloor x \right\rfloor $ は床関数と呼ばれるもので「実数 $x$ に対して $x$ 以下の最大の整数」と定義されます。
例えば

<blockquote>
\begin{aligned}
	\left\lfloor 1.0 \right\rfloor & = 1 \\
	\left\lfloor 0.7 \right\rfloor & = 0 \\
	\left\lfloor -0.5 \right\rfloor & = -1 \\
	\left\lfloor -2.0 \right\rfloor & = -2 \\
\end{aligned}
</blockquote>

となります。
単に小数点を取るだけではないです。
[Go 言語]の [`math`] パッケージには，そのものずばりの `math.Floor()` 関数があります。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("[1.0] = %v\n", math.Floor(1.0))
	fmt.Printf("[0.7] = %v\n", math.Floor(0.7))
	fmt.Printf("[-0.5] = %v\n", math.Floor(-0.5))
	fmt.Printf("[-2.0] = %v\n", math.Floor(-2.0))
}
```

```
C:>go run floor.go
[1.0] = 1
[0.7] = 0
[-0.5] = -1
[-2.0] = -2
```

では換算式を [Go 言語]で実装してみましょう。

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	mm := float64(month - 3)
	y := float64(year) + math.Floor(mm/12.0)
	m := math.Mod(12.0+mm, 12.0)
	d := float64(day - 1)
	fmt.Printf("y = %f\n", y)
	fmt.Printf("m = %f\n", m)
	fmt.Printf("d = %f\n\n", d)

	mjd := math.Floor(365.25*y) + math.Floor(y/400.0) - math.Floor(y/100.0) + math.Floor(30.60*m+0.5) + d - 678881.0
	fmt.Printf("MJD = %f日\n", mjd)
	fmt.Printf("JD = %f日\n", mjd+2400000.5)
}
```

```
C:>go run julian-day-1.go
2015年1月1日

y = 2014.000000
m = 10.000000
d = 0.000000

MJD = 57023.000000日
JD = 2457023.500000日
```

と，これではあまりにひどいので，少し変形。

床関数は正の値に対しては小数点以下の切り捨てと同じ。
Go 言語では int 型の乗除算には int 型の結果が返り小数点以下が切り捨てられることを利用して

```go
package main

import "fmt"

func main() {
	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	y := 1
	m := 1
	if month < 3 {
		y = year - 1
		m = month + 9
	} else {
		y = year
		m = month - 3
	}
	d := day - 1
	fmt.Printf("y = %d\n", y)
	fmt.Printf("m = %d\n", m)
	fmt.Printf("d = %d\n\n", d)

	mjd := (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
```

```
C:>go run julian-day-2.go
2015年1月1日

y = 2014
m = 10
d = 0

MJD = 57023日
JD = 2457023.500000日
```

これなら [`math`] パッケージ自体不要になります。グレゴリオ暦は1582年より前では適用できないのでこれで必要十分です。
以降で使いやすくするために，ここから更に変形して年月日を引数から取得するようにします（引数の Validation は省いています。ゴメンペコン）。

```go
package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int64, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.ParseInt(argsStr[i], 10, 64)
		if err != nil {
			if enum, ok := err.(*strconv.NumError); ok {
				switch enum.Err {
				case strconv.ErrRange:
					fmt.Printf("Bad Range Error")
				case strconv.ErrSyntax:
					fmt.Printf("Syntax Error")
				}
			}
			return
		} else {
			args[i] = num
		}
	}
	fmt.Printf("%v年%v月%v日\n\n", args[0], args[1], args[2])

	y := args[0]
	m := args[1]
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := args[2] - 1
	fmt.Printf("y = %d\n", y)
	fmt.Printf("m = %d\n", m)
	fmt.Printf("d = %d\n\n", d)

	mjd := (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
```

```shell
C:>go run julian-day-2b.go 2015 1 1
2015年1月1日

y = 2014
m = 10
d = 0

MJD = 57023日
JD = 2457023.500000日
```

## UNIX Time からユリウス日を求める

さて，これでグレゴリオ暦の任意の日付からユリウス日を求めることができるようになりました。
これを踏まえて，もう少し簡単にユリウス日を得る方法を考えてみます。

Go 言語で時刻情報を取得・操作するために [`time`] パッケージが用意されています。
[`time`] パッケージをつらつら眺めてみると `Unix()` 関数を使って UNIX Time を得ることができるようです。

> Unix returns t as a Unix time, the number of seconds elapsed since January 1, 1970 UTC.

ということは，1970年1月1日零時 UTC のユリウス日が分かれば，そこを起点に UNIX Time を加算すればいいことになります。
簡単！


```
C:>go run julian-day-2b.go 1970 1 1
1970年1月1日

y = 1969
m = 10
d = 0

MJD = 40587日
JD = 2440587.500000日
```

なので，こうなります。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	const onday = int64(86400)   //seconds
	const baseDay = int64(40587) //Modified Julian Date at January 1, 1970

	year := 2015
	month := 1
	day := 1
	fmt.Printf("%v年%v月%v日\n\n", year, month, day)

	tm := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	ut := tm.Unix()
	fmt.Printf("UNIX Time = %v seconds = %v days and %v seconds\n\n", ut, ut/onday, ut%onday)

	mjd := ut/onday + baseDay
	fmt.Printf("MJD = %d日\n", mjd)
	fmt.Printf("JD = %f日\n", float64(mjd)+2400000.5)
}
```

```
C:>go run julian-day-3.go
2015年1月1日

2015-01-01 00:00:00 +0000 UTC
UNIX Time = 1420070400 seconds = 16436 days and 0 seconds

MJD = 57023日
JD = 2457023.500000日
```

えーっと。
コードを見ればお分かりと思いますが，これだと1970年1月1日より前の日付では正しく動きません。
任意の日付[^7] で正しく動かすには床関数を使います。

[^7]: とはいえ UNIX Time の取りうる値の範囲内での話ですが。

次回は，これをパッケージ化してみましょう。

## ブックマーク

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[`math`]: http://golang.org/pkg/math/
[`time`]: http://golang.org/pkg/time/
