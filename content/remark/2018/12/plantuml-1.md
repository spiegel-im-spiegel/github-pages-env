+++
title = "真面目に PlantUML (1) : PlantUML のインストール"
date = "2018-12-27T19:55:14+09:00"
description = " せっかく時間ができたのだから PlantUML についてちゃんと調べてみることにした。"
image = "/images/attention/kitten.jpg"
tags = [ "java", "plantuml", "uml", "tools" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

1年ほど前に [ATOM] エディタを使って [PlantUML] を使って UML で遊んだのだが， [PlantUML] 自体は「とりあえず動かしてみた」というレベルできちんと調べていなかった。
せっかく時間ができたのだから，この機会にちゃんと調べてみることにした。

## 目次

1. [PlantUML のインストール]({{< ref "/remark/2018/12/plantuml-1.md" >}}) ← イマココ
1. [シーケンス図]({{< ref "/remark/2018/12/plantuml-2-sequence-diagram.md" >}})
1. [クラス図]({{< ref "/remark/2018/12/plantuml-3-class-diagrams.md" >}})
1. [実体関連図]({{< ref "/remark/2019/07/plantuml-4-entity-relationship-diagrams.md" >}})

## [PlantUML] のインストール

[PlantUML] は UML (Unified Modeling Language; 統一モデリング言語) を記述するための DSL (Domain-Specific Language; ドメイン特化言語) と考えてよい。
UML の構文（に相当するもの）は図形で表されるが，それをテキストで記述できるようにしたのが [PlantUML] である。
[PlantUML] では UML 2.0 のうち以下の図（Diagrams）を作成できるようだ。

- 構造図（Structural Diagrams）
    - [クラス図（Class Diagram）](http://plantuml.com/class-diagram) ※同じ機能でパッケージ図も作成可能
    - [オブジェクト図（Object Diagram）](http://plantuml.com/object-diagram)
    - [コンポーネント図（Component Diagram）](http://plantuml.com/component-diagram)
    - [配置図（Deployment Diagram）](http://plantuml.com/deployment-diagram)
- 振る舞い図（Behavioral Diagrams）
    - [ユースケース図（Use Case Diagram）](http://plantuml.com/use-case-diagram)
    - [アクティビティ図（Activity Diagram）](http://plantuml.com/activity-diagram-beta) ※ベータ版
    - [状態遷移図（State Machine Diagram）](http://plantuml.com/state-diagram)
    - [シーケンス図（Sequence Diagram）](http://plantuml.com/sequence-diagram)
    - [タイミング図（Timing Diagram）](http://plantuml.com/timing-diagram)

[PlantUML] の実行モジュールは Java の jar ファイルで提供されている。
[ダウンロードページ](http://plantuml.com/download)から最新の `plantuml.jar` ファイルをダウンロードしてそのまま使える。

[PlantUML] を実行するには Java のランタイム環境を用意する必要がある。
最新の [OpenJDK](http://openjdk.java.net/) を使って問題ない。
[Ubuntu] であれば [APT] で Java ランタイム環境をインストールできる。
たとえば LTS バージョンの Java 11 であれば

```text
$ sudo apt install openjdk-11-jre
```

でOK。

[シーケンス図](http://plantuml.com/sequence-diagram)と[アクティビティ図](http://plantuml.com/activity-diagram-beta)以外は [DOT 言語]にトランス・コンパイルされるため，別途 [Graphviz] もインストールする必要がある。
[Graphviz] は[ダウンロードページ](http://graphviz.org/download/)等から取得できる。
[Ubuntu] であれば [APT] でインストールできる。

```text
$ sudo apt install graphviz
```

全ての用意ができたら `plantuml.jar` ファイルのあるフォルダで以下のコマンドを入力し動作確認しておこう（以下は [Ubuntu] 環境下での例）。

```text
$ java -jar plantuml.jar -version
PlantUML version 1.2019.11 (Sun Sep 22 19:02:15 JST 2019)
(GPL source distribution)
Java Runtime: OpenJDK Runtime Environment
JVM: OpenJDK 64-Bit Server VM
Java Version: 13.0.1+9
Operating System: Linux
Default Encoding: UTF-8
Language: ja
Country: JP
...
```

[Graphviz] との連携をテストするには以下のコマンドを実行する。

```text
$ java -jar plantuml.jar -testdot
The environment variable GRAPHVIZ_DOT has not been set
Dot executable is /usr/bin/dot
Dot version: dot - graphviz version 2.40.1 (20161225.0304)
Installation seems OK. File generation OK
```

## コマンドライン・オプション

[PlantUML] のコマンドライン・オプションは以下の通り。

```text
$ java -jar plantuml.jar -h
Usage: java -jar plantuml.jar [options] -gui
        (to execute the GUI)
    or java -jar plantuml.jar [options] [file/dir] [file/dir] [file/dir]
        (to process files or directories)

You can use the following wildcards in files/dirs:
        *       means any characters but '/'
        ?       one and only one character but '/'
        **      means any characters (used to recurse through directories)

where options include:
    -gui                To run the graphical user interface
    -tpng               To generate images using PNG format (default)
    -tsvg               To generate images using SVG format
    -teps               To generate images using EPS format
    -tpdf               To generate images using PDF format
    -tvdx               To generate images using VDX format
    -txmi               To generate XMI file for class diagram
    -tscxml             To generate SCXML file for state diagram
    -thtml              To generate HTML file for class diagram
    -ttxt               To generate images with ASCII art
    -tutxt              To generate images with ASCII art using Unicode characters
    -tlatex             To generate images using LaTeX/Tikz format
    -tlatex:nopreamble  To generate images using LaTeX/Tikz format without preamble
    -o[utput] "dir"     To generate images in the specified directory
    -DVAR1=value        To set a preprocessing variable as if '!define VAR1 value' were used
    -Sparam1=value      To set a skin parameter as if 'skinparam param1 value' were used
    -r[ecurse]          recurse through directories
    -I/path/to/file     To include file as if '!include file' were used
    -I/path/to/*.puml   To include files with pattern
    -charset xxx        To use a specific charset (default is UTF-8)
    -e[x]clude pattern  To exclude files that match the provided pattern
    -metadata           To retrieve PlantUML sources from PNG images
    -nometadata         To NOT export metadata in PNG/SVG generated files
    -checkmetadata      Skip PNG files that don't need to be regenerated
    -version            To display information about PlantUML and Java versions
    -checkversion       To check if a newer version is available for download
    -v[erbose]          To have log information
    -quiet              To NOT print error message into the console
    -debugsvek          To generate intermediate svek files
    -h[elp]             To display this help message
    -testdot            To test the installation of graphviz
    -graphvizdot "exe"  To specify dot executable
    -p[ipe]             To use stdin for PlantUML source and stdout for PNG/SVG/EPS generation
    -encodesprite 4|8|16[z] "file"      To encode a sprite at gray level (z for compression) from an image
    -computeurl|-encodeurl      To compute the encoded URL of a PlantUML source file
    -decodeurl          To retrieve the PlantUML source from an encoded URL
    -syntax             To report any syntax error from standard input without generating images
    -language           To print the list of PlantUML keywords
    -checkonly          To check the syntax of files without generating images
    -failfast           To stop processing as soon as a syntax error in diagram occurs
    -failfast2          To do a first syntax check before processing files, to fail even faster
    -pattern            To print the list of Regular Expression used by PlantUML
    -duration           To print the duration of complete diagrams processing
    -nbthread N         To use (N) threads for processing
    -nbthread auto      To use 4 threads for processing
    -timeout N          Processing timeout in (N) seconds. Defaults to 15 minutes (900 seconds).
    -author[s]          To print information about PlantUML authors
    -overwrite          To allow to overwrite read only files
    -printfonts         To print fonts available on your system
    -enablestats        To enable statistics computation
    -disablestats       To disable statistics computation (default)
    -htmlstats          To output general statistics in file plantuml-stats.html
    -xmlstats           To output general statistics in file plantuml-stats.xml
    -realtimestats      To generate statistics on the fly rather than at the end
    -loopstats          To continuously print statistics about usage
    -splash             To display a splash screen with some progress bar
    -progress           To display a textual progress bar in console
    -pipeimageindex N   To generate the Nth image with pipe option
    -stdlib             To print standard library info
    -extractstdlib      To extract PlantUML Standard Library into stdlib folder
    -filename "example.puml"    To override %filename% variable
    -preproc            To output preprocessor text of diagrams
    -cypher             To cypher texts of diagrams so that you can share them

If needed, you can setup the environment variable GRAPHVIZ_DOT.
```

かなりの数のオプションがあるが，よく使うのは以下のものだろうか。

| オプション           | 内容                                      |
| -------------------- | ----------------------------------------- |
| `-tpng`              | PNG 形式で出力（既定値）                  |
| `-tsvg`              | SVG 形式で出力                            |
| `-teps`              | EPS 形式で出力                            |
| `-tpdf`              | PDF 形式で出力                            |
| `-tlatex`            | LaTeX/Tikz 形式で出力                     |
| `-tlatex:nopreamble` | LaTeX/Tikz 形式で出力（プリアンブルなし） |
| `-o[utput] "dir"`    | 指定したフォルダに出力                    |
| `-charset xxx`       | 文字エンコーディングの指定                |
| `-nometadata`        | PNG/SVG 出力にメタデータを含めない        |
| `-p[ipe]`            | 入出力をパイプで行う                      |

## みんな大好き Hello World

では簡単なコードを処理してみよう。

`hello.puml` という名前のファイルを用意して以下のように入力する。

```puml
@startuml

Alice->Bob : hello
Alice<-Bob : はろー

@enduml
```

[PlantUML] ではドキュメントの中にコードを埋め込むことを想定して `@startuml ... @enduml` 内の記述のみが処理の対象となる。

ではこのファイルを処理してみよう。

```text
$ java -jar plantuml.jar -charset UTF-8 hello.puml
```

これで `hello.puml` と同じフォルダに `hello.png` ファイルが出力される。
`hello.png` ファイルの内容は以下の通り。

{{< fig-img src="./hello.png" link="./hello.puml" >}}

って，ちっちゃ！
じゃあ，解像度を調節してもう少し大きくしてみよう。

{{< highlight puml "hl_lines=3" >}}
@startuml

skinparam dpi 300

Alice->Bob : hello
Alice<-Bob : はろー

@enduml
{{< /highlight >}}

これを処理するとこうなる。

{{< fig-img src="./hello2.png" link="./hello2.puml" >}}

おおっ。
おっきくなった。

## 見た目の調整

上の例で示したように [PlantUML] では `skinparam` コマンドを使って見た目の調整ができる。
では，もう少し調整を入れて UML っぽい図にしてみる。

{{< highlight puml "hl_lines=3-32" >}}
@startuml

skinparam shadowing false
skinparam backgroundColor transparent
skinparam dpi 300

skinparam arrow {
  FontColor Black
  FontName Noto Serif JP
  FontSize 10
  FontStyle plain
  Color Black
  Thickness 1
}
skinparam participant {
  FontColor Black
  FontName Noto Sans
  FontSize 14
  FontStyle plain
  BackgroundColor WhiteSmoke
  BorderColor Black
  BorderThickness 1
  Padding 20
}
skinparam sequence {
  LifeLineBackgroundColor transparent
  LifeLineBorderColor Black
  LifeLineBorderThickness 1
  MessageAlignment center
}
hide footbox
skinparam BoxPadding 10

Alice->Bob : hello
Alice<-Bob : はろー

@enduml
{{< /highlight >}}

フォントの変更と線の色と太さを調節してみた。
またライフラインの間隔も調整している。
以上を処理するとこんな図になる。

{{< fig-img src="./hello3.png" link="./hello3.puml" >}}

## 見た目の指定を別ファイルに分離する

見た目の指定は別ファイルにして再利用したいところである。
そこで `hello.iuml` ファイルを作成して `skinparam` コマンドの内容をまるっと移動してみる。

```puml
skinparam shadowing false
skinparam backgroundColor transparent
skinparam dpi 300

skinparam arrow {
  FontColor Black
  FontName Noto Serif JP
  FontSize 10
  FontStyle plain
  Color Black
  Thickness 1
}
skinparam participant {
  FontColor Black
  FontName Noto Sans
  FontSize 14
  FontStyle plain
  BackgroundColor WhiteSmoke
  BorderColor Black
  BorderThickness 1
  Padding 20
}
skinparam sequence {
  LifeLineBackgroundColor transparent
  LifeLineBorderColor Black
  LifeLineBorderThickness 1
  MessageAlignment center
}
hide footbox
skinparam BoxPadding 10
```

このファイルをインクルードする形でコードを修正する。

{{< highlight puml "hl_lines=3" >}}
@startuml

!include hello.iuml

Alice->Bob : hello
Alice<-Bob : はろー

@enduml
{{< /highlight >}}

これを処理すると

{{< fig-img src="./hello4.png" link="./hello4.puml" >}}

先程と全く同じ結果が得られる。

また最初の

```puml
@startuml

Alice->Bob : hello
Alice<-Bob : はろー

@enduml
```

に対して

```text
$ java -jar plantuml.jar -charset UTF-8 -config hello.iuml hello.puml
```

と `-config` オプションで `hello.iuml` ファイルを指定しても同じ結果が得られる。
これで再利用しやすくなっただろう。

## 【追記】 [PlantUML] のコメント

[PlantUML] の `@startuml ... @enduml` 領域内にコメントを記述する際には以下の２通りの書き方がある。

```puml
@startuml

' 一行コメント

/'
複数行に渡る
コメント
'/

@enduml
```

なお，行の途中からコメントを挿入する

```puml
@startuml

skinparam dpi 300 '解像度の設定

@enduml
```

のような記述はできない（しかも行がまるごと無視されるだけでエラーにもならない）のでご注意を。

## ブックマーク

- [UML入門 - IT専科](http://www.itsenka.com/contents/development/uml/)
- [Changing colors and fonts](http://plantuml.com/skinparam)
- [call it from your script using command line](http://plantuml.com/command-line)
- [PlantUML の使い方 | プログラマーズ雑記帳](http://yohshiy.blog.fc2.com/blog-entry-152.html)

[ATOM]: https://atom.io/ "Atom"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."
[Graphviz]: http://graphviz.org/ "Graphviz - Graph Visualization Software"
[DOT 言語]: https://graphviz.gitlab.io/_pages/doc/info/lang.html "The DOT Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[APT]: https://debian-handbook.info/browse/ja-JP/stable/apt.html "第 6 章 メンテナンスと更新、APT ツール"
