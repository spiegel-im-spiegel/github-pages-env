+++
title = "Go 1.27 のリリース"
date =  "2026-08-30T14:40:23+09:00"
description = "今はじっくり検証する暇がないので，とりあえずブックマークのみ載せておく。"
isCJKLanguage = true
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "package", "module", "generics", "concurrency", "cryptography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 [Go] 1.27 がリリースされた。

- [Go 1.27 is released - The Go Programming Language](https://go.dev/blog/go1.27)
- [Go 1.27 Release Notes - The Go Programming Language](https://go.dev/doc/go1.27)

個人的には [`encoding/json/v2`] が（実験的実装じゃなくて）正式にリリースされたことが注目点かな。
今はじっくり検証する暇がないので，とりあえずブックマークのみ載せておく。

- [Generic Methods - The Go Programming Language](https://go.dev/blog/generic-methods)
- [Goroutine Leak Profiles - The Go Programming Language](https://go.dev/blog/goroutine-leak-profiles)
- [Go 1.27 から uuid 実装がサポートされる！ので個人的に気になった議論とその着地をまとめてみた](https://zenn.dev/layerx/articles/f7124d4e761c1f)
- [Go 1.27 リリース連載：インデックス+HTTP/3(定期観察)+SIMD(第2弾) | フューチャー技術ブログ](https://future-architect.github.io/articles/20260728a/)
  - [Go 1.27のgo mod tidyの更新点 | フューチャー技術ブログ](https://future-architect.github.io/articles/20260729a/)
  - [Go 1.27 リリース連載：ジェネリクスメソッド (generic methods) | フューチャー技術ブログ](https://future-architect.github.io/articles/20260730a/)
  - [Go 1.27 リリース連載：encoding/json/v2 | フューチャー技術ブログ](https://future-architect.github.io/articles/20260731a/)
  - [Go 1.27リリース連載：goroutineleakプロファイルでリークを検出して修正する | フューチャー技術ブログ](https://future-architect.github.io/articles/20260803a/)
  - [Go 1.27 リリース連載： uuid | フューチャー技術ブログ](https://future-architect.github.io/articles/20260804a/)
  - [Go 1.27 の構造体リテラルのキー拡張で、埋め込みフィールドの初期化が楽になった | フューチャー技術ブログ](https://future-architect.github.io/articles/20260805a/)
  - [Go 1.27 リリース連載：ポスト量子暗号 | フューチャー技術ブログ](https://future-architect.github.io/articles/20260806a/)
  - [Go 1.27 の go fix アップデート | フューチャー技術ブログ](https://future-architect.github.io/articles/20260807a/)
- [Go 1.27 の標準パッケージに uuid が入った！のでいろいろ喋る / go_127_std_go_uuid](https://speakerdeck.com/convto/go-127-std-go-uuid)

`go fix` コマンドについては手持ちのパッケージで一度試したほうがいいかもなぁ。

[Go]: https://go.dev/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Scoop]: https://scoop.sh/
[`encoding/json/v2`]: https://pkg.go.dev/encoding/json/v2 "json package - encoding/json/v2 - Go Packages"
[`encoding/json`]: https://pkg.go.dev/encoding/json "json package - encoding/json - Go Packages"
[`os`]: https://pkg.go.dev/os "os package - os - Go Packages"

## 参考図書

{{< linkcard "9e8d8f717b1d1d85b0da7f07cb10a83b78d4227f" >}} <!-- プログラミング言語Go -->
{{< review-paapi "B0CFL1DK8Q" >}} <!-- Go言語 100Tips -->
{{% review-paapi "B0DNYMMBBQ" %}} <!-- Go言語で学ぶ並行プログラミング -->
