+++
title = "Go のコードでも GitHub Code Scanning が使えるらしい"
date =  "2020-10-01T18:03:46+09:00"
description = "うん。簡単！"
image = "/images/attention/go-logo_blue.png"
tags = [ "security", "vulnerability", "risk", "management", "golang", "github", "continuous-integration" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[GitHub] の Code Scanning 機能が全ユーザで有効になった。

- [Github Code Scanning](https://iamninad.com/github-code-scanning/)
- [Code scanning is now available! - The GitHub Blog](https://github.blog/2020-09-30-code-scanning-is-now-available/)

「えー。どうせ [Go 言語][Go]では使えないんでしょ？」と思ったが， C/C++, C#, [Go], Java, JavaScript/TypeScript, Python で有効らしい。
というわけで，とりあえず手持ちの [Go] パッケージのリポジトリで試してみた。

Code Scanning は各リポジトリの “Security” で設定できる。

{{< fig-img src="./github-security-menu.png" link="./github-security-menu.png" width="1284" >}}

一番下の “Code scanning alerts” の `[Set up code scanning]` ボタンを押す。
次に表示される以下の画面で

{{< fig-img src="./code-scanning-setting.png" link="./code-scanning-setting.png" width="682" >}}

`[Set up this workflow]` ボタンを押す。

すると [GitHub] Actions 用の YAML ファイル編集画面が表示される。
今回は [Go] コードのリポジトリなので，こんな感じの内容になった。

```yaml
# For most projects, this workflow file will not need changing; you simply need
# to commit it to your repository.
#
# You may wish to alter this file to override the set of languages analyzed,
# or to provide custom queries or build logic.
name: "CodeQL"

on:
  push:
    branches: [master]
  pull_request:
    # The branches below must be a subset of the branches above
    branches: [master]
  schedule:
    - cron: '0 6 * * 4'

jobs:
  analyze:
    name: Analyze
    runs-on: ubuntu-latest

    strategy:
      fail-fast: false
      matrix:
        # Override automatic language detection by changing the below list
        # Supported options are ['csharp', 'cpp', 'go', 'java', 'javascript', 'python']
        language: ['go']
        # Learn more...
        # https://docs.github.com/en/github/finding-security-vulnerabilities-and-errors-in-your-code/configuring-code-scanning#overriding-automatic-language-detection

    steps:
    - name: Checkout repository
      uses: actions/checkout@v2
      with:
        # We must fetch at least the immediate parents so that if this is
        # a pull request then we can checkout the head.
        fetch-depth: 2

    # If this run was triggered by a pull request event, then checkout
    # the head of the pull request instead of the merge commit.
    - run: git checkout HEAD^2
      if: ${{ github.event_name == 'pull_request' }}

    # Initializes the CodeQL tools for scanning.
    - name: Initialize CodeQL
      uses: github/codeql-action/init@v1
      with:
        languages: ${{ matrix.language }}
        # If you wish to specify custom queries, you can do so here or in a config file.
        # By default, queries listed here will override any specified in a config file. 
        # Prefix the list here with "+" to use these queries and those in the config file.
        # queries: ./path/to/local/query, your-org/your-repo/queries@main

    # Autobuild attempts to build any compiled languages  (C/C++, C#, or Java).
    # If this step fails, then you should remove it and run the build manually (see below)
    - name: Autobuild
      uses: github/codeql-action/autobuild@v1

    # ℹ️ Command-line programs to run using the OS shell.
    # 📚 https://git.io/JvXDl

    # ✏️ If the Autobuild fails above, remove it and uncomment the following three lines
    #    and modify them (or add more) to build your code if your project
    #    uses a compiled language

    #- run: |
    #   make bootstrap
    #   make release

    - name: Perform CodeQL Analysis
      uses: github/codeql-action/analyze@v1
```

[Go] コードの検査だけならこのままコミットしてしまって構わない。
ちなみにコミットする場合は `master` ブランチ[^br1] に対してではなく pull request 用のブランチを作ってそこにコミットしたほうがよい。
そうすれば PR 時に上記設定の action が走るので動作確認になるだろう。

[^br1]: 2020年10月から [GitHub の新規リポジトリの既定ブランチ名が `main` になるらしい]({{< ref "/remark/2020/08/renaming-default-branch-name-in-github-repositries.md" >}} "GitHub リポジトリの既定ブランチ名が main になるらしい")。ご注意を。

うん。
簡単！

## ブックマーク

- [Workflow syntax for GitHub Actions - GitHub Docs](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstepsrun)
- [GitHub、コードの脆弱性を検出する「Code Scanning」を全ユーザーに提供 - ZDNet Japan](https://japan.zdnet.com/article/35160321/)
- [GitHub、コードの脆弱性などを発見してくれる「GitHub Code Scanning」正式版が提供開始。パブリックリポジトリには無料 － Publickey](https://www.publickey1.jp/blog/20/githubgithub_code_scanning.html)
- [GitHub、開発者がセキュリティ脆弱性を発見するための支援機能「Code Scanning」 | マイナビニュース](https://news.mynavi.jp/article/20201002-1364892/)
- [Go 依存パッケージの脆弱性検査]({{< ref "/golang/check-for-vulns-in-golang-dependencies.md" >}})

[Go]: https://golang.org/ "The Go Programming Language"
[GitHub]: https://github.com/
