+++
title = "Google vs Oracle の訴訟の行方（最終章）"
date =  "2021-04-11T11:00:27+09:00"
description = "これが日本の法制下であれば，おそらく Google は Oracle に勝てないだろう。 "
image = "/images/attention/kitten.jpg"
tags = ["code", "intellectual-property", "copyright", "java", "google", "android"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先月に「[Google vs Oracle の訴訟の行方 3]({{< ref "/remark/2021/03/google-vs-oracle-3.md" >}})」を書いたばかりで，私としては判決が出るのは早くても夏以降だろうと高をくくっていたのだが，4月早々に判決が出ちゃいました。

- [Supreme Court sides with developers in Google v. Oracle - The GitHub Blog](https://github.blog/2021-04-06-supreme-court-sides-with-developers-in-google-v-oracle/)
- [グーグル、米最高裁でオラクルに勝訴--「Android」Javaコード訴訟で - ZDNet Japan](https://japan.zdnet.com/article/35168881/)
- [［速報］10年にわたる著作権訴訟でGoogleがオラクルに勝訴、米連邦最高裁判所で判決。Java SEのコードのコピーはフェアユースの範囲 － Publickey](https://www.publickey1.jp/blog/21/10googlejava_se.html)

これにより Andorid における Java API の利用（コピー）は fair use の範疇であることが認められた。

念のために書いておくけど，ライブラリやフレームワーク等で提供される API (Application Programming Interface) そのものには著作権があり，その利用（複製・配布・改変）について規制がかかる，という点は覆らない。
その上で各 API の利用が fair use の範疇であるか否かについては個別の案件となる。
これが今回の訴訟における最重要ポイントである。

これが日本の法制下であれば，おそらく Google は Oracle に勝てないだろう。
何故なら，日本には fair use 規定がなく，争うとすれば「著作権の制限」から攻めるしかないが，「著作権の制限」には「API の公正利用」などという状況は想定されていないからだ[^fu1]。
これだけ見ても，いかに日本の著作権法が時代遅れで「周回遅れ」も宜なるかな，と実感させられるだろう。

[^fu1]: [2019年初の著作権法改正]({{< ref "/remark/2018/11/copyright-law-is-revised.md" >}})で「著作権の制限」の幾つかが拡張されたが，これは AI 等へのデータ利用を想定したもので API 等のコード要素そのものを想定したものではない。まぁ「AI 等へのデータ利用」についてようやく2019年初で改正された，というのも数年の周回遅れではあるけどね。

もうひとつ，念のために書いておくと，著作権が規制するのは「表現」であって「アイデア」ではない。
故に著作物を「使用」するだけなら著作権は関知しないし[^use1]，先行者優位もないので独立で公表される「表現」がどれだけ似ていてもお互いに規制することはできない。
まぁ，お互いに「独立の表現」であることを証明するのは，今の情報が溢れかえる時代に於いては「悪魔の証明」に近いかもしれないけど（笑）

[^use1]: まぁ，でも， API を「利用」することなく「使用」するという状況は思い浮かばないけど（笑） たとえば Web API のようにプロトコルを規定しているだけの場合でも，プロトコルそのものに著作権が適用されるのなら同じことだ。もちろん API を通してやり取りされるデータにも（データベースの）著作権があるのでご注意を。本当に著作権はエンジニアにとって面倒くさい規制だよ。

## 参考図書

{{% review-paapi "B0CTM1KHDX" %}} <!-- 著作権法 第4版 -->
{{% review-paapi "4757102852" %}} <!-- 著作権２．０ ウェブ時代の文化発展をめざして -->
{{% review-paapi "4757122349" %}} <!-- 〈反〉知的独占 -->
{{% review-paapi "B01CYDGUV8" %}} <!-- CODE VERSION 2.0 -->
{{% review-paapi "4622073455" %}} <!-- 〈海賊版〉の思想‐18世紀英国の永久コピーライト闘争 -->
{{% review-paapi "4797334673" %}} <!-- インターネットの法と慣習 -->
