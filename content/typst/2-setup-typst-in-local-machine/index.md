+++
title = "ローカルで Typst 環境を整える"
date =  "2025-02-21T18:58:37+09:00"
description = "ローカルマシンへの環境構築と日本語フォントの指定について試してみた。"
image = "/images/attention/tools.png"
tags  = [ "typst", "font" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

さっそく[チュートリアル][Tutorial]を見ながら [Typst] を触っていこうと思ったのだが，[ちょっと動かしてみた]({{< relref "/typst/1-about-typst.md#ope" >}})ところ，クラウドサービスの環境では日本語がショボすぎて先へ進める気にならないので，ローカルに環境を作ってしまうことにした。

{{< div-box type="markdown" >}}
**【2025-02-22 追記】**

[前回]の記事に追記して，クラウドサービスの環境で日本語フォントが指定できることを確認した。
改めてクラウドサービスの環境で使えるフォントをよく見てみたが [TeX Live] に収録されている TTF/OTF フォントは全て使えるようだ。
[原ノ味フォント]も使える。

OS にバンドルされているフォントは OS の種類やバージョンによって異なるため，複数人で異なるプラットフォームでドキュメントを作成する場合はフォントを合わせるのに苦労する。
書籍出版とかではなく文字のデザインそのものに拘りがないのなら，クラウドサービスを使った協働作業にするとか Docker 等の仮想環境を使うなど，環境の方を合わせてしまったほうがいいかもしれない。
今回は私の個人的な作業なので好き勝手絶頂にやるけど（笑）

[前回]: {{< relref "./1-about-typst.md" >}} "Typst について"
[TeX Live]: http://www.tug.org/texlive/ "TeX Live - TeX Users Group"
[原ノ味フォント]: https://github.com/trueroad/HaranoAjiFonts "trueroad/HaranoAjiFonts: 原ノ味フォント"
{{< /div-box >}}

今回の環境は以下の通り：

- OS は [Ubuntu] 24.10
- [Typst] は Snap からインストール
- [VS Code] 用の拡張機能 [Tinymist Typst] を導入する[^t1]

[^t1]: ググってみると [VS Code] 用の拡張機能として Typst LSP を紹介しているところが多かったが，2025-02-21 現在では Typst LSP はなくなっていて [Tinymist Typst] に統合されているらしい。

動作確認しておこう。

```text
$ typst --version
typst 0.13.0 (8dce676d)

$ typst --help
Typst 0.13.0 (8dce676d)

Usage: typst [OPTIONS] <COMMAND>

Commands:
  compile  Compiles an input file into a supported output format [aliases: c]
  watch    Watches an input file and recompiles on changes [aliases: w]
  init     Initializes a new project from a template
  query    Processes an input file to extract provided metadata
  fonts    Lists all discovered fonts in system and custom font paths
  help     Print this message or the help of the given subcommand(s)

Options:
      --color <COLOR>  Whether to use color. When set to `auto` if the terminal
                       to supports it [default: auto] [possible values: auto,
                       always, never]
      --cert <CERT>    Path to a custom CA certificate to use when making
                       network requests [env: TYPST_CERT=]
  -h, --help           Print help
  -V, --version        Print version

Resources:
  Tutorial:                 https://typst.app/docs/tutorial/
  Reference documentation:  https://typst.app/docs/reference/
  Templates & Packages:     https://typst.app/universe/
  Forum for questions:      https://forum.typst.app/
```

おっ。
`typst fonts` コマンドで利用可能なフォントの一覧が出るのか。
試してみよう。

```text
$ typst fonts | grep Computer
New Computer Modern
New Computer Modern Math

$ typst fonts | grep "CJK JP"
Noto Sans CJK JP
Noto Sans Mono CJK JP
Noto Serif CJK JP

$ typst fonts | grep Inconsolata
Inconsolata
```

ふむむ。
[Ubuntu] に最初から入ってるフォントは使えるみたいだな。
[Inconsolata] フォントはコードを表示するために後から入れたフォント。
これもちゃんと使えるようだ。
New Computer Modern フォントは OS には入ってない筈だけど (`fc-list` コマンドで確認した) [Typst] 側で持ってるフォントなのかな？

この状態で簡単な文章を組版してみよう。
入力文書は前回と同じ。

```typst
= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

これを `sample-1.typ` ファイルに保存して以下のように `typst compile` する[^v1]。

[^v1]: [VS Code] 用の拡張機能 [Tinymist Typst] を導入している場合はコマンドパレットから “Typst: Exports the Opened File as PDF” を起動することで `typst compile` できる。

```text
$ typst compile sample-1.typ
```

組版結果は以下の通り：

{{< fig-img src="./pdf-sample-1.png" title="組版結果 (1)" link="./sample-1.pdf" width="871" >}}

フォントは以下のようになっていた。

{{< fig-img src="./property-sample-1.png" title="組版結果 (1)" link="./property-sample-1.pdf" >}}

[Firge](https://github.com/yuru7/Firge "yuru7/Firge: Fira Mono と源真ゴシックを合成したプログラミングフォント Firge (ファージ)")Nerd ？ そんなもん入れたっけ？ ~~いやこれ OS に最初っから入ってたやつだな。
てか，入ってるのか。
凄いな [Ubuntu]。~~
勘違い。
このフォントは自分で[入れてた]({{< ref "/remark/2021/09/lualatex-with-firge-font.md" >}} "Firge フォントを使って LuaLaTeX でコードを書く")わ。
よく分からないが，フォントリストから日本語コードが入ってるフォントを（アルファベット順で）探して割り当てたって感じ？

[Typst] では `#set` キーワードとそれに続く関数実行により文書スタイルを設定できる。
フォントの場合は以下のように設定する。

```typst {hl_lines=[1]}
#set text(font: "Noto Serif CJK JP")

= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

この場合の組版結果は以下の通り：

{{< fig-img src="./pdf-sample-2.png" title="組版結果 (2)" link="./sample-2.pdf" width="871" >}}

{{< fig-img src="./property-sample-2.png" title="組版結果 (2)" link="./property-sample-2.pdf" >}}

[Typst] 0.13 ではラテン文字を別のフォントに割り当てられるようになった。
こんな感じに指定するらしい。

```typst {hl_lines=["1-8"]}
// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "Noto Serif CJK JP"
))

= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

組版結果は以下の通り：

{{< fig-img src="./pdf-sample-3.png" title="組版結果 (3)" link="./sample-3.pdf" width="871" >}}

{{< fig-img src="./property-sample-3.png" title="組版結果 (2)" link="./property-sample-3.pdf" >}}

パッと見では分かりにくいだろうけど，数字部分が New Computer Modern フォントに変わっている。

見出しの書体が明朝体のままなのは面白くないので，これをゴシック体に変えたい。
色々と調べてみると以下のようにするのがいいらしい。

```typst {hl_lines=["10-21"]}
// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "Noto Serif CJK JP"
))

// for headings
#let heading_font(body) = { // heading_font 関数を定義
    set text(font: (
      (
        name: "Noto Sans",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
    body
}
#show heading: heading_font  // heading_font を適用する

= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

まず `#let` キーワードを使って `body` のフォントを変更する `heading_font` 関数を定義し，次に `#show` キーワードを使って見出し（`heading`）要素の表示時の処理として `heading_font` 関数を適用する，という流れである。
見出し要素に直接フォントを指定できれば簡単だったろうに，それはできないみたい。

組版結果は以下の通り：

{{< fig-img src="./pdf-sample-4.png" title="組版結果 (4)" link="./sample-4.pdf" width="871" >}}

{{< fig-img src="./property-sample-4.png" title="組版結果 (4)" link="./property-sample-4.pdf" >}}

よしよし。

別の書き方として

```typst {hl_lines=["10-20"]}
// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "Noto Serif CJK JP"
))

// for headings
#show heading: it => {
    set text(font: (
      (
        name: "Noto Sans",
        covers: "latin-in-cjk",
      ),
      "Noto Sans CJK JP"
    ))
	it.body
}

= アルベルト・アインシュタインについて

アルベルト・アインシュタインは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

という書き方もある。
やってることは同じ。
フォントを変更する処理を無名関数として定義しているだけである。

ところで，最近は日本語フォントとしてモリサワ BIZ UD 明朝/ゴシックがお気に入りである。
このブログの Web フォントとしても使っている。

- [モリサワ BIZ UD 明朝/ゴシック Web フォント]({{< ref "/remark/2022/09/morisawa-biz-ud-fonts.md" >}})

[Typst] でもこれを日本語フォントとして使いたい。
また見出しで使う欧文フォントには Helvetica を使いたい。
モリサワ BIZ UD フォントは [Ubuntu] にはないため，手動でインストールする。

- [Morisawa BIZ UDMincho](https://github.com/googlefonts/morisawa-biz-ud-mincho)
- [Morisawa BIZ UDGothic](https://github.com/googlefonts/morisawa-biz-ud-gothic)

ダウンロードした TTF ファイルを `/usr/local/share/fonts/` ディレクトリに入れる。
`fc-cache -fv` コマンドでフォントキャッシュを更新するのを忘れずに。

```text
$ fc-cache -fv

$ fc-list | grep "BIZ UD"
/usr/local/share/fonts/BIZUDMincho-Bold.ttf: BIZ UD明朝,BIZ UDMincho:style=Bold
/usr/local/share/fonts/BIZUDPMincho-Bold.ttf: BIZ UDP明朝,BIZ UDPMincho:style=Bold
/usr/local/share/fonts/BIZUDPGothic-Regular.ttf: BIZ UDPゴシック,BIZ UDPGothic:style=Regular
/usr/local/share/fonts/BIZUDPGothic-Bold.ttf: BIZ UDPゴシック,BIZ UDPGothic:style=Bold
/usr/local/share/fonts/BIZUDPMincho-Regular.ttf: BIZ UDP明朝,BIZ UDPMincho:style=Regular
/usr/local/share/fonts/BIZUDMincho-Regular.ttf: BIZ UD明朝,BIZ UDMincho:style=Regular
/usr/local/share/fonts/BIZUDGothic-Bold.ttf: BIZ UDゴシック,BIZ UDGothic:style=Bold
/usr/local/share/fonts/BIZUDGothic-Regular.ttf: BIZ UDゴシック,BIZ UDGothic:style=Regular
```

[Ubuntu] では Helvetica の代替として Liberation Sans が最初から入っているのでこれを使う[^h1]。

[^h1]: Windows では Helvetica の代替として Arial フォントがよく使われる。

これでフォント指定の最終形を以下とする。

```typst {hl_lines=[7,14,17,22,24]}
// for main text
#set text(font: (
  (
    name: "New Computer Modern",
    covers: "latin-in-cjk",
  ),
  "BIZ UDMincho"
))

// for headings
#show heading: it => {
    set text(font: (
      (
        name: "Liberation Sans",
        covers: "latin-in-cjk",
      ),
      "BIZ UDGothic"
    ))
	it.body
}

= Albert Einsteinについて

Albert Einsteinは1879年3月14日，ドイツ生まれの理論物理学者です。
彼による革命的な3つの論文「光電効果の理論」「ブラウン運動の理論」「特殊相対性理論」が発表された1905年は「奇跡の年」と呼ばれています。
```

フォント指定の効果を確認するために人名部分を英字に置き換えてみた。

組版結果は以下の通り：

{{< fig-img src="./pdf-sample-5.png" title="組版結果 (5)" link="./sample-5.pdf" width="871" >}}

{{< fig-img src="./property-sample-5.png" title="組版結果 (5)" link="./property-sample-5.pdf" >}}

んー。
やっぱモリサワ BIZ UD フォントは見やすくていいねぇ。

というわけで，ローカルマシンへの環境構築と日本語フォントの指定について試してみた。
フォント指定は要素ごとに指定する必要があり，ちょっと面倒くさい感じだな。
数式やコードなどのフォント指定については後日ということで。

## ブックマーク

- [Typstで英数字・日本語文字と見出し・本文の組み合わせごとにフォントを変える #Typst - Qiita](https://qiita.com/rockwell/items/280b4fd2109ef1e3c802)
- [Typstで黒板太字・筆記体・花文字・ギリシャ文字・ドイツ文字・ヘブライ文字(一部)を表記する方法 #組版 - Qiita](https://qiita.com/gomazarashi/items/a7e3d17b13598c1ba143)

[Typst]: https://typst.app/ "Typst: Compose papers faster"
[Typst Documentation]: https://typst.app/docs/ "Typst Documentation"
[Tutorial]: https://typst.app/docs/tutorial "Tutorial – Typst Documentation"

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Tinymist Typst]: https://marketplace.visualstudio.com/items?itemName=myriad-dreamin.tinymist "Tinymist Typst - Visual Studio Marketplace"
[Inconsolata]: https://www.levien.com/type/myfonts/inconsolata.html "Inconsolata"

## 参考文献

{{% review-paapi "B0DPXBNTRS" %}} <!-- Typst完全入門-->
{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
