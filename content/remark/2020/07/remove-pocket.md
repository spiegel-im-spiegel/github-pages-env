+++
title = "Pocket を捨てて Instapaper に出戻る"
date =  "2020-07-31T13:06:48+09:00"
description = "上手く出戻りが成功するといいのだけど。"
image = "/images/attention/kitten.jpg"
tags = [ "web", "surveillance-capitalism" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

普段は，移動中などで見た Web サイトを Pocket に保持っておいて後でパソコンのブラウザ等でじっくり読むスタイルなのだが， Pocket から Web ページに上手く移動できないので「おかしいな」と思ってリロードしたら「サードパーティ cookie を拒否する設定になっているから Pocket を例外にしろ」（← かなり意訳）みたいな表示が出てきた。

{{< div-gen type="markdown" class="center caution" >}}
！！！ふざけんな！！！
{{< /div-gen >}}

しばらくしたら元に戻ったみたいだけど，これで決心がついた。

前々から言っているが「[Mozilla がユーザのプライバシーを重視しているというのは嘘っぱちである]({{< ref "/remark/2019/09/should-disable-doh-in-firefox.md" >}} "Firefox の DoH は無効にすべきか（もしくは水売りと水道局）")」。

まぁ Web ブラウザを提供している営利企業でユーザのプライバシーに誠実なところはひとつもないと断言できると思うので今更な話ではあるが，ついに [Mozilla 傘下の Pocket](https://www.itmedia.co.jp/news/articles/1702/28/news067.html "Mozilla、“後で読む”のPocketを買収 - ITmedia NEWS") も監視資本主義へ邁進し始めたかと思うと感慨深い。

実際，サードパーティ cookie を dis って困ることはほぼない。
むしろ今回みたいなことを言ってくさるサービスを排除できるので便利なくらいだ。

しかし Pocket を捨てるのはいいが代替サービスがない。

- Readability 無くなってるぢゃん
- ネットで見かける「あとで読む」系サービスの紹介記事がどれも古くて，いまだに「Read It Later[^pct1]」を紹介している記事も検索上位にあるし
- Evernote など，大抵のサービスは Pocket との連携が前提になっているらしい。ニンともカンとも
- 「はてブ」はなー。しかも，確か「はてな」の各サービスはサードパーティ cookie 拒否を嫌ってたような...

[^pct1]: “Read It Later” は Pocket の旧サービス名。

というわけで消去法で [Instapaper] になった。
アカウントが残っててよかったよ。
念のため，パスワードだけ変更しておく。

そういえば，なんで [Instapaper] 止めたんだっけ？ 有料化したあたりからかな？ 覚えてない。

[Instapaper] は2018年に Pinterest から独立したらしい。

- [Instapaper is going independent](https://blog.instapaper.com/post/175953870856)
- [「後で読む」サービスのInstapaper、Pinterestから独立へ - CNET Japan](https://japan.cnet.com/article/35122536/)

今は無料版でも使えるようだが，継続的な運用のためにサブスクリプション契約を結ぶことを推奨している。

{{< fig-quote type="markdown" title="The next ten years of Instapaper" lkink="https://blog.instapaper.com/post/176732408411" lang="en" >}}
{{< quote >}}In addition to getting access to Premium features, your Instapaper Premium subscription will help ensure that we can continue developing and operating Instapaper. Our goal is to build a long-term sustainable product and business, without venture capital, and we need your help to achieve that goal{{< /quote >}}.
{{< /fig-quote >}}

志は買うし，年額 29.99 USD ならさして高くもないが，実際に課金するかどうかはもう少し使ってみてからね。

Web ブラウザや他サービスとの連携も問題なさそうだ。

- [Apps](https://www.instapaper.com/apps)
- [How to Save](https://www.instapaper.com/save)
    - [Instapaper – Get this Extension for 🦊 Firefox (en-US)](https://addons.mozilla.org/en-US/firefox/addon/instapaper-official/)

さて，上手く出戻りが成功するといいのだけど。

[Instapaper]: https://www.instapaper.com/

## ブックマーク

- [Flickr は監視資本主義に向かわない]({{< ref "/remark/2019/03/flickr-has-not-turned-to-surveillance-capitalism.md" >}})

## 参考図書

{{% review-tatsujin "infoshare2" %}} <!-- 続・情報共有の未来 -->
{{% review-paapi "B01MZGVHOA" %}} <!-- 超監視社会 -->
