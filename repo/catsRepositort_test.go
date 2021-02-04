package repo_test

import (
	"bytes"
	"cats/executor"
	mocks "cats/executor/mocks"
	"cats/repo"
	model "cats/repo/model"
	util "cats/testutil"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	mockExecutor := new(mocks.Executor)
	factory := new(mocks.RequestFactory)
	repo := repo.NewCatRepository(mockExecutor, factory)

	limit := 25

	req, error := http.NewRequest("GET", "", nil)
	assert.NoError(t, error)

	wrapReq := &executor.Request{Req: req}
	factory.On("Build", 1, limit, "").Return(wrapReq, nil)

	mockExecutor.On("Execute", wrapReq).Return(nil, errors.New("error"))

	_, err := repo.GetCats(1, limit, "")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}

func Test_EmptyResponseEmptyDataSet(t *testing.T) {
	mockExecutor := new(mocks.Executor)
	factory := new(mocks.RequestFactory)
	repo := repo.NewCatRepository(mockExecutor, factory)

	limit := 25
	req, error := http.NewRequest("GET", "", nil)
	assert.NoError(t, error)

	wrapReq := &executor.Request{Req: req}
	factory.On("Build", 1, limit, "").Return(wrapReq, nil)

	emptyResponse, error := json.Marshal(make([]model.CatDto, 0))
	assert.NoError(t, error)

	wrapResponse := &executor.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(emptyResponse))),
	}
	mockExecutor.On("Execute", wrapReq).Return(wrapResponse, nil)

	resp, err := repo.GetCats(1, limit, "")

	assert.Nil(t, err)
	assert.Empty(t, resp)
}

func Test_Response(t *testing.T) {
	mockExecutor := new(mocks.Executor)
	factory := new(mocks.RequestFactory)
	repo := repo.NewCatRepository(mockExecutor, factory)

	limit := 25
	req, error := http.NewRequest("GET", "", nil)
	assert.NoError(t, error)

	wrapReq := &executor.Request{Req: req}
	factory.On("Build", 1, limit, "").Return(wrapReq, nil)

	data := util.GenerateMockedCatsDto(limit)
	dataJson, error := json.Marshal(data)
	assert.NoError(t, error)

	wrapResponse := &executor.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(dataJson))),
	}
	mockExecutor.On("Execute", wrapReq).Return(wrapResponse, nil)

	resp, err := repo.GetCats(1, limit, "")

	assert.Nil(t, err)
	assert.ElementsMatch(t, data, resp)
}

func Test_CannotGenerateRequestReturnNil(t *testing.T) {
	mockExecutor := new(mocks.Executor)
	factory := new(mocks.RequestFactory)
	repo := repo.NewCatRepository(mockExecutor, factory)

	limit := 25

	factory.On("Build", 1, limit, "").Return(nil, errors.New("error"))

	_, err := repo.GetCats(1, limit, "")

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "error")
}
