+++
title = "Windows でも Grep したい"
date = "2018-03-07T19:08:25+09:00"
update = "2018-03-08T17:06:23+09:00"
description = "既に Go 言語製の grep を公開されてる方がいた。パスの再帰検索が地味にありがたい。"
image = "/images/attention/kitten.jpg"
tags        = [ "tools", "grep", "windows" ]

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

個人的に [ATOM] Editor で不満だったのは [grep] 機能（正しくは `project-find`）がアレで微妙に使い勝手が悪いことだった。

いや， [ATOM] Editor の中で作業する分には不満はないんだよ。
でも本当は任意の場所にあるファイルを [grep] して結果をテキストで吐いて欲しかったわけ。
この点では[秀丸]の使い勝手がよかった[^hm1]。

[^hm1]: 他にも [ATOM] Editor では巨大テキストを実質的に扱えないなど微妙だが大事なところで不満があるのだが，概ね使えてる。

UNIX 系プラットフォームには当然のようにある [grep] だが Windows にはない。
[MSYS2] のようなエミュレータを使う手もあるが，やっぱり Windows でも普通に [grep] したいのである。
`findstr.exe` じゃ役不足なんだってば！

いっそ [Go 言語]で自作するか？ いやいや，「車輪の再発明」は絶対ドツボにはまる。
などと葛藤していたのだが，なんと！ 既に [Go 言語]製の [grep] を公開されてる方がいた。

- [Big Sky :: 日本語grepが出来るjvgrepというのを作った。](https://mattn.kaoriya.net/software/lang/go/20110819203649.htm)
    - [mattn/jvgrep: grep for japanese vimmer](https://github.com/mattn/jvgrep)

{{% fig-quote title="日本語grepが出来るjvgrepというのを作った。" link="https://mattn.kaoriya.net/software/lang/go/20110819203649.htm" %}}
“vimgrepの様に複数のエンコーディングに対応していて、検索パターンにマルチバイト文字を含んだ正規表現が使えて、windowsでもちゃんと動いて、ついでといっちゃあなんだが、`"**/*.txt"`で再帰検索してくれる様なgrep無いかなぁと思ってたんですが、やっぱり無いので作りました。”
{{% /fig-quote %}}

ありがたや。

最近のバージョンはバイナリを配布してないようなので，大人しく

```text
$ go get -u -v github.com/mattn/jvgrep
```

でインストールする。
もう [mattn/go-iconv] は使ってないっぽいので `iconv.dll` とかは不要なようだ。

パスの再帰検索が地味にありがたい。

## jvgrep を [NYAGOS] の Alias に組み込む

[NYAGOS] では既定の設定で [grep] を `findstr.exe` の alias として登録している。
これを `~/.nyagos` ファイルで上書きする。
以下の記述を追加して再起動。

```Lua
nyagos.alias.grep = "jvgrep.exe"
```

試し打ち。

```text
$ cd C:\path\to\nyagos

$ grep grep **/*.lua
./nyagos.d/comspec.lua:16:nyagos.alias.grep = "findstr.exe"
```

よーし，うむうむ，よーし。

## ブックマーク

- [Windowsのfindstrで正規表現を検索する：Tech TIPS - ＠IT](http://www.atmarkit.co.jp/ait/articles/0412/18/news018.html)

[grep]: https://linuxjm.osdn.jp/html/GNU_grep/man1/grep.1.html "Man page of GREP"
[ATOM]: https://atom.io/ "Atom"
[秀丸]: http://hide.maruo.co.jp/software/hidemaru.html "秀まるおのホームページ(サイトー企画)－秀丸エディタ"
[MSYS2]: http://msys2.github.io/ "MSYS2 installer"
[Go 言語]: https://golang.org/ "The Go Programming Language"
[mattn/go-iconv]: https://github.com/mattn/go-iconv "mattn/go-iconv: iconv binding for golang"
[`jvgrep`]: https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer"
[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
