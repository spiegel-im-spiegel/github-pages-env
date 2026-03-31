#!/bin/sh
./tagslist.sh || exit 1
hugo --gc --cleanDestinationDir --destination=../text-publishd || exit 1
