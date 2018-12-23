+++
title = "Qiita の90パーセントはカスである"
date =  "2018-11-13T00:37:44+09:00"
description = "変わらず「玉石混交」なのが Qiita のいいところだと思っている。"
image = "/images/attention/kitten.jpg"
tags = [ "media", "cms" ]
draft = true

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

俗に言う「スタージョンの法則[^sr1]」ってあるじゃん。
そのうちのひとつ「どんなものも，その90%はカスだ（ninety percent of everything is crud）」ってかなり真理だと思うのよ。

[^sr1]: 厳密には「スタージョンの黙示（Sturgeon's Revelation）」らしい。引用は [Wikipedia の記事](https://en.wikipedia.org/wiki/Sturgeon%27s_law "Sturgeon's law - Wikipedia")（[日本語](https://ja.wikipedia.org/wiki/%E3%82%B9%E3%82%BF%E3%83%BC%E3%82%B8%E3%83%A7%E3%83%B3%E3%81%AE%E6%B3%95%E5%89%87 "スタージョンの法則 - Wikipedia")）より。

これは1970年代に SF 作品群に向けられた「SFの90%はカスだ（ninety percent of SF is crud）」という批判に対する Theodore Sturgeon の言葉なのだそうだ。
90% という数字の妥当性はともかく，成熟したメディアはこの法則に近似していくと思う。

もし [Qiita] の記事がカスばかり溢れているというのなら，それは [Qiita] がメディアとして成熟してきた証でありユーザとしては歓迎すべき事態なのである。
それよりも，出てくるべくして出てきた「90%のカス」を如何にフィルタリングするか，またフィルタリングされた情報の中に適度な「誤配」を混ぜ込むにはどうすべきか，考えていくのが上策だろう。

まぁ，実際にはどこもそれを上手くこなせないから Google 検索も Twitter も facebook もあんな有様なんだけど。
しかもフィルタリングをユーザ側でコントロールできないから「[フィルターバブル]」化してしまっている。

[Qiita] 側が「90%のカス」に対し特にフィルタリングを行っていないというのは英断と言えるだろう。
それを主導するのはサービス・プロバイダ（あるいはもっと広く技術的ゲートキーパー）ではなく，あくまでもユーザ側であるべきだ。

私は2015年から [Qiita] を利用し始めていて，しかも [Go 言語]を勉強し始めたのも同年からだったのでとても重宝した。
あれから4年近く経つが記事自体の質が特に良くなったり悪くなったりという印象はない。
変わらず「玉石混交」なのが [Qiita] のいいところだと思っている。

確かに一時期 Google 検索で技術系のキーワードを探すと [Qiita] 記事が上位に来ることが多かったが，今はそんなことないよね[^gs1]。
まぁ，私の場合は既定の検索サービスを [DuckDuckGo] にして Google 検索は補助的に利用してるだけなので余計にそう感じるのかも知れないが[^ddg1]。

[^gs1]: Google 検索の結果は個人や地域によってかなり変わるらしいけど。
[^ddg1]: [DuckDuckGo] ってオプションで日本を指定しないと日本語の記事が上位に来ないんだよね。考えてみたら日本語みたいな辺境国の言語が検索結果の上位に来るなんておかしいわけで，それをおかしいと思わないなら既に「[フィルターバブル]」に取り込まれてるってことだ（笑）

「個人サイトのほうが技術証明になる」というのもどうなんかねぇ。
15年前までなら個人サイトを立てて日記やブログを運営するってのは「技術証明」になったかも知れないけど，もはや XaaS が完全に背景化した時代に，そんなのサルでもできるだろ。
[WordPress を個人で運営するなんてリスクやコストばかり]({{< ref "/remark/2016/07/cms.md" >}} "「自分で面倒見られる子」だけが CMS を導入しなさい")。
そんなマゾみたいな活動が「個人サイト」でやりたいことなの？ 個人で CMS を運用するなら Jekyll か Hugo で十分だよ。

「[「オブジェクト指向」の黒歴史]({{< ref "/remark/2018/10/object-oriented-design.md" >}})」でも書いたが，プログラミング言語の進歩や FOSS が一般化したことによりプログラムコードはコミュニケーション手段としても使われるようになった。
[Qiita] はそうしたコミュニケーション手段を提供しているという点でよくできたサービスだと思う。
これはかつての /.J や Q&A サービスなどとは一線を画している。

まぁ，今のエンジニアにとって究極のコミュニケーション・ステージは [GitHub](https://github.com/) だけど。

エンジニアなら，理屈はコードで示せ！ ってなもんである（笑）

## ブックマーク

- [全ての開発者がQiitaへのアウトプットをやめるべき理由 - Qiita](https://qiita.com/qiitadaisuki/items/2160a390ce91283707a1)

[Qiita]: https://qiita.com/
[Go 言語]: https://golang.org/ "The Go Programming Language"
[DuckDuckGo]: https://duckduckgo.com/
[フィルターバブル]: {{< ref "/remark/2016/05/filter-bubble.md" >}} "『フィルターバブル』を読む"
