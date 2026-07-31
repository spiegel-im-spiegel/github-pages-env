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

out_tmp="$(mktemp)"
counts_tmp="$(mktemp)"
trap 'rm -f "$counts_tmp" "$out_tmp"' EXIT

{

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
        # tag<TAB>count
        printf "%s\t%d\n", $0, count
      }' \
    | sort -t $'\t' -k2,2nr -k1,1 > "$counts_tmp"

  echo "tag,count,means"
  awk -F '\t' -v primary_file="$out_file" '
    function csv_parse(line, out,    i, ch, next_ch, in_q, field, n) {
      n=1
      field=""
      in_q=0
      for (i=1; i<=length(line); i++) {
        ch = substr(line, i, 1)
        if (in_q) {
          if (ch == "\"") {
            next_ch = substr(line, i+1, 1)
            if (next_ch == "\"") {
              field = field "\""
              i++
            } else {
              in_q = 0
            }
          } else {
            field = field ch
          }
        } else {
          if (ch == "\"") {
            in_q = 1
          } else if (ch == ",") {
            out[n++] = field
            field = ""
          } else {
            field = field ch
          }
        }
      }
      out[n] = field
      return n
    }
    function load_means(file,    line, n, tag, means) {
      if ((getline line < file) > 0) {
        while ((getline line < file) > 0) {
          n = csv_parse(line, cols)
          tag = cols[1]
          means = (n >= 3 ? cols[3] : "")
          means_map[tag] = means
        }
        close(file)
      }
    }
    BEGIN {
      load_means(primary_file)
    }
    {
      tag = $1
      count = $2
      means = ((tag in means_map) ? means_map[tag] : "")
      gsub(/"/, "\"\"", tag)
      gsub(/"/, "\"\"", means)
      printf "\"%s\",%d,\"%s\"\n", tag, count, means
    }
  ' "$counts_tmp"
} > "$out_tmp"

mv "$out_tmp" "$out_file"

echo "Updated $out_file"
