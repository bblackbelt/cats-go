package adapter

import (
	"cats/domain/model"
	catsDto "cats/repo/model"
)

type CatsMapper interface {
	Map(cats []catsDto.CatDto) []model.Cat
}

type catsMapper struct {
}

func NewCatMapper() *catsMapper {
	return &catsMapper{}
}

func mapBreedsDto(breedsDto []catsDto.CatBreedDto) []model.Breed {
	var retVal = make([]model.Breed, 0)
	for _, breedDto := range breedsDto {
		breed := model.Breed{
			Id:          breedDto.ID,
			Name:        breedDto.Name,
			LifeSpan:    breedDto.LifeSpan,
			Temperament: breedDto.Temperament,
		}
		retVal = append(retVal, breed)
	}
	return retVal
}

func mapCatDto(catDto catsDto.CatDto) model.Cat {
	return model.Cat{
		Height: catDto.Height,
		Width:  catDto.Width,
		Id:     catDto.ID,
		Url:    catDto.URL,
		Breeds: mapBreedsDto(catDto.Breeds),
	}
}

func (mapper *catsMapper) Map(cats []catsDto.CatDto) []model.Cat {
	var retVal = make([]model.Cat, 0)
	for _, catDto := range cats {
		retVal = append(retVal, mapCatDto(catDto))
	}
	return retVal
}
