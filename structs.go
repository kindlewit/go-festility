/* All structures & interfaces go here */
package main

type TMDBlanguage struct {
  Iso   string  `json:"iso_639_1"`
  Name  string  `json:"english_name"`
}

type TMDBcountry struct {
  Iso   string  `json:"iso_3166_1"`
  Name  string  `json:"name"`
}

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

type Movie struct {
  Id              int         `json:"id"`
  Title           string      `json:"title"`
  Directors				[]string    `json:"directors"`
  OriginalTitle		string      `json:"original_title"`
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

type Crew struct {
  Name  string  `json:"name"`
  Job   string  `json:"job"`
}

type Fest struct {
  Id      string    `bson:"id" form:"id" json:"id" binding:"required"`
  Name    string    `bson:"name" form:"name" json:"name" binding:"required"`
  From    int       `bson:"from_date" form:"from_date" json:"from_date" binding:"required"`
  To      int       `bson:"to_date" form:"to_date" json:"to_date" binding:"required"`
  Url     string    `bson:"url" form:"url" json:"url"`
}

type Slot struct {
  // Common for both
  Id              string      `bson:"id" json:"id"`
  Type            string      `bson:"slot_type" binding:"required" json:"slot_type" binding:"required"`
  ScheduleID      string      `bson:"schedule_id" json:"schedule_id"`
  Title           string      `bson:"title" json:"title"`
  Synopsis        string      `bson:"synopsis" json:"synopsis"`
  Start           int         `bson:"start_time" json:"start_time"`
  Duration        int         `bson:"duration" json:"duration"`
  // Movie specific
  Directors       []string    `bson:"directors" json:"directors,omitempty"`
  OriginalTitle		string      `bson:"original_title" json:"original_title,omitempty"`
  Genres          []string    `bson:"genres" json:"genres,omitempty"`
  Languages       []Language  `bson:"languages" json:"languages,omitempty"`
  Countries       []Country   `bson:"countries" json:"countries,omitempty"`
  MovieId         int         `bson:"movie_id" json:"movie_id"`
}

type Schedule struct {
  Id          string  `bson:"id" json:"id"`
  Fest        string  `bson:"fest_id" json:"fest_id"`
  Custom      bool    `bson:"custom" json:"custom"`
  Username    string  `bson:"username" json:"username"`
}
