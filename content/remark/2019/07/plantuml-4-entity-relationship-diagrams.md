+++
title = "真面目に PlantUML (4) : 実体関連図"
date =  "2019-07-14T16:28:21+09:00"
description = "実は実体関連図は UML のクラス図を使って表すことができる。また IE 記法にも一部対応している。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "plantuml", "uml", "class", "entity", "relation" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は実体関連図（entity-relationship diagram）について。

実体関連モデル（entity-relationship model）は関係データベース（relational database）を抽象的に表現する手法のひとつで，実体関連モデルを図で表したものが実体関連図である。

実体関連モデルのうち特に物理データモデルを表した実体関連図は関係データベースの実装とよくマッチしているため，関係データベースのスキーマ設計でよく使われ，逆に関係データベースのスキーマを実体関連図に逆コンパイルするリバースエンジニアリング・ツールもある。

では実体関連図の記法について紹介してみよう。

## 目次

1. [PlantUML のインストール]({{< ref "/remark/2018/12/plantuml-1.md" >}})
1. [シーケンス図]({{< ref "/remark/2018/12/plantuml-2-sequence-diagram.md" >}})
1. [クラス図]({{< ref "/remark/2018/12/plantuml-3-class-diagrams.md" >}})
1. [実体関連図]({{< ref "/remark/2019/07/plantuml-4-entity-relationship-diagrams.md" >}}) ← イマココ

## UML の[クラス図]を使った記法

実体関連モデルは以下の3つの要素で表すことができる。

- 実体（entity）
- 関連（relationship）
- 属性（attribute）

「実体」同士の関係は「関連」で表すことができる。
また「実体」および「関連」はそれぞれ「属性」を持つことができる。

これを見てピンとくる方もいるだろうが，実は実体関連図は UML の[クラス図]を使って表すことができる。
「実体」を表すのがクラスであり，クラス同士を繋いだ関係線が「関連」に相当する。
たとえばこんな感じ。

```text
@startuml

class "Entity A" as A {
  {field} id : int <<PK>>
  {field} name : string
  {field} location : string
}
class "Entity B" as B {
  {field} subid : int <<PK>>
  {field} id : int <<PK,FK>>
  {field} subname : string
}

A "1" <.left. "0..n" B : more info >

@enduml
```

{{< fig-img src="./entity-relationship-1.png" link="./entity-relationship-1.puml" width="973" >}}

これで `Entity B` が `Entity A` の依存エンティティであることが分かる（これに対して外部キーを含まない `Entity A` を独立エンティティと呼ぶ）。
またステレオタイプを使って `PK` や `FK` を表現している。
更に多重度をつかって関連の「濃度（cardinality）」も表現できる。
更に更に [PlantUML] におけるクラスの属性はかなりアドホックな記述ができるため初期値や `NOT NULL` 制約のような付属情報も追記できるだろう。

## IE (Information Engineering) 記法

[PlantUML] は IE 記法にも一部対応している。
たとえばこんな感じ。

```text
@startuml

entity "Entity A" as A {
  id : int
  --
  name : string
  location : string
}
entity "Entity B" as B {
  subid : int
  id : int <<FK>>
  --
  subname : string
}

A ||.right.o{ B : more info >

@enduml
```

{{< fig-img src="./entity-relationship-2.png" link="./entity-relationship-2.puml" width="973" >}}

実体の属性は `--` や `==` や `..` を使って任意に仕切ることができるが，通常は2つに分けて上半分を主キー属性，下半分を非キー属性とする。

関連を示す線は “Crow's Foot” と呼ばれる記法で [PlantUML] では

```text
Entity01 }|..|| Entity02
Entity03 }o..o| Entity04
Entity05 ||--o{ Entity06
Entity07 |o--|| Entity08
```

{{< fig-img src="./crows-foot.png" link="./crows-foot.puml" width="1140" >}}

という感じに記述して濃度を表現する。
これは多重度表現の

```text
Entity01 "1..n" .. "1" Entity02
Entity03 "0..n" .. "0,1" Entity04
Entity05 "1" -- "0..n" Entity06
Entity07 "0,1" -- "1" Entity08
```

{{< fig-img src="./multiplicity.png" link="./multiplicity.puml" width="1140" >}}

と同等である。

なお [PlantUML] では IE 記法の依存エンティティを表現できない。
したがって先ほどの

{{< fig-img src="./entity-relationship-2.png" link="./entity-relationship-2.puml" width="973" >}}

において `Entity A` と `Entity B` のどちらが依存エンティティなのか分かりにくい（まぁ属性を見れば分かるけど）。
独立エンティティと依存エンティティを明示的に区別したいのであれば，苦肉の策ではあるが UML の[クラス図]と IE 記法を混ぜて

```text
@startuml

entity "Entity A" as A {
  id : int
  --
  name : string
  location : string
}
entity "Entity B" as B {
  subid : int
  id : int <<FK>>
  --
  subname : string
}

A "1" <.left. "0..n" B : more info >

@enduml
```

{{< fig-img src="./entity-relationship-3.png" link="./entity-relationship-3.puml" width="973" >}}

などとするのがいいかもしれない。
その辺はお好みで。

## ブックマーク

- [Entity Relationship diagram syntax and features](http://plantuml.com/ie-diagram) : 公式ドキュメント
- [PlantUMLでER図を描く！ - VELTRA Engineering - Medium](https://medium.com/veltra-engineering/how-to-draw-er-diagram-with-plantuml-86ec2095645e)
- [ER図を書くのに疲れたら - Qiita](https://qiita.com/genzouw/items/23cd0119715403e6e110)
- [GitHub - achiku/planter: Generate PlantUML ER diagram textual description from PostgreSQL tables](https://github.com/achiku/planter)

- [GitHub - spiegel-im-spiegel/plantuml-sample: Samples for PlantUML](https://github.com/spiegel-im-spiegel/plantuml-sample)

[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."
[クラス図]: {{< ref "/remark/2018/12/plantuml-3-class-diagrams.md" >}} "真面目に PlantUML (3) : クラス図"
[`skinparams.iuml`]: https://github.com/spiegel-im-spiegel/plantuml-sample/blob/master/skinparams.iuml "plantuml-sample/skinparams.iuml at master · spiegel-im-spiegel/plantuml-sample"
