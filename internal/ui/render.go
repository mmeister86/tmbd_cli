package ui

import (
	"encoding/json"
	"fmt"
	"strings"

	"tmdb-cli/internal/tmdb"

	"github.com/charmbracelet/lipgloss"
)

// Farbschema laut PRD
var (
	primaryColor   = lipgloss.Color("#E50914")
	secondaryColor = lipgloss.Color("#FFD700")
	successColor   = lipgloss.Color("#00D26A")
	mutedColor     = lipgloss.Color("#888888")
	textColor      = lipgloss.Color("#FFFFFF")
)

// Styles
var (
	boxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(textColor)

	taglineStyle = lipgloss.NewStyle().
			Italic(true).
			Foreground(mutedColor)

	labelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(secondaryColor).
			Width(18)

	valueStyle = lipgloss.NewStyle().
			Foreground(textColor)

	ratingStyle = lipgloss.NewStyle().
			Foreground(successColor)

	sectionStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Bold(true)

	errorBoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF0000")).
			Padding(1, 2)

	infoBoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#3498DB")).
			Padding(1, 2)
)

// RenderMovieDetails rendert die Filmdetails
func RenderMovieDetails(movie *tmdb.MovieDetails, short bool) string {
	if short {
		return renderMovieShort(movie)
	}
	return renderMovieFull(movie)
}

func renderMovieFull(movie *tmdb.MovieDetails) string {
	var sb strings.Builder

	// Titel
	year := extractYear(movie.ReleaseDate)
	if movie.Title != movie.OriginalTitle {
		sb.WriteString(titleStyle.Render(fmt.Sprintf("🎬 %s (%s)", movie.Title, movie.OriginalTitle)))
	} else {
		sb.WriteString(titleStyle.Render(fmt.Sprintf("🎬 %s", movie.Title)))
	}
	sb.WriteString("\n")

	// Tagline
	if movie.Tagline != "" {
		sb.WriteString(taglineStyle.Render(fmt.Sprintf("\"%s\"", movie.Tagline)))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	// Details
	sb.WriteString(renderRow("Jahr", year))
	sb.WriteString(renderRow("Laufzeit", fmt.Sprintf("%d Min.", movie.Runtime)))
	sb.WriteString(renderRow("Genre", formatGenres(movie.Genres)))
	sb.WriteString(renderRow("Bewertung", formatRating(movie.VoteAverage, movie.VoteCount)))
	sb.WriteString(renderRow("Status", movie.Status))
	sb.WriteString("\n")

	// Regie
	directors := getDirectors(movie.Credits)
	if len(directors) > 0 {
		sb.WriteString(renderRow("Regie", strings.Join(directors, ", ")))
		sb.WriteString("\n")
	}

	// Budget & Revenue
	if movie.Budget > 0 {
		sb.WriteString(renderRow("Budget", formatMoney(movie.Budget)))
	}
	if movie.Revenue > 0 {
		sb.WriteString(renderRow("Einspielergebnis", formatMoney(movie.Revenue)))
	}
	if movie.Budget > 0 || movie.Revenue > 0 {
		sb.WriteString("\n")
	}

	// Handlung
	if movie.Overview != "" {
		sb.WriteString(renderSection("Handlung"))
		sb.WriteString(wrapText(movie.Overview, 60))
		sb.WriteString("\n\n")
	}

	// Besetzung
	if movie.Credits != nil && len(movie.Credits.Cast) > 0 {
		sb.WriteString(renderSection("Besetzung"))
		for i, cast := range movie.Credits.Cast {
			if i >= 5 {
				sb.WriteString(lipgloss.NewStyle().Foreground(mutedColor).Render("..."))
				sb.WriteString("\n")
				break
			}
			sb.WriteString(fmt.Sprintf("%s als %s\n", cast.Name, cast.Character))
		}
		sb.WriteString("\n")
	}

	// Links
	if movie.ImdbID != "" {
		sb.WriteString(renderSection("Links"))
		sb.WriteString(fmt.Sprintf("IMDb: https://www.imdb.com/title/%s\n", movie.ImdbID))
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

func renderMovieShort(movie *tmdb.MovieDetails) string {
	var sb strings.Builder

	year := extractYear(movie.ReleaseDate)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("🎬 %s (%s)", movie.Title, year)))
	sb.WriteString("\n")

	// Rating, Laufzeit, Genre in einer Zeile
	info := fmt.Sprintf("%s • %d Min. • %s",
		formatRatingCompact(movie.VoteAverage),
		movie.Runtime,
		formatGenres(movie.Genres))
	sb.WriteString(info)
	sb.WriteString("\n\n")

	// Handlung (gekürzt)
	if movie.Overview != "" {
		overview := truncateText(movie.Overview, 200)
		sb.WriteString(overview)
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

// RenderTVDetails rendert die Seriendetails
func RenderTVDetails(tv *tmdb.TVDetails, short bool) string {
	if short {
		return renderTVShort(tv)
	}
	return renderTVFull(tv)
}

func renderTVFull(tv *tmdb.TVDetails) string {
	var sb strings.Builder

	// Titel
	sb.WriteString(titleStyle.Render(fmt.Sprintf("📺 %s", tv.Name)))
	sb.WriteString("\n")

	// Tagline
	if tv.Tagline != "" {
		sb.WriteString(taglineStyle.Render(fmt.Sprintf("\"%s\"", tv.Tagline)))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	// Details
	firstYear := extractYear(tv.FirstAirDate)
	lastYear := extractYear(tv.LastAirDate)
	timeRange := firstYear
	if lastYear != "" && lastYear != firstYear {
		timeRange = fmt.Sprintf("%s - %s", firstYear, lastYear)
	}
	sb.WriteString(renderRow("Zeitraum", timeRange))
	sb.WriteString(renderRow("Staffeln", fmt.Sprintf("%d", tv.NumberOfSeasons)))
	sb.WriteString(renderRow("Episoden", fmt.Sprintf("%d", tv.NumberOfEpisodes)))
	if len(tv.EpisodeRunTime) > 0 {
		sb.WriteString(renderRow("Episodenlänge", fmt.Sprintf("~%d Min.", tv.EpisodeRunTime[0])))
	}
	sb.WriteString(renderRow("Genre", formatGenres(tv.Genres)))
	sb.WriteString(renderRow("Bewertung", formatRating(tv.VoteAverage, tv.VoteCount)))

	// Status mit Icon
	statusIcon := "○"
	if tv.InProduction {
		statusIcon = "●"
	}
	sb.WriteString(renderRow("Status", fmt.Sprintf("%s %s", statusIcon, tv.Status)))

	// Network
	if len(tv.Networks) > 0 {
		networks := make([]string, len(tv.Networks))
		for i, n := range tv.Networks {
			networks[i] = n.Name
		}
		sb.WriteString(renderRow("Sender", strings.Join(networks, ", ")))
	}
	sb.WriteString("\n")

	// Creator
	if len(tv.CreatedBy) > 0 {
		creators := make([]string, len(tv.CreatedBy))
		for i, c := range tv.CreatedBy {
			creators[i] = c.Name
		}
		sb.WriteString(renderRow("Erstellt von", strings.Join(creators, ", ")))
		sb.WriteString("\n")
	}

	// Handlung
	if tv.Overview != "" {
		sb.WriteString(renderSection("Handlung"))
		sb.WriteString(wrapText(tv.Overview, 60))
		sb.WriteString("\n\n")
	}

	// Besetzung
	if tv.Credits != nil && len(tv.Credits.Cast) > 0 {
		sb.WriteString(renderSection("Besetzung"))
		for i, cast := range tv.Credits.Cast {
			if i >= 5 {
				sb.WriteString(lipgloss.NewStyle().Foreground(mutedColor).Render("..."))
				sb.WriteString("\n")
				break
			}
			sb.WriteString(fmt.Sprintf("%s als %s\n", cast.Name, cast.Character))
		}
		sb.WriteString("\n")
	}

	// Staffeln
	if len(tv.Seasons) > 0 {
		sb.WriteString(renderSection("Staffeln"))
		for _, season := range tv.Seasons {
			if season.SeasonNumber == 0 {
				continue // Specials überspringen
			}
			year := extractYear(season.AirDate)
			sb.WriteString(fmt.Sprintf("Staffel %d: %d Episoden (%s)\n",
				season.SeasonNumber, season.EpisodeCount, year))
		}
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

func renderTVShort(tv *tmdb.TVDetails) string {
	var sb strings.Builder

	firstYear := extractYear(tv.FirstAirDate)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("📺 %s (%s)", tv.Name, firstYear)))
	sb.WriteString("\n")

	// Rating, Staffeln, Genre in einer Zeile
	info := fmt.Sprintf("%s • %d Staffeln • %s",
		formatRatingCompact(tv.VoteAverage),
		tv.NumberOfSeasons,
		formatGenres(tv.Genres))
	sb.WriteString(info)
	sb.WriteString("\n\n")

	// Handlung (gekürzt)
	if tv.Overview != "" {
		overview := truncateText(tv.Overview, 200)
		sb.WriteString(overview)
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

// RenderMovieJSON gibt die Filmdaten als JSON aus
func RenderMovieJSON(movie *tmdb.MovieDetails) (string, error) {
	output := tmdb.MovieJSONOutput{
		ID:            movie.ID,
		Title:         movie.Title,
		OriginalTitle: movie.OriginalTitle,
		Year:          extractYear(movie.ReleaseDate),
		Runtime:       movie.Runtime,
		Rating:        movie.VoteAverage,
		VoteCount:     movie.VoteCount,
		Budget:        movie.Budget,
		Revenue:       movie.Revenue,
		Genres:        extractGenreNames(movie.Genres),
		Directors:     getDirectors(movie.Credits),
		Cast:          extractCast(movie.Credits, 5),
		Overview:      movie.Overview,
		ImdbID:        movie.ImdbID,
		ImdbURL:       fmt.Sprintf("https://www.imdb.com/title/%s", movie.ImdbID),
		PosterURL:     formatPosterURL(movie.PosterPath),
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RenderTVJSON gibt die Seriendaten als JSON aus
func RenderTVJSON(tv *tmdb.TVDetails) (string, error) {
	networks := make([]string, len(tv.Networks))
	for i, n := range tv.Networks {
		networks[i] = n.Name
	}

	creators := make([]string, len(tv.CreatedBy))
	for i, c := range tv.CreatedBy {
		creators[i] = c.Name
	}

	output := tmdb.TVJSONOutput{
		ID:           tv.ID,
		Name:         tv.Name,
		OriginalName: tv.OriginalName,
		FirstAirDate: tv.FirstAirDate,
		LastAirDate:  tv.LastAirDate,
		Seasons:      tv.NumberOfSeasons,
		Episodes:     tv.NumberOfEpisodes,
		Rating:       tv.VoteAverage,
		VoteCount:    tv.VoteCount,
		Status:       tv.Status,
		Genres:       extractGenreNames(tv.Genres),
		Networks:     networks,
		Creators:     creators,
		Cast:         extractCast(tv.Credits, 5),
		Overview:     tv.Overview,
		PosterURL:    formatPosterURL(tv.PosterPath),
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// RenderError rendert eine Fehlermeldung
func RenderError(title, message string, hints []string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("❌ %s\n\n", title))
	sb.WriteString(message)

	if len(hints) > 0 {
		sb.WriteString("\n\n")
		for _, hint := range hints {
			sb.WriteString(hint + "\n")
		}
	}

	return errorBoxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

// RenderInfo rendert eine Info-Meldung
func RenderInfo(message string) string {
	return infoBoxStyle.Render(fmt.Sprintf("ℹ️  %s", message))
}

// Hilfsfunktionen

func renderRow(label, value string) string {
	return labelStyle.Render(label) + valueStyle.Render(value) + "\n"
}

func renderSection(title string) string {
	line := strings.Repeat("─", 50)
	return sectionStyle.Render(fmt.Sprintf("─── %s %s", title, line[:50-len(title)-5])) + "\n"
}

func extractYear(date string) string {
	if len(date) >= 4 {
		return date[:4]
	}
	return ""
}

func formatGenres(genres []tmdb.Genre) string {
	names := make([]string, len(genres))
	for i, g := range genres {
		names[i] = g.Name
	}
	return strings.Join(names, ", ")
}

func extractGenreNames(genres []tmdb.Genre) []string {
	names := make([]string, len(genres))
	for i, g := range genres {
		names[i] = g.Name
	}
	return names
}

func formatRating(rating float64, voteCount int) string {
	stars := renderStars(rating)
	return ratingStyle.Render(fmt.Sprintf("%.1f/10 %s (%s Bewertungen)",
		rating, stars, formatNumber(voteCount)))
}

func formatRatingCompact(rating float64) string {
	stars := renderStars(rating)
	return ratingStyle.Render(fmt.Sprintf("%.1f/10 %s", rating, stars))
}

func renderStars(rating float64) string {
	fullStars := int(rating / 2)
	emptyStars := 5 - fullStars
	return strings.Repeat("★", fullStars) + strings.Repeat("☆", emptyStars)
}

func formatMoney(amount int64) string {
	if amount >= 1000000000 {
		return fmt.Sprintf("$%.1f Mrd.", float64(amount)/1000000000)
	}
	if amount >= 1000000 {
		return fmt.Sprintf("$%.1f Mio.", float64(amount)/1000000)
	}
	return fmt.Sprintf("$%s", formatNumber(int(amount)))
}

func formatNumber(n int) string {
	str := fmt.Sprintf("%d", n)
	result := ""
	for i, c := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += "."
		}
		result += string(c)
	}
	return result
}

func getDirectors(credits *tmdb.Credits) []string {
	if credits == nil {
		return nil
	}
	var directors []string
	for _, crew := range credits.Crew {
		if crew.Job == "Director" {
			directors = append(directors, crew.Name)
		}
	}
	return directors
}

func extractCast(credits *tmdb.Credits, limit int) []tmdb.CastOutput {
	if credits == nil || len(credits.Cast) == 0 {
		return nil
	}
	count := limit
	if len(credits.Cast) < limit {
		count = len(credits.Cast)
	}
	cast := make([]tmdb.CastOutput, count)
	for i := 0; i < count; i++ {
		cast[i] = tmdb.CastOutput{
			Name:      credits.Cast[i].Name,
			Character: credits.Cast[i].Character,
		}
	}
	return cast
}

func formatPosterURL(path string) string {
	if path == "" {
		return ""
	}
	return fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", path)
}

func wrapText(text string, width int) string {
	words := strings.Fields(text)
	var lines []string
	var currentLine string

	for _, word := range words {
		if len(currentLine)+len(word)+1 > width {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			if currentLine != "" {
				currentLine += " "
			}
			currentLine += word
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}

	return strings.Join(lines, "\n")
}

func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	return text[:maxLen-3] + "..."
}
