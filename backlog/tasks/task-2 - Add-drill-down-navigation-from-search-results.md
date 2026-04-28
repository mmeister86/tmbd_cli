---
id: TASK-2
title: Add drill-down navigation from search results
status: In Progress
assignee: []
created_date: '2026-04-28 19:35'
updated_date: '2026-04-28 19:47'
labels: []
dependencies: []
priority: high
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Extend the interactive search flow so users can navigate from selected movies, series, and people into related actor, director, and season information without leaving the CLI workflow.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Movie detail flow allows opening cast and director person details from results
- [x] #2 Series detail flow allows opening cast, creator/director-style people, and season details
- [x] #3 Navigation works from interactive search results and has a clear way to go back or exit
- [x] #4 JSON and short output behavior remains script-friendly and non-interactive
- [x] #5 Regression tests cover navigation and rendering changes
<!-- AC:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Design the post-detail action menu and navigation loop.
2. Add tests for menu option construction and season rendering/client behavior.
3. Add TMDB season details types/client method.
4. Add reusable post-detail navigation for movies and series.
5. Keep --json and --short non-interactive.
6. Run gofmt, tests, vet, build, and install dry-run verification.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Context gathered: current commands load a selected movie/series/person and render details once. Movie and TV details already append credits. Person details append combined_credits. Context7 TMDB docs confirm person combined credits and TV season details endpoints: /person/{person_id}/combined_credits and /tv/{tv_id}/season/{season_number}, plus season credits endpoints.

Implemented drill-down primitives and command loops: movie actions for cast/directors, series actions for cast/creator-crew/seasons, season details endpoint and renderer, and non-interactive --json/--short behavior preserved by only opening menus after full normal output.

Verification started after final polish: gofmt plus go test -count=1 ./..., go vet ./..., go build ./..., scripts/test_make_install.sh, and make -n install.

Bugfix: drill-down detail views flashed because the loop immediately reopened a Bubble Tea alt-screen menu after printing details. Added a tested WaitForEnter pause after person and season detail rendering so the view remains visible until the user continues.

Bugfix: initial movie/series details also flashed because the action menu opened immediately in alt-screen. Added an explicit Enter pause after initial full details before opening the drill-down menu; WaitForEnter now accepts a prompt so initial and return-to-menu prompts are clear.
<!-- SECTION:NOTES:END -->
