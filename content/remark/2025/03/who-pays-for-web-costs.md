+++
title = "Web のコストは誰が支払うのか"
date =  "2025-03-11T14:39:34+09:00"
description = "改めて思うのだが Web 2.0 最大の罪はそのコストを安易に他者へ「転嫁」してしまったことだろう。"
image = "/images/attention/kitten.jpg"
tags = [ "google", "firefox", "search"  , "market", "media", "web", "security", "privacy", "risk" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今月（2025-03）に入り Chrome 系のブラウザ拡張機能について Manifest-V2 廃止に基づく粛清がいよいよ始まるようだ？

- [Google’s Chrome extension cull hits more uBlock Origin users | The Verge](https://www.theverge.com/news/622953/google-chrome-extensions-ublock-origin-disabled-manifest-v3)
- [Google Chrome Disables uBlock Origin and Other Extensions | Extremetech](https://www.extremetech.com/computing/google-chrome-disables-ublock-origin-and-other-extensions)

実際には Manifest-V2 の廃止までにはまだ猶予があり，どうやら Google 側が姑息な排除をしようとしているんじゃないかという話もある。

- [Googleが「uBlock Originのサポートは終了しました」とウソをついているとネットが騒然、広告ブロッカーを使い続ける方法はコレ - GIGAZINE](https://gigazine.net/news/20250307-ublock-origin-is-gone/)

本当のところは分からないが（Chrome 使ってないので），今回の V2 から V3 への移行で（利用者から見て）最もインパクトがある拡張機能のひとつが [uBlock Origin] だろう。
[uBlock Origin] は Manifest-V2 の `webRequest` API 仕様に大きく依存していて，新しい Manifest-V3 の API では機能が制限されるらしい。

{{< fig-quote type="markdown" title="uBlock Origin - Free, open-source ad content blocker." link="https://ublockorigin.com/" lang="en" >}}
uBlock Origin relies heavily on the `webRequest` API to block unwanted content before it loads. Under MV3, the `webRequest` API is limited, and extensions are encouraged to use the new `declarativeNetRequest` API instead. This new API allows for predefined rules but lacks some of the dynamic capabilities that uBlock Origin utilizes for advanced content blocking.
{{< /fig-quote >}}

Manifest-V3 でこの制限をうまく回避するる方法はないようで，以下の代替案が提案されている。

{{< fig-quote type="markdown" title="uBlock Origin - Free, open-source ad content blocker." link="https://ublockorigin.com/" lang="en" >}}
1. Continue Using uBlock Origin on Firefox
1. Use uBlock Origin Lite
1. Switch to Browsers Committed to MV2 Support
1. Explore Other Content Blocking Methods
{{< /fig-quote >}}

[uBlock Origin Lite](https://chromewebstore.google.com/detail/ublock-origin-lite/ddkjiahejlhfcafbddmgiahcphecmpfh) であれば Manifest-V3 ベースらしいので Chrome ユーザは取り敢えずこちらに移行する手はある。

一方 Firefox は Manifest-V2 と Manifest-V3 の両方をサポートすると明言しているので，いっそ Firefox に乗り換える手もある。

- [Mozilla’s approach to Manifest V3: What’s different and why it matters for extension users | The Mozilla Blog](https://blog.mozilla.org/en/products/firefox/firefox-manifest-v3-adblockers/)

ブラウザを乗り換えていいのであれば，もうひとつ [Orion] ブラウザに乗り換える手もある。
有料検索サービスの Kagi が提供している [Orion] ブラウザは広告ブロッキング機能を既定で備えていて，他にもプライバシー保護の観点から他のブラウザより有利であると主張している。

{{< fig-quote class="nobox" type="markdown" title="Orion Browser by Kagi" link="https://kagi.com/orion/" lang="en">}}
| Privacy comparison | Orion | Safari | Firefox | Brave | Chrome |
| --- | :---: | :---: | :---: | :---: | :---: |
| Zero telemetry by default | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} |
| Blocking 1st party ads by default | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} |
| Blocking 1st party trackers by default | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} |
| Blocking 3rd party ads by default | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} | {{< emoji "バツ" >}} | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} |
| Blocking 3rd party trackers by default | {{< emoji "チェック" >}} | {{< emoji "チェック" >}} | {{< emoji "チェック" >}} | {{< emoji "チェック" >}} | {{< emoji "バツ" >}} |
{{< /fig-quote >}}

残念ながら [Orion] ブラウザは WebKit ベースのため，今のところ macOS および iOS/iPadOS しか対応していない。
しかし，昨今の情勢を見て Linux 版の開発に着手したようだ。

{{< fig-quote type="markdown" title="March 6th, 2025 - Orion Embarks on Linux Journey & Kagi Doggo Art Celebration" link="https://kagi.com/changelog#6479" lang="en" >}}
We're thrilled to announce that development of Orion Browser for Linux has officially started! Our team is working hard to bring the same speed, privacy, and innovation that Mac users love to the Linux platform.

This is an ambitious project that we expect will take approximately one year to complete. Our target is to achieve feature parity with the current macOS version by March 2026.
{{< /fig-quote >}}

手元にある [MacBook Air]({{< ref "/remark/2024/05/get-a-used-pc-from-workplace.md" >}} "勤務先からの払い下げ PC") には既に導入しているが，今のところ不都合はない。
まぁ macOS のアプリケーションはあまり使わず [Linux 環境]({{< ref "/remark/2025/01/kubuntu-on-macbook-air-m1.md" >}} "MacBook Air M1 に Kubuntu を入れる")での作業が殆どなのだが。
私は Linux 版を待ってます。

[Orion] ブラウザには [Orion+](https://kagi.com/orion/orionplus.html "Support Orion") ってのがあるそうで，一括で150USD払うか月5USD（年50USD）のサブスクリプションに加入することで RC 版の利用が可能になる他，フィードバックにも優先的に対応してもらえるらしい。
Kagi は有料検索サービスとこの [Orion+] の収益で開発・運用を行っているようだ。

その [Kagi の検索サービス](https://kagi.com/ "Kagi Search")だが，現在は Professional プランを月ごとの支払いで運用している。
もうガッツリ使ってるよ。
検索周りの UX が[優秀]({{< ref "/remark/2024/06/kagi-search.md" >}} "Kagi Search を試してみる 〜 検索サービスも有料の時代？")なのは言うまでもないが AI 絡みの機能（Translate, FastGPT[^g1], Universal Summarizer）も便利に使っている。
なので，年単位の契約に切り替えてもいいかなぁ，と考え中ではある。
もしくは Assistant が使える Ultimate プランにするか。
いやでも私の場合 AI アシスタントは GitHub Copilot で足りてるからなぁ。

[^g1]: FastGPT は本当に簡易的な機能のみの提供で，単一の応答しか出来ないし LLM モデルの選択もできない。連続的な「会話」や LLM モデルの切り替え機能が欲しいのであれば Ultimate プラン（月25USD）に加入した上で Assistant 機能を利用する必要がある。

Kagi の中の人が Bluesky で[書いていた](https://bsky.app/profile/kagi.com/post/3lk2gdsm4es2v "Kagi HQ: \"There are only two business models on the web: either you pay for your browser, or someone else does. ...\" — Bluesky")が，ネット上で享受しているサービスなりプラットフォームなりのコストは誰が支払っているのか，って話だ。

改めて思うのだが Web 2.0 最大の罪はそのコストを安易に（本来の利用者ではない）他者へ「転嫁」してしまったことだろう。

そして今になって私達は「[{{% ruby "enshittification" %}}メタクソ化{{% /ruby %}}](https://en.wikipedia.org/wiki/Enshittification "Enshittification - Wikipedia")」という形で（文字通りの）ツケを支払っているわけだ。
広告モデルで実際にコストを支払ってるのは広告主なんだから，サービス側プラットフォーム側が広告主の利益を最大化するよう調整するのは当然と言える。

- [かくしてGoogleはスパマーに敗北した » p2ptk[.]org](https://p2ptk.org/monopoly/4515)
- [Google検索を殺した男――Googleはいつ、どこでメタクソ化に舵を切ったのか » p2ptk[.]org](https://p2ptk.org/monopoly/4541)

だからといって「それ」を許容できるかと問われたら否と答えるけど。

特に Web 上の広告はもはや好悪の問題ではなくセキュリティやプライバシーのリスクの問題になっている。

- [広告ブロッカーは「嫌ならどうする？」の表明である » p2ptk[.]org](https://p2ptk.org/monopoly/2668)
- [米国 FBI は広告ブロッカーを推奨している？]({{< ref "/remark/2022/12/ad-blocker.md" >}})
- [「Google広告からの誘導が6割」との分析結果。より巧妙化し、高齢者を狙う「サポート詐欺」に注意！【被害事例に学ぶ、高齢者のためのデジタルリテラシー】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/column/dlis/1592999.html)

たとえば [Publickey](https://www.publickey1.jp/ "Publickey － Enterprise IT × Cloud Computing × Web Technology / Blog") のように，[ポリシーを持って広告を管理](https://www.publickey1.jp/blog/25/2025publickey.html "年頭のご挨拶：2025年のPublickeyも、読者が安心して記事を読めるように適切な広告だけを掲載します － Publickey")しているサイトは少なく，大抵は Google 等の広告サービスが提供しているものを垂れ流してるだけの脆弱なサイトに見える。
まぁ，一番ダークなのは間違いなく Google 検索サービスのページだけど（笑）

[uBlock Origin] などのツールは単に広告をブロックするツールというだけではなく，そのサイトの広告を許容できるか否かの決定権を，見ているユーザ側に取り返すという大切な手段である。
その上で私達は（広告モデルを含む）他者にコストを転嫁する仕組みから距離を置く必要があると思う。
タダならば何でもいいという世の中じゃなくなったってことかねぇ。

そのビジネスモデルで本当に利用者は幸せになれるのか。
よくよく考えてリソース（お金とは限らない）を投入しないといけないし，私のようなビンボー人は幾つかサービスを諦めることも視野に入れないといけないかもなぁ。

## ブックマーク

- [Kagi is Bringing the Orion Web Browser to Linux - OMG! Ubuntu](https://www.omgubuntu.co.uk/2025/03/kag-orion-web-browser-coming-to-linux)
- [大いなる力には何の責任も伴わなかった » p2ptk[.]org](https://p2ptk.org/monopoly/5369)

[uBlock Origin]: https://ublockorigin.com/ "uBlock Origin - Free, open-source ad content blocker."
[Orion]: https://kagi.com/orion/ "Orion Browser by Kagi"
[Orion+]: https://kagi.com/orion/orionplus.html "Support Orion"

## 参考図書

{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->
{{% review-paapi "B0CK19L1HC" %}} <!-- ハッキング思考 Kindle 版 -->
{{% review-paapi "430924744X" %}} <!-- スパム -->
