+++
title = "脆弱性情報の収集・管理ツール jvnman の最初のリリース"
date = "2018-05-13T19:43:21+09:00"
update = "2018-05-13T21:30:45+09:00"
description = "先月あたりから余暇にコツコツ作ってた脆弱性情報の収集・管理ツール [jvnman] の最初のリリースを行った。"
image = "/images/attention/tools.png"
tags  = [ "tools", "security", "vulnerability", "risk", "management" ]

[author]
  name      = "Spiegel"
  url       = "http://www.baldanders.info/spiegel/profile/"
  avatar    = "/images/avatar.jpg"
  license   = "by-sa"
  github    = "spiegel-im-spiegel"
  twitter   = "spiegel_2007"
  tumblr    = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  flickr    = "spiegel"
  facebook  = "spiegel.im.spiegel"
  linkedin  = "spiegelimspiegel"
  flattr    = "spiegel"

[scripts]
  mathjax = false
  mermaidjs = false
+++

先月あたりから余暇にコツコツ作ってた脆弱性情報の収集・管理ツール [jvnman] の最初のリリースを行った。

- [spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management](https://github.com/spiegel-im-spiegel/jvnman)
    - [Release v0.1.0 · spiegel-im-spiegel/jvnman](https://github.com/spiegel-im-spiegel/jvnman/releases/tag/v0.1.0)

インストールおよび使い方については[リポジトリ](https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management")の [README.md](https://github.com/spiegel-im-spiegel/jvnman/blob/master/README.md "jvnman/README.md at master · spiegel-im-spiegel/jvnman") を参照のこと。

このツールを作った理由は大きく2つ。

ひとつは私が Windows 7 のサポートの切れる2020年を目処に Windows プラットフォームを捨てる予定であり Windows 環境への依存をできるだけ排除する形でセキュリティ管理を行えないか模索していること。
もうひとつは昨今のセキュリティや脆弱性情報に関するメディアの報道にいい加減ウンザリしていることだ。

Windows ユーザはセキィリティ管理に関しては比較的恵まれている。
ベンダ企業である Microsoft によるセキュリティ管理のダメな部分は20世紀のうちに粗方やり尽くしていて，現在は脆弱性の修正サイクルも安定しているため安心して利用できる。
まさに se-cure な状態なのだ。
これが破られたのが [Windows 10 の強制アップロード]({{< ref "/remark/2015/windows-10-upgrade-problem.md" >}} "また Windows 10 にヤラレタ（KB3112343 の恐怖）")のときだったが，それ以外は（年に2,3個ほど？ ヤバいのがある程度で）本当に安定している。
まぁ，あの強制アップロード事件があったから Windows を本気で捨てる気になったんだけどね。

障害管理やセキュリティ・インシデント・レスポンスが稚拙な Apple 製品や Android 端末機は論外としても，たとえば Linux 機でどの程度セキュリティを気にしなければならないか，といった感覚的な部分でよく分からなかったりする。
私がセキュリティ管理者をやってたのはかれこれ20年近く前なので当時の経験は殆ど参考にならない（当時のセキュリティ周りは実に牧歌的だった，今に比べれば）。
かといって今更この歳でセキュリティ管理者みたいなヘヴィな職を目指すつもりもない。

私が学生の頃にある教官から教わった言葉に「情報浴」というのがある。
ある分野で最先端を行きたいなら常にその分野の情報を浴びていなさい，という意味である。
裏を返せば専門でない（興味があるわけでもない）分野の情報を過剰に浴びるのは毒でしかないということである。
セキュリティ情報というのは（専門家以外から見れば）まさに毒なのだ。

煩わしいことに自身の寿命とお金をかけることほど馬鹿げてることはない。
どうせリソースを使うなら煩わしさを解消する方向に使いたい。
セキュリティ情報が過剰に溢れているのなら，もはやそれを人間が扱うのは悪手である。
機械に任せられることは機械にやらせればいいのだ。

そのための調査を兼ねて [jvnman] を運用してみようと考えた次第である。

[jvnman]: https://github.com/spiegel-im-spiegel/jvnman "spiegel-im-spiegel/jvnman: JVN Vulnerability Data Management"
