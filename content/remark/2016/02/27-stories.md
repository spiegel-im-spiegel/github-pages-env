+++
date = "2016-02-27T22:31:08+09:00"
update = "2016-03-14T22:46:02+09:00"
description = "3月9日は皆既日食（日本では部分日食） / サーバサイド CMS はリスクか / 「中高生が知っておきたいサイバーセキュリティ」とは"
draft = false
tags = ["astronomy", "eclipse", "security", "risk", "management", "cms"]
title = "週末スペシャル： 3月9日は皆既日食（日本では部分日食）"

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
  url = "https://baldanders.info/profile/"
+++

1. [3月9日は皆既日食（日本では部分日食）]({{< relref "#eclipse" >}})
1. [サーバサイド CMS はリスクか]({{< relref "#cms" >}})
1. [「中高生が知っておきたいサイバーセキュリティ」とは]({{< relref "#security" >}})

## 3月9日は皆既日食（日本では部分日食）{#eclipse}

- [3月9日は部分日食 | 国立天文台(NAOJ)](http://www.nao.ac.jp/news/topics/2016/20160224-eclipse.html)

来月3月9日は東南アジアやオーストラリアを中心に皆既日食がある。
残念ながら日本は皆既食帯から完全に外れているため部分日食となる。
父島でも食分 0.48 程度らしいので，期待せず楽しみましょう。

ちなみに，当然のことながら太陽を直接見ないこと。
レンズ越しとかもっての外。
文房具の下敷き越しもダメだよ。

オススメは木漏れ日とかかな。

{{< fig-img src="https://farm8.staticflickr.com/7072/7239249856_bda0f91b27.jpg" title="Eclipse Shadows" link="https://www.flickr.com/photos/8106459@N07/7239249856/" >}}

## サーバサイド CMS はリスクか{#cms}

JPCERT/CC からアラートが出ている。

- [改ざんの標的となるCMS内のPHPファイル(2016-02-25)](https://www.jpcert.or.jp/magazine/acreport-cms.html)

改竄された PHP ファイルにより不正コードが挿入されるらしい。
これ自体は別に珍しいことではないが， WordPress をはじめ主要な複数の CMS (Content Management System) がターゲットになっているのが特徴である。

JPCERT/CC でも PHP ファイルの改竄手法は掴んではいないらしい。
割と無差別な感じだし脆弱性のあるサーバを機械的に攻めているのかもしれない（そう見せかけている可能性もあるが）。
そうなら簡単なんだけどねぇ。

EC サイトのようにサーバサイドに機能を乗せざるを得ない場合は別だが，たかがブログ程度でサーバサイドに CMS を設置するのはリスクしかないんじゃないだろうか。
問題は「改竄されたファイル」の検出が難しいことなので， CI (Continuous Integration)  を使ったデプロイ・プロセスでそれを検出・排除できれば，かなりいけるんじゃないかと思うんだけど，どうだろう。

### 追記

- [ユーキャンが Movable Type クラウド版を使う理由 - 導入事例 | シックス・アパート - CMSソフトウェア、サービスを提供](http://www.sixapart.jp/business/u-can.html)

株式会社ユーキャンでは本番環境にサーバサイドプログラムを置かないポリシーなんだそうだ。
そこでクラウド版 Movable Type のサーバー配信機能を使い，ステージング環境と本番環境を分離して運用しているらしい。
これはこれで賢い選択である。

## 「中高生が知っておきたいサイバーセキュリティ」とは{#security}

いや，これダメなんじゃないかなぁ。
脅威を煽るだけで中途半端な対策しか示さない。
これじゃあただの FUD (Fear, Uncertainty and Doubt) だよ。

- [有料アプリがなぜか無料で落とせる、ウイルスだった――中高生が知っておきたいサイバーセキュリティのマンガ、NISCが公開 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160225_745252.html)
- {{< pdf-file title="マンガで学ぶサイバーセキュリティ" link="http://www.nisc.go.jp/security-site/files/CSmanga_JPN.pdf" >}}

### スマートフォンのセキュリティ

例示としてはいわゆる ransomware に引っかかり「身代金」を払ったものの事態は改善しなかったという，まぁ今時ありがちな話。
対処としては「ウイルス対策ソフトの導入」「信頼できるサイトからのアプリの導入」「OS やアプリを最新バージョンに保つ」というもの。

しかし相手は「中高生」である。
PC 用では個人向けに無料のものもあるが，携帯端末用の「ウイルス対策ソフト」の多くは有料（ランニング・コストがかかる）で「中高生」は手を出しづらいだろう。
その上，現時点ではウイルス対策ソフトの導入は「やらないよりマシ」程度のもので優先順位としては低いといえる[^v]。

[^v]: 「ウイルス対策ソフト」を装う malware も多い。そういうのは無料で提供されてたりするので「中高生」がうっかり手を出してしまう可能性もある。

また「信頼できるサイトからのアプリの導入」というのは表現としては手ぬるい。
何を持って信頼できるのか客観的な指標はない。
したがって現時点では Google Play や App Store といった**公式**のストア以外からは導入しない，とすべき。
Malware は Phishing メールなどから「うっかり」導入してしまうことが多いからだ。
もちろん Andorid 端末では「提供元不明アプリのインストールを許可しない」設定を有効にすることを忘れないこと。

結局一番有効なのは「OS やアプリを最新バージョンに保つ」ことであるが，ここでも問題がある。

まず iOS は基本的に機能追加のついででしかセキュリティ・アップデートを行わない。
しかも機能追加によって端末に不具合が起きることも少なくないためアップデートそのものをユーザが忌避することも多い。
少なくとも通常のアップデートとセキュリティ・アップデートは別サイクルで行うべきで，可能であれば iOS のバージョン別で行うのが望ましい。

Android の場合は更に酷い。
OS のアップデートは端末メーカの裁量に任されているため，特に古い機種ではアップデートそのものをオミットするケースが多いのだ。
Google 直轄の Nexus シリーズがほぼ毎月セキュリティ・アップデートを行っているのに対してその他の端末メーカではそういったものが殆ど見られない。
これは明らかに怠慢である。

- [エフセキュアブログ : Androidのセキュリティ上の大問題の1つを示す2つのグラフ](http://blog.f-secure.jp/archives/50763657.html)

そういった状況を無視して脳天気に「OS やアプリを最新バージョンに保つ」と言ってしまうのはいかがのものかと思ってしまう。

あといい加減に「ウイルス」って言うの止めない？ 誤解しか生まないよ，それ。

### 無線 LAN のセキュリティ

無防備な公衆無線 LAN を使ったために通信内容が漏れてしまった，という例。

無線 LAN の暗号化は端末とアクセスポイントの間の通信経路の間でしか行われない。
これは重要なポイントである。
もともと無線 LAN の暗号化は「有線」の代替として考えられているもので通信経路全体や通信データを守るものではない。
そもそも，やり方はどうあれ，「公衆」の「LAN」が信用できる道理はないのだ。

だから対策として「公衆 Wi-Fi に接続する場合は、出来るだけ暗号化された、信頼できる Wi-Fi を利用しよう」というのは完全にミス・リーディングである[^w]。

[^w]: どうでもいいけど「無線 LAN」と「Wi-Fi」で表現が分離してるけど何か意味あるの？ 用語は統一しようよ。

公衆無線 LAN サービスを使う場合には，ちゃんとユーザごとに認証機能（EAP など）のあるものでないとダメ。
公衆無線 LAN サービスの利用方法で SSID とパスワードしか記載されてないようなものは暗号化してないも同然なので，絶対に利用してはならない。
たとえ自治体などが提供しているサービスであってもだ。

安全でない公衆無線 LAN をどうしても使う必要がある場合は VPN (Virtual Private Network) を使って暗号化すること。
ただしまともな VPN サービスは有料（ランニング・コストがかかる）のものが多く「中高生」では金銭的に手が出ないことも多いだろう。

結局「中高生」相手であれば「公衆無線 LAN は絶対に利用しない」くらいに言ってしかるべきだと思う。
もちろん親とかの支援を受けて VPN を含むセキュリティ対策アプリを導入できるのであれば，それに越したことはないが。

### インターネット上の詐欺

「インターネット上の詐欺」というのは要するに social engineering のことである。
Social engineering の手法は日々進歩している。
数年前なら日本語の Phishing メールや誘導先のサイトはひと目で分かるようなものだったが，最近は見た目も手口も巧妙になってきていて本物と見分けがつかなくなってたりする。

- [日本郵政のスタッフを装ったメールに要注意、割と自然な日本語、外部リンクや添付ファイルなどの危険要素も -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20151216_735599.html)
- [Amazon.co.jpをかたるフィッシングサイトに注意、「.co」ドメインで開設 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160201_741724.html)
- [貴様のアカウントの利用中止を避けるために――りそな銀行かたるフィッシングメールに注意 -INTERNET Watch](http://internet.watch.impress.co.jp/docs/news/20160226_745607.html) （いや「貴様」てw）

まぁこういった情報に留意していくしかない。
こういうのは「私は騙されない」と思ってる人ほど騙されるもので「私も騙されるかもしれない」と覚悟した上で「事後」をどうすべきか学校や家族も交えて議論していくべきだと思う。

思うのだが，インターネットに限らず「中高生」向けにセキュリティ事例を共有していくシステムを構築すべきと思う。
むやみに恐れる必要はないが，日頃から心構えをしておくことが大切である。

これは企業とかでよくやるが，定期的に「抜き打ちテスト」するのも効果的だ。
あらかじめ「詐欺」を演出し，その対応を見て個別に指導していくというもの。
この場合も演出する側が「最新の手口」を知っていることが重要。
やはり日頃の情報収集と共有は欠かせないのだ。

### SNS 利用上の注意

例示が古いよ。

今時の子どもは「見知らぬユーザ」にホイホイついていったりしないって。
むしろ注意すべきは「知ってる友人」になりすましているケースだ。
LINE とかで昨年一昨年と騒ぎになって金銭被害とか出たじゃん。
これ作った奴はもっと勉強しろよ。

まぁこれも上の「インターネット上の詐欺」と同じ social engineering である。
なので対策も「インターネット上の詐欺」と同じ。

常に最新情報に留意して「事後」についてもあらかじめきちんと決めておく，程度の対処しかない。
