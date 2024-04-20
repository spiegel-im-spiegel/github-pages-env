+++
title = "オープンソース・プロジェクトの乗っ取りを試みる"
date =  "2024-04-20T17:59:26+09:00"
description = "OpenSSF および OpenJS Foundation からの警告"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "engineering" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Bruce Schneier 先生の記事](https://www.schneier.com/blog/archives/2024/04/other-attempts-to-take-over-open-source-projects.html "Other Attempts to Take Over Open Source Projects - Schneier on Security")経由：

- [Open Source Security (OpenSSF) and OpenJS Foundations Issue Alert for Social Engineering Takeovers of Open Source Projects – Open Source Security Foundation](https://openssf.org/blog/2024/04/15/open-source-security-openssf-and-openjs-foundations-issue-alert-for-social-engineering-takeovers-of-open-source-projects/)
- [Open Source Security (OpenSSF) and OpenJS Foundations Issue Alert for Social Engineering Takeovers of Open Source Projects | OpenJS Foundation](https://openjsf.org/blog/openssf-openjs-alert-social-engineering-takeovers)

先日の [XZ Utils に仕組まれたバックドア][今回]に関連して [OpenSSF (Open Source Security Foundation)][OpenSSF] と [OpenJS Foundation][OpenJS] よりオープンソース・プロジェクトの乗っ取りに関する警告が出ている。
両者とも同じ内容かな。

[今回]のような話は特異なケースというわけではないらしく，以下のような類似例（試みは失敗したようだが）を報告している。

{{< fig-quote type="markdown" title="Open Source Security (OpenSSF) and OpenJS Foundations Issue Alert for Social Engineering Takeovers of Open Source Projects" link="https://openssf.org/blog/2024/04/15/open-source-security-openssf-and-openjs-foundations-issue-alert-for-social-engineering-takeovers-of-open-source-projects/" lang="en" >}}
The OpenJS Foundation Cross Project Council received a suspicious series of emails with similar messages, bearing different names and overlapping GitHub-associated emails. These emails implored OpenJS to take action to update one of its popular JavaScript projects to “address any critical vulnerabilities,” yet cited no specifics. The email author(s) wanted OpenJS to designate them as a new maintainer of the project despite having little prior involvement. This approach bears strong resemblance to the manner in which “Jia Tan” positioned themselves in the XZ/liblzma backdoor.

[...]

The OpenJS team also recognized a similar suspicious pattern in two other popular JavaScript projects not hosted by its Foundation, and immediately flagged the potential security concerns to respective OpenJS leaders, and the Cybersecurity and Infrastructure Security Agency (CISA) within the United States Department of Homeland Security (DHS).
{{< /fig-quote >}}

これらを踏まえ social engineering による乗っ取りのパターンとして以下を挙げている。

{{< fig-quote type="markdown" title="Open Source Security (OpenSSF) and OpenJS Foundations Issue Alert for Social Engineering Takeovers of Open Source Projects" link="https://openssf.org/blog/2024/04/15/open-source-security-openssf-and-openjs-foundations-issue-alert-for-social-engineering-takeovers-of-open-source-projects/" lang="en" >}}
- Friendly yet aggressive and persistent pursuit of maintainer or their hosted entity (foundation or company) by relatively unknown members of the community.
- Request to be elevated to maintainer status by new or unknown persons.
- Endorsement coming from other unknown members of the community who may also be using false identities, also known as “sock puppets.”
- PRs containing blobs as artifacts.
  - For example, the XZ backdoor was a cleverly crafted file as part of the test suite that wasn’t human readable, as opposed to source code.
- Intentionally obfuscated or difficult to understand source code.
- Gradually escalating security issues.
  - For example, the XZ issue started off with a relatively innocuous replacement of safe_fprintf() with fprintf() to see who would notice.
- Deviation from typical project compile, build, and deployment practices that could allow the insertion of external malicious payloads into blobs, zips, or other binary artifacts.
- A false sense of urgency, especially if the implied urgency forces a maintainer to reduce the thoroughness of a review or bypass a control.
{{< /fig-quote >}}

まぁ，悪人顔の悪人はいないってね。
漫画やドラマならともかく，現実の詐欺師は友好的かつ誠実そうな顔をしてやってくる（笑） 企業・組織などへの標的型攻撃もそうだけど，安全な「距離」をはかりながら徐々に侵食していく感じだよね。
これを防ぐのはなかなか難しいだろう。
特に小規模の FOSS プロジェクトなんかでは。

件の記事では，オープンソース・プロジェクトを保護するための手順として

{{< fig-quote type="markdown" title="Open Source Security (OpenSSF) and OpenJS Foundations Issue Alert for Social Engineering Takeovers of Open Source Projects" link="https://openssf.org/blog/2024/04/15/open-source-security-openssf-and-openjs-foundations-issue-alert-for-social-engineering-takeovers-of-open-source-projects/" lang="en" >}}
- Consider following industry-standard security best practices such as [OpenSSF Guides](https://openssf.org/resources/guides/).
- Use strong authentication.
  - Enable two-factor authentication (2FA) or Multifactor Authentication (MFA).
  - Use a secure password manager.
  - Preserve your recovery codes in a safe, preferably offline place.
  - Do not reuse credentials/passwords across different services.
- Have a security policy including a “[coordinated disclosure](https://github.com/ossf/oss-vulnerability-guide)” process for reports.
- Use best practices for merging new code.
  - Enable branch protections and signed commits.
  - If possible, have a second developer conduct code reviews before merging, even when the PR comes from a maintainer.
  - Enforce readability requirements to ensure new PRs are not obfuscated, and use of opaque binaries is minimized.
  - Limit who has npm publish rights.
  - Know your committers and maintainers, and do a periodic review. Have you seen them in your working group meetings or met them at events, for example?
- If you run an open source package repository, consider adopting [Principles for Package Repository Security](https://repos.openssf.org/principles-for-package-repository-security).
- Review “[Avoiding social engineering and phishing attacks](https://www.cisa.gov/news-events/news/avoiding-social-engineering-and-phishing-attacks)” from CISA and/or “[What is ‘Social Engineering’](https://www.enisa.europa.eu/topics/incident-response/glossary/what-is-social-engineering)” from ENISA.
{{< /fig-quote >}}

を挙げている。
比較的大きなコミュニティならこれでもいいんだろうけどねぇ。
サプライチェーンの観点では重要だけどオープンソース・プロジェクトとしては小規模で，殆どワンオペで回してるようなところは難しいかもしれない。
[今回]の XZ Utils のように。

[OpenSSF]: https://openssf.org/ "Open Source Security Foundation – Linux Foundation Projects"
[OpenJS]: https://openjsf.org/ "A safe and modern home for JavaScript technologies | OpenJS Foundation"
[今回]: {{< relref "./xz-utils-backdoor.md" >}} "XZ Utils に仕組まれたバックドアに関する覚え書き"

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
{{% review-paapi "4296001574" %}} <!-- ハッキング思考 -->
