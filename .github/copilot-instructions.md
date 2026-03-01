**Hugo Skills — 操作ガイド**

- **Purpose**: このリポジトリで Copilot が支援する Hugo 関連のスキルと運用上の振る舞いをまとめます。

**Local Commands**

- **`new.sh`**: 投稿作成を補助するスクリプト。
  - `archetypes/*.md` を自動検出して、`hugo new --kind=<kind> <path>` を実行します。
  - 安全なシェルオプション `set -euo pipefail` を採用しています。
  - `remark` と `bookmarks` のデフォルト生成は日付の先頭 `DD-` をファイル名に付けます（例: `01-stories.md`, `01-bookmarks.md`）。
  - 引数の取り扱い: 第一引数がセクション名（例: `remark`）かファイル名かを判定します。存在しない場合は利用可能な archetype を表示します。

**Archetypes / Front Matter**

- `archetypes/*.md` に合わせて front matter を生成します（例: `archetypes/remark.md` のフィールドを踏襲）。
 - 投稿ファイル生成時に `slug` フィールドを自動で追加しません。必要な場合は手動で `slug` を追加してください。
 - 英語タイトル用スラグはハイフン区切りの小文字を想定します（例: `reading-is-dead`）。

**投稿作成の例**

```bash
./new.sh remark                # content/remark/<YYYY>/<MM>/stories.md を作成（archetype: remark があれば使用）
./new.sh bookmarks             # content/bookmarks/<YYYY>/<MM>/bookmarks.md を作成
./new.sh release my-post.md    # content/release/<YYYY>/<MM>/my-post.md を作成
./new.sh some/path.md          # 任意パスで作成（archetype が無ければ default を使用）
```

**README 更新**

- `text/README.md` に `new.sh` の使い方を追記済みです。

**運用上の注意 / 次のステップ**

- デプロイ先と CI（GitHub Actions 等）の自動化は未設定です。必要であれば `publish.sh` と合わせて自動化手順を追加します。
- `new.sh` のさらなる自動化（例: 英語タイトルから自動で `slug` を生成してファイル名に反映）は対応可能です。

---

もしこの要約に追記・修正したい項目があれば教えてください。
