#!/bin/sh
./toptags.sh || exit 1
./tagslist.sh || exit 1
hugo --gc --cleanDestinationDir --destination=../text-publishd || exit 1
