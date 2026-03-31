#!/bin/bash
set -euo pipefail

out_file=".github/workflows/tagslist.csv"

mkdir -p "$(dirname "$out_file")"

extract_tags_from_frontmatter() {
  local file="$1"
  awk '
    BEGIN { fm=0; in_tags=0 }
    $0=="+++" { fm++; next }
    fm==1 {
      if (in_tags) {
        line=$0
        while (match(line, /"[^"]+"/)) {
          print substr(line, RSTART+1, RLENGTH-2)
          line=substr(line, RSTART+RLENGTH)
        }
        if ($0 ~ /\]/) in_tags=0
        next
      }

      if ($0 ~ /^tags[[:space:]]*=/) {
        line=$0
        while (match(line, /"[^"]+"/)) {
          print substr(line, RSTART+1, RLENGTH-2)
          line=substr(line, RSTART+RLENGTH)
        }
        if ($0 !~ /\]/) in_tags=1
      }
    }
  ' "$file"
}

{
  echo "tag,count"
  find content -type f -name '*.md' -print0 \
    | while IFS= read -r -d '' f; do
        extract_tags_from_frontmatter "$f"
      done \
    | sort \
    | uniq -c \
    | awk '{
        count=$1
        $1=""
        sub(/^ +/, "", $0)
        printf "%s\t%s\n", count, $0
      }' \
    | sort -k1,1nr -k2,2 \
    | awk -F '\t' '{
        count=$1
        tag=$2
        gsub(/"/, "\"\"", tag)
        printf "\"%s\",%d\n", tag, count
      }'
} > "$out_file"

echo "Updated $out_file"
