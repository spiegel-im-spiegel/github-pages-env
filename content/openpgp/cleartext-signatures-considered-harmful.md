+++
title = "クリアテキスト署名の危険性"
date = "2026-01-07T16:46:00+09:00"
description = "注意喚起を促すブログ記事の紹介 / GnuPG 最新版へのアップデート"
image = "/images/attention/openpgp.png"
tags = [ "openpgp", "gnupg", "security", "cryptography" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Mastodon の TL で以下のアナウンスを見かけた。

{{< fig-gen >}}
<blockquote class="mastodon-embed" data-embed-url="https://mstdn.social/@GnuPG/115842294964864180/embed" style="background: #FCF8FF; border-radius: 8px; border: 1px solid #C9C4DA; margin: 0; max-width: 540px; min-width: 270px; overflow: hidden; padding: 0;"> <a href="https://mstdn.social/@GnuPG/115842294964864180" target="_blank" style="align-items: center; color: #1C1A25; display: flex; flex-direction: column; font-family: system-ui, -apple-system, BlinkMacSystemFont, 'Segoe UI', Oxygen, Ubuntu, Cantarell, 'Fira Sans', 'Droid Sans', 'Helvetica Neue', Roboto, sans-serif; font-size: 14px; justify-content: center; letter-spacing: 0.25px; line-height: 20px; padding: 24px; text-decoration: none;"> <svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="32" height="32" viewBox="0 0 79 75"><path d="M63 45.3v-20c0-4.1-1-7.3-3.2-9.7-2.1-2.4-5-3.7-8.5-3.7-4.1 0-7.2 1.6-9.3 4.7l-2 3.3-2-3.3c-2-3.1-5.1-4.7-9.2-4.7-3.5 0-6.4 1.3-8.6 3.7-2.1 2.4-3.1 5.6-3.1 9.7v20h8V25.9c0-4.1 1.7-6.2 5.2-6.2 3.8 0 5.8 2.5 5.8 7.4V37.7H44V27.1c0-4.9 1.9-7.4 5.8-7.4 3.5 0 5.2 2.1 5.2 6.2V45.3h8ZM74.7 16.6c.6 6 .1 15.7.1 17.3 0 .5-.1 4.8-.1 5.3-.7 11.5-8 16-15.6 17.5-.1 0-.2 0-.3 0-4.9 1-10 1.2-14.9 1.4-1.2 0-2.4 0-3.6 0-4.8 0-9.7-.6-14.4-1.7-.1 0-.1 0-.1 0s-.1 0-.1 0 0 .1 0 .1 0 0 0 0c.1 1.6.4 3.1 1 4.5.6 1.7 2.9 5.7 11.4 5.7 5 0 9.9-.6 14.8-1.7 0 0 0 0 0 0 .1 0 .1 0 .1 0 0 .1 0 .1 0 .1.1 0 .1 0 .1.1v5.6s0 .1-.1.1c0 0 0 0 0 .1-1.6 1.1-3.7 1.7-5.6 2.3-.8.3-1.6.5-2.4.7-7.5 1.7-15.4 1.3-22.7-1.2-6.8-2.4-13.8-8.2-15.5-15.2-.9-3.8-1.6-7.6-1.9-11.5-.6-5.8-.6-11.7-.8-17.5C3.9 24.5 4 20 4.9 16 6.7 7.9 14.1 2.2 22.3 1c1.4-.2 4.1-1 16.5-1h.1C51.4 0 56.7.8 58.1 1c8.4 1.2 15.5 7.5 16.6 15.6Z" fill="currentColor"/></svg> <div style="color: #787588; margin-top: 16px;">Post by @GnuPG@mstdn.social</div> <div style="font-weight: 500;">View on Mastodon</div> </a> </blockquote> <script data-allowed-prefixes="https://mstdn.social/" async src="https://mstdn.social/embed.js"></script>
{{< /fig-gen >}}

まとめて書かれているが，内容としては

1. 注意喚起を促すブログ記事の紹介
2. GnuPG 最新版へのアップデート

の2点ある。
このうち，前半の注意喚起を促すブログ記事がこちら

- [Cleartext Signatures Considered Harmful](https://www.gnupg.org/blog/20251226-cleartext-signatures.html)

クリアテキスト署名とは，以下のように平文テキストと電子署名が一体化した形式の署名である。

```text
-----BEGIN PGP SIGNED MESSAGE-----
Hash: SHA512

Far out in the uncharted backwaters of the unfashionable  end  of
the  western  spiral  arm  of  the Galaxy lies a small unregarded
yellow sun.
-----BEGIN PGP SIGNATURE-----

iJEEARYKADkWIQSHd0YfKgdOvEgNNZQZzByeCFsQegUCaU5mGhsUgAAAAAAEAA5t
YW51MiwyLjUrMS4xMSwyLDIACgkQGcwcnghbEHq27gEAqE+Cw1FmIlDXyzc1+5K8
3e60/3TjpqpFmkmuot8ZROMBAIrQXMhfN3gr7jdsxLuV1L7+IzHSRyUMlelZSnAs
k+AL
=kCuN
-----END PGP SIGNATURE-----
```

このブログ記事を [Kagi Assistant] に要約してもらった。
こんな感じ。

{{< fig-quote type="markdown" >}}
以下は指定ページの要約です。

- 要旨：PGPの「クリアテキスト署名」は一見便利だが、表示と実際に署名された内容が一致しない可能性があり危険であると指摘している。検証が容易ではなく攻撃に対して脆弱であるため、使用を避けるべきという結論を示している。
- 主な問題点：
  - 端末やビューアのエスケープコードや制御文字により、表示内容が改変され得ること。
  - エンコーディングや改行の扱いの差異で、見た目と署名対象が食い違う恐れがあること。
  - 偽のヘッダーやコメント行の誤解を利用した混乱を招く手法が存在すること。
  - これらの落とし穴は歴史的背景の中で認識されてきたが、現在も実用上のリスクとなっていること。
- 推奨事項：
  - クリアテキスト署名の使用を避け、代わりにデタッチ署名（別ファイルの署名）を使うことを強く推奨すること。
  - 電子メールではPGP/MIMEの利用を推奨すること（テキスト表示の差異による誤検証を防ぐため）。
- 結論：可読性の利点はあるが、セキュリティ上の欠点が重大であるため、より安全な署名形式へ移行すべきである。
{{< /fig-quote >}}

んー。
だいたい合ってるかな。

クリアテキスト署名は大昔の素朴な電子メールなど限られたコミュニケーション手段しかなかった時代の形式で，現代の多様なデータをやり取りする際のエンコーディング手段としては適切とは言えなくなっている。
なので，データと署名を分ける分離署名（detached signature）を用いるか，電子メールであれば PGP/MIME を使うことを推奨する，という内容である。

アナウンスの後半は古いバージョン [GnuPG] には ASCII-Armor 形式のテキストデータの解析においてメモリ破壊があるので最新版にアップデートしてね，というもの。
ちなみに，この不具合は 2025-11-27 にリリースされたバージョン 2.5.14 で修正されている。
バージョン 2.5.16 および 2.4.9 は 2025-12-30 にリリースされている。

なお [GnuPG] 2.5 系は安定版に昇格している。
さらに 2.4 系は 2026-06-30 でサポートが終了する。

{{< fig-quote type="markdown" title="[Announce] GnuPG 2.5.16 released" link="https://lists.gnupg.org/pipermail/gnupg-announce/2025q4/000500.html" lang="en" >}}
Note that the 2.5 series is now declared the stable version of GnuPG.
The oldstable 2.4 series will reach end-of-life in just 6 months.

The main features in the 2.5 series are improvements for 64 bit Windows
and the introduction of Kyber (aka ML-KEM or FIPS-203) as PQC encryption
algorithm.  Other than PQC support the 2.6 series will not differ a lot
from 2.4 because the majority of changes are internal to make use of
newer features from the supporting libraries.
{{< /fig-quote >}}

## ブックマーク

- [Responsible Disclosure Requires Accountability: Putting Recent GnuPG Security Reports into Context](https://gnupg.com/20260122-39C3_reply_gpg_fail.html)

- [GnuPG 2.5.17 のリリース【セキュリティ・アップデート】]({{< ref "/release/2026/01/gnupg-2_5_17-is-released.md" >}})

[OpenPGP]: http://openpgp.org/
[RFC 4880]: https://tools.ietf.org/html/rfc4880 "RFC 4880 - OpenPGP Message Format"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[Kagi Assistant]: https://kagi.com/assistant "The Assistant"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "4900900028" %}} <!-- PGP―暗号メールと電子署名 -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
