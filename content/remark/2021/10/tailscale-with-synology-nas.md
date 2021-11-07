+++
title = "Synology NAS に Tailscale を設定する"
date =  "2021-10-31T15:42:53+09:00"
description = "いよいよ Tailscale に手を出すことにした。"
image = "/images/attention/kitten.jpg"
tags = [ "nas", "vpn", "tailscale", "wireguard" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./nas.md" >}} "秋 NAS は俺に喰わせろ！")からチマチマ作業して，ようやく中身の整理と Git Server の設定までできた。
[以前まで使ってた簡易 NAS]({{< ref "/remark/2018/04/nas.md" >}} "NAS を導入した") では無造作にバックアップを取ってたので，特に音楽ファイルと写真ファイルの整理で時間がかかってしまったよ。

写真ファイルは8K個くらいあったんだけど，この規模だとダブりとか確認しようがないし，途中で面倒くさくなって全部 [Synology Photos](https://www.synology.com/ja-jp/DSM70/SynologyPhotos) サービスに突っ込んでしまった。

一方で音楽ファイルは2.3K個ほど。
これは普通だよね。
アルバム換算で160枚分程度だし。

で，まぁ，そんなこんなで落ち着いて運用できるようになったので，いよいよ [Tailscale] に手を出すことにした。

## Tailscale とは

[Tailscale] とは VPN (Virtual Private Network) 実装のひとつである。

なお，この記事では VPN を物理的な回線・ネットワークの上に構築された「{{% ruby "virtual network" %}}実質的な網{{% /ruby %}}」であると定義し，以下の機能を持つものとして話を進める。

- カプセル化とトンネリング
- 経路の暗号化
- 認証とアクセス制御

旧来の VPN は，拠点間 VPN であれリモートアクセスであれ「[バイパス型](https://baldanders.info/blog/000773/ "VPN に関する基礎学習")」であれ，出口ノードがネットワークそのものになっているのが特徴である。

{{< fig-img-quote class="lightmode" src="./home-network-before.svg" title="Tailscale · Best VPN Service for Secure Networks" link="https://tailscale.com/" width="630" lang="en" >}}

これに対して [Tailscale] では機器間を直接つないでいるのが特徴と言える[^exit1]。

[^exit1]: [Tailscale] 接続先機器を出口ノードとしてネットワークにアクセスすることは可能，[らしい](https://tailscale.com/kb/1019/subnets/ "Subnet routers and traffic relay nodes · Tailscale")。

{{< fig-img-quote class="lightmode" src="./home-network-after.svg" title="Tailscale · Best VPN Service for Secure Networks" link="https://tailscale.com/" width="630" lang="en" >}}

[Tailscale] の実装は [WireGuard] 技術を用いており，サービスとしての [tailscale][Tailscale].com は接続機器のコンフィグレーションを管理するのみで通信そのものには関与しない。
[正しく E2E 暗号化通信]({{< ref "/remark/2021/08/apples-mass-surveillance-plans.md" >}})というわけだ（笑）

2021-10-31 現在，個人ユーザは20接続までの制限付きで無料で利用できる。
また年48USDほど払えば100接続まではいけるらしい。
企業またはチームとして利用する場合はユーザごとに接続先を制御する必要があるだろうし，そのための料金プランも用意されているようだ。

{{< fig-img-quote src="./price-of-tailscale.png" title="Pricing · Tailscale" link="https://tailscale.com/pricing/" width="1357" lang="en" >}}

旧来の VPN はいわゆる ZTA (Zero Trust Architecture) とは相性が悪いと言われているが [Tailscale] を使って機器単位でアクセス制御できるのであれば ZTA にも組み込みやすいだろう。
個人的にはスマート・ホームなど IoT 機器こそ [Tailscale] に対応してもらいたいよなぁ。

ちなみに [Tailscale] の基本ロジックは [Go] で書かれている。
まぁ [WireGuard] のユーザランド実装が [Go] だしね。

## Synology NAS に Tailscale をインストールする

今回，私が導入した Synology NAS には [Tailscale パッケージ]が用意されおり，パッケージセンターからインストールできる。

{{< fig-img src="./tailscale-in-synology-nas.png" link="./tailscale-in-synology-nas.png" width="801" lang="en" >}}

これは現在 GitHub で [Tailscale] 公式のリポジトリとして運営されているようだ。

- [tailscale/tailscale-synology: Synology packages for tailscale.com][tailscale-synology]

さっそくインストールしてパッケージを開くと

{{< fig-img src="./login.png" link="./login.png" width="567" lang="en" >}}

てな感じにログインを促す画面が開く。
早速 `[Log In]` ボタンを押すとブラウザの [tailscale][Tailscale].com のサインイン・ページが開く。

{{< fig-img src="./signin-by-tailscale.png" link="./signin-by-tailscale.png" lang="en" >}}

おっ GitHub アカウントが使えるのか。
よしよし。

このページで認証が完了したら再び [Tailscale パッケージ]を開いて

{{< fig-img src="./connected-2.png" link="./connected-2.png" width="566" lang="en" >}}

てな感じに表示されれば接続完了である。
うん，簡単簡単。

なお IP アドレスの `100.64.0.0/10` 領域は Carrier-Grade NAT の “Shared Address Space” ([RFC 6598](https://datatracker.ietf.org/doc/html/rfc6598 "IANA-Reserved IPv4 Prefix for Shared Address Space")) として確保されているものらしい。
こんなん使って大丈夫か？ と思うが，まぁユーザ側が気にすることではあるまい。

## tailscale-synology のアップデート

Synology NAS から見て [tailscale-synology] はサードパーティ製品なので自動アップデートに対応してないらしい。
故に GitHub の [tailscale/tailscale-synology][tailscale-synology] リポジトリのリリース情報に注意して手動でアップデートを行う必要がある。

インストール後のパッケージセンターの表示を見ると

{{< fig-img src="./tailscale-package-center-1.png" link="./tailscale-package-center-1.png" width="909" lang="en" >}}

となっている。
2021-10-31 時点の最新は [v1.16.2](https://github.com/tailscale/tailscale-synology/releases/tag/v1.16.2 "Release v1.16.2: update to 1.16.2 · tailscale/tailscale-synology") なのでちょっと古いようだ。

“[List of spksrc architecture per Synology model](https://github.com/SynoCommunity/spksrc/wiki/Architecture-per-Synology-model "Architecture per Synology model · SynoCommunity/spksrc Wiki")” によると [DS220j] は “rtd1296 (armv8)” だそうなので `tailscale-armv8-1.16.2-2013-dsm7.spk` ファイルをダウンロードして適用すればいいのかな。
やってみよう。

上の画面の右上にある `[手動インストール]` ボタンを押して「パッケージのアップロード」を開く。

{{< fig-img src="./upload-tailscale-package-1.png" link="./upload-tailscale-package-1.png" width="707" lang="en" >}}

ここで先ほどダウンロードしたファイルを指定して `[次へ]` ボタンを押すと

{{< fig-img src="./upload-tailscale-package-3.png" link="./upload-tailscale-package-3.png" width="705" lang="en" >}}

と警告が出るが，ビビりつつ `[同意する]` ボタンを押す。
すると「設定の確認」が出るので

{{< fig-img src="./upload-tailscale-package-4.png" link="./upload-tailscale-package-4.png" width="707" lang="en" >}}

`[完了]` ボタンを押す。
アップデートが完了すると，先ほどのパッケージセンターの表示が

{{< fig-img src="./tailscale-package-center-2.png" link="./tailscale-package-center-2.png" width="909" lang="en" >}}

とバージョンが上がっていることがわかる。

よーし，うむうむ，よーし。

## Android 端末に Tailscale をインストールする

次は手持ちの Android 機に [Tailscale] をインストールして，外から前節で設定した Synology NAS にアクセスしてみる。

Andorid 版 [Tailscale] は Google Play Store からインストールできる。

{{< fig-img src="./tailscale-in-google-play-store.png" link="./tailscale-in-google-play-store.png" width="1024" lang="en" >}}

いや，3歳は無理やろ（笑）

インストールを完了しアプリを開くとサインイン画面になる。

{{< fig-img src="./signin-by-tailscale-android.png" link="./signin-by-tailscale-android.png" width="1024" lang="en" >}}

「あれ？ Google だけ？」と一瞬思ったが，焦らず `[Sign in with Other]` ボタンを押すと NAS のときと同じ [tailscale][Tailscale].com のサインイン・ページが開く。

{{< fig-img src="./signin-by-tailscale-browser.png" link="./signin-by-tailscale-browser.png" width="1024" lang="en" >}}

Google 以外ならこちらのページで認証を行うとよい。
認証が完了すると

{{< fig-img src="./connected-by-android-2.png" link="./connected-by-android-2.png" width="1024" lang="en" >}}

などと表示される。
自機以外に先程の NAS も表示されているのが分かるだろうか。
なお [Tailscale] で接続済みの機器は，ブラウザでも[機器一覧ページ](https://login.tailscale.com/admin/machines)で確認することができる。

この状態で [FE File Explorer Pro](https://play.google.com/store/apps/details?id=com.skyjos.apps.fileexplorer "File Explorer Pro - file manager and file transfer - Google Play") を使ってアプリに表示されている IP アドレスに対して SMB 接続を確認できた。
めでたし！

ちなみに私は「[バイパス型 VPN](https://baldanders.info/blog/000773/ "VPN に関する基礎学習")」として F-Secure 社の [FREEDOME VPN][FREEDOME] を購入して使っているが，これを起動した状態で [Tailscale] を有効にしても問題なく相手機器に接続できた。
よかった，よかった。

## ブックマーク

- [Tailscale v1.16 · Tailscale](https://tailscale.com/blog/tailscale-v1.16/)
- [Supported SSO identity providers (Google, AzureAD, GitHub, Okta, etc) · Tailscale](https://tailscale.com/kb/1013/sso-providers/)
- [Enabling Synology outbound connections · Tailscale](https://tailscale.com/kb/1152/synology-outbound/)
- [Tailscale for developers: Connect to your resources from GitHub Codespaces · Tailscale](https://tailscale.com/blog/github-codespaces/)
- [100台まで無料のVPNサービス「tailscale」、リンクだけでマシンのシェアも可能!?【イニシャルB】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/column/shimizu/1303751.html)
- [Tailscale VPN を使ってみたので感想 | つくみ島だより](https://blog.tsukumijima.net/article/tailscale-vpn/)
- [リモートワーク向けにセキュアなVPNを構築【tailscale】 | 4b-media](https://4b-media.net/tailscale/)

- [「門番」としての VPN]({{< ref "/remark/2020/07/vpn-as-a-gatekeeper.md" >}})
- [NIST SP 800-207: “Zero Trust Architecture”]({{< ref "/remark/2020/09/nist-sp-800-207-zero-trust-architecture.md" >}})

[WireGuard]: https://www.wireguard.com/ "WireGuard: fast, modern, secure VPN tunnel"
[Go]: https://golang.org/ "The Go Programming Language"
[Tailscale]: https://tailscale.com/ "Tailscale · Best VPN Service for Secure Networks"
[DS220j]: https://www.synology.com/ja-jp/products/DS220j "DS220j | Synology Inc."
[Tailscale パッケージ]: https://tailscale.com/kb/1131/synology/ "Access Synology NAS from anywhere · Tailscale"
[tailscale-synology]: https://github.com/tailscale/tailscale-synology "tailscale/tailscale-synology: Synology packages for tailscale.com"
[FREEDOME]: https://www.f-secure.com/en/home/products/freedome "F-Secure FREEDOME VPN — Secure and private browsing | F-Secure"
