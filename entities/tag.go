package entities

type Tag struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TagRepositoryInterface interface {
	CreateTag(tag Tag) (Tag, error)
}

type TagUsecaseInterface interface {
	CreateTag(tag Tag) (Tag, error)
}