#!/bin/bash
set -euo pipefail

y=$(date '+%Y')
m=$(date '+%m')
d=$(date '+%d')

section="${1:-}"
fname="${2:-}"
p=""

available_kinds=""
for f in archetypes/*.md; do
  [ -e "$f" ] || continue
  available_kinds+="$(basename "$f" .md) "
done

usage() {
  echo "Usage: $0 <section|filename> [filename]"
  echo "Available archetypes: ${available_kinds:-(none)}"
  exit 1
}

if [ -z "$section" ]; then
  usage
elif [ "$section" = "remark" ]; then
  if [ -z "$fname" ]; then
    p="$section/$y/$m/$d-stories.md"
  else
    p="$section/$y/$m/$fname"
  fi
elif [ "$section" = "bookmarks" ]; then
  if [ -z "$fname" ]; then
    p="$section/$y/$m/$d-bookmarks.md"
  else
    p="$section/$y/$m/$fname"
  fi
elif [ "$section" = "release" ]; then
  if [ -z "$fname" ]; then
    echo "入力ファイルを指定してください。"
    exit 1
  else
    p="$section/$y/$m/$fname"
  fi
elif [ "$section" = "cc-licenses" ]; then
  if [ -z "$fname" ]; then
    echo "入力ファイルを指定してください。"
    exit 1
  else
    p="$section/$fname"
  fi
elif [ "$section" = "golang" ] || [ "$section" = "rust-lang" ] || [ "$section" = "hugo" ] || [ "$section" = "openpgp" ] || [ "$section" = "typst" ]; then
  if [ -z "$fname" ]; then
    echo "入力ファイルを指定してください。"
    exit 1
  else
    p="$section/$fname"
  fi
else
  fname="$section"
  section=""
  p="$fname"
fi

kind="default"
if echo " $available_kinds " | grep -qw " $section "; then
  kind="$section"
fi

echo "Create file content/$p (archetype: $kind)"
hugo new --kind="$kind" "$p"
