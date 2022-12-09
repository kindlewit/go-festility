package models

type Language struct {
  Iso   string  `json:"iso"`
  Name  string  `json:"name"`
}

type Country struct {
  Iso   string  `json:"iso"`
  Name  string  `json:"name"`
}

type Genre struct {
  Id    int     `json:"id"`
  Name  string  `json:"name"`
}

type Crew struct {
  Name  string  `json:"name"`
  Job   string  `json:"job"`
}


type Movie struct {
  Id              int         `json:"id"`
  Title           string      `json:"title"`
  Directors       []string    `json:"directors"`
  OriginalTitle   string      `json:"original_title"`
  Date            string      `json:"date"`
  Year            string      `json:"year"`
  Tagline         string      `json:"tagline"`
  Synopsis        string      `json:"synopsis"`
  Genres          []string    `json:"genres"`
  Languages       []Language  `json:"languages"`
  Countries       []Country   `json:"countries"`
  Duration        int         `json:"duration"`
  Imdb            string      `json:"imdb_id"`
  Poster          string      `json:"poster"`
}
