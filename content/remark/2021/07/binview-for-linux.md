+++
title = "バイナリ・データ・ビュア zetamatta/binview の Linux 用バイナリ登場"
date =  "2021-07-07T11:24:20+09:00"
description = "CUI なターミナル・エミュレータ上でちょろんと眺めてさくっと編集できる"
image = "/images/attention/tools.png"
tags  = [ "tools", "windows", "linux", "ubuntu", "editor" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

バイナリ・データ・ビュアの [zetamatta/binview][binview] の v0.2.0 がリリースされたのだが，このリリースで Linux 用バイナリも提供されるようになった。

- [Release v0.2.0 · zetamatta/binview · GitHub](https://github.com/zetamatta/binview/releases/tag/v0.2.0)
- [Release v0.2.1 · zetamatta/binview · GitHub](https://github.com/zetamatta/binview/releases/tag/v0.2.1)

もともと [binview] のバイナリは Windows 用のみの提供だったのだ。
ありがたや。

若い頃はバイナリ・エディタを使うような仕事も多くて [Bz Editor] などのツールのお世話になったものだが[^bz1]，もうバイナリデータを直接イジるような細かい仕事はできなくなりつつあるし，でも CUI なターミナル・エミュレータ上でちょろんと眺めてさくっと編集できるツールがあるといいなぁ，みたいな需要はあるわけですよ。

[^bz1]: 今なら[改造版 Bz](https://gitlab.com/devill.tamachan/binaryeditorbz/ "devill.tamachan / binaryeditorbz · GitLab") がオススメ。

ちなみに Linux には xxd というダンプツールがあり

```text
$ xxd -g 1 ./bindata1 
00000000: a8 03 50 47 50 c3 04 04 03 00 01 c9 38 e7 2d 2f  ..PGP.......8.-/
00000010: b1 f1 0f c3 ce 55 5d b2 8a 4b e8 4f 43 15 6e 7d  .....U]..K.OC.n}
00000020: 90 90 53 6a 9a e3 aa 1c 68 d6 d3 fc 6a 4e 79 a8  ..Sj....h...jNy.
00000030: e7 b1 a5 87 ea cc cc 99 66 31 ad ff e1 a3 03 b6  ........f1......
00000040: 47 85 76 bd 0b                                   G.v..
```

みたいな感じにバイナリデータを眺めることはできる。
あるいは [vim と xxd を組み合わせる](https://kaworu.jpn.org/vim/vim%E3%81%A7%E3%83%90%E3%82%A4%E3%83%8A%E3%83%AA%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB%E3%82%92%E6%89%B1%E3%81%86 "vimでバイナリファイルを扱う - neovim/vim入門")ことでバイナリ・エディタのように使うこともできるけど，あんまりお手軽じゃないのよね。
やり方をしょっちゅう忘れるし（笑） なので [binview] のようなお手軽ツールはありがたいし，最初からマルチプラットフォームな製品として構成されているのもありがたい。

ファイルを指定して [binview] を起動するとこんな感じになる。

{{< fig-img src="./binview-start-1.png" link="./binview-start-1.png" width="714" >}}

入力は標準入力からでもOK。
なので前処理をパイプで繋いで

{{< fig-img src="./binview-start-2.png" link="./binview-start-2.png" width="694" >}}

てな感じにできる。
カーソル位置のコードが UTF-8 (の一部) と判定された場合は下部のステータス行に Unicode 符号点が表示されるのが秀逸である。

終了する場合は `[q]` キー押下で

{{< fig-img src="./binview-quit.png" link="./binview-quit.png" width="709" >}}

のように確認プロンプトが表示されるので `[y]` キー押下で終了できる。

データの保存は `[w]` キー押下で

{{< fig-img src="./binview-write.png" link="./binview-write.png" width="694" >}}

とプロンプトが出るので，保存先のファイル名を指定して `[Enter]` キー押下で確定する。

これで，たとえば処理結果を標準出力に吐くツールと組み合わせてデータをいったん [binview] に流し込んで中身を確認してから任意のファイル名に保存する，てなことも簡単にできるわけだ。
うんうん。

`[r]` キーを押下するとカーソル位置のコードを変更できる。
こんな感じ。

{{< fig-img src="./binview-replace.png" link="./binview-replace.png" width="690" >}}

コマンドキーの一覧は以下の通り。

|    キー    | 機能                                         |
| :--------: | -------------------------------------------- |
|    `q`     | quit                                         |
|    `r`     | replace one byte                             |
|    `i`     | insert `'\0'` on the cursor                  |
|    `a`     | append `'\0'` at the rightside of the cursor |
| `x`, `DEL` | delete and yank one byte on the cursor       |
|    `p`     | paste 1 byte the rightside of the cursor     |
|    `P`     | paste 1 byte the leftside of the cursor      |
|    `w`     | output to file                               |

`[x]` と `[p]`/`[P]` は対になっていて `[x]` で削除したコードを `[p]`/`[P]` で任意の場所に挿入できるようだ。

なお `[ESC]` キー押下でコマンド前の状態に戻る。
なので，状態がよく分からなくなったら `[ESC]` 連打である（vi/vim でありがちな光景w）。

さらにカーソルの移動には `[h]` `[j]` `[k]` `[l]` キーが使える。
私のように vim を起動するとトラウマレベルで指が `[h]` `[j]` `[k]` `[l]` キーを押さえてしまう人には朗報だろう（笑）


[binview]: https://github.com/zetamatta/binview "zetamatta/binview: Binary data viewer on the terminal"
[Bz Editor]: https://www.vcraft.jp/soft/bz.html "Bz - c.mos"
<!-- eof -->
