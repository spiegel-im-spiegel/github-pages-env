+++
date = "2015-12-06T00:22:47+09:00"
update = "2016-10-18T21:25:39+09:00"
description = "「はやぶさ2」地球 Swing-by / GnuPG 2.1.10 released / Go 1.5.2 Released / 11月の Flattr / 機械創作でクリエイターは失業するか / 新刊小説の滅亡"
draft = false
tags = ["astronomy", "jaxa", "hayabusa2", "swing-by", "security", "cryptography", "openpgp", "gnupg", "tools", "golang", "flattr", "artificial-intelligence", "book"]
title = "週末スペシャル： 「はやぶさ2」地球 Swing-by"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

1. [「はやぶさ2」地球 Swing-by]({{< relref "#hayabusa" >}})
1. [GnuPG 2.1.10 released]({{< relref "#gnupg" >}})
1. [Go 1.5.2 Released]({{< relref "#golang" >}})
1. [11月の Flattr]({{< relref "#flattr" >}})
1. [機械創作でクリエイターは失業するか]({{< relref "#machine" >}})
1. [新刊小説の滅亡]({{< relref "#book" >}})

## 「はやぶさ2」地球 Swing-by{#hayabusa}

12月3日に「はやぶさ2」の地球 swing-by が行われた。
Swing-by 自体はうまくいったようだが，今のところ所定の軌道に乗ったかどうか確認中らしい。

- [「はやぶさ２」の地球スイングバイ軌道が確定しました](http://www.hayabusa2.jaxa.jp/topics/20151202/)
- [「はやぶさ2」打ち上げから1年、本日軌道変更(地球スイングバイ)を実施 | 月探査情報ステーションブログ](http://moonstation.jp/ja/blog/archives/1931)
- [JAXA | 小惑星探査機「はやぶさ2」の地球スイングバイ実施について](http://www.jaxa.jp/press/2015/12/20151203_hayabusa2_j.html)
- [はやぶさ2、軌道変更(地球スイングバイ)を実施、成否は数日後に判明 | 月探査情報ステーションブログ](http://moonstation.jp/ja/blog/archives/1938)
- [「はやぶさ２」スイングバイ直前に撮影された地球の画像 | JAXA はやぶさ２プロジェクト](http://www.hayabusa2.jaxa.jp/topics/20151203/)

地球 Swing-by については以下の解説記事が参考になる。

- [探査機をスマートに飛ばすテクニック「スイングバイ」: 小惑星探査機「はやぶさ2」: NECの宇宙開発利用への取り組み「宙への挑戦」 | NEC](http://jpn.nec.com/ad/cosmos/hayabusa2/swingby/)

{{< fig-quote title="探査機をスマートに飛ばすテクニック「スイングバイ」" link="http://jpn.nec.com/ad/cosmos/hayabusa2/swingby/" >}}
<q>もし地球が止まっているなら、近づく時も離れる時も、探査機の速度は変わりません。ですが地球は、秒速30kmという猛スピードで太陽の周りを公転しており、とても大きな運動エネルギーを持っています。「はやぶさ２」は地球の後方を通過することで、そのエネルギーのほんの一部をもらって、前方へ加速していきます。今回のスイングバイで「はやぶさ2」は毎秒1.6km、時速にすると何と約6000kmも速くなります。</q>
{{< /fig-quote >}}

さて，週明けは「あかつき」だよ。
うまくいくといいね。

- [Back from the Brink: Akatsuki Returns to Venus | The Planetary Society](http://www.planetary.org/blogs/guest-blogs/2015/1204-akatsuki-returns-to-venus.html)

### 12月14日 追記

地球 swing-by 後の「はやぶさ2」の軌道計算が完了。
予定通り航行しているようだ。
おめでとう！

- [JAXA | 小惑星探査機「はやぶさ2」の地球スイングバイ実施結果について](http://www.jaxa.jp/press/2015/12/20151214_hayabusa2_j.html)
- [はやぶさ2、軌道変更(スイングバイ)成功を確認 | 月探査情報ステーションブログ](http://moonstation.jp/ja/blog/archives/1948)
- [「はやぶさ2」の軌道変更(スイングバイ)時に撮影された地球の写真 | 月探査情報ステーションブログ](http://moonstation.jp/ja/blog/archives/1951)

## GnuPG 2.1.10 released{#gnupg}

- [[Announce] GnuPG 2.1.10 released](https://lists.gnupg.org/pipermail/gnupg-announce/2015q4/000381.html)

今回はセキュリティ脆弱性関連のアップデートはないが，このバージョンで TOFU (Trust-On-First-Use) 信用モデルをサポートするらしい。
これって SSH のプロトコルで使われている方式かな[^a]。

- [Trust on first use - Wikipedia, the free encyclopedia](https://en.wikipedia.org/wiki/Trust_on_first_use)

あぁ，また勉強しないと。
日本語の記事が見当たらないんだが...

[^a]: `TOFU` でググってて気がついたんだけど， [GnuTLS](http://www.gnutls.org/) ってまだ生きてるんだね。使ってるところってどのくらいあるんだろう。

その他の変更点は以下のとおり。

* gpg: New trust models "tofu" and "tofu+pgp".
* gpg: New command --tofu-policy.  New options --tofu-default-policy and --tofu-db-format.
* gpg: New option --weak-digest to specify hash algorithms which should be considered weak.
* gpg: Allow the use of multiple --default-key options; take the last available key.
* gpg: New option --encrypt-to-default-key.
* gpg: New option --unwrap to only strip the encryption layer.
* gpg: New option --only-sign-text-ids to exclude photo IDs from key signing.
* gpg: Check for ambigious or non-matching key specification in the config file or given to --encrypt-to.
* gpg: Show the used card reader with --card-status.
* gpg: Print export statistics and an EXPORTED status line.
* gpg: Allow selecting subkeys by keyid in --edit-key.
* gpg: Allow updating the expiration time of multiple subkeys at once.
* dirmngr: New option --use-tor.  For full support this requires libassuan version 2.4.2 and a patched version of libadns (e.g. adns-1.4-g10-7 as used by the standard Windows installer).
* dirmngr: New option --nameserver to specify the nameserver used in Tor mode.
* dirmngr: Keyservers may again be specified by IP address.
* dirmngr: Fixed problems in resolving keyserver pools.
* dirmngr: Fixed handling of premature termination of TLS streams so that large numbers of keys can be refreshed via hkps.
* gpg: Fixed a regression in --locate-key [since 2.1.9].
* gpg: Fixed another bug for keyrings with legacy keys.
* gpgsm: Allow combinations of usage flags in --gen-key.
* Make tilde expansion work with most options.
* Many other cleanups and bug fixes.

```bash
C:> gpg --version
gpg (GnuPG) 2.1.10
libgcrypt 1.6.4
Copyright (C) 2015 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
Supported algorithms:
Pubkey: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
Cipher: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
        CAMELLIA128, CAMELLIA192, CAMELLIA256
Hash: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
Compression: Uncompressed, ZIP, ZLIB, BZIP2
```

## Go 1.5.2 Released{#golang}

今週はアップデートが多いなぁ。

- [Go 1.5.2 milestone](https://golang.org/doc/devel/release.html#go1.5.minor)

細かい不具合の修正がメインのようだ。
セキュリティ関連のバグはないかな。

## 11月の Flattr{#flattr}

11月は1件のみ。

- [SETI@home](https://flattr.com/thing/114786/timestamp/1443655861 "SETI@home - Flattr.com")

[Flattr] はまたサービスを[改装](http://blog.flattr.net/2015/11/flattr-developer-update-6-time-to-move/)してるみたいなのだが，今度はどうなるやら。

[Flattr]: https://flattr.com/ "Flattr - Social microdonations"

## 機械創作でクリエイターは失業するか{#machine}

- [人工知能と著作権　～機械創作の普及でクリエイターは失業するのか？～ -INTERNET Watch](http://internet.watch.impress.co.jp/docs/special/fukui/20151201_732993.html)

別に人工知能を使わなくても今あるもので「機械創作」なるものは可能だろう。

たとえば「人工無脳」を使えば無限に言葉を紡ぐことができる。
また，今は亡き Mag. さんが公開されていた「[Chaos von Eschenbach](http://magarchive.halfmoon.jp/vector/chaos/index.html)」を使えば無限のバリエーションで作曲できる[^b]。
あとは初音ミクにでも唄わせればいい。

[^b]: 当時はこのソフトにメチャメチャはまったものである。他にも自動作曲ソフトには様々なものがある。

ほら「機械創作」なんて簡単でしょ。

でも「人工無脳」や「自動作曲」が紡ぐ言葉や曲を以って「表現」しているとは誰も思わないだろう。
何故ならそこに知能（intelligence）は存在しない（しないと思っている）から。
それでもその知能なき創作に私たちは時に感動を見出す。
たとえば自然現象のような。

ということは「機械創作」は別の問題を示唆している。
つまり「機械が創作する」ことが問題なのではなく「創作」することに知能が必要なのかということである。
それはたとえば「知的財産権」の論理基盤を根底から覆す可能性を秘めている。

人工知能は今のところ「問いに答える（応える）」ことはできるが，自ら「問いを立てる」ことはできない[^c]。
私は，それこそが「機械」を知的存在と呼べない（今や最後の）理由だと思う。
IBM の [Watson](http://www.ibm.com/smarterplanet/jp/ja/ibmwatson/) がどんなに優秀でも「エキスパート・システム」のバリエーションに留まる限り知的存在にはならない。

[^c]: 推論の過程でいくつかの仮説をたてることはある。が，それも「反応」の一種に過ぎない。

現実のピノキオはどうやっても人間になれないのである。
もしピノキオが人間になれるなら，そこにいるのは人間ではなく，まったく新しい知的存在と言える。

じゃあ「知的労働者」は将来も安泰なのかといえば，全くそんなことはない。
むしろ人工知能は「知的労働」を侵食していくだろう。
何故なら「知的労働」を行うのが知的存在である必要はないのだから。

決められたコードに従って「創作」しているだけの職業クリエイターこそが人工知能に排除される対象である。

- [自動作曲と著作権](http://magarchive.halfmoon.jp/vector/chaos/copyright.html)
- [自作のビデオに音楽をつけたい人、作曲サービスJukedeckが著作権のない曲を一瞬で作ってくれる | TechCrunch Japan](https://jp.techcrunch.com/2015/12/08/20151207jukedeck/)

## 新刊小説の滅亡{#book}

- [新刊小説は滅亡について考えた方がいい « マガジン航[kɔː]](https://magazine-k.jp/2015/12/04/downfall-of-novels/)

いやぁ

{{< fig-quote title="新刊小説は滅亡について考えた方がいい" link="https://magazine-k.jp/2015/12/04/downfall-of-novels/" >}}
<q>消費者にとって、新刊小説を買わないことには、メリットがあるからである。<br>
新刊小説は、買わない方がいいからである。<br>
なぜ買わない方がいいのか。どんなメリットがあるのか。それは「新刊小説の滅亡」に書いた。</q>
{{< /fig-quote >}}

はウケた。
なかなかのセールストークである。
これ見てうっかり『[本をめぐる物語](http://store.kadokawa.co.jp/shop/g/g321509000668/)』を買う人もいるかもしれない。
いや，私は買わないけど。

まぁ，でも，しかし，続きを読めば「なぜ」の答えは書いてあって

{{< fig-quote title="新刊小説は滅亡について考えた方がいい" link="https://magazine-k.jp/2015/12/04/downfall-of-novels/" >}}
<q>有体にいえば、一回読んだら置き場所に困るような文学など、買わない方がいい、ということだ。</q>
{{< /fig-quote >}}

と結ばれている。

わたしが15年間に溜め込んだマンガやラノベの山（文字通りの山というか蟻塚）を[処分]({{< ref "/remark/2015/0920-diary.md" >}})したら軽トラ山盛り1杯分あった。
これからは紙の本を買うのは仕事に絡む本か余程好きな作家さんのみにしようと心に決めた。
たしかに「一回読んだら置き場所に困るような」紙の本などおよびではない。
買うのなら E ブックで買うか[^d]， E ブックで買えないのなら図書館へ行くべきだ。

[^d]: E ブックの実体はクラウドに置かれる。ユーザは E ブックにアクセスする権利を買っているだけで所有しているわけではない。ちなみに図書館は太古からあるクラウドである。

要するに私たちがしたいことは「本を読む」ことであり「本を所有」することではないということだ。
「所有する」ことにお金を使うバカバカしさに（もはや消費者ではない）ユーザは気づいてしまったのだ。
これはたとえば映像や音楽でも同じ。
ユーザは観て聴いて読むことができるのなら手段は問わない（つまりお金を払っても構わない）し，それを阻もうとするのなら迂回するか，迂回もできないのなら「観て聴いて読む」行為を止めるだけである。

もう「コンテンツを売って儲ける」ビジネスは破綻している。
それは近代でのみ通用する泡沫のビジネスモデルだったのだろう。

音楽業界が失くなっても音楽は存在し続けるし，書籍出版業界が失くなっても本は存在し続ける。
そろそろ気づけ。
