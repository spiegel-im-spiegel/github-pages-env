+++
title = "JSON 形式で Feed を出力する"
date = "2019-02-19T22:54:05+09:00"
description = "Hugo では JSON 形式のファイルも出力できる。"
image = "/images/attention/hugo.jpg"
tags = [ "hugo", "feed", "json" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回も小ネタ。

サイトの feed は RSS や Atom といった XML 形式を使うのが普通だが，他サイトの連携を考えるなら JSON 形式で出力できれば色々と便利そうである。
[Hugo] では JSON 形式のファイルも出力できる。

## JSON 形式のファイルを出力するよう設定する。

まずは JSON 形式のファイルを出力できるよう設定ファイル（`config.toml` など）に以下の記述を追加する。

```toml
[outputs]
  home = ["HTML", "RSS", "JSON"]
```

これでサイトのルートディレクトリに `index.html`, `index.xml`, そして `index.json` ファイルを出力できるようになった。

## JSON ファイルのテンプレートを作成する

テンプレートファイルを格納する `layouts/` フォルダの直下または `layouts/_default/` フォルダに `index.json` または `home.json` ファイルを作成する。
中身は例えばこんな感じ。

```text
{
  "title": {{ .Site.Title | jsonify }},
  "url": {{ .Permalink | jsonify }},{{ with .OutputFormats.Get "RSS" }}
  "feed": {{ .Permalink | jsonify }},{{ end }}
  "entry": [{{ range $i, $v := first 3 .Data.Pages }}{{ if ne $i 0 }},{{ end }}{{ with $v }}
    {
      "title": {{ .Title | jsonify }},{{ with .Section }}
      "section": {{ . | jsonify }},{{ end }}{{ with .Description }}
      "description": {{ . | jsonify }},{{ end }}
      "url": {{ .Permalink | jsonify }},
      "published": {{ .Date.UTC.Format "2006-01-02T15:04:05-07:00" | jsonify }},
      "update": {{ .Lastmod.UTC.Format "2006-01-02T15:04:05-07:00" | jsonify }}
    }{{ end }}{{ end }}
  ]
}
```

これでコンパイルを行うと以下のような内容で `index.json` ファイルが出力される。

```json
{
  "title": "text.Baldanders.info",
  "url": "https://text.baldanders.info/",
  "feed": "https://text.baldanders.info/index.xml",
  "entry": [
    {
      "title": "“Article 13” に関するブックマーク",
      "section": "remark",
      "description": "英語不得手の私としては難儀していたが，ようやく日本語の記事を見かけるようになった。",
      "url": "https://text.baldanders.info/remark/2019/02/bookmarks-for-article-13/",
      "published": "2019-02-19T13:04:51+00:00",
      "update": "2019-02-19T13:05:49+00:00"
    },
    {
      "title": "2019-02-17 のブックマーク",
      "section": "bookmarks",
      "description": "「火星ローバー「オポチュニティ」ミッション終了、15年の活動に幕」他",
      "url": "https://text.baldanders.info/bookmarks/2019/02/17-bookmarks/",
      "published": "2019-02-17T12:25:23+00:00",
      "update": "2019-02-17T12:25:48+00:00"
    },
    {
      "title": "Hugo で Git 情報を取得する",
      "section": "hugo",
      "description": "Hugo の環境を git で管理している場合はコミット情報等をテンプレートに組み込むことができる。",
      "author": "Spiegel",
      "license": "https://creativecommons.org/licenses/by-sa/4.0/",
      "url": "https://text.baldanders.info/hugo/gitinfo/",
      "published": "2019-02-16T13:15:35+00:00",
      "update": "2019-02-16T13:20:26+00:00"
    }
  ]
}
```

[Hugo] では出力ファイルについてかなり細かく制御することが出来るが，今回はこのくらいにしておこう（笑） そのうち詳しい記事を書くこともあるかも知れない。

## ブックマーク

- [Custom Output Formats | Hugo](https://gohugo.io/templates/output-formats/)
- [GitHub - hugojapan/hugoDocs: Hugo 日本語ドキュメント](https://github.com/hugojapan/hugoDocs)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
