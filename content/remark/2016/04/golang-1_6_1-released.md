+++
date = "2016-04-17T17:10:49+09:00"
description = "Go 言語に2つの脆弱性がある。脆弱性に対処した 1.6.1 および 1.5.4 がリリースされている。"
draft = false
tags = ["golang", "security", "vulnerability"]
title = "Go 言語 1.6.1 および 1.5.4 のセキュリティ・アップデート"

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

[Go 言語]に2つの脆弱性がある。
脆弱性に対処した 1.6.1 および 1.5.4 がリリースされている。

{{< fig-quote title="[security] Go 1.6.1 and 1.5.4 are released - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/9eqIHqaWvck" lang="en" >}}
<q>On Windows, Go loads system DLLs by name with LoadLibrary, making it vulnerable to DLL preloading attacks. For instance, if a user runs a Go executable from a Downloads folder, malicious DLL files also downloaded to that folder could be loaded into that executable.<br>
This is CVE-2016-3958 and was addressed by this change: <a href="https://golang.org/cl/21428">https://golang.org/cl/21428</a></q>
{{< /fig-quote >}}

{{< fig-quote title="[security] Go 1.6.1 and 1.5.4 are released - Google グループ" link="https://groups.google.com/forum/#!topic/golang-announce/9eqIHqaWvck" lang="en" >}}
<q>Go's crypto libraries passed certain parameters unchecked to the underlying big integer library, possibly leading to extremely long-running computations, which in turn makes Go programs vulnerable to remote denial of service attacks.  Programs using HTTPS client certificates or the Go SSH server libraries are both exposed to this vulnerability.<br>
This is CVE-2016-3959 and was addressed by this change: <a href="https://golang.org/cl/21533">https://golang.org/cl/21533</a></q>
{{< /fig-quote >}}

具体的に CVSS などを記述したページは見つからなかった。
が，順次更新する予定。

[Go 言語]: https://golang.org/ "The Go Programming Language"

## 参考図書

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/41aCueik45L._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/4621300253/baldandersinf-22/">プログラミング言語Go</a></dt><dd>Alan A.A. Donovan Brian W. Kernighan 柴田 芳樹 </dd><dd>丸善出版 2016-06-15</dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117526/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117526.09._SCTHUMBZZZ_.jpg"  alt="Go言語によるWebアプリケーション開発"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117607/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117607.09._SCTHUMBZZZ_.jpg"  alt="マイクロサービスアーキテクチャ"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4873117402/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4873117402.09._SCTHUMBZZZ_.jpg"  alt="ハイパフォーマンスPython"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/0134190440/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/0134190440.09._SCTHUMBZZZ_.jpg"  alt="The Go Programming Language (Addison-Wesley Professional Computing Series)"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/4774166340/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/4774166340.09._SCTHUMBZZZ_.jpg"  alt="Vim script テクニックバイブル ~Vim使いの魔法の杖"  /></a> </p>
<p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。買おうかどうか悩み中。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-03-12">2016-03-12</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>

<div class="hreview" ><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/"><img src="http://ecx.images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" alt="photo" class="photo"  /></a><dl ><dt class="fn"><a class="item url" href="http://www.amazon.co.jp/exec/obidos/ASIN/B015643CPE/baldandersinf-22/">暗号技術入門 第3版　秘密の国のアリス</a></dt><dd>結城 浩 </dd><dd>SBクリエイティブ 2015-08-25</dd><dd>評価<abbr class="rating" title="5"><img src="http://g-images.amazon.com/images/G/01/detail/stars-5-0.gif" alt="" /></abbr> </dd></dl><p class="similar"><a href="http://www.amazon.co.jp/exec/obidos/ASIN/B0148FQNVC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B0148FQNVC.09._SCTHUMBZZZ_.jpg"  alt="自作エミュレータで学ぶx86アーキテクチャ　コンピュータが動く仕組みを徹底理解！"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLJM/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLJM.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/丸い三角関数"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00Y9EYOIW/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00Y9EYOIW.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/微分を追いかけて"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B012BYBTZC/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B012BYBTZC.09._SCTHUMBZZZ_.jpg"  alt="情報セキュリティ白書2015: サイバーセキュリティ新時代：あらゆる変化へ柔軟な対応を"  /></a> <a href="http://www.amazon.co.jp/exec/obidos/ASIN/B00W6NCLL0/baldandersinf-22/" target="_top"><img src="http://images.amazon.com/images/P/B00W6NCLL0.09._SCTHUMBZZZ_.jpg"  alt="数学ガールの秘密ノート/数列の広場"  /></a> </p>
<p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
<p class="gtools" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="http://www.goodpic.com/mt/aws/index.html" >G-Tools</a>)</p>
</div>
