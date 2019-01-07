+++
date = "2016-05-05T00:35:18+09:00"
description = "どうも Jota+ ってのがいいらしい。"
draft = false
tags = ["tools", "editor", "android"]
title = "Android で使えるテキストエディタ"

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

わけあって Android で使えるテキストエディタを物色中。

- [イマドキのIDE事情 (151) メモからプログラミングまで!? Android上で動作するテキストエディタたち | マイナビニュース](http://news.mynavi.jp/column/ide/151/)

どうも [Jota+] ってのがいいらしい。

[Jota+] 自身は無料だが [PRO-KEY](https://play.google.com/store/apps/details?id=jp.sblo.pandora.jota.plus.prokey) を購入して activation すると色々いいことがある。
まず広告表示を無効にできる。
あとストレージのアクセスに [Storage Access Framework](http://developer.android.com/intl/ja/guide/topics/providers/document-provider.html) が使えるようになる。
これは特にクラウド・ストレージを利用する際に便利だ。

[Jota+] 自身も内臓機能でクラウド・ストレージにアクセスできるのだが，サービスごとに connector アプリを入れる必要がある。
SAF を有効にしておけばバックグラウンドで [Box](https://play.google.com/store/apps/details?id=com.box.android) や [Dropbox](https://play.google.com/store/apps/details?id=com.dropbox.android) を起動した上で SAF 経由で各サービスにアクセスできる。

って，これって少なくともテキスト情報に限れば [Evernote](https://play.google.com/store/apps/details?id=com.evernote) や [Google Keep](https://play.google.com/store/apps/details?id=com.google.android.keep) は要らないってことか。
善き哉。

そうそう。
[Jota+] は既定で [Google Analysis による情報収集](https://sites.google.com/site/aquamarinepandora/jotaplusja/privacypolicyja)を行っている。
これは設定で無効にできるので，気に入らないのであれば無効にすることをお薦めする。

[Jota+]: https://play.google.com/store/apps/details?id=jp.sblo.pandora.jota.plus "Jota+ (Text Editor) - Google Play"
