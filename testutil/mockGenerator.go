package testutil

import (
	catsModel "cats/domain/model"
	catsDtoModel "cats/repo/model"
	"fmt"
)

func GenerateMockedCatBreedDto(n int) []catsDtoModel.CatBreedDto {
	retVal := make([]catsDtoModel.CatBreedDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsDtoModel.CatBreedDto{
			ID:          fmt.Sprintf("ID %d", i),
			Name:        fmt.Sprintf("Name %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateCatBreed(n int) []catsModel.Breed {
	retVal := make([]catsModel.Breed, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsModel.Breed{
			Id:          fmt.Sprintf("ID %d", i),
			Name:        fmt.Sprintf("Name %d", i),
			LifeSpan:    fmt.Sprintf("LifeSpan %d", i),
			Temperament: fmt.Sprintf("Temperament %d", i),
		})
	}
	return retVal
}

func GenerateMockedCatsDto(n int) []catsDtoModel.CatDto {
	retVal := make([]catsDtoModel.CatDto, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsDtoModel.CatDto{
			Height: i,
			ID:     fmt.Sprintf("Id %d", i),
			Width:  i,
			URL:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateMockedCatBreedDto(n),
		})
	}
	return retVal
}

func GenerateCat(n int) []catsModel.Cat {
	retVal := make([]catsModel.Cat, 0)
	for i := 0; i < n; i++ {
		retVal = append(retVal, catsModel.Cat{
			Height: i,
			Id:     fmt.Sprintf("Id %d", i),
			Width:  i,
			Url:    fmt.Sprintf("Url %d", i),
			Breeds: GenerateCatBreed(n),
		})
	}
	return retVal
}
