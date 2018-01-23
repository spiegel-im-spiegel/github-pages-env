+++
title = "「CPU に対するサイドチャネル攻撃」に関する覚え書き"
date =  "2018-01-04T22:01:45+09:00"
update = "2018-01-22T19:54:39+09:00"
description = "投機的実行機能やアウトオブオーダー実行機能を持つ CPU に対してサイドチャネル攻撃を行う手法が報告されている。ただし影響度は低い。"
image = "/images/attention/remark.jpg"
tags = [
  "security",
  "vulnerability",
  "device",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

## 脆弱性の内容

投機的実行機能やアウトオブオーダー実行機能を持つ CPU に対してサイドチャネル攻撃を行う手法が報告されている。
攻撃手法としては “{{< pdf-file title="Meltdown" link="https://meltdownattack.com/meltdown.pdf" >}}” と “{{< pdf-file title="Spectre" link="https://spectreattack.com/spectre.pdf" >}}” の2つがあるようだ。

- [Meltdown and Spectre](https://meltdownattack.com/)
- [Project Zero: Reading privileged memory with a side-channel](https://googleprojectzero.blogspot.jp/2018/01/reading-privileged-memory-with-side.html)

想定される影響としては

{{< fig-quote title="CPU に対するサイドチャネル攻撃" link="http://jvn.jp/vu/JVNVU93823979/" >}}
<q>ユーザ権限で実行中のプロセスから、カーネルメモリの情報を取得される可能性があります。<br>
Spectre 攻撃に関しては、細工された Javascript コードによって、Javascript からは取得できないはずの web ブラウザプロセス中のデータを取得可能であると報告されています。</q>
{{< /fig-quote >}}

とのこと。

## 影響度（CVSS）

- [JVNVU#93823979: CPU に対するサイドチャネル攻撃](http://jvn.jp/vu/JVNVU93823979/)

**CVSSv3 基本評価値 4.7 (`CVSS:3.0/AV:L/AC:H/PR:L/UI:N/S:U/C:H/I:N/A:N`)**

|                            基本評価基準 | 評価値        |
| ---------------------------------------:|:------------- |
|                        攻撃元区分（AV） | ローカル（L） |
|                  攻撃条件の複雑さ（AC） | 高（H）       |
|                  必要な特権レベル（PR） | 低（L）       |
|                  ユーザ関与レベル（UI） | 不要（N）     |
|                           スコープ（S） | 変更なし（U） |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）     |
|   業務停止の可能性（可用性への影響, A） | なし（N）     | 

CVSS については[解説ページ]({{< relref "remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

一般的にサイドチャネル攻撃は条件の複雑さから影響度が低く見積もられる。
この脆弱性は2017年春頃からベンダ企業などには知らされていたが，2018年早々にリークされ批判を浴びている模様。

（2018-01-10 追記： 機密性への影響が「低」から「高」へ格上げ）

## 影響を受ける製品

- プロセッサとしては Intel, AMD, ARM の各製品が影響を受けるとされている
    - Intel については1995年以降に製造された Itanium/Atom を除くプロセッサが対象
    - AMD や ARM については影響度を調査中
- Meltdown についてはExploit Code が公開されており Windows および Linux 環境で動作する
    - ただし AMD, ARM については，公開されているコードでは動作しないらしい？
    - Xen PV などによる仮想化マシンは影響を受ける
    - Docker, LXC, OpenVZ などのコンテナ・ソリューションは影響を受ける
- Spectre については各種 Web ブラウザが影響を受ける

## 対策・回避策

根本的な対策はプロセッサを含めた ハードウェア, ファームウェア, OS を一新するしかないが，プロセッサの再設計と製造には数年かかると言われ現実的ではない。
当面の緩和措置として対応するのであれば，今のところソフトウェア側で行うことになる。
ただしソフトウェア側で対策するとなるとパフォーマンスが犠牲にならざるを得ない。

- ファームウェアのアップデートがある場合はこれを適用すること
- Linux では [KPTI / KAISER パッチ](https://lwn.net/Articles/738975/) によって Meltdown の影響を軽減できる
- Windows については2018年1月のアップデートで対応
    - [ADV180002 | Guidance to mitigate speculative execution side-channel vulnerabilities](https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/ADV180002)
    - [Microsoft 製品の脆弱性対策について(2018年01月)：IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/ciadr/vul/20180110-ms.html)
    - [2018年 1月マイクロソフトセキュリティ更新プログラムに関する注意喚起](https://www.jpcert.or.jp/at/2018/at180002.html)
- Apple 製品については2017年12月の更新で Meltdown については対応済み。 Spectre についてはは2018年1月の更新で対応
    - [ARM や Intel の CPU の投機的実行の脆弱性について - Apple サポート](https://support.apple.com/ja-jp/HT208394)
    - [About the security content of iOS 11.2.2 - Apple サポート](https://support.apple.com/ja-jp/HT208401)
    - [About the security content of Safari 11.0.2 - Apple サポート](https://support.apple.com/ja-jp/HT208403)
    - [About the security content of macOS High Sierra 10.13.2 Supplemental Update - Apple サポート](https://support.apple.com/ja-jp/HT208397)
- Google 製品については1月に各種対応版が出る
    - [Product Status Google’s Mitigations Against CPU Speculative Execution Attack Methods - Google Help](https://support.google.com/faqs/answer/7622138)
- Mozilla Firefox 等のブラウザについても Spectre 対応（軽減）版が出ている
    - [Mitigations landing for new class of timing attack | Mozilla Security Blog](https://blog.mozilla.org/security/2018/01/03/mitigations-landing-new-class-timing-attack/)
- サービス・プロバイダでも対応が進んでいる
    - [【重要】MeltdownおよびSpectre（CPUの脆弱性）による弊社サービスへの影響について | さくらインターネット](https://www.sakura.ad.jp/news/sakurainfo/newsentry.php?id=1832)
        - [MeltdownおよびSpectre（CPUの脆弱性）による弊社サービスへの影響について – さくらのサポート情報](https://help.sakura.ad.jp/hc/ja/articles/115000204541?_bdld=2XXAbU.m3Ec1JI)
    - [WEBサーバーの CPU 脆弱性対応メンテナンスのお知らせ - 2018年01月05日 18時57分 / メンテナンス情報 / お知らせ - レンタルサーバーならロリポップ！](https://lolipop.jp/info/mainte/5905/)
- ウイルス対策ソフトで今回の脆弱性を使った malware を検出するのは（できなくはないが）難しい？

## ブックマーク

- [Cyberus Technology Blog - Meltdown](http://blog.cyberus-technology.de/posts/2018-01-03-meltdown.html)
- [Google Online Security Blog: Today's CPU vulnerability: what you need to know](https://security.googleblog.com/2018/01/todays-cpu-vulnerability-what-you-need.html)
- [Kernel-memory-leaking Intel processor design flaw forces Linux, Windows redesign • The Register](https://www.theregister.co.uk/AMP/2018/01/02/intel_cpu_design_flaw/)
- [Meltdown and Spectre Side-Channel Vulnerability Guidance | US-CERT](https://www.us-cert.gov/ncas/alerts/TA18-004A)
- [Spectre and Meltdown Attacks Against Microprocessors - Schneier on Security](https://www.schneier.com/blog/archives/2018/01/spectre_and_mel_1.html) : 必見
- [CPUの脆弱性 MeltdownとSpectreについてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20180104/1515094046) : 日本語の情報ではここが参考になる
- [プロセッサの脆弱性「Meltdown」と「Spectre」についてまとめてみた ｜ Developers.IO](https://dev.classmethod.jp/security/processor-vulnerability-meltdown-spectre/) : 日本語の情報ではここも参考になる
- [Intel、プロセッサ脆弱性はAMDやARMにもあり、対策で協力中と説明 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1801/04/news009.html)
- [マイクロソフト、CPUの脆弱性対策でAzureの計画メンテを前倒し、全リージョンの仮想マシンを今朝から強制再起動。Googleは対策済みと発表 － Publickey](http://www.publickey1.jp/blog/18/cpuazuregoogle.html)
- [「CPUに深刻なバグ」報道にIntel反論――OSアーキテクチャーに内在する欠陥で他社製チップにも同様の影響  |  TechCrunch Japan](http://jp.techcrunch.com/2018/01/04/2018-01-03-intel-calls-reports-of-major-vulnerability-incorrect/)
- [インテル、ARM、AMDなど多数のCPUに脆弱性--各社が対応急ぐ - ZDNet Japan](https://japan.zdnet.com/article/35112721/)
- [“投機的実行”機能を備えるCPUに脆弱性、Windows向けのセキュリティパッチが緊急公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1099690.html)
- [Google、CPU脆弱性“Meltdown”“Spectre”の緩和策を「Google Chrome 64」へ導入 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1099729.html)
- [Mozilla、「Firefox」v57.0.4を公開 ～CPU脆弱性“Meltdown”“Spectre”へ対策 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1099704.html)
- [マイクロソフト、CPUの脆弱性問題で緩和策を緊急公開 - ZDNet Japan](https://japan.zdnet.com/article/35112758/)
- [Intel、プロセッサ脆弱性対策で「来週末までに過去5年に製造したプロセッサの9割に更新実行」 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1801/05/news027.html)
- [Apple、プロセッサ脆弱性「Meltdown」と「Spectre」の対策について説明 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1801/05/news035.html)
- [インテル、CPUに影響する脆弱性「Meltdown」「Spectre」対策でパッチを公開 - ZDNet Japan](https://japan.zdnet.com/article/35112769/)
- [Apple曰く、メルトダウンとスペクター問題は「全MacシステムとiOSデバイス」に影響を与えるが長くは続かない  |  TechCrunch Japan](http://jp.techcrunch.com/2018/01/05/2018-01-04-apple-says-meltdown-and-spectre-flaws-affect-all-mac-systems-and-ios-devices-but-not-for-long/)
- [プロセッサの脆弱性「Spectre」と「Meltdown」について知っておくべきこと - CNET Japan](https://japan.cnet.com/article/35112771/)
- [ニュース - CPU脆弱性問題でAWSとAzureの対応状況が判明：ITpro](http://itpro.nikkeibp.co.jp/atcl/news/17/010402926/)
- [CPUに内在する脆弱性問題「メルトダウン」「スペクター」への各社の対応まとめ - GIGAZINE](https://gigazine.net/news/20180105-meltdown-spectre-security/)
- [CPUの脆弱性、Linux関係者らの見方や対応 - ZDNet Japan](https://japan.zdnet.com/article/35112767/)
- [「Waterfox」v56.0.2が公開、「Firefox」のCPU脆弱性修正を取り込む - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100016.html)
- [NVIDIA、最新版ドライバーでCPU脆弱性“Spectre”対策を実施 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100044.html)
- [VMware製品へのSpectre/Meltdown脆弱性の影響 - Qiita](https://qiita.com/tsukamoto/items/9259050159e9858c81af)
    - [「VMware Workstation」「VMware Fusion」が更新、“Spectre”脆弱性への緩和策が施される - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100770.html)
- [Opera、CPU脆弱性“Meltdown”の影響を緩和する方法を案内 ～対策版は今月末公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100328.html)
- [Windowsの“Meltdown”“Spectre”パッチでトラブル、AMDデバイスへの提供が一時停止 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100231.html)
- [Microsoft、2018年1月のセキュリティ更新プログラムを公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100321.html)

{{< fig-quote title="Microsoft、2018年1月のセキュリティ更新プログラムを公開" link="https://forest.watch.impress.co.jp/docs/news/1100321.html" >}}
<q>なお、“Meltdown”“Spectre”脆弱性のパッチはウイルス対策ソフトとの互換性問題を抱えている。先にウイルス対策ソフトをアップデートして、互換性が取れた旨をレジストリの特定エントリーにマーキングしないと“Windows Update”から更新プログラムを受信できない場合があるので注意したい。</q>
{{< /fig-quote >}}

- [CPUの脆弱性パッチがLinuxマシンの性能に与える影響--レッドハットが検証 - ZDNet Japan](https://japan.zdnet.com/article/35113008/)
- [IBM、脆弱性「Meltdown/Spectre」への対策を「POWER」CPU向けに公開 - ZDNet Japan](https://japan.zdnet.com/article/35112986/)
- [AMD、RyzenとEPYCのSpectre対策を今週中に提供へ　GPUには影響なし - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1801/12/news071.html)
- [AMD、CPU脆弱性に関する新しい声明を発表 ～配信停止中のWindowsパッチは来週再開か - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1100903.html)
- [AWSもSpectreとMeltdownの対策完了を報告。対策後、Amazon EC2で性能の低下は見られないと － Publickey](http://www.publickey1.jp/blog/18/awsspectremeltdownamazon_ec2.html)
- [GoogleはSpectreとMeltdownの対策を昨年6月に開始し12月には完了していた。性能低下の報告はなし － Publickey](http://www.publickey1.jp/blog/18/googlecpuspectremeltdown612.html)
- [LinuxコアメンバーによるMeltdownとSpectre 対応状況の説明 （1/19更新） - Qiita](https://qiita.com/hogemax/items/008f19fd14eebca474d7)
- [「Spectre」のような脆弱性は「おそらくほかにも」--ArmのCEOが警告 - ZDNet Japan](https://japan.zdnet.com/article/35113201/)
- [Windows 10ミニTips(247) CPUの脆弱性「Meltdown」「Spectre」に関する状態を確認する | マイナビニュース](https://news.mynavi.jp/article/win10tips-247/)
- [「Vivaldi」にCPU脆弱性“Spectre”緩和策を施したマイナーアップデートが提供される - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1102534.html)
- [CPU脆弱性への暫定的な緩和策を加えた「Opera」のセキュリティアップデートが公開 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1102503.html)
- [トーバルズ氏、インテルの「Meltdown/Spectre」への対応を批判 - ZDNet Japan](https://japan.zdnet.com/article/35113527/)
