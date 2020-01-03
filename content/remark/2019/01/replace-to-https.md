+++
title = "（不完全ながら） HTTPS 接続に対応した"
date = "2019-01-07T17:53:06+09:00"
description = "各ページに埋め込まれる他サイトのスクリプトや画像等のマテリアルに HTTPS に対応していないものがあるため不完全な対応になっている点はご了承の程を。"
image = "/images/attention/kitten.jpg"
tags = [ "site", "web", "tls", "cryptography" ]

[scripts]
  mathjax = false
  mermaidjs = false
+++

遅まきながら[本家サイト]の HTTPS 接続に対応した。

いや，随分前に「[さくらのレンタルサーバ](https://www.sakura.ne.jp/)」で [Let's Encrypt] を用いた [TLS 設定が簡単になった](https://www.sakura.ne.jp/function/freessl.html "無料SSLサーバー証明書 Let's Encrypt - レンタルサーバーはさくらインターネット")という話は聞いてたのよ。

- [さくらインターネットの「さくらのレンタルサーバ」、コントロールパネル上の簡単操作で無料SSL証明書「Let’s Encrypt」を設定可能に | さくらインターネット](https://www.sakura.ad.jp/information/pressreleases/2017/10/10/90197/)
- [当ブログもHTTPSになりました！さくらの500円サーバーで「Let's Encrypt」によるSSL化](https://chalow.net/2018-01-14-1.html)
- [今更ではあるが、本家サイトのHTTPS対応を行った - YAMDAS現更新履歴](http://d.hatena.ne.jp/yomoyomo/20180420/https)

ただし，個人的には積極的に HTTPS に切り替える動機はなかった[^tls1] ので放置していた。

[^tls1]: こう言ったら専門家の方には鼻で笑われるだろうが，セキュリティ専門家が言う「木の葉を隠すなら林の中」理論で馬も鹿も HTTPS ってのには首を傾げてしまう。ましてや SEO 対策で HTTPS に切り替えるとか阿呆としか言いようがない。ひとつのドメインやサービスで，ページや状況によって HTTP と HTTPS を使い分けるな，というのならその通りだと思うけれど。まぁ [Let's Encrypt] が正式リリースされたし「面倒臭い」以外に HTTPS にしない理由はないので，今回はいい機会になったと考えている。

ところが Facebook で聞いた話によると米国では ISP が Web の通信に対して何らかの干渉を行う懸念を払拭できないとかで，そういえば米国では[「ネットの中立性」は公式に破れている]({{< ref "/remark/2017/12/hacker-ethic-2.md" >}})し日本もいずれそうなるだろうとは想像に難くない。

そうであるなら，そういった通信に対する干渉を抑止するためにも HTTPS を有効にしておくのは意味があるだろうと考えを改めることにした。
まぁ，ページの広告エリアに採掘コードを差し込むとか，セキュリティと称して HTTPS 通信に中間者攻撃をかまして出歯亀するとかいったことが横行している状況で HTTPS でどこまで安全が担保されるのかは分からないが。

なお，各ページに埋め込まれる他サイトのスクリプトや画像等のマテリアルに HTTPS に対応していないものがあるため不完全な対応になっている点はご了承の程を[^site1]。
これは徐々に改修していく予定。

[^site1]: なにせ1999年以来19年分の垢がこびりついたサイトなので。って今年は（[本格的に Web で活動を始めて](https://baldanders.info/spiegel/log/)から）20周年やん！

このため，当面は HTTP と HTTPS の両方とも有効にして運用し強制的な HTTPS リダイレクトはしない。
上述したように，私は HTTPS をデータの機密性および完全性の観点からはあまり信用してないので「やらないよりはマシ」というスタンスで運用することにする。

## 2019年の目標{#2019}

実は今年の目標というか TODO として[本家サイト]の改造を考えていて，具体的には

1. [本家サイト]を [Hugo] ベースに再構築する（ただし URL はできるだけ変えない）
2. Flickr の写真を引き揚げて[本家サイト]に移転する

と考えている。
今回の HTTP 接続対応はそのついでというわけだ。
上の2つのうち 2 に関しては既に全バックアップをとって手元にデータはあるのだが，[本家サイト]での公開については 1 の作業と併せて行う予定なので，しばらく先のことになるだろう。

これ以外にもうひとつ問題があって， Amazon アフィリエイト・リンク作成サービスの [Amakuri] がサービスを終了するらしい。

- [【Amakuri】来年15日にAmakuriは閉鎖いたします - dadadadoneのメモ帳](https://blog.dadadadone.com/archives/235)

[G-Tools さんに続いて]({{< ref "/remark/2018/10/amazon-affiliate-with-amakuri.md" >}} "Amazon アフィリエイトリンク作成サービスを Amakuri へ移行する")お宅もか `orz`

どうしよう。
たぶん類似サービスは軒並み閉鎖だよね，これ。
自前でツールを用意するしかないのか。
それが面倒くさいから人様のサービスを利用してたのに。

## [ここのブログ]も  HTTPS 接続に対応していた

[ここのブログ]は [GitHub Page](https://github.com/spiegel-im-spiegel/spiegel-im-spiegel.github.io "") なのだが，なにもしないのに何時の間にか HTTPS 接続に対応していたらしい。

- [Custom domains on GitHub Pages gain support for HTTPS | The GitHub Blog](https://blog.github.com/2018-05-01-github-pages-custom-domains-https/)

こちらも当面は HTTP と HTTPS の両方とも有効にして運用するが Flickr の写真と Amazon アフィリエイト・リンクの目処が立てば HTTPS 強制にしてもいいかも知れない。

{{< div-box type="markdown" >}}
*【2019-05-21】* 以下を参考に HTTP 強制にしました。

- [Securing your GitHub Pages site with HTTPS - GitHub Help](https://help.github.com/en/articles/securing-your-github-pages-site-with-https)
{{< /div-box >}}

### 【追記1】 e-Words 用語集止めました

[ここのブログ]ではページ下部に [e-Words] 用語集のブログパーツを貼り付けていたのだが，既にメンテナンスされていないようで，かなり前から説明ページが消えてるし，しかも HTTP 接続オンリーで HTTPS なページに貼り付けるとブラウザから reject されるので，今回を機に削除することにした。

本当は [Disqus] も止めたい。
[Disqus] の埋め込みコードは Firefox からはトラッキング・コードと見做されているため reject される。
設定を緩めればいける[^dq1] のだが，それも微妙な気もする。
[Twitter](https://twitter.com/spiegel_2007) か [Facebook ページ](https://www.facebook.com/baldanders.info/ "Baldanders.info")でフィードバックを，という手もあるが...

[^dq1]: どうも Firefox 側でホワイトリストまたはブラックリストを持っていて，トラッキング・コードの配信元から有効/無効を判断してるっぽい。ホワイトリスト方式またはブラックリスト方式は運用する側にどうしても恣意が混入してしまう。この辺は spam メール対策で懲りているので賢いやり方に見えない。

### 【追記2】 Amazon 関連の URL 変換

Amazon の商品ページへのリンク URL や書影等のイメージの URL について以下のように変換した。

{{< fig-gen title="Amazon 関連の URL 変換" >}}

<style>
main table.amazon th  {
  vertical-align:middle;
  text-align: center;
}
main table.amazon td  {
  vertical-align:middle;
  text-align: left;
}
</style>
<table class="amazon">
<thead>
    <tr>
        <th>変換前</th>
        <th>変換後</th>
    </tr>
</thead><tbody>
    <tr>
        <td><code>http://www.amazon.co.jp/</code></td>
        <td><code>https://www.amazon.co.jp/</code></td>
    </tr><tr>
        <td><code>http://ecx.images-amazon.com/</code></td>
        <td rowspan="3"><code>https://images-fe.ssl-images-amazon.com/</code></td>
    </tr><tr>
        <td><code>http://images.amazon.com/</code></td>
    </tr><tr>
        <td><code>http://g-images.amazon.com/</code></td>
    </tr>
</tbody>
</table>

{{< /fig-gen >}}


## ブックマーク

- [ウェブサイトの常時SSL化普及へ向けたさくらのレンタルサーバの取り組み | さくらインターネット](https://www.sakura.ad.jp/information/news/2018/12/19/1968198392/)
- [GitHub - rinopo/sakura-init: さくらのレンタルサーバを借りたとき最初にすること](https://github.com/rinopo/sakura-init)
- [Gitのリモートリポジトリにプッシュすると自動的にWebサイトが更新されるリポジトリを作成する |  Arcany](https://tapioca-hiroyuki.net/?blog=git0320)
- [さくらレンタルサーバーとSourceTreeでGit環境構築 - Qiita](https://qiita.com/zprodev/items/0a16bc51866ee6a7a102)
- [.htaccessによるアクセス制御 – さくらのサポート情報](https://help.sakura.ad.jp/hc/ja/articles/206054622)
- [さくらのレンタルサーバーでのリダイレクト設定「.htaccess」を使ってhttp→httpsとwww有り無し統一のやり方 | SEO対策なら株式会社ペコプラ](https://pecopla.net/seo-column/sakura-server-redirect-method)
- [さくらインターネットでWordPressを動かす時の定番のトラブル | ゴルディアスの涙目](https://gordiustears.net/trouble-with-wordpress-on-sakura-internet/)

- [Why we’re changing Flickr  free accounts | Flickr Blog](https://blog.flickr.net/en/2018/11/01/changing-flickr-free-accounts-1000-photos/)
- [Flickrの新しいビジネスモデルでCreative Commonsの作品は消されるのか  |  TechCrunch Japan](https://jp.techcrunch.com/2018/11/03/2018-11-02-flickrs-new-business-model-could-see-works-deleted-from-creative-commons/)

[本家サイト]: https://baldanders.info/ "Baldanders.info"
[ここのブログ]: {{< rlnk "/" >}} "text.Baldanders.info"
[Let's Encrypt]: https://letsencrypt.org/ "Let's Encrypt - Free SSL/TLS Certificates"
[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"
[Amakuri]: https://dadadadone.com/amakuri/ "Amakuri [Amazonアフィリエイトリンク作成ツール]"
[e-Words]: http://e-words.jp/ "IT用語辞典 e-Words"
[Disqus]: https://disqus.com/
