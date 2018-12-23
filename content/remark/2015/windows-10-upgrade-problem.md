+++
date = "2015-12-13T13:12:02+09:00"
update = "2016-07-10T18:27:46+09:00"
description = "自宅のマシンをチェックしたら，既に遅かった。"
draft = false
tags = ["windows"]
title = "また Windows 10 にヤラレタ（KB3112343 の恐怖）"

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
  url = "http://www.baldanders.info/spiegel/profile/"
+++

## KB3112343 の恐怖

- [Microsoft様、お願いですからWindows 7のままでいさせてください（KB3112343の恐怖）：海外速報部ログ：オルタナティブ・ブログ](http://blogs.itmedia.co.jp/burstlog/2015/12/microsoftwindows_7kb3112343.html)

これ見て慌てて自宅のマシンをチェックしたら，既に遅かった。

{{< fig-img src="https://farm1.staticflickr.com/630/23083745173_dddce0d481.jpg" title="windows update 201512" link="https://www.flickr.com/photos/spiegel/23083745173/" >}}

Windows Update にこういうの混ぜないで欲しいなぁ。
しょうがない。
削除するか。

## Windows Update によるアップグレードを抑止する

まずはコントロールパネルから「プログラムと機能」を表示し，その中の「インストールされた更新プログラムを表示」する。

{{< fig-img src="https://farm1.staticflickr.com/763/23082675514_d35d3628e3.jpg" title="windows update 201512" link="https://www.flickr.com/photos/spiegel/23082675514/" >}}

これで今までインストールした Windows Update の一覧が表示されるので，この中から以下の項目があればアンインストールする。

- KB2952664
- KB3021917
- KB3035583
- KB3112343
- KB3123862
- KB3173040

一覧から目視で探すのは大変なので，右肩にある検索窓から上に挙げた KBxxxxxxx の番号を入力して探してみると簡単にできる。
アンインストールを行うと再起動を要求されることがあるが，アンインストールを全部行ったあと再起動すればよい。

次にコントロールパネルから「Windows Update」を開く。
利用可能な更新プログラムの中に「Windows 10 にアップグレード」の文言のある項目や上に挙げた KB3112343 等の更新プログラムがあれば全て「非表示」にする。

{{< fig-img src="https://farm6.staticflickr.com/5674/23684736236_3a5b44e179.jpg" title="windows update 201512" link="https://www.flickr.com/photos/spiegel/23684736236/" >}}

トドメにローカル・グループ・ポリシーを変更する。
今度は `gpedit.msc` を起動する（起動の方法はご随意に，多分管理者権限が必要）。

起動したら「コンピュータの構成」→「管理用テンプレート」→「Windows コンポーネント」→「Windows Update」を開く。

{{< fig-img src="https://farm1.staticflickr.com/655/23343726479_f484fe8914.jpg" title="local group policy editor" link="https://www.flickr.com/photos/spiegel/23343726479/" >}}

右側に表示されている一覧から「Turn off the upgrade to the latest version of Windows through Windows Update」をダブルクリックで開く。

{{< fig-img src="https://farm1.staticflickr.com/637/23685672316_5cd228bfc0.jpg" title="Turn off the upgrade to the latest version of Windows through Windows Update" link="https://www.flickr.com/photos/spiegel/23685672316/" >}}

左上にある「有効(E)」のラジオボタンをチェックして「適用」ボタンを押せば OK。
もとの画面で「Turn off the upgrade to the latest version of Windows through Windows Update」の項目が「有効」になっていれば設定が効いている。

Windows 10 が悪いとは言わない（まぁ敢えてオススメはしないけど）。
しかし，その気がない人にまで強制的にアップグレードさせようというのは悪質すぎる。
こういうことやるから Microsoft は嫌われるんだってそろそろ自覚しろよ。

個人的には Windows 7 のサポートが切れる2020年までにメインの環境を Linux に移行する予定なんで（そのために[秀丸から ATOM に乗り換え]({{< ref "/remark/2015/atom-editor.md" >}})たりしてるんだし），要らんことしないでほしい。

## 参考

- [Windows 10 の広告アイコンを消す方法 - Qiita](http://qiita.com/spiegel-im-spiegel/items/bbc91030c26bc3c799f7)
- [【続】 Windows7でWindows10 無償アップデートのアイコンをアンインストール | 空中庭園](http://fortune-work.com/2015/windows10-2.html)
- [Windows 7のWindows Updateで表示される「Windows 10 にアップグレード」を削除する方法 | 空中庭園](http://fortune-work.com/2015/windows10-4.html)
- [エフセキュアブログ : Windows 10を安心して使用するために知っておくべき5つのこと](http://blog.f-secure.jp/archives/50752605.html)
- [Windows 10 で PSK を共有する — しっぽのさきっちょ | text.Baldanders.info]({{< ref "/remark/2015/wifi-sense.md" >}})
- [「Windows 10」へのアップグレード、来年には「推奨される更新プログラム」に“格上げ”へ - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1511/02/news076.html)
- [「Windows 10にアップグレードしませんか？」のポップアップ広告プログラムがさらにアグレッシブに変更される - GIGAZINE](http://gigazine.net/news/20160113-new-kb3035583/)
- [MS、「Windows 10」を「推奨される」更新プログラムとして提供開始 - CNET Japan](http://japan.cnet.com/news/service/35077208/)
- [Windows 7／8.1→Windows 10が“推奨される更新”に - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1602/02/news081.html)
- [Tech TIPS：まだWindows 10へアップグレードしたくない人のための設定まとめ - ＠IT](http://www.atmarkit.co.jp/ait/articles/1603/18/news047.html)
- [「Windows 10」自動アップデートをオフにするツール「Never 10」が公開 - ZDNet Japan](http://japan.zdnet.com/article/35080272/)
- [“Windows 10アップグレード”の公式キャンセル方法をマイクロソフトが解説、6月“Windows 10の日” - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1004723.html)
- [【特集】『朝起きたらOSが勝手にWindows 10になってた！』場合の対処法　前編 - 窓の杜](http://forest.watch.impress.co.jp/docs/special/1004285.html)
    - [【特集】『朝起きたらOSが勝手にWindows 10になってた！』場合の対処法　後編 - 窓の杜](http://forest.watch.impress.co.jp/docs/special/1004476.html)
- {{< pdf-file title="Windows 10への無償アップグレードに関し、確認・留意が必要な事項について" link="http://www.caa.go.jp/adjustments/pdf/160622adjustments_1.pdf" >}}
    - [消費者庁がWindows 10への自動アップグレードについてWindows 7/8.1ユーザーへ注意喚起 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1006675.html)
- [Microsoft、Windows 10へのアップグレード通知画面を改善し“お断りボタン”を追加 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1606/29/news063.html)
- [Windows 10の無償アップグレード問題、ウインドウを閉じても自動更新しない仕様に - CNET Japan](http://japan.cnet.com/news/service/35085243/) : 終わり間際になってようやくか
- [「アップグレードに関する情報発信が不十分だった」――日本マイクロソフト、Windows 10の通知内容を謝罪 - ITmedia PC USER](http://www.itmedia.co.jp/pcuser/articles/1607/05/news157.html)
