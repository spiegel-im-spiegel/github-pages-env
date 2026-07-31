+++
title = "マジで恥ずかしいミス"
date =  "2022-03-05T09:58:12+09:00"
description = "鍵管理は難しいやね。"
image = "/images/attention/openpgp.png"
tags = [ "security", "vulnerability", "cryptography", "k-tai" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ハードウェアに暗号システムを乗っけるというと，どうしても1990年代の「クリッパーチップ」を思い起こさるが[^cc]，時代を経た今でも鍵管理はなかなか大変なようだ。

[^cc]: 当時のクリッパーチップが問題視されたのは「鍵預託（key escrow）」の観点からであるが，クリッパーチップが幕引きとなった直接のトリガーは鍵預託のシステムに欠陥が見つかったからとされている。

というわけで， Bruce Schneier 先生をして {{% quote lang="en" %}}[really embarrassing mistake](https://www.schneier.com/blog/archives/2022/03/samsung-encryption-flaw.html "Samsung Encryption Flaw - Schneier on Security"){{% /quote %}} と言わしめた Samsung 社製 Android 機のハードウェア・キーストアの欠陥について指摘した論文がこれ：

- {{< pdf-file title="Trust Dies in Darkness: Shedding Light on Samsung’s TrustZone Keymaster Design" link="https://eprint.iacr.org/2022/208.pdf" >}}

引用の引用で申し訳ないが，論文の概要（Abstract）はこう。

{{< fig-quote type="markdown" title="Trust Dies in Darkness: Shedding Light on Samsung’s TrustZone Keymaster Design" link="https://eprint.iacr.org/2022/208.pdf" lang="en" >}}
ARM-based Android smartphones rely on the TrustZone hardware support for a Trusted Execution Environment (TEE) to implement security-sensitive functions. The TEE runs a separate, isolated, TrustZone Operating System (TZOS), in parallel to Android. The implementation of the cryptographic functions within the TZOS is left to the device vendors, who create proprietary undocumented designs.<br>
In this work, we expose the cryptographic design and implementation of Android’s Hardware-Backed Keystore in Samsung’s Galaxy S8, S9, S10, S20, and S21 flagship devices. We reversed-engineered and provide a detailed description of the cryptographic design and code structure, and we unveil severe design flaws. We present an IV reuse attack on AES-GCM that allows an attacker to extract hardware-protected key material, and a downgrade attack that makes even the latest Samsung devices vulnerable to the IV reuse attack. We demonstrate working key extraction attacks on the latest devices. We also show the implications of our attacks on two higher-level cryptographic protocols between the TrustZone and a remote server: we demonstrate a working FIDO2 WebAuthn login bypass and a compromise of Google’s Secure Key Import.<br>
We discuss multiple flaws in the design flow of TrustZone based protocols. Although our specific attacks only apply to the ≈100 million devices made by Samsung, it raises the much more general requirement for open and proven standards for critical cryptographic and security designs.
{{< /fig-quote >}}

ポイントのひとつは，この欠陥をリバース・エンジニアリングで発見したことかね。
もう少し詳しく紹介すると，4.1章の冒頭にこう書かれている。

{{< fig-quote type="markdown" title="Trust Dies in Darkness: Shedding Light on Samsung’s TrustZone Keymaster Design" link="https://eprint.iacr.org/2022/208.pdf" lang="en" >}}
As we discussed in Section 3, the wrapping key used to encrypt the key blobs (HDK) is derived using a salt value computed by the Keymaster TA. In v15 and v20-s9 blobs, **the salt is a deterministic function that depends only on the application ID and application data (and constant strings)**, which the Normal World client fully controls. This means that for a given application, all key blobs will be encrypted using the same key. As the blobs are encrypted in AES-GCM mode-of-operation, the security of the resulting encryption scheme depends on its IV values never being reused.
{{< /fig-quote >}}

（強調は私がやりました）

マジですか。
これは恥ずかしいわ（笑）

この論文はたまたまシェアの大きい Galaxy 端末を調べたのだろうが，他の携帯端末や暗号デバイスなんかでも評価していかないとマズいかもねぇ。
鍵管理は難しいやね。

## 参考文献

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
