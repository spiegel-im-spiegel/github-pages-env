+++
title = "Signal への攻撃手口（ソーシャルエンジニアリング）と防御策"
date =  "2026-04-28T13:12:54+09:00"
description = "Signalへの攻撃はハックではなく，偽サポート型フィッシングだった。報道と公式声明をもとに手口と対策を整理する。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "security", "risk", "management", "media", "signal", "phishing", "messaging", "authentication" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

今朝，Bluesky と Mastodon の TL に Signal による長いスレッドが上がっていた。
Bluesky のポストは以下から始まるスレッドである（Mastodon のほうも内容は同じ，多分）。

{{< fig-gen >}}
<blockquote class="bluesky-embed" data-bluesky-uri="at://did:plc:tovhrfml47r6u4idaitaxefk/app.bsky.feed.post/3mkixc3wypc2u" data-bluesky-cid="bafyreicp3qwxcelnd4jqovdw6qp4vniz22tf7wrny3wr53vam5uhgim6ui" data-bluesky-embed-color-mode="system"><p lang="en">A response to recent reporting in Germany, in service of clarity and accountability: 

First, it’s important to be precise when it comes to critical infrastructure like Signal. Signal was not “hacked” — in that our encryption, infrastructure, &amp; the integrity of the app’s code was not compromised. 1/</p>&mdash; Signal (<a href="https://bsky.app/profile/did:plc:tovhrfml47r6u4idaitaxefk?ref_src=embed">@signal.org</a>) <a href="https://bsky.app/profile/did:plc:tovhrfml47r6u4idaitaxefk/post/3mkixc3wypc2u?ref_src=embed">April 28, 2026 at 5:53 AM</a></blockquote><script async src="https://embed.bsky.app/static/embed.js" charset="utf-8"></script>
{{< /fig-gen >}}

[ちょっと何言ってるか分からなかった](https://dic.pixiv.net/a/%E3%81%A1%E3%82%87%E3%81%A3%E3%81%A8%E4%BD%95%E8%A8%80%E3%81%A3%E3%81%A6%E3%82%8B%E3%81%8B%E5%88%86%E3%81%8B%E3%82%89%E3%81%AA%E3%81%84 "ちょっと何言ってるか分からない (ちょっとなにいってるかわからない)とは【ピクシブ百科事典】")ので [Kagi Assistant] と [GitHub Copilot] に手伝ってもらいながら調べてみた。

## もとになってる報道

[Signal] に対する攻撃，特に phishing を含むソーシャルエンジニアリングについては以前から言われていて，ドイツ BSI (Bundesamt für Sicherheit in der Informationstechnik) も [Signal] の利用ガイドラインを公開しているほどだ。

- [Germany warns of Signal account attacks targeting high-profile figures](https://cyberinsider.com/germany-warns-of-signal-account-attacks-targeting-high-profile-figures/)
- [Signal、国家支援型の巧妙なアカウント乗っ取り攻撃が標的に―ドイツ当局が警告](https://innovatopia.jp/cyber-security/cyber-security-news/79762/)
- [BSI - Handlungsleitfaden bei Phishing über den Signal Support](https://www.bsi.bund.de/DE/Themen/Unternehmen-und-Organisationen/Informationen-und-Empfehlungen/Empfehlungen-nach-Angriffszielen/Signal-Support/signal-support_node.html)

にもかかわらずってことなんだろうなぁ。

[Signal] による声明の直接のトリガーとなっているのは先週のドイツでの一連の報道のようだ。
発端は DER SPIEGEL の報道：

- [Julia Klöckner ist Opfer des Signal-Hacks - DER SPIEGEL](https://www.spiegel.de/politik/deutschland/phishing-alarm-in-berliner-regierungsviertel-julia-kloeckner-opfer-des-signal-hacks-a-7f5fc795-d0c2-4325-b726-4109531270bc)

これは有料記事だが，他メディアでもいくつかあった。

- [Was der Angriff auf Signal-Konten für den Bundestag bedeutet | tagesschau.de](https://www.tagesschau.de/inland/kloeckner-signal-phishing-100.html)
  - [Phishing auf Signal: Bundesanwaltschaft ermittelt wegen Spionageverdachts | tagesschau.de](https://www.tagesschau.de/inland/innenpolitik/ermittlungen-bundesanwaltschaft-signal-100.html)
  - [Phishing auf Signal: Bundesregierung von Angriff betroffen | tagesschau.de](https://www.tagesschau.de/inland/signal-ausspaehung-bundesregierung-100.html)
- [Phishing via Signal: Attacke auf Ministerinnen | taz.de](https://taz.de/Angriff-via-Signal/!6174144/)
- [Phishing-Angriffe auf Signal: Bundesregierung sieht offenbar Russland hinter Phishing bei Signal | DIE ZEIT](https://www.zeit.de/digital/2026-04/signal-phishing-bundestag-russland-abgeordnete-gxe)

この中では [Tagesschau] の一連の記事が良さげなので [Kagi Assistant] に要約してもらった。

{{< div-ai type="markdown" >}}
ドイツの政界や重要機関を標的としたSignalでの大規模なフィッシング攻撃は，国家が関与するスパイ工作の疑いへと発展しています。

- **被害の規模と対象**:
    - ドイツ連邦議会のほぼすべての会派の議員が標的となり，ユリア・クレックナー連邦議会議長や，フーベルツ建設相，プリーン家族相といった現職閣僚の被害も確認されています。
    - 政治家だけでなく，ジャーナリスト，軍関係者，さらにはNATO関係者も攻撃対象に含まれています。
- **攻撃の手法**:
    - 攻撃者は「Signalサポート」を装い，偽のセキュリティ警告でユーザーに接触します。
    - ユーザーを騙して**PINコード**を聞き出しアカウントを完全に奪取するか，あるいは**QRコード**をスキャンさせて攻撃者のデバイスを連携（同期）させることで，チャット内容や連絡先，過去45日分のファイルを盗み見ます。
- **当局の捜査と背後関係**:
    - ドイツ連邦検察庁は，本件を「諜報員活動（スパイ容疑）」として正式に捜査を開始しました。
    - ドイツ治安当局（BfVおよびBSI）は，このキャンペーンが「国家主導のサイバー攻撃者」によるものであり，現在も勢いを増していると警告しています。
    - オランダの諜報機関や米国のFBIは，この攻撃の背後に**ロシアの諜報機関**が関与していると特定しています。
- **Signal側の見解**:
    - Signalは，アプリの暗号化やインフラ自体が突破されたわけではなく，あくまでユーザーを欺くソーシャルエンジニアリングによる攻撃であると強調しています。
{{< /div-ai >}}

大体合ってるかなぁ。
ドイツ語はよく分からないので自信がない。

[GitHub Copilot] が各メディアの報道の違いを表にまとめてくれるというのでやってもらった。

{{< div-ai type="markdown" >}}
| 媒体 | 位置づけ | 主なポイント | 閲覧性 |
|---|---|---|---|
| DER SPIEGEL | 初報の可能性が高い | 政治家アカウント被害の最初期報道。見出しで「Signal-Hack」と表現 | 本文はSPIEGEL+（有料） |
| Tagesschau | 検証・整理記事 | 「最初にSPIEGELが報じた」と明記。ハックではなくフィッシングという整理が明確 | 無料で読みやすい |
| taz | 補強報道 | 被害対象，政府見解，手口（偽Signal Support）を具体化。SPIEGEL初報への言及あり | 無料 |
| ZEIT | 続報 | 政府・当局側の見方（ロシア関与の可能性）を短く整理 | 一部制限あり（記事により差） |
{{< /div-ai >}}

## Signal による声明

というわけで [Signal] による SNS での声明に戻る。
これについても [Kagi Assistant] に要約してもらった。

{{< div-ai type="markdown" >}}
- **攻撃の性質**: Signalの暗号化やインフラ自体が「ハッキング」されたわけではありません。実際には，攻撃者が「Signalサポート」を装い，ソーシャルエンジニアリングを用いてユーザーから認証情報を聞き出す巧妙なフィッシング詐欺が発生しています。
- **手口の詳細**: 攻撃者は盗んだ情報でアカウントを乗っ取り，登録電話番号を変更します。被害者には「再登録が必要」と信じ込ませて新しいアカウントを作らせることで，乗っ取りに気づかせないようにします。その後，乗っ取ったアカウントを使って被害者の連絡先にさらなる攻撃を仕掛けます。
- **Signal側の対応**: 今後数週間で，こうした攻撃を阻止するための新機能を導入予定です。
- **ユーザーへの推奨事項**:
    - Signalサポートがメッセージリクエストを送ったり，認証コードやPINを尋ねたりすることはないため，警戒を怠らないこと。
    - 保護を強化するため，設定から「登録ロック」を有効にすること。
{{< /div-ai >}}

んー。
これは大体合ってるはず。

「登録ロック」はアプリの設定の「アカウント」セクションにある項目で，電話番号を変更する際に PIN 入力を要求するらしい。

{{< fig-img-quote src="./signal-account.jpg" title="Signal アカウント設定" link="./signal-account.jpg" >}}

PIN をばらしちゃったらダメだけどね。
ちなみに [Signal] の PIN は英字も含めることができる。

Phishing はシステム侵入の手口としてはありふれていて，しかも人間の行動が関与するから完全には防ぎづらい。
故に多層防衛をしていかないと「蟻の一穴」から簡単に class break してしまうことになる。

[GitHub Copilot] はこうまとめてくれた。

{{< div-ai type="markdown" >}}
本件は，セキュアな暗号技術がそのまま利用者体験の安全を保証するわけではないことを示した事例である。攻撃者は最も弱い層，すなわち人間の心理と手続きを狙う。ゆえに防御も，暗号強度だけでなく，認証情報を渡さない運用，異常時の確認手順，設定レベルの予防策を組み合わせて成立するのである。
{{< /div-ai >}}

全く以ってそのとおり！

## 余談だが...

最初の SNS のポストの内容を元に [Kagi Assistant] と [GitHub Copilot] に元ネタとなるドイツ報道を調べてもらった。
[Kagi Assistant] のほうは自前の検索サービスがあり，一発で [Tagesschau] の記事を引き当てたので問題なかったが， [GitHub Copilot] はだいぶ難儀していた。

まず DuckDuckGo で調べようとして bot 判定され拒否られる。
そこで Bing に切り替えたのだがノイズが多すぎて見つからなかったようだ（「検索エンジンが使いづらい」と愚痴ってた。 AI も愚痴るのかw）。
まぁ， Google を使わなかっただけ偉いと言っておこう（笑）

そんで検索サービスを彷徨うのは一旦諦めて Bluesky の該当スレッドを API 経由で取得し，そこからようやく「Julia Klöckner（独連邦議会議長の人）」というキーワードを引き当てた。
これを使って Bing で絞り込みをかけた結果，意味のある検索結果を得られたという流れだった。

この [GitHub Copilot] による一連の作業をチャット上で眺めながら，人間みたいなことをするんだなぁ，と妙なところで感心してしまった。
そんでもって，ネット上の調べ物は [Kagi Assistant] のもんやねぇ，と思う。

そもそも調べ物をするのに検索窓を叩いて調べるとかせんやろ，いまどきは。
ノイズに塗れた Web で，かつ主要検索サービスがメタクソ化する状況で「行きたい場所」に到達するには，結局 AI を使わなきゃ無理なのよ。
特に今回みたいな漠然とした内容で知らない言語圏の情報を探す場合は。

つまり，もはや AI に認知してもらえないサイトやページは存在しないのと同じで，安易に拒否してるだけでは大海に取り残された某島の生物のようになってしまうだろう。

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"
[GitHub Copilot]: https://github.com/features/copilot "GitHub Copilot · Your AI pair programmer · GitHub"
[Signal]: https://signal.org/ "Signal"
[Tagesschau]: https://www.tagesschau.de/ "tagesschau.de - die erste Adresse für Nachrichten und Information | tagesschau.de"

## 参考

{{% review-paapi "4822283100" %}} <!-- セキュリティはなぜやぶられたのか -->
