+++
title = "またしてもプライバシーと敵対する Meta"
date =  "2025-06-09T13:36:43+09:00"
description = "当時の CA みたいな話に発展しないことを祈るばかりである。"
image = "/images/attention/kitten.jpg"
tags = [ "privacy", "risk", "artificial-intelligence" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

Meta というか Facebook とプライバシーで思い出すのは1918年のアレである。

{{< fig-quote type="markdown" title="善悪の葛藤" link="/remark/2018/03/name-identification/" >}}
簡単に言うと Campbridge Analytica (CA) とかいうところが過去に学術調査用に収集した Facebook の個人情報をトランプ陣営側にあるとされるデータ分析会社が利用していた，とするものだ。この「ランプ陣営側にあるとされるデータ分析会社」の Facebook アカウントはこの話が表沙汰になる前に停止されていたが， CA が関わっていたということで影響範囲が当初考えられていたよりもかなり大きかったようだ。

まぁ，反トランプな方々にとっては怒り心頭というところだろう。この「怒り」がトランプ政権ではなく Facebook に向かっているところがいい感じに誘導されていて面白い。
{{< /fig-quote >}}

まぁ，当時の Facebook (や Instagram) は既にプライバシーに関してはアレな企業だったので，個人的には今更な感じではあった。
それでも，私が Facebook の利用を控えようと思うようになったきっかけの[ひとつ]({{< ref "/remark/2019/03/withdrawal-from-facebook.md" >}} "Facebook はもう駄目か知らん")になったし Instagram については2020年に[完全にアカウントを削除]({{< ref "/remark/2020/08/divorce-from-instagram.md" >}} "Instagram から足を洗う方法")した。

そこからは Facebook 改め Meta のサービスは最低限にしているのだが，かの企業がまたやらしてるらしい。

- [Meta's New AI App Lets You See Feed of Other People's Conversations - Business Insider](https://www.businessinsider.com/meta-ai-app-public-feed-warning-how-2025-05)

1ヶ月前の記事で恐縮っス。

Meta が自社の AI アシスタント・サービスをリリースした際に

{{< fig-quote type="markdown" title="Introducing the Meta AI App: A New Way to Access Your AI Assistant" link="https://about.fb.com/news/2025/04/introducing-meta-ai-app-new-way-access-ai-assistant/" lang="en" >}}
And just like all our platforms, we built Meta AI to connect you with the people and things you care about. The Meta AI app includes a Discover feed, a place to share and explore how others are using AI. You can see the best prompts people are sharing, or remix them to make them your own. And as always, you’re in control: nothing is shared to your feed unless you choose to post it.
{{< /fig-quote >}}

と述べているのだが，実際には自身の発したプロンプトが「大切な人やもの」どころか不特定多数に晒されているそうで，しかもユーザ側はそのことに思い至らないのか

{{< fig-quote type="markdown" title="Meta's New AI App Lets You See Feed of Other People's Conversations" link="https://www.businessinsider.com/meta-ai-app-public-feed-warning-how-2025-05" lang="en" >}}
Like the woman with the sick pet turtle. Or another person who was asking for advice about what legal measures he could take against his former employer after getting laid off. Or a woman asking about the effects of folic acid for a woman in her 60s who has already gone through menopause. Or someone asking for help with their Blue Cross health insurance bill.
{{< /fig-quote >}}

なんてな情報が公開情報になってるそうで「わざとやってるの？」「ホンマにええのん？」みたいな懸念を述べる記事になっている。

こうした状況について Mozilla が

{{< fig-quote type="markdown" title="Meta: Shut Down Your Invasive AI Discover Feed. Now. - Mozilla Foundation" link="https://www.mozillafoundation.org/en/campaigns/meta-shut-down-your-invasive-ai-discover-feed-now/" lang="en" >}}
Meta is quietly turning private AI chats into public content — and too many people don’t realize it’s happening.
{{< /fig-quote >}}

などと警告を発しており Meta に対して以下の要求を行っている。

{{< fig-quote type="markdown" title="Meta: Shut Down Your Invasive AI Discover Feed. Now. - Mozilla Foundation" link="https://www.mozillafoundation.org/en/campaigns/meta-shut-down-your-invasive-ai-discover-feed-now/" lang="en" >}}
1. **Shut down the Discover feed** until real privacy protections are in place.
1. **Make all AI interactions private by default** with no public sharing option unless explicitly enabled through informed consent.
1. **Provide full transparency** about how many users have unknowingly shared private information.
1. **Create a universal, easy-to-use opt-out system** for all Meta platforms that prevents user data from being used for AI training.
1. **Notify all users whose conversations may have been made public**, and allow them to delete their content permanently.
{{< /fig-quote >}}

この辺は Meta に限らない話で（Meta みたいにあけすけでないだけで）いま世の中に溢れる数多の生成 AI および AI アシスタントサービスでも同様の問題があると思う。
まぁ，私はそもそも [AI サービスを信用してない]({{< ref "/remark/2025/06/problem-with-ai-as-a-service.md" >}} "それは「AI」ではなく「サービスとしての AI」の問題")ので GitHub Copilot や [Kagi Assistant] で必要最小限の[利用]({{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}} "Kagi Assistant が全ユーザに解放")という感じであるが。

それにしても，こういった話がまたしてもトランプ政権下で出てくるというのは象徴的ではある。
当時の CA みたいな話に発展しないことを祈るばかりである。

## ブックマーク

- [Mozilla Criticizes Meta's 'Invasive' Feed of Users' AI Prompts, Demands Its Shutdown - Slashdot](https://tech.slashdot.org/story/25/06/08/1929242/mozilla-criticizes-metas-invasive-feed-of-users-ai-prompts-demands-its-shutdown)

[Kagi Assistant]: https://kagi.com/assistant "The Assistant"

## 参考図書

{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
{{% review-paapi "B0125TZSZ0" %}} <!-- つながりっぱなしの日常を生きる -->
{{% review-paapi "B0C9Z7KGRN" %}} <!-- はじめて学ぶ ビデオゲームの心理学 Kindle 版 -->
