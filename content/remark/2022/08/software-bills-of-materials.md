+++
title = "オープンソース製品とソフトウェア部品表"
date =  "2022-08-09T20:31:46+09:00"
description = "ブログってのはこういう粗結合連鎖が面白いと思うのだけどねぇ"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "engineering", "management" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

拙作の [depm](https://github.com/goark/depm "goark/depm: Visualize depndency packages and modules") で利用している [github.com/google/licenseclassifier](https://github.com/google/licenseclassifier "google/licenseclassifier: A License Classifier") パッケージの v2 系モジュールがいい感じにバージョンが上がってきたので，どのくらい使えるようになったのか試してみた記事がこれ。

- [ライセンスファイルからライセンスを推定する](https://zenn.dev/spiegel/articles/20220806-licenseclassifier)

で，この記事について有り難くも「[SPDX License identifier にも触れて欲しい](https://twitter.com/fu7mu4/status/1556141959755886593)」というリクエストを頂いたので [SPDX (Software Package Data Exchange)][SPDX] についてちょろんと紹介する文章を追記した。
これを書くためにサイトを眺めて初めて気がついたのだが [SPDX] が [ISO/IEC 5962:2021] として標準化されていたらしい。

- [SPDX Becomes Internationally Recognized Standard for Software Bill of Materials - Linux Foundation](https://www.linuxfoundation.org/featured/spdx-becomes-internationally-recognized-standard-for-software-bill-of-materials/)

{{< fig-quote type="markdown" title="SPDX Becomes Internationally Recognized Standard for Software Bill of Materials" link="https://www.linuxfoundation.org/featured/spdx-becomes-internationally-recognized-standard-for-software-bill-of-materials/" lang="en" >}}
Between eighty and ninety percent (80%-90%) of a modern application is assembled from open source software components. An SBOM accounts for the software components contained in an application — open source, proprietary, or third-party — and details their provenance, license, and security attributes. SBOMs are used as a part of a foundational practice to track and trace components across software supply chains. SBOMs also help to proactively identify software issues and risks and establish a starting point for their remediation.
{{< /fig-quote >}}

[ISO/IEC 5962:2021] がリリースされたのが2021年8月。
その年の年末に例の [Apache Log4j の脆弱性]({{< ref "/remark/2021/12/log4j-vulnerability.md" >}} "ava 製 Logger “Log4j” の脆弱性について")に端を発したソフトウェア・サプライチェーン脆弱性の問題が大きく取り上げられることになった。

- [Apache Log4jの脆弱性とともに浮かび上がったオープンソースのメンテナの責任範囲の問題 - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20211222/apache-log4j)

どーりで今年（2022年）に入ってやたらと「ソフトウェア部品表（Software Bill of Materials; SBOM）」の話を聞くようになったわけだ。

更にタイミングのいいことに7月下旬に以下の記事が公開されていたのを yomoyomo さんの記事で知った。

- [Open-Source Security: How Digital Infrastructure Is Built on a House of Cards  - Lawfare](https://www.lawfareblog.com/open-source-security-how-digital-infrastructure-built-house-cards)
- [Securing Open-Source Software - Schneier on Security](https://www.schneier.com/blog/archives/2022/07/securing-open-source-software.html)
- [オープンソースのセキュリティ：デジタルインフラは砂上の楼閣に築かれている？ - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20220808/open-source-security)

{{< fig-quote type="markdown" title="Open-Source Security: How Digital Infrastructure Is Built on a House of Cards" link="https://www.lawfareblog.com/open-source-security-how-digital-infrastructure-built-house-cards" lang="en" >}}
Importantly, the culprit was not the developers of the code but the company that [failed to implement a patch](https://techcrunch.com/2018/12/10/equifax-breach-preventable-house-oversight-report/) that promised to prevent the very thing that happened. Many observers complain that Equifax has suffered little consequence for its negligence, highlighting weak oversight and accountability structures.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="オープンソースのセキュリティ：デジタルインフラは砂上の楼閣に築かれている？" link="https://yamdas.hatenablog.com/entry/20220808/open-source-security" >}}
これに対し、最近では連邦取引委員会が Log4Shell の対応パッチの適用が遅い企業を強制措置をかますなど、政府がオープンソースのセキュリティ問題に介入する姿勢を見せつつある。著者はその一環としての SBOM（Software Bill Of Materials：ソフトウェア部品表）を評価するが、それだけでは不十分と断じる。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Open-Source Security: How Digital Infrastructure Is Built on a House of Cards" link="https://www.lawfareblog.com/open-source-security-how-digital-infrastructure-built-house-cards" lang="en" >}}
An SBOM is simply a list of the ingredients, or codebases, that comprise software you purchased. It does not provide a list of vulnerabilities nor does it impose any minimum security requirements on the vendor generating the SBOM. Comparable to a list of ingredients on a snack or medication you purchase, the information is only as useful as your ability to parse it. 

To operationalize an SBOM, a company must be able to read it, which is a challenge as there is no mandated standard format for an SBOM, and actually use it to check databases such as the [National Vulnerability Database (NVD)](https://nvd.nist.gov/) for new vulnerabilities found in the software components the SBOM lists. These activities are costly and cumbersome. While Google and Intel might have the resources and security maturity to demand machine-readable SBOMs and regularly scan databases for new vulnerabilities that impact their systems, there are countless small businesses using open source that cannot.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="オープンソースのセキュリティ：デジタルインフラは砂上の楼閣に築かれている？" link="https://yamdas.hatenablog.com/entry/20220808/open-source-security" >}}
著者は、もはや公共財の性質を持つオープンソースを維持する制度的構造を構築する必要があると訴える。そして、それ自体は目新しい主張ではない。それには効率的な資源配分を確保し、最低基準を課すことが必要になるが、果たしてそれをオープンソースプロジェクトに適用するのはうまくいくかねというのがワタシの感想になる。
{{< /fig-quote >}}

個人的には FOSS 製品を「公共財」と見なす向きには違和感や危うさを感じてしまうのだが，もはや四の五の言ってられねー，って感じなのだろう。
せめて [SPDX] が SBOM の標準としてセキュリティ・リスク管理に上手く組み込まれることを期待したい。

報告される脆弱性の量的評価としての CVSS，ソフトウェア・サプライチェーンの構成を可視化する [SPDX]/SBOM，脆弱性が報告された際のアクションを支援する SSVC[^ssvc1] といった道具・手段を組み合わせて，脆弱性報告から対応までのワークフローがスムーズに流れるようになるといいなぁ，と思ったり。
まぁ，そのワークフロー自体がソフトウェア・サプライチェーンだったりするのだが（笑）

[^ssvc1]: SSVC (Stakeholder-Specific Vulnerability Categorization) については “{{< pdf-file title="Prioritizing Vulnerability Response: A Stakeholder-Specific Vulnerability Categorization (Version 2.0)" link="https://resources.sei.cmu.edu/library/asset-view.cfm?assetid=653459">}}” あたりを参照のこと。

しかし，ブログってのはこういう粗結合連鎖が面白いと思うのだが「[ブログの退潮](https://yamdas.hatenablog.com/entry/20220808/pc-crash "恥さらし文章「ある「パソコンの大先生」の死」に寄せられたありがたいコメントの数々 - YAMDAS現更新履歴")」はもはや避けられぬか（笑）

## [Synk] と SBOM

開発者向けのセキュリティ関連サービスを提供している [Synk] は [SPDX]/SBOM について以下のように述べている。

{{< fig-quote type="markdown" title="Advancing SBOM standards: Snyk and SPDX | Snyk" link="https://snyk.io/blog/advancing-sbom-standards-snyk-spdx/" lang="en" >}}
Snyk integrates with a wide range of different package managers and developer tools to help identify vulnerabilities in the software components used. In doing that we need to build up a SBOM under the hood, normalising the list of software and augmenting that list with additional metadata from other sources. The Snyk tooling mainly focuses on presenting that information alongside information about vulnerabilities, but Snyk customers can access that raw SBOM via our built-in reporting or [powerful API](https://snyk.docs.apiary.io/#reference/dependencies/dependencies-by-organization/list-all-dependencies).

In fact, you can think of Snyk client tools like the CLI and CI/CD plugins as generating an SBOM, while Snyk’s backend takes an SBOM and returns vulnerability data, or provides automation around that data to help you fix issues. It’s this extensive experience that leads to our interest in emerging standards in this space.
{{< /fig-quote >}}

## Microsoft と SBOM

忘れていたが Microsoft も SBOM を自動生成するツールを OSS で出してたんだっけ。

- [Microsoft open sources its software bill of materials (SBOM) generation tool - Engineering@Microsoft](https://devblogs.microsoft.com/engineering-at-microsoft/microsoft-open-sources-software-bill-of-materials-sbom-generation-tool/)
- [マイクロソフト、ビルド時にソフトウェアの部品表（SBOM）を自動生成する「SBOM Tool」、オープンソースで公開 － Publickey](https://www.publickey1.jp/blog/22/sbomsbom_tool.html)

{{< fig-quote type="markdown" title="マイクロソフト、ビルド時にソフトウェアの部品表（SBOM）を自動生成する「SBOM Tool」、オープンソースで公開" link="https://www.publickey1.jp/blog/22/sbomsbom_tool.html" >}}
具体的には、ソフトウェア名やライセンス、作成者、バージョン、個々のファイルのハッシュ値などの情報を、SBOMのフォーマットとしての標準の1つである「SPDX」（Software Package Data Exchange）形式で出力してくれます。

また、ビルドの対象となるソフトウェアにSBOMが用意されている場合、新たに生成されるSBOMにはそのSBOMの情報をきちんととりまとめて生成することが可能。
{{< /fig-quote >}}

まぁ，メインの文章でも紹介したように「SBOM をどう使うか」が重要なんだけどね。

## ブックマーク

- [SSVC（Stakeholder-Specific Vulnerability Categorization）を活用した脆弱性管理 | PwC Japanグループ](https://www.pwc.com/jp/ja/knowledge/column/awareness-cyber-security/stakeholder-specific-vulnerability-categorization.html)
- [Why SBOM Is Not Enough in the Era of Supply Chain Attacks](https://forelens.com/blog/why-sbom-is-not-enough-in-the-era-of-supply-chain-attacks/)
- [IPAがセキュリティ情報の収集ツール「mjcheck4」を公開 ～Adobe AIR製の旧版は廃止 - 窓の杜](https://forest.watch.impress.co.jp/docs/news/1489503.html) : SBOM をインポートできるらしい
- [SSVCを使った脆弱性管理の取り組み](https://www.conoris.jp/blog/ssvc)
- [経済産業省、「ソフトウェア管理に向けたSBOM（Software Bill of Materials）の導入に関する手引」公開。環境構築、SBOM作成、運用管理など解説 － Publickey](https://www.publickey1.jp/blog/23/sbomsoftware_bill_of_materialssbom.html)

- [Go はどのようにしてサプライチェーン攻撃を低減しているか](https://zenn.dev/spiegel/articles/20220402-how-go-mitigates-supply-chain-attacks)
- [「ブログはやはり『死に続けている』」]({{< ref "/remark/2022/07/slow-motion-car-crash.md" >}})

[SPDX]: https://spdx.dev/ "International Open Standard (ISO/IEC 5962:2021) - Software Package Data Exchange (SPDX)"
[ISO/IEC 5962:2021]: https://www.iso.org/standard/81870.html "ISO - ISO/IEC 5962:2021 - Information technology — SPDX® Specification V2.2.1"
[Synk]: https://snyk.io/ "Snyk | Developer security | Develop fast. Stay secure."

## 参考図書

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
