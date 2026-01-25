package cmd

import (
	"errors"
	"fmt"
	"strings"

	"tmdb-cli/internal/i18n"
	"tmdb-cli/internal/tmdb"
	"tmdb-cli/internal/ui"

	"github.com/spf13/cobra"
)

var movieCmd = &cobra.Command{
	Use:     "movie <suchbegriff>",
	Aliases: []string{"m", "film"},
	Short:   "Suche nach Filminformationen",
	Long: `Sucht nach einem Film und zeigt detaillierte Informationen an.

Beispiele:
  tmdb movie "The Matrix"
  tmdb movie "Inception" --short
  tmdb movie "Pulp Fiction" --json
  tmdb m "Fight Club"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runMovie,
}

func runMovie(cmd *cobra.Command, args []string) error {
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
	results, err := client.SearchMovies(query, lang)
	if err != nil {
		return fmt.Errorf("Suche fehlgeschlagen: %w", err)
	}

	// Keine Ergebnisse
	if len(results) == 0 {
		fmt.Println(ui.RenderInfo(i18n.Translatef(i18n.KeyNoMoviesFound, lang, query)))
		return nil
	}

	// Film-ID bestimmen
	var movieID int
	if len(results) == 1 {
		movieID = results[0].ID
	} else {
		// Interaktive Auswahl
		selectedID, err := ui.SelectMovie(results, lang)
		if err != nil {
			return fmt.Errorf("Auswahl fehlgeschlagen: %w", err)
		}
		if selectedID == -1 {
			// Abgebrochen
			return nil
		}
		movieID = selectedID
	}

	// Details laden
	movie, err := client.GetMovieDetails(movieID, lang)
	if err != nil {
		return fmt.Errorf("Details konnten nicht geladen werden: %w", err)
	}

	// Ausgabe
	if jsonOutput {
		output, err := ui.RenderMovieJSON(movie)
		if err != nil {
			return fmt.Errorf("JSON-Ausgabe fehlgeschlagen: %w", err)
		}
		fmt.Println(output)
	} else {
		fmt.Println(ui.RenderMovieDetails(movie, shortOutput, lang))
	}

	return nil
}
