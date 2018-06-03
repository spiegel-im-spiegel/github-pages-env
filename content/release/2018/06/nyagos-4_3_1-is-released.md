+++
title = "NYAGOS 4.3.1 がリリース"
date = "2018-06-03T18:16:56+09:00"
description = "いくつかの起動時オプションが追加された。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows" ]

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

[NYAGOS] 4.3.1 がリリースされた。

- [Release 4.3.1_0 · zetamatta/nyagos](https://github.com/zetamatta/nyagos/releases/tag/4.3.1_0)

いくつかの起動時オプションが追加された。

{{% fig-quote title="Release 4.3.1_0" link="https://github.com/zetamatta/nyagos/releases/tag/4.3.1_0" lang="en" %}}

- Support Windows10's native ESCAPE SEQUENCE processing with `--no-go-colorable` and `--enable-virtual-terminal-processing`
- For #304,#312, added options to search for the executable from the current directory
    - `--look-curdir-first`: do before %PATH% (compatible with CMD.EXE)
    - `--look-curdir-last` : do after %PATH% (compatible with PowerShell)
    - `--look-curdir-never`: never (compatible with UNIX Shells)
- nyagos.prompt can now be assigned string literal as prompt template directly.
- Fix #314 rmdir could not remove junctions.

----

- `--no-go-colorable` と `--enable-virtual-terminal-processing` で、Windows10 ネイティブのエスケープシーケンスをサポート
- #304,#312, カレントディレクトリから実行ファイルを探す時のオプションを追加
    - `--look-curdir-first`: `%PATH%` より前に探す(デフォルト:CMD.EXE互換動作)
    - `--look-curdir-last` : `%PATH%` より後に探す(PowerShell互換動作)
    - `--look-curdir-never`: `%PATH%` だけから実行ファイルを探す(UNIX Shells互換動作)
- nyagos.prompt にプロンプトテンプレートの文字列を直接代入できるようになった。
- #314 rmdir がジャンクションを削除できなかった問題を修正
{{% /fig-quote %}}

実行ファイルの検索オプションだが `--look-curdir-never` について `--look-curdir-last` との違いがよく分からなかったり。
たとえばカレントディレクトリに `hoge.exe` があって PATH 上に同名のファイルが存在しない場合に，どちらのオプションでも

```text
$ hoge
```

で起動してしまう。
私としては `--look-curdir-never` 指定時は `./hoge` のみで起動することを期待したのだが違うのだろうか。
まぁ，大した問題ではないのでよかろう。

にしても，コマンドプロンプトと PowerShell ってそんなところで挙動が違うんだなぁ。
いいのか，それ。

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
