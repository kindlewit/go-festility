package services

import (
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/kindlewit/go-festility/models"
  "github.com/kindlewit/go-festility/constants"
)

const BASE_URL = "https://api.themoviedb.org/3"

// Fetch movie details from TMDB API.
func GetMovie(movieID string) (data models.TMDBmovie, err error) {
  var API_KEY = os.Getenv("API_KEY");

  // Read https://developers.themoviedb.org/3/movies/get-movie-details
  resp, err := http.Get(fmt.Sprintf("%s/movie/%s?api_key=%s", BASE_URL, movieID, API_KEY));
  if err != nil {
    return data, constants.ErrApiFetch;
  }
  defer resp.Body.Close();
  // TODO: status check
  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.ErrApiParse;
  }

  err = json.Unmarshal(body, &data);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.ErrApiParse;
  }

  return data, nil;
}

// Get director details of a movie id from TMDB API.
func GetDirector(movieID string) (data []string, err error) {
  var API_KEY = os.Getenv("API_KEY");

  var castsList struct {
    Crew  []models.Crew  `json:"crew"`
  };

  // Read https://developers.themoviedb.org/3/movies/get-movie-credits
  resp, err := http.Get(fmt.Sprintf("%s/movie/%s/credits?api_key=%s", BASE_URL, movieID, API_KEY));
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.ErrApiFetch;
  }
  defer resp.Body.Close();

  // TODO: status check

  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.ErrApiParse;
  }

  err = json.Unmarshal(body, &castsList);

  for i := 0; i < len(castsList.Crew); i++ {
    if castsList.Crew[i].Job == "Director" {
      data = append(data, castsList.Crew[i].Name);
    }
  }

  return data, nil;
}
