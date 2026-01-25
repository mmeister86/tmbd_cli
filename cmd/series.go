package cmd

import (
	"errors"
	"fmt"
	"strings"

	"tmdb-cli/internal/tmdb"
	"tmdb-cli/internal/ui"

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
				"Fehler: TMDB_API_KEY nicht gesetzt",
				"Setze deinen API Key mit:",
				[]string{
					"  export TMDB_API_KEY='dein-api-key'",
					"",
					"API Key erhältst du unter:",
					"  https://www.themoviedb.org/settings/api",
				},
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
		fmt.Println(ui.RenderInfo(fmt.Sprintf("Keine Serien gefunden für: %s", query)))
		return nil
	}

	// Serien-ID bestimmen
	var tvID int
	if len(results) == 1 {
		tvID = results[0].ID
	} else {
		// Interaktive Auswahl
		selectedID, err := ui.SelectTV(results)
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
		fmt.Println(ui.RenderTVDetails(tv, shortOutput))
	}

	return nil
}
