+++
title = "量子コンピュータで256ビット楕円曲線暗号は敗れるか"
date =  "2022-02-19T13:39:48+09:00"
description = "量子コンピュータってどんどん進歩してるんだねぇ。"
image = "/images/attention/kitten.jpg"
tags = [
  "security",
  "cryptography",
  "ecc",
]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

Bruce Schneier 先生のブログ記事より。

- [The impact of hardware specifications on reaching quantum advantage in the fault tolerant regime: AVS Quantum Science: Vol 4, No 1](https://avs.scitation.org/doi/10.1116/5.0073075)
- [Breaking 256-bit Elliptic Curve Encryption with a Quantum Computer - Schneier on Security](https://www.schneier.com/blog/archives/2022/02/breaking-245-bit-elliptic-curve-encryption-with-a-quantum-computer.html)

この論文で256ビット楕円曲線暗号を破るのに必要な量子コンピュータのサイズを計算したらしい。
概要（ABSTRACT）には

{{< fig-quote type="markdown" title="The impact of hardware specifications on reaching quantum advantage in the fault tolerant regime" link="https://avs.scitation.org/doi/10.1116/5.0073075" lang="en" >}}
{{% quote %}}Finally, we calculate the number of physical qubits required to break the 256-bit elliptic curve encryption of keys in the Bitcoin network within the small available time frame in which it would actually pose a threat to do so. It would require $317 \times 10^6$ physical qubits to break the encryption within one hour using the surface code, a code cycle time of $1\ {\mu}\mathrm{s}$, a reaction time of $10\ {\mu}\mathrm{s}$, and a physical gate error of $10^{−3}$. To instead break the encryption within one day, it would require $13 \times 10^6$ physical qubits{{% /quote %}}.
{{< /fig-quote >}}

とあるので，そういうことなのだろう。
ちなみに2021年末の時点で [IBM が開発した量子コンピュータ](https://www.newscientist.com/article/2297583-ibm-creates-largest-ever-superconducting-quantum-computer/ "IBM creates largest ever superconducting quantum computer | New Scientist")のサイズが127物理キュービットで最大らしいので，まぁ当分は大丈夫というところなのだろう。

日本でも既存暗号に対する量子コンピュータ耐性について継続的に調査と評価が行われていて，毎年の [CRYPTREC] Report に載っている。

2020年2月に公開された「[現在の量子コンピュータによる暗号技術の安全性への影響](https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html)」では

{{< fig-quote type="markdown" title="現在の量子コンピュータによる暗号技術の安全性への影響" link="https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html" >}}
{{% quote %}}例えば、量子コンピュータを用いて2048ビットRSA合成数の素因数分解を行う場合には、量子誤りが一切ないという理想的な環境下でも、4098量子ビットが必要であり、1012～1013回のゲート演算が必要であると見積もられています。また、量子誤りがあるという現実的な環境下では、2000万量子ビットが必要であるという見積もりもあります{{% /quote %}}。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="現在の量子コンピュータによる暗号技術の安全性への影響" link="https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html" >}}
{{% quote %}}量子コンピュータの性能を測る上での指標（量子ビット数、量子誤りの大きさ、演算可能回数など）や、量子コンピュータの開発状況もあわせて考慮にいれると、近い将来に、2048ビットの素因数分解や256ビットの楕円曲線上の離散対数問題が解かれる可能性は低いと考えます{{% /quote %}}。
{{< /fig-quote >}}

とあるが，当時の試算は

{{< fig-quote type="markdown" title="現在の量子コンピュータによる暗号技術の安全性への影響" link="https://www.cryptrec.go.jp/topics/cryptrec-er-0001-2019.html" >}}
{{% quote %}}論文[^a1] で使用されている量子コンピュータは53量子ビットであり、計算は合計1543回のゲート演算で構成されています。このとき、1回当たりの計算時間は、1マイクロ秒程度であると見積もられています。なお、ターゲットとする問題の性質上、量子誤り訂正は組み込まれていません{{% /quote %}}。

[^a1]: [Quantum supremacy using a programmable superconducting processor](https://www.nature.com/articles/s41586-019-1666-5)
{{< /fig-quote >}}

となっている。
量子コンピュータってどんどん進歩してるんだねぇ。
今後も注意して見ていく必要があるかな。

## ブックマーク

- [MIT Tech Review: ポスト量子暗号の実装安全性を向上、東北大学とNTTが考案](https://www.technologyreview.jp/n/2022/02/03/268972)

- [Edwards-curve Digital Signature Algorithm]({{< ref "/remark/2020/06/eddsa.md" >}})

[CRYPTREC]: https://www.cryptrec.go.jp/

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
