+++
title = "住所データに関する覚え書き"
date =  "2025-11-04T17:28:25+09:00"
description = "デジタル庁の「アドレス・ベース・レジストリ」について"
image = "/images/attention/tools.png"
tags = [ "engineering", "politics" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Bluesky の TL で見かけた記事から。

- [住所の「表記揺れ」解決で官民の負担削減、災害支援も　アドレス・ベース・レジストリが目指す「データを自動的に使える未来」 ｜デジタル庁ニュース](https://digital-agency-news.digital.go.jp/articles/2025-10-31-0)
- [アドレス・ベース・レジストリ｜デジタル庁](https://www.digital.go.jp/policies/base_registry_address)

書いてる内容がイマイチよく分からなかったので，順を追って整理してみよう。

## 住所は2つある

実際にやりとりされる住所には通り名とかあったり色々複雑だけど，お役所なんかでやりとりする住所には大きく2つあるらしい。

{{< fig-img-quote src="./20230314-Address-BR-01_v2.png" title="アドレス・ベース・レジストリ｜デジタル庁" link="https://www.digital.go.jp/policies/base_registry_address" width="1190" >}}

「住居表示を実施しているアドレス」というのは住所の最後の方に「○○番○○号」とか付いてるやつ。
一方「住居表示を実施していないアドレス」というのは土地に対する住所だそうで，不動産に絡む手続きなどで使われるらしい。
なんで2つに分かれているかというと根拠となる法律が違うからで，住居表示は「住居表示に関する法律（住居表示法）」，土地の住所は「不動産登記法」に基づいているそうだ[^j1]。

[^j1]: 住居表示がない場合もある。たとえば松江市役所の住所は「島根県松江市末次町86」だが，最後の「86」は街区符号でも住居番号でもなく地番である。

## 住所情報の管理主体は3つある

確かに法律は2つに分かれているけど「別に2つに分ける必要なくね？」って思うよね。
実は上述に示す住所の各要素を管理している主体は3つに分かれているらしい。

{{< fig-img-quote src="./ABR03.png" title="住所の「表記揺れ」解決で官民の負担削減、災害支援も　アドレス・ベース・レジストリが目指す「データを自動的に使える未来」 ｜デジタル庁ニュース" link="https://digital-agency-news.digital.go.jp/articles/2025-10-31-0" width="1190" >}}

じゃけこんな面倒臭いことになってるのか `orz`

都道府県および市区町村は全国地方公共団体コードとしてコード化されている。
街区符号・住居番号・地番はそもそもコードである。

問題は{{% ruby "まちあざ" %}}町字{{% /ruby %}}で，これがコード化されていなのである。

## アドレス・ベース・レジストリ

というわけでアドレス・ベース・レジストリ[^j2]（Address Base Registry; ABR）が整備されたということのようだ。
ABR では町字にコードが付与されただけでなく街区符号・住居番号・地番にも改めてコード（番号）が振られている[^j3]。
各データは[ダウンロードページ](https://dataset.address-br.digital.go.jp/ "アドレス・ベース・レジストリ ダウンロードサイト")から取得できる。

[^j2]: アドレス・ベース・レジストリの法律上の正式名称は「公的基礎情報データベース」と呼ぶらしい。
[^j3]: 例外的に街区符号や地番に数字以外の文字が含まれる場合があるらしい。

たとえば松江市役所の住所は「島根県松江市末次町86」なので

| コード種別 | ID | 名称 |
| :--- | :--- | :--- |
| 全国地方公共団体コード | 322016 | 島根県松江市 |
| 町字コード | 0083000 | 末次町 |
| 住居表示フラグ | 住居表示なし | － |
| 住居表示-街区符号 | － | － |
| 住居表示-住居番号 | － | － |
| 地番 | 000860000000000 | 86 |

となる。
末次町は「住居表示なし」なので住居表示マスターではなく地番マスターから検索する。

なお町字と郵便番号が対応している場合は，町字マスターに郵便番号が付与されている。
郵便番号が対応していない町字もあるので注意。
ホンマ面倒臭いよな。

ABR データの利用規約は以下を参照のこと。

- [アドレス・ベース・レジストリの利用規約｜デジタル庁](https://www.digital.go.jp/policies/base_registry_address_tos)

最後の方に

{{< fig-quote type="markdown" title="アドレス・ベース・レジストリの利用規約｜デジタル庁" link="https://www.digital.go.jp/policies/base_registry_address_tos" >}}
本利用ルールは、[クリエイティブ・コモンズ・ライセンスの表示4.0国際](https://creativecommons.org/licenses/by/4.0/legalcode.ja)に規定される著作権利用許諾条件。以下「CCBY」といいます。）と互換性があり、本利用ルールが適用されるコンテンツはCCBYに従うことでも利用することができます。
{{< /fig-quote >}}

と書かれているので [{{% cc-syms "cc" "by" %}}](https://creativecommons.org/licenses/by/4.0/ "Deed - Attribution 4.0 International - Creative Commons") ライセンス下で利用できるようだ。

## 住所データの正規化

ABR データだけでは「表記揺れ」は解決できない。
何らかの正規化処理が必要になる。

正規化ツールとして「[ABR ジオコーダー（abr-geocoder）](https://lp.geocoder.address-br.digital.go.jp/ "abr-geocoder | ABRジオコーダー（abr-geocoder）は、アドレス・ベース・レジストリ（ABR）で整備された住所・所在地を用いたジオコーダーです。")」が用意されている。

- [digital-go-jp/abr-geocoder: Address Base Registry Geocoder by digital.go.jp](https://github.com/digital-go-jp/abr-geocoder)
- [digital-go-jp/abr-geocoder-web: Address Base Registry Geocoder Web Client by digital.go.jp](https://github.com/digital-go-jp/abr-geocoder-web)

ABR ジオコーダーは API サービスまたは Web サービスを構築するためのツールキットで TypeScript で実装されているようだ。
どこかで API サービスを提供しているってわけじゃないみたい。

{{< fig-img-quote src="./2024072820391722166771.png" title="abr-geocoder | ABRジオコーダー（abr-geocoder）は、アドレス・ベース・レジストリ（ABR）で整備された住所・所在地を用いたジオコーダーです。" link="https://lp.geocoder.address-br.digital.go.jp/" width="1056" >}}

バックエンド DB は SQLite。
ABR データを SQLite にインポートするスクリプトも同梱されている。

正規化ルールは以下の通り：

{{< fig-quote type="markdown" title="正規化仕様" link="https://lp.geocoder.address-br.digital.go.jp/case.html" >}}
- [アドレス・ベース・レジストリ](https://www.digital.go.jp/policies/base_registry_address)の階層構造・表記に基づいた住所の正規化
- 都道府県省略を補完
  - `東京都千代田区` → `東京都千代田区`
  - `千代田区` → `東京都千代田区`
- 住居表示、地番をハイフン表記に正規化
  - `1番3号` → `1-3`
  - `1番地3` → `1-3`
- 全角数字を半角数字に正規化
  - `１番３号` → `1-3`
  - `１−３` → `1-3`
- 全角英字を半角英字に正規化
  - `ＤＩＧＩＴＡＬビル` → `DIGITALビル`
- 表記揺れ正規化
  - [JIS 第2水準 => 第1水準、旧字体 => 新字体 変換](https://github.com/digital-go-jp/abr-geocoder/blob/main/src/usecases/geocode/services/jis-kanji.ts)
  - [半角ｶﾅ => 全角カナ 変換](https://github.com/digital-go-jp/abr-geocoder/blob/main/src/usecases/geocode/services/to-katakana.ts)
  - [漢数字、全角数字 => 半角数字 変換](https://github.com/digital-go-jp/abr-geocoder/blob/main/src/usecases/geocode/services/kan2num.ts)
  - [全角英字 => 半角英字、全角数字 => 半角数字 変換](https://github.com/digital-go-jp/abr-geocoder/blob/main/src/usecases/geocode/services/to-hankaku-alpha-num.ts)
  - [揺らぎ => ひらがな 変換](https://github.com/digital-go-jp/abr-geocoder/blob/main/src/usecases/geocode/services/to-hiragana.ts)
  - 例
    - `壱`, `一`, `1`, `１`などの表記揺れに対応
    - `霞ケ関`, `霞ヶ関`, `霞ガ関` → `霞が関`
    - `篠ﾉ井`, `篠の井`, `篠之井` → `篠ノ井`
{{< /fig-quote >}}

Go でコマンドライン版で書けないかなぁ。
仕事ならやるんだけど（笑）

## 【おまけ1】 Geolonia による住所正規化ツール

[Geolonia] は ABR を使った API サービスを構築するためのツールを提供している。

- [geolonia/japanese-addresses-v2: 全国の住所データAPI](https://github.com/geolonia/japanese-addresses-v2)

さらに正規化ライブラリも提供しているようだ。

- [geolonia/normalize-japanese-addresses: オープンソースの住所正規化ライブラリ。](https://github.com/geolonia/normalize-japanese-addresses)

[Geolonia] では有料のサービスも提供していて実績もあるみたいなので，業務で使うなら検討してもいいかもしれない。

## 【おまけ2】 CSV ファイルエディタ CSVI

[CSVI] はテキストベースのターミナル上で動作する CSV ファイルエディタだ。

- [hymkor/csvi: Terminal CSV Editor](https://github.com/hymkor/csvi)

Go 製のツールで Windows および UNIX 系の OS で動作する。
シングルバイナリ構成で Windows, Linux, FreeBSD および macOS 用のバイナリが[提供](https://github.com/hymkor/csvi/releases "Releases · hymkor/csvi")されている。

Ubuntu 機に入れて使っているが，ビュアーとしてもすごく快適。

{{< fig-img src="./csvi.png" title="csvi スナップショット" link="./csvi.png" width="1920" >}}

かなり大きいファイルでもストレスなく開けるっぽいので，普段ターミナル上で作業している人にはオススメである。

## ブックマーク

- [住所コードと郵便番号に関する覚え書き]({{< ref "/remark/2020/09/address-code-and-postal-code.md" >}})

[GIS]: https://nlftp.mlit.go.jp/
[Geolonia]: https://geolonia.com/ "Geolonia - An Elastic Map Hosting - Geolonia"
[geolonia/japanese-addresses-v2]: https://github.com/geolonia/japanese-addresses "geolonia/japanese-addresses: 全国の町丁目レベル（189,540件）の住所データのオープンデータ"
[CSVI]: https://github.com/hymkor/csvi "hymkor/csvi: Terminal CSV Editor"
