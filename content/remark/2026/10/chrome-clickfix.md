+++
title = "Chrome の新機能と ClickFix"
date =  "2026-10-22T16:10:56+09:00"
description = "まぁ，私はプライバシーに関しては Google を全く信用してないので，そもそも Chrome ブラウザを入れてないんだけど。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "phishing", "web" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今回は Bluesky の TL で見かけた以下の記事を起点にしてみる。

- [安全？ それともプライバシーの侵害？ Chromeに新たに追加されたセーフブラウジング機能の仕様が波紋【やじうまWatch】](https://internet.watch.impress.co.jp/docs/yajiuma/2149836.html)

Chrome ブラウザの新機能についての記事なのだが，話が噛み合っていないように見える。
そこで，以下の元記事を少し詳しく見てみることにした。

- [Chrome will now show a warning before you visit ‘suspicious’ sites](https://www.theverge.com/tech/1014446/google-chrome-suspicious-website-warning)
- [Chrome can now share your clipboard data with Google servers](https://www.androidauthority.com/chrome-enhanced-protection-new-features-3725460/)

## The Verge 記事の評価

最初の The Verge の記事では Web サイトの危険度を「不審（suspicious）」と「危険（dangerous）」の二段階で評価し，視覚的にもわかりやすくしたというもの。
以下は [Kagi Assistant] による要約：

{{< div-ai type="markdown" >}}
- **不審なサイト**：フィッシングやマルウェア，詐欺サイトに似た特徴があるものの，誤判定の可能性もあります。警告は画面を暗くしたポップアップで表示され，誤判定の場合は「安全としてマーク」できます。
- **危険なサイト**：被害のリスクが高いと判断されたサイトです。警告は画面の大部分を占め，ページは赤い背景に置き換わります。記事は，こうしたサイトを避けるべきだと説明しています。
{{< /div-ai >}}

The Verge の記事はこれを好意的に評価している。

{{< fig-quote type="markdown" title="Chrome will now show a warning before you visit ‘suspicious’ sites" link="https://www.theverge.com/tech/1014446/google-chrome-suspicious-website-warning" lang="en" >}}
It’s easy to distinguish between these warnings at a glance. The notifications for suspicious websites will appear as a pop-up window on both desktop and mobile, with the webpage visible but darkened underneath, and an option to “Mark as safe” if the site was flagged mistakenly. The already familiar notification for dangerous sites is much harder to ignore: the warning message takes up almost the entire screen and the website will be replaced with a solid red background.
{{< /fig-quote >}}

「不審」なサイトの誤判定をユーザが「安全」としてマークできる仕組みはいいよね。
誤判定の頻度にもよるけれど。
誤判定が多ければ「不審」判定そのものを信用しなくなるし，それでユーザに安全か否かをマークさせ Google と共有するなら，立派な搾取構造である。

## Android Authority 記事の評価

次の Android Authority の記事では，先ほどの Web 危険度の評価と通知機能に加えて Gemini による詐欺の検出と ClickFix 対策についても言及している。

{{< fig-img-quote src="./server-side-gemini-scam-protection-in-chrome-840w-472h.jpg.webp" title="Chrome can now share your clipboard data with Google servers" link="https://www.androidauthority.com/chrome-enhanced-protection-new-features-3725460/" lang="en" width="840" >}}

「ClickFix って何？」と [Kagi Assistant] に訊いたのだが，以下がその回答。

{{< div-ai type="markdown" >}}
**ClickFix（クリックフィックス）** は，偽の CAPTCHA や「エラーを修正してください」といった画面を使い，利用者自身に悪意のあるコマンドを実行させる詐欺手法です。マルウェアを勝手に送り込むというより，利用者をだまして正規の操作機能を悪用させる「ソーシャルエンジニアリング」の一種です。

**典型的な流れ**

1. メールのリンク，広告，改ざんされたウェブサイトなどから，偽の確認画面やエラー画面に誘導されます。
2. 「人間であることを確認する」「問題を修復する」などと表示し，手順に従うよう促します。
3. ページがコマンドをクリップボードにコピーしたうえで，利用者にそれを Windows の「ファイル名を指定して実行」，PowerShell，Terminal などへ貼り付けて実行させます。
4. 実行すると，マルウェアがダウンロードされ，情報窃取や遠隔操作などにつながることがあります。

見た目には「確認」や「修復」の操作でも，実際にはページ上の説明とクリップボードにコピーされる内容が違う場合があります。つまり，**指示されたコマンドを自分で確認せずに貼り付けて実行させること**が，攻撃の肝です。

**どう見分けるか**

CAPTCHA やウェブサイトの「確認」手順の途中で，**コマンドをコピーして，パソコンの実行画面やターミナルに貼り付けるよう求められたら，実行しないでください。** 画面が有名企業のロゴを使っていたり，もっともらしいエラーを表示したりしていても，安全の根拠にはなりません。攻撃者はそうした表示で不審に思われにくくしようとします。

**参考文献**

- [Think before you Click(Fix): Analyzing the ClickFix social engineering technique | Microsoft Security Blog](https://www.microsoft.com/en-us/security/blog/2025/08/21/think-before-you-clickfix-analyzing-the-clickfix-social-engineering-technique/)
{{< /div-ai >}}

ひゃあ！ そんな手口があるのか。
油断大敵だねぇ。
気をつけましょう，お互いに。

で，この ClickFix を防ぐために Chrome というかローカルの Gemini はユーザマシンのクリップボードを監視して，情報をセーフブラウジング・サーバに送って評価してるらしい。

{{< fig-quote type="markdown" title="Chrome can now share your clipboard data with Google servers" link="https://www.androidauthority.com/chrome-enhanced-protection-new-features-3725460/" lang="en" >}}
Chrome will now detect if a user has been tricked into copying a suspicious command, and will send the clipboard data to Safe Browsing servers for analysis. If the server finds that the command is dangerous, Chrome will display a warning to prevent the user from executing it on their computer. Sharing clipboard data with a server also raises privacy questions, but Google states that Chrome will strip any personally identifiable information from the clipboard before sharing the data with Safe Browsing servers.
{{< /fig-quote >}}

上の引用のように，記事ではクリップボードの内容を送り込むことについてはプライバシー上の懸念があるとしながらも「クリップボードのデータから個人を特定できる情報は抜くから大丈夫じゃね？」と楽観的に考えているように見える。

要するに両者とも（条件付きではあるが）今回の機能追加については概ね好意的に評価していると思われる。
なお，これらの機能を含む Enhanced Protection は既定では無効になっていて，オプトインで有効化する必要がある。

まぁ，私はプライバシーに関しては Google を全く信用してないので，そもそも Chrome ブラウザを入れてないんだけど。

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
