package utils

import "festility/models"

// Helps bind movie data into slot data model.
func BindMovieToSlot(slot models.Slot, movie models.TMDBmovie) models.Slot {
  slot.Title = movie.Title;

  if (movie.OriginalTitle != "" && movie.Title != movie.OriginalTitle) {
    slot.OriginalTitle = movie.OriginalTitle;
  }

  slot.Synopsis = movie.Synopsis;
  slot.Duration = movie.Runtime;
  slot.Languages = make([]models.Language, len(movie.Languages));
  slot.Countries = make([]models.Country, len(movie.Countries));
  slot.Genres = make([]string, len(movie.Genres));

  for i := 0; i < len(movie.Languages); i++ {
    slot.Languages[i] = models.Language(movie.Languages[i]);
  }
  for i := 0; i < len(movie.Countries); i++ {
    slot.Countries[i] = models.Country(movie.Countries[i]);
  }
  for i:= 0; i < len(movie.Genres); i++ {
    slot.Genres[i] = movie.Genres[i].Name;
  }
  return slot;
}