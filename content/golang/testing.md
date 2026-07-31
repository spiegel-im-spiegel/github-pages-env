+++
date = "2015-09-19T23:40:43+09:00"
update = "2018-04-24T20:10:27+09:00"
description = "パッケージ化したのならテストをしましょう。"
draft = false
tags = ["golang", "testing"]
title = "Go 言語のテスト・フレームワーク"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

（初出： [はじめての Go 言語 (on Windows) その7 - Qiita](http://qiita.com/spiegel-im-spiegel/items/64224f22ef17d916dc2d)）

[前回]の続き。

## テストコードを書く

[Go 言語]では最初からテスト・フレームワークが同梱されています。
いまどきの言語はみんなそうですよね。
テストコードを書くには対象のソースファイルと同じフォルダに `*_test.go` という名前のファイルを用意します。
まぁ，説明するより書いた方が早いですね。

```go
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
1. `import` には [`testing`] パッケージを含めます。
1. `Test...` で始まる関数名がテスト実行用の関数です。引数には `t *testing.T` を指定します。
1. `TestMain()` は特別な関数です。テストの最初に呼び出され， `Run()` で他のテスト関数群をキックします。引数には `m *testing.M` を指定します。 `TestMain()` 内で初期化や条件を変えたテストの繰り返しや後始末処理などを行うことができます。

[`testing`] パッケージには，他の言語のテスト・フレームワークによくある [assertion 関数がありません](http://golang.jp/go_faq#assertions)[^1]。 [FAQ](http://golang.jp/go_faq#testing_framework)[^2] によると

> 一般的なテストフレームワークにおいて条件・制御・出力機構を持つ専用のミニ言語が用意される傾向がありますが、Go言語にはすでにこれらが備わっています。これらを再び作成するより、我々はGo言語のテストを進めたかったのです。このようにしたことで余計な言語を覚える必要がなくなり、テストを直接的かつ理解しやすくしています。

とあります。
テスト駆動型開発の場合，テストコードはそれ自体が設計書として機能しますので，この割り切りは妥当と言えます[^3]。
その代わりテストコードを（ドキュメントとして）きちんと書くのは骨が折れますが（笑）

[^1]: オリジナルは [http://golang-jp.org/doc/faq#assertions](http://golang-jp.org/doc/faq#assertions)
[^2]: オリジナルは [http://golang-jp.org/doc/faq#testing_framework](http://golang-jp.org/doc/faq#testing_framework)
[^3]: 私は組み込みエンジニアなので，プログラミングで assert を多用するのは，エンジニアの怠慢だと思ってしまいます。まぁ，ベクタ・テーブルからゴリゴリ書くってのなら別ですが。

テストコードが書けたので早速動かしてみましょう。
環境は[前回]の最後の状態をそのまま引き継ぎます。

テストを行うには `go test` コマンドを使います。
以下の例ではパッケージを指定していますが， `./...` と指定すれば全てのパッケージのテストが対象になります。

```shell
C:\workspace\jd>go test -v github.com/spiegel-im-spiegel/astrocalc/modjulian
=== RUN   TestDayNumber
--- PASS: TestDayNumber (0.00s)
PASS
ok      github.com/spiegel-im-spiegel/astrocalc/modjulian       0.229s
```

これは成功例。じゃあ，[元のコード](https://github.com/spiegel-im-spiegel/astrocalc/blob/master/modjulian/modjulian.go)を少しいじってわざと失敗させてみましょうか（なんだかなぁ）。

```shell
C:\workspace\jd>go test -v github.com/spiegel-im-spiegel/astrocalc/modjulian
=== RUN   TestDayNumber
--- FAIL: TestDayNumber (0.00s)
        modjulian_test.go:35: DayNumber of "1969-12-31 00:00:00 +0000 UTC" = 40587, want 40586.
FAIL
exit status 1
FAIL    github.com/spiegel-im-spiegel/astrocalc/modjulian       1.566s
```

エラーレポートを吐く `Errorf()` は内部で `Fail()` を呼び出し，テスト自体は続行します。
一方 `Errorf()` の代わりに `Fatalf()` を使うと，内部で `FailNow()` を呼び出しテストを中断します。

Go 言語のテスト・フレームワークでは benchmark や coverage もサポートしてますが，今回は割愛します。

## テストの自動化（Continuous Integration）

今回のコードは自動化するほどの規模でもないですが，話のついでに [Travis CI] で自動化しちゃいましょう。
えっと，今回は [Travis CI] の説明は割愛します。
興味のある方は「[ブックマーク]({{< ref "#bookmark" >}})」の項を参考にして下さい。

[Travis CI] でビルド・テストを行うためには `.travis.yml` を書く必要がありますが，テストを行うだけなら `.travis.yml` の記述は簡単です。

```yaml
language: go

go:
  - 1.4
  - 1.5

script:
 - go test -v ./...
```

実行結果は[ここ](https://travis-ci.org/spiegel-im-spiegel/astrocalc)を参照して下さい。

[次回]はドキュメントの話。

## ブックマーク{#bookmark}

- [Go の Test に対する考え方 - Qiita](http://qiita.com/Jxck_/items/8717a5982547cfa54ebc)
- [Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)
- [golang 1.4で追加されたtestingの便利機能(テストの初期化とお片づけ) - Qiita](http://qiita.com/umisama/items/0d589cca7e89b89c29a8)
- [Go + Travis CI + Coveralls でCI環境を作る - Qiita](http://qiita.com/dmnlk/items/3fb4e0abb98e39fee275)
- [GithubにあるリポジトリをTravis CI連携する手順 #junitbook - くりにっき](http://sue445.hatenablog.com/entry/2013/06/01/170607)
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [golangでTravis CIを使ってクロスコンパイルするときにハマったところ #golang #travisci - uchimanajet7のメモ](http://uchimanajet7.hatenablog.com/entry/2015/03/20/211352)
- [Go言語のビルド生活を drone.ioで幸せに暮らす #golang - Qiita](http://qiita.com/atotto/items/b796c31c1755dbec13db)
- [Golang におけるサブテストの並行処理実装について | eureka tech blog](https://developers.eure.jp/tech/go1_7-subtests/)
- [golangのテストはじめ - Qiita](https://qiita.com/tmzkysk/items/8bb37795ac223664d682)

[Go 言語に関するブックマーク集はこちら]({{< relref "bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< relref "packaging.md" >}} "機能のパッケージ化"
[次回]: {{< relref "document.md" >}} "Go 言語のドキュメント・フレームワーク"
[`testing`]: http://golang.org/pkg/testing/
[Travis CI]: https://travis-ci.org/ "Travis CI - Test and Deploy Your Code with Confidence"
