+++
title = "Go 1.17.1 のリリース【セキュリティ・アップデート】"
date =  "2021-10-08T20:10:54+09:00"
description = "wasm_exec.js ファイルの置き換えが必要らしい。ご注意を。"
image = "/images/attention/go-logo_blue.png"
tags  = [ "programming", "language", "golang", "ubuntu", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/7efr4VBoZIw "[security] Go 1.17.2 and Go 1.16.9 pre-announcement")通り [Go] 1.17.1 がリリースされた。

- [[security] Go 1.17.2 and Go 1.16.9 are released](https://groups.google.com/g/golang-announce/c/AEBu9j7yj5A)

今回は1件の脆弱性修正を含んでいる。

{{< fig-quote type="markdown" title="Release History - The Go Programming Language" link="https://golang.org/doc/devel/release#go1.17.minor" lang="en" >}}
{{% quote %}}go1.17.2 (released 2021-10-07) includes a security fix to the linker and `misc/wasm` directory, as well as bug fixes to the compiler, the runtime, the go command, and to the `time` and `text/template` packages. See the [Go 1.17.2 milestone](https://github.com/golang/go/issues?q=milestone%3AGo1.17.2+label%3ACherryPickApproved) on our issue tracker for details{{% /quote %}}.
{{< /fig-quote >}}

## [CVE-2021-38297]

{{< fig-quote type="markdown" title="Go 1.17.1 and Go 1.16.8 are released" link="https://groups.google.com/g/golang-announce/c/dx9d7IOseHw" lang="en" >}}
When invoking functions from WASM modules, built using `GOARCH=wasm GOOS=js`, passing very large arguments can cause portions of the module to be overwritten with data from the arguments.

If using `wasm_exec.js` to execute WASM modules, users will need to replace their copy (as described in [https://golang.org/wiki/WebAssembly#getting-started](https://golang.org/wiki/WebAssembly#getting-started)) after rebuilding any modules.
{{< /fig-quote >}}

というわけで `wasm_exec.js` ファイルの置き換えが必要らしい。
1.17.1 の `wasm_exec.js` ファイルとで diff をとってみたら

```text
$ diff -u go1.17.1/misc/wasm/wasm_exec.js go1.17.2/misc/wasm/wasm_exec.js
--- go1.17.1/misc/wasm/wasm_exec.js    2021-09-10 00:41:20.000000000 +0900
+++ go1.17.2/misc/wasm/wasm_exec.js    2021-10-08 04:58:29.000000000 +0900
@@ -567,6 +567,13 @@
                 offset += 8;
             });
 
+            // The linker guarantees global data starts from at least wasmMinDataAddr.
+            // Keep in sync with cmd/link/internal/ld/data.go:wasmMinDataAddr.
+            const wasmMinDataAddr = 4096 + 4096;
+            if (offset >= wasmMinDataAddr) {
+                throw new Error("command line too long");
+            }
+
             this._inst.exports.run(argc, argv);
             if (this.exited) {
                 this._resolveExitPromise();
```

とチェック用の条件文が追加されているだけのようだ。

（以下未稿）

## 例によって...

[Ubuntu] の APT で管理している [Go] コンパイラは古いので，[ダウンロードページ](https://golang.org/dl/ "Downloads - The Go Programming Language")からバイナリ（[`go1.17.2.linux-amd64.tar.gz`](https://golang.org/dl/go1.17.2.linux-amd64.tar.gz)）を取ってきて手動でインストールすることを強く推奨する。
以下は手動での作業例。

```text
$ cd /usr/local/src
$ sudo curl -L "https://golang.org/dl/go1.17.2.linux-amd64.tar.gz" -O
$ cd ..
$ sudo unlink go # 以前の Go が入っている場合
$ sudo tar xvf src/go1.17.2.linux-amd64.tar.gz
$ sudo mv go go1.17.2
$ sudo ln -s go1.17.2 go
$ go version # /usr/local/go/bin にパスが通っている場合
o version go1.17.2 linux/amd64
```

アップデートは計画的に。

[Go]: https://golang.org/ "The Go Programming Language"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[CVE-2021-38297]: https://nvd.nist.gov/vuln/detail/CVE-2021-38297

## 参考図書

{{% review-paapi "B099928SJD" %}} <!-- プログラミング言語Go -->
