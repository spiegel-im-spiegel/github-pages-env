+++
title = "VTuber ごっこ"
date =  "2026-04-28T20:50:41+09:00"
description = "Ubuntu 環境で VTube Studio と Facetracker を連携し，OBS で VTuber 風の合成映像を作るまでをまとめる。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "streaming", "ubuntu", "vtuber", "tools", "flatpak" ]
pageType = "text"
draft = false

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

[前回]は Web カメラからの映像をゲーム画面と合成して Zoom に出力するところまでやった。
こうなると次は VTuber だよね。

そういや，今って VTuber の定義ってあるんだっけ。
まぁ，定義するのがバカバカしいほど意味が広がってるのは確かだけど。
[日本版 Wikipedia の記事](https://ja.wikipedia.org/wiki/%E3%83%90%E3%83%BC%E3%83%81%E3%83%A3%E3%83%ABYouTuber "バーチャルYouTuber - Wikipedia")を見ても「結局 VTuber ってなんやねん？」という感想しか出ないし。

日頃お世話になってる [Kagi Assistant] に「技術寄りに定義して」と頼んだら，こう返してくれた。

{{< div-ai type="markdown" >}}
2D または 3D の CG アバターを用い、トラッキング技術によって演者の動きをリアルタイムに反映させながら、動画配信プラットフォームで活動する存在
{{< /div-ai >}}

まぁ，この辺が無難だろうな。
そして「技術的な構成要素」として以下の3つを挙げてくれた。

- モーションキャプチャ（身体の動き）
- フェイシャルキャプチャ / トラッキング（表情）
- リップシンク（音声同期）

つまり動く CG キャラを使って上の3つの要素を満たせば，一応は「VTuber」ごっこができるわけだ。
「動画配信プラットフォームで活動」する気はないけど。

既存の Web カメラを使って Live2D[^l2d] のアバターを動かすことはできそうだ。

[^l2d]: 「Live2D」の名称およびロゴは[株式会社 Live2D の登録商標](https://www.live2d.jp/brand/ "ロゴとガイドライン | 株式会社Live2D")だそうな。また，開発用のソフトウェア（SDK および Live2D Cubism Editor）の利用には，利用者の規模や目的に応じた使用許諾契約（EULA）が適用されるとのこと。ちなみに [VTube Studio] にも EULA が設定されているので利用の際はご注意を。

調べてみたところ（調べたのは AI だが），どうやら Ubuntu 環境で Live2D モデルを動かすには [VTube Studio] 一択のようだ。
[VTube Studio] は Steam で提供されている。
使い方と [OBS Studio][OBS] との連携方法は以下のページが参考になった。

- [VTube Studioの使い方。重要なとこだけ設定してVTuberになる](https://vip-jikkyo.net/vtube-studio-tutorial)
- [VTube Studio + OBS設定ガイド｜モデルだけ映して背景透過する方法](https://vip-jikkyo.net/vtube-studio-with-obs)

ただ Ubuntu 環境で使う場合は，そのままでは Web カメラを認識しないので [Facetracker] も導入する必要がある[^ft1]。
[Facetracker] は [Flatpak] で提供されている[^fp1]。

[^ft1]: [VTube Studio] は [OpenSeeFace] に対応している。 [OpenSeeFace] 自体は Python スクリプトのようだが [Facetracker] が [OpenSeeFace] のインタフェースを持っているということらしい。
[^fp1]: [Flatpak] は Ubuntu に既定で入っていない。 [Flatpak] の導入については[前回の記事][前回]を参照のこと。

```text
$ sudo flatpak install flathub de.z_ray.Facetracker
Looking for matches…
Required runtime for de.z_ray.Facetracker/x86_64/stable (runtime/org.gnome.Platform/x86_64/50) found in remote flathub
Do you want to install it? [Y/n]: y

de.z_ray.Facetracker permissions:
    ipc    network    fallback-x11    wayland    x11    devices



        ID                                        Branch         Op    Remote     Download
 1. [✓] de.z_ray.Facetracker.Locale               stable         i     flathub      2.0 kB / 4.1 kB
 2. [✓] org.freedesktop.Platform.GL.default       25.08          u     flathub      4.4 MB / 142.4 MB
 3. [✓] org.freedesktop.Platform.GL.default       25.08-extra    u     flathub      3.0 MB / 142.4 MB
 4. [✓] org.freedesktop.Platform.codecs-extra     25.08-extra    u     flathub    869.0 kB / 14.4 MB
 5. [✓] org.gnome.Platform.Locale                 50             i     flathub    147.7 kB / 385.9 MB
 6. [✓] org.gnome.Platform                        50             i     flathub    254.2 MB / 408.8 MB
 7. [✓] de.z_ray.Facetracker                      stable         i     flathub    167.9 MB / 171.4 MB

Changes complete.
```

[VTube Studio] と連携させるために，以下のディレクトリに設定ファイル `ip.txt` を置く。

- `~/.local/share/Steam/steamapps/common/VTube Studio/VTube Studio_Data/StreamingAssets/`

`ip.txt` の内容は以下の通り。

```text
# To listen for remote connections, change this to 0.0.0.0 or your actual IP on the desired interface.
ip=0.0.0.0

# This is the port the server will listen for tracking packets on.
port=11573
```

起動するとこんな画面が表示される。

{{< fig-img-quote src="./facetracker.png" title="Facetracker" link="./facetracker.png" width="544" >}}

Webcom でカメラを指定する以外は既定のままでも動くが Server Settings はそのままだと多分ヤバいので変えたほうがいいだろう。
Server Settings を変えた場合は先程の `ip.txt` も記述を合わせること。

左上のボタン（アイコンではない）を押さないとトラッキングが始まらないので注意。

[VTube Studio] にカメラを繋いでキャリブレーションなどの設定をし，背景をグリーンバックにした状態がこれ。

{{< fig-img-quote src="./vtube-studio.png" title="VTube Studio" link="./vtube-studio.png" width="1288" >}}

[VTube Studio] にはあらかじめいくつかの Live2D モデルが用意されていて，比較的自由に使っていいみたい（EULA があるので注意）。
もちろん自前で用意したモデルも使える。

下の方で小さいのがちょろちょろ動いてるが，これが watermark になってるらしい。
[DLC](https://store.steampowered.com/app/1520620/VTube_Studio__Remove_Watermark/ "VTube Studio - Remove Watermark on Steam") を購入（1,520円）すると消せるそうな。
買い切り有り難い。

[VTube Studio] をソースとして [OBS Studio][OBS] に取り込んでフィルタを設定したところ。

{{< fig-img-quote src="./obs-studio-vtube-studio.png" title="OBS Studio" link="./obs-studio-vtube-studio.png" width="867" >}}

クロマキーで背景が綺麗に抜けているのが分かる。
[gogh] 画面と合わせるとこんな感じ。

{{< fig-img-quote src="./obs-studio.png" title="OBS Studio" link="./obs-studio.png" width="1286" >}}

うんうん。
いい感じやね。
くたびれたオッサンの絵面がこうなるんだから，そりゃあ VTuber が流行るわけだよ（笑）

さて [OBS] はだいたい遊び尽くしたかな。
[Streamplace] を使った [gogh] 作業配信は気まぐれで続けている。
Bluesky でアナウンスするので，よろしかったらどうぞ。

## ブックマーク

- [RogueRen/Linux-Guide-to-Vtubing: this is a guide to help Vtubers who use Linux or vtubers who want to start using linux - Codeberg.org](https://codeberg.org/RogueRen/Linux-Guide-to-Vtubing)

{{% review-paapi "B0DBZ3QP7J" %}} <!-- VTuber学 -->
{{% review-paapi "B0D1V1WXRH" %}} <!-- VTuberの哲学 -->

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

[前回]: {{< relref "/remark/2026/04/send-obs-composited-video-to-zoom.md" >}} "OBS 合成映像を Zoom に出す"
[gogh]: https://gogh.gg/ "gogh - Focus with Your Avatar"
[OBS]: https://obsproject.com/ "Open Broadcaster Software | OBS"
[Flatpak]: https://flatpak.org/ "Flatpak—the future of application distribution"
[VTube Studio]: https://denchisoft.com/ "VTube Studio – Official Website"
[Facetracker]: https://flathub.org/en/apps/de.z_ray.Facetracker "Install Facetracker on Linux | Flathub"
[OpenSeeFace]: https://github.com/emilianavt/OpenSeeFace "emilianavt/OpenSeeFace: Robust realtime face and facial landmark tracking on CPU with Unity integration"
[Streamplace]: https://stream.place/ "Streamplace"
