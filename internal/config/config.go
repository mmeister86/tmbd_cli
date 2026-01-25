package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

var (
	// ErrConfigNotFound wird zurückgegeben, wenn die Konfigurationsdatei nicht existiert
	ErrConfigNotFound = errors.New("konfigurationsdatei nicht gefunden")
)

// Config repräsentiert die Konfiguration des CLI
type Config struct {
	Language string `json:"language"`
}

// configDir gibt das Konfigurationsverzeichnis zurück
func configDir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".tmdb")
}

// configPath gibt den Pfad zur Konfigurationsdatei zurück
func configPath() string {
	return filepath.Join(configDir(), "config.json")
}

// Load lädt die Konfiguration aus der Datei
// Wenn die Datei nicht existiert, wird eine Standardkonfiguration zurückgegeben
func Load() (*Config, error) {
	path := configPath()

	// Prüfen ob Datei existiert
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, ErrConfigNotFound
	}

	// Datei lesen
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// JSON parsen
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Save speichert die Konfiguration in eine Datei
func Save(cfg *Config) error {
	// Konfigurationsverzeichnis erstellen, falls es nicht existiert
	dir := configDir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Konfiguration als JSON serialisieren
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	// In Datei schreiben
	path := configPath()
	return os.WriteFile(path, data, 0644)
}

// GetDefaultConfig gibt die Standardkonfiguration zurück
func GetDefaultConfig() *Config {
	return &Config{
		Language: "de-DE",
	}
}
