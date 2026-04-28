package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mmeister86/tmbd_cli/internal/i18n"
	"github.com/mmeister86/tmbd_cli/internal/tmdb"
	"github.com/mmeister86/tmbd_cli/internal/ui"

	"github.com/spf13/cobra"
)

var seriesCmd = &cobra.Command{
	Use:     "series <suchbegriff>",
	Aliases: []string{"s", "tv", "show"},
	Short:   "Suche nach Serieninformationen",
	Long: `Sucht nach einer Serie und zeigt detaillierte Informationen an.

Beispiele:
  tmdb series "Breaking Bad"
  tmdb series "Stranger Things" --short
  tmdb tv "The Office" --json
  tmdb s "Dark"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runSeries,
}

func runSeries(cmd *cobra.Command, args []string) error {
	query := strings.Join(args, " ")
	lang := getLanguage()

	// Client erstellen
	client, err := tmdb.NewClient()
	if err != nil {
		if errors.Is(err, tmdb.ErrNoAPIKey) {
			fmt.Println(ui.RenderError(
				i18n.Translate(i18n.KeyErrorNoAPIKey, lang),
				i18n.Translate(i18n.KeyErrorSetAPIKey, lang),
				[]string{
					"  export TMDB_API_KEY='dein-api-key'",
					"",
					i18n.Translate(i18n.KeyGetAPIKeyURL, lang),
				},
				lang,
			))
			return nil
		}
		return err
	}

	// Suche durchführen
	results, err := client.SearchTV(query, lang)
	if err != nil {
		return fmt.Errorf("Suche fehlgeschlagen: %w", err)
	}

	// Keine Ergebnisse
	if len(results) == 0 {
		fmt.Println(ui.RenderInfo(i18n.Translatef(i18n.KeyNoSeriesFound, lang, query)))
		return nil
	}

	// Serien-ID bestimmen
	var tvID int
	if len(results) == 1 {
		tvID = results[0].ID
	} else {
		// Interaktive Auswahl
		selectedID, err := ui.SelectTV(results, lang)
		if err != nil {
			return fmt.Errorf("Auswahl fehlgeschlagen: %w", err)
		}
		if selectedID == -1 {
			// Abgebrochen
			return nil
		}
		tvID = selectedID
	}

	// Details laden
	tv, err := client.GetTVDetails(tvID, lang)
	if err != nil {
		return fmt.Errorf("Details konnten nicht geladen werden: %w", err)
	}

	// Ausgabe
	if jsonOutput {
		output, err := ui.RenderTVJSON(tv)
		if err != nil {
			return fmt.Errorf("JSON-Ausgabe fehlgeschlagen: %w", err)
		}
		fmt.Println(output)
	} else {
		fmt.Println(ui.RenderTVDetails(tv, shortOutput, lang))
		if !shortOutput {
			if err := ui.WaitForEnter(os.Stdin, os.Stdout, "Enter drücken, um weitere Optionen anzuzeigen..."); err != nil {
				return err
			}
			return runSeriesDrillDown(client, tv, lang)
		}
	}

	return nil
}

func runSeriesDrillDown(client *tmdb.Client, tv *tmdb.TVDetails, lang string) error {
	for {
		action, err := ui.SelectAction(ui.TVDrillDownActions(tv, lang), "Weitere Informationen")
		if err != nil {
			return fmt.Errorf("Auswahl fehlgeschlagen: %w", err)
		}
		switch action {
		case "", ui.ActionExit:
			return nil
		case ui.ActionCast:
			if err := renderSelectedPerson(client, ui.PeopleFromCast(tv.Credits.Cast), "Besetzung", lang); err != nil {
				return err
			}
		case ui.ActionCrew:
			if err := renderSelectedPerson(client, ui.TVCrewPeople(tv), "Creator & Crew", lang); err != nil {
				return err
			}
		case ui.ActionSeasons:
			if err := renderSelectedSeason(client, tv, lang); err != nil {
				return err
			}
		}
	}
}

func renderSelectedSeason(client *tmdb.Client, tv *tmdb.TVDetails, lang string) error {
	seasons := ui.SelectableSeasons(tv.Seasons)
	if len(seasons) == 0 {
		fmt.Println(ui.RenderInfo("Keine Staffeln verfügbar"))
		return nil
	}

	seasonNumber, err := ui.SelectSeasonOption(seasons, "Staffel wählen")
	if err != nil {
		return fmt.Errorf("Staffelauswahl fehlgeschlagen: %w", err)
	}
	if seasonNumber == -1 {
		return nil
	}

	season, err := client.GetTVSeasonDetails(tv.ID, seasonNumber, lang)
	if err != nil {
		return fmt.Errorf("Staffeldetails konnten nicht geladen werden: %w", err)
	}
	fmt.Println(ui.RenderSeasonDetails(season, lang))
	return ui.WaitForEnter(os.Stdin, os.Stdout, "Enter drücken, um zum Menü zurückzukehren...")
}
