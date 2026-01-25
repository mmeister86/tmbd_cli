package tmdb

// SearchResponse ist die generische Antwort für Suchanfragen
type SearchResponse[T any] struct {
	Page         int `json:"page"`
	Results      []T `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

// MovieSearchResult repräsentiert ein Suchergebnis für Filme
type MovieSearchResult struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	OriginalTitle string  `json:"original_title"`
	Overview      string  `json:"overview"`
	ReleaseDate   string  `json:"release_date"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
	PosterPath    string  `json:"poster_path"`
	Adult         bool    `json:"adult"`
}

// TVSearchResult repräsentiert ein Suchergebnis für Serien
type TVSearchResult struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	Overview     string  `json:"overview"`
	FirstAirDate string  `json:"first_air_date"`
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
	PosterPath   string  `json:"poster_path"`
}

// MovieDetails enthält alle Details zu einem Film
type MovieDetails struct {
	ID            int      `json:"id"`
	Title         string   `json:"title"`
	OriginalTitle string   `json:"original_title"`
	Tagline       string   `json:"tagline"`
	Overview      string   `json:"overview"`
	ReleaseDate   string   `json:"release_date"`
	Runtime       int      `json:"runtime"`
	Budget        int64    `json:"budget"`
	Revenue       int64    `json:"revenue"`
	VoteAverage   float64  `json:"vote_average"`
	VoteCount     int      `json:"vote_count"`
	Genres        []Genre  `json:"genres"`
	Status        string   `json:"status"`
	Homepage      string   `json:"homepage"`
	ImdbID        string   `json:"imdb_id"`
	PosterPath    string   `json:"poster_path"`
	Credits       *Credits `json:"credits,omitempty"`
}

// TVDetails enthält alle Details zu einer Serie
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
	PosterPath       string    `json:"poster_path"`
	Credits          *Credits  `json:"credits,omitempty"`
}

// Genre repräsentiert ein Film-/Serien-Genre
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Credits enthält Cast und Crew Informationen
type Credits struct {
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}

// CastMember repräsentiert einen Schauspieler
type CastMember struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Character   string `json:"character"`
	Order       int    `json:"order"`
	ProfilePath string `json:"profile_path"`
}

// CrewMember repräsentiert ein Crew-Mitglied
type CrewMember struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Job         string `json:"job"`
	Department  string `json:"department"`
	ProfilePath string `json:"profile_path"`
}

// Network repräsentiert einen TV-Sender
type Network struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LogoPath string `json:"logo_path"`
}

// Creator repräsentiert einen Serien-Ersteller
type Creator struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ProfilePath string `json:"profile_path"`
}

// Season repräsentiert eine Staffel
type Season struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	SeasonNumber int    `json:"season_number"`
	EpisodeCount int    `json:"episode_count"`
	AirDate      string `json:"air_date"`
	Overview     string `json:"overview"`
	PosterPath   string `json:"poster_path"`
}

// MovieJSONOutput ist das Format für die JSON-Ausgabe von Filmen
type MovieJSONOutput struct {
	ID            int          `json:"id"`
	Title         string       `json:"title"`
	OriginalTitle string       `json:"original_title"`
	Year          string       `json:"year"`
	Runtime       int          `json:"runtime"`
	Rating        float64      `json:"rating"`
	VoteCount     int          `json:"vote_count"`
	Budget        int64        `json:"budget"`
	Revenue       int64        `json:"revenue"`
	Genres        []string     `json:"genres"`
	Directors     []string     `json:"directors"`
	Cast          []CastOutput `json:"cast"`
	Overview      string       `json:"overview"`
	ImdbID        string       `json:"imdb_id"`
	ImdbURL       string       `json:"imdb_url"`
	PosterURL     string       `json:"poster_url"`
}

// TVJSONOutput ist das Format für die JSON-Ausgabe von Serien
type TVJSONOutput struct {
	ID           int          `json:"id"`
	Name         string       `json:"name"`
	OriginalName string       `json:"original_name"`
	FirstAirDate string       `json:"first_air_date"`
	LastAirDate  string       `json:"last_air_date"`
	Seasons      int          `json:"seasons"`
	Episodes     int          `json:"episodes"`
	Rating       float64      `json:"rating"`
	VoteCount    int          `json:"vote_count"`
	Status       string       `json:"status"`
	Genres       []string     `json:"genres"`
	Networks     []string     `json:"networks"`
	Creators     []string     `json:"creators"`
	Cast         []CastOutput `json:"cast"`
	Overview     string       `json:"overview"`
	PosterURL    string       `json:"poster_url"`
}

// CastOutput ist das Format für Cast in der JSON-Ausgabe
type CastOutput struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}
