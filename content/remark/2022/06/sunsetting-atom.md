+++
title = "Atom の落日"
date =  "2022-06-11T09:19:04+09:00"
description = "ユーザ間のエコシステムが出来上がっている製品であっても廃れたらあっという間"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "editor", "atom", "vscode", "golang" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

GitHub が Microsoft に買収されて以来，その内そうなるだろうとは思っていたが，ついにこの日が来てしまったか。

{{< fig-quote type="markdown" title="Sunsetting Atom | The GitHub Blog" link="https://github.blog/2022-06-08-sunsetting-atom/" lang="en" >}}
Today, we’re announcing that we are sunsetting Atom and will archive all projects under the organization on December 15, 2022.
{{< /fig-quote >}}

Atom には[思い入れ]({{< rlnk "/tags/atom/">}})がある。
思い返せば2015年は色々と心境の変化があった年で，たとえば Windows 7 のサポート終了をにらんで Windows 依存からの脱却を模索し始めた年で [Go] で遊び始めた年でもある。
「Windows 依存からの脱却」の最大の障害はテキストエディタの秀丸に依存しきっていたことで，マルチプラットフォームで手に馴染むエディタを探すことが[最優先事項](https://dic.pixiv.net/a/%E9%A2%A8%E8%A6%8B%E3%81%BF%E3%81%9A%E3%81%BB "風見みずほ (かざみみずほ)とは【ピクシブ百科事典】")だったのだ。

というわけで，私の中で Atom と [Go] はセットになっていた。
Atom の拡張機能である go-plus の出来が（当時としては）よかったのも大きい。

そういえば Twitter で「VSCode に「中華を初めて統一した始皇帝」みたいな印象ができつつある」みたいな tweet を見かけたが，[エディタ界の始皇帝は vim だろう](https://twitter.com/spiegel_2007/status/1535014771522560001)とか思ってみたり。
このブログでも何度か書いているが，私は vi にトラウマがあって， vi/vim を起動するとペーペーの新人の頃に工場の片隅でガチの VT 端末を前に泣きながらデバッグしていたあの頃がフラッシュバックしてしまうのだ。
なので最初から vim という選択肢はなかった。
まぁ Ubuntu 環境にいると結局は vim も使わざるを得ないのだが（笑）

{{< fig-youtube id="P7LNU9HYr7M" title="プログラミルクボーイ「Vim」 - YouTube" >}}

話を戻そう。

私の中で風向きが変わったのは “[Language Server Protocol (LSP)](https://microsoft.github.io/language-server-protocol/ "Official page for Language Server Protocol")” が登場したあたり。
LSP はホンマに画期的なアイデアで，当然 [Go] 用の lunguage server も登場するんだけど[^lsp1] go-plus はこれを取り込むことができなかった。

[^lsp1]: LSP が発表された頃 [Go] 用の lunguage server としていくつかの実装があったが，現在はほぼ [gopls](https://github.com/golang/tools/tree/master/gopls) 一択である。

加えて，私が IT 業界に再就職して支給された Windows 10 機に Atom を入れたら堪えられないほど遅いのにビックリし，観念して VS Code に乗り換えたのだった。

- [パソコンに Visual Studio Code を導入する（再チャレンジ）]({{< ref "/remark/2021/02/installing-vscode-again.md" >}})

Microsoft Windows を嫌って Ubuntu や Atom にしたというのに，結局は Microsoft に屈してしまったわけだ（「[くっころ](https://dic.pixiv.net/a/%E3%81%8F%E3%81%A3%E3%80%81%E6%AE%BA%E3%81%9B%21 "くっ、殺せ! (くっころ)とは【ピクシブ百科事典】")」とか言わないよ）。
今やすっかり手に馴染んでしまったけどね。

Atom が登場したのは2014年だそうだが，当時は SublimeText のカウンタという位置付けだったと思う。
8年というのはソフトウェア製品としては息が長いほうだと思うけど，テキストエディタは下手すると10年20年と使うものなので簡単に消えられては困るわけですよ。
それでも時代の流れには逆らえない。
ユーザ間のエコシステムが出来上がっている製品であっても廃れたらあっという間ということか。
まさに「落日」だな。

せめて VS Code のエコシステムが永く続くことを祈ろう。

## ブックマーク

- [Why GitHub Is Killing Atom Text Editor](https://www.makeuseof.com/atom-text-editor-why-github-is-killing/)
- [Introducing Zed – A collaborative code editor written in Rust](https://zed.dev/)

[Go]: https://go.dev/ "The Go Programming Language"
