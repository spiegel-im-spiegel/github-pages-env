+++
title = "ClamTk を導入する"
date =  "2024-06-30T15:53:58+09:00"
description = "GUI フロントエンドで（多分）簡単操作"
image = "/images/attention/kitten.jpg"
tags = [ "security", "malware", "risk", "management", "tools", "linux", "ubuntu", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

随分前に [ClamAV] を[導入][前]する話を書いたのだが，今回は GUI フロントエンドである [ClamTk] を導入する。

その前に [ClamAV] 手順が以前とは少し違ってるみたいなので，そこからやり直し。

## ClamAV のインストール

[ClamAV] は APT でインストール可能。
ここは以前と変わらず。

```text
$ sudo apt install clamav clamav-daemon libclamunrar
```

[前]のときはなかった `libclamunrar` は圧縮ファイルの中身やメール添付ファイルを検証するパッケージらしい。
[Ubuntu] では依存パッケージとして `libclamunrar11` もインストールされる。
Debian だと [`libclamunrar12`](https://packages-pkgmirror-csail.debian.org/sid/libclamunrar12 "Debian -- sid の libclamunrar12 パッケージに関する詳細") ってのがあるらいいけど [Ubuntu] の標準リポジトリにはなかった。

バージョンを確認しておこう。

```text
$ clamscan --version
ClamAV 1.0.5/27321/Sat Jun 29 17:39:26 2024
```

おー。
v1 に到達してたのか。

インストール後のサービスの状態を確認する。

```text
$ sudo systemctl status clamav-daemon.service
○ clamav-daemon.service - Clam AntiVirus userspace daemon
     Loaded: loaded (/usr/lib/systemd/system/clamav-daemon.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/clamav-daemon.service.d
             └─extend.conf
     Active: inactive (dead)
TriggeredBy: ○ clamav-daemon.socket
  Condition: start condition unmet at Sun 2024-06-30 09:19:59 JST; 2min 11s ago
             └─ ConditionPathExistsGlob=/var/lib/clamav/daily.{c[vl]d,inc} was not met
       Docs: man:clamd(8)
             man:clamd.conf(5)
             https://docs.clamav.net/
```

ありゃ，状態が `inactive` だな。
サービスを `start` させればいいのかな。

```text
$ sudo systemctl start clamav-daemon.service
$ sudo systemctl status clamav-daemon.service
● clamav-daemon.service - Clam AntiVirus userspace daemon
     Loaded: loaded (/usr/lib/systemd/system/clamav-daemon.service; enabled; preset: enabled)
    Drop-In: /etc/systemd/system/clamav-daemon.service.d
             └─extend.conf
     Active: active (running) since Sun 2024-06-30 09:37:30 JST; 5s ago
TriggeredBy: ● clamav-daemon.socket
       Docs: man:clamd(8)
             man:clamd.conf(5)
             https://docs.clamav.net/
    Process: 1408921 ExecStartPre=/bin/mkdir -p /run/clamav (code=exited, status=0/SUCCESS)
    Process: 1408923 ExecStartPre=/bin/chown clamav /run/clamav (code=exited, status=0/SUCCESS)
   Main PID: 1408925 (clamd)
      Tasks: 1 (limit: 37666)
     Memory: 922.0M (peak: 922.0M)
        CPU: 5.803s
     CGroup: /system.slice/clamav-daemon.service
             └─1408925 /usr/sbin/clamd --foreground=true
```

んー，大丈夫かな。

設定ファイルは `/etc/clamav/clamd.conf` ができているはず。
そのままでも問題ない。
強いて挙げるなら `MaxThreads` の値が CPU コア数になってると思うので，適当に間引いた値にするとストレスになりにくい。

On-Access スキャンの設定については以下を参照のこと。

- [On-Access Scanning - ClamAV Documentation](https://docs.clamav.net/manual/OnAccess.html)

設定変更したらサービスを `restart` してね。

### ClamAV 動作確認

以下のページからテスト用のファイルをダウンロードする。

- [Download Anti Malware Testfile - EICAR](https://www.eicar.org/download-anti-malware-testfile/)

次に，以下のコマンドを実行しダウンロードしたファイルをスキャンさせる。

```text
$ sudo clamscan -r --bell -i --max-filesize=100M -l scan.log .
/path/to/eicar-test/eicar.com.txt: Win.Test.EICAR_HDB-1 FOUND
/path/to/eicar-test/eicarcom2.zip: Win.Test.EICAR_HDB-1 FOUND
/path/to/eicar-test/eicar_com.zip: Win.Test.EICAR_HDB-1 FOUND
/path/to/eicar-test/eicar.com: Win.Test.EICAR_HDB-1 FOUND

----------- SCAN SUMMARY -----------
Known viruses: 8695433
Engine version: 1.0.5
Scanned directories: 2
Scanned files: 11
Infected files: 4
Data scanned: 309.96 MB
Data read: 44.05 MB (ratio 7.04:1)
Time: 42.667 sec (0 m 42 s)
Start Date: 2024:06:30 09:53:06
End Date:   2024:06:30 09:53:49
```

んー，これも問題なさそうかな。
ちなみにオプションの意味は

- `-r` は指定したディレクトリを再帰的に走査する
- `--bell` は malware を見つけたらベルを鳴らす
- `-i` malware に感染したファイルのみを報告する
- `--max-filesize` は指定したファイルより大きいファイルはスキップする（既定値は `25M`）
- `-l` はログファイルの指定

という感じ。

## ClamTk のインストール

[ClamTk] のインストールも APT でOK。

```text
$ sudo apt install clamtk
```

ただしそのままだと色々と不具合があるらしいので，以下のページを参考に，いくつかのファイルを変更する。

- [【Ubuntu】ウィルス対策。ClamTkが上手く動作しないときの対処あれこれ。 - freefielder.jp](https://freefielder.jp/blog/2017/05/ubuntu-clamtk-trouble.html)

まずは `/usr/bin/clamtk` (Perl ソース) ファイルから。

以下の記述を探す（45行目あたり？）。

```perl
setlocale( LC_ALL, '' );
# setlocale( LC_ALL, 'C' );
```

これのコメントを反転させる。

```perl
# setlocale( LC_ALL, '' );
setlocale( LC_ALL, 'C' );
```

もうひとつ。
`/usr/share/perl5/ClamTk/Scan.pm` (これも Perl ソース) ファイルについても以下の記述を探す（200行目あたり？）。

```perl
# Try to avoid scanning emails...
$directive .= ' --scan-mail=no';
```

これの `no` の部分を `yes` に変更する。

```perl
# Try to avoid scanning emails...
$directive .= ' --scan-mail=yes';
```

これって APT で [ClamTk] を更新する度にやるのかねぇ。
面倒な...

まぁ，いいや。
早速起動してみよう。

{{< fig-img src="./clamtk.png" title="ClamTk" link="./clamtk.png" width="554" >}}

修正のせいで英語表記になってるのかな。
これはしょうがないか。

まずは “Update Assistant” から

{{< fig-img src="./update-assistant.png" title="Update Assistant" link="./update-assistant.png" width="554" >}}

“I would like to update signitures myself” にチェックを入れて Apply する。
これは signiture データの更新を後述する “Scheduler” で行うため。

次に “Update” を開く。

{{< fig-img src="./update.png" title="Update" link="./update.png" width="554" >}}

“Check for Updates” を押下して signiture データを更新する（上の図は更新後の状態）。

更に “Whitelist” を開いて `/home/username` ディレクトリ以下でスキャンから除外するディレクトリを指定する。

これでようやく “Scheduler” の設定ができる。

{{< fig-img src="./scheduler.png" title="Scheduler" link="./scheduler.png" >}}

これでスキャンと signiture データの更新を行う時刻を指定できる。
“Scheduler” の設定は crontab に反映されている。
こんな感じ[^s1]。

[^s1]: 私は裏で色々バッチ処理を行ってる関係で，マシンは基本的に立ち上げっぱなしである。

```text
$ crontab -l
1 1 * * * /usr/bin/freshclam --datadir=/home/username/.clamtk/db --log=/home/username/.clamtk/db/freshclam.log # clamtk-defs
10 1 * * * /usr/bin/clamscan --exclude-dir=/home/username/.clamtk/viruses --exclude-dir=smb4k --exclude-dir=/run/user/username/gvfs --exclude-dir=/home/username/.gvfs --exclude-dir=.thunderbird --exclude-dir=.mozilla-thunderbird --exclude-dir=.evolution --exclude-dir=Mail --exclude-dir=kmail --database=/home/username/.clamtk/db -i -r /home/username --log="$HOME/.clamtk/history/$(date +\%b-\%d-\%Y).log" 2>/dev/null # clamtk-scan
```

crontab って一見便利ぽいけど，コマンドライン・オプションの指定ミスで簡単に設定が消えちゃうし [ClamTk] みたいにツールが設定を行う場合もあるので，うっかり触れないんだよなー

愚痴はともかく，これで設定完了。
とりあえず “Scan a directory” で malware がないか確認しておこう。

お疲れ様でした。

## ブックマーク

- [Linux向けウィルス対策ソフトClamAVの使い方](https://zenn.dev/gladevise/articles/clamav-usage)
- [Ubuntu 22.04 LTSでウイルス対策ソフトを実行 ! | 与太郎のLinux道楽](https://ameblo.jp/l-yotarou/entry-12836559211.html)
- [Linux MintにClamAVのGUIフロントエンドであるClamTkを導入する  |  サブログ -sub log-](https://sub-log.jp/2022/12/04/linux-mint%e3%81%abclamav%e3%81%aegui%e3%83%95%e3%83%ad%e3%83%b3%e3%83%88%e3%82%a8%e3%83%b3%e3%83%89%e3%81%a7%e3%81%82%e3%82%8bclamtk%e3%82%92%e5%b0%8e%e5%85%a5%e3%81%99%e3%82%8b/)

[前]: {{< ref "/remark/2019/05/clamav-for-ubuntu.md" >}} "Ununtu に ClamAV を導入する"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[ClamAV]: https://www.clamav.net/ "ClamavNet"
[ClamTk]: https://gitlab.com/dave_m/clamtk/ "Dave M / clamtk · GitLab"

## 参考図書

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
