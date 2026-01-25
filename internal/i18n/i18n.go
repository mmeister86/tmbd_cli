package i18n

import (
	"fmt"
)

// TranslationKeys enthält alle Übersetzungsschlüssel
const (
	KeyMovieIcon          = "movie_icon"
	KeySeriesIcon         = "series_icon"
	KeyLabelYear          = "label_year"
	KeyLabelRuntime       = "label_runtime"
	KeyLabelGenre         = "label_genre"
	KeyLabelRating        = "label_rating"
	KeyLabelStatus        = "label_status"
	KeyLabelBudget        = "label_budget"
	KeyLabelRevenue       = "label_revenue"
	KeyLabelDirector      = "label_director"
	KeyLabelPeriod        = "label_period"
	KeyLabelSeasons       = "label_seasons"
	KeyLabelEpisodes      = "label_episodes"
	KeyLabelEpisodeLength = "label_episode_length"
	KeyLabelNetwork       = "label_network"
	KeyLabelCreatedBy     = "label_created_by"
	KeySectionPlot        = "section_plot"
	KeySectionCast        = "section_cast"
	KeySectionSeasons      = "section_seasons"
	KeySectionLinks       = "section_links"
	KeyNoMoviesFound      = "no_movies_found"
	KeyNoSeriesFound      = "no_series_found"
	KeySelectMovie        = "select_movie"
	KeySelectSeries       = "select_series"
	KeySelectAction       = "select_action"
	KeyCancelAction      = "cancel_action"
	KeyErrorNoAPIKey     = "error_no_api_key"
	KeyErrorSetAPIKey    = "error_set_api_key"
	KeyGetAPIKeyURL      = "get_api_key_url"
)

// translations enthält alle Übersetzungen pro Sprache
var translations = map[string]map[string]string{
	"de-DE": {
		KeyMovieIcon:          "🎬",
		KeySeriesIcon:         "📺",
		KeyLabelYear:          "Jahr",
		KeyLabelRuntime:       "Laufzeit",
		KeyLabelGenre:         "Genre",
		KeyLabelRating:        "Bewertung",
		KeyLabelStatus:        "Status",
		KeyLabelBudget:        "Budget",
		KeyLabelRevenue:       "Einspielergebnis",
		KeyLabelDirector:      "Regie",
		KeyLabelPeriod:        "Zeitraum",
		KeyLabelSeasons:       "Staffeln",
		KeyLabelEpisodes:      "Episoden",
		KeyLabelEpisodeLength: "Episodenlänge",
		KeyLabelNetwork:       "Sender",
		KeyLabelCreatedBy:     "Erstellt von",
		KeySectionPlot:        "Handlung",
		KeySectionCast:        "Besetzung",
		KeySectionSeasons:      "Staffeln",
		KeySectionLinks:       "Links",
		KeyNoMoviesFound:      "Keine Filme gefunden für: %s",
		KeyNoSeriesFound:      "Keine Serien gefunden für: %s",
		KeySelectMovie:        "Wähle einen Film",
		KeySelectSeries:       "Wähle eine Serie",
		KeySelectAction:       "auswählen",
		KeyCancelAction:       "abbrechen",
		KeyErrorNoAPIKey:     "Fehler: TMDB_API_KEY nicht gesetzt",
		KeyErrorSetAPIKey:    "Setze deinen API Key mit:",
		KeyGetAPIKeyURL:      "API Key erhältst du unter:\n  https://www.themoviedb.org/settings/api",
	},
	"en-US": {
		KeyMovieIcon:          "🎬",
		KeySeriesIcon:         "📺",
		KeyLabelYear:          "Year",
		KeyLabelRuntime:       "Runtime",
		KeyLabelGenre:         "Genre",
		KeyLabelRating:        "Rating",
		KeyLabelStatus:        "Status",
		KeyLabelBudget:        "Budget",
		KeyLabelRevenue:       "Revenue",
		KeyLabelDirector:      "Director",
		KeyLabelPeriod:        "Period",
		KeyLabelSeasons:       "Seasons",
		KeyLabelEpisodes:      "Episodes",
		KeyLabelEpisodeLength: "Episode Length",
		KeyLabelNetwork:       "Network",
		KeyLabelCreatedBy:     "Created by",
		KeySectionPlot:        "Plot",
		KeySectionCast:        "Cast",
		KeySectionSeasons:      "Seasons",
		KeySectionLinks:       "Links",
		KeyNoMoviesFound:      "No movies found for: %s",
		KeyNoSeriesFound:      "No series found for: %s",
		KeySelectMovie:        "Select a movie",
		KeySelectSeries:       "Select a series",
		KeySelectAction:       "select",
		KeyCancelAction:       "cancel",
		KeyErrorNoAPIKey:     "Error: TMDB_API_KEY not set",
		KeyErrorSetAPIKey:    "Set your API key with:",
		KeyGetAPIKeyURL:      "Get your API key at:\n  https://www.themoviedb.org/settings/api",
	},
	"fr-FR": {
		KeyMovieIcon:          "🎬",
		KeySeriesIcon:         "📺",
		KeyLabelYear:          "Année",
		KeyLabelRuntime:       "Durée",
		KeyLabelGenre:         "Genre",
		KeyLabelRating:        "Note",
		KeyLabelStatus:        "Statut",
		KeyLabelBudget:        "Budget",
		KeyLabelRevenue:       "Recettes",
		KeyLabelDirector:      "Réalisateur",
		KeyLabelPeriod:        "Période",
		KeyLabelSeasons:       "Saisons",
		KeyLabelEpisodes:      "Épisodes",
		KeyLabelEpisodeLength: "Durée épisode",
		KeyLabelNetwork:       "Réseau",
		KeyLabelCreatedBy:     "Créé par",
		KeySectionPlot:        "Intrigue",
		KeySectionCast:        "Distribution",
		KeySectionSeasons:      "Saisons",
		KeySectionLinks:       "Liens",
		KeyNoMoviesFound:      "Aucun film trouvé pour: %s",
		KeyNoSeriesFound:      "Aucune série trouvée pour: %s",
		KeySelectMovie:        "Sélectionnez un film",
		KeySelectSeries:       "Sélectionnez une série",
		KeySelectAction:       "sélectionner",
		KeyCancelAction:       "annuler",
		KeyErrorNoAPIKey:     "Erreur: TMDB_API_KEY non défini",
		KeyErrorSetAPIKey:    "Définissez votre clé API avec:",
		KeyGetAPIKeyURL:      "Obtenez votre clé API sur:\n  https://www.themoviedb.org/settings/api",
	},
	"es-ES": {
		KeyMovieIcon:          "🎬",
		KeySeriesIcon:         "📺",
		KeyLabelYear:          "Año",
		KeyLabelRuntime:       "Duración",
		KeyLabelGenre:         "Género",
		KeyLabelRating:        "Puntuación",
		KeyLabelStatus:        "Estado",
		KeyLabelBudget:        "Presupuesto",
		KeyLabelRevenue:       "Recaudación",
		KeyLabelDirector:      "Director",
		KeyLabelPeriod:        "Periodo",
		KeyLabelSeasons:       "Temporadas",
		KeyLabelEpisodes:      "Episodios",
		KeyLabelEpisodeLength: "Duración episodio",
		KeyLabelNetwork:       "Red",
		KeyLabelCreatedBy:     "Creado por",
		KeySectionPlot:        "Trama",
		KeySectionCast:        "Reparto",
		KeySectionSeasons:      "Temporadas",
		KeySectionLinks:       "Enlaces",
		KeyNoMoviesFound:      "No se encontraron películas para: %s",
		KeyNoSeriesFound:      "No se encontraron series para: %s",
		KeySelectMovie:        "Selecciona una película",
		KeySelectSeries:       "Selecciona una serie",
		KeySelectAction:       "seleccionar",
		KeyCancelAction:       "cancelar",
		KeyErrorNoAPIKey:     "Error: TMDB_API_KEY no configurada",
		KeyErrorSetAPIKey:    "Configura tu clave API con:",
		KeyGetAPIKeyURL:      "Obtén tu clave API en:\n  https://www.themoviedb.org/settings/api",
	},
	"it-IT": {
		KeyMovieIcon:          "🎬",
		KeySeriesIcon:         "📺",
		KeyLabelYear:          "Anno",
		KeyLabelRuntime:       "Durata",
		KeyLabelGenre:         "Genere",
		KeyLabelRating:        "Valutazione",
		KeyLabelStatus:        "Stato",
		KeyLabelBudget:        "Budget",
		KeyLabelRevenue:       "Incasso",
		KeyLabelDirector:      "Regista",
		KeyLabelPeriod:        "Periodo",
		KeyLabelSeasons:       "Stagioni",
		KeyLabelEpisodes:      "Episodi",
		KeyLabelEpisodeLength: "Durata episodio",
		KeyLabelNetwork:       "Rete",
		KeyLabelCreatedBy:     "Creato da",
		KeySectionPlot:        "Trama",
		KeySectionCast:        "Cast",
		KeySectionSeasons:      "Stagioni",
		KeySectionLinks:       "Collegamenti",
		KeyNoMoviesFound:      "Nessun film trovato per: %s",
		KeyNoSeriesFound:      "Nessuna serie trovata per: %s",
		KeySelectMovie:        "Seleziona un film",
		KeySelectSeries:       "Seleziona una serie",
		KeySelectAction:       "seleziona",
		KeyCancelAction:       "annulla",
		KeyErrorNoAPIKey:     "Errore: TMDB_API_KEY non impostata",
		KeyErrorSetAPIKey:    "Imposta la tua chiave API con:",
		KeyGetAPIKeyURL:      "Ottieni la tua chiave API su:\n  https://www.themoviedb.org/settings/api",
	},
}

// SupportedLanguages gibt alle unterstützten Sprachen zurück
func SupportedLanguages() []string {
	return []string{"de-DE", "en-US", "fr-FR", "es-ES", "it-IT"}
}

// Translate gibt den übersetzten Text für den angegebenen Schlüssel und Sprache zurück
// Wenn die Sprache nicht unterstützt wird, wird auf Deutsch (de-DE) zurückgefallen
func Translate(key, language string) string {
	// Prüfen ob Sprache unterstützt wird
	langMap, ok := translations[language]
	if !ok {
		// Auf Deutsch zurückfallen
		langMap = translations["de-DE"]
	}

	// Übersetzung suchen
	if translation, ok := langMap[key]; ok {
		return translation
	}

	// Schlüssel nicht gefunden, auf Deutsch zurückfallen
	if deMap, ok := translations["de-DE"]; ok {
		if translation, ok := deMap[key]; ok {
			return translation
		}
	}

	// Als letztes den Schlüssel selbst zurückgeben
	return key
}

// Translatef gibt den übersetzten Text mit Formatierung zurück
func Translatef(key, language string, args ...interface{}) string {
	translation := Translate(key, language)
	return fmt.Sprintf(translation, args...)
}

// GetLanguageName gibt den lesbaren Namen einer Sprache zurück
func GetLanguageName(langCode string) string {
	names := map[string]string{
		"de-DE": "Deutsch",
		"en-US": "English",
		"fr-FR": "Français",
		"es-ES": "Español",
		"it-IT": "Italiano",
	}
	if name, ok := names[langCode]; ok {
		return name
	}
	return langCode
}
