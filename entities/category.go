package entities

type Category struct {
	ID   uint
	Name string
}

type CategoryRepositoryInterface interface {
	CreateCategory(category Category) (Category, error)
	GetAllCategories(metadata Metadata) ([]Category, error)
	GetCategoryByID(categoryID uint) (Category, error)
	UpdateCategory(category Category) (Category, error)
	DeleteCategory(categoryID uint) error
}

type CategoryUsecaseInterface interface {
	CreateCategory(category Category) (Category, error)
	GetAllCategories(metadata Metadata) ([]Category, error)
	GetCategoryByID(categoryID uint) (Category, error)
	UpdateCategory(category Category) (Category, error)
	DeleteCategory(categoryID uint) error
}