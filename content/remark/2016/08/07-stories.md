+++
date = "2016-08-07T12:48:32+09:00"
update = "2016-08-07T14:01:59+09:00"
description = "「UNIX コマンドを SQL で抽出できるツール qq」が面白そう / Pokémon GO 最大の失敗は / Kindle Unlimited / その他の気になる記事"
draft = false
tags = ["tools", "golang", "games", "pokemon", "kindle", "e-book"]
title = "週末スペシャル： 「UNIX コマンドを SQL で抽出できるツール qq」が面白そう"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
+++

~~明日~~今日は立秋です[^0]。
めっさ暑いけど。
でも，今日からは「残暑」なんだよねぇ。

[^0]: 明日だと勘違いしてた。ゴメンペコン。

それにしてもカープ。
首脳陣は「リーグ優勝阻止」に向けて本腰を入れてきているようです。
まぁ，いままで「馬力」だけで勝ってきたようなものなので，カープ側の選手に（バットで殴られたりとかして）けが人が出たり相手チームがきっちり対策してくればひとたまりもないわけで，やはり夢は夢で終わるのかと早くも地元カープファンは落胆しています。

まぁ，まかり間違って優勝して某監督が「名将」などと勘違いされても来年以降に何の展望も見えないので，このまま今年限りでとっとと勇退していただきたいところです。

1. [「UNIX コマンドを SQL で抽出できるツール qq」が面白そう]({{< relref "#qq" >}})
1. [Pokémon GO 最大の失敗は]({{< relref "#pg" >}})
1. [Kindle Unlimited]({{< relref "#ku" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 「UNIX コマンドを SQL で抽出できるツール qq」が面白そう{#qq}

- [Big Sky :: UNIX コマンドを SQL で抽出できるツール qq を作った。](http://mattn.kaoriya.net/software/lang/go/20160805190022.htm)

さっそくビルドしようとしたが Windows 環境では gcc がないと怒られた。
多分 [MSYS2]({{< ref "/remark/2016/03/gcc-msys2-1.md" >}}) なら何とかなると思うが，そこまでする気力がない[^a] のでまた今度。

[^a]: 自宅マシンを Linux ベースに換装すると[決めて]({{< ref "/remark/2016/03/13-stories.md#win" >}})から（年内を目標に準備中） Windows 環境のカスタマイズへの興味が薄れつつある。

[qq] では文字エンコーディングを指定できるようで，

```text
$ cat meisai.csv | qq -ic -e cp932 -q "select 単価 from stdin where 品目 = 'みかん'"
```

とかできるらしい。
一瞬「CP932」のエンコーディングができるの？ と思ったが，使われている [go-encoding] を見たら，どうやら Shift-JIS として解釈するようだ。
やっぱ Shift-JIS や EUC-JP の[方言は簡単じゃない]({{< ref "/golang/transform-character-encoding.md" >}})よなぁ。

## Pokémon GO 最大の失敗は{#pg}

[Pokémon GO] 最大の失敗は夏休み前リリースでお馬鹿ユーザが大量に湧いて出たことか？ と Facebook でつぶやいたら「高校生の息子ももう若干飽き始めてるし、休み明けに勉強に身を入れるには良いタイミングだったのかも」と言われてなるほどと思った。

平和公園のポケストップやジムが全部撤去されたそうで，ゲーム画面上は荒涼とした大地が広がっているそうだ。
こればっかりは土地の所有者や管理者の意向が最優先されるべきなのでしょうがないとは思うけど，ポケモンを規制するくらいなら春先のあの乱痴気騒ぎの方をどうにかすべきじゃないのかと思ってしまうのは私だけだろうか[^b]。
原爆ドームを「世界遺産」化して観光地にしてしまったのは彼等だろうに。

[^b]: 被爆二世の知り合いに聞いたことがあるが，平和公園で毎年行われる「花見」と称する乱痴気騒ぎには複雑な思いがあるそうだ。その知り合いも「花見自体は否定しないが私は参加しない」と言っていた。

近所の公園から子どもがいなくなったように，完全なるシステムでは人は排除されるものなのでこういうのも仕方ないのであろう。
夏休みが終わる頃にはどうなっているやら。

## Kindle Unlimited{#ku}

- [ASCII.jp：電子書籍読み放題の「Kindle Unlimited」が月額980円で日本でもスタート](http://ascii.jp/elem/000/001/205/1205073/)
- [月額980円で読み放題　日本版「Kindle Unlimited」提供スタート　 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1608/03/news048.html)
- [Kindle Unlimited開始でKindleダイレクト・パブリッシング（KDP）冬の時代が到来？ 出版社が優遇され過ぎ:見て歩く者 by 鷹野凌](http://www.wildhawkfield.com/2016/08/major-publishers-have-been-special-treatment-on-Kindle-Unlimited.html)
- [キンドル・アンリミテッド登場は何を意味するか « マガジン航[kɔː]](http://magazine-k.jp/2014/07/30/kindle-unlimited/) ： 2014年の記事

実は [paperwhite]({{< ref "/remark/2016/07/kindle-paperwhite.md" >}} "今更 Kindle Paperwhite を買う") では以前から unlimited 対象の書籍が見えるようになっていて，そのラインナップのあまりのつまらなさに今回のニュースにもほとんど食指が動かなかった。

{{< fig-quote title="電子書籍読み放題の「Kindle Unlimited」が月額980円で日本でもスタート" link="http://ascii.jp/elem/000/001/205/1205073/" >}}
<q>読み放題の対象となる作品は、ウェブ上でアイコンが表示されるようになります。読む手順はカンタンで、通常の単体購入時と同じく、読みたい作品をまず自分のクラウドライブラリーに登録し、読みたい端末にダウンロード。Kindle Unlimited対象書籍は一度に10冊まで登録可能で、11冊目を登録する際は事前に1冊分を解除しておく必要があります。</q>
{{< /fig-quote >}}

ということで，同時に10冊までとかナイわ（笑）

海外で launch したときは「水のような読書」を連想したけど，実際にはそうではなく，どうもこれは「定額レンタルサービス」って感じ。
昔あったよね。
TSUTAYA とか DMM とかがやってたやつ。
定額制で DVD 借り放題だけど前のものを返さないと次のものが借りれないシステム。
だから常に新しいものを借り続けないと全く元が取れない。

まぁうちの親のようにジャンル問わず片っ端から「読み捨てる」タイプの人なら良いサービスなんだろう。
私は要らない。

## その他の気になる記事{#other}

- [日本の火星衛星探査計画「MMX」、打ち上げが1〜2年程度遅れて2024年頃に | 月探査情報ステーション](http://moonstation.jp/blog/marsexp/mmx/japanese-mars-satellite-exploration-mmx-will-delay-one-or-two-years)
    - [日本の「火星衛星」サンプルリターン計画、2022年から1〜2年ほど遅れか | sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/030201/2016_08_05_fob.html)
- [EUの一般データ保護規則と改正個人情報保護法--日本企業が気をつけるべきことは - ZDNet Japan](http://japan.zdnet.com/article/35086772/)
- [プラネタリウム番組「ネイチャーリウム オーロラの調べ 神秘の光を探る」特別企画](http://www.pyonta.city.hiroshima.jp/event/detail/id/2942.html) : 9月10日。申し込みは往復はがき9月3日（土）必着
- [WindowsにMicrosoftアカウントが盗まれる既知の脆弱性--概念実証サイトが公開に - ZDNet Japan](http://japan.zdnet.com/article/35086867/)
- [Windows10 Anniversary updateで知らぬ間にSSHdが起動している : やすひでぶろぐ](http://yasuhide.blog.jp/archives/48155574.html)
- [Docker、オーケストレーション機能のSwarmモードを搭載した「Docker 1.12」が正式版に － Publickey](http://www.publickey1.jp/blog/16/dockerswarmdocker_112.html)
    - [Docker for Mac/Windowsが正式版としてリリース － Publickey](http://www.publickey1.jp/blog/16/docker_for_macwindows_1.html)
    - [コンテナー管理ソフト「Docker」のWindows版「Docker for Windows」が正式版に - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1013117.html)
- [原爆の実相を世界へ：英語版も加わり新しくなった「ヒロシマ・アーカイブ」 · Global Voices 日本語](https://jp.globalvoices.org/2016/08/01/42121/)

[qq]: https://github.com/mattn/qq "mattn/qq"
[go-encoding]: https://github.com/mattn/go-encoding "mattn/go-encoding"
[Pokémon GO]: http://www.pokemongo.jp/ "『Pokémon GO』公式サイト"
