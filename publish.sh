#!/bin/bash
./build.sh || exit 1
pushd ../published
git add --all || exit 1
git commit -v -m "Auto commit in $(date -u '+%FT%T+00:00')" || exit 1
git push -u origin master || exit 1
popd
