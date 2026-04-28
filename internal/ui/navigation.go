package ui

import (
	"bufio"
	"fmt"
	"io"

	"github.com/mmeister86/tmbd_cli/internal/tmdb"
)

const (
	ActionCast      = "cast"
	ActionDirectors = "directors"
	ActionCrew      = "crew"
	ActionSeasons   = "seasons"
	ActionExit      = "exit"
)

// DrillDownOption beschreibt einen auswählbaren Eintrag in Drill-down-Menüs.
type DrillDownOption struct {
	ID          string
	Title       string
	Description string
}

// PersonOption beschreibt eine auswählbare Person.
type PersonOption struct {
	ID          int
	Name        string
	Description string
}

// MovieDrillDownActions erzeugt die verfügbaren Aktionen für einen Film.
func MovieDrillDownActions(movie *tmdb.MovieDetails, language string) []DrillDownOption {
	var options []DrillDownOption
	if movie != nil && movie.Credits != nil && len(movie.Credits.Cast) > 0 {
		options = append(options, DrillDownOption{ID: ActionCast, Title: "Besetzung", Description: "Schauspielerinnen und Schauspieler anzeigen"})
	}
	if movie != nil && movie.Credits != nil && len(PeopleFromCrew(movie.Credits.Crew, "Director")) > 0 {
		options = append(options, DrillDownOption{ID: ActionDirectors, Title: "Regie", Description: "Regisseurinnen und Regisseure anzeigen"})
	}
	return append(options, DrillDownOption{ID: ActionExit, Title: "Beenden", Description: "Zurück zur Shell"})
}

// TVDrillDownActions erzeugt die verfügbaren Aktionen für eine Serie.
func TVDrillDownActions(tv *tmdb.TVDetails, language string) []DrillDownOption {
	var options []DrillDownOption
	if tv != nil && tv.Credits != nil && len(tv.Credits.Cast) > 0 {
		options = append(options, DrillDownOption{ID: ActionCast, Title: "Besetzung", Description: "Schauspielerinnen und Schauspieler anzeigen"})
	}
	if tv != nil && len(TVCrewPeople(tv)) > 0 {
		options = append(options, DrillDownOption{ID: ActionCrew, Title: "Creator & Crew", Description: "Kreative Personen anzeigen"})
	}
	if tv != nil && len(SelectableSeasons(tv.Seasons)) > 0 {
		options = append(options, DrillDownOption{ID: ActionSeasons, Title: "Staffeln", Description: "Staffeldetails und Episoden anzeigen"})
	}
	return append(options, DrillDownOption{ID: ActionExit, Title: "Beenden", Description: "Zurück zur Shell"})
}

// PeopleFromCast baut eindeutige Personenoptionen aus Cast-Einträgen.
func PeopleFromCast(cast []tmdb.CastMember) []PersonOption {
	people := make([]PersonOption, 0, len(cast))
	seen := make(map[int]bool)
	for _, member := range cast {
		if member.ID == 0 || seen[member.ID] {
			continue
		}
		seen[member.ID] = true
		people = append(people, PersonOption{
			ID:          member.ID,
			Name:        member.Name,
			Description: member.Character,
		})
	}
	return people
}

// PeopleFromCrew baut eindeutige Personenoptionen für einen Crew-Job.
func PeopleFromCrew(crew []tmdb.CrewMember, job string) []PersonOption {
	people := make([]PersonOption, 0, len(crew))
	seen := make(map[int]bool)
	for _, member := range crew {
		if member.ID == 0 || seen[member.ID] || member.Job != job {
			continue
		}
		seen[member.ID] = true
		people = append(people, PersonOption{
			ID:          member.ID,
			Name:        member.Name,
			Description: member.Job,
		})
	}
	return people
}

// TVCrewPeople baut eindeutige Creator-, Regie- und Schreib-Crew-Optionen.
func TVCrewPeople(tv *tmdb.TVDetails) []PersonOption {
	if tv == nil {
		return nil
	}

	var people []PersonOption
	seen := make(map[int]bool)
	for _, creator := range tv.CreatedBy {
		if creator.ID == 0 || seen[creator.ID] {
			continue
		}
		seen[creator.ID] = true
		people = append(people, PersonOption{ID: creator.ID, Name: creator.Name, Description: "Creator"})
	}
	if tv.Credits == nil {
		return people
	}
	for _, member := range tv.Credits.Crew {
		if member.ID == 0 || seen[member.ID] {
			continue
		}
		if member.Job != "Director" && member.Department != "Writing" {
			continue
		}
		seen[member.ID] = true
		description := member.Job
		if description == "" {
			description = member.Department
		}
		people = append(people, PersonOption{ID: member.ID, Name: member.Name, Description: description})
	}
	return people
}

// SelectableSeasons baut auswählbare Staffeln und überspringt Specials.
func SelectableSeasons(seasons []tmdb.Season) []DrillDownOption {
	options := make([]DrillDownOption, 0, len(seasons))
	for _, season := range seasons {
		if season.SeasonNumber == 0 {
			continue
		}
		description := fmt.Sprintf("%d Episoden", season.EpisodeCount)
		if season.AirDate != "" {
			description = fmt.Sprintf("%s, %s", description, season.AirDate)
		}
		options = append(options, DrillDownOption{
			ID:          fmt.Sprintf("%d", season.SeasonNumber),
			Title:       season.Name,
			Description: description,
		})
	}
	return options
}

// WaitForEnter hält die aktuelle Ansicht sichtbar, bis der Nutzer weitergeht.
func WaitForEnter(input io.Reader, output io.Writer, prompt string) error {
	if prompt == "" {
		prompt = "Enter drücken, um fortzufahren..."
	}
	if _, err := fmt.Fprintf(output, "\n%s", prompt); err != nil {
		return err
	}
	_, err := bufio.NewReader(input).ReadString('\n')
	return err
}
