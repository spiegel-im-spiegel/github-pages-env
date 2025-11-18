+++
title = "ダッシュとかチルダとか"
date =  "2025-11-14T11:58:22+09:00"
description = "文字コードはホンマ面倒くさい"
image = "/images/attention/kitten.jpg"
tags = [ "character", "unicode", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## エンダッシュとエムダッシュを直接入力できる？

Bluesky の TL に面白い記事が流れてたので，今回はこれを起点にしてみる。

- [【朗報？】Windows、「—」と「–」が入力しやすくなる。英文を打つ人にとっては最高のアプデかも | ライフハッカー・ジャパン](https://www.lifehacker.jp/article/2511-microsoft-added-shortcut-for-em-dashes/)

{{< fig-quote type="markdown" title="【朗報？】Windows、「—」と「–」が入力しやすくなる。英文を打つ人にとっては最高のアプデかも" link="https://www.lifehacker.jp/article/2511-microsoft-added-shortcut-for-em-dashes/" >}}
- エンダッシュ（–）を入力したい時：Windowsキーを押しながら、ハイフン（-）を押す
- エムダッシュ（—）を入力したい時：Windowsキー ＋ Shiftキーを押しながら、ハイフン（-）を押す
{{< /fig-quote >}}

「いや，これ， US キーボードの話じゃろ？」と思って，試しに自宅の [Windows 機]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした")で試したら

```text
$ echo -–— | gnkf dump -u
0x002d, 0x2013, 0x2014, 0x000d, 0x000a
```

ホンマじゃ！ ちゃんと日本語キーボードで `–` (EN DASH, U+2013) と `—` (EM DASH, U+2014) が出力された（US キーボードでどうなるかは知らない）。
試しに Ubuntu 機で同じようにやってみたが駄目だった。

エンダッシュやエムダッシュを使うことってあんまりないんだよな。
仕事で書く文書ではほぼ使わない。
TeX や [Typst](https://typst.app/docs/reference/symbols/ "Symbols – Typst Documentation") では `--` や `---` でエンダッシュやエムダッシュに変換されるので，困った覚えはない。
ここのブログでは実体参照の `&ndash;` や `&mdash;` を使うので，これも困ることはないな。

というわけで「へぇ」とは思ったけど，私は使う機会もなく忘れてしまいそう（笑）

ちなみに，よく使う `-` (U+002D) は “HYPHEN-MINUS” でダッシュ記号ではない。
さらに “HYPHEN” は `‐` (U+2010) と定義されているし “MINUS SIGN” は `−` (U+2212) と定義されている。
さらにさらに言うと全角のハイフンマイナスである `－` (U+FF0D) もある。
でも見た目で殆ど違いが分からんし，ハイフンマイナス以外はどうやって入力するのかも分からない。
どうも生成 AI は使い分けてるみたいなんだが，見た目がみんな同じなので，これはこれで[困る](#h-ai)んだよなぁ。

## 全角ダッシュの Unicode マッピング

日本語の文字コード（JIS X 0208 および JIS X 0213）にはいわゆる「全角ダッシュ」があるのだが Unicode では上述の “EM DASH” (U+2014) にマッピングされているそうな。
まぁ，これは真っ当だろう。

しかし一時期，この全角ダッシュを U+2015 (HORIZONTAL BAR) にマッピングしていた時期があったらしい。
これは後に訂正されたようだが，古い処理系だとシフトJISの全角ダッシュ（8160H）を U+2015 に変換したりする場合もあるんだとか。

あと和文の倍角ダッシュに U+2015 を並べて使う場合もあるみたいで，簡単な話ではないっぽい。
たとえば奥村晴彦さんが書かれた『[LaTeX美文書作成入門](https://www.amazon.co.jp/dp/4297138891?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "［改訂第9版］LaTeX美文書作成入門 | 奥村 晴彦, 黒木 裕介 |本 | 通販 | Amazon")』では，倍角ダッシュを U+2015 を使って

{{< fig-quote class="nobox" type="markdown" title="『［改訂第9版］LaTeX美文書作成入門』4.8節" link="https://www.amazon.co.jp/dp/4297138891?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```tex
\def\――{―\kern-.5zw―\kern-.5zw―}
```
{{< /fig-quote >}}

というコマンドを定義すれば文字と文字の間の隙間を潰せると紹介している。
U+2015 の文字を全角2文字分の幅で3つ重ねて繋がってるように見せるのね。
なるほど。

## 波ダッシュと全角チルダ

似たような話に和文の「波ダッシュ」と「全角チルダ」がある。

| 字形 | コードポイント | 名前 |
| :---: | --- | --- |
| 〜 | `U+301C` | WAVE DASH |
| ～ | `U+FF5E` | FULLWIDTH TILDE |

このように見た目はそっくりなのだが Unicode 7.0 までは波ダッシュの字形が異なっていたらしい。
この辺は以下が参考になるだろうか。

- [どうしてこうなった!? 波ダッシュをめぐる考察｜『人文×社会』の中の人](https://note.com/jinbunxshakai/n/n7466cd68db43)

字形に関しては Unicode 8.0 で修正されているのだが， Windows の IME では今だに「から」を変換すると全角チルダのほうが候補に上がるため，波ダッシュであるべきところが全角チルダになってしまうことがあるようだ。
見た目いっしょだし，分からんよね（笑）

文字コードはホンマ面倒くさい。

## 【2025-11-18 追記】 エムダッシュとアポストロフィと AI {#a-ai}

最近また面白い記事を見かけた[^t1]。

[^t1]: [この記事](https://www.itmedia.co.jp/aiplus/articles/2511/16/news017.html "ChatGPTにエムダッシュを使わないよう指示できるとアルトマンCEO【訂正あり】 - ITmedia AI＋")は最初エムダッシュを「日本語では**で表示している」とか意味不明なことを書いていて混乱したが，後に訂正されて意味が通った。 ITmedia はたまにこういうのがあるよな（笑）

- [ChatGPTにエムダッシュを使わないよう指示できるとアルトマンCEO【訂正あり】 - ITmedia AI＋](https://www.itmedia.co.jp/aiplus/articles/2511/16/news017.html)

簡単に言うと ChatGPT にはエムダッシュ（U+2014）を頻繁に使う癖があるそうで ChatGPT-5.1 の有料プランでは（カスタム指示として）これを抑制するよう命令できるらしい。
なんというか果てしなく間抜けな話だが，これを以って “[Small-but-happy win](https://x.com/sama/status/1989193813043069219)” とか言ってるんだから，昨今の AI ブームが如何に狂ってるか分かる（笑）

これに絡むかどうか分からないが Mastodon の TL で

{{< fig-quote class="nobox" >}}
<iframe src="https://mstdn.jp/@zetamatta/115566226345217152/embed" class="mastodon-embed" style="max-width: 100%; border: 0" width="400" allowfullscreen="allowfullscreen"></iframe><script src="https://mstdn.jp/embed.js" async="async"></script>
{{< /fig-quote >}}

という[ポスト](https://mstdn.jp/@zetamatta/115566226345217152)があって思わず脊髄反射してしまったのだが，（実際のところは分からないが）いかにもありそうな話ではある。

ちなみに Unicode では

| 字形 | コードポイント | 名前 |
| :---: | --- | --- |
| `'` | `U+0027` | APOSTROPHE |
| ’ | `U+2019` | RIGHT SINGLE QUOTATION MARK |

と定義されていて，この意味に従うならアポストロフィとして U+2019 を使うのは間違いなんじゃね？ と思ったのだが，話はそう簡単ではないようだ。

{{< fig-quote type="markdown" title="“Chapter 6 – Unicode 17.0.0” 6.2.7 Apostrophes" link="https://www.unicode.org/versions/Unicode17.0.0/core-spec/chapter-6/#G12411" lang="en" >}}
U+0027 `'` APOSTROPHE is the most commonly used character for apostrophe. For historical reasons, U+0027 is a particularly overloaded character. In ASCII, it is used to represent a punctuation mark (such as right single quotation mark, left single quotation mark, apostrophe punctuation, vertical line, or prime) or a modifier letter (such as apostrophe modifier or acute accent). Punctuation marks generally break words; modifier letters generally are considered part of a word.

When text is set, U+2019 ’ RIGHT SINGLE QUOTATION MARK is preferred as apostrophe, but only U+0027 is present on most keyboards. Software commonly offers a facility for automatically converting the U+0027 `'` APOSTROPHE to a contextually selected curly quotation glyph. In these systems, a U+0027 in the data stream is always represented as a straight vertical line and can never represent a curly apostrophe or a right quotation mark.
{{< /fig-quote >}}

つまり U+0027 のほうは多様な意味に使われていてワケ分からんし字形としても適切ではないので U+2019 のほうをアポストロフィとして使うのが望ましいが，キーボード入力では U+0027 しか使えないので，ソフトウェアで適切に変換してね，ということらしい[^h1]。

[^h1]: このブログのジェネレータである [Hugo] は，たとえば `'Hoge'` と入力すれば 'Hoge' のように U+2018 および U+2019 の左右シングルクォーテーションに変換してくれるし， `can't` みたいなのも can't と変換してくれる。ありがたや。

いやいやいや...

さらに

{{< fig-quote type="markdown" title="“Chapter 6 – Unicode 17.0.0” 6.2.7 Apostrophes" link="https://www.unicode.org/versions/Unicode17.0.0/core-spec/chapter-6/#G12411" lang="en" >}}
An implementation cannot assume that users’ text always adheres to the distinction between these characters. The text may come from different sources, including mapping from other character sets that do not make this distinction between the letter apostrophe and the punctuation apostrophe/right single quotation mark. In that case, all of them will generally be represented by U+2019.

The semantics of U+2019 are therefore context dependent. For example, if surrounded by letters or digits on both sides, it behaves as an in-text punctuation character and does not separate words or lines.
{{< /fig-quote >}}

とか言ってくさる。
つまり，私達ユーザが文字の意味をわかって使い分けてるとは限らないので，とりあえず U+2019 を使って，あとは文脈で解釈しろ，ということのようだ。

いやいやいや...

さらにさらに，文字装飾用アポストロフィとして U+02BC (MODIFIER LETTER APOSTROPHE) もあるそうで，これはヘブライ語などの翻字に使ったり国際音声記号として使ったりするらしい。
他にも U+055A (ARMENIAN APOSTROPHE) とか U+07F4 (NKO HIGH TONE APOSTROPHE) とかいった文字もあるそうで，日本語だと全角文字 U+FF07 (FULLWIDTH APOSTROPHE) も別に定義されている。
ホンマ，勘弁してくだい `orz`

それはともかく，生成 AI がアポストロフィやシングルクォーテーションに U+2019 を律儀に使ってる（可能性が高い）のに対し，人間様はキーボード入力の U+0027 をそのまま使ってる（可能性が高い）わけで，文字コード的に真っ当な文章のほうが AI 由来だとして迫害される（かもしれない）のは，なかなかに狂っている。

## 【2025-11-18 追記】 ハイフンと AI {#h-ai}

これ，ハイフン記号もなんだよなぁ。

[PA-API に関する記事]({{< relref "./paapi-access-disabled.md" >}} "Amazon PA-API にアクセスできなくなった")を書いたときに GitHub Copilot のチャットでタイトルを添削してもらったのだが，そのときに提示された “PA‑API” のハイフンは U+2011 (NON-BREAKING HYPHEN) だったのよ。

いや，要らんけぇ，そんなの。
文字コード的には真っ当かもしれんけど，あとで grep で検索するときに U+002D じゃなかったら検索条件が面倒くさくなるので困るじゃろうが。

正しいけど人間に優しくないってのは，いかにも機械っぽいけどさ。

## ブックマーク

- [Unicodeのアポストロフィとシングル引用符 - 帰ってきた💫Unicode刑事〔デカ〕リターンズ](https://moji-memo.hatenablog.jp/entry/20081128/1227866254)
- [アポストロフィの悩み | Okumura's Blog](https://okumuralab.org/~okumura/blog/node/2480)
- [アポストロフィって何だっけ？ 文字コードをめぐる謎｜『人文×社会』の中の人](https://note.com/jinbunxshakai/n/n9fd19668797d)
- [Quotation mark - Wikipedia](https://en.wikipedia.org/wiki/Quotation_mark) : 国ごとの引用符の違いが載っていて面白い


[Hugo]: https://gohugo.io/ "Hugo - The world’s fastest framework for building websites"

## 参考

{{% review-paapi "4297138891" %}} <!-- ［改訂第9版］LaTeX美文書作成入門 -->
