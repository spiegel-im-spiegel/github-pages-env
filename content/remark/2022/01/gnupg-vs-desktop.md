+++
title = "GnuPG の商業的成功と GnuPG VS-Desktop"
date =  "2022-01-03T12:25:37+09:00"
description = "国家や市場による独占ではなく，暗号技術が「自由」の名の下に私達の手の中に残り続けますように。"
image = "/images/attention/kitten.jpg"
tags = ["security", "cryptography", "openpgp", "gnupg", "market"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

昨年末に [GnuPG] による

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">To our donors: Please stop your donations to GnuPG and instead donate to projects which still need financial support</p>&mdash; GNU Privacy Guard (@gnupg) <a href="https://twitter.com/gnupg/status/1472943158346629125?ref_src=twsrc%5Etfw">December 20, 2021</a></blockquote>
{{< /fig-gen >}}

という tweet を[紹介]({{< ref "/remark/2021/12/gnupg-030-years-anniversary.md" >}} "GnuPG 030周年（笑）と寄付の話")したが，これに関する詳細記事が公開されている。

- [A New Future for GnuPG]
- [GnuPGの新しい未来](https://hidekis.notion.site/GnuPG-ab4490fce1c24f6bb8cf0799b494406a)（[斉藤英樹](https://home.hidekisaito.com/)さんによる日本語訳）

この記事では最初にこう述べられている。

{{< fig-quote type="markdown" title="A New Future for GnuPG" link="https://gnupg.org/blog/20220102-a-new-future-for-gnupg.html" lang="en" >}}
In the beginning GnuPG was a fun project I did in my spare time. After a few years this turned out to be a full time job and it was possible to acquire paid projects to maintain and further develop GnuPG.

When the BSI (Germany's Federal Office for Information Security) migrated back from Linux to Windows, a need to migrate their end-to-end encryption solution, based on GnuPG and KMail, was needed. A call for bids for an Open Source solution was issued and our company, g10 Code, along with our friends at Intevation and KDAB received the contract. The outcome was Gpg4win, the meanwhile standard distribution of GnuPG for Windows.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="GnuPGの新しい未来" link="https://hidekis.notion.site/GnuPG-ab4490fce1c24f6bb8cf0799b494406a" >}}
当初はGnuPGは私の楽しみのプロジェクトとして開始しました。数年後、それは専業の仕事としての形となっていき、その後、報酬を受け取るプロジェクトとして、GnuPGの管理や開発を行うことが可能となりました。

BSI（ドイツ連邦電子情報保安局）がLinuxからWindowsに移行するにあたり、GnuPGやKMailなどのエンドツーエンドな暗号ソリューションを移行する必要がありました。オープンソースなソリューションの入札により私たちの会社、g10 Code社及び、仲間の会社である、Intevation社とKDAB社が受注しました。その結果がGpg4winであり、これはしばらくの間、GnuPGのWindows向けディストリビューションの標準となっていました。
{{< /fig-quote >}}

当時を知ってる方は「あぁ，あれか」と思い出すだろうが，これがいわゆる Ägypten 計画である。

- [GnuPG - Project Ägypten](https://www.gnupg.org/aegypten/)
- [GnuPG - Project Ägypten2](https://www.gnupg.org/aegypten2/index.html)

Ägypten 計画の成果のひとつは暗号化電子メールの2大フォーマットである S/MIME と PGP/MIME を透過的に扱えるようにすることだった。
今でこそ鍵管理も含めてそういったことは Thunderbird など手軽にできる MUA があるが，当時は割と画期的だったのだ。

営利企業に踊らされてネットの中立性をかなぐり捨てた某フランチャイズ国家や2020年代に入ってもまだ [PPAP が云々]({{< ref "/remark/2020/06/security-theater.md" >}} "劇場型セキュリティと PPAP")とか言っている某野蛮国家とは比べ物にならないよね（笑）

私自身は MUA を Thunderbird へ乗り換えたこともあり Ägypten 計画以後の [GnuPG] 関連の商業的展開については関心を払ってこなかった。
2015年に [GnuPG] の財政的困窮が報じられてはじめて「まじすか！」と驚いたほどだ。

- [The World's Email Encryption Software Relies on One Guy, Who is Going Broke — ProPublica](https://www.propublica.org/article/the-worlds-email-encryption-software-relies-on-one-guy-who-is-going-broke)
- [Once-starving GnuPG crypto project gets a windfall. Now comes the hard part | Ars Technica](https://arstechnica.com/information-technology/2015/02/once-starving-gnupg-crypto-project-gets-a-windfall-but-can-it-be-saved/)

この記事をきっかけに寄付を中心とした [GnuPG] への継続的財政支援が行われるようになった。

{{< fig-quote type="markdown" title="A New Future for GnuPG" link="https://gnupg.org/blog/20220102-a-new-future-for-gnupg.html" lang="en" >}}
We started to make a real product out of *Gpg4VS-NfD*. Thus we rented a new office to work desk by desk on this and hired staff for sales and marketing. We introduced the brand *GnuPG.com* to have a better recognition of our product than by our legal name g10 Code GmbH. The software itself was re-branded as *GnuPG VS-Desktop<sup>®</sup>* and distributed as an MSI packet for Windows and as an AppImage for Linux. Except for customer specific configuration files *GnuPG VS-Desktop* is and will always be Open Source under the GNU General Public License.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="GnuPGの新しい未来" link="https://hidekis.notion.site/GnuPG-ab4490fce1c24f6bb8cf0799b494406a" >}}
私たちは、*Gpg4VS-NfD*より実際の製品を作り始めました。そのため、協力して仕事ができるように事務所を借り、販売やマーケティングのためのスタッフを雇用しました。私たちは法的な名称である*g10 Code社*よりもブランドとして通りやすい*GnuPG.com*の使用をはじめました。ソフトウェアそのものは*GnuPG VS-Desktop<sup>®</sup>* とブランド化し、こちらはWindows向けのMSIパッケージ及びLinux向けのAppImageとしての配布となっています。特定の顧客に対するカスタム部分設定以外においては*GnuPG VS-Desktop*はGNU General Public Licenseのオープンソースであり、またそれは今後も継続されます。
{{< /fig-quote >}}

さらに “[A New Future for GnuPG]” によると2019年末には BSI による VS-NfD の承認を得られたようだ。
VS-NfD 承認取得時のプレスリリース（ドイツ語）は以下の通り。

{{< fig-quote type="markdown" title="German press release: Gpg4win für VS-NfD freigegeben" link="https://gnupg.com/20200107-freigabe-vs-nfd.html" lang="de" >}}
**Erkrath, 7. Januar 2020**. Das Bundesamt für Sicherheit in der Informationstechnik (BSI) hat am 15. November 2019 die Freigabeempfehlung für Gpg4win, Version 3.x, und Gpg4KDE zur Übertragung und Verarbeitung von nationalen Verschlusssachen bis einschließlich Geheimhaltungsgrad VS-NUR FÜR DEN DIENSTGEBRAUCH (VS-NfD), RESTREINT UE/EU RESTRICTED und NATO RESTRICTED erteilt (BSI-VSA-10400, BSI-VSA-10412).
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="German press release: Gpg4win für VS-NfD freigegeben" link="https://gnupg.com/20200107-freigabe-vs-nfd.html" lang="de" >}}
Dank der Freigabeempfehlung kann GnuPG nun die bisher für VS-NfD verbreitete Software Chiasmus ersetzen. „OpenPGP und S/MIME haben gegenüber Chiasmus einige Vorteile“, sagt Werner Koch, Erfinder und einer der Hauptentwickler von GnuPG. „Man muss nur noch einmal die Zertifikate austauschen und braucht nicht für jede Kommunikation einen eigenen Schlüssel.“ 
{{< /fig-quote >}}

これにより [GnuPG] および [GnuPG VS-Desktop](https://gnupg.com/ "GnuPG.com") はドイツにおいて商業的な事業展開ができるようになった。

{{< fig-quote type="markdown" title="A New Future for GnuPG" link="https://gnupg.org/blog/20220102-a-new-future-for-gnupg.html" lang="en" >}}
Since summer 2021 the phones of our sales team didn't stop ringing and we could bring in the fruits of our work. We were not aware how many different governmental agencies exist and how many of them have a need to protect data at the VS-NfD (restricted) level. And with those agencies also comes a huge private and corporate sector who also have to handle such communication.

Although we support S/MIME, the majority of our customers decided in favor of the OpenPGP protocol, due to its higher flexibility and independence of a centralized public key infrastructure. A minor drawback is that for a quick start and easy migration from Chiasmus, many sites will use symmetric-only encryption (i.e. based on "gpg -c"). However, the now deployed software provides the foundation to move on to a comfortable public-key solution.

In particular, our now smooth integration into Active Directory makes working with OpenPGP under Windows really nice. We were also able to partner with *Rohde & Schwarz Cybersecurity GmbH* for a smooth integration of *GnuPG VS-Desktop* with their smartcard administration system.

We estimate that a quarter million workplaces will be equipped with *GnuPG VS-Desktop* and provide the users state of the art file and mail encryption. Our longer term plan is to equip all public agency workplaces with end-to-end encryption software - not only those with an immediate need for an approved VS-NfD solution. This should also fit well into the announced goal of the new German government to foster the development of Open Source.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="GnuPGの新しい未来" link="https://hidekis.notion.site/GnuPG-ab4490fce1c24f6bb8cf0799b494406a" >}}
2021年の夏以降、私たちのセールスチームの電話が鳴りやむことがなく、完成したものを提供することができました。私たちはどのくらいの数の政府機関が存在していて、その多くがVS-NfD（機密）レベルの保護が必要かを認識していませんでした。それらの政府機関とともに、多くの大手の民間企業や企業セクターがそのような通信を扱っています。

私たちはS/MIMEに対応していますが、顧客のほとんどはその柔軟性やPKIからの独立性を評価し、OpenPGPをより高く評価しました。些細な欠点としてはChiasmusよりの迅速な移行のためには、多くの場所において対象暗号化（gpg -cをベースにした）ものを使用することとなりました。しかしながら、新しく提供されるソフトウェアにおいては快適な公開鍵ソリューションの基盤となるものでした。

例えば、私たちのActive Directoryと統合して動作するソリューションはWindows環境下においてOpenPGPを使用する場合に良いものでした。また、*Rohde & Schwarz Cybersecurity社*とのパートナーシップにより、*GnuPG VS-Desktop*をスマートカード管理システムに統合することもできました。

私たちは25万の事業所において*GnuPG VS-Desktop*が使用され、最新のファイルとメールの暗号化が使用できると予測しています。私たちの長期計画としては例え承認されたVS-NfDソリューションを直ちに使用する必要がない場所を含む、全ての公共機関の事業所においてエンドツーエンドな暗号化ソフトウェアを導入することです。これは、ドイツ政府により発表された、オープンソフト開発を支援するという目標にも合致するものです。
{{< /fig-quote >}}

某日本では DX DX と姦しいが，こうやって下回りを整えつつ，時代に合わせて適宜入れ替えを行えることが重要なんだよ[^oss]。
ドイツで20年以上かけて積み上げてきたものを中身薄っぺらなデジタル庁だかが短期で達成できるわけ無いぢゃん（私は全く期待していない）。

[^oss]: もっと言うと「時代に合わせて適宜入れ替えを行える」ことが商業システムにおける FOSS の大きなメリットなんだよ。これは「それらを容易にするための戦略は、できるだけ長い期間、できるだけ多くの選択肢を残すことである」（『[Clean Architecture](https://www.amazon.co.jp/dp/B07FSBHS2V?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1)』より）というアーキテクチャ設計本来の目的とも合致する（ちなみに「それら」とはソフトウェアシステムの開発・デプロイ・運用・保守のこと）。バグや脆弱性を含めて，自分たちに都合が悪くなったからといって微塵も貢献していない FOSS 製品に文句を言って「直させる」のは根底から間違っている。「文句があるならコードをくれ！」だ。

というわけで “[A New Future for GnuPG]” は締めの言葉として

{{< fig-quote type="markdown" title="A New Future for GnuPG" link="https://gnupg.org/blog/20220102-a-new-future-for-gnupg.html" lang="en" >}}
For many years our work was mainly financed by donations and smaller projects. Now we have reached a point where we can benefit from a continuous revenue stream to maintain and extend the software without asking for donations or grants. This is quite a new experience to us and I am actually a bit proud to lead one of the few self-sustaining free software projects who had not to sacrifice the goals of the movement.

Those of you with SEPA donations, please cancel them and redirect your funds to other projects which are more in need of financial support. The Paypal and Stripe based recurring donations have already been canceled by us.

All you supporters greatly helped us to keep GnuPG alive and to finally setup a sustainable development model.
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="GnuPGの新しい未来" link="https://hidekis.notion.site/GnuPG-ab4490fce1c24f6bb8cf0799b494406a" >}}
長きにわたり、私たちの開発は寄付や小さなプロジェクトにより支えられていました。私たちは寄付や奨励金等に頼らなくとも、ソフトウェアのメンテナンスや拡張を続けられる継続的な収益を受けることができる状態に移行しました。これは非常に貴重な体験であり、その目的を犠牲にすることなく独立して開発ができる体制になった数少ないソフトウェアのリーダーとなったことを誇りに思います。

振り込みにより寄付を行っている方はその寄付を他の金銭による支援を必要としているプロジェクトに付け替えていただければと思います。PaypalとStripeによる継続的な寄付に関しましては私たちの側で解除いたしました。

全ての支援はGnuPGの継続とともに、独立して開発できる形態に移行するのにあたり、大変な助けとなりました。
{{< /fig-quote >}}

と書いている。

国家や市場による独占ではなく，暗号技術が「自由」の名の下に私達の手の中に残り続けますように。
そして OpenPGP やその実装のひとつである [GnuPG] が「自由」のための有効な手段として存在し続けますように。

[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[A New Future for GnuPG]: https://gnupg.org/blog/20220102-a-new-future-for-gnupg.html "A New Future for GnuPG"

## 参考図書

{{% review-paapi "4314009071" %}} <!-- 暗号化 プライバシーを救った反乱者たち -->
{{% review-paapi "B015643CPE" %}} <!-- 暗号技術入門 第3版 -->
