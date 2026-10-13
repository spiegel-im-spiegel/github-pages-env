#!/usr/bin/env bash
set -euo pipefail

branch="master"
cutoff="fd1d324d0423d8a27efea7daa808f28928b50b54"
message="Consolidate deployment history through July 2026"
script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
target_path="${1:-$script_dir/../../text-publishd}"

fail() {
  printf 'Error: %s\n' "$1" >&2
  exit 1
}

[[ $# -le 1 ]] || fail "Usage: $0 [publish-repository-path]"
[[ -d "$target_path" ]] || fail "Repository directory does not exist: $target_path"
repo=$(git -C "$target_path" rev-parse --show-toplevel) || fail "Not a Git repository: $target_path"

[[ "$(git -C "$repo" branch --show-current)" == "$branch" ]] || fail "The publish repository must be on the master branch."
[[ -z "$(git -C "$repo" status --porcelain)" ]] || fail "The publish repository working tree must be clean."
git -C "$repo" cat-file -e "$cutoff^{commit}" || fail "The cutoff commit is not available."

git -C "$repo" fetch origin master:refs/remotes/origin/master
old_head=$(git -C "$repo" rev-parse "$branch")
origin_head=$(git -C "$repo" rev-parse refs/remotes/origin/master)
[[ "$old_head" == "$origin_head" ]] || fail "Local master differs from origin/master; reconcile it before rewriting."
git -C "$repo" merge-base --is-ancestor "$cutoff" "$old_head" || fail "The cutoff is not an ancestor of master."
[[ -z "$(git -C "$repo" rev-list --merges "$cutoff..$old_head")" ]] || fail "Merge commits exist after the cutoff."

old_tree=$(git -C "$repo" rev-parse "$old_head^{tree}")
cutoff_tree=$(git -C "$repo" rev-parse "$cutoff^{tree}")
old_count=$(git -C "$repo" rev-list --count "$cutoff..$old_head")
cutoff_date=$(git -C "$repo" show -s --format=%aI "$cutoff")
backup_branch="backup/master-before-202608-$(date '+%Y%m%d%H%M%S')"

if git -C "$repo" show-ref --verify --quiet "refs/heads/$backup_branch"; then
  fail "Backup branch already exists: $backup_branch"
fi
git -C "$repo" branch "$backup_branch" "$old_head"

aggregate=$(printf '%s\n' "$message" | GIT_AUTHOR_DATE="$cutoff_date" GIT_COMMITTER_DATE="$cutoff_date" git -C "$repo" commit-tree -S "$cutoff_tree")
printf 'Repository: %s\nBackup branch: %s\nAggregate commit: %s\nReplaying %s commits after %s\n' \
  "$repo" "$backup_branch" "$aggregate" "$old_count" "$cutoff"

git -C "$repo" rebase --onto "$aggregate" "$cutoff" "$branch"

new_head=$(git -C "$repo" rev-parse "$branch")
new_tree=$(git -C "$repo" rev-parse "$new_head^{tree}")
new_count=$(git -C "$repo" rev-list --count "$aggregate..$new_head")

[[ "$new_tree" == "$old_tree" ]] || fail "The final tree differs from the original. Restore from $backup_branch if needed."
[[ "$new_count" == "$old_count" ]] || fail "The retained commit count changed. Restore from $backup_branch if needed."
if git -C "$repo" merge-base --is-ancestor "$cutoff" "$new_head"; then
  fail "The old cutoff is still an ancestor. Restore from $backup_branch if needed."
fi
[[ -z "$(git -C "$repo" status --porcelain)" ]] || fail "The publish repository is not clean after rewriting."

printf 'History rewrite verified.\nNew master: %s\nTree: %s\nRetained commits: %s\n' \
  "$new_head" "$new_tree" "$new_count"
