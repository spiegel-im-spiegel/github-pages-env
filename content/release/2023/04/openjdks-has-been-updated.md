+++
title = "OpenJDK のセキュリティ・アップデート（2023-04）"
date =  "2023-04-19T08:11:50+09:00"
description = "影響を受けるバージョンは 20, 17.0.6, 11.0.18, 8u362 およびそれ以前。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2023/04/18](https://openjdk.org/groups/vulnerability/advisories/2023-04-18)

CVE ID ベースで7個の脆弱性修正がある。
影響を受けるバージョンは 20, 17.0.6, 11.0.18, 8u362 およびそれ以前。
7, 13, 15 はついに対象から外れちゃったんだね。

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
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21930">CVE-2023-21930</a></td>
<td style="text-align:left;"><code>security-libs/<br>javax.net.ssl</code></td>
<td>7.4</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21954">CVE-2023-21954</a></td>
<td style="text-align:left;"><code>hotspot/<br>gc</code></td>
<td>5.9</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td></td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21967">CVE-2023-21967</a></td>
<td style="text-align:left;"><code>security-libs/<br>javax.net.ssl</code></td>
<td>5.9</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21939">CVE-2023-21939</a></td>
<td style="text-align:left;"><code>client-libs/<br>javax.swing</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21938">CVE-2023-21938</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.lang</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21937">CVE-2023-21937</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.net</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2023-21968">CVE-2023-21968</a></td>
<td style="text-align:left;"><code>core-libs/<br>java.nio</code></td>
<td>3.7</td>
<td>{{% icons "check" %}}</td>
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
$ sudo curl -L "https://download.java.net/java/GA/jdk20.0.1/b4887098932d415489976708ad6d1a4b/9/GPL/openjdk-20.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-20.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-20.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "20.0.1" 2023-04-18
OpenJDK Runtime Environment (build 20.0.1+9-29)
OpenJDK 64-Bit Server VM (build 20.0.1+9-29, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2023-04-20 に [PlantUML] v1.2023.6 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1167" >}}

よーし，うむうむ，よーし。

## ブックマーク

（あとで補完）

[OpenJDK]: http://openjdk.java.net/
[Adoptium]: https://adoptium.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
