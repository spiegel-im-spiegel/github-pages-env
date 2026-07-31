# tagtools

A small CLI toolset for Hugo tag metadata workflows in this repository.

## Features

- Generate `.github/workflows/tagslist.csv` from front matter tags.
- Generate `data/toptags.json` from posts in the last year.
- Verify generated outputs against golden files.
- Emit machine-readable verify output with `--debug`.

## Requirements

- Go 1.26+

## Build

```bash
go build -o ./bin/tagtools .
```

## Commands

### tagslist

Generate tag counts and preserve `means` values from an existing CSV.

```bash
./bin/tagtools tagslist \
  --content-dir ../content \
  --out ../.github/workflows/tagslist.csv \
  --inherit-means-from ../.github/workflows/tagslist.csv
```

### toptags

Generate sorted top tags from posts in the last year.

```bash
./bin/tagtools toptags \
  --content-dir ../content \
  --out ../data/toptags.json \
  --top-n 15 \
  --today 2026-05-18
```

### all

Run `toptags` and then `tagslist` with default settings.

```bash
./bin/tagtools all
```

### verify

Compare generated outputs against expected files.

```bash
./bin/tagtools verify \
  --content-dir ../content \
  --expected-tagslist testdata/golden/tagslist.csv \
  --expected-toptags testdata/golden/toptags.json \
  --inherit-means-from ../.github/workflows/tagslist.csv \
  --today 2026-05-18
```

When verification fails, the output includes:

- stage (`tagslist` or `toptags`)
- expected and actual file paths
- first different line
- expected/actual line excerpts
- a `reproduce:` command

## Debug JSON output

Add `--debug` to `verify` for JSON output.

```bash
./bin/tagtools verify --debug \
  --content-dir ../content \
  --expected-tagslist testdata/golden/tagslist.csv \
  --expected-toptags testdata/golden/toptags.json \
  --inherit-means-from ../.github/workflows/tagslist.csv \
  --today 2026-05-18
```

Success example:

```json
{"status":"ok","reproduce_command":"./bin/tagtools verify ... --debug"}
```

Failure example:

```json
{
  "status": "error",
  "stage": "toptags",
  "message": "toptags verification failed (first diff line: 1)",
  "expected_path": "testdata/golden/toptags.json",
  "actual_path": "/tmp/tagtools-verify-xxxx/toptags.json",
  "line": 1,
  "expected_line": "...",
  "actual_line": "...",
  "reproduce_command": "./bin/tagtools verify ... --debug"
}
```

## Test

```bash
go test ./...
```
