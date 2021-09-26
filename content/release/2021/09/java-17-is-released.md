+++
title = "Java 17 LTS がリリースされた"
date =  "2021-09-15T20:27:14+09:00"
description = "Java のサポート期限ってどうなってるんだっけ？"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "windows" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

予定通り Java 17 がリリースされた。
Java 11 以来の LTS (Long-Term-Support) バージョンである。

- [JDK 17](https://openjdk.java.net/projects/jdk/17/)
- [JDK 17 GA Release](https://jdk.java.net/17/)
- [JDK 17 Release Notes](https://jdk.java.net/17/release-notes)

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/17/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk17/0d483333a00540d886896bac774ff48b/35/GPL/openjdk-17_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-17_linux-x64_bin.tar.gz
$ sudo ln -s jdk-17 java
$ java -version # すでに PATH が通っている場合
openjdk version "17" 2021-09-14
OpenJDK Runtime Environment (build 17+35-2724)
OpenJDK 64-Bit Server VM (build 17+35-2724, mixed mode, sharing)
```

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-08-30 に [PlantUML] V1.2021.10 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

うむうむ。
ちゃんと動くな。

## Java のサポート期限ってどうなってるんだっけ？

まずは Oracle Java。
これははっきりしている。

- [Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)

これによると

| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 7 (LTS)  | 2011-07 | 2019-07         | 2022-07          |
| 8 (LTS)  | 2014-03 | 2022-03         | 2030-12          |
| 11 (LTS) | 2018-09 | 2023-09         | 2026-09          |
| 16       | 2021-03 | 2021-09         | -                |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-09          |
| 18       | 2022-03 | 2022-09         | -                |
| 19       | 2022-09 | 2023-03         | -                |
| 20       | 2023-03 | 2022-09         | -                |
| 21 (LTS) | 2023-09 | 2028-09         | 2031-09          |

となっている[^java8]（2021-09 時点でサポートが終了しているものは除いている）。

[^java8]: Oracle Java 8 はクライアント側での利用と開発目的での利用のみ許可されている。サーバ側での運用は別途有償契約が必要（の筈）。

一方 [OpenJDK] はよく分からないが [AdoptOpenJDK] については

- [Support | AdoptOpenJDK - Open source, prebuilt OpenJDK binaries](https://adoptopenjdk.net/support.html)

によると

| バージョン | First Availability | End of Availability |
| ---------- | ------------------ | ------------------- |
| 8 (LTS)    | 2014-03            | at least 2026-05    |
| 11 (LTS)   | 2018-09            | 2024-10             |
| 16         | 2021-03            | 2021-09             |
| 17 (LTS)   | 2021-09            | (TBC)               |

となっている。
まだ更新されないのかな？

こういっちゃあ何だが，きょうび Oracle Java をわざわざ選択する理由はないだろう。
どうしても Oracle のクラウドを使いたいというのなら別だが（笑）

Amazon にしたって Microsoft にしたって IBM にしたって，自社のクラウドでそれぞれオープンな製品を用意してる。
つまり「どの Java を使うか」ではなく「どの XaaS 環境を使うか」で必然的に Java も決まってしまうのである。

いや，この前見かけた [#ITアウトレイジ](https://twitter.com/hashtag/IT%E3%82%A2%E3%82%A6%E3%83%88%E3%83%AC%E3%82%A4%E3%82%B8)が面白くて

{{< div-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">とっくにOpenJDKに移行済みなんだよ！今更戻れるわけねーだろ！<a href="https://twitter.com/hashtag/IT%E3%82%A2%E3%82%A6%E3%83%88%E3%83%AC%E3%82%A4%E3%82%B8?src=hash&amp;ref_src=twsrc%5Etfw">#ITアウトレイジ</a><a href="https://t.co/FL8FyhEjSj">https://t.co/FL8FyhEjSj</a> <a href="https://t.co/UAtrrnvABM">pic.twitter.com/UAtrrnvABM</a></p>&mdash; hisa_u (@hisa_u) <a href="https://twitter.com/hisa_u/status/1438057120302190595?ref_src=twsrc%5Etfw">September 15, 2021</a></blockquote>
{{< /div-gen >}}

って，感じだよねぇ（笑）

## ブックマーク

- [3年ぶりの長期サポート版となる「Java 17」正式版がリリース。M1 Macのサポート、Sealed Classの追加など － Publickey](https://www.publickey1.jp/blog/21/3java_17m1_macseald_class.html)
- [オラクル、Oracle JDKを再び無料提供へ、本番環境でも利用可。昨日リリースのJava 17から － Publickey](https://www.publickey1.jp/blog/21/oracle_jdkjava_17.html)
- [Oracle、「Java 17」を発表 ～3年ぶりの長期サポートリリース（LTS） - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1351146.html)
- [Javaの長期サポート（LTS）版、次回は2年後に登場の見通し。オラクルがLTSのサイクルを3年から2年に変更提案 － Publickey](https://www.publickey1.jp/blog/21/javalts2lts32.html)
- [AWS、Java 17対応の無料Javaディストリビューション「Amazon Corretto 17」リリース。ただし長期サポート期間は未定 － Publickey](https://www.publickey1.jp/blog/21/awsjava_17javaamazon_corretto_17.html)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
{{% review-paapi "B06XC6CYRF" %}} <!-- 大鉄人１７ -->
