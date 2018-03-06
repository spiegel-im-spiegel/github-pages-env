+++
title = "バージョン間のコミット・ログを取得する"
date = "2018-03-06T18:58:09+09:00"
description = "今度忘れたときに Google 先生のお世話にならなくても済むよう覚え書きとして残しておく。"
image = "/images/attention/kitten.jpg"
tags        = [ "tools", "git", "lua", "nyagos" ]

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

[GitHub] でリリース・ドキュメントを書くことを考える。

[Go 言語]で書いたツールのバイナリ・リリース時は [GoReleaser] にお任せで全部やってくれるのだが，ソースコードやドキュメントのみのリリースの場合は自前でドキュメントを書く必要がある。
このとき `git log` コマンドでコミット・ログを参照するのだが，滅多に使わないコマンドなので毎回 Google 先生のお世話になる。

今度忘れたときに Google 先生のお世話にならなくても済むよう覚え書きとして残しておく。

## コミット・ログの収集

たとえば v0.6.0 と v0.7.0 の間のコミット・ログを markdown 形式で箇条書きにしたい場合はこうする。

```
$ git log "--pretty=format:- %s %h" v0.6.0..v0.7.0
- Merge pull request #3 from spiegel-im-spiegel/signal-subpackage 01a70c3
- Update Document 3fe7b80
- Add signal subpackage cfff012
```

この出力から “`Merge`” とか “`typo`” とかいった単語を含む行を除きたければ grep と組み合わせればよい。

```
$ git log "--pretty=format:- %s %h" v0.6.0..v0.7.0 | grep -v Merge
- Update Document 3fe7b80
- Add signal subpackage cfff012
```

なお pretty format に使える `%s` などのプレースホルダ等については以下が参考になる。

- [Git - pretty-formats Documentation](https://git-scm.com/docs/pretty-formats)

## コマンドを [NYAGOS] の Alias として組み込む

[NYAGOS] には [Lua] で書いたコードを alias コマンドとして組み込めるという素敵な機能がある。
そこで，先ほどの `git log` コマンドを [NYAGOS] の alias として組み込んでみる。
具体的には `.nyagos` に以下の記述を追加する。

```lua
-- git log
nyagos.alias.gitlog = function(args)
    local form = "--pretty=format:- %s %h"
    if #args < 1 then
        nyagos.rawexec("git", "log", form)
    elseif #args == 1 then
        nyagos.rawexec("git", "log", form, args[1])
    elseif #args == 2 then
        nyagos.rawexec("git", "log", form, args[1]..".."..args[2])
    else
        nyagos.writerr("Usage: gitlog [[<from>] <to>]\n")
    end
end
```

これで `gitlog` コマンドができた。
実際に動かしてみる。

```
$ gitlog v0.6.0 v0.7.0
- Merge pull request #3 from spiegel-im-spiegel/signal-subpackage 01a70c3
- Update Document 3fe7b80
- Add signal subpackage cfff012

$ gitlog v0.6.0 v0.7.0 | grep -v Merge
- Update Document 3fe7b80
- Add signal subpackage cfff012
```

よーし，うむうむ，よーし。

## ブックマーク

- [git logでタグとタグの間のものだけ抽出する - Qiita](https://qiita.com/suin/items/e98cef1409b6525f9bb6)
- [Git tagとGitHub ReleasesとCHANGELOG.mdの自動化について | Web Scratch](http://efcl.info/2014/07/20/git-tag-to-release-github/)
- [Git - Book](https://git-scm.com/book/en/v2)
    - [Git - Revision Selection](https://git-scm.com/book/en/v2/Git-Tools-Revision-Selection)

- [あまり使わないけど，たまに使おうとすると忘れてる Git コマンド集]({{< relref "remark/2015/git-commands.md" >}})
- [NYAGOS で Lua]({{< relref "remark/2015/nyagos-and-lua.md" >}})

[Git]: http://git-scm.com/ "Git"
[GitHub]: https://github.com/ "GitHub"
[GoReleaser]: https://goreleaser.com/ "GoReleaser | Deliver Go binaries as fast and easily as possible"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
[Lua]: https://www.lua.org/ "The Programming Language Lua"
