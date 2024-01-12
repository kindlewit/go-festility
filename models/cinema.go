package models

type Cinema struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name" binding:"required"`
	Address  string `bson:"address" json:"address,omitempty"`
	City     string `bson:"city" json:"city" binding:"required"`
	PlusCode string `bson:"google_plus_code" json:"google_plus_code,omitempty"`
}

type Screen struct {
	Id       string `bson:"id" json:"id"`
	Name     string `bson:"name" json:"name" binding:"required"`
	CinemaId string `bson:"cinema_id" json:"cinema_id"`
}

type CinemaScreen struct {
	Id         string `json:"id"`
	ScreenName string `json:"screen_name"`
	CinemaName string `json:"cinema_name"`
	Address    string `json:"address"`
	City       string `json:"city" binding:"required"`
	PlusCode   string `json:"google_plus_code,omitempty"`
}
