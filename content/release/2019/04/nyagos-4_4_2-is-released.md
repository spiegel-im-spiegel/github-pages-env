+++
title = "NYAGOS 4.4.2 がリリースされた"
date = "2019-04-07T09:16:12+09:00"
description = "今回は盛り沢山だぞ！"
image = "/images/attention/tools.png"
tags  = [ "tools", "nyagos", "shell", "windows", "linux" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[NYAGOS] 4.4.2_0 がリリースされた。
つか，もたもたしてたら 4.4.2_1 も出てたよ。

- [Release 4.4.2_0 · nyaosorg/nyagos · GitHub](https://github.com/nyaosorg/nyagos/releases/tag/4.4.2_0)
- [Release 4.4.2_1 · nyaosorg/nyagos · GitHub](https://github.com/nyaosorg/nyagos/releases/tag/4.4.2_1)

変更点は以下の通り。
盛り沢山だぞ！

{{% fig-quote type="markdown" title="Release 4.4.2_0" link="https://github.com/nyaosorg/nyagos/releases/tag/4.4.2_0" lang="en" %}}
- OLEオブジェクトからLuaオブジェクトへの変換が日付型などでパニックを起こす不具合を修正
- Luaの数値が実数として OLE に渡されるべきだったのに、整数として渡されていた。
- Lua: 関数: `nyagos.to_ole_integer(n)` (数値を OLE 向けの整数に変換)を追加(trash.lua用)
- Lua: OLEObject に列挙用オブジェクトを得るメソッド `_iter()` を追加
- Lua: OLEObject を開放するメソッド `OLEObject:_release()` を追加
- trash.lua が COM の解放漏れを起こしていた問題を修正
- Lua: `create_object`生成された IUnkown インスタンスが解放されていなかった不具合を修正
- 「~ユーザ名」の展開を実装
- バッチファイル以外の実行ファイルの exit status が表示されなくなっていた不具合を修正
- %COMSPEC% が未定義の時に CMD.EXE を用いるエイリアス(ren,mklink,dir,...)が動かなくなっていた不具合を修正
- 全角空白(%U+3000%)がパラメータの区切り文字と認識されていた点を修正
- (#359) -c,-k オプションで CMD.EXE のように複数の引数をとれるようにした
- 「存在しないディレクトリ\何か」を補完しようとすると「The system cannot find the path specified.」と表示される不具合を修正 (Thx! tsuyoshicho)
- (#360) 幅ゼロやサロゲートペアな Unicode は`<NNNNN>` と表示するようにした (Thx! tsuyoshicho)
- サロゲートペアな Unicode をそのまま出力するオプション --output-surrogate-pair を追加
- suコマンドで、ネットワークドライブが失なわれないようにした
- (#197) ソースがディレクトリで -s がない時、`ln` はジャンクションを作成するようにした
- 内蔵の mklink コマンドを実装し、`CMD.exe /c mklink` のエイリアス `mklink` を削除
- ゼロバイトの Lua ファイルを削除(cdlnk.lua, open.lua, su.lua, swapstdfunc.lua )
- (#262) `diskfree` でボリュームラベルとファイルシステムを表示するようにした
- UNCパスがカレントディレクトリでもバッチファイルを実行できるようにした。
- UNCパスがカレントディレクトリの時、ren,assoc,dir,for が動作しない不具合を修正
- (#363) nyagos.alias.COMMAND="string" 中では逆クォート置換が機能しない問題を修正 (Thx! tostos5963 & sambatriste )
- (#259) アプリケーションをダイアログで選んでファイルを開くコマンド `select` を実装
- `diskfree` の出力フォーマットを修正
{{% /fig-quote %}}

{{% fig-quote type="markdown" title="Release 4.4.2_1" link="https://github.com/nyaosorg/nyagos/releases/tag/4.4.2_1" lang="en" %}}
- diskfree: 行末の空白を削除
- `~"\Program Files"` の最初の引用符が消えて、`Files` が引数に含まれない不具合を修正
{{% /fig-quote %}}

新たに追加された `select` コマンドはエクスプローラのコンテキストメニュー「プログラムから開く」に近い機能を提供していて，たとえば

```text
$ select index.html
```

とか打てば

{{< fig-img src="./select.png" link="./select.png" width="672" >}}

てな感じでファイルを開くプログラムの選択ダイアログが開く。
なにそれ素敵！

まるきし余談だが，先々月に [NYAGOS] 4.4.1 が出てて，私も Issue 上げてたのに記事にしてなかった。
1月2月はホンマに余裕がなかったからなぁ。
まぁ今は余裕があるかと言われればそうでもないんだけど...

アップデートは計画的に。

## 【2019-04-13 追記】 NYAGOS 4.4.2_2 がリリースされた

{{% fig-quote type="markdown" title="Release 4.4.2_2" link="https://github.com/nyaosorg/nyagos/releases/tag/4.4.2_2" lang="en" %}}
- Ctrl-RIGHT,ALT-F(次の単語へ), Ctrl-LEFT,ALT-B(前の単語へ)を実装
- インクリメンタルサーチ開始時にトップへ移動する時のバックスペースの数が間違っていた不具合を修正
- (#364) `ESC[0A` というエスケープシーケンスが使われていた不具合を修正
{{% /fig-quote %}}

[NYAGOS]: https://github.com/nyaosorg/nyagos/ "nyaosorg/nyagos: NYAGOS - The hybrid UNIXLike Commandline Shell for Windows"
