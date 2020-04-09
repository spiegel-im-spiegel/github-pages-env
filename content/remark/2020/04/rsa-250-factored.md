+++
title = "RSA-250 解読完了！"
date =  "2020-04-09T10:45:03+09:00"
description = "実際には何万ものマシンを使って数ヶ月で解読したらしい。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "cryptography", "rsa" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

2月末の話で申し訳ないが（Bruce Schneier 先生の記事でさっき知ったのだ）， RSA-250 鍵が因数分解され解読完了したらしい。

- [[Cado-nfs-discuss] Factorization of RSA-250](https://lists.gforge.inria.fr/pipermail/cado-nfs-discuss/2020-February/001166.html)
- [RSA-250 Factored - Schneier on Security](https://www.schneier.com/blog/archives/2020/04/rsa-250_factore.html)

{{< fig-quote type="markdown" title="Factorization of RSA-250" link="https://lists.gforge.inria.fr/pipermail/cado-nfs-discuss/2020-February/001166.html" lang="en" >}}
The total computation time was roughly 2700 core-years, using Intel Xeon
Gold 6130 CPUs as a reference (2.1GHz):

```
RSA-250 sieving:  2450 physical core-years
RSA-250 matrix:    250 physical core-years
```
{{< /fig-quote >}}

実際には何万ものマシンを使って数ヶ月で解読したらしい。

{{< fig-quote type="markdown" title="RSA-250 Factored" link="https://www.schneier.com/blog/archives/2020/04/rsa-250_factore.html" lang="en" >}}
{{< quote >}}The computation involved tens of thousands of machines worldwide, and was completed in a few months{{< /quote >}}.
{{< /fig-quote >}}

[RSA-240 が解けた]({{< ref "/remark/2019/12/rsa-240-factored.md" >}})のって，つい昨年末なんだけどねぇ（笑） RSA は近い将来，量子コンピュータの一般化を待たずにお払い箱になるんだろうね。

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
