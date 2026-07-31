+++
title = "OpenJDK のセキュリティ・アップデート（2022-07）"
date =  "2022-08-14T09:43:22+09:00"
description = "影響を受けるバージョンは 18.0.1, 17.0.3, 15.0.7, 13.0.11, 11.0.15, 8u332, 7u341 およびそれ以前。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

サボってました。
まじすんません。

今更だけど，2022年7月に定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2022/07/19](https://openjdk.org/groups/vulnerability/advisories/2022-07-19)

CVE ID ベースで4個の脆弱性修正がある。
影響を受けるバージョンは 18.0.1, 17.0.3, 15.0.7, 13.0.11, 11.0.15, 8u332, 7u341 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2022/07/19" link="https://openjdk.org/groups/vulnerability/advisories/2022-07-19" lang="en" >}}
<table class="risk-matrix center smaller" summary="Risk matrix">
<tr>
<th rowspan="2">CVE ID</th>
<th rowspan="2">Component</th>
<th rowspan="2"><a href="https://www.first.org/cvss/">CVSSv3.1</a><br>Score</th>
<th colspan="7">Affects ...</th>
</tr>
<tr>
<th>7</th>
<th>8</th>
<th>11</th>
<th>13</th>
<th>15</th>
<th>17</th>
<th>18</th>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-34169">CVE-2022-34169</a></td>
<td style="text-align:left;"><code>xml/jaxp</code></td>
<td>7.5</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21541">CVE-2022-21541</a></td>
<td style="text-align:left;"><code>hotspot/runtime</code></td>
<td>5.9</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21549">CVE-2022-21549</a></td>
<td style="text-align:left;"><code>core-libs/java.util</code></td>
<td>5.3</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>{{% icons "check" %}}</td>
<td></td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21540">CVE-2022-21540</a></td>
<td style="text-align:left;"><code>hotspot/compiler</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

</table>
{{< /fig-quote >}}

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/18/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk18.0.2/f6ad4b4450fd4d298113270ec84f30ee/9/GPL/openjdk-18.0.2_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-18.0.2_linux-x64_bin.tar.gz
$ sudo ln -s jdk-18.0.2 java
$ java -version # すでに PATH が通っている場合
openjdk version "18.0.2" 2022-07-19
OpenJDK Runtime Environment (build 18.0.2+9-61)
OpenJDK 64-Bit Server VM (build 18.0.2+9-61, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-06-22 に [PlantUML] V1.2022.6 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

よーし，うむうむ，よーし。

[OpenJDK]: http://openjdk.java.net/
[Adoptium]: https://adoptium.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
