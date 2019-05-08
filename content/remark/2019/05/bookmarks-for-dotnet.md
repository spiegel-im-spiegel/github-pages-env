+++
title = ".NET 統合に関するブックマーク"
date =  "2019-05-08T22:36:09+09:00"
description = "ノンビリと成り行きを見守ることにしよう。"
image = "/images/attention/kitten.jpg"
tags = [ "engineering", "dot-net", "mono", "bookmark" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

いよいよ Microsoft による .NET 統合の動きが本格化しそうである。
めでたい！

- [Introducing .NET 5 | .NET Blog](https://devblogs.microsoft.com/dotnet/introducing-net-5/)

個人的には [Mono] 上で動くアプリケーションがどうなるか気になるところであるが

{{< fig-quote title="Introducing .NET 5" link="https://devblogs.microsoft.com/dotnet/introducing-net-5/" lang="en" >}}
<q><a href="https://github.com/mono/mono">Mono</a> is the original cross-platform implementation of .NET. It started out as an open-source alternative to .NET Framework and transitioned to targeting mobile devices as iOS and Android devices became popular. Mono is the runtime used as part of Xamarin.<br>
[...]<br>
Taken together, the .NET Core and Mono runtimes have a lot of similarities (they are both .NET runtimes after all) but also valuable unique capabilities. It makes sense to make it possible to pick the runtime experience you want. We’re in the process of making CoreCLR and Mono drop-in replacements for one another. We will make it as simple as a build switch to choose between the different runtime options.</q>
{{< /fig-quote >}}

 .NET 5 のリリース自体は来年（2020年）なのでノンビリと成り行きを見守ることにしよう。
[Go 言語]か Kotlin が .NET をサポートすれば完璧なのだが（笑）

以降では .NET 5 関連のブックマークを挙げていく。
随時更新する予定なのであしからず。

- [［速報］オープンソースの「.NET 5」がすべての.NETを引き継ぐ。.NET Frameworkと.NET CoreとXamarinは「.NET 5」に。Microsoft Build 2019 － Publickey](https://www.publickey1.jp/blog/19/net_5netnet_frameworknet_corexamarinnet_5microsoft_build_2019.html)
- [.NET Core 3.0 の目玉の single exe を preview5 で試してみた - Qiita](https://qiita.com/c-yan/items/a581c165ca1b6a5536ca)
- [[.NET Core その63 - .NET Coreの将来と.NET 5・.NET Frameworkから.NET Coreへ - kledgeb](https://kledgeb.blogspot.com/2019/05/net-core-63-net-corenet-5net.html)]

- [Ubuntu に Mono を導入する]({{< ref "/remark/2019/04/mono-in-ubuntu.md" >}})

[Mono]: https://www.mono-project.com/
[Go 言語]: https://golang.org/ "The Go Programming Language"
