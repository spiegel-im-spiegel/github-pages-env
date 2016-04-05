# ATOM Editor をそろそろ始めようか

私は20年来の「[秀丸](http://hide.maruo.co.jp/software/hidemaru.html)」ユーザでかつ重度の秀丸依存症なのだが，今まではともかく，これからも「Windows で秀丸」というわけには恐らくいかないので，マルチプラットフォームで使えるエディタを探している。

いくつか試してみたが，やはり [ATOM] Editor が相当だろう，と思うようになった[^atm]。 [Visual Studio Code は MS 製品にしては面白い](http://qiita.com/spiegel-im-spiegel/items/76f5965a19c178685b83)と思うし， .NET 5 のリリース時期あたりにはかなり良くなってる可能性もあるのだが，まだできないことのほうが多い。 Sublime Text は（Windows 版はあるけど）なんとなく Mac 用という印象。

そこで， [ATOM] Editor を実用性の観点から評価してみる。観点は以下のとおり。

1. Windows 環境でのインストールや運用の簡易さ
1. できるだけ少ない機能拡張（基本機能でどこまでできるか）
1. 秀丸との比較

[^atm]: [前にも書いた](http://qiita.com/spiegel-im-spiegel/items/76f5965a19c178685b83)が，テキストエディタは機能云々以前に手に馴染むかどうかが絶対的に重要。そのへん分かってない人がライセンスがどうのとか煩すぎる（有料ライセンスにはちゃんとお金を払いましょう。お金払えば気兼ねなく使えるんだから）。秀丸依存症の私が秀丸の劣化コピーでしかないサクラエディタに乗り換えることはないし， GUI でマウスも使えるのが当たり前の環境で（VT 端末を使ってた頃のトラウマがフラッシュバックしそうな） vi/vim に戻る理由もない。

## ATOM Editor のインストール

- [ATOM]

Windows 版は，インストールパッケージ（実行形式）をダウンロードしてインストールすればよい。

2015年6月25日に 1.0 が出た。[おめでたう！](http://qiita.com/spiegel-im-spiegel/items/5c6dafcece9e7118877a) （2016年4月1日に 1.6.2 stable がリリース）

[Beta Channel](https://atom.io/beta) が登場。

- [Introducing the Atom Beta Channel](http://blog.atom.io/2015/10/21/introducing-the-atom-beta-channel.html)
- [オープンソースのテキストエディター「Atom」にベータチャンネルが登場 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20151022_726976.html)

### Proxy 設定

Firewall/Proxy で囲まれている環境ではパッケージの検索・インストールがうまくいかないので， HTTPS の Proxy サーバを指定する必要がある。指定には apm コマンドを使う。

```shell
C:>apm config list -l
; cli configs
globalconfig = "C:\\Users\\username\\.atom\\.apm\\.apmrc"
user-agent = "npm/2.5.1 node/v0.10.35 win32 ia32"
userconfig = "C:\\Users\\username\\.atom\\.apmrc"

; globalconfig C:\Users\username\.atom\.apm\.apmrc
cache = "C:\\Users\\username\\.atom\\.apm"

; node bin location = C:\Users\username\AppData\Local\atom\app-1.0.0\resources\app\apm\bin\node.exe
; cwd = C:\Users\username
; HOME = C:\Users\username
; 'npm config ls -l' to show all defaults.

C:>apm config set https-proxy http://username:password@proxy.exsample.com:8080

C:>apm config list -l
; cli configs
globalconfig = "C:\\Users\\username\\.atom\\.apm\\.apmrc"
user-agent = "npm/2.5.1 node/v0.10.35 win32 ia32"
userconfig = "C:\\Users\\username\\.atom\\.apmrc"

; userconfig C:\Users\username\.atom\.apmrc
https-proxy = "http://username:password@proxy.exsample.com:8080"

; globalconfig C:\Users\username\.atom\.apm\.apmrc
cache = "C:\\Users\\username\\.atom\\.apm"

; node bin location = C:\Users\username\AppData\Local\atom\app-1.0.0\resources\app\apm\bin\node.exe
; cwd = C:\Users\username
; HOME = C:\Users\username
; 'npm config ls -l' to show all defaults.
```

多分 Windows 特有の問題だと思うけど，同じツールをあちこちにインストールさせるのは何とかならないのだろうか。 node.js は既にインストール済みやっちうねん。少なくとも PATH 上に存在してるなら存在の有無とバージョンは確認できるんだから確認しろよ。だれがバージョン管理すると思うとるんや。あと，アプリケーションごとに Proxy の設定をさせるのもウザい。 Windows の設定を使えよ。この辺 Linux や Mac の人とかどうしてるの。

参考： [プロキシで苦しむ人へ　プロキシ設定チートシート - Qiita](http://qiita.com/uzresk/items/bc7c4a9dc764390cd5ce)

## 機能の詳細

### Font Family の選択

エディタのフォントは作業効率に大きく影響する。これは私個人の感覚だが，日本語の地の文章がゴシック体なのは辛い。いくら綺麗だからといって Meiryo フォントをエディタに使う気にはならない。ただし，コードに関しては視認性が一番重要。ということで， “Font Family” の項目に以下を指定してみる。

```
Inconsolata, "MS Mincho"
```

[Inconsolata] は何かと便利なので OpenType フォントを取ってきて「インストール」してしまえばよい。「[游ゴシック 游明朝フォントパック](https://www.microsoft.com/ja-jp/download/details.aspx?id=49116)」が使える場合にはこれを使う手もある。この場合は

```
Inconsolata, "Yu Mincho"
```

とすればよい。ちなみに IPA 明朝は線が細すぎて不向きだった（まぁ印刷用に特化したフォントだからね。 IPA は Web Font 用の IPA 明朝/ゴシックフォントを開発すべき）。

Windows では Tree View のフォントが汚いので、ここは素直に Meiryo UI フォントに変える。 `%USERPROFILE%\.atom\styles.less` を以下のように変更する。

```css:styles.less
.tree-view {
	font-family: "Meiryo UI";
}
```

また

```css:styles.less
atom-workspace {
  font-family: "Meiryo UI";
}
```

とすればタブや Settings 画面のフォントも変えられる。

[Inconsolata]: http://www.levien.com/type/myfonts/inconsolata.html "Inconsolata"

### Soft Wrap が CJK に対応した

[ATOM] 1.2 から Soft Wrap が CJK に対応した。ので [japanese-wrap](https://atom.io/packages/japanese-wrap) はなくても大丈夫になった。なお， [japanese-wrap](https://atom.io/packages/japanese-wrap) と同じ作者による [wrap-style](https://atom.io/packages/wrap-style) では wrap style を細かく制御できる。お勧めである。

- [ATOM 1.2からCJKのSoft Wrapが公式対応されたのでjapanese-wrapのパッケージが必要なくなりました。 - Qiita](http://qiita.com/naru0504/items/fffcc9d5c8e9ef9d1146)
- [Atomで上手にwarpを刻んでくれるwrap-styleを開発しました。 - Qiita](http://qiita.com/raccy/items/4678af4020189366a297)

Markdown ファイルの場合は Soft Wrap を無効にしても， Preferred Line Length を 1024 とか大きい値にしても，アプリケーションの Window 幅で折り返してしまう（プレビューで折り返しが必要だから？）。

### EditorConfig のススメ

- [EditorConfig]

[EditorConfig] は今やコード書きには必須の機能だが，それ以外の人にも是非使って欲しい。人によって手持ちのプラットフォームが異なるのはもはや当たり前なので， [EditorConfig] でスタイルを統一することはとても重要。 [ATOM] では以下のパッケージがある。公開してくださった方に感謝。

- [editorconfig](https://atom.io/packages/editorconfig)

ちなみに[秀丸](http://hide.maruo.co.jp/software/hidemaru.html)でもマクロで [EditorConfig] を読み込み可能。

- [EditorConfigを反映させるマクロ](http://homepage3.nifty.com/_htom/macro/EditorConfigSet.html)

私は，仕事以外の環境での [EditorConfig] 設定を以下のようにしている。

```test:.editorconfig
root = true

[*]
end_of_line = lf
charset = utf-8
indent_style = tab
indent_size = 4
trim_trailing_whitespace = false
insert_final_newline = false

[*.md]
insert_final_newline = true

[*.{js,java,go,c,cpp,tex}]
trim_trailing_whitespace = true
insert_final_newline = true

[*.{md,c,cpp,tex}]
indent_style = space
indent_size = 4
```

- [行末の空白は EditorConfig で始末しましょう - Qiita](http://qiita.com/spiegel-im-spiegel/items/a1b4d1ad2a6693ae33e4) : ファイル読み込み時の文字エンコーディングの自動判別についても紹介

[EditorConfig]: http://editorconfig.org/ "EditorConfig"

### Theme について

個人的に Dark Theme は見づらい。かといって地が真っ白なのも目に痛いので， UI Theme は “One Light”， Syntax Theme は “Solarize Light” にしている。わざわざ Theme をダウンロードする必要はなく、 Core の Theme で充分。まぁ仕事で使う場合は Theme を統一させられるかもしれないけど。

ただし弱視の人や色覚障害の人はそれなりの Theme を導入する必要があるかもしれない。

Tree View のファイルのアイコンがそっけないので， [file-icons](https://atom.io/packages/file-icons) を入れる手もある。公開してくださった方に感謝。

（追記：[atom-monokai](https://atom.io/themes/atom-monokai) Syntax Theme というのがあって， UI Theme によって Dark と Light を自動的に切り替えてくれるようだ。最近はこっちを使っている）

### Git 関連機能

GitHub が作ったという割には git 機能が貧弱なのは何故なのだろう。特に [git-plus](https://atom.io/packages/git-plus) 相当の機能が Core パッケージに入ってないのは解せぬ。というわけで以下をインストールした。公開してくださった方に感謝。

- [git-plus](https://atom.io/packages/git-plus)

[git-plus](https://atom.io/packages/git-plus) は主にコマンドパレットで操作するが（メニューの「Packages」からも辿れる），コマンドパレットの呼び出しを f1 キーに割り当てたらだいぶ使いやすくなった。

[git-control](https://atom.io/packages/git-control) は GUI で最初の頃は便利だったのだが， [git-plus] が手に馴染んできたら使わなくなった。 [git-control](https://atom.io/packages/git-control) が， submodule の処理が得意ではないっぽいのもマイナス。

### 矩形選択

矩形選択用に `%USERPROFILE%\.atom\keymap.cson` ファイルに対して以下のキー設定を行った

```text:keymap.cson
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

いずれのやり方でも，日本語混じりのテキストではうまくいかない。これは上述した行の折り返しの問題と同じと思われる。やれやれ。

### その他 既存パッケージによる操作

私がよく使うコマンドを挙げておく。（個人的な都合で一部 keybind を変更している）

| Keybind         | コマンド                          | 内容                      | 変更した keybind |
|:----------------|:----------------------------------|:--------------------------|:---|
| `f3`            | `find-and-replace:find-next`      | 次を検索                  |    |
| `shift-f3`      | `find-and-replace:find-previous`  | 前を検索                  |    |
| `ctrl-,`        | `application:show-settings`       | 設定画面                  |    |
| `ctrl-f`        | `find-and-replace:show`           | 検索と置換                |    |
| `ctrl-g`        | `go-to-line:toggle`               | 行番号を指定してジャンプ  |    |
| `ctrl-k ctrl-l` | `editor:lower-case`               | 小文字に変換              |    |
| `ctrl-k ctrl-u` | `editor:upper-case`               | 大文字に変換              |    |
| `ctrl-p`        | `fuzzy-finder:toggle-file-finder` | ファイル検索              |    |
| `alt-delete`    | `editor:delete-to-end-of-subword` | 単語を削除                | `ctrl-delete`   |
| `alt-backspace` | `editor:delete-to-beginning-of-subword` | 直前の単語を削除   | `ctrl-backspace`   |
| `alt-left`      | `editor:move-to-previous-subword-boundary` | 直前の単語区切りまで移動 | `ctrl-left`   |
| `alt-right`     | `editor:move-to-next-subword-boundary` | 直前の単語区切りまで移動 | `ctrl-right`   |
| `ctrl-shift-D`  | `editor:duplicate-lines`          | 行の二重化                | `shift-f10`   |
| `ctrl-shift-K`  | `editor:delete-line`              | 行を削除                  |    |
| `ctrl-shift-F`  | `project-find:show`               | grep 相当のファイル検索   |    |
| `ctrl-shift-P`  | `command-palette:toggle`          | コマンドパレット          | `f1`   |

もともと `ctrl-delete` に割り当てられている `editor:delete-to-end-of-word` は挙動が大雑把すぎて使いづらいので，（もともと `alt-delete` に割り当てられてる）`editor:delete-to-end-of-subword` に機能を振りなおした。同様に `ctrl-backspace`, `ctrl-left`, `ctrl-right` も `alt-*` の機能を振りなおした。いや， alt キーって押しにくいのよ。

最終的なキー設定は [Gist に貼り付けた](https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c)ので，よろしかったら参考にどうぞ。

1. 自動保存を行う Core パッケージ [autosave](https://atom.io/packages/autosave) は既定で無効になってる。私は [git-plus](https://atom.io/packages/git-plus) で commit する前に保存するのをどうしても忘れるので，これを有効にしておくと吉。
1. Core パッケージ [tabs](https://atom.io/packages/tabs) の設定で “Enable VCS Coloring” を有効にするとタブに表示されているファイル名が色分けされて表示される。ブラボー！
1. Core パッケージ [git-diff](https://atom.io/packages/git-diff) の設定で “Show Icons In Editor Gutter” を有効にすると差分情報がちょっとだけ見やすくなる。
1. インデントをそろえるために [highlight-column](https://atom.io/packages/highlight-column) を紹介しているページを見かけたけど，インデントをそろえる目的なら，標準の Settings（ctrl-,）で “Show Indent Guide” にチェックを入れておく方が吉。

#### not([mini]) ?

最近の [ATOM] で、以前は `atom-text-editor` カテゴリだった keybind 設定が `atom-text-editor:not([mini])` とかなっているものがある。 `not([mini])` って何？ と思ってググったら

- [What is the purpose of atom-text-editor:not([mini]) in keymap.cson? - support - Atom Discussion](https://discuss.atom.io/t/what-is-the-purpose-of-atom-text-editor-not-mini-in-keymap-cson/18273)

> The :not([mini]) part tells Atom not to use the keybinding on atom-text-editor elements with a mini attribute. Mini text editors are the one-line text inputs found in the settings view (and other places). So you use atom-text-editor:not([mini]) when you want a keybinding to apply only to code text editors.

なんだそうだ。 `%USERPROFILE%\.atom\keymap.cson` の設定を見直さないと。

### その他 不足機能の追加

「秀丸にあるのに ATOM にねーよ」とかいった機能を追加していく。公開してくださった方に感謝。

1. 秀丸では大変重宝したアウトライン解析。[symbols-tree-view](https://atom.io/packages/symbols-tree-view) があればアウトライン解析を行って右側のサイドバーに表示する。
1. [japan-util](https://atom.io/packages/japan-util) : 全角・半角変換機能。英数字を全角で書いたり仮名文字を半角で書いたりするバカがいるので，意外と使うのよ，これ。
1. [highlight-line](https://atom.io/packages/highlight-line) でカーソル行をハイライトにしたり下線を引いたりできる。色のカスタマイズは `%USERPROFILE%\.atom\styles.less` ファイルで行う。
1. [open-recent](https://atom.io/packages/open-recent) は最近開いたファイルやフォルダを覚えておいてくれる便利なやつ。てか，なぜこれが標準で搭載されてないのだ。
1. [show-ideographic-space](https://atom.io/packages/show-ideographic-space) : いわゆる「全角空白」を視覚化してくれる。見せ方は `%USERPROFILE%\.atom\styles.less` ファイルでカスタマイズ可能。実際には IME のプロパティでスペース・キー押下で常に「半角空白」を入力するように設定すればほとんど防げるんだどね（全角空白を入力する場合は `shift-space` 押下）。

### その他 追加したい便利な機能

これまで述べた機能以外で「これあったら便利かなぁ」という機能を挙げてみる。実際に入れるかどうかは別。

- [autocomplete-paths](https://atom.io/packages/autocomplete-paths) は結構便利だったが必須というほどではないな。
- [autoclose-html](https://atom.io/packages/autoclose-html) : HTML 入力でタグを入力すると自動的に閉じタグを補完してくれる。自動補完の機能は色々あって，大抵は [autocomplete-plus](https://atom.io/packages/autocomplete-plus) のサブパッケージだったりするのだが，これだけは毛色が違う（笑）
- [linter](https://atom.io/packages/linter) : lint ツール。多くの言語に対応しているが，実際に動かすには lint 本体を導入する必要あり。
- [merge-conflicts](https://atom.io/packages/merge-conflicts) : conflict 時のヘルパーツール。
- [term2](https://atom.io/packages/term2) : タブ内にターミナルを起動できる。内部で `pty.js` をビルドしようとするが， Windows 環境では Visual Studio 等の開発環境が必要らしい？
    - [Atomにterm 2パッケージを導入してみた【Windows】 - Qiita](http://qiita.com/Ted-HM/items/540a57cc2a14346e4767)
- [web-view](https://atom.io/packages/web-view) : タブ内にブラウザを起動できる。
- スクラッチパッドみたいなんないのん？ と思って探していたのだが， [tempfile](https://atom.io/packages/tempfile) というパッケージが登場した。ファイルタイプを指定できる。またファイルの保存先を指定できるため，適当に書き散らしたものを保存することも可能。
- 特定の単語をマークしておける [quick-highlight](https://atom.io/packages/quick-highlight) が結構使える。 Windows だとキーに割り当てられないので、 toggle を適当なファンクションキーとかに割り当てておくとめっさ便利。
- [tablr](https://atom.io/packages/tablr) : CSV Editor。なにこれ素敵！

## 継続調査（欲しい機能）

順次更新予定。

- grep。 `project-find` ではなくて普通の grep。 Windows には grep は標準ではないので，今まではエディタ内蔵の grep を使ってた。
- [結城浩さんの amazon.mac](http://d.hatena.ne.jp/hyuki/20120413/amazon) が結構便利でよく使うのだが，同等の機能のものがないだろうか。処理自体は簡単だし自前で作る手もあるが，それならちゃんとパッケージについて勉強しないとなぁ。 → [自作しました](http://text.baldanders.info/remark/2015/insert-amazon-url-with-associate-id-in-atom-editor/)
- ファイル読み込み時の文字エンコーディング自動判別は[自前で何とかできた](http://qiita.com/spiegel-im-spiegel/items/a1b4d1ad2a6693ae33e4)が，あとは文字エンコード変換かぁ → [auto-encoding](https://atom.io/packages/auto-encoding) パッケージで自動判別できるようだ

ちなみに，10万行の CSV ファイル（Shit-JIS エンコード）を読み込ませたら，文字エンコード変換中にハングアップした。たかだかそれくらいで動かなくなるとか，まだまだ秀丸は手放せないなぁ sigh...

## ブックマーク

### 設定全般（みんな設定ってどうやってるの？）

- [ATOMエディタのproxy設定 - Qiita](http://qiita.com/tsukamoto/items/cef0f2d7e2b33a26a9e5)
- [EditorConfig で簡単にエディタの設定を共有する | イントフロート スタッフブログ](http://maruta.be/intfloat_staff/164)
- [Atomで快適な生活を送る方法 - Qiita](http://qiita.com/nsatohiro/items/33d87cfb7cfca436f3d2)
- [Atom Editor 俺の設定とプラグインをさらす - Qiita](http://qiita.com/agektmr/items/4da2c362fef6598fc382)
- [Tumbling Dice — Atom やっておきたい設定＋入れておきたいパッケージメモ](http://outofmem.tumblr.com/post/102602584159/atom)
- [Atom でインストールしている package 一覧 - Qiita](http://qiita.com/iorionda/items/b108e73e51ff40049c60)
- [SublimeText から Atom に乗り換えた時にやった設定とインストールしたパッケージ | rythgs.co](http://rythgs.co/archives/2015/04/03/my-atom-settings-packages/)
- [WindowsでAtomをインストールして使ってみる - Qiita](http://qiita.com/purini-to/items/caca2a0c56e984b2f9d8)
- [Atom の package で定義されているキーバインドを変更する方法 - Qiita](http://qiita.com/iorionda/items/0f386408fb9591c0ed7f)
- [Atomで特定のキーバインドによる操作が遅い場合の対策 - Qiita](http://qiita.com/aki77/items/2d3ed9ce0de72209c283)
- [ATOM でファイルを開いたら自動文字コード判定を行う - Qiita](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)
- [Rails開発用のAtomエディタの設定メモ(2015年版) - Qiita](http://qiita.com/knt45/items/d56ac5aa04fe3fff8506)
- [atomでキーマクロ - Qiita](http://qiita.com/nbjiao/items/36ca975a882ddd0f8767)
- [Atom apmの使い方 - Qiita](http://qiita.com/akaimo/items/31c1825f7e2ac4c4fb5e)
- [【Atom】Web開発者達によるWeb開発者の為のエディターについて - Qiita](http://qiita.com/snowsunny/items/31ecb42d156b50b76376)
	- [【超おすすめ！！】Atomのパッケージ、テーマ、キーバインディング、設定を紹介してみる（※随時更新） - Qiita](http://qiita.com/snowsunny/items/f40c3291a580f3215797)
- [Atom のツリーとタブのフォントを変更する - Qiita](http://qiita.com/yuch4n/items/85bc93eb9eba38e45d74)
- [AtomでTypeScriptの環境を構築する - Qiita](http://qiita.com/Sugima/items/113eb16f14ca1e9a6d0f)
- [ブラウザと同じキーマップで、Inspect Element - Qiita](http://qiita.com/t9md/items/013db47ea5fe807ec9d2)
- [Atomエディタでデフォルトのキーバインド（ctrl-shift-K）を無効にする - Qiita](http://qiita.com/madogiwa/items/1b892221c2de92ca4187)
- [Atomのスペースをはっきり見えるようにする - Qiita](http://qiita.com/shibukk/items/e62931a7a3b6dc617623)
- [atomのmarkdown-previewで日本語が豆腐になる(Ubuntu) - Qiita](http://qiita.com/Mokkeee/items/cca4ba3a6bf334fa7934)
- [【atom】シンタックスカラーが気に入らないなら、自分で作ればいいじゃない - Qiita](http://qiita.com/nnmr/items/11642180f592a9ccaf36)
- [Atomエディタで本文１行目をタブに表示する方法 - Qiita](http://qiita.com/nobuhito/items/5219173cf179754a4113)
- [Atomエディタ ver1.0.8で前から欲しかった機能が2つ追加されていた - Qiita](http://qiita.com/knt45/items/9f88e2b54f622e27dd0d)
- [Atomエディタの背景画像を設定するのに苦労したこと【Windows】 - Qiita](http://qiita.com/Ted-HM/items/afe9eddaefc0247859a8)
- [apm loginすると`EINVAL, invalid argument`で止まる - Qiita](http://qiita.com/nureha/items/12a8526835ddc5f7dbb0)
- [Atomでvim-modeのカーソルが見辛いのを解消する - Qiita](http://qiita.com/74th/items/e9bebcdc9d979f1c6b44)
- [1つのキーバインドで、同時に複数のコマンドを実行する - Qiita](http://qiita.com/74th/items/160b011bbfcac77d4e3c)
- [atomで最終行を改行させない - Qiita](http://qiita.com/bouzuao/items/c4b6c49c13c2d1ad54f3)
- [Atomをvimっぽくするプラグイン5選 - Qiita](http://qiita.com/ikaro1192/items/d477c7153a4d3af9f209)
- [Atomのwebview内のフォントを変更する - Qiita](http://qiita.com/from_kyushu/items/3f818a80ba7e58d356dc)
- [atom markdown-preview用のスタイル - Qiita](http://qiita.com/s_k_o_i/items/7ade4a5dda86e5c92163)
- [備忘録：Atom - Qiita](http://qiita.com/nanoco/items/0156ae8aee1a773b933a)
- [Atomに本当のAutoSave機能をつける - Qiita](http://qiita.com/nobuhito/items/3e577127bde175315914) : [autosave](https://atom.io/packages/autosave) を使うんじゃなくてタイマを仕掛けて定期的に保存する方法
- [Atomで行移動(moveLineUp/moveLineDown)時のオートインデントをオフにする方法 - Qiita](http://qiita.com/snowsunny/items/e6408e5ee7806bff133a) : オートインデントの挙動がおかしい場合にどうぞ
- [Atomが”editor is not responding”になったら - Qiita](http://qiita.com/devzooiiooz/items/777da5b63064f85cd096)
- [Atom レジストリエントリを含むアンインストール方法（Windows版） - Qiita](http://qiita.com/masa36/items/c800185174bd77526a7d)
- [Atom Editor を EDITOR 環境変数に設定する - Qiita](http://qiita.com/thermes/items/a38d0e5824c787166871)

### 各種パッケージ設定

- [EditorConfigで文字コード設定を共有して喧嘩しなくなる話。（Frontrend Advent Calendar 2014 – 14日目） | Ginpen.com](http://ginpen.com/2014/12/14/editorconfig/)
- [Atom で超快適に Git を扱うためのプラグインと設定](http://www.geeks-dev.com/atom-%E3%81%A7%E5%BF%AB%E9%81%A9%E3%81%AB-git-%E3%82%92%E6%89%B1%E3%81%86%E3%81%9F%E3%82%81%E3%81%AE%E3%83%97%E3%83%A9%E3%82%B0%E3%82%A4%E3%83%B3%E3%81%A8%E8%A8%AD%E5%AE%9A/)
- [Atomの設定を共有・移行するときはsync-settingsがオススメ - Qiita](http://qiita.com/sagaraya/items/dc2240c32cae2f522613)
- [AtomでMarkdown+数式を利用する - Qiita](http://qiita.com/noppefoxwolf/items/335323b98f0400a6f07d) : `markdown-preview` を `markdown-preview-plus` で置き換える
- [AtomエディタでPHP(＋HTML)がシンタックスハイライトされないのを解消 - Qiita](http://qiita.com/kijtra/items/4247fca754b1d8e89c62) : `file-types` パッケージを使う
	- [Atomで拡張子ごとのハイライトを追加/変更する - Qiita](http://qiita.com/verytired/items/60635ba9b48d8ffb2308)
- [ATOMで開き括弧を入力した時に、閉じ括弧を自動で入力するのをやめる - Qiita](http://qiita.com/74th/items/8d79e06c7099a73f3eea) : Core パッケージの `bracket-matcher` の設定を無効にする
- [atomをevernoteのmarkdownエディタとして使う（EVND） - Qiita](http://qiita.com/kuwa_tw/items/16d1e63b66092601391e)
- [AtomでSublimeTextみたいに保存しないで終了しても再開できるやつ - Qiita](http://qiita.com/ru_shalm/items/453bb63c3fa54cb93f61) : [save-session](https://atom.io/packages/save-session) があるじゃない。という話
- [atom から gist にコードをアップする gist-it を試す - belog](http://kkabetani.hatenablog.com/entry/2014/12/29/235147)
- [【atom】atomでgistを作るには - Qiita](http://qiita.com/Mocacamo/items/42421ae5cf82b1bf6cb3)
- [Atomのパッケージが急に動かなくなった！ - Qiita](http://qiita.com/wh11e7rue/items/545c0a26c609cec9487d) : パッケージの修復方法
- [Atomに現在日時を挿入するコマンドを追加する - Qiita](http://qiita.com/toruot/items/b26fde1a898bb52985e1)
    - [日付を挿入するキーマップをAtomに追加 | Jun Nishii](http://bcl.sci.yamaguchi-u.ac.jp/~jun/ja/blog/150221-insert_date_keymap_to_atom)
- [ATOMでながら作業が捗るパッケージ紹介 - Qiita](http://qiita.com/nekobato/items/5b1e4e2f4f328466cc1d)
- [Atom でリモートのファイルを操作するプラグイン remote-ftp | Lonely Mobiler](http://loumo.jp/wp/archive/20151004000041/) : [remote-ftp](https://atom.io/packages/remote-ftp) の使い方 FTP/FTPS/SFTP が使えるようだ。 Windows 環境では pageant が使えるらしい
	- [Atom Package「Remote-ftp」のローカルでのあれこれ - Qiita](http://qiita.com/makoto1007/items/3fb628796a880ddbea15)
- [Atom v1.2.0以上 で term2 v0.9.21 がビルドエラーで動かなくなる件 - Qiita](http://qiita.com/kaminaly/items/853dc9c40ae7433cf4c1) : 結論は [term3](https://atom.io/packages/term3) を使えということらしい
- [こんなのがあるのはAtomだけ？尖りまくってるパッケージ紹介【随時（たまーに）更新】 - Qiita](http://qiita.com/snowsunny/items/17f35baf0671451d2c09)
- [Atom でエンコーディングを自動判別するパッケージ auto-encoding - Qiita](http://qiita.com/enk/items/59fdcc4409e6eb650952)
- [Atomのウィンドウを半透明にする方法 - Qiita](http://qiita.com/kakinoki/items/37bfee713990f1c6b43c)
- [Windows 10 で Atom の Term3 を入れる(たぶん)最小の方法 - Qiita](http://qiita.com/raccy/items/86bfd002d713b7c4d032)
- [Atomのscriptパッケージでbundle exec付きで実行する - Qiita](http://qiita.com/Kesin11/items/75b400e45220124e1569)
- [JSON・CSVからピボットテーブルを作ってくれるAtomパッケージ - Qiita](http://qiita.com/takeyuichi/items/1ea77ce7f3d0eeb1c206)

### パッケージを作ってみた

- [AtomでローカルのPackageをインストールする方法 - Qiita](http://qiita.com/sagaraya/items/11b4f9fbec7a4fcc4439)
- [toggle つくった - Qiita](http://qiita.com/t9md/items/a77387e08a3305ba2862)
- [atomで単語削除を拡張するパッケージを作った - Qiita](http://qiita.com/pppp403/items/e05fed884eea657fa52e)
- [smalls つくった - Qiita](http://qiita.com/t9md/items/bca96d45af1a5244b5d1)
- [AtomパッケージのTips - Qiita](http://qiita.com/p-baleine@github/items/2a29e82616de3680adf3)
- [MermaidのプレビューができるAtomパッケージを作った - Qiita](http://qiita.com/takeyuichi/items/4b8ed40a598623026a9c)
	- [mermaid.jsが素晴らしいけどなかなか使ってる人見かけないので実例晒す(追記あり) - Qiita](http://qiita.com/uzuki_aoba/items/a01f8b0b52ced69c8092)
- [Package 作成時の便利フレーズ、知識など - Qiita](http://qiita.com/t9md/items/c02da839621a1050bc34)
- [Atomプラグイン作成入門 ハンズオン形式でプレビュー型プラグインを作ってみよう！ - Qiita](http://qiita.com/geek_duck/items/654919124d24fa74a0d3)
- [Atomでパッケージを作ってみた その2 - Qiita](http://qiita.com/tajihiro/items/f69315ee7006d51f2f19)
- [Atomでパッケージを作ってみた その3 - Qiita](http://qiita.com/tajihiro/items/f6e987ce49f5714e76cd)
- [Atom エディタの自作パッケージで子プロセスを呼ぶ方法 - Qiita](http://qiita.com/_meki/items/32aa956ae9d2d8ea5b26)
- [Atomでファイルを読み書きするパッケージを作る際にUTF-8以外の文字コードにも対応する方法メモ - Qiita](http://qiita.com/from_kyushu/items/851a55dd90821d23c497)
- [Riot.js向けのSyntax Highlightをプラグインにしました (Atom用) - Qiita](http://qiita.com/dekimasoon/items/9cd05ba5c95a0a8f02c8)
- [他パッケージのモジュールをロードする - Qiita](http://qiita.com/from_kyushu/items/2082f51416acd0bfab0a)
- [書き捨てファイルを作成するtempfile Packageをリリースした - Qiita](http://qiita.com/from_kyushu/items/0b31e5949da1f03bc690) : なんか上手く動かない
- [Atomで翻訳系API呼び出すパッケージを作った - Qiita](http://qiita.com/MikamiHiroaki/items/7620023e6a870ac17e90)

### テーマを作ってみた

- [Atomでthemeを自作して登録するまでの備忘録 - Qiita](http://qiita.com/inuscript/items/638ebbdde33b6d1181f1)

### ATOM で開発環境を整える

- [AtomでGoを書く環境を整える（Windows） - Technically, technophobic.](http://notchained.hatenablog.com/entry/2014/09/20/104829)
- [Haskell開発環境の構築 (Windows編) - Qiita](http://qiita.com/td2sk/items/9e4b49a4a31b7138d3ad)
- [SFDC：MavensMate for Atomを試してみました - tyoshikawa1106のブログ](http://tyoshikawa1106.hatenablog.com/entry/2015/08/30/012917)
- [AtomでCoffeeScriptの開発環境を整える - Qiita](http://qiita.com/kyo_nanba/items/c8e5e79cb1816e8c9dba)
- [atomをリモートサーバのエディタとして使う(nuclide-server) - Qiita](http://qiita.com/hashrock/items/2668532ece72961f2cfd)
    - [Atomでリモートファイルを更新するパッケージ - Qiita](http://qiita.com/from_kyushu/items/d47536f275c3bf8d544d)
- [快適Haskell環境構築！ (ghc-modがエラーで動かない方へ) - Qiita](http://qiita.com/Asakage/items/5aebc4a53244d5adc335)
- [ATOMをR言語に対応させる方法 - Qiita](http://qiita.com/mark88232/items/bf67d1c4c76c99a9bb94)
- [【新人教育 資料】UMLまでの道 〜クラス図を書いてみよう編〜 - Qiita](http://qiita.com/devopsCoordinator/items/8d2af381c1c469103459) : ATOM でクラス図が描ける！
- [AtomとPlantUMLで爆速UMLモデリング - Qiita](http://qiita.com/nakahashi/items/3d88655f055ca6a2617c)
	- [Atom+PlantUMLで見た目もいい感じのシーケンス図を作成する - Qiita](http://qiita.com/k_nakayama/items/77ca73753ebd049a66de)
- [AtomでRailsを爆速開発する環境を作ってみた - Qiita](http://qiita.com/tacumai/items/e84e586b5bde2979a066)

### ATOM TeX 環境を整える

- [Atom - TeX Wiki](https://texwiki.texjp.org/?Atom)
- [Mac の Atom で latex - Qiita](http://qiita.com/NariseT/items/4b80c80f8e39ea35f597)

### その他 四方山話

- [Atom Editorを使う - Qiita](http://qiita.com/t_mrt/items/f939bce2f97cdee861d2) : Atom API について
- [SublimeText - Atomを使って感じた、やっぱりSublime Textが良い理由 - Qiita](http://qiita.com/nipoko/items/fa9fd1570d3b29e9f6ac)
- [GitHub、コードエディター「Atom」v1.3を公開。インストール不要のポータブル版を追加 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20151214_735200.html)
- [私がどのようにしてAtomの奇妙なバグを修正したか : 正規表現が暴走を起こすとき | プログラミング | POSTD](http://postd.cc/how-i-fixed-atom/)
- [GitHub、オープンソースのテキストエディター「Atom」の最新正式版v1.6.0を公開 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160318_749051.html) : かなり色々変わったらしい

[ATOM]: https://atom.io/ "Atom"
[Visual Studio Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined."
