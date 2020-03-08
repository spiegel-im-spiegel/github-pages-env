+++
title = "SETI@home 20年間の区切り"
date =  "2020-03-08T12:56:44+09:00"
description = "とはいえ SETI@home プロジェクト自体は閉鎖されず，サイトとユーザ・アカウントは保持される。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "seti" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[SETI@home] が休眠するらしい。

- [SETI@home hibernation](https://setiathome.berkeley.edu/forum_thread.php?id=85267)

休眠する最大の理由はあらかたのデータ解析が終わったからのようだ。

{{< fig-quote type="markdown" title="SETI@home hibernation" link="https://setiathome.berkeley.edu/forum_thread.php?id=85267" lang="en" >}}
We're doing this for two reasons:

1. Scientifically, we're at the point of diminishing returns; basically, we've analyzed all the data we need for now.
2. It's a lot of work for us to manage the distributed processing of data. We need to focus on completing the [back-end analysis](http://setiathome.berkeley.edu/nebula/index.php) of the results we already have, and writing this up in a scientific journal paper.
{{< /fig-quote >}}

とはいえ [SETI@home] プロジェクト自体は閉鎖されず，サイトとユーザ・アカウントは保持される。

{{< fig-quote type="markdown" title="SETI@home hibernation" link="https://setiathome.berkeley.edu/forum_thread.php?id=85267" lang="en" >}}
{{% quote %}}However, SETI@home is not disappearing. The web site and the message boards will continue to operate. We hope that other UC Berkeley astronomers will find uses for the huge computing capabilities of SETI@home for SETI or related areas like cosmology and pulsar research. If this happens, SETI@home will start distributing work again. We'll keep you posted about this{{% /quote %}}.
{{< /fig-quote >}}

科学的学術研究テーマとしての SETI (Search for Extra-Terrestrial Intelligence; 地球外知的生命探査) は広範囲かつ長期間の観測が要求される割に実りが少なく，分の悪い賭けといえる。

ボランティアとして参加する私達一般ユーザはともかく，プロジェクトを統括・運営するプロの研究者達は無償というわけにはいかない。
また（プロなのだから）区切られた期間できちんと結果を出さなければならない。
もちろんシステムを維持するためのコストもある（最初の頃はサーバや回線の確保が大変だった）。

もともと [SETI@home] は2年間の期限付きプロジェクトだった。
それが世界中で予想以上の反響を巻き起こし3年4年と延長され，ついには [BOINC] なる汎用プラットフォームを生み出した[^boinc1]。

[^boinc1]: もともと [BOINC] は “Berkeley Open Infrastructure for Network Computing” の略称だったが，現在は単に “[BOINC]” と呼ばれているようだ。

[BOINC] 移行後は当時の熱狂も収まり（資金難に陥ることも度々あったが）ボランティア・ユーザが淡々と解析を行うスタイルで継続し，なんだかんだと正式稼働から20年も続いたわけだ。
天文学関連の学術プロジェクトで20年というのはさして長いわけではないだろうが（太陽観測なら2周期分に満たない`w`），このテーマでよくもこれだけもったというのが正直な感想である。

現在の [SETI@home] は純粋に宇宙人探しをしているわけではなく AstroPulse[^ap1] などの関連プロジェクトに相乗りする形で運営されている。
したがって，今後 [SETI@home] と相乗り可能なプロジェクトが興れば再開される可能性がある。
あるいは全く違うアプローチのプロジェクトになるか。

[^ap1]: AstroPulse は全天で特徴的なパルスを放つ天体を捜索するプロジェクトで [SETI@home] と相性がよかった。 AstroPulse で発見された連星パルサーが [Einstein@Home] の観測対象になったこともある。

{{< fig-quote type="markdown" title="SETI@home hibernation" link="https://setiathome.berkeley.edu/forum_thread.php?id=85267" lang="en" >}}
{{% quote %}}We're extremely grateful to all of our volunteers for supporting us in many ways during the past 20 years. Without you there would be no SETI@home. We're excited to finish up our original science project, and we look forward to what comes next{{% /quote %}}.
{{< /fig-quote >}}

[SETI@home] の[次回作にご期待ください](https://dic.pixiv.net/a/%E5%85%88%E7%94%9F%E3%81%AE%E6%AC%A1%E5%9B%9E%E4%BD%9C%E3%81%AB%E3%81%94%E6%9C%9F%E5%BE%85%E3%81%8F%E3%81%A0%E3%81%95%E3%81%84)，といったところか（笑）

なお，私が参加している天文学関連の [BOINC] プロジェクトは（[SETI@home] 以外では）以下の通り。

- [Einstein@Home]
- [Asteroids@home](http://asteroidsathome.net/)

[BOINC] では複数のプロジェクトを掛け持ちして効率よく回すのがコツである。
「[計算機による信号傍受と解析は人類のためにこそ使いましょう](https://web.archive.org/web/*/http://www.planetary.or.jp/setiathome/)」。

Happy analyzing !!

## ブックマーク

- [地球外知的生命体探査プロジェクト“SETI@home”が3月31日で一旦休止 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1238940.html)

- [Club-HUAA サポートページ](http://huaa.baldanders.info/)

[BOINC]: https://boinc.berkeley.edu/
[SETI@home]: https://setiathome.berkeley.edu/
[Einstein@Home]: https://einsteinathome.org/

## 参考図書

{{% review-paapi "4166600044" %}} <!-- ファースト・コンタクト -->
{{% review-paapi "4898140866" %}} <!-- SETI@homeファンブック -->
