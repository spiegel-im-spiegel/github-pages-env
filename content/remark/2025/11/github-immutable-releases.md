+++
title = "GitHub「変更不可リリース（Immutable Releases）」について"
date =  "2025-11-11T10:25:44+09:00"
description = "私のリポジトリで採用するかどうかは分からないが，覚え書きとして記しておく。"
image = "/images/attention/kitten.jpg"
tags = [ "github", "engineering", "security", "risk", "management" ]
pageType = "text"
draft = true

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

気づくのが遅れたが，先日 GitHub が「変更不可リリース[^ir1]（Immutable Releases）」機能を正式リリースしたらしい。

[^ir1]: GitHub の公式ドキュメントで「変更不可リリース」「不変リリース」「変更できないリリース」など訳語が混在しているが，この記事では「変更不可リリース」としておく。これって生成 AI に翻訳させてるのかね。用語の統一くらいしておけよ。

- [Immutable releases are now generally available - GitHub Changelog](https://github.blog/changelog/2025-10-28-immutable-releases-are-now-generally-available/)
- [変更不可リリース - GitHub ドキュメント](https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/immutable-releases)
  - [リリースに対する変更の防止 - GitHub ドキュメント](https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/preventing-changes-to-your-releases)
  - [リリースの整合性の検証 - GitHub ドキュメント](https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/verifying-the-integrity-of-a-release)
- [GitHub、リリース後のバイナリなどアセットを変更不可にする「Immutable Release」（変更不可リリース）機能を正式リリース － Publickey](https://www.publickey1.jp/blog/25/githubimmutable_release.html)

私のリポジトリでこれを採用するかどうかは分からないが，覚え書きとして記しておく。

リポジトリに対して「変更不可リリース」を設定すると以下の機能が有効になる。

{{< fig-quote type="markdown" title="変更不可リリース - GitHub ドキュメント" link="https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/immutable-releases" >}}
変更できないリリースを有効にすると、次の保護が適用されます。

- **Git タグは移動または削除できません**。変更できないリリースが発行されると、関連付けられている Git タグは特定のコミットにロックされ、変更または削除することはできません。
- **リリースアセットを変更または削除することはできません**:リリースにアタッチされているすべてのファイル (バイナリやアーカイブなど) は、変更または削除から保護されます。

さらに、変更不可リリースを作成すると、**リリース構成証明**が自動的に生成されます。これは、リリース タグ、コミット SHA、およびリリース アセットを含むリリースに関する、暗号で検証可能なレコードです。 コンシューマーはこの構成証明を使用して、使用しているリリースと成果物が、公開された GitHub リリースと完全に一致することを確認できます。
{{< /fig-quote >}}

さらに

{{< fig-quote type="markdown" title="変更不可リリース - GitHub ドキュメント" link="https://docs.github.com/ja/code-security/supply-chain-security/understanding-your-software-supply-chain/immutable-releases" >}}
不変リリースには、リポジトリの復活攻撃に対する保護が含まれます。 リポジトリを削除し、同じ名前の新しいリポジトリを作成した場合でも、元のリポジトリの変更できないリリースに関連付けられたタグを再利用することはできません。
{{< /fig-quote >}}

なんだそうだ。
リポジトリを移動したらどうなるんだろう？

「変更不可リリース」を有効にする前のリリースバージョンにはこれらの機能は適用されない。
逆に「変更不可リリース」有効状態でリリースしたバージョンについては「変更不可リリース」設定を解除した後も機能が残り続ける。

{{< fig-quote type="markdown" title="Immutable releases are now generally available" link="https://github.blog/changelog/2025-10-28-immutable-releases-are-now-generally-available/" lang="en" >}}
Once enabled:

- All new releases are immutable (i.e., assets are locked and tags are protected).
- Existing releases remain mutable unless you republish them.

Disabling immutability doesn’t affect releases created while it was enabled. They remain immutable.
{{< /fig-quote >}}


「変更不可リリース」はリポジトリ単位で設定できる。
リポジトリの “Settings” を開き “General” 項目の “Releases” にあるチェック項目 “Enable release immutability” をチェックする。

{{< fig-img-quote src="./repository-settings.png" title="Repository Settings" link="./repository-settings.png" lang="en" width="1429" >}}

Organization の場合は Settings で「変更不可リリース」のポリシーを設定できる。
“Code, planning, and automation” の “Repository” → “General” 項目を開き “Releases” を “All repositories” または “Selected repositories” に設定する。

{{< fig-img-quote src="./organization-settings.png" title="Organization Settings" link="./organization-settings.png" lang="en" width="1580" >}}

“All repositories” を選択すると Organization 内の全リポジトリに「変更不可リリース」が適用される。
“Selected repositories” を選択すると適用するリポジトリを選択できる。

{{< fig-img-quote src="./selected-repositories.png" title="Selected repositories" link="./selected-repositories.png" lang="en" width="959" >}}

まぁ [Go] のパッケージに関しては， [Go] のエコシステムの中で運用する限り，リポジトリのバージョンタグは実質的に immutable なんだよな。

- [Go はどのようにしてサプライチェーン攻撃を低減しているか](https://zenn.dev/spiegel/articles/20220402-how-go-mitigates-supply-chain-attacks)

それでも GitHub のリリースページでバイナリを配布している場合は「変更不可リリース」機能は有用なので，導入を検討してもいいかもしれない。

[Go]: https://go.dev/

## 参考図書

{{% review-paapi "4814401191" %}} <!-- 初めてのGo言語 第2版 -->
