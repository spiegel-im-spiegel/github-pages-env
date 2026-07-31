+++
title = "Kagi Assistant が全ユーザに解放"
date =  "2025-04-19T18:32:55+09:00"
description = "この方法のいいところは検索結果と LLM による出力が明確に分離されている点だろう。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "kagi", "search", "tools", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

有料の検索サービス [Kagi Search] は検索以外にいくつかのサービスを提供しているが，そのうちのひとつである [Assistant] 機能が全ユーザに解放された。

- [Kagi Assistant is now available to all users! | Kagi Blog](https://blog.kagi.com/assistant-for-all)

2025年4月現在， [Kagi Search] の料金体系は以下のようになっているが

{{< fig-img src="./plans.png" title="Kagi Search Pricing and Plans" link="https://kagi.com/pricing" width="1565" >}}

以前は Ultimate プランでしか [Assistant] 機能が使えなかった。
これを全ユーザで使えるようにしたということらしい。

さっそく試してみた。
こんな感じ。

{{< fig-img src="./the-assistant.png" title="The Assistant" link="./the-assistant.png" width="1281" >}}

ただし Ultimate プランとそれ以外とで一応の差別化はされているみたいで Ultimate プラン以外では使えるモデルに制限があるようだ。

{{< fig-img src="./model-selection.png" title="モデルの選択" link="./model-selection.png" >}}

Ultimate プランなら {{% emoji "X" %}} の Grok も使えるらしい。
モデルの[ベンチマーク](https://help.kagi.com/kagi/ai/llm-benchmark.html "Kagi LLM Benchmarking Project | Kagi's Docs")もあるので参考になるだろう。
まぁ，今のところアップグレードは必要ないか。

[Kagi Assistant][Assistant] ではカスタムアシスタントを複数作成できる。

{{< fig-img src="./custom-assistant.png" title="カスタムアシスタントの作成" link="./custom-assistant.png" width="1281" >}}

これを使って [Kagi Search] のレンズ機能（検索対象の絞り込みをカスタマイズできる）や検索結果のパーソナライズ（検索結果の優先順位付けをユーザが設定できる）をアシスタントに反映させることが出来る。

Bang 機能は検索する際に特定のサービスを指定する機能で，例えば検索クエリの先頭に `!w` を付加すると Wikipedia を検索することができる。

{{< fig-img-quote src="./bang.Bm6fkjvX.gif" title="Bangs | Kagi's Docs" link="https://help.kagi.com/kagi/features/bangs.html" width="800" lang="en" >}}

これを [Assistant] と関連付けることが出来る。
たとえば先ほどの図のように Bang 名を `!foo` に設定して以下のように検索窓で問い合わせてみる。

{{< fig-img src="./query.png" title="Kagi Search with Custom Assistant" link="./query.png" width="1281" >}}

すると，以下のように [Assistant] のほうで応答してくれる。

{{< fig-img src="./response.png" title="Custom Assistant" link="./response.png" width="1281" >}}

この方法のいいところは検索結果と LLM による出力が明確に分離されている点だろう。
ぶっちゃけ Bing や Google Search は LLM が前面に出過ぎていて既に検索サービスじゃなくなってる。
まぁ，検索結果が使いものにならないくらい精度が悪いからだろうけど。

ローカルにあるファイルを読み込ませて問い合わせることも出来る。
たとえば以下の写真を読み込ませて

{{< fig-img src="./54461519855_3ee2153756_e.jpg" title="今日のお花見（カンザン） | Flickr" link="https://www.flickr.com/photos/spiegel/54461519855/" >}}

何の花か問い合わせてみると

{{< fig-img src="./response-2.png" title="花の名は。" link="./response-2.png" width="710" >}}

などと大雑把な答えが返ってきた。
いや，八重桜で間違ってないけどさ（笑） この辺はまだ Google Lens には勝てないか。

なお画像をジブリ風に変換するとかは出来ないらしい。
試してみたが以下のように諭された。

{{< fig-img src="./response-3.png" title="漫画風に変換して" link="./response-3.png" width="710" >}}

まじすんません {{% emoji "ペコン" %}}

もともと [Kagi Search] の検索機能は（Bing や Google Search に比べて）とても優秀なのだが，周辺サービスも充実してきた。
[翻訳][Translate]機能（これも LLM を使った翻訳）も徐々に良くなってる印象（まだ Google 翻訳との併用が必要だけど）。

私は AGI の到来など（近い将来では）全く信用してないが LLM を中心とした技術は便利であることに間違いはないので（懐と相談しつつ）無理のない範囲で利用していくとしよう。

## ブックマーク

- [なぜGoogleは“あなたの不満”を無視できるのか » p2ptk[.]org](https://p2ptk.org/monopoly/4535)
- [検索ツール「Kagi」に課金して世界が変わった10のこと | ライフハッカー・ジャパン](https://www.lifehacker.jp/article/2505-the-best-hidden-features-in-kagi-the-paid-alternative-to-google-search/)

- [Kagi Search を試してみる 〜 検索サービスも有料の時代？]({{< ref "/remark/2024/06/kagi-search.md" >}})
- [Web のコストは誰が支払うのか]({{< ref "/remark/2025/03/who-pays-for-web-costs.md" >}})

{{< fig-youtube id="UoPW-_IH6pE" title="22年連続日本一の庭園を訪れ溢れる狂気に歓喜するらでん [ホロライブ/ReGLOSS/儒烏風亭らでん] - YouTube" >}}

[Kagi Search]: https://kagi.com/ "Kagi Search"
[Assistant]: https://kagi.com/assistant "The Assistant"
[Translate]: https://translate.kagi.com/ "Kagi Translate"
