+++
title = "GitHub Copilot はシェル芸達者"
date =  "2026-04-09T15:57:34+09:00"
description = "GitHub Copilot に少しずつ仕事を振ってみたら，記事の下ごしらえからタグ整理，デプロイ，Hugo 更新まで思った以上に任せられた話。"
isCJKLanguage = true
image = "/images/attention/kitten.jpg"
tags = [ "artificial-intelligence", "site", "hugo", "shell", "github", "tools", "vscode" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = false
+++

最初に予防線を張っておくと，この記事は技術系の内容ではありませんし，私は生成 AI の扱いにも慣れてません。
その辺を割り引いて読んでいただけるとありがたいです。

## 生成 AI にどこまで任せていいのか

最近は GitHub Copilot に開発支援だけじゃなくて，このブログのメンテナンスもさせている。
なにせ，どこまで任せていいのか匙加減がさっぱり分からないので，できそうなところから徐々に任せていく感じ。
気分は新卒社会人の OJT だぜ（笑）

今のところ GitHub Copilot に任せているのは以下の作業：

1. タイトル & slug の提案と記事ファイルの作成
2. 書いた記事の校正とタグの提案
3. 記事ファイルの commit & push
4. ブログのビルドと GitHub Pages への commit & push
5. [Hugo] の更新

記事そのものは書かせてない。

いや，生成 AI に書かせるとホンマに好き勝手な（私の意図とは程遠い）ことを書くのよ。
一度，タイトルだけ決めたまっさらな状態で書かせようとしたら「何の論文を書く気？」って感じのアウトラインを組み始めて，さすがに止めた。
ブログ記事に10個も章立てしてどうする（笑）

よくアフィリエイト記事を AI に書かせて云々みたいな話を聞くけど，ホンマに AI に書かせてるの？ かなり無茶苦茶するぞ。
ある意味，文才があるのかも知れないが。

というわけで「記事を書く」以外の作業をやらせている。

ちなみにモデルは GPT-5.3-Codex を使っている。
あまり Premium Request を消費したくないので最初は GPT-5 mini を使ってたのだが，どうもこの手の「作業」には向かない感じ。
最終的にはローカル LLM を構築したいのだが，お試しレベルならともかく，まともに動くモデルを個人環境で構築するのは難しいか？

## タイトル & slug の提案と記事ファイルの作成

私は最初にタイトルを決めないとブログ記事が書けない。
マイクロブログ程度の短文ならともかく，ブログ記事でタイトルも決めずに書き始めると大抵記事が迷走する（タイトルを決めて書いても迷走しがちなのだが，それはそれ）。

手順としては

1. 日本語のタイトルを決める（ここだけ私の作業）
2. Copilot に英語の翻訳と対応する slug を提案させる（たいてい複数個の候補を提案してくれる）
3. slug が決まったら記事ファイルを作成させる
4. ファイル内の front matter を埋めさせる

という感じ。
日本語タイトルを英訳させるための対話で，色々と気付きがあって面白い。

たとえば日本語タイトルが曖昧で英訳が頓珍漢な文になることがある。
これは生成 AI 以前の機械翻訳サービスを使ってたときもそうだったが，機械と対話することで「文のこの部分が誤解を招いてるのか」と気づいたりできるわけだ。
こういったやり取りが意外と面白かったりする。

記事ファイルの作成用にシェルスクリプトを組んでいるのだが，まずその挙動を理解させるところから始めた。
したら「このスクリプトをリファクタリングするか？」と訊いてきたのでやらせてみたら，なかなかいい感じに仕上げてくれた。
これなら AI にスクリプトを改善させることもできそうだ。

他の既存のスクリプトについても同様にリファクタリングさせ，それらの使い方を `copilot-instructions.md` および `README.md` にまとめさせた。

## 書いた記事の校正とタグの提案

私のブログは意図的に文体を崩して書いているので（設計書やビジネス文書じゃあるまいし），下手に賢い lint では却って使いづらかったりする。
ためしに Copilot に「誤字だけ指摘して」と指示してみたら，いい感じに結果を返してくれた。
以来，校正は Copilot に任せている。

誤字だけでなく慣用句の漢字間違いの指摘とかは結構ありがたかったり。
それでいて崩した文体については一切スルーしてくれる（笑） 偉い偉い。

もうひとつ悩んでいるのが記事に付与するタグの選択だ。
これを普通に AI に任せるとメチャクチャにしてくれるので，まずは過去の記事を全て攫ってタグの一覧を作らせるところから始めた。

したらすごい勢いでシェルスクリプトを書き始めた。
私はそれを「へぇ。なるほど」とか言いながら眺めるだけ。
アレだ。
ペアプログラミングってやつ（笑） 機械相手にペアプロすることになるとは思わなかった。

せっかくいい感じのスクリプトを書いてくれたので，それをファイルにまとめていつでも呼び出せるようにさせた。
タグの一覧を出現数の多いものからソートして CSV ファイルに出力させる。
ここまでできれば，作成した記事に対してタグ一覧の中から妥当なタグを提案するよう指示すればよい。

タグ一覧を作ったメリットはもうひとつあって typo がいくつか見つかったのと似たような意味のタグの重複が見つかった。
見つけたのも Copilot だけど（なにせ10年以上の間にタグが450個以上できてるのだ）。
これらも修正方針を提案させた上で修正させた。
ファイル作業に関してはちゃんと手順を踏めば間違えることなくやってくれる。

## 記事ファイルの commit & push

記事ファイルの commit に関しては，念のため，最初に commit message を提案させて，私が確認してから実際の作業を行わせている。
ここで GPT-5 mini と GPT-5.3-Codex の違いが大きく出た。

いや，ブログ記事の追加なんだから {{% ruby "君" %}}AI{{% /ruby %}} のしたことを書くんじゃなくて，どんな記事を追加したのかを書けよ！

というわけで，GPT-5 mini は向いてないという結論になった。
commit & push するためのコマンドラインの組み立て自体はどちらも大差ないんだけどねぇ。

## ブログのビルドと GitHub Pages への commit & push

ブログのビルドと GitHub Pages への commit & push についても以前からスクリプトファイルを作っていたので，それをそのまま Copilot に使わせている。
まぁ，自分でやったほうが早いのだが，作業の一連を任せることに意味があると思ってるので，敢えて任せている。
この辺も OJT っぽいよな（笑）

指示も最初はできるだけ具体的にしていたが，最近は「デプロイして」だけでも（過去の作業履歴や `copilot-instructions.md` の記述から）適切に処理してくれるようになった。

## Hugo の更新

このブログは静的サイトジェネレータである [Hugo] を使って構築しているのだが，未だ v1 に到達せずベータ版のままである。
古いテンプレート変数や関数のいくつかは deprecated になっているし，稀に破壊的変更が起きることもある。
[Hugo] のバージョンアップのたびにビルドし直して互換性等の確認をしているのだが，だんだん面倒くさくなってきたので，これも Copilot に任せてみることにした。

今回もゴリゴリとスクリプトを書き始める AI。
私も前と同じく様子を眺めていたが，出来がよさそうなのでスクリプトファイルとして書き出してもらった。
こんな感じ（長いぞ！）。

```bash
#!/usr/bin/env bash
set -euo pipefail

# Install latest Hugo Extended .deb from GitHub Releases using apt.
# Usage: ./hugo_inst.sh

REPO="gohugoio/hugo"
API_URL="https://api.github.com/repos/${REPO}/releases/latest"
WORK_DIR="${TMPDIR:-/tmp}/hugo-inst.$$"

cleanup() {
  rm -rf "$WORK_DIR"
}
trap cleanup EXIT

mkdir -p "$WORK_DIR"
cd "$WORK_DIR"

arch="$(dpkg --print-architecture)"
case "$arch" in
  amd64|arm64)
    ;;
  *)
    echo "Unsupported architecture: $arch" >&2
    echo "This script supports: amd64, arm64" >&2
    exit 1
    ;;
esac

echo "Current Hugo:"
if command -v hugo >/dev/null 2>&1; then
  hugo version || true
else
  echo "hugo command not found"
fi

echo "Fetching latest release metadata from ${REPO} ..."
release_json="$(curl -fsSL "$API_URL")"

version="$(awk -F '"' '/"tag_name"/ {print $4; exit}' <<<"$release_json")"
if [[ -z "$version" ]]; then
  echo "Failed to read latest release version." >&2
  exit 1
fi

asset_url="$(awk -F '"' '/"browser_download_url"/ && /hugo_extended_.*_linux-'"$arch"'\.deb/ {print $4; exit}' <<<"$release_json")"
if [[ -z "$asset_url" ]]; then
  echo "Failed to find hugo_extended .deb for architecture: $arch" >&2
  exit 1
fi

deb_file="${asset_url##*/}"

echo "Latest release: ${version}"
echo "Downloading: ${deb_file}"
curl -fL --retry 3 --retry-delay 2 -o "$deb_file" "$asset_url"

# Optional checksum verification when checksums.txt exists in latest release.
checksums_url="$(awk -F '"' '/"browser_download_url"/ && /checksums\.txt/ {print $4; exit}' <<<"$release_json")"
if [[ -n "$checksums_url" ]]; then
  echo "Downloading checksums.txt for verification ..."
  curl -fL --retry 3 --retry-delay 2 -o checksums.txt "$checksums_url"

  if awk '{print $2}' checksums.txt | grep -qx "$deb_file"; then
    echo "Verifying checksum ..."
    grep "  $deb_file$" checksums.txt | sha256sum -c -
  else
    echo "checksums.txt does not include ${deb_file}; skipping verification."
  fi
else
  echo "checksums.txt not found in latest release; skipping verification."
fi

echo "Installing ${deb_file} with apt ..."
sudo apt install -y "./$deb_file"

echo "Installed Hugo:"
hugo version

echo "Done."
```

手順としては

1. インストールするマシンのアーキテクチャを確認して対応する `.deb` ファイルをダウンロードする
2. 可能ならチェックサムの検証も行う
3. ダウンロードした `.deb` ファイルを `apt install` コマンドでインストールする

という感じ。

これを実行して [Hugo] を更新した後，ブログのビルドを行い，ビルド前後の差分を調べて問題がないことを確認する。
問題がなければビルドしたものを commit & push させる。

「問題がないことを確認する」の部分は私が判断しているので完全な自動化ではないが，今のところ問題は起きていない。

## GitHub Copilot はシェル芸達者

という感じで緩々と GitHub Copilot を使っている。

今回の一連で意外とシェル芸達者なのが分かったので，自宅機に [CLI 版](https://github.com/features/copilot/cli "GitHub Copilot CLI · GitHub")を入れようかなぁ。
でも，[結城浩](https://social.hyuki.net/@hyuki "結城浩 / Hiroshi Yuki (@hyuki@social.hyuki.net) - 結城浩のマストドン")さんほど使い倒す勇気はないんだよなぁ。
それよりもローカル LLM 構築に向けて環境を整えるほうが先か？

[Kagi Search]: https://kagi.com/ "Kagi Search - A Premium Search Engine"
[Kagi Translate]: https://translate.kagi.com/ "Kagi Translate"
[Kagi Assistant]: https://assistant.kagi.com/ "Kagi Assistant"

[Hugo]: https://gohugo.io/ "The world's fastest framework for building websites"

## 参考

{{% review-paapi "B07TSZZPWN" %}} <!-- フルスクラッチから1日でCMSを作る_シェルスクリプト高速開発手法入門 -->
