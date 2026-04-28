package ui

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/mmeister86/tmbd_cli/internal/i18n"
	"github.com/mmeister86/tmbd_cli/internal/tmdb"

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
func RenderMovieDetails(movie *tmdb.MovieDetails, short bool, language string) string {
	if short {
		return renderMovieShort(movie, language)
	}
	return renderMovieFull(movie, language)
}

func renderMovieFull(movie *tmdb.MovieDetails, language string) string {
	var sb strings.Builder

	// Titel
	movieIcon := i18n.Translate(i18n.KeyMovieIcon, language)
	year := extractYear(movie.ReleaseDate)
	if movie.Title != movie.OriginalTitle {
		sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s (%s)", movieIcon, movie.Title, movie.OriginalTitle)))
	} else {
		sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s", movieIcon, movie.Title)))
	}
	sb.WriteString("\n")

	// Tagline
	if movie.Tagline != "" {
		sb.WriteString(taglineStyle.Render(fmt.Sprintf("\"%s\"", movie.Tagline)))
		sb.WriteString("\n")
	}
	sb.WriteString("\n")

	// Details
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelYear, language), year))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelRuntime, language), fmt.Sprintf("%d Min.", movie.Runtime)))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelGenre, language), formatGenres(movie.Genres)))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelRating, language), formatRating(movie.VoteAverage, movie.VoteCount)))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelStatus, language), movie.Status))
	sb.WriteString("\n")

	// Regie
	directors := getDirectors(movie.Credits)
	if len(directors) > 0 {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelDirector, language), strings.Join(directors, ", ")))
		sb.WriteString("\n")
	}

	// Budget & Revenue
	if movie.Budget > 0 {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelBudget, language), formatMoney(movie.Budget)))
	}
	if movie.Revenue > 0 {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelRevenue, language), formatMoney(movie.Revenue)))
	}
	if movie.Budget > 0 || movie.Revenue > 0 {
		sb.WriteString("\n")
	}

	// Handlung
	if movie.Overview != "" {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionPlot, language)))
		sb.WriteString(wrapText(movie.Overview, 60))
		sb.WriteString("\n\n")
	}

	// Besetzung
	if movie.Credits != nil && len(movie.Credits.Cast) > 0 {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionCast, language)))
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
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionLinks, language)))
		sb.WriteString(fmt.Sprintf("IMDb: https://www.imdb.com/title/%s\n", movie.ImdbID))
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

func renderMovieShort(movie *tmdb.MovieDetails, language string) string {
	var sb strings.Builder

	movieIcon := i18n.Translate(i18n.KeyMovieIcon, language)
	year := extractYear(movie.ReleaseDate)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s (%s)", movieIcon, movie.Title, year)))
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
func RenderTVDetails(tv *tmdb.TVDetails, short bool, language string) string {
	if short {
		return renderTVShort(tv, language)
	}
	return renderTVFull(tv, language)
}

func renderTVFull(tv *tmdb.TVDetails, language string) string {
	var sb strings.Builder

	// Titel
	tvIcon := i18n.Translate(i18n.KeySeriesIcon, language)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s", tvIcon, tv.Name)))
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
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelPeriod, language), timeRange))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelSeasons, language), fmt.Sprintf("%d", tv.NumberOfSeasons)))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelEpisodes, language), fmt.Sprintf("%d", tv.NumberOfEpisodes)))
	if len(tv.EpisodeRunTime) > 0 {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelEpisodeLength, language), fmt.Sprintf("~%d Min.", tv.EpisodeRunTime[0])))
	}
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelGenre, language), formatGenres(tv.Genres)))
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelRating, language), formatRating(tv.VoteAverage, tv.VoteCount)))

	// Status mit Icon
	statusIcon := "○"
	if tv.InProduction {
		statusIcon = "●"
	}
	sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelStatus, language), fmt.Sprintf("%s %s", statusIcon, tv.Status)))

	// Network
	if len(tv.Networks) > 0 {
		networks := make([]string, len(tv.Networks))
		for i, n := range tv.Networks {
			networks[i] = n.Name
		}
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelNetwork, language), strings.Join(networks, ", ")))
	}
	sb.WriteString("\n")

	// Creator
	if len(tv.CreatedBy) > 0 {
		creators := make([]string, len(tv.CreatedBy))
		for i, c := range tv.CreatedBy {
			creators[i] = c.Name
		}
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelCreatedBy, language), strings.Join(creators, ", ")))
		sb.WriteString("\n")
	}

	// Handlung
	if tv.Overview != "" {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionPlot, language)))
		sb.WriteString(wrapText(tv.Overview, 60))
		sb.WriteString("\n\n")
	}

	// Besetzung
	if tv.Credits != nil && len(tv.Credits.Cast) > 0 {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionCast, language)))
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
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionSeasons, language)))
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

func renderTVShort(tv *tmdb.TVDetails, language string) string {
	var sb strings.Builder

	tvIcon := i18n.Translate(i18n.KeySeriesIcon, language)
	firstYear := extractYear(tv.FirstAirDate)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s (%s)", tvIcon, tv.Name, firstYear)))
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
func RenderError(title, message string, hints []string, language string) string {
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
			if currentLine != "" {
				lines = append(lines, currentLine)
			}
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
	if maxLen <= 3 {
		runes := []rune(text)
		if len(runes) <= maxLen {
			return text
		}
		return string(runes[:maxLen])
	}

	runes := []rune(text)
	if len(runes) <= maxLen {
		return text
	}
	return string(runes[:maxLen-3]) + "..."
}

// RenderPersonDetails rendert die Personendetails
func RenderPersonDetails(person *tmdb.PersonDetails, short bool, language string) string {
	if short {
		return renderPersonShort(person, language)
	}
	return renderPersonFull(person, language)
}

func renderPersonFull(person *tmdb.PersonDetails, language string) string {
	var sb strings.Builder

	// Titel
	personIcon := i18n.Translate(i18n.KeyPersonIcon, language)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s", personIcon, person.Name)))
	sb.WriteString("\n\n")

	// Geburtsdatum
	if person.Birthday != "" {
		birthday := formatBirthday(person.Birthday, language)
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelBirthday, language), birthday))
	}

	// Sterbedatum
	if person.Deathday != "" {
		deathday := formatDeathday(person.Deathday, language)
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelDeathday, language), deathday))
	}

	// Geburtsort
	if person.PlaceOfBirth != "" {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelPlaceOfBirth, language), person.PlaceOfBirth))
	}

	// Beruf/Department
	if person.KnownForDepartment != "" {
		sb.WriteString(renderRow(i18n.Translate(i18n.KeyLabelDepartment, language), person.KnownForDepartment))
		sb.WriteString("\n")
	} else if len(person.Birthday) > 0 || len(person.PlaceOfBirth) > 0 {
		sb.WriteString("\n")
	}

	// Biografie
	if person.Biography != "" {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionPlot, language)))
		sb.WriteString(wrapText(person.Biography, 60))
		sb.WriteString("\n\n")
	}

	// Bekannte Rollen
	if person.CombinedCredits != nil {
		knownFor := getKnownForCredits(person.CombinedCredits)
		if len(knownFor) > 0 {
			sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionKnownFor, language)))
			for i, work := range knownFor {
				if i >= 5 {
					sb.WriteString(lipgloss.NewStyle().Foreground(mutedColor).Render("..."))
					sb.WriteString("\n")
					break
				}
				sb.WriteString(fmt.Sprintf("%s %s\n", getWorkIcon(work.MediaType), formatKnownForWork(work)))
			}
		}
	}

	// IMDb Link
	if person.IMDBID != "" {
		sb.WriteString(renderSection(i18n.Translate(i18n.KeySectionLinks, language)))
		sb.WriteString(fmt.Sprintf("IMDb: https://www.imdb.com/name/%s\n", person.IMDBID))
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

func renderPersonShort(person *tmdb.PersonDetails, language string) string {
	var sb strings.Builder

	personIcon := i18n.Translate(i18n.KeyPersonIcon, language)
	sb.WriteString(titleStyle.Render(fmt.Sprintf("%s %s", personIcon, person.Name)))
	sb.WriteString("\n")

	// Geburtsdatum und Beruf in einer Zeile
	var info string
	if person.Birthday != "" && person.KnownForDepartment != "" {
		birthday := formatBirthday(person.Birthday, language)
		info = fmt.Sprintf("%s • %s", birthday, person.KnownForDepartment)
		sb.WriteString(info)
		sb.WriteString("\n\n")
	} else if person.Birthday != "" {
		birthday := formatBirthday(person.Birthday, language)
		sb.WriteString(birthday)
		sb.WriteString("\n\n")
	}

	// Biografie (gekürzt)
	if person.Biography != "" {
		overview := truncateText(person.Biography, 200)
		sb.WriteString(overview)
	}

	return boxStyle.Render(strings.TrimRight(sb.String(), "\n"))
}

// RenderPersonJSON gibt die Personendaten als JSON aus
func RenderPersonJSON(person *tmdb.PersonDetails) (string, error) {
	var knownForWorks []tmdb.KnownForWorkOutput

	if person.CombinedCredits != nil {
		knownFor := getKnownForCredits(person.CombinedCredits)
		count := len(knownFor)
		if count > 5 {
			count = 5
		}
		knownForWorks = make([]tmdb.KnownForWorkOutput, count)
		for i := 0; i < count; i++ {
			knownForWorks[i] = formatKnownForWorkJSON(knownFor[i])
		}
	}

	output := tmdb.PersonJSONOutput{
		ID:            person.ID,
		Name:          person.Name,
		Birthday:      person.Birthday,
		Deathday:      person.Deathday,
		PlaceOfBirth:  person.PlaceOfBirth,
		KnownFor:      person.KnownForDepartment,
		Biography:     person.Biography,
		KnownForWorks: knownForWorks,
		IMDBID:        person.IMDBID,
		IMDBURL:       fmt.Sprintf("https://www.imdb.com/name/%s", person.IMDBID),
		ProfileURL:    formatPosterURL(person.ProfilePath),
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Hilfsfunktionen für Personen
func getKnownForCredits(credits *tmdb.CombinedCredits) []tmdb.CombinedCast {
	if credits == nil {
		return nil
	}

	// Alle Cast-Einträge sammeln (Priorität: Cast vor Crew)
	var allWorks []tmdb.CombinedCast
	allWorks = append(allWorks, credits.Cast...)

	// Nach Popularität sortieren und duplikate entfernen
	if len(allWorks) == 0 {
		return nil
	}

	// Top 10 nach Popularität
	sort.Slice(allWorks, func(i, j int) bool {
		return allWorks[i].Popularity > allWorks[j].Popularity
	})

	// Begrenzen auf 10, Popularitäts-Reihenfolge beibehalten
	result := make([]tmdb.CombinedCast, 0, 10)
	seenIDs := make(map[int]bool)
	for _, work := range allWorks {
		if len(result) >= 10 {
			break
		}
		if !seenIDs[work.ID] {
			seenIDs[work.ID] = true
			result = append(result, work)
		}
	}

	return result
}

func formatKnownForWork(work tmdb.CombinedCast) string {
	title := work.Title
	if title == "" {
		title = work.Name
	}
	year := extractYear(work.ReleaseDate)
	if year == "" {
		year = extractYear(work.FirstAirDate)
	}

	character := work.Character
	if character == "" {
		character = i18n.Translate(i18n.KeyLabelDepartment, "de-DE")
	}

	return fmt.Sprintf("%s in %s (%s)", character, title, year)
}

func formatKnownForWorkJSON(work tmdb.CombinedCast) tmdb.KnownForWorkOutput {
	title := work.Title
	if title == "" {
		title = work.Name
	}
	year := extractYear(work.ReleaseDate)
	if year == "" {
		year = extractYear(work.FirstAirDate)
	}

	return tmdb.KnownForWorkOutput{
		Title:         title,
		OriginalTitle: work.OriginalTitle,
		Year:          year,
		MediaType:     work.MediaType,
	}
}

func formatBirthday(date string, language string) string {
	if date == "" {
		return ""
	}

	// TMDB Format: YYYY-MM-DD
	parts := strings.Split(date, "-")
	if len(parts) != 3 {
		return date
	}

	year, month, day := parts[0], parts[1], parts[2]

	// Je nach Sprache formatieren
	switch language {
	case "de-DE":
		// 9. Juli 1956
		monthNames := map[string]string{
			"01": "Januar", "02": "Februar", "03": "März",
			"04": "April", "05": "Mai", "06": "Juni",
			"07": "Juli", "08": "August", "09": "September",
			"10": "Oktober", "11": "November", "12": "Dezember",
		}
		return fmt.Sprintf("%s. %s %s", day, monthNames[month], year)
	case "en-US":
		// July 9, 1956
		monthNames := map[string]string{
			"01": "January", "02": "February", "03": "March",
			"04": "April", "05": "May", "06": "June",
			"07": "July", "08": "August", "09": "September",
			"10": "October", "11": "November", "12": "December",
		}
		dayNum := day
		if dayNum[0] == '0' {
			dayNum = dayNum[1:]
		}
		return fmt.Sprintf("%s %s, %s", monthNames[month], dayNum, year)
	case "fr-FR":
		// 9 juillet 1956
		monthNames := map[string]string{
			"01": "janvier", "02": "février", "03": "mars",
			"04": "avril", "05": "mai", "06": "juin",
			"07": "juillet", "08": "août", "09": "septembre",
			"10": "octobre", "11": "novembre", "12": "décembre",
		}
		return fmt.Sprintf("%s %s %s", day, monthNames[month], year)
	case "es-ES":
		// 9 de julio de 1956
		monthNames := map[string]string{
			"01": "enero", "02": "febrero", "03": "marzo",
			"04": "abril", "05": "mayo", "06": "junio",
			"07": "julio", "08": "agosto", "09": "septiembre",
			"10": "octubre", "11": "noviembre", "12": "diciembre",
		}
		return fmt.Sprintf("%s de %s de %s", day, monthNames[month], year)
	case "it-IT":
		// 9 luglio 1956
		monthNames := map[string]string{
			"01": "gennaio", "02": "febbraio", "03": "marzo",
			"04": "aprile", "05": "maggio", "06": "giugno",
			"07": "luglio", "08": "agosto", "09": "settembre",
			"10": "ottobre", "11": "novembre", "12": "dicembre",
		}
		return fmt.Sprintf("%s %s %s", day, monthNames[month], year)
	default:
		return date
	}
}

func formatDeathday(date string, language string) string {
	return formatBirthday(date, language)
}

func getWorkIcon(mediaType string) string {
	switch mediaType {
	case "movie":
		return "🎬"
	case "tv":
		return "📺"
	default:
		return "•"
	}
}
