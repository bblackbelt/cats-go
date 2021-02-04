package executor

type RequestFactory interface {
	Build(page int, size int, breedID string) (*Request, error)
}
