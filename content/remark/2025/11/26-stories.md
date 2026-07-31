+++
title = "諸々の雑感（2025-11-26 の戯れ言）"
date =  "2025-11-26T14:58:25+09:00"
description = "Orion ブラウザがバージョン 1.0 に到達 / Firefox 145 で追加されたプライバシー保護関連機能 / 𝕏 プロファイルの所在地をボカしてみる / KeePassXC 2.7.9 が ANSSI のセキュリティビザを取得"
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "web", "firefox", "twitter", "security", "privacy", "risk", "management", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

## Orion ブラウザがバージョン 1.0 に到達

- [Orion 1.0 ✴︎ Browse Beyond | Kagi Blog](https://blog.kagi.com/orion)

[Orion] ブラウザは WebKit ベースの Web ブラウザで，ゼロテレメトリ，ユーザのプライバシー優先の設計を謳っている。

{{< fig-quote type="markdown" title="Orion 1.0 ✴︎ Browse Beyond" link="https://blog.kagi.com/orion" lang="en" >}}
- A lean, native codebase without ad‑tech bloat.
- Optimized startup, tab switching, and page rendering.
- A UI that gets out of your way and gives you more screen real estate for content.
- **Zero Telemetry**: We don’t collect usage data. No analytics, no identifiers, no tracking.
- **No ad or tracking technology** baked in: Orion is not funded by ads, so there is no incentive to follow you around the web.
- **Built‑in protections**: Strong content blocking and privacy defaults from the first launch.
{{< /fig-quote >}}

生成 AI への態度も明確で[^llm1]

[^llm1]: この件については拙文「[「LLM は詭弁家である」]({{< ref "/remark/2025/11/llms-are-sophists.md" >}})」も参考にどうぞ。

{{< fig-quote type="markdown" title="Orion 1.0 ✴︎ Browse Beyond" link="https://blog.kagi.com/orion" lang="en" >}}
- We are not against AI, and we are conscious of [its limitations](https://blog.kagi.com/llms). We already integrate with AI‑powered services wherever it makes functional sense and will continue to expand those capabilities.
- **We are against rushing insecure, always‑on agents into the browser core.** Your browser should be a secure gateway, not an unvetted co‑pilot wired into everything you do.
{{< /fig-quote >}}

とした上で，実装としては

{{< fig-quote type="markdown" title="Orion 1.0 ✴︎ Browse Beyond" link="https://blog.kagi.com/orion" lang="en" >}}
- Orion ships with **no built‑in AI code** in its core.
- We focus on providing a clean, predictable environment, **especially for enterprises and privacy‑conscious professionals**.
- Orion is designed to connect seamlessly to the AI tools you choose – soon including Kagi’s intelligent features – while keeping a clear separation between your browser and any external AI agents.
{{< /fig-quote >}}

ということだそうだ。

個人的には Web ブラウザのコア機能として AI を組み込むのは時期尚早だと思っているので，上述の実装方針には共感する。
とはいえ AI 機能自体は便利に使っていて，たとえば Firfox 組み込みの AI 機能は使ってないけど “[Kagi Search for Firefox](https://addons.mozilla.org/en-US/firefox/addon/kagi-search-for-firefox/ "Kagi Search for Firefox – Get this Extension for 🦊 Firefox (en-US)")” (表示ページの内容を要約してくれる) や “[Kagi Translate](https://addons.mozilla.org/en-US/firefox/addon/kagi-translate/ "Kagi Translate – Get this Extension for 🦊 Firefox (en-US)")” といった拡張機能はかなり積極的に使っている（お金払ってるんだから使わにゃ損）。

[Orion] ブラウザは現在 macOS, iOS/iPadOS 版が提供されていて Linux 版はアルファ版がある。
Windows 版は開発中とのこと。
私は外出用の [MacBook Air]({{< ref "/remark/2025/01/kubuntu-on-macbook-air-m1.md" >}} "MacBook Air M1 に Kubuntu を入れる") には既に入れているのだが macOS の機能は殆ど使ってないからなぁ。

Linux 版のアルファが取れる日をお待ちしています。

## Firefox 145 で追加されたプライバシー保護関連機能

- [Firefox expands fingerprint protections: advancing towards a more private web](https://blog.mozilla.org/en/firefox/fingerprinting-protections/)

記事を読んでも[ちょっと何言ってるか分からない]のだが（自画自賛？）

{{< fig-quote type="markdown" title="Firefox expands fingerprint protections" link="https://blog.mozilla.org/en/firefox/fingerprinting-protections/" lang="en" >}}
Firefox’s new protections are a balance of disrupting fingerprinters while maintaining web usability. More aggressive fingerprinting blocking might sound better, but is guaranteed to break legitimate website features. For instance, calendar, scheduling, and conferencing tools legitimately need your real time zone. Firefox’s approach is to target the most leaky fingerprinting vectors (the tricks and scripts used by trackers) while preserving functionality many sites need to work normally. The end result is a set of layered defenses that significantly reduce tracking without downgrading your browsing experience.
{{< /fig-quote >}}

と書かれている。
Firefox を使ってる人は，この機会に「プライバシーとセキュリティ」の設定を見直してみるといいかもしれない。

Firefox のこういうところは悪くないと思うのだが，イマイチ[信用しきれん]({{< ref "/remark/2025/11/what-is-an-ai-driven-browser.md" >}} "そもそも「AI ブラウザ」ってなんなん？")のだよなぁ。

## {{% emoji "X" %}} プロファイルの所在地をボカしてみる

最初は NHK の {{% emoji "X" %}} アカウントが US になってるという話を聞いて，これこそ「[ちょっと何言ってるか分からない]」だったが，どうもプロファイルに表示されるようになった「アカウントの所在地」の話らしい。

{{< fig-img-quote src="./x-profile.png" title="プロファイル" link="https://x.com/spiegel_2007" width="595" >}}

ちっ。私は日本在住か。
いや，合ってるけど。

これに端を発して，なかなかに愉快なことになってるようだ。

- [Elon Musk’s Worthless, Poisoned Hall of Mirrors - The Atlantic](https://www.theatlantic.com/technology/2025/11/x-about-this-account/685042/)

あまり真面目に読む気がなかったので [Kagi Assistant] に要約させたのだが

{{< div-ai type="markdown" >}}
- 機能の説明：「このアカウントについて」はアカウント作成国やユーザー名の変更履歴などを表示する新機能です。
- 発見された事実：アメリカの愛国者を名乗る多くのプロMAGAアカウントが、実際にはナイジェリア、ロシア、インド、タイなど外国を拠点にしていることが明らかになりました。
- 表示の問題：ただし、この機能には誤表示や不正確なラベル付けも確認されており、すべての情報が信頼できるわけではありません。
- 著者の主張：筆者は、X（旧Twitter）がエンゲージメントと利益を優先することで、政治的分断や偽情報を増幅し、プラットフォームを「価値のない、毒された万華鏡」のような状態に変えていると論じています。
- 結論：この事例は、真実や安全性よりも利益を優先するソーシャルメディア全体に共通するより大きな問題の一端であると結ばれています。
{{< /div-ai >}}

これで合ってる？

間違いや不正確な情報を元にどれだけ騒ぎ立てても1ミリも説得力がないと思うのだが，やっぱ {{% emoji "X" %}} は碌でもないということは分かる。

{{% emoji "X" %}} のリージョン・チェックがどうなってるか知らないが（壊れてるようにしか見えないけど），表示を緩和させることはできるらしい（完全に隠すことはできない）。

- [How To Hide Your Country Location on X (Twitter) by Switching to Region – Hackread – Cybersecurity News, Data Breaches, Tech, AI, Crypto and More](https://hackread.com/how-to-hide-x-twitter-location-switch-region/)

これによると「設定とプライバシー」から「プライバシーと安全」を選択し，更に「アカウントについて」を選択すると以下の画面が表示される。

{{< fig-img-quote src="./x-setting.png" title="プロファイル" link="https://x.com/settings/about_your_account" width="1041" >}}

ここで「地域/大陸を使用する」を選択すると

{{< fig-img-quote src="./x-profile-2.png" title="プロファイル" link="https://x.com/spiegel_2007" width="595" >}}

おおっ。
所在地が “East Asia & Pacific” になった。
まぁ，大陸レベルで誤魔化したい場合（あるいは大陸を跨いで間違えてる場合）は，これでもダメだろうけど。

なんでこんなこと始めたのかね。
意味のない情報だよな。
GitHub みたいに時差を表示してくれるんならまだ分かるけど。
あれか。
米国とそれ以外を区別しようとしてるのかな？ かの国は「[アメリカ・ファースト](https://en.wikipedia.org/wiki/America_First "America First - Wikipedia")」らしいし。
でもそれで間違った情報をばらまいてるならしょうがないな（笑）

ちなみに私の {{% emoji "X" %}} アカウントは認証されてません。
お金を払ってない（払う気がない）ので。

## KeePassXC 2.7.9 が ANSSI のセキュリティビザを取得

[KeePassXC] 2.7.11 がリリースされた。

- [KeePassXC 2.7.11 released – KeePassXC](https://keepassxc.org/blog/2025-11-23-2.7.11-released/)

この中で， Windows 版 KeePassXC 2.7.9 がフランスの国立情報システムセキュリティー庁（Agence Nationale de la Sécurité des Systèmes d’Information; ANSSI）の第一級セキュリティ認証（Certification de Sécurité de Premier Niveau; CSPN）に合格し，セキュリティビザを取得したと書かれている。

{{< fig-img-quote src="./anssi-security-visa.png" title="KeePassXC Security Audits and Certifications – KeePassXC" link="https://keepassxc.org/audits/#cspn-2025-11-17" width="1227" lang="en" >}}

このセキュリティビザ取得により [KeePassXC] は，今後3年間 ANSSI の後援を受けられるそうな。

ANSSI はフランスにおいて暗号製品の政府調達にも関わってるセクションらしい。
以下のドキュメントにフランスの暗号関連の認証・調達に関する解説が載っている。

- [「暗号利用環境に関する動向調査」報告書（2015年度） | 情報セキュリティ | IPA 独立行政法人 情報処理推進機構](https://www.ipa.go.jp/security/reports/crypto/crypto_survey.html)

また，セキュリティビザに関しては以下をどうぞ。

- [Security VISA | ANSSI](https://cyber.gouv.fr/en/security-visa)


## ブックマーク

- [Linux その288 - KeePassXC 2.7.11 リリース・KeePassXC をインストールするには - kledgeb](https://kledgeb.blogspot.com/2025/11/linux-288-keepassxc-2711-keepassxc.html)

[Orion]: https://orionbrowser.com/ "Orion Browser by Kagi"
[ちょっと何言ってるか分からない]: https://dic.pixiv.net/a/%E3%81%A1%E3%82%87%E3%81%A3%E3%81%A8%E4%BD%95%E8%A8%80%E3%81%A3%E3%81%A6%E3%82%8B%E3%81%8B%E5%88%86%E3%81%8B%E3%82%89%E3%81%AA%E3%81%84 "ちょっと何言ってるか分からない (ちょっとなにいってるかわからない)とは【ピクシブ百科事典】"
[Kagi Assistant]: {{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}} "Kagi Assistant が全ユーザに解放"
[KeePassXC]: https://keepassxc.org/ "KeePassXC Password Manager"
