+++
title = "Ubuntu に Visual Studio Code を導入する"
date =  "2020-05-28T12:31:57+09:00"
description = "継続的に使うなら今後も追記するつもり。"
image = "/images/attention/kitten.jpg"
tags = ["vscode", "editor", "ubuntu", "terminal"]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

[ちうわけで]({{< ref "/remark/2020/05/x-terminal-with-atom.md" >}} "とりあえず ATOM エディタ内ターミナルを x-terminal に乗り換えた") [Visual Studio Code][VS Code]を導入して試してみることにした。

継続的に使うなら今後も追記するつもり。

## [VS Code] のインストール

[Ubuntu] に [VS Code] をインストールするのであれば [Snap] を使うのが一番簡単なのだが，うちの環境では何故か [Snap] 版で日本語入力ができないという致命的な問題があるため，今回は公式サイトから deb ファイルをダウンロードして直接 `apt install` した。

インストールが成功すれば以下のようにコマンドが起動するはず。

```text
$ code -v
1.45.1
5763d909d5f12fe19f215cbfdd29a91c0fa9208a
x64
```

よしよし。
ちゃんと最新バージョンだな。

実際に [VS Code] を起動してアプリケーション情報（メニューの `Help` → `About` を選択）も確認してみる。

{{< fig-img src="./about-vscode.png" link="./about-vscode.png" width="524" >}}

うむうむ。

## [VS Code] のカスタマイズ

カスタマイズはメニューの `File` → `Preferences` から選択できる。

- `File` → `Preference` → `Settings` で設定画面が開く（または `Ctrl+,` キー押下）
- `File` → `Preference` → `Keyboard Shortcuts` でキー割当の確認・変更ができる
    - 既定のキー割当は [ATOM] とほとんど同じなので悩むところは少なかった
    - `File` → `Preference` → `Keymaps` で他エディタの割当を導入することもできる
- `f1` キー押下でもコマンド・パレット（Show All Commands）が開く

### とりあえず導入してみた拡張機能

[ATOM] エディタみたいに印を付けておいて纏めてインストールできるといいんだけどねぇ。

- [EditorConfig for VS Code - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig) : [EditorConfig](https://editorconfig.org/)
- [zenkaku - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=mosapride.zenkaku) : 全角空白文字を色付きで表示できる
- [Zenkaku-Hankaku - Visual Studio Marketplace](https://marketplace.visualstudio.com/items?itemName=masakit.zenkaku-hankaku) : 全角⇔半角文字変換

## ブックマーク

- [Running Visual Studio Code on Linux](https://code.visualstudio.com/docs/setup/linux)
- [UbuntuにVSCodeをインストールする3つの方法 - Qiita](https://qiita.com/yoshiyasu1111/items/e21a77ed68b52cb5f7c8)
- [VS Code でドキュメントの空白文字を見やすくしてみる - Qiita](https://qiita.com/satokaz/items/cb45d82f6f8f1e24c0d6)

[VS Code]: https://code.visualstudio.com/ "Visual Studio Code - Code Editing. Redefined"
[Ubuntu]: https://www.ubuntu.com/ "The leading operating system for PCs, IoT devices, servers and the cloud | Ubuntu"
[Snap]: https://github.com/snapcore/snapd "snapcore/snapd: The snapd and snap tools enable systems to work with .snap files."
[ATOM]: https://atom.io/