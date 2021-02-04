package model

type CatDto struct {
	Breeds []CatBreedDto `json:"breeds"`
	Height int           `json:"height"`
	ID     string        `json:"id"`
	URL    string        `json:"url"`
	Width  int           `json:"width"`
}
