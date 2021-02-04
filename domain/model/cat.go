package model

type Cat struct {
	Breeds []Breed `json:"breeds"`
	Height int     `json:"height"`
	Id     string  `json:"id"`
	Url    string  `json:"url"`
	Width  int     `json:"width"`
}
