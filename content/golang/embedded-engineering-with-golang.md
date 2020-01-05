+++
title = "組込みで Go"
date = "2018-01-21T18:39:17+09:00"
description = "組込み関連の記事を見かけたのでブックマークしておく。"
image = "/images/attention/go-code.png"
tags = [ "embedded", "engineering", "golang", "bookmark" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Qiita] の [@tetsu_koba](https://qiita.com/tetsu_koba "tetsu_koba - Qiita") さんが最近組込み関連の記事を連投されているのでブックマークしておく。

- [組み込みLinuxでGolangのススメ - Qiita](https://qiita.com/tetsu_koba/items/7435ef8d0c77844d751e)
- [Golangから物理メモリを読み書きする - Qiita](https://qiita.com/tetsu_koba/items/dba170bf220c45781428)
- [Golangでioctlのシステムコールを使う - Qiita](https://qiita.com/tetsu_koba/items/decee4d1a6ff621a7d37)
- [GolangでGPIOの割り込み通知を受け取る - Qiita](https://qiita.com/tetsu_koba/items/1928730136736c9dd133)
- [Golangの実行ファイルを複数まとめてトータルのファイルサイズを減らす工夫(busybox方式) - Qiita](https://qiita.com/tetsu_koba/items/53d84286ba5d87de607a)
- [GolangのプロセスをFIFO priorityにセットする - Qiita](https://qiita.com/tetsu_koba/items/1ccca9b3f4bd1e6b7f5c)
- [Golangでシリアルポート(UART)を使う - Qiita](https://qiita.com/tetsu_koba/items/f8afbb8326ee42fd27f5)
- [GolangでBLE (Bluetooth Low Enagy)を使う - Qiita](https://qiita.com/tetsu_koba/items/7d8f2f40e45e1549a6fa)

[Go 言語]で組込みといってもフルスクラッチでシーケンサみたいなのを組むわけではなく， RT Linux 下でのリアルタイム処理を想定しているようだ。

独学での組込みソフトの勉強は，具体的なターゲットがないと「その辺の出来合いのハードで『Lチカ[^lt1]』組んで満足」みたいなことになりかねないので，手を出すのを躊躇っていた。
でも今時は組込みつっても要は標準的な RTOS ( Real-Time Operating System) で動くアプリケーションなので（スマホ・アプリの開発も「組込み」カテゴリらしいし），勉強ならもっと気楽に構えていいのかなと思い直している。

[^lt1]: LED をチカチカ点滅させる組込み開発の “Hello World” みたいなやつ（笑）

とはいえ，私の場合は生活基盤を立て直すことから始めないといけないので，やるにしても暫く先になるかな。

## その他のブックマーク

- [GoでFPGAしてみる(Reconfigure.io) - Qiita](https://qiita.com/mjhd-devlion/items/5e6f6f2f40ecb4ad4217)
- [Goでのシリアル通信でハマった事 - Qiita](https://qiita.com/tomoya0x00/items/d957dc00682c57f96771)
- [Gobotの招きにあひて、徒然なるままにArduinoとRaspberry PiでIoTっぽいことをやってみるなり - Qiita](https://qiita.com/KemoKemo/items/10fb644f9d359c35646a)
- [Go言語のリアルタイムGC　理論と実践 | プログラミング | POSTD](http://postd.cc/golangs-real-time-gc-in-theory-and-practice/)
- [Go言語の低レイテンシGC実現のための取り組み | プログラミング | POSTD](http://postd.cc/gos-march-to-low-latency-gc/)
- [minimumgo: Linuxでgolangの実行ファイルをひとつだけ動かすときに必要最小限の初期化処理をしてくれるgolangのパッケージ - Qiita](https://qiita.com/tetsu_koba/items/059849c0871a7e3bd94f)
- [CHIP-8 &amp; Golang でエミュレータ作成入門してみた - Qiita](https://qiita.com/tuboc/items/b87f9a346fdf522a40fa)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Qiita]: https://qiita.com/

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
