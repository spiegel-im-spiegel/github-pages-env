+++
title = "住所コードと郵便番号に関する覚え書き"
date =  "2020-09-09T18:30:36+09:00"
description = "個人が興味本位で弄るのは無理，という結論になった。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

昨日リリースした [spiegel-im-spiegel/cov19data](https://github.com/spiegel-im-spiegel/cov19data "spiegel-im-spiegel/cov19data: Importing WHO COVID-2019 Cases Global Data") パッケージなんだけど，東京都の PCR 検査陽性者のデータも参照できるようにしている。

- [東京都 新型コロナウイルス陽性患者発表詳細 - データセット - 東京都オープンデータカタログサイト](https://catalog.data.metro.tokyo.lg.jp/dataset/t000010d0000000068)


今までは日付毎に行数をカウントするだけでデータの詳細に興味はなかったのだが，今回改めて CSV ファイルを眺めてみると「全国地方公共団体コード」のカラムがあったので「これなら都道府県名や市区町村名いらなくね？」と気がついた。

ついでにこの住所コードと郵便番号を関連付ければ面白いデータセットが組めるんじゃないかと安直に考えたのだが，どうも簡単な話ではないようだ。
ていうか，個人が興味本位でやるのは無理，という結論になった。

今回は，その辺の話を覚え書きとして記しておく。

## 住所コード

ひとくちに住所コードと言っても，様々なレイヤがある。
以下で細かく見てみよう。

### JIS 都道府県コード

JIS 規格で決められているコードで，都道府県ごとに `01` から `47` までの2桁の数字列で表される。
たとえば，島根県の都道府県コードは `32` である。

### JIS 住所コードと全国地方公共団体コード

市区町村までを表す住所コードは，都道府県コード＋3桁の合計5桁の数字列で表され，これも JIS 規格で決められている。
たとえば，島根県松江市は `32201` である。

更に，住所コードの末尾にチェック・ディジットを付加した6桁を「全国地方公共団体コード」と呼ぶ。
ちなみにチェック・ディジットの計算手順は以下の通り。

1. 住所コード $abcde$ の各桁に対して $a \times 6 + b \times 5 + c \times 4 + d \times 3 + e \times 2$ を求める *(1)*
2. *(1)* で求めた値を $11$ で割ったあまりを求める *(2)*
3. *(2)* で求めた値を $11$ から引いた値の下1桁がチェック・ディジットとなる

たとえば，住所コード `32201` のチェック・ディジットは 

{{< fig-math >}}
\[
11 - \left(\left(3 \times 6 + 2 \times 5 + 2 \times 4 + 0 \times 3 + 1 \times 2 \right) \bmod 11 \right) = 6
\]
{{< /fig-math >}}

なので `322016` が全国地方公共団体コードとなる。

また 都道府県コード＋`000`＋チェック・ディジット で都道府県を表す全国地方公共団体コードになるらしい。
つまり `320005` で島根県を表す。

### 国交省 [GIS] による大字町丁目コード

JIS 住所コードよりも更に詳細な住所コードは色々あるが，たとえば以下のものがある。

- 国土地理協会の全国町・字ファイルで提供される JIS 住所コードを含む11桁のコード体系。通称「11桁コード」。住基ネットや個人番号カードでおなじみ地方公共団体情報システム機構もこれを使っているようだ
- 運輸局で使われる運輸局住所コード。9桁または12桁のコード体系。 JIS 住所コード非互換
- 国土交通省 [GIS] の位置参照情報に含まれる大字町丁目コード。 JIS 住所コードを含む12桁のコード体系

オススメは [GIS] の位置参照情報。
[GIS] で提供されるデータはいわゆるオープンデータになっていて Creative Commons の「[表示  {{< cc-syms "cc" "by" >}}](https://creativecommons.org/licenses/by/4.0/ "Creative Commons — Attribution 4.0 International — CC BY 4.0")」ライセンス条件下で利用することができる。

{{< fig-quote type="markdown" title="利用規約" link="https://nlftp.mlit.go.jp/ksj/other/agreement.html" >}}
{{% quote %}}本利用ルールは、クリエイティブ・コモンズ・ライセンスの表示4.0国際（[https://creativecommons.org/licenses/by/4.0/legalcode.ja](https://creativecommons.org/licenses/by/4.0/legalcode.ja)に規定される著作権利用許諾条件。以下「CCBY」といいます。）と互換性があり、本利用ルールが適用されるコンテンツはCCBYに従うことでも利用することができます{{% /quote %}}。
{{< /fig-quote >}}

### [Japanese-Addresses][japanese-addresses]

[GIS] の位置参照情報を上手く使っているのが [Geolonia] から提供されているオープンデータの [japanese-addresses] である。

- [日本全国の住所マスターデータをオープンデータとして無料公開 - Geolonia](https://geolonia.com/pressrelease/2020/08/05/japanese-addresses.html)
- [geolonia/japanese-addresses: 全国の町丁目レベル（189,540件）の住所データのオープンデータ](https://github.com/geolonia/japanese-addresses)

[japanese-addresses] データは Creative Commons の「[表示  {{< cc-syms "cc" "by" >}}](https://creativecommons.org/licenses/by/4.0/ "Creative Commons — Attribution 4.0 International — CC BY 4.0")」ライセンスで提供されている。
またデータの生成コードは MIT ライセンスで公開されている。

[japanese-addresses] データは [GIS] の位置参照情報をそのまま使っているわけではなく JP の[郵便番号データ]と組み合わせて住所の読み情報も付加しているようだ。
あれ？ じゃあ [GIS] の位置参照情報と[郵便番号データ]を組み合わせて住所コードと郵便番号とを連携させられるんじゃね？ と思った私を誰が責められよう（笑）

## 住所コードと郵便番号

JP の[郵便番号データ]を眺めてみるとこんな感じになっている。

```text
32201,"690  ","6900000","ｼﾏﾈｹﾝ","ﾏﾂｴｼ","ｲｶﾆｹｲｻｲｶﾞﾅｲﾊﾞｱｲ","島根県","松江市","以下に掲載がない場合",0,0,0,0,0,0
32201,"69002","6900261","ｼﾏﾈｹﾝ","ﾏﾂｴｼ","ｱｲｶﾁｮｳ","島根県","松江市","秋鹿町",0,0,0,0,0,0
32201,"690  ","6900022","ｼﾏﾈｹﾝ","ﾏﾂｴｼ","ｱｵﾊﾞﾀﾞｲ","島根県","松江市","青葉台",0,0,0,0,0,0
32201,"690  ","6900015","ｼﾏﾈｹﾝ","ﾏﾂｴｼ","ｱｹﾞﾉｷﾞ","島根県","松江市","上乃木",0,0,1,0,0,0
...
```

一見 JIS 住所コードと郵便番号が

{{< fig-img class="lightmode" src="./entity-relationship-1.png" link="./entity-relationship-1.png" width="687" >}}

のように 1 対 n の関係になっているように見えるけど

- [郵便番号や市区町村データを取り扱うときにはまったこと - Qiita](https://qiita.com/_takwat/items/3a121656425fac7bb820)

によると

- 郵便番号は必ず1つの町名に紐づいているわけではない
- 市区町村をまたいで同じ郵便番号を持つケースがある
- 市区町村はおろか県を飛び越えて同じ郵便番号を持ちうるケースがある

そうで，つまり JIS 住所コードと郵便番号と住所（文字列）の関係は

{{< fig-img class="lightmode" src="./entity-relationship-2.png" link="./entity-relationship-2.png" width="687" >}}

となっていて， JIS 住所コードと郵便番号の間で関係を記述できない。
敢えてやるなら

{{< fig-img class="lightmode" src="./entity-relationship-3.png" link="./entity-relationship-3.png" width="956" >}}

のように第3の固有IDを作って間接的に関連付ける必要がある。

## ...というわけで諦めました

道理で住所コードと郵便番号を関連付ける実装を見かけない筈だよ。

使い方としては，まず構造のない住所（文字列）があって，単にその住所を絞り込む条件として住所コードや郵便番号が使える，というだけなのだろう。
古い閉じたシステムではよくある構成だが，外部データと関連付けようとすると素朴すぎて使えない。

たとえば JIS 住所コードより詳細な「11桁コード」や [GIS] 位置参照情報の「大字町丁目コード」を第3の固有IDとして郵便番号を関連付けることは可能かもしれないが，そのためには文字列の住所を「名寄せ」する必要がある。
しかも[郵便番号データ]って月単位で変更されるので，ほぼ無理ゲーな気がする。
実際 [japanese-addresses] は郵便番号との関連付けを行っていないわけだし。

というわけで，個人が興味本位でやるのは無理，と諦めた。

おあとがよろしいようで `m(_ _)m`

[GIS]: https://nlftp.mlit.go.jp/
[Geolonia]: https://geolonia.com/ "Geolonia - An Elastic Map Hosting - Geolonia"
[japanese-addresses]: https://github.com/geolonia/japanese-addresses "geolonia/japanese-addresses: 全国の町丁目レベル（189,540件）の住所データのオープンデータ"
[郵便番号データ]: https://www.post.japanpost.jp/zipcode/download.html "郵便番号データダウンロード - 日本郵便"
<!-- eof -->
