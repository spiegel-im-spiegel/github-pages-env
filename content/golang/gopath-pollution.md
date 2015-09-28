+++
date = "2015-09-28T15:21:49+09:00"
description = "go get コマンドは外部パッケージの revision 等をコントロールできず，常に repogitory の最新コードを取ってこようとする。ひとつの環境でひとつのプロジェクトを管理していくのならこれでも何とかならないこともないが，複数のプロジェクトが同居している場合は同じ外部パッケージでも異なるリビジョンを要求する場合があり，管理が煩雑になってしまう。"
draft = true
tags = ["golang", "engineering", "environment", "pollution"]
title = "GOPATH 汚染問題"

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

（初出： [そろそろ真面目に Golang 開発環境について考える — GOPATH 汚染問題 - Qiita](http://qiita.com/spiegel-im-spiegel/items/73ebc684b5807277b7e2)，[そろそろ真面目に Golang 開発環境について考える — Internal Packages と Vendoring - Qiita](http://qiita.com/spiegel-im-spiegel/items/baa3671c7e1b8a6594a9)）

`go get` コマンドはとても強力な機能で，私のように Windows と UNIX 系環境の間を渡り歩いてる身としては， make などの tool chain に大きく依存することなく， `go get` コマンドだけで repository の fetch から build/install まで出来てしまうのは非常にありがたい[^a]。

[^a]: それでも git などのコード管理ツールへの依存はどうしても残るのだけれど。

`go get` コマンドは外部パッケージの revision 等をコントロールできず，常に repogitory の最新コードを取ってこようとする。
ひとつの環境でひとつのプロジェクトを管理していくのならこれでも何とかならないこともないが，複数のプロジェクトが同居している場合は同じ外部パッケージでも異なるリビジョンを要求する場合があり，管理が煩雑になってしまう。

で，困ったことに `GOPATH` 環境変数は複数のプロジェクト管理を想定していない。

`GOPATH` 環境変数には複数のパスをセットできる。 Windows 環境ならこんな感じ。

```
SET GOPATH=C:\golib;C:\workspace\project1;C:\workspace\project1;...
```

しかし，これらのパスが全て有効となるのは Go コンパイラが外部パッケージを探す場合であり， `go get` コマンドで repogitory を fetch する場合には `GOPATH` で指定する最初のパス（上述の例なら `C:\golib`）と決められている。
これでは折角プロジェクトごとにフォルダを分けても，外部パッケージは勝手にひとつのフォルダに集約されてしまうことになる。

## 【対策1】 プロジェクトごとに GOPATH を設定し直す

この問題に対する一番安直な答えは「プロジェクトごとに GOPATH を設定し直す」である。例えば[前回]紹介した [gb] をビルドする場合は以下のようにする。

```
C:>mkdir C:\workspace\gb

C:>SET GOPATH=C:\workspace\gb

C:>go get -v github.com/constabulary/gb/...
github.com/constabulary/gb (download)
github.com/constabulary/gb/log
github.com/constabulary/gb
github.com/constabulary/gb/vendor
github.com/constabulary/gb/cmd
github.com/constabulary/gb/cmd/gb
github.com/constabulary/gb/cmd/gb-vendor
```

あとは `GOPATH` 直下の `bin` フォルダにパスを通すか，パスの通ってるフォルダに実行ファイルをコピーすればよい。
実行履歴はバッチファイル（UNIX 系なら shell スクリプト）に保存しておけばいつでも復元できる。

毎回環境をセットアップしないといけないのは面倒だが，プロジェクト管理のためのツールも必要なく， Go コンパイラの標準機能のみで管理できる。
標準機能のみで管理できるというのは結構重要で，たとえば CI ツールを使っている場合は，設定を単純にできるので管理しやすいといえる。

UNIX 系の環境であれば [direnv] を使う手もある[^b]。
[direnv] は `cd` をフックし，ディレクトリごとに環境変数を書き換えることができる。
[direnv] 使うことで プロジェクト・フォルダごとに `GOPATH` を設定し直してくれる。

[^b]: [direnv] は [Go 言語]で組まれている。

### 【対策2】 プロジェクト・ベースの管理ツールを使う

もうひとつは [gb] のようなプロジェクト・ベースでコード管理のできるツールを使う方法である[^c]。
[gb] で作った開発環境はフォルダごと開発メンバに配布・同期することが可能になるため，複数人で環境を合わせることが容易になる。

[^c]: [gb] については[前回]の記事を参照のこと。

### 【対策3】 Go 1.5 の vendoring 機能を使う












## ブックマーク

- [direnvで解決するGOPATHの3つの問題点 - None is None is None](http://doloopwhile.hatenablog.com/entry/2014/06/18/010449)
- [改めて、direnvを使いましょう！ - HDE BLOG](http://blog.hde.co.jp/entry/2015/02/27/182117)
- [さくら - homeにgolang, direnv とvirtualenvを入れて動かす - Qiita](http://qiita.com/aminamid/items/5a0e9461385c80d0c8a6)

[Go 言語に関するブックマーク集はこちら]({{< ref "golang/bookmark.md" >}})。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[前回]: {{< ref "golang/project-based-development.md" >}} "プロジェクト・ベースの開発環境をつくる"
[gb]: http://getgb.io/ "gb - A project based build tool for Go"
[direnv]: http://direnv.net/ "direnv - unclutter your .profile"
