+++
title = "Rust 使いの日英文学翻訳者による辞書ツール"
date =  "2026-11-08T12:28:29+09:00"
description = "売りは Linux GUI 環境で動くこと，検索が速いこと，何より完全にオフラインで使えること"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "tools", "linux", "language" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Linux 特に Ubuntu の情報でお世話になっている [kledgeb](https://kledgeb.blogspot.com/ "kledgeb") の記事で知ったのだが

- [Linux その456 - GNOME 向けアプリとアップデートの紹介・日英辞書アプリや PDF 編集アプリなど](https://kledgeb.blogspot.com/2026/11/linux-456-gnome-pdf.html)

[Kotoba] という日英辞書ツールがあるらしい。
Flatpak で[インストールできる](https://flathub.org/ja/apps/net.trowell.kotoba "KotobaをLinuxにインストール | Flathub")みたいなので，さっそく試してみた。

```text
$ flatpak install flathub net.trowell.kotoba
Looking for matches…
Remote ‘flathub’ found in multiple installations:

   1) system
   2) user

Which do you want to use (0 to abort)? [0-2]: 2
...
```

起動直後はこんな感じ。

{{< fig-img-quote src="./kotoba-01.png" title="Kotoba (1)" link="./kotoba-01.png" width="969" lang="en" >}}

入力は日本語や英語のワードで行うことができる。
仮名やローマ字もOK。
ワイルドカード `*` も使える。

たとえば「箱根」と入力すると，関連するワードが山ほど表示される。

{{< fig-img-quote src="./kotoba-02.png" title="Kotoba (2)" link="./kotoba-02.png" width="969" lang="en" >}}

"Artificial Intelligence" だとこんな感じ。

{{< fig-img-quote src="./kotoba-03.png" title="Kotoba (3)" link="./kotoba-03.png" width="969" lang="en" >}}

ローマ字綴りで "abikyoukan" と入力すると，ちゃんと「阿鼻叫喚」が出る。

{{< fig-img-quote src="./kotoba-04.png" title="Kotoba (4)" link="./kotoba-04.png" width="969" lang="en" >}}

おおっ。
ちゃんと仏教由来の言葉だと説明があるな[^ak1]。
さすがに「[あえんびえん](https://dic.pixiv.net/a/%E3%81%82%E3%81%88%E3%82%93%E3%81%B3%E3%81%88%E3%82%93 "あえんびえん (ただしくはあびきょうかん)とは【ピクシブ百科事典】")」では出てこなかった（笑）

[^ak1]: 阿鼻叫喚は仏教で言う八大地獄の「阿鼻」と「叫喚」を併せた言葉で，井伏鱒二の『黒い雨』で原爆投下時の情景を表す言葉として有名。

秀逸なのは部首などから漢字を特定できる点。
たとえば「霎[^k1]」の意味も読みも分からないとしても，「雨」「立」「女」で構成されていると分かれば，そこから漢字を探すことができる。

[^k1]: 「霎」は「ショウ」または「こさめ」と読む。ひとしきり降る雨。いわゆる小雨のこと。派生語として「霎時（しょうじ; しばらくの間）」といった言葉もある。

{{< fig-img-quote src="./kotoba-05.png" title="Kotoba (5)" link="./kotoba-05.png" width="969" lang="en" >}}

{{< fig-img-quote src="./kotoba-06.png" title="Kotoba (6)" link="./kotoba-06.png" width="969" lang="en" >}}

「霎」では検索結果に何も出てこなかったけどね（笑）
まぁ，個人で作られているツールのようなので網羅性に関しては仕方がないかと。

英単語については注意が必要。
たとえば "compose" は出てくるが "composable" は出てこない。
複数形についても，たとえば "pronouns" だと出てこない。
ちなみに "enshittification" は出てきた。

{{< fig-img-quote src="./kotoba-07.png" title="Kotoba (7)" link="./kotoba-07.png" width="969" lang="en" >}}

なんでやねん（笑）

[Kotoba] の売りは Linux GUI 環境で動くこと，検索が速いこと，何より完全にオフラインで使えることだろう。
運用でクラウドに依存しないツールはイマドキ貴重だからね。

作者の Haydn さんについても簡単に調べてみた。
いや，調べてくれたのは [Kagi Assistant] だけど。
オーストラリア出身で，どうやら本業は日本文学の翻訳者らしい。

{{< fig-quote type="markdown" title="About - Haydn Trowell" link="https://haydntrowell.com/#about" lang="en" >}}
Haydn Trowell is an Australian translator of Japanese literature whose work has appeared in Granta, Electric Literature, and Words Without Borders. He has translated fiction by major prize-winning authors such as Yasunari Kawabata, Maki Kashimada, and Rie Qudan, and he has worked extensively with publishers in Australia, Japan, the UK, and the US. Haydn holds a PhD in translation studies from Monash University and has a background in both translation and creative writing. He delivers expertly crafted English renditions that capture the author’s unique voice and preserve the rich cultural nuances of the original, making each work feel authentic, engaging, and completely at home in its new language and environment.
{{< /fig-quote >}}

私が文学作品に疎いだけかもしれないが，日本語圏での言及はほとんど見当たらない感じ。
ざっと調べてみたところ，9月に出た『[別冊GUNZOU](https://link.amazon/B0hZQvIZy)』には，彩瀬まるさんの「私たちは星」の翻訳が掲載されていることを確認できた。

翻訳者で Rust 使いかぁ。
面白い！

Haydn さんが公開しておられる他の製品としては Typst 用のエディタ（[Typesetter](https://gitlab.gnome.org/haydn/typesetter "Haydn / Typesetter · GitLab")）や画像データから文字情報を抽出する多言語 OCR ([Mojitori](https://gitlab.gnome.org/haydn/mojitori "Haydn / Mojitori · GitLab")) がある。
いずれもオフラインで動作可能らしい。
OCR はそのうち試してみるかな。

[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[Kotoba]: https://gitlab.gnome.org/haydn/kotoba "Haydn / Kotoba · GitLab"
