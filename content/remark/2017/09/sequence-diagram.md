+++
title = "そうだ，シーケンス図を描こう！ （一応クラス図も描けるよ）"
date =  "2017-09-21T16:54:32+09:00"
update =  "2017-09-21T20:41:52+09:00"
description = "久しぶりに本ブログをいじくって mermaid でシーケンス図を描けるようにしてみる。"
tags        = [ "tools", "mermaid", "uml", "hugo", "shortcodes", "web", "site" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  highlightjs = true
  mathjax = false
  mermaidjs = true
+++

久しぶりに[本ブログ]をいじくって [mermaid] でシーケンス図を描けるようにしてみる。

## シーケンス図とは

一応説明しておくと，シーケンス図というのは複数の「もの」の間でのやり取りを時系列で表現したものである。
「もの」は相互作用するものなら何でもよく，以下のような会話も表現できる。

{{< fig-mermaid title="カバとカバン" >}}
sequenceDiagram
    カバ->>+カバン: あなた，泳げまして？
    カバン-->>-カバ: いえ
    カバ->>+カバン: 空は飛べるんですの？
    カバン-->>-カバ: いえ
    カバ->>+カバン: じゃあ，足が速いとか？
    カバン-->>-カバ: いえ
    カバ->>カバン: あなた，何にもできないのねぇ
    loop ひとりヘコむ
        カバン->>カバン: ううっ
    end
{{< /fig-mermaid >}}

縦のラインを「ライフライン[^ll]」と呼び（上から下に時間が流れる），横向きの矢印を「メッセージ」と呼ぶ。
“loop” で囲まれている部分は「複合フラグメント」と呼ばれるもので，シーケンス内の処理のかたまりを指す。
“loop” は文字通り繰り返し処理のこと。
カバンはカバに「何もできない」と言われて悶々としてしまったわけですね（笑）

[^ll]: UML 的にはライフラインは破線じゃないといけないのだけど，どういうわけか [mermaid] では実線になっている。

シーケンス図はオブジェクト指向設計ではとても重宝されている。
クラス間の相互作用を記述するだけじゃなくて，要件定義で「登場人物（人間とは限らない）」の関係を記述するのにも使われる（もちろん UML で使われる図はシーケンス図だけじゃないけどね）。

ただ，これを手描きで作図するのは結構面倒くさい（図を描きながら試行錯誤するし）。
企業向けなら [Astah*](http://astah.change-vision.com/ "astah システム設計、ソフトウェア開発支援ツール | Astah") のような開発支援ツールを使う手もあるけれど，個人では手にあまるものだし，どうせ描くならテキスト入力でサクッとやりたいよね。 

## Mermaid 記法と  mermaid.js

そうした需要に応えてくれるのが [mermaid] である。

- [mermaid]
- [knsv/mermaid: Generation of diagram and flowchart from text in a similar manner as markdown](https://github.com/knsv/mermaid) ： こちらのドキュメントのほうが詳しい

たとえば，先程の「カバとカバン」のやり取りは  [mermaid] 記法で以下のように表せる。
直感的で分かりやすいでしょ[^sd1]。

[^sd1]: 複合フラグメントは loop 以外には alt, opt, par のみサポートされているようだ。個人的には ref と  critical （par があるなら critical は必要）くらいは入れて欲しいところであるが...

```text
sequenceDiagram
    カバ->>+カバン: あなた、泳げまして？
    カバン-->>-カバ: いえ
    カバ->>+カバン: 空は飛べるんですの？
    カバン-->>-カバ: いえ
    カバ->>+カバン: じゃあ、足が速いとか？
    カバン-->>-カバ: いえ
    カバ->>カバン: あなた、何にもできないのねぇ
    loop ひとりヘコむ
        カバン->>カバン: ううっ
    end
```

この記述をもとに Web ページ上で描画を行うには mermaid.js を使う。

###  mermaid.js の設定

まず mermaid.js のセットアップは以下の通り。

```html
<script src="https://unpkg.com/mermaid/dist/mermaid.min.js"></script>
<script>
  mermaid.initialize({startOnLoad: true, theme: 'neutral'});
</script>
```

この記述を HTML の末尾（`</body>` タグの直前）に記述する。

mermaid.js は [CDN (Content Delivery Network) が用意されている](https://unpkg.com/mermaid/)ので利用することをお勧めする。
バージョンを指定する場合は

- https://unpkg.com/mermaid@7.1.0/dist/mermaid.min.js

のように記述する。

`mermaid.initialize()` 関数は初期化と描画を行う。
ページロード時に描画を行う場合は `startOnLoad` を true にする。
また `theme` は今（v7.1.0）のところ以下の4つが用意されている。
お好みでどうぞ。

- `dark`
- `default`
- `forest`
- `neutral`

さらに mermaid.js に記述部分を認識させるために `<div class="mermaid"> ... </div>` で囲む。

```html
<div class="mermaid">
sequenceDiagram
    カバ->>+カバン: あなた、泳げまして？
    カバン-->>-カバ: いえ
    カバ->>+カバン: 空は飛べるんですの？
    カバン-->>-カバ: いえ
    カバ->>+カバン: じゃあ、足が速いとか？
    カバン-->>-カバ: いえ
    カバ->>カバン: あなた、何にもできないのねぇ
    loop ひとりヘコむ
        カバン->>カバン: ううっ
    end
</div>
```

これで最初に挙げたシーケンス図を描画してくれる。

### Hugo 用の shortcode

[Hugo] 用に [shortcode を作ってみた](https://github.com/spiegel-im-spiegel/github-pages-env/blob/master/layouts/shortcodes/fig-mermaid.html)。

```html
<figure style='margin:0 auto;text-align:center;'>
<div class="mermaid">{{ .Inner }}</div>
{{ if .Get "title"}}<figcaption>{{ with .Get "link"}}<a href="{{.}}">{{ end }}{{ .Get "title" }}{{ with .Get "link"}}</a>{{ end }}</figcaption>{{ end }}
</figure>
```

この shortcode を使うなら以下のように記述すればよい。

```text
{{</* fig-mermaid title="カバとカバン" */>}}
sequenceDiagram
    カバ->>+カバン: あなた，泳げまして？
    カバン-->>-カバ: いえ
    カバ->>+カバン: 空は飛べるんですの？
    カバン-->>-カバ: いえ
    カバ->>+カバン: じゃあ，足が速いとか？
    カバン-->>-カバ: いえ
    カバ->>カバン: あなた，何にもできないのねぇ
    loop ひとりヘコむ
        カバン->>カバン: ううっ
    end
{{</* /fig-mermaid */>}}
```

[Hugo] を使っておられる方は参考にどうぞ。

## 複雑な関係を分かりやすくする

[本ブログ]にシーケンス図を組み込もうと思ったのは徳島県のチケット売買詐欺事件のシーケンスが分かりにくかったからだ。

- [徳島県警察の誤認逮捕事件についてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20170910/1505065248)

で，リンク先の図を元にしてシーケンス図を起こしてみたのがこれ。

{{< fig-mermaid title="「[徳島県警察の誤認逮捕事件についてまとめてみた」より" link="http://d.hatena.ne.jp/Kango/20170910/1505065248" >}}
sequenceDiagram
    participant A as 愛知県女性A
    participant Tw as Twitter
    participant B as 京都女子中学生B
    participant C as 徳島県女子高生C
    participant D as 和歌山県女性D
    participant Tk as チケット売買サイト
    participant E as 関東女性E
    A->>Tw: チケット売ります
    C->>Tw: 購入希望
    D->>Tw: 購入希望
    B->>+Tw: 
    Tw-->>-B: A,B,Cのやり取りを発見
    B->>+A: 売買交渉：8万円で購入希望
    B->>+C: Aとして売買交渉：Aの口座に4万円送金指示
    B->>+D: Aとして売買交渉：Aの口座に4万円送金指示
    B->>+Tk: A名義でチケット出品
    C-->>-A: 4万円振込
    D-->>-A: 4万円振込
    E->>Tk: 6万5千円で落札，送金
    B->>A: 送付先をEの住所へ指示
    A-->>-E: チケット送付
    Tk-->>-B: 約6万円入金
{{< /fig-mermaid >}}

[mermaid] 記法ではこんな感じになる。

```text
sequenceDiagram
    participant A as 愛知県女性A
    participant Tw as Twitter
    participant B as 京都女子中学生B
    participant C as 徳島県女子高生C
    participant D as 和歌山県女性D
    participant Tk as チケット売買サイト
    participant E as 関東女性E
    A->>Tw: チケット売ります
    C->>Tw: 購入希望
    D->>Tw: 購入希望
    B->>+Tw: 
    Tw-->>-B: A,B,Cのやり取りを発見
    B->>+A: 売買交渉：8万円で購入希望
    B->>+C: Aとして売買交渉：Aの口座に4万円送金指示
    B->>+D: Aとして売買交渉：Aの口座に4万円送金指示
    B->>+Tk: A名義でチケット出品
    C-->>-A: 4万円振込
    D-->>-A: 4万円振込
    E->>Tk: 6万5千円で落札，送金
    B->>A: 送付先をEの住所へ指示
    A-->>-E: チケット送付
    Tk-->>-B: 約6万円入金
```

しかし，改めて見るとホンマ凄いよねぇ。
これを本当に一人で考えて実行したのならちょとした天才かもしれん。

## シーケンス図以外の図も描けるよ

さて，今（v7.1.0）のところ  [mermaid] 記法および mermaid.js は以下の図をサポートしている。

- フローチャート
- シーケンス図
- ガント図
- クラス図（experimental）
- Git グラフ（experimental）

たとえば，クラス図は以下のように記述する。

```text
classDiagram
    friends<|--serval
    friends<|--raccoon
    friends<|--fennec
    serval : +Waai()
    raccoon : +Omakase-nanoda()
    fennec : +Haiyo()
```

これを mermaid.js で描画すると以下のようになる。

{{< fig-mermaid title="フレンズ" >}}
classDiagram
    friends<|--serval
    friends<|--raccoon
    friends<|--fennec
    serval : +Waai()
    raccoon : +Omakase-nanoda()
    fennec : +Haiyo()
{{< /fig-mermaid >}}

んー。
クラス図を描くなら多重度は必須なんだがなぁ（たとえば多対多の関係は実装できないため「設計が間違っている」と断言できる）。
これに関しては「これからに期待」といったところか。

## 考えながら描く

クラス図やシーケンス図といったものは試行錯誤しながら描いていくものなので，切ったり貼ったりあるいは結合を変えたりといったことが簡単にできないと意味がない。
したがって，とにかく記述が簡単で見た目が直感的であるというのが大事になってくるわけだ，多少は機能を落としてでも。

そういう意味で [mermaid] はバランスのいい製品と言える。
まぁ，[本ブログ]では頻繁に使うものではないかもしれないが，使える道具はあるに越したことはないだろう。

## ブックマーク

- [シーケンス図(Sequence Diagram) - UML入門 - IT専科](http://www.itsenka.com/contents/development/uml/sequence.html)
- [クラス図(Class Diagram) - UML入門 - IT専科](http://www.itsenka.com/contents/development/uml/class.html)
- [簡単にガントチャートとかクラス図とか書けるやつ - Qiita](https://qiita.com/rana_kualu/items/da394fd33ce019bf0ba7)
- [ちょっとしたシーケンス図を書くのにatom-mermaidが便利すぎる話 - Qiita](https://qiita.com/ririli/items/64320ed2918b1982f89d)
- [Drawing Graphs using Dot and Graphviz](http://www.tonyballantyne.com/graphs.html) : 図にこだわるのであれば [Graphviz](http://www.graphviz.org/ "Graphviz | Graphviz - Graph Visualization Software") の DOT 言語を使う手もある
    - [Graphvizとdot言語でグラフを描く方法のまとめ - Qiita](https://qiita.com/rubytomato@github/items/51779135bc4b77c8c20d)

- [エディタ以上ワープロ未満の HackMD]({{< relref "remark/2017/04/hackmd.md" >}})

[本ブログ]: / "text.Baldanders.info"
[mermaid]: https://mermaidjs.github.io/
[Hugo]: https://gohugo.io/ "Hugo | A Fast and Flexible Website Generator"
