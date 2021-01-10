+++
title = "Git GUI フロントエンドをたずねて三千里"
date = "2019-04-21T14:37:37+09:00"
description = "ブランチやタグや履歴を見ながらちょっと込み入った処理をする際はやはり GUI があるといいよねってことで。"
image = "/images/attention/kitten.jpg"
tags = [ "linux", "ubuntu", "git", "gnupg", "openssh" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

私は [Git] の GUI フロントエンドに [Git Extensions] を使っている。
個人的には FOSS の GUI フロントエンドではこれ以上のものはないと思っているが，残念なことに .NET Framework 上で動く製品なのである。
Windows プラットフォームではこれで何の問題もなかったが，どうにも[（.NET Framework の別実装である） Mono]({{< relref "./mono-in-ubuntu.md" >}} "Ubuntu に Mono を導入する") と相性が悪いようで，メジャーバージョンは上げられないしチョットしたことで落っこちてしまう。

というわけで Linux/[Ubuntu] で動作する GUI フロントエンドを探してみることした。
今回の評価ポイントは以下の通り。

1. コミット履歴をブランチ込みでツリー表示でき，簡単に操作できること
2. コミットやタグに対して電子署名を付与でき，かつ署名の検証ができること
3. [GnuPG] や [OpenSSH] を適切に扱えること
4. サブモジュールを簡単に扱えること
5. できれば FOSS 製品であること

VS Code や [ATOM] など最近はテキストエディタや IDE でグラフィカルに [git] を扱える製品も多く `commit`, `push`, `fetch`, `pull` といった基本機能を扱うだけなら GUI フロントエンドを別途用意する必要はないのだが，ブランチやタグや履歴を見ながらちょっと込み入った処理をする際はやはり GUI があるといいよねってことで。

## GUI フロントエンドいろいろ

### [Git]-gui

[Git] の[公式フロントエンド](https://git-scm.com/docs/git-gui "Git - git-gui Documentation")らしい。
APT でインストールできる。

Look&Feel は微妙だが機能的には悪くない。
ただ，タグへの署名と署名検証ができないっぽいんだよなぁ。

### [Giggle]

APT でインストールできる。

GNOME 用ということで Look&Feel は悪くないが，あまり複雑なことはできないようだ。
残念。

[Giggle]: https://wiki.gnome.org/Apps/giggle/ "Apps/giggle - GNOME Wiki!"

### [Gitg]

APT でインストールできる。

これも GNOME 用の製品だが，やっぱり複雑なことはできない感じ。
GNOME 用の製品ってみんなこんな感じなのか？

[Gitg]: https://wiki.gnome.org/Apps/Gitg/ "Apps/Gitg - GNOME Wiki!"

### [GitHub Desktop]

あまり複雑なことはできなくて残念という意味では [GitHub Desktop] もそうか。
元々は Windows および macOS 用だが Linux 向けの fork を公開している方がいるようだ。

- [shiftkey/desktop: Simple collaboration from your desktop](https://github.com/shiftkey/desktop)

[GitHub Desktop] は GitHub 上の Issue や pull request が使いやすくなるよう設計されている。
なので [git] 機能自体のサポートについてはイマイチな感じである。
「初心者向け」と言われるのも宜なるかなというところであろうか。

[GitHub Desktop] の今後には期待している。

[GitHub Desktop]: https://desktop.github.com/ "GitHub Desktop | Simple collaboration from your desktop"

### [GitEye]

FOSS ではないが「無料」で利用できる。
動作には別途 Java ランタイム（JRE）を用意する必要がある[^java1]。

[^java1]: JRE のインストールについては[「Ubuntu で遊ぶ」の OpenJDK インストールの項]({{< ref "/remark/2019/03/using-ubuntu.md#jdk" >}})を参照のこと。

Eclipse を連想させるプロジェクトベースの構成。
でも，これならいっそ Eclipse もしくは [IntelliJ IDEA](https://www.jetbrains.com/idea/ "IntelliJ IDEA: The Java IDE for Professional Developers by JetBrains") を使ったほうがいいんじゃないのか。

[GitEye]: https://www.collab.net/products/giteye "Free Git Client - GitEye | CollabNet VersionOne"

### [GitKraken]

プロプライエタリ・ライセンスで非商用のみ「無料」で利用できる。
Pro 版であれば4.08USD/月のサブスクリプション制で利用できる。

[Git] の GUI フロントエンドといえば大抵これが挙がるくらい有名。
てことは皆これにお金払ってるってことか？

Look&Feel は好みがあるので言及しないとして，機能自体は悪くないのだが [GnuPG] の扱いが雑。
なんでパスフレーズを入力させるのに自前の入力窓を使うかな。
独自の Pinentry を使ってるってわけでもないようだし，まさかパスフレーズをメモリ上に保持ってないよね？

あと，私の環境ではコミットやタグの署名検証ができなかった。
購入すれば見れるのかな？

[GitKraken]: https://www.gitkraken.com/ "Git Client - Glo Boards | GitKraken"

### [SmartGit] 【2021-01-09 訂正】

こちらもプロプライエタリ・ライセンスで非商用のみ「無料」で利用できる。
フルサポートで買うとかなり高い。
サブスクリプションを利用するなら5.99USD/月。

Look&Feel や機能は申し分ない。
以前は ssh 接続が上手くいかなくて諦めていたのだが

- [gpg-agent の設定： GnuPG in Ubuntu]({{< ref "/openpgp/gpg-agent-in-ubuntu.md" >}})

でちゃんと設定したら上手くいくようになった。
最新バージョンの OpenJDK でも問題なくし git も ssh も gpg も問題なく動く。

惜しむらくは署名の検証が出来ないことかな。
[GitKraken] もそうだが，署名はできるのに，なんで署名の検証をサボるかな。
意味ないよな。

[SmartGit]: https://www.syntevo.com/smartgit/ "SmartGit – Git Client for Windows, macOS, Linux"

## というわけで

[SmartGit] は GUI も好みに合ってるしいいかな，と思っている。
しばらく試してみて良さげならお金払ってもいいだろう。

あとは署名検証機能がなぁ。
もうしばらくは [Git Extensions] も手放せないか。

## 【おまけ】 CUI な [Tig]

GUI ではないがキャラクタベースの [tig] というのがあるらしい。
あちこちページを眺めていると，これを推す記事が結構多い。

サーバ等のリモートホストに対してキャラクタ端末でリポジトリにアクセスする場合は便利かもしれない。

[Tig]: https://jonas.github.io/tig/ "Introduction · Tig - Text-mode interface for Git"
[tig]: https://jonas.github.io/tig/ "Introduction · Tig - Text-mode interface for Git"

## ブックマーク

- [Git - GUI Clients](https://git-scm.com/downloads/guis)
- [Interfaces, frontends, and tools - Git SCM Wiki](https://git.wiki.kernel.org/index.php/Interfaces,_frontends,_and_tools)
- [Linuxで使えるGitクライアントを集めてみた](http://note.kurodigi.com/linux-gitclient/)
- [Ubuntu/Linuxで使えるGitのGUIクライアント(無料)まとめ - Qiita](https://qiita.com/solmin719/items/f174aab0fc73ddbc9cdf)

- [Git Commit で OpenPGP 署名を行う]({{< ref "/openpgp/git-commit-with-openpgp-signature.md" >}})

[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Git]: https://git-scm.com/
[git]: https://git-scm.com/
[Git Extensions]: https://gitextensions.github.io/ "Git Extensions | Git Extensions is a graphical user interface for Git that allows you to control Git without using the commandline"
[GnuPG]: https://gnupg.org/ "The GNU Privacy Guard"
[OpenSSH]: http://www.openssh.com/ "OpenSSH"
[ATOM]: https://atom.io/
