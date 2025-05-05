package entities

type User struct {
	ID uint
	Password string
	Username string
	Token string
}

type UserRepositoryInterface interface {
	Register(user User) (User, error)
	Login(user User) (User, error)
}

type UserUsecaseInterface interface {
	Register(user User) (User, error)
	Login(user User) (User, error)
}