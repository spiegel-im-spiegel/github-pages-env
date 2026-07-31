+++
title = "Typst に関する雑多な話"
date =  "2025-03-02T21:55:38+09:00"
description = "ここでは Typst に関する小ネタをまとめて挙げておく（随時更新）"
image = "/images/attention/tools.png"
tags  = [ "typst", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ここでは [Typst] に関する小ネタをまとめて挙げておく（随時更新）。

## [Typst] をビルドする

[Typst] は Windows であれば Winget， macOS なら Homebrew， Linux なら Snap または “[Versions for typst](https://repology.org/project/typst/versions "typst package versions - Repology")” から直接取得できるが，これらの方法でインストールできない場合は [Rust] のビルド環境を導入して [Typst] をビルドする。

とりあえず Linux プラットフォームで [Rust] ビルド環境をインストールするには以下のコマンドでいける[^r1]。

[^r1]: Linux の [Rust] ビルド環境には GCC も必要。 Ubuntu は GCC が既定で入ってないので， `sudo apt build-essential` で GCC を導入する。

```text
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

これで PATH の設定までやってくれる。
その後，以下のコマンドで [Typst] のビルドを行う。

```text
$ cargo install --locked typst-cli
```

[Rust] ビルド環境を標準設定でインストールしているなら `~/.cargo/bin/` ディレクトリ（Linux の場合）に `typst` コマンドが出来ているはずである。
`cargo` コマンドが起動しているならこのディレクトリに PATH が通ってるはずなので，そのまま `typst` コマンドを実行できる。

## 変数をコマンドライン引数で指定する {#input}

[Typst] は `compile` 時に `--input` オプションでキーと値を与えることができる。
`--input` オプションは複数指定できる。

```text
$ typst compile --input key1=value1 --input key2=value2 sys-inputs.typ
```

設定したキー・値のペアは [`sys`]`.inputs` から取得することができる。

```typst {hl_lines=[12]}
#show raw: body => {
    set text(font: (
      (
        name: "Inconsolata",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}

#sys.inputs
```

このコードに対し `--input key1=value1 --input key2=value2` オプションを付けて `compile` すると。

{{< fig-img src="./sys-inputs-1.png" title="変数をコマンドライン引数で指定する" link="./sys-inputs-1.pdf" width="677" >}}

などと連想配列（[`dictionary`]）の形で格納されていることが分かる。
なので，上の例であれば `sys.inputs.key1` または `sys.inputs.at("key1")` で値 `"value1"` を取得できる（値は必ず文字列に解釈される）。
ただし `sys.inputs.key3` のように `--input` オプションで指定していないキーを読もうとすると

```text
$ typst compile --input key1=value1 --input key2=value2 sys-inputs-1b.typ 
error: dictionary does not contain key "key3"
   ┌─ sys-inputs-1b.typ:12:12
   │
12 │ #sys.inputs.key3
   │             ^^^^
```

という感じにコンパイルエラーになる。
厄介なことに [VS Code] の [Tinymist Typst] 拡張機能は，このような immediate なキーの記述に対してエラーを吐いてくれて，けっこう鬱陶しい。
回避策としては

```typst
#let key3 = ""

#if "key3" in sys.inputs {
	key3 = sys.inputs.at("key3")
}

#key3
```

などと記述すればいいようだ。
この場合 `compile` 処理自体は「正常終了」してしまうのがデメリットかな[^e1]。

[^e1]: ~~軽く調べてみたが [Typst] には `exit` や `panic` のようなプロセスを強制終了させる仕組みがない。 `try-catch` の例外処理もなく，当然 `throw` のようなものもない。たとえば処理中に（文法エラーや言語仕様上の致命的エラーではなく）ビジネスロジック上の問題があったときに，その問題をドキュメントとして出力することは可能だが，処理自体は「正常終了」してしまうため，プロセスを監視する側はエラーを感知できず出力結果を目視してはじめてエラーが起きていることが分かることになる。これはちょっと面白くない。~~【2025-03-07 追記】あれから少し勉強して [`assert`] 関数を使えばロジカルなエラーを検出したときに指定したメッセージを吐いてエラー終了できることが分かった。まじすんません {{% emoji "ペコン" %}} ただし，今回の例では [`assert`] 関数を使ってエラーハンドリングしようとしても，結局は [Tinymist Typst] 拡張機能がエラーを吐くので鬱陶しいことには変わりない。

## ブックマーク

ブックマークは「[Typst に関するブックマーク]({{< relref "./0-bookmark.md" >}})」にてまとめています。

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"
[`sys`]: https://typst.app/docs/reference/foundations/sys/ "System Functions – Typst Documentation"
[`dictionary`]: https://typst.app/docs/reference/foundations/dictionary/ "Dictionary Type – Typst Documentation"
[`assert`]: https://typst.app/docs/reference/foundations/assert/ "Assert Function – Typst Documentation"

[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Tinymist Typst]: https://marketplace.visualstudio.com/items?itemName=myriad-dreamin.tinymist "Tinymist Typst - Visual Studio Marketplace"
[Rust]: https://www.rust-lang.org/ "Rust Programming Language"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
