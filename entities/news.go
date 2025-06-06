package entities

type News struct {
	ID         uint
	Content    string
	Category Category
	User       User
	Comments   []Comment
	Tags []Tag
}

type NewsUsecaseInterface interface {
	CreateNews(news News) (News, error)
	GetAllNews(metadata Metadata) ([]News, error)
	GetNewsByID(newsID uint) (News, error)
	UpdateNews(news News) (News, error)
	DeleteNews(newsID uint, userID uint) error
}

type NewsRepositoryInterface interface {
	CreateNews(news News) (News, error)
	GetAllNews(metadata Metadata) ([]News, error)
	GetNewsByID(newsID uint) (News, error)
	UpdateNews(news News) (News, error)
	DeleteNews(newsID uint, userID uint) error
}