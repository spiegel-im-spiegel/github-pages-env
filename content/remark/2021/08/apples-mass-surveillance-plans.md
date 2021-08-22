+++
title = "Apple 監視社会化計画（裏口を穿つ）"
date =  "2021-08-22T12:47:46+09:00"
description = "この件が恐ろしいのは Apple が end-to-end 暗号通信の定義を都合よく書き換えている点だ。"
image = "/images/attention/kitten.jpg"
tags = [ "security", "privacy", "risk", "surveillance-capitalism", "censorship" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

毎度のことながら感度が低くてゴメンペコン。
起点は EFF のこの記事でいいかな。

- [Apple's Plan to "Think Different" About Encryption Opens a Backdoor to Your Private Life | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2021/08/apples-plan-think-different-about-encryption-opens-backdoor-your-private-life)
- [暗号化の「見方を変えた」Appleが監視社会へのバックドアを開く | P2Pとかその辺のお話R](https://p2ptk.org/privacy/3329)

この件に関しては Bruce Schneier 先生も反応されている。

- [Apple Adds a Backdoor to iMessage and iCloud Storage - Schneier on Security](https://www.schneier.com/blog/archives/2021/08/apple-adds-a-backdoor-to-imesssage-and-icloud-storage.html)
- [Apple’s NeuralHash Algorithm Has Been Reverse-Engineered - Schneier on Security](https://www.schneier.com/blog/archives/2021/08/apples-neuralhash-algorithm-has-been-reverse-engineered.html)
- [More on Apple’s iPhone Backdoor - Schneier on Security](https://www.schneier.com/blog/archives/2021/08/more-on-apples-iphone-backdoor.html)

この件が恐ろしいのは Apple が end-to-end 暗号通信の定義を都合よく書き換えている点だ。

{{< fig-quote type="markdown" title="Expanded Protections for Children - Frequently Asked Questions v1.1" link="https://www.apple.com/child-safety/pdf/Expanded_Protections_for_Children_Frequently_Asked_Questions.pdf" lang="en" >}}
Does this break end-to-end encryption in Messages? 

No. This doesn’t change the privacy assurances of Messages, and Apple never gains access to communications as a result of this feature. Any user of Messages, including those with communication safety enabled, retains control over what is sent and to whom. **If the feature is enabled for the child account, the device will evaluate images in Messages and present an intervention if the image is determined to be sexually explicit**. For accounts of children age 12 and under, parents can set up parental notifications which will be sent if the child confirms and sends or views an image that has been determined to be sexually explicit. None of the communications, image evaluation, interventions, or notifications are available to Apple.
{{< /fig-quote >}}

強調部分は私がやりました。
いや，それはもう “end-to-end encryption” じゃないから！ ある意味で（自製品に E2E 暗号機能を組み込もうとしない） Google よりも邪悪だよ。

これだけ見ても Apple が個人のプライバシーに対して敵対的な企業であることが分かる。
こういうことをやらかす企業がいるから [Signal](https://signal.org/) や [GnuPG](https://gnupg.org/ "The GNU Privacy Guard") のような製品で自衛せざるを得ない。
まぁ，私は Apple 信者じゃないので Apple 製品を使うことはないけど GAFA はだいたい似たり寄ったりだからなぁ。

ちうわけで，この件は注視しておこう。

## ブックマーク

- [Child Safety - Apple](https://www.apple.com/child-safety/)

- [If You Build It, They Will Come: Apple Has Opened the Backdoor to Increased Surveillance and Censorship Around the World | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2021/08/if-you-build-it-they-will-come-apple-has-opened-backdoor-increased-surveillance)
  - [それを作れば彼らはやってくる：Appleが開く世界的な監視・検閲へのバックドア | P2Pとかその辺のお話R](https://p2ptk.org/privacy/3334)
- [Speak Out Against Apple’s Mass Surveillance Plans | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2021/08/speak-out-against-apples-mass-surveillance-plans)
  - [Appleの監視社会化計画にNOの声をあげよう | P2Pとかその辺のお話R](https://p2ptk.org/privacy/3340)
- [How LGBTQ+ Content is Censored Under the Guise of "Sexually Explicit" | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2021/08/how-lgbtq-content-censored-under-guise-sexually-explicit)
  - [「性的に露骨」と称して検閲されるLGBTQ+コンテンツ | P2Pとかその辺のお話R](https://p2ptk.org/privacy/3336)
- [EFF Joins Global Coalition Asking Apple CEO Tim Cook to Stop Phone-Scanning | Electronic Frontier Foundation](https://www.eff.org/deeplinks/2021/08/eff-joins-global-coalition-asking-apple-ceo-tim-cook-stop-phone-scanning)

- [One Bad Apple - The Hacker Factor Blog](https://www.hackerfactor.com/blog/index.php?/archives/929-One-Bad-Apple.html)
- [In internal memo, Apple addresses concerns around new Photo scanning features, doubles down on the need to protect children - 9to5Mac](https://9to5mac.com/2021/08/06/apple-internal-memo-icloud-photo-scanning-concerns/)
- [AppleがiPhoneの写真やメッセージをスキャンして児童の性的搾取を防ぐと発表、電子フロンティア財団などから「ユーザーのセキュリティとプライバシーを損なう」という抗議の声も - GIGAZINE](https://gigazine.net/news/20210806-apple-csam-icloud-photo-scan/)
- [国際的な連合は、アップル社にiPhone、iPad等製品への監視機能組み込み計画の中止を求める のコピー - Google ドキュメント](https://docs.google.com/document/d/1aBooqLlOuzIeKw5qCJgnAHrCli91c6x8SKCXjjiNsyY/edit)

- [誰がプライバシーを支配するのか]({{< ref "/remark/2018/04/handling-privacy.md" >}})
- [プライバシーなどク◯喰らえ]({{< ref "/remark/2020/01/privacy.md" >}})

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "B093KFTV53" %}} <!-- 監視資本主義 -->
