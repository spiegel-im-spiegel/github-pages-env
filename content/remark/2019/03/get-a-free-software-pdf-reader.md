+++
title = "いつの間にか Evince Windows 版がなくなっていた"
date = "2019-03-07T23:46:07+09:00"
description = "というわけで，とりあえず MuPDF を試してみることにする。"
image = "/images/attention/kitten.jpg"
tags  = [ "security", "vulnerability", "pdf", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近 PDF への電子署名を騙す脆弱性が公開された。

- [How To Spoof PDF Signatures](https://web-in-security.blogspot.com/2019/02/how-to-spoof-pdf-signatures.html)
- [PDF Signature Spoofing](https://www.pdf-insecurity.org/index.html)
- [Digital Signatures in PDFs Are Broken - Schneier on Security](https://www.schneier.com/blog/archives/2019/03/digital_signatu.html)

脆弱性は大きく3つあるようだ。

- Universal Signature Forgery (USF) : [CVE-2018-16042](https://nvd.nist.gov/vuln/detail/CVE-2018-16042)
    - CVSSv3 深刻度 7.5 (`CVSS:3.0/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:N`)
- Incremental Saving Attack (ISA) : [CVE-2018-18688](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-18688)
- Signature Wrapping Attack (SWA) : [CVE-2018-18689](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2018-18689)

対象となるアプリケーションあるいはサービスは以下の通り。

- [PDF Viewer](https://www.pdf-insecurity.org/signature/viewer.html)
- [PDF Online Validation Services](https://www.pdf-insecurity.org/signature/services.html)

該当するアプリケーションあるいはサービスを利用している方は今後の動向に注視していただきたい。

で，ここからが本番で，私が常用している [Evince] はどうだっけ？ と久しぶりにサイトを覗いてみたら Windows 用のバイナリがなくなってるじゃない。
どうも Windows のサポートは止めたらしい（MSYS2 用のソースコードはあるようだけど...） orz

しょうがない， [PDFreaders.org] から適当に探してみよう。

Windows 用のフリーな PDF リーダーと言えば [Sumatra PDF] が有名だが，こちらの最終更新は2016年で止まっていてどうにも使う気が起きない。
というわけで，とりあえず [MuPDF] を試してみることにする。
活動も活発なようだし Android 版を含めてマルチプラットフォーム対応みたいだし。

他にも色々試して，よさそうのを常用してみよう。

[Evince]: http://projects.gnome.org/evince/
[PDFreaders.org]: https://pdfreaders.org/ "Get a Free Software PDF reader!"
[Sumatra PDF]: https://www.sumatrapdfreader.org/ "Free PDF Reader - Sumatra PDF"
[MuPDF]: https://mupdf.com/
