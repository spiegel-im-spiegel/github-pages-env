+++
title = "「投機的実行機能を持つ CPU に対するサイドチャネル攻撃」に関する覚え書き"
date =  "2018-01-04T22:01:45+09:00"
description = "投機的実行機能を持つ CPU に対してサイドチャネル攻撃を行う手法が報告されている。ただし影響度は低い。"
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

このニュースは帰省中に見てちょっとビックリしたのだが，よく見たら昨年11月に見たこのニュースのことらしい。
脅かすな，マスメディアめ！

- [2015年以降のインテルCPUに遠隔攻撃許す深刻な脆弱性。サーバーからIoTまで、早急なファームウェアの更新を呼びかけ - Engadget 日本版](http://japanese.engadget.com/2017/11/23/2015-cpu-iot/)

## 脆弱性の内容

投機的実行機能を持つ CPU に対してサイドチャネル攻撃を行う手法が報告されている。
攻撃手法としては “{{< pdf-file title="Meltdown" link="https://meltdownattack.com/meltdown.pdf" >}}” と “{{< pdf-file title="Spectre" link="https://spectreattack.com/spectre.pdf" >}}” の2つがあるようだ。

- [Meltdown and Spectre](https://meltdownattack.com/)
- [Project Zero: Reading privileged memory with a side-channel](https://googleprojectzero.blogspot.jp/2018/01/reading-privileged-memory-with-side.html)

## 影響度（CVSS）

- [JVNVU#93823979: 投機的実行機能を持つ CPU に対するサイドチャネル攻撃](http://jvn.jp/vu/JVNVU93823979/)

**CVSSv3 基本評価値 2.5 (`CVSS:3.0/AV:L/AC:H/PR:L/UI:N/S:U/C:L/I:N/A:N`)**

|                            基本評価基準 | 評価値        |
| ---------------------------------------:|:------------- |
|                        攻撃元区分（AV） | ローカル（L） |
|                  攻撃条件の複雑さ（AC） | 高（H）       |
|                  必要な特権レベル（PR） | 低（L）       |
|                  ユーザ関与レベル（UI） | 不要（N）     |
|                           スコープ（S） | 変更なし（U） |
| 情報漏えいの可能性（機密性への影響, C） | 低（L）       |
| 情報改ざんの可能性（完全性への影響, I） | なし（N）     |
|   業務停止の可能性（可用性への影響, A） | なし（N）     | 

CVSS については[解説ページ]({{< relref "remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

一般的にサイドチャネル攻撃は条件の複雑さから影響度が低く見積もられる。
今回はマイクロプロセッサの脆弱性ということで，クラウドサービスを含め影響範囲が広いため話題になっていると思われる。

## 影響を受ける製品

- プロセッサとしては Intel, AMD, ARM の各製品が影響を受けるとされている。
- Meltdown についてはExploit Code が公開されており Windows および Linux 環境で動作する
    - ただし AMD, ARM については，公開されているコードでは動作しないらしい？
    - Xen PV などによる仮想化マシンは影響を受ける
    - Docker, LXC, OpenVZ などのコンテナ・ソリューションは影響を受ける

## 対策・回避策

対策するのであれば，今のところソフトウェア側で行うことになる。
ただしソフトウェア側で対策するとなるとパフォーマンスが犠牲にならざるを得ないだろう。

- ファームウェアのアップデートがある場合はこれを適用すること
- Linux では [KPTI / KAISER パッチ](https://lwn.net/Articles/738975/) によって Meltdown の影響を軽減できる
- Windows については2018年1月のアップデートで対応予定？
    - [ADV180002 | Guidance to mitigate speculative execution side-channel vulnerabilities](https://portal.msrc.microsoft.com/en-US/security-guidance/advisory/ADV180002)

## ブックマーク

- [Cyberus Technology Blog - Meltdown](http://blog.cyberus-technology.de/posts/2018-01-03-meltdown.html)
- [Google Online Security Blog: Today's CPU vulnerability: what you need to know](https://security.googleblog.com/2018/01/todays-cpu-vulnerability-what-you-need.html)
- [Kernel-memory-leaking Intel processor design flaw forces Linux, Windows redesign • The Register](https://www.theregister.co.uk/AMP/2018/01/02/intel_cpu_design_flaw/)
- [Spectre and Meltdown Attacks - Schneier on Security](https://www.schneier.com/blog/archives/2018/01/spectre_and_mel.html) : “I'll be writing something for publication over the next few days.” だそうです
- [Intel、プロセッサ脆弱性はAMDやARMにもあり、対策で協力中と説明 - ITmedia NEWS](http://www.itmedia.co.jp/news/articles/1801/04/news009.html)
- [マイクロソフト、CPUの脆弱性対策でAzureの計画メンテを前倒し、全リージョンの仮想マシンを今朝から強制再起動。Googleは対策済みと発表 － Publickey](http://www.publickey1.jp/blog/18/cpuazuregoogle.html)
- [「CPUに深刻なバグ」報道にIntel反論――OSアーキテクチャーに内在する欠陥で他社製チップにも同様の影響  |  TechCrunch Japan](http://jp.techcrunch.com/2018/01/04/2018-01-03-intel-calls-reports-of-major-vulnerability-incorrect/)
- [インテル、ARM、AMDなど多数のCPUに脆弱性--各社が対応急ぐ - ZDNet Japan](https://japan.zdnet.com/article/35112721/)
