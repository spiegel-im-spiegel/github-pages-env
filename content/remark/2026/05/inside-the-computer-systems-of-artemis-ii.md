+++
title = "Artemis II 宇宙機のコンピュータシステム"
date =  "2026-05-20T08:22:59+09:00"
description = "こうした命に関わるような究極のシステムを設計することはそうそうないと思うが，知識として知っておくのもいいだろう。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "nasa", "engineering", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は Mastodon の自 TL に流れてきた記事を紹介してみる。
先ごろ月の裏側まで行った [Artemis II] ミッションのコンピュータシステムの話らしい。

- [How NASA Built Artemis II’s Fault-Tolerant Computer – Communications of the ACM](https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/)

アポロ計画の有人宇宙機が3機のコンピュータの多数決制をとっていたことは有名な話だが（スペースシャトルも多数決制だっけ？），あれから50年経った [Artemis II] 宇宙機の構成はだいぶ違うようだ。

そもそもの前提としてアポロ計画時代の宇宙機ではコンピュータはシステムの中核ではなく，機械制御や人間の操作がシステムの中心だったそうだ。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
Apollo astronauts navigated to the lunar surface using a computer with a 1-MHz processor and roughly 4 kilobytes of erasable memory, supported by a larger store of fixed “rope” memory. While it was a marvel of 1960s engineering, the Apollo Guidance Computer’s functional scope was focused and not in the control loop for every system. Critical environmental and power controls were managed through manual or electromechanical means, such as switches and relays.
{{< /fig-quote >}}

これに対して [Artemis II] の宇宙機オリオンは，フライトのみならず生命維持を含む全システムの制御をコンピュータが担う構成である。
このため，障害耐性（fault tolerance）においてガッチガチに固められた設計になっている。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
One of the biggest drivers for this redundancy is the harsh radiation environment of space, where high-energy particles can affect avionics and create ‘wrong answers’ that must be filtered out of the flight solution.
{{< /fig-quote >}}

ヴァン・アレン帯の外側の深宇宙では宇宙線が電子機器[^a1] に与える影響がより深刻になり，システムが誤った答えを出す可能性がある。
したがって，その誤りを検出し正す方法が必要になるというわけ[^a2]。
多数決制もそのひとつだったんだけどね。

[^a1]: 元記事にある “avionics” って何？ と思って調べたら “aviation”（航空）と “electronics”（電子機器）の合成語で航空機等に搭載される電子機器全般を指すらしい。
[^a2]: もちろん宇宙機としても宇宙線などの脅威に強い構造になってる筈だし，チップセットやネットワークシステムにおいても強い誤り訂正や冗長性の機能があるそうな。それでも完全ではないからソフトウェア上でも対策が必要になる。

たとえば，オリオンには2台の VMC (Vehicle Management Computer; 宇宙機管理コンピュータ[^sv1]) があり，それぞれに2つの FCM (Flight Control Module; 飛行制御モジュール) が含まれている。
さらに各 FCM は自己診断機能を持つ一対のプロセッサで構成されている。
つまりオリオンのフライトソフトウェアでは，合計8つのプロセッサが並列（parallel）に動作しているということらしい。

[^sv1]: “vehicle” はラテン語の “vehere”（運ぶ）に由来する単語で，自動車だけではなく，広く（運搬する）乗り物を指す言葉として使うようだ。たとえば宇宙機は “[space vehicle](https://en.wikipedia.org/wiki/Space_vehicle "Space vehicle - Wikipedia")” と呼んだりする。あと AI に教えてもらったのだが，ベクトル（vector）の語源も同じラテン語の “vehere” らしい。ベクトルは「ある地点から別の地点へ（情報を）運ぶもの」というニュアンスなんだね（故に方向と距離という2つの属性を持つ）。 AI にウンチクを語られるとは思わなかった（笑）

8つのプロセッサおよび4つの FCM は，多数決ではなく “fail-silent” と呼ばれる設計に基づいて動作している。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
“A faulty computer will fail silent, rather than transmit the ‘wrong answer,’” Uitenbroek explained. This approach simplifies the complex task of the triplex “voting” mechanism that compares results. Instead of comparing three answers to find a majority, the system uses a priority-ordered source selection algorithm among healthy channels that haven’t failed-silent. It picks the output from the first available FCM in the priority list; if that module has gone silent due to a fault, it moves to the second, third, or fourth.
{{< /fig-quote >}}

沈黙した FCM はそのまま捨て置かれるのではなく，リセットし再同期してグループに加わることができるようになってるらしい。

ここで現代の開発手法に対する微妙な dis りがあるのが面白かった。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
Michael Riley, a team lead at Carnegie Mellon’s Software Engineering Institute who previously collaborated with NASA to adapt risk-assessment tools for the Orion mission, noted that while earlier generations worked within strict hardware constraints, modern mission-critical development is different.

“Modern Agile and DevOps approaches prioritize iteration, which can challenge architectural discipline,” Riley explained. “As a result, technical debt accumulates, and maintainability and system resiliency suffer.”
{{< /fig-quote >}}

日本語で言い直すと， DevOps などの現代的な開発手法は反復を優先するため，アーキテクチャの規律を保つのが難しくなり，保守性やシステムの回復力が損なわれることがあるという指摘だ。
個人的には「やり方によるんじゃない？」とも思うが，孫子よろしく「兵は拙速を尊ぶ」でイテレーションを回すと，間違った方向に進んだときに修正が難しくなってしまうというのはあるかもしれない。
似たようなことを前にどっかで聞いたな。
[「回帰不能点」の話]({{< ref "/remark/2025/03/the-point-of-no-return.md" >}} "回帰不能点と死の行軍")だったか。

話を戻そう。

4つの FCM によるシステムは強力ではあるが，オリオンではこれらと独立したバックアップシステム（Backup Flight Software; BFS）を搭載している。
BFS はプライマリシステムとは異なるハードウェアと異なる OS で動作し，独立して開発された簡素な飛行ソフトウェアを利用しているそうな。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
“It is intentionally different to ensure that a common mode software failure in the primary flight software isn’t also implemented incorrectly on the backup,” Uitenbroek said. The BFS runs constantly in the background and automatically takes over via source selection if the primary computers fail. If the system finds itself on the BFS, it can complete all dynamic portions of the mission to reach a quiescent phase, at which point the crew can attempt to recover the primary FCMs.
{{< /fig-quote >}}

もっと言うと “dead bus” と呼ばれる全電源喪失シナリオにおいてもオリオン（およびその乗組員）が生き残れるよう設計されているらしい。

{{< fig-quote type="markdown" title="How NASA Built Artemis II’s Fault-Tolerant Computer" link="https://cacm.acm.org/news/how-nasa-built-artemis-iis-fault-tolerant-computer/" lang="en" >}}
the spacecraft enters a safe mode, in which the vehicle first stabilizes itself and then points its solar arrays at the Sun to recover power. Then, it orients its tail toward the Sun for thermal stability before attempting to re-establish communication with Earth. During such a failure, the crew can also take manual action to configure life support systems or don space suits.
{{< /fig-quote >}}

こうした命に関わるような究極のシステムを設計することはそうそうないと思うが，このような世界もあるのだ，ということを知識として知っておくのもいいだろう。

## ブックマーク

- [ネトフリに再契約しなくてもよかったのに]({{< ref "/remark/2026/04/shouldnt-have-resubscribed-to-netflix-for-artemis-ii.md" >}})

[Artemis]: https://www.nasa.gov/humans-in-space/artemis/ "Moon to Mars | NASA's Artemis Program - NASA"
[Artemis II]: https://www.nasa.gov/mission/artemis-ii/ "Artemis II: NASA’s First Crewed Lunar Flyby in 50 Years - NASA"
