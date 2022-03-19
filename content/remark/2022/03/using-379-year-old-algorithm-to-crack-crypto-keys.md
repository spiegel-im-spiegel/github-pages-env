+++
title = "379年前のアルゴリズムを使って RSA 暗号鍵を破った話"
date =  "2022-03-19T12:20:03+09:00"
description = "例によって Bruce Schneier さんの記事より"
image = "/images/attention/kitten.jpg"
tags = [ "security", "cryptography", "rsa", "vulnerability" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

例によって [Bruce Schneier さんの記事](https://www.schneier.com/blog/archives/2022/03/breaking-rsa-through-insufficiently-random-primes.html "Breaking RSA through Insufficiently Random Primes - Schneier on Security")経由：

- [Fermat Attack on RSA](https://fermatattack.secvuln.info/)
- [Researcher uses 379-year-old algorithm to crack crypto keys found in the wild | Ars Technica](https://arstechnica.com/information-technology/2022/03/researcher-uses-600-year-old-algorithm-to-crack-crypto-keys-found-in-the-wild/)

RSA 公開鍵暗号の仕組みについては結城浩さんの『[暗号技術入門](https://www.amazon.co.jp/dp/B015643CPE?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "暗号技術入門 第3版　秘密の国のアリス")』の第5章に分かりやすく解説されているので，まずはそちらを参照のこと。

今回のポイントとなる部分だけ紹介すると RSA 公開鍵暗号の鍵ペアを生成する際には最初に2つの大きな素数 $(p,q)$ を用意する。
$(p,q)$ を掛け合わせた $N=p \times q$ は公開鍵にも秘密鍵にも使われる値だが「2つの大きな素数を合成した値を（元の素数を知らずに）因数分解するのは難しい」ため $N$ から秘密鍵を推測するのは難しいとされている。
当然ながら素数 $(p,q)$ の組み合わせは第三者に知られないよう管理する必要がある。

{{< fig-quote type="markdown" title="Researcher uses 379-year-old algorithm to crack crypto keys found in the wild" link="https://arstechnica.com/information-technology/2022/03/researcher-uses-600-year-old-algorithm-to-crack-crypto-keys-found-in-the-wild/" lang="en" >}}
{{% quote %}}The security of RSA keys depends on the difficulty of factoring a key's large composite number (usually denoted as N) to derive its two factors (usually denoted as P and Q). When P and Q are known publicly, the key they make up is broken, meaning anyone can decrypt data protected by the key or use the key to authenticate messages{{% /quote %}}.
{{< /fig-quote >}}

ただし，素数 $(p,q)$ が互いに近い値の場合は容易に因数分解できることも大昔から知られている。

{{< fig-quote type="markdown" title="Researcher uses 379-year-old algorithm to crack crypto keys found in the wild" link="https://arstechnica.com/information-technology/2022/03/researcher-uses-600-year-old-algorithm-to-crack-crypto-keys-found-in-the-wild/" lang="en" >}}
Cryptographers have long known that RSA keys that are generated with primes that are too close together can be trivially broken with [Fermat's factorization method](https://en.wikipedia.org/wiki/Fermat%27s_factorization_method). French mathematician Pierre de Fermat [first described this method in 1643](https://madhavamathcompetition.com/tag/fermats-factorization-method/).

Fermat's algorithm was based on the fact that any odd number can be expressed as the difference between two squares. When the factors are near the root of the number, they can be calculated easily and quickly. The method isn't feasible when factors are truly random and hence far apart.
{{< /fig-quote >}}

で，実際に一部の暗号製品で「鍵サイズは大きいけど容易に破られる暗号鍵」を生成してしまうものがあったそうで，これは [CVE-2022-26320] として詳細が公開されている。

{{< fig-quote type="markdown" title="CVE-2022-26320" link="https://nvd.nist.gov/vuln/detail/CVE-2022-26320" lang="en" >}}
{{% quote %}}The Rambus SafeZone Basic Crypto Module before 10.4.0, as used in certain Fujifilm (formerly Fuji Xerox) devices before 2022-03-01, Canon imagePROGRAF and imageRUNNER devices through 2022-03-14, and potentially many other devices, generates RSA keys that can be broken with Fermat's factorization method. This allows efficient calculation of private RSA keys from the public key of a TLS certificate{{% /quote %}}.
{{< /fig-quote >}}

この脆弱性を報告した [Hanno Böck](https://hboeck.de/) さんは，更に SKS PGP 鍵サーバにも今回のような脆弱な RSA 公開鍵があったと言っているらしい（実際に運用している鍵ではなさそうだが）。

{{< fig-quote type="markdown" title="Researcher uses 379-year-old algorithm to crack crypto keys found in the wild" link="https://arstechnica.com/information-technology/2022/03/researcher-uses-600-year-old-algorithm-to-crack-crypto-keys-found-in-the-wild/" lang="en" >}}
Böck also found four vulnerable PGP keys, typically used to encrypt email, on SKS PGP key servers. A user ID tied to the keys implied they were created for testing, so he doesn't believe they're in active use.

Böck said he believes all the keys he found were generated using software or methods not connected to the SafeZone library. If true, other software that generates keys might be easily broken using the Fermat algorithm. It's plausible that the keys were generated manually, "possibly by people aware of this attack creating test data," Böck said.
{{< /fig-quote >}}

今回の脆弱性を公開した [Hanno Böck](https://hboeck.de/) さんは自身の[記事](https://fermatattack.secvuln.info/ "Fermat Attack on RSA")の中で，破られやすい素数の組み合わせとして

{{< fig-quote type="markdown" title="Fermat Attack on RSA" link="https://fermatattack.secvuln.info/" lang="en" >}}
**How "close" do primes need to be in order to be vulnerable?**

With common RSA key sizes (2048 bit) in our tests the Fermat algorithm with 100 rounds reliably factors numbers where p and q differ up to 2^517. In other words it can be said that primes that only differ within the lower 64 bytes (or around half their size) will be vulnerable.

Up to 2^514 it almost always finds the factorization in the first round of the algorithm. It could be argued that the 100 rounds is therefore excessive, however the algorithm is so fast that it practically does not matter much.
{{< /fig-quote >}}

と見積もっている。
ちなみに SSH で生成する RSA 鍵については

{{< fig-quote type="markdown" title="Fermat Attack on RSA" link="https://fermatattack.secvuln.info/" lang="en" >}}
**Is SSH affected?**

There are probably no vulnerable SSH implementations creating such keys, though I obviously cannot proove that.

I checked multiple large collections of both SSH host and user keys. I did not find any vulnerable keys.
{{< /fig-quote >}}

なんだそうな。
よかったね。

なお，記事の最後では

{{< fig-quote type="markdown" title="Fermat Attack on RSA" link="https://fermatattack.secvuln.info/" lang="en" >}}
**What do you recommend?**

If you are running one of the affected devices you should make sure that you update the devices and regenerate the keys.

If you are in a position where external users will supply public RSA keys to you then you might want to implement checks for this vulnerability. A typical case for this are certificate authorities. I shared the exploit code with certificate authorities and are aware that some have implemented checks in their certificate issuance process to avoid accepting keys vulnerable to this attack.

You can find Let's Encrypt's implementation of the check in their Boulder software [in this pull request](https://github.com/letsencrypt/boulder/pull/5853).
{{< /fig-quote >}}

と締めている。
今回のケースは暗号製品や CA など暗号鍵を提供・運用する側の問題でユーザ側でできることは殆どないだろうが，とりあえず怪しげな暗号製品は使うなっちうことやね（笑）

[CVE-2022-26320]: https://nvd.nist.gov/vuln/detail/CVE-2022-26320

## 参考図書

{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
<!-- eof -->
