# Drill-Down Navigation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add interactive drill-down menus after normal movie and series detail output.

**Architecture:** Keep command handlers as orchestration and add small UI helpers for action, person, and season selection. Reuse existing TMDB detail calls for credits, add a season details endpoint, and keep `--json`/`--short` one-shot.

**Tech Stack:** Go, Cobra, Bubble Tea/Bubbles list, Lipgloss, TMDB v3 API.

---

### Task 1: Menu And Rendering Primitives

**Files:**
- Create: `internal/ui/navigation_test.go`
- Create: `internal/ui/navigation.go`
- Modify: `internal/ui/render.go`
- Test: `internal/ui/navigation_test.go`

- [x] Write failing tests for movie actions, director/cast option building, season filtering, and season rendering.
- [x] Implement minimal UI option builders and season renderer.
- [x] Run `go test ./internal/ui`.

### Task 2: TMDB Season Endpoint

**Files:**
- Modify: `internal/tmdb/types.go`
- Modify: `internal/tmdb/client.go`

- [ ] Add `SeasonDetails` and `Episode` response types.
- [ ] Add `GetTVSeasonDetails(tvID, seasonNumber int, language string)`.
- [ ] Run `go test ./internal/tmdb`.

### Task 3: Command Navigation Loop

**Files:**
- Modify: `cmd/movie.go`
- Modify: `cmd/series.go`

- [ ] After normal full output, show an action menu.
- [ ] Movie actions open cast or directors as person details.
- [ ] Series actions open cast, creators/crew as person details, or season details.
- [ ] Skip drill-down for `--json` and `--short`.

### Task 4: Verification

**Files:**
- Modify: `backlog/tasks/task-2 - Add-drill-down-navigation-from-search-results.md`

- [ ] Run `go test -count=1 ./...`.
- [ ] Run `go vet ./...`.
- [ ] Run `go build ./...`.
- [ ] Run `sh scripts/test_make_install.sh` and `make -n install`.
