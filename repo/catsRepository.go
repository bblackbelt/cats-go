package repo

import (
	"cats/executor"
	"cats/repo/model"
	"encoding/json"
)

type CatRepository interface {
	GetCats(page int, limit int, breedID string) ([]model.CatDto, error)
}

type catRepository struct {
	Executor       executor.Executor
	RequestFactory executor.RequestFactory
}

func NewCatRepository(Executor executor.Executor, factory executor.RequestFactory) CatRepository {
	return &catRepository{
		Executor:       Executor,
		RequestFactory: factory,
	}
}

func (catRepository *catRepository) GetCats(page int, limit int, breedID string) ([]model.CatDto, error) {
	req, e := catRepository.RequestFactory.Build(page, limit, breedID)
	if e != nil {
		return nil, e
	}
	response, eww := catRepository.Executor.Execute(req)
	if eww != nil {
		return nil, eww
	}
	defer response.Body.Close()
	var data []model.CatDto
	json.NewDecoder(response.Body).Decode(&data)
	return data, eww
}
