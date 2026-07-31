---
name: tagtools repository migration
description: Checklist for migrating tagtools to an independent repository
title: "[tagtools] Independent repository migration"
labels: ["maintenance", "tooling", "tagtools"]
assignees: []
---

## Summary

- Goal: Migrate `tagtools` into an independent repository while preserving daily operation in this repository.
- Scope: release policy，security workflow，CI，connection strategy，rollback plan.

## Background

- Current code location: `tagtools/`
- Current wrappers: `tagslist.sh`，`toptags.sh`
- Current build switch: `build.sh` with `TAGTOOLS_EXEC` (`wrapper` or `go`)

## 1. Repository Design

- [ ] Create a dedicated `tagtools` repository.
- [ ] Decide module path (`github.com/<owner>/tagtools`).
- [ ] Decide policy: CLI-only or library export.

## 2. Release Policy

- [ ] Define SemVer policy (`v0.x` to `v1.0.0`).
- [ ] Define breaking-change handling and major bump rules.
- [ ] Decide distribution channel (GitHub Releases binary / `go install` / source-only).

## 3. Security Maintenance

- [ ] Enable dependency update automation (Dependabot or Renovate).
- [ ] Add CI checks: `go test`，`go vet`，`govulncheck`.
- [ ] Add checksum publication to release process.
- [ ] Evaluate signing strategy (GPG or cosign).

## 4. Test Strategy Split

- [ ] Split repository-local fixture tests from site-content integration tests.
- [ ] Keep minimal fixtures in the independent repository.
- [ ] Keep integration verification in this repository.

## 5. Connection from This Repository

- [ ] Preserve wrapper compatibility in `tagslist.sh` and `toptags.sh`.
- [ ] Preserve env compatibility (`TOP_N` etc.).
- [ ] Decide fetch/install strategy for external `tagtools` binary.
- [ ] Verify `build.sh` switch behavior for staged rollout.

## 6. Migration Steps

- [ ] Create new repository and baseline CI.
- [ ] Move current `tagtools/` and update module path.
- [ ] Cut trial release (`v0.x`) and run smoke test in this repository.
- [ ] Cut stable release (`v1.0.0`) when verified.
- [ ] Switch wrappers/build to external distribution.

## 7. Rollback Plan

- [ ] Keep fallback scripts in `tagtools/orgsh/` during migration window.
- [ ] Document rollback commands and validation checks.
- [ ] Define rollback trigger criteria and owner.

## Acceptance Criteria

- [ ] Existing outputs remain compatible (`tagslist.csv` and `toptags.json`).
- [ ] CI and security checks are green in the new repository.
- [ ] This repository can switch between wrapper and external distribution without regression.

## Notes

- Related plan document: `docs/go-tag-tools-migration-plan.md`
