---
id: TASK-1
title: Bring codebase up to maintenance baseline
status: In Progress
assignee: []
created_date: '2026-04-28 19:25'
updated_date: '2026-04-28 19:29'
labels: []
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Clean up the TMDB CLI codebase by fixing confirmed CLI bugs, version wiring, rendering/i18n issues, and adding focused regression tests while preserving the existing product scope.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Duplicate language command registration is removed
- [x] #2 Version reported by the CLI is driven consistently by build-time metadata
- [x] #3 Known-for person output no longer duplicates cast entries
- [x] #4 Text truncation and wrapping are safe for UTF-8 input
- [x] #5 Focused regression tests cover the changed behavior
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add regression tests for duplicated command registration, version propagation, person known-for selection, and UTF-8 text helpers.
2. Run tests to confirm the new tests fail for the existing issues.
3. Fix the smallest production surfaces needed: command registration/version wiring and UI helper behavior.
4. Run gofmt, tests, vet, and build.
5. Update Backlog acceptance criteria and notes; leave task In Progress until user confirms manual testing.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Context7 used for Cobra docs: /spf13/cobra. Confirmed AddCommand is the registration point and root command Version drives --version.
Red tests added before fixes: duplicate language command and wrapText long-first-word behavior failed on the original code.
Implemented focused fixes for command registration, build-time version ldflags, UTF-8-safe truncation, long-word wrapping, and duplicate known-for cast collection.

Sharpened known-for regression to verify popularity order remains intact after deduplication; removed the final cast-order resort.

Verification: go test -count=1 ./..., go vet ./..., go build ./..., and make build followed by ./tmdb --version and ./tmdb --help all exited successfully. Help output now lists language once; version output is 1.0.2.

Follow-up install failure: /usr/local existed but /usr/local/bin did not. Added INSTALL_DIR and SUDO Makefile variables, create INSTALL_DIR before moving the binary, and added scripts/test_make_install.sh to regression-test install into a temporary directory without sudo.
<!-- SECTION:NOTES:END -->
