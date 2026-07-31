+++
title = "APT で upgrade/dist-upgrade できなかったとき"
date =  "2022-09-25T13:33:56+09:00"
description = "apt install で無理やりアップグレードすればいいらしい"
image = "/images/attention/kitten.jpg"
tags = [ "ubuntu", "tools", "install" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

実は今回で2度目。
たぶん3度目もあるだろうということで，覚え書きとして残しておく。
つっても大した話ではないが。

APT (Advanced Package Tool) で `apt upgrade` する際「以下のパッケージは保留されます」みたいな感じで保留されることがある。
そういうときは大抵 `apt dist-upgrade` で何とかなると聞いているのだが（本来はディストリビューションのアップグレードで使うコマンド），それでも上手くアップグレードしてくれない場合がある。

諦めてググってみたら「[結局`apt install`で明示的に指定すればいいみたいです](https://kazuhira-r.hatenablog.com/entry/2021/01/10/184337 "Ubuntu Linuxでapt upgradeで保留されたパッケージがあった場合に、アップデートするには？ - CLOVER🍀")」という身も蓋もない解決策を見かけて「ええのん？」と思ったが，他に方法も見当たらないので思い切って `apt install` でアップグレードした。
なお `apt install` を使う場合は対象のパッケージを明示する必要がある。

いや，流石に xwayland や systemd がアップグレードされないのは拙いだろう，ということで。
なんか NAS に SAMBA 経由でアクセスできなくなってたし。

`apt install` で無理やりアップグレードして再起動（厳密にはシャットダウンからのコールド・スタート）後 SAMBA アクセスもちゃんとできるようになった。
まずはめでたし。
ホンマにええんかなぁ...

## 【2022-10-16 追記】 Phased Updates

詳しい[解説記事](https://kledgeb.blogspot.com/2022/10/ubuntu-2204-328-apt.html "Ubuntu 22.04 その328 - aptコマンドとフェーズドアップデートによるパッケージアップデートの保留 - kledgeb")を見つけた。
これによると

{{< fig-quote type="markdown" title="aptコマンドとフェーズドアップデートによるパッケージアップデートの保留" link="https://kledgeb.blogspot.com/2022/10/ubuntu-2204-328-apt.html" >}}
元々フェーズドアップデートは「Ubuntu Desktop 13.04」の「ソフトウェアの更新」で導入された機能です。<br>
その後APT全体でこの機能が有効化され、「Ubuntu 22.04 LTS」でも引き続き有効になっています。

フェーズドアップデートは、特定のパッケージのアップデートをユーザーに段階的に提供する仕組みです。<br>
言い換えれば一度に全ユーザーにアップデートを提供するのではなく、アップデートを提供するユーザーを絞りながら、徐々に提供対象のユーザーを増やしていく機能です。
{{< /fig-quote >}}

らしい。
ということは放っておいてもいずれはアップデートされるということか。

保留されているパッケージが phased update の対象かどうかは以下のコマンドで分かるそうな。

```text
$ apt policy <package name>
```

今度見かけたら試してみよう。
でも xwayland や systemd のアップデートが保留されるのは勘弁して欲しいのだが...

## 参考

{{% review-paapi "4295013498" %}} <!-- Linuxシステムの仕組み -->
