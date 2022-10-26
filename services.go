/* All db services go here */
package main

import (
  "os"
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

const BASE_URL = "https://api.themoviedb.org/3"

// Fetch movie details from TMDB API
func getMovie(movieId string) (doc *TMDBmovie, success bool) {
  var API_KEY = os.Getenv("API_KEY");

  // Read https://developers.themoviedb.org/3/movies/get-movie-details
  resp, err := http.Get(fmt.Sprintf("%s/movie/%s?api_key=%s",BASE_URL, movieId, API_KEY));
  if err != nil {
    panic(err);
    return nil, false;
  }
  defer resp.Body.Close();
  // TODO: status check
  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    panic(err);
    return nil, false;
  }

  err = json.Unmarshal(body, &doc);
  if err != nil {
    panic(err);
    return nil, false;
  }

  return doc, true;
}

// Get director details of a movie id from TMDB API
func getDirector(movieId string) []string {
  var API_KEY = os.Getenv("API_KEY");

  var castsList struct {
    Crew  []Crew  `json:"crew"`
  };

  // Read https://developers.themoviedb.org/3/movies/get-movie-credits
  resp, err := http.Get(fmt.Sprintf("%s/movie/%s/credits?api_key=%s", BASE_URL, movieId, API_KEY));
  if err != nil {
    panic(err);
    return []string{};
  }
  defer resp.Body.Close();
  // TODO: status check
  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    panic(err);
    return []string{};
  }

  err = json.Unmarshal(body, &castsList);

  directors := []string{}

  for i := 0; i < len(castsList.Crew); i++ {
    if castsList.Crew[i].Job == "Director" {
      directors = append(directors, castsList.Crew[i].Name);
    }
  }

  return directors;
}

// func getGenre(ids []string) []Genre {
//   var API_KEY = os.Getenv("API_KEY");

//   resp, err := http.Get(fmt.Sprintf("%s/genre/movie/list?api_key=%s", BASE_URL, API_KEY));
//   if err != nil {
//     panic(err);
//     return []Genre{};
//   }
//   defer resp.Body.Close();
//   // TODO: status check
//   body, err := ioutil.ReadAll(resp.Body);
//   if err != nil {
//     panic(err);
//     return []Genre{};
//   }

//   var list struct {
//     Genres  []Genre `json:"genres"`
//   }

//   err = json.Unmarshal(body, &list);
//   if err != nil {
//     panic(err);
//     return []Genre{};
//   }

//   requiredGenres := []Genre{};

//   for i := 0; i < len(ids); i++ {
//     for j := 0; j < len(list.Genres); j++ { // O(n*n) indeed
//       if list.Genres[j].Id == ids[i] {
//         requiredGenres = append(requiredGenres, Genre(list.Genres[j]))
//       }
//     }
//   }

//   return requiredGenres;
// }

// Get list of movie ids for a list from TMDB API
func getListMovieIds(id string) []string {
  var API_KEY = os.Getenv("API_KEY");

  // Read https://developers.themoviedb.org/3/lists/get-list-details
  resp, err := http.Get(fmt.Sprintf("%s/list/%s?api_key=%s", BASE_URL, id, API_KEY));
  if err != nil {
    panic(err);
    return []string{};
  }
  defer resp.Body.Close();
  // TODO: status check
  body, err := ioutil.ReadAll(resp.Body);
  if err != nil {
    panic(err);
    return []string{};
  }

  var list struct {
    Items   []TMDBmovie `json:"items"`
    Count   int         `json:"item_count"`
  }

  err = json.Unmarshal(body, &list);
  if err != nil {
    panic(err);
    return []string{};
  }

  idList := []string{};
  for i := 0; i < len(list.Items); i++ {
    idList = append(idList, fmt.Sprint(list.Items[i].Id));
  }
  return idList;
}
