#!/bin/bash
./build.sh || exit 1
pushd ../published
comment="Auto commit in $(date -u '+%FT%T+00:00')"
if [ $# -ne 0 ]; then
  comment=$1
fi
git add --all || exit 1
git commit -v -m "$comment" || exit 1
git push -u origin master || exit 1
popd
