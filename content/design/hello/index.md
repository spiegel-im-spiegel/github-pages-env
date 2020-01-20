+++
title = "まずは，みんな大好き Hello World"
date =  "2020-01-20T21:09:37+09:00"
description = "description"
image = "/images/attention/design.png"
tags = [ "programming", "design" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

まずは，みんな大好き “Hello World” から。

“Hello World” はデザイン・パターンではないが，開発環境を構築した際の動作確認の例として今でもよく使われる。
なにせ [“Hello World” を集めた GitHub リポジトリ](https://github.com/leachim6/hello-world "leachim6/hello-world: Hello world in every computer language. Thanks to everyone who contributes to this, make sure to see CONTRIBUTING.md for contribution instructions!")まであるくらいなのだ。

このセクションでは以下の環境を前提に記事を書いていく。

- Linux/[Ubuntu] プラットフォーム
- Java は [OpenJDK] の（記事を書いた時点での）最新バージョン
- [Go] は（記事を書いた時点での）最新バージョン

では，まずは Java から。

## Java{#java}

最近の Java はいろいろ記述を端折れるようになっていて，標準ライブラリを使うだけなら

```java
public class Hello {
	public static void main(String[] args) {
		System.out.println("Hello World");
	}
}
```

の記述のみで OK (クラス名は任意)。

実行についても1ファイルのみなら

```text
$ java Hello.java
Hello World
```

でコンパイル&実行してくれる。

うむうむ。

## [Go]

[Go] の “Hello World” はこんな感じ。

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

実行については

```text
$ go run hello.go 
Hello World
```

とすればコンパイル&実行してくれる。

よーし，うむうむ，よーし。

## ブックマーク

- [結局 OpenJDK をインストールし直すことにした]({{< ref "/remark/2019/07/reinstalling-openjdk.md" >}})

[OpenJDK]: http://openjdk.java.net/
[Go]: https://golang.org/ "The Go Programming Language"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
