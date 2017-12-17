+++
title = "年末なので ATOM Editor を掃除しましょう（もしくは2017年お気に入り ATOM パッケージ）"
date =  "2017-12-16T18:16:31+09:00"
update = "2017-12-17T08:37:00+09:00"
description = "というわけで，唐突に ATOM エディタの掃除とか始めてしまう。ついでに最近のお気に入りパッケージとか紹介してみる。"
image = "/images/attention/remark.jpg"
tags = ["atom", "editor", "tools", "git", "golang", "japanese", "language"]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

なんで日本人はクソ忙しい年末に「大掃除」とかするんでしょうねぇ。
普段からやっとけっての，自分。
分かってます。
現実逃避ってやつですね（笑）

というわけで，唐突に [ATOM] エディタの掃除とか始めてしまう。
ついでに最近のお気に入りパッケージとか紹介してみる。

## [ATOM] のいいところ，悪いところ

[ATOM] のいいところって何だろう。
私の場合，あるエディタが気に入る理由はひとつしかない。
それは

**手に馴染むかどうか**

である。
これは極めて感覚的かつ主観的なので説明しづらい。
でも，これって殆どの人がそうじゃないだろうか。

たとえば私が「[ATOM] の◯◯機能がチョー便利！」とか「◯◯パッケージが素敵！」とか言ったところで「それ◯◯で出来るよ」と返されるのがオチである。
「好き」という感情は感覚的なものなので合理的に説明するのは難しいし，好悪を評価軸に他人を説得するのは更に無理筋というものである。

逆に [ATOM] の悪いところは幾らでも挙げられる[^tsundere1]。
すぐ思いつくところでは

[^tsundere1]: あっ，私ツンデレ属性とかないですから（笑）

1. 起動が遅い
2. 全体的に動きがモッサリしている
3. 巨大ファイルを（事実上）扱えない

といった感じ。
これのせいで私はテキスト・エディタを「[秀丸](http://hide.maruo.co.jp/software/hidemaru.html)」から完全に乗り換えることができないでいる。
今のところ [ATOM] での作業はテキスト表示・入力・編集作業全体の9割5分といったところ。
ちょこっとだけファイルの中身を覗いてみたいとか，10万行くらいの巨大 CSV ファイル（しかも Shift-JIS）を開きたいとかいったことには  [ATOM] は向かない。

まぁ，でも，私がよく使うパッケージ等を紹介することで「[ATOM] エディタってこんな感じ」くらいのことは分かってもらえるのではないかと淡い期待を寄せてみる。

## インストールしたパッケージを整理する

さて，前口上はこのくらいにして，タイトル通り [ATOM] エディタの掃除を始めよう。

[ATOM] は多様なパッケージ群が特徴でインストールも非常に簡単なのはいいのだが [ATOM] を長く使っているとインストールしたまま忘れ去ってしまったパッケージのひとつやふたつはあるものだ。
自分の環境にどんなパッケージが入っているか調べるには以下のコマンドを使う。

```text
$ apm list -i -b
```

ちなみにパッケージだけを見る場合には `-p` オプションを付けて

```text
$ apm list -i -b -p
```

とすればよい（テーマのみなら `-t` オプションを付ける）。
表示される一覧を見て「何だっけ？ これ」っていうのがあったなら削除してしまって構わない。
ただし，パッケージ同士が依存関係にある場合は [ATOM] を再起動した際に再インストールを要求されることがある。

パッケージの削除は [ATOM] の Setting で行うかコマンドラインで

```text
$ apm delete package-name
```

とすればいい。

お気に入りのパッケージについては，各パッケージの Web ページで ☆ (star) を付けることをお勧めする。
Star を付けたパッケージは初期インストール時に

```text
$ apm stars --install
```

で一気にインストールすることが可能になる。
詳しくは「[ATOM Editor に関するメモ]({{< relref "remark/2015/atom-editor.md" >}})」を参照のこと。

## 2017年のお気に入りパッケージ

それでは以下に2017年末時点でよく使うパッケージを紹介する。

### テーマ

最近は [atom-material-syntax-dark]。

本当はライト・テーマにしたいのだが，いいのがないんだよねぇ。
何だってみんなライト・テーマの背景を真っ白にするんだろう。
ちょっとベージュっぽい背景色にすれば，それで充分眼に優しい配色になるのに。

黒背景もそれはそれで視覚的にキツいんだけど，真っ白背景よりはマシなので仕方なくダーク・テーマを使っている。

選択するテーマによるのだが [highlight-line] でカーソル行にアンダーラインを表示させている。
`styles.less` ファイルで簡単に色などを調整できるのがよい。

[atom-material-syntax-dark]: https://atom.io/themes/atom-material-syntax-dark
[highlight-line]: https://atom.io/packages/highlight-line

### Git 関連

[ATOM] 1.18 でコア・パッケージに [github] が導入され，ようやく GUI で git の基本操作が出来るようになった。
ただし，キーボード・ベースの操作では [git-plus] の方が使いやすい。
てか，私は主に [git-plus] のほうを使っている。

他には， commit 履歴を視覚的に見るために [git-time-machine] を使うことがある。
ちょっと癖が強いのが難点なのだが。

[github]: https://atom.io/packages/github
[git-plus]: https://atom.io/packages/git-plus
[git-time-machine]: https://atom.io/packages/git-time-machine

### 自動保存

ファイルの自動保存機能として，コア・パッケージに [autosave] というのがあるのだが，ファイル保存をトリガにして動く機能もあるので（[git-plus] の commit コメントとか），使い所が微妙である。

ちなみに，かつての [save-session] 機能は [ATOM] 本体に取り込まれているそうで，保存を忘れて [ATOM] を閉じても，次に開くときは未保存状態のまま復元されるようだ。
よしよし。

[autosave]: https://atom.io/packages/autosave
[save-session]: https://atom.io/packages/save-session

### go-plus と atom-ide

はっきり言おう。
私が [ATOM] にロックインされている最大の理由が [go-plus] パッケージである。
コードの入力補完と整形，lint，テスト，カバレッジ 等々... これがなかったらコーディング効率半減どころか 70% 減かも。

[go-plus] パッケージを入れると副パッケージとして，以下も併せてインストールされる。

- [go-debug]
- [go-signature-statusbar]
- [hyperclick] （[atom-ide-ui] が入ってる場合は不要？）

また [alecthomas/gometalinter] をはじめとする大量のツールがインストールされるため `GOPATH` および `PATH` 環境変数の設定をしておくこと。

最近の [go-plus] は [atom-ide-ui] を導入するよう勧めてくる。
[atom-ide-ui] は最近発表された [ATOM] 用の IDE (Integrated Development Environment; 統合開発環境) パッケージで [LSP (Language Server Protocol)](http://langserver.org/) に対応しているのが売りである。

- [Atom IDE](https://ide.atom.io/)

どうも [go-plus] は [atom-ide-ui] とも両立するようになったらしい。
[atom-ide-ui] は言語ごとのサブパッケージと組み合わせるようになっていて， [Go 言語]に対応しているものとしては [ide-go] というパッケージがあるのだが，これがウチの環境では動いてくれんのよ。
[ide-go] のバックエンドには [sourcegraph/go-langserver] が動いているのだが，こいつが Windows と相性が悪い気がする。

[Go 言語]に関しては [go-plus] があれば [atom-ide-ui] は全く必要ないのだが，他の言語（たとえば [ide-typescript] とか）用に入れてある。
そのうち統合されることがあるのだろうか。

[go-plus]: https://atom.io/packages/go-plus
[go-debug]: https://atom.io/packages/go-debug
[go-signature-statusbar]: https://atom.io/packages/go-signature-statusbar
[hyperclick]: https://atom.io/packages/hyperclick
[atom-ide-ui]: https://atom.io/packages/atom-ide-ui
[ide-go]: https://atom.io/packages/ide-go
[ide-typescript]: https://atom.io/packages/ide-typescript
[sourcegraph/go-langserver]: https://github.com/sourcegraph/go-langserver "GitHub - sourcegraph/go-langserver: Go language server to add Go support to editors and other tools that use the Language Server Protocol (LSP)"
[alecthomas/gometalinter]: https://github.com/alecthomas/gometalinter "GitHub - alecthomas/gometalinter: Concurrently run Go lint tools and normalise their output"

### [EditorConfig]

コード書きで [EditorConfig] を使わないやつは，もはや evil と言っていいだろう。
チームでソースコードのフォーマットを合わせるなら [EditorConfig] は必須だし，ひとりで作業するときもリポジトリに必ず `.editorconfig` を含める習慣をつけていきたいものである。

[EditorConfig]: http://editorconfig.org/ "EditorConfig"

### 機能を直接埋め込む

[ATOM] では `init.coffee` ファイルで機能を記述し組み込むことが出来る。
パッケージにするまでもない小さな処理などで重宝する。

私が組み込んでいるのは以下の機能。

- [ATOM Editor で現在日時を挿入する]({{< relref "remark/2015/insert-datetime-in-atom-editor.md" >}})
- [ATOM Editor で Amazon Associate ID を含んだ商品 URL を生成する]({{< relref "remark/2015/insert-amazon-url-with-associate-id-in-atom-editor.md" >}})

### 日本語関連のパッケージ

以降は簡単に箇条書きで。

- *[auto-encoding]* ： 文字エンコーディングを自動で判定してくれる。 Shift-JIS や EUC-JP なファイルを開く時に重宝する。たまに間違うのがご愛嬌
- *[japan-util]* ： 日本語用の文字変換パッケージ半角/全角変換や平仮名/片仮名変換とかしてくれる
- *[show-ideographic-space]* ： いわゆる全角空白文字を視覚化してくれる。これがないとコンパイルエラー時にパニクるハメになる[^zs1]。見せ方は `styles.less` ファイルでカスタマイズ可能

[^zs1]: Windows 環境なら，全角空白文字の誤入力は IME のプロパティでスペース・キー押下で常に「半角空白」を入力するように設定すればほとんど防げるんだどね（全角空白を入力する場合は `shift-space` 押下）。

[auto-encoding]: https://atom.io/packages/auto-encoding
[japan-util]: https://atom.io/packages/japan-util
[show-ideographic-space]: https://atom.io/packages/show-ideographic-space

### その他のお気に入りパッケージ

- *[autoclose-html]* ： HTML 入力でタグを入力すると自動的に閉じタグを補完してくれる。便利なのだが HTML を直に弄ることが少なくなったので削除した
- *[file-icons]* ： Tree View やタブのアイコン表示を見やすくしてくれる
- *[language-lua]* ： Lua 言語用のパッケージ。 [atom-ide-ui] に対応するものがないっぽいので。 [NYAGOS] 用のバッチ処理等を書くのに Lua 言語を使うのよ
- *[markdown-table-editor]* ： Markdown のテーブル作成支援パッケージ。めっさ便利
- *[nav-panel-plus]* ： アウトライン表示。 [atom-ide-ui] のアウトラインで表示できない言語（markdown 等）はこちらで。以前は [symbols-tree-view] を使っていたが，こちらに乗り換えた
- *[open-recent]* ： File メニューに “Open Recent” 項目を追加し，最近開いたファイルやプロジェクトを表示してくれる
- *[platformio-ide-terminal]* ： [ATOM] 内でターミナルを起動する。 shell  や環境変数を指定できるのが素敵。私は shell として [NYAGOS] を指定している

[autoclose-html]: https://atom.io/packages/autoclose-html
[file-icons]: https://atom.io/packages/file-icons
[language-lua]: https://atom.io/packages/language-lua
[markdown-table-editor]: https://atom.io/packages/markdown-table-editor
[nav-panel-plus]: https://atom.io/packages/nav-panel-plus
[symbols-tree-view]: https://atom.io/packages/symbols-tree-view
[open-recent]: https://atom.io/packages/open-recent
[platformio-ide-terminal]: https://atom.io/packages/platformio-ide-terminal

## 未インストールだけど気になってるパッケージ

こちらも簡単に箇条書きで。

- *[fdocblockr]* ： `/** */` みたいなブロックコメントを生成してくれる。 Java とかならありがたいが [Go 言語]は微妙
- *[language-plantuml] & [plantuml-viewer]* ： PlantUML 用の言語パッケージと画像表示パッケージ。 SVG や PNG といった画像データとして保存できるのが素敵。 DOT 言語に変換されるので， [Graphviz] を用意する必要がある
    - [AtomとPlantUMLで爆速UMLモデリング - Qiita](http://qiita.com/nakahashi/items/3d88655f055ca6a2617c)
	- [Atom+PlantUMLで見た目もいい感じのシーケンス図を作成する - Qiita](http://qiita.com/k_nakayama/items/77ca73753ebd049a66de)
- *[linter]* ： これ単独では使えなくて，言語ごとに lint パッケージを用意する必要がある。 [Go 言語]では [go-plus] が独自の強力な lint 機能を持っているため不要
    - [AtomにESLint導入した - Qiita](http://qiita.com/pechefamille/items/40966a0c78846f4053c9)
- *[tablr]* ： CSV ファイルをスプレッドシート風に表示・編集できる。小さいファイルならいいんだけど， CSV ファイルって大抵が巨大ファイルだからなぁ。結局 Office ツール使ったほうがよかったり


[fdocblockr]: https://atom.io/packages/docblockr
[language-plantuml]: https://atom.io/packages/language-plantuml
[linter]: https://atom.io/packages/linter
[plantuml-viewer]: https://atom.io/packages/plantuml-viewer
[tablr]: https://atom.io/packages/tablr

## ブックマーク

- [Atom - TeX Wiki](https://texwiki.texjp.org/?Atom)
- [［保存版］Atomエディタ 便利なパッケージ一覧！ 全23社のWebエンジニア・デザイナーがおすすめを紹介 - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://employment.en-japan.com/engineerhub/entry/2017/08/10/110000)

- [ATOM Editor に関するメモ]({{<relref "remark/2015/atom-editor.md" >}})
- [ATOM × NYAGOS ＝ ♥]({{< relref "remark/2016/11/nyagos-with-atom.md" >}})
- [ATOM で Go]({{< relref "golang/golang-with-atom.md" >}})

[ATOM]: https://atom.io/ "Atom"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[Graphviz]: http://www.graphviz.org/ "Graphviz | Graphviz - Graph Visualization Software"
