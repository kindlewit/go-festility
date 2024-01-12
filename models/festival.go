package models

type Fest struct {
	Id   string `bson:"id" form:"id" json:"id" binding:"required"`
	Name string `bson:"name" form:"name" json:"name" binding:"required"`
	From int    `bson:"from_date" form:"from_date" json:"from_date" binding:"required"`
	To   int    `bson:"to_date" form:"to_date" json:"to_date" binding:"required"`
	Url  string `bson:"url" form:"url" json:"url"`
}

type Slot struct {
	// Common for both
	Id         string `bson:"id" json:"id"`
	Type       string `bson:"slot_type" json:"slot_type" binding:"required"`
	ScheduleID string `bson:"schedule_id" json:"schedule_id"`
	ScreenID   string `bson:"screen_id" json:"screen_id" binding:"required"`
	Start      int    `bson:"start_time" json:"start_time"`
	Duration   int    `bson:"duration" json:"duration"`
	// Movie specific
	Title         string     `bson:"title" json:"title"`
	Synopsis      string     `bson:"synopsis" json:"synopsis"`
	Directors     []string   `bson:"directors" json:"directors,omitempty"`
	OriginalTitle string     `bson:"original_title" json:"original_title,omitempty"`
	Year          string     `bson:"year" json:"year,omitempty"`
	Genres        []string   `bson:"genres" json:"genres,omitempty"`
	Languages     []Language `bson:"languages" json:"languages,omitempty"`
	Countries     []Country  `bson:"countries" json:"countries,omitempty"`
	MovieId       int        `bson:"movie_id" json:"movie_id"`
}

type Schedule struct {
	Id     string `bson:"id" json:"id"`
	Fest   string `bson:"fest_id" json:"fest_id"`
	Custom bool   `bson:"custom" json:"custom"`
}

type CustomSchedule struct {
	Username string `bson:"username" json:"username"`
	Schedule
}
