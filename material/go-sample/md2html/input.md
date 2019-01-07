# Go 1.8.7 および 1.9.4 がリリース

[Go 言語]コンパイラの 1.8.7 および 1.9.4 がリリースされている。
このバージョンでは脆弱性 [CVE-2018-6574] の修正が行われている。

## 脆弱性の内容

`#cgo` ディレクティブを含む [Go 言語]コードをビルドする際に任意の処理を実行される可能性がある。

> When cgo is enabled, the build step during “go get” invokes the host C compiler, gcc or clang, adding compiler flags specified in the Go source files.
> Both gcc and clang support a plugin mechanism in which a shared-library plugin is loaded into the compiler, as directed by compiler flags.
> This means that a Go package repository can contain an attack.so file along with a Go source file that says (for example) `// #cgo CFLAGS: -fplugin=attack.so`, causing the attack plugin to be loaded into the host C compiler during the build.
> Gcc and clang plugins are completely unrestricted in their access to the host system.

## 影響度（CVSS）

「[CVE-2018-6574 - Red Hat Customer Portal](https://access.redhat.com/security/cve/cve-2018-6574)」より。

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

CVSS については[解説ページ](https://text.baldanders.info/remark/2015/cvss-v3-metrics-in-jvn/ "JVN が CVSSv3 による脆弱性評価を開始 — しっぽのさきっちょ | text.Baldanders.info")を参照のこと。

[Go 言語]: https://golang.org/ "The Go Programming Language"
[CVE-2018-6574]: https://cve.mitre.org/cgi-bin/cvename.cgi?name=2018-6574
