+++
title = "勤務先からの払い下げ PC"
date =  "2024-05-18T14:26:51+09:00"
description = "こういう機会でもないと Mac パソコンなんか買うこともない。"
image = "/remark/2024/05/get-a-used-pc-from-workplace/53726198297_fc01691b54_o.jpg"
tags = [ "tools", "gear", "apple", "macos" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

GW 明けに勤務先が古い Mac パソコンを大量放出するとアナウンスした。
いや，もっとはよ言うてや。
既に [Chromebook 買っちまった]({{< ref "/remark/2024/03/chromebook-1.md" >}} "Chromebook を導入する 1")ぢゃん `orz`

まぁ，でも，こういう機会でもないと Mac パソコンなんか買うこともないだろうから，手を挙げてみた。
若い人で欲しい人がいればそちらを優先するつもりだったのだが，私以外に誰も手を挙げなかったようで，見事に当選してしまった。
購入したのは2020年製の MacBook Air M1。

- [MacBook Air (M1, 2020) - 技術仕様 - Apple サポート (日本)](https://support.apple.com/ja-jp/111883)

いや，ちゃんとお金払いましたよ，消費税込みで[^t1]。
一応中古ショップの買取価格を参考に値段設定したらしい（転売禁止と念を押されたw）。
無料で下げ渡すと色々と手続きが面倒なんだそうな。

[^t1]: メルカリとか消費税はどうなってるんだろうと思ったが，個人による売買については消費税の対象にならないらしい。個人事業主や法人による売買の場合はかかるんだって。

というわけでいただきました。

{{< fig-img src="./53726198297_6e4bf46aa4_e.jpg" title="勤務先からの払い下げ品 GET！ | Flickr" link="https://www.flickr.com/photos/spiegel/53726198297/" >}}

13インチ Retina ディスプレイ， Apple M1 プロセッサ。
メモリは16GBに換装してある。
ストレージは256GB SSD。
まぁ，仕事で使うわけじゃなし，個人で遊ぶならこの程度のスペックで十分やろ。

業務上のデータは全部引き上げてると聞いていたので，まるっと初期化から。

- [【簡単】MacBookを初期化する方法！10通りの注意点とリセットできない時の対処法 | リネットジャパン通信](https://www.renet.jp/media/macbook-initialization/)

シャットダウンしている状態で `Command+R` キーと Touch ID のキーを押し続けてしばらく待つと「起動オプション」画面が立ち上がる。
そこから「オプション」をクリックして表示された復旧メニューから更に「ディスクユーティリティ」を開く。
ディスクユーティリティの「表示」で「すべてのデバイスを表示」にして，そこで表示される “Apple SSD xxxxxx Media” を「消去」すればいいようだ。
マシンを再起動したら復旧メニューが再表示されるので，そこから「macOS xxxx を再インストール」すればいいようだ。
再インストールの前に Wi-Fi の設定と言語の設定が必要かも。

以下作業しながらつらつらと感想。

- この辺の手順は macOS のバージョンによって微妙に違うらしい。迷惑な
- ネットインストールで[自宅環境]({{< ref "/remark/2024/04/fiber-optic-network.md" >}} "光回線を導入する")でも1時間以上かかった。とほほ
- この作業で初めて知ったのだが MacBook の中の OS ってコンテナ化されてるんだね。合理的ではある
- Apple ID がなくとも初期設定はできるんだな[^ai1]。やっぱ ChromeOS や Windows がクズなだけか
- 起動時の「ジャーン」が五月蝿い。ソッコーで無効にした。今どき Windows でもやらねーよ，そんなの
- 自宅 NAS には簡単につながった。よしよし

[^ai1]: Apple ID を作ったのは最初期の iPad を購入したときで，当時のメールアドレスは（spam がウザいので）失効させてるしクレカ番号も失効してるはず。というわけで，新しく Apple ID を作り直した。使わないで済みますように。

13インチノートなのでプライベート鞄に入るし重量も1.29kgで，この前買った Chromebook よりちょっと軽いんだよね。
なので上手くセッティングできれば外出用PCとして使えるかな？ とは思っている。

{{< ruby "だがしかし" >}}駄菓子菓子{{< /ruby >}}。

今まで Mac なんて全く触ったことがないので何をどうすればいいかさっぱり分からない。
たとえば Homebrew を入れようとしたら `xcode-select` がどうとか言われ（なんじゃそれ？），そこから先に進めないでいる。
最悪は macOS を潰して Linux を入れるのもありかなぁ，とか不穏なことも考えている。

まぁ，のんびり行こう。

{{< div-box type="markdown" >}}
**【2025-01-29 追記】**
仮想環境を導入し Kubuntu を入れて運用を始めた。

- [MacBook Air M1 に Kubuntu を入れる]({{< ref "/remark/2025/01/kubuntu-on-macbook-air-m1.md" >}})
{{< /div-box >}}

## ブックマーク

- [M1などAppleシリコン搭載MacBookの初期化（フォーマット）方法 | いわっちろぐ](https://iwatti.com/format-applesilicon-macbook/)

## 参考

{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B07KJWYQJW" %}} <!-- ANKER PowerExpand USB メディアハブ -->
{{% review-paapi "B088R6SV4Z" %}} <!-- ANKER 充電器 Power Delivery (PD) 対応。65W -->
{{% review-paapi "B08P54PQDB" %}} <!-- メッセンジャーバッグ -->
