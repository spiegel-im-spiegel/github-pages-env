#!/usr/bin/env bash
set -euo pipefail

branch="master"
cutoff="7ece30869822d78b9bdf8a2a4611e5de8998d5b3"
message="Consolidate history through July 2026"

repo_root=$(git rev-parse --show-toplevel)
cd "$repo_root"

fail() {
  printf 'Error: %s\n' "$1" >&2
  exit 1
}

[[ "$(git branch --show-current)" == "$branch" ]] || fail "Run this script on the master branch."
[[ -z "$(git status --porcelain)" ]] || fail "The working tree must be clean."
git cat-file -e "$cutoff^{commit}" || fail "The cutoff commit is not available."

old_head=$(git rev-parse "$branch")
git merge-base --is-ancestor "$cutoff" "$old_head" || fail "The cutoff is not an ancestor of master."
[[ -z "$(git rev-list --merges "$cutoff..$old_head")" ]] || fail "Merge commits exist after the cutoff."

old_tree=$(git rev-parse "$old_head^{tree}")
cutoff_tree=$(git rev-parse "$cutoff^{tree}")
old_count=$(git rev-list --count "$cutoff..$old_head")
cutoff_date=$(git show -s --format=%aI "$cutoff")
backup_branch="backup/master-before-202608-$(date '+%Y%m%d%H%M%S')"

git show-ref --verify --quiet "refs/heads/$backup_branch" && fail "Backup branch already exists: $backup_branch"
git branch "$backup_branch" "$old_head"

aggregate=$(printf '%s\n' "$message" | GIT_AUTHOR_DATE="$cutoff_date" GIT_COMMITTER_DATE="$cutoff_date" git commit-tree -S "$cutoff_tree")
printf 'Backup branch: %s\nAggregate commit: %s\nReplaying %s commits after %s\n' \
  "$backup_branch" "$aggregate" "$old_count" "$cutoff"

git rebase --onto "$aggregate" "$cutoff" "$branch"

new_head=$(git rev-parse "$branch")
new_tree=$(git rev-parse "$new_head^{tree}")
new_count=$(git rev-list --count "$aggregate..$new_head")

[[ "$new_tree" == "$old_tree" ]] || fail "The final tree differs from the original. Restore from $backup_branch if needed."
[[ "$new_count" == "$old_count" ]] || fail "The retained commit count changed. Restore from $backup_branch if needed."
if git merge-base --is-ancestor "$cutoff" "$new_head"; then
  fail "The old cutoff is still an ancestor. Restore from $backup_branch if needed."
fi
[[ -z "$(git status --porcelain)" ]] || fail "The working tree is not clean after rewriting."

printf 'History rewrite verified.\nNew master: %s\nTree: %s\nRetained commits: %s\n' \
  "$new_head" "$new_tree" "$new_count"
