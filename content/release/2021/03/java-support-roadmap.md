+++
title = "Java のサポート期限ってどうなってるんだっけ？"
date =  "2021-03-19T20:25:53+09:00"
description = "アップデートは計画的に。"
image = "/images/attention/tools.png"
tags  = [ "java", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ぶっちゃけ [PlantUML] を動かす程度の利用なら最新の [OpenJDK] だけ気にしていればよかったんだけど，仕事で Java を使うとなればそういうわけにもいかない。
けど，今の Java のサポート期限ってどうなってるんだっけ？

まずは Oracle Java。
これははっきりしている。

- [Oracle Java SE Support Roadmap](https://www.oracle.com/java/technologies/java-se-support-roadmap.html)

これによると

| Release  | GA Date | Premier Support | Extended Support |
| -------- | ------- | --------------- | ---------------- |
| 7        | 2011-07 | 2019-07         | 2022-07          |
| 8        | 2014-03 | 2022-03         | 2030-12          |
| 9        | 2017-09 | 2018-03         | -                |
| 10       | 2018-03 | 2018-09         | -                |
| 11 (LTS) | 2018-09 | 2023-09         | 2026-09          |
| 12       | 2019-03 | 2019-09         | -                |
| 13       | 2019-09 | 2020-03         | -                |
| 14       | 2020-03 | 2020-09         | -                |
| 15       | 2020-09 | 2021-03         | -                |
| 16       | 2021-03 | 2021-09         | -                |
| 17 (LTS) | 2021-09 | 2026-09         | 2029-03          |

ということらしい[^java8]（どのバージョンでも Sustaining Support は無期限で受けれるため省いた）。
つまり Premier Support を受けれるバージョンは現時点（2021-03）で 8, 11, 15/16 が対象となる（15 は 2021-03 まで）。
Extended Support なら 7 も含まれる（Java 7 なら費用は免除）。

[^java8]: Oracle Java 8 はクライアント側での利用と開発目的での利用のみ許可されている。サーバ側での運用は別途有償契約が必要（の筈）。

一方 [OpenJDK] はよく分からない。

Java 10 以降なら（リリースサイクルが半年毎というだけで）特にサポート期限というものはないみたい。
向こうが「やーめた」と言うまではアップデートを出し続けるのかな？ 7u と 8u については記述がなかったが，四半期ごとのマイナーアップデートの対象になってるみたいなので，一応はサポートしていると思われる。
まぁ， Oracle Java や [AdoptOpenJDK], [OpenJ9] などからフィードバックがあるのかもしれない。

ちなみに [AdoptOpenJDK] のサポート期限は

- [Support | AdoptOpenJDK - Open source, prebuilt OpenJDK binaries](https://adoptopenjdk.net/support.html)

によると

| バージョン | First Availability | End of Availability |
| ---------- | ------------------ | ------------------- |
| 8 (LTS)    | 2014-03            | at least 2022-05    |
| 9          | 2017-09            | 2018-03             |
| 10         | 2018-03            | 2018-09             |
| 11 (LTS)   | 2018-09            | 2024-10             |
| 12         | 2019-03            | 2019-09             |
| 13         | 2019-09            | 2020-03             |
| 14         | 2020-03            | 2020-09             |
| 15         | 2020-09            | 2021-03             |
| 16         | 2021-03            | 2021-09             |
| 17 (LTS)   | 2021-09            | (TBC)               |

となっていた。
現時点（2021-03）では 8, 11, 15/16 がサポート対象だ（15 は 2021-03 まで）。

今はクラウドサービスごとに独自のディストリビューションを提供してたりするので， 11 以降ならあまり気にする必要はないのかな。
今だに 8/8u を使ってるサービス（行政サービスとか多そうだw）は，そろそろ未来について（Java を捨てる選択肢も含めて）議論したほうがいいだろう。

なお，年間のリリース・イベントは以下の通り。

|   月 | バージョンアップ・イベント |
| ----:| -------------------------- |
|  1月 | マイナー・バージョンアップ |
|  2月 |                            |
|  3月 | メジャー・バージョンアップ |
|  4月 | マイナー・バージョンアップ |
|  5月 |                            |
|  6月 |                            |
|  7月 | マイナー・バージョンアップ |
|  8月 |                            |
|  9月 | メジャー・バージョンアップ |
| 10月 | マイナー・バージョンアップ |
| 11月 |                            |
| 12月 |                            |

アップデートは計画的に。

## ブックマーク

- [Eclipse Adoptium | projects.eclipse.org](https://projects.eclipse.org/projects/adoptium)
    - [AdoptOpenJDKプロジェクトがEclipse Foundationへの合流を発表。合流後の新プロジェクト名は「Eclipse Adoptium」に － Publickey](https://www.publickey1.jp/blog/20/adoptopenjdkeclipse_foundationeclipse_adoptium.html)
- [AWS、「Java 8」を2026年まで、「Java 11」は2027年まで、現行より3年サポート期間延長を発表。独自JavaディストリビューションのCorretto 8とCorretto 11で － Publickey](https://www.publickey1.jp/blog/20/awsjava_82026java_1120273javacorretto_8corretto_11.html)
- [OpenJDKと各種JDKディストリビューションの情報源まとめ](https://zenn.dev/yamadamn/articles/2e3b388076cbde229655)

[OpenJDK]: http://openjdk.java.net/
[AdoptOpenJDK]: https://adoptopenjdk.net/ "AdoptOpenJDK - Open source, prebuilt OpenJDK binaries"
[OpenJ9]: http://www.eclipse.org/openj9/ "Eclipse OpenJ9"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
