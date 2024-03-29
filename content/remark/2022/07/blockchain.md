+++
title = "Blockchain/Bitcoin は何だったのか"
date =  "2022-07-25T21:56:45+09:00"
description = "技術で社会は変えられない"
image = "/images/attention/kitten.jpg"
tags = [ "code", "blockchain", "bitcoin", "grigori", "internet" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

Twitter で小耳に挟んだ話なのだが，某出版社から出る（出た）いわゆる Web3 本がかなり愉快な内容らしく，[回収](https://book.impress.co.jp/info/20220725.html "書籍「いちばんやさしいWeb3の教本 人気講師が教えるNFT、DAO、DeFiが織りなす新世界」の回収について - インプレスブックス")するとか何とか。
勿体ない。
論文とかなら出し直しもやむなしだろうが，その程度の本なら回収までしなくても「緊急改訂」とかでいいと思うのだが。
それも「歴史」だよ。
本を書いた方は後年黒歴史として悶絶してしまうかもしれないが（笑）

私自身は NFT にも Web3 にも[大して興味はない]({{< ref "/remark/2022/06/this-is-not-web.md" >}} "Web3 ってゆーな！")が，折角なので Blockchain/Bitcoin について色々と放言しておこう。
書籍出版で許されるんだから辺境のブログで書くくらい構わんじゃろ。

{{< div-box type="markdown" class="center" >}}
*この記事については，いつも以上に内容の正しさを保証しません。
<br>マジに受け取らないでね（予防線）*
{{< /div-box >}}

## 脱中央集権「社会」など来ない

「黄色矮星人は2人いれば力比べを始め，3人いれば派閥を作る[^nh1]」

[^nh1]: 「黄色矮星人は2人いれば力比べを始める」というのはSF作家の野尻抱介さんの著作でよく出てくるフレーズ（多分）。後半の「3人いれば派閥を作る」は誰が言ったか思い出せない。うろ覚え。

Blockchain 技術を含むプロダクトについて語られるとき，必ずと言っていいほどよく出てくるフレーズが「非中央集権」あるいは「脱中央集権」である。

しかし，考えてみてほしい。
そもそも「インターネット」は脱中央集権的な思想で設計されたものなのだ。

{{< fig-quote type="markdown" title="介入されないもうひとつのウェブ" link="http://www.nikkei-science.com/201206_074.html" >}}
もし今日のインターネットが実際にこの理論に近い状態であれば，メッシュネットワークは余計ものだったろう。
だがインターネットが当初の学術目的から踏み出して現在のような誰でも使える商業サービスになってから20年以上が経つうちに，そうした蓄積伝送の原理が果たす役割は，一貫して縮小していった。

　この間，ネットワークに加わる新たなノードの圧倒的多数はISPを介してネットに接続する家庭や企業のコンピューターだった。
ISPの接続モデルでは，利用者のコンピューターはデータの中継はしない。
それはネットワークの端末，つまりデータの送受信だけを，常にISPのコンピューターを介して行うターミナル・ノードだ。
言い換えれば，インターネットの爆発的な成長はネットワーク地図に行き止まりのルートを増やしただけで，新たなルートを加えることはほとんどなかった。

　そしてISPなど大量の情報ルートを持つ者は，彼らがルートを提供している何百万ものノードを支配下におくこととなった。
これらのノードは，もしISPがダウンしたり，ネットから遮断されたりすると，その障害を回避する方法がない。
ISPはインターネットが停止しないようにするどころか，実効上は停止スイッチになってしまった。
{{< /fig-quote >}}

そして Blockchain/Bitcoin もインターネットと同じ末路を辿る。

## Bitcoin でスーパーの醤油は買えない

Bitcoin がシステムとして成立する条件は

1. 不特定多数が参加
2. 全てのユーザが唯一の「台帳」の完全なコピーを持つ
3. ユーザ同士の peer な接続
4. PoW への平等参加

あたりだろう。
これをぶち壊したのが「交換所」の存在である。

最初の頃，私は Bitcoin を「デフレ型補完通貨」と定義していた。
補完通貨はゼロ年代に「地域通貨」と呼ばれていたことがある。
流行ったよねぇ，[地域通貨](https://cruel.org/krugman/babysitj.html "経済を子守りしてみると。")（笑）

また Bitcoin は最終的な「総量」が決まっているため，必然的に「デフレ型」となる[^hf1]。

[^hf1]: Bitcoin の総量が決まっているという前提は2017年の hard fork で崩れた。

しかし実際には Bitcoin 自体は補完通貨にならなかった。
昔は Bitcoin で決済できる小売店とかあったと聞くが，今はそんなお店はないだろう。

Bitcoin を日常生活で使うためには法定通貨に「両替」する必要がある。
故に「交換所」の登場は必然だった。

その結果なにが起こったかというと，交換所を利用する多くのユーザは上述の4つの条件を満たす player ではなくなってしまったのだ。
これは Bitcoin システム全体のアクターが hierarchical に構造化したことを意味する。

インターネットが商業化された際に多くのユーザが ISP にぶら下がる「端末」になり下がり，インターネットに「参加」する player ではなくなった状態とまさに同じである。

## それは「通貨」ではない

さらに言えば，2018年の G20 ブエノスアイレス・サミットの首脳宣言で「暗号資産（crypto-assets）」という表現が出て以来「それら」は通貨ではなく資産として認知されるようになった。
こう考えると NFT (Non-Fungible Token) の登場は寧ろ必然だったと言えるだろう。
ぶっちゃけるなら「デジタル証券バブル」の台頭とでも言おうか（笑）

日本政府が妙に Web3 や NFT を推しているのは「バブル時代よ，もう一度」という願いなのだろうか。

## 技術で社会は変えられない

私自身は与しないが，「交換所」を中心とした hierarchical な構造化も「デジタル証券バブル」も批判するほどのものではない。
しかし，もし Blockchain/Bitcoin が「脱中央集権」的かつ国家から独立した「補完通貨」を目指していたとするなら，この結末は完全に失敗だったと言えるだろう。

結局 Bitcoin などの Blockchain 実装を通じて分かったのは「計算機パワーと情報力の強いものが勝つ」という身も蓋もない話で，それは今までの暴力による民衆統治や経済力による市場支配と何ら変わらない。
そこを見ないで「脱中央集権」が云々とかヘソで茶が沸いてしまう。

{{< fig-quote type="markdown" link="https://wirelesswire.jp/2022/06/82564/" title="Web3の「魂」は何なのか？ – WirelessWire News" >}}
ジャック・ドーシーの批判は、「Web3を所有するのはベンチャーキャピタルとその投資先の企業であってウェブユーザーではなく、結局は中央集権型のものに別のラベルを貼っただけ」と続き、Andreessen Horowitzの共同創業者マーク・アンドリーセンにTwitterでブロックされ、それを自慢するというダメなTwitter芸を見せる[オチがつきました](https://www.itmedia.co.jp/news/articles/2112/24/news077.html)
{{< /fig-quote >}}

{{< fig-quote type="markdown" link="https://wirelesswire.jp/2022/06/82564/" title="Web3の「魂」は何なのか？ – WirelessWire News" >}}
「誰かがブロックチェーンで何かを解決できると言う場合、その人はその『何か』を理解していないので無視してかまわない（ウィーバーのブロックチェーンの鉄則）」
{{< /fig-quote >}}

インターネットにしろ Blockchain/Bitcoin にしろ，脱中央集権的な「システム」を構築することはできるのだろう。
しかし私達が暮らすのは「社会」である。
社会が「脱中央集権」を望まないのであれば，土台がどうあれ，そこに到達することはない。

私は「技術で社会は変えられない」と思っている。
社会を変えることができるとすれば，それはあくまでも「人」の行いである（道具・手段としての技術の意義はあると思うけど）。
インターネットや Blockchain/Bitcoin は奇しくも私のこの妄言を補強する事例となっている。

## ブックマーク

- [ĐApps：Web 3.0はどんなものか（ĐApps: What Web 3.0 Looks Like 日本語訳）](https://www.yamdas.org/column/technique/dappsweb3j.html)
- [ジョン・ハンケが語るWeb3、ティム・バーナーズ＝リーが懐疑的なブロックチェーン - YAMDAS現更新履歴](https://yamdas.hatenablog.com/entry/20220628/john-hanke-on-web3)
- [【夏休みの自由研究】そうだブロックチェーンを作ろう！](https://zenn.dev/koduki/articles/52c207493f4d5e)
- [Letter in Support of Responsible Fintech Policy](https://concerned.tech/)
  - [クリプト・Web3業界の誇大広告に踊らされてはならない：1500人超の科学者・エンジニア・技術者が米議会に警告 | p2ptk[.]org](https://p2ptk.org/monopoly/3760)
- [反クリプト公開書簡に反論する【オピニオン】 | coindesk JAPAN | コインデスク・ジャパン](https://www.coindeskjapan.com/150683/)
- [「脱中央集権」の有益かつ重要な分類――ブロックチェーンを越えて：「脱中央集権」勢は一枚岩ではない | p2ptk[.]org](https://p2ptk.org/monopoly/3771)
- [架空の暗号通貨とMLMを掛け合わせたら何が起こるのか | On Off and Beyond](https://chikawatanabe.com/2022/08/26/the_misding_cryptoqueen/)

## 参考図書

{{% review-aozora "4307" %}} <!-- グリゴリの捕縛 -->
{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "4873118964" %}} <!-- マスタリング・イーサリアム Ethereum -->
