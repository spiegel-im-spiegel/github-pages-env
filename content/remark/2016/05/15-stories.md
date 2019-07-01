+++
date = "2016-05-15T18:46:06+09:00"
description = "水星の日面経過 / 今週はセキュリティ・アップデート週間でした / クラウドな仕事道具 / その他の気になる記事"
draft = false
tags = ["astronomy", "mercury", "transit", "security", "vulnerability", "cloud", "tools"]
title = "週末スペシャル： 水星の日面経過"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = ""
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = ""
  twitter = "spiegel_2007"
  url = "https://baldanders.info/profile/"
+++

1. [水星の日面経過]({{< relref "#tran" >}})
1. [今週はセキュリティ・アップデート週間でした]({{< relref "#upd" >}})
1. [クラウドな仕事道具]({{< relref "#cld" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## 水星の日面経過{#tran}

５月9日は水星の日面経過がありました。
ただし日本では見られず。
（以下の動画は再生時に音楽が流れるのでご注意を）

{{< fig-youtube id="8J4LoX3eOWc" title="A Mercury Transit Music Video from SDO - YouTube" width="500" height="281" >}}

特に，経過中に ISS が横切る以下の映像は面白い。

{{< fig-youtube id="Le-B2AnFiWw" title="Double solar transits of Mercury with the ISS and a plane - YouTube" width="500" height="281" >}}

ちなみに「経過（transit）」というのは観測者から見て大きな天体の手前を小さな天体が横切る現象を指す天文学用語。
これの反対が「{{< ruby "えんぺい" >}}掩蔽{{< /ruby >}}」。
ただしメディアや学校の教科書では「通過」と書かれるらしい。
どちらでも間違いではないようだ。

- [「日面経過」の用語について](http://optik2.mtk.nao.ac.jp/~somamt/notes/transit.htm)

### ブックマーク

- [「ひので」衛星が見た水星の太陽面通過](http://hinode.nao.ac.jp/news/160509MercuryTransit/)
- [APOD: 2016 May 11 - A Mercury Transit Music Video from SDO](http://apod.nasa.gov/apod/ap160511.html)
- [APOD: 2016 May 13 - ISS and Mercury Too](http://apod.nasa.gov/apod/ap160513.html)

## 今週はセキュリティ・アップデート週間でした{#upd}

毎月第2週水曜日はセキュリティ・アップデートの日。
主なものは以下のとおり。

- [2016 年 5 月のセキュリティ情報 (月例) – MS16-051 ～ MS16-062, MS16-064 ～ MS16-067 | 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/05/11/201605-security-bulletin/)
    - [2016年 5月 Microsoft セキュリティ情報 (緊急 8件含) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160022.html)
    - [Microsoft 製品の脆弱性対策について(2016年5月)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160511-ms.html)
- [Adobe Reader および Acrobat の脆弱性 (APSB16-14) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160023.html)
    - [Adobe Reader および Acrobat の脆弱性対策について(APSB16-14)(CVE-2016-1045等)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160511-adobereader.html)
- [Adobe Flash Player の脆弱性対策について(APSA16-02)(CVE-2016-4117)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160511-adobeflashplayer.html)
    - [Adobe Flash Player の脆弱性 (APSB16-15) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160024.html)
- [Cisco Talos Blog: Multiple 7-Zip Vulnerabilities Discovered by Talos](http://blog.talosintel.com/2016/05/multiple-7-zip-vulnerabilities.html)
    - [「7-Zip」v16.00には危険な脆弱性の修正も。「PeaZip」にもアップデートが提供される - 窓の杜](http://www.forest.impress.co.jp/docs/news/20160513_757356.html)

というわけできちんと対応しましょう。

あー，早く Windows 10 の無料アップグレード期間が終わらんかな。
こいつのせいで Windows Update を「自動更新」にできない。
鬱陶しいことこの上ない。
[サイバーテロ企業]({{< ref "/remark/2016/03/13-stories.md#win" >}})め！

個人的な印象論で申し訳ないが，こういうのは詐欺と同じで，きちんとアップデートしないユーザは「セキュリティに関心が薄い」とみなされ犯罪者の格好の標的となる[^s]。
無差別の phishing メールの malware が（普通ならとっくに塞がれているはずの）少し古い脆弱性を突いてくるのは，「カモネギ」なユーザか見極めているのではないかと思ったり。

[^s]: 「セキュリティに関心が薄い」ユーザなら侵入が発覚するまでかなりの時間を稼げるからだ。

## クラウドな仕事道具{#cld}

[GitHub] の新しい料金体系。

- [Introducing unlimited private repositories](https://github.com/blog/2164-introducing-unlimited-private-repositories)
- [［速報］GitHub、月額7ドルでプライベートリポジトリを無制限に作成可能に。新料金プランを発表 － Publickey](http://www.publickey1.jp/blog/16/github7.html)

日常生活も仕事も徐々にクラウドに移行しつつあるが，正直に言って「これ！」ってのがない。

PIM (Personal Information Management) に関しては通常は Gmail と Google Calendar で問題ない。
ただし仕事絡みのメールに関しては [ProtonMail](https://protonmail.com/ "Secure email: ProtonMail is free encrypted email.") に移行したいところである。
でもこれも相手がいてこそだからなぁ。

個人的なやり取りは SMS かキャリアメール（あるいは Facebook の Messenger[^l]）でしかやり取りできない人がほとんどなので諦めた。

[^l]: LINE は随分前に捨てた。

仕事のプロジェクトに参加するとそれぞれでアドレスを割り振ってもらえるが，こっちは非公開でしかも特定の groupware と関連付けられていることが多く融通が効かないのがアレである。

Storage は今のところ [Box](https://www.box.com/) がメイン。
最近は広島の企業でも [Box](https://www.box.com/) を利用するところが増えてきつつあるし。
ただし [Dropbox](https://www.dropbox.com/) じゃないと使い勝手が悪いアプリもあるので，そこはしょうがなく。

写真は [Amazon Cloud Drive](https://www.amazon.co.jp/gp/feature.html?docId=3077664656) にぶち込んでる（公になったら社会的に死ぬような写真は撮らない主義なので無問題）。
Prime に加入したので無制限に使えるのが魅力。
まぁこれは完全に「物置」代わり。

Evernote や Google Drive はなんか気持ち悪いので使わなくなった。
Evernote は無料版だと他サービスと連携できなくて使い勝手が悪いってのもあり，最終的には捨てる予定。
でも古い資産が結構残ってるので他所に移さないといけないんだよなぁ。

今借りてるレンタルサーバの容量が余りまくってるので「[さくらぽけっと](https://www.sakura.ad.jp/press/2015/0312_sakurapocket/ "さくらインターネット、さくらのレンタルサーバをオンラインストレージとして使えるスマートフォンアプリ「さくらぽけっと」をiOS／Androidで提供開始")」も考えたのだが，今時 EUC-JP を要求するクソ仕様なので（何故 UTF-8 にしない）使えない。

Workplace は [GitHub] または [Bitbucket]。
[Bitbucket] は private repository が必要な場合のみ利用している。
[GitLab.com](https://gitlab.com/ "Code, test, and deploy together with GitLab open source git repo management software | GitLab") は割とよさ気な感じなので [Bitbucket] にある private repository の一部をこっちに移すのもありかなぁ，と思っている。

[Backlog](http://www.backlog.jp/) は個人的にとても好みなのだが，容量の割にお値段高めで手が出ない。

Storage と workplace に関してはそろそろちゃんとお金を払ってどこかに定住したいのだが，何処にしようか悩み中。
たとえば， [GitHub] に月7ドル払うのなら「[さくらの VPS](http://vps.sakura.ad.jp/ "VPS（仮想専用サーバー）｜さくらインターネット - 無料お試し実施中")」で環境を作ったほうがいいんじゃないか，とも思うわけよ。

あるいは storage だけに絞って [SpiderOak](https://spideroak.com/) に月12ドル払って 1TB の安全な環境を確保する手もある。
[Box](https://www.box.com/) や [Dropbox](https://www.dropbox.com/) に[プライベートなファイルは置きたくない](http://jp.techcrunch.com/2014/10/13/20141011edward-snowden-new-yorker-festival/ "スノーデンのプライバシーに関する助言：Dropboxは捨てろ、FacebookとGoogleには近づくな | TechCrunch Japan")。
[これから実装されるという Project Infinite](http://www.publickey1.jp/blog/16/dropboxproject_infinite.html "Dropbox、クラウドとのファイル同期をファイルへのアクセス時にオンデマンド実行してくれる「Project Infinite」プレビュー、チームのファイル共有向け － Publickey") は魅力だけど，所詮 Windows と Mac だけの話だし，これから Windows を捨てようかという私には関係ない気がする。

まぁ，もう少し考えてみるか。

## その他の気になる記事{#other}

- [go-ole/go-ole: win32 ole implementation for golang](https://github.com/go-ole/go-ole) : Go 言語で OLE/COM を操作できるらしい
- [【解説】 クリエイティブ・コモンズ・ライセンス入門 【知財管理65巻6号掲載】 | 弁護士 増田雅史の記録帳](https://masudalaw.wordpress.com/2016/05/06/ccl-basics/)
- [死を司る「テロメア」とは何なのか？｜WIRED.jp](http://wired.jp/2016/05/08/about-telomere/)
- [DRMの悪夢：KoboをアップグレードするとSonyリーダから受け継いだ電子書籍がライブラリから消えてしまう - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20160508/drmnightmare)
- [News: KeePass 2.33 available! - KeePass](http://keepass.info/news/n160507_2.33.html)
- [2016年5月11日ニュース「太陽系外惑星、新たに1284個発見 地球型も9個」 | SciencePortal](http://scienceportal.jst.go.jp/news/newsflash_review/newsflash/2016/05/20160511_02.html)
    - [NASA Finds 1,284 Alien Planets, Biggest Haul Yet, with Kepler Space Telescope](http://www.space.com/32850-nasa-kepler-telescope-finds-1284-alien-planets.html)
    - [NASA、太陽系外惑星1284個を発見 | TechCrunch Japan](http://jp.techcrunch.com/2016/05/13/20160512astronomers-announce-largest-batch-of-new-planets-ever-discovered/)
- [ダン・ギルモア著『あなたがメディア　ソーシャル新時代の情報術』を全文公開します | 新聞紙学的](https://kaztaira.wordpress.com/2016/05/12/%E3%83%80%E3%83%B3%E3%83%BB%E3%82%AE%E3%83%AB%E3%83%A2%E3%82%A2%E8%91%97%E3%80%8E%E3%81%82%E3%81%AA%E3%81%9F%E3%81%8C%E3%83%A1%E3%83%87%E3%82%A3%E3%82%A2%E3%80%80%E3%82%BD%E3%83%BC%E3%82%B7%E3%83%A3/)
- [パナマ文書事件が明らかにした「第五階級」とは « マガジン航[kɔː]](http://magazine-k.jp/2016/05/13/panama-papers-and-fifth-estate/) : 確かに WikiLeaks とは異なるが， WikiLeaks の次の段階とも言える
- [Amazon.co.jp： 【中東大混迷を解く】 サイクス=ピコ協定 百年の呪縛: 池内 恵: 本](https://www.amazon.co.jp/exec/obidos/ASIN/4106037866/baldandersinf-22/) : とりあえず予約した
    - [【サポートページ開設】『サイクス＝ピコ協定 百年の呪縛』のカテゴリーを設定しました – 中東・イスラーム学の風姿花伝](http://ikeuchisatoshi.com/%E3%80%90%E3%82%B5%E3%83%9D%E3%83%BC%E3%83%88%E3%83%9A%E3%83%BC%E3%82%B8%E9%96%8B%E8%A8%AD%E3%80%91%E3%80%8E%E3%82%B5%E3%82%A4%E3%82%AF%E3%82%B9%EF%BC%9D%E3%83%94%E3%82%B3%E5%8D%94%E5%AE%9A-%E7%99%BE/)
- [「モノのインターネットは、セキュリティの面では悪夢だ」EFFの警告 | TechCrunch Japan](http://jp.techcrunch.com/2016/05/12/20160509the-internet-of-things-is-security-nightmare-warns-eff/)

[Box]: https://www.box.com/
[GitHub]: https://github.com/
[Bitbucket]: https://bitbucket.org/ "Bitbucket — The Git solution for professional teams"