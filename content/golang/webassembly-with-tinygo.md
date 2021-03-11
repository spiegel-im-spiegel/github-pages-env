+++
title = "WebAssembly で Hello, World (ブログ版)"
date =  "2021-03-10T22:02:16+09:00"
description = "description"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "tinygo", "webassembly" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

[TinyGo] は本家 [Go] のサブセットと言えるもので [LLVM] を使った組み込み用途のコンパイラである。
しかも [LLVM] が [WebAssembly] バイナリを直接出力できるということもあって [TinyGo] と [WebAssembly] の相性はとてもよい。

というわけで今回は， [Go] および [TinyGo] を使って [WebAssembly] へのコンパイルを行い， Web ブラウザ上で動作させるところまでやってみることにする。














[Go]: https://golang.org/ "The Go Programming Language"
[TinyGo]: https://tinygo.org/ "TinyGo - Go on Microcontrollers and WASM"
[WebAssembly]: https://webassembly.org/ "WebAssembly"
[LLVM]: https://llvm.org/ "The LLVM Compiler Infrastructure Project"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
