+++
title = "正規表現に関する戯れ言"
date = "2018-04-18T00:28:41+09:00"
description = "いや，便利なのは分かるんだけどさ，「その正規表現ホンマにいるの？」ってのがあるのよ，たまに。"
image = "/images/attention/go-logo_blue.png"
tags = [ "regular-expression", "programming", "golang" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

以下の記事を見て思いついたことを戯れ言として書いておく。
コード少な目でゴメンペコン（どっちのセクションで書こうか悩んだ）。

- [Golangの正規表現がどれぐらいおそいのかをPythonと比較してみた話 - Qiita](https://qiita.com/tj8000rpm/items/b92d7617883639a3e714)

詳しい解説は[ブックマーク]({{< relref "#bookmark" >}})にある記事を参照のこと。

## 正規表現 [regexp] パッケージは（概ね）遅い

そもそも正規表現（regular expression）は，それ自体が言語の一種と言える[^lang1]。
スクリプト言語のように言語仕様の一部として組み込まれている場合は別だが， [Go 言語]のようなコンパイル言語の場合，普通は（言語本体ではなく）ライブラリやフレームワークの一部として組み込まれる。

[^lang1]: 数学の数式だって言語の一種と考えられるのなら正規表現も言語の一種と考えてさしつかえないだろう。

[Go 言語]標準の正規表現エンジンである [regexp] パッケージは（Perl や Ruby といった言語に比べて）概ね遅い。
これは本当である。
[regexp] パッケージの設計方針について以下のような記述がある。

{{< fig-quote title="regexp - The Go Programming Language" link="https://golang.org/pkg/regexp/" lang="en" >}}
<q>The regexp implementation provided by this package is guaranteed to run in time linear in the size of the input.</q>
{{< /fig-quote >}}

まぁぶっちゃけて言うと [regexp] パッケージは「遅くなりすぎない」ように作られているのだ[^re1]。

[^re1]: 詳しくは “[Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)” を参照のこと。

### [strings] パッケージを使え

その代わりと言ってはナニだが [Go 言語]では文字列検索や操作のための [strings] や [unicode] といった標準パッケージが充実している。
例えば，単純な前方一致や後方一致なら [`strings`]`.HasPrefix()` 関数や [`strings`]`.HasSuffix()` 関数を使いやがれ，みたいな感じ。

ただ，正規表現を使うことでコードの可読性が上がる（ことによってバグ混入リスクを減らせる）ことは確かなので，状況によってパッケージを使い分けるのがコツである。

### [regexp] パッケージ利用のコツ

[regexp] パッケージを使う場合はちょっとした工夫が必要となる。

まず正規表現のコンパイルにはそこそこのコストがかかるため，できるだけ使いまわすようにする。
固定の正規表現ならグローバル変数として初期化してしまうのも手である[^re2]。

[^re2]: [`regexp`]`.MustCompile()` 関数では，コンパイルに失敗すると [panic] を投げる。通常は [`regexp`]`.Compile()` 関数を使い，返り値の error をチェックする。

```go
var re = regexp.MustCompile("^[0-9a-fA-F]+$")
```

複数の正規表現を動的に切り替えるのであれば map 等を使ってコンパイル結果をキャッシュするのもいいかもしれない。

さらに正規表現による評価を実行する際は排他制御がかかるため，並行処理下でコンパイル済みインスタンスをそのまま使うとパフォーマンスが落ちる。
そこで [`regexp`]`.Regexp.Copy()` 関数を使ってインスタンスのコピーを生成し，コピーを使って評価を行う。

```go
if !re.Copy().MatchString(s) {
    fmt.Println(s, "is not hexadecimal.")
}
```

{{% md-box %}}
【追記 2019-02-26】 [Go 1.12]({{< ref "/release/2019/02/go-1_12-is-released.md" >}}) からは [`regexp`]`.Regexp.Copy()` 関数は必要なくなった。
ブラボー！
{{% /md-box %}}

## その正規表現ホンマにいるの？

いや，便利なのは分かるんだけどさ，「その正規表現ホンマにいるの？」ってのがあるのよ，たまに。
私もやっつけのコードだと正規表現で済ますことがあるけど。

どうしても正規表現が必要で Perl や Ruby といった他言語のほうがシステム全体のパフォーマンスがいいというのであれば，それらの言語を使えばいいと思う。
もしくは正規表現パッケージを自作するか。

個人的に「言語の選択は適材適所」と思ってるので，無理に [Go 言語]にこだわる必要はないと思う。
もちろん遊びでやるのなら好きにすればいいんだけど。

## ブックマーク{#bookmark}

- [Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html)
- [regexpとの付き合い方 〜 Go言語標準の正規表現ライブラリのパフォーマンスとアルゴリズム〜 | eureka tech blog](https://developers.eure.jp/tech/golang-regexp/)
- [正規表現チェッカー（JavaScript版） | Softel labs](https://www.softel.co.jp/labs/tools/regex/)
- [40行以内で正規表現エンジンを構築 | プログラミング | POSTD](http://postd.cc/build-your-own-regex/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[panic]: http://blog.golang.org/defer-panic-and-recover "Defer, Panic, and Recover - The Go Blog"
[regexp]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"
[`regexp`]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"
[strings]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"
[`strings`]: https://golang.org/pkg/strings/ "strings - The Go Programming Language"
[unicode]: https://golang.org/pkg/unicode/ "unicode - The Go Programming Language"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253"><img src="https://images-fe.ssl-images-amazon.com/images/I/41meaSLNFfL._SL160_.jpg" width="123" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9EGo-ADDISON-WESLEY-PROFESSIONAL-COMPUTING-Donovan/dp/4621300253?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4621300253">プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES)</a></dt>
	<dd>Alan A.A. Donovan, Brian W. Kernighan</dd>
	<dd>柴田 芳樹 (翻訳)</dd>
    <dd>丸善出版 2016-06-20</dd>
    <dd>Book 単行本（ソフトカバー）</dd>
    <dd>ASIN: 4621300253, EAN: 9784621300251</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">著者のひとりは（あの「バイブル」とも呼ばれる）通称 “K&amp;R” の K のほうである。この本は Go 言語の教科書と言ってもいいだろう。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2018-10-20">2018-10-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
