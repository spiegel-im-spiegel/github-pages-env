+++
title = "カレンダーに祝日を入れたいなら国立天文台へ行けばいいじゃない"
date =  "2019-05-14T22:37:13+09:00"
description = "日本の暦はこれが国家公式データである。"
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "ephemeris", "japanese", "calendar" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回のネタ]({{< relref "./tokyo-2020.md" >}})は個人的に色々と不本意だったので，今回は軽い小ネタで。
よく考えたらこの記事が今年初めての天文ネタだよ `orz`

- [Google カレンダーに日本の祝祭日を表示したい！ カレンダーを追加する方法 - 窓の杜](https://forest.watch.impress.co.jp/docs/serial/chrometips/1184245.html)

という記事を見かけたのだが，日本の祝日は毎年国立天文台で発表してるんだから国立天文台のデータを使えばいいじゃない，と思ったり。

日本の暦は[国立天文台]の[暦計算室]で見ることができる。
こんな感じのページ。

{{< fig-img src="./google-ephemeris.png" title="今月のこよみ powered by Google Calendar - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/cande/calendar.html" width="824" >}}

しかも Google カレンダーへインポート可能で，上のページの「＋Google Calendar」の部分をクリックすれば自身の Google カレンダーにインポートできる。

祝日だけでなく二十四節気や朔望月，あるいは日食・月食といった情報が取得できる。
必要に応じてオン・オフを切り替えればいいだろう。

なお，上記データで見れるのは「祝日および休日」で「祭日」のデータはない。
つか，祭日は国家が規定するものではないので，あるわけないのだが。
祭日が知りたければ[高島暦](https://www.amazon.co.jp/dp/B07JKP4CG8?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "令和2年神宮館高島暦 | 神宮館編集部 | 占い | Kindleストア | Amazon")とかを購入することをオススメする（笑）

ちなみに国立天文台では毎年2月1日に翌年の「[暦要項](https://eco.mtk.nao.ac.jp/koyomi/yoko/ "暦要項 - 国立天文台暦計算室")」が官報で発表される。

- [平成31（2019）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2018/20180201-rekiyoko.html)
- [平成32（2020）年暦要項の発表 | 国立天文台(NAOJ)](https://www.nao.ac.jp/news/topics/2019/20190201-rekiyoko.html)

日本の暦はこれが国家公式データである。

[国立天文台]: https://www.nao.ac.jp/ "国立天文台(NAOJ)"
[暦計算室]: https://eco.mtk.nao.ac.jp/koyomi/ "国立天文台 天文情報センター 暦計算室"

## 参考図書

{{% review-paapi "B07Y3FGB99" %}} <!-- 令和2年神宮館高島暦 -->
{{% review-paapi "4805202254" %}} <!-- 天体の位置計算 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
