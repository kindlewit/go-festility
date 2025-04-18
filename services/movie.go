package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
)

const BASE_URL = "https://api.themoviedb.org/3"

// Fetches movie details from TMDB API.
func GetMovie(movieId string) (data models.TMDBmovie, err error) {
	var API_KEY = os.Getenv("API_KEY")
	if API_KEY == "" {
		return data, constants.ErrUnauthorized
	}

	// Read https://developers.themoviedb.org/3/movies/get-movie-details
	urlStructure := "%s/movie/%s?api_key=%s" // {BASE_URL}/movie/{movieId}?api_key={API_KEY}
	url := fmt.Sprintf(urlStructure, BASE_URL, movieId, API_KEY)
	resp, err := http.Get(url)
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
		Crew []models.Crew `json:"crew"`
	}

	// Read https://developers.themoviedb.org/3/movies/get-movie-credits
	urlStructure := "%s/movie/%s/credits?api_key=%s" // {BASE_URL}/movie/{movieId}/credits?api_key={API_KEY}
	url := fmt.Sprintf(urlStructure, BASE_URL, movieId, API_KEY)
	resp, err := http.Get(url)
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
func SearchMovie(term string) (data models.TMDBmovie, err error) {
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
