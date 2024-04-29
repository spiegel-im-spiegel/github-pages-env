+++
title = "整いました！"
date =  "2020-09-24T21:55:47+09:00"
description = "これで新しいリビング PC の完成である。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "android", "linux", "gear" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

久しぶりに Go 言語関連のイベントに参加しようと[エントリした](https://gpl-reading.connpass.com/event/188380/ "第5回『プログラミング言語Go』オンライン読書会 - connpass")のだが，このご時世でオンラインイベントなのですよ。

そんで Zoom 環境を用意しなければならないんだけど，ぶっちゃけ自分のパソコンは使いたくない。
なので[昨年買った中古ノート][昨年]に入れようと久〜しぶりに電源を入れたら OS が起動しない！？ どうやら内蔵ストレージが完全にお亡くなりになっているらしい `orz`

参加をキャンセルすることも考えたが，地元でもオンラインもくもく会やるみたいな話も聞こえてくるし（まだアナウンスがないし立ち消えかな？），諦めて新しいマシンを買うことにした。
失業してるのにお金が出ていくばっかりだよ...

## あまり選択肢はないらしい

取り敢えず今回は予算を「3万円」に決めて，最初に中古ノート PC を物色してみたのだが，今回の反省を踏まえて考えると，1年位で故障しそうなのばかり（笑） 多分この価格帯のマシンは「部品取り」用なのだろう。
ホンマ，[昨年のアレ][昨年]は「安物買いの銭失い」だったんだなぁ...

となるとタブレットか Chromebook か，それともラズパイで自作するか。

自作は興味あるけど絶対にドツボに嵌まるから今回はパス。
Chromebook も同じく興味はあるが「できること」の範囲がよく分からないので，これもパス。
となるとタブレットか。

タブレットの選択肢は3つ。
Android 機か iPad か Windows 機か。
と思ったのだが，予算的に Android 機以外は無理，と判明した（笑）

そこで Android タブレットに絞って探してみたのだが... ぶっちゃけ HUAWEI 社以外の製品は「なし」だわ。
でも HUAWEI 社はなー。
米国大統領が再選したらいよいよヤバいんじゃね？ なんでこんなに選択肢がないのだろう。
しゃーない。
HUAWEI 社にするか。
いいたかねえけどめんどうみよう（あわれ仁吉よ，どこへ行く）

さんざん悩んだ末に [HUAWEI MediaPad M5 lite 10](https://www.amazon.co.jp/dp/B07KJ5S7FS?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1) に決めた。
タブレットが決まれば，あとは五月雨式に決まっていく。

## 整いました！

というわけで

{{< fig-img src="./50376693028_e67cc61968_e.jpg" title="整いました！ | Spiegel | Flickr" link="https://www.flickr.com/photos/spiegel/50376693028/" >}}

Bluetooth 接続のキーボードとマウスと[骨伝導ヘッドセット]({{< ref "/remark/2019/10/15th-anniversary-of-baldandersinfo.md" >}} "Baldanders.info 開設15周年記念に骨伝導スピーカーを買うたった！")は既存のものがちゃんと繋がった。
スタンドは今回新たに購入（今まで使ってたスマホスタンドは100円のやっすいやつだった）。
スマホと兼用だが，10インチのタブレットでも問題なく支えられるようだ。
よしよし。
これで新しいリビング PC の完成である。

タブレットを起動してみて思ったのだが，一応 “Powered by Android” ってなってるけど「[Android 風のなにか](https://mo-no-log.com/emui/ "HUAWEI「EMUI」とは？機能やデザインについて徹底解説!!｜モノログ")」だよね，これ。
Android アプリが使えるってだけで，これを Android 端末と呼ばわるのは詐欺なんじゃないかなぁ（笑） 

スマホもそうだけど，もう “Android” という生態系は崩壊してるのかもしれない。
なんか1990年代の UNIX 戦争を彷彿とさせるよ。
おそらく Google 的には既に生態系としての “Android” を見限っていて「タブレット買うくらいなら Chromebook にしろよ」ってとこなのだろう。

## Android タブレットでコードが書けるか

ところで，折角なので今回はタブレット上でコードが書ける環境を作ろうと思っている。

まずは上の写真のように [Termux] を導入してみた。
[Termux] って Windows でいうところの MSYS2 みたいな位置づけなのかな。
Linux ぽい作りでパッケージ管理に APT が使える優れものだが WSL2 みたいなサブシステムというわけではなさそうだ。
この辺はいつか記事にしよう。

ほんで，テキストエディタ。
軽くググってみたが，以下の4つがよさげ。

- [DroidEdit Pro (code editor) - Apps on Google Play](https://play.google.com/store/apps/details?id=com.aor.droidedit.pro)
- [anWriter text editor - Apps on Google Play](https://play.google.com/store/apps/details?id=com.ansm.anwriter.pro)
- [QuickEdit Text Editor Pro - Writer & Code Editor - Apps on Google Play](https://play.google.com/store/apps/details?id=com.rhmsoft.edit.pro)
- [Turbo Editor PRO | Text Editor - Apps on Google Play](https://play.google.com/store/apps/details?id=com.maskyn.fileeditorpro)

無料版でそれぞれ試してみたところ [QuickEdit] がもっとも手に馴染んだので，君に決めた！ テキストエディタは機能以前に手に馴染むかが[最優先事項よ！](https://ameblo.jp/kikuko-inoue/entry-12212589968.html)

## ブックマーク

- [Termux - Apps on Google Play](https://play.google.com/store/apps/details?id=com.termux)
- [How to install Git on Android - TechRepublic](https://www.techrepublic.com/article/how-to-install-git-on-android/)
- [Termux on AndroidのSSHサーバに接続する方法 | LFI](https://linuxfan.info/termux-sshd)
- [6 Best Android Text Editor for Programming | TechWiser](https://techwiser.com/android-text-editor-for-programming/)
- [いつでもLinuxコマンドが使える！Androidで動くLinux端末「Termux」【Root化不要】 | LFI](https://linuxfan.info/termux)
- [web帳 | Androidに Linuxを簡単にインストールするアプリ「UserLAnd」](https://www.webcyou.com/?p=9476)
- [耳のためにできること。1万円を切る骨伝導ヘッドセットAfterShokz OpenMoveを見てきた。｜塚本 牧生｜note](https://note.com/tsukamoto/n/nacdd57144504)
- [【特集】5万円以下のChromebookは仕事に使えるのか? ～Chromebook活用術【ビジネス編】 - PC Watch](https://pc.watch.impress.co.jp/docs/topic/feature/1270376.html)
- [「Zoom」v5.3.0が公開 ～ユーザー側から参加するブレイクアウトルームを選択可能に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1278352.html)

[昨年]: {{< ref "/remark/2019/12/install-ubuntu-to-second-hand-pc.md" >}} "中古 PC に Ubuntu 環境を導入する"
[Termux]: https://termux.com/
[QuickEdit]: https://play.google.com/store/apps/details?id=com.rhmsoft.edit.pro "QuickEdit Text Editor Pro - Writer & Code Editor - Apps on Google Play"

## 参考

{{% review-paapi "B07KJ5S7FS" %}} <!-- HUAWEI MediaPad M5 lite 10 タブレット -->
{{% review-paapi "B07HWSDG97" %}} <!-- デスクトップスタンド スマホ タブレット -->
{{% review-paapi "B00ZP3503O" %}} <!-- iClever Bluetooth キーボード -->
{{% review-paapi "B00G9NIL7G" %}} <!-- エレコム マウス Bluetooth -->
{{% review-paapi "B07QJB7R13" %}} <!-- Bluetooth イヤホン 骨伝導 ヘッドホン -->
{{% review-paapi "B011LC4LY6" %}} <!-- めんどうみたョ -->


<!-- eof -->
