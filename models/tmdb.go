package models

type TMDBlanguage struct {
  Iso   string  `json:"iso_639_1"`
  Name  string  `json:"english_name"`
}

type TMDBcountry struct {
  Iso   string  `json:"iso_3166_1"`
  Name  string  `json:"name"`
}

type Genre struct {
  Id    int     `json:"id"`
  Name  string  `json:"name"`
}

type TMDBmovie struct {
  Id              int             `json:"id"`
  Title           string          `json:"title"`
  OriginalTitle   string          `json:"original_title"`
  Date            string          `json:"release_date"`
  Synopsis        string          `json:"overview"`
  Tagline         string          `json:"tagline"`
  Genres          []Genre         `json:"genres"`
  Languages       []TMDBlanguage  `json:"spoken_languages"`
  Countries       []TMDBcountry   `json:"production_countries"`
  Runtime         int             `json:"runtime"`
  Imdb            string          `json:"imdb_id"`
  Poster          string          `json:"poster_path"`
}
