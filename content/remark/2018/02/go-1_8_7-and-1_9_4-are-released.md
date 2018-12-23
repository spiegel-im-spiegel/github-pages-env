+++
title = "Go 1.8.7 および 1.9.4 がリリース（セキュリティ・アップデート）"
date = "2018-02-09T20:19:26+09:00"
description = "#cgo ディレクティブを含む Go 言語コードをビルドする際に任意の処理を実行される可能性がある。"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "vulnerability",
  "golang",
]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = ""
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = ""

[scripts]
  mathjax = false
  mermaidjs = false
+++

[Go 言語]コンパイラの 1.8.7 および 1.9.4 がリリースされている。
このバージョンでは脆弱性 [CVE-2018-6574] の修正が行われている。
また 1.10 のリリース候補版も RC2 がリリースされている。 

- [cmd/go: arbitrary code execution during “go get” · Issue #23672 · golang/go · GitHub](https://github.com/golang/go/issues/23672)

## 脆弱性の内容

`#cgo` ディレクティブを含む [Go 言語]コードをビルドする際に任意の処理を実行される可能性がある。

{{< fig-quote title="arbitrary code execution during “go get”" link="(https://github.com/golang/go/issues/23672" lang="en" >}}
<q>When cgo is enabled, the build step during “go get” invokes the host C compiler, gcc or clang, adding compiler flags specified in the Go source files. Both gcc and clang support a plugin mechanism in which a shared-library plugin is loaded into the compiler, as directed by compiler flags. This means that a Go package repository can contain an attack.so file along with a Go source file that says (for example) <code>// #cgo CFLAGS: -fplugin=attack.so</code>, causing the attack plugin to be loaded into the host C compiler during the build. Gcc and clang plugins are completely unrestricted in their access to the host system.</q>
{{< /fig-quote >}}

このため `#cgo` ディレクティブでコンパイラプラグインに関する `-fplugin` 等のオプションを無効にしたようだ。
このことによる副作用を軽減するため，新たに2つの環境変数 `CGO_CFLAGS_ALLOW` および `CGO_CFLAGS_DISALLOW` が用意された。

{{< fig-quote title="arbitrary code execution during “go get”" link="(https://github.com/golang/go/issues/23672" lang="en" >}}
<q>The fix changes “go build” (used during “go get” and “go install”) to limit the flags that can appear in Go source file <code>#cgo</code> directives to a list of allowed compiler flags; <code>-fplugin=</code> and other variants are not on the allowed list. The same restrictions are applied to compiler flags obtained from pkg-config. Flags obtained from the environment variables $CGO_CFLAGS and so on are not restricted, since those variables can only be set by the user running the build. To change the set of allowed compiler flags, new environment variables $CGO_CFLAGS_ALLOW and $CGO_CFLAGS_DISALLOW can set to regular expressions matching additional allowed and disallowed flags. Run “go doc cgo” for details.</q>
{{< /fig-quote >}}


## 影響度（CVSS）

- [CVE-2018-6574 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-6574)

**CVSSv3 基本評価値 8.3 (`CVSS:3.0/AV:N/AC:H/PR:N/UI:R/S:C/C:H/I:H/A:H`)**

|                            基本評価基準 | 評価値            |
| ---------------------------------------:|:----------------- |
|                        攻撃元区分（AV） | ネットワーク（N） |
|                  攻撃条件の複雑さ（AC） | 高（H）           |
|                  必要な特権レベル（PR） | 不要（N）         |
|                  ユーザ関与レベル（UI） | 要（R）           |
|                           スコープ（S） | 変更あり（C）     |
| 情報漏えいの可能性（機密性への影響, C） | 高（H）           |
| 情報改ざんの可能性（完全性への影響, I） | 高（H）           |
|   業務停止の可能性（可用性への影響, A） | 高（H）           |

CVSS については[解説ページ]({{< ref "/remark/2015/cvss-v3-metrics-in-jvn.md" >}})を参照のこと。

## 影響を受ける製品

以下のバージョンが影響を受ける。

- 1.8.6 およびそれ以前
- 1.9.3 およびそれ以前
- 1.10rc1

## 対策・回避策

以下のバージョンにアップデートすること。

- 1.8.7 およびそれ以降
- 1.9.4 およびそれ以降
- 1.10rc2 およびそれ以降

影響度が高いので早めにアップデートを行うこと。

## ブックマーク

- [cgo - The Go Programming Language](https://golang.org/cmd/cgo/)
- [コンパイラプラグイン](http://www.kotha.net/ghcguide_ja/7.6.2/compiler-plugins.html)

[Go 言語]: https://golang.org/ "The Go Programming Language"
[CVE-2018-6574]: https://cve.mitre.org/cgi-bin/cvename.cgi?name=2018-6574
