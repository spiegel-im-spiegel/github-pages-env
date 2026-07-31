+++
title = "Scoop 版 GnuPG 2.3 系の gpg-agent が動かない？"
date =  "2021-12-07T21:35:57+09:00"
description = "こういう調査を Azure Virtual Desktop でしたかったんだよね。"
image = "/images/attention/kitten.jpg"
tags = [ "gnupg", "windows", "azure", "install", "tools", "scoop" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [ようやく Azure Virtual Desktop を導入できた]({{< ref "/remark/2021/12/azure-virtual-desktop.md" >}})
- [Azure Virtual Desktop で遊ぶ]({{< ref "/remark/2021/12/azure-virtual-desktop-2.md" >}})
- [Scoop 版 GnuPG 2.3 系の gpg-agent が動かない？]({{< ref "/remark/2021/12/gpg-agent-for-windows.md" >}}) ← イマココ

今日も [Azure Virtual Desktop] で遊んでいるのだが，どうやら [Scoop] でインストールした 2.3 系 [GnuPG] で gpg-agent が上手く起動しないみたい。
厳密には [GnuPG] v2.3.2 以降。

実は仕事用の Windows マシンでは随分前から気がついていたが，忙しくて碌に検証もできなかったのだ[^gpg1]。
今回の真っさら Windows で試してみて「やっぱ動かんぢゃん！」ってなった。

[^gpg1]: 2.3系では特にセキュリティ・アップデートは発生していないので「後回しでいっか」と思っていた。

詳しく言うと

```text
$ gpg-connect-agent /bye
gpg-connect-agent: gpg-agentが動いていません - 開始します'C:\\Users\\username\\scoop\\apps\\gnupg\\current\\bin\\gpg-agent.exe'
gpg-connect-agent: agent の起動のため、5秒待ちます...
gpg-connect-agent: agent の起動のため、4秒待ちます...
gpg-connect-agent: agent の起動のため、3秒待ちます...
gpg-connect-agent: agent の起動のため、2秒待ちます...
gpg-connect-agent: agent の起動のため、1秒待ちます...
gpg-connect-agent: can't connect to the gpg-agent: IPC connect呼び出しに失敗しました
gpg-connect-agent: 標準オプションを送信エラー: エージェントが動いていません
```

という感じに gpg-agent を起動しようとするも，5秒でタイムアウトになって失敗する。
v2.3.1 に戻すと動くので v2.3.2 で何かあったんだろう。

[v2.3.2 の変更点]({{< ref "/release/2021/08/gnupg-2_3_2-is-released.md" >}})を眺めると

{{< fig-quote type="markdown" title="GnuPG 2.3.2 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2021q3/000462.html" lang="en" >}}
* Under Windows use `LOCAL_APPDATA` for the socket directory.  [#5537]
{{< /fig-quote >}}

とある。
どうもこれっぽい。
でも `gpgconf` で見ると

```text
$ gpgconf --list-dirs | grep socket
socketdir:C%3a\Users\username\scoop\apps\gnupg\current\gnupg
dirmngr-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.dirmngr
keyboxd-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.keyboxd
agent-ssh-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.gpg-agent.ssh
agent-extra-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.gpg-agent.extra
agent-browser-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.gpg-agent.browser
agent-socket:C%3a\Users\username\scoop\apps\gnupg\current\gnupg\S.gpg-agent
```

てな感じで，ちゃんと [Scoop] 配下で設定されてるんだけどねぇ。

## [winget] 版 [GnuPG] を導入する

まぁ，愚痴ってもしょうがないので，とりあえず [Scoop] 版は諦めることにしよう。
幸いなことに [winget] でも [GnuPG] をインストール可能で，こちらは標準のインストーラが動くようだ。

```text
$ winget search gnupg
名前              ID            バージョン ソース
--------------------------------------------------
GNU Privacy Guard GnuPG.GnuPG   2.3.3      winget
Gpg4win           GnuPG.Gpg4win 3.1.16     winget
```

というわけでインストールしてしまおう。

```text
$ winget install GnuPG.GnuPG
見つかりました GNU Privacy Guard [GnuPG.GnuPG] バージョン 2.3.3
このアプリケーションは所有者からライセンス供与されます。
Microsoft はサードパーティのパッケージに対して責任を負わず、ライセンスも付与しません。
Downloading https://gnupg.org/ftp/gcrypt/binary/gnupg-w32-2.3.3_20211012.exe
  ██████████████████████████████  4.59 MB / 4.59 MB
インストーラーハッシュが正常に検証されました
パッケージのインストールを開始しています...
インストールが完了しました

$ gpg --version
gpg (GnuPG) 2.3.3
libgcrypt 1.9.4
Copyright (C) 2021 g10 Code GmbH
License GNU GPL-3.0-or-later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: C:\Users\username\AppData\Roaming\gnupg
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256,
      TWOFISH, CAMELLIA128, CAMELLIA192, CAMELLIA256
AEAD: EAX, OCB
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2

$ which gpg.exe
C:\Program Files (x86)\gnupg\bin\gpg.exe

$ gpgconf --list-dirs | grep socket
socketdir:C%3a\Users\username\AppData\Local\gnupg
dirmngr-socket:C%3a\Users\username\AppData\Local\gnupg\S.dirmngr
keyboxd-socket:C%3a\Users\username\AppData\Local\gnupg\S.keyboxd
agent-ssh-socket:C%3a\Users\username\AppData\Local\gnupg\S.gpg-agent.ssh
agent-extra-socket:C%3a\Users\username\AppData\Local\gnupg\S.gpg-agent.extra
agent-browser-socket:C%3a\Users\username\AppData\Local\gnupg\S.gpg-agent.browser
agent-socket:C%3a\Users\username\AppData\Local\gnupg\S.gpg-agent

$ gpg-connect-agent /bye
gpg-connect-agent: gpg-agentが動いていません - 開始します'C:\\Program Files (x86)\\gnupg\\bin\\gpg-agent.exe'
gpg-connect-agent: agent の起動のため、5秒待ちます...
gpg-connect-agent: agentへの接続が確立しました
```

よーし，うむうむ，よーし。

こういう調査を [Azure Virtual Desktop] でしたかったんだよね。
最悪コワしても仮想マシンを作り直せばいいんだし。

さて [GnuPG] の導入で思いのほか手こずったので git 環境の整備は次回だな（笑）

## ブックマーク

- [GnuPG の HOME はどこにある？]({{< ref "/openpgp/gnupg-home-in-windows.md" >}})

[Azure Virtual Desktop]: https://docs.microsoft.com/ja-jp/azure/virtual-desktop/ "Azure Virtual Desktop のドキュメント | Microsoft Docs"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[winget]: https://github.com/microsoft/winget-cli "microsoft/winget-cli: Windows Package Manager CLI (aka winget)"
[Scoop]: https://scoop.sh/ "Scoop"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
