#!/bin/bash
set -euo pipefail

TOP_N="${TOP_N:-15}"
OUT_FILE="data/toptags.json"

cutoff=$(date -d '1 year ago' +%F)
today=$(date +%F)

mkdir -p "$(dirname "$OUT_FILE")"

extract_tags_last_year() {
  local file="$1"
  awk -v cutoff="$cutoff" -v today="$today" '
    BEGIN { fm=0; in_tags=0; eligible=0 }

    $0=="+++" { fm++; next }

    fm==1 {
      if ($0 ~ /^date[[:space:]]*=/) {
        if (match($0, /"[0-9]{4}-[0-9]{2}-[0-9]{2}/)) {
          post_date=substr($0, RSTART+1, 10)
          eligible=(post_date >= cutoff && post_date <= today)
        }
      }

      if (in_tags) {
        if (eligible) {
          line=$0
          while (match(line, /"[^"]+"/)) {
            print substr(line, RSTART+1, RLENGTH-2)
            line=substr(line, RSTART+RLENGTH)
          }
        }
        if ($0 ~ /\]/) in_tags=0
        next
      }

      if ($0 ~ /^tags[[:space:]]*=/) {
        if (eligible) {
          line=$0
          while (match(line, /"[^"]+"/)) {
            print substr(line, RSTART+1, RLENGTH-2)
            line=substr(line, RSTART+RLENGTH)
          }
        }
        if ($0 !~ /\]/) in_tags=1
      }
    }
  ' "$file"
}

find content -type f -name '*.md' -print0 \
  | while IFS= read -r -d '' f; do
      extract_tags_last_year "$f"
    done \
  | sort \
  | uniq -c \
  | awk '{
      count=$1
      $1=""
      sub(/^ +/, "", $0)
      printf "%s\t%d\n", $0, count
    }' \
  | sort -k2,2nr -k1,1 \
  | head -n "$TOP_N" \
  | cut -f1 \
  | sort \
  | awk '
      BEGIN { printf "[" }
      {
        gsub(/\\/, "\\\\")
        gsub(/"/, "\\\"")
        if (NR > 1) printf ", "
        printf "\"%s\"", $0
      }
      END { print "]" }
    ' > "$OUT_FILE"

echo "Updated $OUT_FILE"
