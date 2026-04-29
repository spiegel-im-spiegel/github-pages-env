**Hugo Skills - Operation Guide**

- **Purpose**: This document defines Hugo-related operational rules for Copilot in this repository.

**Writing Style**

- In Japanese sentences, use `，` for commas and `。` for periods.

**Local Commands**

- **`new.sh`**: Helper script for creating posts.
  - Auto-detects `archetypes/*.md` and runs `hugo new --kind=<kind> <path>`.
  - Uses safe shell options `set -euo pipefail`.
  - For `remark` and `bookmarks`, default file names are prefixed with `DD-`.
  - First argument is treated as either a section name or a file name.
- **`ml`**: Converts a URL to Markdown link format.
  - Usage: `ml "https://example.com/article"`
  - Output: `[Page Title](https://example.com/article)`
- **`rg`** (`ripgrep`): Preferred text search command for article proofreading.
  - Usage examples:
    - `rg "フィッシング" content/remark`
    - `rg "Signal Support|Registration Lock" content/remark/2026/04/signal-attack-patterns.md`
  - Install on Ubuntu/Debian: `sudo apt install -y ripgrep`
  - Alternative install: `sudo snap install ripgrep`
  - Verify: `rg --version`
- **`build.sh`**: Build entry point.
  - Runs `./toptags.sh` and `./tagslist.sh`, then runs Hugo with destination `../text-publishd`.
- **`publish.sh`**: Build and publish helper.
  - Runs `./build.sh` and then commits/pushes in `../text-publishd`.
  - Usage:
    - `./publish.sh`
    - `./publish.sh "your commit message"`
  - If no argument is provided, uses the latest commit subject from `text`.
- **`tagslist.sh`**: Rebuilds `.github/workflows/tagslist.csv` from front matter tags.
- **`toptags.sh`**: Rebuilds `data/toptags.json` from recent posts.
- **`hugo_inst.sh`**: Updates Hugo using latest GitHub Release `.deb`.

**Deploy Playbook**

- Standard order for article publication tasks:
  1. Check changed files in `text`, and confirm whether non-article changes (for example `themes/baldanders-info`) should be included.
  2. Run a light pre-publish check (typos, links, and front matter: `title`, `description`, `tags`, `draft`, `date`).
  3. Propose 3-5 commit message candidates and proceed after one is selected.
  4. Commit source changes in `text`.
  5. Push source changes when requested.
  6. Run `./publish.sh` to build and deploy to `../text-publishd`.
  7. Check both repositories after deploy and report remaining local changes.
- Keep status reports concise and focused on actionable deltas.
- Update front matter `description` only when its current value is the placeholder string `description`.
- After deploy, if `.github/workflows/tagslist.csv` and/or `data/toptags.json` remain uncommitted in `text`, handle them in a separate commit.
- For that separate commit, choose the commit message automatically and do not push.

- Hugo upgrade and redeploy routine:
  1. Check both repositories first (`git status --short` in `text` and `../text-publishd`).
  2. Check current Hugo version (`hugo version`).
  3. Run `./hugo_inst.sh` to install the latest release.
  4. If sudo password is requested, continue interactively and complete the installer.
  5. Verify installed version again (`hugo version`).
  6. Run `./publish.sh` to rebuild and deploy with the new Hugo.
  7. Re-check both repositories after deploy and report remaining local changes.
  8. If `text` has only generated tag artifacts after deploy, create a separate local commit and do not push.
  9. If asked whether publish changes are version-only, review the latest commit with `git show --numstat` and inspect outlier files before concluding.

**Archetypes and Front Matter**

- Front matter is generated according to `archetypes/*.md` (for example, `archetypes/remark.md`).
- `slug` is not automatically added. Add it manually only when needed.
- For English titles, use lowercase hyphen-separated slugs (for example, `reading-is-dead`).
- Post creation examples:

```bash
./new.sh remark                # creates content/remark/<YYYY>/<MM>/stories.md
./new.sh bookmarks             # creates content/bookmarks/<YYYY>/<MM>/bookmarks.md
./new.sh release my-post.md    # creates content/release/<YYYY>/<MM>/my-post.md
./new.sh some/path.md          # creates at arbitrary path
```

**AI Block Rules**

- Use `div-ai` for AI-generated or AI-summarized text blocks.
- Basic wrapper:

```markdown
{{< div-ai type="markdown" >}}
...AI-generated content...
{{< /div-ai >}}
```

- Keep `div-box` for non-AI notes (manual updates, side remarks, generic callouts).
- When updating older posts, replace `div-box` with `div-ai` only for sections explicitly described as AI output.
- To list likely migration candidates, run `./div-ai-check.sh` (or `./div-ai-check.sh content/remark`).
