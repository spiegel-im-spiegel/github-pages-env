+++
title = "Node.js v24 がリリースされた"
date =  "2025-05-09T07:34:18+09:00"
description = "Node.js v24 は 2025-10 以降に LTS になる予定である。"
image = "/images/attention/tools.png"
tags  = [ "nodejs", "programming", "react" ]
pageType = "text"

[scripts]
  mathjax = false
  mermaidjs = false
+++

最近仕事で React なんか弄ってるもんで [Node.js] の動向も押さえておこうかな，と。

- [Node.js — Node v24.0.0 (Current)](https://nodejs.org/en/blog/release/v24.0.0)

この記事について [Kagi Assistant] に要約してもらった。

{{< fig-quote type="markdown" >}}
Node.js 24.0.0のリリース内容を要約します。

- **リリース日**: 2025年5月6日
- **主な更新点**:
  - V8 JavaScriptエンジンがバージョン13.6にアップデート
    - 新機能例: Float16Array, 明示的リソース管理, RegExp.escape, WebAssembly Memory64, Error.isErrorなど
  - npmがバージョン11にアップデート
    - パフォーマンス向上やセキュリティ改善、最新JavaScriptパッケージとの互換性強化
  - WindowsでのビルドにMSVCサポートを廃止し、ClangCLが必須に
  - `AsyncLocalStorage`がデフォルトで`AsyncContextFrame`を使用し、非同期コンテキスト追跡の効率化
  - `URLPattern` APIがグローバルオブジェクトに追加され、URLパターンマッチングが容易に
  - パーミッションモデルの改善（実験的フラグが`--experimental-permission`から`--permission`に変更）
  - テストランナーの改善でサブテストの自動待機が可能に
  - HTTPクライアントライブラリUndiciがバージョン7に更新
- **非推奨・削除**:
  - `url.parse()`のランタイム非推奨（WHATWG URL APIの使用推奨）
  - `tls.createSecurePair`の削除
  - `SlowBuffer`のランタイム非推奨
  - REPLのnewなしインスタンス化の非推奨
  - Zlibクラスのnewなし使用の非推奨
  - `child_process`の`spawn`や`execFile`での`args`パラメータの非推奨
- **LTS移行予定**: 2025年10月にLTSへ移行予定（現在はCurrentリリース）

このリリースはNode.jsのプラットフォームの進化を促進する重要なアップデートであり、開発者は新機能の活用や非推奨APIの置き換えを検討することが推奨されています[^1]。

[^1]: [Node.js — Node v24.0.0 (Current)](https://nodejs.org/en/blog/release/v24.0.0#:~:text=Full%20document,more%20information.)
{{< /fig-quote >}}

というわけで [Node.js] v24 は 2025-10 以降に LTS になる予定である。
[Node.js] の LTS は通常30ヶ月間サポートされる。
現在 LTS として運用されているバージョンは v22 と v20 である（v20 は 2026-04 まで）。
また v18 は 2025-04 に EoL (End-of-Life) になっている。

## ブックマーク

- [React のお勉強]({{< ref "/remark/2025/03/a-study-in-react.md" >}})

[Node.js]: https://nodejs.org/ "Node.js — Run JavaScript Everywhere"
[Kagi Assistant]: {{< ref "/remark/2025/04/kagi-assistant-for-all-users.md" >}} "Kagi Assistant が全ユーザに解放"

## 参考文献

{{% review-paapi "4814401094" %}} <!-- Effective TypeScript 第2版 -->
{{% review-paapi "4297129167" %}} <!-- TypeScriptとReact/Next.jsでつくる実践Webアプリケーション開発 -->
