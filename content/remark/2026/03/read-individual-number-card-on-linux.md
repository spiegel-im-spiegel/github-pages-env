+++
title = "Linux で個人番号カードを読む"
date =  "2026-03-31T20:23:14+09:00"
description = "自宅の Ubuntu 機に IC カードリーダーを接続し，個人番号カードの内容を読み取る。 MPA for Linux については後日に試してみたい。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "my-number", "pki", "certification", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

自宅の Ubuntu 機に IC カードリーダーを接続し，個人番号カードの内容を読み取る。

## 事前準備

今回は IO DATA の非接触式 IC カードリーダー USB-NFC4 を使用する。

- [USB-NFC4 | ICカードリーダーライター | アイ・オー・データ機器 I-O DATA][USB-NFC4]

昨年の夏頃に Amazon で安売りしてたのを買ったのだが，そのまま放置していた。
いざ確定申告で使おうとしたら Windows 機で認識できなくて[使えなかった]({{< ref "/remark/2026/03/tax-return-review.md" >}} "確定申告のふりかえり")。
そのまま捨て置くのはもったいないので Linux で使えるか試そうという話である。

Ubuntu 側で必要なのは以下のソフトウェア

- [OpenSC/OpenSC: Open source smart card tools and middleware. PKCS#11/MiniDriver][OpenSC]
- デバイスドライバ : [CIR315A] 用のドライバで代用
- [jpki/myna: マイナンバーカード・ユーティリティ・JPKI署名ツール · GitHub][myna]

では早速はじめよう。

## インストール

[Wiki](https://github.com/OpenSC/Wiki "OpenSC/Wiki") によると Linux 版の [OpenSC] は自前でビルドしろとあるが， Ubuntu であればバイナリが提供されているっぽいのでそちらを使う。

```text
$ sudo aptitude install opensc opensc-pkcs11 pcscd pcsc-tools libpcsclite1 libusb-1.0-0 libpcsclite-dev libusb-1.0-0-dev
libpcsclite1 は、要求されたバージョン (2.3.3-1) で既にインストールされています
libusb-1.0-0 は、要求されたバージョン (2:1.0.29-2) で既にインストールされています
libpcsclite1 は、要求されたバージョン (2.3.3-1) で既にインストールされています
libusb-1.0-0 は、要求されたバージョン (2:1.0.29-2) で既にインストールされています
以下の新規パッケージがインストールされます:
  libccid{a} libeac3{a} libintl-perl{a} libintl-xs-perl{a} libpcsc-perl{a} libpcsclite-dev libusb-1.0-0-dev libusb-1.0-doc{a} opensc 
  opensc-pkcs11 pcsc-tools pcscd 
更新: 0 個、新規インストール: 12 個、削除: 0 個、保留: 0 個。
アーカイブの 2,824 kB を取得する必要があります。展開後に 12.8 MB のディスク領域が新たに消費されます。
先に進みますか? [Y/n/?] 
```

起動確認だけしておく。

```text
$ opensc-tool -i
OpenSC 0.26.1 [gcc  15.2.0]
Enabled features: locking zlib readline openssl pcsc(libpcsclite.so.1)
```

問題はなさそうかな[^osc1]。

[^osc1]: [OpenSC] v0.27 より前のバージョンには[脆弱性が報告](https://github.com/OpenSC/OpenSC/wiki/OpenSC-security-advisories "OpenSC security advisories · OpenSC/OpenSC Wiki")されている。 2026-03-31 時点では Ubuntu の APT リポジトリには反映されてない模様。頑張って反映させてね。

[CIR315A] の製品ページから「USB Linux インストーラ」をダウンロードする。

{{< fig-img-quote src="./cir315a-driver.png" title="CIR315A関連ファイルのダウンロード" link="https://www.abcircle.com/jp/product/14/CIR315A/%e9%9d%9e%e6%8e%a5%e8%a7%a6%e5%bc%8fic%e3%82%ab%e3%83%bc%e3%83%89%e3%83%aa%e3%83%bc%e3%83%80%e3%83%a9%e3%82%a4%e3%82%bf/" width="1452" >}}

ダウンロードしたファイルの内容は以下の通り：

- `Circle_USB_Linux_Installer_v2.2.2_(driver_v.2.2.2).zip` (2025-07-25 時点)
  - `Generic-Debian`
    - `libabcccid_2.2.2-1_amd64.deb`

この `libabcccid_2.2.2-1_amd64.deb` をインストールする。

```text
$ sudo dpkg -i libabcccid_2.2.2-1_amd64.deb
```

[myna] は GitHub の[リリースページ](https://github.com/jpki/myna/releases "Releases · jpki/myna")からバイナリをダウンロードしてインストールする。

- `myna-v0.6.4-x86_64-unknown-linux-gnu.zip` (2026-03-12 時点)
  - `myna`

ファイル `myna` を `PATH` の通ってるディレクトリに置く。
こちらも起動確認だけしておこう。

```text
$ myna help
Usage: myna [OPTIONS] <COMMAND>

Commands:
  text     券面入力補助AP
  visual   券面確認AP
  test     Test card reader
  jpki     公的個人認証
  pin      Pin operation
  unknown  謎のAP
  help     Print this message or the help of the given subcommand(s)

Options:
  -v...          
  -d, --debug    
  -h, --help     Print help
  -V, --version  Print version
```

## 個人番号カードを読み込む

まずは [USB-NFC4] を繋いだだけの状態で IC カードリーダーが認識されているか確認する。

```text
$ opensc-tool -l
# Detected readers (pcsc)
Nr.  Card  Features  Name
0    No              Circle CIR315 CL [CIR315 CL] (137K231232M2) 00 00
```

[USB-NFC4] から「ピッ！」って音がする。
やっと認識してくれたよ。

ではカードリーダーに個人番号カードを乗っけてみる。

```text
$ opensc-tool -l
# Detected readers (pcsc)
Nr.  Card  Features  Name
0    Yes             Circle CIR315 CL [CIR315 CL] (137K231232M2) 00 00
```

Card 項目が Yes になっている。
よしよし。

次に PIN 情報を取得する。

```text
$ pkcs15-tool --list-pins
Using reader with a card: Circle CIR315 CL [CIR315 CL] (137K231232M2) 00 00
PIN [User Authentication PIN]
	Object Flags   : [0x12], modifiable
	ID             : 01
	Flags          : [0x12], local, initialized
	Length         : min_len:4, max_len:4, stored_len:0
	Pad char       : 0x00
	Reference      : 1 (0x01)
	Type           : ascii-numeric
	Tries left     : 3

PIN [Digital Signature PIN]
	Object Flags   : [0x12], modifiable
	ID             : 02
	Flags          : [0x12], local, initialized
	Length         : min_len:6, max_len:16, stored_len:0
	Pad char       : 0x00
	Reference      : 2 (0x02)
	Type           : ascii-numeric
	Tries left     : 5
```

鍵は取り出せるかな。

```text
$ pkcs15-tool --read-certificate 1
Using reader with a card: Circle CIR315 CL [CIR315 CL] (137K231232M2) 00 00
-----BEGIN CERTIFICATE-----

...

-----END CERTIFICATE-----
```

ありゃ。
暗証番号がなくてもいいのか。
でもちゃんと取り出せてるみたい。

もういっちょ。

```text
$ pkcs15-tool --read-certificate 2 --verify-pin --auth-id 02
Using reader with a card: Circle CIR315 CL [CIR315 CL] (137K231232M2) 00 00
Please enter PIN [Digital Signature PIN]: 
-----BEGIN CERTIFICATE-----

...

-----END CERTIFICATE-----
```

こちらもちゃんと取り出せてるようだな。

“Please enter PIN” には署名用パスワードを入力する[^pw1]。
正しく入力すると以下のポップアップが出る。

{{< fig-img-quote src="./opensc-notify.png" link="./opensc-notify.png" width="524" >}}


[^pw1]: 個人番号カードの暗証番号・パスワードには 署名用パスワード（最大16文字の英数字），利用者証明用パスワード（4文字の数字），券面事項入力補助用パスワード（4文字の数字），個人番号カード用（住民基本台帳用）パスワード（4文字の数字）の4つがある。

[myna] のほうも動かしてみよう。

```text
$ myna pin status
券面入力補助AP 暗証番号: 3
券面入力補助AP 暗証番号A: 10
券面入力補助AP 暗証番号B: 10
券面確認AP 暗証番号A: 10
券面確認AP 暗証番号B: 10
JPKIユーザー認証用 暗証番号: 3
JPKIデジタル署名用 パスワード: 5
```

これでパスワード入力を失敗できる（ロックアウトされまでの）残り回数が分かる。

券面情報を取得してみよう。

```text
$ myna text attrs
暗証番号(4桁): ****
氏名    : **********
住所    : **********
生年月日: ********
性別    : *
```

実際にはちゃんと内容が表示されるが，ここでは伏せ字にしている。
あしからず。
暗証番号には券面事項入力補助用パスワードを入力する。

[myna] を使えば PDF ドキュメントなどに電子署名を付与できる。
こんな感じらしい。

```text
$ myna jpki pdf sign input.pdf -o signed.pdf
```

署名の検証は以下の通り。

```text
$ myna jpki pdf verify signed.pdf
```

「JPKI署名用証明書は4属性(氏名・住所・生年月日・性別)を含みますので注意してください」とあるので，実際に運用する場合はホンマにご注意を。

さらに「[MPA for Linux](https://github.com/jpki/myna/tree/master/mpa)」を使えば Linux のブラウザでマイナポータルや e-Tax などのサイトに個人番号カードを使ってログインできるようだ。
ただし（今のところ） Rust のビルド環境が必要なのとブラウザ拡張を無理やり入れるみたいな操作が必要らしいので，今回は割愛する。
またどこかで試そうか。

今回はここまで。
[次回]({{< ref "/remark/2026/04/login-myna-portal-etax-mpa-for-linux.md" >}} "MPA for Linux でログイン検証（Linux で個人番号カードを読む 2）")へ続く。

## ブックマーク

- [Linux & IO DATA USB-NFC4 & マイナンバーカード](https://zenn.dev/sorehaomosiroi/articles/sorehaomosiroi-2024010800_linux_nfc_iodata-usbnfc4)

[USB-NFC4]: https://www.iodata.jp/product/interface/iccardreader/usb-nfc4/ "USB-NFC4 | ICカードリーダーライター | アイ・オー・データ機器 I-O DATA"
[OpenSC]: https://github.com/OpenSC/OpenSC "OpenSC/OpenSC: Open source smart card tools and middleware. PKCS#11/MiniDriver"
[CIR315A]: https://www.abcircle.com/jp/product/14/CIR315A/%e9%9d%9e%e6%8e%a5%e8%a7%a6%e5%bc%8fic%e3%82%ab%e3%83%bc%e3%83%89%e3%83%aa%e3%83%bc%e3%83%80%e3%83%a9%e3%82%a4%e3%82%bf/ "CIR315A - 非接触式ICカードリーダライタ | AB Cir"
[myna]: https://github.com/jpki/myna "jpki/myna: マイナンバーカード・ユーティリティ・JPKI署名ツール · GitHub"

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
