#!/bin/sh
set -eu

# TAGTOOLS_EXEC controls how tag metadata generators are executed.
# - wrapper (default): call compatibility wrappers (./toptags.sh, ./tagslist.sh)
# - go: call Go binary directly
TAGTOOLS_EXEC="${TAGTOOLS_EXEC:-wrapper}"

if [ "$TAGTOOLS_EXEC" = "go" ]; then
	if [ ! -x ./tagtools/bin/tagtools ]; then
		mkdir -p ./tagtools/bin
		go build -o ./tagtools/bin/tagtools ./tagtools
	fi
	./tagtools/bin/tagtools toptags || exit 1
	./tagtools/bin/tagtools tagslist || exit 1
else
	./toptags.sh || exit 1
	./tagslist.sh || exit 1
fi

hugo --gc --cleanDestinationDir --destination=../text-publishd || exit 1
