+++
title = "CVSS v3.1"
date =  "2020-01-26T17:20:03+09:00"
description = "調べてみたら 3.1 って2019年6月にリリースされてたんだねぇ。 半年以上前の話だよ。 感度低いなぁ orz"
image = "/images/attention/kitten.jpg"
tags = ["security", "vulnerability", "risk", "management", "cvss"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回の記事]({{< relref "./pre-announcement-of-go-1_13_7.md" >}} "Go 1.13.7 リリース予告と CVE-2020-0601")を書いてて気がついたのだが CVSSv3 のバージョンが 3.1 に上がってるぢゃん。

たとえば [CVE-2020-0601] の CVSS ベクタは

- `CVSS:3.0/AV:N/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:N`
- `CVSS:3.1/AV:N/AC:L/PR:N/UI:R/S:U/C:H/I:H/A:N`

の2つが用意されている。

調べてみたら 3.1 って2019年6月にリリースされてたんだねぇ。
半年以上前の話だよ。
感度低いなぁ，私 `orz`

3.0 と 3.1 の仕様は以下のリンクから見れる。

- [CVSS v3.0 Specification Document](https://www.first.org/cvss/v3.0/specification-document)
- [CVSS v3.1 Specification Document](https://www.first.org/cvss/v3.1/specification-document)

差分情報がないのでひっじょーに分かりにくいのだが[^wp1]，各評価基準の項目と値に変更はなく，スコア算出式のみ変更になっているようだ。
しかも変更されているのは環境評価基準（Environmental Metrics）だけのようなので，私達がよく見る基本評価基準（Base Metrics）は変更なしと見てよさそうだ。

[^wp1]: 少なくとも技術文書をワープロで書くのは止めてほしいのだが。最終出力は PDF でも構わないが（PDF で出力するなら PDF/A で）， Markdown でも AsciiDoc でも Org-mode でもいいから，入力はプレイン・テキストで管理して欲しい。したら簡単に差分が取れるでしょ。

なので上述の CVSS ベクタのスコアはいずれも 8.1 で深刻度（Severity Rating）も「重要（High）」となる。

実は [Go 言語]で CVSS のパッケージを作って公開しているのだが

- [spiegel-im-spiegel/go-cvss: Common Vulnerability Scoring System (CVSS) Version 3](https://github.com/spiegel-im-spiegel/go-cvss)

基本評価基準しか実装してないんで大っぴらにしていない。
仕事で使いそうなら続きを作り込もうかと思っていたのだが，職業エンジニア自体が無期休業中だからねぇ（笑）

なお [spiegel-im-spiegel/go-cvss] については，一応 v3.1 のベクタも受け入れるようにした。

## ブックマーク

- [共通脆弱性評価システムCVSS概説：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/vuln/CVSS.html) : v3.1 に対応しているようだ
    - [共通脆弱性評価システムCVSS v3概説：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/vuln/CVSSv3.html) : v3.1 に対応していないように見えるのだが...
- [CVSS（共通脆弱性評価システム）3.0から3.1への変更点：OpenSCAPで脆弱性対策はどう変わる？（7） - ＠IT](https://www.atmarkit.co.jp/ait/articles/2005/19/news002.html)

- [JVN が CVSSv3 による脆弱性評価を開始]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})

[CVE-2020-0601]: https://nvd.nist.gov/vuln/detail/CVE-2020-0601
[Go 言語]: https://golang.org/ "The Go Programming Language"
[spiegel-im-spiegel/go-cvss]: https://github.com/spiegel-im-spiegel/go-cvss "spiegel-im-spiegel/go-cvss: Common Vulnerability Scoring System (CVSS) Version 3"
<!-- eof -->
