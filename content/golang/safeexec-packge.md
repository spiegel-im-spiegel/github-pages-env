+++
title = "github.com/cli/safeexec パッケージを使った外部コマンドの安全な起動"
date =  "2020-12-20T10:00:16+09:00"
description = "このパッケージを使って exec.LookPath() 関数を置き換えることを検討するのもいいだろう。"
image = "/images/attention/go-logo_blue.png"
tags = [ "programming", "golang", "security", "risk", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Hugo v0.79.1 のリリースノート](https://github.com/gohugoio/hugo/releases/tag/v0.79.1 "Release v0.79.1 · gohugoio/hugo")を見て気づいたのだが， GitHub が自身のコマンドライン・ツール用に [github.com/cli/safeexec][`safeexec`] パッケージを公開している。

これは [`os/exec`][`exec`] 標準パッケージ内の [`exec`]`.LookPath()` 関数を置き換えるもので

```go
import (
    "os/exec"
    "github.com/cli/safeexec"
)

func gitStatus() error {
    gitBin, err := safeexec.LookPath("git")
    if err != nil {
        return err
    }
    cmd := exec.Command(gitBin, "status")
    return cmd.Run()
}
```

てな感じに使うようだ。

以前に Zenn で紹介したが

- [Go でサブプロセスを起動する際は LookPath に気をつけろ！](https://zenn.dev/spiegel/articles/20201107-lookpath-by-golang)

Windows 環境で [`os/exec`][`exec`] 標準パッケージを使って外部コマンドをする際，パスなしのコマンド名を指定するとカレントにある同名の実行モジュールが優先的に起動されてしまう。
このブログでは [CVE-2020-27955] の脆弱性として[紹介]({{< ref "/release/2020/11/git-for-windows-2_29_2-2-is-released.md" >}} "Git for Windows 2.29.2 (2) のリリース【セキュリテイ・アップデート】")している。

[`os/exec`][`exec`] 標準パッケージを使って外部コマンドを起動している場合は [github.com/cli/safeexec][`safeexec`] パッケージで `LookPath()` 関数を置き換えることを検討するのもいいだろう。

なお，この問題は既に以下の issue として上がっているようだ（特に10月下旬辺りからの議論）。

- [proposal: os/exec: make LookPath not look in dot implicitly on Windows · Issue #38736 · golang/go · GitHub](https://github.com/golang/go/issues/38736)

どうなることやら。

## ブックマーク

- [Go パッケージ／モジュールの依存関係可視化ツール Depm v0.3.0 をリリースした]({{< ref "/release/2020/11/depm-v0_3_0-is-released.md" >}}) : 拙作では [github.com/cli/safeexec][`safeexec`] パッケージ と同じような方針で [`exec`]`.LookPath()` 関数の問題を回避している

[Go]: https://golang.org/ "The Go Programming Language"
[`safeexec`]: https://github.com/cli/safeexec "cli/safeexec: A safer version of exec.LookPath on Windows"
[`exec`]: https://golang.org/pkg/os/exec/ "exec - The Go Programming Language"
[CVE-2020-27955]: https://nvd.nist.gov/vuln/detail/CVE-2020-27955 "NVD - CVE-2020-27955"

## 参考図書

{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
