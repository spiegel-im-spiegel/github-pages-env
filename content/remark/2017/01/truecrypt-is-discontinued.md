+++
date = "2017-01-10T20:21:52+09:00"
update = "2017-10-11T19:18:24+09:00"
title = "TrueCrypt の代替え製品"
draft = false
tags = ["security", "cryptography", "tools", "truecrypt"]
description = "現在メンテナンスが止まっている TrueCrypt の代替え製品を提案する。"

[author]
  github = "spiegel-im-spiegel"
  license = "by-sa"
  flattr = ""
  facebook = "spiegel.im.spiegel"
  instagram = "spiegel_2007"
  tumblr = ""
  flickr = "spiegel"
  twitter = "spiegel_2007"
  url = "https://baldanders.info/spiegel/profile/"
  name = "Spiegel"
  linkedin = "spiegelimspiegel"
  avatar = "/images/avatar.jpg"
+++

電子メールで教えていただいた。

- [TrueCrypt is discontinued, try these free alternative disk encryption tools](https://www.comparitech.com/blog/information-security/truecrypt-is-discoutinued-try-these-free-alternatives/)

メールをくれた方は，どうやら私の[古い日記の記述](https://baldanders.info/spiegel/log/200506.html#d07_t2)を見てアドバイスを下さったようだ。
感謝。

TrueCrypt はフリーのストレージ暗号化ツールである。

ストレージを丸ごと暗号化して仮想ドライブ等にできるのが特徴で， PC ドライブや USB メモリの暗号化などで重宝していた。
この TrueCrypt の開発が突然停止したのが2014年5月である。

- [暗号化ツールTrueCryptに突然の終了宣言、「セキュリティ問題」を警告 - ITmedia エンタープライズ](http://www.itmedia.co.jp/enterprise/articles/1405/30/news041.html)

当時は様々な憶測が流れたが，単純に開発者が投げ出してしまったようだ。
その後，暫定的に有志による [truecrypt.ch](https://truecrypt.ch/ "TCnext - Site dedicated to the development of the next truecrypt") が立ち上がりバージョン 7.1 をベースにメンテナンスが行われるようになった。

- [開発が停止したと思われていた暗号化ソフトTrueCryptが復活に向けて動きだす - GIGAZINE](http://gigazine.net/news/20140603-truecrypt-not-die/)

ただダウンロードページを見ると当時の 7.1a （7.1 の改良版）のままで

{{< fig-quote title="Downloads - TCnext" link="https://truecrypt.ch/downloads/" lang="en" >}}
<q>Please note: currently unmaintained, which might have security implications</q>
{{< /fig-quote >}}

と記述のあることから既に TrueCrypt 自体のメンテナンスは終了しているということなのだろう。

“[TrueCrypt is discontinued]” ではいくつかの代替え製品を提案している。
以下に挙げておく。

- [**VeraCrypt**](https://veracrypt.codeplex.com/) is open-source and code audited, improves on TrueCrypt, works on Mac and PC, and allows creation of encrypted containers
- [**Bitlocker**](https://technet.microsoft.com/en-us/library/cc732774.aspx) is built into Windows, is not open-source, only encrypts full disks, and has no plausible deniability mechanism
- [**DiskCryptor**](https://diskcryptor.net/wiki/Main_Page) is a Windows-only tool, is open source but not audited, allows the bootloader to be installed on a USB or CD, and works faster than others
- [**Ciphershed**](https://www.ciphershed.org/) is another TrueCrypt fork, works with old TrueCrypt containers, is slow with updates, and works on Mac, PC, and Linux
- [**FileVault 2**](https://support.apple.com/en-us/HT204837) is built into Mac OSX Lion and later, only allows full disk encryption, and is not open source
- [**LUKS**](https://guardianproject.info/code/luks/) is an open-source option for Linux, supports multiple algorithms, but does not offer much support for non-Linux systems

説明によると [VeraCrypt] と [Ciphershed] が TrueCrypt からの fork のようだ。
両者とも外部からのセキュリティ・チェックを受けているようなので概ね安全と言っていいだろう。

暗号化ストレージが必要な場合は上記のツールも検討してみてはどうだろうか。
まぁ私は PC や USB メモリを持ち出すような状況はなくなったので，さして必要ではなくなったんだけど。

## ブックマーク

- [TrueCrypt 終了宣言？｜Der Spiegel im Spiegel｜note](https://note.mu/spiegel/n/na028940a6872)
- [TrueCrypt must not die｜Der Spiegel im Spiegel｜note](https://note.mu/spiegel/n/n655406068534)
- [CIA（および NSA）による iOS および Windows BitLocker への攻撃 – Medium](https://medium.com/@spiegel/-a58a0cae2464)
- [【レビュー】秘密のファイルを隠しておくための暗号化仮想ドライブを作成できる「VeraCrypt」 - 窓の杜](https://forest.watch.impress.co.jp/docs/review/1084645.html)

[TrueCrypt is discontinued]: https://www.comparitech.com/blog/information-security/truecrypt-is-discoutinued-try-these-free-alternatives/ "TrueCrypt is discontinued, try these free alternative disk encryption tools"
[VeraCrypt]: https://veracrypt.codeplex.com/
[Ciphershed]: https://www.ciphershed.org/ "CipherShed | Secure Encryption Software"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE"><img src="https://images-fe.ssl-images-amazon.com/images/I/51t6yHHVwEL._SL160_.jpg" width="113" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E6%9A%97%E5%8F%B7%E6%8A%80%E8%A1%93%E5%85%A5%E9%96%80-%E7%AC%AC3%E7%89%88-%E7%A7%98%E5%AF%86%E3%81%AE%E5%9B%BD%E3%81%AE%E3%82%A2%E3%83%AA%E3%82%B9-%E7%B5%90%E5%9F%8E-%E6%B5%A9-ebook/dp/B015643CPE?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=B015643CPE">暗号技術入門 第3版　秘密の国のアリス</a></dt>
	<dd>結城 浩</dd>
    <dd>SBクリエイティブ 2015-08-25 (Release 2015-09-17)</dd>
    <dd>eBooks Kindle版</dd>
    <dd>ASIN: B015643CPE</dd>
    <dd>評価<abbr class="rating fa-sm" title="5">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i></abbr></dd>
  </dl>
  <p class="description">SHA-3 や Bitcoin/Blockchain など新しい知見や技術要素を大幅追加。暗号技術を使うだけならこれ1冊でとりあえず無問題。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2015-09-20">2015-09-20</abbr> (powered by <a href="https://github.com/spiegel-im-spiegel/amazon-item" >amazon-item</a> v0.2.0)</p>
</div>
