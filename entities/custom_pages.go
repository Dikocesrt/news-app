package entities

type CustomPage struct {
	ID uint
	CustomURL string
	Content string
	User User
}

type CustomPageRepositoryInterface interface {
	GetAllCustomPages(metadata Metadata) ([]CustomPage, error)
	GetCustomPageByID(id uint) (CustomPage, error)
	CreateCustomPage(customPage CustomPage) (CustomPage, error)
	UpdateCustomPage(customPage CustomPage) (CustomPage, error)
	DeleteCustomPage(id uint, userID uint) error
}

type CustomPageUsecaseInterface interface {
	GetAllCustomPages(metadata Metadata) ([]CustomPage, error)
	GetCustomPageByID(id uint) (CustomPage, error)
	CreateCustomPage(customPage CustomPage) (CustomPage, error)
	UpdateCustomPage(customPage CustomPage) (CustomPage, error)
	DeleteCustomPage(id uint, userID uint) error
}