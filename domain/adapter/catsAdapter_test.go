package adapter_test

import (
	"cats/domain/adapter"
	dtoModel "cats/repo/model"
	util "cats/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyMap(t *testing.T) {
	assert.Empty(t, adapter.NewCatMapper().Map(make([]dtoModel.CatDto, 0)))
}

func TestMapLen(t *testing.T) {
	n := 5
	values := util.GenerateMockedCatsDto(n)
	assert.Len(t, adapter.NewCatMapper().Map(values), n)
}

func TestMapContent(t *testing.T) {
	n := 5
	values := util.GenerateMockedCatsDto(n)
	assert.Equal(t, adapter.NewCatMapper().Map(values), util.GenerateCat(n))
}
