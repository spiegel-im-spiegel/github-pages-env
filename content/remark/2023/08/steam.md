+++
title = "Ubuntu に Steam を入れてみた"
date =  "2023-08-10T21:57:16+09:00"
description = "とりあえず「ドラゴンクエストXI」は動いてるっぽい"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "games" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昔話から始めると，2011年にアナログのテレビ電波が停波したじゃない。
まぁ東北の大震災の影響もあって足並みが揃わなかったけど。
そんで，当時はまだブラウン管テレビだったのよ，うちは。

アナログ停波で無用の長物になったブラウン管テレビは廃棄したんだけど，山ほどあったコンシューマゲーム機を繋ぐ先がなくなってしまったのね。
しかもタブレットやスマホで普通にゲームが出来る時代に入ってしまったことで，コンシューマゲーム（機）は全部押し入れに封印してしまった[^cg1]。

[^cg1]: その後，[引っ越し]({{< ref "/remark/2018/12/move-with-trello.md" >}} "Trello で引っ越し")でコンシューマゲーム（機）は全部処分した `orz`

ときは流れて。

近年 VTuber のゲーム実況をチラチラ見るようになって「やっぱ，終わりの見えない無間地獄のようなスマホゲームじゃなくて，お金払ってちゃんとしたタイトルもやりたい」と思うようになった。
とはいえ，いまさらコンシューマゲーム機を買う気はしないし（[メガドライブミニ2は買うた]({{< ref "/remark/2022/10/megadrive-mini-2.md" >}} "メガドライブミニ2買うた")けどね。若かりし頃の思い出ということで），かといって，[自宅機]({{< ref "/remark/2021/06/new-machine-here.md" >}} "自宅マシンを買うた（これで私も人並みに...）")は Ubuntu でゲーム向きの環境じゃないよなぁ，とかグルグル考えはじめた。

そういや NieR シリーズの [Steam] 版があるんだっけ。
あれはやってみたい...

そうこうするうち

- [SteamゲーマーのOSシェアでついにLinuxがmacOSを追い抜く、「Steam Deck」が影響する可能性 - GIGAZINE](https://gigazine.net/news/20230804-steam-mac-linux/)

なんてな記事も見かけるようになった。
なんでも [Proton] (Wine の fork らしい) によって [Windows 用のゲームの多くが Linux 機でも動く](https://gigazine.net/news/20211223-steam-game-played-on-linux/ "Steamで販売されるPCゲームの約8割がLinuxに対応、LinuxはPCゲームプラットフォームとして成長を遂げている - GIGAZINE")ようになったそうな。
こりゃあ試してみなければ。
せっかく夏休み満喫中だし（笑）

## Ubuntu に Steam をインストールする

ググってみると Ubuntu に [Steam] を入れる方法はいくつかあるようだが，以下の[公式ページ](https://store.steampowered.com/about/ "Steam, The Ultimate Online Game Platform")から deb ファイルをダウンロードして APT でインストールするのが一番確実なようだ。

{{< fig-img-quote src="./steam-about.png" title="Steam, The Ultimate Online Game Platform" link="https://store.steampowered.com/about/" width="1193" lang="en" >}}

APT でのインストールで百個以上の関連パッケージを入れるのを見て若干後悔したり。
これ絶対に更新管理が面倒くさいよなぁ...

## Proton を有効にする

[Proton] を有効にするには「設定」の「互換性」の項目で「他のすべてのタイトルでSteam Playを有効化」をONにすればいいらしい（アプリケーションの再起動が要求される）。

{{< fig-img-quote src="./settings.png" title="STEAM設定 互換性" link="./settings.png" width="850" >}}

この状態で試しに「ドラゴンクエストXI 過ぎ去りし時を求めて S」の体験版（無料）をダウンロード&インストール&起動してみた。

{{< fig-img-quote src="./dq-xi-s.png" title="ドラゴンクエストXI　過ぎ去りし時を求めて S" link="https://store.steampowered.com/app/1295510/XI_S/" width="963" >}}

うっ，やっぱその辺に転がってるゲームコントローラではダメか（笑） しゃーない。
買ってくるか！

{{< fig-img src="./53106156501_eda1b49f1e_e.jpg" title="Steam 用のゲームパッドを購入 | Flickr" link="https://www.flickr.com/photos/spiegel/53106156501/">}}

これならちゃんと動くな。
よしよし。
では仕切り直し。

ほうほう。
オープニングムービーは違和感なく動く。
体験版なのでゲーム画面は3D表示のみなのだが，これも問題なく動いている。
これならイケそうかな。

そういや，最近のゲームの流行りとか全く把握してないな。
ン十年ぶりにゲーム雑誌とか買ったほうがいいのだろうか。

{{< div-box type="markdown" >}}
**【2023-08-11 追記】**
3D動作限定かもしれないけど，たまに OS を巻き添えにしてハングアップする。
こうなるとマシンをリセットするしかないので，ゲームをするときは他のアプリケーションを落とした状態でやらないと。
{{< /div-box >}}

{{< div-box type="markdown" >}}
**【2023-08-12 追記】**
ドラクエとどっちにしようか悩んだが，結局 “[NieR:Automata](https://store.steampowered.com/app/524220/NieRAutomata/ "Steam：NieR:Automata™")” を購入。
今のところは問題なく動いている。
まぁ，念のためにゲーム起動中は極力ほかのアプリケーションは落としてるけど。
{{< /div-box >}}

## ブックマーク

- [ProtonDB | Gaming know-how from the Linux and Steam Deck community](https://www.protondb.com/)
- [自分だけのSteam Machineを組み立てる](https://store.steampowered.com/steamos/buildyourown?l=japanese)
- [【Ubuntu日和】【第3回】Ubuntuでもエルデンリングを動かせる！ SteamでWindows用のゲームをプレイしよう  - PC Watch](https://pc.watch.impress.co.jp/docs/column/ubuntu/1409524.html)

[Steam]: https://store.steampowered.com/
[Proton]: https://github.com/ValveSoftware/Proton "ValveSoftware/Proton: Compatibility tool for Steam Play based on Wine and additional components"

## 参考

{{% review-paapi "4571210450" %}} <!-- はじめて学ぶ ビデオゲームの心理学 -->
{{% review-paapi "B00CDG7994" %}} <!-- Steam ゲームコントローラ -->
