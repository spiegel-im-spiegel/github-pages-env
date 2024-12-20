+++
title = "Bluesky 上で起きたなりすましとサイバースクワッティング"
date =  "2024-12-20T22:16:04+09:00"
description = "Bluesky アカウントのために新たにドメインを取得しても本人証明としては大して役に立っていない"
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "bluesky" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

興味深い話なので覚え書きとして記しておく。

- [So, Bluesky Has An Extortion Problem](https://tedium.co/2024/12/17/bluesky-impersonation-risks/)
- [Blueskyの「ドメインを用いた本人証明システム」が悪用され「著名人の名前でドメインを取得して本物に買取りを要求する」という事件が発生 - GIGAZINE](https://gigazine.net/news/20241219-bluesky-handle-domain-name-extortion/)

まず前提として， Bluesky ではユーザ本人の証明として実在のドメイン名を[ハンドル名として紐付け](https://blueskyweb.zendesk.com/hc/en-us/articles/19001802873101-How-to-Set-your-Domain-as-your-Handle "How to Set your Domain as your Handle – Bluesky")ることができる。
例えば私の場合はこんな感じ。

{{< fig-img-quote src="./bluesky-profile.png" title="Spiegel＠旧Twitter難民 (@baldanders.info) — Bluesky" link="https://bsky.app/profile/baldanders.info" width="602" >}}

今回のケースでは，誰かが Bloomberg のコラムニスト [Conor Sen] 氏の名前を使ったドメイン名 `conorsen.com` を取得し，これを使って Bluesky アカウントを作成したことから始まる。
ちなみに [Conor Sen] 氏本人は `@conorsen.bsky.social` というアカウントを運用している。

偽アカウント作成者は [Conor Sen] 氏に対し，このドメインと Bluesky アカウントを買い取るよう要求してきた。
まぁ，ここまではよくあるサイバースクワッティング（cybersquatting）の事例である。

ここからが面白くて，さらに [The Hustle の founder である Sam Parr][Sam Parr] 氏（有名らしい，私は知らんけど）の偽 Bluesky アカウントが登場する。
この偽物氏は上述の脅迫者に同調し，オンライン上の身元を気にするなら [Conor Sen] 氏は金を払うべきだとか言ったらしい。
こうして偽 Sam Parr 氏と他ユーザとのやり合いが始まった。

さらにさらに，ここで本物の [Sam Parr] 氏が Bluesky に登場。
偽物氏は本物の [Sam Parr] 氏にドメインとアカウントを買い取るよう要求する。

このように，有名人の複数偽 Bluesky アカウントを作成・連動させ本物達からお金を巻き上げる手口らしい。
偽アカウントは独自ドメインで紐付けされているため本物より本物らしく見えるところがポイントである。

本物氏は当然ながらこれに抗議して Bluesky 側に偽アカウントをブロックするよう要請した。
そしてどうなったかというと Bluesky 側は本物のアカウントの方をブロックしちゃったらしい。
独自ドメインで紐付けされてる偽物を信用しちゃったわけだ。

{{< div-gen class="center" type="markdown" >}}
**[なぁにぃ？ やっちまったなぁ！](https://www.youtube.com/watch?v=RWcMDnWCwxQ)**
{{< /div-gen >}}

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}。

既に十分なブランド力を持つドメインを Bluesky アカウントと連携させるのは，上手いやり方だろう。
たとえば，私がフォローしている河出書房新社（[`@kawade.co.jp`](https://bsky.app/profile/kawade.co.jp "河出書房新社 (@kawade.co.jp) — Bluesky")）や早川書房（[`@hayakawa-online.co.jp`](https://bsky.app/profile/hayakawa-online.co.jp "早川書房公式 (@hayakawa-online.co.jp) — Bluesky")）などは疑いようもない。
また hololive 所属の VTuber である「[大神ミオ](https://www.youtube.com/@OokamiMio "Mio Channel 大神ミオ - YouTube")」は `hololive.tv` のサブドメインを使って [`@ookamimio.hololive.tv`](https://bsky.app/profile/ookamimio.hololive.tv "大神ミオ@ホロライブゲーマーズ (@ookamimio.hololive.tv) — Bluesky") としている。
企業所属のユーザなんかは参考になると思う。

でも Bluesky アカウントのために新たにドメインを取得しても本人証明としては大して役に立っていない，ということが今回のケースを通して改めて分かった。
なりすましやサイバースクワッティングに対する抑止力になっていないのだ。
それとも Bluesky が詐欺や脅迫の現場となるまでに成長したと褒めるべきだろうか（笑）

特に SNS や（{{% emoji "X" %}} のような）マイクロブログサービスや YouTube/Instagram/TikTok などで囲い込まれて自前ドメインの Web サイトもないユーザは他サービスとの連携が薄い。
例えば「{{% emoji "X" %}} はアカン」と思って Bluesky に行っても {{% emoji "X" %}} のアカウントと Bluesky アカウントが同一ユーザであることを証明する方法はない。
だから今回のようなケースも起こり得るわけだ。

SNS における「ネットワーク効果」によって他サービスへの移行が難しくなるという話はよく聞くが，移行前後でアイデンティティの同一性を保てなくなるというのもスイッチング・コストのひとつなのかも知れない。

Mastodon は Web ページ上の `<link>` 要素や `<a>` 要素に `rel="me"` の属性を付与することでサービス間で緩く連携できる。

{{< fig-img-quote src="./mastodon-profile.png" title="Spiegel＠がんばらない (@spiegel@goark.fedicity.net) - Goark" link="https://goark.fedicity.net/@spiegel" width="602" >}}

でもこれも他の SNS などとはまず連携できない。
{{% emoji "X" %}} から移住してきても同一ユーザであることを証明することはできないのだ。

最初に挙げた [Tedium の記事](https://tedium.co/2024/12/17/bluesky-impersonation-risks/ "So, Bluesky Has An Extortion Problem")では

{{< fig-quote type="markdown" title="So, Bluesky Has An Extortion Problem" link="https://tedium.co/2024/12/17/bluesky-impersonation-risks/" lang="en" >}}
It’s a goddamn mess, and it makes me appreciate why some people may want to skip Bluesky altogether. I’ve been a [booster](https://tedium.co/2024/11/14/bluesky-network-growth-analysis/) of the [network](https://tedium.co/2024/12/05/bluesky-business-models-craigslist/) so far—but they need to get this figured out, or all the prominent people who have put their stake over here may find themselves looking for the exits.
{{< /fig-quote >}}

と締めているが，なかなか厳しいんじゃないかなぁ。

一応 Bluesky 側は先月末の時点でモデレーションチームを4倍（100人）に増やしたそうで，なりすましやサイバースクワッティングに対しても積極的に対処するとのこと。

- [Blueskyがなりすましアカウントやパロディアカウントに対する方針をアップデート、偽物であることが明記されていない場合は削除 - GIGAZINE](https://gigazine.net/news/20241201-bluesky-promises-more-verification-impersonation/)

冤罪の BAN も増えそうな気がするが，どうなることやら。

[Conor Sen]: https://www.bloomberg.com/authors/ATA2uaN4m4A/conor-sen "Conor Sen - Bloomberg"
[Sam Parr]: https://thehustle.co/originals/author/sam-parr "Stories - The Hustle | Sam Parr"
