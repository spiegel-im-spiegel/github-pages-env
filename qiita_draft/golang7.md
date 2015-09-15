# はじめての Go 言語 (on Windows) その7

（[これまでの記事の目次](http://qiita.com/spiegel-im-spiegel/items/dca0df389df1470bdbfa#%E7%9B%AE%E6%AC%A1)）

[前回](http://qiita.com/spiegel-im-spiegel/items/404871d2bafd22bdbb90)の続き。パッケージ化までしたのならテストまで行っちゃおうかと。

## テストコードを書く

そもそもテストコードを書く前に実際のコードを書くなんてのは邪道なのですが，今まで「正しく動かす」ことはあまり意識してなかったので。ゴメンペコン。

Go 言語では最初からテスト・フレームワークが同梱されています。最近の言語はみんなそうですよね。テストコードを書くには対象のソースファイルと同じフォルダに `*_test.go` という名前のファイルを用意します。まぁ，説明するより書いた方が早いですね。

```go:modjulian_test.go
package modjulian

import (
	"os"
	"testing"
	"time"
)

type mjdnTest struct { //test case for DayNumber
	in time.Time //input data
	out int64 //expected result
}

var mjdnTests []mjdnTest  //test cases for DayNumber

func TestMain(m *testing.M) {
	//initialization
 	mjdnTests = []mjdnTest {  //test cases for DayNumber
		{time.Date(1969, 12, 31, 0, 0, 0, 0, time.UTC), int64(40586)},
		{time.Date(1970,  1,  1, 0, 0, 0, 0, time.UTC), int64(40587)},
		{time.Date(2015,  1,  1, 0, 0, 0, 0, time.UTC), int64(57023)},
	}

	//start test
    code := m.Run()

	//termination
    os.Exit(code)
}

func TestModifiedJulianDayNumber(t *testing.T) {
	for _, testCase := range mjdnTests {
		result := DayNumber(testCase.in)
		if result != testCase.out {
			t.Errorf("DayNumber of \"%v\" = %d, want %d.", testCase.in, result, testCase.out)
		}
	}
}
```

1. `package` にはテスト対象のパッケージを指定します。
1. `import` には [testing](http://golang.org/pkg/testing/) パッケージを含めます。
1. `Test...` で始まる関数名がテスト実行用の関数です。引数には `t *testing.T` を指定します。
1. `TestMain()` は特別な関数です。テストの最初に呼び出され， `Run()` で他のテスト関数（群）をキックします。引数には `m *testing.M` を指定します。 `TestMain()` 内で初期化や条件を変えたテストの繰り返しや後始末処理などを行うことができます。

[testing](http://golang.org/pkg/testing/) パッケージには，他の言語のテスト・フレームワークによくある [assertion 関数がありません](http://golang.jp/go_faq#assertions)。 [FAQ](http://golang.jp/go_faq#testing_framework) によると

> 一般的なテストフレームワークにおいて条件・制御・出力機構を持つ専用のミニ言語が用意される傾向がありますが、Go言語にはすでにこれらが備わっています。これらを再び作成するより、我々はGo言語のテストを進めたかったのです。このようにしたことで余計な言語を覚える必要がなくなり、テストを直接的かつ理解しやすくしています。

とあります。テスト駆動型開発の場合，テストコードはそれ自体が設計書として機能しますので，この割り切りは妥当と言えます[^1]。その代わりテストコードを（ドキュメントとして）きちんと書くのは骨が折れますが（笑）

[^1]: 私は組み込みエンジニアなので，プログラミングで assert を多用するのは，エンジニアの怠慢だと思ってしまいます。まぁ，ベクタ・テーブルからゴリゴリ書くってのなら別ですが。

テストコードが書けたので早速動かしてみましょう。テストを行うには `test` コマンドを使います。また引数に `./...` と指定すればカレント・フォルダ以下の全てのテストが対象になります。

```shell
C:>go test -v ./...
=== RUN TestModifiedJulianDayNumber
--- PASS: TestModifiedJulianDayNumber (0.00s)
PASS
ok      _/C_/pathto/astrocalc/modjulian 0.255s
```

これは成功例。じゃあ，[元のコード](https://github.com/spiegel-im-spiegel/astrocalc/blob/master/modjulian/modjulian.go)を少しいじってわざと失敗させてみましょうか（なんだかなぁ）。

```shell
C:\home\spiegel\docs\docs-baldanders\wiki\Draft\golang-src\astrocalc>go test -v ./...
=== RUN TestModifiedJulianDayNumber
--- FAIL: TestModifiedJulianDayNumber (0.00s)
        modjulian_test.go:35: DayNumber of "1969-12-31 00:00:00 +0000 UTC" = 40585, want 40586.
FAIL
exit status 1
FAIL    _/C_/pathto/astrocalc/modjulian 1.729s
```

エラーレポートを吐く `Errorf()` は内部で `Fail()` を呼び出し，テスト自体は続行します。一方 `Errorf()` の代わりに `Fatalf()` を使うと，内部で `FailNow()` を呼び出しテストを中断します。

Go 言語のテスト・フレームワークでは benchmark や coverage もサポートしてますが，今回は割愛します。

## テストの自動化（Continuous Integration）

今回のコードは自動化するほどの規模でもないですが，話のついでに [Travis CI] で自動化しちゃいましょう。えっと，今回は [Travis CI] の説明は割愛します。ネットにいくらでも解説が転がってますし。

テストを行うだけなら .travis.yml の記述は超簡単です。

```yaml:.travis.yml
language: go

go:
  - 1.4

script:
 - go test -v ./...
```

実行結果は[ここ](https://travis-ci.org/spiegel-im-spiegel/astrocalc)を参照して下さい。

[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"

## ブックマーク

- [Go の Test に対する考え方 - Qiita](http://qiita.com/Jxck_/items/8717a5982547cfa54ebc)
- [Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)
- [golang 1.4で追加されたtestingの便利機能(テストの初期化とお片づけ) - Qiita](http://qiita.com/umisama/items/0d589cca7e89b89c29a8)
- [Go + Travis CI + Coveralls でCI環境を作る - Qiita](http://qiita.com/dmnlk/items/3fb4e0abb98e39fee275)
- [GithubにあるリポジトリをTravis CI連携する手順 #junitbook - くりにっき](http://sue445.hatenablog.com/entry/2013/06/01/170607)
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [golangでTravis CIを使ってクロスコンパイルするときにハマったところ #golang #travisci - uchimanajet7のメモ](http://uchimanajet7.hatenablog.com/entry/2015/03/20/211352)
- [Go言語のビルド生活を drone.ioで幸せに暮らす #golang - Qiita](http://qiita.com/atotto/items/b796c31c1755dbec13db)

## 脚注

