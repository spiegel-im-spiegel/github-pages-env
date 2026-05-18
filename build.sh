#!/bin/sh
set -eu

# TAGTOOLS_EXEC controls how tag metadata generators are executed.
# - wrapper (default): call compatibility wrappers (./toptags.sh, ./tagslist.sh)
# - go: call Go binary directly
TAGTOOLS_EXEC="${TAGTOOLS_EXEC:-wrapper}"

ROOT_DIR="$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)"
DEFAULT_BIN="$ROOT_DIR/tagtools/bin/tagtools"
TAGTOOLS_BIN="${TAGTOOLS_BIN:-$DEFAULT_BIN}"
TAGTOOLS_SOURCE="${TAGTOOLS_SOURCE:-external}"
TAGTOOLS_MODULE="${TAGTOOLS_MODULE:-github.com/spiegel-im-spiegel/tagtools}"
TAGTOOLS_VERSION="${TAGTOOLS_VERSION:-v0.1.0}"

if [ "$TAGTOOLS_EXEC" = "go" ]; then
	if [ ! -x "$TAGTOOLS_BIN" ]; then
		mkdir -p "$(dirname "$TAGTOOLS_BIN")"
		if [ "$TAGTOOLS_SOURCE" = "local" ]; then
			go build -o "$TAGTOOLS_BIN" "$ROOT_DIR/tagtools"
		else
			GOBIN="$(dirname "$TAGTOOLS_BIN")" go install "$TAGTOOLS_MODULE@$TAGTOOLS_VERSION"
		fi
	fi
	"$TAGTOOLS_BIN" toptags || exit 1
	"$TAGTOOLS_BIN" tagslist || exit 1
else
	./toptags.sh || exit 1
	./tagslist.sh || exit 1
fi

hugo --gc --cleanDestinationDir --destination=../text-publishd || exit 1
