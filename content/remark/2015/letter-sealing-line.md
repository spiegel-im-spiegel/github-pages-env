+++
date = "2015-10-14T22:44:28+09:00"
description = "昨年 EFF が安全なメッセージングアプリの条件を示した。この条件に照らせば Letter Sealing は最初の2つはクリアしたけど，残りは「もっと頑張りましょう」な状態。"
draft = false
tags = ["security", "privacy", "cryptography", "messaging", "line"]
title = "LINE： Letter Sealing による End-to-End 暗号化"

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

- [LINE、トークを高度な暗号化で保護する「Letter Sealing」を実装　設定方法も紹介 | アプリオ](http://appllio.com/how-to-set-line-letter-sealing)

以前の「[タイマートーク](http://ja.hideki.hclippr.com/2014/08/01/line%E3%81%AE%E3%82%BF%E3%82%A4%E3%83%9E%E3%83%BC%E3%83%88%E3%83%BC%E3%82%AF%E3%81%A7%E6%80%9D%E3%81%86%E3%81%93%E3%81%A8/)」のように「どうせ日本リージョンのユーザには関係ない話だろう」と思っていたら，日本でもやるらしい。
LINE は捨てたので完全に見落としていたのだが，既に9月時点で実装されていたようだ。

- [LINEのE2EE実装「Letter Sealing」初見 | Hideki's Random Stuff Japanese](http://ja.hideki.hclippr.com/2015/09/01/line%E3%81%AEe2ee%E5%AE%9F%E8%A3%85%E3%80%8Cletter-sealing%E3%80%8D%E5%88%9D%E8%A6%8B/)

Letter Sealing に関する詳細情報も公開されている。

- [New generation of safe messaging: “Letter Sealing” « LINE Engineers' Blog](http://developers.linecorp.com/blog/?p=3679)
- [メッセージの安全性新時代：Letter Sealing « LINE Engineers' Blog](http://developers.linecorp.com/blog/ja/?p=3591)

まぁ Google の Hungout や Facebook の Messanger よりマシになったとは言えるだろう。
上の記事で説明されている仕様は以下の通り。

- エンド・ユーザ間の Diffie-Hellman 鍵交換。アルゴリズムは ECDH/256bits。ただし ephemeral かどうかは不明[^0]
- メッセージの暗号化には上の鍵交換で生成した AES/256bits 鍵を使う。モードは CBC
- HMAC によるメッセージの完全性の担保[^1]

[^0]: Ephemeral かどうかは通信が PFS (Perfect Forward Secrecy) を満たすか否かに関わってくる。
[^1]: 「HMACは、ハッシュ関数(hash funciton)という一方向性関数(one-way function)でメッセージに対するデジタル署名(digital signature)を取得し」と書かれているのは何かの冗談だと思う。英語の記事にはそんな記述ないし。

お気づきと思うが，このままでは鍵の真正性を証明（certification）できない。
相手から送られて来る公開鍵情報を無条件に受け入れてしまう。
ぶっちゃけて言うなら LINE サーバ上で中間者攻撃（man-in-the-middle attack）が可能ということだ[^a]。
「[LINEのE2EE実装「Letter Sealing」初見](http://ja.hideki.hclippr.com/2015/09/01/line%E3%81%AEe2ee%E5%AE%9F%E8%A3%85%E3%80%8Cletter-sealing%E3%80%8D%E5%88%9D%E8%A6%8B/)」によるとクライアントごとに異なる鍵を生成するようなので[^b]，なおさら鍵の証明が重要になってくると思うのだけど。

[^a]: ただし中間者攻撃は（昨年の [OpenSSL に対する HeartBleed](http://www.baldanders.info/spiegel/log2/000682.shtml) のように）既に攻撃が可能な状態なら相当な脅威だけど，攻撃を行うための前提条件を揃えるのは結構面倒なので，全体としてのリスクはそれほど高いわけではない。
[^b]: これ自体は妥当。もっと言うなら， ephemeral な鍵であればおそらくセッション毎に異なる。

どうも最近は「暗号鍵がユーザに見えるのはダサい」みたいな風潮があるけど，公開鍵暗号方式を使うなら鍵の証明と管理の scheme が生命線になるので，これを怠ると今回のような中途半端な実装になってしまう。

昨年， EFF が安全なメッセージング・アプリの条件を示した。

- [Secure Messaging Scorecard | Electronic Frontier Foundation](https://www.eff.org/secure-messaging-scorecard)
- [ASCII.jp：実は危険だらけのメッセージアプリ](http://ascii.jp/elem/000/000/958/958626/)

条件は7つ。

1. Encrypted in transit?
2. Encrypted so the provider can't read it?
3. Can you verify contacts' identities?
4. Are past comms secure if your keys are stolen?
5. Is the code open to independent review?
6. Is security design properly documented?
7. Has there been any recent code audit?

この条件に照らせば Letter Sealing は最初の2つはクリアしたけど[^c]，残りは「もっと頑張りましょう」な状態。
せめて条件4までは全て対応していただきたいところだ。

LINE の場合，アカウントの乗っ取りや成りすましが（手を変え品を変え）現在も続いているようなので，（脆弱なユーザの purge を含めて）そろそろ抜本的に対策していかないと「暗号化したけど意味ないじゃん」ってことになりかねないと思う[^d]。

[^c]: PFS を満たすなら条件4も満たす。
[^d]: ていうか， LINE はユーザを増やしすぎたよね。これは Facebook とかでも同じだけど。総じて LINE は真正性について甘い。むしろ意図的に避けてる印象。真正性を担保できない状態で，どれだけ暗号化に力を入れようとも無意味。

## 参考

- [安全なメッセージング・アプリとは — Baldanders.info](http://www.baldanders.info/spiegel/log2/000782.shtml)
- [安全なメッセージング・アプリ（ちょびっと改訂） — Baldanders.info](http://www.baldanders.info/spiegel/log2/000800.shtml)
- [LINEのLetter Sealingに関するフォローアップ | Hideki's Random Stuff Japanese](http://ja.hideki.hclippr.com/2015/10/14/line%E3%81%AEletter-sealing%E3%81%AB%E9%96%A2%E3%81%99%E3%82%8B%E3%83%95%E3%82%A9%E3%83%AD%E3%83%BC%E3%82%A2%E3%83%83%E3%83%97/)
- [【注意喚起】 友人や知人になりすまして電話番号やSMS認証番号を聞き出すメッセージにご注意ください : LINE公式ブログ](http://official-blog.line.me/ja/archives/39021529.html) : 今年の夏ごろに流行った乗っ取り手口
