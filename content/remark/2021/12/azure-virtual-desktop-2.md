+++
title = "Azure Virtual Desktop で遊ぶ"
date =  "2021-12-06T22:15:10+09:00"
description = "Tailscale で NAS とも繋がった。"
image = "/images/attention/kitten.jpg"
tags = [ "windows", "azure", "nas", "vpn", "tailscale", "install", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[前回]({{< relref "./azure-virtual-desktop.md" >}} "ようやく Azure Virtual Desktop を導入できた")でようやく [Azure Virtual Desktop] が導入できたので，これを使って Windows でボチボチ遊んでいこう。

## とりあえず [Scoop] と [NYAGOS] は入れておく

なにはともあれ [Scoop] と [NYAGOS] は入れておかないとね。

- [scoop / nyagos で始めるコマンドライン生活](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146)

この辺は問題なし。

## winget を試してみる

[winget] は Microsoft 公式のアプリケーション・パッケージ・マネージャで，たぶん [Chocolatey] あたりと置き換えて使えるように構成されていると思われる。

- [microsoft/winget-cli: Windows Package Manager CLI (aka winget)][winget]
- [winget ツールを使用したアプリケーションのインストールと管理 | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows/package-manager/winget/)

せっかく真っサラな環境が手に入ったんだから [winget] を試してみるとしよう。
と思ったが，最近の Windows には最初から入ってるのね。

```text
$ winget -v
v1.1.12986
```

ない場合は Microsoft Store から入れるのがいいようだ。
以下の「アプリ インストーラー」を入れる。

{{< fig-img src="./winget-by-ms-store.png" link="./winget-by-ms-store.png" width="1131" >}}

とりあえず git があるか検索してみよう。

```text
$ winget search git
'msstore' ソースを使用するには、使用する前に次の契約を表示する必要があります。
Terms of Transaction: https://aka.ms/microsoft-store-terms-of-transaction
ソースが正しく機能するには、現在のマシンの地域を送信する必要があります。

すべてのソース契約条件に同意しますか?
[Y] はい  [N] いいえ:
```

おうふ。
まぁ，いいか。
`[Y]` を入力して先に進む。

```text
名前                                  ID                                 バージョン    一致     ソース
-------------------------------------------------------------------------------------------------------
My Git                                9NLVK2SL2SSP                       Unknown                msstore
GitFiend                              9NMNKLTSZNKC                       Unknown                msstore
GitCup                                9NBLGGH4XFHP                       Unknown                msstore
GitVine                               9P3BLC2GW78W                       Unknown                msstore
GitIt                                 9NBLGGH40HV7                       Unknown                msstore
GitHub Zen                            9NBLGGH4RTK3                       Unknown                msstore
GitLooker                             9PK6TGX9T87P                       Unknown                msstore
Git                                   Git.Git                            2.34.1                 winget
GitNote                               zhaopengme.gitnote                 3.1.0         Tag: git winget
Agent Git                             Xidicone.AgentGit                  1.85          Tag: Git winget
TortoiseSVN                           TortoiseSVN.TortoiseSVN            1.14.29085    Tag: git winget
TortoiseGit                           TortoiseGit.TortoiseGit            2.12.0.0      Tag: git winget
Sublime Merge (Dev)                   SublimeHQ.SublimeMerge.Dev         2057          Tag: git winget
Sublime Merge                         SublimeHQ.SublimeMerge             2063          Tag: git winget
Git Credential Manager Core           Microsoft.GitCredentialManagerCore 2.0.567.18224 Tag: git winget
glab                                  GLab.GLab                          1.21.1        Tag: git winget
Git LFS                               GitHub.GitLFS                      3.0.2         Tag: git winget
GitHub Desktop                        GitHub.GitHubDesktop               2.9.5         Tag: git winget
Git Extensions                        GitExtensionsTeam.GitExtensions    3.5.4.12724   Tag: git winget
Fork - a fast and friendly git client Fork.Fork                          1.69.4        Tag: git winget
GitKraken                             Axosoft.GitKraken                  8.1.1         Tag: git winget
Sourcetree                            Atlassian.Sourcetree               3.4.7         Tag: git winget
Gource                                acaudwell.Gource                   0.51          Tag: git winget
DevSidecar                            docmirror.dev-sidecar              1.7.1         Tag: git winget
GitFiend                              TobySuggate.GitFiend               0.28.0                 winget
Snagit                                TechSmith.Snagit                   21.4.4.12541           winget
GitHubReleaseNotes                    StefHeyenrath.GitHubReleaseNotes   1.0.7.1                winget
GitBlade                              PirinelLtd.GitBlade                00.00.8.9              winget
GVFS                                  Microsoft.VFSforGit                1.0.21229.1            winget
Microsoft Git                         Microsoft.Git                      2.33.0.0.3             winget
Gitify                                manosim.gitify                     4.3.1                  winget
Logitech Unifying Software            Logitech.UnifyingSoftware          2.50.25                winget
Logitech Solar App                    Logitech.Solar                     1.20.28                winget
Logitech SetPoint                     Logitech.SetPoint                  6.70.55                winget
Logitech Options                      Logitech.Options                   9.40.86                winget
Logitech Gaming Software              Logitech.LGS                       9.02.65                winget
Logitech Camera Settings              Logitech.CameraSettings            2.12.8.0               winget
AnimeBack                             LeGitHubDeTai.AnimeBack            8.0.8                  winget
gitg                                  gnome.gitg                         3.32.1                 winget
Thermal                               gitthermal.thermal                 0.0.4                  winget
eDEX-UI                               GitSquared.edex-ui                 2.2.8                  winget
Gitter                                Gitlab.Gitter                      5.0.1                  winget
GitHub CLI                            GitHub.cli                         2.3.0                  winget
classroom-assistant                   GitHub.ClassroomAssistant          2.0.4                  winget
Atom Beta                             GitHub.Atom.Beta                   1.59.0-beta0           winget
Atom                                  GitHub.Atom                        1.58.0                 winget
GitAhead                              GitAhead.GitAhead                  2.6.3                  winget
MaxTo                                 DigitalCreations.MaxTo             2.2.1                  winget
APDigitalExams                        CollegeBoard.APDigitalExams        0.9.1                  winget
2021 Digital AP Exams                 CollegeBoard.2021DigitalAPExams    0.9.4.5                winget
Fundels                               CartamundiDigital.Fundels          3.0.10                 winget
Logi Tune                             Logitech.LogiTune                  2.206.12.0             winget
Duplicate Cleaner Free                DigitalVolcanoSoftware.DuplicateC… 4.1.2                  winget
Duplicate Cleaner Pro                 DigitalVolcanoSoftware.DuplicateC… 4.1.4                  winget
GitHub Desktop Beta                   GitHub.GitHubDesktop.Beta          2.9.3-beta1            winget
Logitech G HUB                        Logitech.GHUB                      2021.3.9205            winget
GetData Graph Digitizer               getdata.getdata                    2.26                   winget
Renoise                               Renoise.Renoise                    3.3.2         Tag: di… winget
Cacher                                PenguinLabs.Cacher                 2.42.2        Tag: gi… winget
AppInstaller File Builder(Preview)    Microsoft.AppInstallerFileBuilder  1.2020.221.0  Tag: Gi… winget
FL Studio                             ImageLine.FLStudio                 20.8.3.2304   Tag: di… winget
Gridea                                getgridea.gridea                   0.9.2         Tag: gi… winget
MarkRight                             dvcrn.markright                    0.1.11        Tag: gi… winget
Dolt                                  DoltHub.Dolt                       0.34.7        Tag: gi… winget
DAWG                                  dawg.dawg                          0.2.3         Tag: di… winget
REAPER                                Cockos.REAPER                      6.42          Tag: di… winget
CloudShow Launcher                    BinaryFortress.CloudShow           5.7.3.0       Tag: di… winget
Lepton                                hackjutsu.Lepton                   1.10.0        Tag: Gi… winget
Monkey's Audio                        MonkeysAudio.MonkeysAudio          7.20          Tag: di… winget
```

多いわ！ どんだけ引っかかるねん。

まぁ，ともかく， git のバージョンは最新ぽいな。 [Git Extensions](https://gitextensions.github.io/ "Git Extensions | Git Extensions is a standalone UI tool for managing Git repositories") も最新かな。 [PuTTY](https://www.chiark.greenend.org.uk/~sgtatham/putty/ "PuTTY: a free SSH and Telnet client") はどうだろう。

```text
$ winget search putty
名前               ID                          バージョン   一致       ソース
------------------------------------------------------------------------------
Metro PuTTY        9WZDNCRDNDFH                Unknown                 msstore
Putty (Unofficial) 9N8PDN6KS0F8                Unknown                 msstore
PuTTY              PuTTY.PuTTY                 0.76.0.0                winget
MTPuTTY            TTYPlus.MTPutty             1.8          Tag: putty winget
NETworkManager     BornToBeRoot.NETworkManager 2021.11.30.0 Tag: putty winget
PortX              NetSarangComputer.PortX     2.0.3        Tag: putty winget
```

`PuTTY.PuTTY` なら大丈夫かな。なんかノイズが多いなぁ...

おっ， [KeePassXC](https://keepassxc.org/ "KeePassXC Password Manager") は [winget で入れれる](https://keepassxc.org/download/#windows "Download - KeePassXC")ようだ。

```
$ winget install keepassxc
```

ものによって scoop と winget で使い分ければいいかもな。

## ストアアプリ版「秀丸」

いつまでも「メモ帳」なわけにもいかないが Web クライアント越しに VS Code はたぶん辛いので vim あたりを突っ込んでおきたいところだ。

[vim のインストールは Scoop で可能](https://zenn.dev/zetamatta/books/5ac80a9ddb35fef9a146/viewer/c181ec "git , kaoriya-vim のインストール｜scoop / nyagos で始めるコマンドライン生活")。

でも GUI も欲しいよね。
というわけで，かつて20年はお世話になった「秀丸」を入れようと思う。

- [秀まるおのホームページ(サイトー企画)－秀丸エディタ(ストアアプリ版)](http://hide.maruo.co.jp/software/hidemaru_appx.html)

ふむむ。
ストアアプリ版は別ライセンスで Microsoft アカウント単位でお金を払えばいいのか。
Microsoft アカウントに紐付いている決済手段が使えるので楽っちゃあ楽。

というわけで4K円ほど払い込みました。

{{< fig-img src="./hidemaru-ms-store.png" link="./hidemaru-ms-store.png" >}}

[v9.00](https://forest.watch.impress.co.jp/docs/news/1370220.html "「秀丸エディタ」が10年以上ぶりのメジャーバージョンアップ ～v9.00が正式版に - 窓の杜") じゃないのか。
残念（笑）

おそらく「どこでもエディタ」をやりたいのなら，こんな仮想マシンを用意するのではなく [GitHub Codespaces](https://github.com/features/codespaces) や [vscode.dev](https://vscode.dev/) を使えっちう話なんよね。

時代は変わっていく。

## [Tailscale] で NAS につなぐ

さていよいよ外に繋ぐことにする。
といっても外部から [Azure Virtual Desktop] 内のリソースに触る気はないので [Azure Virtual Desktop] 側から [Tailscale] で[自宅 NAS]({{< ref "/remark/2021/10/tailscale-with-synology-nas.md" >}} "Synology NAS に Tailscale を設定する") に繋ぐことを考える。

Windows 版 [Tailscale] は [winget] でインストールできる。

```text
$ winget search tailscale
名前      ID                  バージョン ソース
------------------------------------------------
Tailscale tailscale.tailscale 1.18.0     winget
```

バージョンも問題なさそうだな。

インストールしたら仮想デスクトップから [Tailscale] にログインすれば OK。
[Tailscale] が指定した IP アドレスで NAS に入れることを確認できた。

簡単簡単。
やっぱマシン同士を peer に繋げられるのはいいわ。

というわけで NAS とも繋がったので，更にいろいろ遊べるようになった。

## ブックマーク

- [アプリのインストールや更新をコマンドでサクっと実行「winget」の使い方：Windows 10 The Latest - ＠IT](https://atmarkit.itmedia.co.jp/ait/articles/2106/11/news021.html)
- [wingetを使おう 前編：wingetのインストールと基本操作 | AsTechLog](https://astherier.com/blog/2021/09/winget-usage-1/)
- - [マイクロソフト、WebIDEの「Visual Studio Codespaecs」を「GitHub Codespaces」に統合へ － Publickey](https://www.publickey1.jp/blog/20/webidevisual_studio_codespaecsgithub_codespaces.html)
- [「GitHub Codespaces」が法人向けに提供開始 ～［.］キーでリポジトリを「Visual Studio Code」開くショートカットも - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1344025.html)
- [Github CodeSpace 触ってみた - Qiita](https://qiita.com/Alt225/items/5d904fafc779e6505768)
- [Webブラウザーで「Visual Studio Code」が完全動作 ～「vscode.dev」にアクセスするだけ - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1360147.html)

[Azure Virtual Desktop]: https://docs.microsoft.com/ja-jp/azure/virtual-desktop/ "Azure Virtual Desktop のドキュメント | Microsoft Docs"
[Scoop]: https://scoop.sh/ "Scoop"
[NYAGOS]: https://github.com/nyaosorg/nyagos "nyaosorg/nyagos: NYAGOS - The hybrid Commandline Shell betweeeeeeen UNIX & DOS"
[Chocolatey]: https://chocolatey.org/ "Chocolatey Software | Chocolatey - The package manager for Windows"
[winget]: https://github.com/microsoft/winget-cli "microsoft/winget-cli: Windows Package Manager CLI (aka winget)"
[Tailscale]: https://tailscale.com/ "Tailscale · Best VPN Service for Secure Networks"
