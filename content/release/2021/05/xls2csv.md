+++
title = "Excel → CSV 変換ツールを作ってみた"
date =  "2021-05-17T19:16:06+09:00"
description = "CLI ツールで，機能としては Excel データを CSV 形式に変換するだけの簡単なお仕事。 "
image = "/images/attention/tools.png"
tags  = [ "tools", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先日 Zenn に

- [Go で簡単 Excel → CSV 変換](https://zenn.dev/spiegel/articles/20210516-excel-to-csv)

という記事を書いたが，同等の機能の汎用ツールを作ってみた。

- [spiegel-im-spiegel/xls2csv: Export CSV Text from Excel Data](https://github.com/spiegel-im-spiegel/xls2csv)

CLI ツールで，機能としては Excel データの行・列の情報を単純に読み込んで CSV 形式に変換するだけの簡単なお仕事。
Pure Go で書かれているので Excel がない環境（たとえば Linux）でも問題なく動く。

```text
$ xls2csv -h
Export CSV text from Excel data.

Usage:
  xls2csv [flags] <Excel file>

Flags:
      --debug             for debug
  -h, --help              help for xls2csv
  -o, --output string     path of output CSV file
  -p, --password string   password in Excel file
  -s, --sheet string      sheet name in Excel file
  -t, --tsv               output with TSV format
  -v, --version           output version of xls2csv
  -w, --win-newline       output with CRLF newline

$ xls2csv conv/testdata/test-pw.xlsx -p passwd
名前,年齢
Alice,18
Bob,19
太郎,20
花子,21
```

このようにパスワードロックされている Excel ファイルも読み込める[^pw1]。

[^pw1]: Shell のプロンプトやバッチファイルに直接パスワードを記述するのはオススメできないが， Excel のパスワード・ロックなんてザルだからなぁ

なお [Excelize] パッケージを使っているので Excel 2007 までの古い形式（拡張子が `.xls` のファイル）には対応していない。
あしからず。

文字エンコードは UTF-8 で改行コードは LF の CSV 形式のテキストを標準出力に返す。
先頭の BOM は付かない（付ける気もない）。
なお改行コードは `-w` オプションで CRLF に変更可能である。
文字エンコードを変える機能はないので， Shift-JIS とかにする必要があるなら拙作の [gnkf]({{< ref "/release/gnkf.md" >}} "GNKF: Network Kanji Filter by Golang
") との組み合わせでどうぞ。

機能がニッチ過ぎるので需要はないだろうけど，まっ，自分用ということで（笑）

[Excelize] パッケージのドキュメントって日本語版もあるんだねぇ。

- [Excelize ドキュメンテーション](https://xuri.me/excelize/ja/)

いろいろと遊べそうである。

[Excelize]: https://xuri.me/excelize/ "Excelize Official Docs"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
