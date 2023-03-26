+++
title = "AI は「創作者様」を引きずり下ろすか — 『人権と文化コモンズ』を流し読む"
date =  "2022-10-23T15:06:30+09:00"
description = "この記事を含めて，ブログではこの本の感想を細切れで書いていくことになると思う。"
image = "/images/attention/kitten.jpg"
tags = [ "book", "intellectual-property", "copyright", "artificial-intelligence", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

ここのところ[図書館][島根県立図書館]で読んでたのは山田奨治さんの『[著作権は文化を発展させるのか: 人権と文化コモンズ][人権と文化コモンズ]』だったり。

{{% review-paapi "B099RTG3J7" %}} <!-- 著作権は文化を発展させるのか: 人権と文化コモンズ -->

[またか]({{< ref "/remark/2022/10/library-2.md" >}})とお思いでしょうが，またなんです。
だってまとまった時間をとって読書をするのに図書館は最適の空間（affordance）なんだもの。

で，まぁ，読みながら思ったのは「これ，一回読んだだけじゃ分からんわ」だったり。
なので今のところ「完全に理解した」は諦めて「最後まで読み通す」ことに注力することにした。
なので今後は，何度か読み直すことになるだろう。

この記事を含めて，ブログではこの本の感想を細切れで書いていくことになると思う。
理解の不備の訂正も同様。

## ユーザの「表現」

『[人権と文化コモンズ]』で提示されている内容は私が著作権に対して抱いている問題意識と似ている，と感じている。

7年前に書き始めた「[改訂3版： CC Licenses について]({{< rlnk "/cc-licenses/" >}})」で私はこう書いている。

{{< fig-quote type="markdown" title="著作権と著作権法" link="/cc-licenses/01-copyright/" >}}
著作者や出版社にとっての著作権システムとは，著作物の「独占」状態を如何にコントロールするかであった。
しかし，ユーザの場合は行動原理が異なる。それはひとことで言うなら「共有」である。
何故ならユーザはコミュニケーションの手段（メディア）として著作物（＝表現）を利用するからだ。
{{< /fig-quote >}}

一方で『[人権と文化コモンズ]』では表現を行使する全ての人の「人権」という観点から著作権を捉え直すところからスタートしている。
つまり

{{< fig-quote type="markdown" title="著作権は文化を発展させるのか: 人権と文化コモンズ" link="https://www.amazon.co.jp/dp/B099RTG3J7?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
人間の内面は、幾多の著作物でできている（p.64）
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="著作権は文化を発展させるのか: 人権と文化コモンズ" link="https://www.amazon.co.jp/dp/B099RTG3J7?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
ユーザーの人権は、受け取った作品が身体化していることにその根拠がある（p.81）
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="著作権は文化を発展させるのか: 人権と文化コモンズ" link="https://www.amazon.co.jp/dp/B099RTG3J7?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
著作者とユーザーの、両方の身体の中間に作品は立ち現れるのだ（p.89）
{{< /fig-quote >}}

ときて

{{< fig-quote type="markdown" title="著作権は文化を発展させるのか: 人権と文化コモンズ" link="https://www.amazon.co.jp/dp/B099RTG3J7?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
いま文化産業を名乗る彼らは、「文化」の「創り手＝送り手」を特権化して強力な保護を求め、ユーザーを消費者の位置に押し込めて、その自由を奪おうとしてきた。彼らのいう「文化」とは、その本質的な意味からは程遠い「消費財」でしかない。（p.153）
{{< /fig-quote >}}

とまで書いている。
私は不覚にも図書館内で笑いそうになった。

## プログラム・コードという表現

ところで『[人権と文化コモンズ]』ではプログラム・コードについての言及がない（読み落としがなければ）。

現代はプログラム・コードも表現のひとつとみなされ著作権システムに組み入れられている。
「著作物」という観点でプログラム・コードが小説・音楽・絵画等と違うのは「作家性」の強度だろうか。

たとえば「Linux は Linus Torvalds 作」みたいな感じで製品と作者がダイレクトに繋がっている場合もあるが，実際には一から十まで彼が全部プログラム・コードを書いているわけではなく，多数の参加者による製品に対する貢献（contribution）の上に成り立っている。
この辺の話は，今や古典となった「[伽藍とバザール](http://cruel.org/freeware/cathedral.html)」あたりを読むのが手っ取り早いだろう。
FOSS は「バザール型」を積極的に受け入れ，ひいては IT 産業全体の構造をも変えたわけだ。

2022年に入って AI によるコード生成支援サービスである GitHub Copilot が騒がしいが

- [Give Up GitHub: The Time Has Come! - Conservancy Blog - Software Freedom Conservancy](https://sfconservancy.org/blog/2022/jun/30/give-up-github-launch/)
  - [「GitHubの利用を中止しよう」 SFCが提言、AI開発ツールに疑念 | TECH+](https://news.mynavi.jp/techplus/article/20220701-2385316/)
- [ソースコードの続きを自動補完するGitHub Copilotに「著作権で保護されたコードを出力している」という指摘が寄せられる - GIGAZINE](https://gigazine.net/news/20221018-github-copilot-emit-copyrighted-code/)

GitHub Copilot の問題は「コードをパクっている」ことではなく FOSS を支える「自由のライセンス」まで洗い流され「ボクが考えたコード」みたいな顔をして何の前提もなく提示してしまうことにある（GitHub Copilot が知識ソースとしているのは GitHub 上の公開された FOSS コードである）。
こうした「どんぶり勘定」を放っておくなら Microsoft/GitHub はいずれ困った事態に直面すると思うのだが，どうだろう[^gc1]。

[^gc1]: 私個人としては GitHub Copilot 問題の関心領域は少しずれたところにある。詳しくは拙文「[GitHub Copilot は貢献者から貢献を奪うか？]({{< ref "/remark/2022/07/code-launderring.md" >}})」をどうぞ。

## AI は「創作者様」を引きずり下ろすか

もうひとつ AI 界隈で騒がしいのは「AI 絵画」だろうか。
こうした話は突然出てきたように思えるかもしれないが，楽曲の世界では以前からあったし，文章を書く AI みたいな製品も存在する。

ただ「AI 絵画」を含む一連が象徴するのは「[機械学習技術の民主化]」であろう。

{{< fig-gen >}}
<blockquote class="twitter-tweet"><p lang="ja" dir="ltr">mimicのような物議然り、このフィールドはそれなりに仕組みを応用できる人がPCでも実現できる世界にすでになっているため、恐らく今後もこういうケースは続出するし、多分もう流れをせき止めることはできないと思う。良くも悪くも機械学習技術が民主化されてしまった影響。</p>&mdash; Hideki Saito @ メタバース技術者＆Vtuber (@hideki) <a href="https://twitter.com/hideki/status/1564509002205052931?ref_src=twsrc%5Etfw">August 30, 2022</a></blockquote>
{{< /fig-gen >}}

例を挙げると Microsoft は GitHub Copilot と同じ技術基盤を用いて Azure 上で Dall·E 2 の提供を開始した。

- [［速報］マイクロソフト、Azure OpenAI Serviceで「Dall·E 2」の提供を開始、テキストから画像を生成可能。Ignite 2022 － Publickey](https://www.publickey1.jp/blog/22/azure_openai_servicedalle_2ignite_2022.html)

こうした動きは加速することはあっても巻き戻すことはもはや不可能だ。

『[人権と文化コモンズ]』で著者が指摘している通りに，特権者たる「創り手＝送り手」が生み出しているものが文化とは名ばかりの「その本質的な意味からは程遠い「消費財」でしかない」のであれば，この「[機械学習技術の民主化]」がもたらす **翻案の大量生産** は「来たるべき未来」と言うほかない。
私が連想するのはスープ缶の絵を並べたあのポップアートだが，「AI 絵画」はアレの延長線上にあるように見える。

今まで「創作者様」として特権的な位置から生み出していた「創作物」が実は機械によって大量生産可能な「消費財」に過ぎなかったと， AI によって気付かされるわけだ。

人が「汗をかいて」成し遂げてきたことが機械に取って代わられるなんてことは過去に幾度も繰り返されてきたし，それによって失職する人が出たり新たな職業が生まれたりするのも歴史の必然だろう。
大昔にあった「コーダー」という職業がコンピュータとソフトウェア開発環境の劇的な向上に伴って消え失せたように，将来「プログラマ」という職業もなくなるかもしれない。

情報のデジタル符号化とインターネットの台頭は，間違いなく現代におけるパラダイムシフトと言える。
今ですら著作権システムはこれに適応できていないのに，さらに「[機械学習技術の民主化]」が進めばどうなってしまうのか。

...今回はこの辺で。

[島根県立図書館]: https://www.library.pref.shimane.lg.jp/
[人権と文化コモンズ]: https://www.amazon.co.jp/dp/B099RTG3J7?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "著作権は文化を発展させるのか: 人権と文化コモンズ"
[機械学習技術の民主化]: {{< ref "/remark/2022/08/ai-drawing.md" >}}

## ブックマーク

- [生成AIと『猿の手』の呪い：著作権を強化してもクリエイターは救われない | p2ptk[.]org](https://p2ptk.org/copyright/4349)
- [AIアートが簡単・安価に生み出される世界で、それでも著作権は必要なのか？ | p2ptk[.]org](https://p2ptk.org/copyright/4352)
- [著作権は必要不可欠だった、でもこれからもそう？ | p2ptk[.]org](https://p2ptk.org/copyright/4360)

## 参考文献

{{% review-paapi "4622073455" %}} <!-- 〈海賊版〉の思想‐18世紀英国の永久コピーライト闘争 -->
{{% review-aozora "227" %}} <!-- 伽藍とバザール -->
