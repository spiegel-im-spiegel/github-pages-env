+++
title = "性能とはなにか【『効率的な Go』読書会 初回】"
date =  "2024-04-21T13:18:52+09:00"
description = "『効率的な Go』読書会が始まった"
image = "/images/attention/go-logo_blue.png"
tags = [ "book", "golang", "engineering" ]
pageType = "text"

[scripts]
  mathjax = true
  mermaidjs = false
+++

冬の間忙しくてできなかった読書会への参加を再開することにした。

- [第58回横浜Go読書会（オンライン）- Part 2 - connpass](https://yokohama-go-reading.connpass.com/event/313675/)

「横浜Go読書会」とあって田舎暮らしの私としては躊躇するところではあったが主催の柴田芳樹さんより「[全国から参加可能](https://twitter.com/yoshiki_shibata/status/1760774448217723104)」と言っていただけたので参加することに決めた。

題材となる書籍はこれ。

{{% review-paapi "4814400535" %}} <!-- 効率的なGo : Efficient Go -->

[版元が O'Reilly Japan](https://www.oreilly.co.jp/books/9784814400539/ "O'Reilly Japan - 効率的なGo") ということで PDF で購入した。

タイトルを見たとき「ついに Effective Go を冠する本が出たのか」と思ったが[翻訳者の方](https://ymotongpoo.hatenablog.com/entry/efficient-go "『効率的なGo』という本が出版されました #efficient_go - YAMAGUCHI::weblog")に「["Efficient Go" だよ](https://bsky.app/profile/ymotongpoo.com/post/3knziwolx2c27)」と指摘していただいた恥ずかしい過去は内緒である。

普通は本を買ったら最低でも流し読みくらいをしておくものだけど，これを買った2月頃はホンマに忙しくてねぇ。
目次くらいしか眺めてなかった。
まぁ読書会で読むし，と後回しにしてたら結局初見で読書会に望むことになってしまった。
反省。

『[効率的な Go]』は11章で構成されていて500ページちょっとというボリューム。
読書会では読み終わるまで1年くらいかかるかなぁという見通しだった。
1回目の今回は都合で時間短めだったこともあり，序文と1章の途中まで。

読み始めた最初の感想は「講釈が長い！」だった。
[Go] のコードは殆ど出てこないし。

読書会のあとで後ろの章をさっくり眺めてみたが，コードは少なめで地の文章が多い。
なので，よくあるリファレンス本と思って読むと面食らうかもしれない。
どっちかというとソフトウェア工学，プログラミング（の考え方）を学ぶ教科書的な位置づけかなぁ。
その教材としてのプログラミング言語として [Go] が選ばれているという感じ。
実際に序文でも

{{< fig-quote type="markdown" title="『効率的なGo』序文" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
最適化の考え方やオブザーバビリティパターンを示すためにGoを例題言語として使っていますが、本書の11 章のうち8章は言語に中立的に書かれています。Java、C#、Scala、Python、C++、Rust、Haskellなど、他の言語で書かれたソフトウェアを改善するために、同じテクニックを使えます。
{{< /fig-quote >}}

と書かれている。
なので，この本の内容を [Go] 以外の言語で応用するにはどうすればいいか考えながら読むのも面白いかもしれない。

そういえば「訳者まえがき」には

{{< fig-quote type="markdown" title="『効率的なGo』訳者まえがき" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
私が直近で翻訳に携わった『オブザーバビリティ・エンジニアリング』（2023、オライリー・ジャパン、ISBN9784814400126）、『SLO サービスレベル目標』（2023、オライリー・ジャパン、ISBN9784814400348）の2冊に続いて、本書が加わることで、システムの性能改善というテーマにおいてマクロからミクロまで、全体の理解を深められる書籍が揃ったからです。したがって先の2冊を併せて読むことで、広い視点を保ったまま、プログラムの性能改善に取り組む意味が見えてくると思います。
{{< /fig-quote >}}

と書かれている。
この辺も読んでおくといいかも。

では，1章の前半で刺さったフレーズをつまみ食いしてみよう。

{{< fig-quote type="markdown" title="『効率的なGo』p.3" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
よく言われるように、「完璧は善の敵（Perfect is the enemy of good、https://oreil.ly/OogZF）」ですが、まずはそのバランスの取れた善を見つけなければなりません。
{{< /fig-quote >}}

これは個人的に

{{< fig-quote type="markdown" title="はやぶさ―不死身の探査機と宇宙研の物語" link="https://www.amazon.co.jp/dp/4344980158?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
理学は、真理の探究であり、工学は善の実現である。そして、藝術は美の表現である－－これで所謂「真美善」が揃う
{{< /fig-quote >}}

を座右の銘のひとつとしている私としては気になるフレーズである。
ちなみに [`https://oreil.ly/OogZF`](https://oreil.ly/OogZF) は Wikipedia へのリンクになっていて，そのバリエーションとして

{{< fig-quote type="markdown" title="Perfect is the enemy of good - Wikipedia" link="https://en.wikipedia.org/wiki/Perfect_is_the_enemy_of_good" lang="en" >}}
"If you never miss a plane, you're spending too much time at the airport"
{{< /fig-quote >}}

というフレーズがあって笑ってしまった。他に Donald Knuth 博士の

{{< fig-quote type="markdown" title="Perfect is the enemy of good - Wikipedia" link="https://en.wikipedia.org/wiki/Perfect_is_the_enemy_of_good" lang="en" >}}
"Premature optimization is the root of all evil"
{{< /fig-quote >}}

も紹介されていたが，これは『[効率的な Go]』でも言及されている。

{{< fig-quote type="markdown" title="『効率的なGo』p.13" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
しかし、この名言はずいぶん前になされたものです。過去から一般的なプログラミングについて多くを学ぶことができる一方で、1974年から大幅に改善されたことも多くあります。たとえば、当時はリスト1-5に紹介されているように、変数名に変数の型に関する情報を追加することが一般的でした。
{{< /fig-quote >}}

「リスト1-5」ってのがこれ

{{< fig-quote class="nobox" type="markdown" title="『効率的なGo』p.14" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
```go
type structSystem struct {
    sliceU32Numbers []uint32
    bCharacter      byte
    f64Ratio        float64
}
```
{{< /fig-quote >}}

いわゆるハンガリアン記法ってやつですな。

{{< fig-quote type="markdown" title="『効率的なGo』p.14" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
当時はコンパイラーや統合開発環境（IDE）があまり成熟していなかったので、ハンガリアン記法が便利でした。しかし現在では、IDEやGitHubのようなリポジトリサイトでも、変数にカーソルを合わせるとすぐにその型を認識できるようになりました。数ミリ秒で変数の定義にたどり着き、解説を読み、すべての呼び出しと変更を見つけられます。賢いコードの提案、高度なハイライト、1990 年代半ばに開発されたオブジェクト指向プログラミングの優位性により、実用的な可読性に大きな影響を与えることなく、機能や効率の最適化（複雑化）を追加できるツールを手に入れました。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.14 脚注" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
最近では、IDEの機能に対応しやすいようにコードを書くことが推奨されていることは、特筆に値するでしょう。たとえば、コード構造は「連結」グラフ（https://oreil.ly/mFzH9）であるべきです。これは、IDE が支援できる方法で関数を接続することを意味します。動的ディスパッチ、コードインジェクション、遅延読み込みは、これらの機能を無効にするので、厳密に必要でない限り避けるべきです。
{{< /fig-quote >}}

こういった感じにツールの助け等によってコードに対する認知負荷が変わってきているため「性能の最適化」も昔とは変わってきている，ということらしい。

じゃあ，そもそも「性能（performance）」って何？ って話になる。
この辺の話が今回読んだ中で一番面白かった。

{{< fig-quote type="markdown" title="『効率的なGo』p.4" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
私の経験では、多くの人がソフトウェア開発において性能（パフォーマンス）という言葉を速度（スピード）の代名詞と考えています。他の人々にとっては、この言葉の本来の定義である実行の全体的な質を意味します 。この現象は「意味拡散（semantic diffusion、https://oreil.ly/Qx9Ft）」と呼ばれることもあり、ある言葉がより大きな集団によって、本来持っていた意味と異なる意味で使われ始めるときに起こります。
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="『効率的なGo』p.6" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
結論から言うと、性能とは少なくともこの3つの要素の組み合わせです。

\[ 性能 = ( 精度 \times 効率 \times 速度 ) \]
{{< /fig-quote >}}

ここで性能は単にソフトウェアの機能を指すものではないというのが重要なポイント[^e1]。
日々要件が変わるソフトウェア開発においては如何にして変化をプロセスに組み込み小刻みにイテレーションを回していくかが重要になる。

[^e1]: 『[効率的な Go]』では 性能＝上手くできているか？ とし，その内訳として 精度＝間違いを犯してないか？ 効率＝余計な仕事をしてないか / 資源を使いすぎてないか？ 速度＝早くできているか？ といった感じに噛み砕いて説明している。

{{< fig-quote type="markdown" title="『効率的なGo』p.1" link="https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1" >}}
たとえば、RockstarGamesが人気ゲーム「グランド・セフト・オートV」を開発するのに5年の歳月と250人のエンジニアを要し、その費用は1億3750万ドルと推定されて（https://oreil.ly/0CRW2）います。また、Appleは、実用的で商品化されたOSを作るために、2001年にmacOSを初めてリリースするまでに5億ドル（https://oreil.ly/hQhiv）をはるかに超える資金を費やさなければなりませんでした。
{{< /fig-quote >}}

故に1章の「ソフトウェア効率性が重要」という見出しになるわけだ。

...という感じで読み進めていくことになりそうだ。
次回の読書会の感想を書くかどうかはわからないけど（笑）

[Go]: https://go.dev/
[効率的な Go]: https://www.amazon.co.jp/dp/4814400535?tag=baldandersinf-22&linkCode=ogi&th=1&psc=1 "効率的なGo ―データ指向によるGoアプリケーションの性能最適化 | Bartłomiej Płotka, 山口 能迪 |本 | 通販 | Amazon"

## 参考

{{% review-paapi "4873118468" %}} <!-- Go言語による並行処理 -->
{{% review-paapi "4621300253" %}} <!-- プログラミング言語Go -->
