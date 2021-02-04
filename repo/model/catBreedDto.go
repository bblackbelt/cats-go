package model

type CatBreedDto struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	LifeSpan    string `json:"life_span"`
	Temperament string `json:"temperament"`
}
