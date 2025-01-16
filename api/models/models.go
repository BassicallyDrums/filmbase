package models

type Film struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	VoteAverage float32 `json:"vote_average"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	GenreIds    []int   `json:"genre_ids"`
}

type TmdbResponse struct {
	Page    int    `json:"page"`
	Results []Film `json:"results"`
}

type Auth struct {
	APIKey   string
	APIToken string
}
