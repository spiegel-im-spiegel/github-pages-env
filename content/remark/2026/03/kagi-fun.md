+++
title = "Kagi 翻訳によるおもしろ翻訳"
date =  "2026-03-20T10:23:14+09:00"
description = "儀礼的な文章・会話がいかにアホみたいかよく分かるよね。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "web", "kagi" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## Kagi 翻訳によるおもしろ翻訳

先日の [Kagi Search 変更履歴](https://kagi.com/changelog "変更履歴 - Kagi Search")に

{{< fig-quote type="markdown" title="March 19th, 2026 - Small Web Expansion and Translate goes viral" link="https://kagi.com/changelog#10140" lang="en" >}}
On March 16, we launched our latest fun language on Kagi Translate, [LinkedIn Speak](https://translate.kagi.com/?from=en&to=LinkedIn+speak), and it quickly went viral on social media, generating millions of engagements.
{{< /fig-quote >}}

とあり，早速試してみた。

{{< fig-img-quote src="./kagi-translate.png" link="./kagi-translate.png" width="1174" >}}

いや「Z世代」とか（笑） これは試してみるしかあるまい。

{{< fig-img-quote src="./kagi-translate-z.png" link="./kagi-translate-z.png" width="1174" >}}

んー。
面白いけど普通かな。
じゃあ「LinkedIn Speak」だとどうなるか。

{{< fig-img-quote src="./kagi-translate-linkedin.png" link="./kagi-translate-linkedin.png" width="1174" >}}

ぶはは！ あれだ。
日本で言う「おじさん構文」だな，これ。
調子に乗って「ビジネス用語」でいってみよう。

{{< fig-img-quote src="./kagi-translate-business-1.png" link="./kagi-translate-business-1.png" width="1174" >}}

なげーよ！ 確かにビジネス挨拶だわ（笑） つか spam メールの序文だな。
もっと簡単に “Hi” だとどうなるか。

{{< fig-img-quote src="./kagi-translate-business-2.png" link="./kagi-translate-business-2.png" width="1174" >}}

同じだ同じ。
漫画とかで元の会話と翻訳でボリュームが全く釣り合ってないというシチュエーションがあるけど，まさにこれだな。
ちなみに上の翻訳文を更に日本語に訳すとこうなる。

{{< fig-img-quote src="./kagi-translate-business-jp.png" link="./kagi-translate-business-jp.png" width="1174" >}}

メールでこんな文章来たらソッコーでゴミ箱行きだな。
まぁ，儀礼的な文章・会話がいかにアホみたいかよく分かるよね。

## 【おまけ】 Kagi 翻訳に “Hi” を日本語に訳させてみる

[Kagi 翻訳][Kagi Translate]は翻訳に関して細かくチューニングすることができる。
たとえば，先程の “Hi” をデフォルト設定で日本語に訳すとこうなる。

{{< fig-img-quote src="./kagi-translate-hi-jp1.png" link="./kagi-translate-hi-jp1.png" width="1020" >}}

まぁ，無難。
最後の「最近どう？」は意訳しすぎやろと思うけど（笑）

ここで以下のように設定を変えてみる。

{{< fig-img-quote src="./kagi-translate-hi-conf.png" link="./kagi-translate-hi-conf.png" width="1020" >}}

1. 謙譲語
2. 話し手および聞き手はともに男性
3. 勤務先で部下から上司への挨拶

この設定で翻訳させると以下のようになる。

{{< fig-img-quote src="./kagi-translate-hi-jp2.png" link="./kagi-translate-hi-jp2.png" width="1020" >}}

コンテキスト通り日本の職場っぽくなった。
日本語って同じ意味でもシチュエーションによって全然違う表現になったりするので，こういう翻訳オプションは面白い。

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
