+++
title = "Ununtu に ClamAV 導入する"
date =  "2019-05-06T17:39:20+09:00"
description = "こんなに面倒臭いとは思わなかった。もうしばらく運用してから評価してみるですよ。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "malware", "risk", "management", "tools", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[ClamAV] を [Ubuntu] に導入したのだが，思いのほか面倒臭かったので，覚え書きとして記しておく。

## [ClamAV] をインストールする

[ClamAV] のインストール自体は APT (Advanced Package Tool) で簡単にできる。

```text
$ sudo apt install clamav clamav-daemon
```

なお `clamav-daemon` は [ClamAV] をサービス化するためのパッケージである。

現時点（2019年5月）で [ClamAV] の最新は 0.101.x なのだが APT で提供されているバージョンは 0.100.x の古いバージョンのようだ。
一応，以下の脆弱性への対応はしてあるみたいだが

- [ClamAV® blog: ClamAV 0.101.2 and 0.100.3 patches have been released!](https://blog.clamav.net/2019/03/clamav-01012-and-01003-patches-have.html)

このバージョンの古さは後々に影響を及ぼすかもしれない。
Debian には最新の[ソースパッケージがある](https://packages.debian.org/source/unstable/clamav)ようだが，セキュリティ製品だしソースコードからビルドするのは流石に...

## データベースを更新する

まずは手動でデータベースを更新する。

設定ファイル `/etc/clamav/freshclam.conf` から `NotifyClamd` の行をコメントアウトした上で `freshclam` コマンドを実行する。
コマンドラインで書くとこんな感じ。

```text
$ sudo systemctl stop clamav-freshclam

$ sudo sed -i -e "s/^NotifyClamd/#NotifyClamd/g" /etc/clamav/freshclam.conf

$ sudo freshclam
Sun May  5 13:36:34 2019 -> ClamAV update process started at Sun May  5 13:36:34 2019
Sun May  5 13:36:34 2019 -> ^Your ClamAV installation is OUTDATED!
Sun May  5 13:36:34 2019 -> ^Local version: 0.100.3 Recommended version: 0.101.2
Sun May  5 13:36:34 2019 -> DON'T PANIC! Read https://www.clamav.net/documents/upgrading-clamav
Sun May  5 13:36:34 2019 -> main.cvd is up to date (version: 58, sigs: 4566249, f-level: 60, builder: sigmgr)
Sun May  5 13:36:34 2019 -> daily.cvd is up to date (version: 25439, sigs: 1562310, f-level: 63, builder: raynman)
Sun May  5 13:36:34 2019 -> bytecode.cvd is up to date (version: 328, sigs: 94, f-level: 63, builder: neo)

$ sudo sed -i -e "s/^#NotifyClamd/NotifyClamd/g" /etc/clamav/freshclam.conf

$ sudo systemctl start clamav-freshclam
```

データベースの更新自体はうまく行ってるようだが `OUTDATED` とか言われてるよ `orz`

...見なかったことにしよう。

なお `clamav-freshclam.service` の状態を確認するには以下のコマンドを起動すればよい。

```text
$ sudo systemctl status clamav-freshclam.service
```

## 手動で動作確認

Malware 対策ソフト用の無害なウイルス（笑）があるので，これを使って `clamscan` コマンドで動作確認する。

```text
$ curl http://www.eicar.org/download/eicar.com -O

$ clamscan --infected --remove ./eicar.com
./eicar.com: Eicar-Test-Signature FOUND
./eicar.com: Removed.

----------- SCAN SUMMARY -----------
Known viruses: 6120222
Engine version: 0.100.3
Scanned directories: 0
Scanned files: 1
Infected files: 1
Data scanned: 0.00 MB
Data read: 0.00 MB (ratio 0.00:1)
Time: 20.181 sec (0 m 20 s)
```

よーし，うむうむ，よーし。

## Scan On-Access の設定

ファイルアクセスへの常時監視を行う Scan On-Access を有効にするために `/etc/clamav/clamd.conf` ファイルを修正する。
本当は `ScanOnAccess` を `true` にするだけでいい筈なのだが[^rs1]，以下のコマンドで

[^rs1]: サービスのリスタートには `sudo systemctl restart clamav-daemon.service` とすればよい。

```text
$ sudo systemctl status clamav-daemon.service
```

サービスの状態を確認すると以下のログを吐いて停止していた。

```text
ERROR: ScanOnAccess: fanotify_init failed: Operation not permitted
ScanOnAccess: clamd must be started by root
```

`/etc/clamav/clamd.conf` ファイルの他の部分にも手を入れる必要があるようだ。
以下に変更・追記が必要な項目を挙げておく。

{{< highlight text "hl_lines=1-3" >}}
LocalSocketGroup root
User root
ScanOnAccess true
OnAccessMountPath /home/username
OnAccessExcludePath /home/username/nocheck
# VirusEvent /usr/local/bin/clamd-response
OnAccessPrevention false
OnAccessExtraScanning true
OnAccessExcludeUID 0
{{< /highlight >}}

最初の3行が変更が必要な項目，以降が追加項目である。
`OnAccessMountPath` でスキャン対象のディレクトリを `OnAccessExcludePath` で除外対象を指定する。
どちらも複数指定できる。

### VirusEvent が効かない！？

上の設定で `VirusEvent` をコメントアウトしているが，どうやら 0.100 では Scan On-Access を有効にしても `VirusEvent` の設定が効かないようなのだ（0.101 はどうなんだろう）。

`VirusEvent` には malware を発見した際に起動するコマンドを指定するのだが，これが動かないってかなり致命的じゃないのか？

次善の策として以下のスクリプトを書いて cron で回すことにした。
Malware を検出すれば画面に通知が表示されるはずである。
cron で回してるだけなのでタイムラグが発生する。

```bash
#!/bin/bash

logfile="/home/username/.local/log/clamd-found.log"
if [ -r $logfile ]; then
  PRV_CNT=`cat $logfile | wc -l`
else
  PRV_CNT=0
fi
CNT=`grep ScanOnAccess /var/log/syslog | grep FOUND | grep -v "(null)" | tee $logfile | wc -l`
if [ $CNT != $PRV_CNT ]; then
  DATESTR=`date`
  sudo -u username DISPLAY=:0 DBUS_SESSION_BUS_ADDRESS=unix:path=/run/user/1000/bus /usr/bin/notify-send -t 10000 "$DATESTR Virus Found $CNT"
fi
```

ダサいけどしょうがない。
ちなみに `username` の部分には cron を回すログインユーザ名を， `1000` の部分にはそのユーザ ID をセットする[^uid1]。

[^uid1]: ユーザ ID が分からない場合は `id username` とコマンドを打てば分かるだろう。

## もう一度，動作確認

`clamav-daemon.service` がうまく機能するか試してみよう。

先ほどの無害なウイルス（笑）をダウンロードする。

```text
$ curl http://www.eicar.org/download/eicar.com -O
```

これでログ等を見て検知しているか確認する（上の通知も出れば完璧）。
確認できたらとっとと削除すること。

今回は `clamdscan` のほうを使ってみよう。

```text
$ clamdscan --infected --remove ./eicar.com
/home/spiegel/work/./eicar.com: Eicar-Test-Signature FOUND
/home/spiegel/work/./eicar.com: Removed.

----------- SCAN SUMMARY -----------
Infected files: 1
Time: 0.365 sec (0 m 0 s)
```

圧倒的に速くなった。
まぁ，バックグラウンドにサービスがいるんだから当たり前だが（笑）

## Thunderbird と連携したかったが...

MUA である Thunderbird と連携できるアドオンがあると聞いたのだが

- [clamdrib LIN :: Thunderbird 向けアドオン](https://addons.thunderbird.net/ja/thunderbird/addon/clamdrib-lin/)

うまくインストールできない。

[clamdrib LIN] の最新バージョンは2017年のもので最近のバージョン 60 以降には対応してないっぽい（互換性チェック回避のオプションも試したがうまく行かなかった）。
諦めるしかないようだ。

こんなんばっかだな ＞ [ClamAV]

## 「ウイルス対策ソフト」は必要か？

なんだか知らないが「Linux にウイルス対策ソフトは必要か？」などという頭の悪い議論があるそうで，しかも「必要ない」という人の主張は「Linux ユーザは Windows や macOS に比べて規模が小さいから狙われにくい」という，これまた頭が悪いとしか言いようがないものらしい。

「ウイルス」という表現は古臭くて的を得ていないので，この記事では色々ひっくるめて malware (malicious software) という表現に統一させてもらうが，ぶっちゃけて言うなら使っている OS にかかわらず「**ユーザレベルでは malware 対策はセキュリティ管理の中核ではない**」というのが多分正しい。

じゃあ malware 対策は必要ないのかと言えば「さにあらず」で，しかしそういったものは外部化され不可視になっている。
今どきの言葉で言うなら “Security as a Service” とでもいうような状況である。
そしてそのプラットフォーム OS として Linux 等のセキュリティ要件はむしろ高まっていると言えるだろう。
今回 [ClamAV] を触ってみて感じたことだが，一般ユーザにとって [ClamAV] がちょっと残念な感じになっているのはサーバ用途に最適化されつつあるからではないかと思うのだ。

セキュリティのトレンドは「防衛」中心から「監視」中心にシフトしていて，単純に malware をバラ撒いただけではすぐに検知され対策されてしまう。
故に攻撃側も時間をかけてでも密やかに確実にターゲットを追い詰める social engineering を駆使した各種ターゲティング攻撃へとシフトしている。

Phishing メールやそれに含まれる malware 検知などはサービスプロバイダが（有料のものも含めて）やってくれるし， Web についても Phishing サイトやマイニング・コード等を含んだサイトはブラウザレベルでかなりいい感じにブロックできるようになりつつある[^ps1]。
そういう意味で私にとって malware 対策ソフトは「うっかり転んだときのための少額保険」みたいなもので，それ以外ではあまり出しゃばって欲しくなかったりする。

[^ps1]: はっきり言ってフェイク・ニュースなんかより Phishing ページなどセキュリティ・リスクの高いコンテンツに対する監視の方にもっとリソースを割いてほしいのだけど。かといって中間車攻撃で暗号通信を覗き見るなんてのは論外だが。

実際にそれまで使ってた Windows 7 でも “Microsoft Security Essentials” で必要十分な性能だったし [ClamAV] でも同程度の性能を満たしていれば，今回はそれでよかったのだ。
それがこんなに面倒臭いとは思わなかった。

かといって malware 対策ソフトそのものに金銭は払いたくないし[^fd1]，もうしばらく運用してから評価してみるですよ。

[^fd1]: Android 端末には [F-Secure 社の FREEDOME](https://www.f-secure.com/en/home/products/freedome) を入れている。 FREEDOME 自体は VPN ソフトだが malware 対策機能も付いていて追跡コードもある程度ブロックしてくれる優れもの。実はこれの Linux 版があれば買ってもよかったのだがなかったのだ。仮想ネットワークについては I2P を試してみるか，と思っていたりする。

## ブックマーク

- [The 8 Best Free Anti-Virus Programs for Linux](https://www.tecmint.com/best-antivirus-programs-for-linux/)
    - [Linux向けアンチウイルスソフト8選 | マイナビニュース](https://news.mynavi.jp/article/20170420-a074/)
- [Clam Antivirusに関するメモ](http://clamav-jp.osdn.jp/jdoc/clamav.html)
- [ClamAV - Community Help Wiki](https://help.ubuntu.com/community/ClamAV)
- [ClamAVをUbuntu MATE 18.04 LTSにセットアップ](https://hnakamur.github.io/blog/2018/05/21/setup-clamav-on-ubuntu-mate-18.04-lts/)
- [Ubuntu 18.04 LTS DesktopでClamAVによるウィルスチェックを実行 – LAB4ICT](https://lab4ict.com/system/archives/934)
- [ClamAV - ArchWiki](https://wiki.archlinux.org/index.php/ClamAV)

- [Ubuntu16.04(さくらVPS)で１からcronの設定をし、pythonプログラムを定期実行する - Qiita](https://qiita.com/gano/items/802519add83c524e3019)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[ClamAV]: https://www.clamav.net/ "ClamavNet"
[clamdrib LIN]: https://addons.thunderbird.net/ja/thunderbird/addon/clamdrib-lin/ "clamdrib LIN :: Thunderbird 向けアドオン"
