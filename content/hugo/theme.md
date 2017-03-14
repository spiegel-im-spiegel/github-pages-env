+++
title = "Theme を利用する"
date = "2015-09-11T17:59:18+09:00"
update = "2015-10-02T07:56:00+09:00"
description = "前々回，前回とちょっと飛ばしすぎたので今回は軽めに theme の話題を。"
tags = [ "hugo", "theme" ]
draft = false

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

[前々回]，[前回]とちょっと飛ばしすぎたので今回は軽めに theme の話題を。

## Hugo の Theme{#overview}

[Hugo] には theme 機能がある。
Theme 機能の利点はもちろん画面の構成等を共有できる点にある。
また複数人で作業している場合は theme として構成を管理することで作業環境を統一できる利点もある。

[Hugo] の repository には既に多くの theme が寄せられている。

- [spf13/hugoThemes](https://github.com/spf13/hugoThemes)

また repository に登録されているもの以外にも色々なテーマが公開されているようである。

## Theme をインストールして利用する{#install}

Theme を利用するには作業環境上に repository を `git clone` する。

```
C:\hugo-env\www>git clone --recursive https://github.com/spf13/hugoThemes.git themes
```

この repository ではユーザが公開している theme を git の submodule として結合しているため， `--recursive` オプションを付けて clone する。
現時点（2015年9月）で56個の theme が登録されているようである。

また単独の theme をインストールする場合は，作業環境に `themes` フォルダを作り，その下に theme の repository を clone する

```
C:\hugo-env\www>git clone https://github.com/spiegel-im-spiegel/hugo-theme-text.git themes/hugo-theme-text
Cloning into 'themes/hugo-theme-text'...
remote: Counting objects: 174, done.
remote: Compressing objects: 100% (133/133), done.
Receiving objecemote: Total 174 (delta 82), reused 123 (delta 35), pack-reused 0
Receiving objects:  49%
Receiving objects: 100% (174/174), 25.34 KiB | 0 bytes/s, done.
Resolving deltas: 100% (82/82), done.
Checking connectivity... done.
```

Theme を使ってビルドする場合には `-t` または `--theme` オプションを使って theme 名を指定する。
Theme 名は `themes` フォルダ直下のフォルダ名である。

```
C:\hugo-env\www>hugo --theme="html5"
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 72 ms
```

また `config.toml` に theme 名を記述する方法もある（ただしコマンドライン・オプションのほうが優先される）。

```toml:config.toml
baseurl = "http://hello.example.com/"
languageCode = "ja"
title = "Hello World!"
theme = "html5"
```

## Theme を作成する{#make}

もちろん Theme は自身で作成することもできる。
作成するには `hugo new theme` コマンドを使う。

```
C:\hugo-env\www>hugo new theme sample-theme

C:\hugo-env\www>tree /f .
C:\HUGO-ENV\WWW
│  .editorconfig
│  .gitignore
│  config.toml
│  README.md
│
├─archetypes
├─content
│  └─practice
│          hello.md
│
├─data
├─layouts
│  │  404.html
│  │  index.html
│  │
│  ├─practice
│  │      single.html
│  │
│  ├─section
│  │      practice.html
│  │
│  └─_default
│          list.html
│          single.html
│
├─public
│  │  404.html
│  │  index.html
│  │  index.xml
│  │  sitemap.xml
│  │
│  ├─categories
│  │  └─hugo
│  │          index.html
│  │          index.xml
│  │
│  ├─css
│  │      reset.css
│  │      text.css
│  │
│  ├─practice
│  │  │  index.html
│  │  │  index.xml
│  │  │
│  │  └─hello
│  │          index.html
│  │
│  └─tags
│      ├─hello
│      │      index.html
│      │      index.xml
│      │
│      └─world
│              index.html
│              index.xml
│
├─static
└─themes
    └─sample-theme
        │  LICENSE.md
        │  theme.toml
        │
        ├─archetypes
        │      default.md
        │
        ├─layouts
        │  │  index.html
        │  │
        │  ├─partials
        │  │      footer.html
        │  │      header.html
        │  │
        │  └─_default
        │          list.html
        │          single.html
        │
        └─static
            ├─css
            └─js

C:\hugo-env\www>hugo --theme="sample-theme"
0 draft content
0 future content
1 pages created
0 paginator pages created
2 tags created
1 categories created
in 37 ms
```

いくつかファイルが作成されているが `LICENSE.md` と `theme.toml` 以外は空のファイルである。

`LICENSE.md` はライセンスファイルで MIT ライセンスの雛形が出力されている。
そのまま使うのであればライセンス発行者の名前を記述する。
他のライセンスを使うのであれば他のファイルに差し替えるなり削除するなりすればよい。

`theme.toml` には theme の情報を記述する。
初期値は以下のようになっている。

```toml:themes/sample-theme/theme.toml
# theme.toml template for a Hugo theme
# See https://github.com/spf13/hugoThemes#themetoml for an example

name = "Sample Theme"
license = "MIT"
licenselink = "https://github.com/yourname/yourtheme/blob/master/LICENSE.md"
description = ""
homepage = "http://siteforthistheme.com/"
tags = ["", ""]
features = ["", ""]
min_version = 0.14

[author]
  name = ""
  homepage = ""

# If porting an existing theme
[original]
  name = ""
  homepage = ""
  repo = ""
```

これに必要な項目を書き込めばいい。

## Theme をカスタマイズする{#customaize}

[前々回]を読んだ方はピンと来ると思うが， theme のフォルダ構成（`archetypes`, `layouts`, `static`）は作業環境のそれとほぼ同じである。
実は theme の構成は作業環境で上書きできる。
従って theme を自サイト用にカスタマイズするのであれば， `themes` フォルダ内は触らないようにして，作業環境側で上書きするのがスマートである。

## 拙作 Theme “hugo-theme-text”{#mytheme}

今回，このサイトを構築するにあたり，シンプルなテキスト・サイト用の theme を作成・公開した。

- [spiegel-im-spiegel/hugo-theme-text](https://github.com/spiegel-im-spiegel/hugo-theme-text)

次回からはこの theme をベースに説明をしていきたいと思う。

## ブックマーク{#bookmark}

- [Hugoサイト構築 | Watanabe-DENKI Inc. 渡辺電気株式会社](http://wdkk.co.jp/lab/hugo/) : お勧め！

[Hugo に関するブックマークはこちら]({{< ref "hugo/bookmark.md" >}})。

[Hugo]: https://gohugo.io/ "Hugo :: A fast and modern static website engine"
[前々回]: {{< ref "hugo/hello.md" >}} "インストールから Hello World まで"
[前回]: {{< ref "hugo/section.md" >}} "Categories, Tags そして Section"
