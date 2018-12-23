+++
date = "2016-06-19T19:49:11+09:00"
update = "2017-06-21T16:13:55+09:00"
description = "GnuPG 2.1.13 および Libgcrypt 1.7.1 がリリース / 学校で「プログラミング」を学ばせる必要はない / 今週はセキュリティ・アップデート週間でした / Pokémon GO / その他の気になる記事"
draft = false
tags = ["security", "cryptography", "tools", "openpgp", "gnupg", "games", "ingress", "pokemon", "programming"]
title = "週末スペシャル： GnuPG 2.1.13 および Libgcrypt 1.7.1 がリリース"

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
  url = "http://www.baldanders.info/spiegel/profile/"
+++

1. [GnuPG 2.1.13 および Libgcrypt 1.7.1 がリリース]({{< relref "#gpg" >}})
1. [学校で「プログラミング」を学ばせる必要はない]({{< relref "#ed" >}})
1. [今週はセキュリティ・アップデート週間でした]({{< relref "#upd" >}})
1. [Pokémon GO]({{< relref "#pg" >}})
1. [その他の気になる記事]({{< relref "#other" >}})

## GnuPG 2.1.13 および Libgcrypt 1.7.1 がリリース{#gpg}

- [[Announce] Libgcrypt 1.7.1 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q2/000389.html)
- [[Announce] GnuPG 2.1.13 released](https://lists.gnupg.org/pipermail/gnupg-announce/2016q2/000390.html)

今回もセキュリティ・アップデートはなし。

GnuPG 2.1.13 の変更点は以下のとおり。

* gpg: New command `--quick-addkey`.  Extend the `--quick-gen-key` command.
* gpg: New `--keyid-format` "`none`" which is now also the default.
* gpg: New option `--with-subkey-fingerprint`.
* gpg: Include Signer's UID subpacket in signatures if the secret key has been specified using a mail address and the new option `--disable-signer-uid` is not used.
* gpg: Allow unattended deletion of a secret key.
* gpg: Allow export of non-passphrase protected secret keys.
* gpg: New status lines `KEY_CONSIDERED` and `NOTATION_FLAGS`.
* gpg: Change status line `TOFU_STATS_LONG` to use '`~`' as a non-breaking-space character.
* gpg: Speedup key listings in Tofu mode.
* gpg: Make sure that the current and total values of a `PROGRESS` status line are small enough.
* gpgsm: Allow the use of AES192 and SERPENT ciphers.
* dirmngr: Adjust WKD lookup to current specs.
* dirmngr: Fallback to LDAP v3 if v2 is is not supported.
* gpgconf: New commands `--create-socketdir` and `--remove-socketdir`, new option `--homedir`.
* If a /run/user/$UID directory exists, that directory is now used for IPC sockets instead of the GNUPGHOME directory.  This fixes problems with NFS and too long socket names and thus avoids the need for redirection files.
* The Speedo build systems now uses the new versions.gnupg.org server to retrieve the default package versions.
* Fix detection of libusb on FreeBSD.
* Speedup fd closing after a fork.

一方， Libgcrypt 1.7.1 の変更点は以下のとおり。

* Bug fixes:
    - Fix ecc_verify for cofactor support.
    - Fix portability bug when using gcc with Solaris 9 SPARC.
    - Build fix for OpenBSD/amd64
    - Add OIDs to the Serpent ciphers.
* Internal changes:
    - Use getrandom system call on Linux if available.
    - Blinding is now also used for RSA signature creation.
  - Changed names of debug envvars

インストール後はこんな感じ。

```text
$ gpg --version
gpg (GnuPG) 2.1.13
libgcrypt 1.7.1
Copyright (C) 2016 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <https://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Home: ********
サポートしているアルゴリズム:
公開鍵: RSA, ELG, DSA, ECDH, ECDSA, EDDSA
暗号方式: IDEA, 3DES, CAST5, BLOWFISH, AES, AES192, AES256, TWOFISH,
    CAMELLIA128, CAMELLIA192, CAMELLIA256
ハッシュ: SHA1, RIPEMD160, SHA256, SHA384, SHA512, SHA224
圧縮: 無圧縮, ZIP, ZLIB, BZIP2
```

## 学校で「プログラミング」を学ばせる必要はない{#ed}

- [小学校段階におけるプログラミング教育の在り方について（議論の取りまとめ）：文部科学省](http://www.mext.go.jp/b_menu/shingi/chousa/shotou/122/attach/1372525.htm)

単純に特定のプログラミング言語で「コーディングを覚えることが目的ではない」というのは一安心だが

{{< fig-quote title="小学校段階におけるプログラミング教育の在り方について" link="http://www.mext.go.jp/b_menu/shingi/chousa/shotou/122/attach/1372525.htm" >}}
<q>プログラミング教育とは、子供たちに、コンピュータに意図した処理を行うよう指示することができるということを体験させながら、将来どのような職業に就くとしても、時代を超えて普遍的に求められる力としての「プログラミング的思考」などを育むことであり</q>
{{< /fig-quote >}}

というのは全く以ていただけない。
「プログラミング的思考」に「時代を超えて普遍的に求められる力」はない。
はっきり言うが，それは全く無駄なことである。
それがやりたいのならコンピュータ系の専門学校にでも行かせればいい話である（まぁコンピュータ系の専門学校へ行ってモノになるのは全体の一割弱だろうけど）。

前に「[「だれもがプログラミングを学ぶべき」ではない]({{< ref "/remark/2016/05/lets-programming.md" >}})」でも書いたが，プログラミングもプログラミング言語も手段や道具でしかない。

プログラミングやプログラミング言語には前提となる学問体系が存在する。
それはたとえば「コンピュータ科学（computer science）」などと言われたりする。
また手段や道具を使うからには，それらを使う目的があるはずである。
たとえば「[ドリルを買う人が欲しいのは「穴」である](http://marketing-campus.jp/lecture/noyan/045.html "レビット博士が提唱したマーケティングで最も重要な格言とは | ノヤン先生のマーケティング講座 | 講座 | マーケティングキャンパス")」といったようなことだ。

コンピュータオタクには偏執者が多く，彼等は手段や道具そのものに異常な執着を見せるが，一般の人達には関係のないことだし，そんな執着心を真似すべきではない。

コンピュータ科学や社会学・経済学（の基礎となるもの）について小さい頃から学ばせる事自体は賛成だ。
しかし，何度でも言うが「[掛け算は順序が大事](http://www.baldanders.info/spiegel/log2/000744.shtml "日本の「算数」は壊れてる？ — Baldanders.info")」などと言ってはばからない未開人の国がどうやってそれらを教えるというのだ。

ちなみに「人工知能」の台頭を意識しているのなら，なおさら「プログラミング教育」は不要である。
なぜなら「問題を解決する仕事」はこれからどんどん機械が奪っていくから。
そうなった時に望まれる人材は「問題を解決できる人」ではなく「正しい問いを立てられる人」である。
親や学校教師に言われた通りのことしかできない子どもは，大人社会の中では機械以下の底辺でしか生きられなくなる。

一時の流行り廃りに流されることなく，きちんと足元を見て，50年100年のタイムスケールで「教育」というものを考えて欲しい。

### 参考

- [頼むからプログラミングを学ばないでくれ | TechCrunch Japan](http://jp.techcrunch.com/2016/05/17/20160510please-dont-learn-to-code/)

## 今週はセキュリティ・アップデート週間でした{#upd}

- [2016 年 6 月のセキュリティ情報 (月例) – MS16-063, MS16-068 ～ MS16-082 | 日本のセキュリティチーム](https://blogs.technet.microsoft.com/jpsecurity/2016/06/15/201606-security-bulletin/)
    - [2016年 6月 Microsoft セキュリティ情報 (緊急 5件含) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160025.html)
    - [Microsoft 製品の脆弱性対策について(2016年6月)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160615-ms.html)
    - [（緊急）Microsoft Windows DNSの脆弱性（リモートでのコード実行）について（CVE-2016-3227）](https://jprs.jp/tech/security/2016-06-17-msdns-vuln-remotecodeexec.html)
- [JVNDB-2016-003040 Apache Struts における任意のコードを実行される脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003040.html)
- [JVNDB-2016-003113 Android のメディアサーバの libstagefright における権限を取得される脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003113.html)
- [JVNDB-2016-003134 Android の sdcard/sdcard.c における権限を取得される脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003134.html)
- [JVNDB-2016-003093 GNU C Library の sunrpc/clnt_udp.c の clntudp_call 関数におけるスタックベースのバッファオーバーフローの脆弱性 - JVN iPedia - 脆弱性対策情報データベース](http://jvndb.jvn.jp/ja/contents/2016/JVNDB-2016-003093.html)
- [Adobe Flash Player の脆弱性対策について(APSA16-03)(CVE-2016-4171)：IPA 独立行政法人 情報処理推進機構](http://www.ipa.go.jp/security/ciadr/vul/20160615-adobeflashplayer.html)
    - [Adobe Flash Player の脆弱性 (APSB16-18) に関する注意喚起](https://www.jpcert.or.jp/at/2016/at160026.html)
    - [「Adobe Flash Player」にゼロデイ脆弱性、攻撃者にシステムを乗っ取られる恐れ - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1005272.html)
- [「Google Chrome 51」に3件の脆弱性、修正が施されたv51.0.2704.103が公開 - 窓の杜](http://forest.watch.impress.co.jp/docs/news/1005785.html)

あと，最近は不正アクセスやアカウント情報漏洩に絡む記事が多い。

- [Tumblrと（懐かしの）MySpaceから大量のログイン情報が流出 | Kaspersky Daily - カスペルスキー公式ブログ](https://blog.kaspersky.co.jp/myspace-tumbler-data-breach/11619/)
- [3200万のTwitterアカウントのパスワードがハックされてリークした、かもしれない | TechCrunch Japan](http://jp.techcrunch.com/2016/06/09/20160608twitter-hack/)
- [GitHub Security Update: Reused password attack](https://github.com/blog/2190-github-security-update-reused-password-attack)
    - [GitHub、不正アクセスを確認--他サービスから流出したパスワードを試行の可能性 - ZDNet Japan](http://japan.zdnet.com/article/35084420/)
- [JTBへの不正アクセスについてまとめてみた - piyolog](http://d.hatena.ne.jp/Kango/20160614/1465925330)
    - [JTB、793万人分の個人情報流出か--外部への通信で不正アクセスと判明 - ZDNet Japan](http://japan.zdnet.com/article/35084254/)


“[Have I been pwned?](https://haveibeenpwned.com/ "Have I been pwned? Check if your email has been compromised in a data breach")” のサイトには過去漏洩したアカウント情報のデータベースがあり，自分のアカウントが漏洩しているかどうか調べれることができる。
ご利用の際は自己責任でどうぞ。

なお，パスポート番号については

{{< fig-quote title="パスポート番号漏えい、偽造のリスクは？　外務省に聞いた - ITmedia ニュース" link="http://www.itmedia.co.jp/news/articles/1606/15/news102.html" >}}
<q>担当者は「番号や取得日が流出していても、パスポートの冊子じたいを手元にお持ちであれば、過剰な心配には及ばない」と話し、番号が漏れてしまった人も、再発行手続きなどは「不要だと考えている」という。外務省としても、今回の問題に伴う特別な対応などは「考えていない」という。</q>
{{< /fig-quote >}}

なんだそうで，あくまでもパスポートの原本が大事，ということらしい。

パスワード認証についてはいつものとおり。

1. 可能な限り2要素認証を使うこと
2. サービス間で同じパスワードを使いまわさないこと
3. パスワードは推測の難しい充分に複雑なものにすること

を最低限守っておけばいざというときの被害が少なくて済む[^a]。
パスワードの強度については，いつものあの表を引用するが，

[^a]: アカウント情報を含む個人情報漏洩はもはや「あり得べきもの」として事後をふまえた備えが必要。これを「未然防止」という。

<figure>
<table>
<thead>
<tr>
<th colspan='4'>利用する文字種類数と内訳</th>
<th colspan='4'>パスワード長</th>
</tr>
<tr>
<th>種類数</th>
<th>数字</th>
<th>文字</th>
<th>シンボル</th>
<th>4文字</th>
<th>8文字</th>
<th>12文字</th>
<th>16文字</th>
</tr>
</thead>
<tbody>
<tr><td>10種</td><td>0-9</td><td>なし</td>      <td>なし</td><td>1円未満（$2^{13.3}$）</td><td>1円未満（$2^{26.6}$）</td>  <td>約35円（$2^{39.9}$）</td>     <td>約35万円（$2^{53.2}$）</td></tr>
<tr><td>36種</td><td>0-9</td><td>a-z</td>       <td>なし</td><td>1円未満（$2^{20.7}$）</td><td>約100円（$2^{41.4}$）</td>  <td>約1.65億円（$2^{62.0}$）</td> <td>約276兆円（$2^{82.7}$）</td></tr>
<tr><td>62種</td><td>0-9</td><td>a-z<br>A-Z</td><td>なし</td><td>1円未満（$2^{23.8}$）</td><td>約7,500円（$2^{47.6}$）</td><td>約1,120億円（$2^{71.5}$）</td><td>約165京円（$2^{95.3}$）</td></tr>
<tr><td>94種</td><td>0-9</td><td>a-z<br>A-Z</td><td><code style='font-size:smaller;'>! " # $ %<br>&amp; ' ( ) =<br>~ | - ^ `<br>¥ { @ [<br>+ * ] ; :<br>} &lt; &gt; ? _<br>, . /</code></td>
                                                             <td>1円未満（$2^{26.2}$）</td><td>約21万円（$2^{52.4}$）</td> <td>約16.5兆円（$2^{78.7}$）</td> <td>約129,000京円（$2^{104.9}$）</td></tr>
</tbody>
</table>
<figcaption>パスワード解読の想定コスト例（<q><a href='https://www.ipa.go.jp/files/000026760.pdf'>情報漏えいを防ぐためのモバイルデバイス等設定マニュアル 解説編 <sup><i class='far fa-file-pdf'></i></sup></a></q> 2.4.2.2項より）</figcaption>
</figure>

12文字を目安に数字と英字（大文字小文字）を絡めたランダムな文字列にする。
人間の脳は無意識に規則性を求めてしまうものなので，横着せずパスワード生成ツールを使ったほうがよい。
もっと言うなら「[パスワードを覚えるなんて脳みその無駄遣い](http://www.baldanders.info/spiegel/log2/000739.shtml "「パスワードを覚える」なんて脳みその無駄遣い — Baldanders.info")」である。
[KeePass](http://keepass.info/ "KeePass Password Safe") のようなパスワード管理ツールを使うこと。

ちなみに [GitHub](https://github.com/) は2要素認証が使えるが，2要素認証を使うと https で git push する際にパスワードが有効でなくなる。
この場合は access token を取得してパスワードとして使えばよい。
また master branch にプロテクトを掛けることで直接 master branch に push させないようにする手もある（pull request でのみ merge 可能にする）。

{{< fig-img flickr="true" src="https://c8.staticflickr.com/8/7352/27154130743_8690a88078.jpg" title="Branch Protection" link="https://www.flickr.com/photos/spiegel/27154130743/" >}}

更に，前にも紹介したが， [commit に OpenPGP 署名を付与する]({{< ref "/remark/2016/04/git-commit-with-openpgp-signature.md" >}} "Git Commit で OpenPGP 署名を行う")ことで不正な merge を検出しやすくできるかもしれない。

[GitHub](https://github.com/) アカウントはエンジニアにとっては Twitter や Facebook よりも重要な identity なので，きっちり管理していきたいものである。

## Pokémon GO{#pg}

すっかり忘れていたが，いよいよ出るらしい。

- [『Pokémon GO』開発現場を直撃──もうすぐぼくらは、ポケモンを現実世界でゲットできる｜WIRED.jp](http://wired.jp/2016/06/18/pokemon-go-will-soon-bring/)

{{< fig-youtube id="lKUwVYUKii4" width="500" height="281" title="【公式】『Pokémon GO』　初公開映像 - YouTube" >}}

いやぁ [NIANTIC](https://www.nianticlabs.com/ "Niantic, Inc.") が絡んでて面白くならないわけがないよね。

{{< fig-quote title="『Pokémon GO』開発現場を直撃──もうすぐぼくらは、ポケモンを現実世界でゲットできる" link="http://wired.jp/2016/06/18/pokemon-go-will-soon-bring/" >}}
<q>イングレスで「ポータル」として指定されている地点は、いまや1,500万にもなる。図書館、博物館、歴史的な場所、彫像、パブリックアートワークといったそれらのランドマークが、Pokémon GOではポケストップやジムとして使われるわけだ。
<br>各地のデータを集めてくれたイングレスプレイヤーたちのおかげでPokémon GOはそのスタートから「ロケットスタート」を迎えられる。ちなみに、先週末カリフォルニアのマウントディアブロで登山をしていたハンケは、その途中でポケストップやジムになっているピースマーカー（山頂の基準点）やトレイルマーカー（道標）を見つけた。</q>
{{< /fig-quote >}}

さらにこれは朗報。

{{< fig-quote title="『Pokémon GO』開発現場を直撃──もうすぐぼくらは、ポケモンを現実世界でゲットできる" link="http://wired.jp/2016/06/18/pokemon-go-will-soon-bring/" >}}
<q>Pokémon GOを詳しく知るための、もう1つのパズルのピースが「Pokémon GO Plus」だ。Bluetooth接続するタイプのウェアラブルデヴァイスで、これを持ち歩けばスマートフォンを見続けずに済む。「そう、頭を下げて歩き回る必要はないのです」と、ハンケも言う。
<br>「ポケストップやポケモンに近づくと、点滅し振動します。アイテムを取得し、ポケモンを捕獲することもできます。スマホを見ることなく、毎朝の通勤中に使えます」</q>
{{< /fig-quote >}}

[Ingress](https://www.ingress.com/) でもこれがあれば良かったのに。
ねぇ。
そういえば，脚を痛めて以来 [Ingress](https://www.ingress.com/) はご無沙汰だったんだよな。
忙しいからオフ会とかも出れないし。

というわけで，今更だけど[フィールドテスト](http://www.pokemon.co.jp/redirect/index.php/3901/)に申し込んでみた。
もう無理かな。

さて，世のお父さんお母さんたちはどうなる（笑） どえらいおもちゃが出てしまいますよ。

## その他の気になる記事{#other}

- [Release v0.16 · gohugoio/hugo](https://github.com/gohugoio/hugo/releases/tag/v0.16)
- [新しい pLaTeX の話：非公式リリースノート 2016 年版 (1) - Acetaminophen’s diary](http://acetaminophen.hatenablog.com/entry/new-platex-20160507-01)
    - [新しい pLaTeX の話：非公式リリースノート 2016 年版 (2) - Acetaminophen’s diary](http://acetaminophen.hatenablog.com/entry/new-platex-20160507-02)
- [フロリダとパリのテロはグローバル・ジハードの拡散の典型：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41274)
    - [フロリダ州オーランドのゲイ・ナイトクラブへのテロに対するオバマの言葉：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41273)
    - [グローバル・ジハードの「触発」と「指示」の相違：池内恵 | 池内恵の中東通信 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41275)
    - [フロリダのテロが示す、宗教規範と世俗・リベラルな規範の間の「葛藤」：池内恵 | 中東の部屋 | 新潮社　Foresight(フォーサイト) | 会員制国際情報サイト](http://www.fsight.jp/articles/-/41276)
- [Amazon.co.jp： 増補新版 イスラーム世界の論じ方: 池内 恵: 本](http://www.amazon.co.jp/exec/obidos/ASIN/4120048349/baldandersinf-22/)
- [ASCII.jp：脆弱性情報のサイトでよく目にする「CVE」とは？](http://ascii.jp/elem/000/001/178/1178273/)
- [Apple、iPadでコーディングして学べる「Swift Playgrounds」を無償提供 - ITmedia ニュース](http://www.itmedia.co.jp/news/articles/1606/14/news071.html)
- [MicrosoftがLinkedInを262億ドルで買収、エンタープライズ向けソーシャルメディアに参入 | TechCrunch Japan](http://jp.techcrunch.com/2016/06/14/20160613microsoft-to-buy-linkedin-for-26b-in-cash-makes-big-move-into-enterprise-social-media/)
- [Instagramで売りに出される銃と、ソーシャルメディアのモラル｜WIRED.jp](http://wired.jp/2016/06/17/instagram-gunsforsale/)
- [ソフトウェアのバグをなくすには？--プログラミングの際に避けるべき10の失敗 - ZDNet Japan](http://japan.zdnet.com/article/35083529/)
- [k16's note: Markdown原稿をGitHubで管理して本にする仕組みが出版社で導入されないわけ](http://note.golden-lucky.net/2016/06/markdowngithub.html)
