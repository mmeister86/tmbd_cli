package ui

import (
	"fmt"
	"io"
	"strconv"

	"github.com/mmeister86/tmbd_cli/internal/i18n"
	"github.com/mmeister86/tmbd_cli/internal/tmdb"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Item-Styles
var (
	itemTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Bold(true)

	itemDescStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#888888"))

	selectedItemStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), false, false, false, true).
				BorderForeground(lipgloss.Color("#E50914")).
				Padding(0, 0, 0, 1)

	paginationStyle = list.DefaultStyles().PaginationStyle.
			PaddingLeft(4)

	helpStyle = list.DefaultStyles().HelpStyle.
			PaddingLeft(4).
			PaddingBottom(1)
)

// searchItem implementiert list.Item für die Suchergebnisse
type searchItem struct {
	id       int
	title    string
	year     string
	rating   float64
	overview string
}

func (i searchItem) FilterValue() string { return i.title }

// itemDelegate ist der benutzerdefinierte Delegate für die Liste
type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 3 }
func (d itemDelegate) Spacing() int                            { return 1 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(searchItem)
	if !ok {
		return
	}

	title := formatSearchItemTitle(item)

	// Beschreibung kürzen
	desc := item.overview
	desc = truncateText(desc, 60)

	var str string
	if index == m.Index() {
		// Ausgewähltes Item
		str = selectedItemStyle.Render(
			fmt.Sprintf("> %s\n  %s",
				itemTitleStyle.Render(title),
				itemDescStyle.Render(desc)))
	} else {
		// Normales Item
		str = fmt.Sprintf("  %s\n  %s",
			itemTitleStyle.Render(title),
			itemDescStyle.Render(desc))
	}

	fmt.Fprint(w, str)
}

func formatSearchItemTitle(item searchItem) string {
	title := item.title
	if item.year != "" {
		title = fmt.Sprintf("%s (%s)", title, item.year)
	}
	if item.rating > 0 {
		title = fmt.Sprintf("%s ★ %.1f", title, item.rating)
	}
	return title
}

// selectModel ist das Bubble Tea Model für die Auswahl
type selectModel struct {
	list     list.Model
	choice   int
	quitting bool
}

func (m selectModel) Init() tea.Cmd {
	return nil
}

func (m selectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height - 4)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			m.quitting = true
			m.choice = -1
			return m, tea.Quit
		case "enter":
			item, ok := m.list.SelectedItem().(searchItem)
			if ok {
				m.choice = item.id
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m selectModel) View() string {
	if m.quitting {
		return ""
	}
	return "\n" + m.list.View()
}

// SelectMovie zeigt eine interaktive Auswahl für Filme
func SelectMovie(results []tmdb.MovieSearchResult, language string) (int, error) {
	items := make([]list.Item, len(results))
	for i, r := range results {
		year := ""
		if len(r.ReleaseDate) >= 4 {
			year = r.ReleaseDate[:4]
		}
		items[i] = searchItem{
			id:       r.ID,
			title:    r.Title,
			year:     year,
			rating:   r.VoteAverage,
			overview: r.Overview,
		}
	}

	return runSelect(items, i18n.Translate(i18n.KeySelectMovie, language))
}

// SelectTV zeigt eine interaktive Auswahl für Serien
func SelectTV(results []tmdb.TVSearchResult, language string) (int, error) {
	items := make([]list.Item, len(results))
	for i, r := range results {
		year := ""
		if len(r.FirstAirDate) >= 4 {
			year = r.FirstAirDate[:4]
		}
		items[i] = searchItem{
			id:       r.ID,
			title:    r.Name,
			year:     year,
			rating:   r.VoteAverage,
			overview: r.Overview,
		}
	}

	return runSelect(items, i18n.Translate(i18n.KeySelectSeries, language))
}

// SelectPerson zeigt eine interaktive Auswahl für Personen
func SelectPerson(results []tmdb.PersonSearchResult, language string) (int, error) {
	items := make([]list.Item, len(results))
	for i, r := range results {
		// Für Personen: Name + bekannt für
		knownFor := ""
		if len(r.KnownFor) > 0 {
			work := r.KnownFor[0]
			if work.Title != "" {
				knownFor = work.Title
			} else if work.Name != "" {
				knownFor = work.Name
			}
		}
		items[i] = searchItem{
			id:       r.ID,
			title:    r.Name,
			year:     knownFor, // year-Feld für "bekannt für"
			rating:   r.Popularity,
			overview: "",
		}
	}

	return runSelect(items, i18n.Translate(i18n.KeySelectPerson, language))
}

func runSelect(items []list.Item, title string) (int, error) {
	delegate := itemDelegate{}

	l := list.New(items, delegate, 80, 20)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#E50914")).
		MarginLeft(2)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// Hilfstext anpassen
	l.KeyMap.Quit.SetKeys("q", "esc")
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "select"),
			),
		}
	}

	m := selectModel{list: l}
	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		return -1, err
	}

	result := finalModel.(selectModel)
	if result.choice == -1 {
		return -1, nil // Abgebrochen
	}

	return result.choice, nil
}

func runOptionSelect(items []list.Item, title string) (int, error) {
	return runSelect(items, title)
}

// SelectAction zeigt eine interaktive Auswahl für Drill-down-Aktionen.
func SelectAction(options []DrillDownOption, title string) (string, error) {
	items := make([]list.Item, len(options))
	for i, option := range options {
		id, err := strconv.Atoi(option.ID)
		if err != nil {
			id = i + 1
		}
		items[i] = searchItem{
			id:       id,
			title:    option.Title,
			year:     "",
			rating:   0,
			overview: option.Description,
		}
	}

	selected, err := runOptionSelect(items, title)
	if err != nil || selected == -1 {
		return "", err
	}
	for i, option := range options {
		id, convErr := strconv.Atoi(option.ID)
		if convErr != nil {
			id = i + 1
		}
		if id == selected {
			return option.ID, nil
		}
	}
	return "", nil
}

// SelectPersonOption zeigt eine interaktive Personenauswahl.
func SelectPersonOption(people []PersonOption, title string) (int, error) {
	items := make([]list.Item, len(people))
	for i, person := range people {
		items[i] = searchItem{
			id:       person.ID,
			title:    person.Name,
			year:     person.Description,
			rating:   0,
			overview: "",
		}
	}
	return runOptionSelect(items, title)
}

// SelectSeasonOption zeigt eine interaktive Staffelauswahl.
func SelectSeasonOption(seasons []DrillDownOption, title string) (int, error) {
	items := make([]list.Item, len(seasons))
	for i, season := range seasons {
		id, err := strconv.Atoi(season.ID)
		if err != nil {
			continue
		}
		items[i] = searchItem{
			id:       id,
			title:    season.Title,
			year:     "",
			rating:   0,
			overview: season.Description,
		}
	}
	return runOptionSelect(items, title)
}

// SelectLanguage zeigt eine interaktive Auswahl für Sprachen
func SelectLanguage() (string, error) {
	languages := i18n.SupportedLanguages()
	items := make([]list.Item, len(languages))

	for i, lang := range languages {
		langName := i18n.GetLanguageName(lang)
		items[i] = searchItem{
			id:       i,
			title:    langName,
			year:     lang,
			rating:   0,
			overview: "",
		}
	}

	delegate := itemDelegate{}

	l := list.New(items, delegate, 40, 10)
	l.Title = "🌍 Wähle eine Sprache"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#E50914")).
		MarginLeft(2)
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// Hilfstext anpassen
	l.KeyMap.Quit.SetKeys("q", "esc")
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "auswählen"),
			),
		}
	}

	m := selectModel{list: l}
	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		return "", err
	}

	result := finalModel.(selectModel)
	if result.choice == -1 {
		return "", nil // Abgebrochen
	}

	// Sprache anhand des Index zurückgeben
	if result.choice >= 0 && result.choice < len(languages) {
		return languages[result.choice], nil
	}

	return "", nil
}
