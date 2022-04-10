+++
title = "spring-beans パッケージの RCE 脆弱性について"
date =  "2022-04-01T21:16:22+09:00"
description = "通称 “Spring4shell”"
image = "/images/attention/kitten.jpg"
tags  = [ "programming", "java", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

通称 “Spring4shell” に絡む情報は錯綜していて，実は私も勘違いしていたのだが，情報の整理のために記事をのこしておく。

Java の Spring Framework に関する脆弱性報告は大きく2つあって，どちらも RCE (Remote Code Execution) 脆弱性だからか，両者がとっちらかっている印象を受ける。

報告された脆弱性のひとつは Spring Cloud Function に絡むもので（[CVE-2022-22963]），これに関して詳しくは VMWare の記事を見ていただきたい。
これはこれで結構ヤバいっす。

- [CVE-2022-22963: Remote code execution in Spring Cloud Function by malicious Spring Expression | Security | VMware Tanzu](https://tanzu.vmware.com/security/cve-2022-22963)
- [Remote Code Execution in org.springframework.cloud:spring-cloud-function-context | CVE-2022-22963 | Snyk](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORKCLOUD-2436645)

で，もうひとつが本命の org.springframework:spring-beans パッケージの RCE 脆弱性，通称 “Spring4shell” である（[CVE-2022-22965]）。

- [CVE-2022-22965: Spring Framework RCE via Data Binding on JDK 9+ | Security | VMware Tanzu](https://tanzu.vmware.com/security/cve-2022-22965)
- [Spring4Shell: What we know about the Java RCE vulnerability | Snyk](https://snyk.io/blog/is-there-such-a-thing-as-spring4shell/)
- [Remote Code Execution in org.springframework:spring-beans | CVE-2022-22965 | Snyk](https://security.snyk.io/vuln/SNYK-JAVA-ORGSPRINGFRAMEWORK-2436751)

[Snyk のブログ記事][Spring4Shell]によると

{{< fig-quote type="markdown" title="Spring4Shell: What we know about the Java RCE vulnerability" link="https://snyk.io/blog/is-there-such-a-thing-as-spring4shell/" lang="en" >}}
If you’ve used the @Autowired annotation or utilized the magic of constructor injection, you’ve encountered dependency injection in the Spring ecosystem.

In affected versions, an RCE is achievable by manipulating the ClassLoader via a carefully composed HTTP POST request.

At this time, the exploit is only known to be possible with a Java Runtime Environment (JRE) version 9 or greater AND Tomcat version 9 or greater.
{{< /fig-quote >}}

だそうな。
これはヤバいわ。

対策としては，既に出ている Spring Framework の最新バージョン（5.2.20 または5.3.18 以降）に更新すればよい。
みんな大好き Spring Boot については，これを含んだバージョンが既に登場している。

- [Release v2.5.12 · spring-projects/spring-boot · GitHub](https://github.com/spring-projects/spring-boot/releases/tag/v2.5.12)
- [Release v2.6.6 · spring-projects/spring-boot · GitHub](https://github.com/spring-projects/spring-boot/releases/tag/v2.6.6)

というわけで，慌てず騒がず，冷静に対処しましょう。
アップデートは計画的に。

- `CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H`
- 深刻度: 緊急 (Score: 9.8)

| 基本評価基準 | 評価値 |
|--------|-------|
| 攻撃元区分 | ネットワーク |
| 攻撃条件の複雑さ | 低 |
| 必要な特権レベル | 不要 |
| ユーザ関与レベル | 不要 |
| スコープ | 変更なし |
| 機密性への影響 | 高 |
| 完全性への影響 | 高 |
| 可用性への影響 | 高 |

## ブックマーク

- [Spring Framework RCE, Early Announcement](https://spring.io/blog/2022/03/31/spring-framework-rce-early-announcement)
- [VU#970766 - Spring Framework insecurely handles PropertyDescriptor objects with data binding](https://kb.cert.org/vuls/id/970766)
- [Spring Frameworkの任意のコード実行の脆弱性（CVE-2022-22965）について](https://www.jpcert.or.jp/newsflash/2022040101.html)
- [Spring Frameworkの脆弱性 CVE-2022-22965（Spring4shell）についてまとめてみた - piyolog](https://piyolog.hatenadiary.jp/entry/2022/04/01/065946)
- [SpringShell RCE vulnerability: Guidance for protecting against and detecting CVE-2022-22965 - Microsoft Security Blog](https://www.microsoft.com/security/blog/2022/04/04/springshell-rce-vulnerability-guidance-for-protecting-against-and-detecting-cve-2022-22965/)
- [CVE-2022-22947: Spring Cloud Gateway Code Injection Vulnerability | Security | VMware Tanzu](https://tanzu.vmware.com/security/cve-2022-22947)

## 参考図書

{{% review-paapi "B0893LQ5KY" %}} <!-- Spring Boot 2 入門 -->
{{% review-paapi "B07CKHR8C1" %}} <!-- Spring Data JPAプログラミング入門 -->
{{% review-paapi "4621303252" %}} <!-- Effective Java 第3版 -->

[CVE-2022-22963]: https://nvd.nist.gov/vuln/detail/CVE-2022-22963
[CVE-2022-22965]: https://nvd.nist.gov/vuln/detail/CVE-2022-22965
[Spring4Shell]: https://snyk.io/blog/is-there-such-a-thing-as-spring4shell/ "Spring4Shell: What we know about the Java RCE vulnerability | Snyk"
