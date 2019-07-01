+++
date = "2015-10-11T23:06:14+09:00"
update = "2015-10-22T08:46:00+09:00"
description = "TPP 大筋合意 / 「クリエイティブ・コモンズ・ライセンスについて」改訂版 / Git for Windoes 2.6.1"
draft = false
tags = ["code", "politics", "tpp", "intellectual-property", "copyright", "creative-commons", "git"]
title = "今日の戯れ言：TPP 大筋合意"

[scripts]
  mathjax = false
  mermaidjs = false
+++

1. [TPP 大筋合意]({{< relref "#tpp" >}})
1. [「クリエイティブ・コモンズ・ライセンスについて」改訂版]({{< relref "#cc" >}})
1. [Git for Windoes 2.6.1]({{< relref "#git" >}})

## TPP 大筋合意{#tpp}

- [TPP大筋合意との報に際して « マガジン航](http://magazine-k.jp/2015/10/10/soramoyou-tpp/)

さてさて，困ったもんですね。

{{< fig-quote title="TPP大筋合意との報に際して" link="http://magazine-k.jp/2015/10/10/soramoyou-tpp/" >}}
<q>TPPに関して今後、条約の締結や国内法の整備などが進められていくことでしょう。とはいえそのなかで、ひとりひとりが粘り強く声を上げ、自分たちの文化がどうあるべきなのか、あきらめずに議論を続けることも必要です。<br>
今ようやく芽生えてきたパブリックドメインによる豊かで多様な共有文化が損なわれないような、柔軟な著作権のあり方を切に望みます。</q>
{{< /fig-quote >}}

という点には全く同意。
日本の国会は，一度法案が上がると撤回するのは難しい。
しかし，まだ「合意」しただけで「条約」すらない状態なのだから，必要なことに関しては今のうちからきちんと声を上げて行くことが必要。

## 「クリエイティブ・コモンズ・ライセンスについて」改訂版{#cc}

そういうわけでもないのだが，[本家サイト](https://baldanders.info/)の「[クリエイティブ・コモンズ・ライセンスについて](https://baldanders.info/spiegel/cc-license/)」の改訂版を作成中[^a]。

[^a]: [Hugo] の Section 機能を使うとこういった特集記事を組みやすい。いやぁ，ホンマに [Hugo] にしてよかったよ。

私はその筋の専門家ではないので，今まで著作権や著作権法に大きく踏み込む記述は（敢えて）しなかったのだが，それだと CC Licenses の背景や Free Culture について説明しきれない気がして。
本当は，こういうのって [creativecommons.jp] がすべきなんだろうけど， FAQ もいまだに古いままだし， Free Culture の F の字もないし[^b]。

[^b]: 理事である [Dominique Chen さん](http://creativecommons.jp/about/people/#chen)の本は出たけど，それっきり。まぁ彼の関心領域も変わっちゃったみたいだからなぁ。

今年は CC0 や CC Licenses の日本語訳を公開したりして [creativecommons.jp] の成果には敬意を払うものだけど，どうも「ライセンス・ツールを提供すれば終わり」みたいな態度が透けていて，活動母体である [commonsphere] もただの「寄せ鍋」状態で活動実態がまるで見えない。
この前まで本家 [creativecommons.org] がやってた crowd-funding に連動するような動きもなかったし，今月あるソウルでの Global Summit で何かするってわけでもないようだし。
せめて本家 [creativecommons.org] や [EFF] の活動くらいはちゃんとフォローして欲しいけど，そういうのもないからなぁ。

個人的には [creativecommons.jp] はもう死んだものとして本家 [creativecommons.org] の活動のみを追いかけていく所存です（英語不得手なのに orz）。

改訂版はこれから[徐々に公開](/cc-licenses)していく。
[作業環境 repository](https://github.com/spiegel-im-spiegel/github-pages-env "spiegel-im-spiegel/github-pages-env") には草稿版があるので興味のある方はどうぞ。

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[creativecommons.jp]: http://creativecommons.jp/ "クリエイティブ・コモンズ・ジャパン"
[creativecommons.org]: https://creativecommons.org/ "Creative Commons"
[commonsphere]: http://commonsphere.jp/ "commonsphere | コモンスフィア"
[EFF]: https://www.eff.org/ "Electronic Frontier Foundation | Defending your rights in the digital world"

## Git for Windoes 2.6.1{#git}

- [Release Git for Windows 2.6.1 - git-for-windows/git](https://github.com/git-for-windows/git/releases/tag/v2.6.1.windows.1)

特に問題なさそう。
脆弱性や大きな瑕疵がない限り無理に追随する必要はないのかもしれないけど，まぁ一応。
