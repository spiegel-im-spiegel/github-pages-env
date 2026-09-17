+++
title = "“Astronomy Picture of the Day” が移転・リニューアル"
date =  "2026-09-17T15:19:41+09:00"
description = "今月中に対応しろということらしい。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "astronomy", "photography", "nasa", "web", "api" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

NASA の “Astronomy Picture of the Day”（以下 APOD）は普段 Bluesky や Mastodon で更新を眺めているのだが，先日ひさしぶりに Web ページに行ってみたら

{{< fig-quote type="markdown" title="Astronomy Picture of the Day" link="https://apod.nasa.gov/" lang="en" >}}
**APOD's main NASA site is [moving](https://asterisk.apod.com/viewtopic.php?t=45023)**: From [apod.nasa.gov](https://apod.nasa.gov/) to [science.nasa.gov/apod](https://science.nasa.gov/apod)
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="APOD's main NASA site is moving - Starship Asterisk*" link="https://asterisk.apod.com/viewtopic.php?t=45023" lang="en" >}}
During 2026 August and September, APOD's main NASA website will be moving from https://apod.nasa.gov/ to https://science.nasa.gov/apod . Please be sure to update your bookmarks and other web pointers before the end of September.
{{< /fig-quote >}}

とアナウンスされていた。
今月中に対応しろということらしい。
おそらくだけど新しい方は WordPress で作成・運用されてるっぽい。
RSS フィードも一新されていて

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<rss version="2.0" xmlns:content="http://purl.org/rss/1.0/modules/content/" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:apod="https://science.nasa.gov/apod/">
<channel>
	<title>APOD Basic</title>
	<description>APOD Basic markup feed.</description>
	<link>https://science.nasa.gov/</link>
	<atom:link href="https://science.nasa.gov/feed/apod-basic/" rel="self" type="application/rss+xml" />
	<language>en-us</language>
	<item>
		<title>Webb&#039;s View of M64</title>
		<link>https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/</link>
		<guid isPermaLink="true">https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/</guid>
		<pubDate>Wed, 16 Sep 2026 04:05:00 +0000</pubDate>
		<apod:url>https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/</apod:url>
		<apod:hdurl>https://assets.science.nasa.gov/dynamicimage/assets/science/cds/apod/apod/2026/september/M64_Webb.jpg?w=3853&#038;h=4070&#038;fit=clip&#038;crop=faces%2Cfocalpoint</apod:hdurl>
		<apod:credit><![CDATA[NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)]]></apod:credit>
		<apod:copyright><![CDATA[NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)]]></apod:copyright>
		<apod:alt>Hubble and JWST image of M64. A massive spiral galaxy glows with a yellow core, surrounded by arms full of orange-brown dust and pink and blue patches of star formation. Framed by a haze of dark dust, the galaxy shines against black space dotted with a few stars.</apod:alt>
		<apod:primary_image_alt>Hubble and JWST image of M64. A massive spiral galaxy glows with a yellow core, surrounded by arms full of orange-brown dust and pink and blue patches of star formation. Framed by a haze of dark dust, the galaxy shines against black space dotted with a few stars.</apod:primary_image_alt>
		<apod:secondary_image_alt>Hubble image of M64. A massive spiral galaxy glows with a yellow core, surrounded by arms full of dark dust and pink and blue patches of star formation. Framed by a haze of dark dust, the galaxy shines against black space dotted with a few stars.</apod:secondary_image_alt>
		<apod:explanation><![CDATA[ ... ]]></apod:explanation>
		<description><![CDATA[... ]]></description>
		<content:encoded><![CDATA[<!doctype html> ... </html>]]></content:encoded>
	</item>
    ...
</channel>
</rss>
```

みたいな感じに `apod` の語彙を使って構造化されている。
`<item>` 内の主なフィールドは以下の通り。

{{< fig-quote class="nobox" type="markdown" title="APOD Feed And API User Guide" link="https://schlotterer.notion.site/APOD-Feed-And-API-User-Guide-39697d8747c38015a53edfdde76d4f5e" lang="en" >}}
| Field | Meaning |
| --- | --- |
| `title` | APOD title. |
| `link` | APOD post URL. |
| `guid` | Permanent item URL. |
| `pubDate` | WordPress publication date. |
| `apod:explanation` | APOD explanation text. |
| `apod:url` | APOD post URL. |
| `apod:hdurl` | Full-size featured image URL when available. |
| `description` | RSS-standard description field. It uses the same value as `apod:explanation`. |
| `content:encoded` | Full APOD Basic HTML document. |
{{< /fig-quote >}}

Atom はこういう拡張ができるからいいよな（XML データの解析は面倒くさいけど）。

なお RSS フィードでは年月を指定して過去のデータを取得することもできるようだ。

```text
$ curl -sS https://science.nasa.gov/feed/apod-basic/?year=2026&month=9
```

[NASA API] も大きく変わった。

{{< fig-quote type="markdown" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" >}}
The APOD API has been updated. It now returns information from https://science.nasa.gov/wp-json/wp/v2/apod-basic. The legacy APOD API will be archived on December 1, 2026. Full documentation on the new APOD API can be found [here](https://schlotterer.notion.site/APOD-Feed-And-API-User-Guide-39697d8747c38015a53edfdde76d4f5e).
{{< /fig-quote >}}

まず，上の引用のように API エンドポイントが変わった（旧エンドポイントは 2026-12-01 まで）。
これも WordPress かな？ 以前必要だった API キーも要らないようだ。
パラメータなしでリクエストを発行すると，最新25件の APOD データが取得できる。

```text
$ curl -sS https://science.nasa.gov/wp-json/wp/v2/apod-basic/ | jq .
[
  {
    "date": "2026-09-16",
    "post_id": 1339001,
    "title": "Webb's View of M64",
    "permalink": "https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/",
    "media_type": "image",
    "explanation": "...",
    "credit": "NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)",
    "copyright": "NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)",
    "alt": "Hubble and JWST image of M64. A massive spiral galaxy glows with a yellow core, surrounded by arms full of orange-brown dust and pink and blue patches of star formation. Framed by a haze of dark dust, the galaxy shines against black space dotted with a few stars.",
    "url": "https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/",
    "hdurl": "https://assets.science.nasa.gov/dynamicimage/assets/science/cds/apod/apod/2026/september/M64_Webb.jpg?w=3853&h=4070&fit=clip&crop=faces%2Cfocalpoint",
    "basic_html": "<!doctype html> ... </html>",
    "basic_html_url": "https://science.nasa.gov/wp-json/wp/v2/apod-basic/260916/html"
  },
  ...
]
```

また，パスに日付を `YYMMDD` 形式で指定して1件のみ取得することも可能である。

```text
$ curl -sS https://science.nasa.gov/wp-json/wp/v2/apod-basic/260916 | jq .
{
  "date": "2026-09-16",
  "post_id": 1339001,
  "title": "Webb's View of M64",
  "permalink": "https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/",
  "media_type": "image",
  "explanation": "...",
  "credit": "NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)",
  "copyright": "NASA, CSA, ESA, F. Belfiore (ESO), J. Lee (STScI), A. Leroy (OSU), and D. Thilker (JHU); <em>Processing</em>: G. Kober (NASA/Catholic University)",
  "alt": "Hubble and JWST image of M64. A massive spiral galaxy glows with a yellow core, surrounded by arms full of orange-brown dust and pink and blue patches of star formation. Framed by a haze of dark dust, the galaxy shines against black space dotted with a few stars.",
  "url": "https://science.nasa.gov/image-article/apod-2026-september-16-webbs-view-of-m64/",
  "hdurl": "https://assets.science.nasa.gov/dynamicimage/assets/science/cds/apod/apod/2026/september/M64_Webb.jpg?w=3853&h=4070&fit=clip&crop=faces%2Cfocalpoint",
  "basic_html": "<!doctype html> ... </html>",
  "basic_html_url": "https://science.nasa.gov/wp-json/wp/v2/apod-basic/260916/html"
}
```

他に指定できるパラメータは以下の通り。

{{< fig-quote class="nobox" type="markdown" title="APOD Feed And API User Guide" link="https://schlotterer.notion.site/APOD-Feed-And-API-User-Guide-39697d8747c38015a53edfdde76d4f5e" lang="en" >}}
| Parameter | What it does |
| --- | --- |
| `page` | Page number for list results. |
| `per_page` | Number of results per page. Capped at 25. |
| `date_from` | Start date using legacy `YYMMDD` format. |
| `date_to` | End date using legacy `YYMMDD` format. |
{{< /fig-quote >}}

レスポンスデータの主なフィールドは以下の通り。

{{< fig-quote class="nobox" type="markdown" title="APOD Feed And API User Guide" link="https://schlotterer.notion.site/APOD-Feed-And-API-User-Guide-39697d8747c38015a53edfdde76d4f5e" lang="en" >}}
| Field | Meaning |
| --- | --- |
| `date` | APOD date in `YYYY-MM-DD` format. |
| `post_id` | WordPress post ID. |
| `title` | APOD post title. |
| `permalink` | URL of the APOD post on the site. |
| `media_type` | Normalized APOD media type, such as `image`, `video`, or `iframe`. |
| `explanation` | APOD explanation text. |
| `url` | APOD post URL. This matches the permalink. |
| `hdurl` | Full-size featured image URL when available. |
| `basic_html` | APOD Basic HTML document as a JSON string. |
| `basic_html_url` | Plain-text HTML source URL for easier copy/paste workflows. |
{{< /fig-quote >}}

データが存在しない場合は以下のようにエラーが返ってくる。

```text
$ curl -sS https://science.nasa.gov/wp-json/wp/v2/apod-basic/260930 | jq .
{
  "code": "apod_basic_not_found",
  "message": "APOD not found.",
  "data": {
    "status": 404
  }
}
```

おー。
ちゃんとしてる。

最初にちょろんと書いたように APOD は Mastodon や Bluesky でも運用されている。
以下は公式が認めているアカウントである。

- [`@apod@reentry.codl.fr`](https://reentry.codl.fr/@apod "Astronomy Picture of the Day (@apod@reentry.codl.fr)") (mastodon)
- [`@apod.shinyakato.dev`](https://bsky.app/profile/apod.shinyakato.dev "Astronomy Picture of the Day 🪐 (@apod.shinyakato.dev) — Bluesky") (Bluesky)

非公式のものを含めれば，おそらく相当数のアカウントがあると思う。
ちなみに私も非公式で運用している。

- [`@fanapod.baldanders.info`](https://bsky.app/profile/did:plc:hbzmqswkx5pbg5fhr33pw4iw "Astronomy Picture of the Day (unofficial bot) (@fanapod.baldanders.info) — Bluesky") (Bluesky)
  - [`@fanapod.baldanders.info@bsky.brid.gy`](https://bsky.brid.gy/r/https://bsky.app/profile/fanapod.baldanders.info) : Mastodon へのブリッジ

こちらは API を使ってボット運用している。
文字数の関係で説明文（explanation）は省いているが，よろしかったら見てやって下さい。

[NASA API]: https://api.nasa.gov/ "NASA Open APIs"

## 参考

{{< linkcard "6b8ffab8684f95758ee5d25b952682638ba7889e" >}} <!-- 天文年鑑 2026年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
