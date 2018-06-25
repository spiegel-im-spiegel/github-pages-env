+++
title = "Thunderbird でメール暗号化"
date = "2018-06-25T21:32:25+09:00"
description = "今回は Thunderbird で OpenPGP 暗号化メールをやり取りする手順について簡単に紹介する。"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "security", "cryptography", "mua", "thunderbird", "messaging" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

フィードバックで少しやり取りして分かったのだが JPCERT/CC の記事は古いまま更新されていないようだ。

- [はじめての暗号化メール (Thunderbird編)](https://www.jpcert.or.jp/magazine/security/pgpquick.html)

なので今回は [Thunderbird] で [OpenPGP] 暗号化メールをやり取りする手順について簡単に紹介する。
なお [GnuPG] の導入と鍵の生成は完了しているものとする。

- [GnuPG for Windows インストール編]({{< relref "openpgp/using-gnupg-for-windows-1.md" >}})
    - [GnuPG for Windows : gpg-agent について]({{< relref "openpgp/using-gnupg-for-windows-2.md" >}})
- [GnuPG チートシート（鍵作成から失効まで）]({{< relref "openpgp/gnupg-cheat-sheet.md" >}})

## [Enigmail] の導入

[Thunderbird] で [OpenPGP] 暗号化メールをやり取りするには [Enigmail] アドオンを導入する。

{{< fig-img src="https://farm2.staticflickr.com/1832/29123820638_1ba7aaecf2.jpg" title="Search Enigmail" link="https://www.flickr.com/photos/spiegel/29123820638/" >}}

この状態で `[インストール]` ボタンを押してインストールを完了させる。

## [Enigmail] の設定

インストールが完了したらメニューから [Enigmail] の「セットアップウィザード」を開く。

{{< fig-img src="https://farm2.staticflickr.com/1798/41186104180_4107cef79b.jpg" title="Enigmail menu" link="https://www.flickr.com/photos/spiegel/41186104180/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1767/42946809382_ef30dc6208.jpg" title="Enigmail setup-wizard (1)" link="https://www.flickr.com/photos/spiegel/42946809382/" >}}

ここでは初心者向けの設定を紹介する。

{{< fig-img src="https://farm2.staticflickr.com/1795/42278121714_574e5a8af5.jpg" title="Enigmail setup-wizard (2)" link="https://www.flickr.com/photos/spiegel/42278121714/" >}}

メールアカウントに対応する鍵を既に作成している場合はその鍵を選択する（[Enigmail] はメールアカウントと [OpenPGP] 鍵との関連をメールアドレスで行っている）。
対応する鍵がない場合には「メッセージ署名/暗号に使用する鍵ペアを新規に作成する」を選択して [OpenPGP] 鍵を作成する[^key1]。

[^key1]: セットアップウィザードでは RSA 4096ビット長の [OpenPGP] 鍵を生成する。

{{< fig-img src="https://farm2.staticflickr.com/1826/29124707788_1f1008d37f.jpg" title="Enigmail setup-wizard (3)" link="https://www.flickr.com/photos/spiegel/29124707788/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1763/42947411712_144d47a711.jpg" title="Enigmail setup-wizard (4)" link="https://www.flickr.com/photos/spiegel/42947411712/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1823/42278604034_64dbcfcb14.jpg" title="Enigmail setup-wizard (5)" link="https://www.flickr.com/photos/spiegel/42278604034/" >}}

ここで失効証明書を作っておく[^rvk1]。
`[失効証明書の生成]` ボタンを押して失効証明書を保存すると `[次へ]` ボタンが有効になる。

[^rvk1]: 失効証明書は鍵を失効させる際に必要となる。失効した公開鍵を配布するのを忘れずに。詳しくは「[GnuPG チートシート（鍵作成から失効まで）]({{< relref "openpgp/gnupg-cheat-sheet.md" >}})」を参照のこと。

{{< fig-img src="https://farm2.staticflickr.com/1785/42997052861_13d6919259.jpg" title="Enigmail setup-wizard (6)" link="https://www.flickr.com/photos/spiegel/42997052861/" >}}

これでセットアップ完了。
なお，アカウント設定ではもう少し詳細な設定が可能である。

{{< fig-img src="https://farm2.staticflickr.com/1807/29125191098_253270a7d2.jpg" title="Enigmail setup-wizard (6)" link="https://www.flickr.com/photos/spiegel/29125191098/" >}}

## 暗号化メールの作成と表示

では自分自身宛にメールを作成して動作確認してみよう。

{{< fig-img src="https://farm2.staticflickr.com/1771/29125311968_87bc4a8803.jpg" title="Make mail" link="https://www.flickr.com/photos/spiegel/29125311968/" >}}

矢印で示した部分で暗号化および電子署名の有無を指定する。
ここでは暗号化と電子署名を有効にしてメールを送信する。

送信時に件名を保護するか問い合わせがある。

{{< fig-img src="https://farm2.staticflickr.com/1776/42948011252_d01053b35c.jpg" title="Encrypt title" link="https://www.flickr.com/photos/spiegel/42948011252/" >}}

`[件名を保護する]` を選択すると件名を “Encrypted Message” に置き換えて送信する。
本来の件名は復号時に表示される。

では実際に受信したメールを復号して表示してみよう。

{{< fig-img src="https://farm2.staticflickr.com/1792/29125520728_f8573e31c7.jpg" title="Receive message" link="https://www.flickr.com/photos/spiegel/29125520728/" >}}

メールの復号と署名検証が正しく行われていることが確認できた。

## [OpenPGP] 鍵の管理

メニューから [Enigmail] の「鍵の管理」を開く。

{{< fig-img src="https://farm2.staticflickr.com/1768/41188654830_a1cbaf817e.jpg" title="menu key management" link="https://www.flickr.com/photos/spiegel/41188654830/" >}}

{{< fig-img src="https://farm2.staticflickr.com/1791/42949496932_72d496e909.jpg" title="Enigmail key management" link="https://www.flickr.com/photos/spiegel/42949496932/" >}}

鍵の管理で公開鍵のインポート・エクスポートや電子署名等の操作ができる。

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[RFC 4880bis]: https://datatracker.ietf.org/doc/draft-ietf-openpgp-rfc4880bis/ "draft-ietf-openpgp-rfc4880bis - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Thunderbird]: https://www.thunderbird.net/ "Thunderbird — Software made to make email easier. — Mozilla"
[Enigmail]: https://addons.mozilla.org/thunderbird/addon/enigmail/ "Enigmail :: Add-ons for Thunderbird"
