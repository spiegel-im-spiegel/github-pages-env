+++
date = "2015-11-21T07:48:29+09:00"
update = "2015-11-25T21:06:23+09:00"
description = "The TPP must be rejected. / 広島大学天文学研究会プラネタリウム & 天体観望会 / NTPsec ベータ版が公開 / Visual Studio Code がベータ版に到達"
draft = false
tags = ["code", "intellectual-property", "copyright", "tpp", "creative-commons", "astronomy", "planetarium", "security", "ntp", "engineering", "editor", "vscode", "dot-net"]
title = "週末スペシャル： The TPP must be rejected."

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [The TPP must be rejected.]({{< relref "#tppip" >}})
1. [広島大学天文学研究会プラネタリウム & 天体観望会]({{< relref "#huaa" >}})
1. [NTPsec ベータ版が公開]({{< relref "#ntpsec" >}})
1. [Visual Studio Code がベータ版に到達]({{< relref "#vscode" >}})

## The TPP must be rejected.{#tppip}

[Creative Commons] から TPP の分析とかなり強い調子の提言。

- [Trans-Pacific Partnership Would Harm User Rights and the Commons - Creative Commons](http://creativecommons.org/weblog/entry/46455) : 要旨
- [Trans-Pacific Partnership Would Harm User Rights and the Commons - Creative Commons](https://creativecommons.org/campaigns/trans-pacific-partnership-would-harm-user-rights-and-the-commons)

要点は以下のとおり（英語のままでごめん）

1. 20-year copyright term extension is unnecessary and unwarranted
2. The mention of the public domain is lip service, at best
3. Enforcement provisions are mandatory, while exceptions and limitations are optional
4. Potentially drastic infringement penalties, even for non-commercial sharing
5. Criminal penalties for circumventing digital rights management on works
6. Investor-state dispute settlement mechanism may be leveraged for intellectual property claims

最初の著作権期限の延長および2番目については，[これまで述べた]({{< ref "/remark/2015/1114-diary.md#tppip" >}})ように「公有財産の私有化」と言えるものだが，残りの4つはまさに「[コモンズの悲劇](https://ja.wikipedia.org/wiki/%E3%82%B3%E3%83%A2%E3%83%B3%E3%82%BA%E3%81%AE%E6%82%B2%E5%8A%87)」を地で行くような内容であり [Creative Commons] が強い調子に出るのも当然と言える。

「所有」の概念も「共有」の概念も希薄[^a] な日本では，このような強い調子に出る人や団体はない。
しかし TPP 知財は1国の問題ではなく，かつ今後100年単位で世界に影響を及ぼしうるものだ。

[^a]: 「所有」の概念が薄いからこそ「公有」や「共有」の概念も希薄だと言えるが。

さて，どうなるやら。

### 追記

- [著作権侵害の非親告罪化には「慎重であるべき」　文芸家協会が声明 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1511/25/news078.html)
- [日本ではパロディー認める判決は出ていない、コミケなど摘発の可能性――TPPによる著作権侵害の非親告罪化になおも懸念、日本文藝家協会が声明 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20151125_732191.html)

## 広島大学天文学研究会プラネタリウム & 天体観望会{#huaa}

- [広島大学天文学研究会プラネタリウム＆天体観望会を開催します | 広島大学](http://hiroshima-u.jp/moto/news/2015-11-13-1057)

わお。
広大のサイトで告知だよ。
偉くなったなぁ，天文研。
私の頃のようなヌルいサークル活動じゃなくなってるのかもな。

11月20日（金）の回は終わったけど，12月8日（火）の回はこれからなので，興味のある方は是非。
西条おっと東広島キャンパスの夜空はいいですよ。

私は平日は無理 orz

## NTPsec ベータ版が公開{#ntpsec}

- [NTP Security Project announces public development release](https://www.ntpsec.org/pressrelease-20151116.html)
- [Network Time Protocol（NTP）の脆弱性を改善する「NTPsec」、ベータ版が公開 － Publickey](http://www.publickey1.jp/blog/15/ntpsec09.html)

{{< fig-quote title="NTP Security Project announces public development release" link="https://www.ntpsec.org/pressrelease-20151116.html" lang="en" >}}
<q>The goal of the project is to harden against security vulnerabilities and especially against “amplification attacks” that threaten the stability of the entire Internet. The project welcomes the participation of information security researchers, and practices Responsible Disclosure.</q>
{{< /fig-quote >}}

NTP における「増幅攻撃（amplification attacks）」とは2014年に発生した NTP を使った DDoS を指していて

{{< fig-quote title="DNSよりも高い増幅率の「理想的なDDoSツール」：NTP増幅攻撃で“史上最大規模”を上回るDDoS攻撃発生 - ＠IT" link="http://www.atmarkit.co.jp/ait/articles/1402/12/news140.html" >}}
<q>というのも、monlistで問い合わせを送ると、そのNTPサーバーが過去に通信したマシン、最大600台分のIPアドレスを返答する。わずか234バイトの問い合わせパケットに対し、帰ってくる応答パケットのサイズは数十倍、数百倍という大きさだ。オープンなNTPサーバーに送信元を偽装したmonlistリクエストを送り付けると、標的には膨大なトラフィックが押し寄せる。<br>
この構図は、2013年のDDoS攻撃に悪用された「DNSリフレクション攻撃」（DNS amp攻撃）と同様だ。ただ、DNSリフレクション攻撃における増幅率が8倍であるのに対し、NTPのmonlistを悪用した増幅攻撃では、19倍から206倍という数字がはじき出せると、CloudFlareは説明している。</q>
{{< /fig-quote >}}

というものだ。
これから徐々に [NTPsec] に置き換わっていくのかな。

## Visual Studio Code がベータ版に到達{#vscode}

- [Announcing Visual Studio Code Beta - Visual Studio Code - Site Home - MSDN Blogs](http://blogs.msdn.com/b/vscode/archive/2015/11/17/announcing-visual-studio-code-beta.aspx)
- [［速報］無償のコードエディタ「Visual Studio Code」が、Go言語/Pascal/Reactなどに対応。ベータ版にも到達。Microsoft Connect(); 2015 － Publickey](http://www.publickey1.jp/blog/15/visual_studio_code_go_pascal.html)

さらに

- [［速報］オープンソース版.NETがリリース候補版に到達。Windows、MacOS X、Linuxで同一の.NETアプリが実行可能に。Microsoft Connect(); 2015 － Publickey](http://www.publickey1.jp/blog/15/netwindowsmacos_xlinux.html)
- [Microsoft、開発者向けの無償プログラム“Visual Studio Dev Essentials”を発表 - 窓の杜](http://www.forest.impress.co.jp/docs/news/20151119_731485.html)

大盤振る舞い！ まぁでも，残念ながらエディタは [ATOM に乗り換えちゃった]({{< ref "/remark/2015/atom-editor.md" >}})ので [Visual Studio Code] は当分出番なしかな。
仕事でも使う予定はないし。

新しい .NET はかなり期待している。
希望としては [Git Extensions](http://gitextensions.github.io/) や [KeePass](http://keepass.info/) など今まで [Mono](http://www.mono-project.com/) で動いていたものが新しい .NET に置き換えれるなら私が嬉しい。

[Creative Commons]: http://creativecommons.org/ "Creative Commons"
[NTPsec]: https://www.ntpsec.org/ "Welcome to NTPsec"
[Visual Studio Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined."
