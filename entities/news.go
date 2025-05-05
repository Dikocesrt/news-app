package entities

type News struct {
	ID         uint   `json:"id"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
}

type NewsUsecaseInterface interface {
	CreateNews(news News) (News, error)
	GetAllNews(metadata Metadata) ([]News, error)
	GetNewsByID(newsID uint) (News, error)
	UpdateNews(news News) (News, error)
	DeleteNews(newsID uint) error
}

type NewsRepositoryInterface interface {
	CreateNews(news News) (News, error)
	GetAllNews(metadata Metadata) ([]News, error)
	GetNewsByID(newsID uint) (News, error)
	UpdateNews(news News) (News, error)
	DeleteNews(newsID uint) error
}