**Hugo Skills - Operation Guide**

- **Purpose**: This document summarizes Hugo-related skills and operational behavior supported by Copilot in this repository.

**Writing Style**

- In Japanese sentences, use `，` for commas and `。` for periods.

**Local Commands**

- **`new.sh`**: A helper script for creating posts.
  - It auto-detects `archetypes/*.md` and runs `hugo new --kind=<kind> <path>`.
  - It uses the safe shell options `set -euo pipefail`.
  - Default generation for `remark` and `bookmarks` prefixes file names with `DD-` (for example, `01-stories.md`, `01-bookmarks.md`).
  - Argument handling: The first argument is treated as either a section name (for example, `remark`) or a file name. If no valid section is found, available archetypes are shown.
- **`ml`**: A helper command that converts a URL into Markdown link format.
  - Basic usage: `ml "https://example.com/article"`
  - Output format: `[Page Title](https://example.com/article)`
  - Clipboard example (Linux): `out=$(ml "https://example.com/article") && printf "%s" "$out" | xsel --clipboard --input`
- **`publish.sh`**: A helper script for building and publishing generated site content.
  - It runs `./build.sh` first and stops immediately on failure.
  - It moves to `../text-publishd`, then runs `git add --all`, `git commit`, and `git push -u origin master`.
  - Basic usage: `./publish.sh`
  - With commit message: `./publish.sh "your commit message"`
  - If no argument is provided, it uses an auto-generated UTC commit message like `Auto commit in 2026-03-31T03:00:00+00:00`.
- **`tagslist.sh`**: A helper script for tag frequency export.
  - It scans front matter tags from `content/**/*.md`.
  - It writes CSV output to `.github/workflows/tagslist.csv`.
  - The CSV schema is `tag,count,means`.
  - The `means` column is preserved from the existing `tagslist.csv` when regenerating counts.
  - Output is sorted by descending count, with alphabetical tag order as a tie-breaker.
  - The generated `tagslist.csv` is used as a reference when deciding which tags to assign to future posts.
- **`toptags.sh`**: A helper script for recent top-tags export.
  - It targets posts within the last one year based on front matter `date`.
  - It writes a JSON array of tag names to `data/toptags.json`.
  - Default output size is top 15 tags; override with `TOP_N` environment variable.
  - The generated `toptags.json` is used as source data for the “最近の主な tags” section on `layouts/_default/home.html`.
- **`build.sh`**: Build entry point.
  - It runs `./toptags.sh` first and stops on error.
  - It then runs `./tagslist.sh` and stops on error.
  - It finally runs Hugo with `hugo --gc --cleanDestinationDir --destination=../text-publishd`.
- **`hugo_inst.sh`**: Hugo updater script using the latest GitHub Release `.deb`.
  - It fetches the latest release metadata from `gohugoio/hugo`.
  - It selects `hugo_extended_*_linux-<arch>.deb` for `amd64` or `arm64`.
  - It optionally verifies the downloaded file with `checksums.txt` when available.
  - It installs the package with `sudo apt install -y ./<deb-file>`.
  - Basic usage: `./hugo_inst.sh`

**Current Workflow Note (Update Hugo with hugo_inst.sh)**

- To update Hugo from the latest GitHub Release, run:

```bash
./hugo_inst.sh
```

- Example post-update verification workflow:

```bash
hugo version
./build.sh
cd ../text-publishd && git status --short
```

- If generated files are valid, continue with commit/push in `../text-publishd`.

**Current Workflow Note (Publish with latest commit message)**

- To reuse the latest commit subject from `text` as the publish commit message, run:

```bash
./publish.sh "$(git log -1 --pretty=%s)"
```

- This executes the site build, commits generated output in `../text-publishd`, and pushes to `origin/master`.

**Archetypes / Front Matter**

- Front matter is generated according to `archetypes/*.md` (for example, fields from `archetypes/remark.md`).
- The `slug` field is not automatically added when creating a post file. Add `slug` manually if needed.
- For English titles, use lowercase hyphen-separated slugs (for example, `reading-is-dead`).

**Post Creation Examples**

```bash
./new.sh remark                # creates content/remark/<YYYY>/<MM>/stories.md (uses archetype: remark when available)
./new.sh bookmarks             # creates content/bookmarks/<YYYY>/<MM>/bookmarks.md
./new.sh release my-post.md    # creates content/release/<YYYY>/<MM>/my-post.md
./new.sh some/path.md          # creates at an arbitrary path (uses default if no matching archetype)
```

**Current Workflow Note (Create remark with slug)**

- For the title "aptitude wo insutoru suru", the selected English slug is `install-aptitude`.
- To use the slug as the file name in the `remark` section, run:

```bash
./new.sh remark install-aptitude.md
```

- Output path: `content/remark/<YYYY>/<MM>/install-aptitude.md`
- For this case, do not add or modify front matter (use the generated content as-is).

**Current Workflow Note (`div-ai` shortcode usage)**

- Use `div-ai` for AI-generated or AI-summarized text blocks in article bodies.
- Basic markdown wrapper:

```markdown
{{< div-ai type="markdown" >}}
...AI-generated content...
{{< /div-ai >}}
```

- Keep `div-box` for non-AI notes such as manual updates, side remarks, and generic callouts.
- When updating older posts, replace `div-box` with `div-ai` only for sections that are explicitly described as AI output (for example, generated summaries by assistants).
- To list likely migration candidates, run `./div-ai-check.sh` (or pass a path like `./div-ai-check.sh content/remark`).

**README Update**

- Usage of `new.sh` has already been added to `text/README.md`.

**Operational Notes / Next Steps**

- Deployment target and CI automation (for example, GitHub Actions) are not configured yet. If needed, add automation steps together with `publish.sh`.
- Further automation for `new.sh` is possible (for example, auto-generating a `slug` from an English title and applying it to the file name).

---

If you want to add or revise any part of this summary, let me know.
