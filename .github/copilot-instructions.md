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

**README Update**

- Usage of `new.sh` has already been added to `text/README.md`.

**Operational Notes / Next Steps**

- Deployment target and CI automation (for example, GitHub Actions) are not configured yet. If needed, add automation steps together with `publish.sh`.
- Further automation for `new.sh` is possible (for example, auto-generating a `slug` from an English title and applying it to the file name).

---

If you want to add or revise any part of this summary, let me know.
