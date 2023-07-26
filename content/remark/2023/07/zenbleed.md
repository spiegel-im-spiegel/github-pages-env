+++
title = "AMD Zen 2 CPU の脆弱性について"
date =  "2023-07-26T20:50:32+09:00"
description = "とりあえずの回避策として脆弱性に対応した microcode がリリースされている。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

覚え書きとして記しておく。

## [CVE-2023-20593] Cross-Process Information Leak

発見者は "Zenbleed” と名付けているようだ。

- [Zenbleed](https://lock.cmpxchg8b.com/zenbleed.html)

これによると対象となるプロセッサは以下の通り。


- AMD Ryzen 3000 Series Processors
- AMD Ryzen PRO 3000 Series Processors
- AMD Ryzen Threadripper 3000 Series Processors
- AMD Ryzen 4000 Series Processors with Radeon Graphics
- AMD Ryzen PRO 4000 Series Processors
- AMD Ryzen 5000 Series Processors with Radeon Graphics
- AMD Ryzen 7020 Series Processors with Radeon Graphics
- AMD EPYC “Rome” Processors

この脆弱性は2023年5月には AMD に報告されている。

- [Cross-Process Information Leak](https://www.amd.com/en/resources/product-security/bulletin/amd-sb-7008.html)

とりあえずの回避策として脆弱性に対応した microcode がリリースされている。
Ubuntu でも APT で提供済みで，手元のマシンではすでに対応した（`amd64-microcode` ってパッケージがそれ）。

- [kernel/git/firmware/linux-firmware.git - Repository of firmware blobs for use with the Linux kernel](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/commit/?id=0bc3126c9cfa0b8c761483215c25382f831a7c6f)
- [USN-6244-1: AMD Microcode vulnerability | Ubuntu security notices | Ubuntu](https://ubuntu.com/security/notices/USN-6244-1)

ファームウェアの改修は2023年末まで待つ必要があるらしい。

{{< fig-quote class="nobox" type="markdown" title="Zen 2コアに脆弱性。修正は2023年10月以降 - PC Watch" link="https://pc.watch.impress.co.jp/docs/news/1518766.html" >}}
| プロセッサ                                        | ファームウェア                                            | 時期           |
| ------------------------------------------------- | --------------------------------------------------------- | -------------- |
| Ryzen 3000(Matisse)                               | ComboAM4v2PI_1.2.0.C<br>ComboAM4PI_1.0.0.C                | 2023年12月目標 |
| Ryzen 4000(Renoir)                                | ComboAM4v2PI_1.2.0.C                                      | 2023年12月目標 |
| Ryzen Threadripper 3000(Castle Peak)              | CastlePeakPI-SP3r3 1.0.0.A                                | 2023年10月目標 |
| Ryzen Threadripper PRO 3000WX(Castle Peak WS SP3) | CastlePeakWSPI-sWRX8<br>ChagallWSPI-sWRX8 1.0.0.7 1.0.0.C | 2023年11月目標 |
| Ryzen 5000(Lucienne、モバイル)                    | CezannePI-FP6_1.0.1.0                                     | 2023年12月目標 |
| Ryzen 4000(Renoir、モバイル)                      | RenoirPI-FP6_1.0.0.D                                      | 2023年11月目標 |
| Ryzen 7020(Mendocino、モバイル)                   | MendocinoPI-FT6_1.0.0.6                                   | 2023年12月目標 |
{{< /fig-quote >}}

忘れそうだな（笑）

## ブックマーク

- [security-research/pocs/cpus/zenbleed at master · google/security-research · GitHub](https://github.com/google/security-research/tree/master/pocs/cpus/zenbleed)
- [RyzenシリーズなどAMD CPUの一部にプロセス間情報漏洩の脆弱性 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1518841.html)
- [Zen 2コアに脆弱性。修正は2023年10月以降  - PC Watch](https://pc.watch.impress.co.jp/docs/news/1518766.html)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2023-20593]: https://nvd.nist.gov/vuln/detail/CVE-2023-20593
