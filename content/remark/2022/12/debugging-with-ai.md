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

## 【2022-12-04 追記】 AI エコー・チェンバー {#ec}

なるほどな，と思う tweet を見かけた。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">文章作成AIによる解説ブログ、近いうちに乱立してめちゃくちゃノイズになると思う。<br><br>ボリュームがそこそこある検索ワードを引っこ抜いて、検索ワードごとに文章を自動生成するだけ。<br>技術的には簡単だし、低コスト・低リスクで収入が期待できてしまうという。</p>&mdash; catnose (@catnose99) <a href="https://twitter.com/catnose99/status/1599014159320252421?ref_src=twsrc%5Etfw">December 3, 2022</a></blockquote>
{{< /fig-gen >}}

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">こういうサイトが乱立して、Googleの機械学習でも大量のゴミデータが食わされるようになった時がヤバそう</p>&mdash; catnose (@catnose99) <a href="https://twitter.com/catnose99/status/1599014161908105216?ref_src=twsrc%5Etfw">December 3, 2022</a></blockquote>
{{< /fig-gen >}}

「エコー・チェンバー効果」と呼ばれるものがある。
閉じたコミュニティなどで文字通り「反響室（echo chamber）」のように特定の思想や意見が増幅・固定されてしまうことを指すらしい。
まぁ，コミュニティ内でそういった「増幅」が見られるのは普通だし，[だからどうしたって感じ]({{< ref "/remark/2021/09/echo-chamber.md" >}} "「エコーチェンバーの中にいる可能性が高い」")ではある。
世捨て人じゃあるまいし，人が社会の中で生きていく限り，多かれ少なかれ，こうした影響からは免れない。

しかし，この「エコー・チェンバー」が AI と組み合わさると面白いことになりそうだ。

昔で言うところの（今でも言うかもしれないが）「[1円ライター]({{< ref "/remark/2016/12/is-curation-media-dead.md" >}} "「キュレーション」じゃなく，はっきり「注目の搾取」と言えよ！")」がネタ出しに [ChatGPT] みたいなツールを使うことはありそうな話だし，ツールの出力結果が Web ページとして「洗浄」されてしまえば判別は難しいだろう。

人間から見れば，今でさえ S/N 比が底辺状態の Web に「翻案の大量生産」が加わったとしてもどうってことはないだろうが，意図せず AI の出力結果が AI にフィードバックされる本当の「エコー・チェンバー効果」が起きたらどうなるのか。
なかなか興味深い。

## ブックマーク

- [OpenAI API](https://beta.openai.com/)
- [質問に答えてくれる言語モデル「ChatGPT」--プレビューは無料公開 - ZDNet Japan](https://japan.zdnet.com/article/35196862/)
- [結城浩とChatGPTの対話 · GitHub](https://gist.github.com/hyuki/65ebb23855d31731ee2342e0920bcf9f)
- [フィリップ・グラスが人工知能と芸術について語る - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20210902/philip-glass-on-ai)
- [ChatGPTを使って画像生成AIのプロンプトを生成する · GitHub](https://gist.github.com/hyuki/0943a22a166923e1ee0cda56c3ca1012)
- [対話AI「ChatGPT」が大学生レベルの試験の自由記述問題に合格してしまう - GIGAZINE](https://gigazine.net/news/20221205-chatgpt-passes-ap-computer-science-a/)
- [会話AI「ChatGPT」の回答の投稿がコーディングQ＆AサイトのStack Overflowで一時的に禁止される - GIGAZINE](https://gigazine.net/news/20221206-chatgpt-banned-temporary/) : 「「ChatGPT」の回答」の判断基準を示してくれる記事が何処にも見当たらないのだが，これって体のいい検閲じゃねーの？
- [GitHub - humanloop/awesome-chatgpt: Curated list of awesome tools, demos, docs for ChatGPT and GPT-3](https://github.com/humanloop/awesome-chatgpt)
- [AIの応用における「検索」と「生成」の類似性 - 結城浩の連ツイ](https://rentwi.hyuki.net/?1601462336899796993s)
- [2022年と、AI戦争の歴史](https://youkoseki.com/text/2022_ai) : 名作！
- [AI assistants help developers produce code that's insecure • The Register](https://www.theregister.com/2022/12/21/ai_assistants_bad_code/)

[ChatGPT]: https://chat.openai.com/

## 参考図書

{{% review-paapi "B0791XCYQG" %}} <!-- AI vs. 教科書が読めない子どもたち -->
