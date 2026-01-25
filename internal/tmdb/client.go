package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	baseURL        = "https://api.themoviedb.org/3"
	defaultTimeout = 10 * time.Second
)

var (
	// ErrNoAPIKey wird zurückgegeben, wenn kein API Key gesetzt ist
	ErrNoAPIKey = errors.New("TMDB_API_KEY nicht gesetzt")
	// ErrNoResults wird zurückgegeben, wenn keine Ergebnisse gefunden wurden
	ErrNoResults = errors.New("keine Ergebnisse gefunden")
	// ErrAPIError wird bei API-Fehlern zurückgegeben
	ErrAPIError = errors.New("TMDB API Fehler")
)

// Client ist der TMDB API Client
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient erstellt einen neuen TMDB Client
func NewClient() (*Client, error) {
	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		return nil, ErrNoAPIKey
	}

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}, nil
}

// GetLanguage gibt die konfigurierte Sprache zurück
func GetLanguage() string {
	lang := os.Getenv("TMDB_LANGUAGE")
	if lang == "" {
		return "de-DE"
	}
	return lang
}

// SearchMovies sucht nach Filmen
func (c *Client) SearchMovies(query string, language string) ([]MovieSearchResult, error) {
	endpoint := fmt.Sprintf("%s/search/movie", baseURL)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("query", query)
	params.Set("include_adult", "false")

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var result SearchResponse[MovieSearchResult]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return result.Results, nil
}

// GetMovieDetails lädt die Details zu einem Film
func (c *Client) GetMovieDetails(id int, language string) (*MovieDetails, error) {
	endpoint := fmt.Sprintf("%s/movie/%d", baseURL, id)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("append_to_response", "credits")

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var movie MovieDetails
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return &movie, nil
}

// SearchTV sucht nach Serien
func (c *Client) SearchTV(query string, language string) ([]TVSearchResult, error) {
	endpoint := fmt.Sprintf("%s/search/tv", baseURL)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("query", query)

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var result SearchResponse[TVSearchResult]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return result.Results, nil
}

// GetTVDetails lädt die Details zu einer Serie
func (c *Client) GetTVDetails(id int, language string) (*TVDetails, error) {
	endpoint := fmt.Sprintf("%s/tv/%d", baseURL, id)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("append_to_response", "credits")

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var tv TVDetails
	if err := json.NewDecoder(resp.Body).Decode(&tv); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return &tv, nil
}

// SearchPeople sucht nach Personen
func (c *Client) SearchPeople(query string, language string) ([]PersonSearchResult, error) {
	endpoint := fmt.Sprintf("%s/search/person", baseURL)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("query", query)
	params.Set("include_adult", "false")

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var result SearchResponse[PersonSearchResult]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return result.Results, nil
}

// GetPersonDetails lädt die Details zu einer Person
func (c *Client) GetPersonDetails(id int, language string) (*PersonDetails, error) {
	endpoint := fmt.Sprintf("%s/person/%d", baseURL, id)

	params := url.Values{}
	params.Set("api_key", c.apiKey)
	params.Set("language", language)
	params.Set("append_to_response", "combined_credits")

	resp, err := c.httpClient.Get(fmt.Sprintf("%s?%s", endpoint, params.Encode()))
	if err != nil {
		return nil, fmt.Errorf("API-Anfrage fehlgeschlagen: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: Status %d", ErrAPIError, resp.StatusCode)
	}

	var person PersonDetails
	if err := json.NewDecoder(resp.Body).Decode(&person); err != nil {
		return nil, fmt.Errorf("JSON-Dekodierung fehlgeschlagen: %w", err)
	}

	return &person, nil
}
