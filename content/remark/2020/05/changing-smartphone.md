+++
title = "ついカッとなって機種変した，反省はしない"
date =  "2020-05-15T00:02:50+09:00"
description = "もはやスマホにこれ以上パラダイム・シフトの夢を見ることはないし，スペックも最小限で無問題。"
image = "/images/attention/kitten.jpg"
tags = [ "k-tai", "gear", "android", "authentication", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[5年前に買ったスマートフォン](https://baldanders.info/blog/000852/ "HTC J butterfly HTV31（もしくは最近のスマホ？）がなかなか酷い")なのだが，相変わらず[バッテリ周りが酷くて]({{< ref "/remark/2016/10/new-phone.md" >}} "電話機を交換してもらったのですよ")今のが3台目だったのですよ。
その3台目も2年と経たずにイカれてしまい（勝手に強制リブートを繰り返すようになった）ついカッとなって機種変更した。
失業して銭のないときになんちう迷惑な。

近所（田舎の近所なので察してください）の au ショップに駆け込んで「テザリングが使えるいっちゃん安い機種を」と言ったら [Galaxy A20] を勧められた。
端末代が税込で 33,000JPY とのことで即金で買いましたよ。
もはやスマホにこれ以上パラダイム・シフトの夢を見ることはないし，スマホでゲームはしないからスペックも最小限で無問題だよね。

## 2要素認証するならリカバリ・コードは控えておくこと

今までの反省から，旧端末から認証情報やデータを移行できない可能性を考慮して準備しておいたのは助かった。

特に2要素認証[^2fa] の2要素目で TOTP を使っている場合は認証不能になる場合があるので，サービス側が発行するリカバリ・コードを必ずダウンロードして控えておくこと。
なんなら紙に印刷して厳重にしまっておけば確実だろう。

[^2fa]: 今だに「2段階認証」とかぬかす馬鹿メディアがあるみたいだが「2段階」では不十分だから「2要素」が要求されるのだ，ということを分かっているのだろうか。

言い換えると2要素認証を推奨しているのにリカバリ・コードの提供すらしないサービスはダメなサービスだと断言していいだろう。
まぁ SMS に一時パスワードを垂れ流して「**2段階** 認証だから安全」とか言ってくさるサービスとかあるけどな（笑）

最近流行りの認証デバイスを使えばそんな面倒もないんだろうけど，個人的には紛失・盗難リスクが怖くてノートパソコンや携帯端末に認証デバイスを使う気にならないんだよねぇ。

## 最初にすること {#first}

まずは OS のアップデートを行うこと。
最近の端末は最初からストレージの暗号化がされてるんだね。
よーし，うむうむ，よーし。

で，アプリのアップデートを行う前に既定で入ってるアプリで使わないものは削除する。
中には削除できないものもあるが，そういうのは，アップデート前であれば，見分けがつくので最初にやってしまおうってわけ。

なんで LINE や Facebook や Twitter のアプリが最初から意味もなく入ってるんだろうねぇ。
ぜんぶ削除ですよ。
あと端末メーカー製やキャリア製のアプリで明らかに使わないものは可能な限り削除する。
ついでに Google 製の不要アプリもザクザク削除。

これで，すっきり！

## 雑多な作業 {#misc}

「[最初にすること]({{< relref "#first" >}})」が終わったらパスワード管理アプリを入れる。

- [Keepass2Android](https://play.google.com/store/apps/details?id=keepass2android.keepass2android)

私の場合，データベースファイルをクラウド・ストレージに置いているので，クラウド・ストレージにアクセスするアプリも併せて導入する。
暗号鍵は USB で PC に直結して端末にコピってしまう。
暗号鍵とパスワードでデータベースファイルを二重にロックしておけば大丈夫だろう。

その後

- [File Explorer Pro](https://play.google.com/store/apps/details?id=com.skyjos.apps.fileexplorer)

を導入して LAN 上の NAS に入れるよう設定すれば一段落。

他にセキュリティ関連アプリとして

- [Google 認証システム](https://play.google.com/store/apps/details?id=com.google.android.apps.authenticator2)
- [FREEDOME](https://play.google.com/store/apps/details?id=com.fsecure.freedome.vpn.security.privacy.android)
- [Signal](https://play.google.com/store/apps/details?id=org.thoughtcrime.securesms)

を導入してセットアップする。
[Signal](https://play.google.com/store/apps/details?id=org.thoughtcrime.securesms) は既定の SMS アプリとしても設定できるので置き換える。

ブラウザは

- [Firefox Focus](https://play.google.com/store/apps/details?id=org.mozilla.focus)
- [Firefox](https://play.google.com/store/apps/details?id=org.mozilla.firefox)

を導入し [Firefox Focus](https://play.google.com/store/apps/details?id=org.mozilla.focus) の方を既定のブラウザにする。
[Firefox](https://play.google.com/store/apps/details?id=org.mozilla.firefox) は予備系とし Chrome や他のブラウザは使わないようにする。
もちろん検索サービスにはどちらも [DuckDuckGo](https://duckduckgo.com/) を指定する。

Input method は

- [Gboard](https://play.google.com/store/apps/details?id=com.google.android.inputmethod.latin)

で無問題。
ていうか，これ以外使いたくない。

あとは好みで

- [Feedly](https://play.google.com/store/apps/details?id=com.devhd.feedly)
- [Pocket](https://play.google.com/store/apps/details?id=com.ideashower.readitlater.pro)
- [Slack](https://play.google.com/store/apps/details?id=com.Slack)
- [Trello](https://play.google.com/store/apps/details?id=com.trello)
- [Flickr](https://play.google.com/store/apps/details?id=com.flickr.android)
- [RealCalc Plus](https://play.google.com/store/apps/details?id=uk.co.nickfines.RealCalcPlus)
- [Simplenote](https://play.google.com/store/apps/details?id=com.automattic.simplenote)
- [時計](https://play.google.com/store/apps/details?id=com.google.android.deskclock)
- [Googleカレンダー](https://play.google.com/store/apps/details?id=com.google.android.calendar)
- [Camera FV-5](https://play.google.com/store/apps/details?id=com.flavionet.android.camera.pro)
- [tenki.jp](https://play.google.com/store/apps/details?id=jwa.or.jp.tenkijp3)

あたりを順次入れていく。
[5年前](https://baldanders.info/blog/000852/ "HTC J butterfly HTV31（もしくは最近のスマホ？）がなかなか酷い")に比べればだいぶ顔ぶれが変わったなぁ。

## [Microsoft Launcher] を導入してみた {#launcher}

各端末メーカーが既定で入れてるランチャってなんであんなにダサいのかね。
いや，デザイン・センス皆無の私に言われたくないだろうけど。

今までは「どうせすぐ壊れるから」と手を付けなかったんだけど，今回は試しに [Microsoft Launcher] を導入してみた。
セットアップ時に Microsoft account を要求するのだが，大昔に登録したのがまだ有効だったようで，問題なく行けた。
位置情報も要求されるが，許可しなくても無問題（もしくは後から許可を取り消せる）。

ちなみに検索バーと連携する検索サービスに [DuckDuckGo](https://duckduckgo.com/) を指定できる。
使わんけどね。

## 「[デバイスを探す]」で避難訓練{#drill}

ひととおり設定が終わったら「[デバイスを探す]」で一度は避難訓練をしておくとよいだろう。
携帯端末の捜索についてはキャリアも端末メーカーもサービスを提供しているが，キャリアはともかく，端末メーカーのサービスは無用である。

「[デバイスを探す]」では検索対象の携帯端末に対して「音を鳴らす」「デバイスのロック」「データの消去」といった操作ができる。
「音を鳴らす」と「デバイスのロック」については一度は試してみることをお勧めする。

- [端末を探す](https://play.google.com/store/apps/details?id=com.google.android.apps.adm)

[Galaxy A20]: https://www.galaxymobile.jp/galaxy-a20/ "Galaxy A20（ギャラクシーA20）- Galaxy公式（日本）"
[Microsoft Launcher]: https://play.google.com/store/apps/details?id=com.microsoft.launcher "Microsoft Launcher - Google Play"
[デバイスを探す]: https://www.google.com/android/find
<!-- eof -->
