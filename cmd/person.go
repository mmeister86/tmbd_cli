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

var personCmd = &cobra.Command{
	Use:     "person <suchbegriff>",
	Aliases: []string{"p", "actor"},
	Short:   "Suche nach Personeninformationen",
	Long: `Sucht nach einer Person und zeigt detaillierte Informationen an.

Beispiele:
  tmdb person "Tom Hanks"
  tmdb person "Meryl Streep" --short
  tmdb person "Leonardo DiCaprio" --json
  tmdb p "Brad Pitt"`,
	Args: cobra.MinimumNArgs(1),
	RunE: runPerson,
}

func runPerson(cmd *cobra.Command, args []string) error {
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
	results, err := client.SearchPeople(query, lang)
	if err != nil {
		return fmt.Errorf("Suche fehlgeschlagen: %w", err)
	}

	// Keine Ergebnisse
	if len(results) == 0 {
		fmt.Println(ui.RenderInfo(i18n.Translatef(i18n.KeyNoPersonsFound, lang, query)))
		return nil
	}

	// Personen-ID bestimmen
	var personID int
	if len(results) == 1 {
		personID = results[0].ID
	} else {
		// Interaktive Auswahl
		selectedID, err := ui.SelectPerson(results, lang)
		if err != nil {
			return fmt.Errorf("Auswahl fehlgeschlagen: %w", err)
		}
		if selectedID == -1 {
			// Abgebrochen
			return nil
		}
		personID = selectedID
	}

	// Details laden
	person, err := client.GetPersonDetails(personID, lang)
	if err != nil {
		return fmt.Errorf("Details konnten nicht geladen werden: %w", err)
	}

	// Ausgabe
	if jsonOutput {
		output, err := ui.RenderPersonJSON(person)
		if err != nil {
			return fmt.Errorf("JSON-Ausgabe fehlgeschlagen: %w", err)
		}
		fmt.Println(output)
	} else {
		fmt.Println(ui.RenderPersonDetails(person, shortOutput, lang))
	}

	return nil
}
