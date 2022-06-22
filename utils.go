package main

import "strings"

const POSTER_BASE_URL = "https://image.tmdb.org/t/p/w300_and_h450_bestv2";

func reformat(doc *TMDBmovie, directors []string) Movie {
	resp := Movie{
		Id: doc.Id,
		Title: doc.Title,
		Directors: directors,
		OriginalTitle: doc.OriginalTitle,
		Date: doc.Date,
		Synopsis: doc.Synopsis,
		Tagline: doc.Tagline,
		Runtime: doc.Runtime,
		Imdb: doc.Imdb,
		Poster: POSTER_BASE_URL + doc.Poster,
	};


	// Convert genres
	genreList := []string{};
	for i := 0; i < len(doc.Genres); i++ {
		genreList = append(genreList, doc.Genres[i].Name)
	}
	resp.Genres = genreList

	// Convert languages
	langList := []Language{};
	for i := 0; i < len(doc.Languages); i++ {
		// langList = append(langList, Language(doc.Languages[i]))
		langList = append(langList, Language{
			Name: doc.Languages[i].Name,
			Iso: strings.ToUpper(doc.Languages[i].Iso),
		})
	}
	resp.Languages = langList;

	// Convert countries
	cntList := []Country{};
	for i := 0; i < len(doc.Countries); i++ {
		cntList = append(cntList, Country(doc.Countries[i]));
	}
	resp.Countries = cntList;

	return resp;
}

