+++
title = "プログラミング言語の暗黙ルール"
date =  "2017-10-25T12:32:16+09:00"
description = "これは Scrapbox に書いた記事の再構成です。内容はほぼ同じ"
tags        = [ "programming", "language" ]

[author]
  name      = "Spiegel"
  url       = "https://baldanders.info/profile/"
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

（これは [Scrapbox に書いた記事](https://scrapbox.io/spiegel-branch/%E3%80%8C%E3%83%97%E3%83%AD%E3%82%B0%E3%83%A9%E3%83%9F%E3%83%B3%E3%82%B0%E8%A8%80%E8%AA%9E%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%80%8D%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6)の再構成です。内容はほぼ同じです）

- [プログラミング言語について | プログラミング | POSTD](http://postd.cc/the-language-of-programming/)

英語不得手な私としてはこの記事で延々と聞かせられる「愚痴」に深く同情はするが，しょせん「愚痴」は「愚痴」でしかない。

これで思い出したのが「[頼むからプログラミングを学ばないでくれ](https://jp.techcrunch.com/2016/05/17/20160510please-dont-learn-to-code/ "頼むからプログラミングを学ばないでくれ | TechCrunch Japan")」である。
この記事については私も以前にブログ記事を書いた。

- [「だれもがプログラミングを学ぶべき」ではない]({{< ref "/remark/2016/05/lets-programming.md" >}})

ちょうど「小学生の学校教育カリキュラムに『プログラミング』を導入する」などという頭の悪い教育政策を聞かされた頃だったので，反発する意味で書いたのだが

{{< fig-quote title="「だれもがプログラミングを学ぶべき」ではない" link="/remark/2016/05/lets-programming/" >}}
<q>日本語や英語を習得するのに文法から習う人はいないだろう（日本の学校教育は違うかもw）。 たくさんの言葉を聞いて話して書いて読んで，そうして少しずつ語彙を飲み込んでいって習得していくものだ。<br>
プログラミング言語は違う。 プログラミング言語で決定的に重要なのは言語仕様つまり文法である。 何故ならプログラミングとは，究極的には，ゼロ知識から論理を積み上げていくことであり，プログラミング言語はそのための道具であり手段なのだ。</q>
{{< /fig-quote >}}

が主張の全てである。

「[プログラミング言語について](http://postd.cc/the-language-of-programming/ "プログラミング言語について | プログラミング | POSTD")」で唯一面白かった指摘は，プログラミング言語には言語仕様とは別に暗黙的に決められているルールというか規範のようなものが存在するという点。
それが「名前」である。
この点については日本語圏でも参考になるページが色々ある。
たとえば以下のページ。

- [正しいコーディングが身につくエンジニア英語の手引き 〜文法とクラス／メソッド、命名規則〜 | Find Job! Startup](https://www.find-job.net/startup/english-for-engineers-naming-conventions)
- [モデルやメソッドに名前を付けるときは英語の品詞に気をつけよう - Qiita](https://qiita.com/jnchito/items/459d58ba652bf4763820)

実は [Go 言語]にも「名前」に関する議論がある。
たとえば

- [Big Sky :: Names](https://mattn.kaoriya.net/software/20160126101358.htm)

のような話だが，ここでも英語の語彙が組み込まれている（ということに今気がついた）。

つまり，英語圏で作られたプログラミング言語では暗黙的に英語の語彙を要求しているのである。
こればっかりはいくらプログラミングを勉強しても身につかないし，英語を無視して作ったものはまさに「[中国語の部屋](https://ja.wikipedia.org/wiki/%E4%B8%AD%E5%9B%BD%E8%AA%9E%E3%81%AE%E9%83%A8%E5%B1%8B "中国語の部屋 - Wikipedia")」と同じく機械的なものにならざるを得ない[^lang1]。

[^lang1]: 言い換えれば日本語のプログラミング言語は日本語の語彙を前提にしているわけで，そんなもん日本（語圏）人以外には使えない。ただでさえ日本語は難しいと外国語圏からは言われているのに。いまや IT 後進国の日本が自国語のプログラミング言語を作るとかヘソで茶が沸いてしまう。

だから若者たちよ。
悪いことは言わないから英語を習得しなはれ。
語彙を育むためには

{{< fig-quote title="「だれもがプログラミングを学ぶべき」ではない" link="/remark/2016/05/lets-programming/" >}}
<q>たくさんの言葉を聞いて話して書いて読んで，そうして少しずつ語彙を飲み込んでいって習得していく</q>
{{< /fig-quote >}}

しかないのだ。
まぁ，私はこれが大変苦手なのだが。
だからこの歳でもいまだ英語不得手のままである。
とほほ。

## ブックマーク

- [プログラミング言語との付き合い方]({{< ref "/remark/2015/programming-language.md" >}})
- [プログラミングで「計算論的思考」は身につかない]({{< ref "/remark/2016/09/programming.md" >}})
- [Codic ATOM Package - codic blog](http://blog.codic.jp/2015/11/20/release-codic-atom-package/)

[Go 言語]: https://golang.org/ "The Go Programming Language"
