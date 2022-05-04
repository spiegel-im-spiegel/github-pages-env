+++
title = "週末スペシャル： まじめに規制に従っている人ほど馬鹿を見る社会"
date = "2016-04-10T18:44:29+09:00"
description = "まじめに規制に従っている人ほど馬鹿を見る社会 / Linux サブシステムは Windows の終わりの始まり / 鍵管理システム CONIKS / Go 言語を使うようになって変わったこと / その他の気になる記事"
tags = ["code", "security", "risk", "management", "cryptography", "windows", "linux", "messaging", "pki", "engineering", "golang", "grigori"]

[scripts]
  mathjax = false
  mermaidjs = false
+++

3月は去りました。
春になっちゃったよ。

うっかり左手首を痛めてしまった（疲労がたまるとたまになる）のでいろいろ控えてた。
溜まりまくった小ネタを消化しないと。

1. [まじめに規制に従っている人ほど馬鹿を見る社会]({{< relref "#code" >}})
1. [Linux サブシステムは Windows の終わりの始まり]({{< relref "#bash" >}})
1. [鍵管理システム CONIKS]({{< relref "#pki" >}})
1. [Go 言語を使うようになって変わったこと]({{< relref "#go" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## まじめに規制に従っている人ほど馬鹿を見る社会{#code}

もう何度も書いているが「警察にできることは犯罪者にもできる」。
問題は犯罪者にできることが警察にもできるかどうか駄菓子菓子...

- [米政府によるスマホデータ取り出しの協力要請、ACLUが実態調査 - CNET Japan](http://japan.cnet.com/news/society/35080404/)
- [FBIのiPhoneロック解除方法、Appleに知らされない可能性も (1/3) - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/01/news114.html)
- [FBI長官、「購入したロック解除ツールはiPhone 5sでは機能しなかった」 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/08/news060.html)
- [司法省がまたAppleにiPhoneアンロック要求、今度はAppleが“相手を間違えた”国を訴訟か | TechCrunch Japan](https://jp.techcrunch.com/2016/04/09/20160408justice-department-keeps-pushing-apple-to-unlock-iphone-in-new-york-drug-case/)

FBI が端末を突破するのに外部企業を使ったということ，そして企業がそれに応じたことは重要だ。
もちろん実は NSA の息のかかった企業だった，としても驚かないけど。

企業は利があると思えば警察にも犯罪者にだって加担する。
今回の件のポイントは「犯罪者にできることが警察にできるとは限らない」と証明してしまったことだ。
セキュリティ企業は新しい時代の「死の商人」になるかもしれない。

警察が優位に立てるのは犯罪者よりもパワー（暴力・権力を含む）を有している場合のみである。
コンピュータ・ネットワーク技術あるいは暗号技術において政府・警察は優位に立てない。
米国司法省は法規制によって優位に立てると思ってるようだが，こんなもの最初から「法の外」にいる犯罪者やテロリストに対しては効力がない。

- [暗号化解除をめぐる米法案、司法当局へのバックドア提供を義務付け - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/09/news022.html)
- [バックドア提供を拒む企業に制裁金を--米国で法案が公開 - CNET Japan](http://japan.cnet.com/news/society/35080962/)

これは「飲酒運転を減らすために飲酒運転規制を厳罰化する」というのとは話が違う。
犯罪者にはインパクトがないし，まじめに規制に従っている人ほど「馬鹿を見る」ことになる。

有害なルールに従う必要はないし，それに従うことはむしろリスクを高めることになる。

## Linux サブシステムは Windows の終わりの始まり{#bash}

- [Build 2016で驚きの発表―Microsoftはこの夏Windows 10でBashシェルをサポート | TechCrunch Japan](https://jp.techcrunch.com/2016/03/31/20160330be-very-afraid-hell-has-frozen-over-bash-is-coming-to-windows-10/)
- [「Windows 10」で動作するUbuntuのBashシェル--その実現方法 - CNET Japan](http://japan.cnet.com/news/service/35080406/)
- [開発者がWindows 10でBashシェルとユーザー モードのUbuntu Linuxバイナリを実行可能に | S/N Ratio (by SATO Naoki)](https://satonaoki.wordpress.com/2016/03/31/bash-ubuntu-windows/)
- [MariaDB、カラム型データベースエンジン「MariaDB ColumnStore」発表。OLAPへ参入 － Publickey](http://www.publickey1.jp/blog/16/mariadbmariadb_columnstoreolap.html)

~~もともと Windows は POSIX サブシステムを持っている。
今回はそれに加えて~~ Ubuntu ベースの Linux サブシステムを組み込むということらしいが子亀の上に親亀を乗っけるようなものだ。

Windows の基本的な設計思想は20～25年くらい前の古いものだ。
しかも DOS/Windows はもともとシングルユーザ用に設計されたもので UNIX 等のマルチユーザ向けの OS とは全く異なる。

Linux のベースとなっている UNIX もそうとう古いが，マルチユーザを前提とした考え方は今でも通用するし，なにより Linux はもはや UNIX に縛られない。

- [Linux公開25周年を受けたリーナス・トーバルズのインタビュー - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20160331/linux25years)
- [Linux創始者トーバルズ氏、IoTを語る--「セキュリティは二の次」と警鐘 - ZDNet Japan](http://japan.zdnet.com/article/35080722/)

Windows は永遠に Windows に縛られ続ける。
Microsoft が満を持して出した Windows 10 も結局は Windows に縛られている。

Windows が時代遅れなのは明らかである。
Microsoft 自らこういう無茶をすること自体が「Windows の終わりの始まり」だ。
個人的に2020年までに自宅 PC のメインを Linux 機に換装する予定だが，ちょっと計画を前倒ししたほうがいいかもしれない。

- [目的別のおすすめLinuxディストリビューション--あなたにぴったりなのはどれ？ - ZDNet Japan](http://japan.zdnet.com/article/35080364/)

### 追記

- [Big Sky :: Windows ユーザは cmd.exe で生きるべき。](http://mattn.kaoriya.net/software/why-i-use-cmd-on-windows.htm)

激しく同意。
もっとも私は [ConEmu & NYAGOS]({{< ref "/remark/2015/conemu-and-nyagos.md" >}}) だけど（笑）

## WhatsApp がついに Signal ベースの E2E 暗号化を実装する{#sig}

- [Open Whisper Systems >> Blog >> WhatsApp's Signal Protocol integration is now complete](https://whispersystems.org/blog/whatsapp-complete/)
- [WhatsApp completes end-to-end encryption rollout | TechCrunch](http://techcrunch.com/2016/04/05/whatsapp-completes-end-to-end-encryption-rollout/)
- [Facebook傘下のWhatsApp、完全暗号化を完了　「政府もわれわれも解除できない」 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1604/06/news069.html)
- [WhatsApp、全てのプラットフォームのエンドツーエンド暗号化を完了 | TechCrunch Japan](https://jp.techcrunch.com/2016/04/06/20160405whatsapp-completes-end-to-end-encryption-rollout/)

もともと WhatsApp が Signal ベースの暗号化システムを実装することは予告されていた。

- [Open Whisper Systems](https://whispersystems.org/) ([GitHub](https://github.com/WhisperSystems))
    - [Is it private? Can I trust it? – Support](http://support.whispersystems.org/hc/en-us/articles/212477768-Is-it-secure-Can-I-trust-it-)

Signal 自体は SMS アプリを置き換えることのできる優れたアプリなのだが SNS ベースのメッセンジャー・アプリとしては機能的に劣る。
WhatsApp がその辺を埋めることになるかどうか。
でも日本のユーザにはウケないかなぁ。

メールは ProtonMail， SMS ベースのメッセンジャーには Signal，それ以外のメッセンジャーには WhatsApp と，だいぶ揃ってきたねぇ。

## 鍵管理システム CONIKS{#pki}

- [CONIKS](https://coniks.cs.princeton.edu/)
- [CONIKS - Schneier on Security](https://www.schneier.com/blog/archives/2016/04/coniks.html)

{{< fig-quote title="CONIKS" link="https://coniks.cs.princeton.edu/" lang="en" >}}
<q>CONIKS is a key management system for end users capable of integration in end-to-end secure communication services. The main idea is that users should not have to worry about managing encryption keys when they want to communicate securely, but they also should not have to trust their secure communication service providers to act in their interest.</q>
{{< /fig-quote >}}

とりあえずメーリング・リストに入ってみた。

## Go 言語を使うようになって変わったこと{#go}

内容自体にさほど文句があるわけではないが（細かい部分は置いておいて），「interface を中心に設計する」という記述が気になって。

<script async class="speakerdeck-embed" data-id="947e9a6ef68c4310baf21afdec4fcfab" data-ratio="1.33333333333333" src="//speakerdeck.com/assets/embed.js"></script>

私はそんなにたくさんの言語を知っているわけではないが， [Go 言語]を勉強するようになって設計の考え方が少し変わった。
まさに「制約は構造を生む」（by 結城浩「数学ガール」シリーズより）が如く，言語仕様によって思考も影響を受けるのである。
以下にいくつか例を挙げよう。

### Value Object から考える

さて，いつもの図。

{{< fig-img src="./DDD.svg" title="Domain-Driven Design" link="./DDD.svg" width="640" >}}

Domain Layer の中身は Domain Service, Entity, そして Value Object に分類される。
ビジネスロジックは図の右側，つまり Entity や Value Object に記述されるのが良い設計だと言われている（記述の重複を避けられるため）。
ただし Value Object はしばしば省略されることが多い。

Value Object の特徴は以下のとおり。

- 内部状態を持たず不変である
- 属性（property）の比較のみでオブジェクト同士が等価かどうか決定できる

そして実装上の要件としては「軽量」であることが求められる。
何故なら Value Object は Entity の属性として使われることが多く Value Object がボトルネックになるとシステム全般へのインパクトが大きいからだ。

実は [Go 言語]はこの Value Object の実装にとても向いている。

- [Go 言語における「オブジェクト」 — プログラミング言語 Go]({{< ref "/golang/object-oriented-programming.md" >}})
- [関数とポインタ — プログラミング言語 Go]({{< ref "/golang/function-and-pointer.md" >}})

[Go 言語]の特徴である「強い型付け」も Value Object を念頭に置いて考えるなら合理的な仕様であることが分かるだろう。

### 多態性を「振る舞い」から考える

[Go 言語]の多態性（polymorphism）は振る舞いによってのみ規定される（構造的部分型）。
つまり「猫」のように振る舞うのであれば実体がロボットだろうがコスプレイヤーだろうが全部「猫」として括れるのである。
そして「猫」のようにあるためにロボットやコスプレイヤーの identity を書き換える必要はない。
これはとても重要な事である。

たとえば「猫」を実装する際に，それに多態性を持たせなければならないかどうかは設計の割と早い段階で決めなければならないことが多い。
そうして先に `interface` などを決めなければ具体的なクラスを記述することができない。
しかし [Go 言語]ではアプローチが逆になる。
先にロボットやコスプレイヤーといった具体的な型（[type]）をバンバン作り，個々の振る舞いを見て，あとから「あっ，これ「猫」で括れるぢゃん♥」となるわけだ。
言い方を変えるなら refactoring 向きであるとも言える。

### 要件定義からコードを書く

これは [Go 言語]に限らないが， refactoring しやすい言語は prototyping に向いている言語であるとも言える。
Prototyping に向いているということはプロジェクトのかなり早い段階（たとえば要件定義）からコードを書けるということでもある。
結局エンジニアにとって信用できるのは百万語を連ねた設計書より「動くコード」なのである。

## その他の気になる記事{#other}

- [Transmission Releases Long-Awaited BitTorrent Client For Windows - TorrentFreak](https://torrentfreak.com/transmission-releases-long-awaited-bittorrent-client-for-windows-160327/)
- [NPMとleft-pad : 私たちはプログラミングのやり方を忘れてしまったのか？ | プログラミング | POSTD](http://postd.cc/npm-and-left-pad/)
- [WindowsにBitTorrentクライアントの決定版Transmissionがやってくる | TechCrunch Japan](https://jp.techcrunch.com/2016/03/29/20160328windows-users-finally-have-a-good-bittorrent-client/)
- [IIJ、Webサイトにおけるユーザ認証のセキュリティを強化する 「IIJ SmartKeyマネージメントサービス」を提供開始 | 2016年 | IIJ](http://www.iij.ad.jp/news/pressrelease/2016/0329-2.html)
- [IPAテクニカルウォッチ「公衆無線LAN利用に係る脅威と対策」：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/technicalwatch/201600330.html)
- [高度サイバー攻撃(APT)への備えと対応ガイド～企業や組織に薦める一連のプロセスについて](https://www.jpcert.or.jp/research/apt-guide.html)
- [著作権削除要請の28％が「疑わしい」との研究結果 – P2Pとかその辺のお話R](http://p2ptk.org/copyright/231)
- [国立極地研究所情報図書室、ウェブサイトをCC BYで公開 | カレントアウェアネス・ポータル](http://current.ndl.go.jp/node/31200)
- [Ｘ線天文衛星「ひとみ」（ASTRO-H）の状況について - 20160408_hitomi.pdf](http://fanfun.jaxa.jp/jaxatv/files/20160408_hitomi.pdf)
    - [X線天文衛星「ひとみ」、回転は破片を誤認？米軍発表 | Sorae.jp : 宇宙（そら）へのポータルサイト](http://sorae.jp/030201/2016_04_02_jspoc.html)
- [定時帰宅のススメ — Medium](https://medium.com/@tsukamoto/-f42bf7b5e25e)
- [SpaceXのFalcon 9ロケット、洋上のドローン艀への軟着陸についに成功 | TechCrunch Japan](https://jp.techcrunch.com/2016/04/09/20160408spacex-just-landed-a-rocket-on-a-drone-ship-for-the-first-time/)
- [GoogleがTLSでの採用を提唱している共通鍵暗号方式「ChaCha」についてまとめた - sonickun.log](http://sonickun.hatenablog.com/entry/2016/04/03/183220)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[type]: https://go.dev/ref/spec#Properties_of_types_and_values "Properties of types and values"

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
