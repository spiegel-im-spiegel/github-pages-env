+++
title = "絶対インターネットになんか負けたりしない!!"
date =  "2019-08-01T22:28:11+09:00"
description = "アメリカには勝てなかったよ orz"
image = "/images/attention/kitten.jpg"
tags = [ "internet", "github" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

- [GitHubがイランなどからアクセス不可に、米国の経済制裁により。CEOのフリードマン氏「望んでやっているのではない」 － Publickey](https://www.publickey1.jp/blog/19/githubceo.html)

この記事を見て GitHub の CEO がレイプ目で「アメリカには勝てなかったよ」とか言ってるシーンを妄想してしまった私は薄汚れてますね，すみません。

今回の記事に関して

{{< fig-quote type="md" title="「インターノット」の敗北 | おごちゃんの雑文" link="http://www.nurs.or.jp/~ogochan/essay/archives/5487"  >}}
{{< quote >}}多くの場合「マスターリポジトリ」はGitHubにあったりする。社内の開発であってもGitHubのプライベートサービスを使って行なわれたりする(今回アクセスが遮断されたのはそこである。OSSな公開リポジトリは関係ない)。開発の中心にGitHubがあるために、GitHubにアクセス出来ないと協調が出来なくなってしまう。元々、そういったことがないように分散リポジトリを提供していたはずのgitが、GitHubを使うことによって「ちょっと便利なcvs」くらいになってしまった{{< /quote >}}
{{< /fig-quote  >}}

とかいった話も聞かれるが，はっきり言って git には「マスターリポジトリ」などというものは存在しない。
上流にあるリポジトリも下流にあるリポジトリも実質的に等価であり，入れ替えも分岐も簡単にできるし上流リポジトリを分散させることも可能だ。
これがそれまで主流だった CVS や Subversion などといったものとは異なる部分である。

もし GitHub に BAN されたら開発体制が決定的に破綻してしまうというのなら，この機会に見直しを行ったほうがいいだろう。
職業エンジニアなら特定のサービスにロックインしてしまうことの危うさは身に沁みているはずである。

商用利用が解禁されて以降，つまり20年以上前からインターネットは「分散ネットワーク」でも「非中央集権ネットワーク」でもなくなった。
これを印象づけたのが2010年末から始まった「アラブの春」である。
市民運動で揺れるエジプトではたった5本の電話で8000万人のインターネット・アクセスが遮断された。

{{< fig-quote title="介入されないもうひとつのウェブ" link="http://www.nikkei-science.com/201206_074.html" >}}
<q>インターネットが当初の学術目的から踏み出して現在のような誰でも使える商業サービスになってから20年以上が経つうちに，そうした蓄積伝送の原理が果たす役割は，一貫して縮小していった。<br />
　この間，ネットワークに加わる新たなノードの圧倒的多数はISPを介してネットに接続する家庭や企業のコンピューターだった。
ISPの接続モデルでは，利用者のコンピューターはデータの中継はしない。
それはネットワークの端末，つまりデータの送受信だけを，常にISPのコンピューターを介して行うターミナル・ノードだ。
言い換えれば，インターネットの爆発的な成長はネットワーク地図に行き止まりのルートを増やしただけで，新たなルートを加えることはほとんどなかった。<br />
　そしてISPなど大量の情報ルートを持つ者は，彼らがルートを提供している何百万ものノードを支配下におくこととなった。
これらのノードは，もしISPがダウンしたり，ネットから遮断されたりすると，その障害を回避する方法がない。
ISPはインターネットが停止しないようにするどころか，実効上は停止スイッチになってしまった。</q>
（p.77）
{{< /fig-quote >}}

つまり GitHub も「停止スイッチ」を持っていて，国家はそのスイッチを押させることができるということだ。

これから本気で「分散ネットワーク」とか「非中央集権ネットワーク」とかいったものを構築したいのなら，それは最低でも「インターネットがなくとも機能する」ことが要件となるだろう。
インターネットは（国家や社会や企業といったものの思惑に）敗北したのだから。

## ブックマーク

- [“The Shadow Web” （再掲載）]({{< ref "/remark/2016/11/the-shadow-web.md" >}})

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/%E3%82%B5%E3%82%A4%E3%83%90%E3%83%BC%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3-%E5%88%A5%E5%86%8A%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9-%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%E7%B7%A8%E9%9B%86%E9%83%A8/dp/4532512123?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532512123"><img src="https://images-fe.ssl-images-amazon.com/images/I/51gurnOqhiL._SL160_.jpg" width="120" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/%E3%82%B5%E3%82%A4%E3%83%90%E3%83%BC%E3%82%BB%E3%82%AD%E3%83%A5%E3%83%AA%E3%83%86%E3%82%A3-%E5%88%A5%E5%86%8A%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9-%E6%97%A5%E7%B5%8C%E3%82%B5%E3%82%A4%E3%82%A8%E3%83%B3%E3%82%B9%E7%B7%A8%E9%9B%86%E9%83%A8/dp/4532512123?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4532512123">サイバーセキュリティ (別冊日経サイエンス)</a></dt>
    <dd>日経サイエンス編集部 (編集)</dd>
    <dd>日本経済新聞出版社 2016-04-22 (Release 2016-04-22)</dd>
    <dd>ムック</dd>
    <dd>4532512123 (ASIN), 9784532512125 (EAN)</dd>
    <dd>Rating<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">『日経サイエンス』2012年6月号の「介入されないもうひとつのウェブ」が収録されている。その他にも2010年代前半におけるセキュリティ問題についてよくまとめられている。</p>
  <p class="powered-by" >reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2016-11-05">2016-11-05</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home" >PA-API</a>)</p>
</div>
