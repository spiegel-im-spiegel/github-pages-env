#!/bin/sh
set -eu

ROOT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
DEFAULT_BIN="$ROOT_DIR/tagtools/bin/tagtools"
TAGTOOLS_BIN="${TAGTOOLS_BIN:-$DEFAULT_BIN}"
TAGTOOLS_SOURCE="${TAGTOOLS_SOURCE:-external}"
TAGTOOLS_MODULE="${TAGTOOLS_MODULE:-github.com/spiegel-im-spiegel/tagtools}"
TAGTOOLS_VERSION="${TAGTOOLS_VERSION:-v0.3.0}"

if [ ! -x "$TAGTOOLS_BIN" ]; then
  mkdir -p "$(dirname "$TAGTOOLS_BIN")"
  if [ "$TAGTOOLS_SOURCE" = "local" ]; then
    go build -o "$TAGTOOLS_BIN" "$ROOT_DIR/tagtools"
  else
    GOBIN="$(dirname "$TAGTOOLS_BIN")" go install "$TAGTOOLS_MODULE@$TAGTOOLS_VERSION"
  fi
fi

if [ -n "${TOP_N:-}" ]; then
  set -- --top-n "$TOP_N" "$@"
fi

if [ -n "${TOPTAGS_WINDOW:-}" ]; then
  set -- --window "$TOPTAGS_WINDOW" "$@"
fi

exec "$TAGTOOLS_BIN" toptags "$@"
