+++
title = "NYAGOS 4.4.0 がリリースされた"
date = "2019-01-12T17:23:12+09:00"
description = "実験レベルながら Linux に対応したですよ。"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows", " linux" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[NYAGOS] 4.4.0 がリリースされた。

- [Release 4.4.0_0 · zetamatta/nyagos · GitHub](https://github.com/zetamatta/nyagos/releases/tag/4.4.0_0)

変更点は以下の通り。

{{% fig-quote title="Release 4.4.0_beta" link="https://github.com/zetamatta/nyagos/releases/tag/4.4.0_beta" lang="en" %}}
- Support Linux (experimental)
- Fix the problem that current directories per drive were not inherited to child processes.
- Use the library "mattn/go-tty" instead of "zetamatta/go-getch"
- Stop using msvcrt.dll via "syscall" directly
- On linux, the filename NUL equals /dev/null
- Add lua-variable nyagos.goos
- (#341) Fix an unexpected space is inserted after wide characters
    - On Windows10, enable stdout virtual terminal processing always
    - If git.exe push disable virtual terminal processing, enable again.
- (#339) Fix that wildcard pattern .??* matches ..
    - It requires github.com/zetamatta/go-findfile tagged 20181223-2

----

- Linux サポート(実験レベル)
- ドライブ毎のカレントディレクトリが子プロセスに継承されなかった問題を修正
- ライブラリ "zetamatta/go-getch" のかわりに "mattn/go-tty" を使うようにした
- msvcrt.dll を直接syscall経由で使わないようにした。
- Linux でも NUL を /dev/null 相当へ
- Lua変数 nyagos.goos を追加
- (#341) Windows10で全角文字の前に文字を挿入すると、不要な空白が入る不具合を修正
    - それに伴い、Windows10 では virtual terminal processing を常に有効に
    - git.exe pushが無効にしても再び有効にする
- (#339) ワイルドカード .??* が .. にマッチする問題を修正
    - 要 github.com/zetamatta/go-findfile tagged 20181230-2
{{% /fig-quote %}}


{{% fig-quote title="Release 4.4.0_0" link="https://github.com/zetamatta/nyagos/releases/tag/4.4.0_0" lang="en" %}}
- Remove beta
- To call a batchfile, stop to use `/V:ON` for CMD.EXE

----

- βを外した
- バッチファイルを呼ぶ時に、`/V:ON` を CMD.EXE に使わないようにした
{{% /fig-quote %}}

実験レベルながら Linux に対応したですよ。
これって Windows を捨てようとしている私へのご褒美？ 正直に言って Linux に移行する際の shell をどうしようか悩んでたのよ[^lua1]。
まぁ，当分先の話だが。

[^lua1]: 普段は bash でいいのだが，現在 [NYAGOS] 上で運用している Lua スクリプトを Linux でも使いたいなぁ，なんて。

## ブックマーク

- [「GNU Bash 5.0」リリース、10年ぶりのメジャーバージョンアップ － Publickey](https://www.publickey1.jp/blog/19/gnu_bash_5010.html)

[NYAGOS]: https://github.com/zetamatta/nyagos/ "zetamatta/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
