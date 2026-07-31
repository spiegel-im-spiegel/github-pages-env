+++
title = "そろそろ Vuls を唱えるか"
date =  "2019-06-02T22:22:06+09:00"
description = "セットアップ自体はうまく行ったが，このままでは使えないなぁ。しょうがない。少しずつ調べてみるか。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "vulnerability", "risk", "management", "tools", "vuls", "linux", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Ubuntu] 関連の脆弱性情報は以下のサイトで収集できる。

- [Debian -- Security Information](https://www.debian.org/security/)
- [Ubuntu security notices](https://usn.ubuntu.com/)

これらのサイトの脆弱性情報はとても有用なのだが，パッケージ単位の情報なのでパッケージ間の依存関係が分かってないと何がどう影響するか分かりにくい。
更に上のサイトを見ただけでは深刻度が分からない（CVE 番号から調べることは可能）。
更に更に言うと [Ubuntu] の各パッケージは古いバージョンに対してバックポートティングをかけて脆弱性を手当している場合もあり，パッケージのバージョン番号を見ただけでは分かりにくかったりする。

つまり CVE や OVAL や各ディストリビューションが提供する情報を統合管理する必要があるのだが，手動でこれをやるのは骨が折れる。
今回は個人のデスクトップ PC だけだから管理も比較的楽だけど，複数のサーバ（クラウド環境も併せて）を管理するとか気が遠くなる。
つか，もはや手動で管理とかあり得ないだろう。

というわけで，そろそろ [Vuls] を唱えてみるか。

## [Vuls] とは

- [Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD](https://vuls.io/)

[Vuls] は脆弱性の検知を行うツールで，ほぼ [Go 言語]で組まれているのが特徴である。
複数のサーバに対して脆弱性の有無を一括スキャンすることが可能という実にサーバ管理者に優しいつくりになっている。

{{< fig-img src="./vuls-abstract.png" title="future-architect/vuls: Agent-less vulnerability scanner for Linux/FreeBSD/WordPress/Programming language libraries/Network devices" link="https://github.com/future-architect/vuls" width="1100" >}}

[Vuls] 本体はシングル・バイナリだが CVE や OVAL などの脆弱性情報を収集するツール群と組み合わせて使うことを前提としている。
また最近では [Vuls] を含むツール群を SaaS 化した FutureVuls なる商用サービスも登場している。

- [FutureVuls - Vulsクラウドサービス [Vuls SaaS]](https://vuls.biz/)

今回は [Vuls] の基本機能でローカル PC の脆弱性管理を行ってみたいと思う。

## 【前準備】 ディレクトリの作成

まず前準備として [Vuls] を含むツール群が出力するデータベースを置くディレクトリ `~/vuls` を作成しておく。

```text
$ mkdir ~/vuls
```

次に各ツールが吐き出すログの出力先ディレクトリを作っておく。

```text
$ sudo mkdir /var/log/vuls
$ sudo chown username:username /var/log/vuls
$ sudo chmod 700 /var/log/vuls
```

```text
$ sudo mkdir /var/log/gost
$ sudo chown username:username /var/log/gost
$ sudo chmod 700 /var/log/gost
```

```text
$ sudo mkdir /var/log/go-exploitdb
$ sudo chown username:username /var/log/go-exploitdb
$ sudo chmod 700 /var/log/go-exploitdb
```

## CVE 情報の収集

CVE 情報の収集には [go-cve-dictionary] を使う。

### [go-cve-dictionary] のセットアップ

ソースコードからのインストール方法については色々説明があるが，最近のバージョンはバイナリも提供しているようである。

- [Releases · kotakanbe/go-cve-dictionary · GitHub](https://github.com/kotakanbe/go-cve-dictionary/releases)

これをありがたく使わせてもらおう。
いやぁ [Go 言語]のツールは取り回しが楽でいいね。

### [go-cve-dictionary] による CVE 情報の収集

たとえば2002年以降の CVE 情報を収集し SQLite のデータベースファイル `~/vuls/cve.sqlite3` へ格納するには以下の一行スクリプトを実行する。

```text
$ for i in `seq 2002 $(date +"%Y")`; do go-cve-dictionary fetchnvd -dbpath ~/vuls/cve.sqlite3 -years $i; done
```

ものごっつ時間がかかるので，ここらでお茶の時間にしよう。

## OVAL 情報の収集

OVAL 情報の収集には [goval-dictionary] を使う。

### [goval-dictionary] のセットアップ

[goval-dictionary] もバイナリがリリースされているのだが，バイナリ・リリースされているバージョンでは [Ubuntu] 19 に対応していないようなので，最新ソースを取ってきてビルドする必要がある。

```text
$ cd $GOPATH/src/github.com/kotakanbe
$ git clone https://github.com/kotakanbe/goval-dictionary.git
$ cd gost/
$ make install
```

実行モジュールは `$GOPATH/bin` ディレクトリにインストールされる。

### [goval-dictionary] による OVAL 情報の収集

ディストリビューションおよびそのバージョンごとに  OVAL 情報を収集する。
たとえば [Ubuntu] なら

```text
$ goval-dictionary fetch-ubuntu -dbpath ~/vuls/oval.sqlite3 12 14 16 18 19
```

などとする。
他にも `fetch-alpine`, `fetch-amazon`, `fetch-debian`, `fetch-oracle`, `fetch-redhat`, `fetch-suse` といったディストリビューションに対応している。

## ディストリビューション別セキュリティ情報の収集

ディストリビューション毎のセキュリティ情報の収集には [gost (go-security-tracker)] を使う。

### [gost] のセットアップ

[gost] はソースコードのみのリリースなのでビルドを行う。

```text
$ cd $GOPATH/src/github.com/knqyf263
$ git clone https://github.com/knqyf263/gost.git
$ cd gost/
$ make install
```

実行モジュールは `GOPATH/bin` ディレクトリにインストールされる。

### [gost] によるセキュリティ情報の収集

[gost] では `redhat`, `debian`, `microsoft` のセキュリティ情報の収集を行う。

```text
$ gost fetch debian --dbpath ~/vuls/gost.sqlite3
```

[Ubuntu] のセキュリティ情報収集については TODO 扱いになっているようだ。
今後に期待しよう。

## Exploit 情報の収集

Exploit 情報の収集には [go-exploitdb] を使う。

### [go-exploitdb] のセットアップ

[go-exploitdb] はソースコードのみのリリースなのでビルドを行う。

```text
$ cd $GOPATH/src/github.com/mozqnet
$ git clone https://github.com/mozqnet/go-exploitdb.git
$ cd go-exploitdb/
$ make install
```

実行モジュールは `GOPATH/bin` ディレクトリにインストールされる。

### [go-exploitdb] によるセキュリティ情報の収集

[go-exploitdb] では `awesomepoc`, `exploitdb`, `githubrepos` を対象に Exploit 情報の収集を行う。

```text
$ go-exploitdb fetch exploitdb --dbpath ~/vuls/go-exploitdb.sqlite3
```

## [Vuls] でローカル PC をスキャンする

さて，いよいよ [Vuls] を唱えるときが来た（笑）

### [Vuls] のセットアップ

[Vuls] はバイナリがリリースされている。

- [Releases · future-architect/vuls](https://github.com/future-architect/vuls/releases)

ありがたや。

次にスキャン対象のマシンへのアクセスを定義する `~/vuls/config.toml` 設定ファイルを作成する。
今回はローカル PC のみのスキャンなので以下のような記述にする。

```toml
[servers]

[servers.localhost]
host = "localhost"
port = "local"
```

スキャン対象は複数指定可能で，たとえば対象に ssh 接続する場合は以下のような内容になるようだ。

```toml
[servers.remotehost]
host         = "remotehost"
port        = "22"
user        = "username"
keyPath     = "/home/username/.ssh/id_rsa"
```

設定が正しいかどうかチェックしておこう。

```text
$ cd ~/vuls
$ vuls configtest
[Jun  2 18:07:05]  INFO [localhost] Validating config...
[Jun  2 18:07:05]  INFO [localhost] Detecting Server/Container OS... 
[Jun  2 18:07:05]  INFO [localhost] Detecting OS of servers... 
[Jun  2 18:07:05]  INFO [localhost] (1/1) Detected: localhost: ubuntu 19.04
[Jun  2 18:07:05]  INFO [localhost] Detecting OS of containers... 
[Jun  2 18:07:05]  INFO [localhost] Checking Scan Modes...
[Jun  2 18:07:05]  INFO [localhost] Checking dependencies...
[Jun  2 18:07:05]  INFO [localhost] Dependencies... Pass
[Jun  2 18:07:05]  INFO [localhost] Checking sudo settings...
[Jun  2 18:07:05]  INFO [localhost] sudo ... No need
[Jun  2 18:07:05]  INFO [localhost] It can be scanned with fast scan mode even if warn or err messages are displayed due to lack of dependent packages or sudo settings in fast-root or deep scan mode
[Jun  2 18:07:05]  INFO [localhost] Scannable servers are below...
localhost 
```

んー，エラーにはなっていないみたいだし，こんな感じでいいのかな。

### [Vuls] でローカル PC をスキャンする

`configtest` も問題なさそうだし，実際にスキャンをかけてみよう。

```text
$ cd ~/vuls
$ vuls scan
[Jun  2 20:24:26]  INFO [localhost] Start scanning
[Jun  2 20:24:26]  INFO [localhost] config: /home/username/vuls/config.toml
[Jun  2 20:24:26]  INFO [localhost] Validating config...
[Jun  2 20:24:26]  INFO [localhost] Detecting Server/Container OS... 
[Jun  2 20:24:26]  INFO [localhost] Detecting OS of servers... 
[Jun  2 20:24:26]  INFO [localhost] (1/1) Detected: localhost: ubuntu 19.04
[Jun  2 20:24:26]  INFO [localhost] Detecting OS of containers... 
[Jun  2 20:24:26]  INFO [localhost] Checking Scan Modes... 
[Jun  2 20:24:26]  INFO [localhost] Detecting Platforms... 
[Jun  2 20:24:27]  INFO [localhost] (1/1) localhost is running on other
[Jun  2 20:24:27]  INFO [localhost] Scanning vulnerabilities... 
[Jun  2 20:24:27]  INFO [localhost] Scanning vulnerable OS packages...
[Jun  2 20:24:27]  INFO [localhost] Scanning in fast mode


One Line Summary
================
localhost	ubuntu19.04	2173 installed




To view the detail, vuls tui is useful.
To send a report, run vuls report -h.
```

よしよし。
うまく行ったようである。

スキャン結果を簡易表示してみよう。

```text
$ cd ~/vuls
$ vuls report -format-one-line-text
[Jun  2 20:30:18]  INFO [localhost] Validating config...
[Jun  2 20:30:18]  INFO [localhost] Loaded: /home/username/vuls/results/2019-06-02T20:30:15+09:00
[Jun  2 20:30:18]  INFO [localhost] Validating db config...
INFO[0000] -cvedb-type: sqlite3, -cvedb-url: , -cvedb-path: /home/username/vuls/cve.sqlite3 
INFO[0000] -ovaldb-type: sqlite3, -ovaldb-url: , -ovaldb-path: /home/username/vuls/oval.sqlite3 
INFO[0000] -gostdb-type: sqlite3, -gostdb-url: , -gostdb-path: /home/username/vuls/gost.sqlite3 
INFO[0000] -exploitdb-type: sqlite3, -exploitdb-url: , -exploitdb-path: /home/username/vuls/go-exploitdb.sqlite3 
INFO[06-02|20:30:18] Opening DB.                              db=sqlite3
INFO[06-02|20:30:18] Migrating DB.                            db=sqlite3
INFO[06-02|20:30:18] Opening Database.                        db=sqlite3
INFO[06-02|20:30:18] Migrating DB.                            db=sqlite3
[Jun  2 20:30:18]  INFO [localhost] OVAL is fresh: ubuntu 19.04 
[Jun  2 20:30:26]  INFO [localhost] localhost: 220 CVEs are detected with OVAL
[Jun  2 20:30:26]  INFO [localhost] localhost: 0 CVEs are detected with CPE
[Jun  2 20:30:26]  INFO [localhost] localhost: 0 CVEs are detected with GitHub Security Alerts
[Jun  2 20:30:26]  INFO [localhost] localhost: 0 unfixed CVEs are detected with gost
[Jun  2 20:30:26]  INFO [localhost] Fill CVE detailed information with CVE-DB
[Jun  2 20:30:27]  INFO [localhost] Fill exploit information with Exploit-DB
[Jun  2 20:30:27]  INFO [localhost] localhost: 13 exploits are detected
[Jun  2 20:30:27]  INFO [localhost] localhost: en: 0, ja: 0 alerts are detected


One Line Summary
================
localhost	Total: 220 (High:26 Medium:172 Low:22 ?:0)	0/220 Fixed	2173 installed	13 exploits	en: 0, ja: 0 alerts
```

ありゃりゃ。
最新状態にしているにも関わらず220もひっかかるとか。

もう少し詳細に見るには TUI モードにするとよい[^tui]。

[^tui]: TUI モードを終了するには `ctrl+c` を入力すれば良い。

```text
$ cd ~/vuls
$ vuls tui
```

{{< fig-img src="./vuls-tui.png" link="./vuls-tui.png" width="1272" >}}

うーむ。
どうも `report` がまともに機能してないっぽい？ これってやっぱ [gost] が [Ubuntu] に対応してないせいなのかな（状態のほぼ全てが `unfixed` になっている）。

このままではちょっと使えないなぁ。
しょうがない。
少しずつ調べてみるか。

今回はここまで。

## ブックマーク

- [あなたのサーバは本当に安全ですか？今もっともイケてる脆弱性検知ツールVulsを使ってみた - Qiita](https://qiita.com/sadayuki-matsuno/items/0bb8bb1689425bb9a21c)
- [FutureVuls(脆弱性管理サービス) Advent Calendar 2018 - Qiita](https://qiita.com/advent-calendar/2018/futurevuls)
- [Tutorial - Local Scan Mode · Vuls](https://vuls.io/docs/ja/tutorial-local-scan.html)
- [脆弱性検知ツールVulsの、きちんと動く構築手順](https://zenn.dev/myuki/articles/d0c1f780d45e8b)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Vuls]: https://vuls.io/ "Vuls · Agentless Vulnerability Scanner for Linux/FreeBSD"
[go-cve-dictionary]: https://github.com/kotakanbe/go-cve-dictionary "kotakanbe/go-cve-dictionary: Build a local copy of CVE (NVD and Japanese JVN). Server mode for easy querying."
[goval-dictionary]: https://github.com/kotakanbe/goval-dictionary "kotakanbe/goval-dictionary: Build a local copy of OVAL. Server mode for easy querying."
[gost]: https://github.com/knqyf263/gost "knqyf263/gost: Build a local copy of Security Tracker. Notify via E-mail/Slack if there is an update."
[gost (go-security-tracker)]: https://github.com/knqyf263/gost "knqyf263/gost: Build a local copy of Security Tracker. Notify via E-mail/Slack if there is an update."
[go-exploitdb]: https://github.com/mozqnet/go-exploitdb "mozqnet/go-exploitdb"
