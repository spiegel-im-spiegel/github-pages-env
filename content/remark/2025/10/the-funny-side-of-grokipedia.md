+++
title = "Grokipedia とやら（笑）"
date =  "2025-10-29T18:32:36+09:00"
description = "これに関する報道がめっちゃ面白い。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering", "politics" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++


[「宇宙の本質を理解する」](https://www.bloomberg.co.jp/news/articles/2023-07-12/RXP9DBT0G1KW01 "イーロン・マスク氏、オープンＡＩに対抗－新会社xAIの設立発表 - Bloomberg")ことを目標として某マスク氏が設立した企業 [xAI] が [Grokipedia] なるものをローンチしたらしい。

正直に言って，同じく [xAI] による [Grok] について私はあまり評価してないので [Grokipedia] 自体も関心が薄いのだが，これに関する報道がめっちゃ面白い。

最初に見かけたのがこれ：

- [“イーロン・マスク版Wikipedia”初期版公開 特徴はAIファクトチェック 「すでにWikipediaより優秀」 - ITmedia NEWS](https://www.itmedia.co.jp/news/articles/2510/28/news077.html)

記事を読めばわかるが「すでにWikipediaより優秀」と言ってるのは某マスク氏で，当然これに対する批判もあるのだが，このセリフを見出しに持っていくことで如何にも Grokipedia を礼賛しているように見えるのが面白い。
あれかな。
炎上マーケティングでも狙ってるのかな（笑）

こうなると Wikipedia と比較するような記事も出てくる。
マーケティング的には計画通り？

- [イーロン・マスクがWikipediaより優秀と豪語する百科事典サイト「Grokipedia」が一般公開される、Wikipediaから丸ごとコピーした記事も多数 - GIGAZINE](https://gigazine.net/news/20251028-grokipedia-xai/)
- [AIが事実確認をおこなう百科事典、Grokpediaが登場 | gihyo.jp](https://gihyo.jp/article/2025/10/grokpedia-v0.1)

ゴシップネタ大好きな GIGAZINE の記事はともかく gihyo.jp 記事の見出しにある「AIが事実確認」て（笑） なんだその面白ワードは。

{{< fig-quote type="markdown" title="AIが事実確認をおこなう百科事典、Grokpediaが登場" link="https://gihyo.jp/article/2025/10/grokpedia-v0.1" >}}
用語の説明内容の正しさを検証するために、同社のAI、Grokによる事実確認が行われており、ページではいつ事実確認を行ったかの情報が記されているのが特徴だ（また、各ページの文章の作成などにもGrokが使われているはずだが、現在のところ、このあたりの詳しい公式情報が一般向けに案内されていない⁠）⁠。
{{< /fig-quote >}}

何を以って「事実」としてるのかね？ AI の言うことは正しい？ 笑かす。

また Wikipedia の記事を丸パクリなものも結構あるそうだ。

{{< fig-quote type="markdown" title="イーロン・マスクがWikipediaより優秀と豪語する百科事典サイト「Grokipedia」が一般公開される、Wikipediaから丸ごとコピーした記事も多数" link="https://gigazine.net/news/20251028-grokipedia-xai/" >}}
いろいろ検索してみたところ、「[宮崎県](https://grokipedia.com/page/Miyazaki_Prefecture)」「[大分県](https://grokipedia.com/page/%C5%8Cita_Prefecture)」「[たこ焼き](https://grokipedia.com/page/Takoyaki)」「[お好み焼き](https://grokipedia.com/page/Okonomiyaki)」「[Firefox](https://grokipedia.com/page/Firefox)」「[GIMP](https://grokipedia.com/page/GIMP)」など多数の記事がWikipediaと同一の内容でした。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="AIが事実確認をおこなう百科事典、Grokpediaが登場" link="https://gihyo.jp/article/2025/10/grokpedia-v0.1" >}}
なお、Grokpediaには、Wikipediaの内容を利用したページもあり、そのページでは、ページ末尾にWikipediaの内容を利用した旨が記載されている（Wikipediaのページの文章のライセンスは基本的にCC BY-SA 4.0で、適切なクレジットの表示・ライセンスの継承が必要になる⁠）⁠。
{{< /fig-quote >}}

ちなみに “Creative Commons [Attribution-ShareAlike 4.0](https://creativecommons.org/licenses/by-sa/4.0/ "Deed - Attribution-ShareAlike 4.0 International - Creative Commons") License” ({{% cc-syms "cc" "by" "sa" %}}) は copyleft 的なライセンス構成で，実際に [GNU GPLv3 との互換性もある](https://creativecommons.org/share-your-work/licensing-considerations/compatible-licenses/ "Compatible Licenses - Creative Commons")。
上の記事で例示されている “[Okonomiyaki](https://grokipedia.com/page/Okonomiyaki "Okonomiyaki | Grokipedia")” を見ると，元の Wikipedia 記事へのリンクも張られてるし，ええんちゃうかなという気がしないでもない。
生成 AI で copyleft なコンテンツを「{{% ruby "adapt" %}}翻案{{% /ruby %}}」するのって，まさに「ウイルス的」で面白そうだ[^cc1]。

[^cc1]: Creative Commons および CC Licenses については，拙文ではあるが「[改訂3版： CC Licenses について]({{< rlnk "/cc-licenses/">}})」セクションの記事を参考にどうぞ。

日本の新聞系メディアは「[ウィキペディアの記事はリベラルな左派寄りに偏っているとして、自身らの保守的な考えを反映させる](https://www.nikkei.com/article/DGXZQOGN282UM0Y5A021C2000000/ "イーロン・マスク氏、Wikipedia対抗サイト「Grokipedia」開始 記事が「左派に偏り」と批判 - 日本経済新聞")」などと報じている。
自分で Grokipedia がプロパガンダ・コンテンツだって言っちゃってるよ（笑）

一方，最近私がよく参照する英語圏のメディアである [404 Media の記事](https://www.404media.co/grokipedia-is-the-antithesis-of-everything-that-makes-wikipedia-good-useful-and-human/ "Grokipedia Is the Antithesis of Everything That Makes Wikipedia Good, Useful, and Human")では

{{< fig-quote type="markdown" title="Grokipedia Is the Antithesis of Everything That Makes Wikipedia Good, Useful, and Human" link="https://www.404media.co/grokipedia-is-the-antithesis-of-everything-that-makes-wikipedia-good-useful-and-human/" lang="en" >}}
It is a fully robotic, heartless regurgitation machine that cynically and indiscriminately sucks up the work of humanity to serve the interests, protect the ego, amplify the viewpoints, and further enrich the world’s wealthiest man.
[...]
It is a totem of what Wikipedia could and would become if you were to strip all the humans out and hand it over to a robot; in that sense, Grokipedia is a useful warning because of the constant pressure and [attacks by AI slop purveyors](https://www.404media.co/jimmy-wales-wikipedia-ai-chatgpt/) to push [AI-generated content into Wikipedia](https://www.404media.co/wikipedia-pauses-ai-generated-summaries-after-editor-backlash/).
{{< /fig-quote >}}

と痛烈である。
要するに Grokipedia の存在自体が， Wikipedia がビッグテック & 生成 AI に汚染されたらこうなる，という反面教師だというのだ。

ただ，その後に続く

{{< fig-quote type="markdown" title="Grokipedia Is the Antithesis of Everything That Makes Wikipedia Good, Useful, and Human" link="https://www.404media.co/grokipedia-is-the-antithesis-of-everything-that-makes-wikipedia-good-useful-and-human/" lang="en" >}}
One needs only spend a few minutes clicking around the launch version of Grokipedia to understand that it lacks the human touch that makes Wikipedia such a valuable resource.
{{< /fig-quote >}}

を読んで Wikipedia が登場した頃によく言われた批判を思い出した。
曰く「Wikipedia は百科事典未満の品質だ」というものだ。
各分野の専門家が吟味して編纂した百科事典が，その辺の有象無象が好き勝手絶頂に書き散らかした Wikipedia に負けるわけないやろ。
というわけ。

若い頃（IT 業界に入る前）に百科事典の訪問販売をしてたこともある私としては[^e1] この辺の議論には思うところもあるんだけど，ぶっちゃけ，どっちもどっちなんだよ。
とするなら Grokipedia も「どっちもどっち」のひとつに過ぎないって感想になる。

[^e1]: 百科事典の訪問販売してたけど，売るのは下手くそだった。まぁ，そこでセールスマンとして才能開花してたら IT 業界には入らなかったろうけど（笑）

百科事典や Wikipedia に書かれている内容って「情報」であって「知識」ではない。
たとえば百科事典は多くの専門家が関わって彼らが持ってる知識を基に編纂されているのかも知れないが，記述自体はその知識から切り離された情報でしかない。
一方で情報を享受する側からすれば，自身が持つ知識と照らし合わせることで初めて（事実であれフェイクであれ）その情報に価値が生まれるのだ（フェイクと認識するなら，フェイクが記述されているということに意味が生まれる。陰謀論とかw）。
教科書や百科事典に書かれている情報を鵜呑み丸呑みするのは小学生まで（笑）

私は昭和時代の年寄りなので AI と聞くとどうしても「エキスパートシステム」を連想してしまう。
エキスパートシステムは一種の知識データベースで，最初から人類の構築した知識体系を機械の中に構築・再現するという目標があった。

今世紀に入って大きく変わったのは「情報とルールを組み合わせれば勝手に知識体系が構築されるんじゃね？」という考え方だろう。
これはクラウドという情報ダムが台頭したことが大きいと思う。
2010年代に話題になった IBM の Watson とか典型だ。
日本の東ロボくんもこれに入れていいかも知れない。
これらは前世紀に夢見たエキスパートシステムとは似て非なるものだ。

今の生成 AI ブームもこれの延長線上にあるもので，人の記憶と認知の仕組みを一部模倣することで，[人生相談](https://fukushishimbun.com/series08/42282 "「相談先はAI」３割 生きづらいこども調査〈自殺対策支援センターライフリンク〉 - 福祉新聞Web")してしまうほど人間臭いインターフェイスになった。
機械翻訳の進歩にも大きく寄与してるよね。
しかし，実際に生成 AI がやってることは「翻案の大量生産」であり，そこに自律性[^ai1] はない。
ある意味で「情報を（ルールで）いくら組み合わせても知識にはならない」ということを証明してしまってるんじゃないだろうか。

[^ai1]: 『[そろそろ、人工知能の真実を話そう](https://www.amazon.co.jp/dp/B071FHBGW8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』では「自律というのは哲学的な意味であり、自らが行動する際の基準と目的を明確を持ち、自ら規範を作り出すことができることをいう」と解説している（[参考]({{< ref "/remark/2017/09/the-myth-of-the-singularity.md" >}} "『シンギュラリティの神話』を読む")）。

「[AGI 実現まであと10年](https://note.com/trans_n_ai/n/n11cf35257ff8 "Andrej Karpathy最新インタビュー：AGI実現まであと10年、強化学習は実はかなり不十分、AGIは経済の劇的な発展にはつながらない｜Trans-N")」という話もあるみたいだが，まだ何か重要なピースが足りないのではないかという気がしている。
それが10年以内に埋まって AGI が登場するかどうかは分からないが。
まぁ，知識なんかどうでもよく，耳障りのいい「クイズ解答マシン」が欲しいだけなら，今のままスケールを上げていけばいつかは AGI になるかも知れない。

というわけで Grokipedia に関する報道が面白い，というお話でした。

まったくの余談だが，最近 “rnicrosoft.com” ドメインからの phishing メールについて「[老眼の人を狙い撃ち](https://www.itmedia.co.jp/news/articles/2510/27/news106.html "老眼の人を狙い撃ち？　「microsoft」ならぬ「rnicrosoft」からの詐欺メールが話題に - ITmedia NEWS")」とか失礼な記事（笑）を見かけたが， Grokipedia の “r” をうっかり見失って{{% ruby "イニシャル" %}}頭文字{{% /ruby %}}Gなサービスかと思ったというポストにはちょっと共感しそうになった。

## ブックマーク

- [アンドレイ・カーパシーの最新インタビューに見るAI技術の現在「AGIの実現まで10年はかかる」 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20251027/andrej-karpathy)

[xAI]: https://x.ai/ "xAI"
[Grok]: https://grok.com/ "Grok"
[Grokipedia]: https://grokipedia.com/ "Grokipedia"
