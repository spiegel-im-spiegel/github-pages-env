+++
title = "Windows に Ubuntu を入れる"
date = "2019-03-27T22:54:51+09:00"
description = "ここまで来るのにものすごい試行錯誤しちゃったよ。これでようやく色々試せる。"
image = "/images/attention/kitten.jpg"
tags = [ "windows", "virtualbox", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Microsoft は Windows 7 を2020年1月まで **ちゃんと** サポートすると言っていたのに，今の時点で Windows 10 への移行を強制するようだ。

- [Windows 7に未解決の脆弱性、Google Chromeのゼロデイ攻撃で組み合わせて悪用 - ITmedia エンタープライズ](https://www.itmedia.co.jp/enterprise/articles/1903/12/news065.html) : Windows 7 の脆弱性は放置するらしい。つか直せる人がもういないんだろうな
- [Windows 7セキュリティアップデートの終了期限をマイクロソフトが通達  |  TechCrunch Japan](https://jp.techcrunch.com/2019/03/21/2019-03-20-windows-7-security-updates/)

しかし，ここはひとつ寛大な心で「Microsoft は親切にも早く私が Windows を捨てるよう応援してくれてる」と解釈して Linux への移行を前倒しすることにした。
本当は新しい PC をどーんと買っていろいろ弄り倒して遊びたいところだが，今は[経済的・物理空間的な理由でできない]({{< ref "/remark/2018/12/i-am-a-sunday-programmer.md" >}} "どうも，日曜プログラマの Spiegel です")ため，今ある Windows 環境を Linux に換装することを考える。

といっても環境が違いすぎるためまずは簡単な仮想環境を構築していろいろ実験していきたいと思う。
仮想環境には [VirtualBox] を，デスクトップ用のディストリビューションには [Ubuntu] を使うことにする。
バージョンは2019年3月時点で [VirtualBox] 6.0.4, [Ubuntu] 18.10[^u1] である。

[^u1]: [Ubuntu] 18.10 は 2019年7月までサポートされる。長期サポートが欲しいなら 18.04 がオススメ。こちらは2023年4月までサポートされる。

日本語環境として [Ubuntu Japanese Team](http://www.ubuntulinux.jp/) から[「Ubuntu Desktop 日本語 Remix」をダウンロード](http://www.ubuntulinux.jp/download/ja-remix "Ubuntu Desktop 日本語 Remixのダウンロード | Ubuntu Japanese Team")して利用する。
ありがたや。

## [Ubuntu] 用仮想環境の作成

[VirtualBox] のインストール手順は割愛する。
インストールが完了したら早速 [VirtualBox] マネージャを起動する。

{{< fig-img src="./virtualbox-6-manager.png" link="./virtualbox-6-manager.png" width="976" >}}

この画面から「新規」アイコンをクリックする。

{{< fig-img src="./new-guest-expert.png" link="./new-guest-expert.png" width="568" >}}

メモリサイズは多めに2048MBほど取っておく。

{{< fig-img src="./virtual-hdd.png" link="./virtual-hdd.png" width="508" >}}

仮想ディスクは固定サイズで50GBほど。
まぁ検証用だし。

これで仮想環境を作成すると最終的に以下のようになる。

{{< fig-img src="./end-of-setup.png" link="./end-of-setup.png" width="976" >}}

## [Ubuntu] のインストール

ここで「起動」アイコンをクリックして仮想環境を起動する。
空の仮想ハードディスクなので起動用のドライブを訊いてくる。
ここで「[Ubuntu Desktop 日本語 Remixのダウンロード](http://www.ubuntulinux.jp/download/ja-remix)」ページからダウンロードしたディスク・イメージファイルを指定する[^iso1]。

{{< fig-img src="./assign-iso-image.png" link="./assign-iso-image.png" >}}

[^iso1]: ダウンロードしたイメージ・ファイルからインストール用 DVD を作成して使ってもいいのだが（つか実際の運用は DVD か USB からインストールすることになるだろう），めっさ遅いので注意すること。つか実際に遅すぎて使い物にならなかったのでインストールを中断した。本番はどうなることやら。

[Ubuntu] インストーラが起動した状態が以下の通り。

{{< fig-img src="./install-ubuntu-01.png" link="./install-ubuntu-01.png" width="816" >}}

ここからインストールを進めていくのだが，長ったらしくなるので途中は省いて，結果は以下の通りとなる。

{{< fig-img src="./install-ubuntu-08.png" link="./install-ubuntu-08.png" width="816" >}}

これで再起動すれば一応は完了である。
再起動のあとも最新版へのアップデートが走ったりするのだが，これも割愛する。

## 画面の解像度を変更する

この状態でも問題なく使えるのだが（ネットワーク設定も必要なし）仮想ディスプレイが SVGA のままで変更できない。

{{< fig-img src="./ratio-01.png" link="./ratio-01.png" width="822" >}}

このままでは不便なので解像度を変更できるようにする。
解像度を変更するには Guest Additions を導入する。

メニューバーの「デバイス」から「Guest Additions CD イメージの挿入...」を選択する。

{{< fig-img src="./ratio-02.png" link="./ratio-02.png" width="818" >}}

ここで [VirtualBox] をインストールしたフォルダ[^vb1] にある `VBoxGuestAdditions.iso` を選択する。
[Ubuntu] 側の管理者権限でプログラムが自動起動しようとするので，大人しくしたがって実行する。

[^vb1]: 既定のままインストールすれば `C:\Program Files\Oracle\VirtualBox` にインストールされる

{{< fig-img src="./ratio-04.png" link="./ratio-04.png" width="818" >}}

これで以下のように複数の解像度を選べるようになった。

{{< fig-img src="./ratio-05.png" link="./ratio-05.png" width="826" >}}

マウントした仮想ドライブはもう不要なので「取り出す」でアンマウントしておこう。

{{< fig-img src="./ratio-06.png" link="./ratio-06.png" width="816" >}}

いやぁ，ここまで来るのにものすごい試行錯誤しちゃったよ。
これでようやく色々試せる。

## ブックマーク

- [「Oracle VM VirtualBox 6.0」が正式版に ～「Hyper-V」フォールバックをサポート - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1159338.html)
- [VirtualBox を Windows にインストール - 山崎屋の技術メモ](https://www.shookuro.com/entry/2018/01/28/162252)
- [VirtualBox で作成した仮想マシンに CentOS をインストール - 山崎屋の技術メモ](https://www.shookuro.com/entry/2018/02/03/165526)
- [VirtualBox その93 - VirtualBox 6.0.4がリリースされました・VirtualBoxをインストールするには - kledgeb](https://kledgeb.blogspot.com/2019/01/virtualbox-93-virtualbox-604virtualbox.html)
- [VirtualBoxで仮想マシンを作成してみた（Windows10） | あんらぶぎーくどっとこむ](https://anlovegeek.com/create-virtual-machine/)
- [VirtualBoxの仮想マシンにUbuntuをインストールしてみた（18.10） | あんらぶぎーくどっとこむ](https://anlovegeek.com/virtualbox-install-ubuntu/)
- [VertualBox ゲストOSのスクリーンショットを撮る。](http://www.invisible-works.com/archives/2016/01/post-300/)
- [VirtualBox にインストールした Ubuntu の画面サイズ（解像度）を640×480以外に変更する方法 | mogi2fruitsどっとねっと](https://mogi2fruits.net/blog/os-software/windows/2389/)
- [その他　Windows上のVirtualBoxで解像度とマルチモニタの設定の仕方](https://www.oborodukiyo.info/etc/2016/ETC-MultiMonitorOnVirtualBox)
- [UbuntuをUSBメモリからインストール | Linuxサーバより愛を込めて。](https://chee-s.net/ubuntu%e3%82%92usb%e3%83%a1%e3%83%a2%e3%83%aa%e3%81%8b%e3%82%89%e3%82%a4%e3%83%b3%e3%82%b9%e3%83%88%e3%83%bc%e3%83%ab)

[VirtualBox]: https://www.virtualbox.org/ "Oracle VM VirtualBox"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%EF%BC%BB%E5%90%88%E6%9C%AC%E7%89%88%EF%BC%BD%E9%80%99%E3%81%84%E3%82%88%E3%82%8C%EF%BC%81%E3%83%8B%E3%83%A3%E3%83%AB%E5%AD%90%E3%81%95%E3%82%93-%E5%85%A8%EF%BC%91%EF%BC%92%E5%B7%BB-GA%E6%96%87%E5%BA%AB-%E9%80%A2%E7%A9%BA-%E4%B8%87%E5%A4%AA-ebook/dp/B072ML1SB2?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B072ML1SB2"><img src="https://images-fe.ssl-images-amazon.com/images/I/61juhv2SfIL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%EF%BC%BB%E5%90%88%E6%9C%AC%E7%89%88%EF%BC%BD%E9%80%99%E3%81%84%E3%82%88%E3%82%8C%EF%BC%81%E3%83%8B%E3%83%A3%E3%83%AB%E5%AD%90%E3%81%95%E3%82%93-%E5%85%A8%EF%BC%91%EF%BC%92%E5%B7%BB-GA%E6%96%87%E5%BA%AB-%E9%80%A2%E7%A9%BA-%E4%B8%87%E5%A4%AA-ebook/dp/B072ML1SB2?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B072ML1SB2">［合本版］這いよれ！ニャル子さん　全１２巻 (GA文庫)</a></dt>
	<dd>逢空 万太</dd>
	<dd>狐印 (イラスト)</dd>
    <dd>SBクリエイティブ 2017-06-08 (Release 2017-06-08)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B072ML1SB2</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">ニャル子さん合本版が半額セールしてたので衝動買いした。紙の本に換算して3327ページ（笑） こういう無茶が出来るのがEブックのいいところだよね。でも，ニャル子さんを読む（観る）と仮面ライダーが見たくなるんだよなぁ。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-27">2019-03-27</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> 0.2.1)</p>
</div>
