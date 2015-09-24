# はじめての Go 言語 (on Windows) その6

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa#%E7%9B%AE%E6%AC%A1)）

[前回](http://qiita.com/spiegel-im-spiegel/items/e743d63ef5165d750eff)のユリウス日の続き。なんだけど、今回はパッケージのお話。

## ユリウス日計算関数をパッケージ化する

ユリウス日の計算なんて簡単なので今まで `main()` の中にゴリゴリ書いてたんだけど、今後のことを考えてパッケージ化を行うことにします。

まずは、[前回](http://qiita.com/spiegel-im-spiegel/items/e743d63ef5165d750eff)のコードから計算処理部分をきちんと分離します。

```go:jd4.go
package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

func ModifiedJulianDayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return mjdnUnix(t)
	} else {
		return mjdnGregorian(t)
	}
}

func mjdnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func mjdnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
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
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", ModifiedJulianDayNumber(tm))
}
```

```shell
C:>go run jd4.go 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:>go run jd4.go 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:>go run jd4.go 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

ユリウス日の端数が鬱陶しいので修正ユリウス日の整数部分のみ計算しています[^1]。あと1970年1月1日を境界として計算方法を変えています。本当はユリウス暦の場合の計算も含めるべきなんでしょうけど、今回は無視。

[^1]: ちなみに1日未満の端数を含む場合を「ユリウス日（Julian Date）」、端数を含まない場合を「ユリウス通日（Julian Day Number）」と呼び分けているようです。

さて上のコードのうち修正ユリウス日計算関数を別ファイルにしてパッケージ化します。

```go:modjulian.go
package modjulian

import "time"

func DayNumber(t time.Time) int64 {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	} else {
		return dnGregorian(t)
	}
}

func dnGregorian(t time.Time) int64 {
	y := int64(t.Year())
	m := int64(t.Month())
	if m < 3 {
		y -= 1
		m += 9
	} else {
		m -= 3
	}
	d := int64(t.Day()) - 1
	return (1461*y)/4 + y/400 - y/100 + (153*m+2)/5 + d - 678881
}

func dnUnix(t time.Time) int64 {
	const (
		onday   = int64(86400) //seconds
		baseDay = int64(40587) //Modified Julian Date at January 1, 1970
	)
	return (t.Unix())/onday + baseDay
}
```

パッケージ名は `modjulian` としました。ちなみにパッケージ内の関数等は名前の先頭が大文字のものだけが外部から参照可能です。

### Go 言語における名前の問題

Go 言語およびそのコミュニティは名前にうるさいようです。たとえば「[実践Go言語](http://golang.jp/effective_go)」によると

> 慣例では、パッケージ名は小文字でひとつの単語です。アンダースコアや大文字が混ざって(mixedCaps)はいけません。パッケージ使用者がその名前をタイプすることを考慮して、簡潔すぎるぐらいにしてください。

なんだそうで、本当はパッケージ名を `julianDate` としたかったんだけど、それじゃだめらしい。

もうひとつ。

> もう一つの慣例は、パッケージ名がそのソースディレクトリのベース名であるということです。たとえばsrc/pkg/encoding/base64に置かれているパッケージは、"encoding/base64"としてインポートし、名前はbase64となります。encoding_base64やencodingBase64とはなりません。

なのでパッケージの指定はハンガリアン記法とかではなく単語をディレクトリで区切って階層化することで整理できそうです。これって Java とかに慣れてる人には比較的とっつきやすき仕組みかもしれません。

> 長い名前は、慣れたとしても読みやすくなることはありません。複雑もしくは微妙なニュアンスを持つものに名前をつけるときは、すべての情報を名前で表現しようとするより、通常は役立つドキュメントコメントを書いたほうがよいでしょう。

## パッケージの配置と呼び出し

パッケージを呼び出すのには（標準のパッケージと同じく） `import` を使えばいいのですが、パッケージを記述したファイルをどこに配置するかで多少記述が変わります。

たとえば

```go
import "./modjulian"
```

と記述した場合は、呼び出し元のファイルの場所にある `modjulian` フォルダを探します。一方

```go
import "modjulian"
```

と記述した場合には、環境変数 `GOROOT` および `GOPATH` で指定されるフォルダ以下を探すようです。 

```shell
C:>go run jd4c.go 1969 12 31
jd4c.go:4:2: cannot find package "modjulian" in any of:
        C:\goroot\src\modjulian (from $GOROOT)
        C:\gopath\src\modjulian (from $GOPATH)
```

まずは前者で行きましょう。カレント・フォルダに `modjulian` フォルダを掘って、先ほどの `modjulian.go` を入れます。呼び出し側のコードは以下のようにします。

```go:jd4b.go
package main

import (
	"./modjulian"
	"flag"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
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
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

```shell
C:>go run jd4b.go 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:>go run jd4b.go 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:>go run jd4b.go 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

同じ結果が得られることを確認しました。

### パッケージを GitHub 上に配置する

インターネット上にあるパッケージを `import` することもできます。この場合、以下のように指定します。

```go
import "github.com/spiegel-im-spiegel/astrocalc/modjulian"
```

ここでは GitHub の [spiegel-im-spiegel/astrocalc](https://github.com/spiegel-im-spiegel/astrocalc) リポジトリに `modjulian` パッケージを置いてみました。

ただし build や run コマンドではリポジトリ上のパッケージを自動的にとってきません。まずはパッケージを get します。（get コマンドについては「[その3](http://qiita.com/spiegel-im-spiegel/items/a52a47942fd3946bb583)」を参照）

```shell
C:>go get -v github.com/spiegel-im-spiegel/astrocalc/modjulian
github.com/spiegel-im-spiegel/astrocalc (download)
github.com/spiegel-im-spiegel/astrocalc/modjulian
```

これで `%GOPATH%` 上にパッケージを取ってきました。

呼び出し側は `import` 以外は先ほどと同じです。

```go:jd4c.go
package main

import (
	"flag"
	"fmt"
	"github.com/spiegel-im-spiegel/astrocalc/modjulian"
	"strconv"
	"time"
)

func main() {
	//引数のチェック
	flag.Parse()
	argsStr := flag.Args()
	if len(argsStr) < 3 {
		fmt.Printf("年月日を指定してください")
		return
	}
	args := make([]int, 3)
	for i := 0; i < 3; i++ {
		num, err := strconv.Atoi(argsStr[i])
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
	tm := time.Date(args[0], time.Month(args[1]), args[2], 0, 0, 0, 0, time.UTC)
	fmt.Printf("%v\n", tm)
	fmt.Printf("MJD = %d日\n", modjulian.DayNumber(tm))
}
```

```shell
C:>go run jd4c.go 1969 12 31
1969-12-31 00:00:00 +0000 UTC
MJD = 40586日

C:>go run jd4c.go 1970 1 1
1970-01-01 00:00:00 +0000 UTC
MJD = 40587日

C:>go run jd4c.go 2015 1 1
2015-01-01 00:00:00 +0000 UTC
MJD = 57023日
```

同じ結果が得られることを確認しました。

## さて...

パッケージについてもう少し掘り下げるべきか。それとも暦計算に戻るか。

## 脚注
