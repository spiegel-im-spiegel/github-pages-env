+++
title = "技術的負債としての「オブジェクト指向」"
date =  "2019-07-28T20:47:23+09:00"
description = "「オブジェクト指向」が云々なんてのは些末なことで，その時のトレンドによって要求は変わるし，要求の変化に伴って設計やプログラミングも変化する。そのための DevOps や OODA や PDCA なんじゃないの"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "object-oriented", "design", "programming" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

「[オブジェクト指向プログラミングは1兆ドル規模の大失敗なのか？ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20190723/isoopdisaster)」より。

- [Object-Oriented Programming — The Trillion Dollar Disaster](https://medium.com/codeiq/-92a4b666c7c7)
- [ブログ: オブジェクト指向プログラミング -- 1兆ドル規模の大失敗](https://okuranagaimo.blogspot.com/2019/07/1.html)

邦訳を公開してくださった方，マジ感謝です。
英語不得手な私には大変ありがたい。

「1兆ドル」という見積もりが妥当かどうかはともかく[^ns1]，オブジェクト指向設計あるいはプログラミングが究極ではないという意味なら同意できるし「その先」に進むべきというのも分かる。
更にこれからプログラミング言語を学ぶなら関数型を学べという点についても激しく同意する。

[^ns1]: 「null 参照」の損失を[10億ドルと見積った人](https://en.wikipedia.org/wiki/Tony_Hoare "Tony Hoare - Wikipedia")もいたな。こういう{{< ruby "経済アナリスト" >}}競馬の予想屋{{< /ruby >}}みたいなやり口は好かんのだがなぁ。

{{< fig-quote type="md" title="Java はやめておけ" link="/remark/2016/07/java/" >}}
{{< quote >}}どういうわけか日本人は Java が大好きで，確かにあと数年は飯の種になるだろうが，将来を見据えれば Java は間違いなく COBOL と同じ道をたどる。{{< /quote >}}
{{< /fig-quote >}}

そういう意味で Java のようなレガシーなプログラミング言語が将来的な技術的負債になるのは間違いないと思う。
まぁ[「みずほ」のような大災害](https://note.mu/tsukamoto/n/n66e1be6b616b "みずほ銀行の新基幹システムが全面稼働へ｜塚本 牧生｜note")はそうそう起こらないだろうけど（笑）

でもそれが「オブジェクト指向」故の問題かと言われれば，ちょっと考えるとところでもある。

## オブジェクト指向設計とオブジェクト指向プログラミング

ひとくちに「オブジェクト指向」と言うが「設計」と「プログラミング」は分けて考える必要がある。
「プログラミング」はどうしても言語仕様に依存してしまうが，「設計」は言語に依存しないしするべきではない。

研究分野についてはよく知らないのでスルーするが，産業分野に於いて「オブジェクト指向」とは制御単位としての「オブジェクト」と「オブジェクト」同士の関係を「記述」することだ。
つまり「制御」を主体に設計するならば「オブジェクト」そのものに着目せざるを得ない。

これが，それまでの「構造化プログラミング」からの発展としての「オブジェクト指向」である。

### さぁ，例外（おまえ）の罪を数えろ！

私が数年前に [Go 言語]を習い始めて驚いたのは「例外がない」ことだった[^pa1]。

[^pa1]: [Go 言語]には panic という「例外」に似た仕組みはあるが，これは文字通りプログラム続行不可能な状態で強制終了させるための仕組みであり「例外」のように運用すべきではないとされている。

{{< fig-quote type="md" title="エラー・ハンドリングについて" link="/golang/error-handling/" >}}
{{< quote >}}例外が抱える問題というのは本質的に goto 文の問題と同じ。 オブジェクトの状態ごと脱出するため，（脱出前ではなく）脱出後にオブジェクトの後始末を記述せざるを得ないし，記述するためには脱出前の状態（の可能性）を「知識」として知っていなければならない。 もし後始末をきちんとしないと，それがバグやリークやその他の脆弱性のもとになる。{{< /quote >}}
{{< /fig-quote >}}

例外処理を行うためにはオブジェクトの内部状態に外部からアクセスする手段を実装する必要がある。
これが「オブジェクト指向」プログラミングを狂わせる要因のひとつになっている，と思う。

[Go 言語]ではエラー・ハンドリングの問題を tupple と defer 構文というシンプルな言語仕様によって解決している。
これを知ったときはマジで感動したね。
ほんでもって，私が [Go 言語]に傾倒する理由のひとつになっている。

### 具象から抽象へ記述する

これも[以前に書いた]({{< ref "/remark/2017/03/generics-vs-duck-typing.md" >}} "きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」")ことだが [Go 言語]では汎化・特化の関係を記述する手段として interface 型が存在する。
Interface 型は [duck typing](https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia") と勘違いされることもあるが（私もそう思っていた），実は「構造的部分型（structural subtyping）」と呼ばれるもので，オブジェクトの「ふるまい」のみを定義する[^dt1]。

[^dt1]: [Duck typing](https://en.wikipedia.org/wiki/Duck_typing "Duck typing - Wikipedia") は主に動的型付け言語における型推論方式（のひとつ）で，クラス間の関係を記述するものではないらしい。ちなみに構造的部分型に対する言葉として「公称型（nominal subtyping）」というのがあって C++ や Java におけるテンプレート・クラスやインタフェース・クラスを使った汎化・特化関係を指す場合に使うそうだ。

{{< fig-quote type="md" title="きみは Generics がとくいなフレンズなんだね，または「制約は構造を生む」" link="/remark/2017/03/generics-vs-duck-typing/" >}}
{{< quote >}}たとえばウォータフォール型の開発スタイルでは実装を開始するまでに設計が終わることが（建前上は）保証されているため「抽象→具象」へと書き進めることが容易な言語が向いている。一方，要件が絶えず変わったり実験的な製品の場合は設計が終わるまで待っていられないため Go 言語のような言語が向いてるかもしれない。{{< /quote >}}
{{< /fig-quote >}}

私が「プログラマも要件定義から参加すべき」と思うようになったのは [Go 言語]を知ってからだ。
他の流行りの言語でもそうだろうが，今やプログラミングは文芸的になりつつありプログラマ間のコミュニケーション「言語」にもなりつつある。

### Don't Talk to Strangers

「オブジェクト指向」設計およびプログラミングには {{< quote lang="en" >}}don't talk to strangers{{< /quote >}} の原則があり，いわゆる「友達の友達」のことは知らないふりをするのがよい設計と言われている。
つまり「オブジェクト指向」で重要なのは「オブジェクト」ではなく，その「関係」であるというわけだ。

{{< fig-quote title="Object-Oriented Programming — The Trillion Dollar Disaster" link="https://medium.com/@ilyasz/-92a4b666c7c7" lang="en" >}}
<q>I’m sorry that I long ago coined the term “objects” for this topic because it gets many people to focus on the lesser idea. The big idea is messaging.<br>
- Alan Kay, the inventor of OOP</q>
{{< /fig-quote >}}

引用の引用でスマン。

たとえば REST (REpresentational State Transfer) と呼ばれる分散システムにおける設計原則がある。
でも REST って上で言うところの「オブジェクト指向」そのものだよね。
REST の背景には多分「オブジェクト指向」の考え方があるのかもしれないが，そもそも分散システムの設計における必然から生まれたものでもある。

## 結局のところ...

「オブジェクト指向」が云々なんてのは些末なことで，その時のトレンドによって要求は変わるし，要求の変化に伴って設計やプログラミングも変化するのだから拘っても仕方がない。

大事なことは，そうした「変化」のタイムスケールが，かつて感覚で考えていたよりもずいぶん小さいものであり，サボることなく対応していかなければあっという間に技術的負債になってしまうということだろう。
そのための DevOps や OODA や PDCA なんじゃないの，と思うわけだ。

まぁ，既に職業エンジニアを廃業し現場から遠ざかったロートルが言うことじゃないけどさ。

## ブックマーク

- [「オブジェクト指向」の黒歴史]({{< ref "/remark/2018/10/object-oriented-design.md" >}})

[Go 言語]: https://golang.org/ "The Go Programming Language"
