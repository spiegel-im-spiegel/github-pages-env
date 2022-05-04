+++
title = "技術的負債と FOSS"
date =  "2021-07-26T21:13:02+09:00"
description = "不完全でおｋ"
image = "/images/attention/kitten.jpg"
tags = [ "code", "engineering", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回の戯れ言の起点はここから。

- [The end of open source? – TechCrunch](https://techcrunch.com/2021/07/18/the-end-of-open-source/)
- [オープンソースの終焉？ ではなく次代の（技術、ガバナンス）モデルに移るべきという話 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20210726/the-end-of-open-source)

発端は2021年春頃に発覚した話で，知っている人は知っていると思うが簡単に言うと，どこぞの大学関係者が論文のネタに Linux Kernel にわざと脆弱性を含むパッチを提供し，それがどのように解消されていくか（あるいはされないのか）を観察するという，なかなかに鬼畜なことをやらかしたらしい。

私もこの話は知っていたが，あまり重視していなかった（[ブックマーク]({{< rlnk "/bookmarks/" >}})にも入れてなかったし）。
いや，これって大学関係者がやらかしたから物珍しかっただけで[^acd1]，脆弱性のあるコードの混入自体は別に珍しくないだろ。
悪意や害意はないにしても Ubuntu の APT ベースで月に何回 Kernel を更新してるんだって話である。

[^acd1]: 今回の話を聞いて私はむしろ「ありがち」と思ったよ。大学が「善」なわけない。真理に善悪はないのだから。たとえば1990年代の米国暗号シーンに於いて大学は何をしてたのかってことだし，もっと言うと，ゼロ年代の Winny 狂躁は（お上による弾圧のほうがクローズアップされるが）明らかに「余計なこと」で，日本はアレで P2P に関して周回遅れになったし，日本企業が「リモートワーク」に及び腰なのはゼロ年代のこの悪夢が尾を引いてるよね。

伽藍であれバザールであれ，隠された悪意や悪意なき害意を見つけるのは困難である。
こういうのは人的な統制をしても無駄で，結局はコードで証明していくしかない。
まぁ，昔に比べれば典型的なバグや脆弱性の検出についてはずいぶん楽になったけどね。

大昔にソフトウェアを「[大量生産される一品もの](https://baldanders.info/spiegel/log/200511.html#d20_t1)」と評したことがある。
当時は XaaS とか考慮してなかったので当てはまらないことも多いと思うが，私は製品を作るときは「開発」と「製造」を分けるべきだと思っている。
何故ならそれぞれに要求される技術要素が異なるからだ。
でもソフトウェアはしばしば「開発」と「製造」が混濁する。

Linux Kernel の場合は（営利企業を主体とする）ディストリビュータが「製造」で要求される品質を担保していると思っていたが，そういうわけでもないらしい。

ところで，上の記事を見て連想したのが「技術的負債」というフレーズだ。
「技術的負債」については，同じく yomoyomo さんの記事が参考になる。

- [お前も技術的負債にしてやろうか！ もしくは技術的負債と和田卓人さんをめぐるシンクロニシティ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20201210/technical-debt)

「技術的負債」を，お金を前借りするように「より早く実現するために引き換えにするもの」とするなら， FOSS と「技術的負債」は本来は相性がいいと言える。
ただ，負債の清算（＝リファクタリング）が上手く回らないのであれば，何らかの見直しが必要だろう。
これは FOSS 全体というよりは個々の製品・サービス毎に考えるべきことだ。

FOSS 最大の特徴は「*不完全でおｋ*」という点だと思う。
不完全さを埋めようとする限り，その製品には伸びしろがある。
その変化のプロセスが重要なのである。

さて，最初に挙げた記事では進化する「オープンソースへの脅威」への対応策として

{{< fig-quote type="markdown" title="The end of open source?" link="https://techcrunch.com/2021/07/18/the-end-of-open-source/" lang="en" >}}
- Limit the spread of monocultures. Stuff like Alva Linux and AWS’ Open Distribution of ElasticSearch are good, partly because they keep widely used FOSS solutions free and open source, but also because they inject technical diversity.
- Reevaluate project governance, organization and funding with an eye toward mitigating complete reliance on the human factor, as well as incentivizing for-profit companies to contribute their expertise and other resources. Most for-profit companies would be happy to contribute to open source because of its openness, and not despite it, but within many communities, this may require a culture change for existing contributors.
- Accelerate commodification by simplifying the stack and verifying the components. Push appropriate responsibility for security up into the application layers.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="オープンソースの終焉？ ではなく次代の（技術、ガバナンス）モデルに移るべきという話" link="https://yamdas.hatenablog.com/entry/20210726/the-end-of-open-source" >}}
- モノカルチャーの広がりを制限する。オープンソースであるというだけでなく、技術的多様性（technical diversity）の導入が重要
- プロジェクトのガバナンス、組織、資金を見直し、特定の人に完全に依存しないようにしながら、営利企業が専門知識などのリソースを提供する動機付けを行う
- スタックを単純化し、コンポーネントを検証してコモディティ化を促進。セキュリティ面のしかるべき責任をアプリケーション層に押し上げる
{{< /fig-quote >}}

とか書かれているが FOSS 関係なくね？ yomoyomo さんも「タイトルはいくらなんでも釣り」と呆れているが，そのとーり（by 財津一郎）だと思う。
最後のやつとか，普通にアーキテクチャ設計の基本の話だよね。

ちうわけで，そろそろ『[Clean Architecture](https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "Clean Architecture　達人に学ぶソフトウェアの構造と設計 (アスキードワンゴ) | Ｒｏｂｅｒｔ Ｃ．Ｍａｒｔｉｎ, 角 征典, 高木 正弘 | 工学 | Kindleストア | Amazon")』の読書感想文を書かなあかんかなぁ，と思ったりする（予定は未定）。

## ブックマーク

- [【コラム】オープンソースの終焉が来ているのだろうか？  |  TechCrunch Japan](https://techcrunch.com/2021/07/18/the-end-of-open-source/) : よーやく日本語記事が（笑）

- [伽藍とバザール（The Cathedral and the Bazaar）](http://cruel.org/freeware/cathedral.html)
- [技術的負債とハッカー]({{< ref "/remark/2020/12/technical-debt-and-hacker.md" >}})

## 参考文献

{{% review-paapi "4309242456" %}} <!-- リナックスの革命 Hacker Ethic -->
{{% review-paapi "B07FSBHS2V" %}} <!-- Clean Architecture -->
<!-- eof -->
