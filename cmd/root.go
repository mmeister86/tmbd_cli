package cmd

import (
	"fmt"
	"os"

	"tmdb-cli/internal/tmdb"

	"github.com/spf13/cobra"
)

// Version wird beim Build gesetzt
var Version = "1.0.0"

// Globale Flags
var (
	jsonOutput  bool
	shortOutput bool
	language    string
)

// rootCmd ist der Basis-Command
var rootCmd = &cobra.Command{
	Use:   "tmdb",
	Short: "CLI Tool für TMDB Film- und Serieninformationen",
	Long: `tmdb ist ein Command-Line Tool zum Abrufen von Informationen
über Filme und Serien von The Movie Database (TMDB).

Beispiele:
  tmdb movie "The Matrix"
  tmdb series "Breaking Bad"
  tmdb m "Inception" --short
  tmdb tv "Dark" --json`,
	Version: Version,
}

// Execute führt den Root-Command aus
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Globale Flags
	rootCmd.PersistentFlags().BoolVar(&jsonOutput, "json", false, "Ausgabe als JSON")
	rootCmd.PersistentFlags().BoolVarP(&shortOutput, "short", "s", false, "Kompakte Ausgabe")
	rootCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "Sprache überschreiben (z.B. en-US)")

	// Subcommands hinzufügen
	rootCmd.AddCommand(movieCmd)
	rootCmd.AddCommand(seriesCmd)
}

// getLanguage gibt die zu verwendende Sprache zurück
func getLanguage() string {
	if language != "" {
		return language
	}
	return tmdb.GetLanguage()
}
