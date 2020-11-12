+++
title = "そうだ！ Zenn の記事や本にバッヂを贈ろう"
date =  "2020-11-12T21:17:02+09:00"
description = "バッヂを贈られた記事は「金を払う価値のある記事」という実績を示せるし，贈った側も「オラが贈ったバッヂだべや」と悦に入ることができる"
image = "/images/attention/kitten.jpg"
tags = [ "web", "market" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

仕事でまた Windows を弄るようになって[^win10] 改めてウンザリしていること：

[^win10]: 春までいた職場でも Windows 10 は使っていたが，あれは PC は PC でも Programmable Controller のほうなので，指定されたアプリや操作以外は許容されていなかった。

1. ディレクトリ区切り文字がバックスラッシュ
    - エスケープ文字と被るのでチョー面倒
2. 改行コードが2バイト（CR+LF）
    - ちゃんと2バイトにしないとバッチが動かない
3. ダブルバイトの呪い。また Shift-JIS と付き合わにゃならんのか `orz`
    - テキストボックスに漢字10文字入るからって半角英数字で20文字入るわけぢゃねーよ！ いつの時代だ

そんなこんなで，習作（study）として手遊びで作った [gnkf] が実務で役に立つ日が来ようとは...

たとえば Shift-JIS のテキストを UTF-8 で「読める」ようにするために

```text
$ gnkf e -s shift_jis -f sjis.txt | gnkf nl -n lf > utf8.txt
```

としたり，逆に

```text
$ gnkf e -d shift_jis -f utf8.txt | gnkf nl -n crlf > sjis.txt
```

とか Shift-JIS に変換したりするわけなのだが，いちいち手で打つのは面倒なのでバッチファイルにしたりするよね。
そんで，バッチファイルにうっかりそのまま1行だけ

```bat
gnkf e -s shift_jis -f %1 | gnkf nl -n lf
```

と書いて

```text
$ conv-utf8.cmd sjis.txt > utf8.txt
```

とか実行したら大変なことになるのですよ。
分かるかな？ ここではじめて「そういやバッチ書くときって先頭に `@echo off` って付けにゃアカンかったっけ」と思い出すのである。

はい。
長い前置きでした。
ここからが本題。

こんなときに役に立つのが以下の本。

- [/bin/shに慣れた人に贈るバッチファイルの書き方](https://zenn.dev/zetamatta/books/c84cbe23093eee1b5830)

マジで役に立ちました，過ぎ去った過去を思い出す意味でも。
これを読んで最終的にはこんな感じに書き直した。

```bat
@ECHO OFF
SETLOCAL
SET "INFILE=%~1"
IF "%INFILE%" == "" GOTO ENDPROC
SET "OUTFILE=%~2"
IF "%OUTFILE%" == "" (
  gnkf e -s shift_jis -f "%INFILE%" | gnkf nl -n lf
) ELSE (
  gnkf e -s shift_jis -f "%INFILE%" | gnkf nl -n lf -o "%OUTFILE%"
)
:ENDPROC
ENDLOCAL
EXIT /b
```

こんな有用な本が無料ですよ。
もったいねー

でも，ご安心を。
[Zenn] では「サポート」機能を使って無料の記事や本に送金できるのである。

{{< fig-img src="./support-button-in-zenn.png" link="./support-button-in-zenn.png" width="712" >}}

このボタンを押すと

{{< fig-img src="./support-it.png" link="./support-it.png" width="527" >}}

てなダイアログが表示され，500円単位で支払いできる。

注目は下半分に表示されているバッヂだ。
送金時にこのバッヂを選択することで，実際にサポート受けた記事では

{{< fig-img src="./badges.png" link="./badges.png" >}}

てな感じで贈られたバッヂが表示される。

金銭的なやり取り以上に，バッヂを贈られた記事は「金を払う価値のある記事」という実績を示せるし，もちろん贈った側も「オラが贈ったバッヂだべや」と悦に入ることができる（ちなみに誰がバッヂを贈ったかは外部には非公開）。

このバッヂ機能，シンプルだけど強力な仕掛けだよね。
単なる LGTM[^lgtm1] よりも実利的な意味を持つし note.com のサポート機能より更に一歩進んでいると思う。
こういうところが新興サービスの面白いところだよな。

[^lgtm1]: “Looks Good To Me” の略。

というわけで，お世話になった無料本のいくつかにバッヂを贈りました。
お世話になりました。

[gnkf]: {{< ref "/release/gnkf.md" >}} "NKF: Network Kanji Filter by Golang"
[Zenn]: https://zenn.dev/ "Zenn｜プログラマーのための情報共有コミュニティ"
<!-- eof -->
