+++
title = "GnuPG 030周年（笑）と寄付の話"
date =  "2021-12-21T06:36:50+09:00"
description = "寄付は「消費」ではない。 個々人の意思による明確な社会参加の形である。"
image = "/images/attention/kitten.jpg"
tags = [ "code", "engineering", "gnupg" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

12月20日は最初の [GnuPG] がリリースされた日ということで記念の tweet が上がっていたが

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">It&#39;s now 030 years since the first version of GnuPG was released. (Back then under its working title g10.)<br>After all these years we are now on solid financial grounds and do not anymore need to ask for donations. Many Thanks to everyone who made this possible. <a href="https://t.co/gr3Mc0LuRA">pic.twitter.com/gr3Mc0LuRA</a></p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1472942349865136130?ref_src=twsrc%5Etfw">December 20, 2021</a></blockquote>
{{< /fig-gen >}}

とあり，さらに

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">To our donors: Please stop your donations to GnuPG and instead donate to projects which still need financial support</p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1472943158346629125?ref_src=twsrc%5Etfw">December 20, 2021</a></blockquote>
{{< /fig-gen >}}

と tweet されている。
8進数は思いつかなんだ。
再来年は私も真似しよう（笑）

寄付云々の話はピンとこない人もいるかもしれないが，実は [GnuPG] プロジェクトは財政的に困窮していた時期があり，作者の Werner Koch はプロジェクトを放棄することも考えていたらしい。

- [The World's Email Encryption Software Relies on One Guy, Who is Going Broke — ProPublica](https://www.propublica.org/article/the-worlds-email-encryption-software-relies-on-one-guy-who-is-going-broke)
- [Once-starving GnuPG crypto project gets a windfall. Now comes the hard part | Ars Technica](https://arstechnica.com/information-technology/2015/02/once-starving-gnupg-crypto-project-gets-a-windfall-but-can-it-be-saved/)

当時は Edward Snowden による NSA 内部告発や [OpenSSL の脆弱性](https://baldanders.info/blog/000682/ "パスワード変更は計画的に")で発覚したオープンソース・プロジェクトにおける人的・経済的リソースの不足などが背景としてあり，個人・法人を問わず寄付が集まったおかげで [GnuPG] で人を雇えるほど財政が持ち直した経緯がある。

[前にも書いた]({{< ref "/remark/2021/12/open-source-is-unbreakable.md" >}} "オープンソースは砕けない")けど，オープンソースの功績・貢献をお金に換えることについては昔から議論があるが，一般的に「これだ！」と言えるものはない。
一時期はデジタル補完通貨[^cc1] と FinTech で皆ウハウハみたいな話もあったが，あれって本当に儲かってるのは miner と山師だけだろう。

[^cc1]: 暗号通貨とか暗号資産とか言ってやらん（笑）

寄付は「消費」ではない。
個々人の意思による明確な[社会参加](http://shinta.tea-nifty.com/nikki/2005/01/donation.html "寄付する前に立ち止まれ")の形である。
この手の話は今後も泡沫のように出てくる話題なのだろう。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
