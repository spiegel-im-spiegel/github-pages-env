+++
title = "オープンソースはコードの多様性を担保しない"
date = "2018-12-10T22:28:16+09:00"
description = "オープンソースは自由かもしれないが多様とは限らない。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "engineering", "firefox", "chromium", "webkit" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

まぁ Microsoft は営利企業として賢明な判断をしたと思うよ。

- [Microsoft Edge、Chromiumベースに――旧Windowsでも作動、macOS版も登場へ  |  TechCrunch Japan](https://jp.techcrunch.com/2018/12/07/2018-12-06-microsoft-edge-goes-chromium-and-macos/)
- [EdgeブラウザがついにChromiumを採用へ　Mozillaは「独占は危険」と警告 - ITmedia PC USER](http://www.itmedia.co.jp/pcuser/articles/1812/09/news016.html)
- [マイクロソフトのEdgeがChromiumベースになることによりウェブから多様性が失われることを懸念するのでモジラ財団に寄付した - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20181210/mozilla)
- [どうなる、「Microsoft Edge」？ ～「Chromium」ベースになってよい点、悪い点 - やじうまの杜 - 窓の杜](https://forest.watch.impress.co.jp/docs/serial/yajiuma/1157671.html)

これでブラウザエンジンはいよいよ [WebKit] と [Blink]/[V8] と [Servo] に絞られたか。
確か W3C で標準と認められるには2種類以上の実装実績が必要だったはず。
親戚同士[^ch1] の [WebKit] と [Blink] で実装してしまえば，なし崩し的に標準になるわけだ（笑）
しかも携帯端末の世界では既に [WebKit]/[Blink] でほぼ独占状態。
更に JavaScript エンジンに至っては今や [V8] 一択だろうし[^e1]， [Mozilla] の言い分は今更すぎる話ではある。

[^ch1]: [Chromium] と [V8] はもともと [WebKit] のポーティングのひとつだったが，レンダリングエンジンや JavaScript エンジンの一部を fork して独立したプロジェクトになった。 [Blink] は [Chromium] プロジェクトで開発しているレンダリングエンジンという位置づけ。ちなみに Chrome for iOS は [Chromium] プロジェクトのひとつだが iOS では [WebKit] を使うことが要求されているため [Blink] を使わない実装になっている。
[^e1]: そういえば [ATOM](https://atom.io/) エディタなどで採用されている [Electron](https://electronjs.org/ "Electron | Build cross platform desktop apps with JavaScript, HTML, and CSS.") は [Chromium] ベースだったっけ。ブラウザのみならずデスクトップ・アプリでも [Chromium] による侵食は進行しているわけやね。他にも Google は Unicode で大きな発言権を持ってるし（Unicode に絵文字を追加することを最初に提案したのは Google）自前で開発したプログラミング言語もいくつかある。更に AI 技術についても大きなクラウド資源を持っていて，それを十全に発揮している。本当に，本当に今更な話で，これが[「情報力」が台頭した現代]({{< ref "/remark/2017/10/too-many-ghosts.md" >}} "今こそ「グリゴリの捕縛」を読め！ または遍在する草薙素子")における Google の強さ・凄みであり怖さと言える。市場の独占はそれに付随するものに過ぎない。

{{< fig-quote title="Microsoft Edge、Chromiumベースに――旧Windowsでも作動、macOS版も登場へ" link="https://jp.techcrunch.com/2018/12/07/2018-12-06-microsoft-edge-goes-chromium-and-macos/" >}}
<q>数日前にEdgeがリニューアルされるといいう情報が流れたとき、一部の専門家はChromiumプロジェクトが力を持ちすぎることになるという懸念を示した。<br>
この懸念には理由があることは認めるものの、MicrosoftはどのみちEdgeのシェアは低いのでChromium化がオープンソース・コミュニティーにドラスティックな影響を与えることはないという説得力のある反論をしている。MicrosoftがChromiumコミュニティーに参加してウェブの標準化を推進する側に回り、Chromiumにイノベーションを吹き込むことになればメリットは大きいだろう。</q>
{{< /fig-quote >}}


フリーな製品が市場にインパクトを与えたという点ではまさに Chrome を含む [Chromium] プロジェクトがそれである。
ならば当初 [Mozilla]/Firefox が目指していたプロプライエタリな IE へのカウンタとしての役目は [Chromium] による独占で達成されたわけであり， [Mozilla] が文句を言う筋合いではないんじゃないだろうか。

Linux より優れた OS をフルスクラッチで作ろうという人は少ないだろう。
ブラウザだって今更フルスクラッチで作り直そうなんて（しかもそのために[プログラミング言語](https://www.rust-lang.org/ "Rust programming language")から開発するとか）酔狂なことをするのは [Mozilla] くらいなものである。
私はそういう狂った人は好きだけど（笑）

まっとうなエンジニアなら（遊びや勉強用ならともかく）優れた製品を前にわざわざ「車輪の再発明」なんかしない。
それがオープンソースなら尚更。
昔に比べれば今はオープンソース・プロジェクトの fork に対して寛容になってるけど，それでも頻繁に起きることではない。
オープンソースは自由かもしれないが多様とは限らないということだよね。

ブラウザエンジンのシェア独占が問題だというのなら，たとえば iOS で [WebKit] 使用を要求していることとか Android の WebView によって実質的に [Chromium] 以外が締め出されている（選択肢に上らない）こととか[^ff1]，そういったことをもっと議論すべきなんじゃないの？

[^ff1]: プライバシーに配慮したとされる Firefox Focus の Android 版では独自のレンダリングエンジンとして [GeckoView を作った](https://support.mozilla.org/ja/kb/geckoview-firefox-focus "Firefox Focus の GeckoView | Firefox Focus ヘルプ")らしい。

そういう風に考えたら今回の Edge の話は，単にフロントエンド・エンジニアが楽できるということではなく，プロプライエタリ OS ベンダが自前のブラウザエンジンを捨てて「自由なソフトウェア」を選択したという意味で（Microsoft にしては）めっさ英断だと思うけどね。
最終的にどうなるかは知らないが。

## ブックマーク

- [iOS版Chromeがオープンソース化される - GIGAZINE](https://gigazine.net/news/20170201-ios-chrome-open-source/)
- [「Edgeで既存のChrome拡張機能をサポートする意向」とMicrosoftの担当者がRedditで - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1812/11/news072.html)

- [技術的負債としての Web ブラウザ]({{< ref "/remark/2018/06/web-browser-as-the-technical-debt.md" >}})

[Chromium]: https://www.chromium.org/ "The Chromium Projects"
[Blink]: https://www.chromium.org/blink "Blink - The Chromium Projects"
[V8]: https://v8.dev/ "V8 JavaScript engine"
[WebKit]: https://webkit.org/
[Mozilla]: https://www.mozilla.org/
[Servo]: https://servo.org/ "Servo, the parallel browser engine"

## 参考図書

{{% review-paapi "B009LFBL4Y" %}} <!-- グーグル　ネット覇者の真実 -->
