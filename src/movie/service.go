package movie

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/kindlewit/go-festility/src/constants"
)

const BASE_URL = "https://api.themoviedb.org/3"
const POSTER_URL = "https://www.themoviedb.org/t/p/w600_and_h900_bestv2"

// Fetches movie details from TMDB API.
func GetMovie(movieId string) (data TMDBmovie, err error) {
	var API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		return data, constants.ErrUnauthorized
	}

	// Read https://developers.themoviedb.org/3/movies/get-movie-details
	urlStructure := "%s/movie/%s?api_key=%s" // {BASE_URL}/movie/{movieId}?api_key={API_KEY}
	url := fmt.Sprintf(urlStructure, BASE_URL, movieId, API_KEY)
	resp, err := http.Get(url) // #nosec G107 -- URL is constructed from env config and path param
	if err != nil {
		return data, constants.ErrMissingApiKey
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return data, constants.ErrNoSuchRecord
	}
	if resp.StatusCode != 200 {
		return data, constants.ErrApiFetch
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiParse
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiParse
	}

	return data, nil
}

// Fetches directors of a movie from TMDB API.
func GetDirector(movieId string) (data []string, err error) {
	var API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		return data, constants.ErrMissingApiKey
	}

	var castsList struct {
		Crew []Crew `json:"crew"`
	}

	// Read https://developers.themoviedb.org/3/movies/get-movie-credits
	urlStructure := "%s/movie/%s/credits?api_key=%s" // {BASE_URL}/movie/{movieId}/credits?api_key={API_KEY}
	url := fmt.Sprintf(urlStructure, BASE_URL, movieId, API_KEY)
	resp, err := http.Get(url) // #nosec G107 -- URL is constructed from env config and path param
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiFetch
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return data, constants.ErrApiFetch
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiParse
	}

	err = json.Unmarshal(body, &castsList)

	for i := 0; i < len(castsList.Crew); i++ {
		if castsList.Crew[i].Job == "Director" {
			data = append(data, castsList.Crew[i].Name)
		}
	}

	return data, nil
}

// Searches for a movie in TMDB
func SearchMovie(term string) (data TMDBmovie, err error) {
	var API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		return data, constants.ErrUnauthorized
	}
	// Read https://developer.themoviedb.org/reference/search-movie
	urlStructure := "%s/search/movie?query=%s" // {BASE_URL}/search/movie?query={term}
	url := fmt.Sprintf(urlStructure, BASE_URL, term)
	authTokenHeader := fmt.Sprintf("Bearer: %s", API_KEY)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, constants.ErrApiFetch
	}
	req.Header.Set("Authorization", authTokenHeader)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return data, constants.ErrMissingApiKey
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return data, constants.ErrApiFetch
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiParse
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(err.Error())
		return data, constants.ErrApiParse
	}

	return data, nil
}

// Converts a TMDBmovie record into the Movie model.
func convertTMDBtoMovie(tmdb TMDBmovie) Movie {
	bound := Movie{
		Id:            tmdb.Id,
		Title:         tmdb.Title,
		OriginalTitle: tmdb.OriginalTitle,
		Date:          tmdb.Date,
		Synopsis:      tmdb.Synopsis,
		Tagline:       tmdb.Tagline,
		Duration:      tmdb.Runtime,
		Imdb:          tmdb.Imdb,
		Poster:        fmt.Sprintf("%s/%s", POSTER_URL, tmdb.Poster),
	}

	// Include Genres
	for i := 0; i < len(tmdb.Genres); i++ {
		bound.Genres = append(bound.Genres, tmdb.Genres[i].Name)
	}

	// Include Languages
	for i := 0; i < len(tmdb.Languages); i++ {
		lang := Language{
			Iso:  strings.ToUpper(tmdb.Languages[i].Iso),
			Name: tmdb.Languages[i].Name,
		}
		bound.Languages = append(bound.Languages, lang)
	}

	// Include Countries
	for i := 0; i < len(tmdb.Countries); i++ {
		cntry := Country{
			Iso:  tmdb.Countries[i].Iso,
			Name: tmdb.Countries[i].Name,
		}
		bound.Countries = append(bound.Countries, cntry)
	}

	// Year from date
	if hasYear, _ := regexp.MatchString(`\d\d\d\d`, tmdb.Date); hasYear {
		yearRegex := regexp.MustCompile(`\d\d\d\d`)
		bound.Year = yearRegex.FindString(tmdb.Date)
	}

	return bound
}
