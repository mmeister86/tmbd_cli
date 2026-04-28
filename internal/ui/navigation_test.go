package ui

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mmeister86/tmbd_cli/internal/tmdb"
)

func TestMovieDrillDownActionsReflectAvailableCredits(t *testing.T) {
	movie := &tmdb.MovieDetails{
		Credits: &tmdb.Credits{
			Cast: []tmdb.CastMember{{ID: 1, Name: "Keanu Reeves"}},
			Crew: []tmdb.CrewMember{{ID: 2, Name: "Lana Wachowski", Job: "Director"}},
		},
	}

	actions := MovieDrillDownActions(movie, "de-DE")

	assertActionIDs(t, actions, []string{"cast", "directors", "exit"})
}

func TestUniquePeopleFromCastRemovesDuplicateIDs(t *testing.T) {
	cast := []tmdb.CastMember{
		{ID: 1, Name: "Actor", Character: "Neo"},
		{ID: 1, Name: "Actor", Character: "Thomas Anderson"},
		{ID: 2, Name: "Other Actor", Character: "Morpheus"},
	}

	people := PeopleFromCast(cast)

	if len(people) != 2 {
		t.Fatalf("expected 2 unique people, got %d: %#v", len(people), people)
	}
	if people[0].Description != "Neo" {
		t.Fatalf("expected first role to be preserved, got %q", people[0].Description)
	}
}

func TestPeopleFromCrewFiltersDirectors(t *testing.T) {
	crew := []tmdb.CrewMember{
		{ID: 1, Name: "Director", Job: "Director"},
		{ID: 2, Name: "Producer", Job: "Producer"},
	}

	people := PeopleFromCrew(crew, "Director")

	if len(people) != 1 {
		t.Fatalf("expected 1 director, got %d: %#v", len(people), people)
	}
	if people[0].ID != 1 {
		t.Fatalf("expected director ID 1, got %#v", people[0])
	}
}

func TestSelectableSeasonsSkipSpecials(t *testing.T) {
	seasons := []tmdb.Season{
		{SeasonNumber: 0, Name: "Specials"},
		{SeasonNumber: 1, Name: "Season 1", EpisodeCount: 10},
	}

	got := SelectableSeasons(seasons)

	if len(got) != 1 {
		t.Fatalf("expected 1 selectable season, got %d: %#v", len(got), got)
	}
	if got[0].ID != "1" {
		t.Fatalf("expected season number 1 as ID, got %#v", got[0])
	}
}

func TestRenderSeasonDetailsIncludesEpisodes(t *testing.T) {
	season := &tmdb.SeasonDetails{
		Name:         "Season 1",
		SeasonNumber: 1,
		AirDate:      "1999-01-01",
		Overview:     "The beginning.",
		Episodes: []tmdb.Episode{
			{EpisodeNumber: 1, Name: "Pilot", AirDate: "1999-01-01", VoteAverage: 8.5},
		},
	}

	got := RenderSeasonDetails(season, "en-US")

	for _, want := range []string{"Season 1", "The beginning.", "1. Pilot", "8.5/10"} {
		if !strings.Contains(got, want) {
			t.Fatalf("expected rendered season to contain %q, got %q", want, got)
		}
	}
}

func TestFormatSearchItemTitleOmitsEmptyMetadata(t *testing.T) {
	got := formatSearchItemTitle(searchItem{title: "Besetzung"})

	if got != "Besetzung" {
		t.Fatalf("expected bare menu title, got %q", got)
	}
}

func TestWaitForEnterPrintsPromptAndBlocksUntilNewline(t *testing.T) {
	input := strings.NewReader("\n")
	var output bytes.Buffer

	err := WaitForEnter(input, &output, "Weiter mit Enter...")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !strings.Contains(output.String(), "Weiter mit Enter...") {
		t.Fatalf("expected custom prompt, got %q", output.String())
	}
}

func assertActionIDs(t *testing.T, actions []DrillDownOption, want []string) {
	t.Helper()
	if len(actions) != len(want) {
		t.Fatalf("expected %d actions, got %d: %#v", len(want), len(actions), actions)
	}
	for i := range want {
		if actions[i].ID != want[i] {
			t.Fatalf("action %d: expected %q, got %q", i, want[i], actions[i].ID)
		}
	}
}
