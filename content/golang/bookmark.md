+++
title = "Go 言語に関するブックマーク"
date = "2015-09-11T17:58:42+09:00"
description = "主に公式情報や書籍情報について纏めていくことにした。"
tags = ["golang", "bookmark"]
image = "/images/attention/go-logo_blue.png"

[scripts]
  mathjax = false
  mermaidjs = false
+++

主に公式情報や書籍情報について纏めていくことにした。
それ以外の情報は [bookmarks]({{< rlnk "/bookmarks/" >}}) セクションで扱っている。
なお，過去分は[こちらに退避]({{< ref "/bookmarks/golang.md" >}})させている。

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

### Go 1.15 is released

- [Proposals for Go 1.15 - The Go Blog](https://blog.golang.org/go1.15-proposals)
- [Go 1.15 Release Notes - The Go Programming Language](https://golang.org/doc/go1.15)
    - [Go 1.15 リリースノート 日本語訳 - Qiita](https://qiita.com/c-yan/items/dad49c9dce27e77a94e7)
- [Dive deep into Go 1.15 changes to testing package - Speaker Deck](https://speakerdeck.com/hgsgtk/dive-deep-into-go-1-dot-15-changes-to-testing-package)
- [Go1.15 からのreflectパッケージの挙動の違いについて - 一寸先は/dev/null](https://nokute.hatenablog.com/entry/2020/08/14/112132)

### Go 1.16 is released

- [Go 1.16 Release Candidate 1 is released](https://groups.google.com/g/golang-announce/c/U_FUHY4wuSc/m/3_Vw3oqpAgAJ)
- [Go 1.16 Release Notes - The Go Programming Language](https://tip.golang.org/doc/go1.16) : ドラフト版
    - [Go 1.16 ドラフトリリースノート 日本語訳  - Qiita](https://qiita.com/c-yan/items/f6f504d5e1c1caceace4)
- [Go on ARM and Beyond - The Go Blog](https://blog.golang.org/ports)
- [Working with Embed in Go 1.16 Version](https://lakefs.io/working-with-embed-in-go/)
- [Go1.16からの go get と go install について - Qiita](https://qiita.com/eihigh/items/9fe52804610a8c4b7e41)
- [go:embed 詳解 - 使用編 -](https://zenn.dev/koya_iwamura/articles/53a4469271022e)

### Go 1.17 is released

- [Go 1.17 Release Notes - The Go Programming Language](https://golang.org/doc/go1.17)
- [Go 1.17連載が始まります: コンパイラとgo mod | フューチャー技術ブログ](https://future-architect.github.io/articles/20210810a/)
  - [Go1.17のencoding/csv | フューチャー技術ブログ](https://future-architect.github.io/articles/20210811a/)
  - [Go 1.17のtesting新機能 | フューチャー技術ブログ](https://future-architect.github.io/articles/20210812a/)
  - [Go 1.17からの負のruneの扱い | フューチャー技術ブログ](https://future-architect.github.io/articles/20210817a/)
  - [Go1.17で警告されるようになったerror#Is/As/Unwrap | フューチャー技術ブログ](https://future-architect.github.io/articles/20210819b/)
  - [Go 1.17の sync/atomic パッケージ更新点と CompareAndSwap | フューチャー技術ブログ](https://future-architect.github.io/articles/20210820a/)
- [Go 1.17 Release祝い 細かすぎて伝わらないMinor change in flag package](https://zenn.dev/hgsgtk/articles/9f662a4c96fa3f)
- [Big Sky :: unsafe.Add と unsafe.Slice が入った。](https://mattn.kaoriya.net/software/lang/go/20210504002438.htm)
- [Go 1.17 リリースパーティ - connpass](https://gocon.connpass.com/event/216361/)
  - [GoのGenerics関連プロポーザル最新状況まとめと簡単な解説 (2021年8月版)](https://zenn.dev/syumai/articles/c42hdg1e0085btnen5hg)
  - [初めての型セット - Speaker Deck](https://speakerdeck.com/nobishino/introduction-to-type-sets)

### Go 1.18 RC

- [Go 1.18 Release Notes - go.dev](https://tip.golang.org/doc/go1.18)
  - [Go 1.18 Beta 1 is available, with generics - go.dev](https://go.dev/blog/go1.18beta1)
- [Big Sky :: Go の http パッケージに MaxBytesHandler が入った。](https://mattn.kaoriya.net/software/lang/go/20211224005655.htm)
- [go1.18で入ったhttp.MaxBytesHandlerの中身を見てみた](https://zenn.dev/hiroyukim/articles/4b4f5b482c0c2d)
- [Go1.18から導入されるnetip package/netip-package - Speaker Deck](https://speakerdeck.com/sonatard/netip-package)
- [What's new in Go 1.18? - Speaker Deck](https://speakerdeck.com/syumai/whats-new-in-go-1-dot-18)
- [What Is the Go Workspace Mode - Speaker Deck](https://speakerdeck.com/110y/what-is-the-go-workspace-mode)
- [[shared] 20220218 Go 1.18 Fuzzing - Google スライド](https://docs.google.com/presentation/d/e/2PACX-1vQ2PX-s-As01o_fvGi9qdx9ZCpQS6RePDWw6rkznN-lI3z8bc4JJ601HLzb1fujo4uf0wSK0Wzl_oc1/pub?resourcekey=0-R72bI85HvGnbMfYmAP_77g#slide=id.g405a9dc47b_0_0)
- [Go 1.18で追加されるstrings/bytes.Cutとsync.Mutex.TryLockについて - Google スライド](https://docs.google.com/presentation/d/1iaEMhXHQa5chIK7Zqqcv6sugXoOEYDQnvldlZlxhJjw/edit#slide=id.p)
- [Go言語のジェネリクス入門(1)](https://zenn.dev/nobishii/articles/type_param_intro)
  - [Go言語のジェネリクス入門(2) インスタンス化と型推論](https://zenn.dev/nobishii/articles/type_param_intro_2)

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
- [Draft designを読む · GitHub](https://gist.github.com/tenntenn/fe8995c347a5e1000832d3c6942f1fbe)
- [A Proposal for Adding Generics to Go - The Go Blog](https://blog.golang.org/generics-proposal)

### Go Cloud Development Kit

- [What's new in the Go Cloud Development Kit - The Go Blog](https://blog.golang.org/gcdk-whats-new-in-march-2019)
- [The New Go Developer Network - The Go Blog](https://blog.golang.org/go-developer-network)
- [Pkg.go.dev is open source! - The Go Blog](https://blog.golang.org/pkgsite)
- [Redirecting godoc.org requests to pkg.go.dev - The Go Blog](https://blog.golang.org/godoc.org-redirect)

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4908686033" %}} <!-- Goならわかるシステムプログラミング -->
{{% review-paapi "B07VPSXF6N" %}} <!-- 改訂2版 みんなのGo言語 -->
{{% review-paapi "4873117526" %}} <!-- Go言語によるWebアプリケーション開発 -->
