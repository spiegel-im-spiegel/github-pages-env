+++
date = "2015-09-15T21:00:13+09:00"
update = "2015-12-12T12:32:44+09:00"
description = "ATOM Editor に関するメモ。 Windows 環境が前提になっているのであしからず。"
draft = false
tags = ["atom", "editor", "tools"]
title = "ATOM Editor に関するメモ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

この記事は [ATOM] Editor に関する情報を Qiita に投稿した記事から再構成したものである。

- [ATOM Editor をそろそろ始めようか - Qiita](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)
- [行末の空白は EditorConfig で始末しましょう - Qiita](http://qiita.com/spiegel-im-spiegel/items/a1b4d1ad2a6693ae33e4)
- [ATOM 1.0 リリースおめでたう記念に最初からインストールし直してみた - Qiita](http://qiita.com/spiegel-im-spiegel/items/5c6dafcece9e7118877a)

そうそう。
Windows 環境が前提になっているのであしからず。

## ATOM 1.0 をインストールする

[ATOM] サイトからインストールパッケージ `AtomSetup.exe` をダウンロードして起動すればよい。
インストールが成功するとインストールフォルダ `C:\Users\username\AppData\Local\atom\bin` に PATH が通る。
これでコマンドプロンプトからも `atom` および `apm` コマンドが使えるようになる。

```
C:>atom -v
[7696:0915/112859:INFO:CONSOLE(0)] 1.0.11


C:>apm -v
apm  1.0.4
npm  2.13.3
node 0.10.40
python
git 2.5.2.windows.2
visual studio
```

上記の環境では Python と Visual Studio は入れてないのでバージョンが入ってないのかな。
たしか node-gyp をビルドするのに（Windows 環境では） Python と Visual Studio が要るんだよね。
ううむ。

- [Windowsでnode-gypのビルドを通す - なにか作る](http://create-something.hatenadiary.jp/entry/2014/07/13/021655)

### ATOM をアンインストールする

アンインストール自体はコンパネから「プログラムと機能」を開いて「Atom」をアンインストールする。アンインストールを開始するもほぼ無言で完了。男前（笑）

ただしユーザのフォルダ内には [ATOM] 関連のファイルがかなり残っているので手動で掃除する。対象となるのは以下のフォルダ。

- `C:\Users\username\.atom`
- `C:\Users\username\AppData\Local\atom`
- `C:\Users\username\AppData\Local\Temp`
- `C:\Users\username\AppData\Roaming\Atom`

`AppData` フォルダは既定では不可視になっているのでご注意を。 `C:\Users\username\.atom` フォルダには `keymap.cson` などの設定ファイルが入ってるので，バックアップを取っておくと安全。

`C:\Users\username\AppData\Local\Temp` には `Atom Crashes` フォルダがある。どうやらクラッシュ・レポートはここに吐かれるらしい。テンポラリ・フォルダにある古い日付のフォルダ・ファイルは，大概は削除して大丈夫なのだが，たまにヤバいやつもあるので掃除は慎重に。

## apm stars でテーマ・パッケージを一気にインストールする

`apm` には star を付けたテーマ・パッケージを一気にインストールするコマンドがある。

この機能を使うには，まず `apm` にアカウントのトークンを登録する必要がある。アカウントのトークンは [Account](https://atom.io/account) ページから取得できる。（[GitHub] のアカウントを持っていれば，そのまま [ATOM] にも sign in できるのだが，持ってない人はどうするんだろう？）

取得したトークンを `apm login` コマンドで登録すれば OK。

```
C:>apm login
Welcome to Atom!

Before you can publish packages, you'll need an API token.

Visit your account page on Atom.io https://atom.io/account,
copy the token and paste it below when prompted.

Press [Enter] to open your account page on Atom.io.
Token> ****************
Saving token to Keychain done
```

Star を付けたテーマ・パッケージは `apm stars` コマンドで見ることができる。

```bash
C:>apm stars
Packages starred by you (13)
├── atom-monokai Monokai syntax theme for Atom Dark & Light UI, One Dark & Light, and Seti UI (27359 downloads, 39 stars)
├── autoclose-html Automates closing of HTML Tags (83771 downloads, 296 stars)
├── editorconfig Helps developers maintain consistent coding styles between different editors (55732 downloads, 446 stars)
├── file-icons Assign file extension icons and colours for improved visual grepping (305041 downloads, 1561 stars)
├── git-plus Do git things without the terminal (264729 downloads, 861 stars)
├── highlight-line Highlights the current line in the editor (54346 downloads, 423 stars)
├── japan-util utilities for Japanese (858 downloads, 12 stars)
├── open-recent Open recent files in the current window, and recent folders (optionally) in a new window. (11595 downloads, 142 stars)
├── quick-highlight Highlight text quickly. (943 downloads, 11 stars)
├── show-ideographic-space Show ideographic space (known as 全角スペース) (4255 downloads, 57 stars)
├── symbols-tree-view A symbols view like taglist (19099 downloads, 181 stars)
├── tablr Edit CSV files using a table editor (538 downloads, 16 stars)
└── wrap-style Select word warp style. (10 downloads, 2 stars)

Use `apm stars --install` to install them all or visit http://atom.io/packages to read more about them.
```

さらに `--install` オプションを付ければ一気にインストールできる。

```bash
C:>apm stars --install
Installing atom-monokai to C:\Users\username\.atom\packages done
Installing autoclose-html to C:\Users\username\.atom\packages done
Installing editorconfig to C:\Users\username\.atom\packages done
Installing file-icons to C:\Users\username\.atom\packages done
Installing git-plus to C:\Users\username\.atom\packages done
Installing highlight-line to C:\Users\username\.atom\packages done
Installing japan-util to C:\Users\username\.atom\packages done
Installing open-recent to C:\Users\username\.atom\packages done
Installing quick-highlight to C:\Users\username\.atom\packages done
Installing show-ideographic-space to C:\Users\username\.atom\packages done
Installing symbols-tree-view to C:\Users\username\.atom\packages done
Installing tablr to C:\Users\username\.atom\packages done
Installing wrap-style to C:\Users\username\.atom\packages done
```

これで複数マシンへの環境構築が随分楽になると思う。なお star の管理は `apm star` または `apm unstar` コマンドでできるが，テーマ・パッケージのページでも可能。

（[ATOM] の star が [GitHub] の star のように [Flattr](https://flattr.com/) と連動すれば面白いんだけどねぇ。とりあえず flattr ボタンを貼り付ける手もあるけど）

「[ATOM の Theme / Package の感想文（2015-06-10） - Qiita](http://qiita.com/spiegel-im-spiegel/items/115fea37ad2e515f6641)」にテーマ・パッケージの感想を書いている。個人的な印象なんであんまり参考にならないかもだけど，よろしかったらどうぞ。

## ATOM の設定

### Proxy 設定

Intranet 上のマシンで外部との接続が阻まれている場合は Proxy 設定を行う。
設定には `apm` コマンドを使う。

```bash
C:>apm config set https-proxy http://username:password@proxy.exsample.com:8080
```

Firewall のなかには，セキュリティ上の理由から， SSL/TLS 暗号通信を中間者攻撃[^b] でのぞき見するものがある。
このタイプの Firewall/Proxy は SSL/TLS の証明書を書き換えてしまうため， `apm` が通信エラーになる。
この場合は以下の設定を行って強制的に SSL/TLS を通すようにするとよいらしい（取扱注意）。

[^b]: Deep Packet Inspection とか言うらしいけど，どう見たって Man-in-the-Middle Attack だろ（笑）

```bash
C:>apm config set strict-ssl false
```

やれやれ。

### Font Family の選択

エディタのフォントは作業効率に大きく影響する。
これは私個人の感覚だが，日本語の地の文章がゴシック体なのは辛い。
いくら綺麗だからといって Meiryo フォントをエディタに使う気にはならない。
ただし，コードに関しては視認性が一番重要。
ということで，フォントの指定は以下で無問題[^a]。

```
"Inconsolata", "MS Mincho"
```

[^a]: 日本語フォントの指定が MS 明朝なのは，どんな日本語 Windows 環境でも MS 明朝は必ず入ってるからという理由だけなので，自分の感覚で見易いフォントがあればそちらを使うべき。ちなみに IPA 明朝は線が細すぎて不向きだった。 [Inconsolata] は OpenType フォントを取ってきて「インストール」してしまえばよい。

Windows では Tree View のフォントが汚いので，ここは素直に Meiryo UI フォントに変える。 `%USERPROFILE%\.atom\styles.less` を以下のように変更する。

```css
.tree-view {
	font-family: "Meiryo UI";
}
```

また

```css
atom-workspace {
  font-family: "Meiryo UI";
}
```

とすればタブや Settings 画面のフォントも変えられる。

### EditorConfig

[EditorConfig] があればタブや改行コードなどの設定を統一できる。
これは特に複数人で作業する場合に威力を発揮する。
たとえば，このサイトの作業環境では以下のように設定している。

```ini
root = true

[*]
end_of_line = lf
charset = utf-8
indent_style = tab
indent_size = 4
trim_trailing_whitespace = true
insert_final_newline = true

[*.html]
insert_final_newline = false

[*.md]
indent_style = space
indent_size = 4
trim_trailing_whitespace = false
```

[EditorConfig] はフォルダを遡って `.editorconfig` ファイルを探し，フォルダの上から順番に評価していく。
`root = true` の記述がないとどこまでも上の階層に遡っていくので，プロジェクトのトップ・フォルダの `.editorconfig` には必ずこれを記述すること。

- `[*]` は対象となるファイルを指定している。 `[*]` なら全てのファイルが対象である。 `[*.html]` は拡張子 `html` のファイルが対象となる。
- `end_of_line` では改行コードを指定する。 `lf`, `cr`, `crlf` から選択できる。
- `indent_style` はインデント（タブ）のスタイルを指定する。 `tab` または `space` を指定する。 `space` にすると，いわゆる「ソフトタブ」になる。
- `indent_size` はタブの幅を指定する。 `indent_style` と `indent_size` は常にセットで指定すると間違いがない。
- `trim_trailing_whitespace` を `true` にすると行末の空白文字を削除してくれる。残念なことに [ATOM] の場合は，いわゆる「全角空白」を空白文字と見なしてくれない。まぁ全角空白を空白文字と見なす実装のほうが少ないけど。ただし現在は機能していない。
- `insert_final_newline` を `true` にするとファイルの末尾が改行文字ではない場合に補完してくれる。ただし現在は機能していない。

[EditorConfig] は多くのテキスト・エディタや統合開発環境に対応していて，もちろん [ATOM] にも対応パッケージがある。

- [editorconfig](https://atom.io/packages/editorconfig)

#### 文字エンコーディングについて残念なお知らせ

`charset` では文字エンコーディングを指定するが，標準では `latin1`, `utf-8`, `utf-8-bom`, `utf-16be`, `utf-16le` しかサポートしていない。
それ以外の文字エンコーディングは実装依存ということになる。

しかも [ATOM] の場合，ファイルを新規作成する場合にこの設定が効かないようで，たとえば [ATOM] 側の “File Encoding” が `shiftjis` で  `.editorconfig` ファイル側が `charset = utf-8` の場合，新規作成ファイルは `shiftjis` にセットされ，そのまま保存される。
しかも次にそのファイルを開く場合は（`.editorconfig` ファイル側の設定が効いてしまうので） `utf8` で開いてしまい，結果派手に文字化けする（`ctrl-shift-U` で文字エンコーディングを変更すれば元に戻るけど）。

新規作成時の初期の文字エンコーディングは今のところどうしようもないが，プロジェクトごとに “File Encoding” を変更して対応するか，新規作成ファイルが開いた直後に `ctrl-shift-U` で文字エンコーディングを変更することで何とかなるだろう（ダサいけど）。

#### ファイル読み込み時に文字エンコーディングを自動判別する

もうひとつの緩和策としては，既存ファイル読み込み時に [ATOM] に文字エンコードを自動判別させることだ。ただし，この機能を持つ Package は今のところ存在しないようなので自前で何とかするしかない。

- [ATOM でファイルを開いたら自動文字コード判定を行う - Qiita](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)

この設定を行うには node.js のフルパッケージが必要。

（Linux や Mac な人は依存関係で node.js がインストールされると思うけど， Windows では [ATOM] インストール時に一部機能が同梱されているだけなので（しかもバージョンが古い），フル機能を使うには別途インストールする必要あり）

Windows の場合は `%USERPROFILE%\.atom` フォルダに移動する。その後， `npm` コマンドを使って `iconv-lite` と `jschardet` をインストールする。

```bash
C:>cd C:\Users\username\.atom
C:\Users\username\.atom>npm install iconv-lite
iconv-lite@0.4.10 node_modules\iconv-lite

C:\Users\username\.atom>npm install jschardet
jschardet@1.1.1 node_modules\jschardet
```

すると `%USERPROFILE%\.atom\node_modules` フォルダが作成され，その中に `iconv-lite` と `jschardet` がインストールされているはずである。

次は `%USERPROFILE%\.atom\init.coffee` ファイルに以下の記述を追加する。

```coffee
fs = require('fs')

atom.workspace.onDidOpen ->
  editor = atom.workspace.getActiveTextEditor()

  try
    filePath = editor.getPath()
  catch error
    return
  return unless fs.existsSync(filePath)

  jschardet = require 'jschardet'
  iconv = require 'iconv-lite'
  fs.readFile filePath, (error, buffer) =>
    return if error?
    {encoding} = jschardet.detect(buffer) ? {}
    encoding = 'utf8' if encoding is 'ascii'
    return unless iconv.encodingExists(encoding)

    encoding = encoding.toLowerCase().replace(/[^0-9a-z]|:\d{4}$/g, '')
    editor.setEncoding(encoding)
```

「[ATOM でファイルを開いたら自動文字コード判定を行う](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)」によると，これは [encoding-selector](https://atom.io/packages/encoding-selector) からの流用らしい。ただし現在， [encoding-selector](https://atom.io/packages/encoding-selector) は Core Package に入ってるので atom フォルダをひっくり返してもソースコードは見当たらない。ので，GitHub repository [atom/encoding-selector](https://github.com/atom/encoding-selector) にある [lib/encoding-list-view.coffee](https://github.com/atom/encoding-selector/blob/master/lib/encoding-list-view.coffee) を参考にするといいだろう。 `detectEncoding:` のあたりである。

これで [ATOM] を起動して既定の文字エンコーディングでない適当なファイルを読み込ませてみれば確認できる。ただし，自動判別は万能じゃない（たまに間違う）ので，その辺は悪しからずってことで。

一番いいのは [encoding-selector](https://atom.io/packages/encoding-selector) がファイル読み込み時に自動判別する機能を付けてくれることなんだけど。誰かやらないかな。個人的には [ATOM] 開発に積極的に commit する気はないので，完全に他人任せなのだが。

### Git 関連機能

GitHub が作ったという割には git 機能が貧弱なのは何故なのだろう。特に [git-plus](https://atom.io/packages/git-plus) 相当の機能が Core パッケージに入ってないのは解せぬ。というわけでインストールした。公開してくださった方に感謝。

- [git-plus](https://atom.io/packages/git-plus)

[git-plus](https://atom.io/packages/git-plus) は主にコマンドパレットで操作するが（メニューの「Packages」からも辿れる），コマンドパレットの呼び出しを f1 キーに割り当てたらだいぶ使いやすくなった。

[git-control](https://atom.io/packages/git-control) は GUI で最初の頃は便利だったのだが， [git-plus](https://atom.io/packages/git-plus) が手に馴染んできたら使わなくなった。
submodule の処理が得意ではないっぽいのもマイナス。

### 矩形選択

矩形選択用に `%USERPROFILE%\.atom\keymap.cson` ファイルに対して以下のキー設定を行った

```cson
'body':
  'alt-shift-down': 'editor:add-selection-below'
  'alt-shift-left': 'core:select-left'
  'alt-shift-right': 'core:select-right'
  'alt-shift-up': 'editor:add-selection-above'
```

これで Alt+Shift キーを押しながらカーソルを上下左右に動かせば矩形選択ができる。実はこれらの操作のキー割り当ては元々こうなっている。

| Keybind         | コマンド                     |
|:----------------|:-----------------------------|
| `ctrl-alt-down` | `editor:add-selection-below` |
| `ctrl-alt-up`   | `editor:add-selection-above` |
| `shift-left`    | `core:select-left`           |
| `shift-right`   | `core:select-right`          |

ただし，うちのパソコンでは Ctrl+Alt キーを押しながらカーソルを動かすと（Windows がキーを横取りして）ディスプレイの向きが変わってしまうので，かなり切ないことになってしまう。

実際には `editor:add-selection-below` および `editor:add-selection-above` はマルチカーソル・モードのコマンドである。マルチカーソル・モードは癖がある感じだが慣れれば結構使えるかも。

キーボード操作ではなく，マウス操作で矩形選択がしたい場合は [sublime-style-column-selection](https://atom.io/packages/Sublime-Style-Column-Selection) を導入するとよい。 Windows ではマウスのセンターボタンを押しながらマウスを動かす。

いずれのやり方でも，日本語混じりのテキストではうまくいかない。
やれやれ。

### その他

1. 自動保存を行う Core パッケージ [autosave](https://atom.io/packages/autosave) は既定で無効になってる。私は [git-plus](https://atom.io/packages/git-plus) で commit する前に保存するのをどうしても忘れるので，これを有効にしておくと吉。ただし [git-plus](https://atom.io/packages/git-plus) で，自動保存により勝手に commit が走るので注意。
1. Core パッケージ [tabs](https://atom.io/packages/tabs) の設定で “Enable VCS Coloring” を有効にするとタブに表示されているファイル名が色分けされて表示される。
1. Core パッケージ [git-diff](https://atom.io/packages/git-diff) の設定で “Show Icons In Editor Gutter” を有効にすると差分情報がちょっとだけ見やすくなる。
1. インデントをそろえるために [highlight-column](https://atom.io/packages/highlight-column) を紹介しているページを見かけたけど，インデントをそろえる目的なら，標準の Settings で “Show Indent Guide” にチェックを入れておく方が吉。
1. 秀丸では大変重宝したアウトライン解析。[symbols-tree-view](https://atom.io/packages/symbols-tree-view) があればアウトライン解析を行って右側のサイドバーに表示する。
1. [japan-util](https://atom.io/packages/japan-util) : 全角・半角変換機能。英数字を全角で書いたり仮名文字を半角で書いたりするバカがいるので，意外と使うのよ，これ。
1. [highlight-line](https://atom.io/packages/highlight-line) でカーソル行をハイライトにしたり下線を引いたりできる。色のカスタマイズは `%USERPROFILE%\.atom\styles.less` ファイルで行う。
1. [open-recent](https://atom.io/packages/open-recent) は最近開いたファイルやフォルダを覚えておいてくれる便利なやつ。てか，なぜこれが標準で搭載されてないのだ。
1. [show-ideographic-space](https://atom.io/packages/show-ideographic-space) : いわゆる「全角空白」を視覚化してくれる。見せ方は `%USERPROFILE%\.atom\styles.less` ファイルでカスタマイズ可能。実際には IME のプロパティでスペース・キー押下で常に「半角空白」を入力するように設定すればほとんど防げるんだどね（全角空白を入力する場合は `shift-space` 押下）。
1. [autoclose-html](https://atom.io/packages/autoclose-html) : HTML 入力でタグを入力すると自動的に閉じタグを補完してくれる。自動補完の機能は色々あって，大抵は [autocomplete-plus](https://atom.io/packages/autocomplete-plus) のサブパッケージだったりするのだが，これだけは毛色が違う（笑）
1. 特定の単語をマークしておける [quick-highlight](https://atom.io/packages/quick-highlight) が結構使える。 Windows だとキーに割り当てられないので、 toggle を適当なファンクションキーとかに割り当てておくとめっさ便利。
- むむっ。 [latex](https://atom.io/packages/latex) パッケージなるものがあるなぁ。
1. [tablr](https://atom.io/packages/tablr) : CSV Editor。なにこれ素敵！

### 現在の設定

現在の設定は [Gist に貼り付け](https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c)ている。
よろしかったら参考にどうぞ。

{{< fig-gist "https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c" >}}

## ブックマーク

[ATOM] に関するブックマークは [Qiita でメンテナンス](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)している。
こちらも併せてどうぞ。

[ATOM]: https://atom.io/ "Atom"
[GitHub]: https://github.com/ "GitHub"
[Inconsolata]: http://www.levien.com/type/myfonts/inconsolata.html "Inconsolata"
[EditorConfig]: http://editorconfig.org/ "EditorConfig"
