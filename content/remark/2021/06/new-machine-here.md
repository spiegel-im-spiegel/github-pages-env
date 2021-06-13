+++
title = "自宅マシンを買うた（これで私も人並みに...）"
date =  "2021-06-13T17:05:03+09:00"
description = "なんちゃらとパソコンは新しいものに限る。"
image = "/images/attention/kitten.jpg"
tags = [ "tools", "ubuntu" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++


ふっふっふ，歩が3つですョ。
ついにパソコン買うたですョ。

本当なら2020年までに買い換える予定だったのが，諸事情で田舎に引っ込むことになり，収入ガタ落ちで職業エンジニアに復帰することもないかなぁとしみじみしてたところに，思いがけず拾っていただいた会社があって，よーやくパソコンが買えるまでになった。
そのうち FOSS プロジェクトに寄付できるくらいにまで懐事情が回復すればいいのに，と希望的観測を述べてみる。

{{< ruby "それはさておき" >}}閑話休題{{< /ruby >}}

[8年前](https://baldanders.info/blog/000632/ "これで私も人並みに")と同じく，今回も[ドスパラ](https://www.dospara.co.jp/)でミニタワーを見繕った。
その前（2005年）がソフマップの「牛丼パソコン」（←知ってる人いるかなぁ）だったし，8年ごとに買い替えてる勘定か。

コアは [Linus Torvalds 氏も絶賛](https://linux.slashdot.org/story/20/05/25/020240/linus-torvalds-dumps-intel-for-32-core-amd-ryzen-on-his-personal-pc "Linus Torvalds Dumps Intel For 32-core AMD Ryzen On His Personal PC - Slashdot")する AMD Ryzen だよ。
いや，ごめん。
うちのは6コア12スレッドの AMD Ryzen5 PRO 4650G なのでだいぶ落ちる[^amd1]。
調子こいてすんません。

[^amd1]: [Linus Torvalds 氏が絶賛](https://linux.slashdot.org/story/20/05/25/020240/linus-torvalds-dumps-intel-for-32-core-amd-ryzen-on-his-personal-pc "Linus Torvalds Dumps Intel For 32-core AMD Ryzen On His Personal PC - Slashdot")している自宅マシンのコアは AMD Threadripper 3970x なのでだいぶ違う（笑）

Windows バンドルは外してメモリを32GBに，ストレージを512GB SSD に換装した。
これで9万円＋消費税。
更にスピーカ付き24インチモニタ（2万円＋税）を付けても12万円弱。
やっと自宅マシンも複数モニタになったよ。

今回の構成を職場の（その筋の）オタク同僚に評価してもらったが，今どきのゲーミングマシンは48GBメモリが普通だし，デュアル・モニタにするくらいなら50インチを買ったほうがいいと言われてしまった（笑） いいもん。私はパソコンでゲームしないし。

OS は [Ubuntu 21.04 日本語 Remix](https://kledgeb.blogspot.com/2021/05/ubuntu-2104-25-ubuntu-2104-remix.html "Ubuntu 21.04 その25 - Ubuntu 21.04 日本語 Remixがリリースされました・ディスクイメージのダウンロード - kledgeb") を導入した。
そういや Ubuntu 21.04 から Active Directory に対応したんだよね。
あとは MS Office を捨てれるのなら，次の LTS 版リリースあたりから本格的に Windows は要らなくなるんじゃないのかなぁ。

今回は前のマシンから環境をそのままコピってくればよかったので楽に移行できた。
開発環境も GCC と Go コンパイラを入れたくらいだが，とりあえず VSCode でブログが書けるようになったので，ここからは少しのんびりやる予定。
完全に移行が完了したら前のマシンは実験用（GnuPG 2.3 を試したりとか）に使うつもり。

[Hugo] が速くなってありがたい。
2Kページ以上フルビルドを2秒未満で処理できる。

```text
Start building sites … 

                   |  JA   
-------------------+-------
  Pages            | 2138  
  Paginator pages  |    0  
  Non-page files   |  532  
  Static files     |   59  
  Processed images |    0  
  Aliases          |    0  
  Sitemaps         |    1  
  Cleaned          |    0  

Total in 1385 ms
```

つか，本来はこのくらいのスピードなんだな。
これで私もよーやく人並みだよ。

## ブックマーク

- [ブログ: リーナス・トーバルズ、個人用PCのCPUをIntelから32コアAMD Ryzenに置き換え](https://okuranagaimo.blogspot.com/2020/05/pccpuintel32amd-ryzen.html)

[Hugo]: https://gohugo.io/ "The world’s fastest framework for building websites | Hugo"

## 参考

{{% review-paapi "B087745J5B" %}} <!-- 電源タップ -->
{{% review-paapi "B07TS24BFM" %}} <!-- USB切替器 -->
{{% review-paapi "B07H55YTSV" %}} <!-- 有線静音マウス -->
{{% review-paapi "B086X5S1TG" %}} <!-- CAT8 LANケーブル -->
{{% review-paapi "B08Q7TTY5Q" %}} <!-- 椅子 -->
{{% review-paapi "B01G4JSTYO" %}} <!-- 机 -->
{{% review-paapi "B07NDLJ6Q9" %}} <!-- テンキーレス キーボード -->
