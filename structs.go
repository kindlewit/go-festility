/* All structures & interfaces go here */
package main

type TMDBlanguage struct {
	Iso		string	`json:"iso_639_1"`
	Name	string	`json:"english_name"`
}

type TMDBcountry struct {
	Iso		string	`json:"iso_3166_1"`
	Name	string	`json:"name"`
}

type Language struct {
	Iso		string		`json:"iso"`
	Name	string		`json:"name"`
}

type Country struct {
	Iso		string	`json:"iso"`
	Name	string	`json:"name"`
}

type Genre struct {
	Id		int			`json:"id"`
	Name	string	`json:"name"`
}

type TMDBmovie struct {
	Id							int							`json:"id"`
	Title						string					`json:"title"`
	OriginalTitle		string					`json:"original_title"`
	Date						string					`json:"release_date"`
	Synopsis				string					`json:"overview"`
	Tagline					string					`json:"tagline"`
	Genres					[]Genre					`json:"genres"`
	Languages				[]TMDBlanguage	`json:"spoken_languages"`
	Countries				[]TMDBcountry		`json:"production_countries"`
	Runtime					int							`json:"runtime"`
	Imdb						string					`json:"imdb_id"`
	Poster					string					`json:"poster_path"`
}

type Movie struct {
	Id							int					`json:"id"`
	Title						string			`json:"title"`
	Directors				[]string		`json:"directors"`
	OriginalTitle		string			`json:"original_title"`
	Date						string			`json:"date"`
	Tagline					string			`json:"tagline"`
	Synopsis				string			`json:"synopsis"`
	Genres					[]string		`json:"genres"`
	Languages				[]Language	`json:"languages"`
	Countries				[]Country		`json:"countries"`
	Runtime					int					`json:"runtime"`
	Imdb						string			`json:"imdb_id"`
	Poster					string			`json:"poster"`
}

type Crew struct {
	Name		string		`json:"name"`
	Job			string		`json:"job"`
}

type Fest struct {
	Id			string		`form:"id" json:"id" binding:"required"`
	Name		string		`form:"name" json:"name" binding:"required"`
	From		int				`form:"from_date" json:"from_date" binding:"required"`
	To			int				`form:"to_date" json:"to_date" binding:"required"`
	Url			string		`form:"url" json:"url"`
}
