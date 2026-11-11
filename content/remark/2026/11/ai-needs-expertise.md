+++
title = "AI を教師役にしてはいけないらしい"
date =  "2026-11-11T14:25:29+09:00"
description = "見識って難しいよな。単なる知識じゃないもん。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "security" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は Cory Doctorow さんの記事からはじめよう。

- [Pluralistic: Discernment (28 Jul 2026) – Pluralistic: Daily links from Cory Doctorow](https://pluralistic.net/2026/07/28/hitl-ers/)
- [見識 » p2ptk[.]org](https://p2ptk.org/ai/5674) (邦訳版)

ちなみに "discernment" を [Kotoba]({{< relref "./kotoba-dictionary.md" >}} "Rust 使いの日英文学翻訳者による辞書ツール") で引いたら以下が出た。

{{< fig-img-quote src="./discernment.png" title="Kotoba" link="./discernment.png" width="965" lang="en" >}}

いきなり引用しておこうか。

{{< fig-quote type="markdown" title="見識" link="https://p2ptk.org/ai/5674" >}}
ある主題を本当によく理解していれば――有用なボット出力と問題のある出力を見分けられるほどに理解していれば――ボットは役に立ちうる。ときにはきわめて有用に、たいていは並程度に有用に。

[...]

繰り返すが、チャットボットから有用な仕事を引き出すなら、見識こそがすべてである。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="見識" link="https://p2ptk.org/ai/5674" >}}
生産的なAI利用にとって見識が中心的な役割を果たすことは自明である。だからこそ私には、AIをティーチングアシスタントとして――最悪なのは教師として――活用できるのだという主張が、そら恐ろしいほど不可解に思える。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="見識" link="https://p2ptk.org/ai/5674" >}}
AIは普通のテクノロジーに過ぎない。つまり、役に立つ場面もあれば、無意味な場面や大いに有害な場面もあるということだ。チャットボットに関する経験則の一つは、LLMが必ず吐き出す欠陥のある出力を無視できる見識を備えた専門家に対してのみ、有用な情報を提供しうる、というものである。
{{< /fig-quote >}}

言う言う（笑）

「教師」がピンとこないなら「メンター」とか「カウンセラー」とかで置き換えて想像してみると分かりやすいかもしれない。
私はプログラマなのでそちらの比喩で考えるなら LLM は上級プログラマというよりは IT 土方って感じだろうか。
ゼロ年代に流行ったなぁ「IT 土方」ってワード。
泥臭いテストコードも文句言わずに書いてくれるので助かっている。

でも，機械（ソフトウェア）だからチューニングはできるが（自ら）成長はしない。

そういえば某通信会社が AI に自社システムを診断させたところ，15K件の脆弱性が見つかったので頑張って潰したのだが，その後，最新モデルで再チェックしたら更に35K件の脆弱性が見つかって `orz` な感じ[らしい](https://japan.cnet.com/article/35254620/ "ソフトバンク社長「自社システムをGPT-6 Astraで診断したら3.5万件の追加脆弱性を発見」と明かす - CNET Japan")。
多分，その35K件を潰し終わる頃には新しいモデルが出るよ，知らんけど。
もう完全にモグラ叩き状態（笑）

余談はさておき，問題はどうやってその「見識（discernment）」を身に付けるかである。
既に AI という道具がある状況で「私たちの世代を見習え」とは到底言えない。
でも何らかの形で「学び」は必要なんだよ。

{{< fig-quote type="markdown" title="見識" link="https://p2ptk.org/ai/5674" >}}
定義上、学生は自らが学んでいる主題の専門家ではない。だからこそ学んでいるのだ。知識を、ひいては見識を身につけるために学んでいる。チャットボットの説明を通じて学べと学生に求めるのは、道理に合わないどころか危険ですらある。
{{< /fig-quote >}}

見識って難しいよな。
単なる知識じゃないもん。
