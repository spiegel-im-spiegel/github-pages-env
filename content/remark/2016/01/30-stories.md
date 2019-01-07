+++
date = "2016-01-30T19:19:39+09:00"
update = "2016-10-18T21:25:39+09:00"
description = "フェルミのパラドックス / GnuPG 2.1.11 released / Go 1.6 Release Candidate 1 / MIAU からの意見書"
draft = false
tags = ["astronomy", "seti", "security", "cryptography", "openpgp", "gnupg", "tools", "golang", "code", "politics", "intellectual-property"]
title = "週末スペシャル： フェルミのパラドックス"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

1. [フェルミのパラドックス]({{< relref "#seti" >}})
1. [GnuPG 2.1.11 released]({{< relref "#gpg" >}})
1. [Go 1.6 Release Candidate 1]({{< relref "#golang" >}})
1. [MIAU からの意見書]({{< relref "#miau" >}})

## フェルミのパラドックス{#seti}

人類の歴史は（宇宙誕生からの時間に比べれば）極々短いものだが，この歴史の中で私たちが ETI（Extra-terrestrial Intelligence; 地球外生命体）と直接的・間接的に接触したという記録はない。
宇宙に ETI やその文明がありふれているのなら，今まで地球人が接触しなかったのは何故？ というのが「フェルミのパラドックス」の内容であり，これが ETI 非存在を示す間接的な材料となっている。

最近このフェルミのパラドックスを説明する新しい説が出たらしい。

- [The Case for a Gaian Bottleneck: The Biology of Habitability](http://online.liebertpub.com/doi/10.1089/ast.2015.1387)
- [The aliens are silent because they are extinct | Science, Medicine and Health](http://science.anu.edu.au/whats-on/all-news/aliens-are-silent-because-they-are-extinct)
- [地球外生命からコンタクトがない理由 - アストロアーツ](http://www.astroarts.co.jp/news/2016/01/27exolife/index-j.shtml)

{{< fig-quote title="地球外生命からコンタクトがない理由" link="http://www.astroarts.co.jp/news/2016/01/27exolife/index-j.shtml" >}}
<q>もし金星や火星に初期の微生物が存在していたとしても、急速な環境の変化を安定にすることはできませんでした。一方、おそらく地球上の生物は、惑星の気候を安定させる重要な役割を果たしたのでしょう。いまだに地球外生命体の存在を示す兆候を見つけられていないのは、生物または知的生命体の誕生というよりも、惑星表面におけるフィードバック・サイクルの生物学的な制御が急速に起こることは稀だ、ということと関係が深いと思われます</q>
{{< /fig-quote >}}

銀河系内にどのくらい（地球人と交信可能な）知的文明が存在するか推定する方程式がある。
これが「ドレイク方程式」と呼ばれるものだ。

{{< fig-quote >}}
\[
    N = R_{*} \times f_p \times  n_e \times f_l \times f_i \times f_c \times L
\]
{{< /fig-quote >}}

各項の意味は

- $R_{*}$ ： 銀河系内で年間で誕生する恒星の数
- $f_p$ ： その恒星が惑星を持つようになる確率
- $n_e$ ： それらの中で生命の発生し得る条件を備えた惑星の数
- $f_l$ ： その惑星の中に実際に生命が発生し得る確率
- $f_i$ ： その生命が知能を持つに至る確率
- $f_c$ ： 彼らが実際に恒星間電波通信を行うまでに進歩する確率
- $L$ ： その文明の寿命

である。

この式の評価は色々あるのだが，金子隆一著『[ファースト・コンタクト―地球外知性体と出会う日](http://www.amazon.co.jp/exec/obidos/ASIN/4166600044/baldandersinf-22/)』では

- $R_{*} = 20$
- $f_p = 0.25$
- $n_e = 1.5$
- $f_l = 0.75$
- $f_i = 0.5$
- $f_c = 0.5$

と見積もられていた。
つまり $N = 1.4L$ となる。
仮に恒星間電波通信が可能な文明の寿命を1万年とするなら約1万4千の地球外文明が存在し得ることになる。
ちなみに地球人が電波を通信手段として使うようになってからまだ150年も経っていない。
さらに宇宙に向けて「アレシボ・メッセージ」を送ったのは1974年末。
「つい最近」の出来事である。

後半の $f_i$, $f_c$, $L$ 以外はほぼ物理的要因で決まると言ってよい。
また $L$ の唯一のサンプルは地球文明なので，地球文明が長く続ければ続くほど $L$ が大きく見積もれることになる。

しかし今回の論文から， $n_e$ や $f_l$ といった項が実はかなり小さい値なのではないか，と考えることもできる。
さて，現在の学者さん達はドレイク方程式をどのように解釈するだろうか。

そうそう。
[SETI@home](http://setiathome.ssl.berkeley.edu/) はバージョン 8 がリリースされている。
自宅マシンでは順調に稼働中。
Android 版もあるよ。

## GnuPG 2.1.11 released{#gpg}

[GnuPG] 2.1.11 が出た。

- [[Announce] GnuPG 2.1.11 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q1/000383.html)

セキュリティ・アップデートはなし。
主な変更点は以下のとおり。

* gpg: New command --export-ssh-key to replace the gpgkey2ssh tool.
* gpg: Allow to generate mail address only keys with --gen-key.
* gpg: "--list-options show-usage" is now the default.
* gpg: Make lookup of DNS CERT records holding an URL work.
* gpg: Emit PROGRESS status lines during key generation.
* gpg: Don't check for ambigious or non-matching key specification in the config file or given to --encrypt-to.  This feature will return in 2.3.x.
* gpg: Lock keybox files while updating them.
* gpg: Solve rare error on Windows during keyring and Keybox updates.
* gpg: Fix possible keyring corruption. (bug#2193)
* gpg: Fix regression of "bkuptocard" sub-command in --edit-key and remove "checkbkupkey" sub-command introduced with 2.1.  (bug#2169)
* gpg: Fix internal error in gpgv when using default keyid-format.
* gpg: Fix --auto-key-retrieve to work with dirmngr.conf configured keyservers. (bug#2147).
* agent: New option --pinentry-timeout.
* scd: Improve unplugging of USB readers under Windows.
* scd: Fix regression for generating RSA keys on card.
* dirmmgr: All configured keyservers are now searched.
* dirmngr: Install CA certificate for hkps.pool.sks-keyservers.net. Use this certiticate even if --hkp-cacert is not used.
* gpgtar: Add actual encryption code.  gpgtar does now fully replace gpg-zip.
* gpgtar: Fix filename encoding problem on Windows.
* Print a warning if a GnuPG component is using an older version of gpg-agent, dirmngr, or scdaemon.

## Go 1.6 Release Candidate 1{#golang}

[Go 言語]の 1.6 RC 版が登場。

- [Go 1.6 Release Candidate 1 is released - Google グループ](https://groups.google.com/forum/#!topic/golang-announce/pXcUuFHc4YE)

{{< fig-quote title="Go 1.6 Release Candidate 1 is released" link="https://groups.google.com/forum/#!topic/golang-announce/pXcUuFHc4YE" lang="en" >}}
<q>Our goal is to release the final version of Go 1.6 in around two weeks.</q>
{{< /fig-quote >}}

だそうで，楽しみである。

## MIAU からの意見書{#miau}

MIAU から「知的財産推進計画2016」が公開されている。

- [MIAU : 知的財産戦略本部「知的財産推進計画2016」策定に当たっての意見募集に、意見書を提出しました。](http://miau.jp/index1453962972.phtml)

内容については概ね同意。
つか， MIAU って仕事してるんだねぇ（笑）

でもこれって効果あるのかねぇ。
やらないよりはマシだろうけど。

[GnuPG]: https://www.gnupg.org/ "The GNU Privacy Guard"
[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4166600044/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41GPXP2HRVL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4166600044/baldandersinf-22/">ファースト・コンタクト―地球外知性体と出会う日 (文春新書)</a></dt><dd>金子 隆一 </dd><dd>文藝春秋 1998-10</dd><dd>評価<abbr class="rating" title="4"><img src="http://g-images.amazon.com/images/G/01/detail/stars-4-0.gif" alt="" /></abbr> </dd></dl><p class="similar"></p>
<p class="description">地球外文明探査の歴史を俯瞰する良書。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-01-30">2016-01-30</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
