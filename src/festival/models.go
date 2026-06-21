package festival

type Fest struct {
	Id   string `bson:"id" form:"id" json:"id" binding:"required"`
	Name string `bson:"name" form:"name" json:"name" binding:"required"`
	From int    `bson:"from_date" form:"from_date" json:"from_date" binding:"required"`
	To   int    `bson:"to_date" form:"to_date" json:"to_date" binding:"required"`
	Url  string `bson:"url" form:"url" json:"url"`
}
