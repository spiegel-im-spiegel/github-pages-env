+++
title = "Kagi Assistant と Google AI モード"
date =  "2025-09-09T19:50:39+09:00"
description = "両者の違いについて AI に訊いてみた"
image = "/images/attention/kitten.jpg"
tags = [ "web", "search", "google", "kagi", "tools", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今年（2025年）2月から[社内テスト](https://gigazine.net/news/20250207-google-test-search-ai-mode/ "Googleが新しい検索「AIモード」のテストを開始、一体どんな検索機能になるのか？ - GIGAZINE")が始まり，5月に[米国で一般提供](https://news.mynavi.jp/article/20250521-3331682/ "Google、“AI検索”に本腰、「AI Mode」米国から一般提供開始 | マイナビニュース")が開始された Google AI モードだが，日本での提供も開始されたようだ。

- [グーグル「AIモード」検索、日本でスタート 「検索は知性になる」 - Impress Watch](https://www.watch.impress.co.jp/docs/news/2045506.html)

私の環境ではまだ使えないのだが

{{< fig-quote type="markdown" title="グーグル「AIモード」検索、日本でスタート 「検索は知性になる」" link="https://www.watch.impress.co.jp/docs/news/2045506.html" >}}
AIモードの実現には、「クエリ ファンアウト」と呼ばれる技術を採用する。この技術では、質問をサブトピックに分解し、利用者に代わり、サブクエリに対して検索を実行。これにより、従来のGoogle 検索よりはるかに深くウェブを探索できるようになり、個別の質問に最も適した、関連性の高いコンテンツを見つけられるようにする。
{{< /fig-quote >}}

らしい。

一方で，我らが [Kagi Search] も [Kagi Assistant][Assistant] を以前から Ultimate プランのユーザに提供していたが，今年の4月に全ユーザ[^ku1] に開放された。

[^ku1]: [Kagi Search] を利用するためにはユーザ登録が必要。無料プランもあるが制限あり。検索機能を制限無しで使うには Professional 以上のプランに加入する必要がある。また [Assistant] で高機能モデルを使うには Ultimate プランに加入する必要がある。

- [Kagi Assistant が全ユーザに解放]({{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}})

両者ともプロンプトから複数の検索クエリを生成して検索し，その結果を基に回答を生成する，という仕組みのようだ。

というわけで [Kagi Assistant][Assistant] と Google AI モードとの違いを [Kagi Assistant][Assistant] に訊いてみた[^pub1]。

[^pub1]: [Kagi Assistant][Assistant] はスレッドごとに公開タグを付けることで結果を公開できるようになった。助かる。

- [Kagi Assistant vs. Google AI Mode - Kagi Assistant](https://kagi.com/assistant/7f40e789-3ce7-4d76-af80-245bbb06d900)

生成 AI が出力した結果だから鵜呑みにしないようにね。
ちなみに [Kagi Assistant][Assistant] では回答の末尾に

{{< fig-img src="./kagi-assistant.png" title="Kagi Assistant vs. Google AI Mode - Kagi Assistant" link="https://kagi.com/assistant/7f40e789-3ce7-4d76-af80-245bbb06d900" width="667" >}}

という感じに参考ページを示してくれるが，関連性の高いページほど行末の%の値が高くなるようだ。
これを参考に Web ページを確認するって感じだろうか。
行末の%表示は今月（2025-09）から追加されたのだが

{{< fig-quote type="markdown" title="September 4th, 2025 - Kagi Summarize goes mobile, Kagi Assistant adds source attribution and study mode" link="https://kagi.com/changelog#8164" lang="en" >}}
More importantly, this technology paves the way down the road for Kagi to share profits with publishers participating in our AI answers. This would happen automatically for all websites, with no deals, no contracts needed.
{{< /fig-quote >}}

などと述べている。
Google AI モードについては

{{< fig-quote type="markdown" title="グーグル「AIモード」検索、日本でスタート 「検索は知性になる」" link="https://www.watch.impress.co.jp/docs/news/2045506.html" >}}
ビジネス展開については、現在回答の下に広告をつけるテストを行なっており、今後も検討を進めていくとした。
{{< /fig-quote >}}

などと言ってるそうで，ここでも広告がついてまわるようだ。
まぁ，収益構造が違うからね。

両者を比較すると [Kagi Assistant][Assistant] があくまで[検索サービス][Kagi Search]とユーザとの仲立ちとして機能し，パーソナライズに関しても検索結果や AI 出力のカスタマイズおよびプライバシーへの配慮に重きを置いているのに対し， Google の AI サービスはおそらく Google 全サービスを統合してユーザに提供することを目論んでいるように見える。
AI モードは機能の一部に過ぎず，そこだけ見て判断できないってことかな。
まぁ Google が気にしてるのは Microsoft や Apple のような OS やクラウド市場で影響力のある企業だろうし，当然だよな。

私の環境で Google AI モードが使えるようになったら試してみたいところではあるが， Google 各サービスへの依存度を減らしたい私としては追加料金を払ってまで積極的に使わないかなぁ。

## ブックマーク

- [Web のコストは誰が支払うのか]({{< ref "/remark/2025/03/who-pays-for-web-costs.md" >}})

[Kagi Search]: https://kagi.com/ "Kagi Search"
[Assistant]: https://kagi.com/assistant "The Assistant"
[Translate]: https://translate.kagi.com/ "Kagi Translate"
<!-- eof -->
