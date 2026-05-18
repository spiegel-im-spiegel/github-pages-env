#!/bin/sh
set -eu

ROOT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
DEFAULT_BIN="$ROOT_DIR/tagtools/bin/tagtools"
TAGTOOLS_BIN="${TAGTOOLS_BIN:-$DEFAULT_BIN}"

if [ ! -x "$TAGTOOLS_BIN" ]; then
  mkdir -p "$ROOT_DIR/tagtools/bin"
  go build -o "$DEFAULT_BIN" "$ROOT_DIR/tagtools"
  TAGTOOLS_BIN="$DEFAULT_BIN"
fi

exec "$TAGTOOLS_BIN" tagslist "$@"
