# TMDB CLI

Ein plattformübergreifendes Command-Line Tool zum Abrufen von Film- und Serieninformationen von [The Movie Database (TMDB)](https://www.themoviedb.org/).

## Features

- 🎬 Suche nach Filmen mit detaillierten Informationen
- 📺 Suche nach Serien mit Staffelübersicht
- 🎨 Ansprechende, farbige Terminal-Ausgabe
- 📋 Interaktive Auswahlliste bei mehreren Ergebnissen
- 📄 JSON-Ausgabe für Skript-Integration
- 🌍 Mehrsprachige Unterstützung

## Installation

### Voraussetzungen

- Go 1.22 oder höher
- TMDB API Key (kostenlos erhältlich)

### Via Go

```bash
go install github.com/mmeister86/tmbd_cli@latest
```

### Manuell bauen

```bash
git clone https://github.com/mmeister86/tmbd_cli.git
cd tmbd_cli
make build
make install
```

### Binaries herunterladen

Vorkompilierte Binaries für verschiedene Plattformen findest du unter [Releases](https://github.com/mmeister86/tmbd_cli/releases).

## API Key einrichten

1. Erstelle einen Account auf [TMDB](https://www.themoviedb.org/signup)
2. Gehe zu [API Settings](https://www.themoviedb.org/settings/api)
3. Erstelle einen API Key (v3 auth)
4. Setze die Umgebungsvariable:

**Bash/Zsh:**
```bash
export TMDB_API_KEY='dein-api-key'

# Permanent (in ~/.bashrc oder ~/.zshrc)
echo "export TMDB_API_KEY='dein-api-key'" >> ~/.zshrc
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

## Verwendung

### Filme suchen

```bash
tmdb movie "The Matrix"
tmdb m "Inception"
tmdb film "Pulp Fiction"
```

### Sprache konfigurieren

```bash
# Sprachauswahlliste anzeigen
tmdb language

# Sprache für einzelne Anfrage überschreiben
tmdb movie "The Matrix" --language en-US
```

### Serien suchen

```bash
tmdb series "Breaking Bad"
tmdb s "Dark"
tmdb tv "The Office"
tmdb show "Stranger Things"
```

### Optionen

| Flag | Kurzform | Beschreibung |
|------|----------|--------------|
| `--help` | `-h` | Hilfe anzeigen |
| `--version` | `-v` | Version anzeigen |
| `--json` | | Ausgabe als JSON |
| `--short` | `-s` | Kompakte Ausgabe |
| `--language` | `-l` | Sprache überschreiben (z.B. `en-US`) |

### Sprache konfigurieren

Die Standard-Sprache ist Deutsch (`de-DE`). Du kannst die Sprache dauerhaft ändern:

```bash
# Interaktive Sprachauswahl
tmdb language
```

Unterstützte Sprachen:
- Deutsch (de-DE)
- Englisch (en-US)
- Französisch (fr-FR)
- Spanisch (es-ES)
- Italienisch (it-IT)

Alternativ kannst du die Sprache für einzelne Anfragen überschreiben:

```bash
# Englische Ergebnisse
tmdb movie "The Matrix" --language en-US

# Französische Ergebnisse
tmdb series "Les Misérables" --language fr-FR
```

Die Sprachauswahl wird in `~/.tmdb/config.json` gespeichert.

### Beispiele

```bash
# Detaillierte Filminfos
tmdb movie "The Matrix"

# Kompakte Ausgabe
tmdb movie "Inception" --short

# JSON-Ausgabe (z.B. für jq)
tmdb movie "Fight Club" --json | jq '.rating'

# Englische Ergebnisse
tmdb series "Game of Thrones" --language en-US
```

## Konfiguration

Die Konfiguration wird in `~/.tmdb/config.json` gespeichert.

| Umgebungsvariable | Beschreibung | Standard |
|-------------------|--------------|----------|
| `TMDB_API_KEY` | TMDB API Key (v3 auth) | *Pflicht* |
| `TMDB_LANGUAGE` | Sprache für Ergebnisse (wenn keine Config vorhanden) | `de-DE` |

**Sprachpriorität:**
1. `--language` Flag (einzelne Anfrage)
2. Config-Datei (`~/.tmdb/config.json`)
3. Umgebungsvariable (`TMDB_LANGUAGE`)
4. Standard (`de-DE`)

## Entwicklung

```bash
# Dependencies installieren
go mod tidy

# Bauen
make build

# Tests ausführen
make test

# Formatieren und Linting
make lint

# Cross-Compilation für alle Plattformen
make build-all
```

## Technologie-Stack

- [Go](https://golang.org/) - Programmiersprache
- [Cobra](https://cobra.dev/) - CLI Framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI Framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal Styling
- [TMDB API](https://developer.themoviedb.org/) - Datenquelle

## Lizenz

MIT License - siehe [LICENSE](LICENSE) für Details.

## Danksagungen

- [The Movie Database](https://www.themoviedb.org/) für die kostenlose API
- [Charmbracelet](https://charm.sh/) für die fantastischen Terminal-Libraries
