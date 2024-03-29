+++
title = "ハード屋が書く C ソースコードが凄まじかった思い出（再掲載）"
date = "2016-06-04T09:08:09+09:00"
description = "この件では，若いころのほろ苦い思い出がある。"
tags = ["engineering", "programming"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

これは[ちょうど1年前に Medium で書いた記事](https://medium.com/@spiegel/-1ca9e4895f4c)の再掲載である。
今は Medium を全く利用しないので，昔書いたもので（私が）面白いと思った記事は少しずつこっちに移転しようかな。

----

{{< fig-quote title="トヨタの車のソースコードはスパゲッティコード山盛り？ - YAMDAS現更新履歴" link="http://d.hatena.ne.jp/yomoyomo/20150604/toyota" >}}
<q>この記事で面白いのは、Michael Barr が20ヶ月以上にわたりトヨタ車で使われているソースコードを、Philip Koopman カーネギーメロン大学教授がトヨタのエンジニアリングの安全プロセスを精査した話で、両者ともトヨタのソフトウェアがスパゲッティコード山盛りなことを証言している。
<br>
トヨタの生産方式はアジャイル方面においてソフトウェア開発手法に多大な影響を与えている。ところでそのトヨタが開発するソフトウェアの品質はどうなんだろう、というのは多くの人の頭に浮かぶ疑問だろう。組み込みソフトウェアのエキスパートによると、ものすごく複雑で、複雑すぎてテストもメンテもできない関数がたくさんあるとか、グローバル変数が1万個以上あるとかなかなか壮絶らしい……。マジかよ。</q>
{{< /fig-quote >}}

この件では，若いころのほろ苦い思い出がある。

私は若いころは「システムハウス」と呼ばれる類の会社にいたのだが，そこではハードとソフトの両面で開発を進められるのが「売り」だった。
[私はハードは壊滅的にダメ](https://baldanders.info/blog/000529/ "私はこうしてプログラミングを覚えた — Baldanders.info")なのでソフトウェア担当。
当時のハード屋は自分が組んだ回路の実証のために自分でもプログラムを組んで動作確認する。
私たちソフト屋はそのコードをもらって実際のコードを書くわけだが，この実証コードが凄まじかった。

まず変数は全てグローバル変数。
スコープとかカプセル化なんて知るか！ という気概が感じられる。
そして関数は果てしなく長い main 関数のみ。
無間地獄のネスト。
goto 文で飛びまくり。
なのに異常系の記述は皆無。
世に聞く「スパゲッティ・コード」とはこのようなものなのかと感嘆したよ。

一番凄かったのは，とあるチップを使った30次のバンドパスフィルタを組むのに「サンプルがあるから簡単でしょ」と言われてサンプルを見たら世にも悍ましいコードで，解析するだけで半月もかかってしまった。

まぁ，ハード屋がこういうコードを書くのは理由があって，変数を記述するときはメモリ上のマッピングをそのまま置き換えようとするし，ロジックも基本的にマシン語のインストラクションをそのまま C に置き換えようとするから「関数」という概念がそもそもないことが多い[^a]。

[^a]: ハード屋は call インストラクションを「特殊なジャンプ」程度にしか認識していない（まぁ確かにそうなのだが）。あるプロジェクトで見せてもらったアセンブラコードでは call で積んだスタックをいじって戻り先アドレスを変えて return する記述が頻出していた。これはアセンブラ・コードにパッチを当てる際の基本テクニックらしい。大昔の話だよ（笑）

おかげで私は「他人のコードを読む」ことがすんごい得意になってしまった。
アレに比べればソフト屋の書くコードなんて絵本を読むように分かりやすい（笑） でも，こんなしょうもない特技でも後年ちゃんと役に立ってるんだから世の中というのは分からないものである。

{{% review-paapi "B01IGW5IIW" %}} <!-- リーン開発の現場 -->
