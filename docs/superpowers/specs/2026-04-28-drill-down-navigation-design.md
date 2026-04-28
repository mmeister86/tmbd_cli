# Drill-Down Navigation Design

## Goal

After an interactive movie or series search result is selected, the CLI should keep the user in context and offer related information without requiring a new command. The existing detail view remains the first thing shown.

## Interaction

- `tmdb movie <query>` with normal output loads and renders the selected movie as today, then shows an action menu.
- Movie actions:
  - Cast: choose a cast member and render that person's detail view.
  - Directors: choose a director from movie credits and render that person's detail view.
  - Back/Exit: return to the action menu after a detail drill-down, or leave the command.
- `tmdb series <query>` with normal output loads and renders the selected series as today, then shows an action menu.
- Series actions:
  - Cast: choose a cast member and render that person's detail view.
  - Creators/Crew: choose from creators plus directing/writing crew when available and render person details.
  - Seasons: choose a season and render season details with episodes.
  - Back/Exit: return to the action menu after a drill-down, or leave the command.

## Non-Interactive Modes

`--json` and `--short` remain one-shot outputs. They do not open drill-down menus, preserving script-friendly behavior and compact output.

## API

Movie and TV detail calls already use `append_to_response=credits`, which provides cast and crew IDs. Person details already append `combined_credits`. Add a TV season details call using `/tv/{tv_id}/season/{season_number}` with the selected language.

## Code Shape

Keep the existing command files as orchestration. Add small UI selection helpers for actions, people, and seasons. Add TMDB season response types and a season renderer. Avoid a broad refactor of the HTTP client during this feature.

## Testing

Cover pure behavior with unit tests: action option construction, unique person option construction, season rendering, and client URL construction where practical. Run the existing full Go verification before completion.
