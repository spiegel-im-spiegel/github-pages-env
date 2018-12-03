+++
date = "2016-10-15T15:20:41+09:00"
update = "2016-10-16T15:54:04+09:00"
description = "今回は完全に失敗だった。次に機種変更するときはもう HTC は選ばない。"
draft = false
tags = ["k-tai"]
title = "電話機を交換してもらったのですよ"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

この夏は大変でしたが，少しずつ再起動中です。
いろいろ試してみたいこととかあるのですがなかなか時間が取れず。
というか，余裕がないというか。
しかも広島がリーグ優勝なんかしちゃったものだから体重が少し増えてしまったじゃないですか（←他人のせいにしている）。

それはさておき。
買ってまだ1年ちょっとしか経ってないのに端末機のバッテリがイカれちまったですよ。

{{< fig-youtube id="AW4MiYudAJI" title="忌野清志郎　雨あがりの夜空に - YouTube" width="500" height="281" >}}

まぁ7月くらいから兆候はあったのだが「そのうちそのうち」と先延ばしにしてたら，電源繋がないと5分も経たずに OS が落ちるようになった。
まぁ多分バッテリがヘタれて所定のパワーが出なくなってるのだろう。
フル充電状態でもカメラとか GPS とか起動すると確実に落ちるので。

au ショップに持ち込んだら「修理か交換か」と言われたが，交換のほうが早いらしいので交換で。

前持ってた機種が（OS のサポート以外は）優秀だったので今の機種も HTC 社製にしたのだが，今回は完全に失敗だった。
HTC は今 VR に力を入れているらしいが，ケータイも満足に作れないメーカーが VR とかヘソでお茶が沸いてしまう。
次に機種変更するときはもう HTC は選ばない。

そうして新しい端末が来たのだが，[前にも書いた](http://www.baldanders.info/spiegel/log2/000852.shtml "HTC J butterfly HTV31（もしくは最近のスマホ？）がなかなか酷い")とおり，ホンマに面倒くさい。
どうでもいいアプリは山ほどプリインストールしてくさるのに，なぜ「[Google 認証システム]」が最初から入ってないのだ。

Google Play を開くには Google アカウントが必要で，2要素認証でログインするには「[Google 認証システム]」が必要になる[^iij]。
でも「[Google 認証システム]」をインストールするためには Google Play にログインしなきゃならないんだよ。
こんなアホな話があるかっての。
まぁ Google の場合はいくつか代替手段があるのでまだマシだけど[^sms]。

[^iij]: IIJ の「[IIJ SmartKey](https://play.google.com/store/apps/details?id=jp.ad.iij.smartkey2)」でもよい。
[^sms]: ちなみに SMS を使った認証はもはや推奨されない。網間プロトコルである通称「No.7」に脆弱性があり，これを使って通話やテキスト通信が比較的簡単に傍受できることが知られている（参照： [SMSを使った二要素認証を非推奨〜禁止へ、米国立技術規格研究所NISTの新ガイダンス案 | TechCrunch Japan](http://jp.techcrunch.com/2016/07/26/20160725nist-declares-the-age-of-sms-based-2-factor-authentication-over/)）。実際にこの脆弱性を利用できるプレイヤーは限られると思うが，今だに2要素認証に SMS を使わせようとするサービス・プロバイダがいたら「あぁ，ここはユーザのセキュリティやプライバシーを軽視してるんだな」と蔑みの目で見てよい。

私の場合は以下の手順で最初のセットアップを行った。
覚え書きのためにメモしておく。

1. どうにかして Google Play にログインする（ここが一番面倒くさかった）
1. 「[Google 認証システム]」と「[QRコードスキャナー](https://play.google.com/store/apps/details?id=com.google.zxing.client.android)」をインストールする
1. 2要素認証が必要なサービス毎に [TOTP] の secret を取得して「[Google 認証システム]」にセットする（大抵の場合 secret は QR コードで示される）
1. 「[Keepass2Android]」をインストールする
1. 「[Keepass2Android]」にパスワード情報の入ったデータベースファイル（と必要なら暗号鍵）を読み込ませる
    - 私の場合はデータベースファイルをストレージ・サービス上に置いて端末間で共有しているので，そのサービスにアクセスするためのアプリもインストールしてログイン処理を済ませておく必要がある。暗号鍵はさすがに怖くてクラウドに置けないのでローカル環境から手動でコピった

これで最初の準備が完了。
次のステップでは端末内の環境を整理する。

1. 無線 LAN の接続設定を行う（接続に必要なパスワード情報は「[Keepass2Android]」に入れている。あんな長いパスワード，手動で入力できるかっての）
1. プリインストールされているアプリで不要なものをアンインストールまたは無効化する（システムの構成上，削除しちゃダメなアプリやサービスもあるので作業は慎重に。自信がないなら無理にやらなくてよい）
1. Google Play や他のストア・アプリからインストール済みアプリのアップデートを行う。私の場合は47個ほどあった気がする
1. ストレージの暗号化を行う

このストレージの暗号化が毎回ドキドキするんだよなぁ。
まぁ失敗じることはないだろうけど。

そうそう。
この機会に「避難訓練」しておくことをお薦めする。

- [「Googleアカウントがアップデート」]({{< ref "/remark/2016/06/05-stories.md#dm" >}} "週末スペシャル： 古いパソコンに Ubuntu を入れようと思ったが...")

端末の捜索と鳴動だけでも試してみるといいだろう。
端末のロックはまたの機会にでも。

ここまでくればあとは自由にアプリを入れられる。
以下は個人的なオススメ。

- [Swarm](https://play.google.com/store/apps/details?id=com.foursquare.robin), [Instagram](https://play.google.com/store/apps/details?id=com.instagram.android) : SNS アプリ。ちなみに Facebook と Twitter は最初から入っていた。 LINE は使わないので削除
- [Signal], [Messenger] : メッセンジャーアプリ。 [Signal] は1対1の E2E 暗号化通信を行い，標準の SMS アプリと置き換えることもできる。 Facebook の [Messenger] は最近 [E2E 暗号化に対応](http://japanese.engadget.com/2016/10/05/facebook-messenger/ "Facebook Messengerの暗号化機能、全ユーザーで利用可能に。チャット開始時に「秘密のスレッド」を指定 - Engadget Japanese")した。 [Signal] と同じ方式。めでたい！
- [Box](https://play.google.com/store/apps/details?id=com.box.android), [Dropbox](https://play.google.com/store/apps/details?id=com.dropbox.android), [Amazon Drive](https://play.google.com/store/apps/details?id=com.amazon.drive)/[Amazon Photos](https://play.google.com/store/apps/details?id=com.amazon.clouddrive.photos),  [SpiderOakONE](https://play.google.com/store/apps/details?id=com.spideroak.android) : Storage アプリ。 [SpiderOakONE](https://play.google.com/store/apps/details?id=com.spideroak.android) は E2E 暗号化を行う。オススメ。私は Prime 会員なので Amazon のストレージへ写真を無制限にアップロードできる。なので写真は [Amazon Drive](https://www.amazon.co.jp/clouddrive/)，重要なデータは [SpiderOak](https://spideroak.com/)，それ以外は用途に応じて [Box](https://www.box.com/) か [Dropbox](https://www.dropbox.com/) を使うようにしている
- [Firefox](https://play.google.com/store/apps/details?id=org.mozilla.firefox) & [Pocket](https://play.google.com/store/apps/details?id=com.ideashower.readitlater.pro) : ブラウザは必要。 [Pocket](https://play.google.com/store/apps/details?id=com.ideashower.readitlater.pro) は「あとで読む」ために使っている
- [Inbox](https://play.google.com/store/apps/details?id=com.google.android.apps.inbox), [ProtonMail](https://play.google.com/store/apps/details?id=ch.protonmail.android) : メールソフト。 Gmail は [Inbox](https://www.google.co.jp/inbox/) に乗り換えた。メールを TODO 管理のように使えるのが特徴（自分でリマインダを発行することもできる）。暗号化メールが必要な場合は [ProtonMail](https://protonmail.com/ "Secure email: ProtonMail is free encrypted email.") を使う。
- [Freedome](https://play.google.com/store/apps/details?id=com.fsecure.freedome.vpn.security.privacy.android) : F-Secure 社製の VPN クライアント。有料だが日本にローミング・ポイントがある。 [Freedome の運用ポリシーはこちら](http://blog.f-secure.jp/archives/50745430.html "エフセキュアブログ : プライバシーに対するFreedomeのアプローチ")を参考にどうぞ。お金を払わなくてもアプリのセキュリティ検査はやってくれる。
- [Amazon Music](https://play.google.com/store/apps/details?id=com.amazon.mp3) : 音楽アプリはこれに統一した。私の端末には2千2百曲ちょっと入っている。若い頃買った CD とか全部つっこんでるからなぁ
- [Simplenote](https://play.google.com/store/apps/details?id=com.automattic.simplenote) : 名前の通りシンプルなテキストエディタ。「ファイル」という概念がないのが特徴。保存先を [simplenote.com](https://simplenote.com/) にして複数端末（と Web）で共有することができる。テキスト入力のみの Evernote みたいなものって感じだろうか
- [RealCalc](https://play.google.com/store/apps/details?id=uk.co.nickfines.RealCalc) : 電卓アプリ。 N 進数演算ができるのが個人的にお気に入り

これでようやく落ち着いたよ。

## ブックマーク

- [HTC J butterfly HTV31（もしくは最近のスマホ？）がなかなか酷い](http://www.baldanders.info/spiegel/log2/000852.shtml)

[Google 認証システム]: https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2
[TOTP]: https://tools.ietf.org/html/rfc6238 "RFC 6238 - TOTP: Time-based One-time Password Algorithm"
[Keepass2Android]: https://play.google.com/store/apps/details?id=keepass2android.keepass2android "Keepass2Android Password Safe"
[Signal]: https://play.google.com/store/apps/details?id=org.thoughtcrime.securesms "Signal Private Messenger"
[Messenger]: https://play.google.com/store/apps/details?id=com.facebook.orca "Messenger"

