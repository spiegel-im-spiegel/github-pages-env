# Tag Absorption Plan (2026-05)

## Purpose
- Count=1 tags and typo-like tags are consolidated into broader tags.
- This document is the canonical plan before bulk edits.

## Decision Notes
- `sha-3` is replaced with `hash` (not `cryptography`).
- Astronomy mission tags are handled with source-organization enrichment:
  - `akatsuki`, `kaguya` -> add `astronomy` + `jaxa`
  - `juno`, `voyager` -> add `astronomy` + `nasa`
- When adding tags, do not duplicate existing tags in the same front matter list.

## Replacement Rules
### Cryptography / Security
- `aead` -> add `cryptography`, remove `aead`
- `misty1` -> add `cryptography`, remove `misty1`
- `rc4` -> add `cryptography`, remove `rc4`
- `sha-3` -> add `hash`, remove `sha-3`

### Astronomy
- `akatsuki` -> add `astronomy`,`jaxa`, remove `akatsuki`
- `kaguya` -> add `astronomy`,`jaxa`, remove `kaguya`
- `juno` -> add `astronomy`,`nasa`, remove `juno`
- `voyager` -> add `astronomy`,`nasa`, remove `voyager`
- `tmt` -> add `astronomy`, remove `tmt`
- `planet9` -> add `astronomy`, remove `planet9`
- `interstellar-object` -> add `astronomy`, remove `interstellar-object`
- `satellite-constellation` -> add `astronomy`, remove `satellite-constellation`
- `comet` -> add `astronomy`, remove `comet`
- `mercury` -> add `astronomy`, remove `mercury`
- `venus` -> add `astronomy`, remove `venus`
- `jupiter` -> add `astronomy`, remove `jupiter`

### Math
- `galois-theory` -> add `math`, remove `galois-theory`
- `gaussian` -> add `math`, remove `gaussian`
- `integral` -> add `math`, remove `integral`
- `matrix` -> add `math`, remove `matrix`
- `vector` -> add `math`, remove `vector`
- `permutation` -> add `math`, remove `permutation`
- `greatest-common-divisor` -> add `math`, remove `greatest-common-divisor`

### Linux / Infra
- `cifs` -> add `linux`, remove `cifs`
- `nfs` -> add `linux`, remove `nfs`
- `file-system` -> add `linux`, remove `file-system`
- `pacman` -> add `linux`, remove `pacman`
- `glibc` -> add `linux`, remove `glibc`
- `gnu` -> add `linux`, remove `gnu`

### Programming
- `closure` -> add `programming`, remove `closure`
- `functional-options` -> add `programming`, remove `functional-options`
- `reference` -> add `programming`, remove `reference`
- `relation` -> add `programming`, remove `relation`
- `stack` -> add `programming`, remove `stack`
- `singleton` -> add `programming`, remove `singleton`

### Security incidents
- `blueborne` -> add `security`, remove `blueborne`
- `botnet` -> add `security`, remove `botnet`
- `class-break` -> add `security`, remove `class-break`
- `ddos` -> add `security`, remove `ddos`
- `ransomware` -> add `security`, remove `ransomware`

### Web / Service and policy overrides
- `gitlab` -> remove only (no replacement)
- `trello` -> add `tools`, remove `trello`
- `tumblr` -> add `media`, remove `tumblr`
- `open-graph` -> add `web`, remove `open-graph`
- `punycode` -> add `web`, remove `punycode`
- `codespaces` -> add `github`, remove `codespaces`

### Typo / cleanup
- `requirement-revelopment` -> add `design`, remove `requirement-revelopment`
- `study` -> remove only (no replacement)
- `review` -> remove only (no replacement)

## Execution Policy
- Run updates file-by-file from the batch manifest.
- Keep tags unique after add/remove operations.
- Preserve order as much as practical; append newly added tags at the end when missing.

## Batch Script Used

```bash
cd /home/spiegel/ws/hugo/text

update_tags() {
  local file="$1" remove_list="$2" add_list="$3"
  local line inner
  line=$(rg -n '^tags\s*=' "$file" | head -n1 | cut -d: -f2-)
  if [[ -z "$line" ]]; then
    echo "[skip:no-tags] $file"
    return
  fi
  inner=$(echo "$line" | sed -E 's/^tags\s*=\s*\[(.*)\]\s*$/\1/')

  mapfile -t tags < <(echo "$inner" | tr ',' '\n' | sed -E 's/^[[:space:]]*"?([^" ]+)"?[[:space:]]*$/\1/' | sed '/^$/d')

  if [[ -n "$remove_list" ]]; then
    IFS='|' read -r -a removes <<< "$remove_list"
    local filtered=()
    for t in "${tags[@]}"; do
      local drop=0
      for r in "${removes[@]}"; do
        [[ "$t" == "$r" ]] && drop=1 && break
      done
      [[ $drop -eq 0 ]] && filtered+=("$t")
    done
    tags=("${filtered[@]}")
  fi

  if [[ -n "$add_list" ]]; then
    IFS='|' read -r -a adds <<< "$add_list"
    for a in "${adds[@]}"; do
      [[ -z "$a" ]] && continue
      local exists=0
      for t in "${tags[@]}"; do
        [[ "$t" == "$a" ]] && exists=1 && break
      done
      [[ $exists -eq 0 ]] && tags+=("$a")
    done
  fi

  local rebuilt='tags = [ '
  local first=1
  for t in "${tags[@]}"; do
    if [[ $first -eq 1 ]]; then
      rebuilt+="\"$t\""
      first=0
    else
      rebuilt+=", \"$t\""
    fi
  done
  rebuilt+=' ]'

  sed -i -E "s|^tags\s*=\s*\[.*\]$|$rebuilt|" "$file"
  echo "[updated] $file"
}

while IFS=',' read -r file remove_tags add_tags notes; do
  [[ "$file" == "file" ]] && continue
  file=${file#\"}; file=${file%\"}
  remove_tags=${remove_tags#\"}; remove_tags=${remove_tags%\"}
  add_tags=${add_tags#\"}; add_tags=${add_tags%\"}
  update_tags "$file" "$remove_tags" "$add_tags"
done < docs/tag-absorption-batch-2026-05.csv
```
