+++
title = "おそるべき絵文字"
date =  "2021-03-30T19:58:27+09:00"
description = "なんでこんなカオスになっちゃったのかねぇ。"
image = "/images/attention/kitten.jpg"
tags = [ "character", "unicode", "engineering", "emoji" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

きっかけは以下の tweet:

{{< fig-gen class="nobox" >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">ほら <a href="https://twitter.com/hashtag/golang?src=hash&amp;ref_src=twsrc%5Etfw">#golang</a> ならこんなに簡単に絵文字を取り除けますよ！8ビット単位だとややこしい書き方が必要になっちゃう。<a href="https://t.co/AiELa8RlvS">https://t.co/AiELa8RlvS</a></p>&mdash; のぼのぼ📡 (@nobonobo) <a href="https://twitter.com/nobonobo/status/1376186097881866240?ref_src=twsrc%5Etfw">March 28, 2021</a></blockquote>
{{< /fig-gen >}}

この tweet の元ネタである「[Go言語の正規表現regexpが遅すぎる](https://twitter.com/c_shii41/status/1376163845647986688)」というのは概ねその通りで，私も[記事で書いた]({{< ref "/golang/regular-expression.md" >}} "正規表現に関する戯れ言") が，これはそもそも [`regexp`] パッケージの設計方針が

{{< fig-quote type="markdown" title="regexp - The Go Programming Language" link="https://golang.org/pkg/regexp/" lang="en" >}}
{{% quote %}}The regexp implementation provided by this package is guaranteed to run in time linear in the size of the input{{% /quote %}}.
{{< /fig-quote >}}

となっているためで，「遅くなりすぎない」ことのトレードオフとして（Ruby や Perl などと比べて）全体的に遅いのである[^regexp1]。

[^regexp1]: [`regexp`] パッケージの設計については {{% quote %}}[Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html){{% /quote %}} を参照のこと。

ただ，[サンプルコード](https://play.golang.org/p/brhrHn6dm3A)にある

```go
var emojiRegex = regexp.MustCompile("[\U0001f000-\U0001ffff]")
```

という正規表現はさすがに端折りすぎで，つい「[そんな単純じゃねーよ](https://twitter.com/spiegel_2007/status/1376301767411990529)」と脊髄反射してしまったのは許してほしい。

でも，よく考えたら，ここのブログや [Zenn] でも絵文字の話をよく書くようになったが，まとまった記事としては書いてない。
というわけで，この記事である程度まとめてみたいと思う。

以上，前説おわり。

## Unicode は1コード＝1文字じゃない

まず大前提として Unicode は1コード＝1文字じゃない。
これは日本語圏で苦労している私たちには自明だと思うが，たとえば「ペンギン」という単語にしたって

- ペ：`U+30D8` + `U+309A`
- ン：`U+30F3`
- ギ：`U+30AD` + `U+3099`
- ン：`U+30F3`

と[濁点・半濁点が分離している場合がある]({{< ref "/golang/unicode-normalization.md" >}} "Go 言語と Unicode 正規化") 。
これは絵文字でも同様で，たとえば「土下座する男性 {{< emoji ":bowing_man:" >}}」を拙作の [gnkf] を使って Unicode 符号点（code point）に分解すると

```text
$ echo 🙇‍♂️ | gnkf dump --unicode
0x0001f647, 0x200d, 0x2642, 0xfe0f, 0x000a
```

と，4つの符号点の合成列として構成されているのが分かる。
ちなみに各符号点の内容は

| 符号点    |              字形               | 名称                  |
| --------- | :-----------------------------: | --------------------- |
| `U+1F647` | {{< emoji ":person_bowing:" >}} | PERSON BOWING DEEPLY  |
| `U+200D`  |                                 | ZERO WIDTH JOINER     |
| `U+2642`  |            &#x2642;             | MALE SIGN             |
| `U+FE0F`  |                                 | VARIATION SELECTOR-16 |

となっている。
これだけ見ても，絵文字を

```go
var emojiRegex = regexp.MustCompile("[\U0001f000-\U0001ffff]")
```

などと安直に括り出すのが如何にヤバいか分かるであろう。

Unicode 文字列中の「文字」を真面目に分解して取り扱いたいなら，たとえば [Go] 言語なら `rune` ではなく `[]rune` のように符号点の可変配列として取り扱う必要があるだろう。
かつ，各符号点の意味を考えながら「どこまでが1文字か」を解析するロジックが必要になる。

ちなみに [Go] では [rivo/uniseg] パッケージを使うといい感じに分解してくれるらしい。

## 符号とフォントと入出力

これも言ってしまえば当たり前のことなのだが，絵文字は符号とフォントと入出力が揃ってはじめて絵文字として機能する。
先ほどの {{< emoji ":bowing_man:" >}} も

```text
$ echo 🙇‍♂️ | gnkf dump --unicode
0x0001f647, 0x200d, 0x2642, 0xfe0f, 0x000a
```

という符号点の並びがひとつの絵文字であると解釈し，かつその文字に対応するフォントとグリフがあり，さらにそのグリフで表示できるアプリケーションがあってはじめて {{< emoji ":bowing_man:" >}} と表示できる。

この記事中の絵文字も，もしかしたらブラウザの種類やバージョン等によって 🙇♂ と2文字に見えているかもしれない。
もしくは変な記号に「文字化け」するとか，そもそも全く表示されないとか...

たとえば Twitter ではユーザ側の環境やアプリケーション間の差異を軽減するため，絵文字を独自の画像データに置き換えるツールキットを公開している。

- [Twemoji](https://twemoji.twitter.com/)

[Zenn] のアテンション用絵文字もこれを使って表示している筈。

それぞれの絵文字がどのように表示されるか（または表示されないか）については Unicode 公式ページにある

- [Full Emoji List]

が参考になるだろう。
2021年3月時点の最新版である v13.1 では1,816個の絵文字が定義されている。

多いよ `orz`

### ターミナル・エミュレータと絵文字

ブラウザとかスマホ・アプリならまだマシだと思うが， CUI ベースのターミナル・エミュレータは絵文字周りが特に冷遇されている印象がある。

私は CLI ツールで対話モードを組むときに [nyaosorg/go-readline-ny] パッケージ[^ny1] のお世話になるのだが，[リリース情報](https://github.com/nyaosorg/go-readline-ny/releases "Releases · nyaosorg/go-readline-ny")を見るとホンマに絵文字で苦労されているのが分かる。
その一端が以下の [Zenn] 本に記されているので，是非ご覧あれ。

- [Windows と Unicode とボク](https://zenn.dev/zetamatta/books/b820d588f4856bcf836c)

プログラマが読んだら涙で前が見えなくなるよ（笑）

[^ny1]: [nyaosorg/go-readline-ny] は同じ作者による [NYAGOS] からのスピンオフで，ターミナル・エミュレータからの入力制御に特化している。Emacs 風のキー・バインドでヒストリ機能を付けることもできる。元々は Windows 用だと思うが Ubuntu のターミナル・エミュレータでも問題なく機能するので重宝している。

## “#” の Keycap はあるのに “A” の Keycap はないのか

上で紹介した “[Full Emoji List]” を眺めると，たとえば絵文字 {{< emoji ":keycap_#:" >}} は

| 符号点   | 字形  | 名称                       |
| -------- | :---: | -------------------------- |
| `U+0023` |  `#`  | NUMBER SIGN                |
| `U+FE0F` |       | VARIATION SELECTOR-16      |
| `U+20E3` |  ` ⃣`  | COMBINING ENCLOSING KEYCAP |

という並びになっている。
もう

```go
var emojiRegex = regexp.MustCompile("[\U0001f000-\U0001ffff]")
```

という正規表現に微塵もかからない（笑）

いや，そうじゃなくて。
たしかに普通の半角記号である `#` が合成列の先頭に来ているのも驚きだが，その直後に絵文字用の異体字セレクタが来ているのにビックリした。
必ずしも異体字セレクタが最後に付くんじゃないのか。

まぁ，でも意味は分かる。
{{< emoji ":keycap_#:" >}} は `#` 記号の絵文字異体字ってことで，先頭の符号点だけ見れば `#` と等価な文字として扱えるってことだよね。
そしてその後ろに keycap を表す結合文字がくっ付いている，と。

そんじゃあ `#` の代わりに `A` とかの任意の文字でも keycap の絵文字が作れるんじゃね？ と一瞬思ったが “[Full Emoji List]” を見る限りそんなことはないようだ[^blloda1]。

[^blloda1]: 処理系によっては任意のコードを組み合わせて勝手に絵文字を作れたりするのかもしれないが，互換性がなくなるので「情報交換用」としては使えない。ちなみに {{< emoji ":A_button_(blood_type):" >}} という絵文字はあるが，これは血液型（A型）を表す絵文字らしい。

どういうルールなんだろうねぇ。

## 肌色の異体字

最初に紹介した “PERSON BOWING DEEPLY” {{< emoji ":person_bowing:" >}} には肌色の情報をくっ付けることができる。

こんな感じ。

| 合成列                |                 字形                  | 名称                                  |
| --------------------- | :-----------------------------------: | ------------------------------------- |
| `U+1F647` + `U+1F3FB` | {{< emoji ":person_bowing_tone1:" >}} | person bowing: light skin tone        |
| `U+1F647` + `U+1F3FC` | {{< emoji ":person_bowing_tone2:" >}} | person bowing: medium-light skin tone |
| `U+1F647` + `U+1F3FD` | {{< emoji ":person_bowing_tone3:" >}} | person bowing: medium skin tone       |
| `U+1F647` + `U+1F3FE` | {{< emoji ":person_bowing_tone4:" >}} | person bowing: medium-dark skin tone  |
| `U+1F647` + `U+1F3FF` | {{< emoji ":person_bowing_tone5:" >}} | person bowing: dark skin tone         |

この肌色って人種を指すものじゃなくて “[fitzpatrick skin typing](https://en.wikipedia.org/wiki/Fitzpatrick_scale "Fitzpatrick scale - Wikipedia")” と呼ばれる紫外線への感受性を基にした分類らしい。

| 符号点    | 名称                                |
| --------- | ----------------------------------- |
| `U+1F3FB` | EMOJI MODIFIER FITZPATRICK TYPE-1-2 |
| `U+1F3FC` | EMOJI MODIFIER FITZPATRICK TYPE-3   |
| `U+1F3FD` | EMOJI MODIFIER FITZPATRICK TYPE-4   |
| `U+1F3FE` | EMOJI MODIFIER FITZPATRICK TYPE-5   |
| `U+1F3FF` | EMOJI MODIFIER FITZPATRICK TYPE-6   |

このコードは結合文字の一種のように見えるが，どんな絵文字にもくっ付くわけではないらしい。
有効な組み合わせは

- [Full Emoji Modifier Sequences](https://www.unicode.org/emoji/charts/full-emoji-modifiers.html)

を参考にするといいだろう。

## ZWJ による絵文字の合成

最初の「土下座する男性 {{< emoji ":bowing_man:" >}}」に戻る。
男性があるのだから女性バージョン {{< emoji ":bowing_woman:" >}} も当然ある。

{{< emoji ":bowing_woman:" >}} の内容は以下の通り。

| 符号点    |              字形               | 名称                  |
| --------- | :-----------------------------: | --------------------- |
| `U+1F647` | {{< emoji ":person_bowing:" >}} | PERSON BOWING DEEPLY  |
| `U+200D`  |                                 | ZERO WIDTH JOINER     |
| `U+2640`  |            &#x2640;             | FEMALE SIGN           |
| `U+FE0F`  |                                 | VARIATION SELECTOR-16 |

つまり “PERSON BOWING DEEPLY” に性差を示す “MALE SIGN” または “FEMALE SIGN” を “ZERO WIDTH JOINER” (ZWJ) を介してくっ付けることで異体字を構成している。

このように ZWJ を使って文字をくっ付けて異体字や新たな絵文字を作る方法にはかなりのバリエーションがあるようだが，どうもプラットフォームごとに勝手にコードを組み合わせているらしく上手く表示できないパターンがあるようだ。

このため [Zenn] では，以下のリストにない絵文字を排除することで[対応](https://github.com/zenn-dev/zenn-roadmap/issues/224#issuecomment-767231796)しているそうだ。

- [Recommended Emoji ZWJ Sequences](https://www.unicode.org/emoji/charts/emoji-zwj-sequences.html)

この一覧によると “PERSON BOWING DEEPLY” {{< emoji ":person_bowing:" >}} には，前節の肌色バリエーションと男女性差のバリエーションを組み合わせて，合計10個の異体字がある，ということになる。

## 国旗絵文字

個人的にこれが一番ワケワカメ。

REGIONAL INDICATOR なるコードがあって

| 符号点    |   字形    | 名称                               |
| --------- | :-------: | ---------------------------------- |
| `U+1F1E6` | &#x1F1E6; | REGIONAL INDICATOR SYMBOL LETTER A |
| `U+1F1E7` | &#x1F1E7; | REGIONAL INDICATOR SYMBOL LETTER B |
| `U+1F1E8` | &#x1F1E8; | REGIONAL INDICATOR SYMBOL LETTER C |
| ...       |    ...    | ...                                |
| `U+1F1EF` | &#x1F1EF; | REGIONAL INDICATOR SYMBOL LETTER J |
| ...       |    ...    | ...                                |
| `U+1F1F5` | &#x1F1F5; | REGIONAL INDICATOR SYMBOL LETTER P |
| ...       |    ...    | ...                                |
| `U+1F1FF` | &#x1F1FF; | REGIONAL INDICATOR SYMBOL LETTER Z |

これを使って ISO 3166 (日本では JIS X 0304) の国別コードを組むと国旗の絵文字になるらしい。
日本の国別コードは JP なので

| 合成列                |            字形           | 名称        |
| --------------------- | :-----------------------: | ----------- |
| `U+1F1EF` + `U+1F1F5` | {{< emoji ":flag_jp:" >}} | flag: Japan |

となる。

ところでこれ，ちゃんと国旗に見えてます？ かなり環境依存度が大きいようで Windows とかでは軒並みアウトっぽい。

さらにアレなのが Unicode 10.0 で追加された [subdivision-flag](https://emojipedia.org/emoji-tag-sequence/ "🏴󠁧󠁢󠁳󠁣󠁴󠁿 Emoji Tag Sequences for Subdivision Flags") に分類されているもので，たとえば「イングランドの国旗 {{< emoji ":flag_England:" >}}」は

| 符号点     |   字形    | 名称                     |
| ---------- | :-------: | ------------------------ |
| `U+x1F3F4` | &#x1F3F4; | WAVING BLACK FLAG        |
| `U+E0067`  |           | TAG LATIN SMALL LETTER G |
| `U+E0062`  |           | TAG LATIN SMALL LETTER B |
| `U+E0065`  |           | TAG LATIN SMALL LETTER E |
| `U+E006E`  |           | TAG LATIN SMALL LETTER N |
| `U+E0067`  |           | TAG LATIN SMALL LETTER G |
| `U+E007F`  |           | CANCEL TAG               |

という並びになっている。
一般的には “WAVING BLACK FLAG” と “CANCEL TAG” の間に TAG LATIN SMALL LETTER A～Z (`U+E0061` ～ `U+E007A`) を5つ並べる構成になっているらしい。

なにその面倒くさいやつ。

## 絵文字シーケンスのまとめ

以上をまとめると，絵文字の分類は以下のようになっているらしい（「[Unicode 絵文字にまつわるあれこれ](https://qiita.com/_sobataro/items/47989ee4b573e0c2adfc "Unicode 絵文字にまつわるあれこれ (絵文字の標準とプログラム上でのハンドリング) - Qiita")」を参考に分類）。

- singleton ： 単体のコードポイントからなる絵文字
- emoji sequence ： 複数のコードポイントからなる絵文字
  - emoji core sequence ： 通常の絵文字
    - emoji combining sequence ： 囲み文字
    - emoji modifier sequence ： skin tone 絵文字
    - emoji flag sequence ： 国旗絵文字
  - emoji zwj sequence ： 家族絵文字、職業絵文字など
  - emoji tag sequence ：タグ絵文字

この分類でいくと Keycap の絵文字は “emoji combining sequence” に相当する。
肌色の異体字は “emoji modifier sequence” やね。
ZWJ にとる文字の結合は “emoji zwj sequence” で，国旗絵文字はそのまま “emoji flag sequence”。
ただし subdivision-flag は “emoji tag sequence” に分類されるようだ。

しかもこれってただの分類なので，実際の絵文字がどれに分類されるかは「知識」としてあらかじめ知ってないといけない。
プログラマ的に言うならロジックを書くだけでは足りなくてロジックを駆動する知識をデータベースとして何処かに保持っておく必要がある。

なんでこんなカオスになっちゃったのかねぇ。
本当に「おそるべき絵文字」だよ。

## 【余談】元号記号も絵文字と見做せばいいぢゃん

絵文字を調べてて思い出したのが，4年前に書いた

- [新元号「文字」という技術的負債]({{< ref "/remark/2017/12/character-of-the-new-era-name.md" >}})

という記事。
「[技術的負債]({{< ref "/remark/2020/12/technical-debt-and-hacker.md" >}} "技術的負債とハッカー")」と口走ったのは私の黒歴史としてスルーしていただけるとありがたいが，よく考えたら「㍻ `U+337B`」とかを「文字」だと思うから不合理に感じるのであって「絵文字」と思えばアリなのか。
いや， Unicode では絵文字と定義しているわけじゃないけど。

でも「{{< emoji ":copyright:" >}} `U+00A9`」だって絵文字扱いなんだから ㍻ も絵文字でいいよね（笑）

## ブックマーク

Unicode における「絵文字」黎明期の話は，[小形克宏](https://ogwata.hatenadiary.org/ "もじのなまえ")さんの一連の記事が参考になる。

- [絵文字が開いてしまった「パンドラの箱」第1回--日本の携帯電話キャリアが選んだ道 - CNET Japan](https://japan.cnet.com/article/20389042/)
- [絵文字が開いてしまった「パンドラの箱」第2回--Googleの開けてしまった箱の中味 - CNET Japan](https://japan.cnet.com/article/20389453/)
- [絵文字が開いてしまった「パンドラの箱」第3回--Unicode提案の限界とメリット - CNET Japan](https://japan.cnet.com/article/20390204/)
- [絵文字が開いてしまった「パンドラの箱」第4回--絵文字が引き起こしたUnicode-MLの“祭り” - CNET Japan](https://japan.cnet.com/article/20394318/)
- [絵文字が開いてしまった「パンドラの箱」第5回--絵文字と日本マンガの親密な関係 - CNET Japan](https://japan.cnet.com/article/20398174/)
- [絵文字が開いてしまった「パンドラの箱」第6回--Google・Apple提案とそのシナリオ - CNET Japan](https://japan.cnet.com/article/20407674/)
- [絵文字が開いてしまった「パンドラの箱」第7回--そして舞台はダブリンから東京へ - CNET Japan](https://japan.cnet.com/article/20407951/)
- [グーグルが絵文字を世界標準に提案した理由--国際化エンジニアに聞くプロジェクトの舞台裏（前編） - CNET Japan](https://japan.cnet.com/article/20409186/)
- [グーグルが絵文字を世界標準に提案した理由--国際化エンジニアに聞くプロジェクトの舞台裏（前編） - CNET Japan](https://japan.cnet.com/article/20409186/)
- [「絵文字に平等をサポートしてください」人種差別の指摘にゆれるUnicode - INTERNET Watch Watch](https://internet.watch.impress.co.jp/docs/special/670150.html)
- [どんな人々がUnicodeの絵文字に「民族の多様性」を求めているのか？ - INTERNET Watch Watch](https://internet.watch.impress.co.jp/docs/special/678029.html)
- [絵文字を「符号」として処理する難しさ～日本のモバイルウェブのカオスぶり - INTERNET Watch Watch](https://internet.watch.impress.co.jp/docs/special/379269.html)
- [絵文字を「語」として処理する難しさ～定義通りとは限らない、絵文字の意味 - INTERNET Watch Watch](https://internet.watch.impress.co.jp/docs/special/380149.html)
- [絵文字を「語」として処理する難しさ～「ビール」と「飲み会」見分ける技術 - INTERNET Watch Watch](https://internet.watch.impress.co.jp/docs/special/380180.html)

その他のブックマーク：

- [Let's EMOJI｜絵文字一覧と絵文字検索🎉](https://lets-emoji.com/)
- [絵文字と異体字と Markdown]({{< ref "/remark/2020/10/emoji-variation-and-markdown.md" >}})
- [Unicode 文字種の判別]({{< ref "/golang/unicode-rangetables.md" >}})
- [やっかいな日本語](https://zenn.dev/spiegel/articles/20210118-characters)

[Go]: https://go.dev/
[`regexp`]: https://golang.org/pkg/regexp/ "regexp - The Go Programming Language"
[gnkf]: https://github.com/spiegel-im-spiegel/gnkf "spiegel-im-spiegel/gnkf: Network Kanji Filter by Golang"
[rivo/uniseg]: https://github.com/rivo/uniseg "rivo/uniseg: Unicode Text Segmentation for Go (or: How to Count Characters in a String)"
[nyaosorg/go-readline-ny]: https://github.com/nyaosorg/go-readline-ny "nyaosorg/go-readline-ny: Readline library for golang , used in nyagos"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
[Zenn]: https://zenn.dev/ "Zenn｜エンジニアのための情報共有コミュニティ"
[Full Emoji List]: https://unicode.org/emoji/charts/full-emoji-list.html
