# ATOM の Theme / Package の感想文（2015-06-10）

[ATOM] Editor を定常的に使うようになって2週間程度かな。現時点で使ってる Theme と Package の感想を書いておく。
[ATOM] について書いた記事は「[ATOM Editor をそろそろ始めようか](http://qiita.com/spiegel-im-spiegel/items/3d41d98dacc107d73431)」をどうぞ（ブックマークもこちら）。ちなみに Windows 環境。

## Themes

```shell
C>apm list -i -b -t
atom-monokai@0.8.0
 → Theme にこだわりはないのだが、atom-monokai は UI Theme に反応して自動的に Dark と Light を切り替えてくれるので便利に使っている。
```

## Packages

```shell
C:>apm list -i -b -p
autoclose-html@0.17.1
→ 私はブログ記事を書くときは HTML べた書きなので、まぁまぁ助かる。
editorconfig@1.0.0
→ 仕事で ATOM を使うときは必須。
file-icons@1.5.7
→ 見栄えは大事。
git-control@0.3.0
→ git-plus が手に馴染んだら使わなくなった。 submodule を上手く処理してくれないのもマイナス。
git-plus@5.1.0
→ 手に馴染んだらとても便利なことに気が付いた。 GUI が要るなら GIT Extensions か SourceTree 併用で。
highlight-line@0.10.2
→ Windows キャレットをよく見失うのよ。これないと作業効率が下がる。
japan-util@0.1.1
→ 全角の英数字や半角カナを使う馬鹿が後を絶たないので必要。
japanese-wrap@0.2.10
→ 日本語環境では必須なのだが、たまにうまく動かないんだよなぁ。特に Markdown 編集時。
open-recent@2.2.4
→ 最近開いたファイルとフォルダを覚えてくれる賢いやつ。てか、何故これを Core に入れないのだ。
show-ideographic-space@1.0.0
→ 全角空白を可視化する。コード書きには必須アイテム。
symbols-tree-view@0.9.3
→ アウトライン解析。コード書きには必須アイテム。
```

[ATOM]: https://atom.io/ "Atom"
