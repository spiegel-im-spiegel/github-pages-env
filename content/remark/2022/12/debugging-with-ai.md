+++
title = "デバッグする AI"
date =  "2022-12-02T21:12:03+09:00"
description = "今日も楽しく遊びました。"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

朝から私の Twitter TL が騒がしいなぁ，と思ったら [ChatGPT] なるものが流行ってるらしい。

- [ChatGPT: Optimizing Language Models for Dialogue](https://openai.com/blog/chatgpt/)
- [ChatGPT使い方総まとめ - Qiita](https://qiita.com/sakasegawa/items/82069c97a1ee011c2d1e)

たとえば「goroutine を使って並行処理で複数の URL から CSV ファイルをダウンロードするコードを書いてください。」と入力するとちゃんと[動くっぽいコードを返してくれる](https://twitter.com/mattn_jp/status/1598474642209275904)みたいだし，コードを示して「バグを探してくれ」と書いたら[バグの内容と修正方法までアドバイス](https://twitter.com/yuroyoro/status/1598538264126050304)してきたらしい，流暢な日本語で。

まぁ，今は[コミットメッセージを AI に書かせる](https://note.com/sakasegawa/n/n9f63e82ef391)時代らしいし，これって前世紀の人が夢見た，まさに「エキスパートシステム」だねぇ。
凄い凄い。

一方，私はといえば...

「語尾に「にゃん」を付けて可愛らしい言い方で Mastodon について解説して」って頼んだら。

{{< fig-img-quote src="./chatgpt-1.png" title="ChatGPT" link="https://chat.openai.com/" width="805" lang="en" >}}

とか返してきた。
思わず「[そうだけどそうじゃねー](https://fedibird.com/@spiegel/109441682281643776)」ってツッコんじゃったよ（笑） まぁ，個人的に一番面白かったのは「「やせ蛙負けるな一茶これにあり」の俳句を解説してください。」に対して

{{< fig-img-quote src="./chatgpt-2.png" title="ChatGPT" link="https://chat.openai.com/" width="835" lang="en" >}}

と返したことだったけどね。
最後の「この俳句は、「人間の弱さ」をテーマにした俳句であると言えます」のところだけ合ってるような気がしないでもないのがにんともかんとも（笑）

[前にも書いた]({{< ref "/remark/2022/10/cultural-commons.md" >}} "AI は「創作者様」を引きずり下ろすか — 『人権と文化コモンズ』を流し読む")が，現代の AI 技術が齎すものは **翻案の大量生産** だろう。
20世紀が夢見た「エキスパートシステム」も，今となってはそのバリエーションのひとつに過ぎなくなっている。
「翻案機」としての AI をどう上手く使うかがこれからの「人材」に求められているのかもしれない。

なんてな。
今日も楽しく遊びました。

## ブックマーク

- [結城浩とChatGPTの対話 · GitHub](https://gist.github.com/hyuki/65ebb23855d31731ee2342e0920bcf9f)
- [フィリップ・グラスが人工知能と芸術について語る - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20210902/philip-glass-on-ai)

[ChatGPT]: https://chat.openai.com/

## 参考図書

{{% review-paapi "B0791XCYQG" %}} <!-- AI vs. 教科書が読めない子どもたち -->
