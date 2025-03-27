package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kindlewit/go-festility/models"
)

const POSTER_URL string = "https://www.themoviedb.org/t/p/w600_and_h900_bestv2"

// Helps bind movie data into slot data model.
func BindMovieToSlot(slot models.Slot, movie models.TMDBmovie) models.Slot {
	slot.Title = movie.Title

	if movie.OriginalTitle != "" && movie.Title != movie.OriginalTitle {
		slot.OriginalTitle = movie.OriginalTitle
	}

	// Year from date
	if hasYear, _ := regexp.MatchString(`\d\d\d\d`, movie.Date); hasYear {
		yearRegex := regexp.MustCompile(`\d\d\d\d`)
		slot.Year = yearRegex.FindString(movie.Date)
	}

	// slot.Synopsis = movie.Synopsis;
	slot.Duration = movie.Runtime
	slot.Languages = make([]models.Language, len(movie.Languages))
	slot.Countries = make([]models.Country, len(movie.Countries))
	slot.Genres = make([]string, len(movie.Genres))

	for i := 0; i < len(movie.Languages); i++ {
		slot.Languages[i] = models.Language(movie.Languages[i])
	}
	for i := 0; i < len(movie.Countries); i++ {
		slot.Countries[i] = models.Country(movie.Countries[i])
	}
	for i := 0; i < len(movie.Genres); i++ {
		slot.Genres[i] = movie.Genres[i].Name
	}
	return slot
}

// Helps bind cinema & screen data into CinemaScreen model.
func BindCinemaToScreen(screen models.Screen, cinema models.Cinema) models.CinemaScreen {
	bound := models.CinemaScreen{
		CinemaName: cinema.Name,
		Address:    cinema.Address,
		City:       cinema.City,
		PlusCode:   cinema.PlusCode,
		ScreenName: screen.Name,
		Id:         screen.Id,
	}

	return bound
}

// Helps bind directors to TMDBmovie record.
func ConvertTMDBtoMovie(movie models.TMDBmovie) models.Movie {
	bound := models.Movie{
		Id:            movie.Id,
		Title:         movie.Title,
		OriginalTitle: movie.OriginalTitle,
		Date:          movie.Date,
		Synopsis:      movie.Synopsis,
		Tagline:       movie.Tagline,
		Duration:      movie.Runtime,
		Imdb:          movie.Imdb,
		Poster:        fmt.Sprintf("%s/%s", POSTER_URL, movie.Poster),
	}

	// Include Genres
	for i := 0; i < len(movie.Genres); i++ {
		genre := movie.Genres[i]
		bound.Genres = append(bound.Genres, genre.Name)
	}

	// Include Languages
	for i := 0; i < len(movie.Languages); i++ {
		tmdbLang := movie.Languages[i]
		lang := models.Language{
			Iso:  strings.ToUpper(tmdbLang.Iso),
			Name: tmdbLang.Name,
		}
		bound.Languages = append(bound.Languages, lang)
	}

	// Include Countries
	for i := 0; i < len(movie.Countries); i++ {
		tmdbCountry := movie.Countries[i]
		cntry := models.Country{
			Iso:  tmdbCountry.Iso,
			Name: tmdbCountry.Name,
		}
		bound.Countries = append(bound.Countries, cntry)
	}

	// Year from date
	if hasYear, _ := regexp.MatchString(`\d\d\d\d`, movie.Date); hasYear {
		yearRegex := regexp.MustCompile(`\d\d\d\d`)
		bound.Year = yearRegex.FindString(movie.Date)
	}

	return bound
}
