+++
title = "節分の狂躁"
date =  "2025-02-02T19:47:10+09:00"
description = "今回は過去に書いた節分・立春絡みの記事を再構成してお送りする"
image = "/images/attention/kitten.jpg"
tags = [ "osanpo-camera", "photography", "matsue", "bike", "calendar", "ephemeris", "golang" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

今回は過去に書いた節分・立春絡みの記事を再構成し，更にいくつか記事を足してお送りする。

## 立春も動き出す

「立春」は二十四節気のひとつで，現代の定義は「太陽黄経が315度になる瞬間を含む日」である。
春分点が黄経0度で冬至点が黄経270度（$\frac{3}{2}\pi = \frac{3}{4}\tau\\,\mathrm{rad}$）で[^tau1]，立春はその中間という感じ。
分かりやすい。

[^tau1]: $\tau$ については「[π は間違ってる？]({{< ref "/remark/2023/03/pi-is-wrong.md" >}})」を参照のこと。

{{< fig-img-quote src="/remark/2025/02/setsubun/Celestial_sphere_in_Japanese.png" title="Explanation about a celestial sphere in Japanese" link="https://commons.wikimedia.org/wiki/File:Celestial_sphere(in_Japanese).png" lang="en" >}}

「節分」は元々は（暦上の）各季節の終わりにあったが，現代では雑節のひとつとして「立春の前日」と定義されている。
他の季節は（公式には）なくなっちゃったのね（笑）

今年は閏年の翌年ということで節分&立春が1日前倒しになっている。
これは前回の2021年からの傾向で少なくとも今世紀中は続く。
つか，だんだん2月3日が立春の年が普通になっていくんだけどね。

{{< fig-img-quote src="./topics2021_2a.png" title="節分の日が動き出す - 国立天文台暦計算室" link="https://eco.mtk.nao.ac.jp/koyomi/topics/html/topics2021_2.html" width="1000" >}}

詳しくは4年前に書いた節分を参照のこと。

- [立春も動き出す]({{< ref "/remark/2021/01/the-beginning-of-spring.md" >}})

## 節分の狂躁

私の子供の頃は「恵方巻」を食べる習慣などなかったのだが，広島で独り暮らしをしている間に実家では定番行事になっていた。
まぁ，いまどき玄関にイワシの頭を飾る家のほうが珍しいのかもしれないが，えっと変わらんよね。

というわけでおつかいで近所のスーパーへ恵方巻を買いに行く。

{{< fig-img src="./54300241862_4898676717_e.jpg" title="節分の狂躁 | Flickr" link="https://www.flickr.com/photos/spiegel/54300241862/" >}}

つくづく馬鹿な風習だよな，と思う。
色街の下品なお遊び程度に留めておけばよかったのに（あれも「諸説」のひとつに過ぎないけどさ）。

恵方巻については以下の調査が面白い。

- [食卓の縁起に関する研究 I　－恵方巻の受容とその背景－ – 煌道 生活文化研究所](https://seikatsubunka.biz/archives/81)
  - {{< pdf-file title="食卓の縁起に関する研究 I　－恵方巻の受容とその背景－" link="http://seikatsubunka.biz/main/wp-content/uploads/2017/04/12.pdf" >}}

こうやって「{{% ruby "forklore" %}}民俗{{% /ruby %}}」というのは形成されてくんだねぇ。

ちなみに今日の我が家の夕食は，買ってきた恵方巻と買い置きの出雲蕎麦と（従兄の畑でできた大根で）ふろふき大根。
あと正月からの飲み残しの「[七冠馬](https://sake-hikami.co.jp/guide/ "七冠馬を知る - 簸上清酒合名会社 | 島根県奥出雲町にある「七冠馬」「玉鋼」の簸上（ひかみ）清酒")」。
ハレの日だからね（笑）

おおごっつぉでした。

## 2025年の恵方

さて恵方巻といえば今年の恵方はどっち？ ということで。

「恵方」というのはその年の歳神様のおられる方位で，居住地から見て恵方にあたる社寺に詣ることを「恵方詣り」と言うらしいのだが，明治以降の鉄道の発達で長距離移動が容易になり，方位に依存する「恵方詣り」が廃れ，代わりに「初詣」が主流になっていったとのこと。
その一方で，節分の「恵方巻」みたいな文化が平成以降に台頭してくるのは面白い。

恵方はその年の十干で決定するらしい。

{{< fig-img-quote src="/golang/favourable-direction/Ehou-direction.png" title="File:Ehou-direction.png - Wikimedia Commons" link="https://commons.wikimedia.org/wiki/File:Ehou-direction.png" lang="en" >}}

って，4方向を十干で分けるのか。
中途半端やなぁ。

拙作の [`github.com/goark/koyomi`](https://github.com/goark/koyomi "goark/koyomi: 日本のこよみ") パッケージで指定した年の干支と恵方を取得できるようにした。
こんな感じ。

```go
package main

import (
    "fmt"

    "github.com/goark/koyomi/zodiac"
)

func main() {
    year := 2025
    干, 支 := zodiac.ZodiacYearNumber(year)
    fmt.Printf("%d年は%v%v，恵方は%v (%v°)", year, 干, 支, 干.DirectionJp(), 干.Direction())
}
```

これを[実行](https://go.dev/play/p/4U7xf-Fted2)すると「2025年は乙巳，恵方は西南西微西 (255°)」と出力される。
まぁ，ほぼ西。

というわけで，今年はほぼ西を向いてモグモグしましょう。
...ホンマ，馬鹿な風習だよな。

## 節分祭

昨日はまるっとオフだったので，いつものように[八雲温泉][八雲温泉ゆうあい熊野館]へ行ったのだが，となりの[熊野大社]では既に一部で屋台が出ていた。

{{< fig-img src="./54299712939_38cd1a0fa4_e.jpg" title="節分祭前日（熊野大社） | Flickr" link="https://www.flickr.com/photos/spiegel/54299712939/" >}}

節分祭は翌日だというのに商魂たくましいというか。
栗まんじゅうをひとつ所望してしまった。

{{< fig-img src="./54299730953_c4b241edc5_e.jpg" title="栗まんじゅうゲットだぜ！ | Flickr" link="https://www.flickr.com/photos/spiegel/54299730953/" >}}

あぁ，糖分が染み渡る。

明けて今日。
午後から出掛けられそうなのだが，2日連続で[熊野大社]というのも芸がない。
そうだ。
[佐太神社]へ行ってみよう。

{{< fig-img src="./54300552172_ec9b3b2063_e.jpg" title="佐太神社（別に祭じゃなかった？） | Flickr" link="https://www.flickr.com/photos/spiegel/54300552172/" >}}

あれ？ いつもと変わらん？ いや境内まで行けば何かやってるかも知れんけど。
神事そのものには興味が薄いので[^s1]，スルーして道の駅でお茶した。

[^s1]: 民俗としての神道や仏教やその他の宗教・信仰には興味があるけど，宗教そのものへの関心は薄い。というか無神論者だし，私。いや，他者の信仰を否定するつもりはないのよ。日本みたいに宗教メーカーが弱くて怪しげなカルト（もどき）や詐欺師が横行するよりは宗教に縛られるほうが100倍はマシだと思うし。まぁ，米国みたいに「[金持ちこそ正義]({{< ref "/remark/2016/09/anti-intellectualism.md" >}} "ようやく『反知性主義』を読んだ")」みたいな宗教観も「なんだかなぁ」と思うけど。

{{< fig-img src="./54300567637_ea5f183b58_e.jpg" title="糖分補給中（道の駅 秋鹿なぎさ公園） | Flickr" link="https://www.flickr.com/photos/spiegel/54300567637/" >}}

## 明日から春

明日から暦の上では春っスよ。
つっても4日から松江も氷点下予報だけどね。
光熱費高騰の折ではありますが，暖かくしてお過ごしくださいませ。

これでこっぽし。

## ブックマーク

- [今年の恵方を取得する]({{< ref "/golang/favourable-direction.md" >}})
- [お散歩カメラ 2024-02-03]({{< ref "/remark/2024/02/03-osanpo-camera.md" >}})

{{< fig-youtube id="6SQVZgKea-0" title="節分って2月3日じゃないの？ ―みんな気になる暦のふしぎ― - YouTube" >}}

[八雲温泉ゆうあい熊野館]: https://www.kumanokan.jp/ "八雲温泉ゆうあい熊野館"
[熊野大社]: http://www.kumanotaisha.or.jp/ "出雲國一之宮　熊野大社"
[佐太神社]: http://www.sadajinjya.jp/ "佐太神社公式ホームページ"

## 参考

{{% review-paapi "B0191845R0" %}} <!-- 鉄道が変えた社寺参詣 -->
{{% review-paapi "B00EUVZHX0" %}} <!-- 神道入門 -->
{{% review-paapi "B012VRLPRG" %}} <!-- 反知性主義 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
{{% review-paapi "B08L4WKDZ7" %}} <!-- PowerShot ZOOM -->
{{% review-paapi "B08DFNBWBM" %}} <!-- 恵根餅 Enemoti 三種お試し用 -->
{{% review-paapi "B09TQLY5VZ" %}} <!-- アミノバイタル ゼリードリンク -->
{{% review-paapi "B0BP9PNX2P" %}} <!-- 冬木透：交響詩ウルトラセブン on Brass -->
{{% review-paapi "B099ZXJWMV" %}} <!-- 冬木透 帰ってきたウルトラマン オリジナル・サウンドトラック<ウルトラサウンド殿堂シリーズ> -->

## 作業中の BGV (メン限配信以外)

- [おしゃべり唐あげあげ太くん 「遠回りの人生」 - YouTube](https://www.youtube.com/watch?v=LZuqc7xqXSI)
- [ムーンライト / 星街すいせい covered by ReGLOSS 【歌ってみた / hololive DEV_IS】 - YouTube](https://www.youtube.com/watch?v=m6T71r7hRMI)
- [晴る / ヨルシカ  covered by ReGLOSS 【歌ってみた / hololive DEV_IS】 - YouTube](https://www.youtube.com/watch?v=Mb2Pdf7P-a4)
- [【雑談】「しらたきください」って言ったらはんぺん出てきたりLINEスタンプ考えたりする＋後半スパチャ読み【儒烏風亭らでん #ReGLOSS 】 - YouTube](https://www.youtube.com/watch?v=XoHPk1Lqahc)
- [追悼 ウルトラ音楽の父・冬木透先生 - YouTube](https://www.youtube.com/watch?v=NmJQlvR0_7s)
- [Ultrasiete Concierto de Toru Fuyuki Ultraseven Concert 冬木透 CONDUCT ウルトラセブン - YouTube](https://www.youtube.com/watch?v=I7xBu-IqlLQ)
- [【💬 #ピノまど観測所】生き物大好きお嬢様を観測せよ！ 実はサシの雑談コラボ、初めてです【カルロ・ピノ / 星見まどか】 - YouTube](https://www.youtube.com/watch?v=NUH_vy5jXUg)
- [【🌃 2月の星空案内】夜空に惑星たちが大集合！ ではあるけれども……🌟宇宙大好きVTuberが天文現象や星座の情報などをお届け！【星見まどか】 - YouTube](https://www.youtube.com/watch?v=EfW2vMMlUpU)
- [【#昭和100年歌枠リレー】いにしえの鬼が昭和の懐かしい名曲たちを歌います！【古代日本史VTuber きら子】 - YouTube](https://www.youtube.com/watch?v=xcW6-wECZ5M)
