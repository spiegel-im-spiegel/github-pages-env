# ATOM 1.0 リリースおめでたう記念に最初からインストールし直してみた

まずは [ATOM] 1.0 リリースおめでとうございます。

- [Atom 1.0](http://blog.atom.io/2015/06/25/atom-1-0.html)
- [Atom 1.0 リリース - Qiita](http://qiita.com/tomoyaton/items/f4be24cc7152ce8b4107)

あらかじめ [ATOM] がインストールされていれば起動時に勝手にアップグレードされるのだが，いろいろお試しでパッケージを入れたり消したりしてたのでフォルダが随分汚くなってしまった。そこで，この際，最初から [ATOM] をインストールし直してみることにした。あっ，ちなみに Windows 環境ね。

## ATOM をアンインストールする

まずは [ATOM] をアンインストールしてフォルダを掃除する。

アンインストール自体はコンパネから「プログラムと機能」を開いて「Atom」をアンインストールする。アンインストールを開始するもほぼ無言で完了。男前（笑）

ただしユーザのフォルダ内には [ATOM] 関連のファイルがかなり残っているので手動で掃除する。対象となるのは以下のフォルダ。

- `C:\Users\username\.atom`
- `C:\Users\username\AppData\Local\atom`
- `C:\Users\username\AppData\Local\Temp`
- `C:\Users\username\AppData\Roaming\Atom`

`AppData` フォルダは既定では不可視になっているのでご注意を。 `C:\Users\username\.atom` フォルダには `keymap.cson` などの設定ファイルが入ってるので，バックアップを取っておくと安全。

`C:\Users\username\AppData\Local\Temp` には `Atom Crashes` フォルダがある。どうやらクラッシュ・レポートはここに吐かれるらしい。テンポラリ・フォルダにある古い日付のフォルダ・ファイルは，大概は削除して大丈夫なのだが，たまにヤバいやつもあるので掃除は慎重に。

## 改めて ATOM をインストールする

インストールは今までと変わりなく。 [ATOM] サイトからインストールパッケージ `AtomSetup.exe` をダウンロードして起動すればよい。前にインストールしたときはエクスプローラからのクリックでうまくいかなかったのだが（コマンドプロンプトから起動したらうまくいった），今回は無問題。

インストールが成功するとインストールフォルダ `C:\Users\username\AppData\Local\atom\bin` に PATH が通る。これでコマンドプロンプトから `atom` および `apm` コマンドが使えるようになる。

```shell
C:>atom -v
[6568:0626/113656:INFO:CONSOLE(0)] 1.0.0

C:>apm -v
apm  1.0.1
npm  2.5.1
node 0.10.35
python
git 1.9.5.msysgit.0
visual studio
```

上記の環境では Python と Visual Studio は入れてないのでバージョンが入ってないのかな。
たしか node-gyp をビルドするのに（Windows 環境では） Python と Visual Studio が要るんだよね。ううむ。

- [Windowsでnode-gypのビルドを通す - なにか作る](http://create-something.hatenadiary.jp/entry/2014/07/13/021655)

### Proxy の設定

Intranet 上のマシンで外部との接続が阻まれている場合は Proxy 設定を行う。設定には apm コマンドを使う。

```shell
C:>apm config set https-proxy http://username:password@proxy.exsample.com:8080
```

（イマドキは Intranet が安全とは限らないので URL にパスワードを含めるのはどうかと思うんだけど，まぁしょうがない）

### テーマ・パッケージのインストール

テーマやパッケージについては拙文「[ATOM Editor をそろそろ始めようか](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)」を参考にどうぞ。 keybind や style については [Gist に設定を貼り付けている](https://gist.github.com/spiegel-im-spiegel/e6e9c7340987f1607b2c)。併せてどうぞ。

ちみに今回入れたテーマ・パッケージ（と一言感想）は以下の通り。

```shell
C:>apm list -i -b -t
atom-monokai@0.8.0
 → Theme にこだわりはないのだが、atom-monokai は UI Theme に反応して自動的に Dark と Light を切り替えてくれるので便利に使っている。
```

```shell
C:>apm list -i -b -p
autoclose-html@0.18.0
 → 私はブログ記事を書くときは HTML べた書きなので、まぁまぁ助かる。
editorconfig@1.0.0
 → 仕事で ATOM を使うときは必須。
file-icons@1.5.8
 → 見栄えは大事。
git-plus@5.2.1
 → 手に馴染んだらとても便利なことに気が付いた。 GUI が要るなら GIT Extensions か SourceTree 併用で。ちなみに git-control は外してしまった。
highlight-line@0.11.0
 → Windows キャレットをよく見失うのよ。これないと作業効率が下がる。
japan-util@0.1.1
 → 全角の英数字や半角カナを使う馬鹿が後を絶たないので必要。
japanese-wrap@0.2.7
 → 日本語環境では必須なのだが、たまにうまく動かないんだよなぁ。特に Markdown 編集時。ちなみにこれをインストールしようとしたら “Version 0.2.7 is not the latest version available for this package.” と警告が出た。ただし使用上は問題ない。
open-recent@2.2.4
 → 最近開いたファイルとフォルダを覚えてくれる賢いやつ。てか、何故これを Core に入れないのだ。
quick-highlight@0.1.6
 → 任意の単語をマークできる。結構便利。 toggle を f5 に割り当てて使っている。
show-ideographic-space@1.0.1
 → 全角空白を可視化する。コード書きには必須アイテム。
symbols-tree-view@0.9.3
 → アウトライン解析。コード書きには必須アイテム。
```

### apm stars でパッケージを一気にインストールする

`apm` には star を付けたテーマ・パッケージを一気にインストールするコマンドがある。

この機能を使うには，まず `apm` にアカウントのトークンを登録する必要がある。アカウントのトークンは [Account](https://atom.io/account) ページから取得できる。（[GitHub] のアカウントを持っていれば，そのまま [ATOM] にも sign in できるのだが，持ってない人はどうするんだろう？）

取得したトークンを `apm login` コマンドで登録すれば OK。

```shell
C:>apm login
Welcome to Atom!

Before you can publish packages, you'll need an API token.

Visit your account page on Atom.io https://atom.io/account,
copy the token and paste it below when prompted.

Press [Enter] to open your account page on Atom.io.
Token> ****************
Saving token to Keychain done
```

Star を付けたテーマ・パッケージは `apm stars` コマンドで見ることができる。

```shell
C:>apm stars
Packages starred by you (12)
├── atom-monokai Atom syntax theme inspired by monokai. Works with Atom Dark, Light and Seti UI. (15492 downloads, 20 stars)
├── autoclose-html Automates closing of HTML Tags (57086 downloads, 184 stars)
├── editorconfig Helps developers maintain consistent coding styles between different editors (29681 downloads, 301 stars)
├── file-icons Assign file extension icons and colours for improved visual grepping (172513 downloads, 1058 stars)
├── git-plus Do git things without the terminal (160434 downloads, 574 stars)
├── highlight-line Highlights the current line in the editor (40060 downloads, 294 stars)
├── japan-util utilities for Japanese (213 downloads, 5 stars)
├── japanese-wrap Word wrap for Japanese text (37025 downloads, 234 stars)
├── open-recent Open recent files in the current window, and recent folders (optionally) in a new window. (5777 downloads, 78 stars)
├── quick-highlight Highlight text quickly. (189 downloads, 4 stars)
├── show-ideographic-space Show ideographic space (known as 全角スペース) (2035 downloads, 33 stars)
└── symbols-tree-view A symbols view like taglist (12046 downloads, 118 stars)

Use `apm stars --install` to install them all or visit http://atom.io/packages to read more about them.
```

さらに `--install` オプションを付ければ一気にインストールできる。

```shell
C:>apm stars --install
Installing atom-monokai to C:\Users\username\.atom\packages done
Installing autoclose-html to C:\Users\username\.atom\packages done
Installing editorconfig to C:\Users\username\.atom\packages done
Installing file-icons to C:\Users\username\.atom\packages done
Installing git-plus to C:\Users\username\.atom\packages done
Installing highlight-line to C:\Users\username\.atom\packages done
Installing japan-util to C:\Users\username\.atom\packages done
Installing japanese-wrap to C:\Users\username\.atom\packages done
Installing open-recent to C:\Users\username\.atom\packages done
Installing quick-highlight to C:\Users\username\.atom\packages done
Installing show-ideographic-space to C:\Users\username\.atom\packages done
Installing symbols-tree-view to C:\Users\username\.atom\packages done
```

これで複数マシンへの環境構築が随分楽になると思う。なお star の管理は `apm star` または `apm unstar` コマンドでできるが，テーマ・パッケージのページでも可能。

（[ATOM] の star が [GitHub] の star のように [Flattr](https://flattr.com/) と連動すれば面白いんだけどねぇ。とりあえず flattr ボタンを貼り付ける手もあるけど）

## 関連する記事

- [ATOM Editor をそろそろ始めようか - Qiita](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431) : この記事については今後もメンテする予定
- [行末の空白は EditorConfig で始末しましょう - Qiita](http://qiita.com/spiegel-im-spiegel/items/a1b4d1ad2a6693ae33e4)
- [ATOM の Theme / Package の感想文（2015-06-10） - Qiita](http://qiita.com/spiegel-im-spiegel/items/115fea37ad2e515f6641)

[ATOM]: https://atom.io/ "Atom"
[GitHub]: https://github.com/ "GitHub"
