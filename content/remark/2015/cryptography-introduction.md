+++
date = "2015-09-20T21:43:17+09:00"
update = "2016-01-07T13:09:45+09:00"
description = "今回は大幅改訂版なので，以前のを持ってる人も買っておいて損はない。"
draft = false
tags = ["book", "cryptography", "security", "trust", "sha-3", "blockchain", "ecc"]
title = "『暗号技術入門 第3版』をななめ読み"

[scripts]
  mathjax = false
  mermaidjs = false
+++

結城浩さんの『[暗号技術入門 第3版](http://www.hyuki.com/cr/)』がついに登場。
前の第2版のときは細々した追記が主だったような気がするが，今回は大幅改訂版なので，以前のを持ってる人も買っておいて損はない。
主な改訂内容としては

- SHA-3 について詳しく解説
- HeartBleed や POODLE など，最近の攻撃手法について言及
- 付録で楕円曲線暗号（Elliptic Curve Cryptography）について詳しく解説
- Bitcoin というか Bitcoin の中の重要な技術要素である Block Chain について詳しく解説

他にも [GnuPG](https://www.gnupg.org/) の記述が modern version に対応してたり，認証つき暗号（AEAD; Authenticated Encryption with Associated Data）およびその実装である GCM (Galois/Counter Mode) への言及があったり，いろいろ細かく手直しされている。

特に楕円曲線暗号の解説は秀逸で，入門レベルでの解説の中では一番分かりやすかった。
あと Block Chain の解説もお勧め。
Bitcoin や Block Chain に関する解説本はすでにたくさん出ているが，暗号技術の観点からきちんと解説しているものは少なく，「信用モデル」にまで話を展開しているものは更に少ない。

結局，暗号システムの実装というのは究極的には「信用モデル」の開発であると言える。
問題は「信用モデル」はロジックだけでは成立しない，ということだ。
『[暗号技術入門 第3版](http://www.hyuki.com/cr/)』では信用モデルの例として hierarchical PKI の典型である X.509 と OpenPGP の Web of Trust，そして Block Chain を挙げているが，それぞれ特徴があって比較すると面白い。
たとえば Block Chain はユーザ間に働く経済的 incentive を巧妙に使うが，それだけにパラメータの調整が難しいし， Mt. Gox のような経済リスクも考慮しなくてはならない。

そもそも信用というのは過去の事実に対してのみ評価可能なのに，実際に評価したいのは現在及び未来についてなのだ。
これって本来は不能解だよね。
でも信用が評価できなくては世の中は回らない。
だから，どうにかして実用可能なレベルにまで近似できないか，と専門家やエンジニアは頭を悩ますわけ。

そういったことを頭の隅に入れながら読めば，この本は単なるリファレンス本ではないことに気づくと思う。

最後にちょっとだけ注文をつけるなら「前方秘匿性（PFS; Perfect Forward Secrecy）」および OTR (Off-the-Record) Messaging の肝である「否認可能（Deniability）」についても言及が欲しかった[^1]。
メッセージングの世界においてはこのふたつが重要な要件になってきているからだ。

[^1]: PFS についてはもしかしたら見落としてるかもしれないが。なにせ斜め読みだったから。

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4757143044" %}} <!-- 信頼と裏切りの社会 -->
