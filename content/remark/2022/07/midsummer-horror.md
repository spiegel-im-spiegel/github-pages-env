+++
title = "クラウドストレージのアカウント情報がクラウドストレージにある罠"
date =  "2022-07-29T20:41:52+09:00"
description = "真夏のホラー"
image = "/images/attention/kitten.jpg"
tags = [ "risk", "management", "nas" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

久しぶりにざんぞがさばる（出雲弁）記事を見た。

- [ある「パソコンの大先生」の死 – WirelessWire News](https://wirelesswire.jp/2022/07/82884/)

## 今回のお題

今回の条件は割とありがちではないだろうか。

1. Windows パソコン1台のみ
2. パスワード管理ツールあり
   - ただし Windows 専用プロプライエタリ・ツール
   - パスワード情報はクラウド・ストレージにバックアップあり
3. ローカル NAS はなし

この唯一のパソコンが故障して復旧不能となった場合にどうなるかという話。

{{< fig-quote type="markdown" title="ある「パソコンの大先生」の死" link="https://wirelesswire.jp/2022/07/82884/" >}}
しかし、そのフリーウェアはWindows用のプロプライエタリソフトウエアのため、ワタシのスマホではデータを閲覧できません。それに気付いたのは、スマホで新しいPCを購入する手続きを進め、最後にクレジットカードの会員向けサービスのパスワードを聞かれ、途方に暮れたときです。これでは新しいPCが調達できないじゃないか！
{{< /fig-quote >}}

ざんぞがさばる！

## NAS とクラウドストレージ

最近は「クラウドストレージがあるならローカルの NAS とかいらなくね？」という向きも多いようだ。

でもクラウドストレージにアクセスするためにはアカウント情報が必要で，更に多要素認証を使うのであればその分だけ認証手段を確保しておく必要がある。
それをどこに保持しておくか。

個人的にはローカルに NAS を置いておくのはまだ意味があると思う。
もっと言えばローカル NAS とクラウドストレージを連携できるならなおよし。

- [秋 NAS は俺に喰わせろ！]({{< ref "/remark/2021/10/nas.md" >}})

私もクラウドストレージは機密保持の観点からはあまり信用してない。
個人向けのクラウドストレージ・サービスの多くは他者と情報共有しやすいよう設計されているため「うっかり漏洩」してしまう可能性も排除できない。
法人向けならユーザごとに細かく権限を設定できるので逆に安心なんだけどね。

そういう意味じゃ信用してないのはクラウドサービスではなくて自分自身か（笑）

まぁ，クラウドストレージは「バックアップ用」と割り切って sensitive な情報については置かないようにするか暗号化するのがいいんだろうね。

## マルチプラットフォームは重要

今回のもうひとつのポイントはパスワード管理ツールが Windows 専用だったことだろう。
プロプライエタリなツールであることは今回のケースでは問題ではない。

最近はセキュリティ企業も独自のパスワード管理ツールを出しているが，これがもし特定 OS 専用なら考え直したほうがいいかもしれない（実際のところは知らない）。

私は昔から [KeePass](https://keepass.info/ "KeePass Password Safe") 派なのだが，最近は派生ツールである [KeePassXC](https://keepassxc.org/ "KeePassXC Password Manager") および Android 版の [Keepass2Android](https://play.google.com/store/apps/details?id=keepass2android.keepass2android "Keepass2Android Password Safe - Apps on Google Play") を愛用している。

- [KeePassXC Password Manager](https://keepassxc.org/)
- [GitHub - keepassxreboot/keepassxc: KeePassXC is a cross-platform community-driven port of the Windows application “Keepass Password Safe”.](https://github.com/keepassxreboot/keepassxc)
- [Keepass2Android Password Safe - Apps on Google Play](https://play.google.com/store/apps/details?id=keepass2android.keepass2android&hl=en_US&gl=JP)

## パスワードを紙に書くのは案外悪くない

（特に同居人が居る場合は）管理に気を使う必要はあるが，パスワード情報を紙に書き出すというのは案外悪くないと思っている。
もちろん全て書き出す必要はなくて，いざというときに最低限必要な情報のみに厳選する。

私の自宅環境はパソコン（Linux 機），NAS，スマホの3機構成でこれらが同時に使えなくなる事態は真面目に考えてなかったのだが，今回の記事を読んで，考えを改めた。

{{< fig-quote title="マーフィーの法則" >}}
起こる可能性のあることは，いつか実際に起こる（If it can happen, it will happen）
{{< /fig-quote >}}

## 作業記録は大事

上述のようにパスワードを紙に書き出すとしても，どのアカウント情報を書き出すべきかはよくよく検討しないといけない。
その一環としてスマホやパソコンのセットアップ時の作業記録を録っておくことをおすすめする。

以下はスマホの機種変更した際の作業記録の概要。

- [ついカッとなって機種変した，反省はしない]({{< ref "/remark/2020/05/changing-smartphone.md" >}})

それを見れば必要な情報が分かるはず。
Android 機のセットアップには Google アカウントが絶対に必要とか。

今回の記事を書かれた yomoyomo さんにはご愁傷さまとしか言いようがないが，私としてはこれを好機として自宅環境を見直すとしようか。

## 参考

{{% review-paapi "B0855LMP81" %}} <!-- Synology DS220j -->
{{% review-paapi "B08V8LNR2H" %}} <!-- HDD WD Red Plus -->
