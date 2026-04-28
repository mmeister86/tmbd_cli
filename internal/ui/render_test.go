package ui

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/mmeister86/tmbd_cli/internal/tmdb"
)

func TestTruncateTextKeepsUTF8Valid(t *testing.T) {
	input := strings.Repeat("ä", 80)

	got := truncateText(input, 24)

	if !utf8.ValidString(got) {
		t.Fatalf("truncateText returned invalid UTF-8: %q", got)
	}
	if !strings.HasSuffix(got, "...") {
		t.Fatalf("truncateText should append ellipsis, got %q", got)
	}
}

func TestWrapTextDoesNotEmitEmptyLineForLongFirstWord(t *testing.T) {
	got := wrapText("supercalifragilisticexpialidocious short", 10)

	if strings.HasPrefix(got, "\n") {
		t.Fatalf("wrapText emitted an empty first line: %q", got)
	}
}

func TestKnownForCreditsDoesNotDuplicateCastEntries(t *testing.T) {
	credits := &tmdb.CombinedCredits{
		Cast: []tmdb.CombinedCast{
			{ID: 1, Title: "Older Work", Popularity: 1, Order: 1},
			{ID: 2, Title: "Popular Work", Popularity: 100, Order: 99},
		},
	}

	got := getKnownForCredits(credits)

	if len(got) != 2 {
		t.Fatalf("expected 2 known-for works, got %d: %#v", len(got), got)
	}
	if got[0].ID == got[1].ID {
		t.Fatalf("known-for works should not contain duplicate cast entries: %#v", got)
	}
	if got[0].ID != 2 {
		t.Fatalf("known-for works should keep popularity order, got %#v", got)
	}
}
