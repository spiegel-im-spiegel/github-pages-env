+++
title = "Einstein@Home とパルサー"
date =  "2019-08-18T13:05:54+09:00"
description = "「計算機・ネットワークは有益な目的に使おう」"
image = "https://apod.nasa.gov/apod/image/1908/CannonSupernova_English_8404.jpg"
tags = [ "astronomy", "boinc", "pulsar", "gravity-wave" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

（「[星が好きな人のための新着情報](https://news.local-group.jp/20190815.html#p07)」より）

[Astronomy Picture of the Day (APOD)](https://apod.nasa.gov/) で面白い写真が公開されていた。

{{< fig-img src="https://apod.nasa.gov/apod/image/1908/CannonSupernova_English_8404.jpg" title="APOD: 2019 August 13 - Supernova Cannon Expels Pulsar J0002" link="https://apod.nasa.gov/apod/ap190813.html" width="8404" >}}

超新星爆発の残骸である CTB 1 から秒速千キロメートルで遠ざかっていくパルサー。
CTB 1 からはじき出されたと考えられるが詳しいプロセス等はまだ分かっていないようだ。

キャプションには

{{< fig-quote type="markdown" title="APOD: 2019 August 13 - Supernova Cannon Expels Pulsar J0002" link="https://apod.nasa.gov/apod/ap190813.html" lang="en" >}}
{{< quote >}}About 10,000 years ago, the [supernova](https://www.nasa.gov/subject/7226/supernova/) that created the nebular remnant [CTB 1](https://astrodonimaging.com/gallery/ctb-1-supernova-remnant/) not only destroyed a massive star but blasted its newly formed [neutron star](https://svs.gsfc.nasa.gov/12605) core -- a [pulsar](https://apod.nasa.gov/apod/ap090709.html) -- out into the Milky Way Galaxy. [The pulsar](https://public.nrao.edu/news/cannonball-pulsar/), spinning 8.7 times a second, was [discovered](https://ui.adsabs.harvard.edu/abs/2019ApJ...876L..17S/abstract) using downloadable software [Einstein@Home](https://einsteinathome.org/) searching through data taken by NASA's orbiting [Fermi Gamma-Ray Observatory](https://en.wikipedia.org/wiki/Fermi_Gamma-ray_Space_Telescope).{{< /quote >}}
{{< /fig-quote >}}

と書かれていて，どうやら [Einstein@Home] の成果を含んでいるらしい。
うんうん。

[Einstein@Home] は学術系分散コンピューティング基盤である [BOINC (Berkeley Open Infrastructure for Network Computing)](https://boinc.berkeley.edu/) 上で現在も運用されている科学プロジェクトである。

[Einstein@Home] プロジェクトの開始は2005年。
この年は「国際物理年」で様々なイベントが開催されたが，そのうちのひとつが [Einstein@Home] というわけだ。
一過性のお祭りに終わらず14年も運用が続いている（そして成果が出ている）というのは素晴らしいことである[^sah1]。

[^sah1]: 余談だが [SETI@home] は（[BOINC] 以前も含めると）正式運用開始から今年で[20周年](https://www.theringer.com/tech/2019/5/24/18637942/seti-home-aliens-citizen-science-extraterrestrial-life-20th-anniversary "SETI@home, the Alien-Hunting Project for the Nonscientist, Turns 20 - The Ringer")である。資金難で難儀したり最近ではアレシボ望遠鏡が被災したりして大変ではあるが，様々な副プロジェクトを含みつつ継続できているのはいいことだと思う。

なんで2005年が国際物理年だったかというと，2005年のちょうど100年前，1905年がアインシュタインによって3つの論文が公開された「奇跡の年」だったからだ。
その論文が

- 光電効果の理論
- 特殊相対性理論（当時は「相対性原理」と呼ばれていた）
- ブラウン運動の理論

である。
[Einstein@Home] では一般相対性理論が予言した「重力波」の直接検出を目指している。

{{< fig-quote type="markdown" title="Einstein@Home のすすめ -- Club-HUAA" link="http://huaa.baldanders.info/log/000181.shtml" >}}
{{< quote >}}1974年に連星パルサーの公転周期の軌道が短くなっていく現象が観測され，これが重力波の存在を示す間接的な証拠となっています。（連星パルサーは重力波を放出することによりエネルギーを失い公転周期が短くなるとかなんとか。で，実際に重力波があるとして計算した周期減少率が実際の観測と誤差の範囲内で一致したとかなんとか。ちなみにこの研究を行ったグループは1993年にノーベル物理学賞を受賞しています）{{< /quote >}}
{{< /fig-quote >}}

{{< fig-quote type="markdown" title="Einstein@Home のすすめ -- Club-HUAA" link="http://huaa.baldanders.info/log/000181.shtml" >}}
{{< quote >}}Einstein@Home では LIGO などの重力波干渉計でパルサーを観測します。
パルサーとは高速回転する中性子星のことで周期的（数ミリ秒から数秒）に電磁波（電波や可視光線あるいはX線など）を放射します。
もし完全に軸対称形でないパルサーがあればそれは重力波を放出するはずです。
ただし，もし重力波が放出されているとしてもそれは非常に微弱なものであろうと予測されています。
ノイズに埋もれた小さな重力波シグナルを検出するのには膨大な計算が必要です。
そこで「膨大な計算」の部分を私たちのコンピュータの空き時間を使って分担して行おうというわけです。
重力波干渉計はアメリカの LIGO （LIGO Hanford Observatory （LHO）および LIGO Livingston Observatory （LLO）の2箇所あります）以外にもドイツの GEO600 のデータも使われるようです。{{< /quote >}}
{{< /fig-quote >}}

（ちなみに日本にも重力波干渉計がいくつかあるが，それらがプロジェクトに参加するという話は今だに聞いたことがない）

重力波を検出するには観測対象となり得るパルサーを数多く捜索することが必要で， [Einstein@Home] ではそうした捜索も分散コンピューティングで行っている。
今回の「撃ち出されるパルサー」も [Einstein@Home] によるパルサー捜索の成果というわけだ。

[BOINC] はプロジェクトを構築するためのサーバ側のキットと参加ユーザに配布されるクライアント・ツールで構成されている。
クライアント・ツールはマルチプラットフォームに対応していて [Ubuntu] の場合は APT でインストールできる。

```text
$ sudo apt install boinc-client boinc-manager
```

{{< fig-img src="/remark/2019/04/academic-distributed-computing-projects-by-boinc/boinc-manager.jpg" title="OS 移行のため中断していたが BOINC による学術分散コンピューティング・プロジェクトの活動を再開した" link="https://www.flickr.com/photos/spiegel/33863325058/" >}}

クライアント・ツールでは参加プロジェクトの設定や計算機への負荷の調整等もできる。
ラズパイで専用機を組むのも面白いかもしれない。
無理のないところで参加していただければ幸いである。

「[**計算機・ネットワークは有益な目的に使おう**](https://baldanders.info/blog/000581/)」

## ブックマーク

- [BOINC による学術分散コンピューティング・プロジェクトでの活動を再開した]({{< ref "/remark/2019/04/academic-distributed-computing-projects-by-boinc.md" >}})
- [週末スペシャル： LIGO が重力波の「直接検出」に成功する！]({{< ref "/remark/2016/02/14-stories.md" >}})
- [Einstein@Home のすすめ -- Club-HUAA](http://huaa.baldanders.info/log/000181.shtml) : BOINC クライアントの使い方など一部内容が古くなっているのでご注意を

[APOD]: https://apod.nasa.gov/ "Astronomy Picture of the Day"
[Einstein@Home]: https://einsteinathome.org/
[SETI@home]: https://setiathome.berkeley.edu/
[BOINC]: https://boinc.berkeley.edu/
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"

## 参考図書

<div class="hreview">
  <div class="photo"><a class="item url" href="https://www.amazon.co.jp/SETI-home%E3%83%95%E3%82%A1%E3%83%B3%E3%83%96%E3%83%83%E3%82%AF%E2%80%95%E3%81%8A%E3%81%86%E3%81%A1%E3%81%AE%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%81%A7%E5%AE%87%E5%AE%99%E4%BA%BA%E6%8E%A2%E3%81%97-%E9%87%8E%E5%B0%BB-%E6%8A%B1%E4%BB%8B/dp/4898140866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4898140866"><img src="https://images-fe.ssl-images-amazon.com/images/I/51A74XV7MDL._SL160_.jpg" width="108" alt="photo"></a></div>
  <dl class="fn">
    <dt><a href="https://www.amazon.co.jp/SETI-home%E3%83%95%E3%82%A1%E3%83%B3%E3%83%96%E3%83%83%E3%82%AF%E2%80%95%E3%81%8A%E3%81%86%E3%81%A1%E3%81%AE%E3%83%91%E3%82%BD%E3%82%B3%E3%83%B3%E3%81%A7%E5%AE%87%E5%AE%99%E4%BA%BA%E6%8E%A2%E3%81%97-%E9%87%8E%E5%B0%BB-%E6%8A%B1%E4%BB%8B/dp/4898140866?SubscriptionId=AKIAJYVUJ3DMTLAECTHA&tag=baldandersinf-22&linkCode=xm2&camp=2025&creative=165953&creativeASIN=4898140866">SETI@homeファンブック―おうちのパソコンで宇宙人探し</a></dt>
    <dd>野尻 抱介</dd>
    <dd>ローカス</dd>
    <dd>単行本</dd>
    <dd>4898140866 (ASIN), 9784898140864 (EAN)</dd>
    <dd>評価<abbr class="rating fa-sm" title="4">&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="fas fa-star"></i>&nbsp;<i class="far fa-star"></i></abbr></dd>
  </dl>
  <p class="description">内容は古いけど当時の「熱」を伝えた名著だと思うけどなぁ。著者の方が自己出版で Kindle で出してくれたらいいのに。</p>
  <p class="powered-by">reviewed by <a href='#maker' class='reviewer'>Spiegel</a> on <abbr class="dtreviewed" title="2019-03-28">2019-03-28</abbr> (powered by <a href="https://affiliate.amazon.co.jp/assoc_credentials/home">PA-API</a>)</p>
</div>
