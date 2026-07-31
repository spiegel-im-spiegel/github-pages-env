# 最近の私的 Golang 開発環境

あらかじめ予防線を張っておくと [Go 言語]の開発環境で「これ！」という正解はない。特にチームで開発している場合は，チームの流儀に従うのが最善だと思っている。なので，この記事は「こういうやり方もあるよ」という参考程度に見ていただけるとありがたい。

## GOPATH の構造

皆さん御存知の通り，環境変数 `GOPATH` は [Go 言語]パッケージや開発環境を指定するものだが，実は複数のパスを指定できる。Windows 環境ならこんな感じにセミコロン（`;`）で区切って指定する。

```
SET GOPATH=C:\path\to\go-packages;C:\go-wkspace
```

Linux 等ならこんな感じにコロン（`:`）で区切ればいいだろうか。

```
export GOPATH=/usr/local/go-packages:~/go-wkspace
```

`GOPATH` に複数のパスを指定した場合， `go get` コマンドで導入した（依存関係を含む）パッケージ群は最初に指定したパスの配下に格納される。上の例なら `go-packages` フォルダがそれに該当する。

今回はこれを利用する。

## [Go 言語]ツールチェーンを分離する

つまり開発に必要で `go get` コマンドで導入するツール群（ここでは仮に「[Go 言語]ツールチェーン」と呼ぶことにする）は `go-packages` フォルダに集約させ，自身の開発環境は `go-wkspace` フォルダ配下で行うようにするのである。もちろん導入した [Go 言語]ツールチェーンに `PATH` を通しておくのも忘れずに。

```
SET PATH=%PATH%;C:\path\to\go-packages\bin
```

この方法なら [Go 言語]ツールチェーン導入時に依存関係で導入されるパッケージ群によって開発環境が汚染されるのを防ぐことができる。

## 開発で使う外部パッケージは [dep] で管理する

じゃあ自身が作るパッケージが依存している外部パッケージはどうやって管理するのかというと， `go get` コマンドではなく， [dep] を使う。 [dep] で管理するパッケージを追加するには，以下のようにすればよい。

```
$ dep ensure -add github.com/pkg/errors
```

もしくは `Gopkg.toml` ファイルを直接いじって `dep ensure` で更新するか。

```toml:Gopkg.toml
[[constraint]]
  name = "github.com/pkg/errors"
  version = "^0.8.0"
```

これで外部パッケージは `vendor` フォルダ以下に一括管理できる。

## 利点と欠点

このやり方の利点は「プロジェクトごとに環境を切り替えなくていい」ということに尽きる。自身の開発環境は（上の例で言うところの） `go-wkspace` フォルダ配下に集約されるため「[GOPATH 汚染](https://text.baldanders.info/golang/gopath-pollution/ "GOPATH 汚染問題 — プログラミング言語 Go | text.Baldanders.info")」みたいなことは一切気にしなくてよくなる。

また開発環境の管理に使用するツールとその手順が少ないのも利点に挙げていいかもしれない。 [dep] は事実上の公式ツールなので，余程のことがない限りなくなることはないだろう（go コマンドに組み込まれることはあるかもしれないがw）。

欠点は利点の裏返しで，自身の開発リポジトリがどうしてもモノリシックな構成になってしまうことである。複数のプロジェクトを同時に抱えている場合は，そのプロジェクト同士の関係管理が煩雑になりがちである。

もっともモノリシックなリポジトリ構成は近年見直されつつあるようなので，運用次第では帳消しにできる欠点ではある。それに今はプロジェクト環境の分離は docker などで仮想環境ごとやっちゃうのが主流みたいだし。

## 参考

- [GOPATH 汚染問題 — プログラミング言語 Go | text.Baldanders.info](https://text.baldanders.info/golang/gopath-pollution/)
- [Glide から Dep への移行を検討する — プログラミング言語 Go | text.Baldanders.info](https://text.baldanders.info/golang/consider-switching-from-glide-to-dep/)
- [ATOM で Go — プログラミング言語 Go | text.Baldanders.info](https://text.baldanders.info/golang/golang-with-atom/) : この記事で紹介している [go-plus](https://atom.io/packages/go-plus) がバックグラウンドで自動的に [Go 言語]ツールチェーンを導入・更新するので `GOPATH` の指定を動かしづらいというのも正直に言ってあった
- [Travis CI でクロス・コンパイル（GoReleaser 編） — プログラミング言語 Go | text.Baldanders.info](https://text.baldanders.info/golang/cross-compiling-in-travis-ci-with-goreleaser/) : この記事の後半で [Travis CI](https://travis-ci.org/) で [dep] を使う方法をしれっと紹介している 
- [モノリシックなバージョン管理の利点 | プログラミング | POSTD](http://postd.cc/monorepo/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[dep]: https://github.com/golang/dep "golang/dep: Go dependency management tool"
