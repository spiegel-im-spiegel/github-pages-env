+++
title = "劇場型セキュリティと PPAP"
date =  "2020-06-25T18:05:01+09:00"
description = "つか，今だにこんなことやってる企業あるの？ ホンマに？"
image = "/images/attention/kitten.jpg"
tags = [ "security", "cryptography", "risk", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近セキュリティ関連で PPAP なる謎の隠語を聞くのだが，どうも[「情報処理」2020年7月号]の特集記事が発端らしい？

命名の元ネタはもちろんアレである（笑）

{{< fig-youtube id="0E00Zuayv9Q" title="PPAP（Pen-Pineapple-Apple-Pen Official）ペンパイナッポーアッポーペン／PIKOTARO(ピコ太郎) - YouTube" >}}

PPAP については特集記事を見ていただくとして，簡単にいうと，暗号化 zip ファイルとその復号パスワードを相手に電子メールで送りつける日本独特の商慣行（笑）を揶揄した言葉のようだ。

つか，今だにこんなことやってる企業あるの？ ホンマに？

## 劇場型セキュリティ（Security Theater）

[「情報処理」2020年7月号]でも説明されているが， security theater というのはゼロ年代初頭に起きた同時多発テロ（9.11）を受けて打ち出されたセキュリティ対策の間抜けさ加減を指した言葉で， Bruce Schneier さん著作の『[セキュリティはなぜやぶられたのか]』にも登場する。

{{< fig-quote type="markdown" title="セキュリティはなぜやぶられたのか" link="https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
{{% quote %}}セキュリティは、ある程度、どう感じるかという問題だと第1章で説明した。そうであるなら、実体だけでなく守られているという感覚も得られる対策が望ましい。しかし実際には、実体がなくて感覚だけの対策もある。それでは単なる**セキュリティ芝居**である。言い訳にすぎないのだ{{% /quote %}}。
{{< /fig-quote >}}

最近でも COVID-2019 絡みで

{{< fig-quote type="markdown" title="The Public Is Being Misled by Pandemic Technology That Won’t Keep Them Safe" link="https://onezero.medium.com/the-public-is-being-misled-by-pandemic-technology-that-wont-keep-them-safe-1966ed740a87" lang="en" >}}
{{% quote %}}As scientists and health care professionals in general agree, thermal imaging systems are an inappropriate and ineffective strategy for identifying individual cases of Covid-19 and preventing its spread. Beyond the wasted costs, installing thermal scanning as part of a “back to work” process might create a false sense of security that convinces some to prematurely return to their jobs and emboldens others to relax more effective strategies, such as social distancing and responsible contact tracing efforts{{% /quote %}}.
{{< /fig-quote >}}

なんてな話もある。
PPAP とか典型的な security theater と言えるだろう。

## 暗号化 zip とか使うな！

そもそもの話として，暗号化 zip ファイルを解読するツールはそこらに普通に出回ってるけどね（笑）

まともな暗号化圧縮が使いたいなら [gpgtar] を使いましょう。

- [複数ファイルをまとめて署名・暗号化する]({{< ref "/openpgp/gpgtar.md" >}})

Windows 環境に限定するなら「[アタッシェケース](https://hibara.org/software/attachecase/ "アタッシェケース#3 | HiBARA Software")」等を使う手もあり（パスワード配布の問題は残るけど）。

## 電子メールとか使うな！

電子メールはよく葉書に喩えられる。
例として，年賀状に厨二病的な文言を書いて相手に「それ」が渡るまで，どれだけ見られる可能性があるか想像したら軽く死ねるだろう。

いわんや，国家権力をや，である。

大昔，電子メールは重要なメッセージング基盤だった。
だから PGP をはじめとして安全でない電子メールの内容をいかに安全に相手に届けるか，を真剣に考える必要があった。
しかし今日においては，苦労して電子メールなんか使わなくても秘密を安全に配信する手段はあるのだ。

## むしろ電子署名をちゃんとしようよ

私は [GnuPG] を日常的に使ってるが `git commit` 時の電子署名が主な使い道である。
職業エンジニアだった頃も機密文書の共有は Box や Dropbox 等のクラウド・ストレージを使うのが当たり前だったし，そこで暗号化が必要なら [LibreOffice] 文書に [GnuPG] で署名・暗号化すればよかった。

- [Git Commit で OpenPGP 署名を行う]({{< ref "/openpgp/git-commit-with-openpgp-signature.md" >}})
- [LibreOffice と OpenPGP （仕切り直し）]({{< ref "/openpgp/libreoffice-with-openpgp.md" >}})

大袈裟な基盤は必要なく [OpenPGP] の peer-to-peer で緩やかな信用モデルの下に運用できていれば通常業務には十分である。

- [OpenPGP の電子署名は「ユーザーの身元を保証し」ない]({{< ref "/openpgp/web-of-trust.md" >}})

最近

- [PDFベースの共同編集・電子署名サービスのAnvilがグーグル系VCから5億円超を調達  |  TechCrunch Japan](https://jp.techcrunch.com/2020/06/04/2020-06-03-paperwork-automation-platform-anvil-raises-5-million-from-googles-gradient-ventures/)
- [電子署名サービスを今夏提供へ--ドロップボックスが2020年の事業戦略 - ZDNet Japan](https://japan.zdnet.com/article/35155164/)

のようなニュースを見かけるようになり「リモートワークがメインになれば要るよなぁ」と思った一方で，よりにもよって「情報処理」みたいな雑誌が PPAP がどうのとか言ってるのを見るとガッカリしてしまう。
日本の企業はどこまで周回遅れなのか。

もはや暗号化なんてどうとでもできるけど，ゼロトラスト云々な時代なら「そのデータは本当に『正しい』よね？」ってのがより重要なんじゃないのか。

とっとと次に行こうぜ。

[「情報処理」2020年7月号]: https://www.amazon.co.jp/dp/B089N3VX86?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "情報処理 2020年7月号 | 情報処理学会 | 科学・テクノロジー | Kindleストア | Amazon"
[セキュリティはなぜやぶられたのか]: https://www.amazon.co.jp/dp/4822283100?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "セキュリティはなぜやぶられたのか | ブルース・シュナイアー, 井口 耕二 |本 | 通販 | Amazon"
[OpenPGP]: http://openpgp.org/
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[gpgtar]: https://www.gnupg.org/documentation/manuals/gnupg/gpgtar.html "Using the GNU Privacy Guard: gpgtar"
[LibreOffice]: https://www.libreoffice.org/ "LibreOffice - Free Office Suite - Fun Project - Fantastic People"

## 参考図書

{{% review-paapi "B089N3VX86" %}} <!-- 情報処理 2020年7月号 -->
{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
