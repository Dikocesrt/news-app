package entities

type User struct {
	ID uint
	Password string
	Username string
	Token string
}

type RepositoryInterface interface {
	Register(user User) (User, error)
	Login(user User) (User, error)
}

type UsecaseInterface interface {
	Register(user User) (User, error)
	Login(user User) (User, error)
}