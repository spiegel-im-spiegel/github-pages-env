+++
title = "TechCrunch Japan 終了後の後始末"
date =  "2022-05-04T16:03:59+09:00"
description = "翻訳記事 URL を可能な限り原文記事 URL に書き換えてみる。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "media", "web", "vscode", "regexp" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2月に [TechCrunch Japan が終了してバックナンバーも残さず消滅する話を書いた]({{< ref "/remark/2022/02/the-nation-of-amnesia.md" >}} "記憶喪失の国")。
んで，実際に GW 中にサイトが消滅したわけだが，以前の URL を叩いてみたところ 404 ではなく[本家 TechCrunch][TechCrunch] にリダイレクトされるようだ。

いや，そこまでしてくれるなら，せめて翻訳記事は[本家][TechCrunch]の原文記事にリダイレクトしてくれよ `orz`

まぁ，愚痴ってもしょうがない。
こちらで可能な限り URL の書き換えを試みることにしよう。

まずこのブログ・サイトの[作業リポジトリ](https://github.com/spiegel-im-spiegel/github-pages-env "spiegel-im-spiegel/github-pages-env: Document Environment for spiegel-im-spiegel.github.io")上で TechCrunch Japan の URL がどのくらいあるか軽く `grep` してみる[^grep1]。

[^grep1]: 私の環境では [mattn/jvgrep](https://github.com/mattn/jvgrep "mattn/jvgrep: grep for japanese vimmer") を `grep` に alias して使っている。ファイル指定を `"content/**/*.md"` などと再帰的に指定できるのが嬉しい。

```text
$ grep -c "jp\.techcrunch\.com" "content/**/*.md"
1121
```

おぅふ。
アホほどあるがな `orz`

## TechCrunch Japan 記事の URL を機械的に変換できるか

たとえば TechCrunch Japan 記事の URL を

> `https://jp.techcrunch.com/2020/08/14/2020-08-13-instagram-delete-photos-messages-servers/`

とする。
この記事に対する原文記事の URL は

> [`https://techcrunch.com/2020/08/13/instagram-delete-photos-messages-servers/`](https://techcrunch.com/2020/08/13/instagram-delete-photos-messages-servers/)

である。
ドメインが `jp.techcrunch.com` → [`techcrunch.com`][TechCrunch] なのは当然として

1. 翻訳記事と原文記事では URL パスの日付部分が違う
2. 原文記事の日付は翻訳記事の slug に含まれている
3. 日付部分を除く slug の文字列は翻訳記事と原文記事で同じ

これくらいなら正規表現を使った置換処理で何とかなりそうだ。
最近のテキスト・エディタは置換処理で正規表現が使えるものが多いが，私が愛用している [VS Code] でも正規表現を使った一括置換が可能である。

- [Visual Studio Codeを用いた簡単な正規表現検索 - Qiita](https://qiita.com/kgsi/items/a88662c6e43fa5311288)

## {{% ruby "AMP" %}}例外{{% /ruby %}}を潰す

私の作業環境で2箇所ほど例外というか間違いがあって

- `https://jp.techcrunch.com/2017/12/12/2017-12-11-some-hp-laptops-are-hiding-a-deactivated-keylogger/amp/`
- `https://jp.techcrunch.com/2020/01/03/2020-01-02-ex-google-policy-chief-dumps-on-the-tech-giant-for-dodging-human-rights/amp/?__twitter_impression=true`

などと，うっかり AMP 用の URL を載せちゃったみたいで，しかも片方は変なパラメータがくっついている。
これらも機械的に置換できなくはないのだが，2箇所だけだし，手作業で原文記事の URL に書き換えた。

AMP ページはマジで滅びて欲しい。
なんでこんな下らないことで Google に気を使わにゃならんの。
メディアが気を遣うべき相手は私ら閲覧者だろうが。
本末転倒だよ。

あと，古い URL でスキーマが HTTP のままになってるのが結構あったので，これは `http://jp.techcrunch.com` → `https://jp.techcrunch.com` に一括置換した。

## Slug パターン

前節の例外を排除したことで TechCrunch Japan 記事のURL

> `https://jp.techcrunch.com/yyyy/mm/dd/slug/`

のうち slug 部分にのみ注目すればよくなった。
この Slug 部分も複数のパターンが見受けられるので整理しておく

### パターン1: 日付情報 yyyy-mm-dd を含む Slug

最初に挙げた例の通り `yyyy-mm-dd-originalslug` に要素分解できるパターン。
このパターンには別のバリエーションがあって

- `https://jp.techcrunch.com/2020/07/15/x2020-07-14-harvard-mit-sue-ice-student-visas-rule/`
- `https://jp.techcrunch.com/2020/11/21/https-techcrunch-com-2020-11-20-google-facebook-and-twitter-threaten-to-leave-pakistan-over-censorship-law/`

のように日付情報の前に余分な文字列がくっついている。
2番目のとか原文記事の URL そのままぢゃん。
「なにすんねん」ってツッコんじゃったよ（笑）

### パターン2: 日付情報 yyyymmdd を含む Slug

以下のような URL パターン：

- `https://jp.techcrunch.com/2017/09/13/20170912new-bluetooth-vulnerability-can-hack-a-phone-in-ten-seconds/`
- `https://jp.techcrunch.com/2016/07/08/automotive-fortune-tesla20160706tesla-says-drivers-using-autopilot-remain-safer-than-regular-drivers/`

パターン1のハイフンが抜けた状態。

### パターン3: Slug に日付情報がない

- `https://jp.techcrunch.com/2021/06/10/netflix-cowboy-bebop-streaming-this-fall/`

`jp.techcrunch.com` → [`techcrunch.com`][TechCrunch] と置換するだけで行けるかなぁと思ったが駄目だった（[本家サイト][TechCrunch]が404になる）。
原文記事の日付情報が得られないので置換不可。

### パターン4: Slug が[パーセント・エンコーディング]({{< ref "/golang/uri-encoding.md" >}} "URI エンコーディングについて")されている

- `https://jp.techcrunch.com/2017/03/13/%e3%80%8c%e6%b3%95%e4%bb%a4%e4%b8%8a%e9%81%95%e5%8f%8d%e3%81%ae%e5%8f%af%e8%83%bd%e6%80%a7%e3%80%81%e5%80%ab%e7%90%86%e7%9a%84%e3%81%ab%e3%82%82%e5%95%8f%e9%a1%8c%e3%80%8ddena%e3%81%8cwelq%e5%95%8f/`

多分，というか間違いなく日本版オリジナル記事だよね。
これは置換対象外とした。

## 置換用正規表現

というわけで，今回はパターン1と2のみが対象となる。
置換処理は [VS Code] を使っている。

パターン1の検索・置換正規表現は以下の通り。

|  | 正規表現 |
| --- | --- |
| 検索 | `https://jp\.techcrunch\.com/\d{4}/\d{2}/\d{2}/.*(\d{4})-(\d{2})-(\d{2})-(.+)/` |
| 置換 | `https://techcrunch.com/$1/$2/$3/$4/` |
| 対象ファイル | `*.md` |

パターン2の検索・置換正規表現は以下の通り。

|  | 正規表現 |
| --- | --- |
| 検索 | `https://jp\.techcrunch\.com/\d{4}/\d{2}/\d{2}/.*(\d{4})(\d{2})(\d{2})(.+)/` |
| 置換 | `https://techcrunch.com/$1/$2/$3/$4/` |
| 対象ファイル | `*.md` |

もう少し頑張ればひとつにまとめられたかもしれないが，副作用が出るのが嫌だったので分けた。
これで未変換の TechCrunch Japan 記事の URL は118個まで減ったが，今のところ，これ以上は無理なので，放置ということで。

どっとはらい

## ブックマーク

- [基本的な正規表現一覧 | murashun.jp](https://murashun.jp/article/programming/regular-expression.html)

[TechCrunch]: https://techcrunch.com/ "TechCrunch – Startup and Technology News"
[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
