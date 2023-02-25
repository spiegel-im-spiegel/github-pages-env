+++
title = "NASA API を使って “Astronomy Picture of the Day” のデータを取得するツールを作った"
date =  "2023-02-25T14:40:10+09:00"
description = "本当に作ってしまった（笑）"
image = "/images/attention/go-logo_blue.png"
tags = [ "tools", "astronomy", "golang", "photography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[先日の記事]({{< ref "/remark/2023/02/api-for-astronomy-picture-of-the-day.md" >}} "NASA API を使って “Astronomy Picture of the Day” のデータを取得する")で「これなら自作してみるのもいいかもなぁ」と書いたのだが，本当に作ってみた（笑）

- [goark/apod: CLI Tool for Astronomy Picture of the Day with NASA API](https://github.com/goark/apod)

このツールを動かすには NASA API キーが必要[^ak1] だが，ない場合はデモ用の `DEMO_KEY` を使って動作する。
ただし `DEMO_KEY` を使ったアクセスには以下の制限があるのでご注意を。

[^ak1]: NASA API キーの取得方法については「[NASA API を使って “Astronomy Picture of the Day” のデータを取得する]({{< ref "/remark/2023/02/api-for-astronomy-picture-of-the-day.md" >}})」の前半部分を参考にどうぞ。

{{< fig-quote type="markdown" title="NASA Open APIs" link="https://api.nasa.gov/" lang="en" >}}
- Hourly Limit: 30 requests per IP address per hour
- Daily Limit: 50 requests per IP address per day
{{< /fig-quote >}}

## 簡単な使い方

とりあえず，ヘルプはこんな感じ。

```text
$ apod -h
CLI Tool for Astronomy Picture of the Day with NASA API.

Usage:
  apod [flags]
  apod [command]

Available Commands:
  download    Download NASA APOD data
  help        Help about any command
  lookup      Look up NASA APOD data
  version     Print the version number

Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/username/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
  -h, --help                help for apod
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail

Use "apod [command] --help" for more information about a command.
```

`--date`, `--start-date`, `--end-date`, `--count`, `--thumbs`, `--api-key` 各フラグは APOD API の以下のパラメータに対応している。

| API Parameter | Flags | Description |
| --- | --- | --- |
| `date`       | `--date` | the date of the APOD image to retrieve |
| `start_date` | `--start-date` | the start of a date range |
| `end_date`   | `--end-date` | the end of the date range |
| `count`      | `--count` | randomly chosen images will be returned |
| `thumbs`     | `--thumbs` | return the URL of video thumbnail |
| `api_key`    | `--api-key` | `api.nasa.gov` key |

更にこれらのフラグの値は設定ファイル（Linux なら `$XDG_CONFIG_HOME/apod/config.yaml` が既定[^cfg1]）で指定することも可能。
こんな感じ。

[^cfg1]: `XDG_CONFIG_HOME` 環境変数がない場合は `$HOME/.config/apod/config.yaml` が既定ファイルになる。 Windows なら `%AppData%\apod\config.yaml`, macOS なら `$HOME/Library/Application/apod/config.yaml` を既定ファイルにしている。

```yaml
api-key: your_api_key_string
thumbs: true
```

## APOD 情報を取得する

```text
$ apod lookup -h
Look up NASA APOD data.

Usage:
  apod lookup [flags]

Aliases:
  lookup, look, l

Flags:
  -h, --help   help for lookup
      --raw    Output raw data from APOD API

Global Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/username/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail
```

とりあえず [jq] コマンドと組み合わせて

```text
$ apod lookup | jq .
[
  {
    "copyright": "Serge\nBrunier, Jean-François Bax, David Vernet",
    "date": "2023-02-24",
    "explanation": "Planetary nebula Jones-Emberson 1 is the death shroud of a dying Sun-like star. It lies some 1,600 light-years from Earth toward the sharp-eyed constellation Lynx. About 4 light-years across, the expanding remnant of the dying star's atmosphere was shrugged off into interstellar space, as the star's central supply of hydrogen and then helium for fusion was finally depleted after billions of years. Visible near the center of the planetary nebula is what remains of the stellar core, a blue-hot white dwarf star.  Also known as PK 164 +31.1, the nebula is faint and very difficult to glimpse at a telescope's eyepiece. But this deep broadband image combining 22 hours of exposure time does show it off in exceptional detail. Stars within our own Milky Way galaxy as well as background galaxies across the universe are scattered through the clear field of view. Ephemeral on the cosmic stage, Jones-Emberson 1 will fade away over the next few thousand years. Its hot, central white dwarf star will take billions of years to cool.",
    "hdurl": "https://apod.nasa.gov/apod/image/2302/jonesemberson1.jpg",
    "media_type": "image",
    "service_version": "v1",
    "title": "Jones-Emberson 1",
    "url": "https://apod.nasa.gov/apod/image/2302/jonesemberson1_1024.jpg"
  }
]
```

てな感じに JSON データを取得できる。
引数なしなら最新日のデータを， `--date` フラグで日付（`YYYY-MM-DD` 形式）を指定すれば指定した日付のデータを取得する。
`--start-date` および `--end-date` フラグで日付の範囲を指定することもできる。

## APOD 画像データをダウンロードする

```text
$ apod download -h
Download NASA APOD data.

Usage:
  apod download [flags]

Aliases:
  download, dl, d

Flags:
  -d, --base-dir string   Base directory for daownload (default "./apod")
  -h, --help              help for download
      --include-nopd      Download no public domain images or videos
      --overwrite         Overwrite Download files

Global Flags:
      --api-key string      NASA API key
      --config string       Config file (default /home/username/.config/apod/config.yaml)
      --count int           count randomly chosen images
      --date string         date of the APOD image to retrieve (YYYY-MM-DD)
      --debug               for debug
      --end-date string     end of a date range (YYYY-MM-DD)
      --start-date string   start of a date range (YYYY-MM-DD)
      --thumbs              return the URL of video thumbnail
```

`--base-dir`, `--include-nopd`, `--overwrite` 各フラグの値は設定ファイルで指定することも可能。

とりあえず

```text
$ apod download --include-nopd
```

とすれば，最新のデータと画像ファイルを `./apod` ディレクトリ以下にダウンロードできる。

```text
$ LANG=C ls -l apod
total 4
drwxrwxr-x 2 username username 4096 Feb 25 13:58 2023-02-24

$ LANG=C ls -l apod/2023-02-24
total 3376
-rw-rw-r-- 1 username username 3094863 Feb 25 13:58 jonesemberson1.jpg
-rw-rw-r-- 1 username username  352679 Feb 25 13:58 jonesemberson1_1024.jpg
-rw-rw-r-- 1 username username    1365 Feb 25 13:58 metadata.json
```

日付ごとにディレクトリを掘ってメタデータ（JSON 形式）と画像ファイルをダウンロードしているのがお分かりであろうか。

既に日付ディレクトリがある場合は何もしないが， `--overwrite` フラグが指定されていれば上書きダウンロードする。
また，基本的にクレジットがある画像データはダウンロードしない（public domain の画像データのみダウンロードする）が， `--include-nopd` フラグを指定すれば問答無用でダウンロードするようにした。

## 今後は...

基本機能は実装できた。
今後の方針としては

1. `apod download` 実行時のログを出力するようにする
2. ダウンロードを並行処理化
3. 前回からの差分データを自動的に取得

辺りを実装できればバッチ処理化できるかな。
NASA API は色々と面白いことができるみたいなので API 周りは別パッケージにして括り出したいところだが，これは当分先の話としよう。

まぁ，のんびりやります。

[jq]: https://stedolan.github.io/jq/

## ブックマーク

- [NASA Open APIs](https://api.nasa.gov/)

## 参考図書

{{% review-paapi "4416522940" %}} <!-- 天文年鑑 2023年版 -->
{{% review-paapi "4627275110" %}} <!-- 天体物理学 -->
