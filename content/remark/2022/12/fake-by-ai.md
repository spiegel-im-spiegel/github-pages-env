+++
title = "人間だと思った？ 残念！ ボットちゃんでした"
date =  "2022-12-05T10:07:22+09:00"
description = "AI は考えないし感じない。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いや，人間同士の会話では分からないことを「分からない」で済まさないことも多いぢゃん。
まぁ，興味がないなら「分からない」で打ち切る手もあるけど，大抵は持ってる情報や知識から蓋然性の高い推論を立ててみたり，あるいは全くの思いつきを口にしたり，などするでしょう？ それに対する相手の反応を見ながら会話や議論を継続していく（もしくは継続を断念する）わけだ。

...

たとえば，先日 [ChatGPT] で遊んだときに

{{< fig-img-quote src="/remark/2022/12/debugging-with-ai/chatgpt-2.png" title="ChatGPT" link="https://chat.openai.com/" width="835" lang="en" >}}

が[面白かった]({{< relref "./debugging-with-ai.md" >}})という話をしたけど，これの面白さは，内容以上に AI が知識として持ってない（であろう）ことを，元ネタが分からないレベルの「創作」で以って補って提示している点だろう。
人間で言うなら「口からでまかせ」というところだろうが， AI はこれを「考えている」んじゃない。

AI は考えないし感じない。
少なくとも現代の AI 技術は数学の道具 論理・確率・統計 を駆使した入力と出力の組み合わせである。
それは「機械」の域を出ていない。
上の創作（笑）や[結城浩さんが Gist で公開](https://gist.github.com/hyuki/65ebb23855d31731ee2342e0920bcf9f "結城浩とChatGPTの対話 · GitHub")されている事例が示すのは，機械で以って人間側が受容できるレベルの対話もどきが成立している点だろう。

ならば，機械が嘘を吐く可能性がないと言えようか（反語）。

そう考えると [@hyuki@social.hyuki.net](https://social.hyuki.net/@hyuki "結城浩@social.hyuki.net (@hyuki@social.hyuki.net) - 結城浩のマストドン") さんの

{{< fig-quote class="nobox center" >}}
<iframe src="https://social.hyuki.net/@hyuki/109457705957545683/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe>
{{< /fig-quote >}}

も色んな意味を含んでいることが分かる。
相手の反応を引き出すための意図的なフェイクをビジネスロジックとして組み込むことは可能である，ということだ。
何故なら嘘に対する反応も AI にとっては有用な「入力」だから。

噂を誘導するためにネットにデコイをばら撒くってのは SAC 版の「攻殻機動隊」で出たネタだったか。
ネットにデコイをばら撒くってのは既にロシアとかがやってるとかいう噂もあるし（噂だよ，念のため），それを今の AI 技術でやったらさぞかし面白いことになるだろう。
しかもそれを受ける側も AI ボットだったりする可能性もあるわけだ。
まさに [AI エコー・チェンバー]({{< relref "./debugging-with-ai.md#ec" >}})（笑）

[ChatGPT]: https://chat.openai.com/

## 参考図書

{{% review-paapi "B01JMEDX8A" %}} <!-- 攻殻機動隊 STAND ALONE COMPLEX (SAC) -->
