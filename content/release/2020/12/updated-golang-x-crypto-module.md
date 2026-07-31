+++
title = "golang.org/x/crypto/ssh パッケージのセキュリティ・アップデート"
date =  "2020-12-17T19:44:22+09:00"
description = "ヌルポか。似たような話を最近よく聞くな（笑）"
image = "/images/attention/go-logo_blue.png"
tags  = [ "golang", "package", "security", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[予告](https://groups.google.com/g/golang-announce/c/CqSxrm7Mpr0 "[security] golang.org/x/crypto/ssh fix pre-announcement")通り， [golang.org/x/crypto/ssh][ssh] パッケージのセキュリティ・アップデートが行われた。

{{< fig-quote type="markdown" title="[security] Vulnerability in golang.org/x/crypto/ssh" link="https://groups.google.com/g/golang-announce/c/ouZIlBimOsE" lang="en" >}}
Version v0.0.0-20201216223049-8b5274cf687f of [golang.org/x/crypto](http://golang.org/x/crypto) fixes a vulnerability in the [golang.org/x/crypto/ssh](http://golang.org/x/crypto/ssh) package which allowed clients to cause a panic in SSH servers.

An attacker can craft an authentication request message for the “`gssapi-with-mic`” method which will cause `NewServerConn` to panic via a nil pointer dereference if `ServerConfig.GSSAPIWithMICConfig` is nil.
{{< /fig-quote >}}

ヌルポか。
似たような話を最近よく聞くな（笑）

ちうわけで [golang.org/x/crypto/ssh][ssh] パッケージを使っている場合は，最新版を get し直して再ビルドしませう。
その際， `go.mod` ファイル内に記述されている [golang.org/x/crypto][crypto] モジュールのバージョンに注意。
とりあえず

```text
require (
    golang.org/x/crypto latest
)
```

とかにして，テストかビルドをし直せばいいかな？

余談だが [Go] はインスタンス（への参照）が nil でもメソッド自体は呼び出せる。
まぁメソッド内でインスタンスの要素にアクセスしようとすればヌルポで panic になるんだけど。

そんで，実は最近また Java コードを弄ってるんだけど， Java では null 参照インスタンスのメソッドを呼び出した時点でヌルポで落ちることをすっかり忘れていて，エラい目に遭ったですよ（笑）

というわけで，アップデートは計画的に。

[Go]: https://go.dev/
[crypto]: https://pkg.go.dev/golang.org/x/crypto "crypto · pkg.go.dev"
[ssh]: https://pkg.go.dev/golang.org/x/crypto/ssh "ssh · pkg.go.dev"
<!-- eof -->
