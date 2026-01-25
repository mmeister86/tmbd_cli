# Product Requirements Document (PRD)
# TMDB CLI Tool

**Version:** 1.0  
**Datum:** 25. Januar 2026  
**Status:** Draft  

---

## 1. Übersicht

### 1.1 Produktbeschreibung
Ein plattformübergreifendes Command-Line Tool in Go, das Informationen über Filme und Serien von The Movie Database (TMDB) abruft und ansprechend im Terminal darstellt.

### 1.2 Ziele
- Schneller Zugriff auf Film-/Serieninformationen direkt aus dem Terminal
- Ansprechende, farbige Darstellung der Ergebnisse
- Einfache Bedienung mit intuitiven Subcommands
- Plattformübergreifende Kompatibilität (Windows, macOS, Linux)

### 1.3 Nicht-Ziele
- Keine GUI-Anwendung
- Keine Verwaltung von Watchlists oder Favoriten (v1)
- Keine Offline-Funktionalität
- Keine Benutzerauthentifizierung bei TMDB

---

## 2. Technische Spezifikation

### 2.1 Technologie-Stack
| Komponente | Technologie | Begründung |
|------------|-------------|------------|
| Sprache | Go 1.22+ | Plattformübergreifende Kompilierung, Single Binary |
| CLI Framework | [Cobra](https://github.com/spf13/cobra) | De-facto Standard für Go CLIs |
| Terminal UI | [Bubble Tea](https://github.com/charmbracelet/bubbletea) | Interaktive Auswahllisten |
| Styling | [Lipgloss](https://github.com/charmbracelet/lipgloss) | Moderne Terminal-Styles |
| API | TMDB API v3 | Kostenlos, umfangreich, gut dokumentiert |

### 2.2 Projektstruktur
```
tmdb-cli/
├── cmd/
│   ├── root.go          # Haupt-Command, globale Flags
│   ├── movie.go         # Movie Subcommand
│   └── series.go        # Series Subcommand
├── internal/
│   ├── tmdb/
│   │   ├── client.go    # HTTP Client, API Requests
│   │   └── types.go     # API Response Types
│   └── ui/
│       ├── render.go    # Ausgabe-Formatierung
│       └── select.go    # Interaktive Auswahl
├── main.go
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### 2.3 Konfiguration
| Variable | Beschreibung | Pflicht |
|----------|--------------|---------|
| `TMDB_API_KEY` | TMDB API Key (v3 auth) | Ja |
| `TMDB_LANGUAGE` | Sprache für Ergebnisse (default: `de-DE`) | Nein |

---

## 3. Funktionale Anforderungen

### 3.1 Commands

#### 3.1.1 Root Command
```bash
tmdb [flags]
tmdb [command]
```

**Beschreibung:** Zeigt Hilfe und verfügbare Commands an.

**Globale Flags:**
| Flag | Kurzform | Beschreibung |
|------|----------|--------------|
| `--help` | `-h` | Hilfe anzeigen |
| `--version` | `-v` | Version anzeigen |
| `--json` | | Ausgabe als JSON |
| `--short` | `-s` | Kompakte Ausgabe |
| `--language` | `-l` | Sprache überschreiben (z.B. `en-US`) |

---

#### 3.1.1.1 Language Command
```bash
tmdb language
```

**Beschreibung:** Zeigt die aktuelle Sprache an oder ermöglicht die interaktive Auswahl einer neuen Sprache.

**Unterstützte Sprachen:**
- Deutsch (de-DE)
- Englisch (en-US)
- Französisch (fr-FR)
- Spanisch (es-ES)
- Italienisch (it-IT)

**Beispiel:**
```bash
$ tmdb language
Aktuelle Sprache: Deutsch (de-DE)

🌍 Wähle eine Sprache

  > Deutsch
    English
    Français
    Español
    Italiano

✓ Sprache auf English (en-US) geändert
```

---

#### 3.1.2 Movie Command
```bash
tmdb movie <suchbegriff> [flags]
```

**Aliases:** `m`, `film`

**Beschreibung:** Sucht nach einem Film und zeigt detaillierte Informationen an.

**Beispiele:**
```bash
tmdb movie "The Matrix"
tmdb movie "Inception" --short
tmdb movie "Pulp Fiction" --json
tmdb m "Fight Club"
```

**Verhalten:**
1. Suche nach dem Suchbegriff via TMDB Search API
2. Bei mehreren Ergebnissen: Interaktive Auswahlliste anzeigen
3. Bei einem Ergebnis: Direkt Details laden
4. Bei keinem Ergebnis: Fehlermeldung anzeigen

**Ausgabe (Standard):**
```
╭──────────────────────────────────────────────────────────────╮
│ 🎬 The Matrix (The Matrix)                                   │
│ "Unfortunately, no one can be told what the Matrix is."      │
│                                                              │
│ Jahr           1999                                          │
│ Laufzeit       136 Min.                                      │
│ Genre          Action, Science Fiction                       │
│ Bewertung      8.2/10 ★★★★☆ (25,432 Bewertungen)            │
│ Status         Released                                      │
│                                                              │
│ Regie          Lana Wachowski, Lilly Wachowski              │
│                                                              │
│ Budget         $63.0 Mio.                                    │
│ Einspielergebnis $463.5 Mio.                                 │
│                                                              │
│ ─── Handlung ───────────────────────────────────────────     │
│ Thomas A. Anderson ist ein Mann, der zwei Leben lebt...      │
│                                                              │
│ ─── Besetzung ──────────────────────────────────────────     │
│ Keanu Reeves als Neo                                         │
│ Laurence Fishburne als Morpheus                              │
│ Carrie-Anne Moss als Trinity                                 │
│ Hugo Weaving als Agent Smith                                 │
│ ...                                                          │
│                                                              │
│ ─── Links ──────────────────────────────────────────────     │
│ IMDb: https://www.imdb.com/title/tt0133093                   │
╰──────────────────────────────────────────────────────────────╯
```

**Ausgabe (--short):**
```
╭──────────────────────────────────────────────────────────────╮
│ 🎬 The Matrix (1999)                                         │
│ 8.2/10 ★★★★☆ • 136 Min. • Action, Science Fiction           │
│                                                              │
│ Thomas A. Anderson ist ein Mann, der zwei Leben lebt...      │
╰──────────────────────────────────────────────────────────────╯
```

**Ausgabe (--json):**
```json
{
  "id": 603,
  "title": "The Matrix",
  "original_title": "The Matrix",
  "year": "1999",
  "runtime": 136,
  "rating": 8.2,
  "vote_count": 25432,
  "budget": 63000000,
  "revenue": 463517383,
  "genres": ["Action", "Science Fiction"],
  "directors": ["Lana Wachowski", "Lilly Wachowski"],
  "cast": [
    {"name": "Keanu Reeves", "character": "Neo"},
    {"name": "Laurence Fishburne", "character": "Morpheus"}
  ],
  "overview": "Thomas A. Anderson ist ein Mann...",
  "imdb_id": "tt0133093",
  "imdb_url": "https://www.imdb.com/title/tt0133093",
  "poster_url": "https://image.tmdb.org/t/p/w500/..."
}
```

---

#### 3.1.3 Series Command
```bash
tmdb series <suchbegriff> [flags]
```

**Aliases:** `s`, `tv`, `show`

**Beschreibung:** Sucht nach einer Serie und zeigt detaillierte Informationen an.

**Beispiele:**
```bash
tmdb series "Breaking Bad"
tmdb series "Stranger Things" --short
tmdb tv "The Office" --json
tmdb s "Dark"
```

**Verhalten:** Analog zu Movie Command.

**Ausgabe (Standard):**
```
╭──────────────────────────────────────────────────────────────╮
│ 📺 Breaking Bad                                              │
│ "Remember my name."                                          │
│                                                              │
│ Zeitraum       2008 - 2013                                   │
│ Staffeln       5                                             │
│ Episoden       62                                            │
│ Episodenlänge  ~47 Min.                                      │
│ Genre          Drama, Crime                                  │
│ Bewertung      9.5/10 ★★★★★ (13,245 Bewertungen)            │
│ Status         ○ Ended                                       │
│ Sender         AMC                                           │
│                                                              │
│ Erstellt von   Vince Gilligan                                │
│                                                              │
│ ─── Handlung ───────────────────────────────────────────     │
│ Der Chemielehrer Walter White erfährt, dass er an...         │
│                                                              │
│ ─── Besetzung ──────────────────────────────────────────     │
│ Bryan Cranston als Walter White                              │
│ Aaron Paul als Jesse Pinkman                                 │
│ Anna Gunn als Skyler White                                   │
│ ...                                                          │
│                                                              │
│ ─── Staffeln ───────────────────────────────────────────     │
│ Staffel 1: 7 Episoden (2008)                                 │
│ Staffel 2: 13 Episoden (2009)                                │
│ Staffel 3: 13 Episoden (2010)                                │
│ Staffel 4: 13 Episoden (2011)                                │
│ Staffel 5: 16 Episoden (2012)                                │
╰──────────────────────────────────────────────────────────────╯
```

**Serien-spezifische Felder:**
- Creator(s)
- Anzahl Staffeln/Episoden
- Episodenlänge
- Status (laufend/beendet) mit visueller Unterscheidung
- Sender/Network
- Staffelübersicht

---

### 3.2 Interaktive Auswahl

Bei mehreren Suchergebnissen wird eine interaktive Liste angezeigt:

```
🎬 Wähle einen Film

  > The Matrix (1999) ★ 8.2
    Thomas A. Anderson ist ein Mann, der zwei Leben...
    
    The Matrix Reloaded (2003) ★ 7.2
    Neo und seine Verbündeten kämpfen gegen die...
    
    The Matrix Revolutions (2003) ★ 6.7
    Die Schlacht um Zion beginnt...
    
    The Matrix Resurrections (2021) ★ 5.7
    Neo lebt wieder ein normales Leben...

↑/↓: Navigieren • Enter: Auswählen • /: Filtern • q: Abbrechen
```

**Bedienung:**
| Taste | Aktion |
|-------|--------|
| `↑`/`k` | Nach oben |
| `↓`/`j` | Nach unten |
| `Enter` | Auswählen |
| `/` | Filter/Suche |
| `q`/`Esc` | Abbrechen |

---

### 3.3 Fehlerbehandlung

| Szenario | Verhalten |
|----------|-----------|
| Kein API Key | Fehlermeldung mit Anleitung zum Setzen |
| API nicht erreichbar | Timeout-Fehler mit Retry-Hinweis |
| Keine Ergebnisse | Info-Meldung "Keine Ergebnisse für: X" |
| Ungültige Eingabe | Hilfetext für Command anzeigen |

**Fehlerausgabe-Format:**
```
❌ Fehler: TMDB_API_KEY nicht gesetzt

Setze deinen API Key mit:
  export TMDB_API_KEY='dein-api-key'

API Key erhältst du unter:
  https://www.themoviedb.org/settings/api
```

---

## 4. TMDB API Integration

### 4.1 Verwendete Endpoints

| Endpoint | Verwendung |
|----------|------------|
| `GET /search/movie` | Film-Suche |
| `GET /search/tv` | Serien-Suche |
| `GET /movie/{id}` | Film-Details |
| `GET /tv/{id}` | Serien-Details |

### 4.2 Query Parameter

**Für alle Requests:**
- `api_key`: TMDB API Key
- `language`: Sprache (default: `de-DE`)

**Für Details:**
- `append_to_response`: `credits` (Cast & Crew in einem Request)

### 4.3 Beispiel-Requests

**Film-Suche:**
```
GET https://api.themoviedb.org/3/search/movie?api_key=XXX&language=de-DE&query=Matrix
```

**Film-Details:**
```
GET https://api.themoviedb.org/3/movie/603?api_key=XXX&language=de-DE&append_to_response=credits
```

**Serien-Suche:**
```
GET https://api.themoviedb.org/3/search/tv?api_key=XXX&language=de-DE&query=Breaking+Bad
```

**Serien-Details:**
```
GET https://api.themoviedb.org/3/tv/1396?api_key=XXX&language=de-DE&append_to_response=credits
```

---

## 5. Datenmodelle

### 5.1 Movie Response (vereinfacht)

```go
type MovieDetails struct {
    ID            int       `json:"id"`
    Title         string    `json:"title"`
    OriginalTitle string    `json:"original_title"`
    Tagline       string    `json:"tagline"`
    Overview      string    `json:"overview"`
    ReleaseDate   string    `json:"release_date"`
    Runtime       int       `json:"runtime"`
    Budget        int64     `json:"budget"`
    Revenue       int64     `json:"revenue"`
    VoteAverage   float64   `json:"vote_average"`
    VoteCount     int       `json:"vote_count"`
    Genres        []Genre   `json:"genres"`
    Status        string    `json:"status"`
    Homepage      string    `json:"homepage"`
    ImdbID        string    `json:"imdb_id"`
    PosterPath    string    `json:"poster_path"`
    Credits       *Credits  `json:"credits,omitempty"`
}

type Genre struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Credits struct {
    Cast []CastMember `json:"cast"`
    Crew []CrewMember `json:"crew"`
}

type CastMember struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Character string `json:"character"`
    Order     int    `json:"order"`
}

type CrewMember struct {
    ID         int    `json:"id"`
    Name       string `json:"name"`
    Job        string `json:"job"`
    Department string `json:"department"`
}
```

### 5.2 TV Response (vereinfacht)

```go
type TVDetails struct {
    ID               int       `json:"id"`
    Name             string    `json:"name"`
    OriginalName     string    `json:"original_name"`
    Tagline          string    `json:"tagline"`
    Overview         string    `json:"overview"`
    FirstAirDate     string    `json:"first_air_date"`
    LastAirDate      string    `json:"last_air_date"`
    Status           string    `json:"status"`
    NumberOfSeasons  int       `json:"number_of_seasons"`
    NumberOfEpisodes int       `json:"number_of_episodes"`
    EpisodeRunTime   []int     `json:"episode_run_time"`
    VoteAverage      float64   `json:"vote_average"`
    VoteCount        int       `json:"vote_count"`
    Genres           []Genre   `json:"genres"`
    Networks         []Network `json:"networks"`
    CreatedBy        []Creator `json:"created_by"`
    Homepage         string    `json:"homepage"`
    InProduction     bool      `json:"in_production"`
    Seasons          []Season  `json:"seasons"`
    Credits          *Credits  `json:"credits,omitempty"`
}

type Network struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Creator struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type Season struct {
    ID           int    `json:"id"`
    Name         string `json:"name"`
    SeasonNumber int    `json:"season_number"`
    EpisodeCount int    `json:"episode_count"`
    AirDate      string `json:"air_date"`
}
```

---

## 6. UI/UX Design

### 6.1 Mehrsprachigkeit

Das Tool unterstützt 5 Sprachen:
- Deutsch (de-DE) - Standard
- Englisch (en-US)
- Französisch (fr-FR)
- Spanisch (es-ES)
- Italienisch (it-IT)

Sprachpriorität:
1. `--language` Flag (einzelne Anfrage)
2. Konfigurationsdatei (`~/.tmdb/config.json`)
3. Umgebungsvariable (`TMDB_LANGUAGE`)
4. Standard (`de-DE`)

### 6.2 Farbschema

| Element | Farbe | Hex Code |
|---------|-------|----------|
| Primary (Borders, Titel) | Rot | `#E50914` |
| Secondary (Labels) | Gold | `#FFD700` |
| Success (Ratings, Running) | Grün | `#00D26A` |
| Muted (Zusatzinfo) | Grau | `#888888` |
| Text | Weiß | `#FFFFFF` |

### 6.2 Typografie

- **Titel:** Bold
- **Taglines:** Italic
- **Labels:** Bold, Gold
- **Werte:** Normal, Weiß
- **Sekundärinfo:** Dim, Grau

### 6.3 Icons/Emojis

| Element | Icon |
|---------|------|
| Film | 🎬 |
| Serie | 📺 |
| Suche | 🔍 |
| Fehler | ❌ |
| Info | ℹ️ |
| Stern (gefüllt) | ★ |
| Stern (leer) | ☆ |
| Status Running | ● |
| Status Ended | ○ |

---

## 7. Build & Distribution

### 7.1 Makefile

```makefile
BINARY_NAME=tmdb
VERSION=1.0.0

.PHONY: all build clean install test

all: build

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

build-all:
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)-darwin-arm64 .
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY_NAME)-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows-amd64.exe .

install: build
	mv $(BINARY_NAME) /usr/local/bin/

clean:
	rm -f $(BINARY_NAME)
	rm -rf dist/

test:
	go test -v ./...
```

### 7.2 Installation

**Via Go:**
```bash
go install github.com/matthias/tmdb-cli@latest
```

**Via Homebrew (macOS):**
```bash
brew tap matthias/tap
brew install tmdb-cli
```

**Manuell:**
```bash
# Binary herunterladen
curl -L https://github.com/matthias/tmdb-cli/releases/latest/download/tmdb-linux-amd64 -o tmdb
chmod +x tmdb
sudo mv tmdb /usr/local/bin/
```

---

## 8. Qualitätssicherung

### 8.1 Tests

| Test-Typ | Beschreibung |
|----------|--------------|
| Unit Tests | API Client, Rendering Functions |
| Integration Tests | TMDB API Calls (mit Mock) |
| E2E Tests | CLI Commands ausführen |

### 8.2 Code Quality

- `go fmt` für Formatierung
- `go vet` für statische Analyse
- `golangci-lint` für erweiterte Checks

---

## 9. Zukünftige Erweiterungen (v2+)

| Feature | Priorität | Beschreibung |
|---------|-----------|--------------|
| Caching | Hoch | Suchergebnisse lokal cachen |
| Person Search | Mittel | `tmdb person "Tom Hanks"` |
| Watchlist | Mittel | Lokale Merkliste |
| Poster ASCII | Niedrig | Poster als ASCII-Art anzeigen |
| Similar/Recommendations | Niedrig | Ähnliche Filme/Serien vorschlagen |
| Offline Mode | Niedrig | Gecachte Daten offline anzeigen |
| Shell Completions | Niedrig | Bash/Zsh/Fish Autocompletion |

---

## 10. Referenzen

- [TMDB API Dokumentation](https://developer.themoviedb.org/docs)
- [Cobra CLI Framework](https://cobra.dev/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Lipgloss](https://github.com/charmbracelet/lipgloss)

---

## Anhang A: API Key einrichten

1. Account erstellen: https://www.themoviedb.org/signup
2. Einloggen und zu Settings navigieren
3. API Section öffnen: https://www.themoviedb.org/settings/api
4. "Create" klicken und Formular ausfüllen
5. API Key (v3 auth) kopieren
6. Environment Variable setzen:

**Bash/Zsh:**
```bash
# Temporär
export TMDB_API_KEY='dein-api-key'

# Permanent (in ~/.bashrc oder ~/.zshrc)
echo "export TMDB_API_KEY='dein-api-key'" >> ~/.bashrc
```

**Fish:**
```fish
set -Ux TMDB_API_KEY 'dein-api-key'
```

**Windows (PowerShell):**
```powershell
$env:TMDB_API_KEY = 'dein-api-key'

# Permanent
[Environment]::SetEnvironmentVariable("TMDB_API_KEY", "dein-api-key", "User")
```

---

## Anhang B: Beispiel-Session

```bash
$ export TMDB_API_KEY='abc123...'

$ tmdb movie "matrix"
# → Zeigt Auswahlliste mit allen Matrix-Filmen
# → User wählt "The Matrix (1999)"
# → Zeigt detaillierte Infos

$ tmdb series "dark" --short
# → Zeigt kompakte Info zu "Dark"

$ tmdb movie "inception" --json | jq '.rating'
8.4

$ tmdb m "nicht existierender film"
ℹ️ Keine Filme gefunden für: nicht existierender film
```
