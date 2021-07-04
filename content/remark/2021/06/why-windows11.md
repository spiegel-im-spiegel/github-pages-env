+++
title = "Windows 10 が最後って言ったぢゃん！"
date =  "2021-06-26T13:42:39+09:00"
description = "Windows 10 を最後にすると約束したな。あれは嘘だ！"
image = "/images/attention/kitten.jpg"
tags = [ "windows", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

{{< div-gen class="center" type="markdown" >}}
**Windows 10 を最後にすると約束したな<br>あれは嘘だ！**
{{< /div-gen >}}

ちっくしょー！

...気を取り直して。

ビル・ゲイツ氏が引退してからかなりまともになった Microsoft だが，どうにも Windows に関しては「やらかす」度合いが酷い気がする。
というわけで，2021年内に Windows 11 を出すとか何とか。
**最後の** Windows だったはずの 10 は2025年でサポートが切れるとか何とか。

自宅マシンはとっくに [Ubuntu に換装済み]({{< ref "/remark/2019/04/hello-ubuntu.md" >}} "Windows とともに平成は去り Ubuntu とともに令和を迎える")なので Windows がどうなろうと知ったことではないのだが，職場のマシンや関わってるプロジェクトのシステム要件を見直す必要があるのは面倒な話である。
そこで今後のために Windows 11 の要件をまとめておこう。

## ハードウェア要件

一番重要な点は 32bit アーキテクチャが対象外となることだろう。
まぁ，この辺は時代の流れというやつで仕方がない。
「Ubuntu にすればいいじゃない」という声も聞こえるが， 32bit アーキテクチャをサポートする LTS の最終版である 18.04 (Bionic Beaver) のサポート期限が2023年4月なので[^ubunt1]，乗り換えるメリットはないだろう。

[^ubunt1]: Ubuntu は LTS 版について[10年の延長セキュリティ・メンテナンスを発表](https://ubuntu.com/blog/ubuntu-16-04-lts-transitions-to-extended-security-maintenance-esm "Ubuntu 16.04 LTS transitions to Extended Security Maintenance | Ubuntu")しているが， 32bit アーキテクチャは対象に含まれていないようだ。

今ある 32bit 機は4年以内に引退させ，物理的にオフラインにしてソリティア専用機として余生を送らせるのがいいだろう。

その他の要件は以下の通り[^soc1]。

{{< fig-quote class="nobox" type="markdown" title="「Windows 11」は32bit CPUをサポートせず ～セキュアブート、TPM 2.0も必須に - 窓の杜" link="https://forest.watch.impress.co.jp/docs/news/1333957.html" >}}
|                      | Windows 11                            | Windows 10                   |
| -------------------- | ------------------------------------- | ---------------------------- |
| プロセッサー         | 1GHz以上、2コア以上、<br>64bit互換CPU/SoC | 1GHz以上のCPU/SoC            |
| メインメモリ         | 4GB                                   | 2GB（64bit）、<br>1GB（32bit）   |
| ストレージ           | 64GB以上                              | 32GB（64bit）、<br>16GB（32bit） |
| グラフィックスカード | DirectX 12以降（WDDM 2.0）            | DirectX 9以降（WDDM 1.0）    |
| ディスプレイ         | 対角9インチ以上、<br>8bitカラー、720p     | 800x600                      |
{{< /fig-quote >}}

[^soc1]: SoC は “System on a Chip” の略。 CPU や GPU あるいはメモリやモデムといった機能をワンチップに組み込んだ構成を指すらしい。

64bit アーキテクチャでシングルコアってことはないだろうから，プロセッサ要件はいいか。
古いミニノートだとディスプレイ要件を満たさないのか？

Windows 10 のメモリ要件が 2GB 以上というのは知らなかった。
でも現時点でも起動直後に 4GB のメモリを専有してるんだけど，これで 11 になったら 8GB でも足らなくなるんじゃないの？

## ファームウェア要件

ファームウェアについては UEFI (Unified Extensible Firmware Interface)，セキュアブート，および TPM (Trusted Platform Module) 2.0 に対応していることが要件となる。
これ 64bit 機でも古いマシンだと要件を満たさないものが結構あるんじゃないだろうか。

この前買った自宅マシンは TPM 2.0 に対応してるのかと思ったが，大丈夫ぽい。

{{< fig-img src="./firmware.png" link="./firmware.png" width="1920" >}}

AMD Ryzen シリーズは “AMD Memory Guard” としてチップレベルで対応してるんだってさ。

{{< fig-quote class="nobox" type="markdown" title="AMD MEMORY GUARD" link="https://www.amd.com/system/files/documents/amd-memory-guard-white-paper.pdf" lang="en" >}}
{{< fig-img src="./amd-memory-guard.png" link="https://www.amd.com/system/files/documents/amd-memory-guard-white-paper.pdf" width="848" >}}
{{< /fig-quote >}}

うんうん。

### TPM の確認

私の周辺の反応を見ると TPM 2.0 未対応で非互換と判定されることが多いらしい。

Windows 環境で TPM 対応の有無をチェックするには [`tpm.msc` を起動](https://www.nextofwindows.com/do-i-have-tpm-2-0-on-my-computer-to-run-windows-11 "Do I Have TPM 2.0 on My Computer to Run Windows 11? - NEXTOFWINDOWS.COM")するといいようだ。
TPM に対応していないマシンでは「互換性のある TPM が見つかりません」とか表示される。

Linux 環境で確認する場合は，カーネル等の構成によると思うが，たとえば

```text
$ sudo dmesg | grep -i tpm
[    0.004558] ACPI: Reserving TPM2 table memory at [mem 0xba8c4000-0xba8c404b]
```

てな感じにカーネル起動時のメッセージログから調べることもできる。

SoC に組み込まれている fTPM の場合は，取り扱いに若干の注意が必要らしい。

{{< fig-quote type="markdown" title="【特集】Windows 11で必須になった「TPM 2.0」って何？TPMの役割や確認方法を紹介 - PC Watch" link="https://pc.watch.impress.co.jp/docs/topic/feature/1334277.html" >}}
{{% quote %}}Ryzenの場合、CPUを交換するとTPMのデータも置き換わる(CPUのSoCにTPMが搭載されているため)。Windows 11導入の際に、CPU交換なども考えている場合は、鍵の扱いに注意が必要だ。BitLockerなどはあらかじめ解除してからCPUを交換しないと起動できなくなる可能性がある{{% /quote %}}
{{< /fig-quote >}}

ふむふむ。

### ファームウェアに関するブックマーク

- [TPM 2.0 Library | Trusted Computing Group](https://trustedcomputinggroup.org/resource/tpm-library-specification/)
- [Windows での TPM の使われ方 - Microsoft 365 Security | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows/security/information-protection/tpm/how-windows-uses-the-tpm)
- [セキュリティ プロセッサ (TPM) ファームウェアを更新する](https://support.microsoft.com/ja-jp/windows/-94205cbc-a492-8d79-cc55-1ecd6b0a8022)
- [【特集】Windows 11で必須になった「TPM 2.0」って何？TPMの役割や確認方法を紹介  - PC Watch](https://pc.watch.impress.co.jp/docs/topic/feature/1334277.html)
- [UEFI/GPT ベースのハード ドライブ パーティション | Microsoft Docs](https://docs.microsoft.com/ja-jp/windows-hardware/manufacture/desktop/configure-uefigpt-based-hard-drive-partitions) : 実は GPT (GUID Partition Table) によるストレージのパーティション管理が隠し要件になっているらしい。旧来の MBR で構成している Windows 環境に対してアップグレードしようとすると NG になるかもしれない

## アプリケーション

標準で入るアプリケーションはかなり淘汰されるようだ。

{{< fig-quote type="markdown" title="「Windows 11」で廃止されるアプリや機能--「Cortana」も姿を消すことに - ZDNet Japan" link="https://japan.zdnet.com/article/35172927/" >}}
- 管理者と一般ユーザーの両方が、Windows Store以外からアプリケーションをインストールできないようにしたり、多くのWindows管理ツールをブロックしたりする「Sモード」機能は、Windows 11のHomeエディションでしか利用できなくなる。
 - 「Internet Explorer」（IE）はWindows 11に搭載されない。一部の古いウェブサイトや業務アプリケーションにアクセスする必要がある企業顧客は、新しい「Microsoft Edge」に組み込まれた「IEモード」を使う必要がある。
 - 「Cortana」にも別れを告げることになる。かつては「Siri」と「Alexa」をライバル視していたCortanaだが、もはや初回起動時に音声案内が流れることもなければ、タスクバーにアイコンがピン留めされることもない。
{{< /fig-quote >}}

[Cortana は要らん]({{< ref "/remark/2020/10/stop-cortana-kun.md" >}} "ストップ!! Cortana 君！")よねぇ。
職場で Windows 10 機を支給されたときに真っ先に潰したし（笑）

他に削除・制限される機能は以下の通り（「[次期OS「Windows 11」は「Internet Explorer」無効 ～廃止・削除される機能が案内 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1333959.html)」より抜粋）。

- 壁紙の同期 → 同期できなくなる
- 数式入力パネル → 削除
- ニュースと関心事項 → ウィジェットに置き換え
- ロック画面の簡易ステータス → 廃止
- Skype の「今すぐ会議」 → Microsoft Teams のチャット機能に置き換え[^mt1]
- Snipping Tool → 「切り取り & スケッチ」機能に置き換え？
- ［スタート］画面
  - タイルをグループ化する機能はなくなる。レイアウトは今のところサイズ変更できない
  - Windows 10からの移行時にピン留めアプリ・サイトは移行されない
  - ライブタイルは使用できなくなる。後継機能は新しいウィジェットとなる
  - タブレットモードは廃止される。代わりの機能が追加されるようだ
- タスクバー
  - 「People」がタスクバーからなくなる
  - アップグレード前にカスタマイズしたものを含む一部のアイコンは、アップグレードしたデバイスのタスクトレイに表示されなくなる
  - タスクバーの位置は画面下部にのみ。上や横には表示できない
  - アプリはタスクバー エリアをカスタマイズできなくなる
  - タイムラインはなくなる。「Microsoft Edge」にそれを補完する機能が導入される
  - 18インチ以上のモニターでタッチキーボードはキーボードレイアウトのドック・アンドックが不能となる
  - 「Wallet」は削除される
- 以下のプリインストールアプリは廃止（Microsoft Store からインストールは可能）
  - 「3D ビューアー」
  - Windows 10 向けの「OneNote」
  - 「ペイント 3D」
  - 「Skype」

プリインストールアプリは本当に最小限にしていただきたいものである。

[^mt1]: Microsoft Teams は[個人で無料アカウントを取れる](https://www.publickey1.jp/blog/21/microsoft_teamstodo30024.html "Microsoft Teamsの個人向け無償提供が正式にスタート。友人や家族とのチャット、ToDoリストの共有、当面は300人24時間まで無料のビデオ会議など提供 － Publickey")ようになった。

## 【2021-07-04 追記】 「第8世代プロセッサ」が前提？

2021-06-28 の Microsoft のブログでシステム要件の前提として以下の3つの原則を挙げている。

{{< fig-quote type="markdown" title="Update on Windows 11 minimum system requirements | Windows Insider Blog" link="https://blogs.windows.com/windows-insider/2021/06/28/update-on-windows-11-minimum-system-requirements/" lang="en" >}}
1. Security. Windows 11 raises the bar for security by requiring hardware that can enable protections like Windows Hello, Device Encryption, virtualization-based security (VBS), hypervisor-protected code integrity (HVCI) and Secure Boot. The combination of these features has been shown to [reduce malware by 60%](https://www.microsoft.com/security/blog/2021/01/11/new-surface-pcs-enable-virtualization-based-security-vbs-by-default-to-empower-customers-to-do-more-securely/) on tested devices. To meet the principle, all Windows 11 supported CPUs have an embedded TPM, support secure boot, and support VBS and specific [VBS capabilities](https://techcommunity.microsoft.com/t5/virtualization/virtualization-based-security-enabled-by-default/ba-p/890167).
2. Reliability. Devices upgraded to Windows 11 will be in a supported and reliable state. By choosing CPUs that have adopted the new [Windows Driver model](https://docs.microsoft.com/en-us/windows-hardware/drivers/develop/getting-started-with-windows-drivers) and are supported by our OEM and silicon partners who are achieving a 99.8% crash free experience.
3. Compatibility. Windows 11 is designed to be compatible with the apps you use. It has the fundamentals of >1GHz, 2-core processors, 4GB memory, and 64GB of storage, aligning with our minimum system requirements for Office and Microsoft Teams.
{{< /fig-quote  >}}

更に

{{< fig-quote type="markdown" title="Update on Windows 11 minimum system requirements | Windows Insider Blog" link="https://blogs.windows.com/windows-insider/2021/06/28/update-on-windows-11-minimum-system-requirements/" lang="en" >}}
{{% quote %}}Using the principles above, we are confident that devices running on Intel 8th generation processors and AMD Zen 2 as well as Qualcomm 7 and 8 Series will meet our principles around security and reliability and minimum system requirements for Windows 11. As we release to Windows Insiders and partner with our OEMs, we will test to identify devices running on Intel 7th generation and AMD Zen 1 that may meet our principles{{% /quote %}}.
{{< /fig-quote  >}}

と謳っていて，いわゆる「第8世代プロセッサ」を前提としているようだ（第7世代はこれから検証）。
というわけで，[2018年に公表された CPU 脆弱性]({{< ref "/remark/2018/01/reading-privileged-memory-with-a-side-channel.md" >}} "「CPU に対するサイドチャネル攻撃」に関する覚え書き")への連想も成り立つわけだ。

- [ASCII.jp：Windows 11にアップグレード可能なCPUは基本はやっぱり第8世代／Zen+以降になりそう？ (1/2)](https://ascii.jp/elem/000/004/061/4061479/)

今後パソコンを購入するのであれば（特に中古パソコンを買うのなら）この辺が目安になってくるだろう。

## ブックマーク

- [新しい Windows 11 OS へのアップグレード | Microsoft](https://www.microsoft.com/ja-jp/windows/windows-11)
- [Windows 11 の仕様 - Microsoft](https://www.microsoft.com/ja-jp/windows/windows-11-specifications)
- [Windows 11発表。年内提供予定でWindows 10からは無償アップグレード  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1333729.html)
- [「Windows 11」は32bit CPUをサポートせず ～セキュアブート、TPM 2.0も必須に - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1333957.html)
- [マイクロソフトがWindows 11を正式発表、アマゾンと驚きの提携でAndroidアプリも利用可能に、年末商戦までに一般発売  |  TechCrunch Japan](https://jp.techcrunch.com/2021/06/25/2021-06-24-microsoft-announces-windows-11-generally-available-by-the-holidays/)
- [「Windows 11」、機能アップデートは年1回に - ZDNet Japan](https://japan.zdnet.com/article/35172929/)
- [「Windows 11」で廃止されるアプリや機能--「Cortana」も姿を消すことに - ZDNet Japan](https://japan.zdnet.com/article/35172927/)
- [「Windows 11」はAndroidアプリにも対応 ～ストアはAmazonのものを利用 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1333960.html)
- [次期OS「Windows 11」は「Internet Explorer」無効 ～廃止・削除される機能が案内 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1333959.html)
- [Windows 10 ユーザーよ、2025年10月のサポート終了に備える時が来ました | ギズモード・ジャパン](https://www.gizmodo.jp/2021/06/end-of-windows-10-support.html)
- [［速報］マイクロソフト、Windows 11を発表。UIを洗練、Windows Updateは40％小さく、マルチモニタ環境が便利に － Publickey](https://www.publickey1.jp/blog/21/windows_11uiwindows_update40.html)
- [［速報］Windows 11でAndroidアプリが実行可能に、マイクロソフトが発表 － Publickey](https://www.publickey1.jp/blog/21/windows_11android.html)
- [自作AMD RyzenマシンでTPMは利用できるのか  |  ちりつもノート](https://chiritsumon.net/contents/archives/22)
- [「Windows Subsystem for Android」に言及--「Windows 11」の「Android」アプリ対応 - ZDNet Japan](https://japan.zdnet.com/article/35172926/)
- [Windows 11のアプリ配信の仕組みから、マイクロソフトの「大きな変化」が見えてくる | WIRED.jp](https://wired.jp/2021/06/26/microsoft-windows-11-app-store-android/)
- [GitHub - rcmaehl/WhyNotWin11: Detection Script to help identify why your PC isn't Windows 11 ready](https://github.com/rcmaehl/WhyNotWin11) : 対応度合いを詳しく表示してくれるのでベンチだが，実行にはデバイスへのアクセス権限が必要。ご利用は自己責任で
  - [「Windows 11」へアップグレードできない理由を親切に教えてくれる「WhyNotWin11」【レビュー】 - 窓の杜](https://forest.watch.impress.co.jp/docs/review/1334420.html)
- [GitHub - builtbybel/ReadySunValley: Ready for Windows11/Sun Valley update?](https://github.com/builtbybel/ReadySunValley) : コマンドライン版。実行にはデバイスへのアクセス権限が必要。ご利用は自己責任で
  - [Check with the ReadySunValley app if your device works with Windows11/Sun Valley update | built by Bel ❤︎](https://www.builtbybel.com/blog/19-apps/41-check-with-the-readysunvalley-app-if-your-device-works-with-windows11-sun-valley-update)
  - [「Windows 11」へ更新できない原因の調査ツール「ReadySunValley」 ～親切なUIが魅力【レビュー】 - 窓の杜](https://forest.watch.impress.co.jp/docs/review/1335219.html)
- [Microsoftが「Windows 11」のシステム要件を見直しへ ～「PC 正常性チェック」ツールはいったん撤回 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1334623.html)
- [ARMネイティブの速度とx64相互運用を両立 ～Microsoft、「ARM64EC」を発表 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1334821.html)

## 参考図書

{{% review-paapi "B00WAMAKZQ" %}} <!-- コマンドー -->
