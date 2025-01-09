+++
title = "Mini PC を衝動買いした"
date =  "2025-01-09T18:46:24+09:00"
description = "これで Windws 11 Pro がバンドルされてて26K円ちょっと"
image = "/remark/2025/01/win11pro-on-minipc/54253033452_cc05918a02_o.jpg"
tags = [ "engineering", "windows", "linux" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今年は散財せず慎ましく生活していこうと思ったのに `orz`

以下の記事を見て

- [【西川和久の不定期コラム】邪魔なACアダプタがなくなった！？Twin LakeのIntel N150搭載ミニPC「Beelink EQ14」  - PC Watch](https://pc.watch.impress.co.jp/docs/column/nishikawa/1652152.html)

衝動買いしてしまった（←他人のせいにしようとしている）

{{< fig-img src="./54253033452_728a6fa959_e.jpg" title="ミニPCを衝動買い | Flickr" link="https://www.flickr.com/photos/spiegel/54253033452/" >}}

ハードウェアのスペックは以下の通り。

{{< fig-quote class="nobox" type="markdown" title="【西川和久の不定期コラム】邪魔なACアダプタがなくなった！？Twin LakeのIntel N150搭載ミニPC「Beelink EQ14」 - PC Watch" link="https://pc.watch.impress.co.jp/docs/column/nishikawa/1652152.html" >}}
|      |      |
| :--- | :--- |
| プロセッサ | Intel N150(4コア4スレッド、クロック最大 3.6GHz、キャッシュ 6MB、TDP 最大25W) |
| メモリ | 16GB(SO-DIMM DDR4 3200MHz) |
| ストレージ | M.2 2280 SSD PCIe 3.0 500GB(PCIe 3.0 x4)、PCIe 3.0 x1(空き) |
| グラフィックス | Intel UHD Graphics 24EU/HDMI 2.0(4K@60Hz) 2基 |
| ネットワーク | Gigabit Ethernet×2、Wi-Fi 6、Bluetooth 5.2 |
| インターフェイス | USB 3.2 Gen 2 3基、USB 3.2 Gen 2 Type-C、USB 2.0、3.5mmジャック |
| サイズ/重量 | 126×126×39mm、490g |
{{< /fig-quote >}}

これで Windws 11 Pro がバンドルされてて26K円ちょっとで Amazon で売っていた。

{{% review-paapi "B0DNS713WK" %}} <!-- Mini PC ミニPC N150 Beelink EQ14 Win11pro -->

すンごいハイスペックというわけではないけど，これ採算とれてるの？
いや，中国メーカーだし赤字覚悟な値段なわけないか？ ミニPCの相場を知らないので適正価格が分からない。
まぁ，ヤバそうなら潰して Linux 機にするか（笑）

以前 [Azure Virtual Desktop で遊]({{< ref "/remark/2021/12/azure-virtual-desktop-2.md" >}} "Azure Virtual Desktop で遊ぶ")んだんだけど，継続的に使おうとするとランニングコストが馬鹿にならなくて。
ホビーで利用するには割に合わないと感じて削除したんだよね。
今回買ったミニPCなら Windows 環境下の動作検証用として必要十分かな，って感じ。

電源ケーブルと HDMI ケーブルは付属してるけど（当然ながら）モニタやキーボード・マウスや LAN ケーブル（必要なら）は自前で用意する必要がある。
昨年，衝動で[メカニカルキーボードを買った]({{< ref "/remark/2024/11/mechanical-keyboard.md" >}} "はじめてのメカニカルキーボード")ときに，余ったキーボードを捨てずに取っておいてよかったよ。

セットアップは滞りなく。
でも，せっかく省スペースのミニPCを買ったんだから，直にモニタやキーボード等を繋げるんじゃなくリモートデスクトップで Ubuntu 機からアクセスできないかな？ と思ったら Ubuntu デスクトップに標準で入ってる Remmina を使えばいいらしい。

{{< fig-img src="./remmina-setup-1.png" title="Remote Connection Profile (1)" link="./remmina-setup-1.png" width="877" >}}

中途半端に日本語化されているが（日本語化作業有難うございます）これで問題なく接続できた。

{{< fig-img src="./remote-desktop.png" title="リモートデスクトップ on Ubuntu デスクトップ" link="./remote-desktop.png" width="1920" >}}

（標準のブラウザを Firefox に入れ替えている）

クリップボードの共有も問題なくできてるし，音声もリモートオーディオで問題なく鳴っている。
よーし，うむうむ，よーし。

ちなみに上図はリモートデスクトップを最高品質で表示している。

{{< fig-img src="./remmina-setup-2.png" title="Remote Connection Profile (2)" link="./remmina-setup-2.png" width="877" >}}

最高品質より下だと画像が間引きされて大昔のファミコンゲームの画面みたいになった（笑）

これで準備完了。
ミニPC本体はデスクサイドラックに追いやって机上スッキリ。
これから Windows を私好みにカスタマイズするかね。
まずは [Scoop] 入れて Windows Terminal のシェルを [NYAGOS] に換装するところからかな。

## ブックマーク

- [【Ubuntu】UbuntuからWindowsをリモートデスクトップで操作する - Milkのメモ帳](https://www.milkmemo.com/entry/remote_desktop_ubuntu_win)
- [第661回　リモートデスクトップビューアー、Remminaを使用する | gihyo.jp](https://gihyo.jp/admin/serial/01/ubuntu-recipe/0661)

[Scoop]: https://scoop.sh/ "Scoop"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"

## 参考

{{% review-paapi "B07KXCJ8Q1" %}} <!-- デスクサイドラック -->
{{% review-paapi "B089GN1VKY" %}} <!-- BUFFALO 8ポート スイッチングハブ -->
{{% review-paapi "B087745J5B" %}} <!-- 電源タップ -->
