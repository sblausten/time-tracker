package models

type Session struct {
	UserId      string `json:"userId" bson:"userId"`
	SessionName string `json:"name" bson:"name"`
	Start       string `json:"start" bson:"start"`
	End         string `json:"end" bson:"end"`
	Duration    int64 `json:"duration" bson:"duration"`
}
