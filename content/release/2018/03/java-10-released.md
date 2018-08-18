+++
title = "Java 10 Released"
date = "2018-03-21T15:33:51+09:00"
update = "2018-05-13T13:24:15+09:00"
description = "時間はあまり残されていない。 まだ決めかねている方は早期の決断をオススメする。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "engineering" ]

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

Java 10 がリリースされた。

- [Oracle Java SE 10 Release Arrives](https://www.oracle.com/corporate/pressrelease/Java-10-032018.html)
- [Java SE 10 (18.3) ( JSR 383)](http://cr.openjdk.java.net/~iris/se/10/fr/java-se-10-fr-spec/)

Oracle Java の JDK は既に入手可能である。

- [Java SE Development Kit 10- - Downloads](http://www.oracle.com/technetwork/java/javase/downloads/jdk10-downloads-4416644.html)

[OpenJDK] 版は以下から取得できるようだ。

- [JDK Builds from Oracle](http://jdk.java.net/)
    - [JDK 10 GA Release](http://jdk.java.net/10/)

[OpenJDK] 版を起動すると以下の表示になる（Windows の場合）。

```text
$ java -version
openjdk version "10" 2018-03-20
OpenJDK Runtime Environment 18.3 (build 10+46)
OpenJDK 64-Bit Server VM 18.3 (build 10+46, mixed mode)
```

「[Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})」でも書いたが， Oracle が Java を無償サポートするのはこの Java 10 で最後である[^lcs1]。
2018年9月にリリースされる Java 11 を以って Java 10 の無償サポートは終了し， Java 11 以降は有償の LTS (Long Term Support) のみとなる。
なお Java 1.8 系（Java 8）の無償サポートも2019年1月で終了する（個人ユーザへの無償サポートは2020年12月まで）。
つまり，現時点では Java 環境に関する3つの選択肢がある。

[^lcs1]: 有償・無償というのはサポートに関するもので Java コンパイラおよび標準ライブラリは今のところ無償で入手可能である。 Oracle Java は [Oracle Binary Code License](http://jdk.java.net/java-se-ri/10-bcl)， [OpenJDK] は [GPLv2 w/ CE](http://openjdk.java.net/legal/gplv2+ce.html) の下に提供されている。しかし Java のメンテナンスは事実上 Oracle が主導権を握っているため，今後 [OpenJDK] がどこまでそれを肩代わりできるかがポイントとなる。

1. 2019年2月以降も Oracle にお金を払って Java 1.8 系を利用し続ける
2. 2018年9月から Oracle にお金を払って Java 11 の LTS を受ける
3. [OpenJDK] に乗り換え半年ごとの短期リリースに追従する

ちなみに Oracle Java の有償サポートはコア単位で数十万円の桁（ユーザ単位なら10万円ちょっと）なので中小企業以下では全く見合わないし，お役所が自身のサービスに Oracle Java を使うのは税金の無駄遣いであると断言する。
かといって [OpenJDK] で半年ごとのリリースサイクルに追従するのはなかなか大変である（それで顧客からお金を取る手もあるけどね。 SIer さん，出番ですよ`w`）。
Java 9 以降では，それ以前のバージョンとの後方互換性は失われている。
今後もそうならないとは限らない。
Java の特徴とされた "[Write once, run anywhere](https://en.wikipedia.org/wiki/Write_once,_run_anywhere)” はもはや存在しないのだ[^wqda1]。

[^wqda1]: そういや昔， "Write once, debug anywhere” というネガティブキャンペーンがあったな（笑）

時間はあまり残されていない。
まだ決めかねている方は早期の決断をオススメする。

## ブックマーク

- [Java 10が本日付で正式リリース。ローカル変数の型推論、ガベージコレクタが入れ替え可能、不揮発性メモリ対応など。Java 9は早くもサポート期間終了 － Publickey](http://www.publickey1.jp/blog/18/java_10java_9.html)
- [Java 10新機能まとめ - Qiita](https://qiita.com/nowokay/items/d9bc4b3f715d17c2830d)
- [Eclipse Foundationが、Java EEの新ブランドとなった「Jakarta EE」のロゴをコミュニティから募集中。締め切りは3月21日 － Publickey](http://www.publickey1.jp/blog/18/eclipse_foundationjava_eejakarta_ee321.html)
- [OpenJDK入手先まとめ - Qiita](https://qiita.com/ykubota/items/582caa8621a5fc86d0a1)
- [Java 10 で変わる Java のバージョン表記 - Qiita](https://qiita.com/YujiSoftware/items/2c5a9117a577700ea540)
- [2017年末において、Oracle JDK と OpenJDKに性能差はあるのか？ - Qiita](https://qiita.com/ukiuni@github/items/5bb523c31bb4b4f9a278)
- [Java 環境のリリースとサポートに関する覚え書き]({{< ref "/remark/2018/02/release-cycle-of-java-environment.md" >}})
- [Java 10.0.1 がリリース（セキュリティ・アップデート）]({{< ref "/release/2018/04/java-10_0_1-is-released.md" >}})

[OpenJDK]: http://openjdk.java.net/
