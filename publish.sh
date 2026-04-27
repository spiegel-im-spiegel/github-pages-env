#!/bin/bash
./build.sh || exit 1
pushd ../text-publishd
if [ $# -ne 0 ]; then
  comment=$1
else
  comment="$(git log -1 --pretty=%s)"
fi
git add --all || exit 1
git commit -v -m "$comment" || exit 1
git push -u origin master || exit 1
popd
