+++
date = "2015-11-07T22:01:10+09:00"
title = "gibo による .gitignore 生成【改訂版】"
description = "【2020-09-21】 元の内容から全面的に改訂しました。なお，同じものを Zenn でも公開しています。"
tags = ["git", "tools"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-box type="markdown" >}}
**【2020-09-21 更新】**
元の内容から全面的に改訂しました。
なお，同じものを [Zenn でも公開](https://zenn.dev/spiegel/articles/20200921-gitignore-boilerplates "gibo による .gitignore 生成 | Zenn")しています。
{{< /div-box >}}

Git のリポジトリを作る際に `.gitignore` の内容をどうするかは悩みどころだが（大抵は他所からコピってくるのだが）， `gibo` というツールを使えば `.gitignore` のひな型をいい感じに生成してくれる。

- [simonwhitaker/gibo: Easy access to gitignore boilerplates](https://github.com/simonwhitaker/gibo)

## インストールと準備

導入方法は `README.md` に書かれているのでそちらを参考にすればいいのだが、最低限必要なのは `gibo` または `gibo.bat` スクリプトのみなので，どうにでもなるだろう。

とりあえず件のスクリプトを PATH が通るディレクトリに放り込んで `gibo help` を起動してみる（`-h` オプションでも可）。

```text
$ gibo help
gibo 2.2.4 by Simon Whitaker <sw@netcetera.org>
https://github.com/simonwhitaker/gibo

Fetches gitignore boilerplates from https://github.com/github/gitignore

Usage:
    gibo [command]

Example:
    gibo dump Swift Xcode >> .gitignore

Commands:
    dump BOILERPLATE...   Write boilerplate(s) to STDOUT
    help                  Display this help text
    list                  List available boilerplates
    root                  Show the directory where gibo stores its boilerplates
    search STR            Search for boilerplates with STR in the name
    update                Update list of available boilerplates
    version               Display current script version
```

準備として `.gitignore` 情報のアップデートから。

```text
$ gibo update
Cloning https://github.com/github/gitignore.git to /home/username/.gitignore-boilerplates
Cloning into '/home/username/.gitignore-boilerplates'...
...
```

これで `$HOME/.gitignore-boilerplates` ディレクトリに情報がセットされた。ちなみに Windows では `%APPDATA%\.gitignore-boilerplates` フォルダにセットされる。また `GIBO_BOILERPLATES` 環境変数が指定されていれば，この環境変数が指すディレクトリにセットされる。

## .gitignore 情報の取得

`.gitignore` 情報を取得するには `gibo dump` コマンドを使う。たとえば Go 言語なら

```text
$ gibo dump go
### https://raw.github.com/github/gitignore/218a941be92679ce67d0484547e3e142b2f5f6f0/Go.gitignore

# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/
```

となる。

プログラミング言語以外にもプラットフォームでの指定もできる。たとえば `go` と `vim` を両方指定して

```text
$ gibo dump go vim
### https://raw.github.com/github/gitignore/218a941be92679ce67d0484547e3e142b2f5f6f0/Go.gitignore

# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with `go test -c`
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/


### https://raw.github.com/github/gitignore/218a941be92679ce67d0484547e3e142b2f5f6f0/Global/Vim.gitignore

# Swap
[._]*.s[a-v][a-z]
!*.svg  # comment out if you don't need vector files
[._]*.sw[a-p]
[._]s[a-rt-v][a-z]
[._]ss[a-gi-z]
[._]sw[a-p]

# Session
Session.vim
Sessionx.vim

# Temporary
.netrwhist
*~
# Auto-generated tag files
tags
# Persistent undo
[._]*.un~
```

などとできる。まぁ，それぞれの情報を単純に連結してるだけなんだけどね（笑）

`.gitignore` ファイルへはリダイレクトで保存すればよい。

```text
$ gibo dump go vim > .gitignore
```

`gibo dump` コマンドで指定できる名前の一覧は `gibo list` コマンドで取得できる（めちゃめちゃ多いので動作例はパス）。

## gibo コマンドラインの補完

[simonwhitaker/gibo](https://github.com/simonwhitaker/gibo "simonwhitaker/gibo: Easy access to gitignore boilerplates") リポジトリの `shell-completions` ディレクトリ内のスクリプトファイルを使って bash, zsh, fish 上で `gibo` コマンドラインの補完機能を追加できる。

たとえば bash であれば `gibo-completion.bash` ファイルを `/etc/bash_completion.d` ディレクトリに放り込んでおけば OK。

## 参考

- [gemignore](https://rubygems.org/gems/gemignore "gemignore | RubyGems.org | your community gem host") : Ruby/Gem にも同様のツールがあるようだ
    - [gitignoreの雛形を用意する - Qiita](http://qiita.com/nakaken0629/items/cd25b722d9eb15b4efcb)
- [bash-completionでserviceコマンドなどの補完を強化しよう - インフラエンジニアway - Powered by HEARTBEATS](https://heartbeats.jp/hbblog/2013/06/bash-completion.html)

[github/gitignore]: https://github.com/github/gitignore "github/gitignore"
