+++
title = "Ubuntu で音楽 CD のリッピング"
date =  "2019-06-10T22:59:23+09:00"
description = "Windows のときは文字エンコーディングで苦労したからねぇ。 UNIX 系は楽でいいわ（笑）"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "music", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いやぁ，ひーさしぶりに音楽 CD 買ったですよ。
最近はストリーミング・サービスか MP3 で買うかだったもんねぇ。

ちうわけで [Ubuntu] でリッピング・ツールを探したら [Asunder] というのがいいらしい。
インストールは APT でできる。

```text
$ sudo apt install asunder lame
```

`lame` は MP3 エンコーディングを扱うのに必要らしい。

起動したらちゃんと日本語で表示できてた。
えらいえらい！

{{< fig-img src="./asunder.png" link="./asunder.png" width="817" >}}

CDDB のプロキシ・サーバに [`freedbtest.dyndns.org:80`](https://freedbtest.dyndns.org/ "freedb 日本語") を指定すれば日本の CD 情報もちゃんと取れるようだ。

Windows のときは文字エンコーディングで苦労したからねぇ。
UNIX 系は楽でいいわ（笑）

## ブックマーク

- [UbuntuやArch Linuxで音楽CDの取り込み Asunderの紹介!! - BobbyQuineのブログ(備忘録)](http://bobbyquine.hatenablog.com/entry/2018/01/16/005131)
- [Ubuntuで音楽CDのPCへの取り込み（リッピング） - adbird（広告鳥） 備忘録](https://adbird.hatenablog.com/entry/2020/12/26/192332)

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Asunder]: http://www.littlesvr.ca/asunder/

## 参考図書

{{% review-paapi "B07PBB8W1X" %}} <!-- アルティメット☆MAGIC -->
{{% review-paapi "B007EEGL58" %}} <!-- 魔法の天使 クリィミーマミ 公式トリビュート・アルバム -->
