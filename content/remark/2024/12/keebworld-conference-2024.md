+++
title = "「ミリしら」で KeebWorld Conference 2024 へ行ってみた"
date =  "2024-12-07T23:27:29+09:00"
description = "ただ聴きに来ただけという感じになってしまった。発表に対して何を質問していいのかすら分からん"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "ruby", "matsue" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 「ミリしら」で KeebWorld Conference 2024 へ行ってみた

最初に予防線を張っておくと，私は自作キーボードとか全く知らないのよ。
そりゃ，そういうことをする人たちがいるのは知ってるよ。
でも，それだけ。
でもまぁ，せっかく松江でやるってのなら見てみたいし，そもそも技術系のオフラインイベントが「コロナ」以降では初参加なので[^e1]，リハビリにはなるかなぁと（笑）

[^e1]: 技術系以外のイベントならある。「[島根の歴史文化講座 2024]」とか。これも今年に入ってからだけどさ。

毎年松江市内で RubyWorld Conference が開かれるが，今年はその流れで[カンファレンス][RubyWorld Conference 2024]の翌日の今日に開催された。

- [KeebWorld Conference 2024 松江](https://keebkaigi.org/2024/)

さっそく松江駅前の[オープンソースラボ](https://team-place.com/space/494 "松江オープンソースラボ | TeamPlace - 「人でつながる」ワークプレイスプラットフォーム")まで出かける。
自転車は[オーバーホール]({{< ref "/remark/2024/12/overhaul.md#overhaul" >}} "自転車をオーバーホールに出した")に出してるので公共交通機関で。

{{< fig-img src="./54188441314_fa8148d34e_e.jpg" title="今日の会場 カルト？が煩い | Flickr" link="https://www.flickr.com/photos/spiegel/54188441314/" >}}

いや，なんか，なんかワクチンがどうたらとか叫ぶカルト集団？ がビルの前に居たので，大回りの遠巻きにしながら入っていった。
マジ迷惑。
病院前も迷惑だろうし，休日にやるならお役所前でやってくれ。

開始10分前に入ったのだが，既にほぼ満席状態だった。
[RubyWorld Conference][RubyWorld Conference 2024] からの流れとはいえ，結構ビックリした。

{{< fig-img src="./54188425748_05f83ef656_e.jpg" title="KeebWorld Conference 2024 (聴衆モード) | Flickr" link="https://www.flickr.com/photos/spiegel/54188425748/" >}}

最初に述べたように自作キーボードに関しては本当に「ミリしら」状態なので，ただ聴きに来ただけという感じになってしまった。
発表に対して何を質問していいのかすら分からん（笑）

例えば，なんで [RubyWorld Conference][RubyWorld Conference 2024] の流れでこれを開催するのか，今回の一連の発表でようやく分かった。
どうやら Ruby 製の [picoruby/prk_firmware] というファームウェア環境があって，それを使って自作キーボードをカスタマイズするってのが Ruby コミュニティ内で随分前から流行ってるらしい。
そういう基本的なことも知らないレベルな私である[^gb1]。

[^gb1]: 確か [TinyGo] でキーボードのファームウェアを組むみたいな話も聞いたことがあるが，この辺の情報も全く知らない。いやぁ，組込み制御ソフトウェアに全くタッチしなくなって長いからなぁ...

以下は覚え書で（全ての発表を網羅しているわけではないのはご容赦）：

### [Have fun why not](https://docs.google.com/presentation/d/1XPRzvgzHdGFtaawGHmVlssijzvD40zISW7wQGv3uf8o/edit?usp=sharing) by [原悠](https://yhara.jp/) ([@yhara](https://x.com/yhara))

自作キーボードについてのガイダンス的な発表。
60%とか40%とか30%とかいうキーボードの分類の仕方を初めて知った（たとえば有名な [HHKB](https://happyhackingkb.com/ "Happy Hacking Keyboard Microsite | PFU") は60%）。
特に面白かったのは，キーボードの話にはあまり関係ないが

{{< fig-img-quote src="./have-fun-why-not.png" title="Have fun why not" link="https://docs.google.com/presentation/d/1XPRzvgzHdGFtaawGHmVlssijzvD40zISW7wQGv3uf8o/edit?usp=sharing" width="1263" >}}

プログラミング言語を包丁に喩えて，普通は切れ味や軽さや丈夫さといった点で評価するが Ruby は「持ち手がすべすべだと嬉しいよね」的な評価になりがちってのがめがっさウケた。

考えてみれば [Go] 言語も（Ruby とはある意味真逆なのに）似たようなもんだと思ったり。

### [キーボードからロボットへ](https://github.com/kurod1492/keebworldconf-2024/blob/main/slide.md) by [Akihiro Kurotani](https://qiita.com/kurod1492) ([@kurod1492](https://x.com/kurod1492))

次に紹介する発表に絡むが，ファームウェアのカスタマイズ（書き換え）が容易という話は興味深かった。

自作キーボードと関係ないが，来年 2024-01-12 に[松江テルサ]で「[Matz葉がにロボコン](https://www.shimane-oss.org/kani-robo/ "Matz葉がにロボコン|かにロボ連盟 ご当地こども向けプログラミングコンテスト")」ってのが開催されるらしい。
でも関係者以外の見学は無理そう？

### ["Actual" Security in Microcontroller Ruby!?](https://speakerdeck.com/sylph01/actual-security-in-microcontroller-ruby) by [sylph01](https://s01.ninja/) ([@s01](https://x.com/s01))

自作キーボードの制御に使われているらしい [Raspberry Pi Pico] および [Raspberry Pi Pico 2] では WiFi 接続ができるそうな。
しかもファームウェアである [picoruby/prk_firmware] も WiFi 制御に対応してるんだと。
さらに [Raspberry Pi Pico 2] のほうは SRAM がほぼ倍（520MB）になったそうで，単純な制御ではなくもっと色々なことができるようになったわけだ。
しかも前節にあるようにファームウェアの書き換えは容易（スクリプトキディには無理かもしれないが）。

これで何が起きるかというと，自作キーボード経由で不正アクセスがしやすくなるということだ。
これを防ぐのが Secure Boot や Encryptd Boot らしい。

まぁ，キーボードが programmable になれば当然そうなるよな。
面白い！

### わたしのキーボード by [まつもとゆきひろ](https://matz.rubyist.net/) ([@yukihiro_matz](https://x.com/yukihiro_matz))

スライドの公開はないのかな。
うろ覚えですまん。

欲しいキーボードのスペックは決まってるけど，完成品でも自作（金に飽せて作ってもらったらしい）でも満足するものがなくて，キーボードを買いまくってる話が面白かった。
カンファレンスのあと（私は参加してないが）使わなくなったキーボードを放出したそうな。

キーレイアウトは日本語キーボードが欲しいんだけどマッピングは US 配列がいいらしい。
理由は US キーボードはキーの数が少ないから。

さらにファンクションキーは独立したものが欲しいらしい。
これに関しては同意する。
ファンクションキーは結構使うし，コンビネーション・キーと組み合わせたキーアサインは少ないほういいと私も思う。

### 30%キーボード発想法 by [みなも♨️30%](https://scrapbox.io/self-made-kbds-ja/minamo) ([@X___MOON___X](https://x.com/X___MOON___X))

これもスライドの公開はないのかな。
うろ覚えですまん。

30%キーボードってのは概ね以下のようなキーで構成されているキーボードらしい。

{{< fig-img-quote src="./have-fun-why-not-2.png" title="Have fun why not" link="https://docs.google.com/presentation/d/1XPRzvgzHdGFtaawGHmVlssijzvD40zISW7wQGv3uf8o/edit?usp=sharing" width="1230" >}}

たとえば60%だとキーの配列の自由度がなくなって，どれも似たようなレイアウトになるらしい。
せいぜい分割にするとか。
キーの数を極端に減らすことでタイプライタ由来のレイアウトの軛から外れ自由な設計ができる，という話が興味深かった。

## 感想

[カンファレンス][KeebWorld Conference 2024]では自作キーボードも色々展示されていたが，あまり見てない。

{{< fig-img src="./54188541799_2b0851cf6c_e.jpg" title="KeebWorld Conference 2024 | Flickr" link="https://www.flickr.com/photos/spiegel/54188541799/" >}}

やっぱ分割は基本なんだな。
まぁ分割しないでファームウェアのカスタマイズも要らないならメーカーから完成品を買えばいいだろうし。

私は若い頃にガチの VT 端末で「現地調整」させられてたことがあって US キー配列でも日本語キー配列でもどっちでもいけるのよ。
でも60%キーボードは当時のトラウマがフラッシュバックするので使いたくない（あと vi/vim もトラウマがフラッシュバックするw）。
それに「わたしのキーボード」でも言及されていたが，コンビネーション・キーと組み合わせたキーアサインは少ないほうがいい。
そもそもメカニカルキーボード自体を使い出したのが[ごく最近]({{< ref "http://localhost:1313/remark/2024/11/mechanical-keyboard.md" >}} "はじめてのメカニカルキーボード")なんだよな。

あと手先が不器用なのでハンダ付けとかも駄目。
私の場合，パーツ数の少ないガンプラを組み立てるくらいが限界である。
母親は手先が器用でハンダ付けも得意なのに，その辺は遺伝してくれなかったようだ（不器用さは親父似）。
もしやるなら金に飽せて得意な人にやってもらうって感じになるだろうか。

まぁ，しばらくはキーボードには手を出すまいと思った。
ファームウェアをいじるのは面白そうだし「キーボードガチ勢でもないのに全部他人の仕事でごそごそやっていたら新聞に載ってしまった件 (by [東裕人](https://github.com/HirohitoHigashi))」で紹介されているような特殊用途のキーボードとかなら興味あるけど。

広島にいた頃から Ruby 案件には恵まれてなくて，仕事で全く使ったことがない。
仕事以外では C/C++ か Java か [Go] で，今は [Go] 一辺倒なので，松江地元にいてこれでいいのか？ という気はチョットしている。
一方で，歳とってから Ruby かぁ，という気もしないではないが。

あと Ruby コミュニティでは SNS は {{% emoji "X" %}} が主流なのかね。
Bluesky とかで見かけない感じ？ まぁ，今回の[カンファレンス][KeebWorld Conference 2024]は Bluesky の TL で知ったんだけど。
Mastodon は他サーバのユーザは探しにくいので動向が分からない。

タイムテーブルをちゃんと見てなくて「日没までには帰ろう」とか呑気に考えてたら17時半過ぎまでみっちりだった。
逢魔時のバス停で帰りのバスを待っていたら号外が配られているのに気づく。

{{< fig-img src="./54188796995_09454faa67_e.jpg" title="号外！ | Flickr" link="https://www.flickr.com/photos/spiegel/54188796995/" >}}

おー。
いよいよか。
新聞記者さんに捕まってインタビューを求められたが断った。
私は数年前から政治的無関心を装う（特に政治的なポジショントークはしない）ことを年間目標にしてるので（ときどき胡乱なことを口走るのはご容赦），応えないほうがいいだろう（笑）

## ブックマーク

- [はじめてのメカニカルキーボード]({{< ref "http://localhost:1313/remark/2024/11/mechanical-keyboard.md" >}})

[KeebWorld Conference 2024]: https://keebkaigi.org/2024/ "KeebWorld Conference 2024 松江"
[RubyWorld Conference 2024]: https://2024.rubyworld-conf.org/ "RubyWorld Conference 2024"
[picoruby/prk_firmware]: https://github.com/picoruby/prk_firmware "picoruby/prk_firmware: A keyboard firmware platform in PicoRuby"
[Raspberry Pi Pico]: https://www.raspberrypi.com/products/raspberry-pi-pico/ "Buy a Raspberry Pi Pico – Raspberry Pi"
[Raspberry Pi Pico 2]: https://www.raspberrypi.com/products/raspberry-pi-pico-2/ "Buy a Raspberry Pi Pico 2 – Raspberry Pi"
[Go]: https://go.dev/
[TinyGo]: https://tinygo.org/ "TinyGo"
[松江テルサ]: https://www.matsue-terrsa.jp/ "松江テルサ"
[島根の歴史文化講座 2024]: https://shimane-kodaibunka.jp/sympo/sympo-3424/ "島根の歴史文化講座　2024 | 島根県古代文化センター"

## 参考

{{% review-paapi "B0CJT7S6D6" %}} <!-- テンキーレス キーボード メカニカル -->
