package entities

type Comment struct {
	ID        uint
	Name      string
	Comment   string
	News      News
}

type CommentRepositoryInterface interface {
	CreateComment(comment Comment) (Comment, error)
}

type CommentUsecaseInterface interface {
	CreateComment(comment Comment) (Comment, error)
}