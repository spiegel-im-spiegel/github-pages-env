+++
title = "OpenJDK のセキュリティ・アップデート"
date =  "2022-04-23T12:34:49+09:00"
description = "影響を受けるバージョンは 18, 17.0.2, 15.0.6, 13.0.10, 11.0.14, 8u322, 7u331 およびそれ以前。"
image = "/images/attention/tools.png"
tags  = [ "programming", "language", "java", "ubuntu", "plantuml", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

定例の Java マイナー・バージョンアップが行われた。

- [OpenJDK Vulnerability Advisory: 2022/04/19](https://openjdk.java.net/groups/vulnerability/advisories/2022-04-19)

CVE ID ベースで16個の脆弱性修正がある。
影響を受けるバージョンは 18, 17.0.2, 15.0.6, 13.0.10, 11.0.14, 8u322, 7u331 およびそれ以前。

{{< fig-quote class="nobox" title="OpenJDK Vulnerability Advisory: 2022/04/19" link="https://openjdk.java.net/groups/vulnerability/advisories/2022-04-19" lang="en" >}}
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
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21476">CVE-2022-21476</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>7.5</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td></td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21449">CVE-2022-21449</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>7.5</td>
<td></td>
<td></td>
<td></td>
<td></td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21496">CVE-2022-21496</a></td>
<td style="text-align:left;"><code>core-libs/javax.naming</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21434">CVE-2022-21434</a></td>
<td style="text-align:left;"><code>core-libs/java.lang</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21426">CVE-2022-21426</a></td>
<td style="text-align:left;"><code>xml/jaxp</code></td>
<td>5.3</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
<td>{{% icons "check" %}}</td>
</tr>

<tr>
<td style="text-align:left;"><a href="https://nvd.nist.gov/vuln/detail/CVE-2022-21443">CVE-2022-21443</a></td>
<td style="text-align:left;"><code>security-libs/java.security</code></td>
<td>3.7</td>
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

このうち，2番目の [CVE-2022-21449](https://nvd.nist.gov/vuln/detail/CVE-2022-21449) は ECDSA に関するヤバいやつで
Bruce Schneier 先生もこの脆弱性について[言及](https://www.schneier.com/blog/archives/2022/04/java-cryptography-implementation-mistake-allows-digital-signature-forgeries.html "Java Cryptography Implementation Mistake Allows Digital-Signature Forgeries - Schneier on Security")されている。
元ネタの記事によると

{{< fig-quote type="markdown" title="Major cryptography blunder in Java enables “psychic paper” forgeries | Ars Technica" link="https://arstechnica.com/information-technology/2022/04/major-crypto-blunder-in-java-enables-psychic-paper-forgeries/" lang="en" >}}
The vulnerability, which [Oracle patched on Tuesday](https://www.oracle.com/security-alerts/cpuapr2022.html), affects the company’s implementation of the [Elliptic Curve Digital Signature Algorithm](https://en.wikipedia.org/wiki/Elliptic_Curve_Digital_Signature_Algorithm) in Java versions 15 and above. ECDSA is an algorithm that uses the principles of [elliptic curve cryptography](https://arstechnica.com/information-technology/2013/10/a-relatively-easy-to-understand-primer-on-elliptic-curve-cryptography/) to authenticate messages digitally. A key advantage of ECDSA is the smaller size of the keys it generates, compared to RSA or other crypto algorithms, making it ideal for use in standards including [FIDO-based 2FA](https://arstechnica.com/information-technology/2020/07/apple-has-finally-embraced-key-based-2fa-so-should-you/), the [Security Assertion Markup Language](https://en.wikipedia.org/wiki/Security_Assertion_Markup_Language), [OpenID](https://en.wikipedia.org/wiki/OpenID#OpenID_Connect_(OIDC)), and [JSON](https://tools.ietf.org/html/rfc7519).
{{< /fig-quote >}}

だそうな。
もう少し詳細を抜き出すと

{{< fig-quote type="markdown" title="Major cryptography blunder in Java enables “psychic paper” forgeries | Ars Technica" link="https://arstechnica.com/information-technology/2022/04/major-crypto-blunder-in-java-enables-psychic-paper-forgeries/" lang="en" >}}
ECDSA signatures rely on a pseudo-random number, typically notated as K, that’s used to derive two additional numbers, R and S. To verify a signature as valid, a party must check the equation involving R and S, the signer’s public key, and a cryptographic hash of the message. When both sides of the equation are equal, the signature is valid.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Major cryptography blunder in Java enables “psychic paper” forgeries | Ars Technica" link="https://arstechnica.com/information-technology/2022/04/major-crypto-blunder-in-java-enables-psychic-paper-forgeries/" lang="en" >}}
For the process to work correctly, neither R nor S can ever be a zero. That’s because one side of the equation is R, and the other is multiplied by R and a value from S. If the values are both 0, the verification check translates to 0 = 0 X (other values from the private key and hash), which will be true regardless of the additional values. That means an adversary only needs to submit a blank signature to pass the verification check successfully.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Major cryptography blunder in Java enables “psychic paper” forgeries | Ars Technica" link="https://arstechnica.com/information-technology/2022/04/major-crypto-blunder-in-java-enables-psychic-paper-forgeries/" lang="en" >}}
Guess which check Java forgot?

That’s right. Java’s implementation of ECDSA signature verification didn’t check if R or S were zero, so you could produce a signature value in which they are both 0 (appropriately encoded) and Java would accept it as a valid signature for any message and for any public key. The digital equivalent of a blank ID card.
{{< /fig-quote >}}

いやいやいや。
なんだその間抜けなバグ。

[OpenJDK] を APT で管理するメリットはないので，実行バイナリを[リリースページ](https://jdk.java.net/18/)から直接ダウンロードして配置する。
以下は完全手動でのインストール（笑）

```text
$ cd /usr/local/src
$ sudo curl -L "https://download.java.net/java/GA/jdk18.0.1/3f48cabb83014f9fab465e280ccf630b/10/GPL/openjdk-18.0.1_linux-x64_bin.tar.gz" -O
$ cd ..
$ sudo unlink java # 以前のバージョンの Java 環境がある場合
$ sudo tar xvf src/openjdk-18.0.1_linux-x64_bin.tar.gz
$ sudo ln -s jdk-18.0.1 java
$ java -version # すでに PATH が通っている場合
openjdk version "18.0.1" 2022-04-19
OpenJDK Runtime Environment (build 18.0.1+10-24)
OpenJDK 64-Bit Server VM (build 18.0.1+10-24, mixed mode, sharing)
```

LTS 版 Java バイナリが欲しいなら [Adoptium](https://adoptium.net/) で取得できる。

私としては [PlantUML] が動けばいいので，試しておく[^puml1]。

[^puml1]: 2021-04-09 に [PlantUML] V1.2022.4 が[リリース](http://plantuml.com/changes)されている。 [PlantUML] の使い方等については拙文「[真面目に PlantUML]({{< ref "/remark/2018/12/plantuml-1.md" >}})」シリーズを参考にどうぞ。

{{< fig-img src="./factory-method-pattern.png" title="Factory Method Pattern" link="./factory-method-pattern.puml" width="1160" >}}

よーし，うむうむ，よーし。

## ブックマーク

- [Oracle Critical Patch Update Advisory - April 2022](https://www.oracle.com/security-alerts/cpuapr2022.html)
- [Critical cryptographic Java security blunder patched – update now! – Naked Security](https://nakedsecurity.sophos.com/2022/04/20/critical-cryptographic-java-security-blunder-patched-update-now/)
- [2022年4月Oracle製品のクリティカルパッチアップデートに関する注意喚起](https://www.jpcert.or.jp/at/2022/at220012.html)
- [「Java」に署名検証がフリーパスになってしまう危険な脆弱性 ～影響は計り知れず - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1404535.html)

[OpenJDK]: http://openjdk.java.net/
[Adoptium]: https://adoptium.net/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[PlantUML]: http://plantuml.com/ "Open-source tool that uses simple textual descriptions to draw UML diagrams."

## 参考図書

{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->
{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
