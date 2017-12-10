+++
title = "「神」と「神」と Chrome の文字化け"
date =  "2017-12-10T12:35:07+09:00"
update =  "2017-12-10T22:06:56+09:00"
description = "Google Chrome では「神」と「神」が同じ「神」に見える，という報告をいただきまして，確かにこちらの環境でも再現するので，ちょっと調べてみました。"
image = "https://farm5.staticflickr.com/4586/38061039025_a61bb35c85.jpg"
tags        = [ "web", "site", "font", "unicode", "normalization", "character" ]

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
  mathjax = false
  mermaidjs = false
+++

「[Go 言語による Unicode 半角/全角変換]({{< relref "golang/character-width-for-unicode.md" >}})」のページを Google Chrome で見ると「神と神」が同じ「神」に見える，という報告をいただきまして，確かにこちらの環境でも再現するので，ちょっと調べてみました。

（追記： Chrome だけじゃなくて [Android 用の Firefox Focus](https://play.google.com/store/apps/details?id=org.mozilla.focus) でも同じ現象を確認した。 Firefox Focus って Firefox とエンジンが違うのか？ ひょっとして WebKit？ iOS から先にリリースされてるし。もしかして Safari でも同じ現象が起きる？）

ちなみに「神と神」のうち前者の「神」は所謂 CJK 互換文字と呼ばれるもので Unicode コードポイントで「`U+FA19 '神'`」にマッピングされるものです。

まず，[本サイト]の状態を説明しておくと，サイトの各ページの文字は Web フォントで表示されている。
ブラウザ側で Web フォントを無効にしていない（または経路途中でフィルタリングされない）限りは，こちらで指定したフォントで表示されるはずである。

- [Web フォントに関する覚え書き]({{< relref "remark/2015/web-font-family.md" >}})
- [Web フォントに関する覚え書き（明朝体編）]({{< relref "remark/2016/10/japanese-serif-fonts-by-google-cdn.md" >}})

本文を含む地の文章には「[さわらび明朝]」を利用している。
[Google Fonts Early Access](https://fonts.google.com/earlyaccess) で唯一ふだん使いできる明朝体の Web フォントだったのが選択した理由である[^ns1]。

[^ns1]: 最近発表された Noto Serif の和文フォントはまだ Web フォントとしては提供されていない。

ただし「[さわらび明朝]」は，個人の方がボランティアで作成されているということもあり，全ての和文文字を網羅しているわけではなく，幾つか表示できない文字がある。
「`U+FA19 '神'`」もそのうちのひとつである。
どうもその辺に原因がありそうである。

そこで，テスト用に以下の HTML コードを用意し Chrome に表示させてみた。

```html
<html>
<head>
<link rel='stylesheet' href='http://fonts.googleapis.com/earlyaccess/sawarabimincho.css' type='text/css'>
<title>神と神</title>
<style>
#localonly {
  font-family: sans-serif;
}
#sawarabi {
  font-family: 'Sawarabi Mincho', sans-serif;
}
</style>
</head>
<body>
<div id="localonly">
	<p>	「神」と「神」（後者が CJK 互換文字）</p>
	<p>草「彅」剛（JIS 第一・第二水準にない漢字）</p>
	<p>	「繋」ぐ（さわらび明朝にない文字）</p>
</div>
<div id="sawarabi">
	<p>	「神」と「神」（後者が CJK 互換文字）</p>
	<p>草「彅」剛（JIS 第一・第二水準にない漢字）</p>
	<p>	「繋」ぐ（さわらび明朝にない文字）</p>
</div>
</body>
</html>
```

これを表示させた結果は以下の通り。

{{< fig-img src="https://farm5.staticflickr.com/4729/38947212091_dceb4778f1.jpg" title="Google Chrome (1)" link="https://www.flickr.com/photos/spiegel/38947212091/" >}}

最初の3行が Chrome デフォルトの Sans Serif で表示させたもの，次の3行が「[さわらび明朝]」で表示させたものである。
ただし「[さわらび明朝]」にない文字は Sans Serif で表示するよう設定している。

「神」と「彅」と「繋」は「[さわらび明朝]」に収録されていない文字なのだが，何故か「神」だけが明朝体で「神」と表示されているのがお分かりだろうか[^sm1]。
ふむむむむ。

[^sm1]: デフォルトの Sans Serif が使われていないことは Chrome のデベロッパー・ツールでも確認済み。

ここで思い出すのが Unicode 正規化による「神」の正規化である。

- [Go 言語と Unicode 正規化]({{< relref "golang/unicode-normalization.md" >}})

この記事で説明している通り， CJK 互換文字である「`U+FA19 '神'`」は正規化によって「`U+795E '神'`」に変換されてしまうのである。

ならば今度は CJK 互換文字に絞って以下の HTML コードを組んでみる。

```html
<html>
<head>
<link rel='stylesheet' href='http://fonts.googleapis.com/earlyaccess/sawarabimincho.css' type='text/css'>
<title>神と神</title>
<style>
#localonly {
  font-family: sans-serif;
}
#sawarabi {
  font-family: 'Sawarabi Mincho', sans-serif;
}
</style>
</head>
<body>
<div id="localonly">
	<p>	「朗」と「朗」（後者が CJK 互換文字，正規化される）</p>
	<p>	「欄」と「欄」（後者が CJK 互換文字，正規化される）</p>
	<p>	「崎」と「﨑」（後者が CJK 互換文字，正規化されない）</p>
</div>
<div id="sawarabi">
	<p>	「朗」と「朗」（後者が CJK 互換文字，正規化される）</p>
	<p>	「欄」と「欄」（後者が CJK 互換文字，正規化される）</p>
	<p>	「崎」と「﨑」（後者が CJK 互換文字，正規化されない）</p>
</div>
</body>
</html>
```

これを表示させた結果は以下の通り。

{{< fig-img src="https://farm5.staticflickr.com/4586/38061039025_a61bb35c85.jpg" title="Google Chrome (2)" link="https://www.flickr.com/photos/spiegel/38061039025/" >}}

ビンゴ！

- 「朗」と「朗」は「神」と「神」の関係と同じ。 CJK 互換文字「朗」は「[さわらび明朝]」に収録されてなく， Unicode 正規化によって「朗」に変換される
- 「欄」は Unicode 正規化によって「欄」に変換されるが，どちらの文字も「[さわらび明朝]」に収録されていないため， Chrome デフォルトのフォントで表示されている
- 「﨑」は CJK 互換文字で「[さわらび明朝]」に収録されていないが Unicode 正規化の対象ではない（「崎」は似た字形の文字を並べてみただけ，テヘペロ）ため，そのまま Chrome デフォルトのフォントで表示されている

つまり

1. フォントに対象の文字が収録されていない
2. Unicode 正規化によって別の文字に変換可能
3. 変換された文字がフォントに収録されている

という条件が揃った時に「文字化け」が発生する，ということのようだ[^vs1]。

[^vs1]: 今回は「異体字セレクタ（Variation Selector）」については検証していない。そもそも異体字セレクタによる文字表示はアプリケーションとフォントの両方が対応していないと成り立たない。メジャー・ブラウザは異体字セレクタに対応しているそうだが，私は対応するフォントを持ってないし，対応できる自由な Web フォントが（CDN 等で）提供されているという話も聞かない。したがって今回はスルーの方向で。

いや，これは酷いだろう。
同じ文字化けするなら「豆腐 □」や「ゲタ 〓」のほうがまだマシである。
多分 Chrome の中の人は CJK 互換文字や異体字について（歴史背景などを）よく理解しないまま「正規化できる文字で代替えすればいいぢゃん♥」と軽い感じで実装してしまったのだろう。
「[Go 言語と Unicode 正規化]({{< relref "golang/unicode-normalization.md" >}})」では Apple の旧ファイルシステム HFS+ を散々に貶したが， **Google もタイガイである**。（Chrome だけじゃなかった）

ちなみに，この現象は Firefox （またはその系列のブラウザ）では発生しない。
私のメインブラウザは Firefox (Quantum マンセー！) なので，指摘されるまで全く気づかなかった。
ゴメンペコン[^c1]。

[^c1]: 一応このサイトは Firefox (Quantum) と Chrome ならレイアウトが崩れないように考えているつもり。旧 Firefox や IE や Safari などレガシーなブラウザはスルーで。 Edge は対応したいところだが手元に環境を持ってないので確認＆対応できない。 Win10 を導入する予定は，今のところはない。換装するなら Linux 系のデスクトップにするつもり。

とはいえ，これは相当にレアケースなので滅多なことでは発現しないだろう。
個人的には Noto Serif の和文フォントが Web フォントで登場するまでは「[さわらび明朝]」で頑張るつもりである。
なので，特に対策しない予定。

ご不便をおかけすることもあるかもですが，悪しからずご了承の程を。

## ブックマーク

- [「Noto Serif Japanese」を使ってみたので、自分用メモ。 - Qiita](https://qiita.com/umeume66/items/01291353fc43c17da992)

[本サイト]: / "text.Baldanders.info"
[さわらび明朝]: http://sawarabi-fonts.osdn.jp/ "さわらびフォント"
