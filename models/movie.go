package models

type Language struct {
  Iso   string  `json:"iso"`
  Name  string  `json:"name"`
}

type Country struct {
  Iso   string  `json:"iso"`
  Name  string  `json:"name"`
}

type Movie struct {
  Id              int         `json:"id"`
  Title           string      `json:"title"`
  Directors       []string    `json:"directors"`
  OriginalTitle   string      `json:"original_title"`
  Date            string      `json:"date"`
  Tagline         string      `json:"tagline"`
  Synopsis        string      `json:"synopsis"`
  Genres          []string    `json:"genres"`
  Languages       []Language  `json:"languages"`
  Countries       []Country   `json:"countries"`
  Duration        int         `json:"duration"`
  Imdb            string      `json:"imdb_id"`
  Poster          string      `json:"poster"`
}

type MovieSlot struct {
	Directors       []string    `bson:"directors" json:"directors,omitempty"`
  OriginalTitle   string      `bson:"original_title" json:"original_title,omitempty"`
  Genres          []string    `bson:"genres" json:"genres,omitempty"`
  Languages       []Language  `bson:"languages" json:"languages,omitempty"`
  Countries       []Country   `bson:"countries" json:"countries,omitempty"`
  MovieId         int         `bson:"movie_id" json:"movie_id"`
	Slot
}
