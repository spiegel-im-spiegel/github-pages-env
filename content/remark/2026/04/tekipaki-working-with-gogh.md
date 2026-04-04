+++
title = "gogh で てきぱきワーキング"
date =  "2026-04-04T19:08:09+09:00"
description = "gogh は OBS と組み合わせてゲーム内でストリーム配信することもできる。また YouTube などのサービスにも配信できる。 Streamplace でもできそう。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "games", "ubuntu", "windows", "atproto", "linux", "media" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

いわゆる作業支援ゲームあるいはもっと簡単に「作業ゲー」と呼ばれるものは色々と登場しているが [gogh] (ゴッホ) で自身の作業ルームを公開されている方を Bluesky で見かけたので私も試してみた。

- [gogh - Focus with Your Avatar][gogh]
- [ambr｜note](https://note.com/ambr) : [gogh] の制作会社によるブログ
- [gogh ゴッホ - YouTube](https://www.youtube.com/channel/UCecfiDpkMIMNoYWE-lBM2kQ) : [gogh] の公式 YouTube チャンネル
- [@goghjpn.bsky.social on Bluesky](https://bsky.app/profile/goghjpn.bsky.social) : [gogh] の公式 Bluesky アカウント

[Steam] 版の [gogh] は作業ルームをかなり細かくカスタマイズできて，しかもそれを他のユーザと共有できるというので，クリエイター寄りの方々にも人気らしい。

{{< fig-quote type="markdown" title="世界30万本を超えたSteam版goghを振り返る - 開発編｜ambr" link="https://note.com/ambr/n/ncc7ca92fc108" >}}
実際、マルチプレイヤー機能が入ってからは本当に爆発的に広まりました。これはgogh配信の大きな特徴なのですが、想像もしていなかったような**有名なイラストレーターさんや漫画家さんが次々に作業配信をしてくれました**。開発チームには漫画好きやイラストレーターファンが多いので「え…？あの先生も…！？」とチームのSlackに毎日驚きと喜びが溢れていました。
{{< /fig-quote >}}

モバイル版（Android/iOS）が無料で提供されているにもかかわらず有料の [Steam] 版がこれだけ売れてるというのは，この辺に理由があるのかも知れない。

試しに（マルチプレイヤーでも使えるように）大部屋をひとつ作ってみたのだが，うっかり興が乗ってしまい夜ふかし（というかほとんど徹夜）してしまった。
何やってるんだか。

{{< fig-img-quote src="./gogh-my-room.png" title="gogh" link="./gogh-my-room.png" width="1603" >}}

画像を絵や写真オブジェクトとして室内に設置できる。
上の部屋では [Flickr で公開](https://www.flickr.com/photos/spiegel/ "Yasuhiro ARAKAWA | Flickr")している写真を展示してみた。
まぁ，飾るような写真じゃないんだけどね。
雰囲気だよ雰囲気（笑）

## gogh と OBS

[gogh] は [OBS] と組み合わせてゲーム内でストリーム配信することもできる。
たとえばパソコンモニタやプロジェクターといったオブジェクトに [OBS] の仮想カメラの映像を映し出すことができる。

今まで配信とか無縁の生活だったが，面白そうなので  [OBS] Studio をインストールして試してみた。
のだが，どうしても上手くいかない。
最初は [OBS] 側のシーンやソースの設定の問題かと思って色々弄ってみたが全然ダメで，ふと思いついて Web カメラをソースとしてみたが，これもダメだった。
これって多分 [gogh] 側に仮想カメラの映像が渡ってないな。

私のメインマシンは Ubuntu なのだが，よく考えたら [Steam] 版 [gogh] は Windows 版のみの提供で Linux/Ubuntu 環境は [Proton] でエミュレーションして動かしているのだった。
だから [OBS] との連携ができなくても仕方ないのかな？

しょうがない。
[ミニ PC (Windows 機)]({{< ref "/remark/2025/01/win11pro-on-minipc.md" >}} "Mini PC を衝動買いした") に [OBS] Studio を入れるか。
あまりパワーがないのでまともに動くか不安だけど。

{{< fig-img-quote src="./gogh-live.png" title="gogh" link="./gogh-live.png" width="1316" >}}

おっ。
あっさり動いた。
自分で自分を映してるので合わせ鏡みたいになってるが，まぁよかろう。
一応 [gogh] とは別のゲームを起動して仮想カメラで映してみたが，これも問題なかった。
パワーがないのでちょっとカクカクしてたけど（笑）

.｡oO（マルチプレイヤー用に部屋を公開したらストリーミングが切れてしまった。これってこういうもの？）

[OBS] を使えば YouTube などのサービスにも配信できる。
まぁ，でも，これで稼ぐ気はないし，何となく Google のサービスは嫌なので [atproto] エコシステム下のサービスである [Streamplace] を使うことを考えてみようかな。

- [Start streaming with OBS | Streamplace Docs](https://stream.place/docs/guides/start-streaming/obs/)

上手くいったらまた報告する。

## ブックマーク

- [gogh、Steam版の販売本数が30万本を突破！ | 株式会社 ambrのプレスリリース](https://prtimes.jp/main/html/rd/p/000000070.000043299.html)
- [PC版『gogh（ゴッホ）』スタートガイド： アバター・ルーム・マルチプレイで集中空間をつくる | XRメモランダム](https://orecen.com/apps/gogh-app-steam/)
  - [OBS で YouTube 配信する話](https://orecen.com/apps/gogh-app-steam/)
- [gogh ゴッホの感想｜環境音とポモドーロで作業に集中｜瀧本祐ヰ](https://note.com/yuwitakimoto/n/n4745e154a187)
- [Xユーザーのgogh公式｜Steam版発売中＆アプリ配信中さん: 「[Steam版 Tips] gogh × OBSで「gogh内作業配信」する方法をまとめました！...](https://x.com/goghJPN/status/1920318480496173191)

[gogh]: https://gogh.gg/ "gogh - Focus with Your Avatar"
[OBS]: https://obsproject.com/ "Open Broadcaster Software | OBS"
[Steam]: https://store.steampowered.com/ "Welcome to Steam"
[Proton]: https://github.com/ValveSoftware/Proton "ValveSoftware/Proton: Compatibility tool for Steam Play based on Wine and additional components"
[Streamplace]: https://stream.place/ "Streamplace"
[atproto]: https://atproto.com/ "The AT Protocol"

## 参考

{{% review-paapi "4757700423" %}} <!-- てきぱきワーキン♡ラブ -->
