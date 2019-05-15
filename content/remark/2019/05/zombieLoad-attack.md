+++
title = "Intel チップの Microarchitectural Data Sampling 脆弱性について"
date =  "2019-05-15T23:07:45+09:00"
description = "一般的にサイドチャネル攻撃の深刻度はそれほど高くならない。アップデートは計画的に。"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "vulnerability",
  "device",
]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨日あたりから [Ubuntu security notices](https://usn.ubuntu.com/) による Linux kernel 周りの脆弱性報告で騒がしいと思ったら Intel チップをターゲットとした新たなサイドチャネル攻撃が公開されたようだ。

- [Intel Side Channel Vulnerability MDS](https://www.intel.com/content/www/us/en/architecture-and-technology/mds.html)

割り振られている CVE 番号は以下の通り。

- [CVE-2018-12126](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-12126)：Microarchitectural Store Buffer Data Sampling (MSBDS)
- [CVE-2018-12127](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-12127)：Microarchitectural Fill Buffer Data Sampling (MFBDS)
- [CVE-2018-12130](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-12130)：Microarchitectural Load Port Data Sampling (MLPDS)
- [CVE-2019-11091](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2019-11091)：Microarchitectural Data Sampling Uncacheable Memory (MDSUM)

特に [CVE-2018-12130](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-12130) の方は “[ZombieLoad Attack](https://zombieloadattack.com/)” として PoC (Proof of Concept) コードも公開されているようだ。

- [Cyberus Technology - ZombieLoad: Cross Privilege-Boundary Data Leakage](https://www.cyberus-technology.de/posts/2019-05-14-zombieload.html)

[Ubuntu] 19.04 の kernel (Linux kernel v5.0 系) も対象になっている。

- [USN-3979-1: Linux kernel vulnerabilities | Ubuntu security notices](https://usn.ubuntu.com/3979-1/)

もちろん Linux/[Ubuntu] だけでなく Windows や macOS などの主要 OS も影響を受ける。

- [ADV190013 | Microsoft Guidance to mitigate Microarchitectural Data Sampling vulnerabilities](https://portal.msrc.microsoft.com/en-us/security-guidance/advisory/adv190013)
- [Additional mitigations for speculative execution vulnerabilities in Intel CPUs - Apple Support](https://support.apple.com/en-us/HT210107)
- [VMSA-2019-0008 (VMware Security Advisories)](https://www.vmware.com/security/advisories/VMSA-2019-0008.html)

一般的にサイドチャネル攻撃は前提条件が複雑で深刻度もそれほど高くならない。
今回も CVSSv3 で 3.8 から 6.5 程度のようだ（7 以上からが慌てるレベル）。
とはいえ PoC コードも公開されているし放置するわけにもいかないので，計画的にアップデートを行うことをお勧めする。

なお，ソフトウェア側の対応はあくまでも「緩和（mitigation）」なので，本格的な対応はハードウェア側で行う必要がある。
まぁ，可能ならパソコンは1,2年ほど買い控えたほうがいいかもねぇ。

そうそう。
Windows は MDS よりヤバい脆弱性があるみたいで

- [Windows リモートデスクトップサービスの脆弱性 CVE-2019-0708についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2019/05/15/120313)
- [Microsoft 「Windows Server 2008」および「Windows Server 2003」のリモートデスクトップサービスにおける脆弱性について | さくらインターネット](https://www.sakura.ad.jp/information/announcements/2019/05/15/1968200218/)

既にサポートが終了している WinXP まで適用されるとか，本当にご愁傷さまです（他人事）。

## ブックマーク

- [Intel 製品の複数の脆弱性 (INTEL-SA-00213) に関する注意喚起](https://www.jpcert.or.jp/at/2019/at190024.html)
- [インテルCPUに重大バグZombieLoad発見、各社がパッチを相次いでリリース  |  TechCrunch Japan](https://jp.techcrunch.com/2019/05/15/2019-05-14-intel-chip-flaws-patches-released/)
- [Intel製CPUに新たな脆弱性 ～“Microarchitectural Data Sampling（MDS）”が公表される - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1184519.html)
- [インテルのチップに新たな脆弱性--「Microarchitectural Data Sampling（MDS）」 - CNET Japan](https://japan.cnet.com/article/35136950/)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
