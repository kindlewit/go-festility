package services

import (
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "festility/models"
  "festility/constants"
)

const BASE_URL = "https://api.themoviedb.org/3"

// Fetch movie details from TMDB API
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
