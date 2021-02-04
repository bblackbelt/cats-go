package executor

import (
	"fmt"
	"net/http"
)

type doggoRequest struct {
}

func NewHttpRequestFactory() RequestFactory {
	return &doggoRequest{}
}

func (doggoRequest *doggoRequest) Build(page int, size int, breedID string) (*Request, error) {
	req, err := http.NewRequest("GET", newUrl(page, size, breedID), nil)
	if err != nil {
		return nil, err
	}
	return &Request{
		Req: req,
	}, nil
}

func urlBuilder(page int, limit int) string {
	return fmt.Sprintf("https://api.thecatapi.com/v1/images/search?page=%d&limit=%d&mime_types=image/jpeg", page, limit)
}

func newUrl(page int, limit int, breedID string) string {
	url := urlBuilder(page, limit)
	if breedID != "" {
		url += fmt.Sprintf("&breed_id=%s", breedID)
	}
	return url
}
