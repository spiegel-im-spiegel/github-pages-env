+++
title = "Go 言語に関するブックマーク"
date = "2015-09-11T17:58:42+09:00"
description = "本業が忙しくて Go 言語をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。"
tags = ["golang", "bookmark"]
image = "/images/attention/go-logo_blue.png"

[scripts]
  mathjax = false
  mermaidjs = false
+++

本業が忙しくて [Go 言語]をかまってあげる暇がないのだが，ブックマークばっかり溜まっていくので，定期的に吐き出しておく。
なお，ここにない場合は「[未整理分]({{< relref "bookmark2.md" >}})」にあるかも。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 公式サイト

- [The Go Programming Language](https://golang.org/)
    - [git repositories (Google)](https://go.googlesource.com/)
    - [git repositories (GitHub)](https://github.com/golang) : mirror
- [go.dev](https://go.dev/)
    - [Go.dev: a new hub for Go developers - The Go Blog](https://blog.golang.org/go.dev)
- [A Tour of Go](https://go-tour-jp.appspot.com/) : 日本語版

## リリース情報

- [Go Turns 10 - The Go Blog](https://blog.golang.org/10years)

### Go 1.5 Released. 

- [Go 1.5 is released - The Go Blog](https://blog.golang.org/go1.5)
- [Go 1.5 Release Notes - The Go Programming Language](https://golang.org/doc/go1.5)
    - [Go 1.4 "Internal" Packages](https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit) : Internal Packages は 1.5 で GOPATH まで拡張された
    - [Go 1.5 Vendor Experiment](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo/edit)

- [GVM で go1.5rc1 のインストール - Qiita](http://qiita.com/msaito3/items/3aef86e9864990b16b4c)
- [goを1.5にアップデートして1.4とベンチを取る - Qiita](http://qiita.com/masahikoofjoyto/items/4ced298989e6ab346f15)
- [Go 1.3 から 1.5 へのアップデートでエラー - Qiita](http://qiita.com/taji-taji/items/4c43e126e67d65a219e3) : 古いバージョンからアップデートする際は，いったん 1.4 に上げてから 1.5 にアップデートするとよい
- [Big Sky :: golang 1.5 の internal パッケージの使い方。](http://mattn.kaoriya.net/software/lang/go/20150820102400.htm)
    - [「golang 1.5 の internal パッケージの使い方。」を試してみた - Qiita](http://qiita.com/qt-luigi/items/d0f52b3b0906b35e6027)

### Go 1.6 is released

- [Go 1.6 is released - The Go Blog](https://blog.golang.org/go1.6)
- [Go 1.6 Release Notes - The Go Programming Language](https://golang.org/doc/go1.6)
    - [Go1.6の新機能 - Qiita](http://qiita.com/ksato9700/items/5505e506c20b6048c218)

- [Goのバージョンを1.6rc2にアップデートしてみた - Qiita](http://qiita.com/kanuma1984/items/245f7efafeaee5728523)
- [Goで良い感じに日時をパースするライブラリdatemakiの話とGo 1.6 - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/12/22/000011)
- [Go1.6でポインタをcgoの関数へ渡す際の注意点 - Qiita](http://qiita.com/saturday06/items/84535c61a3328c02032c)
- [Go1.6でポインタをcgoの関数へ渡す際に発生するcgoCheckPointerを回避する方法 - Qiita](http://qiita.com/mattn/items/90c8558d5fff05a2ba0c)
- [Goのバージョンを1.4.3→1.6にアップグレードできなかった - Qiita](http://qiita.com/hirocueki2/items/3ec4b409a3ed2cbea681)
- [Go 1.6 templateパッケージ新機能 - Qiita](http://qiita.com/hoshi-k/items/f2eaff298f93f089e10d)

### Go 1.7 is released

- [Go 1.7 is released - The Go Blog](https://blog.golang.org/go1.7)
- [Go 1.7 Release Notes - The Go Programming Language](https://golang.org/doc/go1.7)

- [go1.7 の気になるところを試した - Qiita](http://qiita.com/tutuming/items/16df6fc7bf82fb5d7eb0)
- [Sierraでgo1.7.3のコンパイル - Qiita](http://qiita.com/lufia/items/8a71a29fa6e1089735f2)
- [【作業メモ】さくらインターネットの標準CGIサーバでGoをビルド【古典的レンタル鯖】 - Qiita](http://qiita.com/zetamatta/items/143849cf34db4ffaad4b)
- [Go1.7からSSAが導入された - flyhigh](http://shinpei.github.io/blog/2016/08/13/what-ssa-brings-to-go-17/)

### Go 1.8 is released

- [Go 1.8 is released - The Go Blog](https://blog.golang.org/go1.8)
- [Go 1.8 Release Notes - The Go Programming Language](https://golang.org/doc/go1.8)

- [Go 1.8のpluginパッケージを試してみる - Qiita](http://qiita.com/qt-luigi/items/47a7913145fc747da0c7)

### Go 1.9 is released

- [Go 1.9 is released - The Go Blog](https://blog.golang.org/go1.9)
- [Go 1.9 Release Notes - The Go Programming Language](https://golang.org/doc/go1.9)

- [go言語1.9で追加予定の新機能 型エイリアス - Qiita](http://qiita.com/weloan/items/8abbb4003cfa1031a9e9)

### Go 1.10 is released

- [Go 1.10 is released - The Go Blog](https://blog.golang.org/go1.10)
- [Go 1.10 Release Notes - The Go Programming Language](https://golang.org/doc/go1.10)

- [go1.10beta1の標準パッケの大きな変更点を確認しておく。 - Qiita](https://qiita.com/A_Resas/items/59bf6cda976e29751890)
- [Go1.10で入るstrings.Builderを検証した #golang - Qiita](https://qiita.com/tenntenn/items/94923a0c527d499db5b9)
- [go1.10のtest2jsonを試してみた - Qiita](https://qiita.com/dproject21/items/c406b0044280508b41ff)

### Go 1.11 is released

- [Go 1.11 is released - The Go Blog](https://blog.golang.org/go1.11)
- [Go 1.11 Release Notes - The Go Programming Language](https://golang.org/doc/go1.11)

- [Go 1.11 リリースノート（和訳）](https://qiita.com/pokeh/items/c6511ca15c9a33b48fcc)
- [サクッと Go → WebAssembly を試す](https://qiita.com/cia_rana/items/bbb4112b480636ab9d87)

### Go 1.12 is released

- [Go Modules in 2019 - The Go Blog](https://blog.golang.org/modules2019)
- [Go 1.12 is released - The Go Blog](https://blog.golang.org/go1.12)
- [Go 1.12 Release Notes - The Go Programming Language](https://golang.org/doc/go1.12)
- [Debugging what you deploy in Go 1.12 - The Go Blog](https://blog.golang.org/debugging-what-you-deploy)

### Go 1.13 is released

- [Go 1.13 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.13)
    - [ドラフトリリースノート - Go 1.13 の紹介 - Qiita](https://qiita.com/c-yan/items/b2f5e5c168d517594eb2)
- [Next steps toward Go 2 - The Go Blog](https://blog.golang.org/go2-next-steps)
- [スライドまとめ on Go 1.13 Release Party in Tokyo（非公式） - Qiita](https://qiita.com/usk81/items/b2803b47e658a905af98)
- [Go1.13で導入されたsyscall/jsのjs.CopyBytesToGoとjs.CopyBytesToJSを試してみた - Qiita](https://qiita.com/neko-suki/items/a39dc56523b3d761621b)

### Go 1.14 is released

- [Go 1.14 is released - The Go Blog](https://blog.golang.org/go1.14)
- [Go 1.14 Release Notes - The Go Programming Language](https://golang.org/doc/go1.14)
    - [Go 1.14 リリースノート 日本語訳 - Qiita](https://qiita.com/c-yan/items/3dd0c63ea4c3041728cc)
- [Changes to interfaces in Go1.14 - Dylan Meeus - Medium](https://medium.com/@meeusdylan/changes-to-interfaces-in-go1-14-821ab533f62)
- [A Go Module Testbed « null program](https://nullprogram.com/blog/2020/02/13/)
- [Inlined defers in Go · Go, the unwritten parts](https://rakyll.org/inlined-defers/)
- [Go1.14で来るGo Modules関連の変更を見てみる - Qiita](https://qiita.com/pi9min/items/9dac44c6663e0e933968)
- [Go1.14のcontextは何が変わるのか - Qiita](https://qiita.com/tutuz/items/963a6118cec63a4cd2f3)
- [What's new Go 1.14?](https://talks.godoc.org/github.com/qt-luigi/talks/2020/whats-new-go-114.slide)

### Go 1.15 Beta

- [Proposals for Go 1.15 - The Go Blog](https://blog.golang.org/go1.15-proposals)
- [Go 1.15 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.15) : for beta version

### Go 2 Draft

次期 Go 言語の仕様について（一部は 1.11 以降に取り込まれている）

- [Go 2 Draft Designs - The Go Blog](https://blog.golang.org/go2draft)
- [Go 2 Draft Designs](https://go.googlesource.com/proposal/+/master/design/go2draft.md)
    - [Error Inspection — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-error-inspection.md)
    - [Generics — Problem Overview](https://go.googlesource.com/proposal/+/master/design/go2draft-generics-overview.md)
    - [Contracts — Draft Design](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md)
- [Go 2のgenerics/contract簡易まとめ](https://qiita.com/lufia/items/242d25e8c93d88e22a2e)
- [Go 2, here we come! - The Go Blog](https://blog.golang.org/go2-here-we-come)
- [The Next Step for Generics - The Go Blog](https://blog.golang.org/generics-next-step)

### Go Cloud Development Kit

- [What's new in the Go Cloud Development Kit - The Go Blog](https://blog.golang.org/gcdk-whats-new-in-march-2019)
- [The New Go Developer Network - The Go Blog](https://blog.golang.org/go-developer-network)
- [Pkg.go.dev is open source! - The Go Blog](https://blog.golang.org/pkgsite)

## 言語仕様および標準パッケージに関すること

### はじめての [Go 言語]

- [はじめてのGo―シンプルな言語仕様，型システム，並行処理：特集｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/feature/01/go_4beginners)
- [Golangの基本文法をおさえてみる - Qiita](http://qiita.com/kazusa-qooq/items/40f9ea3e72406d845b10)
- [忙しい人のためのA Tour of Go - Qiita](http://qiita.com/makoto_kw/items/0638c0af1002647e3f7a)
- [import 書き方まとめ - Qiita](http://qiita.com/taji-taji/items/5a4f17bcf5b819954cc1)
- [GoのEnumイディオム - Qiita](http://qiita.com/awakia/items/c81c7816b9aea5422250) : あらかじめ Enum 用の type を作成し、その type に対する `String()` メソッドを定義する
- [Big Sky :: golang では for ループの中で defer してはいけない。](http://mattn.kaoriya.net/software/lang/go/20151212021608.htm) : ループ内で defer が必要になるということは refactoring のチャンス
- [Goのfor rangeで思った値が取れなかった話 - Qiita](http://qiita.com/modal_soul/items/e49480e5692597fda975) : ちょっとしたミス
- [Goで再帰使うと遅くなりますがそれが何だ - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2015/02/23/165341)
- [Go言語(Golang) はまりどころと解決策](http://www.yunabe.jp/docs/golang_pitfall.html)
- [Big Sky :: Golang のオフィシャルが提供するインタフェースまとめ](http://mattn.kaoriya.net/software/lang/go/20140501172821.htm)
- [Go の定数の話 - Qiita](http://qiita.com/hkurokawa/items/a4d402d3182dff387674)
- [init関数のふしぎ #golang - Qiita](http://qiita.com/tenntenn/items/7c70e3451ac783999b4f)
    - [packageに複数のinitがあるときの挙動 - Qiita](http://qiita.com/astronoka/items/aa2f271d280863cedf5e)
- [Big Sky :: golang では変数の宣言位置が大事](http://mattn.kaoriya.net/software/lang/go/20170406003909.htm)
- [Big Sky :: Names](http://mattn.kaoriya.net/software/20160126101358.htm) : Golang の開発者 Russ Cox 氏による記事の抄訳。「変数名の長さ」について
- [Go言語のキーワードが少ない理由 - Qiita](https://qiita.com/weloan/items/ce6e6dce36a6f774d841)
- [Go の、型のない定数の精度 - Qiita](https://qiita.com/Nabetani/items/21ea926bb532cb0ac094)
- [Goを読むその1：go installの流れ](https://qiita.com/junjis0203/items/de55bb69945f08a9166c)
    - [Goを読むその2：compileコマンド（構文解析まで）](https://qiita.com/junjis0203/items/616c00086eb336153f4f)
- [[Go]imported and not usedエラー・declared and not usedエラーとの向き合いかた - My External Storage](https://budougumi0617.github.io/2019/10/06/imported-declared-not-used-error/)
- [C から Go へコードを移植してハマった話 (そして言語仕様へ) - Qiita](https://qiita.com/akif999/items/a94dd8a6fe29dda2b560)
- [Go's Declaration Syntax - The Go Blog](https://blog.golang.org/gos-declaration-syntax)
- [[Go] タグなし switchは switch true {...}と等しい - My External Storage](https://budougumi0617.github.io/2019/11/10/switch-statement-in-go/)
- [Go の命名規則 | micnncim](https://micnncim.com/post/2019/12/11/go-naming-conventions/)
- [Go: How Are Loops Translated to Assembly? - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-how-are-loops-translated-to-assembly-835b985309b3)
- [Go: Inlining Strategy & Limitation - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-inlining-strategy-limitation-6b6d7fc3b1be)
- [Why you shouldn't use func main in Go by Mat Ryer - PACE.](https://pace.dev/blog/2020/02/12/why-you-shouldnt-use-func-main-in-golang-by-mat-ryer)
- [proxy.golang.org の挙動調査メモ — KaoriYa](https://www.kaoriya.net/blog/2020/06/16/)

### オブジェクトに関すること

- [オブジェクト指向言語としてGolangをやろうとするとハマること - Qiita](http://qiita.com/shibukawa/items/16acb36e94cfe3b02aa1)
    - [オブジェクト指向言語としてGolangをやろうとするとハマる点を整理してみる - Qiita](http://qiita.com/sona-tar/items/2b4b70694fd680f6297c)
- [Go言語に継承は無いんですか【golang】 - DRYな備忘録](http://otiai10.hatenablog.com/entry/2014/01/15/220136)
- [Goにatexitやグローバルなデストラクタがない理由 - Qiita](http://qiita.com/ruiu/items/22910a4bae6cb716a391)

### メモリ管理

- [Go: Memory Management and Allocation - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-memory-management-and-allocation-a7396d430f44)
- [Go: How Does the Garbage Collector Mark the Memory?](https://medium.com/a-journey-with-go/go-how-does-the-garbage-collector-mark-the-memory-72cfc12c6976)
- [Go: Memory Management and Memory Sweep - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-memory-management-and-memory-sweep-cc71b484de05)

### 型と [interface]

[type]: https://golang.org/ref/spec#Properties_of_types_and_values "Properties of types and values"
[interface]: https://golang.org/doc/effective_go.html#interfaces_and_types "Effective Go - The Go Programming Language"

- [Go言語の型とreflect - Qiita](http://qiita.com/atsaki/items/3554f5a0609c59a3e10d)
- [埋込みとインタフェース #golang - Qiita](http://qiita.com/tenntenn/items/bba69d84a1e0b67c4db8)
- [Golang: nil Pointer Receiverの話 - Qiita](http://qiita.com/stsn/items/73714caf8458b1d973f2)
- [Go 言語の値レシーバとポインタレシーバ | Step by Step](https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/)
- [Big Sky :: Go言語でインタフェースの変更がそれ程問題にならない理由](http://mattn.kaoriya.net/software/lang/go/20130919023425.htm)
- [Go で interface {} の中身がポインタならその参照先を取得する - Qiita](http://qiita.com/chimatter/items/b0879401d6666589ab71)
- [Go で型を抽象的に扱うには - Qiita](https://qiita.com/nirasan/items/d1b80ecc7a7a7f0af2b7)
- [GoのInterfaceとは何者なのか #golang #go - niconegoto Blog](http://niconegoto.hatenadiary.jp/entry/2017/12/03/222922)
- [インタフェースの実装パターン #golang - Qiita](https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8)
- [Big Sky :: Go で型がインタフェースを実装している事を保証するには](https://mattn.kaoriya.net/software/lang/go/20190702104455.htm)
- [Goの基本的な型変換 - Qiita](https://qiita.com/lostfind/items/ad7bfc1a4860bb108b9c)
- [Sliceを含んだ構造体が等値演算子（==）でpanicを引き起こすのを回避する #golang - My External Storage](https://budougumi0617.github.io/2019/07/07/prevent-runtime-error-by-pointer/)
- [Go言語のInterfaceの考え方、Accept interfaces,return structs - Qiita](https://qiita.com/weloan/items/de3b1bcabd329ec61709)
- [Named typeとType aliasを使い分ける - My External Storage](https://budougumi0617.github.io/2020/02/01/go-named-type-and-type-alias/)

#### 数値型

- [Golang の 数値型 - Qiita](http://qiita.com/tanaka0325/items/9c61a022cd32be0c65a6)
- [Go言語での初期化における&amp;とnewの挙動の違い - Qiita](https://qiita.com/wannabe/items/87200a2cfc62cd7859bb)
- [Using Go Interfaces for Testable Code - The Startup - Medium](https://medium.com/swlh/using-go-interfaces-for-testable-code-d2e11b02dea)

#### 文字列型と操作・変換

[string]: http://golang.org/ref/spec#String_types

- [Strings, bytes, runes and characters in Go - The Go Blog](http://blog.golang.org/strings)
- [Go言語のstring, runeの正体とは？ - golang - The Round](http://knightso.hateblo.jp/entry/2014/06/24/090719)
- [Go小ネタ: 正規表現を全角半角問わずマッチするよう変換する - Qiita](http://qiita.com/tutuming/items/d8aaaf5442d84a7961e1)
- [Go言語は空文字に対してstrings.splitを掛けると1要素の配列を返す - Qiita](http://qiita.com/Sheile/items/ba51ac9091e09927b95c) : コメントに別解あり
- [golang - Goでマルチバイトが混在した文字列をtruncateする - Qiita](http://qiita.com/hokaccha/items/3d3f45b5927b4584dbac)
- [Golangでの文字列・数値変換 - 小野マトペの納豆ペペロンチーノ日記](http://matope.hatenablog.com/entry/2014/04/22/101127)
- [golang - Go言語で文字コード変換 - Qiita](http://qiita.com/uchiko/items/1810ddacd23fd4d3c934)
- [Go で euc-jp や sjis の csv ファイルを読み込むには変換用のリーダーを1つかませるだけでよかった - Qiita](http://qiita.com/ikawaha/items/540c2af34b1c381c37c1)
- [Goでは文字列連結はコストの高い操作 - Qiita](http://qiita.com/ruiu/items/2bb83b29baeae2433a79)
- [Goの文字列結合のパフォーマンス - Qiita](http://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37)
- [Go言語で SplitMultiSep (複数種の区切り文字列で分解) - Qiita](http://qiita.com/yoya/items/23ac2c490625c5d47ad9)
- [Go で UTF-8 の文字列を扱う - Qiita](http://qiita.com/masakielastic/items/01a4fb691c572dd71a19)
- [Go言語 Gmailのsubjectの日本語文字化けに対応する - Qiita](http://qiita.com/yyoshiki41/items/79882e269ca6af4c2236)
- [Text normalization in Go - The Go Blog](https://blog.golang.org/normalization) : Unicode 正規化について
- [Go言語で文字列の変換(全角・半角、ひらがな・カタカナ)をする : Serendip - Webデザイン・プログラミング](http://www.serendip.ws/archives/6307)
- [Goで全角英数字を半角にする - Qiita](http://qiita.com/ktashiro/items/da5cbee3129acc74e5d7)
- [Goのstring型が思ったより容量食いだった話 - Qiita](https://qiita.com/hnw/items/ec3da327c37e3ad8c875)
- [golang で string を []byte にキャストしてもメモリコピーが走らない方法を考えてみる - Qiita](https://qiita.com/mattn/items/176459728ff4f854b165) : `unsafe` パッケージで無理やり処理（笑）
- [utf8としてvalidなバイト列を判定する方法をGoから見る - Qiita](https://qiita.com/catatsuy/items/bccc2c76be501e98382a)
    - [Go言語で文字列がASCIIコード内であるか判定したい - Qiita](https://qiita.com/catatsuy/items/7a9773f9ea3db7069fc1)
- [String interpolation in Golang – Ly Channa – Medium](https://medium.com/@channaly/string-interpolation-in-golang-ecb3bcb2d600)
    - [Golang の文字列内で変数を展開する方法（各種） - Qiita](https://qiita.com/KEINOS/items/baef1be88f15515026ec)
- [Big Sky :: Go で大文字小文字無視の文字列比較ベンチマーク](https://mattn.kaoriya.net/software/lang/go/20190806152526.htm)
- [Go の strings.Index の内部実装と Rabin-Karp アルゴリズム - Qiita](https://qiita.com/po3rin/items/07d51249629390a6201a)
- [Go: String & Conversion Optimization - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-string-conversion-optimization-767b019b75ef)

### 配列と [slice]

[slice]: http://golang.org/ref/spec#Slice_types

- [Go のスライスでハマッたところ - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/golang-slice-internals2)
- [golang - go言語のslice操作をまとめてみた（shiftしたりpushしたり） - Qiita](http://qiita.com/egnr-in-6matroom/items/282aa2fd117aab9469bd)
- [sliceの重複チェックを高速化 - Qiita](http://qiita.com/hi-nakamura/items/5671eae147ffa68c4466)
- [Goのarrayとsliceを理解するときがきた - Qiita](http://qiita.com/seihmd/items/d9bc98a4f4f606ecaef7) : この説明は分かりやすい。おススメ
- [uint64型を[]bytes型に変換する - Qiita](http://qiita.com/joniyjoniy/items/cbfb7d5c49aec5bf63c0)
- [golangのequalityの評価について - podhmo's diary](http://pod.hatenablog.com/entry/2016/07/30/204357)
    - [Goで違うmapであることをテストする - Qiita](https://qiita.com/karupanerura/items/03d6766fd8568c15fc90)
- [文字列をn個後ろにずらす処理を本気でやる - Qiita](https://qiita.com/smith_30/items/eec0ba2e4ec63fe879a0)
- [Golangのmapとsliceはどちらが速いのか - 逆さまにした](https://cipepser.hatenablog.com/entry/go-map-slice)
- [GitHub - imdario/mergo: Mergo: merging Go structs and maps since 2013.](https://github.com/imdario/mergo)
- [Go理解度チェック - Google スライド](https://docs.google.com/presentation/d/1oqfPIEJlw1u0GStHPEJFMtq2FPFUGDVSgMW2sLsXZF8/mobilepresent#slide=id.g7a74c7c8ae_0_298)
- [[Go] 省略記号（...）を使った配列宣言の仕方 - My External Storage](https://budougumi0617.github.io/2020/03/06/go-array-with-the-ellipsis/)
- [Go: Slice and Memory Management - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-slice-and-memory-management-670498bb52be)

### 入出力処理

- [Go ファイルや標準入力から一行ずつ読み込む - Qiita](http://qiita.com/hnakamur/items/a53b701c8827fe4bfec7)
- [Goでファイルを読んで別のgoroutineに渡す - Qiita](http://qiita.com/ono_matope/items/5c8cfce81933c5eb9fd0)
- [「連結されたgzipを1行ずつ見る」をgolangでやったらハマった - Qiita](http://qiita.com/kroton/items/431e6dad9e5e4dbc44cf) : [compress/gzip](https://golang.org/pkg/compress/gzip/) と入出力処理の関係
- [bufio.Scannerのend-of-line判断を変更してみる - Qiita](http://qiita.com/curious-eyes/items/2d4b6c20ea47e3efc47b)
- [KOBE GDG: Go言語　バイナリファイルを扱う](http://kobegdg.blogspot.jp/2013/05/go.html) : 任意のオブジェクトをバイト列に変換してファイルに格納
- [ファイル書き込みの度にファイルを開いたらどれくらい遅いのか - Qiita](http://qiita.com/catatsuy/items/41d3c49248b517b5af96)
- [Golangで標準入力がパイプで渡されたものか判定する - Qiita](http://qiita.com/tanksuzuki/items/e712717675faf4efb07a)
- [Go言語: ファイルの存在をちゃんとチェックする実装? - Qiita](http://qiita.com/suin/items/b9c0f92851454dc6d461)
- [Golangで、ファイル一覧取得（最新順出力） - Qiita](http://qiita.com/shinofara/items/e5e78e6864a60dc851a6)
- [大きなファイルのアップロードを省メモリで行いたい(io.Pipeを使う) - Qiita](http://qiita.com/m0a/items/bba395b2fc9cd160e441)
- [Big Sky :: net/http でレスポンスの内容を確認したいなら io.TeeReader を使おう](https://mattn.kaoriya.net/software/lang/go/20171026101727.htm)
    - [`io.TeeReader`](https://golang.org/pkg/io/#TeeReader)
- [Big Sky :: golang で UNIX コマンドパイプラインを扱う](https://mattn.kaoriya.net/software/lang/go/20151030131242.htm)


### 並行処理と並列処理

- [Visualizing Concurrency in Go · divan's blog](http://divan.github.io/posts/go_concurrency_visualize/)
- [Go の並行処理 - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20130414/1365960707)
- [Go: 計算なしのFizzBuzz - Qiita](http://qiita.com/suin/items/eca21ed935115e5da2e8) : channel の説明するのにいいかも
- [Goのchannelの送受信用の型について - Qiita](http://qiita.com/yuki2006/items/3f90e53ce74c6cff1608)
- [Go言語のChannelは送信時にもブロックする - Qiita](http://qiita.com/hondata/items/64776c79063e93bea9ed) : 意外と見落とす channel 送信時のブロック
- [select loop の小ネタ - Qiita](http://qiita.com/Jxck_/items/da3ca2db58734a966cac)
- [Goのforとgoroutineでやりがちなミスとたった一つの冴えたgo vetと - Qiita](http://qiita.com/sudix/items/67d4cad08fe88dcb9a6d)
- [x/net/context の実装パターン - Qiita](http://qiita.com/tutuming/items/c0ffdd28001ee0e9320d) : [golang.org/x/net/context](https://godoc.org/golang.org/x/net/context) を使って並行処理を細かく制御。 Domain-Driven あるいは Context-Driven な設計でも使えそう。
- [Go言語でチャネルとselect - Qiita](http://qiita.com/najeira/items/71a0bcd079c9066347b4)
- [golangでシグナルを拾ってgracefulにgoroutineを停めたい - Qiita](http://qiita.com/arc279/items/c44d4a18a851ff454c64)
- [golang の channel のブロックがよくわからん - Qiita](http://qiita.com/arc279/items/bc55cdf436c544e91c05)
- [GoのChannelを使いこなせるようになるための手引 - Qiita](http://qiita.com/awakia/items/f8afa070c96d1c9a04c9)
- [Goでスレッド（goroutine）セーフなプログラムを書くために必ず注意しなければいけない点 - Qiita](http://qiita.com/ruiu/items/54f0dbdec0d48082a5b1) : `sync.Mutex` にも言及
- [複数のgroutineが生えてるStructの安全な終了方法 - Qiita](http://qiita.com/shunsukeaihara/items/f9ef7c8d430f63d79d40)
- [Go言語の並行性を映像化する | プログラミング | POSTD](http://postd.cc/go_concurrency_visualize/)
- [Go言語でプロセス間同期処理 - Qiita](http://qiita.com/shanxia1218/items/7fb15f50ec645f114bc7) : Windows の Mutex を使ってプロセス間通信を行う
- [Goで並行処理のベンチマークをとる - Qiita](http://qiita.com/hhatto/items/c8eb987b0516a45db754)
- [go言語初心者が図を書きながらgo routineやgo channelを理解する(Part 1) - Qiita](http://qiita.com/Vermee81/items/88c9e28dec83d43e7883)
    - [go言語初心者が図を書きながらgoroutineやgo channelを理解する(Part2) - Qiita](http://qiita.com/Vermee81/items/30ad42a7265375b1b7b1)
- [Goroutineと channelから はじめるgo言語](http://www.slideshare.net/takuyaueda967/goroutine-channel-go)
- [Go言語の並行処理デザインパターン by Rob Pike 前編 - Qiita](http://qiita.com/tfutada/items/a289628d8b2d0af6152d)
    - [Go言語の並行処理デザインパターン by Rob Pike 後編 - Qiita](http://qiita.com/tfutada/items/dc8db894ac270a79ef2b)
- [goroutine を安全に止める方法 - Qiita](http://qiita.com/castaneai/items/7815f3563b256ae9b18d)
- [Goの同時関数呼び出しを１回で済ませられるライブラリ 「SingleFlight」 が便利 - Qiita](https://qiita.com/igtm/items/8b5343272bc35cd3bc0b)
- [【コンピュータ将棋】ゴルーチンでお手軽持ち時間管理＆並行探索 - Qiita](http://qiita.com/32hiko/items/3be36dad2d651399ba1b)
- [Goのワークスティーリング型スケジューラ | プログラミング | POSTD](http://postd.cc/scheduler/)
- [goの並行処理パターンについてのリンク集 - Qiita](https://qiita.com/ryskiwt/items/b721f6aecc5b3d680462)
    - [chanの使い方パターンメモ。 - GolangRdyJp](http://golang.rdy.jp/2015/03/25/chan_tips/)
    - [Big Sky :: golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)
- [go1.9のsyncmapを試してみた - Qiita](https://qiita.com/arihitohagiwara/items/4bb2ae6a1a43384b4f60) : 並行処理に使える同期型の map
- [Handling CTRL-C (interrupt signal) in Golang Programs | I care, I share, I'm Nathan LeClaire.](https://nathanleclaire.com/blog/2014/08/24/handling-ctrl-c-interrupt-signal-in-golang-programs/)
- [Go context.Context interfaceに Cancelが含まれていない理由 - Qiita](https://qiita.com/YmgchiYt/items/abc6c0a8f57b47fdfcae)
- [Goroutineハンターが過労死する前に - Qiita](https://qiita.com/i_yudai/items/3336a503079ac5749c35)
- [or-done-channelでコードの可読性を上げる - YAMAGUCHI::weblog](http://ymotongpoo.hatenablog.com/entry/2017/12/04/091403)
- [go1.9から追加されたsync.Mapを使う - Qiita](https://qiita.com/meta_closure/items/dd228e49aef8b67e872e)
- [Go1.9から追加されたsync.Mapのパフォーマンス – Straightforward](https://tanksuzuki.com/entries/golang-syncmap/)
- [Big Sky :: 簡単に goroutine の実行個数を制限する方法](https://mattn.kaoriya.net/software/lang/go/20171221111857.htm)
- [Big Sky :: goroutine でドハマリした。](https://mattn.kaoriya.net/software/lang/go/20180124171404.htm)
- [Big Sky :: Go 言語の非同期パターン](https://mattn.kaoriya.net/software/lang/go/20180531104907.htm)
- [goroutineとチャネルの動きを図を使って理解する(和訳) - Qiita](https://qiita.com/hirano00o/items/828393342efcd80aa2e6)
- [goleakでgoroutine leakのリスクを減らす - Qiita](https://qiita.com/smith-30/items/67c4f894bec45a6fd82a)
- [Go の -race option は内部で何をしているのか。何を検知しないのか。 - Qiita](https://qiita.com/po3rin/items/18a6621f39e9a6f6f7c4)
- [Achieving concurrency in Go - RunGo - Medium](https://medium.com/rungo/achieving-concurrency-in-go-3f84cbf870ca)
- [Go: Goroutine, OS Thread and CPU Management - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-goroutine-os-thread-and-cpu-management-2f5a5eaf518a)
- [Go: GOMAXPROCS & Live Updates - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-gomaxprocs-live-updates-407ad08624e1)
- [Go: Goroutine and Preemption - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-goroutine-and-preemption-d6bc2aa2f4b7)
- [Go: Concurrency & Scheduler Affinity - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-concurrency-scheduler-affinity-3b678f490488)
- [Go: How Does Go Recycle Goroutines? - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-how-does-go-recycle-goroutines-f047a79ab352)
- [Go: How Does a Goroutine Start and Exit? - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-how-does-a-goroutine-start-and-exit-2b3303890452)

Go 言語で複数 CPU を使った並列処理を行うには明示的な設定が必要。

- [Go言語でCPU数に応じて並列処理数を制限する | SOTA](http://deeeet.com/writing/2014/07/30/golang-parallel-by-cpu/)
    - [やはり俺のgolangがCPUを一つしか使わないのはまちがっている。 - Qiita](http://qiita.com/ymko/items/554e3630fefdc29393a8)
- [Goでお手軽に行列の積を爆速並列計算 - Qiita](http://qiita.com/hama_du/items/fce4ee1e4b5c2c2d24df)

### Logging

- [Golang logging library - Qiita](http://qiita.com/kosuda/items/988c505c2abc5321aba8)
- [go言語におけるロギングについて](http://blog.satotaichi.info/logging-frameworks-for-go/)
- [Golangで簡単にログを吐くことを考える - Qiita](http://qiita.com/Ets/items/49e8f781990a3b0b3821) : [seelog](https://github.com/cihub/seelog) について解説している。私は XML には全くアレルギーがないので無問題
- [Go言語でdebugログの実現方法 - Qiita](http://qiita.com/sbjib/items/2cef51e572eef0795bc2)
- [loggingについて話そう - Qiita](http://qiita.com/methane/items/cedbf546ae2db8a63c3d)
- [golangでlogを標準出力とテキストファイルの2箇所の出力する - Qiita](http://qiita.com/74th/items/441ffcab80a6a28f7ee3)
- [gorpで実行されるSQLをログに出力する - Qiita](http://qiita.com/reiki4040/items/303a0bfa9f0296eb544f)
- [Goのバッチで統計を取得するAPIを用意しておくと便利 - Qiita](http://qiita.com/sudix/items/c542e1b59bc94dc741e3)
- [golangのloggerを作ってみた - Qiita](http://qiita.com/kazuma1107/items/009454fca4f56af6e411)
- [golangの高速な構造化ログライブラリ「zap」の使い方 - Qiita](http://qiita.com/emonuh/items/28dbee9bf2fe51d28153)
- [zapでログレベルでログの出力先を振り分ける方法 - Qiita](http://qiita.com/emonuh/items/cb3a730979dba7d76920)
- [go-logging における "module" の意味 - Qiita](https://qiita.com/shinsa82/items/c05a0e8544ecdb87c442)
- [Golangでlogのタイムスタンプをマイクロ秒単位にする方法 - Qiita](https://qiita.com/tetsu_koba/items/8401f99a39c9757fcff1)
- [ええっ！？　文字列で書くの！？　ログレベル付きロガーhashicorp/logutilsのご紹介 - Qiita](https://qiita.com/mackee_w/items/3c0940733b6c0922554c)
- [rs/zerolog: Zero Allocation JSON Logger](https://github.com/rs/zerolog) : JSON 形式でログを吐く。おススメ
- [hnakamur/ltsvlog: a minimalist LTSV logging library in Go](https://github.com/hnakamur/ltsvlog) : LTSV 形式でログを吐く
    - [GoでLTSV形式でログ出力するライブラリを書いた](https://hnakamur.github.io/blog/2016/06/13/wrote_go_ltsvlog_library/)
    - [zerologを参考にしてltsvlogを改良してみた](https://hnakamur.github.io/blog/2017/05/28/improve-ltsvlog-with-referring-to-zerolog/)

### エラーハンドリング

[error]: http://blog.golang.org/error-handling-and-go "Error handling and Go - The Go Blog"

- [または私は如何にして例外するのを止めて golang を愛するようになったか — KaoriYa](http://www.kaoriya.net/blog/2014/04/17/)
- [Big Sky :: golang で複数のエラーをハンドリングする方法](http://mattn.kaoriya.net/software/lang/go/20140416212413.htm)
- [DSAS開発者の部屋:Go ではエラーを文字列比較する？という話について](http://dsas.blog.klab.org/archives/go-errors.html) : エラーハンドリングには，定数との比較， conversion 構文による型の比較，エラー文字列の比較がある
- [panicはともかくrecoverに使いどころはほとんどない - Qiita](http://qiita.com/ruiu/items/ff98ded599d97cf6646e)
- [go で AggregationException(.NET)的なことをする - Qiita](http://qiita.com/kyoh86/items/6cadd79de08cc597b65a) : ループ等でエラーを集約してからまとめて処理する方法
- [echoのAPIサーバ実装とエラーハンドリングの落とし穴 - Qiita](http://qiita.com/tienlen/items/5f2bcfe06eb83830ee55)
- [Golangのエラー処理とpkg/errors | SOTA](http://deeeet.com/writing/2016/04/25/go-pkg-errors/)
    - [Golangでエラー時にスタックトレースを表示する - Qiita](http://qiita.com/deeeet/items/dacc71932393ab35d9f8)
- [Ginのミドルウェアを使ったエラーハンドリング - Qiita](https://qiita.com/safu9/items/b8c94bb911cb0a39d5aa)

### Struct タグについて

- [Goのencoding/jsonでタグが反映されなくてハマったしょうもない話 - Qiita](http://qiita.com/modal_soul/items/83b0930d90d44e006768)
- [Go で struct のタグ情報を取得する - hiyosi's blog](http://hiyosi.tumblr.com/post/100922038678/go-%E3%81%A7-struct-%E3%81%AE%E3%82%BF%E3%82%B0%E6%83%85%E5%A0%B1%E3%82%92%E5%8F%96%E5%BE%97%E3%81%99%E3%82%8B)
- [Goのjson.Marshal/Unmarshalの仕様を整理してみる - I Will Survive](http://blog.restartr.com/2014/08/13/golang-json-marshal-unmarshal/)
- [GoでJsonファイルを読み込んで構造体として扱う。 - Qiita](http://qiita.com/niiyz/items/3f522c0e5a32de916ec6)
- [BurntSushi/tomlを使ってハマったこと - Qiita](http://qiita.com/reiki4040/items/6556d4eba797329e9f51) : [BurntSushi/toml](https://github.com/BurntSushi/toml) にバグがあるという話
- [GoでJSONの一部分を利用者が定義した構造体に読み込める便利な手法を見つけた - Qiita](http://qiita.com/hnakamur/items/ba363e82332d4dbdf34a)
- [golang は ゆるふわに JSON を扱えまぁす! — KaoriYa](https://www.kaoriya.net/blog/2016/06/25/)
- [Go言語でJSONに泣かないためのコーディングパターン - Qiita](http://qiita.com/minagoro0522/items/dc524e38073ed8e3831b)
- [Go 言語 1つの構造体に複数の validation を適応する - Qiita](http://qiita.com/iktakahiro/items/2e240147ca3188948a17)
- [Go で関数の引数用構造体のバリデーションと初期化をするメソッドをタグから生成するツールを作った - Qiita](http://qiita.com/nirasan/items/7b96f080cd286c324033)
- [Goでsql.NullStringを含む構造体をjson.Marshalする方法 - Qiita](http://qiita.com/oikyn/items/26808b03997797f1b342)
    - [GoでJSONのnullをいい感じに扱いたい - Qiita](https://qiita.com/catatsuy/items/745237b4e797c5d9d4fb)
- [Go言語でJSON内の整数は10進数6桁しか表現できない - Qiita](http://qiita.com/toast-uz/items/52f0c86716493ad3ca12)
- [JSONSchemaからstructのようなコードを生成する"structr"というのを書いた - Qiita](http://qiita.com/damele0n/items/92a9b845c991b1b29aea)
- [GolangでEnumをフィールドに持つstructをいい感じにjsonエンコード / デコードする - 一から勉強させてください(￣ω￣;)](http://dangerous-animal141.hatenablog.com/entry/2017/01/19/004650)
- [golang xml.Marshal でxmlタグで出力する - Qiita](https://qiita.com/dolpher/items/024af13c37926218c3f5)
- [golangでajaxを使用してrssのデータを取得する - m_shige1979のささやかな抵抗と欲望の日々](http://m-shige1979.hatenablog.com/entry/2016/02/19/080000)
- [Big Sky :: GolangでAPI Clientを実装する、の続き](https://mattn.kaoriya.net/software/lang/go/20161101151118.htm)
    - [GolangでAPI Clientを実装する | SOTA](https://deeeet.com/writing/2016/11/01/go-api-client/)

### [`time`] パッケージ

[`time`]: http://golang.org/pkg/time/ "time - The Go Programming Language"

- [Golang 日付のフォーマットでハマった話 - Qiita](http://qiita.com/masa23/items/e781124a7e0305bc40c4)
- [golangのtime.Timeの当日00:00:00を取得する方法とベンチマーク - Qiita](http://qiita.com/ushio_s/items/3e270933641710bbd88e)
- [Golangで周期的に実行するときのパターン - Qiita](https://qiita.com/tetsu_koba/items/1599408f537cb513b250)
- [Go: Timers’ Life Cycle. ℹ️ This article is based on Go 1.14. | by Vincent Blanchon | A Journey With Go | Jul, 2020 | Medium](https://medium.com/a-journey-with-go/go-timers-life-cycle-403f3580093a)

### [`reflect`] および [`unsafe`] パッケージ

[`reflect`]: https://golang.org/pkg/reflect/ "reflect - The Go Programming Language"
[`unsafe`]: https://golang.org/pkg/unsafe/ "unsafe - The Go Programming Language"

- [Go と reflect と generate - Qiita](http://qiita.com/naoina/items/7966f73f3a807b3d25d6)
- [unsafe が unsafe なケース (1) - Qiita](http://qiita.com/kwi/items/185bb3fe0d60ca765ab0)
    - [unsafe が unsafe なケース (2) - Qiita](http://qiita.com/kwi/items/d06f49c9cf7e5ace8692)
- [Goで関数型プログラミング - Qiita](http://qiita.com/taksatou@github/items/d721a62158f554b8e399) : [reflect](https://golang.org/pkg/reflect/ "reflect - The Go Programming Language") パッケージを使って高階関数を表現できる

### [`context`] パッケージ

[`context`] はバージョン 1.7 から標準パッケージに組み込まれた。

[`context`]: https://golang.org/pkg/context/ "context - The Go Programming Language"

- [Go1.7のcontextパッケージ | SOTA](http://deeeet.com/writing/2016/07/22/context/)
- [Golangのcontext.Valueの使い方 | SOTA](http://deeeet.com/writing/2017/02/23/go-context-value/)
- [contextの使い方 - Qiita](http://qiita.com/taizo/items/69d3de8622eabe8da6a2)
- [context.Contextでリクエストスコープな値を持ち回す - Qiita](http://qiita.com/hogedigo/items/a26d816395b7545ce5f8) : [context](https://golang.org/pkg/context/) の使い方って（名前からいって）本来こっちだよね
- [goroutine にシグナルを送信する - Qiita](http://qiita.com/Kei-Kamikawa/items/620f9504daf2ec53f0b5)
- [Go言語のContextパッケージのTODO( )って何？いつ使うの？ - Qiita](https://qiita.com/po3rin/items/3556c36182e0c635791c)
- [Go 言語 context パッケージ誕生の背景と使用方法](https://ayasuda.github.io/pages/what_is_context_at_go.html)
- [context.TODO()を使って漸進的にcontext対応を始める - My External Storage](https://budougumi0617.github.io/2020/02/21/use-context/)

## 開発支援

- [GoのSSA最適化制御オプション - Qiita](https://qiita.com/tooru/items/a55bcdac0500d9a93f39)
- [Big Sky :: gocode やめます(そして Language Server へ)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)
- [Big Sky :: Go 言語の Language Server「gopls」が completeUnimported に対応した。](https://mattn.kaoriya.net/software/lang/c/20191112100330.htm)

### デバッガ

- [golang でビルド時に最適化をオフにする - tetsuok の旅 blog](http://tetsuok.hatenablog.com/entry/2012/07/01/062325) : gdb でデバッグする際は最適化を off にするといいという話
- [Go言語のトラブルシューティング用機能](http://www.slideshare.net/satorutakeuchi18/go-53685632)
- [Go で利用できるプロファイリングツール pprof の読み方 - Qiita](http://qiita.com/ikawaha/items/e3b35f09fb49e9217924)
- [Goでfunctionが実行された順番を追いかける - sgykfjsm.github.com](http://sgykfjsm.github.io/blog/2016/01/20/golang-function-tracing/)
- [Go言語でプリント文デバッグするときのTips - Qiita](http://qiita.com/ohac/items/0aa8eb6ff8ee5f599dcd)
- [GitHub - derekparker/delve: Delve is a debugger for the Go programming language.](https://github.com/derekparker/delve)
    - [Golangのデバッガdelveの使い方 - Qiita](https://qiita.com/minamijoyo/items/4da68467c1c5d94c8cd7)
- [Go言語のバリデーションチェックライブラリ(ozzo-validation)を分かりやすくまとめてみた - Qiita](https://qiita.com/gold-kou/items/201a19d9d0c760cc2104)
- [go-playground/validator リクエストパラメータ向けValidationパターンまとめ - Qiita](https://qiita.com/RunEagler/items/ad79fc860c3689797ccc)
- [Go言語のモック(gomock)を触ってみた - Qiita](https://qiita.com/gold-kou/items/81562f9142323b364a60)

### テスト・フレームワーク

- [Go の Test に対する考え方 - Qiita](http://qiita.com/Jxck_/items/8717a5982547cfa54ebc)
- [Goでテストを書く - 成らぬは人の為さぬなりけり](http://straitwalk.hatenablog.com/entry/2014/09/18/232810)
- [golang 1.4で追加されたtestingの便利機能(テストの初期化とお片づけ) - Qiita](http://qiita.com/umisama/items/0d589cca7e89b89c29a8)
- [gojiのレスポンス結果をテストする - Qiita](http://qiita.com/taizo/items/11897f6284159919f65a)
- [Go Mockでインタフェースのモックを作ってテストする #golang - Qiita](http://qiita.com/tenntenn/items/24fc34ec0c31f6474e6d)
- [Go でベンチマーク - Block Rockin’ Codes](http://jxck.hatenablog.com/entry/20131123/1385189088)
- [go言語でベンチマーク - Qiita](http://qiita.com/Mulyu/items/ed585f2777496f29a725)
- [プロダクト開発でのGoのテストとモック活用事例 - Qiita](http://qiita.com/peroli-hirayama/items/f1419db7264fa9f9fe8f)
- [Go言語でファジング | SOTA](http://deeeet.com/writing/2015/12/21/go-fuzz/)
- [GAE/GoとGojiの組み合わせでテストを書く - Qiita](http://qiita.com/yosukesuzuki/items/c9e5c19df97d2ad5595a)
- [`go test -count n -timeout t` - Qiita](http://qiita.com/AkihiroSuda/items/0fd83df29182d4f5cdef)
- [Big Sky :: Re: golangでIOへのテストを行う](http://mattn.kaoriya.net/software/lang/go/20160328114637.htm)
- [Golang におけるサブテストの並行処理実装について | eureka tech blog](https://developers.eure.jp/tech/go1_7-subtests/)
- [GAE/Goのテストを実行するために必要なこと - Qiita](http://qiita.com/ttyokoyama/items/5b99299ec112b580f03b)
- [Golangでテストしやすいコードをかく - Qiita](http://qiita.com/chisso/items/1dcc52f404b88d274f29)
- [Go な WebAPI のテスト＆ドキュメントの模索 - Qiita](https://qiita.com/zaru/items/0bee6c19b056dc72948d)
- [GAE/Goで書いたアプリのテストが何故か固まるのでgo testのコードをチラ見したメモ - utahta >> log](http://www.ninxit.com/blog/2017/12/08/091108)
- [Go言語のHTTPサーバのテスト事始め - Qiita](https://qiita.com/theoden9014/items/ac8763381758148e8ce5)
- [golang で実際にメール送信せず smtp.SendMail を試すためのモック作り - Qiita](https://qiita.com/Blufe/items/09b63eb113b5ba0064a8)
- [Goのテスト結果をCIでちょっと詳しく表示する - Qiita](https://qiita.com/take_cheeze/items/4c987cbde09807f8806b)
- [go vet の shadow を知る – Eureka Engineering – Medium](https://medium.com/eureka-engineering/-dabccd4571ea)
- [Ginkgoでgolangのビヘイビア駆動(BDD)開発入門 - Qiita](https://qiita.com/myoshimi/items/62bc89b8065e08834b02)
- [Goのtestを理解する in 2018 #go - My External Storage](https://budougumi0617.github.io/2018/08/19/go-testing2018/)
- [Go: Fuzz Testing in Go - A Journey With Go - Medium](https://medium.com/a-journey-with-go/go-fuzz-testing-in-go-deb36abc971f)
- [Goのtestを理解する in 2019 - My External Storage](https://budougumi0617.github.io/2019/10/30/go-testing2019/)
- [Goのtestを理解する - httptestサブパッケージ編 - My External Storage](https://budougumi0617.github.io/2020/05/29/go-testing-httptest/)

### ドキュメント・フレームワーク

- [Go で godoc を使えるようにする〜godoc のインストール〜 - Qiita](http://qiita.com/megu_ma/items/2066aef2f8c7f0ce2cc3)
- [Go言語のコードレビュー | SOTA](http://deeeet.com/writing/2014/05/26/go-code-review/)
- [Go コードのレビュー時によくされるコメント - build error](http://papaeye.tumblr.com/post/92328649161/go)
- [testingパッケージのExamplesについて - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20131101/1383285018)
- [GoのExampleテストが便利 : swdyh](http://swdyh.tumblr.com/post/55771477125/go-example)
- [godoc.org への掲載方法を調べた - taknb2nchのメモ](http://d.hatena.ne.jp/taknb2nch/20140225/1393302743)
- [[swaggo]GoのGoDocを書いたら、Swaggerを出せるやばいやつ - Qiita](https://qiita.com/pei0804/items/3a0b481d1e47e5a72078)
- [Go CodeReviewComments 日本語翻訳 #golang  - Qiita](https://qiita.com/knsh14/items/8b73b31822c109d4c497)
- [チョットできるGoプログラマーになるための詳解GoDoc - Qiita](https://qiita.com/shibukawa/items/8c70fdd1972fad76a5ce)

### Refactoring

- [ジェネレートしたコードを gofmt / goimports する - Qiita](http://qiita.com/minodisk/items/d96a0673223f36315ce7)
- [Big Sky :: golang のリファクタリングには gofmt ではなく、gorename を使おう。](http://mattn.kaoriya.net/software/lang/go/20150113141338.htm) : リファクタリングには gofmt よりも gorename が使えるという話
- [struct にアノテーションつけてたら go vet . すべき - Qiita](http://qiita.com/amanoiverse/items/fcd25db64f341ad2471f)
- [これからGo言語を書く人への三種の神器 - Qiita](http://qiita.com/osamingo/items/d5ec42fb8587d857310a) : `go vet`, `goimports`, `golint` で正しいコードを書きましょう。
- [golintと闘いたいけど心の折れてしまった勇者のための隠しダンジョン - Qiita](http://qiita.com/kyoh86/items/1f2022b63372b84f1a27)
- [gometalinter で楽々 lint - Qiita](http://qiita.com/spiegel-im-spiegel/items/238f6f0ee27bdf1de2a0) : 各種 lint を統合的に扱える
- [golangのループ変数の使い方をチェックするlinter作ってみた - Qiita](http://qiita.com/kyoh86/items/97911180d6254d5fc90c)

### Continuous Integration

- [Go + Travis CI + Coveralls でCI環境を作る - Qiita](http://qiita.com/dmnlk/items/3fb4e0abb98e39fee275)
- [GithubにあるリポジトリをTravis CI連携する手順 #junitbook - くりにっき](http://sue445.hatenablog.com/entry/2013/06/01/170607)
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [golangでTravis CIを使ってクロスコンパイルするときにハマったところ #golang #travisci - uchimanajet7のメモ](http://uchimanajet7.hatenablog.com/entry/2015/03/20/211352)
- [Go言語のビルド生活を drone.ioで幸せに暮らす #golang - Qiita](http://qiita.com/atotto/items/b796c31c1755dbec13db)
- [Go+Webアプリケーション+CircleCIで静的解析・テスト・バイナリリリースを効率良く行なう - Qiita](http://qiita.com/kaneshin/items/163626c09c1ad9818c6c)
- [Circle CI 2.0の基礎的な設定まとめてみた(GAE/Goのサンプル付き) - Qiita](https://qiita.com/Sekky0905/items/7f9aa94261e17e4fd040)
- [goreleaser/goreleaser: Deliver Go binaries as fast and easily as possible](https://github.com/goreleaser/goreleaser)
    - [goreleaserを使ってGoで書いたツールのバイナリをGithub Releasesで配布する - $shibayu36->blog;](http://blog.shibayu36.org/entry/2017/10/04/193000)
    - [goreleaser と Travis CI で Golang のバイナリ配布を自動化する - /storage/tummy.log](http://rnitame.hatenablog.com/entry/automate-golang-binary-distribution)
- [Go言語利用のレポジトリをInspecodeで静的解析＋テストしてみる - Qiita](https://qiita.com/ks888/items/65d3be0f05e1aecc817b)
- [【GitHub Actions】Go言語の自動テストからリリースまでを作ってみた - Qiita](https://qiita.com/x-color/items/f60025c20a547a7355b5)

### クロス環境

- [Goはクロスコンパイルが簡単 - unknownplace.org](http://unknownplace.org/archives/golang-cross-compiling.html)
- [Go のクロスコンパイル環境構築 - Qiita](http://qiita.com/Jxck_/items/02185f51162e92759ebe)
- [golang でのクロスコンパイルの留意事項 — KaoriYa](http://www.kaoriya.net/blog/2015/03/06/) : Windows 環境でクロス環境を構築する際の注意点。
- [Golang + Raspberry Pi + LPS331AP で気圧・温度を測定してみた - Qiita](http://qiita.com/yanolab/items/5a6dfb3c07c94f7c760d) : Raspberry Pi 用のクロス環境例。
- [Gobot - Golang framework for robotics, physical computing, and the Internet of Things](http://gobot.io/)
- [複数プラットフォームにGoアプリケーションを配布する | SOTA](http://deeeet.com/writing/2014/05/19/gox/)
    - [Go1.5はクロスコンパイルがより簡単 | SOTA](http://deeeet.com/writing/2015/07/22/go1_5-cross-compile/)
- [MacOS X でGo言語のクロスコンパイルを試したらハマった - Qiita](http://qiita.com/ttsuzo/items/64e29dd7caa635ac7863) : [gox](https://github.com/mitchellh/gox) を使う方法
- [Goで64bitと32bitの実行ファイルを同一Windows機で作成するために講じたこと - Qiita](http://qiita.com/zetamatta/items/e44961a8bcbb2578cfe7)
- [Travis-CI で Go の Windows 用バイナリを Github release に登録する - Qiita](http://qiita.com/methane/items/f8c5a5f2209739daf44e)
- [gopherjs + electron テスト - Qiita](http://qiita.com/shizu/items/c8a28e0d2299868dafa9) : [`gopherjs/gopherjs`](https://github.com/gopherjs/gopherjs) を使って Go のコードから javaScript コードを生成できるらしい
- [Raspberry PI ２ 用の consul を作る (201512版 - Qiita](http://qiita.com/rerofumi/items/d6a8ba08270acb61b31c) : Raspberry PI 上でビルドするより Linux のクロス環境を使ったほうが速いらしい
- [CI-as-a-ServiceでGo言語プロジェクトの最新ビルドを継続的に提供する | SOTA](http://deeeet.com/writing/2014/10/16/golang-in-ci-as-a-service/)
- [RaspberryPi1(2とzeroも)で動かすgolang製アプリをクロスコンパイル(onMac) - Qiita](http://qiita.com/m0a/items/d933982293dcadd4998c)
- [GoでFPGAしてみる(Reconfigure.io) - Qiita](https://qiita.com/mjhd-devlion/items/5e6f6f2f40ecb4ad4217)
- [GoでCLIツールを作成してRaspberryPiで実行する - Qiita](https://qiita.com/marumaru-n/items/13ef552adad2d3a4ed1e)

### C 言語との連携

- [cgoでGoのコードからCの関数を利用する - 1000ch.net](https://1000ch.net/posts/2014/c-in-golang-with-cgo.html)
- [cgoでGolangとC++ライブラリをリンクするとき、何が起きているのか - beatsync.net](https://beatsync.net/main/log20141022.html)
- [GO 1.5 と C++ を SWIG でブリッジさせる方法 - Qiita](http://qiita.com/Satachito/items/5a0d7dd228d3272e0780)
- [cgoを使ったCとGoのリンクの裏側 (1) - Qiita](http://qiita.com/yugui/items/e71d3d0b3d654a110188)
    - [cgoを使ったCとGoのリンクの裏側 (1) - Qiita](http://qiita.com/yugui/items/e71d3d0b3d654a110188#_reference-b69de9de7311c6e17e7f)
- [Golang で Static Library を作る際、stringをparameterで受け取るならコピーしよう。 - Qiita](http://qiita.com/flowtumn/items/2df066ca776023bcc687)
- [マルチプラットフォーム対応したライブラリ Golang - Qiita](http://qiita.com/nothingcosmos/items/b5dc76aa953222bbdb5c)
- [Big Sky :: golang の Windows 版が buildmode=c-archive をサポートした。](http://mattn.kaoriya.net/software/lang/go/20160405114638.htm)
- [Golang で過去の遺物的(cp932)DLLを利用する - Qiita](http://qiita.com/asm/items/6184cf5dcca637670e0e)
- [Big Sky :: Golang で Windows の DLL を作る方法](http://mattn.kaoriya.net/software/lang/go/20160921010820.htm)
- [GoでShared Libraryをビルドしてみた(簡単ドキュメント指向DB) - Qiita](http://qiita.com/umisama/items/48ba74a58f1e6530e305)
- [cgoやってみた - Qiita](https://qiita.com/shunsukuda/items/e9646e1a6acc863f3594)
- [GoとCの間のポインタ渡し - Qiita](https://qiita.com/74th/items/0362bea2012ef253c539)
- [Cgoを使ったパッケージと「Docker as Bug Report/Reproduce」というOSS運用について - DRYな備忘録](http://otiai10.hatenablog.com/entry/2017/12/04/013819)

### 組込み開発

主なものは「[組込みで Go]({{< relref "embedded-engineering-with-golang.md" >}})」に移動。
以下は携帯端末開発に関するブックマーク：

- [go 1.5でgomobile(android) - unokun’s blog](http://unokun.hatenablog.jp/entry/2015/08/01/150628)
- [gomobileでiOSアプリをビルドする手順まとめ - GolangRdyJp](http://golang.rdy.jp/2015/09/21/ios-gomobile/)
- [gomobileで日本語フォントを扱ってみる - Qiita](http://qiita.com/bowz_standard/items/5a9c987f9242777fff30)
- [GoでBenchmarking Raspberry Pi Zero W GPIO Speed - Qiita](https://qiita.com/WaToI/items/16938a611f7c026c477d)
- [Raspberry Pi 3 Model B+のUbuntu 64bitに golang の導入を試みて失敗した話 - Qiita](https://qiita.com/shibta/items/97369fa4ca8c5361ddf3)

### ビルド時に情報を各種埋め込みたい

- [Go言語: ビルド時にバージョン情報を埋め込みたい - Qiita](http://qiita.com/suin/items/d643a0ccb6270e8e3734)
- [Golangビルド時に、サブパッケージ内の変数をいじる - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/09/08/211626)
- [Goでビルドバージョン情報を参照できるようにする(Go1.5) - Qiita](http://qiita.com/reiki4040/items/6b32370532c3eafe1f0e)
- [go-bindata でコンパイル時にリソースを埋め込んじゃおう！ - Qiita](http://qiita.com/ikawaha/items/c02d84cfd00f8f442500)
- [ソースを実行ファイルに埋め込む方法 - Qiita](http://qiita.com/ymko/items/c2e3c8fe25bce425136d)
- [Golangで静的ファイルをバイナリに含めるライブラリを書いてみた - Qiita](http://qiita.com/konohazuku/items/131b251a5fa29213ac5e)
- [GolangのGin/bindataでシングルバイナリを試してみた(+React) - Qiita](http://qiita.com/wadahiro/items/4173788d54f028936723)
- [Goで任意のbuild tagsをつけてビルドファイルを切り替える - Qiita](http://qiita.com/ueokande/items/fac0d1219dbbc8f44db7)
- [Goでビルドしたバイナリに製品名やファイルバージョンを追加する方法](https://qiita.com/Targityen/items/6125d4fa83fca28879a9) : Windows 専用

### Build Tools for Golang

[constabulary/gb](https://github.com/constabulary/gb) を使ってプロジェクトベースの環境構築

- [golang - gbを知ったのでgojiを使ったWEBアプリケーションプロジェクトを管理してみた - Qiita](http://qiita.com/shinofara/items/ac0591fef95c2c6e936e)
- [Building Go projects with gb - Supermighty](https://walledcity.com/supermighty/building-go-projects-with-gb)
- [Go言語のDependency/Vendoringの問題と今後．gbあるいはGo1.5 | SOTA](http://deeeet.com/writing/2015/06/26/golang-dependency-vendoring/)

[FiloSottile/gvt](https://github.com/FiloSottile/gvt) というのがあるらしい。

[Masterminds/glide](https://github.com/Masterminds/glide) と Go 1.5 の Vendoring 機能を組み合わせてパッケージ管理できる。

- [glide - パッケージ管理のお困りの方へ - - Qiita](http://qiita.com/tienlen/items/8e192e68d6b18bec3b4a)

パッケージ依存解決ツールというのがあるらしい。

- [Big Sky :: golang オフィシャル謹製のパッケージ依存解決ツール「dep」](http://mattn.kaoriya.net/software/lang/go/20170125023240.htm)
- [dep(Go dependency tool)を自作ライブラリに使ってみた - Qiita](http://qiita.com/kwmt@github/items/0f77e083494ca94d782b)

### direnv で開発環境を切り替える

[direnv - unclutter your .profile](http://direnv.net/)

- [direnv/direnv](https://github.com/direnv/direnv)
- [direnvで解決するGOPATHの3つの問題点 - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/06/18/010449)
- [改めて、direnvを使いましょう！ - HDE BLOG](http://blog.hde.co.jp/entry/2015/02/27/182117)
- [さくら - homeにgolang, direnv とvirtualenvを入れて動かす - Qiita](http://qiita.com/aminamid/items/5a0e9461385c80d0c8a6)

### Integrated Development Environment (IDE)

- [EclipseでGoプログラミング！ GoClipseのインストールとGojiフレームワークを使ったWeb APIの作成 （1/6）：CodeZine](http://codezine.jp/article/detail/8374)
- [WindowsでGolang開発環境構築　IntelliJ IDEA - Qiita](http://qiita.com/ngsm3/items/67620fc4e39219235a23)
    - [第1回　Android Studio，そしてベースとなる「IntelliJ IDEA」とは何か？：Android Studio最速入門～効率的にコーディングするための使い方｜gihyo.jp … 技術評論社](http://gihyo.jp/dev/serial/01/android_studio/0001) : IntelliJ IDEA についての解説
- [Go入門@環境構築編(IntelliJ IDEA 14.1.4) - Qiita](http://qiita.com/makoto2468/items/3ced77de947ea997e72f)
- [WindowsにIntelliJ IDEAでgoの開発環境を作りましたよ　ついでにgit連携もいたしましたよ - Qiita](http://qiita.com/peka2/items/cdc9d508dcbf4a131271)
- [Cloud9にGo言語でHelloWorldを実装・実行してみた - Qiita](http://qiita.com/ryo-endo/items/e8cb4987b4af9ddc24d5)
- [GoLand: Capable and Ergonomic Go IDE by JetBrains](https://www.jetbrains.com/go/)

#### ATOM で Go

- [AtomでGoを書く環境を整える（Windows） - Technically, technophobic.](http://notchained.hatenablog.com/entry/2014/09/20/104829)
- [AtomでのGo言語開発環境セットアップ - Qiita](http://qiita.com/MakoTano/items/3d807a96c3933ac8aa13)

#### Visual Studio Code で Go

- [VisualStudioCode - Visual Studio CodeでGo言語の設定 - Qiita](http://qiita.com/evalphobia/items/f68396d573c7caf2065b)
- [VSCode+Golang](https://gist.github.com/llaughlin/16305f0b91356e93a1c4) : task.json の設定
- [VisualStudioCodeでGAE/Goの環境設定 - Qiita](http://qiita.com/m0a/items/ebcf964effbc7d50a489)

#### Emacs で Go

- [emacsでGoの環境構築をやる - Qiita](http://qiita.com/koki_cheese/items/2e2ead918a1f1ac5bf6e)

#### Vim で Go

- [Vim で golang を書く環境を整える - Qiita](http://qiita.com/masa23/items/db184871c78311566434)
- [Big Sky :: Vim で Go 言語を書くために行った引越し作業 2020年度版](https://mattn.kaoriya.net/software/vim/20200106103137.htm)
- [gopls 0.4.3で構造体を初期化（"fillstruct"）しようとしても、"No code actions found"とだけ表示される - My External Storage](https://budougumi0617.github.io/2020/07/18/use_fillstruct_of_goplus_on_vim/)

## CLI; Command Line Interface

- [Go言語によるCLIツール開発とUNIX哲学について - ゆううきブログ](http://yuuki.hatenablog.com/entry/go-cli-unix)
- [開発者から見た UNIX 哲学とコマンドラインツールと Go言語 - TELLME.TOKYO](http://tellme.tokyo/post/2015/06/23/unix_cli_tool_go/) （[Qiita 版](http://qiita.com/b4b4r07/items/df660d82e2de715acda5)）
- [コマンドライン引数 - Qiita](http://qiita.com/uokada/items/f0e069a751679dcf616d)
- [Go言語のflagパッケージを使う - uragami note](http://ryochack.hatenablog.com/entry/2013/04/17/232753)
- [Go を使ってコマンドラインツール wordc を作ってみた - Qiita](http://qiita.com/flaflasun/items/df5ebb057697da062a08) : [codegangsta/cli](https://github.com/codegangsta/cli) についての言及あり
- [GoでCLIツール作るのに便利そうなパッケージを集めてみました - Qiita](http://qiita.com/isaoshimizu/items/71dd2ca2a08ddb607e31)
- [C言語とGo言語で標準出力が端末を参照しているかどうかを判定する - uragami note](http://ryochack.hatenablog.com/entry/2013/07/15/232207)
- [Go言語のCLIツールのpanicをラップしてクラッシュレポートをつくる | SOTA](http://deeeet.com/writing/2015/04/17/panicwrap/)
- [flag 並にシンプルでより強力な CLI パーサ kingpin の紹介 - Qiita](http://qiita.com/kumatch/items/258d7984c0270f6dd73a)
- [Goで外部コマンドをパイプして実行する - Qiita](http://qiita.com/yuroyoro/items/9358cd25b5f7fe9dd37f)
    - [Big Sky :: golang で UNIX コマンドパイプラインを扱う](http://mattn.kaoriya.net/software/lang/go/20151030131242.htm)
- [Golangで外部コマンドを実行する方法まとめ - Qiita](http://qiita.com/tanksuzuki/items/9205ff70c57c4c03b5e5)
- [spf13/cobra: A Commander for modern Go CLI interactions](https://github.com/spf13/cobra)
    - [GolangでCLIの場合にcobraを使うことにした件 | FLAMA技術Blog](http://lab.flama.co.jp/archives/1536/)
- [GolangでwebサービスのAPIを叩くCLIツールを作ろう - Qiita](http://qiita.com/minamijoyo/items/cfd22e9e6d3581c5d81f)
- [GoでMySQLにCA証明書を使ってアクセスする - Qiita](http://qiita.com/Peranikov/items/4376633bb72492051336)
- [Golangのコマンドライブラリcobraを使って少しうまく実装する - Qiita](https://qiita.com/unchemist/items/3cdeafcde2bd98612428)

## GUI; Graphical User Interface

- [Go用のGoogle製のGUIツールキットgxuiのインストール(Windows版) - Qiita](http://qiita.com/sago35/items/cc9ed3dc38d0b2f19bf9)
- [goでwindowsでwindow - Qiita](http://qiita.com/ohisama@github/items/20bc61175ce4a33b7365)
- [今後イケそうなデスクトップGUIフレームワーク - Qiita](http://qiita.com/hachi8833/items/463264f531474a856064) : [go-thrust](https://github.com/miketheprogrammer/go-thrust/) について言及
- [GolangのOpenGL事情(WebGLも含むよ) - Qiita](http://qiita.com/shibukawa/items/58f6a421462b93dec471)
- [Goで3Dモデル変換してプレビュー - Qiita](http://qiita.com/tetuyoko/items/746599e36ca4985d9e1a)
- [Go最後の秘宝「GUI」を探しに行く - Qiita](http://qiita.com/shibukawa/items/cd8d122dfeb41e1608d1)
- [golangのGUIパッケージgo-gtkを試す - Qiita](http://qiita.com/intelf___/items/2207c02c306a495d99e6)
- [gocui の基本的な使い方 - Qiita](https://qiita.com/zenwerk/items/97ebd5e470bdafdb6e42) : CUI (Console User Interface)
- [GolangでクロスプラットフォームGUIアプリを作る - Qiita](https://qiita.com/nozo_moto/items/40e0fd89bd7fc3eb8b5d)
- [GoCV - Golang Computer Vision Using OpenCV 4](https://gocv.io/)
    - [golangでOpenCVを使おう - Qiita](https://qiita.com/besood/items/0045c62b3bc09332c421)

## Web Microframework for Golang

- [Mithril＋golang Gin を試す - Qiita](http://qiita.com/masatsugumatsus/items/e28254ff52963705ce7f)
- [Big Sky :: golang で最近お気に入りの WAF「Goji」](http://mattn.kaoriya.net/software/lang/go/20141021134209.htm)
- [gojiのMiddlewareの使い方 - Qiita](http://qiita.com/reiki4040/items/a038f1b99e0caee97d3e)
- [Gojiを使ってWebビーコン作る - Qiita](http://qiita.com/sys_cat/items/1b8581de1344cc5db6bb)
- [Google App EngineでGoのウェブアプリケーションをまず動かしてみる - Qiita](http://qiita.com/taizo/items/845fcfc58cfd0b30020a)
- [Go言語(Go-Json-Rest)でAPIサーバーを立てるときのCORS設定 (Basic認証機能付きも) - Qiita](http://qiita.com/n0bisuke/items/45ab414fc11959fc27c7)
- [Goでwebサーバー作るときに考えたこと - Qiita](http://qiita.com/taizo/items/e6597c66c3494d545686)
- [Golang の net/url で # を含む文字列を Parse() する - Qiita](http://qiita.com/tchssk/items/cb208f9ccd0a1819bbfa)
- [Golang で OpenStreetMap ファイル（osm.pbf）の読み込み（osmpbf 利用） - Qiita](http://qiita.com/kkdd/items/bd653e3d02546d1aa340)
- [GoのWEBアプリケーション運用について - Qiita](http://qiita.com/masahikoofjoyto/items/f60188f4252e455541d4)
- [go-qmlのWebViewでローカルWebサイトを表示するときの注意 - Qiita](http://qiita.com/hachi8833/items/315642ceecb378cabcb4)
- [http.Clientのタイムアウトの時間を変更 - Qiita](http://qiita.com/eielh/items/2e5fabb707355253b187)
- [golangでhttpを監視するscriptを書いて結果をslackにpostする - Qiita](http://qiita.com/kenjiszk/items/7ae842415ec392822612)
- [Go言語(Go-Json-Rest)のCORSでのハマり。Safariだけでハマった話。 - Qiita](http://qiita.com/n0bisuke/items/65c4a473a1fdbbf931f0)
- [GoでGoogleCalendarAPIv3を叩く（APIキー） - Qiita](http://qiita.com/yodatomato/items/8013f728bbf4358e9425)
- [Golang で iso-2022-jp メールのデコード - Qiita](http://qiita.com/curious-eyes/items/3dae99e5e0feb6b9f642)
- [Go言語の練習用にTwitterのOAuth認証をフルスクラッチしてみた - Qiita](http://qiita.com/mpyw/items/cb0f824d618d8fed384e)
- [Goとtesseractで簡易OCRサーバを作る - Qiita](http://qiita.com/fumizp/items/63243cf418d27898f208)
- [Go言語で簡単にHTTPリクエストを送ってJSONをパースするサンプル - Qiita](http://qiita.com/ikuwow/items/c8f494bbd16adf6db142)
- [UnixListener.Closeでソケットファイルが消えて困っている - Qiita](http://qiita.com/hiratara/items/0f0b6103a0dc9280cea9)
- [Goフレームワークのパフォーマンス比較 - Qiita](http://qiita.com/ToruFukui/items/eb0d3593b20a5e2f15c6)
- [一定時間だけ立ち上がるサーバーを書いた - Qiita](http://qiita.com/uokada/items/c30e26cd8bdee7dfe7eb)
- [Big Sky :: golang のミドルウェアとして組み込むだけでパフォーマンス改善が見込める「HTTP Coala」](http://mattn.kaoriya.net/software/lang/go/20151120113437.htm)
- [Go付属のツールでオレオレ証明書を生成する - Qiita](http://qiita.com/ono_matope/items/6c0de3e31642dfd17695) : なるほど（笑）
- [Go http.RoundTripper 実装ガイド - Qiita](http://qiita.com/tutuming/items/6006e1d8cf94bc40f8e8)
- [Go で静的 HTTP サーバを作る - Qiita](http://qiita.com/skitaoka/items/a2b55cb08060aa8d6a52)
- [[golang]RevelでCSRFの対策 - Qiita](http://qiita.com/ponchi/items/7e474041484841ee1ce6)
- [net/httpでポート443のHTTPSサーバーを立ち上げるまで - Qiita](http://qiita.com/ryurock/items/f55db5944397619735bf)
- [概観からGoのWebFrameworkを選ぶ（2016/02） - Qiita](http://qiita.com/jumbOS5/items/45f86db15a5a6c8a0622)
- [Revel(Golang)でViewを理解する - Qiita](http://qiita.com/jumbOS5/items/d817bc95279877e72b24)
- [Big Sky :: golang で画像アップロードが簡単に出来る go-imageupload を使ってみた。](http://mattn.kaoriya.net/software/lang/go/20160224103638.htm)
- [Go言語のスクレイピング系ライブラリまとめ - Qiita](http://qiita.com/okataitai/items/db6999ea1ab39ec0bd3e)
- [Go の echo ってWebサーバーでサクッと REST しちゃう - Qiita](http://qiita.com/keika299/items/62e806ae42828bb3567a)
- [nginx+circus+gojiによるgolang webアプリケーションの動作環境構築 - Qiita](http://qiita.com/reiki4040/items/795a008d1b12ee657d9a)
- [Big Sky :: golang で HTTP を使った処理を4倍速くする](http://mattn.kaoriya.net/software/lang/go/20160329094503.htm)
- [GolangのフレームワークEchoの話とHelloWorldサンプルを読み解くだけ - Qiita](http://qiita.com/CST_negi/items/5ea3395b35e68fd0d3b3)
- [Golang+Echo+dbrでMySQLのCRUDをする／JSONでDBの値を返却する話 - Qiita](http://qiita.com/CST_negi/items/5e276ddc0412cefef7e3)
- [Gin(Golang)におけるHTMLテンプレート記述方法 - Qiita](http://qiita.com/lanevok/items/dbf591a3916070fcba0d)
- [GoLang で html/template を使っていて遭遇したエラーとその解決 - Qiita](http://qiita.com/kent_ocean/items/45e153d2c5467501a20b)
- [GAE/Go (echoフレームワーク)で Line Message API 使って Bot を作る。 - Qiita](http://qiita.com/naoki_koreeda/items/8c818a3e9f6138ddbb87)
- [Go製のフレームワークechoを使ってJSONを返すWebサーバーを作り、GoogleAppEngineで動かす - Qiita](http://qiita.com/qube81/items/701279c43b33ce923613)
- [go+ginでローカルで作ったWebアプリをGoogleAppEngineに載せる時に注意することまとめ - Qiita](http://qiita.com/CST_negi/items/bcb4730c6efeb838c4a7)
- [社内のバックエンド開発にgRPCを導入してみた - Qiita](http://qiita.com/nozaq/items/9cd9bf7ee6118779bda9)
- [Echoはver.3で結局のところ何が変わったのか？ - Qiita](http://qiita.com/tienlen/items/85b06c0856c33e716c75)
- [Go言語のWEBフレームワークRevelを使用してセキュアなAPIを作成 - Qiita](http://qiita.com/GushiSnow/items/97c1f64c003b27c6b98a)
- [Go でツール書くときの Makefile 晒す - Qiita](http://qiita.com/dtan4/items/8c417b629b6b2033a541)
- [goa でデザイン・ファーストをシュッとする - Qiita](http://qiita.com/ikawaha/items/6638ee8b6978aef50d65)
- [echo 初心者でも簡単!! echo で扱うアセットファイル群を簡単にバイナリにまとめて使ってみる - Qiita](http://qiita.com/Kei-Kamikawa/items/a6cb72251b95c8f5baa1)
- [Go ライブラリによる CGIプログラム内ルーチング - Qiita](http://qiita.com/jyagaimo_qiita_/items/93d195ca65982b75e205)
- [Go 言語で Apache Bench (ab) を実装してみた - takatoshiono's blog](http://takatoshiono.hatenablog.com/entry/2017/02/06/013323)
- [HTTP/2 Server Push - The Go Blog](https://blog.golang.org/h2push)
- [Go + echoでfetch APIを使ってPOSTする - Qiita](http://qiita.com/yutasuzuki/items/fa7e78a8f4d8cb3e83b3)
- [Big Sky :: golang の http.Client を速くする](http://mattn.kaoriya.net/software/lang/go/20170112181052.htm)
- [Big Sky :: Re: Go でシングルバイナリな Web アプリを開発しているときに webpack --watch をうまいところやる](http://mattn.kaoriya.net/software/lang/go/20170119180147.htm)
- [gopher-lua でサーバーの設定を動的に変更する - Qiita](http://qiita.com/zenwerk/items/729ab53ad6925d80dafa)
- [【Go言語】gothでWebアプリを外部サービス認証ログインできるようにする - Qiita](http://qiita.com/momotaro98/items/90d12c10a655f4026d82)
- [Goでマイクロサービスやってみる〜gokit〜 - Qiita](https://qiita.com/miya-masa/items/1fefa42458857013b519)
- [GoでHTTPクライアントを書く時のURLの組み立て方 - Qiita](https://qiita.com/yyoshiki41/items/a0354d9ad70c1b8225b6)
- [goでWebサーバを書くためのシンプルなライブラリchiの紹介 - Qiita](https://qiita.com/tjun/items/3eea798905b597ec83c0)
- [Golangでパーセントエンコーディング - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/07/29/083729)
    - [Golangでパーセントエンコーディングをデコードする - 逆さまにした](http://cipepser.hatenablog.com/entry/2017/08/05/095807)
    - [encodeURIComponentが世界基準だと誤解してた話 - Qiita](https://qiita.com/shibukawa/items/c0730092371c0e243f62)
- [gRPCとREST APIでスループットを比較する - Qiita](https://qiita.com/muroon/items/1c9ad59653c00d8d5e3d)
- [Big Sky :: Golang と Vue.js で簡単なアプリケーションを作ってみた。](https://mattn.kaoriya.net/software/lang/go/20180330093346.htm)
- [【echo】ファイルのアップロード方法で詰まった話 - Qiita](https://qiita.com/witchy/items/85768165eb1038e045ec)
- [いつの間にか go の http/net が renegotiation に対応していた話 - Qiita](https://qiita.com/yossy6954/items/a0afd2e1082d2f0cafa3)
- [Big Sky :: golang の html/template でレイアウトを扱う方法](https://mattn.kaoriya.net/software/lang/go/20180418222903.htm)
- [ぼくのかんがえたさいきょうのまいくろさーびすあーきてくちゃ - Qiita](https://qiita.com/gaku3601/items/afb4bcd6e854e93e67e1)
- [【Go】ファイルアップロード - Qiita](https://qiita.com/huji0327/items/b8fee669323777a6d41a)
- [Go製WebToolKit Buffalo[概要編] - Qiita](https://qiita.com/k-kurikuri/items/f46356b70fe3e7e8da7d)
- [Go+Echoの環境にNewRelicを導入する - Qiita](https://qiita.com/daisukeoda/items/b02aefa4f464e63729b5)
- [golangのechoで静的なサイトを建てるならNowがお手軽という話 - DEV Community 👩‍💻👨‍💻](https://dev.to/fk2000/golangechonow-3kic)
- [golangでWebアプリケーションのルーティングを実装する！ - Qiita](https://qiita.com/micropig3402/items/ff2a3fd7673e849c5982)
- [Goのhttpルーター「Chi」の紹介 - Qiita](https://qiita.com/minoritea/items/afaf10de3c5b6ebafa84)
- [go で gorma を使ってAPI開発してみる。まずはgoa編 - Qiita](https://qiita.com/keiichi-hikita/items/0cebf00e85f47858c948)
- [grpc-gatewayでgRPCのREST対応を試しました - Qiita](https://qiita.com/nisitanisubaru/items/abe11ff4248997d1fee2)
- [Go言語のRESTサーバーを Clean Architecture で作ってみる - Qiita](https://qiita.com/nijuya_o/items/392cd3e4fa5641b3dec1)
- [SSOサービスKeycloakとgolangのHTTPサーバを連携する - Qiita](https://qiita.com/myoshimi/items/8337aab1b17d89938be0)
- [Go 言語のフレームワーク比較 - Qiita](https://qiita.com/inexp_eng4432/items/50d05262eb6ccbcfda1f)
- [Go言語のnet/httpクライアントでリダイレクトをやめる - Qiita](https://qiita.com/daijuk/items/d5c52b780e90387f2390)
- [ブラウザレンダリングの仕組み - Qiita](https://qiita.com/goemp/items/91dcc8b50d7a61ce98bc)
- [HTTPレスポンスボディの内容をログに残したい - My External Storage](https://budougumi0617.github.io/2020/06/21/record_response_body/)

## 他サービスとの連携

- [マイナンバーのチェックデジットをGoで計算する - Qiita](http://qiita.com/qube81/items/f66a38b28ec58bc5c4da)
- [golangでImageMagickを触りたい - Qiita](http://qiita.com/arc279/items/5f277aa5cce3de5247e5)
- [Go言語でRedshiftとつなぐ（というかただのPostgreSQL） - Qiita](http://qiita.com/otiai10/items/83b186596897705ce392)
- [Gmail API for Goで、下書き生成ツールを作りました。 - Qiita](http://qiita.com/yyoshiki41/items/1159e1a70ffaa8fd84ed)
- [Big Sky :: Windows からも ssh でリモートコマンド実行したい、それ golang で出来るよ](http://mattn.kaoriya.net/software/lang/go/20170111165324.htm)
- [コマンドラインからググれてもいいと思ったので作った - Qiita](http://qiita.com/ieee0824/items/13435fc6de5f22cdb2f4)
- [GoでSpreadsheetを操作するパッケージを作った - Qiita](http://qiita.com/Iwark/items/726dfb2d15a883e389b6)
- [Go言語で東京メトロAPIを叩く - Qiita](http://qiita.com/nayuneko/items/ca2651e3a613c8e3256a)
- [Twilio使って、入力した電話番号をチェックする方法 - Qiita](http://qiita.com/dbsparkle/items/72fc1ad50ba6000f2630)
- [Windows+GoでNFC/Felicaにアクセスしてみた - Qiita](http://qiita.com/mau4x/items/424fe7964e70a3a99965)
- [UnityのネイティブプラグインをGoで書く #golang #unity - Qiita](http://qiita.com/tenntenn/items/4d3316490a571e5d79ed)
- [radikoの録音ツールをGoで書いた - Qiita](http://qiita.com/yyoshiki41/items/f81442d7dc2d0ddcf15b)
- [Golang の DB 操作 ORM をいろいろしらべてみたい - Qiita](https://qiita.com/mochizukikotaro/items/b09116e0ad2d30e37098)
- [100万回のWebSocket接続とGo | プログラミング | POSTD](http://postd.cc/100%E4%B8%87%E5%9B%9E%E3%81%AEwebsocket%E6%8E%A5%E7%B6%9A%E3%81%A8go/)
- [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c)

### DB 連携

- [Go言語でBigQueryのクエリを実行してみる - Qiita](http://qiita.com/najeira/items/8310fecf4b70c918f855)
- [Go の DB アクセス用のパッケージを作った - Qiita](http://qiita.com/chimatter/items/1a5fb2f03477f2ada520)
- [MySQLでのトランザクション処理をGolang+dbrで実現してゆく話 - Qiita](http://qiita.com/CST_negi/items/d79600c34191adb09c79)
- [(メモ)mattn/go-sqlite3を使ってみた - Qiita](https://qiita.com/tukiyo3/items/89c773fd7ffce2adfadd)
- [GOのORM sqlboiler 使ってみた　 - Qiita](https://qiita.com/gougyan/items/5295e4a30697a73868b5)
- [GoでMySQLにアクセスしてみる（gorp編） - Qiita](https://qiita.com/hanenao/items/103774f76abdbc853abf)
- [go-sqlrow](https://blog.web-apps.tech/go-sqlrow/)
- [Go言語のために、すべてのORMに対応できるデータベースシャーディングライブラリを作った - Qiita](https://qiita.com/goccy/items/a54af6db3b8623e90c38)
- [GoのGORMでiterate - Qiita](https://qiita.com/iz-j/items/09097c494292e0f03636)
- [Big Sky :: SQLite3 でロジスティック回帰](https://mattn.kaoriya.net/software/ml/20190512003509.htm)
- [golang-migrate/migrateパッケージを使ってみる - Qiita](https://qiita.com/daijuk/items/2b43781c5a38923f864b)
- [GoのアプリケーションにDatadogAPMを導入する。 - Qiita](https://qiita.com/istsh/items/72b0f2c0ef345e57aaf4)
- [GoのDBライブラリについて調べてみた - Speaker Deck](https://speakerdeck.com/ryer/gofalsedbraiburarinituitediao-betemita)

### DB 連携 with GraphQL

- [graphql-go/graphql: An implementation of GraphQL for Go / Golang](https://github.com/graphql-go/graphql)
    - [golangでGraphQLの素振りを行った - Qiita](https://qiita.com/mitubaEX328/items/77ccc4f6ac0ad2e76996)

### GAE/Go

- [GAE/GoでCMSつくった - Qiita](http://qiita.com/hogedigo/items/342217982f267ccd234d)
- [Go+GAE+Cloud Datastoreで簡単なREST APIを構築 - Qiita](http://qiita.com/silverfox/items/81769425e51f24e676d2])
- [Google App Engine SDK for Goを使ってGAE上でアプリを動かすまで - Qiita](http://qiita.com/walkers/items/e407386d7ef184ec830a)
- [GAE/Goで形態素解析してみた - Qiita](http://qiita.com/mako0715/items/259659e5e2935d2afc10)
- [GAE/Go+glide的な構成での環境構築 ~ローカルサーバー立ち上げまで~ - Qiita](http://qiita.com/ryutah/items/d864310c62f0385d876d)
- [Go言語の依存管理ツールを作って、開発環境構築を覚えた - Qiita](http://qiita.com/eaglesakura/items/b7e92281735569c528a6)
- [GAE/Go で Google Cloud Spanner を操作する（前編） - Qiita](http://qiita.com/wezardnet/items/daf520b82e2199d16f4f)
- [GAE/Goでもgoroutine使おうぜ！というハナシ - Qiita](http://qiita.com/hogedigo/items/f0f409ee944c4b2107c3) : GAE/Go の API には非同期版がないので積極的に goroutine 使おうよ，という話
- [CircleCI 2.0でGlide管理のGoをGAEにデプロイする - Qiita](https://qiita.com/kenichi_odo/items/074f1b9541dd0487f1d8)
- [GAE(Google App Engine) で Golang を開発するための環境を構築する #golang - Qiita](https://qiita.com/ynozue/items/d12c1da13bbdb9213ba4)
- [GAE(Google App Engine) で Golang 初めての REST API #golang - Qiita](https://qiita.com/ynozue/items/22967c3c8e12129d7527)
- [Go×GAE×Dockerで作るGoogleOAuth認証アプリ - Qiita](https://qiita.com/gotokatsuya/items/7f3bd1f71aa825264851)
- [GAE/Goでメールを受信する](https://qiita.com/keitaro_1020/items/667bb1396015d32df09e)
- [GAE Go 開発環境の構築からテストアプリのデプロイまで【MacOS】](https://qiita.com/IJN-Penguin/items/782a23662d68aab1316c)
- [Go 言語で Google Cloud Storage の既存のバケットにオブジェクトを保存する](https://qiita.com/shinkiro/items/6d79b12d06de34119b46)
- [GAE/Go1.11試行（その1：「クイックスタート」） - Qiita](https://qiita.com/sky0621/items/8a42ee24cb417940228c)
- [Go言語でGoogle Drive APIとGmail APIを使う方法 - Qiita](https://qiita.com/KMim/items/f6e14cdaed8ad1907930)

### AWS Lambda

- [AWS LambdaのGoサポートについて今知れるいくつかのこと - Qiita](https://qiita.com/curepine/items/bd67276c9cae543bc0b8)

## Excel との連携

- [Go で簡単に Excelを作成するライブラリ。 色、罫線、網掛けを事前定義済 - Qiita](http://qiita.com/mikeshimura/items/b60823e923fb6d0840c0)
- [Go言語でエクセルファイル (.xlsx) を読み込む - Qiita](http://qiita.com/kaorumori/items/fa37130065d0450d6342) : [`github.com/tealeg/xlsx`](https://github.com/tealeg/xlsx) パッケージを使用
- [Go言語でExcel操作ライブラリを書いてみた - Qiita](http://qiita.com/tebakane/items/2f2ed2558357c274c478) : [`github.com/loadoff/excl`](https://github.com/loadoff/excl) パッケージの説明
- [Go言語でExcelファイルを処理するのが超簡単だった | 非IT企業に勤める中年サラリーマンのIT日記](http://pineplanter.moo.jp/non-it-salaryman/2017/06/18/go-read-excel/)

## Go で数学

- [大学入試問題をGoで解いてみる - Qiita](http://qiita.com/qube81/items/c47b9e3ea8d028e95588) : [math/big](https://golang.org/pkg/math/big/) パッケージを使って大きな数を計算する。
- [golangで数独を解いた - Qiita](http://qiita.com/ciruzzo/items/144bc1874947441f9fb8)
- [golang で AB x CD / E - F * G * H = 2016 になる全パターン洗い出し - Qiita](http://qiita.com/amanoiverse/items/06fff7b224d77517c08f)
- [Go で 0 から始まる連続する n 個の整数を重複無く k 個選んだ時の組み合わせの列挙 - Qiita](http://qiita.com/yumura_s/items/68760d6b902aee9c78f0)
- [ピーマンとハトと数学を Go 言語で試す - Qiita](http://qiita.com/nirasan/items/69643d0ddf8a7345cf7c)
- [golangでニュートン法を使って平方根の計算をする - Qiita](http://qiita.com/k-yamada@github/items/0a7baa61bd668c3cb3dc)
- [[Go]重み付き乱択アルゴリズムを整数だけで完結させる - Qiita](https://qiita.com/cia_rana/items/dca5b008fcee67adda50)
- [宣教師と人食い人種の問題をGolangで再帰するクロージャ使って書いてみた。 - Qiita](https://qiita.com/jun68ykt/items/2b7d788d21a4cc5c9a56)
- [golangによるグラフ理論ライブラリの実装](https://qiita.com/G0nta/items/7455fd0656693dd17ef6)
- [Big Sky :: golang で tensorflow のススメ](https://mattn.kaoriya.net/software/lang/go/20180825013735.htm)
    - [Big Sky :: Golang だけでやる機械学習と画像分類](https://mattn.kaoriya.net/software/lang/go/20181108123756.htm)
- [Big Sky :: TensorFlow Lite の Go binding を書いた。](https://mattn.kaoriya.net/software/lang/go/20190307190947.htm)
- [Go による機械学習 推論フレームワークの最新動向 2019 - Qiita](https://qiita.com/mattn/items/b01f9bb5c2fa3678734a)
    - [Big Sky :: Go 言語で TensorFlow の学習](https://mattn.kaoriya.net/software/ml/20190516010115.htm)
- [Go + gonum を使った行列計算まとめ - Qiita](https://qiita.com/po3rin/items/82c9c0195f9e3e38e2f1)
- [Goで作った数値計算の関数をまとめとく - Qiita](https://qiita.com/ken_sumi1019/items/cdf5ee01a4e2c725c185)

### ソートアルゴリズム

- [sliceのシャッフル - Qiita](http://qiita.com/sugyan/items/fd7138a756c1a409f5fd) : [Fisher–Yates shuffle](https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle) というアルゴリズムらしい
- [Go言語でバイトニックソート実装してみた - Qiita](http://qiita.com/nyamadandan/items/2c82011801b148c98e52)
- [Goでバケットソートアルゴリズム(ビット列を使用) - Qiita](http://qiita.com/ohkawa/items/269507985b3ae10cbff9)
- [interface{} をソートする - Qiita](http://qiita.com/tchssk/items/b61f1f06d22a6232d4c8)
- [Big Sky :: golang の sort インタフェース難しい問題が解決した](http://mattn.kaoriya.net/software/lang/go/20161004092237.htm)

### 暗号技術関連

- [Go言語と暗号技術（AESからTLS） | SOTA](http://deeeet.com/writing/2015/11/10/go-crypto/)
    - [tcnksm/go-crypto](https://github.com/tcnksm/go-crypto) : サンプルコード

以下は [Soshi Katsuta](https://github.com/skatsuta) さんによるテキスト。
丁寧な内容でとても参考になる。

- [Go 言語で学ぶ『暗号技術入門』Part 1 -DES, Triple DES- | Step by Step](http://skatsuta.github.io/2016/01/02/hyuki-crypt-book-go-1/)
- [Go 言語で学ぶ『暗号技術入門』Part 2 -AES- | Step by Step](http://skatsuta.github.io/2016/01/19/hyuki-crypt-book-go-2/)
- [Go 言語で学ぶ『暗号技術入門』Part 3 -CBC Mode- | Step by Step](http://skatsuta.github.io/2016/03/06/hyuki-crypt-book-go-3/)

## Garbage Collection の話

- [Go言語のリアルタイムGC　理論と実践 | プログラミング | POSTD](http://postd.cc/golangs-real-time-gc-in-theory-and-practice/)
- [Go言語の低レイテンシGC実現のための取り組み | プログラミング | POSTD](http://postd.cc/gos-march-to-low-latency-gc/)
- [引数で既存メソッドを渡した場合とlambdaで渡した場合でGCAllocに差がでた - Qiita](https://qiita.com/dimolto/items/3dbc5c20ce3c18d976c1)

## その他 四方山話

- [GASCII.jp：Goならわかるシステムプログラミング](http://ascii.jp/elem/000/001/235/1235262/)
- [O'Reilly Concurrency in Goの読書メモ - Qiita](https://qiita.com/YmgchiYt/items/420eaf2bcf7bee4ae152)
- [Goでクリーンアーキテクチャを試す | プログラミング | POSTD](http://postd.cc/golang-clean-archithecture/)
- [Golangの最新版をソースからビルドする - Qiita](https://qiita.com/tetsu_koba/items/bb3f2801373f64fbc661)
- [#golang CodeReviewComments 日本語翻訳  - Qiita](https://qiita.com/knsh14/items/8b73b31822c109d4c497)
- [Goにおける等値と等価の考察(struct1==struct2と&struct1==&struct2とreflect.DeepEqual(struct1,struct2)とreflect.DeepEqual(&struct1,&struct2)) - Qiita](https://qiita.com/Sekky0905/items/1ff4979d80b163e0aeb6)
- [golangのruntimeからアクセスする/etcのファイル - Qiita](https://qiita.com/tetsu_koba/items/d9447eadf9c419264603)
- [Go言語がWebAssemblyをサポートへ。GOARCHは「wasm」、GOOSは「js」に － Publickey](http://www.publickey1.jp/blog/18/gowebassemblygoarchwasmgoosjs.html)
- [Go にとても長い式を食べさせると死ぬ - Qiita](https://qiita.com/Nabetani/items/e850ad92ba79640bfcd7)
- [Language and Locale Matching in Go - The Go Blog](https://blog.golang.org/matchlang)
- [Go言語(Golang) はまりどころと解決策](http://www.yunabe.jp/docs/golang_pitfall.html)
- [go - Creating call graph in golang - Stack Overflow](https://stackoverflow.com/questions/31362332/creating-call-graph-in-golang#31369718)
- [Goのインクリメントとデクリメントのベンチマーク比較 - Qiita](https://qiita.com/qushot/items/306a9c9c0321304def8e)
- [Goの shortcircuit 最適化パスを読んだので解説する - Qiita](https://qiita.com/osamu329/items/3ac05c3b71f495fc1e61)
- [goumlでgoプロジェクトのUML図を出力する - Qiita](https://qiita.com/rubyu/items/d78470be2ddd9e86ebb5)
- [Goとrdtscの謎を追う - Qiita](https://qiita.com/kubo39/items/4319fa243fd18acc0981)
- [Go言語のアプリケーション設定・環境変数をStructにまとめる - Qiita](https://qiita.com/yuukive/items/27593cd6f3e7f264516b)
- [How a Go Program Compiles down to Machine Code - Better Programming - Medium](https://medium.com/better-programming/how-a-go-program-compiles-down-to-machine-code-e4532dc8b8ca)
- [[発表資料] 第138回RITS技術交流会『なぜ私たちはGoを書くのか。今あらためて考えるGo言語の良さと実際』 - My External Storage](https://budougumi0617.github.io/2019/10/05/jrits-why-go-how-is-go/)
- [Language and Locale Matching in Go - The Go Blog](https://blog.golang.org/matchlang)
- [OWASP/Go-SCPを読んでセキュアプログラミングとGoを学ぶ - My External Storage](https://budougumi0617.github.io/2019/12/04/introduce_go-scp/)
- [The Go Playground（play.golang.org）のショートカットキー - My External Storage](https://budougumi0617.github.io/2020/03/13/shortcut-keys-on-the-go-playground/)
- [The Zen of Go | Dave Cheney](https://dave.cheney.net/2020/02/23/the-zen-of-go)
    - [ブログ: Goの禅 (The Zen of Go)](https://okuranagaimo.blogspot.com/2020/03/go-zen-of-go.html)
- [「プログラミング言語Go完全入門」の期間限定公開のお知らせ - Mercari Engineering Blog](https://tech.mercari.com/entry/2020/03/17/120137)
- [Go vs Rust : 特徴量DBに適するのはどっち！？ - ABEJA Tech Blog](https://tech-blog.abeja.asia/entry/2020/04/09/115152)

### 他言語との比較もしくは移行

- [ErlangとGolangを比較してみる - Qiita](http://qiita.com/soranoba/items/68d57b4635a2917f3c73)
- [Gopherの道を歩む – Node.jsからGo言語への移行 | プログラミング | POSTD](http://postd.cc/the-way-of-the-gopher/)
- [Why Go? | Dave Cheney](https://dave.cheney.net/2017/03/20/why-go)
    - [[翻訳] Why Go? - Qiita](http://qiita.com/methane/items/b627f20457873a504638)
- [RubyからGoの関数をつかう → はやい - Qiita](http://qiita.com/grj_achm/items/679b3f3af2cf377f0f02)
    - [Perl6からGoの関数をつかう → はやい? - Qiita](http://qiita.com/B73W56H84/items/20a67b74bb646d140f7d)
    - [GroovyからGoの関数を使う→はやい - Qiita](http://qiita.com/yujiorama_at_github/items/3f7cab906969764cc805)
    - [Big Sky :: RubyからGoの関数をつかわなくても再帰をやめる → はやい](http://mattn.kaoriya.net/software/20151106194958.htm)
- [Ruby + mecabが遅いのでGoを経由する - Qiita](http://qiita.com/EastResident/items/f41fd0285fe270e7d3d5)

#### [Go 言語]に Generics がない理由

- [Big Sky :: golang と Generics と私](http://mattn.kaoriya.net/software/lang/go/20170309201506.htm)
    - [golang と Generics と吾 - Qiita](http://qiita.com/yuroyoro/items/6bf33f3cd4bb35469e0b)
    - [Java の Generics にもの思い - Qiita](http://qiita.com/t2y/items/139c6a38173d7750ddfc)

#### なぜ [Go 言語]はイケてないか？

- [なぜGo言語は設計が悪いのか – Go愛好者の見地から | 未分類 | POSTD](http://postd.cc/why-go-is-a-poorly-designed-language/)
- [Go言語がダメな理由 | プログラミング | POSTD](http://postd.cc/why-go-is-not-good/)
- [Why Everyone Hates Go · npf.io](Why Everyone Hates Go · npf.io)
    - [[翻訳]なんでGoってみんなに嫌われてるの？ - Qiita](http://qiita.com/hirokidaichi/items/adccebb41f77eaa6132f)

### 小手先のテクニックもしくはプログラミングパターン

- [Big Sky :: golang の遅いコードをたった1行で高速化するテクニック](http://mattn.kaoriya.net/software/lang/go/20160804131744.htm) : 実際にはちょっと速くなるくらいらしい
- [Go言語で作った実行ファイルを小さくしよう！ - Qiita](http://qiita.com/circus/items/450254c59d194cbf22d7)
- [Goのプログラミングパターン](http://www.infoq.com/jp/news/2016/03/go-patterns)
- [6年間におけるGoのベストプラクティス | プログラミング | POSTD](http://postd.cc/go-best-practices-2016/)
- [Go言語のFunctional Option Pattern - Qiita](http://qiita.com/weloan/items/56f1c7792088b5ede136)
- [Big Sky :: Golang で物理ファイルの操作に path/filepath でなく path を使うと爆発します。](https://mattn.kaoriya.net/software/lang/go/20171024130616.htm) : URL の操作には `http.ServeFile` を使うとかあるらしい
- [golang の 引数、戻り値、レシーバをポインタにすべきか、値にすべきかの判断基準について迷っている - pospomeのプログラミング日記](http://pospome.hatenablog.com/entry/2017/08/12/195032) : ケースバイケースだよねぇ。もしくはプロジェクトごとにポリシーを決めるか
- [Goを始めて1年間で最高にお世話になったGo関連ブックマークを晒します。 - Qiita](https://qiita.com/po3rin/items/0d8fef14bfe222f334b7)
- [Practical Go | Dave Cheney](https://dave.cheney.net/practical-go)

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
{{% review-paapi "4873117526" %}} <!-- Go言語によるWebアプリケーション開発 -->
