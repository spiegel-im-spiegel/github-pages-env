+++
date = "2015-09-15T21:00:13+09:00"
update = "2017-12-17T08:37:00+09:00"
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

この記事は [ATOM] エディタに関する情報を Qiita に投稿した記事から再構成したものである。

- [ATOM Editor をそろそろ始めようか - Qiita](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)
- [行末の空白は EditorConfig で始末しましょう - Qiita](http://qiita.com/spiegel-im-spiegel/items/a1b4d1ad2a6693ae33e4)
- [ATOM 1.0 リリースおめでたう記念に最初からインストールし直してみた - Qiita](http://qiita.com/spiegel-im-spiegel/items/5c6dafcece9e7118877a)

そうそう。
Windows 環境が前提になっているのであしからず。

## ATOM 1.x をインストールする

[ATOM] サイトからインストールパッケージ `AtomSetup.exe` または `AtomSetup-x64.exe` をダウンロードして起動すればよい。
インストールが成功するとインストールフォルダ `%USERPROFILE%\AppData\Local\atom\bin` に PATH が通る。
これでコマンドプロンプトからも `atom` および `apm` コマンドが使えるようになる。

```text
$ atom -v

Atom    : 1.23.1
Electron: 1.6.15
Chrome  : 56.0.2924.87
Node    : 7.4.0

$ apm -v
apm  1.18.11
npm  3.10.10
node 6.9.5 x64
atom 1.23.1
python
git 2.15.1.windows.2
visual studio 2013
```

たしか node-gyp をビルドするのに（Windows 環境では） Python と Visual Studio が要るんだよね。
ううむ。

- [Windowsでnode-gypのビルドを通す - なにか作る](http://create-something.hatenadiary.jp/entry/2014/07/13/021655)

まぁ，その辺はとりあえずスルーということで。

### ATOM をアンインストールする

アンインストール自体はコンパネから「プログラムと機能」を開いて「Atom」をアンインストールする。
アンインストールを開始するもほぼ無言で完了。
男前（笑）

ただしユーザのフォルダ内には [ATOM] 関連のファイルがかなり残っているので手動で掃除する。
対象となるのは以下のフォルダ。

- `%USERPROFILE%\.atom`
- `%USERPROFILE%\AppData\Local\atom`
- `%USERPROFILE%\AppData\Local\Temp`
- `%USERPROFILE%\AppData\Roaming\Atom`

（`USERPROFILE` 環境変数は `C:\Users\username` などと設定されているはず）

`AppData` フォルダは既定では不可視になっているのでご注意を。
`%USERPROFILE%\.atom` フォルダには `keymap.cson` などの設定ファイルが入ってるので，バックアップを取っておくと安全。

`%USERPROFILE%\AppData\Local\Temp` に `Atom Crashes` フォルダがある。
どうやらクラッシュ・レポートはここに吐かれるらしい。
テンポラリ・フォルダにある古い日付のフォルダ・ファイルは大抵は削除して大丈夫なのだが，たまにヤバいやつもあるので掃除は慎重に。

*参考：* [Atom レジストリエントリを含むアンインストール方法（Windows版） - Qiita](http://qiita.com/masa36/items/c800185174bd77526a7d)

## apm stars でテーマ・パッケージを一気にインストールする

`apm` には star (☆) を付けたテーマ・パッケージを一気にインストールするコマンドがある。

この機能を使うには，まず `apm` にアカウントのトークンを登録する必要がある。アカウントのトークンは [Account](https://atom.io/account) ページから取得できる。（[GitHub] のアカウントを持っていれば，そのまま [ATOM] にも sign in できるのだが，持ってない人はどうするんだろう？）

取得したトークンを `apm login` コマンドで登録すれば OK。

```text
$ apm login
Welcome to Atom!

Before you can publish packages, you'll need an API token.

Visit your account page on Atom.io https://atom.io/account,
copy the token and paste it below when prompted.

Press [Enter] to open your account page on Atom.io.
Token> ****************
Saving token to Keychain done
```

Star を付けたテーマ・パッケージは `apm stars` コマンドで見ることができる。

```text
$ apm stars
Packages starred by you (17)
├── atom-ide-ui A collection of Atom UIs to support language services. (249978 downloads, 242 stars)
├── atom-material-syntax-dark A darker syntax theme for Atom that uses Google's Material Design color palette (147753 downloads, 184 stars)
├── auto-encoding Auto detect (61100 downloads, 137 stars)
├── editorconfig Helps developers maintain consistent coding styles between different editors (374933 downloads, 1214 stars)
├── file-icons Assign file extension icons and colours for improved visual grepping (3948068 downloads, 4808 stars)
├── git-plus Do git things without the terminal (1843260 downloads, 2408 stars)
├── git-time-machine Visually interact with git commit history for a file (281943 downloads, 933 stars)
├── go-plus Makes working with Go in Atom awesome. (439855 downloads, 479 stars)
├── highlight-line Highlights the current line in the editor (206856 downloads, 907 stars)
├── ide-typescript TypeScript and JavaScript language support for Atom-IDE. (78960 downloads, 81 stars)
├── japan-util utilities for Japanese (5068 downloads, 39 stars)
├── language-lua Add syntax highlighting and snippets to Lua files in Atom (157120 downloads, 165 stars)
├── markdown-table-editor Markdown table editor/formatter (11218 downloads, 39 stars)
├── nav-panel-plus Side panel with automatic, configurable markers for easy navigation with support for persistent bookmarks. (2052 downloads, 12 stars)
├── open-recent Open recent files in the current window, and recent folders (optionally) in a new window. (143665 downloads, 643 stars)
├── platformio-ide-terminal A terminal package for Atom, complete with themes, API and more for PlatformIO IDE. Fork of terminal-plus. (817652 downloads, 533 stars)
└── show-ideographic-space Show ideographic space (known as 全角スペース) (42647 downloads, 262 stars)

Use `apm stars --install` to install them all or visit http://atom.io/packages to read more about them.
```

さらに `--install` オプションを付ければ一気にインストールできる。

```text
$ apm stars --install
Installing atom-ide-ui to C:\Users\username\.atom\packages done
Installing atom-material-syntax-dark to C:\Users\username\.atom\packages done
Installing auto-encoding to C:\Users\username\.atom\packages done
Installing editorconfig to C:\Users\username\.atom\packages done
Installing file-icons to C:\Users\username\.atom\packages done
Installing git-plus to C:\Users\username\.atom\packages done
Installing git-time-machine to C:\Users\username\.atom\packages done
Installing go-plus to C:\Users\username\.atom\packages done
Installing highlight-line to C:\Users\username\.atom\packages done
Installing ide-typescript to C:\Users\username\.atom\packages done
Installing japan-util to C:\Users\username\.atom\packages done
Installing language-lua to C:\Users\username\.atom\packages done
Installing markdown-table-editor to C:\Users\username\.atom\packages done
Installing nav-panel-plus to C:\Users\username\.atom\packages done
Installing open-recent to C:\Users\username\.atom\packages done
Installing platformio-ide-terminal to C:\Users\username\.atom\packages done
Installing show-ideographic-space to C:\Users\username\.atom\packages done
```

これで複数マシンへの環境構築が随分楽になると思う。
なお star の管理は `apm star` または `apm unstar` コマンドでできるが，テーマ・パッケージのページでも可能[^star1]。

[^star1]: [ATOM] の star が [GitHub] の star のように [Flattr](https://flattr.com/) と連動すれば面白いんだけどねぇ。とりあえず flattr ボタンを貼り付ける手もあるけど。

以下の記事でテーマ・パッケージの感想を書いている。
個人的な印象なんであんまり参考にならないかもだけど，よろしかったらどうぞ。

- [年末なので ATOM Editor を掃除しましょう（もしくは2017年お気に入り ATOM パッケージ）]({{< ref "/remark/2017/12/favorite-atom-packages-2017.md" >}})

## ATOM の設定

### Proxy 設定

Intranet 上のマシンで外部との接続が阻まれている場合は Proxy 設定を行う。
設定には `apm config` コマンドを使う。

```text
$ apm config set https-proxy http://username:password@proxy.exsample.com:8080
```

ちゃんと設定できているかどうかは以下で確認可能。

{{< highlight text "hl_lines=8" >}}
$ apm config list -l
; cli configs
globalconfig = "C:\\Users\\username\\.atom\\.apm\\.apmrc"
user-agent = "npm/3.10.10 node/v6.9.5 win32 x64"
userconfig = "C:\\Users\\username\\.atom\\.apmrc"

; userconfig C:\Users\username\.atom\.apmrc
https-proxy = "http://username:password@proxy.exsample.com:8080"

; globalconfig C:\Users\username\.atom\.apm\.apmrc
cache = "C:\\Users\\username\\.atom\\.apm"
progress = false
{{< /highlight >}}

ファイアウォールのなかには，セキュリティ上の理由から， SSL/TLS 暗号通信を中間者攻撃[^b] でのぞき見するものがある。
このタイプのファイアウォールは SSL/TLS の証明書を書き換えてしまうため， `apm` が通信エラーになる。
この場合は以下の設定を行って強制的に SSL/TLS を通すようにするとよいらしい（取扱注意）。

[^b]: Deep Packet Inspection とか言うらしいけど，どう見たって Man-in-the-Middle Attack だろ（笑）

```text
$ apm config set strict-ssl false
```

やれやれ。

### Font Family の選択

エディタのフォントは作業効率に大きく影響する。
これは私個人の感覚だが，日本語の地の文章がゴシック体なのは辛い。
いくら綺麗だからといって Meiryo フォントをエディタに使う気にはならない。
ただし，コードに関しては視認性が一番重要。
というわけで， “Font Family” の項目に以下を指定してみる。

```text
Inconsolata, "MS Mincho"
```

[Inconsolata] は Windows の標準環境では入ってないが，何かと便利なので OpenType フォントを取ってきて「インストール」してしまえばよい。
「[游ゴシック 游明朝フォントパック](https://www.microsoft.com/ja-jp/download/details.aspx?id=49116 "Download 游ゴシック 游明朝フォントパック from Official Microsoft Download Center")」が使える場合にはこれを使う手もある。
この場合は

```text
Inconsolata, "Yu Mincho"
```

とすればよい。
ちなみに IPA 明朝は線が細すぎて不向きだった[^ipaf1]。

[^ipaf1]: まぁ印刷用に特化したフォントだからね。 IPA は Web フォント用の IPA 明朝/ゴシックフォントを開発すべき。

Windows では Tree View のフォントが汚いので，ここは素直に Meiryo UI フォントに変える。
`%USERPROFILE%\.atom\styles.less` を以下のように変更すればよい。

```css
@ui-fonts: "Meiryo UI";
.tree-view {
  font-family: @ui-fonts;
}
atom-workspace {
  font-family: @ui-fonts;
}
```

[Inconsolata]: http://www.levien.com/type/myfonts/inconsolata.html "Inconsolata"

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

[*.css]
indent_style = space
indent_size = 2

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
- `trim_trailing_whitespace` を `true` にすると行末の空白文字を削除してくれる。残念なことに [ATOM] の場合は，いわゆる「全角空白」を空白文字と見なしてくれない。まぁ全角空白を空白文字と見なす実装のほうが少ないけど（現在は[条件付きで対応]({{< ref "/remark/2016/10/warnig-from-editorconfig-at-atom.md" >}} "【ATOM Editor】 EditorConfig を使うなら Whitespace は不要")）。
- `insert_final_newline` を `true` にするとファイルの末尾が改行文字ではない場合に補完してくれる（現在は[条件付きで対応]({{< ref "/remark/2016/10/warnig-from-editorconfig-at-atom.md" >}} "【ATOM Editor】 EditorConfig を使うなら Whitespace は不要")）。

[EditorConfig] は多くのテキスト・エディタや統合開発環境に対応していて，もちろん [ATOM] にも対応パッケージがある。

- [editorconfig](https://atom.io/packages/editorconfig)

[EditorConfig]: http://editorconfig.org/ "EditorConfig"

#### 文字エンコーディングについて残念なお知らせ

`charset` では文字エンコーディングを指定するが，標準では `latin1`, `utf-8`, `utf-8-bom`, `utf-16be`, `utf-16le` しかサポートしていない。
それ以外の文字エンコーディングは実装依存ということになる。

しかも [ATOM] の場合，ファイルを新規作成する場合にこの設定が効かないようで，たとえば [ATOM] 側の “File Encoding” が `shiftjis` で  `.editorconfig` ファイル側が `charset = utf-8` の場合，新規作成ファイルは `shiftjis` にセットされ，そのまま保存される。
しかも次にそのファイルを開く場合は（`.editorconfig` ファイル側の設定が効いてしまうので） `utf8` で開いてしまい，結果派手に文字化けする（`ctrl-shift-U` で文字エンコーディングを変更すれば元に戻るけど）。

新規作成時の初期の文字エンコーディングは今のところどうしようもないが，プロジェクトごとに “File Encoding” を変更して対応するか，新規作成ファイルが開いた直後に `ctrl-shift-U` で文字エンコーディングを変更することで何とかなるだろう（ダサいけど）。

#### ファイル読み込み時に文字エンコーディングを自動判別する

もうひとつの緩和策として，既存ファイル読み込み時に [ATOM] に文字エンコードを自動判別させる手がある。
文字エンコードを自動判別を行うには [auto-encoding] を入れるとよい。
自動判別はコマンドパレットからオフにすることもできる。
ありがとー！

（古い記述だが，文字エンコードを自動判別するコードを直接 `init.coffee` に記述する方法もある： [ATOM でファイルを開いたら自動文字コード判定を行う - Qiita](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)）

[auto-encoding]: https://atom.io/packages/auto-encoding

### Git 関連機能

[ATOM] 1.18 でようやく公式に git 機能が組み込まれた。

- [ATOM 1.18 stable がリリース]({{< ref "/remark/2017/06/atom-1_18.md" >}})

ただし，キーボード操作なら [git-plus] のほうが便利である。

[git-plus] の機能は専用のコマンドパレットから呼び出すことが出来るが（メニューの「Packages」からも辿れる），呼び出しのキー割り当てを `shift-f1` にしたらだいぶ楽になった。

[git-plus]: https://atom.io/packages/git-plus

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

1. Core パッケージ [tabs](https://atom.io/packages/tabs) の設定で “Enable VCS Coloring” を有効にするとタブに表示されているファイル名が色分けされて表示される
1. Core パッケージ [git-diff](https://atom.io/packages/git-diff) の設定で “Show Icons In Editor Gutter” を有効にすると差分情報がちょっとだけ見やすくなる
1. インデントをそろえるために [highlight-column](https://atom.io/packages/highlight-column) を紹介しているページを見かけたけど，インデントをそろえる目的なら，標準の Settings で “Show Indent Guide” にチェックを入れておく方が吉

### 現在の設定

現在の設定は [Gist に貼り付け](https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c)ている。
よろしかったら参考にどうぞ。

{{< fig-gist "https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c" >}}

## ブックマーク

- [Atom - TeX Wiki](https://texwiki.texjp.org/?Atom)
- [GitHub Wiki の Markdown を Atom で編集するアレコレ - Qiita](https://qiita.com/alt-core/items/491357aadcd856c7ea5a)
- [GitHub - increments/atom-qiita-syntax: Qiita theme syntax for Atom](https://github.com/increments/atom-qiita-syntax)

- [ATOM で Go — プログラミング言語 Go]({{< ref "/golang/golang-with-atom.md" >}})

Qiita の「[ATOM Editor をそろそろ始めようか](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)」はもうメンテしてません。
ゴメンペコン。

[ATOM]: https://atom.io/ "Atom"
[GitHub]: https://github.com/ "GitHub"
