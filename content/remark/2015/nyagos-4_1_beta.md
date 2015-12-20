+++
date = "2015-12-20T16:31:06+09:00"
description = "4.1 で大きく変わったのはコールバック関数の扱いのようだ。"
draft = false
tags = ["windows", "tools", "nyagos", "shell", "lua"]
title = "NYAGOS 4.1-beta がリリース"

[author]
  avatar = "/images/avatar.jpg"
  facebook = "spiegel.im.spiegel"
  flattr = "spiegel"
  flickr = "spiegel"
  github = "spiegel-im-spiegel"
  instagram = "spiegel_2007"
  license = "by-sa"
  linkedin = "spiegelimspiegel"
  name = "Spiegel"
  tumblr = "spiegel-im-spiegel"
  twitter = "spiegel_2007"
  url = "http://www.baldanders.info/spiegel/profile/"
+++

ここのところ忙しくしてたら，いつの間にか [NYAGOS] の [4.1-beta](https://github.com/zetamatta/nyagos/releases/tag/4.1-beta) が出てた。
4.1 で大きく変わったのはコールバック関数の扱いのようだ。

- クラッシュ回避のため、全てのLua のコールバック関数はそれぞれの Lua インスタンスを持つようにした。（つまり、.nyagos で定義されたグローバル変数は、全てのコールバック関数から見ることができなくなった）
- コールバック関数と .nyagos 間で値を共有するため、テーブル share[] を作った

コールバック関数の挙動が変わったのは

{{< fig-quote title="nyagosスクリプト解説 - CMD.EXEで化けさせず、nyagosの中だけプロンプトをカラー化" link="http://qiita.com/zetamatta/items/c08586c85fa73c182a7a" >}}
<q>この制限は、クラッシュ回避のため、コールバック関数ごとに別の Lua インスタンスを用意しているためです。エイリアス等は別の goroutine 内で呼ばれるのですが、4.0 では、この時同一の Lua インスタンスを使用していたため、時にスタックに矛盾が発生して、クラッシュすることがあったのです。</q>
{{< /fig-quote >}}

ということらしい。

実は [NYAGOS] で表示するプロンプトは以下の記事を参考にオリジナルの `%PROMPT%` から変えている。

- [nyagosスクリプト解説 - CMD.EXEで化けさせず、nyagosの中だけプロンプトをカラー化 - Qiita](http://qiita.com/zetamatta/items/c08586c85fa73c182a7a)

以前のプロンプト定義は以下のような感じだった。

```lua
-- Simple Prompt for CMD.EXE
set{
    PROMPT='$L'.. nyagos.getenv('COMPUTERNAME') .. ':$P$G$_$$$s'
}

-- Coloring Prompt for NYAGOS.exe
local prompter=nyagos.prompt
nyagos.prompt = function(this)
    return prompter('$e[36;40;1m'..this..'$e[37;1m')
end
```

この中の `prompter` がコールバック関数から見えなくなったということらしい。
そこで以下のようにコードを変更するのだそうだ。

```lua
-- Simple Prompt for CMD.EXE
nyagos.env.prompt='$L'.. nyagos.getenv('COMPUTERNAME') .. ':$P$G$_$$$s'

-- Coloring Prompt for NYAGOS.exe
share.org_prompter=nyagos.prompt
nyagos.prompt = function(this)
    return share.org_prompter('$e[36;40;1m'..this..'$e[37;1m')
end
```

ポイントは `share.org_prompter=nyagos.prompt` の部分。
オリジナルの `nyagos.prompt()` 関数を `share[]` テーブルに退避させている。

これでめでたく

```
C:\program\nyagos>nyagos.exe
Nihongo Yet Another GOing Shell 4.1-beta-amd64 Powered by go1.5.2 & Lua 5.3
Copyright (c) 2014,2015 HAYAMA_Kaoru and NYAOS.ORG
<VENUS:C:/program/nyagos>
$ ls
Doc/               lua53.dll          nyagos.lua*        specialfolders.js*
catalog.d/         makeicon.cmd*      nyole.dll
license.txt        nyagos.d/          readme.md
lnk.js*            nyagos.exe*        readme_ja.md
<VENUS:C:/program/nyagos>
$
```

と表示できるようになった（カラーでお見せできないのが残念です）[^a]。

[^a]: そういや昔の UNIX ワークステーションはホスト名に venus とか惑星名を付けてるところが多かったな。私がネットワーク管理者をしてた時はプロキシサーバに janus とか付けてた。若気の至りである（笑）

実は私もプロンプトは折り返す派。
開発環境ではフォルダがかなり深くなることがあり，既存の `$P$G` では見づらいのだ。
まぁこれは Windows に限らないのだが。

[NYAGOS]: http://www.nyaos.org/index.cgi?p=NYAGOS "NYAOS.ORG - NYAGOS"
[Lua]: http://www.lua.org/ "The Programming Language Lua"
