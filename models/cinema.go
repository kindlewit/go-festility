package models

type Cinema struct {
  Id        string    `bson:"id" json:"id"`
  Name      string    `bson:"name" binding:"required" json:"name" binding:"required"`
  Address   string    `bson:"address" json:"address,omitempty"`
  City      string    `bson:"city" binding:"required" json:"city" binding:"required"`
  PlusCode  string    `bson:"google_plus_code" json:"google_plus_code,omitempty"`
}

type Screen struct {
  Id          string    `bson:"id" json:"id"`
  Name        string    `bson:"name" binding:"required" json:"name" binding:"required"`
  CinemaID    string    `bson:"cinema_id" binding:"required" json:"cinema_id" binding:"required"`
}
