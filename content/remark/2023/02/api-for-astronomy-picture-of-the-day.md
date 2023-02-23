+++
title = "NASA API を使って “Astronomy Picture of the Day” のデータを取得する"
date =  "2023-02-23T12:27:00+09:00"
description = "これなら自作してみるのもいいかもなぁ。"
image = "/images/attention/kitten.jpg"
tags = [ "astoronomy", "golang", "photography" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
+++

Mastodon の自 TL を眺めてたら

- [marcusziade/apod-cli: A command-line tool to browse the NASA Astronomy Picture of the Day archive.](https://github.com/marcusziade/apod-cli)

というリポジトリが紹介されていた。
NASA が毎日更新している “[Astronomy Picture of the Day][APOD]” のデータを取得できるツールだそうな。
コマンドライン・インタフェースで pure [Go] で書かれているらしい。
当然マルチプラットフォーム。

ふむむ。
さっそく試してみよう。

## NASA API キーを取得する

NASA では自身が保有するデータを Web API を使って取得できるらしい。

- [NASA Open APIs](https://api.nasa.gov/)

件のツールはこの API を使うようで，まずは API キーを取得する必要がある。

なお，お試し用に `DEMO_KEY` という API キーが用意されている。

{{< fig-quote type="markdown" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" >}}
In documentation examples, the special `DEMO_KEY` api key is used. This API key can be used for initially exploring APIs prior to signing up, but it has much lower rate limits, so you’re encouraged to signup for your own API key if you plan to use the API (signup is quick and easy). The rate limits for the `DEMO_KEY` are:

- Hourly Limit: 30 requests per IP address per hour
- Daily Limit: 50 requests per IP address per day
{{< /fig-quote >}}

とりあえず軽く試すのであれば，これを使ってもいいかもしれない。

API の取得は簡単。
“[NASA Open APIs](https://api.nasa.gov/)” ページの以下のフォームに必要事項を入力して `[Signup]` すればいい。

{{< fig-img-quote src="./generate-api-key-1.png" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" width="1019" >}}

するとこんな感じに結果が返ってくる（大事なとこ塗りつぶしてます，ご容赦）。
メールアドレス以外の機微データは要求しないので比較的お気軽に使える。

{{< fig-img-quote src="./generate-api-key-2.png" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" width="1009" >}}

同じ内容が登録したメールにも送られる。
早速試してみよう（便宜上，キーを `DEMO_KEY` で表記している）。

```text
$ curl "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY" | jq .
{
  "copyright": "Mehmet Ergün",
  "date": "2023-02-22",
  "explanation": "Our Sun is becoming a busy place.  Only two years ago, the Sun was emerging from a solar minimum so quiet that months would go by without even a single sunspot.  In contrast, already this year and well ahead of schedule, our Sun is unusually active, already nearing solar activity levels seen a decade ago during the last solar maximum.  Our increasingly active Sun was captured two weeks ago sporting numerous interesting features. The image was recorded in a single color of light called Hydrogen Alpha, color-inverted, and false colored.  Spicules carpet much of the Sun's face.  The brightening towards the Sun's edges is caused by increased absorption of relatively cool solar gas and called limb darkening.  Just outside the Sun's disk, several scintillating prominences protrude, while prominences on the Sun's face are known as filaments and show as light streaks.  Magnetically tangled active regions are both dark and light and contain cool sunspots.  As our Sun's magnetic field winds toward solar maximum over the next few years, whether the Sun's high activity will continue to increase is unknown.",
  "hdurl": "https://apod.nasa.gov/apod/image/2302/SunHalphaC_Ergun_2065.jpg",
  "media_type": "image",
  "service_version": "v1",
  "title": "Our Increasingly Active Sun",
  "url": "https://apod.nasa.gov/apod/image/2302/SunHalphaC_Ergun_960.jpg"
}
```

おー。
できたできた。

[APOD] データ取得用 Web API の URL クエリパラメータとして以下のものが使える。

{{< fig-quote class="nobox" type="markdown" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" >}}

| Parameter | Type | Default | Description |
| --- | --- | --- | --- |
| `date` | `YYYY-MM-DD` | today | The date of the APOD image to retrieve |
| `start_date` | `YYYY-MM-DD` | none | The start of a date range, when requesting date for a range of dates. Cannot be used with date. |
| `end_date` | `YYYY-MM-DD` | today | The end of the date range, when used with start_date. |
| `count` | `int` | none | If this is specified then `count` randomly chosen images will be returned. Cannot be used with `date` or `start_date` and `end_date`. |
| `thumbs` | `bool` | False | Return the URL of video thumbnail. If an APOD is not a video, this parameter is ignored. |
| `api_key` | `string` | `DEMO_KEY` | `api.nasa.gov` key for expanded usage |
{{< /fig-quote >}}

では，昨日（2023-02-21）今日（2023-02-22）の2日間のデータを取ってみよう

```text
$ curl "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY&start_date=2023-02-21&end_date=2023-02-22" | jq .
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2522  100  2522    0     0   2208      0  0:00:01  0:00:01 --:--:--  2208
[
  {
    "copyright": "Tara Mostofi",
    "date": "2023-02-21",
    "explanation": "They are both falling. The water in Yosemite Falls, California, USA, is falling toward the Earth. Comet ZTF is falling toward the Sun. This double cosmic cascade was captured late last month as fading Comet C/2022 E3 (ZTF) had just passed its closest to planet Earth. The orange star just over the falls is Kochab. With the exception of a brief encounter with a black bear, the featured image was a well-planned composite of a moonlit-foreground and long-duration background exposures - all designed to reconstruct a deep version of an actual single sight. Although Comet ZTF is now fading as it glides back to the outer Solar System, its path is determined by gravity and so it can be considered to still be falling toward the Sun -- but backwards.    Comet ZTF Gallery: Notable Submissions to APOD",
    "hdurl": "https://apod.nasa.gov/apod/image/2302/CometZtfYosemite_Mostofi_1639.jpg",
    "media_type": "image",
    "service_version": "v1",
    "title": "Comet ZTF  over Yosemite Falls",
    "url": "https://apod.nasa.gov/apod/image/2302/CometZtfYosemite_Mostofi_960.jpg"
  },
  {
    "copyright": "Mehmet Ergün",
    "date": "2023-02-22",
    "explanation": "Our Sun is becoming a busy place.  Only two years ago, the Sun was emerging from a solar minimum so quiet that months would go by without even a single sunspot.  In contrast, already this year and well ahead of schedule, our Sun is unusually active, already nearing solar activity levels seen a decade ago during the last solar maximum.  Our increasingly active Sun was captured two weeks ago sporting numerous interesting features. The image was recorded in a single color of light called Hydrogen Alpha, color-inverted, and false colored.  Spicules carpet much of the Sun's face.  The brightening towards the Sun's edges is caused by increased absorption of relatively cool solar gas and called limb darkening.  Just outside the Sun's disk, several scintillating prominences protrude, while prominences on the Sun's face are known as filaments and show as light streaks.  Magnetically tangled active regions are both dark and light and contain cool sunspots.  As our Sun's magnetic field winds toward solar maximum over the next few years, whether the Sun's high activity will continue to increase is unknown.",
    "hdurl": "https://apod.nasa.gov/apod/image/2302/SunHalphaC_Ergun_2065.jpg",
    "media_type": "image",
    "service_version": "v1",
    "title": "Our Increasingly Active Sun",
    "url": "https://apod.nasa.gov/apod/image/2302/SunHalphaC_Ergun_960.jpg"
  }
]
```

うんうん。
問題なさそうだね。

## [apod-cli][`apod-cli`] を動かしてみる

準備が整ったので，いよいよ [`apod-cli`] を使ってデータを取得してみよう。

[`apod-cli`] のバイナリはリポジトリの[最新リリースページ](https://github.com/marcusziade/apod-cli/releases/latest)から取得できる。
Windows, macOS (Darwin), Linux 版が用意されているので，それぞれのプラットフォームに合ったものをダウンロードして，中身の `apod-cli` ファイルをパスが通るフォルダかカレントフォルダに配置する。
macOS であれば Homebrew も使えるらしい。
一覧にないプラットフォームであれば自前でビルドするしかないかな。

何も設定しないで動かすと

```test
$ ./apod-cli 
API key not found in Keys.json
Please sign up for an API key at https://api.nasa.gov/#signUp
Once you have your API key, enter it below:
```

と怒られるので， `Keys.json` ファイルを作成する。
中身はこんな感じ（値は伏せ字にしています，ご容赦）。

```json
{
    "APIKey": "*****************"
}
```

このファイルをカレントフォルダに置いておけばいいようだ。
あらためて起動すると

```text
$ ./apod-cli 
Fetching APODs...


The Hydra Cluster of Galaxies
February 16, 2023
https://apod.nasa.gov/apod/image/2302/ABELL1060_LRGB_NASA.jpg

2023 CX1 Meteor Flash
February 17, 2023
https://apod.nasa.gov/apod/image/2302/gijsDSC_1917(2x3)1600px.jpg

Barred Spiral Galaxy NGC 1365 from Webb
February 18, 2023
https://apod.nasa.gov/apod/image/2302/JWSTMIRI_ngc1365.png

Seven Dusty Sisters in Infrared
February 19, 2023
https://apod.nasa.gov/apod/image/2302/Pleiades_WiseAntonucci_5000.jpg

NGC 1850: Not Found in the Milky Way
February 20, 2023
https://apod.nasa.gov/apod/image/2302/Ngc1850_HubbleOzsarac_2048.jpg

Comet ZTF  over Yosemite Falls
February 21, 2023
https://apod.nasa.gov/apod/image/2302/CometZtfYosemite_Mostofi_1639.jpg

Our Increasingly Active Sun
February 22, 2023
https://apod.nasa.gov/apod/image/2302/SunHalphaC_Ergun_2065.jpg
```

おぉ。
直近一週間分のデータを取るのか（期間はコマンドライン引数で指定できるらしい）。
画像ファイルを落としてきたりとかはしないんだな。
まぁ，こんなもんか。

NASA API 自体は簡単な RESTful API みたいだし，これなら自作してみるのもいいかもなぁ。

## ブックマーク

- [GitHub - HelixSpiral/apod: Golang wrapper for the Astronomy Picture of the Day API from NASA](https://github.com/HelixSpiral/apod)
- [GitHub - dynafa/apod: Golang NASA APOD Web App](https://github.com/dynafa/apod)
- [Astronomy Picture of the Day (@APoD@botsin.space) - botsin.space](https://botsin.space/@APoD) : 非公式 bot

[Go]: https://go.dev/
[APOD]: https://apod.nasa.gov/ "Astronomy Picture of the Day"
[`apod-cli`]: https://github.com/marcusziade/apod-cli "marcusziade/apod-cli: A command-line tool to browse the NASA Astronomy Picture of the Day archive."

## 参考文献

{{% review-paapi "4416522940" %}} <!-- 天文年鑑 2023年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
{{% review-paapi "B0BW9BDRP5" %}} <!-- Way to go : アニメ「神達に拾われた男2」OP曲。 -->
