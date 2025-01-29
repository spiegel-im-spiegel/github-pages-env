+++
title = "Macbook Air M1 に Kubuntu を入れる"
date =  "2025-01-29T22:11:22+09:00"
description = "いつまでも放置というわけにもいくまい！ ということで"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "ubuntu", "macos", "nas", "tailscale" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨年勤務先から払い下げてもらった [MacBook Air (M1, 2020)][MacBook] なのだが， [Homebrew] のインストールで「面倒くさい」となって最初の1フィートで挫折していた。
外出先では (その前に買った) [Chromebook]({{< ref "/remark/2024/03/chromebook-1.md" >}} "Chromebook を導入する 1") で十分賄えてたというのもある。

とはいえ，いつまでも放置というわけにもいくまい！ と今年に入ってようやく重い腰を上げることにした。
以下の優れた記事を見かけたというのもある。

- [新MacでLinuxシリーズ 目次](https://zenn.dev/armcore/articles/maclinux_index)

痒いところに手が届く良記事だった。
是非 Zenn 本の体裁にして（もしくはガチの本として再構成して）いただきたいものである。
感謝 {{% emoji "ペコン" %}}

## 構成

上で挙げた記事を参考に，今回は以下のように構成した。

- **ハードウェア：** [MacBook Air (M1, 2020)][MacBook]
  - *チップ：* apple M1, 8コア
  - *メモリ：* 16GB
  - *ストレージ：* 256GB
  - *モニタ：* 内蔵13インチ Retina ディスプレイ
- **ホストOS：** macOS Sequoia バージョン 15.3
- **仮想環境：** [UTM for Mac] 4.6.4
  - [QEMU] 9.1 ARM Virtual machine (ハイパーバイザー Type 2)
- **ゲストOS：** [Ubuntu Server] for [ARM](https://ubuntu.com/download/server/arm "Ubuntu for ARM | Download | Ubuntu") 24.04.1 LTS ＋ [Kubuntu] Package

自宅機が Ubuntu 機なので馴染みのある構成にしたかったのだが Ubuntu desktop の最新 LTS の ARM 用インストール・パッケージはないみたいなので，まずは省構成の Ubuntu サーバを入れて，その後 Kubuntu パッケージを組み込んで，ちまちまとチューニングしていく感じ。

[Kubuntu] は Ubuntu のフレーバーのひとつで，デスクトップ UI に [KDE Plasma] を使っているのが特徴。
良く言えばカスタマイズしがいのある，悪く言えば（Ubuntu desktop に比べて）手のかかるディストリビューションと言える。
とはいえ，今回みたいなよく分からないハードに載せるのならチューニングは必至だろう，という判断で今回の構成にした。

仮想環境については，軽くググってみたところ [Parallels] を買え！ という意見が多かったように見えた（あくまで印象，定量評価ではない）。
[Parallels] のほうが安定しているそうな。
私としてはお金を払ってまでやる気はなかったし，最初に挙げた記事を見る限り [UTM][UTM for Mac] でも大きな問題はなさそうと判断して [UTM][UTM for Mac] を導入することに決めた。

ちなみに [UTM][UTM for Mac] は Apple Store で「購入」することもできる（1.5K円）。
この場合，支払ったお金は開発資金に回されるそうな。

## 仮想マシンのセットアップ

[UTM][UTM for Mac] のインストールや使い方については「[新MacでLinuxシリーズ 目次](https://zenn.dev/armcore/articles/maclinux_index)」からリンクされている各記事を見ていただくとして，ここでは仮想マシン作成時の構成について覚え書きとして残しておく。

{{< fig-img src="./setup-new-machine-1.png" title="新規仮想マシンを作成 - Linux" link="./setup-new-machine-1.png" width="2236" >}}

[QEMU (Quick EMUlator)][QEMU] を使うので「Apple仮想化を使用」はチェックしない。
[Ubuntu Server] の ISO イメージはあらかじめダウンロードしておく（ARM 版を取得すること）。

{{< fig-img src="./setup-new-machine-2.png" title="新規仮想マシンを作成 - ハードウェア" link="./setup-new-machine-2.png" width="2236" >}}

既定ではハードウェア・リソースの半分を専有しようとする。
私の場合は Linux 環境がメインなので，メモリ10GBでコアも6つ専有することにした（内部で JIT が走るのでガメすぎないこと）。
ちなみにこの設定で仮想環境を立ち上げた状態で macOS 側で YouTube が問題なく聴けた。

「ハードウェア OpenGL アクセラレーションを有効にする」をチェックしたのだが，失敗だったかもしれない（これについては後で述べる）。
この項目はあとで仮想マシンを作成したあとは修正できないので注意。

{{< fig-img src="./setup-new-machine-3.png" title="新規仮想マシンを作成 - ストレージ" link="./setup-new-machine-3.png" width="2236" >}}

指定したサイズを最初から専有するわけでないのでご安心を。
今回は最大サイズの半分を指定した。

以下は「[新MacでLinuxシリーズ 目次](https://zenn.dev/armcore/articles/maclinux_index)」に従って [Kubuntu] のインストールが完了して再起動した状態。

{{< fig-img src="./kubuntu-1.png" title="Kubuntu (1)" link="./kubuntu-1.png" width="1504" >}}

その後も色々とセットアップしてこんな感じに（以下は全画面表示）。

{{< fig-img src="./kubuntu-2.png" title="Kubuntu (2)" link="./kubuntu-2.png" width="2880" >}}

とりあえずこんなもんかな。

## チューニングは続く

### Tailscale for Linux

外で使うマシンなので，出先から[自宅 NAS]({{< ref "/remark/2021/10/nas.md" >}} "秋 NAS は俺に喰わせろ！") にアクセスするために [Tailscale] を[導入](https://tailscale.com/kb/1031/install-linux "Setting up Tailscale on Linux · Tailscale Docs")する。

つっても

```text
$ curl -fsSL https://tailscale.com/install.sh | sh
```

を実行するだけ。
これで環境に合わせて apt リポジトリのセットアップから [Tailscale] のインストールまでやってくれる。
あとは

```text
$ sudo tailscale up
```

とすれば systemd で動き始める。

[Tailscale] 越しの NAS アクセスを `/etc/fstab` で定義してマウントしまえばいいと思ったのだが，マシン起動時に [Tailscale] サービスへ認証を行う方法を思いつかなかったので `noauto` を指定して起動時にマウントしないようにした。

- [CIFS 経由で NAS に接続する]({{< ref "/remark/2019/03/common-internet-file-system.md" >}})

### UTM はマルチディスプレイに対応しているが...

[UTM][UTM for Mac] はマルチディスプレイに対応してるそうだが，自動的に複数のディスプレイを認識して分割（またはミラーで）表示するるわけではなく，あらかじめ仮想マシンの設定で

{{< fig-img src="./setup-new-machine-6d.png" title="仮想マシンの編集" link="./setup-new-machine-6d.png" width="2236" >}}

こんな感じにディスプレイを必要なだけ追加していく。
これで仮想マシンを起動した際に設定したディスプレイの数だけウィンドウが起動する。

ただ，セカンダリのディスプレイのウィンドウ画面のチラツキが酷くて使い物にならなかったのでマルチディスプレイは諦めた。
まぁ，出先で運よくモニタを借りれるとは限らないからな。

### もしかして「ハードウェア OpenGL アクセラレーションを有効」にしたのは失敗だった？

なんか VS Code と相性が悪いのかたまにウィンドウが真っ黒になるんだよな。
一度こうなると再起動しない限り回復しない。
前節のマルチディスプレイでちらつきが酷い件も併せ，原因は分からないが，今になって考えるに

- 「ハードウェア OpenGL アクセラレーションを有効にする」にチェックを入れないほうがよかった？
- ホストOS側で YouTube を流しっぱなしにしてたのがあかんかった？

あたりは検証したほうがいいかもなぁ。

それでも日常使いする程度には整ってきたので，今後は Chromebook じゃなくてこちらを持ち歩くことになるかな。

## Mac なんも分からん

作業していて痛感するのだが，マジで macOS さっぱり分からん。
たとえば，ブラウザで [UTM][UTM for Mac] のバイナリパッケージ `UTM.dmg` をダウンロードしたまではよかったが，そこからどうすればいいのか分からない。
恐る恐るダブルクリックしたらこんなウィンドウが出る。

{{< fig-img src="./install-utm.png" title="どうすりゃいいの？" link="./install-utm.png" width="1544" >}}

しょうがないのでググってはじめて「左のアイコンを右にD&Dすればいいのか」と分かった。
せめて説明文を書いてくれ `orz`

一事が万事こんな調子なのよ。
私のような年寄りには macOS を使うためのチュートリアルが必要だな！

それにしても，メインで使う OS をゲスト OS として仮想化して使うのは筋がよくないな。
まぁ Macbook を完全に Linux 機として換装できなさそうだからこういう手段を執らざるを得なかったんだけど。
お高い Macbook を買って，わざわざ Linux 機や Windows 機として使うやつはおらんと思う。
私の場合は，たまたま払い下げ品を安く手に入れられただけで，定価なら（自腹では）絶対に買わない。
Apple 信者じゃないので，私。

ある意味で贅沢な使い方かもな（笑）

[MacBook]: https://support.apple.com/ja-jp/111883 "MacBook Air (M1, 2020) - 技術仕様 - Apple サポート (日本)"
[Homebrew]: https://brew.sh/ "Homebrew — The Missing Package Manager for macOS (or Linux)"
[Ubuntu Server]: https://ubuntu.com/download/server "Get Ubuntu Server | Download | Ubuntu"
[Kubuntu]: https://kubuntu.org/ "Kubuntu | Friendly Computing"
[KDE Plasma]: https://kde.org/plasma-desktop/ "Plasma - KDE Community"
[Parallels]: https://www.parallels.com/ "Parallels: Mac & Windows Virtualization, Remote Application Server, Mac Management Solutions"
[UTM for Mac]: https://mac.getutm.app/ "UTM | Virtual machines for Mac"
[QEMU]: https://www.qemu.org/ "QEMU"
[Tailscale]: https://tailscale.com/ "Tailscale · Best VPN Service for Secure Networks"

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
{{% review-paapi "B079MCPJGH" %}} <!-- カメラ 目隠し シャッター -->
{{% review-paapi "B08LMYWKZD" %}} <!-- Bluetooth 無線静音マウス -->
{{% review-paapi "B07KJWYQJW" %}} <!-- ANKER PowerExpand USB メディアハブ -->

## 作業中の BGV (メン限配信以外)

- [【前編】【LIVE】歴史探訪アプリ「たまむすび」を一緒に使ってみよう！＋奈良旅の思い出話【古代日本史VTuber きら子】 - YouTube](https://www.youtube.com/watch?v=rYU30qXDo3M)
