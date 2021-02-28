+++
title = "パソコンに Visual Studio Code を導入する（再チャレンジ）"
date =  "2021-02-27T14:23:23+09:00"
description = "まずはインストールと初期設定"
image = "/images/attention/kitten.jpg"
tags = ["vscode", "editor", "ubuntu", "windows"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今の職場では当然のように Windows 10 機を支給されていてテキスト・エディタ等の普段使いのツールであれば割と自由に使わせてくれるのはいいのだが，6年ほど使い込んで手に馴染んでる筈の [ATOM] が使えなくてねぇ。
理由は大きく2つ。

1. 起動が遅すぎる（もともと遅いが堪えられないほど遅い）
2. [Go] 言語開発支援の [go-plus] が実質的に動かない

特に2番目が致命的。
私の観測範囲が狭いせいか，この手の話をほとんど聞かないのだが Windows で [ATOM] で [Go] で開発をしてる人はいないってことなのかねぇ。

さらに言うと [ATOM] を開発している GitHub が [VS Code] を開発している Microsoft に買収され，オフィシャルの Go Language Server である [gopls] が，これまた [Google オフィシャルの VS Code 用拡張機能](https://marketplace.visualstudio.com/items?itemName=golang.go "Go - Visual Studio Marketplace")の[既定の Language Server になった](https://blog.golang.org/gopls-vscode-go "Gopls on by default in the VS Code Go extension - The Go Blog")というのも大きい。
[ATOM] の [go-plus] なんて今だに gocode 使ってるんだぜ。
完全に廃れているよなぁ `orz`

というわけで，観念して [VS Code] を導入することにした。

実は昨年5月に自宅の [Ubuntu 機に VS Code を入れた]({{< ref "/remark/2020/05/installing-vscode-in-ubuntu.md" >}} "Ubuntu に Visual Studio Code を導入する")んだけど，結局挫折して削除しちゃったんだよねぇ。
今回は職場の Windows 機で先行して導入している。
なので Windows 機と [Ubuntu] 機を比較しながら，何回かに分けて記事を書いていく予定である。

1. [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}}) ← イマココ
2. [Go と VS Code]({{< ref "/remark/2021/02/golang-with-vscode.md" >}})
3. [Markdown と VS Code]({{< ref "/remark/2021/02/markdown-with-vscode.md" >}})

## [VS Code] のインストール

Windows 版のインストールは[サイト][VS Code]からインストーラをダウンロードすればよい。

問題は [Ubuntu] だが Snap 版と APT 版がある。
昨年の話だが Snap 版を入れたら日本語入力が壊滅してたので，今回も選択肢には入れない。

APT 版は Microsoft がリポジトリを公開しているので，これを設定してインストールする。

{{< fig-quote class="nobox" type="markdown" title="VS Codeエディタ本体のインストール方法｜VS Codeエディタ入門" link="https://zenn.dev/karaage0703/books/80b6999d429abc8051bb/viewer/5b814b" >}}
```text
$ curl https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor > microsoft.gpg
$ sudo mv microsoft.gpg /etc/apt/trusted.gpg.d/microsoft.gpg
$ sudo sh -c 'echo "deb [arch=amd64] http://packages.microsoft.com/repos/vscode stable main" > /etc/apt/sources.list.d/vscode.list'
$ sudo apt update
$ sudo apt install -y code
```
{{< /fig-quote >}}

起動は Shell やコマンド・プロンプト等でファイル名またはディレクトリ名を指定して

```text
$ code .
```

とすれば起動する[^win1]。
Windows 版の [ATOM] はコマンド・プロンプトから任意の場所を開けなくて往生したんだよなぁ。

[^win1]: Windows 版でコマンド・プロンプトから開くには，インストール時に PATH を通す設定をする必要がある。また Windows 版ではエクスプローラのコンテキスト・メニューから開くオプションもあるので積極的に利用していいだろう。

[VS Code] 本体以外の各種ファイルは以下のディレクトリに格納される。

- Windows 版
    - `%USERPROFILE%\.vscode\`
    - `%APPDATA%\Code\`
- [Ubuntu] 版
    - `$HOME/.vscode/`
    - `$HOME/.config/Code/`

`settings.json` や `keybindings.json` といった設定ファイルやスニペットを格納する `snippets/` ディレクトリは `Code/User/` ディレクトリ直下にあるので，必要に応じてバックアップを取るのがいいだろう。

ぶっちゃけ日本語化は不要。
つか，コマンドパレットや検索フィルタを常用するなら下手に日本語化しても混乱するだけだろう。

## テレメトリの無効化

これは任意だが，ベンダ・メーカによるテレメトリを嫌うなら，最初の起動時に無効化の設定をしておく。
具体的には設定画面（`[Ctrl+,]` 押下で起動）で `telemetry` を入力し，出てきた項目のチェックを外す。

{{< fig-img src="./telemetry.png" title="telemetry" link="./telemetry.png" width="1044" >}}

## キー設定

まずはキー設定を確認しておく。
キー割当の公式情報が PDF で公開されているので，参考にするとよいだろう。

- {{< pdf-file title="Keyboard Shortcuts for Windows" link="https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf" >}}
- {{< pdf-file title="Keyboard Shortcuts for Linux" link="https://code.visualstudio.com/shortcuts/keyboard-shortcuts-windows.pdf" >}}

### コマンドパレット等

この辺は最低限おぼえておくべき。

| キー割当                   | 機能                                 |
| -------------------------- | ------------------------------------ |
| `[Ctrl+Shift+P]`<br>`[F1]` | Show All Commands (コマンドパレット) |
| `[Ctrl+,]`                 | Open Settings (UI)                   |
| `[Ctrl+K]`&nbsp;`[Ctrl+S]`        | Open Keyboard Shortcuts              |

最初から `[F1]` キーにコマンドパレットが割り当てられているのはありがたい。
まぁ，コマンドパレットさえ使えれば，あとはうろ覚えでも何とかなる（笑）

### マルチカーソル

きょうびのテキスト・エディタでマルチカーソルが使えないのはク◯だろう（下品でごめん）。

| キー割当                            | 機能                                 |
| ----------------------------------- | ------------------------------------ |
| `[Ctrl+Alt+↑]`<br>`[Shift+Alt+↑]` | Add Cursor Above                     |
| `[Ctrl+Alt+↓]`<br>`[Shift+Alt+↓]` | Add Cursor Below                     |
| `[Ctrl+D]`                          | Add Selection To Next Find Match     |
| `[Ctrl+Shift+L]`                    | Select All Occurrences of Find Match |
| `[Ctrl+F2]`                         | Change All Occurrences               |
| `[Shift+Alt+I]`                     | Add Cursors to Line Ends             |

Windows 版ではマルチカーソルの追加（上の2つの機能ね）が`[Ctrl+Alt+↑]` および `[Ctrl+Alt+↓]` にしか割り当てられてないんだけど，それやったらうちの環境ではモニタ表示の向きが変わってしまうんだよね。
うちだけの現象なのだろうか。

ちうわけで Windows 版では `[Shift+Alt+↑]`, `[Shift+Alt+↓]` に変更する。
`keybindings.json` にはこんな感じに設定されるようだ。

```json
// Place your key bindings in this file to override the defaultsauto[]
[
    {
        "key": "shift+alt+up",
        "command": "editor.action.insertCursorAbove",
        "when": "editorTextFocus"
    },
    {
        "key": "ctrl+alt+up",
        "command": "-editor.action.insertCursorAbove",
        "when": "editorTextFocus"
    },
    {
        "key": "shift+alt+down",
        "command": "editor.action.insertCursorBelow",
        "when": "editorTextFocus"
    },
    {
        "key": "ctrl+alt+down",
        "command": "-editor.action.insertCursorBelow",
        "when": "editorTextFocus"
    }
]
```

### 行の2重化

何故この機能がキーに割り当てられてないのだ！ というわけで `[Ctrl+F10]` に割り当てた。

| キー割当      | 機能                |
| ------------- | ------------------- |
| `[Ctrl+F10]` | Duplicate Selection |

`keybindings.json` にはこんな感じに設定される。

```json
// Place your key bindings in this file to override the defaultsauto[]
[
    {
        "key": "ctrl+f10",
        "command": "editor.action.duplicateSelection"
    }
]
```

### Integrated Terminal

[VS Code] には標準でターミナル機能が付いている。
ありがたや。

| キー割当         | 機能                       |
| ---------------- | -------------------------- |
| `[Ctrl+Shift+@]` | Toggle Integrated Terminal |


この機能は頻繁に使うので `[Ctrl+F1]` キーに割り当て直す。
`keybindings.json` にはこんな感じに設定された。

```json
// Place your key bindings in this file to override the defaultsauto[]
[
    {
        "key": "ctrl+f1",
        "command": "workbench.action.terminal.toggleTerminal"
    },
    {
        "key": "ctrl+shift+[BracketLeft]",
        "command": "-workbench.action.terminal.toggleTerminal"
    }
]
```

Windows 版なら [NYAGOS] を shell として使いたいものである。
そこで `settings.json` に以下の設定を直接書き込む。

```json
{
    "terminal.integrated.shell.windows": "C:\\Users\\username\\scoop\\apps\\nyagos\\current\\nyagos.exe",
}
```

（[Scoop で NYAGOS をインストール]({{< ref "/remark/2020/10/windows-terminal-and-nyagos-and-scoop.md" >}} "Windows Terminal × NYAGOS × Scoop ＝ ♥")した場合）

### Search Editor

簡単な検索と置換は `[Ctrl+F]` と `[Ctrl+H]` で可能だがファイルを跨いだ grep や grep 置換を行う場合は `[Ctrl+Shift+F]` または `[Ctrl+Shift+H]` でサイドバーを検索に切り替えた上で `[Ctrl+Shift+J]` で詳細項目を展開する。

ただサイドバーの操作ってマウス前提だし使い勝手がよくないんだよねぇ。

と思ったら設定に Search Editor という項目があって，“Search: Mode” 項目を変更することで  `[Ctrl+Shift+F]` キー押下時にどちらを起動するか選べるようだ。
これを `newEditor` に変更したら `settings.json` にはこんな感じに設定された。

```json
{
    "search.mode": "newEditor"
}
```

これでサイドバーではなくエディタのタブとして検索画面が開く。
検索結果から検索元のファイル：行にジャンプするには `[F12]` キー押下でOK（ソースコードのシンボル定義元へのジャンプと同じ）。

Grep 置換機能には該当の項目がなくサイドバーから行うしかないようだ。
まぁ，ファイルを跨いでの一括置換処理は滅多にしないけどな（笑）

## カラー・テーマは... 入れなくてもいいか

歳をとるとだんだん目が弱くなってくる。
Windows にせよ Ubuntu にせよ，パソコンのデスクトップ画面は基本的にダーク・テーマにしているのだが， [VS Code] もこれに合わせて自動的にダークテーマになるようだ。
これならカラー・テーマを別途入れる必要はないか。

ただし標準のままではカーソル行が分かりにくいので `settings.json` に以下の設定を直接書き込む。

```json
{
    "workbench.colorCustomizations": {
        "editor.lineHighlightBackground":"#303030"
    }
}
```

よーし，うむうむ，よーし。

## 必須の拡張機能

えっと，私にとって「ないと困る機能」だからね。
念のため。

ちなみに

```text
$ code --install-extension <package name>
```

とすることでも拡張機能をインストールできる。
Shell スクリプト（またはバッチ・ファイル）にまとめておけば再インストールの際にちょっとは楽になると思う。

### [EditorConfig for VS Code](https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig "EditorConfig for VS Code - Visual Studio Marketplace")

```text
$ code --install-extension EditorConfig.EditorConfig
```

コード書きなら [EditorConfig] は MUST だろう。

- [インデントおよび行末は EditorConfig で始末する](https://zenn.dev/spiegel/articles/20200922-editorconfig)

### [zenkaku](https://marketplace.visualstudio.com/items?itemName=mosapride.zenkaku "zenkaku - Visual Studio Marketplace")

```text
$ code --install-extension mosapride.zenkaku
```

全角空白を見やすく表示してくれる優れもの。
つか，滅びろ全角空白！

### [Zenkaku-Hankaku](https://marketplace.visualstudio.com/items?itemName=masakit.zenkaku-hankaku "Zenkaku-Hankaku - Visual Studio Marketplace")

```text
$ code --install-extension masakit.zenkaku-hankaku
```

全角半角変換。
日本語入力環境では必須。
というか半角カナとか全角英数とかマジで勘弁してほしい。

### [Render Line Endings](https://marketplace.visualstudio.com/items?itemName=medo64.render-crlf "Render Line Endings - Visual Studio Marketplace")

```text
$ code --install-extension medo64.render-crlf
```

改行コードの可視化。
色々あったけどユーザが多そうだったのでコレにした。
行末の余分な空白文字も目立たせてくれるスグレモノ。
つか，改行コードを表示するのに拡張機能が必要なのかよ `orz`

`settings.json` はこんな感じに設定している。

```json
{
    "code-eol.highlightExtraWhitespace": true,
    "code-eol.highlightNonDefault": true
}
```

ついでに設定の “Render Control Characters” も ON にしている。

```json
{
    "editor.renderControlCharacters": true
}
```

### [Duplicate action](https://marketplace.visualstudio.com/items?itemName=mrmlnc.vscode-duplicate "Duplicate action - Visual Studio Marketplace")

```text
$ code --install-extension mrmlnc.vscode-duplicate
```

[VS Code] のファイル・エクスプローラー機能がショボい。
特にファイルの duplicate 機能がないのは不便って思ってたら拡張機能にあった。
何故これを標準装備しないのだ！

## あるとよさげな拡張機能

### [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme "Material Icon Theme - Visual Studio Marketplace")

```text
$ code --install-extension PKief.material-icon-theme
```

よい。

### [Bracket Pair Colorizer 2](https://marketplace.visualstudio.com/items?itemName=CoenraadS.bracket-pair-colorizer-2 "Bracket Pair Colorizer 2 - Visual Studio Marketplace")

```text
$ code --install-extension CoenraadS.bracket-pair-colorizer-2
```

括弧の対応を色付きで見やすくしてくれる。
ホンマに見やすいな。
2 のほうが出来がいいらしい。

### [Git Graph](https://marketplace.visualstudio.com/items?itemName=mhutchie.git-graph "Git Graph - Visual Studio Marketplace")

```text
$ code --install-extension mhutchie.git-graph
```

`git commit` などの基本的なコマンドはコマンドパレットから簡単に呼び出せるが， GUI で操作したいときもあるので。
しかも以下のオプションを付ければ電子署名の検証も表示してくれる。

```json
{
    "git-graph.repository.commits.showSignatureStatus": true
}
```

というわけで採用。
つか，これがあれば他の git GUI ツール要らなくね？

キーボードで操作しやすいよう `[Shift+F1]` キー押下で表示するようにした。

```json
// Place your key bindings in this file to override the defaultss
[
    {
        "key": "shift+f1",
        "command": "git-graph.view"
    }
]
```

### [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens "GitLens — Git supercharged - Visual Studio Marketplace")

```text
$ code --install-extension eamodio.gitlens
```

多分，チーム運用で相互レビューするときなんかには重宝するんだろう。
かなり詳細な情報が見れるんだけど，私個人にはちょっと過剰な機能なんだよなぁ。
とりあえず保留。

### [Excel Viewer](https://marketplace.visualstudio.com/items?itemName=GrapeCity.gc-excelviewer "Excel Viewer - Visual Studio Marketplace")

```text
$ code --install-extension GrapeCity.gc-excelviewer
```

GrapeCity が公開してるのか。
どーりで（笑）

仕事で使うならよさげだけど，個人レベルじゃ特に要らないかな。

- [【レビュー】「Visual Studio Code」でExcelスプレッドシートやCSVデータを表示「Excel Viewer」 - 窓の杜](https://forest.watch.impress.co.jp/docs/review/1046119.html)

## 現在日時を設定するスニペット

現在日時をセットする方法はいくつかあるようだが，スニペットを使うのがお手軽な感じである。

`Code/User/` ディレクトリ直下に `snippets/datetime.code-snippets` というファイルを作って[^snpt1]，以下の内容をセットすればOK。

[^snpt1]: 汎用で使うスニペットを定義する場合，拡張子が `.code-snippets` なファイルであればいいらしい。

```json
{
    "Today": {
        "prefix": ["today"],
        "body": ["$CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE"],
        "description": "Today (RFC3339)"
    },
    "Now": {
        "prefix": ["now"],
        "body": ["$CURRENT_HOUR:$CURRENT_MINUTE:$CURRENT_SECOND"],
        "description": "Now time (local time)"
    },
    "TodayFull": {
        "prefix": ["todaytime"],
        "body": ["$CURRENT_YEAR-$CURRENT_MONTH-${CURRENT_DATE}T$CURRENT_HOUR:$CURRENT_MINUTE:$CURRENT_SECOND+09:00"],
        "description": "Today and time (RFC3339)"
    }
}
```

これで `today`, `now`, `todaytime` の補完候補として表示される。
自動で表示されない場合は `[Ctrl+space]` で候補一覧が出る。

単純な置換ならスニペットのほうがお手軽なので積極的に使っていきたいところである。

- [Visual Studio Code スニペットで簡単日付入力 - はんなりと、ゆるやかに](https://iucstscui.hatenablog.com/entry/2021/01/13/080000)
- [VsCodeのスニペットのススメ - Qiita](https://qiita.com/xx2xyyy/items/fd333368db548167f15a)
- [Snippets in Visual Studio Code](https://code.visualstudio.com/docs/editor/userdefinedsnippets)

## ブックマーク

- [VS Codeエディタ入門](https://zenn.dev/karaage0703/books/80b6999d429abc8051bb)
- [【随時更新】使ってるVSCodeの拡張機能のまとめ](https://zenn.dev/catnose99/scraps/36c04be9fb1209)
- [Visual Studio Codeのうれしい機能を使いこなして、初心者を最速で脱出する！《VSCode実践入門》 - エンジニアHub｜若手Webエンジニアのキャリアを考える！](https://eh-career.com/engineerhub/entry/2019/06/21/103000)
- [VSCode に入れている拡張機能 2020年版 – 未来をデザインするマーケティング会社 -ハイロックス](https://hirocks.jp/vscode-%e3%81%ab%e5%85%a5%e3%82%8c%e3%81%a6%e3%81%84%e3%82%8b%e6%8b%a1%e5%bc%b5%e6%a9%9f%e8%83%bd-2020%e5%b9%b4%e7%89%88/)
- [VSCode使い必見！？使って便利な Visual Studio Code 拡張機能10選 | ソフトウェア開発のギークフィード](https://www.geekfeed.co.jp/geekblog/vscode_extension)
- [Visual Studio Codeで現在の日時を入力するショートカットを設定する方法 - JavaScript勉強会](https://jsstudy.hatenablog.com/entry/How-to-set-a-shortcut-to-enter-the-current-date-and-time-in-Visual-Studio-Code)

[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."
[ATOM]: https://atom.io/
[Go]: https://golang.org/ "The Go Programming Language"
[go-plus]: https://atom.io/packages/go-plus
[gopls]: https://pkg.go.dev/golang.org/x/tools/gopls "gopls · pkg.go.dev"
[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
[EditorConfig]: https://editorconfig.org/

## 参考図書

{{% review-paapi "B08CZ2C3NZ" %}} <!-- Software Design (2020年8月号) -->
