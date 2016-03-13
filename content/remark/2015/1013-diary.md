+++
date = "2015-10-13T22:54:48+09:00"
update = "2016-03-14T01:12:59+09:00"
description = "「ロボット法学会」設立準備研究会 / 『パクリ経済――コピーはイノベーションを加速するか(仮)』 / GnuPG 2.1.9 released / wrap-style がなかなかよい"
draft = false
tags = ["code", "law", "engineering", "artificial-intelligence", "book", "freakonomics", "security", "cryptography", "openpgp", "gnupg", "tools", "atom", "editor"]
title = "今日の戯れ言：ロボット法学会"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  medium = "@spiegel"
  name = "Spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [「ロボット法学会」設立準備研究会]({{< relref "#law" >}})
1. [『パクリ経済――コピーはイノベーションを加速するか(仮)』]({{< relref "#copy" >}})
1. [GnuPG 2.1.9 released]({{< relref "#gnupg" >}})
1. [wrap-style がなかなかよい]({{< relref "#atom" >}})


## 「ロボット法学会」設立準備研究会{#law}

- [「ロボット法学会」設立準備研究会 | Robot Law @ Japan](http://robotlaw.jp/)
- [「ロボット法学会」設立準備研究会 | Peatix](http://peatix.com/event/115206)
- [「法律が普及の足かせになってはいけない」ロボット法学会の設立目指し、準備会開催|弁護士ドットコムニュース](https://www.bengo4.com/internet/n_3805/)
- [「ロボットの社会導入に向けて、法律家も技術者もともに議論を」——「ロボット法学会」の設立準備イベント開催 - WirelessWire News（ワイヤレスワイヤーニュース）](https://wirelesswire.jp/2015/10/46992/)
- [新保史生「何故に『ロボット法』なのか」(2015年10月11日）報告資料 | 「ロボット法学会」設立準備研究会](http://robotlaw.jp/archives/66)
- [ロボット法って何 | IT Research Art](http://www.itresearchart.biz/?p=442)
- [「ロボット法学会設立準備会」に寄せて](http://blogos.com/article/138762/)

最初の2つの記事では「ロボット法 新8原則」を紹介している。
曰く

1. 人間第一の原則
2. 命令服従の原則
3. 秘密保持の原則
4. 利用制限の原則
5. 安全保護の原則
6. 公開・透明性の原則
7. 個人参加の原則
8. 責任の原則

これって既存の「機械」の概念を超えるなと言ってるのと同じだよね。
エンジニア的には面白くはないんだけど，まぁ多分これが大方の人たちの「気分」なんだろう。

あと

{{< fig-quote title="「法律が普及の足かせになってはいけない」ロボット法学会の設立目指し、準備会開催" link="https://www.bengo4.com/internet/n_3805/" >}}
<q>人間はロボットよりも肉体能力が高いけれど、知的能力は結構、負け始めている。頭はロボット、体は人間という仕事の体制ができることになる。つまり、トップの人間がロボットに基本的な命令を出し、ロボットがゴーグルを通じて下級の労働者に命令を出して、人間がその通りに仕事をすることになる。言葉もいらず、すごく安く人間を使える。この危険性を指摘している人は少ない</q>
{{< /fig-quote >}}

とあるが，現時点において人工知能が「知的」かというと首をひねらざるをえない。
たとえば IBM の Watson は，いわゆる「エキスパート・システム」としては現時点での究極と言えるかもしれないが，「知的」ではない。
Watson は「問いを解く」ことに関して人を凌駕しつつあるかもしれないが，「問いを立てる」ことはできないからだ[^1]。

[^1]: まぁ人間だって自分で「問いを立てる」ことのできる者は少なそうだが。特に日本の学校教育は意図的にそういうことを spoil してるからね。できない者はニコラ・テスラに「君は automaton だ」とか言われそう（笑）

自律的に「問いを立てる」ことができるようにならないかぎり，人工知能は危険でも脅威でもない。
でも，もしそれができるようになれば，まさしく「[進化]({{< relref "remark/2015/0917-diary.md#ai" >}})」だし，そうなれば進化の階梯を機械に譲ることになっても仕方ないだろう。

もうひとつ考えるべきは「遠隔操作ロボット」である。
これは問題提起としては昨年末に既に出ている。

- [ロボットを「奴隷的に拘束」したら憲法違反？ 「ロボット法」の可能性を研究者が議論|弁護士ドットコムニュース](https://www.bengo4.com/other/1146/1287/n_2352/)

システムが複雑化すれば，身体的なものであれ，精神的なものであれ，機械と人間との間に軋轢が起きる。
両者の媒として働くと考えられているのも人工知能である。
個人的にはこちらのほうが社会に普及しやすいと考える。

{{< fig-quote title="ロボットを「奴隷的に拘束」したら憲法違反？ 「ロボット法」の可能性を研究者が議論" link="https://www.bengo4.com/other/1146/1287/n_2352/" >}}
<q>身体的なものをロボットが模倣する方向で動いていると思うが、はたして身体になりうるのか。たとえば憲法18条では、身体について、奴隷的な拘束は認められていないが、テレイグジスタンスが高度に発達して、ロボットが拘束された場合は身体性を犯されたことになるのかという疑問がある</q>
{{< /fig-quote >}}

{{< fig-quote title="ロボットを「奴隷的に拘束」したら憲法違反？ 「ロボット法」の可能性を研究者が議論" link="https://www.bengo4.com/other/1146/1287/n_2352/" >}}
<q>また、ロボット法政策研究者でキャンペナーの工藤郁子氏は「製造物責任は『無過失責任』で、ミスがなくても製造者側に責任を負わせるようになっている。しかし、無過失責任を採用する場合、ユーザーにとってメリットはあるが、メーカーにとって重い負担になる」とした上で、「ロボット自体が自分でユーザーとのインタラクションを経て学習すると、もしかしたら技術者が把握している範囲を超えて活動するかもしれない。そのまま製造物責任を導入していいかという問題があるだろう」と語った</q>
{{< /fig-quote >}}

これに直接応えるものではないが「[「ロボット法学会設立準備会」に寄せて](http://blogos.com/article/138762/)」はなかなか興味深い。
この記事では「4つの NEW」を提示している。

1. NEW Machine
1. NEW Relationship
1. NEW Law
1. NEW Generation

これらを軸に議論していくと面白いのではないかと思う。

## 『パクリ経済――コピーはイノベーションを加速するか(仮)』{#copy}

- [この本の帯文は佐野研二郎氏に依頼すればいいのではと思った『パクリ経済――コピーはイノベーションを加速するか(仮)』 - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20151012/pakuri)

うーん。
3,780円。
392ページ。

今年は高めの本は（主に経済的な理由で）遠慮してるのだが[^a]，これは読みたいなぁ。
でもこの前，[泣く泣く本を処分した]({{< relref "remark/2015/0920-diary.md" >}})ばっかりだしなぁ。
これ以上魔窟の進行を許すわけには...

[^a]: 積ん読も溜まりまくってるので消化しないと。

ちうわけで，とりあえず Kindle リクエストを出してみた。

ところで，この「（仮）」ってのは正式タイトルなのだろうか。
タイトルの後ろに「（仮）」って付いてると，昔あった某どどエロアニメを連想してしまうんだが（笑）

## GnuPG 2.1.9 released{#gnupg}

- [GnuPG 2.1.9 released](https://lists.gnupg.org/pipermail/gnupg-announce/2015q4/000380.html)

今回もセキュリティ脆弱性に絡む修正はなし。

- gpg: Allow fetching keys via OpenPGP DANE (--auto-key-locate).  New option --print-dane-records.
- gpg: Fix for a problem with PGP-2 keys in a keyring.
- gpg: Fail with an error instead of a warning if a modern cipher algorithm is used without a MDC.
- agent: New option --pinentry-invisible-char.
- agent: Always do a RSA signature verification after creation.
- agent: Fix a regression in ssh-add-ing Ed25519 keys.
- agent: Fix ssh fingerprint computation for nistp384 and EdDSA.
- agent: Fix crash during passprase entry on some platforms.
- scd: Change timeout to fix problems with some 2.1 cards.
- dirmngr: Displayed name is now Key Acquirer.
- dirmngr: Add option --keyserver.  Deprecate that option for gpg.  Install a dirmngr.conf file from a skeleton for new installations.

## wrap-style がなかなかよい{#atom}

[ATOM Editor]({{< relref "remark/2015/atom-editor.md" >}}) は実は日本語混じりのテキストが苦手で，まともに使おうと思ったら相応のパッケージを導入する必要がある。
今までは [japanese-wrap] を使ってたのだが，同じ作者による [wrap-style] に乗り換えた。

- [Atomで上手にwarpを刻んでくれるwrap-styleを開発しました。 - Qiita](http://qiita.com/raccy/items/4678af4020189366a297)

作者ご本人は「遅い。遅すぎる」と書かれているが，なかなかどうして。
実用上は問題ない。

[japanese-wrap]: https://atom.io/packages/japanese-wrap "japanese-wrap"
[wrap-style]: https://atom.io/packages/wrap-style "wrap-style"
