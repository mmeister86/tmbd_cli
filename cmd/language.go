package cmd

import (
	"fmt"

	"github.com/mmeister86/tmbd_cli/internal/config"
	"github.com/mmeister86/tmbd_cli/internal/i18n"
	"github.com/mmeister86/tmbd_cli/internal/ui"

	"github.com/spf13/cobra"
)

var languageCmd = &cobra.Command{
	Use:   "language",
	Short: "Sprache auswählen oder anzeigen",
	Long: `Zeigt die aktuelle Sprache an oder ermöglicht die interaktive Auswahl einer neuen Sprache.

Die Sprachauswahl wird in der Konfigurationsdatei gespeichert und für alle zukünftigen
Anfragen verwendet. Mit dem --language Flag kann die Sprache für einzelne Anfragen überschrieben werden.`,
	RunE: runLanguage,
}

func init() {
	rootCmd.AddCommand(languageCmd)
}

func runLanguage(cmd *cobra.Command, args []string) error {
	// Aktuelle Konfiguration laden
	cfg, err := config.Load()
	if err != nil {
		if err == config.ErrConfigNotFound {
			cfg = config.GetDefaultConfig()
		} else {
			return fmt.Errorf("konnte Konfiguration nicht laden: %w", err)
		}
	}

	// Aktuelle Sprache anzeigen
	currentLang := cfg.Language
	currentLangName := i18n.GetLanguageName(currentLang)
	fmt.Printf("Aktuelle Sprache: %s (%s)\n\n", currentLangName, currentLang)

	// Interaktive Sprachauswahl
	selectedLang, err := ui.SelectLanguage()
	if err != nil {
		return fmt.Errorf("Sprachauswahl fehlgeschlagen: %w", err)
	}

	if selectedLang == "" {
		// Abgebrochen
		fmt.Println(ui.RenderInfo("Sprachauswahl abgebrochen"))
		return nil
	}

	// Sprache speichern
	cfg.Language = selectedLang
	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("konnte Sprache nicht speichern: %w", err)
	}

	selectedLangName := i18n.GetLanguageName(selectedLang)
	fmt.Printf("✓ Sprache auf %s (%s) geändert\n", selectedLangName, selectedLang)

	return nil
}
