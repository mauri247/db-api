package models

type PayloadData struct {
	Name      string `json:"name"`
	Character string `json:"character"`
}

type Character struct {
	Name string `json:"name" bson:"name"`
	Race string `json:"race" bson:"race"`
	Ki   string `json:"ki" bson:"ki"`
}

type CharacterApi struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Ki          string `json:"ki"`
	MaxKi       string `json:"maxKi"`
	Race        string `json:"race"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Affiliation string `json:"affiliation"`
	DeletedAt   string `json:"deletedAt"`
}
