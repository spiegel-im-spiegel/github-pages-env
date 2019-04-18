+++
title = "Ubuntu に Mono を導入する"
date = "2019-04-14T19:36:23+09:00"
description = "今回は Mono および Mono で動くアプリケーションを導入する。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "password", "management", "tools", "git", "git-extensions" ,"dot-net", "mono" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

今回は [Mono] および [Mono] で動くアプリケーションを導入する。

## [Mono] の導入

[Mono] を導入するにはリポジトリの登録から始める必要がある。
`gnupg` と `ca-certificates` は既に入ってるみたいなので公開鍵のインポートとソースリストの追加から。

```text
$ sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys 3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF
$ echo "deb https://download.mono-project.com/repo/ubuntu stable-bionic main" | sudo tee /etc/apt/sources.list.d/mono-official-stable.list
$ sudo apt update
```

続いて `mono-devel` を導入する。

```text
$ sudo apt install mono-devel
```

インストール処理が完了したら念のため動作確認をしておこう。

```text
$ mono -V
Mono JIT compiler version 5.20.1.19 (tarball Thu Apr 11 09:02:17 UTC 2019)
Copyright (C) 2002-2014 Novell, Inc, Xamarin Inc and Contributors. www.mono-project.com
	TLS:           __thread
	SIGSEGV:       altstack
	Notifications: epoll
	Architecture:  amd64
	Disabled:      none
	Misc:          softdebug 
	Interpreter:   yes
	LLVM:          yes(600)
	Suspend:       hybrid
	GC:            sgen (concurrent by default)
```

よーし，うむうむ，よーし。

## [KeePass] の導入

[KeePass] は OSS のパスワード管理ツールで，基本的には Windows 用だが各種プラットフォーム用のバージョンが公開されている。
[Ubuntu] では APT で導入可能である。

```text
$ sudo apt show keepass2
Package: keepass2
Version: 2.39.1+dfsg-1
Priority: optional
Section: universe/utils
Origin: Ubuntu
...
```

ちょっとバージョンが古いがプラットフォームが違うし，まぁいいだろう。
インストールしてしまおう。

```text
$ sudo apt install keepass2
```

インストールが完了するとドックのアプリボタンで表示されるアプリ一覧にアイコンが追加される。
あとは Windows 版と同じように使える。

パスワードを覚えるなんて脳みその無駄遣い。
適切な管理ツールでパスワード管理を行いましょう。

## [Git Extensions] の導入

[Git Extensions] は .NET Framework 上で動く [git] 用の GUI ツールだが， [Mono] 版もあって他プラットフォームでも利用できる。
ただし [Mono] 版はバージョン 2 系までしか対応していないので注意が必要である。

ダウンロードページから [Mono] 版の[最新バージョン](https://github.com/gitextensions/gitextensions/releases/tag/v2.51.05 "Release Version 2.51.05 · gitextensions/gitextensions")をダウンロードする。

適当なディレクトリにダウンロードした zip ファイルを `unzip` コマンドで展開する。
`GitExtensions` ディレクトリが作成されてファイルが展開されている筈である。

展開されたファイルの中に `gitext.sh` があるので，これに実行権限を付与して起動する。
今後のためにパスを通しておいたほうがいいだろう。

最初に [Git Extensions] を起動すると言語の選択ウィンドウが表示される。

{{< fig-img src="./language.png" link="./language.png" width="642" >}}

日本語があってよかったね。
あとは Windows 版と同じ。

けど Windows 版と比べて操作感がイマイチなんだよなぁ。
まぁ [Git Extensions] をメインに作業するわけではないので，これでもいいか。

## .NET Framework と [Mono] と .NET Core

.NET Framework と [Mono] と .NET Core の関係がいまひとつ分からなかったのでちょろんと調べてみた。
かなり大雑把な説明なのはご容赦。

### .NET Framework

.NET Framework が最初に公開されたのは2001年。
基盤となる CLI (Common Language Infrastructure) や CLR (Common Language Runtime) は標準化され Windows 以外のプラットフォームでも実装可能となるよう設計されたが，実際にはリファレンス実装である .NET Framework は Windows 以外には対応しなかった。
まぁ，当時の Microsoft は FOSS に敵対的だったし，それもやむなしというところだろうか。

### [Mono] と Xamarin

当然ながら，.NET Framework に対抗する OSS 製品を作ろうという動きがあり，そのうちのひとつが [Mono] Project であった。
[Mono] はマルチプラットフォームで動作するアプリケーション基盤として実装されているのが特徴である。

[Mono] を巡っては紆余曲折があったが（生臭い話は省略），最終的に Xamarin が開発母体となった。

Xamarin は  [Mono] Project のオーナー企業であると同時に [Mono] を基盤とした製品群を指すようだ。
なので今後は [Mono]/Xamarin と一括りにしてしまおう。

企業としての Xamarin は後に Microsoft に買収され [Mono]/Xamarin は Microsoft 傘下である [.NET Foundation] において今も開発が続けられている。

### .NET Core

一方，実質 Windows でしか動作しない .NET Framework の派生として .NET Core が Microsoft からリリースされた。
.NET Core はマルチプラットフォームで動作する OSS 製品として，こちらも [.NET Foundation] で開発が行われている。

ここに於いて Microsoft は経営方針の大転換を果たしたわけだ。

### .NET Core と [Mono]/Xamarin

現在 Microsoft は OSS の .NET シリーズとして .NET Core と [Mono]/Xamarin の2系統の製品を持っている。
両者の棲み分けは以下のような感じらしい。

{{< fig-quote title="ASCII.jp：.NET Core / .NET Framework / Xamarin / Monoの関係を整理する" link="https://ascii.jp/elem/000/001/156/1156721/" >}}
{{% fig-img src="https://ascii.jp/elem/000/001/156/1156949/ph07_2000x1125.jpg" link="https://ascii.jp/elem/000/001/156/1156949/img.html" width="2000" %}}
{{< /fig-quote >}}

つまり Linux や macOS，あるいは iOS や Android の GUI アプリケーションとしては [Mono]/Xamarin，サーバ・サイドあるいはクラウド向けには .NET Core ということのようだ。
さらに将来的には Windows アプリケーション向けの .NET Framework，マルチプラットフォーム向けの .NET Core と [Mono]/Xamarin の基盤を共通化すると表明されている。

{{< fig-quote title="ASCII.jp：.NET Core / .NET Framework / Xamarin / Monoの関係を整理する" link="https://ascii.jp/elem/000/001/156/1156721/" >}}
{{% fig-img src="https://ascii.jp/elem/000/001/156/1156951/ph09_2000x1125.jpg" link="https://ascii.jp/elem/000/001/156/1156951/img.html" width="2000" %}}
{{< /fig-quote >}}

そうなったら [Mono] も役目を終える感じになるのかねぇ。

## ブックマーク

- [C#7に完全対応した「Mono 5.0」が公開 | OSDN Magazine](https://mag.osdn.jp/17/05/22/150000)
- [ASCII.jp：.NET Core / .NET Framework / Xamarin / Monoの関係を整理する (1/3)](https://ascii.jp/elem/000/001/156/1156721/)
- [.NET Core について | Microsoft Docs](https://docs.microsoft.com/ja-jp/dotnet/core/about)
- [.NETの派生を理解する](https://www.infoq.com/jp/articles/varieties-dotnet)
- [オープンソースのMonoと.NET Coreを比較](https://www.ossnews.jp/compare/Mono/dotNET_Core)
- [Xamarin 最近どうよ？ - Qiita](https://qiita.com/amay077/items/399002a02c1abf9d620b)

- [How To: run Git Extensions on Linux · gitextensions/gitextensions Wiki](https://github.com/gitextensions/gitextensions/wiki/How-To%3A-run-Git-Extensions-on-Linux)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Mono]: https://www.mono-project.com/
[KeePass]: https://keepass.info/ "KeePass Password Safe"
[Git Extensions]: https://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"
[git]: https://git-scm.com/
[.NET Foundation]: https://www.dotnetfoundation.org/
