+++
title = "OpenJDK のセキュリティ・アップデート（2023-07）"
date =  "2023-07-19T21:37:42+09:00"
description = "影響を受けるバージョンは 20.0.1, 17.0.7, 11.0.19, 8u372 およびそれ以前。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2023/07/18](https://openjdk.org/groups/vulnerability/advisories/2023-07-18)

CVE ID ベースで8個の脆弱性修正がある。
影響を受けるバージョンは 20.0.1, 17.0.7, 11.0.19, 8u372 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2023/04/18" link="https://openjdk.org/groups/vulnerability/advisories/2023-04-18" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3.1</a><br>Score</th>
<th colspan="7">Affects ...</th>
</tr>
<tr>
<th>8</th>
<th>11</th>
<th>17</th>
<th>20</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22041">CVE-2023-22041</a></td>
<td style="text-align:left;"><code>hotspot/<br>compiler</code></td>
<td>5.1</td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-25193">CVE-2023-25193</a></td>
<td style="text-align:left;"><code>client-libs/<br>2d</code></td>
<td>3.7</td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22044">CVE-2023-22044</a></td>
<td style="text-align:left;"><code>hotspot/<br>compiler</code></td>
<td>3.7</td>
<td></td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22045">CVE-2023-22045</a></td>
<td style="text-align:left;"><code>hotspot/<br>compiler</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22049">CVE-2023-22049</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.io</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22036">CVE-2023-22036</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.util</code></td>
<td>3.7</td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22006">CVE-2023-22006</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.net</code></td>
<td>3.1</td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-22043">CVE-2023-22043</a></td>
<td style="text-align:left;"><code>javafx/<br>graphics</code></td>
<td>5.9</td>
<td>-</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/20/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk20.0.2/6e380f22cbe7469fa75fb448bd903d8e/9/GPL/openjdk-20.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-20.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-20.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "20.0.2" 2023-07-18
OpenJDK Runtime Environment (build 20.0.2+9-78)
OpenJDK 64-Bit Server VM (build 20.0.2+9-78, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2023-07-12 に [PlantUML] V1.2023.10 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1167" >}}

よーし，うむうむ，よーし。



## ブックマーク

- [Oracle Critical Patch Update Advisory - July 2023](https://www.oracle.com/security-alerts/cpujul2023.html)
- [Oracle Java の脆弱性対策について(CVE-2023-22043等) | 情報セキュリティ | IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/security-alert/2023/0719-jre.html)
- [2023年7月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2023/at230012.html)

[OpenJDK]: http://openjdk.java.net/
[Adoptium]: https://adoptium.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
