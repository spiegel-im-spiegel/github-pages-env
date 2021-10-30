+++
title = "秋 NAS は俺に喰わせろ！"
date =  "2021-10-26T23:50:36+09:00"
description = "金曜の夜にポチって中1日で届くとは思わざりき（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "nas", "samba", "ubuntu", "cloud", "nfs", "git" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

勤務先で LAN 上の NAS やクラウド・ストレージの構成を見直す動きが出ていて「そういや私が簡易 NAS を導入したのっていつだっけ？」と過去記事を漁ってみたら3年半も前だったよ。

- [NAS を導入した]({{< ref "/remark/2018/04/nas.md" >}})

その半年後に[田舎に{{% ruby "かえ" %}}帰郷{{% /ruby %}}る羽目になった]({{< ref "/remark/2018/12/i-am-a-sunday-programmer.md" >}} "どうも，日曜プログラマの Spiegel です")のだが，当時は Windows を捨てる準備に入ろうかというところで，バックアップと書き戻し用に（OS を跨いで）アクセスできる共有ストレージが欲しかったんだよね。

導入して良かった点は LAN 上の NAS とクラウド・ストレージ・サービスを連携させることで，そこそこ安全で快適な環境を手に入れられたことだろう。
とはいえ RAID 0 すら組んでない簡易設備でこのまま運用を続けるのはもう限界だろうと思い直したわけだ。

で，まぁ，あれこれ迷った挙句に買ったのは [Synology DS220j][DS220j] ってやつ。
今回も[ドスパラ](https://www.dospara.co.jp/)さんのお世話になりました。
[HDD 2基と3点セット](https://www.dospara.co.jp/5shopping/detail_set.php?bg=&br=&sbr=&camp=1804&ft=synology&lf=2 "【HDD+ドライブケース】Synology NAS + 選べる W.D HDD 3点セット｜｜｜通販のドスパラ")になってるやつをポチった。

本当はポチってから数日あけて来ることを想定してたんだけど，金曜の夜にポチって中1日で（松江みたいな片田舎まで）届くとは思わざりき（笑）

{{< fig-img src="./51625630824_e1ae072032_e.jpg" title="秋 NAS 来た！ | Flickr" link="https://www.flickr.com/photos/spiegel/51625630824/" >}}

ありがたいことに「NAS 初心者ガイド」という日本語の冊子がついていた。
これのおかげでセットアップ自体は比較的スムーズにできた。
ブラウザで Web Assistant を開いたときに最初はうまく検出できなかったくらい（ブラウザを立ち上げ直したら上手くいった）かな？

## RAID (Redundant Arrays of Inexpensive Disks)

RAID は SHR (Synology Hybrid RAID) にした。
RAID 1 とかだと ストレージの容量やメーカーを揃える必要があるが， SHR はその辺は柔軟に運用できるらしい。

- [Synology Hybrid RAID（SHR）とは？ - Synology ナレッジセンター](https://kb.synology.com/ja-jp/DSM/tutorial/What_is_Synology_Hybrid_RAID_SHR)

## QuickConnect は無効にした

最初は Synology アカウントを取得して QuickConnect も使ってたんだが，どうにも気持ち悪いので QuickConnect はいったん無効にした。
Synology アカウントは有効にしてるので，必要になったらまた復活させればいいだろう。

VPN 接続は [Tailscale](https://www.synology.com/ja-jp/dsm/packages/Tailscale?os_ver=7.0 "Tailscale - アドオン パッケージ | Synology Inc.") を検討中...

## Ubuntu ファイルマネージャから SMB サービスに入れない

[Ubuntu] 標準のファイルマネージャ（`nautilus`）は LAN 上の Windows 共有フォルダや [Samba] サービスを自動的に検出・表示してくれる。
[DS220j] も SMB (Server Message Block) が使えるので Windows ネットワークと同じように表示されるんだけど，クリックしても何故か入れないのよね。

{{< fig-img src="./samba-error.png" link="./samba-error.png" width="626" >}}

なんやねん「無効な引数」て `orz`

これは現在に至るも改善しない（理由がわからないので）。
ただし `mount.cifs` でマウントはできるし `/etc/fstab` ファイルへの記述もちゃんと効いたので実害はない。
多分ファイルマネージャ側の問題なんだろうなぁ。
CIFS 経由で NAS に接続する方法は以下の通り。

- [CIFS 経由で NAS に接続する]({{< ref "/remark/2019/03/common-internet-file-system.md" >}})

ちなみに Android の [FE File Explorer Pro](https://play.google.com/store/apps/details?id=com.skyjos.apps.fileexplorer "File Explorer Pro - file manager and file transfer - Google Play") からは普通に見れた。

## Ubuntu から NFS サービスに入れない

[Ubuntu] に NFS クライアントを導入する。

```text
$ sudo apt install nfs-common
```

[DS220j] 側の NFS サービスを有効化して [Ubuntu] からアクセスしようとしたが上手くいかない。
まぁ，考えてみたら当たり前で Kerberos を使ってアカウントのマッピングでもするのならともかく，家庭内 LAN ごときでそこまでしないし，しないなら見えなくても仕方ない。

{{< fig-quote type="markdown" title="NFS 権限を割り当てる | DSM - Synology ナレッジセンター" link="https://kb.synology.com/ja-jp/DSM/help/DSM/AdminCenter/file_share_privilege_nfs?version=7" >}}
**セキュリティ フレーバー：**

特定のユーザー アカウントで NFS を介して共有フォルダにアクセスする場合： 

- **AUTH_SYS** セキュリティが導入されている場合：クライアントは NFS クライアントと Synology NAS で全く同じ UID（ユーザー識別子）と GID（グループ識別子）を持っていなければなりません。そうしなければ、クライアントが共有フォルダにアクセスするときに、**「others」**という許可が割り当てられます。権限が競合しないように、**[Squash]** から **[全ユーザーを admin にマップ]** を選択するか、共有フォルダに｢Everyone｣権限を割り当てます。
- **Kerberos (krb5、krb5i、krb5p)** セキュリティが導入されている場合：**[ファイル サービス]** > **[NFS]** > **[NFS サービスの有効化]** > **[詳細設定]** > **[Kerberos の設定]** の順にを選択して、NFS クライアントを特定のユーザーに割り当てなければなりません。または、相当するユーザー アカウントで Windows/LDAP ドメインに加わります。そうしなければ、クライアントが共有フォルダにアクセスするときに、「guest」という許可が割り当てられます。
- 共有フォルダを作成する拡張デバイスのファイル システムが NTFS または FAT の場合は、**[すべてのユーザーを admin にマップ]** オプションが強制的に適用されます。

Synology NAS に接続するために Kerberos セキュリティを使用するには、**[ファイル サービス]** > **[NFS]** > **[NFS サービスの有効化]** > **[詳細設定]** > **[Kerberos の設定]** の順に選択して、Kerberos 認証が設定されていなければなりません。
{{< /fig-quote >}}

というわけで，今回は 権限を「カスタマイズ」して `Everyone` 権限を追加・設定した。
複数人で使う場合は NFS は使わんほうがいいかもなぁ。

ちなみに `/etc/fstab` ファイルを使ってマウントする場合には

```text
ds220j-hostname:/volumeX/homes/username /home/username/nas-home nfs rsize=8192,wsize=8192,timeo=14,intr 0 0
```

などとすればよい。

## Cloud Sync

パッケージ・センターから [Cloud Sync] をインストールしてストレージ・クラウドと連携できるようにした。

[Cloud Sync] は Google Drive, Dropbox, Box, Microsoft OneDrive あるいは Azure Storage, Google Cloud Storage, S3 storage といったメジャー・サービスはもちろん，一般的な WebDAV とも連携できるらしい。

さらに [Cloud Sync] では

{{< fig-img src="./cloud-storage-encrypt.png" link="./cloud-storage-encrypt.png" width="630" >}}

てな感じに同期データの暗号化ができる。
具体的には [Cloud Sync] からアップロードする際に暗号化された状態でアップロードされる。
元々クラウド・ストレージ側に置いてあるファイルは暗号化の対象にならない。
暗号化ファイルは NAS 側では復号された状態で格納されており，クラウド・ストレージ側で暗号化されているか否かに関係なく透過的に扱うことができる。

クラウド・ストレージって便利で使い勝手がいいけど，他人のデータを覗き見る痴漢野郎がいるぢゃん。
Google とか [Apple]({{< ref "/remark/2021/08/apples-mass-surveillance-plans.md" >}}) とか。
だから積極的には使わないようにしてたんだけど，これで踏ん切りがついたわ。

物置用のクラウド・ストレージ・サービスとして Google Drive をデータ暗号化した上で利用することにした。
これに伴い [Google One](https://one.google.com/) の[有料プラン](https://one.google.com/about/plans "プランと料金設定 - Google One")へアップグレードした。

運用中の各種クラウド・ストレージ・サービスに置いてあるファイルを整理しないとな。

## Git Server と SSH

パッケージ・センターで [Git Server] を見つけたのでインストールしてみたのだが...

どうやら SSH ベースのサーバのようだが

- [Git Server - Synology ナレッジセンター](https://kb.synology.com/ja-jp/DSM/help/Git/git?version=7)

そもそも [DS220j] は管理者権限を持つアカウントしか SSH を使えないようになっている。
なので [Git Server] へのアクセス専用にグループとユーザを作成し，専用の共有フォルダを設定する必要がある。

更に言うとユーザ・ホーム・サービスを有効にしておく必要がある。

{{< fig-img src="./user-home-service.png" link="./user-home-service.png" width="642" >}}

これがないと公開鍵認証ができないらしい。

- [Synology DiskStation で SSH 接続を公開鍵認証方式にする - Qiita](https://qiita.com/shimizumasaru/items/56474d98e723ea1b5ae3)

まぁ，ぼちぼち進めるか。

### 【2021-10-30 追記】 Synology Git Server に公開鍵認証でログインする

まずはおとなしくパスワード認証で入る。
前節で述べたように，あらかじめ専用のユーザとグループを作って SSH ログイン可能な状態にすること（今回はユーザ：グループ名を `git:git` とした[^admin]）。

[^admin]: Git Server 用に作成したユーザを `administrators` グループにも所属させて管理者権限を付与すること。

```text
$ ssh git@ds220j-hostname
git@ds220j-hostname's password: 
```

ログイン直後のホーム・ディレクトリはこんな感じ。

```text
$ pwd
/var/services/homes/git

$ ll
drwxrwxrwx+ 2 git  users 4096 Oct 30 00:00 .
drwxrwxrwx+ 9 root root  4096 Oct 30 00:00 ..
```

このままでは拙いので，ホーム・ディレクトリのパーミッションを変更する。

```text
chmod 755 /var/services/homes/git
```

次にホーム・ディレクトリ直下に `.ssh` ディレクトリを作成する。

```text
$ mkdir .ssh

$ ll
drwxr-xr-x  3 git  users 4096 Oct 30 00:00 .
drwxrwxrwx+ 9 root root  4096 Oct 30 00:00 ..
drwxr-xr-x  2 git  users 4096 Oct 30 00:00 .ssh
```

この `.ssh` ディレクトリ直下に `authorized_keys` ファイルを作り SSH 公開鍵を登録する。
ちなみに公開鍵は ECC 鍵が使える。
ECC 鍵を作ってサーバに登録する方法は以下の通り。

- [OpenSSH 鍵をアップグレードする]({{< ref "/remark/2020/06/upgrade-openssh-key.md" >}})

また SSH 鍵を GnuPG で作成・管理する方法は以下を参考にどうぞ。

- [OpenSSH の認証鍵を GunPG で作成・管理する]({{< ref "/openpgp/ssh-key-management-with-gnupg.md" >}})

`.ssh` ディレクトリおよび `.ssh/authorized_keys` ファイルのパーミッションに注意。

```text
$ chmod 700 ~/.ssh
$ chmod 600 ~/.ssh/authorized_keys
```

これで適当なターミナルを立ち上げて（作業中の SSH セッションは維持しておくこと。しくじってアクセスできなくなったら困るからね）

```text
$ ssh git@ds220j-hostname
```

でパスワード認証ではなく公開鍵認証が動き出したら成功である。

### 【2021-10-30 追記】 試しにリポジトリを作ってみる

Git Server 用に専用の共有フォルダを作る。
今回は `gitrepos` という名前で共有フォルダを作った。
前節の `git` ユーザで読み書きできるようにしておくこと（他のユーザはアクセスできなくてよい）。

さっそく `gitrepos` 共有フォルダに入ってみる。

```text
$ ssh git@ds220j-hostname
$ cd /volume1/gitrepos/
$ sudo chown git:git /volume1/gitrepos
$ chmod 755 /volume1/gitrepos
```

たぶん必須ではないが気持ち悪いので，所有者とパーミッションを変更して `git` ユーザ専用にした。
では，適当なベア・リポジトリを作ってみよう。

```text
$ mkdir hello-repos
$ cd hello-repos/
$ git init --bare
Initialized empty Git repository in /volume1/gitrepos/hello-repos/
```

これで空のベア・リポジトリができた。
リモートから SSH 接続でリポジトリを clone するには

```text
$ git clone ssh://git@ds220j-hostname/volume1/gitrepos/hello-repos
Cloning into 'hello-repos'...
warning: You appear to have cloned an empty repository.
```

とすればよい。

## 次は Tailscale か？

これは次の休みにでもやってみることにしよう。
もともと [Tailscale] は興味があってやってみようとは思っていたのだ。

- [100台まで無料のVPNサービス「tailscale」、リンクだけでマシンのシェアも可能!?【イニシャルB】 - INTERNET Watch](https://internet.watch.impress.co.jp/docs/column/shimizu/1303751.html)
- [Tailscale VPN を使ってみたので感想 | つくみ島だより](https://blog.tsukumijima.net/article/tailscale-vpn/)
- [リモートワーク向けにセキュアなVPNを構築【tailscale】 | 4b-media](https://4b-media.net/tailscale/)

## ブックマーク

- [RAIDの構成を「SHR」から「RAID 1」へ変更できますか？ | Synology Community](https://community.synology.com/jpn/forum/2/post/6)
- [知らないと損! Synology社員が教えるNASの裏ワザ【前編】セキュリティーと信頼性を高める | マイナビニュース](https://news.mynavi.jp/kikaku/20181221-741225/)
- [特別企画：知っておきたい！NAS購入から活用までの疑問あれこれ - デジカメ Watch](https://dc.watch.impress.co.jp/docs/review/special/1209064.html)
- [NASを使ったほうがいい人、使わないほうがいい人：Synologyで始めるNAS生活（1/3 ページ） - ITmedia PC USER](https://www.itmedia.co.jp/pcuser/articles/1603/18/news096.html)
- [Synology アカウントの取得 - アラコキからの Raspberry Pi 電子工作](https://arakoki70.com/?p=2317)
- [Ubuntu 20.04にNFSマウントをセットアップする方法 | DigitalOcean](https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nfs-mount-on-ubuntu-20-04-ja)
- [特別企画：NASに保存した写真データを効率的に管理…「Synology Photos」先行レビュー - デジカメ Watch](https://dc.watch.impress.co.jp/docs/review/special/1325567.html)

[DS220j]: https://www.synology.com/ja-jp/products/DS220j "DS220j | Synology Inc."
[Samba]: https://www.samba.org/ "Samba - opening windows to a wider world"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Cloud Sync]: https://www.synology.com/ja-jp/dsm/feature/cloud_sync "Cloud Sync | Synology Inc."
[Git Server]: https://www.synology.com/ja-jp/dsm/packages/Git?os_ver=7.0 "Git Server - アドオン パッケージ | Synology Inc."
[Tailscale]: https://tailscale.com/ "Tailscale · Best VPN Service for Secure Networks"

## 参考

{{% review-paapi "B0855LMP81" %}} <!-- Synology DS220j -->
{{% review-paapi "B08V8LNR2H" %}} <!-- HDD WD Red Plus -->
{{% review-paapi "B09GS5ZCHN" %}} <!-- Voice of Angel (太田貴子) -->
