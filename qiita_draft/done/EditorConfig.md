# 行末の空白は EditorConfig で始末しましょう

最近は作業の8割くらいは [ATOM] でまかなえるようになって，だいぶ手に馴染んできた。で，相変わらず[秀丸](http://hide.maruo.co.jp/software/hidemaru.html)にあって [ATOM] にない機能を[ちょこちょこと補完](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)してるんだけど，その中で「行末の空白を削除する」パッケージとかマクロ（他のエディタね）とかが目につきだした。みんな苦労してるんだねぃ。

でも，昔はともかく，今は [EditorConfig] に対応しているエディタなら余計なパッケージやマクロなど不要である。

-  [EditorConfig]

たとえば統合環境の IntelliJ IDEA はネイティブに [EditorConfig] に対応してるし，プラグイン対応でなら Visual Studio や NetBeans も対応している。普通のエディタでも ATOM や Sublime Text はもちろん，みんな大好き（私は使わないけど） Emacs や Vim も [EditorConfig] に対応可能である。つまり [EditorConfig] を使わない手はないのである。

## EditorConfig による設定

[EditorConfig] を使うには（既にプラグイン等は導入済みであると仮定して）プロジェクト管理下のトップ・フォルダに `.editorconfig` ファイルを設置する。（[前にも紹介した](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)けど）私の場合はこんな感じ。

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

[EditorConfig] はフォルダを遡って `.editorconfig` ファイルを探し，フォルダの上から順番に評価していく。 `root = true` の記述がないとどこまでも上の階層に遡っていくので，プロジェクトのトップ・フォルダの `.editorconfig` には必ずこれを記述すること。

- `[*]` は対象となるファイルを指定している。 `[*]` なら全てのファイルが対象である。 `[*.{md,c,cpp,tex}]` は拡張子が md, c, cpp, tex のファイルが対象となる。
- `end_of_line` では改行コードを指定する。 `lf`, `cr`, `crlf` から選択できる。
- `indent_style` はインデント（タブ）のスタイルを指定する。 `tab` または `space` を指定する。 `space` にすると，いわゆる「ソフトタブ」になる。
- `indent_size` はタブの幅を指定する。 `indent_style` と `indent_size` は常にセットで指定すると間違いがない。
- `trim_trailing_whitespace` を `true` にすると行末の空白文字を削除してくれる。残念なことに [ATOM] の場合は，いわゆる「全角空白」を空白文字と見なしてくれない。まぁ全角空白を空白文字と見なす設計のほうが少ないけど。
- `insert_final_newline` を `true` にするとファイルの末尾が改行文字ではない場合に補完してくれる。でもこれって使いどころが難しいのよね。

## 文字エンコーディングについて残念なお知らせ

`charset` では文字エンコーディングを指定するが，標準では `latin1`, `utf-8`, `utf-8-bom`, `utf-16be`, `utf-16le` しかサポートしていない。それ以外の文字エンコーディングは実装依存ということになる。

しかも [ATOM] の場合，ファイルを新規作成する場合にこの設定が効かないようで，たとえば [ATOM] 側の “File Encoding” が `shiftjis` で  `.editorconfig` ファイル側が `charset = utf-8` の場合，新規作成ファイルは `shiftjis` にセットされ，そのまま保存される。しかも次にそのファイルを開く場合は（`.editorconfig` ファイル側の設定が効いてしまうので） `utf8` で開いてしまい，結果派手に文字化けする（`ctrl-shift-U` で文字エンコーディングを変更すれば元に戻るけど）。

### ファイル読み込み時に文字エンコーディングを自動判別する

新規作成時の初期の文字エンコーディングは今のところどうしようもないが，プロジェクトごとに “File Encoding” を変更して対応するか，新規作成ファイルが開いた直後に `ctrl-shift-U` で文字エンコーディングを変更することで何とかなるだろう（ダサいけど）。

もうひとつの緩和策としては，既存ファイル読み込み時に文字エンコードを自動判別するようにすることだ。ただし，この機能を持つ Package は今のところ存在しないようなので自前で何とかするしかない。

- [ATOM でファイルを開いたら自動文字コード判定を行う - Qiita](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)

この設定を行うには node.js のフルパッケージが必要。

（Linux や Mac な人は依存関係で node.js がインストールされると思うけど， Windows では [ATOM] インストール時に一部機能が同梱されているだけなので（しかもバージョンが古い），フル機能を使うには別途インストールする必要あり）

Windows の場合は `%USERPROFILE%\.atom` フォルダに移動する。その後， `npm` コマンドを使って `iconv-lite` と `jschardet` をインストールする。

```shell
C:>cd C:\Users\username\.atom
C:\Users\username\.atom>>npm install iconv-lite
iconv-lite@0.4.10 node_modules\iconv-lite

C:\Users\username\.atom>>npm install jschardet
jschardet@1.1.1 node_modules\jschardet
```

すると `%USERPROFILE%\.atom\node_modules` フォルダが作成され，その中に `iconv-lite` と `jschardet` がインストールされているはずである。

次は `%USERPROFILE%\.atom\init.coffee` ファイルに以下の記述を追加する。

```coffee:init.coffee
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

「[ATOM でファイルを開いたら自動文字コード判定を行う](http://qiita.com/tokudiro/items/bc232c7d36261dc45936)」によると，これは [encoding-selector](https://atom.io/packages/encoding-selector) からの流用らしい。ただし現在， [encoding-selector](https://atom.io/packages/encoding-selector) は Core Package に入ってるので atom フォルダをひっくり返してもソースコードは見当たらない。ので，GitHub にある [atom/encoding-selector](https://github.com/atom/encoding-selector) にある [lib/encoding-list-view.coffee](https://github.com/atom/encoding-selector/blob/master/lib/encoding-list-view.coffee) を参考にするといいだろう。 `detectEncoding:` のあたりである。

これで [ATOM] を起動して既定の文字エンコーディングでない適当なファイルを読み込ませてみれば確認できる。ただし，自動判別は万能じゃない（たまに間違う）ので，その辺は悪しからずってことで。

一番いいのは [encoding-selector](https://atom.io/packages/encoding-selector) がファイル読み込み時に自動判別する機能を付けてくれることなんだけど。誰かやらないかな。個人的には [ATOM] 開発に積極的に commit する気はないので，完全に他人任せなのだが。

あとは文字変換か...

## ブックマーク

- [EditorConfig で簡単にエディタの設定を共有する | イントフロート スタッフブログ](http://maruta.be/intfloat_staff/164)
- [EditorConfigで文字コード設定を共有して喧嘩しなくなる話。（Frontrend Advent Calendar 2014 – 14日目） | Ginpen.com](http://ginpen.com/2014/12/14/editorconfig/)
- [どんなエディタでもEditorConfigを使ってコードの統一性を高める - Qiita](http://qiita.com/naru0504/items/82f09881abaf3f4dc171)

[ATOM]: https://atom.io/ "Atom"
[EditorConfig]: http://editorconfig.org/ "EditorConfig"
