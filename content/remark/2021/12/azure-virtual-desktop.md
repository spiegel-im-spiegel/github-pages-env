+++
title = "ようやく Azure Virtual Desktop を導入できた"
date =  "2021-12-05T23:16:02+09:00"
description = "最初の頃はデプロイが失敗しまくりで往生したですよ。"
image = "/images/attention/kitten.jpg"
tags = [ "windows", "azure", "install", "tools" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

よーやく [Azure Virtual Desktop] を導入できたよ。
Windows 11 が使えるということで個人用の Azure アカウントで試そうとしたんだけど，何だか情報が断片的で分かりにくくてさ。

手順を覚えているうちに覚え書きとして残しておく。
なお，法人ユーザの場合は [Windows 365] のほうがオススメである。
今回はあくまで個人ユーザの場合ね。

## [Azure Active Directory] を確認する

個人で Azure アカウントをとった場合でもフリー版の [Azure Active Directory] が用意されている。
[Azure Virtual Desktop] を構成するためには [Azure Active Directory] をチェックしてユーザ・プリンシパル名（User Principal Name; UPN）を確認してサインイン可能にしておく必要がある。

注意しないといけないのは Azure アカウントを取得したときのアカウント名と UPN が異なっている（場合がある）という点。
これに気が付かなくてしーばらく悩んだよ。

## [Azure Virtual Network] を構成する

[Azure Virtual Desktop] を構成する際に [Azure Virtual Network] を指定する必要がある。
これも気が付かなくて，何度もやり直す羽目になった。
Microsoft Docs には色々書かれているが，単純に [Azure Virtual Desktop] を使いたいだけなら BastionHost も Firewall も不要である[^bh1]。

[^bh1]: 仮想マシン間でやり取りしたいなら BastionHost を設定したほうがいいかも。外部から SSH で入る場合も必要かな。

- [クイック スタート: 仮想ネットワークを作成する - Azure portal - Azure Virtual Network | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/virtual-network/quick-create-portal)

## ホストプールを先に作る

さて，いよいよ [Azure Virtual Desktop] の作成を始めるのだが

{{< fig-img src="./azure-virtual-desktop-1.png" link="./azure-virtual-desktop-1.png" width="1003" >}}

「作業の開始」ではなく「ホスト プールの作成」から始めるとよい。
この中で「ホストプール」「仮想マシン」「ワークスペース」を決めていくのが効率がいいようだ。
最初の頃は「作業の開始」から始めて，どうやってもデプロイが失敗するので往生したですよ。

なお，ホストプールだけ先に作っておいて仮想デスクトップとワークスペースは後から追加する手もある。
私は仮想デスクトップを試行錯誤でいくつか潰してしまった（笑）

仮想デスクトップを構成する際に

{{< fig-img src="./azure-virtual-desktop-2.png" link="./azure-virtual-desktop-2.png" width="732" >}}

てな感じに「仮想マシンの管理者アカウント」を決めるのだが，この情報はどこかに控えておくこと。
最初に仮想デスクトップにログインする際はこの管理者アカウント情報が必要になる。
いや，うっかりアカウント情報を紛失しまって仮想デスクトップを無駄にしてしまったのよ[^reset1]。

[^reset1]: 仮想デスクトップのアカウント名だけでも分かっていれば Azure 側からパスワードリセットすることは可能。

なお，仮想デスクトップを追加する際にはホストプールの「セッション ホスト」で「追加」すればよい。

ところで，作成した仮想デスクトップを見ると

{{< fig-img src="./azure-virtual-desktop-3.png" link="./azure-virtual-desktop-3.png" width="502" >}}

ってなってるんだけど Windows 11 って 10 のバリエーションってことでOK？

作成した仮想デスクトップに対して [Azure Active Directory] のユーザを割り当てる。
基本的にひとつの仮想デスクトップに対して1ユーザを割り当てる。

- [Azure Virtual Desktop の個人用デスクトップの割り当ての種類 - Azure | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/virtual-desktop/configure-host-pool-personal-desktop-assignment-type)

ちゃんと割り当てておかないとログイン時に割り当てられたリソースがないって怒られる。

## Web クライアントで接続する

作成した仮想デスクトップにつなぐには各種クライアント・ツールが使えるのだが

- [Azure Virtual Desktop のユーザー ドキュメント | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/virtual-desktop/user-documentation/)

Linux 環境では Web クライアントを使うしかないみたい。

- [Web クライアントを使用して Azure Virtual Desktop に接続する - Azure | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/virtual-desktop/user-documentation/connect-web)

接続先は「[Azure Resource Manager 統合バージョンの Azure Virtual Desktop Web クライアント](https://rdweb.wvd.microsoft.com/arm/webclient)」でいいようだ。
サインインは仮想デスクトップを割り当ててている [Azure Active Directory] の UNP でおこなう。
以下の画面が表示されればサインイン成功。

{{< fig-img src="./azure-virtual-desktop-4.png" link="./azure-virtual-desktop-4.png" width="500" >}}

“SessionDesktop” を選択すると

{{< fig-img src="./azure-virtual-desktop-5.png" link="./azure-virtual-desktop-5.png" width="623" >}}

などとプロンプトが出るので適当に「許可」するとようやく

{{< fig-img src="./azure-virtual-desktop-6.png" link="./azure-virtual-desktop-6.png" width="621" >}}

ログイン画面になる。
ここで「仮想マシンの管理者アカウント」を使ってログインすればよい。

{{< fig-img src="./azure-virtual-desktop-7.png" link="./azure-virtual-desktop-7.png" width="1920" >}}

よーし。
よーやくここまできたよ。

試しに Edge ブラウザで Google のページを開いてみる。

{{< fig-img src="./azure-virtual-desktop-8.png" link="./azure-virtual-desktop-8.png" width="1920" >}}

おー。
これが噂の[ブラウザ勧誘合戦](https://japanese.engadget.com/microsoft-edge-try-stop-user-download-chrome-050051364.html "マイクロソフトEdge、ユーザーにChromeのダウンロードを止めるよう呼びかける - Engadget 日本版")か。
アホだ（笑）

というわけで，やっと Windows 11 で遊び倒す準備ができた。
まずは大量の Windows Update をどうにかしないと（笑）

ちなみにめちゃめちゃ重い。
普段遣いにはできないな，これ。

## [Azure Virtual Desktop] のコスト

2021-12-31 までは無料で [Azure Virtual Desktop] を利用できる。
以降については以下のサイトが参考になるかな。

- [Azure Virtual Desktop の価格 | Microsoft Azure](https://azure.microsoft.com/ja-jp/pricing/details/virtual-desktop/)
- [Azure Virtual Desktop のデプロイ コストの合計について - Azure | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/virtual-desktop/remote-app-streaming/total-costs)

[Azure Virtual Desktop]: https://docs.microsoft.com/ja-jp/azure/virtual-desktop/ "Azure Virtual Desktop のドキュメント | Microsoft Docs"
[Windows 365]: https://www.microsoft.com/ja-jp/windows-365 "Windows 365 クラウド PC | Microsoft"
[Azure Active Directory]: https://docs.microsoft.com/ja-jp/azure/active-directory/fundamentals/active-directory-whatis "Azure Active Directory とは - Azure Active Directory | Microsoft Docs"
[Azure Virtual Network]: https://docs.microsoft.com/ja-jp/azure/virtual-network/ "Azure Virtual Network のドキュメント - チュートリアル、クイックスタート、API リファレンス | Microsoft Docs"
