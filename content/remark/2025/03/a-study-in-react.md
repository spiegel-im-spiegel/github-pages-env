+++
title = "React のお勉強"
date =  "2025-03-14T18:44:25+09:00"
description = "もう本当に基礎の基礎。 JSX から始めよう"
image = "/images/attention/kitten.jpg"
tags = [ "programming", "engineering", "react" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
  jsx = true
+++

[React] について勉強中。

いや，今までにも [React] / [TypeScript] なプロジェクトに参加したことはあるんだけど，ある程度お膳立てされた状態での join だったので，正直 [React] や [TypeScript] を知らなくても（やっつけの知識で）コードは書けるしテストも出来るわけよ。
知らない言語でもコードなんて見れば分かるし。
でも，それってプロジェクトが終われば忘れるぢゃん，身に付いてないんだから。

というわけで，ちょうど新しいプロジェクトのための事前学習を始めたタイミングということもあり，勤務先の有識者に最初から教えてもらえることになった。
ただし，その前に公式サイトの[学習ページ（日本語があるのか）](https://ja.react.dev/learn "クイックスタート – React")は読んでおいてね，と言われ，スキマ時間で勉強を始めたのだが... JSX から始めるのか。
よしやってみよう。

JSX で簡単なコードを書くだけなら特別な開発環境も要らないしフレームワークも要らないらしい。
準備として HTML の `<head>` 要素に以下の記述を加えればよい。

```html
<script src="https://unpkg.com/@babel/standalone/babel.min.js"></script>
<script async src="https://ga.jspm.io/npm:es-module-shims@1.7.0/dist/es-module-shims.js"></script>
<script type="importmap">
{
  "imports": {
    "react": "https://esm.sh/react?dev",
    "react-dom/client": "https://esm.sh/react-dom/client?dev"
  }
}
</script>
```

[@babel/standalone] は JSX から JavaScript (`React.createElement` の式) への変換を行うものらしい。
コードはこんな感じに記述する。

```html {hl_lines=[7,12]}
<div id="root1"></div>
<script type="text/babel" data-type="module">
  import React from 'react';
  import { createRoot } from 'react-dom/client';

  const MyApp = function() {
    return <strong>Hello, world!</strong>;
  }

  const container = document.getElementById('root1');
  const root = createRoot(container);
  root.render(<MyApp />);
</script>
```

強調している部分が JSX の特徴的な部分で HTML タグ風の記述（JSX 要素）が使えるため Web デザイナーでもとっつきやすいというのが売りなんだそうな。
ユーザが定義した `MyApp()` 関数の名前をタグのように使えるのがポイント。
タグは XML と同じ仕様なので void タグについては `<MyApp />` のように明示する必要がある。

このコードを含む HTML ファイルをブラウザで表示すると以下のような表示になる。

{{< fig-gen class="box left" title="Hello, World by React/JSX" >}}
<div id="root1"></div>
<script type="text/babel" data-type="module">
  import React from 'react';
  import { createRoot } from 'react-dom/client';

  const MyApp = function() {
    return <strong>Hello, world!</strong>;
  }

  const container = document.getElementById('root1');
  const root = createRoot(container);
  root.render(<MyApp />);
</script>
{{< /fig-gen >}}

うんうん。
ちゃんと動いてるな。

HTML タグ記法を JSX の記法に変換する “[HTML to JSX](https://transform.tools/html-to-jsx)” というツールもある。

[チュートリアル](https://ja.react.dev/learn)のコードをローカルで試したい場合。

まずは node.js のインストールを済ませておくこと。
私は [NodeSource](https://github.com/nodesource "NodeSource") から v22 LTS 版を入れている。
最近は [Bun](https://bun.sh/ "Bun — A fast all-in-one JavaScript runtime") の話もよく聞くようになった。

{{< fig-img-quote src="./tit-tat-toe-1.png" title="チュートリアル：三目並べ – React" link="https://ja.react.dev/learn/tutorial-tic-tac-toe" width="713" >}}

右上にある `Fork` を押下すると CodeSandbox のページが開く。

{{< fig-img-quote src="./codesandbox.png" title="Preview - nodebox - CodeSandbox" link="./codesandbox.png" width="713" lang="en" >}}

左上にある **□** を押下して表示されるメニューから “Download Sandbox” を選択してコードをまるごと（zip 形式）ダウンロードし，ローカルの適当なディレクトリに展開する。

`npm install` を実行して依存ライブラリをインストールしようとしたのだがエラーになって怒られたので `--legacy-peer-deps` オプションを付けてやり直す。
今度はうまく行ったが deprecated の嵐（笑）

いよいよ `npm start` でローカルサーバを起動しようとしたのだが，またしてもエラーで起動せず `orz` どうも `npm audit fix --force` で依存関係を更新しないといけないらしい。
しかもこれを実行するたびに状況が変わるってどうなってるの？？？

ホンマ，この辺のテキトーさが node.js 開発系が嫌な理由なんだよな。
しかも依存が深すぎて？ 何をどうしていいのか分からない。
みんなこれちゃんとメンテ出来てるの？

ぶちぶちと愚痴を垂れつつ `npm start` を再実行。
とりあえず動いたかな。

{{< fig-img-quote src="./npm-start.png" title="npm start" link="./npm-start.png" width="713" lang="en" >}}

{{< fig-img-quote src="./tit-tat-toe-2.png" title="三目並べ 実行結果" link="./tit-tat-toe-2.png" width="713" >}}

動いてるっぽいな。

さて，お勉強を続けますか。

## ブックマーク

- [@babel/standalone · Babel](https://babeljs.io/docs/babel-standalone)
  - [babel/standaloneの使い方(文字列に格納したjavascriptソースをブラウザ内でトランスパイルする) #React - Qiita](https://qiita.com/murasuke/items/8dbe350c0c1c1fe269bf)
- [es-module-shims - npm](https://www.npmjs.com/package/es-module-shims)

- [新・日本一わかりやすいReact入門【基礎編】 - YouTube](http://www.youtube.com/playlist?list=PLX8Rsrpnn3IWPoM7-1YPDksRRkamRY25k)
- [Hugo で React + TypeScript を利用してサクッとウェブサイトに RSS リーダーを追加する](https://zenn.dev/nikaera/articles/hugo-react-dev)
- [BunでReact関連の開発環境を構築する](https://zenn.dev/yamakenji24/articles/bcba86de05e5d5)
- [ReactをBunで使用する方法 - Web開発における知見共有系ページ](https://job-info.hateblo.jp/entry/2024/07/30/145804)
- [package.jsonで時々見かけるbrowserslistとは](https://zenn.dev/taketaku/articles/ffb239c3da8613)

[React]: https://react.dev/ "React"
[TypeScript]: https://www.typescriptlang.org/ "TypeScript: JavaScript With Syntax For Types."
[@babel/standalone]: https://babeljs.io/docs/babel-standalone "@babel/standalone · Babel"

## 参考文献

{{% review-paapi "4297129167" %}} <!-- TypeScriptとReact/Next.jsでつくる実践Webアプリケーション開発 -->
